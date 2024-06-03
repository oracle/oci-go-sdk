from jira import JIRA, JIRAError
import os


class JiraTicketReporter:
    JIRA_SERVER = 'https://jira.oci.oraclecorp.com'

    JIRA_OPTIONS = {
        'server': JIRA_SERVER,
        'verify': True
    }

    TICKET_SEARCH_QUERY = """
        project={project} and (Description ~ 'IssueRoutingInfo tag {tag}' or worklogComment ~ 'IssueRoutingInfo tag {tag}') and resolution = Unresolved and status not in (\"done\", \"closed\", \"duplicate\", \"won't fix\", \"resolved\") and labels in (\"{label}\")
    """

    JIRA_TICKET_LABEL = "SDK-Nightly-Run-Failure"

    TICKET_UPDATE_DESCRIPTION = """
    
    Code coverage failure:{codecoverage_text}
    """

    CODE_COVERAGE_TEXT = """
    {package}
    Expected code coverage: {baseline_codecoverage}, Actual code coverage: {actual_codecoverage}
    Complete code coverage report can be found at: {report_url}
    """

    TICKET_CREATE_DESCRIPTION = """
    This ticket was opened because the nightly run failed to meet the code coverage criteria.{codecoverage_text}

    """

    TICKET_CREATE_SUMMARY = """
    Java SDK Nightly Build Error
    """

    FAILURE_GENERAL_COMMENT = """
        There were failures. Please check worklog for details.

    """

    FAILURE_CREATE_COMMENT = """
        There were failures. Please check worklog for details.

        Please do not delete this line: {{ IssueRoutingInfo tag {tag}}}
    """

    jira_client = None

    def __init__(self):
        username = os.environ.get('JIRA_USERNAME')
        password = os.environ.get('JIRA_PASSWORD')
        self.jira_client = JIRA(self.JIRA_OPTIONS, basic_auth=(username, password))

    def report_codecov_to_jira_ticket(self, project, tag, package_label, baseline_codecoverage, actual_codecoverage, report_url, dry_run):
        codecoverage_text = self.CODE_COVERAGE_TEXT.format(package=package_label, baseline_codecoverage=baseline_codecoverage, actual_codecoverage=actual_codecoverage, report_url=report_url)
        try:
            query = self.TICKET_SEARCH_QUERY.format(project=project, tag=tag, label=self.JIRA_TICKET_LABEL)
            print(query)
            issues = self.jira_client.search_issues(query, fields=["summary","status","assignee","key"]) 
            if not issues:
                desc = self.TICKET_CREATE_DESCRIPTION.format(codecoverage_text=codecoverage_text)
                if not dry_run:
                    print("creating a new issue")
                    new_issue = self.jira_client.create_issue(project=project, summary=self.TICKET_CREATE_SUMMARY, description=self.FAILURE_CREATE_COMMENT.format(tag=tag), issuetype={'name': 'Bug'})
                    print(new_issue)
                    self.jira_client.add_worklog(new_issue, timeSpentSeconds=60, comment=desc)
                    labels = [self.JIRA_TICKET_LABEL]
                    new_issue.update(fields={"labels": labels})
                else:
                    print("dry run, NOT creating ticket with details: " + desc)
            else:
                if len(issues) > 1:
                    # more than one open issue found, we log the "issue" and just pick the first one returned
                    print("more than one active issue found for {project}-{tag}".format(project=project, tag=tag))
                print(issues)
                comm = self.TICKET_UPDATE_DESCRIPTION.format(codecoverage_text=codecoverage_text)
                if not dry_run:
                    self.jira_client.add_worklog(issues[0], timeSpentSeconds=60, comment=comm)
                    self.jira_client.add_comment(issues[0], self.FAILURE_GENERAL_COMMENT)
                else:
                    print("dry run, NOT updating ticket with details in worklog: " + comm)
                    print("general comments: " + self.FAILURE_GENERAL_COMMENT)
        except JIRAError as e:
            print("error logging issue to jira: " + str(e.status_code) + " " + e.text)
