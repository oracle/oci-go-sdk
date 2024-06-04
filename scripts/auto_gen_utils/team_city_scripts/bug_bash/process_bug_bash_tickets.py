
import os
import sys
import argparse

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../..'))

from sdk_regions_updater.region_updater_utils import get_issues_with_special_regions_to_be_ignored, get_region_from_storekeeper # noqa: ignore=F402
import util  # noqa: ignore=F402
import config  # noqa: ignore=F402

ALL_TEST_PASS_PREFIX = 'All tests passed'
REGION_TYPES_TO_IGNORE = ['(ONSR)', 'Development']
USER_OCID_PREFIX = os.environ.get('USER_OCID_PREFIX') or 'User OCID -'
TENANCY_OCID_PREFIX = os.environ.get('TENANCY_OCID_PREFIX') or 'Tenancy OCID -'
ENV_FILE_LOCATION = os.environ.get('ENV_FILE_LOCATION') or r'/tmp/region_info.txt'
USER_AUTH_RESULT_FILE = os.environ.get('USER_AUTH_RESULT_FILE') or r'/tmp/user_auth_test.txt'
INSTANCE_AUTH_RESULT_FILE = os.environ.get('INSTANCE_AUTH_RESULT_FILE') or r'/tmp/instance_auth_test.txt'
FORMAT_BLOCK = os.environ.get('FORMAT_BLOCK') or '{noformat}'
OCI_SDK_KEY_URL = os.environ.get('OCI_SDK_KEY_URL') or 'https://bitbucket.oci.oraclecorp.com/projects/SDK/repos/oci-sdk-keys/browse/oci_sdk_user_key_public.pem'

ISSUE_ID = 'ISSUE_ID'
REGION_TO_TEST = 'REGION_TO_TEST'
REALM_DOMAIN = 'REALM_DOMAIN'
OCI_TENANCY_ID = 'OCI_TENANCY_ID'
OCI_USER_OCID = 'OCI_USER_OCID'
PRE_RUN = 'pre-run'
POST_RUN = 'post-run'
MAX_COMMENT_LENGTH = 32000
BUILD_COUNTER = os.environ.get('BUILD_COUNTER') or 'Latest Build ID'
TEAM_CITY_URL = os.environ.get('TEAM_CITY_URL') or 'https://teamcity.oci.oraclecorp.com/buildConfiguration/Sdk_SelfService_BugBash'
TEAM_CITY_LOG = 'Team city URL: {} Build-Counter: {}.'.format(TEAM_CITY_URL, BUILD_COUNTER)
USER_AUTH_TEST_ALT_COMMENT = 'API key authentication tests have passed. The result log was too big to be added as comment, please see Team City for detailed logs. '
INSTANCE_AUTH_TEST_ALT_COMMENT = 'Instance Principals Authentication Test has passed. The result log was too big to be added as comment, please see Team City for detailed logs. '
CLI_TEST_PREFIX = 'Testing Python-CLI'
CLI_ALT_COMMENT = 'The result log was too big to be added as comment, please see Team City for detailed logs for CLI result.'
CLI_TEST_SUFFIX = 'Please check teamcity logs for detailed result.'
DEFAULT_BUGBASH_AUTOMATION_OWNER = os.environ.get('DEFAULT_BUGBASH_AUTOMATION_OWNER') or 'anurggar'

def get_jira_test_comment(prefix, test_results):
    return '{}\n{}\n{}\n{}'.format(prefix, FORMAT_BLOCK, test_results, FORMAT_BLOCK)


def format_env_variable(name, value):
    return '{}={}{}'.format(name, value, os.linesep)


def get_issue_id_from_env_file():
    issue_id = ''
    if os.path.exists(ENV_FILE_LOCATION):
        with open(ENV_FILE_LOCATION, 'r') as env_file:
            for line in env_file:
                if ISSUE_ID in line:
                    issue_id = line.split('{}{}'.format(ISSUE_ID, "="), 1)[1].strip()
                    break
    else:
        print('ERROR: File {} not found'.format(ENV_FILE_LOCATION))
    return issue_id


def build_env_file(issue_id, region_id, user_ocid, tenancy_ocid, realm):
    if os.path.exists(ENV_FILE_LOCATION):
        print("File exists.. Removing old env file")
        os.remove(ENV_FILE_LOCATION)

    with open(ENV_FILE_LOCATION, 'w') as env_file:
        env_file.write(format_env_variable(ISSUE_ID, issue_id))
        env_file.write(format_env_variable(REGION_TO_TEST, region_id))
        env_file.write(format_env_variable(REALM_DOMAIN, realm))
        env_file.write(format_env_variable(OCI_USER_OCID, user_ocid))
        env_file.write(format_env_variable(OCI_TENANCY_ID, tenancy_ocid))


