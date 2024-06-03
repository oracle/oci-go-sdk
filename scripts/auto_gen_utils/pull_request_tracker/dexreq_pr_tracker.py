import six
import sys
import getpass

import stashy
import os
import oci
import argparse
import vcr
import requests

from datetime import datetime
from datetime import timedelta
from datetime import date
from dateutil import tz
import pytz
import re
import json
import time

import jinja2
from jira import JIRA

# Add the root of the package, one directory up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '..'))

import util  # noqa: ignore=F402

BITBUCKET_SERVER_URL = "https://bitbucket.oci.oraclecorp.com"
PROJECT_NAME = "SDK"

BITBUCKET_BUILD_STATUS_API_URL_PATH_FORMAT = "/rest/build-status/1.0/commits/{}"
BITBUCKET_PULL_REQUEST_CHANGES_API_URL_PATH_FORMAT = "/rest/api/1.0/projects/SDK/repos/{}/pull-requests/{}/changes"

TICKET_TYPE_PUBLIC = "PUBLIC"
TICKET_TYPE_PREVIEW = "PREVIEW"
TICKET_TYPE_ALL = "ALL"

MASTER_BRANCH_REF = 'refs/heads/master'
PREVIEW_BRANCH_REF = 'refs/heads/preview'

STATE_OPEN = "OPEN"
STATE_MERGED = "MERGED"

# how far back to look in PRs on bitbucket
PULL_REQUEST_PAGINATION_LIMIT = 45

REPO_SLUG_JAVA_SDK = 'java-sdk'
REPO_SLUG_CLI = 'python-cli'
REPO_SLUG_TESTING_SERVICE = 'oci-testing-service'
REPO_SLUG_RUBY_SDK = 'ruby-sdk'
REPO_SLUG_PYTHON_SDK = 'python-sdk'
REPO_SLUG_GO_SDK = 'oci-go-sdk'

REPOS_TO_REPORT = [
    REPO_SLUG_JAVA_SDK,
    REPO_SLUG_CLI,
    REPO_SLUG_TESTING_SERVICE,
    REPO_SLUG_RUBY_SDK,
    REPO_SLUG_PYTHON_SDK,
    REPO_SLUG_GO_SDK
]

OUTPUT_FORMAT_TEXT = "text"
OUTPUT_FORMAT_HTML = "html"
OUTPUT_FORMAT_JSON = "json"

JIRA_SERVER = 'https://jira.oci.oraclecorp.com'
JIRA_SERVER_REST_API_VERSION = 2

JIRA_OPTIONS = {
    'server': JIRA_SERVER,
    'rest_api_version': JIRA_SERVER_REST_API_VERSION,
    'verify': True
}

DEFAULT_JIRA_ISSUE_FIELDS = ['summary', 'description']

PUBLIC_IN_PROGRESS_QUERY = 'project = DEXREQ and issuetype = "Public" and status = "In Progress"'
PUBLIC_READY_FOR_RELEASE_QUERY = 'project = DEXREQ and issuetype = "Public" and status = "Ready for Release"'

PUBLIC_DEXREQ_ISSUE_FROM_UDX_TICKET_QUERY_FORMAT = 'project = DEXREQ and (summary ~ "{udx_ticket}" or "UDX Ticket" ~ "{udx_ticket}") and issuetype = "Public"'
PREVIEW_DEXREQ_ISSUE_FROM_UDX_TICKET_QUERY_FORMAT = 'project = DEXREQ and (summary ~ "{udx_ticket}" or "UDX Ticket" ~ "{udx_ticket}") and issuetype = "Preview"'

PUBLIC_WITH_GA_DATE_QUERY_FORMAT = 'project = DEXREQ and issuetype = "Public" and "SDK/CLI GA Date" = "{}"'

CUSTOM_FIELD_UDX_CLI_GA_DATE = "customfield_11120"
CUSTOM_FIELD_UDX_SDK_GA_DATE = "customfield_11126"
CUSTOM_FIELD_UDX_CONSOLE_GA_DATE = "customfield_12197"
CUSTOM_FIELD_ID_UDX_TICKET = 'customfield_13596'

CUSTOM_FIELD_FEATURE_AVAILABLE_IN_PROD = 'customfield_13780'

OUTPUT_DIRECTORY = 'reports'
UPLOAD_TO_OBJECT_STORAGE = False
REPORT_NAME = 'report'

# PAR with write access on bucket to add reports
OBJECT_STORAGE_DEXREQ_REPORTS_PAR = "https://objectstorage.us-phoenix-1.oraclecloud.com/p/ZZoZkNKUfyEyPQ_VVF3eA8HF28YobPfGBAuVSeJA2C8/n/dex-us-phx-cli-1/b/dexreq_reports_collection/o/"
OBJECT_STORAGE_DEXREQ_REPORTS_INDEX_RW_PAR = "https://objectstorage.us-phoenix-1.oraclecloud.com/p/L8mVstLyLPSGmVO5-vFP5VKukTQGhj1W5d0y1NCVbE4/n/dex-us-phx-cli-1/b/dexreq_reports/o/reports_index.json"

