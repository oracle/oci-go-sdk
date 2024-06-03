import getpass
import json
import requests
import os
import re
import tempfile
import datetime
import time
import sys
import config

from git import Repo, GitCommandError
from six import text_type


verbose = True
dry_run = False


#
# Bitbucket

BITBUCKET_PR_QUERY_URL_WITHOUT_START_ADDITIONAL_PARAMS = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/{project}/repos/{repo}/pull-requests?state={state}&direction={direction}{at_str}'
BITBUCKET_PR_QUERY_URL_WITH_START_ADDITIONAL_PARAMS = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/{project}/repos/{repo}/pull-requests?state={state}&direction={direction}&start={start}{at_str}'
BITBUCKET_PR_QUERY_URL_WITHOUT_START = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/{}/repos/{}/pull-requests?state={}'
BITBUCKET_PR_QUERY_URL_WITH_START = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/{}/repos/{}/pull-requests?state={}&start={}'


def printv(s):
    global verbose
    if verbose:
        if s:
            if not isinstance(s, str) and not isinstance(s, text_type):
                # we only need to turn it into a string if it's not a string already
                s = str(s)
            if sys.version.startswith("2"):
                s = s.encode('utf-8')
        print(s)


def setup_bitbucket(args):
    global bitbucket_auth
    username = os.environ.get('BITBUCKET_USERNAME')
    if not username:
        print("BITBUCKET_USERNAME not set, falling back to using JIRA_USERNAME.")
        username = os.environ.get('JIRA_USERNAME')
    if not username:
        print("JIRA_USERNAME not set either.")
        username = getpass.getpass("Bitbucket API USERNAME:")

    password = os.environ.get('BITBUCKET_PASSWORD')
    if not password:
        print("BITBUCKET_PASSWORD not set, falling back to using JIRA_PASSWORD.")
        password = os.environ.get('JIRA_PASSWORD')
    if not password:
        print("JIRA_PASSWORD not set either.")
        password = getpass.getpass("Bitbucket API Password:")

    bitbucket_auth = (username, password)


def filter_for_pr_matching_branch_suffix(pr, branch_suffix):
    if 'fromRef' not in pr or 'id' not in pr['fromRef']:
        return False

    source_branch = pr['fromRef']['id']
    return branch_suffix in source_branch


def filter_for_pr_matching_substring(pr, substring):
    substring = substring.lower()
    result = ('title' in pr and substring in pr['title'].lower()) or ('description' in pr and substring in pr['description'].lower()) or ('fromRef' in pr and 'id' in pr['fromRef'] and substring in pr['fromRef']['id'].lower())
    # printv("#### filter_for_pr_matching_substring: substring '{}' in {} pr {}: result {}, title '{}', description '{}', fromRef '{}'".format(substring, pr['toRef']['repository']['name'],
    #     pr['id'], result, pr['title'].encode('utf8') if 'title' in pr else "<none>", pr['description'].encode('utf8') if 'description' in pr else "<none>", pr['fromRef']['id'].encode('utf8')))
    return result


def stop_listing_if_pr_older_than_date(pr, dt):
    # need to convert date in "2019-11-13T03:06:09.000+0000" (JIRA) form to
    # "1573763297726" (Bitbucket) milliseconds form
    utc_time = datetime.datetime.strptime(dt, "%Y-%m-%dT%H:%M:%S.%f+0000")
    epoch = datetime.datetime.utcfromtimestamp(0)
    milliseconds = int((utc_time - epoch).total_seconds() * 1000.0)
    # If the PR was last updated before the DEXREQ ticket was opened, it can't belong to the DEXREQ ticket
    if 'updatedDate' in pr and int(pr['updatedDate']) < milliseconds:
        # printv("#### stop_listing_if_pr_older_than_date, dt {} ({}), pr updatedDate {}".format(milliseconds, dt, pr['updatedDate']))
        return True
    else:
        return False


def get_newest_pullrequest_with_string_after(project, repo, title_substring, dt, use_cache=True):
    return get_newest_pullrequest_matching(project, repo, lambda pr: filter_for_pr_matching_substring(pr, title_substring),
        stop_listing=lambda pr: stop_listing_if_pr_older_than_date(pr, dt), use_cache=use_cache)


