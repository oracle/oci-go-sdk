import os
import sys

import constants
import git_util
import jira_util


def execute_local_changes():
    issue_key = sys.argv[1]  # Design review ticket
    try:
        manual_change = sys.argv[2]  # Manual Change Command
    except Exception:
        print("manual change not supplied fetch from jira ticket")
        issue = jira_util.get_jira_issue(issue_key)
        manual_change = jira_util.check_comments_for_manual_change(issue)
        print("retrieved manual change is ", manual_change)

    self_service_branch = git_util.checkout_self_service_branch(git_util.PYTHON_CLI_WORKING_DIRECTORY, 'preview',
                                                                issue_key)
    rtv = 0
    if jira_util.RENAME_ROOT_GROUP in manual_change.upper():
        print("rename root command")
        rtv = os.system(constants.MANUAL_CHANGE_SCRIPT.format(operation=constants.RENAME_ROOT_GROUP,
                                                              issue=issue_key,
                                                              branch=self_service_branch))
        if rtv != 0:
            print("root command changes have completed please check")
            exit()

    for supported_change in jira_util.SUPPORTED_MANUAL_CHANGES:
        if supported_change in manual_change.upper():
            try:
                rtv = jira_util.execute_manual_change(manual_change)
            except Exception as e:
                print(e)
                # handle all raised exceptions
                jira_util.handle_exceptions(issue_key)
                # aggregate list
                jira_util.add_jira_comment(issue_key,
                                           jira_util.MANUAL_CHANGE_FAILED_TEMPLATE.format(e, self_service_branch,
                                                                                          self_service_branch.replace(
                                                                                              '-', '\-')))  # noqa: W605
        break
    if rtv == 0:
        print("supported command changes have completed please check")
        exit()


if __name__ == '__main__':
    execute_local_changes()
