import argparse

import util
import config

import re
import json
import sys

process_preview_jira_queue = __import__("1_process_preview_jira_queue")

DEXREQ_PREVIEW_ISSUES_QUERY = 'project = {JIRA_PROJECT} and labels = {PREVIEW_LABEL} and resolution = Unresolved and labels != NewPipelineTesting'.format(JIRA_PROJECT=config.JIRA_PROJECT, PREVIEW_LABEL=config.PREVIEW_TICKET_LABEL)
DEXREQ_PUBLIC_ISSUES_QUERY = 'project = {JIRA_PROJECT} and labels = {PUBLIC_LABEL} and resolution = Unresolved and labels != NewPipelineTesting'.format(JIRA_PROJECT=config.JIRA_PROJECT, PUBLIC_LABEL=config.PUBLIC_TICKET_LABEL)

OLD_CUSTOM_STATUS_MORE_INFORMATION_NEEDED = 'More Information Needed'  # -> changed to SDK Status = Failure + Overall Status = Service Team Failure Investigation
OLD_CUSTOM_STATUS_MANUAL_ATTENTION_REQUIRED = 'Manual Attention Required'  # -> changed to SDK Status = Failure + Overall Status = DEX Support Required
OLD_CUSTOM_STATUS_PENDING_MERGE = 'Pending Merge'  # -> changed to SDK Status = Success + (Overall Status = 'Service Team Review Required' in preview and Overall Status = 'Service Team Work Required' in public)

OLD_CUSTOM_FIELD_ID_JAVA_SDK_STATUS = 'customfield_12535'
OLD_CUSTOM_FIELD_ID_PYTHON_SDK_STATUS = 'customfield_12536'
OLD_CUSTOM_FIELD_ID_RUBY_SDK_STATUS = 'customfield_12537'
OLD_CUSTOM_FIELD_ID_GO_SDK_STATUS = 'customfield_13104'
OLD_CUSTOM_FIELD_ID_CLI_STATUS = 'customfield_12538'

OLD_SDK_CLI_STATUS_FIELDS = [
    OLD_CUSTOM_FIELD_ID_JAVA_SDK_STATUS,
    OLD_CUSTOM_FIELD_ID_PYTHON_SDK_STATUS,
    OLD_CUSTOM_FIELD_ID_RUBY_SDK_STATUS,
    OLD_CUSTOM_FIELD_ID_GO_SDK_STATUS,
    OLD_CUSTOM_FIELD_ID_CLI_STATUS,
    'labels',
    'status'
]

OLD_PREVIEW_SIGN_OFF_LABEL = 'PreviewReviewComplete'

OLD_STATUS_READY_FOR_WORK = 'Ready for Work'
OLD_STATUS_IN_PROGRESS = 'In Progress'

STATUS_BACKLOG = 'Backlog'

OLD_SDK_CLI_STATUS_FIELD_ID_TO_NEW_MAPPINGS = [
    (OLD_CUSTOM_FIELD_ID_JAVA_SDK_STATUS, config.CUSTOM_FIELD_ID_JAVA_SDK_STATUS, config.CUSTOM_FIELD_NAME_JAVA_SDK_STATUS),
    (OLD_CUSTOM_FIELD_ID_PYTHON_SDK_STATUS, config.CUSTOM_FIELD_ID_PYTHON_SDK_STATUS, config.CUSTOM_FIELD_NAME_PYTHON_SDK_STATUS),
    (OLD_CUSTOM_FIELD_ID_RUBY_SDK_STATUS, config.CUSTOM_FIELD_ID_RUBY_SDK_STATUS, config.CUSTOM_FIELD_NAME_RUBY_SDK_STATUS),
    (OLD_CUSTOM_FIELD_ID_GO_SDK_STATUS, config.CUSTOM_FIELD_ID_GO_SDK_STATUS, config.CUSTOM_FIELD_NAME_GO_SDK_STATUS),
    (OLD_CUSTOM_FIELD_ID_CLI_STATUS, config.CUSTOM_FIELD_ID_CLI_STATUS, config.CUSTOM_FIELD_NAME_CLI_STATUS)
]