def get_spec_pr_branch_reference(pipeline, branch_suffix):
    return "refs/heads/spec-{debug_dexreq_branch_prefix}auto-v2-{pipeline}-{suffix}-diff".format(
        debug_dexreq_branch_prefix=config.get_debug_dexreq_branch_prefix(),
        pipeline=pipeline,
        suffix=branch_suffix)


def get_newest_pullrequest_matching_branch_suffix(project, repo, branch_suffix, pipeline="preview", use_cache=True):
    return get_newest_pullrequest_matching(project, repo, lambda pr: filter_for_pr_matching_branch_suffix(pr, branch_suffix),
        direction="outgoing", at=get_spec_pr_branch_reference(pipeline, branch_suffix), use_cache=use_cache)


def get_newest_pullrequest_matching(project, repo, filter, stop_listing=lambda pr: False, direction="incoming", state="ALL", at=None, use_cache=True):
    prs = get_all_pullrequest_matching(project, repo, filter, stop_listing, direction, state, at, use_cache=use_cache)
    candidate = None

    # We want the newest as in "most recently created"
    # But Bitbucket returns in order of "most recently updated"
    # So we keep track of the highest matching id
    for pr in prs:
        if filter(pr):
            if not candidate or candidate['id'] < pr['id']:
                candidate = pr

    return candidate


def get_all_pullrequest_with_string_after(project, repo, title_substring, dt, use_cache=True):
    return get_all_pullrequest_matching(project, repo, lambda pr: filter_for_pr_matching_substring(pr, title_substring),
        stop_listing=lambda pr: stop_listing_if_pr_older_than_date(pr, dt), use_cache=use_cache)


def get_all_pullrequest_matching(project, repo, filter, stop_listing=lambda pr: False, direction="incoming", state="ALL", at=None, use_cache=True):
    prs = get_all_pullrequest_unfiltered(project, repo, stop_listing, direction, state, at, use_cache=use_cache)

    # Search the PRs to see if there are any that matches the filter
    return [pr for pr in prs if filter(pr)]


pullrequest_cache = {}


# Note: If stop_listing is being used to stop listing PRs older than a particular DEXREQ ticket, make sure to prime the cache.
#       Request once with the oldest DEXREQ ticket.
def get_all_pullrequest_unfiltered(project, repo, stop_listing=lambda pr: False, direction="incoming", state="ALL", at=None, use_cache=True):
    global bitbucket_auth
    url = BITBUCKET_PR_QUERY_URL_WITHOUT_START_ADDITIONAL_PARAMS.format(
        project=project,
        repo=repo,
        state=state,
        direction=direction,
        at_str="&at={}".format(at) if at else "")

    initial_url = url

    if use_cache and initial_url in pullrequest_cache:
        printv("Returning cached PRs for {}".format(initial_url))
        return pullrequest_cache[initial_url]

    headers = {
        "X-Atlassian-Token": "no-check",
        "Content-Type": "application/json"
    }

    prs = []

    is_last_page = False
    while not is_last_page:
        printv("Querying PRs from {}".format(url))
        response = requests.get(url, verify=False, auth=bitbucket_auth, headers=headers)

        print("***************************************************")
        print(response)
        print(response.status_code)
        if len(response.content) == 0:
            print("no response content returned")
        else:
            requested_truncated_length = 250
            truncated_response_content_length = requested_truncated_length if requested_truncated_length < len(response.content) else len(response.content)
            print(response.content[0:truncated_response_content_length - 1])
        print("***************************************************")
        
        count_of_429s = 0
        max_retries = 10
        if response.status_code == 429:
            print("Received 429, waiting before resuming fetching pull requests")
            time.sleep(30)
            count_of_429s = count_of_429s + 1

            if count_of_429s < max_retries:
                continue

        is_last_page = 'isLastPage' in response.json() and response.json()['isLastPage']
        if not is_last_page:
            url = BITBUCKET_PR_QUERY_URL_WITH_START_ADDITIONAL_PARAMS.format(
                project=project,
                repo=repo,
                state=state,
                direction=direction,
                at_str="&at={}".format(at) if at else "",
                start=response.json()['nextPageStart'])

        if 'values' not in response.json():
            continue

        for pr in response.json()['values']:
            printv("Query returned: {}/{} PR {}".format(project, repo, pr['id']))
            prs.append(pr)

            if stop_listing(pr):
                printv("Stop listing")
                is_last_page = True
                break

    pullrequest_cache[initial_url] = prs

    return prs