TIME_ZONE_FOR_DATE_FORMATS = 'America/Los_Angeles'
DATE_FORMAT = '%h %d %I:%M %p %Z'

VCR_INSTANCE = None

SLEEP_BETWEEN_REQUESTS = 0.2

PYTHON_CLI_REPO_SLUG = 'python-cli'
JAVA_SDK_REPO_SLUG = 'java-sdk'
RUBY_SDK_REPO_SLUG = 'ruby-sdk'
GO_SDK_REPO_SLUG = 'oci-go-sdk'
PYTHON_SDK_REPO_SLUG = 'python-sdk'
OCI_TESTING_SERVICE_REPO_SLUG = 'oci-testing-service'

# this was for a hack day so I'm allowed to do ugly things like this :)
DEX_TEAM_APPROVER_EMAILS = [
    # sdk team
    'joe.levy@oracle.com',
    'jyoti.s.saini@oracle.com',
    'mathias.ricken@oracle.com',
    'mingchi.stephen.mak@oracle.com',
    'omkar.p.patil@oracle.com',
    'peng.p.liu@oracle.com',
    'sun.yan@oracle.com',
    'vyas.bhagwat@oracle.com',
    'walt.tran@oracle.com',
    'ziyao.qiao@oracle.com',
    'yash.chandra@oracle.com',
    'vidhi.bhansali@oracle.com',
    'anurag.g.garg@oracle.com',
    'josh.hunter@oracle.com',
    'swetha.krishnan@oracle.com',
    'rakesh.kumar.parappa@oracle.com',
    'swarnava.s.sarkar@oracle.com',
    'ram.kishan.v.vooka@oracle.com',
    'nivedita.parihar@oracle.com',
    'kalpana.ramasamy@oracle.com',
    'joshua.r.ramirez@oracle.com',
    'eric.pendergrass@oracle.com',
    
    # CLI team
    'hamada.ibrahim@oracle.com',
    'kern.lee@oracle.com',
    'manoj.meda@oracle.com',
    'mike.c.ross@oracle.com',
    'srikanth.reddy.kumbham@oracle.com',
    'viral.modi@oracle.com',
    'vishwas.bhat@oracle.com',
    'h.harsh.kumar@oracle.com',
    'arun.swarnam@oracle.com',
    'varun.mankal@oracle.com',
    'zhongwan.wang@oracle.com',
    'alex.t.le@oracle.com',
    'mandy.tsai@oracle.com'
]

REPORTS_NAMESPACE = 'dex-us-phx-cli-1'
REPORTS_BUCKET_NAME = 'dexreq_reports_collection'

CONFIG_FILE_LOCATION = os.path.join('resources', 'config')

AUTO_GEN_API_KEY_PASS_PHRASE_ENV_VAR = 'AUTO_GEN_PASS_PHRASE'


def unix_date_format(timestamp):
    utc_datetime = datetime.utcfromtimestamp(timestamp / 1000)
    return utc_datetime.replace(tzinfo=tz.tzutc()).astimezone(pytz.timezone(TIME_ZONE_FOR_DATE_FORMATS)).strftime(DATE_FORMAT)


def render_html(data):
    environment = jinja2.Environment(
        loader=jinja2.FileSystemLoader('./templates')
    )

    environment.filters['unix_date_format'] = unix_date_format

    template = environment.get_template('report_table.html')
    return template.render(data)


def get_all_pull_requests(bb_client, repo_slug, at, state):
    prs = []
    count = 0

    for pull_request in bb_client.projects[PROJECT_NAME].repos[repo_slug].pull_requests.all(at=at, state=state):
        # LOG.debug('Checking PR: ' + pull_request.get('title'))
        if (pull_request.get('state', '') != 'DECLINED'):
            prs.append(pull_request)

        count = count + 1
        if PULL_REQUEST_PAGINATION_LIMIT is not None and count > PULL_REQUEST_PAGINATION_LIMIT:
            # print('Stopping paginating early for PRs, slug: {}, target={}, issue={}'.format(repo_slug, at, dexreq_issue_key))
            break

    _throttle()

    return prs


def build_clients():
    username = os.environ.get('JIRA_USERNAME')
    if not username:
        sys.stderr.write('Bitbucket / JIRA username: ')
        username = six.moves.input()

    password = os.environ.get('JIRA_PASSWORD')
    if not password:
        password = getpass.getpass('Bitbucket / JIRA password: ', sys.stderr)

    bb_client = stashy.connect(BITBUCKET_SERVER_URL, username, password)

    # set environment variables so that JIRA client works
    os.environ['JIRA_USERNAME'] = username
    os.environ['JIRA_PASSWORD'] = password
    
    jira_client = JIRA(JIRA_OPTIONS, basic_auth=(username, password))
    return bb_client, jira_client


