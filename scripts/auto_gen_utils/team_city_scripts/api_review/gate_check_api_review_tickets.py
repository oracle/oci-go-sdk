import os
import sys
import argparse
import re
import requests
import logging

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../..'))

import util  # noqa: ignore=F402
import config  # noqa: ignore=F402
from shared.bitbucket_utils import get_pullrequest, setup_bitbucket, get_pullrequest_diff, get_pullrequest_changes, get_pullrequest_merge_status, get_file_content_from_commit_id_and_path # noqa: ignore=F402

# Constants
API_SPEC_VALIDATION_OUTPUT_FILE_PATTERN = 'validation-output'
API_SPEC_CONFIG_FILE_PATTERN = 'validation-config'
API_SEPC_FILE_PATTERN = 'cond.yaml'
PR_SEARCH_PATTERN = '^.*bitbucket.*/projects/([^/]*)/repos/([^/]*)/pull-requests/([0-9]*).*$'
SPEC_VALIDATOR_PATTERN = '`(.*)`'
SPEC_VALIDATOR_VERSION = 'Validator version'
SPEC_ERRORS = 'Errors'
SPEC_WARNINGS = 'Warnings'
SPEC_SUPPRESSED_ERRORS = 'Suppressed errors'
SPEC_SUPPRESSED_WARNINGS = 'Suppressed warnings'
GATE_CHECK_PASSED = True
MANUAL_CHECK_REQUIRED = False
GATE_CHECK_FAIL_COMMENTS = []
GATE_CHECK_PASS_COMMENTS = []

# Configurable TC Job environment variables
GATE_CHECK_PASS_LABEL = os.environ.get('GATE_CHECK_PASS_LABEL') or 'APIRB-Ready-Check-Pass'
BYPASS_GATE_CHECK_LABEL = os.environ.get('BYPASS_GATE_CHECK_LABEL') or 'ByPass-APIRB-Ready-Check'
BYPASS_UDX_CHECK_LABEL = os.environ.get('BYPASS_UDX_CHECK_LABEL') or 'Bypass-APIRB-Ready-Check-UDX'
MANUAL_CHECK_LABEL = os.environ.get('MANUAL_CHECK_LABEL') or 'manual-check-required'
TICKETS_IN_TRIAGE_QUERY = 'project = APIReviewBoard AND status in ("Needs Triage") and (labels not in ({}, {}) or labels is EMPTY)'.format(GATE_CHECK_PASS_LABEL, BYPASS_GATE_CHECK_LABEL)
QUERY_TEMPLATE = os.environ.get('QUERY') or TICKETS_IN_TRIAGE_QUERY
DEFAULT_SPEC_VALIDATOR_TOOL_VERSION = os.environ.get('DEFAULT_SPEC_VALIDATOR_TOOL_VERSION') or '1.95.0'
VALIDATOR_TOOL_VERSION_CUTOFF = int(os.environ.get('VALIDATOR_TOOL_VERSION_CUTOFF')) if 'VALIDATOR_TOOL_VERSION_CUTOFF' in os.environ else 3
API_REVIEW_PR_PREFIX = os.environ.get('API_REVIEW_PR_PREFIX') or 'API Review PR:'
DEFAULT_TICKET_TRIAGE_ASSIGNEE = os.environ.get('DEFAULT_TICKET_TRIAGE_ASSIGNEE') or 'arshanm'
AUTOMATION_USER = os.environ.get('BITBUCKET_USERNAME') or 'gear-dexreq-automation'

# create logger
logging.basicConfig()
logger = logging.getLogger('APIRB_TICKET_CHECK')
logger.setLevel(logging.DEBUG)


def init_gate_check():
    global GATE_CHECK_PASSED
    global MANUAL_CHECK_REQUIRED
    global GATE_CHECK_FAIL_COMMENTS
    global GATE_CHECK_PASS_COMMENTS
    GATE_CHECK_PASSED = True
    MANUAL_CHECK_REQUIRED = False
    GATE_CHECK_FAIL_COMMENTS = []
    GATE_CHECK_PASS_COMMENTS = []