def get_pullrequest(project, repo, pullrequest_id):
    global bitbucket_auth
    url = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/{}/repos/{}/pull-requests/{}'.format(
        project, repo, pullrequest_id)

    headers = {
        "X-Atlassian-Token": "no-check",
        "Content-Type": "application/json"
    }

    response = requests.get(url, verify=False, auth=bitbucket_auth, headers=headers)
    return response


def get_pullrequest_from_url(url):
    m = re.search("^.*bitbucket.*/projects/([^/]*)/repos/([^/]*)/pull-requests/([0-9]*).*$", url)
    if m:
        return get_pullrequest(m.group(1), m.group(2), m.group(3))
    else:
        return None


def get_pullrequest_commits(project, repo, pullrequest_id):
    global bitbucket_auth
    url = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/{}/repos/{}/pull-requests/{}/commits'.format(
        project, repo, pullrequest_id)

    headers = {
        "X-Atlassian-Token": "no-check",
        "Content-Type": "application/json"
    }

    response = requests.get(url, verify=False, auth=bitbucket_auth, headers=headers)
    return response


def get_pullrequest_merge_status(project, repo, pullrequest_id):
    global bitbucket_auth
    url = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/{}/repos/{}/pull-requests/{}/merge'.format(
        project, repo, pullrequest_id)

    headers = {
        "X-Atlassian-Token": "no-check",
        "Content-Type": "application/json"
    }

    response = requests.get(url, verify=False, auth=bitbucket_auth, headers=headers)
    return response


def get_file_content_from_commit_id_and_path(project, repo, file_path, commit_id):
    global bitbucket_auth
    url = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/{}/repos/{}/raw/{}?at={}'.format(
        project, repo, file_path, commit_id)

    headers = {
        "X-Atlassian-Token": "no-check",
        "Content-Type": "application/json"
    }

    response = requests.get(url, verify=False, auth=bitbucket_auth, headers=headers)
    return response


def get_pullrequest_changes(project, repo, pullrequest_id):
    global bitbucket_auth
    url = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/{}/repos/{}/pull-requests/{}/changes'.format(
        project, repo, pullrequest_id)

    headers = {
        "X-Atlassian-Token": "no-check",
        "Content-Type": "application/json"
    }
    params = {"start": "0", "limit": "1000"}
    response = requests.get(url, verify=False, auth=bitbucket_auth, headers=headers, params=params)
    return response


def get_pullrequest_diff(project, repo, pullrequest_id):
    global bitbucket_auth
    url = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/{}/repos/{}/pull-requests/{}/diff'.format(
        project, repo, pullrequest_id)

    headers = {
        "X-Atlassian-Token": "no-check",
        "Content-Type": "application/json"
    }

    response = requests.get(url, verify=False, auth=bitbucket_auth, headers=headers)
    return response


def get_commit_diff(project, repo, commit_hash, commit_path):
    global bitbucket_auth
    url = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/{}/repos/{}/commits/{}/diff'.format(
        project, repo, commit_hash)
    if commit_path:
        url = url + "/{}".format(commit_path)

    headers = {
        "X-Atlassian-Token": "no-check",
        "Content-Type": "application/json"
    }

    response = requests.get(url, verify=False, auth=bitbucket_auth, headers=headers)
    return response


def get_repo(project, repo):
    global bitbucket_auth
    url = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/{}/repos/{}'.format(
        project, repo)

    printv("get_repo URL: {}".format(url))

    headers = {
        "X-Atlassian-Token": "no-check",
        "Content-Type": "application/json"
    }

    response = requests.get(url, verify=False, auth=bitbucket_auth, headers=headers)
    return response


MAX_COMMENT_LENGTH = 32000


DEFAULT_COMMENT_HEADER = "(continued from comment above)\n\n"
DEFAULT_COMMENT_FOOTER = "\n\n(continued in comment below)"