def get_dexreq_tickets_for_udx_tickets(jira_client, udx_ticket_keys):
    udx_tickets = []
    for udx_ticket_key in udx_ticket_keys:
        udx_issue = jira_client.issue(udx_ticket_key)

        udx_cli_ga_date = getattr(udx_issue.fields, CUSTOM_FIELD_UDX_CLI_GA_DATE)
        udx_sdk_ga_date = getattr(udx_issue.fields, CUSTOM_FIELD_UDX_SDK_GA_DATE)
        udx_console_ga_date = getattr(udx_issue.fields, CUSTOM_FIELD_UDX_CONSOLE_GA_DATE)
        udx_status = udx_issue.fields.status.name

        public_issues = util.search_dexreq_issues(PUBLIC_DEXREQ_ISSUE_FROM_UDX_TICKET_QUERY_FORMAT.format(udx_ticket=udx_ticket_key), fields=(DEFAULT_JIRA_ISSUE_FIELDS + [CUSTOM_FIELD_FEATURE_AVAILABLE_IN_PROD]))
        preview_issues = util.search_dexreq_issues(PREVIEW_DEXREQ_ISSUE_FROM_UDX_TICKET_QUERY_FORMAT.format(udx_ticket=udx_ticket_key))

        udx_ticket = {
            'key': udx_issue.key,
            'summary': udx_issue.fields.summary,
            'status': udx_status,
            'console_ga_date': udx_console_ga_date,
            'sdk_ga_date': udx_sdk_ga_date,
            'cli_ga_date': udx_cli_ga_date
        }

        if len(public_issues) > 0:
            udx_ticket['public_issues'] = []
            for issue in public_issues:
                field_value = getattr(issue.fields, CUSTOM_FIELD_FEATURE_AVAILABLE_IN_PROD)
                available_in_prod = field_value and field_value.value == 'Yes'
                udx_ticket['public_issues'].append({
                    'key': issue.key,
                    'summary': issue.fields.summary,
                    'available_in_prod': available_in_prod
                })

        if len(preview_issues) > 0:
            udx_ticket['preview_issues'] = []
            for issue in preview_issues:
                udx_ticket['preview_issues'].append({
                    'key': issue.key,
                    'summary': issue.fields.summary
                })

        udx_tickets.append(udx_ticket)

    return udx_tickets


def does_pr_match_any_dexreqs(pull_request, dexreq_public_issue_keys, dexreq_preview_issue_keys, udx_issue_key):
    # do some loose matching, sometimes people will improperly reference the ticket (e.g. UDX123 instead of UDX-123)
    dexreq_public_issue_matches = False
    dexreq_preview_issue_matches = False

    for public_key in dexreq_public_issue_keys:
        public_issue_key_without_dash = public_key if public_key is None else public_key.replace('-', '')
        dexreq_public_issue_matches = dexreq_public_issue_matches or is_string_referenced_by_pr(public_key, pull_request) or is_string_referenced_by_pr(public_issue_key_without_dash, pull_request)

    for preview_key in dexreq_preview_issue_keys:
        preview_issue_key_without_dash = preview_key if preview_key is None else preview_key.replace('-', '')
        dexreq_preview_issue_matches = dexreq_preview_issue_matches or is_string_referenced_by_pr(preview_key, pull_request) or is_string_referenced_by_pr(preview_issue_key_without_dash, pull_request)

    udx_issue_key_without_dash = udx_issue_key.replace('-', '')
    udx_issue_matches = is_string_referenced_by_pr(udx_issue_key, pull_request) or is_string_referenced_by_pr(udx_issue_key_without_dash, pull_request)

    return dexreq_public_issue_matches or dexreq_preview_issue_matches or udx_issue_matches


def is_string_referenced_by_pr(text, pull_request):
    return text is not None and (text in pull_request.get('title', '').lower() or text in pull_request.get('description', '').lower() or text in pull_request['fromRef']['id'].lower())


def add_build_info_for_pr(pull_request):
    latest_commit = pull_request['fromRef']['latestCommit']

    headers = {
        "X-Atlassian-Token": "no-check",
        "Content-Type": "application/json"
    }

    response = requests.get(
        url=BITBUCKET_SERVER_URL + BITBUCKET_BUILD_STATUS_API_URL_PATH_FORMAT.format(latest_commit),
        auth=(os.environ['JIRA_USERNAME'], os.environ['JIRA_PASSWORD']),
        headers=headers,
        verify=False
    )

    r = json.loads(response.content.decode('UTF-8'))
    pull_request['builds'] = r['values']

    _throttle()


def add_approver_info(pull_request):
    for reviewer in pull_request['reviewers']:
        if reviewer['approved']:
            if 'emailAddress' in reviewer['user'] and reviewer['user']['emailAddress'] in DEX_TEAM_APPROVER_EMAILS:
                pull_request['has_dex_approver'] = True
            else:
                pull_request['has_non_dex_approver'] = True

        if reviewer['status'] == 'NEEDS_WORK':
            pull_request['needs_work'] = True


# this is slightly confusing since if both are missing we only say 'Non DEX approver' is missing
# if there is a DEX approver then we don't really care that we are missing a non-dex approver
# these checks are all used solely to determine if the PR is ready to be reviewed by DEX, and then
# separately we check if it has a DEX approver to move it into the final READY state
def check_has_any_approver(pull_request):
    check_pass = pull_request.get('has_non_dex_approver') or pull_request.get('has_dex_approver')
    if not check_pass:
        pull_request['missing'].append('No non-DEX approver')
    
    return check_pass