def query_actionable_tickets():
    query = QUERY_TEMPLATE
    actionable_tickets = util.jira_search_issues(query)
    print('Total of {} actionable tickets found matching query `{}`.'.format(len(actionable_tickets), query))
    return actionable_tickets


def is_udx_check_bypassed(api_ticket):
    labels = api_ticket.fields.labels
    if BYPASS_UDX_CHECK_LABEL in labels:
        return True
    return False


def udx_ticket_check(api_ticket):
    if is_udx_check_bypassed(api_ticket):
        print('Skipping UDX ticket check for ticket:{}'.format(api_ticket.key))
        return

    links = []
    for link in api_ticket.fields.issuelinks:
        if hasattr(link, 'outwardIssue'):
            links.append(link.outwardIssue)
        if hasattr(link, 'inwardIssue'):
            links.append(link.inwardIssue)
    udx_tickets = list(filter(lambda x: 'UDX' in x.key, links))
    if len(udx_tickets) == 0:
        comment = 'There is no UDX ticket linked to this ticket. Please add the link to the UDX ticket this API Review' \
                  ' is for, under the *issue links* section of this ticket. If this review is for an internal API,' \
                  'then, you may skip the UDX check by adding the label: {} to the ticket.'.format(BYPASS_UDX_CHECK_LABEL)
        add_gate_check_fail_comment(comment)
    else:
        print("Found related UDX ticket:{}\n".format(udx_tickets[0]))


def get_api_spec_pr_link_from_ticket(api_ticket):
    """
    Get PR link from ticket
    @param api_ticket: API Review Board ticket being checked for API Review ticket readiness
    @return: pr_link
    """
    pr_link = None
    description = api_ticket.fields.description
    if description is None:
        return pr_link
    # Parse description to get bitbucket API PR Review link
    for line in description.splitlines():
        line = line.strip().replace('*', '')
        if API_REVIEW_PR_PREFIX in line:
            try:
                pr_link = line.split(API_REVIEW_PR_PREFIX, 1)[1].strip()
                print('PR link found: {}'.format(pr_link))
                return pr_link
            except IndexError:
                break
    return None


def is_pr_approved(pr):
    """
    Check if PR has been peer reviewed by at-least one person
    @param pr: Bitbucket Pull Request
    @return: bool
    """
    if 'reviewers' not in pr.json():
        # No reviewers assigned to PR
        return False
    for reviewer in pr.json()['reviewers']:
        if reviewer['status'] == 'APPROVED':
            add_get_check_pass_comment('PR has been approved by {}'.format(reviewer['user']['displayName']))
            return True
    return False


def get_pr_info(url):
    m = re.search(PR_SEARCH_PATTERN, url)
    if m:
        return m.group(1), m.group(2), m.group(3)
    else:
        return None, None, None


def get_latest_spec_validation_tool_version():
    global MANUAL_CHECK_REQUIRED
    version = DEFAULT_SPEC_VALIDATOR_TOOL_VERSION
    object_storage_path = 'https://objectstorage.us-phoenix-1.oraclecloud.com'
    par = os.environ.get('PAR') or 'bfg3NAGcV8PHbO6teLBE9NWSVDZG99DHTF0o6SrQ6RU9Y-5hSJ_HKD0ghNhSAQBL'
    namespace = 'dex-us-phoenix-1'
    bucket = 'generated_markdown'
    version_file_path = 'bmc-sdk-swagger-validator-latest-non-snapshot-version.txt'
    url = object_storage_path + '/p/' + par + '/n/' + namespace + '/b/' + bucket + '/o/' + version_file_path
    response = requests.get(url)
    if response.status_code == 200:
        version = re.search(SPEC_VALIDATOR_PATTERN, response.text.split(" ").pop()).group(1)
        print('Found the latest spec validator version: {}'.format(version))
    else:
        comment = 'Manual Check Required: Call to get latest version of the Spec validator tool failed, Defaulting to: {}'.format(version)
        add_get_check_pass_comment(comment)
        MANUAL_CHECK_REQUIRED = True
    return version