def make_general_comment(project, repo, pullrequest_id, text, parent_id=None, comment_header=DEFAULT_COMMENT_HEADER, comment_footer=DEFAULT_COMMENT_FOOTER):
    # $ curl -k -u mricken https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/SDK/repos/bmc-sdk-swagger/pull-requests/257/comments -X POST -H "X-Atlassian-Token: no-check" -H "Content-Type: application/json" -d '{"text":"A test comment."}'
    # Enter host password for user 'mricken':
    # {"properties":{"repositoryId":3103},"id":122246,"version":0,"text":"A test comment.","author":{"name":"mricken","emailAddress":"mathias.ricken@oracle.com","id":1776,"displayName":"Mathias Ricken","active":true,"slug":"mricken","type":"NORMAL","links":{"self":[{"href":"https://bitbucket.oci.oraclecorp.com/users/mricken"}]}},"createdDate":1524504470238,"updatedDate":1524504470238,"comments":[],"tasks":[],"permittedOperations":{"editable":true,"deletable":true}}

    url = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/{}/repos/{}/pull-requests/{}/comments'.format(
        project, repo, pullrequest_id)

    continued = False
    first_response = None
    in_code_block = False
    while text:
        if len(text) <= MAX_COMMENT_LENGTH:
            first_text = "{}{}{}".format(
                comment_header if continued else "",
                "```\n" if in_code_block else "",
                text)
            text = ""
        else:
            break_pos = text[:MAX_COMMENT_LENGTH].rfind("-----\n")
            if break_pos == -1:
                break_pos = text[:MAX_COMMENT_LENGTH].rfind("\n")
            if break_pos == -1:
                break_pos = text[:MAX_COMMENT_LENGTH].rfind(" ")
            if break_pos == -1:
                break_pos = MAX_COMMENT_LENGTH

            partial_text = text[0:break_pos]
            ends_in_code_block = in_code_block
            for line in partial_text.split("\n"):
                if "```" in line:
                    ends_in_code_block = not ends_in_code_block

            first_text = "{}{}{}{}{}".format(
                comment_header if continued else "",
                "```\n" if in_code_block else "",
                partial_text,
                "\n```\n" if ends_in_code_block else "",
                comment_footer)
            text = text[break_pos + 1:]
            continued = True
            in_code_block = ends_in_code_block

        if parent_id:
            data = '''{{
                "text": "{}",
                "parent": {{
                    "id": {}
                }}
            }}'''.format(json.dumps(first_text)[1:-1], parent_id)  # JSON escape, but this encloses in "" -- get rid of first and last character
        else:
            data = '''{{
                "text": "{}"
            }}'''.format(json.dumps(first_text)[1:-1])  # JSON escape, but this encloses in "" -- get rid of first and last character

        headers = {
            "X-Atlassian-Token": "no-check",
            "Content-Type": "application/json"
        }

        if dry_run:
            print("DRY-RUN: {}".format(first_text))
        else:
            response = requests.post(url, data=data, verify=False, auth=bitbucket_auth, headers=headers)
        
            printv(response.text)
            if response.status_code >= 400:
                printv(response.text)

            if not first_response:
                first_response = response

            printv(first_response.text)

            if not parent_id:
                parent_id = first_response.json()['id']

    return first_response


def get_pr_target_branch(pr):
    base_ref = pr.json()['toRef']
    base_branch = base_ref['displayId']

    return base_branch


def get_pr_source_branch(pr):
    change_ref = pr.json()['fromRef']
    change_branch = change_ref['displayId']

    return change_branch


def get_pr_source_project(pr):
    return pr.json()['fromRef']['repository']['project']['key']


def get_pr_source_repo(pr):
    return pr.json()['fromRef']['repository']['name']


def get_pr_source_self_url(pr):
    links = pr.json()['fromRef']['repository']['links']

    for self_link in links['self']:
        return self_link['href']
    raise ValueError('No self URL found')


def get_repo_permissions_url(pr):
    self_url = get_pr_source_self_url(pr)
    return self_url.replace('/browse', '/permissions')


def get_pr_source_clone_ssh_url(pr):
    links = pr.json()['fromRef']['repository']['links']

    if 'clone' not in links:
        return None

    for clone_link in links['clone']:
        if clone_link['name'] == 'ssh':
            return clone_link['href']
    raise ValueError('No clone URL found')