OLD_SDK_CLI_STATUS_VALUES_TO_NEW = {
    config.CUSTOM_STATUS_TODO: config.CUSTOM_STATUS_TODO,
    OLD_CUSTOM_STATUS_MORE_INFORMATION_NEEDED: config.CUSTOM_STATUS_FAILURE,
    OLD_CUSTOM_STATUS_MANUAL_ATTENTION_REQUIRED: config.CUSTOM_STATUS_FAILURE,
    OLD_CUSTOM_STATUS_PENDING_MERGE: config.CUSTOM_STATUS_SUCCESS,
    config.CUSTOM_STATUS_DONE: config.CUSTOM_STATUS_DONE
}

ISSUE_TYPE_ENUM_PUBLIC = "Public"
ISSUE_TYPE_ENUM_PREVIEW = "Preview"

SCRIPT_MODE_PRE_PROCESS = 'PreProcess'
SCRIPT_MODE_UPDATE_ISSUES = 'UpdateIssues'


def process_old_sdk_status_fields_and_overall_status(issue, issue_type_enum, new_issue_state):
    all_fields_in_todo = True
    all_fields_in_pending_merge = True
    all_fields_in_pending_merge_or_done = True
    for old_field_id, new_field_id, new_field_name in OLD_SDK_CLI_STATUS_FIELD_ID_TO_NEW_MAPPINGS:
        old_value = getattr(issue.fields, old_field_id).value
        new_value = OLD_SDK_CLI_STATUS_VALUES_TO_NEW[old_value]

        if not old_value == config.CUSTOM_STATUS_TODO:
            all_fields_in_todo = False

        if not old_value == OLD_CUSTOM_STATUS_PENDING_MERGE:
            all_fields_in_pending_merge = False

            if not old_value == config.CUSTOM_STATUS_DONE:
                all_fields_in_pending_merge_or_done = False

        if old_value == OLD_CUSTOM_STATUS_MORE_INFORMATION_NEEDED:
            # if the issue has already been set to DEX Support Required, that takes precedence, we don't want to override to 
            # 'Service Team Failure Investigation'
            if not issue.fields.status == config.STATUS_DEX_SUPPORT_REQUIRED:
                new_issue_state['status'] = config.STATUS_SERVICE_TEAM_FAILURE_INVESTIGATION
        elif old_value == OLD_CUSTOM_STATUS_MANUAL_ATTENTION_REQUIRED:
            new_issue_state['status'] = config.STATUS_DEX_SUPPORT_REQUIRED

        new_issue_state['fields'][new_field_id] = {'value': new_value}

    if issue_type_enum == ISSUE_TYPE_ENUM_PREVIEW:
        if all_fields_in_todo:
            new_issue_state['status'] = config.STATUS_PROCESSING_REQUESTED
        elif all_fields_in_pending_merge:
            if issue.fields and issue.fields.labels and OLD_PREVIEW_SIGN_OFF_LABEL in issue.fields.labels:
                new_issue_state['status'] = config.STATUS_READY_FOR_PREVIEW
            else:
                # service team still needs to give explicit sign off
                new_issue_state['status'] = config.STATUS_SERVICE_TEAM_REVIEW_REQUIRED
    elif issue_type_enum == ISSUE_TYPE_ENUM_PUBLIC:
        if all_fields_in_todo and issue.fields and issue.fields.status and issue.fields.status.name == OLD_STATUS_READY_FOR_WORK:
            new_issue_state['status'] = config.STATUS_PROCESSING_REQUESTED
        elif (all_fields_in_pending_merge or all_fields_in_pending_merge_or_done) and (not new_issue_state.get('status') == config.STATUS_DEX_SUPPORT_REQUIRED):
            # DEX support Required takes precedence over any of the states that would be set based on PR progress
            process_pull_request_progress(issue, new_issue_state, 'master')


