import config
from jira import JIRA
import sys
import requests
import os
import textwrap
import re
import json
import six
import datetime
import pytz
import shared.bitbucket_utils

from functools import reduce
from recordclass import recordclass
from jira import JIRAError

IS_VERBOSE = False
_JIRA_CLIENT = None
PACIFIC_TIME_ZONE = pytz.timezone("America/Los_Angeles")

JIRA_ISSUE_FIELDS_REQUIRED_BY_UTIL = ["issuetype", "project", config.CUSTOM_FIELD_ID_PREVIEW_ISSUE]

PUBLIC_TICKET_CONFLUENCE_URL = "https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=43683000"
PUBLIC_TICKET_CONFLUENCE_URL_PREVIEW_ISSUE_FIELD_ANCHOR = PUBLIC_TICKET_CONFLUENCE_URL + "#RequestingapublicSDK/CLI-PreviewIssueField"
PUBLIC_TICKET_CONFLUENCE_URL_LOADING_FIELDS_FROM_PREVIEW_ISSUE_ANCHOR = PUBLIC_TICKET_CONFLUENCE_URL + "#RequestingapublicSDK/CLI-LoadingFieldsFromPreviewIssue"

# These were already in flight.
PUBLIC_DEXREQ_TICKETS_NOT_LOADED_FROM_PREVIEW = ['DEXREQ-663', 'DEXREQ-668']

# CLI sprint name
CLI_SPRINT_NAME_PREFIX = 'DEX CLI'

REGION_SUPPORT_TICKET_SUMMARY_PREFIX = 'Next set of region support in'
BUG_BASH_TICKET_SUMMARY_PREFIX = 'SDK ACTION REQUIRED'

# Release day is the day of week value, starting from Monday = 0
RELEASE_DAY = 1


def JIRA_CLIENT():
    global _JIRA_CLIENT

    if _JIRA_CLIENT:
        return _JIRA_CLIENT

    # attempt to log in using user name and password if present, if not use config.JSESSIONID
    if config.USERNAME and config.PASSWORD:
        print('Building JIRA client with username / password auth')
        _JIRA_CLIENT = JIRA(config.JIRA_OPTIONS, basic_auth=config.JIRA_BASIC_AUTH)
    elif config.JSESSIONID:
        print('Building JIRA client with cookie based auth')
        cookie_options = dict(config.JIRA_OPTIONS)
        cookie_options['cookies'] = {
            'JSESSIONID': config.JSESSIONID
        }

        _JIRA_CLIENT = JIRA(cookie_options)
    else:
        sys.exit('Could not authenticate with JIRA server. Must specify environment variables for either config.JSESSIONID or JIRA_USERNAME and JIRA_PASSWORD.')

    return _JIRA_CLIENT


def printv(s, flush=False):
    if IS_VERBOSE:
        print(s)
        if flush:
            # Flush, so we make sure the output of the issue key is already visible
            # NOTE: This is to help debug for DEX-6382
            sys.stdout.flush()


def get_jira_issue_keys(s, jira_project_key):
    if not s:
        return []

    parts = [x.strip() for x in s.split(',')]
    issue_keys = []
    for part in parts:
        issue_keys.extend(re.findall("({}-[0-9]+)".format(jira_project_key), part))

    return list(set(issue_keys))


def get_dexreq_issue_keys(s):
    return get_jira_issue_keys(s, "DEXREQ")


def add_error(errors, message, kind="ERROR"):
    printv("{}: {}".format(kind, message))
    if errors is not None:
        errors.append(message)

    return errors


def get_preview_dexreq_key(public_issue, errors=None, warnings=None):
    preview_issues_value = getattr(public_issue.fields, config.CUSTOM_FIELD_ID_PREVIEW_ISSUE).strip().upper()
    issue_keys = get_jira_issue_keys(preview_issues_value,"")
    preview_issue_keys = get_dexreq_issue_keys(preview_issues_value)
    printv("NOTE: Public ticket '{}' refers to preview ticket(s): '{}'".format(public_issue.key, ", ".join(str(k) for k in preview_issue_keys)))

    if "DEXREQ" not in preview_issues_value:
        if len(issue_keys) > 1:
            add_error(errors, "A public DEXREQ ticket must refer a preview DEXREQ ticket. {} are not DEXREQ tickets.".format(preview_issues_value))
        else:
            add_error(errors, "A public DEXREQ ticket must refer a preview DEXREQ ticket. {} is not a DEXREQ ticket.".format(preview_issues_value))

    if len(preview_issue_keys) == 0:
        add_error(errors, "A public DEXREQ ticket must refer to exactly one preview DEXREQ ticket. Add the preview ticket in the '{}' field. See {} .".format(
            config.CUSTOM_FIELD_NAME_FOR_ID[config.CUSTOM_FIELD_ID_PREVIEW_ISSUE],
            PUBLIC_TICKET_CONFLUENCE_URL_PREVIEW_ISSUE_FIELD_ANCHOR))

        return None

    if len(preview_issue_keys) > 1:
        add_error(errors, "A public DEXREQ ticket must refer to exactly one preview DEXREQ ticket. This public ticket referred to {} preview tickets in the '{}' field ('{}'). See {} .".format(
            len(preview_issue_keys), config.CUSTOM_FIELD_NAME_FOR_ID[config.CUSTOM_FIELD_ID_PREVIEW_ISSUE], ", ".join(preview_issue_keys),
            PUBLIC_TICKET_CONFLUENCE_URL_PREVIEW_ISSUE_FIELD_ANCHOR))

        return None

    return preview_issue_keys[0]


