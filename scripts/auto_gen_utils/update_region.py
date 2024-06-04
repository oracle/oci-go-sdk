import argparse
import config
import datetime
import os
import urllib3
import util
from sdk_regions_updater.region_updater_utils import get_new_regions_info_from_issues, get_issues_with_special_regions_to_be_ignored
from sdk_regions_updater.sdk_regions_updater import SdkRegionsUpdater
from git import Repo
from shared.bitbucket_utils import setup_bitbucket, get_pullrequest  # noqa: ignore=F402

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)


TOOL_ARGUMENT_ALL = 'All'
BRANCH_TIMESTAMP = datetime.datetime.now().strftime('%Y-%m-%d-%H-%M-%S')
PR_COMMENT_PREFIX = 'PR posted for'
# Region type information in Storekeeper is stored as free form texts. 
# The ONSR region type contains "(ONSR)" substring
# The DRCC region type contains "DedicatedRegion" substring
# The Lab Testing region type contains "Development" substring
REGION_TYPES_TO_IGNORE = ['(ONSR)', 'Development']
DEFAULT_REGION_AUTOMATION_OWNER = os.environ.get('DEFAULT_REGION_AUTOMATION_OWNER') or 'anurggar'


def checkout_branch(repo_path, base_branch, new_branch):
    # Checkout branch and create commit
    repo = Repo.init(repo_path)
    repo.git.checkout(base_branch)
    repo.git.checkout(B=new_branch)


def push_change_to_remote(repo_path, tool_name, new_region_ids):
    repo = Repo.init(repo_path)
    repo.git.add(A=True)
    message = 'Region update for {} with new region(s): {}'.format(tool_name, new_region_ids)
    repo.git.commit("-m", message, "--allow-empty")
    repo.git.push('-u','origin','HEAD')


def create_pull_request(base_branch, new_branch, tool_name, new_region_ids, issues):
    # create PR
    repo_name = config.REPO_NAMES_FOR_TOOL[tool_name][0]
    repo_id = util.get_repo_id(repo_name)
    print('repo id for {} is {}'.format(config.REPO_NAMES_FOR_TOOL[tool_name], repo_id))
    # TODO: Need to update title and description
    title = 'New region support for {} - {}'.format(tool_name, new_region_ids)
    issue_keys = (issue.key for issue in issues)
    description = 'This PR includes new region(s) from {}.'.format(', '.join(issue_keys))
    pr_url = util.create_pull_request(repo_name, new_branch, title, description, repo_id, repo_name, base_branch)
    print("Automatically generated pull request: {}".format(pr_url))
    return pr_url


def check_in_region_update_for_sdk(base_branch, tool_name, new_regions, issues):
    region_ids = []
    for region in new_regions:
        region_ids.append(region['regionIdentifier'])
    region_ids_str = ', '.join(region_ids)
    new_branch = '{}-region-update-{}'.format(tool_name.lower(), BRANCH_TIMESTAMP)
    repo_path = os.environ.get('{}_path'.format(tool_name))
    if repo_path is None:
        repo_path = config.REPO_RELATIVE_LOCATION_FOR_TOOL[tool_name]
    checkout_branch(repo_path, base_branch, new_branch)
    print('Branch {} checked out.'.format(new_branch))
    sdk_updater = SdkRegionsUpdater(tool_name)
    sdk_updater.update(new_regions)
    print('region information updated.')
    push_change_to_remote(repo_path, tool_name, region_ids_str)
    print('changes pushed to remote branch.')
    pr_url = create_pull_request(base_branch, new_branch, tool_name, region_ids_str, issues)
    print('PR created.')

    for issue in issues:
        comment = '{} {}: {}'.format(PR_COMMENT_PREFIX, tool_name, pr_url)
        util.add_jira_comment(issue.key, comment)


def close_tickets_if_all_prs_are_merged():
    print('Checking if there are new region support tickets to be closed.')
    in_progress_issues = util.get_in_progress_region_support_tickets()
    for issue in in_progress_issues:
        print('Checking if {} can be closed.'.format(issue.key))
        pr_urls = util.get_all_pr_urls_from_comment(issue.key, PR_COMMENT_PREFIX)
        if (len(pr_urls) > 0):
            if util.are_all_prs_merged(pr_urls):
                comment = 'All PRs have been merged. Closing the ticket.'
                util.add_jira_comment(issue.key, comment)
                util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_DONE)
            else:
                print('Not all PRs have been merged.')
        else:
            print('No PRs found for the ticket.')


def process_new_tickets_with_special_regions(base_branch):
    issues = util.get_unprocessed_region_suppport_tickets(base_branch)
    issues_to_ignore, issues_with_invalid_regions = get_issues_with_special_regions_to_be_ignored(issues, REGION_TYPES_TO_IGNORE)
    for issue in issues_to_ignore:
        region_id = issue.raw['fields']['summary'].split()[-1]
        print('Region {} in issue {} should not be included in SDK. Closing the issue directly.'.format(region_id, issue.key))
        comment = 'Region {} should not be included in SDK. Closing the issue directly.'.format(region_id)
        util.add_jira_comment(issue.key, comment)
        util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_CLOSED)
    for issue in issues_with_invalid_regions:
        region_id = issue.raw['fields']['summary'].split()[-1]
        contact = issue.fields.assignee.name if issue.fields.assignee else DEFAULT_REGION_AUTOMATION_OWNER
        comment = '[~{}],\nRegion {} in issue {} is invalid. Please check with the Region build team to verify if this region can be added or not'.format(contact, region_id, issue.key)
        print(comment)
        util.add_jira_comment(issue.key, comment)
        util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_MORE_INFORMATION_NEEDED)


def process_new_tickets(tool_name, base_branch):
    print('Checking if there are new region support tickets to be processed.')
    issues = util.get_region_support_tickets_to_process(base_branch)
    new_regions = get_new_regions_info_from_issues(issues)
    if len(new_regions) == 0:
        print('No new regions to update.')
    else:
        if (tool_name == TOOL_ARGUMENT_ALL):
            for tool in config.SDKS_SUPPORTING_REGION_UPDATE:
                check_in_region_update_for_sdk(base_branch, tool, new_regions, issues)
        else:
            check_in_region_update_for_sdk(base_branch, tool_name, new_regions, issues)
        for issue in issues:
            comment = 'Transitioning ticket status to In Progress.'
            util.add_jira_comment(issue.key, comment)
            util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_IN_PROGRESS)
        issue_keys = (issue.key for issue in issues)
        print('{} new regions updated for tickets {}.'.format(len(new_regions), ', '.join(issue_keys)))


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('--tool',
                        required=True,
                        choices=config.SDKS_SUPPORTING_REGION_UPDATE + [TOOL_ARGUMENT_ALL],
                        help='The tool for which to generate the preview. Accepted values: {}.'.format(', '.join(config.SDKS_SUPPORTING_REGION_UPDATE + [TOOL_ARGUMENT_ALL])))
    parser.add_argument('--base-branch',
                        default='preview',
                        help='The base branch to start from (default = preview).')
    args = parser.parse_args()
    tool_name = args.tool
    base_branch = args.base_branch

    setup_bitbucket(None)

    # Close tickets in progress if all PRs have been merged
    close_tickets_if_all_prs_are_merged()

    # Close tickets for regions that are not supposed to be added to SDK
    process_new_tickets_with_special_regions(base_branch)

    # Find new tickets to process and update regions
    process_new_tickets(tool_name, base_branch)
