import util
import config
import argparse
import sys
import os
import re

from autogen_issue_advisor_shared import printv
from create_cli_design_review_ticket import get_cli_design_ticket
from add_or_update_scripts.cli_add_or_update_spec import determine_pom_location  # noqa: ignore=F402
from shared.buildsvc_tc_compatibility import build_log_link


FIX_FAILED_TESTS_COMMENT = """Following tests were disabled/failed during the PythonCLI generation. {action}:
{{code:title=Disabled/Failed Tests}}
{failed_tests}
{{code}}
"""
ERROR_MESSAGE_TEMPLATE = """The job failed to determine failed tests during CLI generation. {exception}.

The full build log can be found {build_log_link}.

For build log and artifact access see the [self-service FAQ|https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=132774928#SDK/CLISelfServiceFrequentlyAskedQuestions-HowdoIgetaccesstologsandartifacts?].
"""

PUBLIC_TESTS_FIX_MESSAGE = "Please use the generated PythonCLI branch to fix the disabled tests"

PREVIEW_TESTS_FIX_MESSAGE = "Please open a PR against PythonCLI's preview branch to fix the disabled tests"

FAILED_TEST_PATH = "../python-cli/failed_tests.txt"

CHANGED_SERVICE = "env.CHANGED_SERVICE"
SERVICES_DIR = "services"
POM_LOCATION_PATTERN = "services/(.*)/pom.xml"


def add_cli_tests_failed_label(issue_key):
    if config.IS_DRY_RUN:
        print("DRY-RUN: Not adding label to {}".format(issue_key))
    else:
        issue = util.get_dexreq_issue(issue_key)
        printv("Adding CLI_FAILED_TESTS_LABEL label to: " + issue_key)
        issue.add_field_value('labels', config.CLI_FAILED_TESTS_LABEL)


def remove_cli_tests_failed_label(issue_key):
    if config.IS_DRY_RUN:
        return
    else:
        issue = util.get_dexreq_issue(issue_key)
        if config.CLI_FAILED_TESTS_LABEL in issue.fields.labels:
            issue.fields.labels.remove(config.CLI_FAILED_TESTS_LABEL)
            issue.update(fields={"labels": issue.fields.labels})


def check_for_json_test_failures(build_type):
    cli_repo = config.CLI_REPO
    head_commit = cli_repo.head.commit
    if build_type == config.BUILD_TYPE_INDIVIDUAL_PUBLIC:
        head_commit_origin = cli_repo.commit('origin/master')
    else:
        head_commit_origin = cli_repo.commit('origin/preview')
    diff = head_commit_origin.diff(head_commit, create_patch=True)

    for mdiff in diff.iter_change_type('M'):
        if 'tests/resources/json_ignore_command_list.txt' in mdiff.a_path:
            print("Json test failures found")
            return True

    print("No Json test failures found")
    return False


def revert_changes_to_json_ignore_command_list_file(build_type):
    if not config.IS_DRY_RUN:
        try:
            cli_repo = config.CLI_REPO
            if build_type == config.BUILD_TYPE_INDIVIDUAL_PUBLIC:
                cli_repo.git.checkout(['origin/master', 'tests/resources/json_ignore_command_list.txt'], force=True)
            else:
                cli_repo.git.checkout(['origin/preview', 'tests/resources/json_ignore_command_list.txt'], force=True)
            cli_repo.git.commit("--amend", "-m", cli_repo.head.commit.message, '--allow-empty')
            cli_repo.git.push('-u', 'origin', 'HEAD', '-f')
        except Exception as e:
            print("Unable to remove local changes to tests/resources/json_ignore_command_list.txt file")
            print(e)


def get_service_name_from_issue(dexreq_issue_ticket):
    jira_obj = util.get_dexreq_issue(dexreq_issue_ticket)
    spec_name = getattr(jira_obj.fields, config.CUSTOM_FIELD_ID_SPEC_FRIENDLY_NAME)
    artifact_id = getattr(jira_obj.fields, config.CUSTOM_FIELD_ID_ARTIFACT_ID)
    services_root_dir = os.path.join(config.CLI_REPO_RELATIVE_LOCATION, SERVICES_DIR)

    path = determine_pom_location(artifact_id, spec_name, services_root_dir)
    return get_service_name_from_path(path)


def get_service_name_from_path(file_path):
    result = re.search(POM_LOCATION_PATTERN, file_path)
    return result.group(1)


