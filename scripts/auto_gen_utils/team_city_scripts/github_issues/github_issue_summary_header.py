import re


class GitHubIssueSummaryHeader:
    """Class that maps a Github issue to a specific Jira issue and vice-versa.

       The contract for mapping is based on the jira issue summary being in a specific format.
       TODO: This contract is brittle.  In the future, consider adding specific Jira fields
         so that the contract is explicit.
    """

    # Format is repo_name, github_issue_num, github_issue_summary
    # Example: '[GitHub Issue (oci-java-sdk #66)]: Empty compartment id in audit events'
    JIRA_ISSUE_SUMMARY_TEMPLATE = '[GitHub Issue ({} #{})]: {}'

    def __init__(self, repo_name, issue_num, summary=''):
        if not repo_name:
            raise ValueError('A GitHub repository name must be defined')
        if not issue_num:
            raise ValueError('An issue number must be defined')

        self.repo_name = repo_name.encode('utf-8').strip()
        self.issue_num = int(issue_num)
        self.summary = summary.encode('utf-8').strip()

    def get_repo_name(self):
        return self.repo_name

    def get_issue_num(self):
        return self.issue_num

    def get_summary(self):
        return self.summary

    def __str__(self):
        return 'GitHubIssueSummaryHeader: [repo_name: {}, issue_num: {}, summary: {}]'.format(self.repo_name,
                                                                                              self.issue_num,
                                                                                              self.summary)

    def __eq__(self, other):
        if not isinstance(other, self.__class__):
            return False

        if self.repo_name != other.repo_name:
            return False

        if self.issue_num != other.issue_num:
            return False

        # Explicitly ignore the summary in case the customer edited it between script executions

        return True

    def __ne__(self, other):
        return not self.__eq__(other)

    def __hash__(self):
        return hash(self.repo_name) + hash(self.issue_num)

    def to_jira_issue_summary(self):
        return self.JIRA_ISSUE_SUMMARY_TEMPLATE.format(self.repo_name, self.issue_num, self.summary)

    @staticmethod
    def from_jira_issue_summary(jira_summary):
        """Creates a new GitHubIssueSummaryHeader from the given jira issue summary.

            :param str jira_summary: The jira summary formatted as defined by JIRA_ISSUE_SUMMARY_TEMPLATE
            :return: the header, or None if the summary could not be parsed
            :rtype: GitHubIssueSummaryHeader
        """
        repo_name = None
        search = re.search(r'GitHub Issue \((.*?) #', jira_summary)
        if search:
            repo_name = search.group(1)

        issue_num = None
        search = re.search(r'#(.*?)\)', jira_summary)
        if search:
            issue_num = search.group(1)

        summary = None
        search = re.search(r']: (.*)$', jira_summary)
        if search:
            summary = search.group(1)

        if repo_name and issue_num and summary:
            return GitHubIssueSummaryHeader(repo_name, issue_num, summary)
        else:
            return None

    @staticmethod
    def from_github_issue(repo_name, github_issue):
        """Creates a new GitHubIssueSummaryHeader from the given repository name and github issue.

            :param str repo_name: The Github repository name
            :param github.Issue github_issue: The issue retrieved from Github
            :return: The header
            :rtype: GitHubIssueSummaryHeader
        """
        return GitHubIssueSummaryHeader(repo_name, github_issue.number, github_issue.title)