def handle_public_ticket(issue, fields=None, expand=None, errors=None, warnings=None):
    ticket_type_id = issue.fields.issuetype.id
    if ticket_type_id != config.PUBLIC_ISSUE_TYPE_ID:
        return issue

    if issue.key in PUBLIC_DEXREQ_TICKETS_NOT_LOADED_FROM_PREVIEW:
        print("NOTE: Public ticket {} is not being loaded from the linked preview ticket because it was already in flight".format(issue.key))
        return issue

    # This is a public ticket, get the preview ticket
    preview_issue_key = get_preview_dexreq_key(issue, errors=errors, warnings=warnings)
    if not preview_issue_key:
        return issue

    try:
        printv("Loading preview issue '{}' from JIRA".format(preview_issue_key))
        preview_issue = JIRA_CLIENT().issue(preview_issue_key, fields=fields, expand=expand)
        printv("Done loading preview issue '{}' from JIRA".format(preview_issue_key))
        if not preview_issue.fields.project.name == config.JIRA_PROJECT:
            add_error(errors, 'The preview ticket provided was not in the DEXREQ project: {}.'.format(preview_issue))

            return issue

        if not preview_issue.fields.issuetype.id == config.PREVIEW_ISSUE_TYPE_ID:
            add_error(errors, 'The DEXREQ ticket provided was not a preview ticket: {}.'.format(preview_issue))

            return issue

        any_public_fields_ignored = False
        # preview_issue is a DEXREQ preview ticket
        for custom_field_id in config.CUSTOM_FIELD_IDS_READ_FROM_PREVIEW_TICKET:
            public_value = None
            public_value_str = "<missing>"
            if hasattr(issue.fields, custom_field_id):
                public_value = getattr(issue.fields, custom_field_id)
                if public_value:
                    public_value_str = "'{}'".format(public_value.encode('utf8'))
                else:
                    public_value_str = "<empty>"

            preview_value = None
            preview_value_str = "<missing>"
            if hasattr(preview_issue.fields, custom_field_id):
                preview_value = getattr(preview_issue.fields, custom_field_id)
                if preview_value:
                    preview_value_str = "'{}'".format(preview_value.encode('utf8'))
                else:
                    preview_value_str = "<empty>"

            public_value_differs = (public_value != preview_value)
            public_value_missing_preview_exists = (public_value_str == "<missing>" and preview_value_str != "<missing>")

            if public_value == preview_value:
                printv("NOTE: The value {} for '{}' in public ticket was the same as in the preview ticket {}.".format(
                    public_value_str, config.CUSTOM_FIELD_NAME_FOR_ID[custom_field_id], preview_issue_key))
            else:
                if public_value_missing_preview_exists or (public_value_differs and public_value is None):
                    # Just put this in the log: it's expected that public ticket has fields missing or empty now
                    printv("NOTE: Using value {} for '{}' from preview ticket {} instead, ignoring value {} from public ticket.".format(
                        preview_value_str, config.CUSTOM_FIELD_NAME_FOR_ID[custom_field_id], preview_issue_key, public_value_str))
                elif public_value_differs:
                    add_error(warnings, "Using value {} for '{}' from preview ticket {} instead, ignoring value {} from public ticket.".format(
                        preview_value_str, config.CUSTOM_FIELD_NAME_FOR_ID[custom_field_id], preview_issue_key, public_value_str), kind="WARN")
                    any_public_fields_ignored = True

                if preview_value_str == "<missing>":
                    # make sure it's removed in the public ticket too
                    if hasattr(issue.fields, custom_field_id):
                        delattr(issue.fields, custom_field_id)
                    issue.raw['fields'].pop(custom_field_id, None)
                else:
                    setattr(issue.fields, custom_field_id, preview_value)
                    issue.raw['fields'][custom_field_id] = preview_value

        if any_public_fields_ignored:
            add_error(warnings, "For more information about the 'using value from preview ticket instead' warnings above, see {} .".format(
                PUBLIC_TICKET_CONFLUENCE_URL_LOADING_FIELDS_FROM_PREVIEW_ISSUE_ANCHOR), kind="WARN")
    except JIRAError as e:
        if e.status_code == 404:
            add_error(errors, 'The preview ticket provided could not be resolved. Please ensure {} is a valid JIRA issue.'.format(preview_issue))
        else:
            add_error(errors, 'The preview ticket provided: {} could not be resolved. Error: {}.'.format(preview_issue, str(e)))

    return issue


# Wrapper function for jira_client.issue(), which takes care of reading values from the Preview ticket, if it's a Public ticket
def get_dexreq_issue(issue_key, fields=None, expand=None, errors=None, warnings=None):
    if fields and not set(JIRA_ISSUE_FIELDS_REQUIRED_BY_UTIL).issubset(set(fields)):
        # Need issue type to determine if we need special processing for Public tickets
        fields = fields + JIRA_ISSUE_FIELDS_REQUIRED_BY_UTIL
    issue = JIRA_CLIENT().issue(issue_key, fields=', '.join(fields) if fields else None, expand=', '.join(expand) if expand else None)
    return handle_public_ticket(issue, fields, expand, errors, warnings)


def jira_search_issues(query, fields=None, expand=None):
    block_size = 50
    block_num = 1

    issues = []

    while True:
        print("Querying block {}".format(block_num))
        more_issues = JIRA_CLIENT().search_issues(query, startAt=len(issues), maxResults=block_size, fields=fields, expand=expand)
        print("Block {} returned {} issues".format(block_num, len(more_issues)))
        if more_issues:
            issues.extend(more_issues)
            if len(more_issues) != block_size:
                break
        else:
            break
        block_num += 1

    print("Returned {} issues".format(len(issues)))
    return issues