def get_bitbucket_build_status_for_commit(commit):
    # https://bitbucket.oci.oraclecorp.com/rest/build-status/1.0/commits/97a464dfa9933e6842fa9eb3e954d603df502e78
    url = 'https://bitbucket.oci.oraclecorp.com/rest/build-status/1.0/commits/{commit}'.format(commit=commit)
    r = requests.get(url, auth=bitbucket_auth)

    if r.status_code >= 300:
        printv(r.json())

        raise ValueError("Failed to get Bitbucket build status: {}".format(r.json()))

    return r.json()


def get_bitbucket_build_status_for_pr(pr):
    change_ref = pr['fromRef']
    latest_commit = change_ref['latestCommit']

    return get_bitbucket_build_status_for_commit(latest_commit)


# This is currently being used in PythonCLI PR builder.
# Ref: https://bitbucket.oci.oraclecorp.com/projects/SDK/repos/python-cli/browse/scripts/comment_on_pr.py
def get_pr_from_branch(project, repo, branch, state="OPEN"):
    global bitbucket_auth
    url = BITBUCKET_PR_QUERY_URL_WITHOUT_START.format(project, repo, state)
    headers = {
        "X-Atlassian-Token": "no-check",
        "Content-Type": "application/json"
    }

    # Search the PRs to see if there are any that matches the filter
    is_last_page = False
    while not is_last_page:
        response = requests.get(url, verify=False, auth=bitbucket_auth, headers=headers)
        printv(response.json())

        for pr in response.json()['values']:
            if pr['fromRef']['displayId'] == branch:
                return pr['id']

        is_last_page = response.json()['isLastPage']

        if not is_last_page:
            url = BITBUCKET_PR_QUERY_URL_WITH_START.format(project, repo, state, response.json()['nextPageStart'])

    return None

    
def get_git_clone_url_from_ref(ref):
    for clone_link in ref['repository']['links']['clone']:
        if clone_link['name'] == 'ssh':
            return clone_link['href']
    raise ValueError('No clone URL found')


def git_clone(clone_url, branch, to_dir):
    printv("Cloning {}".format(clone_url))
    repo = Repo.clone_from(clone_url, to_dir)
    printv("Checking out {}".format(branch))
    repo.git.checkout(branch)
    return repo


def clone_target_branch(pr_id, repo):
    pr = get_pullrequest("SDK", repo, pr_id)
    printv(pr.text)
    base_ref = pr.json()['toRef']
    base_clone_url = get_git_clone_url_from_ref(base_ref)
    base_branch = base_ref['displayId']

    output_dir = tempfile.mkdtemp(prefix='warn_about_backward_incompatible_changes')
    base_dir = '{}/base'.format(output_dir)

    printv("Checking out base in {}".format(base_dir))

    try:
        git_clone(base_clone_url, base_branch, base_dir)
    except GitCommandError as e:
        print(e)
        print("ERROR, could not check out commit or branch '{}' from {}".format(base_branch, base_clone_url))
        return None
    return base_dir


def decline_pr(project, repo, pullrequest_id, version):
    url = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/{}/repos/{}/pull-requests/{}/decline?version={}'.format(
        project, repo, pullrequest_id, version)

    headers = {
        "X-Atlassian-Token": "no-check",
        "Content-Type": "application/json"
    }

    response = None
    if dry_run:
        print("DRY-RUN: not declining {}/{} PR {}".format(project, repo, pullrequest_id))
    else:
        response = requests.post(url, verify=False, auth=bitbucket_auth, headers=headers)
        
        printv(response)
        if response.status_code >= 400:
            printv(response.json())

        print("Declined {}/{} PR {}".format(project, repo, pullrequest_id))

    return response


def merge_pr(project, repo, pullrequest_id, version):
    url = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/{}/repos/{}/pull-requests/{}/merge?version={}'.format(
        project, repo, pullrequest_id, version)

    headers = {
        "X-Atlassian-Token": "no-check",
        "Content-Type": "application/json"
    }

    response = None
    if dry_run:
        print("DRY-RUN: not merging {}/{} PR {}".format(project, repo, pullrequest_id))
    else:
        response = requests.post(url, verify=False, auth=bitbucket_auth, headers=headers)
        
        printv(response)
        if response.status_code >= 200 and response.status_code < 300:
            print("Merged {}/{} PR {}".format(project, repo, pullrequest_id))
        else:
            print(response.json())

    return response
