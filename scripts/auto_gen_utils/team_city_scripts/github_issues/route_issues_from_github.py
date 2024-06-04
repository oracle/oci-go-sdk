import argparse
import six
import sys
import util

from github import Github
from jira_wrapper import JiraWrapper, JiraTransitionError
from jira.exceptions import JIRAError
from issue_config import IssueConfig
from util import debug_print, get_partial_print_func, get_partial_debug_print_func
from github_issue_summary_header import GitHubIssueSummaryHeader

# Service Team Jira configuration for closing an issue
# Each service team has different work flows configured; thus, different field requirements
# upon transitioning a jira issue to a resolved/closed state.
JIRA_SERVICE_TEAM_CLOSE_TRANSITION_CONFIG = {
    'BMI': {
        'do_add_resolution': False
    },
    'DEX': {
        'do_add_resolution': False
    }
}

LINE_SEP_SIZE = 80
CLOSED_JIRA_COMMENT_TEMPLATE = """
Closing this issue since the Github issue has been closed.

{}
"""

REOPEN_JIRA_COMMENT_TEMPLATE = """
Reopening this issue since the Github issue is still in an open state.

If this issue has been resolved, ensure that the Github issue has been closed.

{}
"""

FALLBACK_TO_DEX_JIRA_DESCRIPTION_TEMPLATE = """
Failed to create a JIRA ticket for Github issue with {labels} label.
Move this ticket to the {project} queue.

Please review and respond to the customer issue reported via Github:

Customer Github issue:
    SDK Repository: {repo_name}
    Issue #: {issue_num}
    Title: {title}
    Labels: {labels}
    URL: {url}


Note that you'll need to close the Github issue prior to closing this Jira issue to prevent it from being automatically reopened.
"""

NEW_JIRA_ISSUE_DESCRIPTION_TEMPLATE = """
Please review and respond to the customer issue reported via Github:

Customer Github issue:
    SDK Repository: {repo_name}
    Issue #: {issue_num}
    Title: {title}
    Labels: {labels}
    URL: {url}


Note that you'll need to close the Github issue prior to closing this Jira issue to prevent it from being automatically reopened.
"""

UNLABELED_JIRA_ISSUE_DESCRIPTION_TEMPLATE = """
The following customer Github issues need to be labeled for proper issue routing or answered + closed:

{}
"""

closed_issue_date_threshold = 30
jira_wrapper = None
# Indented print/debug print functions
print_indent_4 = get_partial_print_func(4)
print_indent_8 = get_partial_print_func(8)
print_indent_12 = get_partial_print_func(12)
debug_print_indent_4 = get_partial_debug_print_func(4)
debug_print_indent_8 = get_partial_debug_print_func(8)
debug_print_indent_12 = get_partial_debug_print_func(12)


def get_github_issues(gh_repo):
    """Retrieves the list of customer filed  GitHub issues from a given repository.

    :param Repository gh_repo: the Github repository instance
    :return: the list of closed customer issues and the list of open customer issues
    :rtype: list github.Issue, list github.Issue
    """
    all_issues = gh_repo.get_issues(state='all')
    # Filter out all PRs as GitHub considers PRs as issues.
    customer_issues = []
    for gh_issue in all_issues:
        try:
            if not gh_issue.pull_request:
                customer_issues.append(gh_issue)
        except AttributeError:
            continue

    closed_issues = []
    open_issues = []
    for c_issue in customer_issues:
        if c_issue.state == 'closed':
            closed_issues.append(c_issue)
        elif c_issue.state == 'open':
            open_issues.append(c_issue)
        else:
            raise ValueError("Unrecognized GitHub issue state: {}".format(c_issue))
    return closed_issues, open_issues