# This script will be used in PythonCLI DEXREQ pipeline to verify if any JSON skeleton tests(tests/test_json_skeleton_command_coverage.py) were disabled
# as part of CLI generation.
# For public ticket: If any new tests were disabled, CLI-ManualChangesRequired label is added to the issue.
# For preview ticket: Comments on Design Review ticket with the failed tests.
# Reverts the changes to tests/resources/json_ignore_command_list.txt file.
if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Checks for disabled JSON unit tests during the CLI generation')
    parser.add_argument('--build-id',
                        required=True,
                        help='The team city build id for the build that is running this script.  This is used to update the relevant JIRA tickets with links to the team city build')
    parser.add_argument('--tool', default=config.CLI_NAME)
    parser.add_argument('--build-type',
                        default=config.BUILD_TYPE_INDIVIDUAL_PUBLIC,
                        help='The build type to use')
    parser.add_argument('--dry-run',
                        default=False,
                        action='store_true',
                        help='Perform a dry-run')

    args = parser.parse_args()
    tool_name = args.tool
    build_id = args.build_id
    build_type = args.build_type
    config.IS_DRY_RUN = args.dry_run

    if build_type not in [config.BUILD_TYPE_INDIVIDUAL_PUBLIC, config.BUILD_TYPE_INDIVIDUAL_PREVIEW] or tool_name != config.CLI_NAME:
        print('Cannot find failed tests for :{}, {}'.format(build_type, tool_name))
        sys.exit(0)

    generation_pass, build_pass = util.were_steps_successful(tool_name)
    if not (generation_pass and build_pass):
        print('Generation or Build did not pass, not proceeding.')
        sys.exit(0)

    last_commit_message = util.get_last_commit_message(tool_name)
    issue_keys = util.parse_issue_keys_from_commit_message(last_commit_message)
    if len(issue_keys) != 1:
        print('More than one DEXReq issues found {}, exiting!'.format(', '.join(issue_keys)))
        sys.exit(0)

    dexreq_issue = issue_keys[0]
    failed_json_tests = check_for_json_test_failures(build_type)

    failed_tests = []
    if failed_json_tests:
        failed_tests = ['tests/test_json_skeleton_command_coverage.py']
        revert_changes_to_json_ignore_command_list_file(build_type)  # this is not really required for individual_preview builds.

    # Check for failed integration tests:
    if os.path.isfile(FAILED_TEST_PATH):

        with open(FAILED_TEST_PATH, "r") as file_handle:
            failed_integ_tests = file_handle.readlines()

        failed_integ_tests = [name.strip() for name in failed_integ_tests]

        if len(failed_integ_tests) > 0:
            print("Found failed integ tests:", failed_integ_tests)
            failed_tests.extend(failed_integ_tests)

    try:
        if failed_tests:
            if build_type == config.BUILD_TYPE_INDIVIDUAL_PUBLIC:
                util.add_jira_comment(dexreq_issue, FIX_FAILED_TESTS_COMMENT.format(action=PUBLIC_TESTS_FIX_MESSAGE, failed_tests="\n".join(failed_tests)), config.COMMENT_TYPE_ERROR)
                add_cli_tests_failed_label(dexreq_issue)
            else:
                print("Adding comment to Design Review ticket")
                dexreq_issue_ticket = util.get_dexreq_issue(dexreq_issue)
                design_review_issue = get_cli_design_ticket(dexreq_issue_ticket, False)
                spec_name = getattr(dexreq_issue_ticket.fields, config.CUSTOM_FIELD_ID_SPEC_FRIENDLY_NAME)

                if design_review_issue:
                    if spec_name:
                        modified_service = get_service_name_from_issue(issue_keys)
                        for test in failed_tests:
                            test_split = test.split('.')
                            if len(test_split) >= 2 and test_split[0] == 'services' and\
                                    test_split[1] != modified_service:
                                failed_tests.remove(test)
                    if failed_tests:
                        add_cli_tests_failed_label(design_review_issue.key)
                        util.add_jira_comment(design_review_issue, FIX_FAILED_TESTS_COMMENT.format(action=PREVIEW_TESTS_FIX_MESSAGE, failed_tests="\n".join(failed_tests)), config.COMMENT_TYPE_ERROR)
                        if design_review_issue.fields.status.name in [config.STATUS_DONE, config.STATUS_CLOSED]:
                            util.transition_issue_overall_status(util.JIRA_CLIENT(), design_review_issue, config.STATUS_NEEDS_TRIAGE_STATUS)

        else:
            if build_type == config.BUILD_TYPE_INDIVIDUAL_PUBLIC:
                remove_cli_tests_failed_label(dexreq_issue)
            else:
                print("Found non individual_public build_type: " + build_type)
                design_review_issue = get_cli_design_ticket(util.get_dexreq_issue(dexreq_issue), False)
                if design_review_issue:
                    print("Removing CLI-FailedTests from design_review_issue")
                    remove_cli_tests_failed_label(design_review_issue)
                    print("Removed CLI-FailedTests from design_review_issue")
            print("No tests failures found.")
    except Exception as e:
        issue = util.get_dexreq_issue(dexreq_issue)
        util.add_jira_comment(
            issue.key,
            ERROR_MESSAGE_TEMPLATE.format(
                exception=str(e),
                build_log_link=build_log_link(build_id)
            ),
            comment_type=config.COMMENT_TYPE_ERROR
        )