def get_spec_prefix(value):
    return 'INFO  ' + value


def parse_spec_validator_file(spec_validator_file):
    """
    Parse Spec validator file for getting different values like Spec validator tool version, Errors, etc
    @param spec_validator_file: Spec validation file contents
    @return: mapping for parsed contents
    """
    spec_validator_mapping = dict()
    for line in spec_validator_file:
        if get_spec_prefix(SPEC_VALIDATOR_VERSION) in line:
            spec_validator_mapping[SPEC_VALIDATOR_VERSION] = line.split(" ").pop()
        elif get_spec_prefix(SPEC_ERRORS) in line:
            spec_validator_mapping[SPEC_ERRORS] = line.split(" ").pop()
        elif get_spec_prefix(SPEC_WARNINGS) in line:
            spec_validator_mapping[SPEC_WARNINGS] = line.split(" ").pop()
        elif get_spec_prefix(SPEC_SUPPRESSED_ERRORS) in line:
            spec_validator_mapping[SPEC_SUPPRESSED_ERRORS] = line.split(" ").pop()
        elif SPEC_SUPPRESSED_WARNINGS in line:
            spec_validator_mapping[SPEC_SUPPRESSED_WARNINGS] = line.split(" ").pop()

    return spec_validator_mapping


def is_spec_validation_tool_version_valid(version_used, latest_version):
    print('Spec Validation tool version used: {} Latest tool version: {}'.format(version_used, latest_version))
    version_used = version_used.split('.')
    latest_version = latest_version.split('.')
    if version_used[0] == latest_version[0] \
            and (int(version_used[1]) + VALIDATOR_TOOL_VERSION_CUTOFF) >= int(latest_version[1]):
        return True
    return False


def is_spec_key_value_zero(file_under_test, spec_mapping, key):
    check = True
    try:
        if key in spec_mapping and int(spec_mapping[key]) > 0:
            comment = 'For file {}, API Review ticket readiness check found {} {} in Spec Validation output which needs to be fixed ' \
                      'or suppressed before a review by the API Review Board'.format(file_under_test, spec_mapping[key], key)
            add_gate_check_fail_comment(comment)
    except ValueError:
        comment = 'ERROR: Parsing the Spec Validator mapping content {}'.format(key)
        add_gate_check_fail_comment(comment)
        check = False
    comment = 'For file {}, API Review ticket readiness check found no {}'.format(file_under_test, key)
    add_get_check_pass_comment(comment)
    return check


def is_suppressed_warnings_or_suppressed_errors_increased(file_under_test, original_mappings, updated_mappings):
    check = False
    try:
        if SPEC_SUPPRESSED_ERRORS in original_mappings and SPEC_SUPPRESSED_ERRORS in updated_mappings:
            if int(original_mappings[SPEC_SUPPRESSED_ERRORS]) < int(updated_mappings[SPEC_SUPPRESSED_ERRORS]):
                print('For file {}, Increase in {} was found'.format(file_under_test, SPEC_SUPPRESSED_ERRORS))
                check = True
            else:
                print('For file {}, No Increase in {} was found'.format(file_under_test, SPEC_SUPPRESSED_ERRORS))
        else:
            print('Error: For file {}, Spec Validator Mapping key {} not found'.format(file_under_test, SPEC_SUPPRESSED_ERRORS))

        if SPEC_SUPPRESSED_WARNINGS in original_mappings and SPEC_SUPPRESSED_WARNINGS in updated_mappings:
            if int(original_mappings[SPEC_SUPPRESSED_WARNINGS]) < int(updated_mappings[SPEC_SUPPRESSED_WARNINGS]):
                print('For file {}, Increase in {} was found'.format(file_under_test, SPEC_SUPPRESSED_WARNINGS))
                check = True
            else:
                print('For file {}, No Increase in {} was found'.format(file_under_test, SPEC_SUPPRESSED_WARNINGS))
        else:
            print('Error: For file {}, Spec Validator Mapping key {} not found'.format(file_under_test, SPEC_SUPPRESSED_WARNINGS))

    except ValueError:
        print('Error: For file {}, Error parsing the Spec Validator Mapping values {}/{}'.format(file_under_test, SPEC_SUPPRESSED_ERRORS, SPEC_SUPPRESSED_WARNINGS))
    return check