def process_udx_ticket_from_summary(issue, new_issue_state):
    summary = issue.fields.summary
    udx_ticket = re.findall(r'UDX-\d+', summary, re.IGNORECASE)

    if not udx_ticket or not len(udx_ticket) == 1:
        if issue.fields and issue.fields.labels and config.BYPASS_CHECK_UDX_TICKET_LABEL in issue.fields.labels:
            print('INFO: Allowing issue {} with no UDX ticket because it has bypass label'.format(issue.key))
        else:
            raise ValueError("Issue: {} did not contain exactly 1 UDX tickets in summary. Summary: {}".format(issue.key, issue.fields.summary))
    else:
        new_issue_state['fields'][config.CUSTOM_FIELD_ID_UDX_TICKET] = udx_ticket[0]


def process_pull_request_progress(issue, new_issue_state, target_branch_filter=None):
    pr_status = util.get_pr_status_for_tools(util.JIRA_CLIENT(), issue, config.TOOLS_HANDLED_BY_SERVICE_TEAMS + [config.TESTING_SERVICE_NAME], target_branch_filter=target_branch_filter)
    print('pr status: ' + json.dumps(pr_status, indent=2))
    if pr_status['all_prs_merged']:
        new_issue_state['status'] = config.STATUS_TO_DEPLOY
    elif pr_status['all_prs_initiated']:
        new_issue_state['status'] = config.STATUS_RELEASE_REQUESTED


def calculate_new_issue_states(issues, output_file):
    # grab all open DEXREQ tickets
    preview_issues = util.jira_search_issues(process_preview_jira_queue.limit_query_to_issue_keys(DEXREQ_PREVIEW_ISSUES_QUERY, issues, False), fields=', '.join(
        process_preview_jira_queue.DEFAULT_JIRA_ISSUE_FIELDS + process_preview_jira_queue.CUSTOM_JIRA_ISSUE_FIELDS_PER_BUILD_TYPE[config.BUILD_TYPE_INDIVIDUAL_PREVIEW] + OLD_SDK_CLI_STATUS_FIELDS)
    )
    public_issues = util.jira_search_issues(process_preview_jira_queue.limit_query_to_issue_keys(DEXREQ_PUBLIC_ISSUES_QUERY, issues, False), fields=', '.join(
        process_preview_jira_queue.DEFAULT_JIRA_ISSUE_FIELDS + process_preview_jira_queue.CUSTOM_JIRA_ISSUE_FIELDS_PER_BUILD_TYPE[config.BUILD_TYPE_INDIVIDUAL_PUBLIC] + OLD_SDK_CLI_STATUS_FIELDS)
    )

    print('Found total {} open DEXREQ tickets. Preview: {}, Public: {}.'.format(
        len(preview_issues) + len(public_issues),
        len(preview_issues),
        len(public_issues)
    ))

    new_issue_states = []

    errors = []

    print('Processing Preview Issue(s)...')
    print('=====================')
    for issue in preview_issues:
        print('{} - {}'.format(issue.key, issue.fields.summary))

        new_issue_state = {
            'status': STATUS_BACKLOG,
            'key': issue.key,
            'issuetype': config.PREVIEW_ISSUE_TYPE_NAME,
            'fields': {}
        }

        try:
            process_udx_ticket_from_summary(issue, new_issue_state)
            process_old_sdk_status_fields_and_overall_status(issue, ISSUE_TYPE_ENUM_PREVIEW, new_issue_state)

            new_issue_states.append(new_issue_state)
        except Exception as e:
            errors.append('Could not process ticket {}. Error: {}'.format(issue.key, str(e)))

    print('\n')

    print('Public Issue(s):')
    print('=====================')
    for issue in public_issues:
        print('{} - {}'.format(issue.key, issue.fields.summary))

        new_issue_state = {
            'status': STATUS_BACKLOG,
            'key': issue.key,
            'issuetype': config.PUBLIC_ISSUE_TYPE_NAME,
            'fields': {}
        }

        try:
            process_udx_ticket_from_summary(issue, new_issue_state)
            process_old_sdk_status_fields_and_overall_status(issue, ISSUE_TYPE_ENUM_PUBLIC, new_issue_state)

            new_issue_states.append(new_issue_state)
        except Exception as e:
            errors.append('Could not process ticket {}. Error: {}'.format(issue.key, str(e)))

    print('\n')

    with open(output_file, 'w+') as f:
        f.write(json.dumps(new_issue_states, indent=2))

    if errors:
        print('ERROR: The following errors occurred:')
        for error in errors:
            print('ERROR: {}'.format(error))

    print("""Manually update all tickets referenced above to their appropriate new ticket states (Preview or Public).
Then re-run this script with mode == {} and --input-file {}""".format(
        SCRIPT_MODE_UPDATE_ISSUES,
        output_file)
    )

    print('Query for bulk migrating PREVIEW issues:')
    print('\t' + ' OR '.join(['key = {}'.format(issue.key) for issue in preview_issues]))

    print('Query for bulk migrating PUBLIC issues:')
    print('\t' + ' OR '.join(['key = {}'.format(issue.key) for issue in public_issues]))


