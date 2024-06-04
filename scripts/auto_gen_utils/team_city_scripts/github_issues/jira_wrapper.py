import traceback
import six

from github_issue_summary_header import GitHubIssueSummaryHeader
from jira import JIRA
from jira.exceptions import JIRAError
from util import debug_print, get_partial_debug_print_func, get_partial_print_func


class JiraTransitionError(ValueError):
    """Raised when unable to transition a jira issue to a specific status."""
    pass


class JiraWrapper:
    """Class to wrap calls to Jira"""
    JIRA_LABEL = 'Customer_Github_Issue'
    JIRA_EXISTING_ISSUES_FOR_PROJECT_QUERY = 'project in ({}) and labels = ' + JIRA_LABEL
    JIRA_EXISTING_ISSUES_QUERY = 'labels = ' + JIRA_LABEL
    JIRA_OPTIONS = {
        'server': 'https://jira.oci.oraclecorp.com',
        'rest_api_version': 2,
        'verify': True
    }
    # Different service team projects configure different transition states in their work flows.
    JIRA_CLOSED_STATUS_LIST = ['Done', 'Resolved', 'Resolve Issue', 'Closed', 'Close Issue']
    JIRA_OPEN_STATUS_LIST = ['Backlog', 'Needs Triage', 'To Do', 'Reopen']

    JIRA_COMMENT_TYPE_INFO = "INFO"
    JIRA_COMMENT_TYPE_ERROR = "ERROR"
    JIRA_COMMENT_TYPE_SUCCESS = "SUCCESS"
    JIRA_COMMENT_TYPE_TO_COLOR = {
        JIRA_COMMENT_TYPE_INFO: '#707070',  # gray
        JIRA_COMMENT_TYPE_ERROR: '#FF0000',  # red
        JIRA_COMMENT_TYPE_SUCCESS: '#14892c',  # green
    }

    def __init__(self, username=None, password=None, session_id=None, commit=False):
        if username and password:
            self.jira_client = JIRA(self.JIRA_OPTIONS, basic_auth=(username, password))
        elif session_id:
            cookie_options = dict(self. JIRA_OPTIONS)
            cookie_options['cookies'] = {'JSESSIONID': session_id}
            self.jira_client = JIRA(cookie_options)
        else:
            raise ValueError("Either JIRA Credentials or a session ID must be provided")
        self.commit = commit

    def get_jira_client(self):
        return self.jira_client

    def get_all_jira_issues(self):
        """Queries Jira for all existing issues that have the 'Customer_GitHub_Issue' label.

        :return: dictionary of summaries to jira issues
        :rtype: dict([GitHubIssueSummaryHeader|str], jira.Issue)
        """
        # When searching for existing issues from Jira, ensure that all of the results are paginated and combined
        # into a list for processing.
        returned_issue_count_from_search = 50
        total_issue_count_from_search = 0
        all_jira_issues = []
        while returned_issue_count_from_search == 50:
            returned_issues_page = self.jira_client.search_issues(self.JIRA_EXISTING_ISSUES_QUERY,
                                                                  startAt=total_issue_count_from_search,
                                                                  maxResults=50,
                                                                  fields='summary,description,status,resolution,type',
                                                                  expand='expand')
            all_jira_issues.extend(returned_issues_page)
            returned_issue_count_from_search = len(returned_issues_page)
            total_issue_count_from_search += returned_issue_count_from_search

        summary_to_jira_issue_dict = dict()
        for jira_issue in all_jira_issues:
            jira_issue_summary_text = six.text_type(jira_issue.fields.summary).strip()
            jira_summary = GitHubIssueSummaryHeader.from_jira_issue_summary(jira_issue_summary_text)
            if jira_summary:
                summary_to_jira_issue_dict[jira_summary] = jira_issue
            else:
                # If the issue summary isn't for a service team issue, it could be for a DEX filed unlabeled issue.
                # If so, add the issue to the dictionary if the issue has a non-closed status.
                jira_issue_status = str(jira_issue.fields.status)
                if jira_issue_status not in self.JIRA_CLOSED_STATUS_LIST:
                    debug_print('Found non-service team jira issue [{}] '
                                'with non closed status [{}]'.format(jira_issue,
                                                                     jira_issue_status))
                    summary_to_jira_issue_dict[jira_issue_summary_text] = jira_issue
                else:
                    print('[WARN] Skipping jira issue.  Unable to parse service team Jira issue summary for '
                          '[Key: {}, Summary: {}]'.format(jira_issue.key, jira_issue.fields.summary))

        return summary_to_jira_issue_dict

    def create_jira_issue(self, project, summary, description, service_issue=True, indent=0):
        """ Creates a new jira issue.

        :param str project: the jira project
        :param str summary: the summary for the issue
        :param str description: the description for the issue
        :param (optional) int indent: the number of spaces to indent debug logging
        """
        jira_summary_max_length = 254
        if len(summary) > jira_summary_max_length:
            summary = summary[0: jira_summary_max_length - 4] + '...'

        if self.commit:
            labels = [self.JIRA_LABEL, 'PP']
            if service_issue:
                labels.append('Service-Owned-PP')

            new_issue = self.jira_client.create_issue(project=project,
                                                      summary=summary,
                                                      description=description,
                                                      issuetype={'name': 'Task'},
                                                      labels=labels)
            print_indent = get_partial_print_func(indent)
            print_indent('Created new jira issue [project: {}, '
                         'key: {}, summary: {}]: {}'.format(project,
                                                            new_issue.key,
                                                            summary,
                                                            new_issue.permalink()))
        else:
            debug_print_indent = get_partial_debug_print_func(indent)
            debug_print_indent('**** DRY RUN **** Not creating a new jira issue: '
                               '[project: {}, summary: {}, description:\n{}'.format(project,
                                                                                    summary,
                                                                                    description))

    def transition_jira_issue_status(self, issue, status_list, do_add_resolution=False, indent=0):
        """Transitions a jira issue to a status contained within the given status list.  A list of status values is
        done as different jira projects define different issue work flows.  Finding the first matching status provides
        more coverage to ensure that a given jira issue can be transitioned to a viable status.

        :param jira.Issue issue: the jira issue
        :param list(str) status_list: the list of possible status values to transition to
        :param (optional) boolean do_add_resolution: if True, then a Done resolution will be applied, else, False.
        :param (optional) int indent: the number of spaces to indent debug logging
        :raises JiraTransitionError: if no matching status was found for the issue's project work flow
        """

        transitions_from_jira = []
        transition_names_from_jira = []  # Used for error handling
        transitions = self.jira_client.transitions(issue)
        for transition in transitions:
            transitions_from_jira.append(transition)
            transition_names_from_jira.append(transition['name'])

        transition_to_apply = None
        for status in status_list:
            for transition in transitions_from_jira:
                if transition['name'].lower() == status.lower():
                    transition_to_apply = transition
                    break
            if transition_to_apply:
                break

        debug_print_indent = get_partial_debug_print_func(indent)
        try:
            if issue.fields.status.name == status:
                debug_print_indent('Not transitioning issue {} to status "{}" ({}) '
                                   'because it already has that status'.format(issue.key,
                                                                               transition_to_apply['name'],
                                                                               transition_to_apply['id']))
        except AttributeError:
            debug_print_indent('Not transitioning issue {} to status "{}" ({}) '
                               'Unable to retrieve jira issue status '
                               '(AttributeError)'.format(issue.key,
                                                         transition_to_apply['name'],
                                                         transition_to_apply['id']))
            traceback.print_exc()
        except ValueError:
            debug_print_indent('Not transitioning issue {} to status "{}" ({})'
                               'Unable to retrieve jira issue status (ValueError)'.format(issue.key,
                                                                                          transition_to_apply['name'],
                                                                                          transition_to_apply['id']))
            traceback.print_exc()

        if transition_to_apply:
            if self.commit:
                debug_print_indent('Transitioning {} to {} using transition {} (do_add_resolution: {})'.format(
                    issue.key,
                    transition_to_apply['name'],
                    transition_to_apply['id'],
                    do_add_resolution)
                )
                try:
                    if do_add_resolution:
                        self.jira_client.transition_issue(issue, transition_to_apply['id'], resolution={'name': 'Done'})
                    else:
                        self.jira_client.transition_issue(issue, transition_to_apply['id'])
                except JIRAError as e:
                    raise JiraTransitionError("An error occurred while transitioning issue {} to '{}' ({}): {}".format(
                        issue.key,
                        transition_to_apply['name'],
                        transition_to_apply['id'],
                        e.message)
                    )
            else:
                debug_print_indent('**** DRY RUN **** Not transitioning [{}] to "{}" '
                                   'using transition {}'.format(issue.key,
                                                                transition_to_apply['name'],
                                                                transition_to_apply['id']))
        else:
            raise JiraTransitionError(
                "Don't know how to transition this issue to any of the status values [{}].  Jira returned: [{}]".format(
                    ', '.join(status_list),
                    ', '.join(transition_names_from_jira))
            )

    def add_jira_comment(self, issue_key, comment, comment_type=None, indent=0):
        """Adds a comment to the given issue

        :param number issue_key: the identifier for the jira issue
        :param str comment: The comment to add
        :param str comment_type: The comment type to control the color rendering in the jira comment.
                                 See JIRA_COMMENT_TYPE_TO_COLOR
        :param (optional) int indent: the number of spaces to indent debug logging
        """
        debug_print_indent = get_partial_debug_print_func(indent)
        color = self.JIRA_COMMENT_TYPE_TO_COLOR.get(comment_type, None)
        if color:
            comment = '{{color:{color}}}{comment}{{color}}'.format(color=color, comment=comment)

        if self.commit:
            debug_print_indent("Making the following comment for {}".format(issue_key))
            debug_print_indent(comment)
            self.jira_client.add_comment(issue_key, comment)
        else:
            debug_print_indent("**** DRY RUN **** Not making the following comment for {}\n{}".format(issue_key,
                                                                                                      comment))