# Wrapper function for jira_client.search_issues(), which takes care of reading values from the Preview ticket, if it's a Public ticket
# errors and warnings should be dictionaries from issue key to list of strings, so you can look up errors and warnings for each issue separately
def search_dexreq_issues(query, fields=None, expand=None, errors=None, warnings=None):
    if fields and not set(JIRA_ISSUE_FIELDS_REQUIRED_BY_UTIL).issubset(set(fields)):
        # Need issue type to determine if we need special processing for Public tickets
        fields = fields + JIRA_ISSUE_FIELDS_REQUIRED_BY_UTIL

    issues = jira_search_issues(query, fields=', '.join(fields) if fields else None, expand=', '.join(expand) if expand else None)

    for issue in issues:
        if errors is not None:
            errors_list = []
        else:
            errors_list = None

        if warnings is not None:
            warnings_list = []
        else:
            warnings_list = None

        handle_public_ticket(issue, fields, expand, errors_list, warnings_list)

        if errors is not None:
            errors[issue.key] = errors_list
        if warnings is not None:
            warnings[issue.key] = warnings_list

    return issues


# updates the SDK / CLI 'Status' field of an issue to the specified value
def transition_issue_per_tool_status(jira_client, issue, status, tool_name):
    if not is_tool_jira_reportable(tool_name):
        print("Not transitioning issue {issue_key} for tool status {status}: "
              "Field for tool {tool_name} does not exist and is not reportable".format(issue_key=issue.key,
                                                                                       status=status,
                                                                                       tool_name=tool_name))
    else:
        custom_field_id_for_tool = config.CUSTOM_FIELD_ID_FOR_TOOL[tool_name]
        custom_field_name_for_tool = config.CUSTOM_FIELD_NAME_FOR_TOOL[tool_name]
        if config.IS_DRY_RUN:
            print('DRY-RUN: not transitioning issue {issue_key} to state {status_field_name}: {status}'.format(
                issue_key=issue.key,
                status_field_name=custom_field_name_for_tool,
                status=status
            ))
        else:
            print('Transitioning issue {issue_key} to state {status_field_name}: {status}'.format(
                issue_key=issue.key,
                status_field_name=custom_field_name_for_tool,
                status=status
            ))

            kwargs = {}
            kwargs[custom_field_id_for_tool] = {'value': status}
            issue.update(**kwargs)


def transition_issue_overall_status(jira_client, issue, status):
    transitions = JIRA_CLIENT().transitions(issue)
    transition_to_apply = None
    for transition in transitions:
        if transition['name'] == status:
            transition_to_apply = transition
            break

    if issue and issue.fields and issue.fields.status and issue.fields.status.name == status:
        print('Not transitioning issue {} to status "{}" because it already has that status'.format(issue.key, status))

    if transition_to_apply:
        if config.IS_DRY_RUN:
            print("DRY-RUN: not transitioning {} to {} using transition {}".format(issue.key, transition_to_apply['name'], transition_to_apply['id']))
        else:
            print("Transitioning {} to {} using transition {}".format(issue.key, transition_to_apply['name'], transition_to_apply['id']))
            JIRA_CLIENT().transition_issue(issue, transition_to_apply['id'])
    else:
        print("\tDon't know how to transition this issue to '{}'".format(status))


def transition_issue_overall_status_if_not_in_status(jira_client, issue, desired_status, blacklisted_status):
    if issue and issue.fields and issue.fields.status and issue.fields.status.name == blacklisted_status:
        print('Not transitioning issue {} to status "{}" because this transition was explicitly disallowed'.format(issue.key, desired_status))
        return

    transition_issue_overall_status(jira_client, issue, desired_status)


def update_issue_fields_helper(jira_client, issue, update_list, issue_key, preview_note=""):
    if not update_list:
        # Nothing to do
        return None

    if config.IS_DRY_RUN:
        for update in update_list:
            print('DRY-RUN: not changing issue {issue_key}{preview_note} field {field_name} (id {field_id}) to "{new_value}"'.format(
                issue_key=issue_key,
                preview_note=preview_note,
                field_name=update['field_name'],
                field_id=update['field_id'],
                new_value=update['new_value']
            ))
    else:
        fields = {}
        for update in update_list:
            print('Changing issue {issue_key}{preview_note} field {field_name} (id {field_id}) to "{new_value}"'.format(
                issue_key=issue_key,
                preview_note=preview_note,
                field_name=update['field_name'],
                field_id=update['field_id'],
                new_value=update['new_value']
            ))
            fields[update['field_id']] = update['new_value']

        issue.update(fields=fields)

    return None


# updates a fields of an issue to specified values
# update_list is a list of {field_name: "field_name", field_id: "field_id", new_value: "new_value"} triples
def update_issue_fields(jira_client, issue, update_list):
    ticket_type_id = issue.fields.issuetype.id

    preview_updates = []
    public_updates = []

    if ticket_type_id == config.PUBLIC_ISSUE_TYPE_ID:
        for update in update_list:
            if update['field_id'] in config.CUSTOM_FIELD_IDS_READ_FROM_PREVIEW_TICKET:
                # This field is actually in the preview ticket
                preview_updates.append(update)
            else:
                public_updates.append(update)

        update_list = public_updates

        if preview_updates:
            # This is a public ticket, and some fields are actually in the preview ticket
            errors = []
            preview_issue_key = get_preview_dexreq_key(issue, errors=errors)
            if preview_issue_key:
                fields = [u['field_id'] for u in preview_updates]
                preview_issue = jira_client.issue(preview_issue_key, fields=fields)
                return update_issue_fields_helper(jira_client, preview_issue, preview_updates, issue.key, preview_note=" (in preview ticket {})".format(preview_issue.key))
            else:
                return errors[0]

    return update_issue_fields_helper(jira_client, issue, update_list, issue.key)