def check_no_reviewers_marked_needs_work(pull_request):
    check_pass = not pull_request.get('needs_work')
    if not check_pass:
        pull_request['missing'].append('PR marked Needs Work')
    
    return check_pass


def check_no_merge_conflict(pull_request):
    check_pass = not (pull_request['properties'].get('mergeResult') and pull_request['properties']['mergeResult']['outcome'] == "CONFLICTED")
    if not check_pass:
        pull_request['missing'].append('Merge conflict')

    return check_pass


def check_all_builds_passing(pull_request):
    check_pass = True
    for build in pull_request['builds']:
        if build['state'] == 'FAILED':
            check_pass = False
            break
    
    if not check_pass:
        pull_request['missing'].append('Not all builds passing')

    return check_pass


def check_any_build_passing(pull_request):
    check_pass = False
    for build in pull_request['builds']:
        if build['state'] == 'SUCCESSFUL':
            check_pass = True
            break
    
    if not check_pass:
        pull_request['missing'].append('No build passing')

    return check_pass


def check_has_samples(pull_request):
    check_pass = pull_request.get('has_samples')
    if not check_pass:
        pull_request['missing'].append('No samples')

    return check_pass


def check_has_recordings(pull_request):
    check_pass = pull_request.get('has_recordings')
    if not check_pass:
        pull_request['missing'].append('No recordings')

    return check_pass


def check_has_changelog(pull_request):
    # changelog is not required for preview
    if pull_request['toRef']['id'] == 'refs/heads/preview':
        return True

    check_pass = pull_request.get('has_changelog')
    if not check_pass:
        pull_request['missing'].append('No changelog')

    return check_pass


# this function determines which state the PR is in so that we can color it appropriately
# states:
#    - MERGED - GRAY - PR is merged
#    - READY FOR MERGE - GREEN - all pre-reqs are met on call could press 'MERGE'
#    - READY FOR DEX PR - BLUE - all pre-reqs are met EXCEPT DEX approved
#    - MISSING ANY PRE-REQS - YELLOW
def add_pr_overall_status(pull_request):
    ready_for_dex_review_check_per_repo = {
        PYTHON_CLI_REPO_SLUG: [
            check_has_any_approver,
            check_no_merge_conflict,
            check_has_samples,
            check_has_recordings,
            check_has_changelog,
            check_any_build_passing,
            check_no_reviewers_marked_needs_work
        ],
        JAVA_SDK_REPO_SLUG: [
            check_has_any_approver,
            check_no_merge_conflict,
            check_has_samples,
            check_all_builds_passing,
            check_no_reviewers_marked_needs_work
        ],
        OCI_TESTING_SERVICE_REPO_SLUG: [
            check_has_any_approver,
            check_no_merge_conflict,
            check_all_builds_passing,
            check_no_reviewers_marked_needs_work
        ],
        PYTHON_SDK_REPO_SLUG: [
            check_has_any_approver,
            check_no_merge_conflict,
            check_all_builds_passing,
            check_no_reviewers_marked_needs_work
        ],
        GO_SDK_REPO_SLUG: [
            check_has_any_approver,
            check_no_merge_conflict,
            check_all_builds_passing,
            check_no_reviewers_marked_needs_work
        ],
        RUBY_SDK_REPO_SLUG: [
            check_has_any_approver,
            check_no_merge_conflict,
            check_all_builds_passing,
            check_no_reviewers_marked_needs_work
        ]
    }

    pull_request['missing'] = []
    repo_slug = pull_request['toRef']['repository']['slug']

    ready_for_dex_review = True
    for check in ready_for_dex_review_check_per_repo[repo_slug]:
        ready_for_dex_review = check(pull_request) and ready_for_dex_review

    ready_for_merge = ready_for_dex_review and pull_request.get('has_dex_approver')

    if pull_request['state'] == 'MERGED':
        # PR is merged, that status takes precedence over everyone else
        pull_request['dexreq_status_class'] = 'list-group-item-secondary'
    elif ready_for_merge:
        # EVERY pre-req is met for this, all on-call has to do is click merge
        # only exception is for CLI, we allow the build against master PySDK to be failing
        pull_request['dexreq_status_class'] = 'list-group-item-success'
    elif ready_for_dex_review:
        # every pre-req is met EXCEPT review by DEX team member
        pull_request['dexreq_status_class'] = 'list-group-item-primary'
    else:
        # missing some pre-reqs (will be displayed in 'Missing')
        pull_request['dexreq_status_class'] = 'list-group-item-warning'


def has_cli_samples(change):
    components = change['path']['components']
    return len(components) > 2 and components[0] == 'scripts' and components[1] == 'examples'


def has_java_samples(change):
    components = change['path']['components']
    return len(components) > 0 and components[0] == 'bmc-examples'


