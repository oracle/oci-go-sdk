#!/usr/bin/env python3
# Copyright (c) 2016, 2020, Oracle and/or its affiliates.  All rights reserved.
# This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

import os

import python_cli.git_util as git
import python_cli.jira_config as jira_config
import python_cli.jira_util as jira_util
import python_cli.self_service_manual_change as self_service
from python_cli.constants import RENAME_ROOT_GROUP, EXECUTE_MANUAL_CHANGES, MANUAL_CHANGE_SCRIPT, CLI_REINSTALL_SCRIPT

design_reviews = jira_util.get_open_design_review_tickets()

for design_review in design_reviews:
    issue = jira_util.get_jira_issue(design_review.key)
    manual_change = jira_util.check_comments_for_manual_change(issue)
    if manual_change is not None:
        branch = jira_util.get_branch_with_changes(issue)
        if branch is None:
            continue
        self_service_branch = git.checkout_self_service_branch(git.PYTHON_CLI_WORKING_DIRECTORY, branch,
                                                               design_review.key)
        self_service.run_command(git.PYTHON_CLI_WORKING_DIRECTORY, CLI_REINSTALL_SCRIPT)
        rtv = 0
        if jira_util.RENAME_ROOT_GROUP in manual_change.upper():
            rtv = os.system(MANUAL_CHANGE_SCRIPT.format(operation=RENAME_ROOT_GROUP,
                                                        issue=design_review.key,
                                                        branch=self_service_branch))
        if rtv != 0:
            git.push_all_generated_changes_to_remote(git.PYTHON_CLI_WORKING_DIRECTORY, design_review.key)
            continue
        for supported_change in jira_util.SUPPORTED_MANUAL_CHANGES:

            if supported_change in manual_change.upper():
                rtv = os.system(MANUAL_CHANGE_SCRIPT.format(operation=EXECUTE_MANUAL_CHANGES,
                                                            issue=design_review.key,
                                                            branch=self_service_branch))
                break

        if rtv == 0:
            jira_util.add_jira_comment(design_review.key,
                                       jira_util.MANUAL_CHANGE_COMPLETED_TEMPLATE.format(self_service_branch,
                                                                                         self_service_branch.replace(
                                                                                             '-', '\-')))  # noqa: W605
        git.push_all_generated_changes_to_remote(git.PYTHON_CLI_WORKING_DIRECTORY, design_review.key)
        issue.add_field_value('labels', jira_config.CLI_SELF_SERVICE_LABEL)
        issue.add_field_value('labels', jira_config.CLI_MANUAL_CHANGES_LABEL)
        issue.update()