# updates a field of an issue to the specified value
def update_issue_field(jira_client, issue, field_name, field_id, new_value):
    return update_issue_fields(jira_client, issue, [{'field_name': field_name, 'field_id': field_id, 'new_value': new_value}])


def apply_issue_transition(issue, transition):
    if config.IS_DRY_RUN:
        print('DRY-RUN: not applying transition: {transition_name} to issue {issue_key}'.format(
            issue_key=issue.key,
            transition_name=transition['name']
        ))
    else:
        print('Applying transition: {transition_name} to issue {issue_key}'.format(
            issue_key=issue.key,
            transition_name=transition['name']
        ))

        JIRA_CLIENT().transition_issue(issue, transition['id'])


def get_udx_issue_keys(udx_issue_keys):
    return get_jira_issue_keys(udx_issue_keys, "UDX")


def parse_issue_keys_from_commit_message(commit_message):
    return parse_issue_keys_from_specific_commit_message(commit_message, '[[')


# bracket_prefix allows you to narrow this down to a certain kind of message,
# e.g. "Running generation for: [["
def parse_issue_keys_from_specific_commit_message(commit_message, bracket_prefix):
    issue_keys = []
    while bracket_prefix in commit_message:
        commit_message = commit_message[commit_message.index(bracket_prefix) + len(bracket_prefix):]

        if "]]" not in commit_message:
            break

        substring = commit_message[:commit_message.index("]]")]
        parts = [part.strip() for part in substring.split(',')]
        issue_keys.extend(parts)

        commit_message = commit_message[commit_message.index("]]") + 2:]

    # filter out duplicates
    seen = set()
    seen_add = seen.add  # speed up method resolution
    return [x for x in issue_keys if not (x in seen or seen_add(x))]


def add_jira_comment(issue_key, comment, comment_type=None):
    color = config.COMMENT_TYPE_TO_COLOR.get(comment_type, None)
    if color:
        comment = '{{color:{color}}}{comment}{{color}}'.format(color=color, comment=comment)

    if config.IS_DRY_RUN:
        print("DRY-RUN: not making the following comment for {issue_key}".format(issue_key=issue_key))
        print(comment)
    else:
        print("Making the following comment for {issue_key}".format(issue_key=issue_key))
        print(comment)
        JIRA_CLIENT().add_comment(issue_key, comment)


def safe_delete_branch(repo, branch):
    for prefix in config.BRANCH_PREFIXES_SAFE_FOR_DELETION:
        if branch.startswith(prefix):
            if config.IS_DRY_RUN:
                print('DRY RUN: Not deleting branch: {}'.format(branch))
            else:
                print('Deleting branch: {}'.format(branch))
                repo.git.push('--delete', 'origin', branch)
            return

    print('Refusing to delete branch: {}'.format(branch))


def join(l):
    if len(l) < 1:
        return ""
    if len(l) == 1:
        return l[0]

    all_but_last = ', '.join(l[:-1])
    last = str(l[-1])

    return ' and '.join([all_but_last, last])


def get_last_commit_message(tool_name):
    current_branch = [branch.strip()[2:] for branch in config.REPOS_FOR_TOOL[tool_name][-1].git.branch().split('\n') if branch.startswith('* ')][0]

    last_commit_messages = {}
    for name, repo in zip(config.REPO_NAMES_FOR_TOOL[tool_name], config.REPOS_FOR_TOOL[tool_name]):
        # check out equivalent branch everywhere
        repo.git.checkout(current_branch)

        # parse DEXREQ issue out of last commit message (so we can post in the issue which build is running)
        # this build is distinct from the first build id we post in the ticket which just updates the pom.xml
        repo_last_commit_message = repo.git.log(n=1, format='%s%n%b')
        last_commit_messages[name] = repo_last_commit_message

    return last_commit_messages[config.REPO_NAMES_FOR_TOOL[tool_name][-1]]


def get_repo_id(repo_slug):
    url = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/SDK/repos/{repo}'.format(repo=repo_slug)
    r = requests.get(url, auth=config.JIRA_BASIC_AUTH)
    return r.json()['id']


def get_default_reviewers(repo_slug, source_repo_id, source_branch, target_repo_id, target_branch):
    url = 'https://bitbucket.oci.oraclecorp.com/rest/default-reviewers/1.0/projects/SDK/repos/{repo}/reviewers?sourceRepoId={source_repo_id}&targetRepoId={target_repo_id}&sourceRefId=refs/heads/{source_branch}&targetRefId=refs/heads/{target_branch}'.format(
        repo=repo_slug,
        source_repo_id=source_repo_id,
        target_repo_id=target_repo_id,
        source_branch=source_branch,
        target_branch=target_branch)
    r = requests.get(url, auth=config.JIRA_BASIC_AUTH)

    if r.status_code >= 300:
        print(r.json())

        raise ValueError("Failed to get default reviewers: {}".format(r.json()))

    return r.json()