def has_cli_recordings(change):
    components = change['path']['components']
    yml_extension = 'extension' in change['path'] and change['path']['extension'] == 'yml'
    return yml_extension and \
        len(components) > 5 and \
        'cassettes' in components and \
        'tests' in components


def has_cli_changelog(change):
    components = change['path']['components']
    return len(components) > 0 and components[0] == 'changelog_entries'


def add_pr_pre_req_checks_to_pr(bb_client, pull_request):
    pr_id = pull_request['id']
    repo_slug = pull_request['toRef']['repository']['slug']

    detect_samples_func_for_repo = {
        PYTHON_CLI_REPO_SLUG: has_cli_samples,
        JAVA_SDK_REPO_SLUG: has_java_samples
    }

    detect_recordings_func_for_repo = {
        PYTHON_CLI_REPO_SLUG: has_cli_recordings
    }

    detect_changelog_func_for_repo = {
        PYTHON_CLI_REPO_SLUG: has_cli_changelog
    }

    detect_samples_func = detect_samples_func_for_repo.get(repo_slug)
    detect_recordings_func = detect_recordings_func_for_repo.get(repo_slug)
    detect_changelog_func = detect_changelog_func_for_repo.get(repo_slug)

    headers = {
        "X-Atlassian-Token": "no-check",
        "Content-Type": "application/json"
    }

    is_last_page = False
    start = 'null'

    has_samples = False
    has_recordings = False
    has_changelog = False

    # we can quit paginating early if we find what we're looking for
    while not is_last_page and (not has_samples or not has_recordings or not has_changelog):
        response = requests.get(
            url=BITBUCKET_SERVER_URL + BITBUCKET_PULL_REQUEST_CHANGES_API_URL_PATH_FORMAT.format(repo_slug, pr_id) + '?start={}'.format(start),
            auth=(os.environ['JIRA_USERNAME'], os.environ['JIRA_PASSWORD']),
            headers=headers,
            verify=False
        )

        r = json.loads(response.content.decode('UTF-8'))
        for change in r['values']:
            if detect_samples_func:
                if detect_samples_func(change):
                    has_samples = True

            if detect_recordings_func:
                if detect_recordings_func(change):
                    has_recordings = True

            if detect_changelog_func:
                if detect_changelog_func(change):
                    has_changelog = True

        is_last_page = r['isLastPage']
        start = r['nextPageStart']

    pull_request['has_samples'] = has_samples
    pull_request['has_recordings'] = has_recordings
    pull_request['has_changelog'] = has_changelog

    _throttle()