def handle_existing_jira_issue_for_closed_issue(jira_issue, gh_issue):
    """Handles an existing jira issue for the given closed customer Github issue by closing the jira issue.

    :param jira.Issue jira_issue: the jira issue that corresponds to the customer issue
    :param github.Issue gh_issue: the customer Github issue
    """

    # First determine if the 'do_add_resolution' boolean flag needs to be set
    do_add_resolution = True
    for project_prefix, jira_service_config in JIRA_SERVICE_TEAM_CLOSE_TRANSITION_CONFIG.items():
        if jira_issue.key.startswith(project_prefix):
            do_add_resolution = jira_service_config.get('do_add_resolution', True)

    # Check the state to ensure that the issue status values are in sync.  GitHub is source of truth.
    jira_status = str(jira_issue.fields.status)
    if jira_status not in jira_wrapper.JIRA_CLOSED_STATUS_LIST:
        # Close jira issue with comment
        try:
            print_indent_12('Closing jira issue: {} (Opts: do_add_resolution: {})'.format(jira_issue.key,
                                                                                          do_add_resolution))
            jira_wrapper.transition_jira_issue_status(jira_issue,
                                                      jira_wrapper.JIRA_CLOSED_STATUS_LIST,
                                                      do_add_resolution=do_add_resolution,
                                                      indent=12)
            jira_wrapper.add_jira_comment(jira_issue,
                                          CLOSED_JIRA_COMMENT_TEMPLATE.format(gh_issue.html_url),
                                          comment_type=jira_wrapper.JIRA_COMMENT_TYPE_INFO,
                                          indent=12)
        except JiraTransitionError as e:
            print_indent_12('Unable to close jira issue {}: {}'.format(jira_issue, six.text_type(e.message)))
    else:
        print_indent_12('Skipping jira issue [{}] as it already closed [{}]'.format(jira_issue, jira_status))


def process_closed_issues_for_repo(gh_repo_name, closed_gh_issues, all_jira_issues_dict):
    """Processes all closed customer Github issues by closing any jira issue that is open, if it exists.

    :param str gh_repo_name: the Github repository name
    :param list(github.Issue) closed_gh_issues:  the list of closed customer Github issues
    :param dict(GitHubIssueSummaryHeader, jira.Issue) all_jira_issues_dict: the dictionary of jira issues
    """
    print_indent_4('Closed Github issues for {}'.format(gh_repo_name))
    for gh_issue in closed_gh_issues:
        gh_issue_summary = GitHubIssueSummaryHeader.from_github_issue(repo_name, gh_issue)
        jira_issue = all_jira_issues_dict.get(gh_issue_summary)
        if jira_issue:
            debug_print_indent_8('Found existing Jira issue: {}'.format(jira_issue.key))
            handle_existing_jira_issue_for_closed_issue(jira_issue, gh_issue)
        else:
            debug_print_indent_8('Skipping. No jira issue found for {}'.format(gh_issue_summary.to_jira_issue_summary()))
    print_indent_4('Finished processing closed Github issues for {}'.format(gh_repo_name))
    debug_print('-' * LINE_SEP_SIZE)


def get_labels_from_github_issue(gh_issue):
    """Gets all of the labels from the given Github issue.

    :param github.Issue gh_issue: the customer Github issue
    :return: the list of labels
    :rtype: list(str)
    """
    labels = []
    for label in gh_issue.labels:
        labels.append(label.name)
    return labels


def find_intersecting_service_routing_configs(labels_list, routing_config):
    """Finds all intersection issue routing configurations based on the list of labels.

    :param list(str) labels_list: the list of labels
    :param dict routing_config: the issue routing configuration
    :return: the list of issue routing configurations that match
    :rtype: list(dict)
    """
    intersecting_routing_info_list = []
    for routing_config in routing_config.get('serviceLabels'):
        if routing_config.get('label') in labels_list:
            intersecting_routing_info_list.append(routing_config)
    return intersecting_routing_info_list