# CLI PR is required for the following cases:
#  * Manual changes suggested in Design review ticket
#  * Test failures found during CLI Generation.
def is_cli_pr_required(issue):
    for label in config.CLI_PR_REQUIRED_LABELS:
        if label in issue.fields.labels:
            return True

    return False


def create_pull_request(repo, branch, title, description, target_repo_id, target_repo, target_branch, additional_reviewers=[]):
    repo_id = get_repo_id(repo)

    reviewers = [{"user": r} for r in additional_reviewers]
    if config.IS_DRY_RUN:
        print('DRY-RUN: not getting default reviewers, because the branch has not been pushed')
    else:
        reviewers.extend([{"user": r} for r in get_default_reviewers(repo, repo_id, branch, target_repo_id or repo_id, target_branch)])

    url = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/SDK/repos/{target_repo}/pull-requests'.format(
        target_repo=target_repo)
    headers = {'content-type': 'application/json'}

    json_data = {
        "title":title,
        "description":description,
        "state":"OPEN",
        "open":True,
        "closed":False,
        "fromRef":{
            "id":"refs/heads/{}".format(branch),
            "repository":{
                "slug":repo,
                "name":None,
                "project":{
                    "key":"SDK"
                }
            }
        },
        "toRef":{
            "id":"refs/heads/{}".format(target_branch),
            "repository":{
                "slug":target_repo,
                "name":None,
                "project":{
                    "key":"SDK"
                }
            }
        },
        "locked":False,
        "reviewers": reviewers
    }

    if config.IS_DRY_RUN:
        print('DRY-RUN: not issueing POST to {url} to with headers {headers} and body:\n{body}'.format(
            url=url,
            headers=headers,
            body=json_data
        ))

        return None
    else:
        r = requests.post(url, headers=headers, json=json_data, auth=config.JIRA_BASIC_AUTH)
        if r.status_code >= 300:
            print(r)
            print(r.json())

            raise ValueError("Failed to post pull request: {}".format(r.json()))

        return r.json()['links']['self'][0]['href']


def add_reviewer_to_pull_request(pr_id, target_repo, additional_reviewers=[]):
    reviewers = []

    url = 'https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/SDK/repos/{target_repo}/pull-requests/{pr_id}'.format(
        target_repo=target_repo, pr_id=pr_id)
    headers = {'content-type': 'application/json'}

    r = requests.get(url, headers=headers, auth=config.JIRA_BASIC_AUTH)

    if 'reviewers' in r.json():
        for reviewer in r.json()['reviewers']:
            reviewers.append(reviewer)

    for reviewer_name in additional_reviewers:
        reviewers.append({"user": {"name": reviewer_name}})

    json_data = {
        "reviewers": reviewers,
        "version": r.json()['version']
    }

    if config.IS_DRY_RUN:
        print('DRY-RUN: not issueing PUT to {url} to with headers {headers} and body:\n{body}'.format(
            url=url,
            headers=headers,
            body=json_data
        ))

        return None
    else:
        r = requests.put(url, headers=headers, json=json_data, auth=config.JIRA_BASIC_AUTH)
        if r.status_code >= 300:
            print(r)
            print(r.json())

            raise ValueError("Failed to post pull request: {}".format(r.json()))

        return r.json()['links']['self'][0]['href']


def were_steps_successful(tool_name):
    # this script runs after both generation and build have completed
    # it will run even 'If previous steps have failed' to ensure we can report failures
    generation_success = {}
    for name, success_generation_file in zip(config.REPO_NAMES_FOR_TOOL[tool_name], config.SUCCESS_GENERATION_FILES_FOR_TOOL[tool_name]):
        generation_success[name] = os.path.exists(success_generation_file)

    build_success = {}
    for name, success_build_file in zip(config.REPO_NAMES_FOR_TOOL[tool_name], config.SUCCESS_BUILD_FILES_FOR_TOOL[tool_name]):
        build_success[name] = os.path.exists(success_build_file)

    # report to JIRA tasks whether or not jobs succeeded
    # if either failed, just give link to build log and say that generation failed for some reason
    # this will be hard for external users to investigate so we want to cover easy errors earlier in the process with explicit errors:
    #   - spec artifact / group / version doesnt exist in artifactory
    #   - invalid param set
    #   - relative spec path doesn't point at a spec (yaml file)
    print("Generation successful? {}".format(generation_success))
    print("Build successful? {}".format(build_success))

    generation_pass = all(x is True for x in generation_success.values())
    build_pass = all(x is True for x in build_success.values())

    return generation_pass, build_pass


def update_file_with_feature_ids(file, feature_ids, line_regex, added_line_pattern, new_template):
    if not os.path.exists(file):
        with open(file, 'w') as f:
            f.write(new_template)

    with open(file, 'r') as f:
        content = f.read()

    content_array = content.lower().split('\n')

    need_new_line = not content.endswith('\n')
    with open(file, 'a') as f:
        if feature_ids:
            for feature_id in feature_ids:
                already_exists = False
                for line in content_array:
                    # Remove comments
                    if '#' in line:
                        line = line[line.index('#') + 1:]
                    line = line.strip()
                    if re.match(line_regex.format(feature_id.lower()), line):
                        already_exists = True
                        break
                if not already_exists:
                    if need_new_line:
                        f.write('\n')
                        need_new_line = False

                    f.write(added_line_pattern.format(feature_id.lower()))
                else:
                    print('Feature ID: {} was already contained in: {}'.format(feature_id, file))