def get_validation_config_file_from_output_file(output_file):
    file_name = os.path.splitext(output_file)[0]
    return file_name.replace(API_SPEC_VALIDATION_OUTPUT_FILE_PATTERN, API_SPEC_CONFIG_FILE_PATTERN) + '.yaml'


def get_pr_change_nodes_for_pattern(pr_changes, pattern):
    api_spec_node_list = []
    for node in pr_changes.json()['values']:
        if node and node.get('type') != 'DELETE' and node.get('path') and pattern in node['path'].get('name'):
            api_spec_node_list.append(node)
    return api_spec_node_list


def is_spec_validation_output_valid(project, repo, pr_id, pr_changes):
    global MANUAL_CHECK_REQUIRED
    check = True
    api_spec_nodes = get_pr_change_nodes_for_pattern(pr_changes, API_SPEC_VALIDATION_OUTPUT_FILE_PATTERN)
    if not api_spec_nodes:
        check = False
        comment = 'API Spec Validation output file was not found. ' + \
                  'Please make sure that the validation output file name contains the {} pattern'.format(API_SPEC_VALIDATION_OUTPUT_FILE_PATTERN) + \
                  'You may need to copy the validator output file from the target directory if your local build only sometimes copies it.'
        add_gate_check_fail_comment(comment)
        return check
    else:
        comment = 'API Spec Validation output file is Present'
        add_get_check_pass_comment(comment)

    # 1. Check if API spec diff contains valid spec-validation-output
    for api_spec_node in api_spec_nodes:
        file_under_test = api_spec_node['path']['name']
        print("Validating Spec validation output file:{}".format(file_under_test))
        updated_spec_validator_mappings = None
        original_spec_validator_mappings = None
        pr_diff = get_pullrequest_diff(project, repo, pr_id)
        if pr_diff and pr_diff.json() and not hasattr(pr_diff.json(), 'errors'):
            updated_pr_spec_validator_file = get_file_content_from_commit_id_and_path(project, repo, api_spec_node['path']['toString'], pr_diff.json()['toHash'])
            if updated_pr_spec_validator_file and updated_pr_spec_validator_file.content:
                updated_spec_validator_mappings = parse_spec_validator_file(updated_pr_spec_validator_file.content.split('\n'))

            original_pr_spec_validator_file = get_file_content_from_commit_id_and_path(project, repo, api_spec_node['path']['toString'], pr_diff.json()['fromHash'])
            if original_pr_spec_validator_file and original_pr_spec_validator_file.content:
                original_spec_validator_mappings = parse_spec_validator_file(original_pr_spec_validator_file.content.split('\n'))

        if updated_spec_validator_mappings:
            # 1.1 Check if validation tools used is within the VALIDATOR_TOOL_VERSION_CUTOFF of latest version
            print('For file {}, Spec validation Tool results found: {}'.format(file_under_test, updated_spec_validator_mappings))
            latest_spec_validation_tool_version = get_latest_spec_validation_tool_version()
            if is_spec_validation_tool_version_valid(updated_spec_validator_mappings[SPEC_VALIDATOR_VERSION], latest_spec_validation_tool_version):
                comment = 'For file {}, Spec Validation tool version used passes check, ' \
                          'Version used: {} and Latest Version is: {}'.format(file_under_test, updated_spec_validator_mappings[SPEC_VALIDATOR_VERSION], latest_spec_validation_tool_version)
                add_get_check_pass_comment(comment)
            else:
                check = False
                comment = 'For file {}, Spec Validation tool version used is older than allowed. ' \
                          'Please update it to version *{}* or later, re-run the tool and add its output to the PR '.format(file_under_test, latest_spec_validation_tool_version)
                add_gate_check_fail_comment(comment)

            # 1.2 Check if there are Non Zero errors in Spec Validation output
            if not is_spec_key_value_zero(file_under_test, updated_spec_validator_mappings, SPEC_ERRORS):
                check = False

            # 1.3 Check if there are Non Zero warnings in Spec Validation output
            if not is_spec_key_value_zero(file_under_test, updated_spec_validator_mappings, SPEC_WARNINGS):
                check = False

        else:
            print('Error: For file {}, Unable to parse Updated spec validation output file to mappings!'.format(file_under_test))

        # 1.4 Check if there was an increase in Suppressed Warnings and/or errors
        if original_spec_validator_mappings:
            if is_suppressed_warnings_or_suppressed_errors_increased(file_under_test, original_spec_validator_mappings, updated_spec_validator_mappings):
                print('For file {}, Spec Validation tool detected an increase in suppressed warnings and suppressed errors'.format(file_under_test))
                # 2.4.1 Check if Spec Validator Yaml file is present
                validation_config_file_to_search = get_validation_config_file_from_output_file(file_under_test)
                print('Searching For file {}, in PR'.format(validation_config_file_to_search))
                validation_config_files = get_pr_change_nodes_for_pattern(pr_changes, validation_config_file_to_search)
                if not validation_config_files:
                    comment = 'There was an increase in suppressed warnings/errors in Spec Validation output file {}, '\
                              'however, spec validation config file {} was not included in the PR. ' \
                              'Please add the corresponding spec validation config files to the PR. ' \
                              'If the validation config file is present with some other name, then you will need to ' \
                              'rename that file to match this pattern and update the PR'.format(file_under_test, validation_config_file_to_search)
                    add_gate_check_fail_comment(comment)
                    check = False
                else:
                    comment = 'There was an increase in suppressed warnings/errors in Spec Validation output file {} and ' \
                              'corresponding Spec Validation config file {} was found'.format(file_under_test, validation_config_file_to_search)
                    add_get_check_pass_comment(comment)
            else:
                comment = 'For file {}, No increase in suppressed warnings or suppressed errors was found'.format(file_under_test)
                add_get_check_pass_comment(comment)

        else:
            comment = 'Manual Check Required: For file {}, Unable to find Original version' \
                      'to check if there was an increase in Suppressed errors/warnings'.format(file_under_test)
            add_get_check_pass_comment(comment)
            MANUAL_CHECK_REQUIRED = True
    return check