def get_user_ocid_tenancy_ocid_from_description(description):
    user_ocid = ""
    tenancy_ocid = ""
    if description:
        # Parse description to get bitbucket API PR Review link
        for line in description.splitlines():
            line = line.strip()
            if TENANCY_OCID_PREFIX in line:
                tenancy_ocid = line.split(TENANCY_OCID_PREFIX, 1)[1].strip()
            if USER_OCID_PREFIX in line:
                user_ocid = line.split(USER_OCID_PREFIX, 1)[1].strip()
    return user_ocid, tenancy_ocid


def setup_issue_for_bugbash_testing(issue):
    region_id = issue.raw['fields']['summary'].split()[-1]
    reporter = issue.fields.reporter.name
    assignee = issue.fields.assignee.name
    description = issue.fields.description
    fail_comment = 'Hi [~{}], [~{}]\n'.format(reporter, assignee)

    if region_id and description:
        user_ocid, tenancy_ocid = get_user_ocid_tenancy_ocid_from_description(description)
        region_info = get_region_from_storekeeper(region_id)
        realm = region_info['realmDomainComponent'] if region_info else None
        if user_ocid and tenancy_ocid and realm:
            build_env_file(issue.key, region_id, user_ocid, tenancy_ocid, realm)
            comment = 'Starting Bug Bash testing for region: {}'.format(region_id)
            print(comment)
            if not config.IS_DRY_RUN:
                util.add_jira_comment(issue.key, comment)
        else:
            fail_comment += "ERROR: Description is missing either {}:{} or {}:{} or {}:{}".format(USER_OCID_PREFIX, user_ocid, TENANCY_OCID_PREFIX, tenancy_ocid, REALM_DOMAIN, realm)
            print(fail_comment)
            if not config.IS_DRY_RUN:
                util.add_jira_comment(issue.key, fail_comment)
    else:
        fail_comment += 'ERROR: Region id: {} or description:{} is invalid. Please make sure to use full region identifier in the ticket title preceeded by a -'.format(region_id, description)
        print(fail_comment)
        if not config.IS_DRY_RUN:
            util.add_jira_comment(issue.key, fail_comment)

    if not config.IS_DRY_RUN:
        # Transition the issue to In-Progress so that the ticket is not Picked up again
        print('INFO: Transitioning issue {} to IN-PROGRESS.'.format(issue.key))
        util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_IN_PROGRESS)


def process_new_tickets_with_invalid_regions(issues_with_invalid_regions):
    if not issues_with_invalid_regions:
        print('No invalid bug bash tickets found')
        return
    for issue in issues_with_invalid_regions:
        region_id = issue.raw['fields']['summary'].split()[-1]
        contact = issue.fields.assignee.name if issue.fields.assignee else DEFAULT_BUGBASH_AUTOMATION_OWNER
        comment = '[~{}],\nRegion {} is invalid. Please check with Region build team to see what is the issue'.format(contact, region_id)
        print(comment)
        if not config.IS_DRY_RUN:
            util.add_jira_comment(issue.key, comment)
            util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_IN_REVIEW)


def process_new_tickets_with_special_regions(issues_to_ignore):
    if not issues_to_ignore:
        print('No new Special regions bug bash tickets found')
        return
    for issue in issues_to_ignore:
        region_id = issue.raw['fields']['summary'].split()[-1]
        comment = 'Region {} is a secret region that will not be tested in Bug Bash. Closing the issue directly.'.format(
            region_id)
        print(comment)
        if not config.IS_DRY_RUN:
            util.add_jira_comment(issue.key, comment)
            util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_CLOSED)


