# This script parses the issue routing configuration and updates the configured GitHub repositories to ensure
#   that the defined labels exist in each repository, respectively.
import argparse
import util

from github import Github, GithubObject
from issue_config import IssueConfig
from util import get_partial_print_func, get_partial_debug_print_func

print_indent_4 = get_partial_print_func(4)
debug_print_indent_4 = get_partial_debug_print_func(4)


def determine_missing_service_labels_for_repo(gh_repo, configured_labels):
    """ Queries a given GitHub repository' configured labels and determines which labels from the given list are
        missing.

    :param gh_repo: the GitHub repository instance
    :param configured_labels: the list of service labels
    :return: the list of missing service labels
    :rtype: list str
    """
    labels_from_repo = []
    for label_from_repo in gh_repo.get_labels():
        labels_from_repo.append(label_from_repo.name)

    missing_labels = []
    for configured_label in configured_labels:
        if configured_label not in labels_from_repo:
            missing_labels.append(configured_label)

    return missing_labels


def add_missing_labels_to_repo(gh_repo, missing_labels, label_config):
    """ Adds the list of missing labels to the given GitHub repository.

    :param gh_repo: the GitHub repository instance.
    :param missing_labels: the list of missing labels.
    :param label_config: the label configuration as a dict
    """
    for missing_label in missing_labels:
        label_cfg_to_add = None
        for label_cfg in label_config:
            if label_cfg.get('label') == missing_label:
                label_cfg_to_add = label_cfg
        if label_cfg_to_add is None:
            continue

        label_name = label_cfg_to_add.get('label')
        label_color = label_cfg_to_add.get('color')
        label_description = GithubObject.NotSet
        if 'description' in label_cfg_to_add:
            label_description = label_cfg_to_add.get('description')

        status_msg = "Creating label ['{}', #{}, {}] in repo [{}]".format(label_name,
                                                                          label_color,
                                                                          label_description,
                                                                          gh_repo.name)

        if commit:
            print_indent_4(status_msg)
            repo.create_label(label_name, label_color, description=label_description)
        else:
            print_indent_4("**** DRY RUN ****" + status_msg)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Parses the issue routing and GitHub repository configuration to '
                                                 'update the labels in GitHub')
    parser.add_argument('--issue-routing-config',
                        default=None,
                        help='Path to the issue routing configuration yaml file')
    parser.add_argument('--repo-config',
                        default=None,
                        help='Path to the GitHub repository configuration yaml file')
    parser.add_argument('--github-access-token',
                        required=True,
                        help="The GitHub access token that has access rights to the configured repositories")
    parser.add_argument('--debug',
                        default=False,
                        action='store_true',
                        help="Enable debugging console logging")
    parser.add_argument('--commit',
                        default=False,
                        action='store_true',
                        help="True if labels in GitHub are to be updated; else, False (default)")
    args = parser.parse_args()
    util.debug = args.debug
    commit = args.commit

    # Load Config
    config = IssueConfig(args.issue_routing_config, args.repo_config)
    if util.debug:
        config.print_config()

    gh = Github(args.github_access_token)
    config.verify_github_user_access(gh)
    issue_routing_config = config.get_issue_routing_config()

    standard_labels = config.get_standard_labels()
    standard_label_config = issue_routing_config.get('standardLabels')
    service_labels = config.get_service_labels()
    service_label_config = issue_routing_config.get('serviceLabels')

    for repo_name in config.get_repo_config().get('supportedRepos'):
        print("Processing labels for [{}]...".format(repo_name))
        full_repo_name = config.get_full_repo_name(repo_name)
        debug_print_indent_4("Fetching repo for [{}]".format(full_repo_name))
        repo = gh.get_repo(full_repo_name)

        missing_service_labels_for_repo = determine_missing_service_labels_for_repo(repo, service_labels)
        missing_standard_labels_for_repo = determine_missing_service_labels_for_repo(repo, standard_labels)

        add_missing_labels_to_repo(repo, missing_service_labels_for_repo, service_label_config)
        add_missing_labels_to_repo(repo, missing_standard_labels_for_repo, standard_label_config)