def add_get_check_pass_comment(comment):
    global GATE_CHECK_PASS_COMMENTS
    print(comment)
    GATE_CHECK_PASS_COMMENTS.append(comment)


def add_gate_check_fail_comment(comment):
    global GATE_CHECK_PASSED
    global GATE_CHECK_FAIL_COMMENTS
    print(comment)
    GATE_CHECK_PASSED = False
    GATE_CHECK_FAIL_COMMENTS.append(comment)


def api_spec_pr_check(api_ticket):
    """
    Check if API Ticket has API Spec change PR which conforms to the requirements for API Review.
    @param api_ticket:
    @return: None
    """
    global GATE_CHECK_PASSED
    global GATE_CHECK_FAIL_COMMENTS

    pr_link = get_api_spec_pr_link_from_ticket(api_ticket)
    if pr_link is None:
        comment = 'API Review PR with prefix *{}* is missing in description. ' \
                  'Please add the PR link with this prefix in a separate line in the ticket description'.format(API_REVIEW_PR_PREFIX)
        add_gate_check_fail_comment(comment)

    else:
        project, repo, pr_id = get_pr_info(pr_link)
        # Check if PR is valid for our use-case
        if not project or not repo or not pr_id:
            comment = 'PR link: {} is invalid. The PR link should match the pattern described in the issue template'.format(pr_link)
            add_gate_check_fail_comment(comment)
            return

        pr = get_pullrequest(project, repo, pr_id)
        # Check if automation user has access to PR Repo
        if not pr or not pr.json() or hasattr(pr.json(), 'errors'):
            logger.error('Cannot access PR: {}'.format(pr))
            if pr.json():
                logger.error(pr.json())
            comment = 'The automation user *{}* does not have access to the Repo {}. ' \
                      'Kindly grant temporary read-only access to this user for API Review ticket readiness check automation validation'.format(AUTOMATION_USER, repo)
            add_gate_check_fail_comment(comment)
            return

        # Check if PR is Open if not put a comment that it is closed
        if pr.json()['state'] != 'OPEN':
            comment = 'PR link: {} is not in *Open* state. The PR should be in *Open* state for the API review process.'.format(pr_link)
            add_gate_check_fail_comment(comment)

        # Check if PR can be merged
        merge_status = get_pullrequest_merge_status(project, repo, pr_id)
        if not merge_status or not merge_status.json() or hasattr(merge_status.json(), 'errors') or merge_status.json().get('conflicted'):
            comment = 'PR link {} has a merge conflict or has been already merged. Please resolve the merge conflict, ' \
                      'update the PR and make sure that it is in *OPEN* state'.format(pr_link)
            add_gate_check_fail_comment(comment)

        # Check if PR has at-least one approval
        if not is_pr_approved(pr):
            comment = 'PR link {} has not been approved by at-least 1 member of your team. ' \
                      'The PR should be peer-approved by at-least 1 person from your team before API Review Board will review'.format(pr_link)
            add_gate_check_fail_comment(comment)

        # Check if valid diff is present
        pr_changes = get_pullrequest_changes(project, repo, pr_id)
        if not pr_changes and not pr_changes.json() and pr_changes.json()['values'] is not None:
            comment = 'No changes in found in the PR link attached to the ticket! Please update it with a valid PR link'
            add_gate_check_fail_comment(comment)
            return

        # Check if at-least one spec file is present
        spec_files = get_pr_change_nodes_for_pattern(pr_changes=pr_changes, pattern=API_SEPC_FILE_PATTERN)
        if spec_files:
            add_get_check_pass_comment('A spec file with file name pattern:{} was found!'.format(API_SEPC_FILE_PATTERN))
        else:
            comment = 'No spec file found matching the file name pattern {}. '.format(API_SEPC_FILE_PATTERN) + \
                      'Please make sure you have a spec file attached in the PR under review. ' + \
                      'Please request a manual bypass if no spec file can match this pattern in your PR'
            add_gate_check_fail_comment(comment)

        # Check if API spec is valid
        if is_spec_validation_output_valid(project, repo, pr_id, pr_changes):
            print('Spec validation is valid for this PR')
        else:
            print('Spec validation is Invalid for this PR')


