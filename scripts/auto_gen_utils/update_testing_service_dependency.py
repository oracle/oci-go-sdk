import util
import argparse
import config
from git import Repo
from shared.bitbucket_utils import setup_bitbucket, get_pullrequest  # noqa: ignore=F402


def push_change_to_remote(repo_path):
    repo = Repo.init(repo_path)
    repo.git.add(A=True)
    message = 'Pom Dependencies Update'
    repo.git.commit("-m", message, "--allow-empty")
    repo.git.push('-u','origin','HEAD')


def create_pull_request(base_branch, new_branch):
    # create PR
    repo_name = config.REPO_NAMES_FOR_TOOL["TestingService"][0]
    repo_id = util.get_repo_id(repo_name)
    print('repo id for {} is {}'.format(config.REPO_NAMES_FOR_TOOL["TestingService"], repo_id))
    # TODO: Need to update title and description
    title = 'OCI Testing Service Update Pom dependencies'
    description = 'This PR includes new dependencies from java sdk'
    pr_url = util.create_pull_request(repo_name, new_branch, title, description, repo_id, repo_name, base_branch)
    print("Automatically generated pull request: {}".format(pr_url))
    return pr_url


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('--base-branch',
                        default='preview',
                        help='The base branch to start from (default = preview).')
    parser.add_argument('--branch-name',
                    help='The branch name')
    parser.add_argument('--repo-path',
                    help='The repo location')
    args = parser.parse_args()
    base_branch = args.base_branch
    branch_name = args.branch_name
    repo_path = args.repo_path

    setup_bitbucket(None)
    push_change_to_remote(repo_path)
    pr_link = create_pull_request(base_branch, branch_name)
    print(pr_link)