def create_new_jira_for_issue(gh_issue, gh_issue_summary, intersecting_cfgs, routing_config):
    """Creates a new jira issue for the given Github customer issue.

    :param github.Issue gh_issue: the customer issue
    :param GitHubIssueSummaryHeader gh_issue_summary: the issue summary
    :param list(dict) intersecting_cfgs: the list of issue routing configurations that maps by label to the customer
                                         Github issue
    :param dict routing_config: the issue routing configuration
    """
    cfg = intersecting_cfgs[0]
    if intersecting_cfgs and len(intersecting_cfgs) > 1:
        print_indent_8('Found more than one issue routing config.  Using: {}'.format(cfg))
        debug_print_indent_12('Configurations: {}'.format(''.join(intersecting_cfgs)))
    jira_routing_info = cfg.get('teamRoutingInfo')[0].get('jira')

    gh_issue_labels = []
    for gh_label in gh_issue.labels:
        gh_issue_labels.append(gh_label.name)

    # If there is an issue while creating a new JIRA ticket, try to create a JIRA ticket to the SDK team (DEX queue)
    # as fallback and then manually move the ticket to the service team.
    try:
        description = NEW_JIRA_ISSUE_DESCRIPTION_TEMPLATE.format(repo_name=gh_issue_summary.get_repo_name(),
                                                             issue_num=gh_issue_summary.get_issue_num(),
                                                             title=gh_issue_summary.get_summary(),
                                                             labels=', '.join(gh_issue_labels),
                                                             url=gh_issue.html_url)
        jira_wrapper.create_jira_issue(jira_routing_info.get('project'),
                                       gh_issue_summary.to_jira_issue_summary(),
                                       description,
                                       indent=8)
    except JIRAError:
        print('Failed to create JIRA ticket in {} queue. As a fallback option, creating a ticket in {} queue'.format(
            jira_routing_info.get('project'), routing_config.get('unlabeledRoutingInfo').get('jira').get('project')))
        description = FALLBACK_TO_DEX_JIRA_DESCRIPTION_TEMPLATE.format(project=jira_routing_info.get('project'),
                                                            labels=', '.join(gh_issue_labels),
                                                            repo_name=gh_issue_summary.get_repo_name(),
                                                            issue_num=gh_issue_summary.get_issue_num(),
                                                            title=gh_issue_summary.get_summary(),
                                                            url=gh_issue.html_url)
        jira_wrapper.create_jira_issue(routing_config.get('unlabeledRoutingInfo').get('jira').get('project'),
                                       gh_issue_summary.to_jira_issue_summary(),
                                       description,
                                       indent=8)


def create_or_update_jira_for_unlabeled_issues(repository_name, unlabeled_gh_issues, jira_issues_dict, routing_config):
    """Creates a new Jira ticket, or adds a comment to an already open Jira ticket  so that all unlabeled customer
         Github issues can be manually labeled for the next execution.

    :param str repository_name: the Github repository name
    :param list(github.Issue) unlabeled_gh_issues: the list of customer Github issues missing matching routing labels
    :param dict([GitHubIssueSummaryHeader|str], jira.Issue) jira_issues_dict: the dictionary of all existing jira issues
    :param dict routing_config: the issue routing configuration
    """
    jira_issue_summary = 'Unlabeled Github issues for {}'.format(repository_name)
    jira_project = routing_config.get('unlabeledRoutingInfo').get('jira').get('project')
    issue_list_str = []
    for gh_summary, gh_issue in unlabeled_gh_issues.items():
        issue_list_str.append('{}: {}\n'.format(gh_summary.to_jira_issue_summary(), gh_issue.html_url))
    content_to_add_or_update = UNLABELED_JIRA_ISSUE_DESCRIPTION_TEMPLATE.format("\n".join(issue_list_str))

    existing_open_jira_issue_for_repo = jira_issues_dict.get(jira_issue_summary)
    if existing_open_jira_issue_for_repo:
        debug_print_indent_12('Found existing open jira issue [{}]'.format(existing_open_jira_issue_for_repo))
        jira_wrapper.add_jira_comment(existing_open_jira_issue_for_repo,
                                      content_to_add_or_update,
                                      comment_type=jira_wrapper.JIRA_COMMENT_TYPE_INFO,
                                      indent=12)
    else:
        debug_print_indent_12('No existing open jira issue for repo [{}]'.format(repository_name))
        jira_wrapper.create_jira_issue(jira_project, jira_issue_summary, content_to_add_or_update, indent=12)