def generate_report(bb_client, jira_client, issues):
    print('Fetching PRs for all tickets...')
    all_testing_service_prs_preview = get_all_pull_requests(bb_client, REPO_SLUG_TESTING_SERVICE, PREVIEW_BRANCH_REF, "ALL")
    all_java_prs_preview = get_all_pull_requests(bb_client, REPO_SLUG_JAVA_SDK, PREVIEW_BRANCH_REF, "ALL")
    all_cli_prs_preview = get_all_pull_requests(bb_client, REPO_SLUG_CLI, PREVIEW_BRANCH_REF, "ALL")
    all_ruby_prs_preview = get_all_pull_requests(bb_client, REPO_SLUG_RUBY_SDK, PREVIEW_BRANCH_REF, "ALL")
    all_python_prs_preview = get_all_pull_requests(bb_client, REPO_SLUG_PYTHON_SDK, PREVIEW_BRANCH_REF, "ALL")
    all_go_prs_preview = get_all_pull_requests(bb_client, REPO_SLUG_GO_SDK, PREVIEW_BRANCH_REF, "ALL")

    all_testing_service_prs_master = get_all_pull_requests(bb_client, REPO_SLUG_TESTING_SERVICE, MASTER_BRANCH_REF, "ALL")
    all_java_prs_master = get_all_pull_requests(bb_client, REPO_SLUG_JAVA_SDK, MASTER_BRANCH_REF, "ALL")
    all_cli_prs_master = get_all_pull_requests(bb_client, REPO_SLUG_CLI, MASTER_BRANCH_REF, "ALL")
    all_ruby_prs_master = get_all_pull_requests(bb_client, REPO_SLUG_RUBY_SDK, MASTER_BRANCH_REF, "ALL")
    all_python_prs_master = get_all_pull_requests(bb_client, REPO_SLUG_PYTHON_SDK, MASTER_BRANCH_REF, "ALL")
    all_go_prs_master = get_all_pull_requests(bb_client, REPO_SLUG_GO_SDK, MASTER_BRANCH_REF, "ALL")

    udx_tickets = get_dexreq_tickets_for_udx_tickets(jira_client, issues)
    for udx_ticket in udx_tickets:
        udx_ticket_key = udx_ticket['key'].lower()

        preview_prs = {}
        dexreq_preview_issue_keys = []
        if 'preview_issues' in udx_ticket:
            dexreq_preview_issue_keys = [x['key'].lower() for x in udx_ticket['preview_issues']]

        public_prs = {}
        dexreq_public_issue_keys = []
        if 'public_issues' in udx_ticket:
            dexreq_public_issue_keys = [x['key'].lower() for x in udx_ticket['public_issues']]

        preview_prs['testing'] = []
        preview_prs['java'] = []
        preview_prs['cli'] = []
        preview_prs['ruby'] = []
        preview_prs['python'] = []
        preview_prs['go'] = []

        public_prs['testing'] = []
        public_prs['java'] = []
        public_prs['cli'] = []
        public_prs['ruby'] = []
        public_prs['python'] = []
        public_prs['go'] = []

        # find all preview PRs that match either the public dexreq, preview dexreq, or UDX
        preview_prs['testing'].extend([pr for pr in all_testing_service_prs_preview if does_pr_match_any_dexreqs(pr, dexreq_public_issue_keys, dexreq_preview_issue_keys, udx_ticket_key)])
        preview_prs['java'].extend([pr for pr in all_java_prs_preview if does_pr_match_any_dexreqs(pr, dexreq_public_issue_keys, dexreq_preview_issue_keys, udx_ticket_key)])
        preview_prs['cli'].extend([pr for pr in all_cli_prs_preview if does_pr_match_any_dexreqs(pr, dexreq_public_issue_keys, dexreq_preview_issue_keys, udx_ticket_key)])
        preview_prs['ruby'].extend([pr for pr in all_ruby_prs_preview if does_pr_match_any_dexreqs(pr, dexreq_public_issue_keys, dexreq_preview_issue_keys, udx_ticket_key)])
        preview_prs['python'].extend([pr for pr in all_python_prs_preview if does_pr_match_any_dexreqs(pr, dexreq_public_issue_keys, dexreq_preview_issue_keys, udx_ticket_key)])
        preview_prs['go'].extend([pr for pr in all_go_prs_preview if does_pr_match_any_dexreqs(pr, dexreq_public_issue_keys, dexreq_preview_issue_keys, udx_ticket_key)])

        # find all master PRs that match either the public dexreq, preview dexreq, or UDX
        public_prs['testing'].extend([pr for pr in all_testing_service_prs_master if does_pr_match_any_dexreqs(pr, dexreq_public_issue_keys, dexreq_preview_issue_keys, udx_ticket_key)])
        public_prs['java'].extend([pr for pr in all_java_prs_master if does_pr_match_any_dexreqs(pr, dexreq_public_issue_keys, dexreq_preview_issue_keys, udx_ticket_key)])
        public_prs['cli'].extend([pr for pr in all_cli_prs_master if does_pr_match_any_dexreqs(pr, dexreq_public_issue_keys, dexreq_preview_issue_keys, udx_ticket_key)])
        public_prs['ruby'].extend([pr for pr in all_ruby_prs_master if does_pr_match_any_dexreqs(pr, dexreq_public_issue_keys, dexreq_preview_issue_keys, udx_ticket_key)])
        public_prs['python'].extend([pr for pr in all_python_prs_master if does_pr_match_any_dexreqs(pr, dexreq_public_issue_keys, dexreq_preview_issue_keys, udx_ticket_key)])
        public_prs['go'].extend([pr for pr in all_go_prs_master if does_pr_match_any_dexreqs(pr, dexreq_public_issue_keys, dexreq_preview_issue_keys, udx_ticket_key)])

        udx_ticket['pull_requests'] = {
            "preview": preview_prs,
            "master": public_prs
        }

        print('Fetching build info for all PRs for {}...'.format(udx_ticket['key']))
        for lang,prs in six.iteritems(preview_prs):
            for pr in prs:
                add_build_info_for_pr(pr)
                add_pr_pre_req_checks_to_pr(bb_client, pr)
                add_approver_info(pr)
                add_pr_overall_status(pr)

        for lang,prs in six.iteritems(public_prs):
            for pr in prs:
                add_build_info_for_pr(pr)
                add_pr_pre_req_checks_to_pr(bb_client, pr)
                add_approver_info(pr)
                add_pr_overall_status(pr)

    seattle_now = datetime.now(pytz.timezone(TIME_ZONE_FOR_DATE_FORMATS))
    response = {
        "report_name": REPORT_NAME,
        "generated_time": seattle_now.strftime(DATE_FORMAT),
        "tickets": udx_tickets
    }

    udx_ticket_stamp = '_'.join(issues)  # noqa:F841
    timestamp = datetime.utcnow().strftime('%m_%d_%H_%M')  # noqa:F841

    if not os.path.exists(OUTPUT_DIRECTORY):
        os.makedirs(OUTPUT_DIRECTORY)

    preview_report_location = os.path.abspath(os.path.join(OUTPUT_DIRECTORY, '{report_name}_preview.html'.format(report_name=REPORT_NAME)))
    public_report_location = os.path.abspath(os.path.join(OUTPUT_DIRECTORY, '{report_name}_public.html'.format(report_name=REPORT_NAME)))
    full_report_location = os.path.abspath(os.path.join(OUTPUT_DIRECTORY, '{report_name}_full.html'.format(report_name=REPORT_NAME)))

    # create 3 reports, one for public only, one for preview, and one for combined
    with open(preview_report_location, 'w') as preview_report_file:
        response['show_preview'] = True
        response['show_public'] = False

        preview_report_file.write(render_html(response))

    with open(public_report_location, 'w') as public_report_file:
        response['show_preview'] = False
        response['show_public'] = True

        public_report_file.write(render_html(response))

    with open(full_report_location, 'w') as public_report_file:
        response['show_preview'] = True
        response['show_public'] = True

        public_report_file.write(render_html(response))

    if UPLOAD_TO_OBJECT_STORAGE:
        upload_reports_to_object_storage([
            {
                "location": preview_report_location,
                "report_type": "preview"
            },
            {
                "location": public_report_location,
                "report_type": "public"
            },
            {
                "location": full_report_location,
                "report_type": "full"
            }
        ])

        print('Reports uploaded to object storage. Tenancy: {}. Bucket: {}'.format(REPORTS_NAMESPACE, REPORTS_BUCKET_NAME))

    print('Preview report written to: {}'.format(preview_report_location))
    print('Public report written to: {}'.format(public_report_location))
    print('Full report written to: {}'.format(full_report_location))