def decorate_comment_for_jira(api_ticket, comment_list, gate_check_status, reporter=None):
    reporter = api_ticket.fields.reporter.name if reporter is None else reporter
    reporter_greeting = 'Hi [~{}],\n'.format(reporter)
    if comment_list:
        comment_list[0] = '\n* ' + comment_list[0]
    if gate_check_status:
        return reporter_greeting + 'The API Review ticket readiness check has passed. ' \
            + 'Specific things checked are mentioned below, but this is fyi only and does not require action from your side:-\n' \
            + '\n* '.join(comment_list) \
            + '\n\nThe ticket is ready for assignment to an API Review Board member.'
    else:
        return reporter_greeting \
            + 'The API Review ticket readiness check has failed! ' \
            + 'Transitioning the ticket status to *More information Needed* and re-assigning it back to you. ' \
            + '\n\nPlease Fix the issues mentioned below, update the PR if needed and put the ticket back into ' \
            + '*Needs Triage* status so that the API Review ticket readiness check can run again.' \
            + '\n\nIssues found:-' \
            + '\n* '.join(comment_list)


def gate_check_passed(api_ticket):
    # Reassign ticket to reporter Default Triage Reviewer and API Review ticket readiness check pass comment
    print('DEFAULT_TICKET_TRIAGE_ASSIGNEE is {}'.format(DEFAULT_TICKET_TRIAGE_ASSIGNEE))
    user = util.JIRA_CLIENT().user(DEFAULT_TICKET_TRIAGE_ASSIGNEE)
    if user:
        print('Re-assigning ticket {} to {}'.format(api_ticket.key, user.displayName))
        labels = api_ticket.fields.labels
        print('Adding label {} to ticket {} with existing labels {}'.format(GATE_CHECK_PASS_LABEL, api_ticket.key, labels))
        global MANUAL_CHECK_REQUIRED
        if MANUAL_CHECK_REQUIRED:
            print('Adding label {} to ticket {}'.format(MANUAL_CHECK_LABEL, api_ticket.key))
        global GATE_CHECK_PASS_COMMENTS
        comment = decorate_comment_for_jira(api_ticket, GATE_CHECK_PASS_COMMENTS, True, user.name)
        print('\nComment for {}:\n\n{}\n\n'.format(api_ticket.key, comment))
        if not config.IS_DRY_RUN:
            util.JIRA_CLIENT().assign_issue(api_ticket.key, user.name)
            labels.append(GATE_CHECK_PASS_LABEL)
            if MANUAL_CHECK_REQUIRED:
                labels.append(MANUAL_CHECK_LABEL)
            api_ticket.update(fields={"labels": labels})
            util.add_jira_comment(api_ticket.key, comment)
    else:
        print('Error: User {} not found!'.format(DEFAULT_TICKET_TRIAGE_ASSIGNEE))