def handle_existing_jira_issue_for_open_gh_issue(jira_issue, gh_issue):
    """ Handles an existing jira issue for the given open Github customer issue.

    Logic is as follows:
      If the existing jira issue is in a closed state, then try to reopen the jira issue.
      Else, do nothing.

    :param jira.Issue jira_issue: the existing jira issue that corresponds to the customer Github issue
    :param github.Issue gh_issue: the associated customer Github issue
    """
    # Next check state to ensure that they are in sync.  GitHub is source of truth.
    jira_status = str(jira_issue.fields.status)
    if jira_status in jira_wrapper.JIRA_CLOSED_STATUS_LIST:
        # Reopen with comment
        try:
            print_indent_12('Reopening jira issue: {}'.format(jira_issue))
            jira_wrapper.transition_jira_issue_status(jira_issue, jira_wrapper.JIRA_OPEN_STATUS_LIST, indent=12)
            jira_wrapper.add_jira_comment(jira_issue,
                                          REOPEN_JIRA_COMMENT_TEMPLATE.format(gh_issue.html_url),
                                          comment_type=jira_wrapper.JIRA_COMMENT_TYPE_INFO,
                                          indent=12)
        except JiraTransitionError as e:
            print_indent_12('Unable to reopen jira issue {}: {}'.format(jira_issue, six.text_type(e.message)))
    else:
        print_indent_12('Skipping jira issue [{}] as it already is in a non-closed status [{}]'.format(jira_issue,
                                                                                                       jira_status))