def post_bug_bash_processing():
    issue_id = get_issue_id_from_env_file()
    print('Found issue id:{} from env file:{}'.format(issue_id, ENV_FILE_LOCATION))
    if issue_id:
        issue = util.JIRA_CLIENT().issue(id=issue_id)
        if issue:
            print('Issue:{} exists'.format(issue.key))
            user_auth_comment, instance_auth_comment = '', ''
            sdk_user_auth_comment, cli_user_auth_comment = '', ''
            user_auth_test_passed, instance_auth_test_passed = False, False

            if os.path.exists(USER_AUTH_RESULT_FILE):
                with open(USER_AUTH_RESULT_FILE, 'r') as user_auth_file:
                    #user_auth_comment = get_jira_test_comment(prefix='API Key Authentication tests results:-', test_results=user_auth_file.read())
                    user_auth_comment = user_auth_file.read()
                    print('*' * 30, '\n', 'User Auth comment', '\n', '*' * 30)
                    print(user_auth_comment)
                    sdk_user_auth_comment, cli_user_auth_comment = user_auth_comment.split(CLI_TEST_PREFIX, 1)
                    sdk_user_auth_comment = get_jira_test_comment(prefix='API Key Authentication tests results:-',
                                                                  test_results=sdk_user_auth_comment)
                    add_result_log_comment(issue_key=issue_id, comment=sdk_user_auth_comment,
                                           alt_comment=USER_AUTH_TEST_ALT_COMMENT)
                    # CLI test result
                    cli_user_auth_comment = get_jira_test_comment(prefix=CLI_TEST_PREFIX,
                                                                  test_results=cli_user_auth_comment)
                    add_result_log_comment(issue_key=issue_id,
                                           comment=cli_user_auth_comment + CLI_TEST_SUFFIX + TEAM_CITY_LOG,
                                           alt_comment=CLI_ALT_COMMENT)

            else:
                print('WARN: User auth result file not found at {}'.format(USER_AUTH_RESULT_FILE))

            if os.path.exists(INSTANCE_AUTH_RESULT_FILE):
                with open(INSTANCE_AUTH_RESULT_FILE, 'r') as instance_auth_file:
                    instance_auth_comment = get_jira_test_comment(prefix='Instance Principals Authentication Test Results:-',
                                                                  test_results=instance_auth_file.read())
                    print('*' * 30, '\n', 'Instance Auth comment', '\n', '*' * 30)
                    print(instance_auth_comment)
                    add_result_log_comment(issue_key=issue_id, comment=instance_auth_comment, alt_comment=INSTANCE_AUTH_TEST_ALT_COMMENT)
            else:
                print('WARN: Instance auth result file not found at {}'.format(INSTANCE_AUTH_RESULT_FILE))

            user_auth_test_passed = ALL_TEST_PASS_PREFIX in user_auth_comment
            region_id = issue.raw['fields']['summary'].split()[-1]
            instance_auth_test_passed = region_id in instance_auth_comment

            if user_auth_test_passed and instance_auth_test_passed:
                if not config.IS_DRY_RUN:
                    util.add_jira_comment(issue_key=issue_id, comment=ALL_TEST_PASS_PREFIX)
                    util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_CLOSED)
            else:
                fail_comment = ''
                if not user_auth_test_passed:
                    fail_comment += '* API key authentication tests have failed! Make sure the following details are correct:-\n' \
                        + '** The region identifier in the summary is correct\n' \
                        + '** The Tenancy in {} is correct and matches the value for the user mentioned in the description\n'.format(TENANCY_OCID_PREFIX) \
                        + '** The User in {} is correct and matches the USER OCID value for the user mentioned in the description\n'.format(USER_OCID_PREFIX) \
                        + '** Check if the User has the OCI SDK public key listed under them. If not please add the key under this user from here: {}\n'.format(OCI_SDK_KEY_URL)
                if not instance_auth_test_passed:
                    fail_comment += '* Instance Principals Authentication Test Failed! Please make sure that the Regions Build team has added capacity in this tenancy to create new Instances! \n'

                add_fail_comment(issue, fail_comment)
        else:
            print('ERROR: issue-id:{} not found by JIRA client'.format(issue_id))


def add_result_log_comment(issue_key, comment, alt_comment):
    if not config.IS_DRY_RUN:
        if len(comment) >= MAX_COMMENT_LENGTH:
            util.add_jira_comment(issue_key=issue_key, comment=alt_comment + TEAM_CITY_LOG)
        else:
            util.add_jira_comment(issue_key=issue_key, comment=comment)


def add_fail_comment(issue, fail_comment):
    reporter = issue.fields.reporter.name
    assignee = issue.fields.assignee.name
    comment = 'Hi [~{}], [~{}]\nThe following tests have failed:-\n'.format(reporter, assignee)
    comment += fail_comment
    comment += '\nSee Team City for more details on this failure. ' + TEAM_CITY_LOG
    print(comment)
    if not config.IS_DRY_RUN:
        util.add_jira_comment(issue_key=issue.key, comment=comment)


def process_new_tickets():
    issues = util.get_unprocessed_bug_bash_tickets()
    if not issues:
        print('No new unprocessed bug bash tickets found')
        return
    issues_to_ignore, issues_with_invalid_regions = get_issues_with_special_regions_to_be_ignored(issues, REGION_TYPES_TO_IGNORE)
    process_new_tickets_with_special_regions(issues_to_ignore)
    process_new_tickets_with_invalid_regions(issues_with_invalid_regions)
    for issue in [i for i in issues if i not in issues_to_ignore and i not in issues_with_invalid_regions]:
        setup_issue_for_bugbash_testing(issue)
        # TODO find a way to run this for multiple tickets
        break


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='API Review Reminder')
    parser.add_argument('--dry-run',
                        default=False,
                        action='store_true',
                        help='Perform a dry-run')
    parser.add_argument('--run-type',
                        required=True,
                        choices=[PRE_RUN, POST_RUN],
                        help='Run type for processing bug bash tickets, {}: before processing ticket, {}: for closing the ticket'.format(PRE_RUN, POST_RUN))

    args = parser.parse_args()
    config.IS_DRY_RUN = args.dry_run
    run_type = args.run_type

    if config.IS_DRY_RUN:
        print('Running in dry-run mode')

    if run_type == PRE_RUN:
        # Find new tickets to process and do bug bash testing
        process_new_tickets()
    elif run_type == POST_RUN:
        post_bug_bash_processing()
