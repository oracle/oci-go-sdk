# - this script is run as the first build step of a job triggered on commits to SDK/CLI branches
# of the form auto-preview-* or auto-public-*
# - process_preview_jira_queue.py will make a commit to an auto-preview-* or auto-public-* branch in the
# SDK (for Python CLI, in Python SDK AND the CLI)
# - this script will check out the equivalent branch

import argparse
import util
import config

BUILD_TYPE_DESCRIPTION = {
    config.BUILD_TYPE_INDIVIDUAL_PREVIEW: "preview SDK",
    config.BUILD_TYPE_INDIVIDUAL_PUBLIC: "public SDK",
    config.BUILD_TYPE_BULK_PENDING_MERGE_PREVIEW: "bulk build of the preview SDK",
    config.BUILD_TYPE_BULK_PENDING_MERGE_PUBLIC: "bulk build of the public SDK"
}

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Post "beginning" status update to DEXREQ tickets.')
    parser.add_argument('--build-id',
                        required=True,
                        help='The team city build id for the build that is running this script.  This is used to update the relevant JIRA tickets with links to the team city build')
    parser.add_argument('--tool',
                        default=config.CLI_NAME,
                        help='The tool for which to generate the build. Accepted values: {}'.format(', '.join(config.TOOL_NAMES)))
    parser.add_argument('--build-type',
                        required=False,
                        help='The build type to use, can be one of the following: {}'.format(', '.join(config.VALID_BUILD_TYPES)))
    parser.add_argument('--dry-run',
                        default=False,
                        action='store_true',
                        help='Perform a dry-run')
    parser.add_argument('--issue',
                        action='append',
                        help='Only DEXREQ issues in this filter will be affected: --issue DEXREQ-123')

    args = parser.parse_args()
    build_id = args.build_id
    tool_name = args.tool
    build_type = args.build_type
    issues_whitelist = args.issue
    if build_type:
        build_description = " for the {}".format(BUILD_TYPE_DESCRIPTION[build_type])
    else:
        build_description = ""
    config.IS_DRY_RUN = args.dry_run

    # get current branch of the first repo
    current_branch = [branch.strip()[2:] for branch in config.REPOS_FOR_TOOL[tool_name][0].git.branch().split('\n') if branch.startswith('* ')][0]

    last_commit_message = util.get_last_commit_message(tool_name)

    # will be of the form 'Updating pom.xml for DEXREQ-123: Preview for RQS
    print('Found last commit: {}'.format(last_commit_message))

    issue_keys = util.parse_issue_keys_from_commit_message(last_commit_message)
    if issue_keys:
        for issue_key in issue_keys:
            if issues_whitelist and issue_key not in issues_whitelist:
                print('Skipping processing ticket {} because it was not included in --issue filter'.format(issue_key))
                continue
    else:
        print('Did not find any issue keys in commit message (text between double brackets: "[[issue-key]]"). Not updating any JIRA issues.')