# - Uploads the report files to object storage in the dexreq_reports bucket
# - If OCI Python SDK is configured, will also attempt to update the reports_index
#   which contains PARs for all existing reports
# - This is used to render the reports homepage
def upload_reports_to_object_storage(reports):
    result = requests.get(url=OBJECT_STORAGE_DEXREQ_REPORTS_INDEX_RW_PAR)
    reports_index = json.loads(result.content.decode('UTF-8'))
    reports_from_server = reports_index['reports']

    new_reports = {}
    for report in reports:
        file_path = report['location']
        file_name = os.path.basename(report['location'])
        report_type = report['report_type']

        # upload report to object storage
        with open(file_path, 'r') as f:
            requests.put(
                url=OBJECT_STORAGE_DEXREQ_REPORTS_PAR + file_name,
                data=f.read(),
                headers={
                    "Content-Type": "text/html"
                }
            )

        # check if PAR for this report already exists in reports index, if not, create it
        if isinstance(reports_from_server, dict):
            if REPORT_NAME in reports_from_server and report_type in reports_from_server[REPORT_NAME]:
                print('Report {}: {} already exists in index, skipping creating new PAR...'.format(REPORT_NAME, report_type))
                continue
        else:
            print('Failed to retrieve reports index from object storage, not creating PAR or adding to index')
            continue

        # no existing PAR for this reoprt so create one
        # catch all exceptions because this part is optional if user hasn't configured OCI Python SDK
        try:
            config = oci.config.from_file(file_location=CONFIG_FILE_LOCATION)
            if AUTO_GEN_API_KEY_PASS_PHRASE_ENV_VAR not in os.environ:
                raise ValueError('No passphrase specified for API key. Please populate environment varaible: {}'.format(AUTO_GEN_API_KEY_PASS_PHRASE_ENV_VAR))

            config['pass_phrase'] = os.environ[AUTO_GEN_API_KEY_PASS_PHRASE_ENV_VAR]
            object_storage_client = oci.object_storage.ObjectStorageClient(config)
            create_par_details = oci.object_storage.models.CreatePreauthenticatedRequestDetails(
                name='Read Access for DEXREQ Report: {} - {}'.format(REPORT_NAME, report_type),
                object_name=file_name,
                access_type=oci.object_storage.models.CreatePreauthenticatedRequestDetails.ACCESS_TYPE_OBJECT_READ,
                time_expires=datetime.today() + timedelta(2 * 365 / 12)  # two months from now
            )

            response = object_storage_client.create_preauthenticated_request(REPORTS_NAMESPACE, REPORTS_BUCKET_NAME, create_par_details)
            if response.status != 200:
                print('Failed creating par: ' + str(response.data))
                continue

            par = response.data
            new_reports[report_type] = 'https://objectstorage.us-phoenix-1.oraclecloud.com' + par.access_uri

        except Exception as e:
            print('Failed to create new PAR for report')
            print(str(e))
            continue

    # if there are any new reports, update the report index with them
    if new_reports and isinstance(reports_from_server, dict):
        reports_from_server[REPORT_NAME] = new_reports

        print('Updated reports index: ' + str(reports_from_server))

        result = requests.put(
            url=OBJECT_STORAGE_DEXREQ_REPORTS_INDEX_RW_PAR,
            headers={'Content-Type': 'application/json'},
            data=json.dumps(reports_index)
        )


def get_udx_tickets_with_sdk_cli_ga_date(bb_client, jira_client, ga_date):
    all_udx_tickets = set()
    public_issues_with_ga_date = util.search_dexreq_issues(PUBLIC_WITH_GA_DATE_QUERY_FORMAT.format(ga_date))
    for public_issue in public_issues_with_ga_date:
        udx_keys = re.findall("UDX-[0-9]+", public_issue.fields.summary)
        if not udx_keys:
            udx_ticket_field_value = getattr(public_issue.fields, CUSTOM_FIELD_ID_UDX_TICKET)
            if udx_ticket_field_value:
                udx_keys = [x.strip() for x in udx_ticket_field_value.split(',')]

        if not udx_keys:
            print('WARNING: Did not find any corresponding UDX tickets for: {}'.format(public_issue.key))

        all_udx_tickets.update(udx_keys)

    return all_udx_tickets


