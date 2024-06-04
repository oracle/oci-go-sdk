import os
import sys
import argparse
from datetime import datetime

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../..'))

import util  # noqa: ignore=F402
import config  # noqa: ignore=F402

cutoff = os.environ.get('CUTOFF') or '1d'
REMINDER_TEMPLATE = """\
Hello {stakeholders},
This ticket has not been updated in the past {cutoff}. Can you let us know the latest status on this API Review?"""
QUERY_TEMPLATE = 'project = "Developer Experience" AND labels in (APIConsistency) AND (status = "In Progress" OR (status = "Ready for Work" AND assignee != aarthurs)) AND updated < -{cutoff} ORDER BY updated ASC'


def query_actionable_tickets():
    query = QUERY_TEMPLATE.format(cutoff=cutoff)
    actionable_tickets = util.jira_search_issues(query)

    print('Total of {} actionable tickets found matching query `{}`.'.format(len(actionable_tickets), query))
    for ticket in actionable_tickets:
        print('{} - `{}` was last updated on {}'.format(ticket.key, ticket.fields.summary, ticket.fields.updated))

    return actionable_tickets


def find_stakeholders(ticket):
    stakeholders = find_all_assignees(ticket.key)
    stakeholders.add(ticket.fields.creator.name)
    return stakeholders


def post_reminder(ticket):
    print('Posting reminder on ticket {}:'.format(ticket.key))
    stakeholders = find_stakeholders(ticket)
    reminder = REMINDER_TEMPLATE.format(stakeholders=decorate(stakeholders), cutoff=cutoff)
    print(reminder)
    if not config.IS_DRY_RUN:
        util.add_jira_comment(ticket.key, reminder)


def find_all_assignees(key):
    assignees = set()
    issue = util.JIRA_CLIENT().issue(key, expand='changelog')
    changelog = issue.changelog
    for history in changelog.histories:
        for item in history.items:
            if item.field == 'assignee':
                assignees.add(getattr(item, 'from'))
                assignees.add(item.to)
    assignees.discard(None)
    return assignees


def decorate(names):
    return ', '.join('[~{}]'.format(name) for name in names)


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
    print('Current time is {}'.format(datetime.now()))
    print('Using cutoff {}'.format(cutoff))

    actionable_tickets = query_actionable_tickets()
    for ticket in actionable_tickets:
        post_reminder(ticket)
