import util
import argparse
import config
from jira import JIRAError

if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description='change go DEXREQ ticket from release requested to release approved before BULK PUBLIC.')
    parser.add_argument('--list-of-go-tickets',
                        action='append',
                        help='After the Go-No go meeting, provide a comma seperated list (No space) of all go tickets while running bulk-public: --list-of-go-tickets DEXREQ-123,DEXREQ-124,DEXREX-125')
    parser.add_argument('--dry-run', default=False, action='store_true', help='Dry-run, do not actually transition issues')
    args = parser.parse_args()
    list_of_go_tickets = args.list_of_go_tickets
    dry_run = args.dry_run
    try:
        if list_of_go_tickets:
            for issue in list_of_go_tickets[0].split(','):
                jira_issue = util.get_dexreq_issue(issue)
                if jira_issue.fields.status.name != config.STATUS_RELEASE_REQUESTED:
                    print('{} is not in Release/Processing requested status'.format(jira_issue))
                if not dry_run:
                    print("Changing status to Release Approved of {}".format(jira_issue))
                    # set all issues from release requested to release approved
                    util.transition_issue_overall_status(util.JIRA_CLIENT(), jira_issue, config.STATUS_RELEASE_APPROVED)
                else:
                    print("DRY-RUN: Would have changed status to Release Approved of {}".format(jira_issue))
                    # set all issues from release requested to release approved
                    util.transition_issue_overall_status(util.JIRA_CLIENT(), jira_issue, config.STATUS_RELEASE_APPROVED)

    except JIRAError as e:
        print('{} {}'.format(e.status_code, e.text))
        raise
