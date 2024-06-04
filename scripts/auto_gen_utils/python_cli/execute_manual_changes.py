#!/usr/bin/env python3

import python_cli.jira_util as jira_util
import sys

issue_key = sys.argv[1]
self_service_branch = sys.argv[2]
issue = jira_util.get_jira_issue(issue_key)
manual_change = jira_util.check_comments_for_manual_change(issue)
try:
    jira_util.execute_manual_change(manual_change)

except Exception as e:
    print(e)
    # handle all raised exceptions
    jira_util.handle_exceptions(issue_key)
    # aggregate list
    jira_util.add_jira_comment(issue_key, jira_util.MANUAL_CHANGE_FAILED_TEMPLATE.format(e, self_service_branch,
                                                                                         self_service_branch.replace('-', '\-')))  # noqa: W605
    sys.exit(1)
