import argparse
import textwrap
import sys
import traceback

import config
import util

from autogen_issue_advisor_shared import DEFAULT_JIRA_ISSUE_FIELDS, CUSTOM_JIRA_ISSUE_FIELDS


def query_all_issues(issues):
    query = 'project = {JIRA_PROJECT} AND issuetype = "{ISSUE_TYPE}"'.format(
        JIRA_PROJECT=config.JIRA_PROJECT,
        ISSUE_TYPE=config.PUBLIC_ISSUE_TYPE_NAME)

    if issues:
        query = query + " AND key in (" + ", ".join(issues) + ")"

    # We really do want the values from the public ticket, so don't use util.search_dexreq_issues
    all_issues = util.jira_search_issues(query, fields=DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS)

    for issue in all_issues:
        print('{} - {}'.format(issue.key, issue.fields.summary))

    return all_issues


REMOVED_FIELDS_TEMPLATE = """\
There is nothing you need to do in response to this message. It is for archival purposes only.

We are changing the schema of the '{public_issuetype}' issue type and are removing some fields, which will now be loaded from the referenced preview ticket.

The values for the fields will still be available via REST API, but to make it easier for users to understand what values were contained in this field, this comment will also display them in the JIRA web interface.
{fields}
"""


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Ticket advisor (preview and public).')
    parser.add_argument('--issue',
                        action='append',
                        help='By default, we query JIRA. This allows you to specify a DEXREQ issue to process instead: --issue DEXREQ-123')
    parser.add_argument('--dry-run',
                        default=False,
                        action='store_true',
                        help='Perform a dry-run')
    parser.add_argument('--force',
                        default=False,
                        action='store_true',
                        help='Force an update')
    parser.add_argument('--verbose',
                        default=False,
                        action='store_true',
                        help='Verbose logging')

    failed = []

    args = parser.parse_args()
    verbose = args.verbose
    config.IS_DRY_RUN = args.dry_run
    force = args.force

    issues = query_all_issues(args.issue)

    commented_issues = []
    for issue in issues:
        print(textwrap.dedent("""\
        ========================================
        Issue: {}""").format(issue.key))
        try:
            # Flush, so we make sure the output of the issue key is already visible
            # NOTE: This is to help debug for DEX-6382
            sys.stdout.flush()

            fields = ""
            for removed_field_id in config.CUSTOM_FIELD_IDS_READ_FROM_PREVIEW_TICKET:
                name = config.CUSTOM_FIELD_NAME_FOR_ID[removed_field_id]
                if hasattr(issue.fields, removed_field_id):
                    v = getattr(issue.fields, removed_field_id)
                    if v is not None:
                        value = "'" + v + "'"
                    else:
                        value = "<empty>"
                else:
                    value = "<empty>"
                fields = fields + "\n'{}': {}".format(name, value)

            text = REMOVED_FIELDS_TEMPLATE.format(
                public_issuetype=config.PUBLIC_ISSUE_TYPE_NAME,
                fields=fields)

            if text:
                # comment on issue
                util.add_jira_comment(
                    issue.key,
                    text,
                    comment_type=config.COMMENT_TYPE_INFO
                )
                commented_issues.append(issue.key)
        except Exception as error:
            exception_string = traceback.format_exc()
            print("Unexpected error: {}\n{}".format(type(error), exception_string))
            failed.append(issue.key)

    if config.IS_DRY_RUN:
        print("DRY-RUN: Would have left {} comment(s)".format(len(commented_issues)))
    else:
        print("Left {} comment(s)".format(len(commented_issues)))

    if commented_issues:
        print("Commented on the following issues:\n{}".format("\n".join(commented_issues)))

    if failed:
        print("The following issues failed:\n{}".format("\n".join(failed)))
        sys.exit(1)