def validate_date_format(date_text):
    try:
        datetime.strptime(date_text, '%Y-%m-%d')
    except ValueError:
        return False

    return True


def _throttle():
    # no need to sleep if we are running against recordings
    if VCR_INSTANCE.record_mode != "none":
        time.sleep(SLEEP_BETWEEN_REQUESTS)


# 0 = Monday, 1=Tuesday, 2=Wednesday...
def next_weekday(d, weekday):
    days_ahead = weekday - d.weekday()
    if days_ahead <= 0:  # Target day already happened this week
        days_ahead += 7
    return d + timedelta(days_ahead)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Print information about in progress and complete pull requests for a given DEXREQ issue.')
    parser.add_argument('--issues',
                        help='A comma separated list of UDX tickets to generate a report for')
    parser.add_argument('--limit',
                        type=int,
                        help='The maximum number of PRs to look back over. This is to avoid querying the entire PR history unnecessarily')
    parser.add_argument('--output',
                        default='html',
                        help='Output format. Options are: HTML')
    parser.add_argument('--output-dir',
                        default=OUTPUT_DIRECTORY,
                        help='Output directory. Default is {}'.format(OUTPUT_DIRECTORY))
    parser.add_argument('--report-name',
                        default=REPORT_NAME,
                        help="""The name of the report to generate.  This will be generated in the output directory.
                        Reports will have the name {{report-name}}_full.html, {{report-name}}_preview.html, an {{report-name}}_public.html.""")
    parser.add_argument('--upload-to-object-storage',
                        default=False,
                        action='store_true',
                        help="""Whether or not to upload reports to object storage in dexreq_reports bucket.""")
    parser.add_argument('--vcr-record-mode',
                        default='all',
                        help="""VCR record mode to use: https://vcrpy.readthedocs.io/en/latest/usage.html#record-modes""")
    parser.add_argument('--sdk-cli-ga-date',
                        help="""Create a report for all DEXREQ tickets with the given GA Date (YYYY-MM-DD).
                        This uses the SDK / CLI GA Date field on the public DEXREQ ticket and then looks up the corresponding UDX tickets.
                        This parameter may not be used with --issues""")
    parser.add_argument('--next-release-n',
                        type=int,
                        default=-1,
                        help="""Create a report for all DEXREQ tickets for the nth release date from now.
                        For example '0' indicates the next release, '1' indicates 2 releases from now.""")

    args = parser.parse_args()
    issues = args.issues
    limit = args.limit
    output = args.output
    output_dir = args.output_dir
    report_name = args.report_name
    upload_to_object_storage = args.upload_to_object_storage
    vcr_record_mode = args.vcr_record_mode
    sdk_cli_ga_date = args.sdk_cli_ga_date
    next_release_n = args.next_release_n

    print('Starting...')

    if issues:
        issues = [issue.strip() for issue in issues.split(',')]
    else:
        issues = []

    if not issues and not sdk_cli_ga_date and next_release_n == -1:
        sys.exit('Must specify either --issues or --sdk-cli-ga-date or --next-release-n parameter')

    if not (bool(issues) ^ bool(sdk_cli_ga_date) ^ bool(next_release_n != -1)):
        sys.exit('Can only use one of --issues, --sdk-cli-ga-date, and --next-release-n parameters')

    if sdk_cli_ga_date and not validate_date_format(sdk_cli_ga_date):
        sys.exit('--sdk-cli-ga-date was not in valid date format (YYYY-MM-DD)')

    if limit:
        PULL_REQUEST_PAGINATION_LIMIT = limit

    if output_dir:
        OUTPUT_DIRECTORY = output_dir

    if upload_to_object_storage:
        UPLOAD_TO_OBJECT_STORAGE = True

    if report_name:
        REPORT_NAME = report_name

    if next_release_n != -1:
        today = date.today() + timedelta(7 * next_release_n)
        next_tuesday_date_str = next_weekday(today, 1).isoformat()
        REPORT_NAME = 'release_{}'.format(next_tuesday_date_str.replace('-', '_'))
        sdk_cli_ga_date = next_tuesday_date_str
        print('Fetching tickets for release date: {}'.format(next_tuesday_date_str))

    VCR_INSTANCE = vcr.VCR(
        serializer='yaml',
        record_mode=vcr_record_mode
    )

    with VCR_INSTANCE.use_cassette('cassettes/dexreq_pr_tracker.yaml'):
        bb_client, jira_client = build_clients()
        udx_tickets = issues
        if sdk_cli_ga_date:
            udx_tickets = get_udx_tickets_with_sdk_cli_ga_date(bb_client, jira_client, sdk_cli_ga_date)

        print("Generating report for tickets: {}".format(', '.join(udx_tickets)))
        generate_report(bb_client, jira_client, udx_tickets)
