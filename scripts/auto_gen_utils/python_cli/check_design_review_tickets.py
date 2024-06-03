#!/usr/bin/env python3
# Copyright (c) 2016, 2020, Oracle and/or its affiliates.  All rights reserved.
# This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

import util
import python_cli.jira_util as jira_util
import python_cli.constants as constants

design_reviews = jira_util.get_open_design_review_tickets()
for design_review in design_reviews:
    issue = jira_util.get_jira_issue(design_review.key)
    if not hasattr(issue.fields, 'assignee') or not hasattr(issue.fields.assignee, 'key'):
        continue
    
    if jira_util.is_new_design_review(issue):
        jira_util.add_jira_comment(issue, jira_util.NEW_DESIGN_REVIEW_TEMPLATE.format(user=issue.fields.assignee.key))
    
    if jira_util.no_changes_required(issue, constants.CLI_TEAM_MEMBERS):
        jira = jira_util.JIRA_CLIENT()
        util.transition_issue_overall_status(jira, issue, 'Done')
        
        issue = jira_util.get_jira_issue(design_review.key)
        if issue.fields.status.name == 'Done':
            jira_util.add_jira_comment(issue, jira_util.NO_MANUAL_CHANGES_TICKET_CLOSED.format(user=issue.fields.reporter.key))
        
    if jira_util.not_updated_in_seven_days(issue):
        udx_ticket = jira_util.get_udx_ticket_from_design_review(issue)
        if udx_ticket:
            public_ticket = jira_util.get_public_ticket_from_udx(udx_ticket)
            if public_ticket:
                days = jira_util.is_ga_date_within_month(issue)
                if days:
                    jira_util.add_jira_comment(issue, jira_util.UPCOMING_GA_DATE_TEMPLATE.format(user=issue.fields.assignee.key, days=days))