def update_issues(issues_filter, input_file):
    new_issue_states = []
    with open(input_file, 'r') as f:
        content = f.read()
        new_issue_states = json.loads(content)
    
    for issue_state in new_issue_states:
        if issues_filter and (not issue_state['key'] in issues_filter):
            print('Skipping issue: {} because it is not in issues filter'.format(issue_state['key']))
            continue

        issue = util.JIRA_CLIENT().issue(issue_state['key'])

        if issue_state['issuetype'] != issue.fields.issuetype.name:
            print('Skipping issue {} because it has not yet been converted to issuetype: {}'.format(issue.key, issue_state['issuetype']))
            continue

        util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, issue_state['status'])

        if config.IS_DRY_RUN:
            print('DRY-RUN: Not updating the following fields in issue {}: {}'.format(issue.key, json.dumps(issue_state['fields'])))
        else:
            print('Updating the following fields in issue {}: {}'.format(issue.key, json.dumps(issue_state['fields'])))
            issue.update(fields=issue_state['fields'])


# Known limitations:
# - The script will not detect if bulk preview PRs are out for preview, so don't do the migraiton while the weekly bulk preview is in progress
# - The script does not attempt to handle converting into the new 'Processing' and 'Processing - Bulk' states, so don't do the migration while the current 'Process queue' job is running
if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Process the DEXREQ JIRA queue (preview and public).')
    parser.add_argument('--issue',
                        action='append',
                        help='By default, we query JIRA. This allows you to specify a DEXREQ issue to process instead: --issue DEXREQ-123')
    parser.add_argument('--dry-run',
                        default=False,
                        action='store_true',
                        help='Perform a dry-run')
    parser.add_argument('--mode',
                        default=SCRIPT_MODE_PRE_PROCESS,
                        help=""""The mode to run the script in. {} will determine and output 
                        the new desired states for each issue, {} will take in the output of a PreProcess job and update issues
                        accordingly""".format(SCRIPT_MODE_PRE_PROCESS, SCRIPT_MODE_UPDATE_ISSUES))
    parser.add_argument('--input-file',
                        help=""""The input file to read the desired issue states from. This option is only applicable for the mode: {}""".format(SCRIPT_MODE_UPDATE_ISSUES))
    parser.add_argument('--output-file',
                        help=""""The output file to write the desired issue states to. This option is only applicable for the mode: {}""".format(SCRIPT_MODE_PRE_PROCESS))

    args = parser.parse_args()
    issues = args.issue
    output_file = args.output_file
    input_file = args.input_file
    mode = args.mode
    config.IS_DRY_RUN = args.dry_run

    if mode == SCRIPT_MODE_PRE_PROCESS:
        if not output_file:
            sys.exit('--output-file is required for mode == {}'.format(SCRIPT_MODE_PRE_PROCESS))

        calculate_new_issue_states(issues, output_file)
    elif mode == SCRIPT_MODE_UPDATE_ISSUES:
        if not input_file:
            sys.exit('--input-file is required for mode == {}'.format(SCRIPT_MODE_UPDATE_ISSUES))

        update_issues(issues, input_file)
    else:
        print('Invalid script mode')