# TODO: add comment to file with JIRA title or link to preview ticket
def update_feature_id_file(dir, feature_ids, issue_key):
    update_file_with_feature_ids(os.path.join(dir, "{}.yaml".format(issue_key)), feature_ids, r'^-\s{{1,}}{}$', '  - {}\n',
        textwrap.dedent("""\
        # Add whitelisted feature ids here. Please include a comment describing what the feature is.
        # Example:
        # whitelisted:
        #   # comment
        #   - udx-123
        #   # comment
        #   - udx-456
        whitelisted:
        """))


def update_pre_processor_file(dir, feature_ids, issue_key):
    update_file_with_feature_ids(os.path.join(dir, "{}.txt".format(issue_key)), feature_ids, '^{}$', '{}\n',
        textwrap.dedent("""\
        # This configuration file determines which conditional groups are enabled.
        # Use --group_file enabled.txt or -f enabled.txt
        #
        # The hash '#' character starts a comment
        #
        # Group names can contain the characters A-Z, a-z, 0-9 and the underscore '_' and the hyphen '-'.
        # Whitespace before or after the group name is ignored.
        #
        # GROUP1 # comment
        # GROUP2 # comment
        """))


# TODO: this is currently broken -- https://jira.oci.oraclecorp.com/browse/DEX-7328
def get_dev_status_info_for_issue(jira_client, issue):
    jira_internal_session = JIRA_CLIENT()._session

    issue_dev_status_url = config.JIRA_DEV_STATUS_REST_API_URL_FORMAT.format(issue.id)
    print("issue_dev_status_url: {}".format(issue_dev_status_url))
    return json.loads(jira_internal_session.get(issue_dev_status_url).content)


# TODO: this is currently very inefficient -- https://jira.oci.oraclecorp.com/browse/DEX-7328
def get_pull_requests_for_issue(jira_client, issue):
    created_date = getattr(issue.fields, 'created')

    prs = []
    repos = get_jira_reportable_repo_names_for_tool()
    all_repo_names = []
    for tool_name, repo_names in repos.items():
        all_repo_names.extend(repo_names)

    all_repo_names.append(config.DEXREQ_REPO_NAME)
    all_repo_names = list(set(all_repo_names))

    # The spec diff PR can't be older than the DEXREQ ticket, so only search that far
    printv("To get all spec diff PR, listing all PRs in {} newer than {}".format(all_repo_names, created_date))

    for repo in all_repo_names:
        result = shared.bitbucket_utils.get_all_pullrequest_with_string_after('SDK', repo, issue.key, created_date)
        printv("Found {} PRs for repo {}".format(len(result), repo))
        prs.extend(result)

    printv("Found {} PRs total".format(len(prs)))

    return prs


def timestamp_to_utc_string(timestamp):
    # 'dateAdded' is a timestamp in milliseconds
    # Turn it into a UTC datetime
    build_added_datetime = datetime.datetime.utcfromtimestamp(timestamp / 1000.0)
    # Then turn it into a string, like the other dates
    build_added = "{}.{:03.0f}".format(
        build_added_datetime.strftime('%Y-%m-%dT%H:%M:%S'),
        build_added_datetime.microsecond / 1000.0
    )
    return build_added


def find_dex_tools_active_sprint_id():
    boards = JIRA_CLIENT().boards()
    dex_board = next((board for board in boards if board.name.startswith('DEX Tools')), None)
    if dex_board is None:
        return None
    sprints = JIRA_CLIENT().sprints(dex_board.id)
    active_sprint = next((sprint for sprint in sprints if (sprint.state == 'ACTIVE' and sprint.name.upper().startswith(CLI_SPRINT_NAME_PREFIX.upper()))), None)

    if not active_sprint:
        # get the first future sprint!
        future_sprints = [sprint for sprint in sprints if (sprint.state == 'FUTURE' and sprint.name.upper().startswith(CLI_SPRINT_NAME_PREFIX.upper()))]
        if len(future_sprints) > 0:
            future_sprints.sort(key=lambda sprint: sprint.sequence)
            return future_sprints[0].id
        return None

    return active_sprint.id


PrStatusForTools = recordclass('PrStatusForTools', 'tools prs_per_tool all_prs_initiated all_prs_approved all_prs_merged last_pr_update last_build_update last_update')
PrsPerTool = recordclass('PrsPerTool', 'merged approved open')
PrAndUrl = recordclass('PrAndUrl', 'pr url tool_name')


