import git
import os


PYTHON_CLI_WORKING_DIRECTORY = os.environ.get('PYTHON_CLI_DIR')
PYTHON_SDK_WORKING_DIRECTORY = os.environ.get('PYTHON_SDK_DIR')
COMMIT_MESSAGE = "{} Self Service Manual Changes"


def checkout_branch(directory, branch):
    repo = git.Repo(directory)
    repo.remote().fetch()
    repo.remote().pull()
    try:
        repo.git.checkout(branch)
        print("Successfully checked out branch: {}".format(branch))
    except Exception as e:
        print(e)


def checkout_self_service_branch(directory, branch, ticket):
    repo = git.Repo(directory)
    repo.remote().fetch()
    repo.remote().pull()
    self_service_branch = branch + '-' + ticket
    try:
        repo.git.reset('--hard')
        repo.git.checkout(branch)
        try:
            repo.git.branch('-D', self_service_branch)
            repo.git.checkout('-b', self_service_branch)
        except:  # noqa: E722
            repo.git.checkout('-b', self_service_branch)
    except Exception as e:
        print(e)
    print("Successfully checked out branch for manual changes: {}".format(self_service_branch))
    return self_service_branch


def push_all_generated_changes_to_remote(directory, ticket):
    try:
        repo = git.Repo(directory)
        repo.git.add(A=True)
        repo.git.reset('--', 'setup.py')
        repo.git.reset('--', 'requirements.txt')
        repo.index.commit(COMMIT_MESSAGE.format(ticket))
        repo.git.push('-u', 'origin', repo.active_branch, force=True)
        print("Pushed all changes to remote branch: {}".format(repo.active_branch))
    except Exception as e:
        print(e)