def process_open_issues_for_repo(gh_repo_name, open_gh_issues, existing_jira_issues_dict, routing_config):
    """ Processes all open customer Github issues.
    Logic is as follows:
      For every open Github issue, check to see if it exists in Jira
        If exists, then process the existing Jira.  See handle_existing_jira_issue_for_open_gh_issue.
        If not exists, then resolve the routing config based on the labels.
            If routing config can't be resolved, then include in batch ticket to DEX team to label GH issues.
            If routing config is resolved, then create a new Jira to the service team.
      Finally, create a jira to the DEX team for all open issues that couldn't be routed to the service teams.

    :param github.Repository gh_repo_name: the Github repository
    :param list(github.Issue) open_gh_issues: the list of open customer Github issues
    :param dict([GitHubIssueSummaryHeader|str], jira.Issue) existing_jira_issues_dict: the dictionary
    :param dict from IssueConfig routing_config: the issue routing configuration
    """
    print_indent_4('Open Github issues for {}'.format(gh_repo_name))
    unlabeled_issues = dict()
    for gh_issue in open_gh_issues:
        gh_issue_summary = GitHubIssueSummaryHeader.from_github_issue(gh_repo_name, gh_issue)
        jira_issue = existing_jira_issues_dict.get(gh_issue_summary)
        if jira_issue:
            debug_print_indent_8('Found existing Jira issue: {}'.format(jira_issue.key))
            handle_existing_jira_issue_for_open_gh_issue(jira_issue, gh_issue)
        else:
            jira_issue_summary = gh_issue_summary.to_jira_issue_summary()
            debug_print_indent_8('No existing jira issue found for: {}'.format(jira_issue_summary))
            gh_issue_labels = get_labels_from_github_issue(gh_issue)
            intersecting_configs = find_intersecting_service_routing_configs(gh_issue_labels, routing_config)
            # Handle unlabeled issues later to create a single Jira issue assigned to DEX.
            # Note: Issues with a label that isn't recognized will also be treated the same.
            if not gh_issue_labels or not intersecting_configs:
                debug_print_indent_12('No github issue label or intersecting issue routing configurations')
                unlabeled_issues[gh_issue_summary] = gh_issue
                continue

            # Create a new Jira issue.
            create_new_jira_for_issue(gh_issue, gh_issue_summary, intersecting_configs, routing_config)

    # Now process all unlabeled issues
    if unlabeled_issues:
        debug_print_indent_8('Processing unlabeled issues ({})'.format(len(unlabeled_issues)))
        create_or_update_jira_for_unlabeled_issues(repo_name,
                                                   unlabeled_issues,
                                                   existing_jira_issues_dict,
                                                   routing_config)
    else:
        debug_print_indent_8('No unlabeled issues to process')
    print_indent_4('Finished processing open Github issues for {}'.format(gh_repo_name))
    debug_print('-' * LINE_SEP_SIZE)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Queries for issues from the configured repositories and '
                                                 'creates a Jira issues for the corresponding configured project')
    parser.add_argument('--issue-routing-config',
                        default=None,
                        help='(Optional) Path to the issue routing configuration yaml file')
    parser.add_argument('--repo-config',
                        default=None,
                        help='(Optional) Path to the GitHub repository configuration yaml file')
    parser.add_argument('--github-access-token',
                        required=True,
                        help='[Required] The GitHub access token that has access rights to the configured repositories')
    parser.add_argument('--jira-username',
                        default=None,
                        help='(Optional if --jira-session-id is defined) The JIRA username')
    parser.add_argument('--jira-password',
                        default=None,
                        help='(Optional if --jira-session-id is defined) The JIRA password')
    parser.add_argument('--jira-session-id',
                        default=None,
                        help='(Optional if --jira-username and --jira-password ard defined) '
                             'The active JIRA session ID used for client auth')
    parser.add_argument('--debug',
                        default=False,
                        action='store_true',
                        help='(Optional) Enable debugging console logging')
    parser.add_argument('--commit',
                        default=False,
                        action='store_true',
                        help='(Optional) True if JIRA issues are to be created; else, False (default)')
    args = parser.parse_args()
    util.debug = args.debug
    commit = args.commit

    # Load Config
    config = IssueConfig(args.issue_routing_config, args.repo_config)
    if util.debug:
        config.print_config()
        print('=' * LINE_SEP_SIZE)
        sys.stdout.flush()

    # GitHub
    gh = Github(args.github_access_token)
    config.verify_github_user_access(gh)
    issue_routing_config = config.get_issue_routing_config()

    # Jira
    jira_wrapper = JiraWrapper(args.jira_username, args.jira_password, args.jira_session_id, commit)
    all_existing_jira_issues_dict = jira_wrapper.get_all_jira_issues()
    if util.debug:
        print('Existing Jira issues:')
        for summary, issue in all_existing_jira_issues_dict.items():
            issue_summary = six.text_type(issue.fields.summary).encode('utf-8', 'ignore')
            print('    [{}]:\n'
                  '        Link: {}\n'
                  '        Status: {},\n'
                  '        Resolution: {},\n'
                  '        Summary: {}'.format(issue.key,
                                               issue.permalink(),
                                               issue.fields.status,
                                               issue.fields.resolution,
                                               issue_summary))
        print('=' * LINE_SEP_SIZE)
        sys.stdout.flush()

    for repo_name in config.get_repo_config().get('supportedRepos'):
        try:
            print('Processing GitHub issues for [{}]...'.format(repo_name))
            full_repo_name = config.get_full_repo_name(repo_name)
            repo = gh.get_repo(full_repo_name)

            closed_github_issues, open_github_issues = get_github_issues(repo)
            if closed_github_issues:
                process_closed_issues_for_repo(repo_name, closed_github_issues, all_existing_jira_issues_dict)
            else:
                debug_print_indent_4('No closed issues to process for repo: {}'.format(repo_name))
            debug_print('+' * LINE_SEP_SIZE)

            if open_github_issues:
                process_open_issues_for_repo(repo_name,
                                             open_github_issues,
                                             all_existing_jira_issues_dict,
                                             issue_routing_config)
            else:
                debug_print_indent_4('No open issues to process for repo: {}'.format(repo_name))
        except Exception:
            print('Error while processing GitHub issues for [{}]. Routing github issue for other SDKs, if any.'.format(repo_name))
            continue
        debug_print('+' * LINE_SEP_SIZE)
    debug_print('=' * LINE_SEP_SIZE)
    print('Finished')
