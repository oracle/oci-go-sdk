from datetime import datetime
import re
import git

import util
import config

import argparse
from jira import JIRAError


DEXREQ_AUTOMATION_NAME_AND_EMAIL = 'DEXREQ Automation noreply+dexreq@oracle.com'


def init_branches():
    global DEBUG_DEXREQ_BRANCH_PREFIX
    global GENERATED_AUTO_PREVIEW_BRANCH_PATTERN
    global AUTO_PREVIEW_BRANCH_PATTERN
    global GENERATED_AUTO_PUBLIC_BRANCH_PATTERN
    global AUTO_PUBLIC_BRANCH_PATTERN
    global AUTO_PREVIEW_BRANCH_PREFIX
    global AUTO_PUBLIC_BRANCH_PREFIX

    DEBUG_DEXREQ_BRANCH_PREFIX = config.get_debug_dexreq_branch_prefix()

    GENERATED_AUTO_PREVIEW_BRANCH_PATTERN = 'refs/remotes/origin/generated-{debug_dexreq_branch_prefix}auto-v2-preview-*'.format(
        debug_dexreq_branch_prefix=DEBUG_DEXREQ_BRANCH_PREFIX)
    AUTO_PREVIEW_BRANCH_PATTERN = 'refs/remotes/origin/{debug_dexreq_branch_prefix}auto-v2-preview-*'.format(
        debug_dexreq_branch_prefix=DEBUG_DEXREQ_BRANCH_PREFIX)
    GENERATED_AUTO_PUBLIC_BRANCH_PATTERN = 'refs/remotes/origin/generated-{debug_dexreq_branch_prefix}auto-v2-public-*'.format(
        debug_dexreq_branch_prefix=DEBUG_DEXREQ_BRANCH_PREFIX)
    AUTO_PUBLIC_BRANCH_PATTERN = 'refs/remotes/origin/{debug_dexreq_branch_prefix}auto-v2-public-*'.format(
        debug_dexreq_branch_prefix=DEBUG_DEXREQ_BRANCH_PREFIX)

    AUTO_PREVIEW_BRANCH_PREFIX = DEBUG_DEXREQ_BRANCH_PREFIX + 'auto-v2-preview-'
    AUTO_PUBLIC_BRANCH_PREFIX = DEBUG_DEXREQ_BRANCH_PREFIX + 'auto-v2-public-'


init_branches()

# no reason to keep 'merge to/from GitHub' branches longer than a week
MAX_AGE_IN_DAYS_GITHUB_MERGE_BRANCH = 7

# no reason to keep a bulk preview branch more than 2 weeks old
MAX_AGE_IN_DAYS_BULK_PREVIEW_BRANCH = 14

# these are only used internally in the auto generation process so we can delete them very promptly
# (the generated-auto-preview-* branches are the ones we give to partner teams)
MAX_AGE_IN_DAYS_AUTO_PREVIEW_BRANCH = 2

# these are only used internally in the auto generation process so we can delete them very promptly
# (the generated-auto-public-* branches are the ones we give to partner teams)
MAX_AGE_IN_DAYS_AUTO_PUBLIC_BRANCH = 2

DEXREQ_PREFIX = 'DEXREQ-'
DEXREQ_TOOL_NAME = "DEXREQ"

CLEAN_TOOL_NAMES = config.TOOL_NAMES + [DEXREQ_TOOL_NAME]
CLEAN_TOOL_REPOS_FOR_TOOL = {}
CLEAN_TOOL_REPOS_FOR_TOOL.update(config.REPOS_FOR_TOOL)
CLEAN_TOOL_REPOS_FOR_TOOL[DEXREQ_TOOL_NAME] = [config.DEXREQ_REPO]

TERMINAL_STATES = [config.STATUS_DONE, config.STATUS_WITHDRAWN, config.STATUS_CLOSED]


def clean_auto_preview_branches(tool_name, issues, branches):
    # start by making sure we are on the base branch with no other changes present -- TODO: why is this important?
    # checkout_sdk_and_cli_branches(base_branch, tool_name)
    for repo in CLEAN_TOOL_REPOS_FOR_TOOL[tool_name]:
        repo.git.reset('HEAD','--hard')
        repo.git.clean('-f')

        # prune remote branches so we only try to delete branches that still exist
        repo.git.remote('update', 'origin', '--prune')

    for repo in CLEAN_TOOL_REPOS_FOR_TOOL[tool_name]:
        if branches:
            remote_refs = branches
        else:
            prefix = 'origin/'

            remote_refs = []

            safe_branch_prefixes = []
            safe_branch_prefixes.extend(config.BRANCH_PREFIXES_SAFE_FOR_DELETION)
            if tool_name == DEXREQ_TOOL_NAME:
                # Hack to clear up old DEXREQ branches
                safe_branch_prefixes.extend(["public-DEXREQ", "preview-DEXREQ", "bulk_public-DEXREQ", "bulk_preview-DEXREQ"])

            for branch_prefix in safe_branch_prefixes:
                pattern = 'refs/remotes/origin/' + branch_prefix + '*'
                remote_refs.extend(repo.git.for_each_ref(pattern, format='%(refname:short)').split('\n'))

            remote_refs = [remote_ref[len(prefix):] for remote_ref in remote_refs if remote_ref.startswith(prefix)]

            if len(remote_refs) == 0:
                print('No remote branches found for repo: {}. Skipping.'.format(repo.working_dir))
                continue

        for ref in remote_refs:
            if issues:
                found = False
                for i in issues:
                    if i in ref:
                        found = True
                        break
                if not found:
                    continue

            can_be_deleted, message = branch_ok_for_deletion(ref)
            if can_be_deleted:
                print('Branch {} can be deleted - {}'.format(ref, message))
            else:
                print('Branch {} not ready for deletion - {}'.format(ref, message))
                continue

            try:
                util.safe_delete_branch(repo, ref)
            except git.exc.GitCommandError as e:
                print('Failed to delete ref: {}. Exception: {}', ref, str(e))