# TODO: this is currently not very efficient -- https://jira.oci.oraclecorp.com/browse/DEX-7328
def get_pr_status_for_tools(jira_client, issue, tool_names, target_branch_filter=None,
        filter=None):
    pr_statuses_for_tool = PrStatusForTools(
        tools={},
        prs_per_tool={},
        all_prs_initiated=True,
        all_prs_approved=True,
        all_prs_merged=True,
        last_pr_update=None,
        last_build_update=None,
        last_update=None
    )

    for tool_name in tool_names:
        empty = PrsPerTool(
            merged=[],
            approved=[],  # Open and has at least one sign-off -- this is since service teams should sign off first.
            open=[]       # Open, but no sign-off.
        )
        pr_statuses_for_tool.prs_per_tool[tool_name] = empty

    pull_requests = get_pull_requests_for_issue(jira_client, issue)

    # if a filter is supplied, only consider PRs to that branch
    if target_branch_filter:
        pull_requests = [pr for pr in pull_requests if deep_get(pr, 'toRef.id') == 'refs/heads/{}'.format(target_branch_filter)]

    if filter:
        pull_requests = [pr for pr in pull_requests if filter(pr)]

    printv("pull_requests after filtering: {}".format(len(pull_requests)))

    if pull_requests:
        for pr in pull_requests:
            pr_url = None
            pr_and_url = None
            hrefs = deep_get(pr, 'links.self')
            if hrefs:
                pr_url = deep_get(hrefs[0], 'href')

            printv(pr_url)
            # Take the timestamps the validation builds were created into account
            # and store the latest one.
            pr_build_status = shared.bitbucket_utils.get_bitbucket_build_status_for_pr(pr)
            for build in pr_build_status['values']:
                build_added = timestamp_to_utc_string(build['dateAdded'])
                if not pr_statuses_for_tool.last_build_update or pr_statuses_for_tool.last_build_update < build_added:
                    pr_statuses_for_tool.last_build_update = build_added

            # Keep track of the last time the PR was changed (code commits, not comments)
            # and store the latest one.
            pr_last_update = timestamp_to_utc_string(pr['updatedDate'])
            if not pr_statuses_for_tool.last_pr_update or pr_statuses_for_tool.last_pr_update < pr_last_update:
                pr_statuses_for_tool.last_pr_update = pr_last_update

            pr_status = pr['state']
            target_repo_name = deep_get(pr, 'toRef.repository.name')
            corresponding_tool_name = config.REPO_NAME_TO_PRIMARY_TOOL.get(target_repo_name)
            if not corresponding_tool_name:
                print('Ignored pull request: {} for unrecognized repo'.format(target_repo_name))
                continue

            if corresponding_tool_name not in tool_names:
                continue

            if pr_url:
                pr_and_url = PrAndUrl(pr=pr, url=pr_url, tool_name=corresponding_tool_name)

            prs_for_this_tool = pr_statuses_for_tool.prs_per_tool[corresponding_tool_name]

            # if there are ANY open PRs, we override the status to 'OPEN' so we don't count it as done
            if pr_status == config.PULL_REQUEST_STATUS_OPEN:
                pr_statuses_for_tool.tools[corresponding_tool_name] = pr_status
                # Check reviewers[i].approved
                approved = False
                for reviewer in pr['reviewers']:
                    if reviewer['approved']:
                        approved = True
                        break

                if approved and pr_and_url:
                    prs_for_this_tool.approved.append(pr_and_url)
                else:
                    prs_for_this_tool.open.append(pr_and_url)
                    pr_statuses_for_tool.all_prs_approved = False

            if pr_status == config.PULL_REQUEST_STATUS_MERGED:
                prs_for_this_tool.merged.append(pr_and_url)

            # it may be that there are some PRs merged, and some PRs still open
            # we only want to return status = MERGED if ALL PRs are merged, so if status is already set to 'OPEN' leave it that way
            if not (pr_statuses_for_tool.tools.get(corresponding_tool_name) == config.PULL_REQUEST_STATUS_OPEN) and pr_status == config.PULL_REQUEST_STATUS_MERGED:
                pr_statuses_for_tool.tools[corresponding_tool_name] = pr_status

    for tool, status in six.iteritems(pr_statuses_for_tool.tools):
        if status != config.PULL_REQUEST_STATUS_MERGED:
            pr_statuses_for_tool.all_prs_merged = False

        # it's possible that some PRs are merged, and others are still open, in this case we still
        # want to report all_prs_initiated = True
        if status != config.PULL_REQUEST_STATUS_OPEN and status != config.PULL_REQUEST_STATUS_MERGED:
            pr_statuses_for_tool.all_prs_initiated = False

    # If there's a tool that we didn't find, we can't possibly have all PRs initiated, approved or merged
    for tool in tool_names:
        if tool not in pr_statuses_for_tool.tools:
            pr_statuses_for_tool.all_prs_initiated = False
            pr_statuses_for_tool.all_prs_approved = False
            pr_statuses_for_tool.all_prs_merged = False
            break

    if pr_statuses_for_tool.last_pr_update > pr_statuses_for_tool.last_build_update:
        pr_statuses_for_tool.last_update = pr_statuses_for_tool.last_pr_update
    else:
        pr_statuses_for_tool.last_update = pr_statuses_for_tool.last_build_update

    printv("pr_statuses_for_tool: {}".format(pr_statuses_for_tool))

    return pr_statuses_for_tool


def get_branch_prefix_for_spec_diff(build_type):
    prefix = config.SPEC_BRANCH_PREFIX + "-"
    if build_type == config.BUILD_TYPE_INDIVIDUAL_PREVIEW:
        return prefix + config.INDIVIDUAL_PREVIEW_BRANCH_PREFIX
    elif build_type == config.BUILD_TYPE_BULK_PENDING_MERGE_PREVIEW:
        return prefix + config.BULK_PREVIEW_BRANCH_PREFIX
    elif build_type == config.BUILD_TYPE_BULK_PENDING_MERGE_PUBLIC:
        return prefix + config.BULK_PUBLIC_BRANCH_PREFIX
    elif build_type == config.BUILD_TYPE_INDIVIDUAL_PUBLIC:
        return prefix + config.INDIVIDUAL_PUBLIC_BRANCH_PREFIX
    else:
        raise ValueError("Unknown build type: ".format(build_type))


def is_tool_jira_reportable(tool_name):
    return tool_name in config.CUSTOM_FIELD_NAME_FOR_TOOL


def get_jira_reportable_tool_names():
    tool_names = list(config.TOOL_NAMES)
    return list(filter(is_tool_jira_reportable, tool_names))


