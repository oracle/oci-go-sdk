# Manages the configuration for all of the GitHub issue routing related scripts.
import os
import yaml
from pprint import PrettyPrinter

pp = PrettyPrinter(indent=2)


def load_yaml(yaml_file):
    """Loads a given YAML file.

    :param str yaml_file:
        The path of the yaml file to load.
    :return: The dictionary containing the parsed yaml file
    :rtype: dict
    """
    with open(yaml_file, 'r') as stream:
        return yaml.safe_load(stream)


class IssueConfig:
    """Loads the configuration yaml files in order to access the configuration for GitHub issue routing.

    :param str issue_routing_config_file: (optional)
        The path to the issue routing configuration yaml file. If undefined, then the default configuration file
        bundled with this script will be loaded.
    :param str repo_config_file: (optional)
        The path to the supported GitHub issues configuration yaml file.  If undefined, then the default configuration
        file bundled with this script will be loaded.
    """
    def __init__(self, issue_routing_config_file=None, repo_config_file=None):
        base_path = os.path.dirname(os.path.realpath(__file__))
        if issue_routing_config_file is None:
            self.issue_routing_config = load_yaml(os.path.join(base_path, "github_label_routing_config.yaml"))
        else:
            self.issue_routing_config = load_yaml(issue_routing_config_file)

        if repo_config_file is None:
            self.repo_config = load_yaml(os.path.join(base_path, "github_repo_config.yaml"))
        else:
            self.repo_config = load_yaml(repo_config_file)

    def get_issue_routing_config(self):
        """Gets the parsed issue routing configuration.

        :return: The parsed issue routing configuration.
        :rtype: dict
        """
        return self.issue_routing_config

    def get_repo_config(self):
        """Gets the parsed GitHub repository configuration.

        :return: The parsed repository configuration.
        :rtype: dict
        """
        return self.repo_config

    def get_standard_labels(self):
        """Gets the list of standard labels.

        :return: the list of configured standard labels.
        :rtype: list str
        """
        labels = []
        for label_entry in self.get_issue_routing_config().get('standardLabels'):
            labels.append(label_entry.get('label'))
        return labels

    def get_service_labels(self):
        """Gets the list of service labels.

        :return: the list of configured service labels.
        :rtype: list str
        """
        labels = []
        for label_entry in self.get_issue_routing_config().get('serviceLabels'):
            labels.append(label_entry.get('label'))
        return labels

    def get_service_teams_jira_projects(self):
        """Gets the list of service team JIRA projects.

        :return: the list of JIRA projects
        :rtype: list str
        """
        jira_projects = []
        for service_label_cfg in self.get_issue_routing_config().get('serviceLabels'):
            for team_routing_info in service_label_cfg.get('teamRoutingInfo'):
                jira_projects.append(team_routing_info.get('jira').get('project'))
        return jira_projects

    def print_config(self):
        """Pretty prints the configuration to stdout."""
        print('Issue Routing Config:')
        pp.pprint(self.get_issue_routing_config())

        print('GitHub Repo Config:')
        pp.pprint(self.get_repo_config())

    def verify_github_user_access(self, github):
        """Verifies that the current GitHub instance has access to the configured list of supported repositories.

           :param Github github: the GitHub instance
           :raise ValueError if the configured GitHub instance does not have access to all of the configured repositories.
           """
        gh_user_repo_names = []
        for repo in github.get_user().get_repos():
            gh_user_repo_names.append(repo.name)

        inaccessible_repos = []
        for configured_repo_name in self.get_repo_config().get('supportedRepos'):
            if configured_repo_name not in gh_user_repo_names:
                inaccessible_repos.append(configured_repo_name)

        if inaccessible_repos:
            raise ValueError(
                "The configured GitHub access credentials does not have access to the following configured "
                "repos: {}".format(inaccessible_repos))

    def get_full_repo_name(self, short_repo_name):
        """Gets the full GitHub repository name based on the configured owner and provided repository name

        :param str short_repo_name: the repository name
        :return: the full repository name
        :rtype: str
        """
        return "{}/{}".format(self.get_repo_config().get('owner'), short_repo_name)