# Returns True/False and a message
def check_dexreq_in_terminal_state(branch_name):
    # find dexreq-<<number>> from the branch name
    if DEXREQ_PREFIX in branch_name:
        start = branch_name.index(DEXREQ_PREFIX)
        try:
            end = branch_name.index('-', start + len(DEXREQ_PREFIX))
            issue = util.get_dexreq_issue(branch_name[start:end])

            if issue.fields.status:
                if issue.fields.status.name in TERMINAL_STATES:
                    return True, "terminal '{}'".format(issue.fields.status.name)
                else:
                    return False, "non-terminal '{}'".format(issue.fields.status.name)
            else:
                return False, "unknown issue status"
        except ValueError as e:
            # e.g. for "auto-preview-JavaSDK-DEXREQ-143" without timestamp
            print(e)
        except JIRAError as e:
            print(e)
            if "404" in str(e):
                return True, "does not exist in JIRA"
            else:
                raise e

    return False, "no DEXREQ issue"


# Returns True/False and a message
def branch_ok_for_deletion(branch_name):
    # delete bulk branches older than some threshold: MAX_AGE_IN_DAYS_BULK_PREVIEW_BRANCH
    days_old = 0
    match = re.search('.*(20[0-9]{2}-[0-1]{1}[0-9]{1}-[0-3]{1}[0-9]{1}-[0-2]{1}[0-9]{1}-[0-5]{1}[0-9]{1}-[0-5]{1}[0-9]{1})', branch_name)
    if match:
        groups = match.groups()
        if len(groups) == 1:
            branch_timestamp = datetime.strptime(groups[0], '%Y-%m-%d-%H-%M-%S')
            days_old = (datetime.now() - branch_timestamp).days
            if '-bulk-' in branch_name:
                if days_old >= MAX_AGE_IN_DAYS_BULK_PREVIEW_BRANCH:
                    return True, ">= {} days".format(MAX_AGE_IN_DAYS_BULK_PREVIEW_BRANCH)
                else:
                    return False, "< {} days".format(MAX_AGE_IN_DAYS_BULK_PREVIEW_BRANCH)
            elif 'merge' in branch_name and 'github' in branch_name:
                if days_old >= MAX_AGE_IN_DAYS_GITHUB_MERGE_BRANCH:
                    return True, ">= {} days".format(MAX_AGE_IN_DAYS_GITHUB_MERGE_BRANCH)
                else:
                    return False, "< {} days".format(MAX_AGE_IN_DAYS_GITHUB_MERGE_BRANCH)
    else:
        print('Could not parse timestamp from branch name: {}'.format(branch_name))

    # the auto-preview branches are only used temporarily to trigger the generation + build process
    # after making the commit with the pom updates so we can delete them very promptly
    if branch_name.startswith(AUTO_PREVIEW_BRANCH_PREFIX) and days_old >= MAX_AGE_IN_DAYS_AUTO_PREVIEW_BRANCH:
        return True, ">= {} days".format(MAX_AGE_IN_DAYS_AUTO_PREVIEW_BRANCH)

    if branch_name.startswith(AUTO_PUBLIC_BRANCH_PREFIX) and days_old >= MAX_AGE_IN_DAYS_AUTO_PUBLIC_BRANCH:
        return True, ">= {} days".format(MAX_AGE_IN_DAYS_AUTO_PUBLIC_BRANCH)

    # Delete generated branches if DEXREQ ticket is in terminal state i.e Done / Withdrawn
    return check_dexreq_in_terminal_state(branch_name)


# def checkout_sdk_and_cli_branches(base_branch, tool_name):
#     for repo in CLEAN_TOOL_REPOS_FOR_TOOL[tool_name]:
#         repo.git.checkout(base_branch)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Clean up all auto generated branches for a given tool and base branch.')
    parser.add_argument('--base-branch',
                        default='preview',
                        help='The base branch to start from')
    parser.add_argument('--tool',
                        default='ALL',
                        help='The tool for which to generate the preview. Accepted values: ALL, {}'.format(', '.join(CLEAN_TOOL_NAMES)))
    parser.add_argument('--dry-run',
                        default=False,
                        action='store_true',
                        help='Perform a dry-run')
    parser.add_argument('--issue',
                        action='append',
                        help='By default, we query JIRA. This allows you to specify a DEXREQ issue to process instead: --issue DEXREQ-123')
    parser.add_argument('--branch',
                        action='append',
                        help='By default, we query Bitbucket. This allows you to specify a branch to process instead: --branch generated-auto-preview-JavaSDK-DEXREQ-40-2018-06-28-22-32-06')

    args = parser.parse_args()

    if args.branch and args.issue:
        raise ValueError("Cannot use --issue and --branch together.")

    base_branch = args.base_branch
    tool_name = args.tool
    config.IS_DRY_RUN = args.dry_run

    if tool_name != 'ALL' and tool_name not in CLEAN_TOOL_NAMES:
        raise ValueError("Tool name must be one of: ALL, {}".format(', '.join(CLEAN_TOOL_NAMES)))

    if tool_name == 'ALL':
        for tool in CLEAN_TOOL_NAMES:
            print('Cleaning branches for tool: {}'.format(tool))
            clean_auto_preview_branches(tool, args.issue, args.branch)
    else:
        clean_auto_preview_branches(tool_name, args.issue, args.branch)
