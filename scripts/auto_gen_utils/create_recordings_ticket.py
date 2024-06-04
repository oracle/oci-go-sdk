import util
import config
import os
import sys
import argparse
from shared.buildsvc_tc_compatibility import build_log_link

FAILED_TEST_PATH = "../python-cli/failed_tests.txt"

CLI_RECORDINGS_DESCRIPTION_TEMPLATE = """
The tests posted in the comments were disabled while generating the CLI due to compatibility issues with the newly generated code.
Please re-test and re-create the recordings for the skipped tests, and include them as part of a PR for the CLI preview branch.
Include a link to this ticket as part of your pull request comments. The Jira ticket may be closed once the corresponding PR has been merged.
"""

ADDED_RECORDINGS_TEMPLATE = """
{{code:title=Disabled Tests}}
{tests}
{{code}}
"""
CLI_RECORDINGS_FAILURE_TEMPLATE = """The job failed to create/update CLI Recordings ticket. {exception}.

The full build log can be found {build_log_link}.

For build log and artifact access see the [self-service FAQ|https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=132774928#SDK/CLISelfServiceFrequentlyAskedQuestions-HowdoIgetaccesstologsandartifacts?].
"""

TESTS_FAILED_COMMENT = """Following tests have been failed while generating the CLI. Please use the generated branch to fix the failed tests.
{{code:title=Failed Tests}}
{tests}
{{code}}
"""

CLI_RECORDINGS_SUCCESS_MESSAGE_TEMPLATE = 'The following CLI Create Test Recordings ticket has been opened and needs your attention -- https://jira.oci.oraclecorp.com/browse/{}'


def is_issue_summary_matches_cli_recordings_issue(issue_summary):
    return issue_summary and 'CLI Create Test Recordings'.lower() in issue_summary.lower()


def create_recordings_ticket(dexreq_issue):
    print('Creating new CLI Recordings Issue')
    sprint_id = util.find_dex_tools_active_sprint_id()

    fields = {
        'project': 'DEX',
        'summary': 'CLI Create Test Recordings for {}'.format(dexreq_issue.key),
        'description': CLI_RECORDINGS_DESCRIPTION_TEMPLATE,
        'issuetype': {'name': 'Task'},
        'components': [{'name': 'CLI'}],
        'labels': ['CLITestRecordings'],
        'assignee': {'name': dexreq_issue.fields.creator.name},
        config.CUSTOM_FIELD_ID_SPRINT: sprint_id
    }

    recordings_ticket = util.JIRA_CLIENT().create_issue(fields)
    util.JIRA_CLIENT().create_issue_link(config.IS_ACTION_ITEM_OF_LINK, recordings_ticket, dexreq_issue)

    return recordings_ticket


def get_cli_test_recordings_ticket(dexreq_issue):
    if hasattr(dexreq_issue.fields, 'issuelinks'):
        for link in dexreq_issue.fields.issuelinks:
            if hasattr(link, 'outwardIssue'):
                issue = util.JIRA_CLIENT().issue(link.outwardIssue.key, fields='description, summary')
                if is_issue_summary_matches_cli_recordings_issue(issue.fields.summary):
                    print('Found CLI Recordings ticket: {}'.format(issue.key))
                    return issue

    return create_recordings_ticket(dexreq_issue)


# Comments on the DEXREQ ticekt and halts the DEXREQ pipeline.
def handle_non_generated_tests_failures(tests, dexreq_issue):
    tests_str = "".join(tests)
    util.add_jira_comment(dexreq_issue, TESTS_FAILED_COMMENT.format(tests=tests_str), config.COMMENT_TYPE_ERROR)


# Creates/Updates a CLI Test recordings ticket and continues the DEXREQ pipeline.
def handle_generated_tests_failures(tests, dexreq_issue):
    tests_str = "".join(tests)
    if tests_str:
        recordings_ticket = get_cli_test_recordings_ticket(dexreq_issue)
        util.add_jira_comment(recordings_ticket, ADDED_RECORDINGS_TEMPLATE.format(tests=tests_str))
        util.add_jira_comment(dexreq_issue, CLI_RECORDINGS_SUCCESS_MESSAGE_TEMPLATE.format(recordings_ticket.key))


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='CLI Design review issue creation!')
    parser.add_argument('--build-id',
                        required=True,
                        help='The team city build id for the build that is running this script.  This is used to update the relevant JIRA tickets with links to the team city build')
    parser.add_argument('--tool',
                        default=config.CLI_NAME,
                        help='The tool for which to generate the preview. Accepted values: {}'.format(', '.join(config.TOOL_NAMES)))
    parser.add_argument('--build-type',
                        default=config.BUILD_TYPE_INDIVIDUAL_PREVIEW,
                        help='The build type to use, can be one of the following: {}'.format(', '.join(config.VALID_BUILD_TYPES)))
    parser.add_argument('--dry-run',
                        default=False,
                        action='store_true',
                        help='Perform a dry-run')
    parser.add_argument('--create-jira',
                        default=False,
                        action='store_true',
                        help='Opens a JIRA ticket with the failed tests and links it to DEXREQ ticket')

    args = parser.parse_args()
    tool_name = args.tool
    build_id = args.build_id
    build_type = args.build_type
    config.IS_DRY_RUN = args.dry_run

    if build_type != config.BUILD_TYPE_INDIVIDUAL_PREVIEW or tool_name != config.CLI_NAME:
        print('CLI Create Test Recordings issue generation not required for build type: {}'.format(build_type))
        sys.exit(0)

    last_commit_message = util.get_last_commit_message(tool_name)
    issue_keys = util.parse_issue_keys_from_commit_message(last_commit_message)
    if len(issue_keys) != 1:
        print('More than one DEXReq issues found {}. Not generating CLI Test Recordings ticket', issue_keys)
        sys.exit(0)
    dexreq_issue_key = issue_keys[0]

    try:
        if not os.path.isfile(FAILED_TEST_PATH):
            print("No test failures found, exiting!")
            sys.exit(0)

        with open(FAILED_TEST_PATH, "r") as file_handle:
            failed_tests = list(file_handle)
            dexreq_issue = util.get_dexreq_issue(dexreq_issue_key)
            # filter unit/integ/extended test failures.
            not_generated_tests = list(filter((lambda test: '/tests/generated/' not in test), failed_tests))
            if not_generated_tests and len(not_generated_tests) > 0:
                handle_non_generated_tests_failures(not_generated_tests, dexreq_issue)

            if args.create_jira:
                generated_tests = list(filter(lambda test: '/tests/generated/' in test, failed_tests))
                if generated_tests and len(generated_tests) > 0:
                    handle_generated_tests_failures(generated_tests, dexreq_issue)

    except Exception as e:
        issue = util.get_dexreq_issue(dexreq_issue_key)
        util.add_jira_comment(
            issue.key,
            CLI_RECORDINGS_FAILURE_TEMPLATE.format(
                exception=str(e),
                build_log_link=build_log_link(build_id)
            ),
            comment_type=config.COMMENT_TYPE_ERROR
        )