def gate_check_failed(api_ticket):
    issue = api_ticket.key
    # Comment on the ticket failed API Review ticket readiness check with list of errors
    global GATE_CHECK_FAIL_COMMENTS
    comment = decorate_comment_for_jira(api_ticket, GATE_CHECK_FAIL_COMMENTS, False)
    print('\nComment for {}:\n\n{}\n\n'.format(api_ticket.key, comment))
    if not config.IS_DRY_RUN:
        util.add_jira_comment(api_ticket.key, comment)

    # Transition issue to More Information Needed.
    print('Transitioning {} to More Information Needed'.format(issue))
    if not config.IS_DRY_RUN:
        util.transition_issue_overall_status(util.JIRA_CLIENT(), api_ticket, config.STATUS_MORE_INFORMATION_NEEDED)

    # Reassign ticket back to reporter
    print('Re-assigning {} to {}'.format(issue, api_ticket.fields.reporter.displayName))
    if not config.IS_DRY_RUN:
        util.JIRA_CLIENT().assign_issue(api_ticket.key, api_ticket.fields.reporter.name)


def gate_check(api_ticket):
    print('Clearing API Review ticket readiness check counters')
    init_gate_check()
    print('Running API Review ticket readiness check on ticket {}'.format(api_ticket.key))
    print("=" * 30)
    setup_bitbucket(None)
    # Check for UDX ticket in related issues
    udx_ticket_check(api_ticket)
    # Check for API SPEC PR
    api_spec_pr_check(api_ticket)
    global GATE_CHECK_PASSED
    if GATE_CHECK_PASSED:
        gate_check_passed(api_ticket)
    else:
        gate_check_failed(api_ticket)
    print('=' * 30)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='API Review Reminder')
    parser.add_argument('--dry-run',
                        default=False,
                        action='store_true',
                        help='Perform a dry-run')

    args = parser.parse_args()
    config.IS_DRY_RUN = args.dry_run

    if config.IS_DRY_RUN:
        print('Running in dry-run mode')

    # Find all tickets to API Review ticket readiness check
    tickets_in_triage = query_actionable_tickets()
    if len(tickets_in_triage) == 0:
        print('No actionable tickets found for Query: {}'.format(QUERY_TEMPLATE))
    else:
        # API Review ticket readiness check all tickets
        for ticket in tickets_in_triage:
            try:
                gate_check(ticket)
            except Exception as err:
                logger.error('Exception occurred while processing ticket:{} exception:{}'.format(ticket.key, err))
    print('API Review ticket readiness check Job finished!')