def get_jira_reportable_repo_names_for_tool():
    tool_names = get_jira_reportable_tool_names()
    filtered_repo_names = dict()
    for tool_name, repo_names in config.REPO_NAMES_FOR_TOOL.items():
        if tool_name in tool_names:
            filtered_repo_names[tool_name] = repo_names

    return filtered_repo_names


def get_jira_custom_field_ids_for_tool():
    tool_names = get_jira_reportable_tool_names()
    filtered_field_ids = dict()
    for tool_name, repo_names in config.CUSTOM_FIELD_ID_FOR_TOOL.items():
        if tool_name in tool_names:
            filtered_field_ids[tool_name] = repo_names

    return filtered_field_ids


def deep_get(dictionary, keys, default=None):
    return reduce(lambda d, key: d.get(key, default) if isinstance(d, dict) else default, keys.split("."), dictionary)


def get_all_pr_urls_from_comment(issue, pr_comment_prefix=None):
    pr_urls = []
    comments = JIRA_CLIENT().comments(issue)
    for comment in reversed(comments):
        if pr_comment_prefix:
            comment_body = comment.body
            if comment_body.startswith(pr_comment_prefix):
                pr_url = comment_body.split()[-1].strip('[]')
                print(pr_url)
                if pr_url.startswith(config.BITBUCKET_PR_URL_PREFIX):
                    pr_urls.append(pr_url)
        else:
            urls = re.findall(r'({}[^\s\]]+)'.format(config.BITBUCKET_PR_URL_PREFIX), comment.body)
            pr_urls.extend(urls)
    print('Found {} PRs from comments.'.format(len(pr_urls)))
    return pr_urls


def are_all_prs_merged(pr_urls):
    for pr_url in pr_urls:
        pr = shared.bitbucket_utils.get_pullrequest_from_url(pr_url)
        if pr.json() and 'state' in pr.json() and pr.json()['state'].lower() == 'merged':
            print('PR merged: {}'.format(pr_url))
        else:
            print('PR not merged: {}'.format(pr_url))
            return False
    return True


def get_next_release_day(release_day):
    d = datetime.datetime.now()
    while d.weekday() != 1:
        d += datetime.timedelta(days=release_day)
    d = d.replace(hour=0, minute=0, second=0, microsecond=0)
    print('Next Tuesday is {}'.format(d))
    return d


def get_in_progress_region_support_tickets():
    query = 'project = "DEX" AND summary ~ "{}" AND status in ("{}")'.format(REGION_SUPPORT_TICKET_SUMMARY_PREFIX, config.STATUS_IN_PROGRESS)
    return jira_search_issues(query)


# Assumption: Bug Bash ticket is in the format of
# "[SDK] ACTION REQUIRED: DAL_TEST Customer Validation BugBash us-gov-fortworth-3"
# Ticket status should not be Done, Closed, or In Progress.
def get_unprocessed_bug_bash_tickets():
    query = ('project = "Bug Bash" AND summary ~ "{}" AND status not in ("{}", "{}", "{}", "{}")'
             .format(BUG_BASH_TICKET_SUMMARY_PREFIX, config.STATUS_CLOSED, config.STATUS_DONE, config.STATUS_IN_PROGRESS, config.STATUS_IN_REVIEW))
    return jira_search_issues(query)


# Assumption: DEX ticket is in the format of 
# "Next set of region support in Public SDKs - mx-queretaro-1" for master branch or 
# "Next set of region support in preview SDKs - mx-queretaro-1" for preview branch
# Ticket status should not be Done, Closed, or In Progress. 
def get_unprocessed_region_suppport_tickets(branch):
    if branch.lower() == 'preview':
        branch_in_query = 'preview'
    elif branch.lower() == 'master':
        branch_in_query = 'Public'
    else:
        raise ValueError('Branch must be either preview or master')
    query = 'project = "DEX" AND summary ~ "{} {} SDKs" AND status not in ("{}", "{}", "{}", "{}")'.format(
        REGION_SUPPORT_TICKET_SUMMARY_PREFIX, branch_in_query, config.STATUS_CLOSED, config.STATUS_DONE, config.STATUS_IN_PROGRESS, config.STATUS_MORE_INFORMATION_NEEDED)
    return jira_search_issues(query)


# Assumption: DEX ticket uses "due date" field to specify the expected release date of the new region
# If the due date is before the coming Tuesday, it is ready be processed.
def get_region_support_tickets_to_process(branch):
    issues = get_unprocessed_region_suppport_tickets(branch)
    # For public SDK tickets, check due date and only return tickets that are due less than a week before the next SDK release.
    # Otherwise, for preview SDK tickets, return all unprocessed ones.
    if branch == 'master':
        ready_issues = []
        for issue in issues:  
            for key in issue.raw['fields'].keys():
                if str(key).startswith('duedate'):
                    date_str = issue.raw['fields'][key]
                    date = datetime.datetime.strptime(date_str, '%Y-%m-%d')
                    today = datetime.datetime.now()
                    if date >= today:
                        next_tuesday = get_next_release_day(RELEASE_DAY)
                        if date <= next_tuesday:
                            print('Ticket {} with due date {} should be included in the next release.'.format(issue, date_str))
                            ready_issues.append(issue)
        print('Found {} issues ready to be added.'.format(len(ready_issues)))
        return ready_issues
    else:
        return issues


# Assumption: DEX ticket is in the format of "Next set of region support in Public SDKs - mx-queretaro-1"
def get_region_id_from_dex_tickets(issues):
    region_ids = []
    for issue in issues:
        region_id = issue.raw['fields']['summary'].split()[-1]
        region_ids.append(region_id)
    return region_ids
