import argparse
import config
import util
import requests
import os
import re
import xml.etree.ElementTree as ET
from git.exc import GitCommandError
import urllib

USERNAME = os.environ.get('TEAMCITY_USERNAME')
PASSWORD = os.environ.get('TEAMCITY_PASSWORD')
if USERNAME and PASSWORD:
    TC_BASIC_AUTH = (USERNAME, PASSWORD)
else:
    TC_BASIC_AUTH = None


def get_last_successful_build_id(build_conf_name, teamcity_branch):
    branch_filter = ""
    if teamcity_branch:
        branch_filter = ",branch:{teamcity_branch}".format(teamcity_branch=urllib.quote_plus(teamcity_branch))

    url = "https://teamcity.oci.oraclecorp.com/httpAuth/app/rest/builds/?locator=buildType:id:{build_conf_name},status:SUCCESS{branch_filter},count:1".format(
        build_conf_name=build_conf_name,
        branch_filter=branch_filter)

    print("get_last_successful_build_id url: {}".format(url))

    response = requests.get(url, auth=TC_BASIC_AUTH, verify=False)

    root = ET.fromstring(response.text.encode('utf-8'))
    build_nodes = root.findall("build")
    if len(build_nodes) != 1:
        return None
    else:
        id = build_nodes[0].attrib['id']
        print("Last successful build id for '{}': {}".format(build_conf_name, id))
        return id


def get_commits_of_build(build_id):
    url = "https://teamcity.oci.oraclecorp.com/httpAuth/app/rest/builds/id:{build_id}".format(
        build_id=build_id)

    response = requests.get(url, auth=TC_BASIC_AUTH, verify=False)

    root = ET.fromstring(response.text.encode('utf-8'))

    revision_nodes = root.findall(".//revision")
    commits = [r.attrib['version'] for r in revision_nodes]
    print("Commits of build {}: {}".format(build_id, commits))
    return commits


def get_commit_message_for_commit(repo, commit_hash):
    return repo.git.log('-n 1', "--format='%s%n%b'", commit_hash)


def get_issue_keys_from_repo(repo, last_successful_commits, commit_lookback_range, issue_key_extractor):
    delimiter = '++++++++'
    last_commit_messages = []

    # Find all the commit messages up to (excluding) the last successfully-built commit (up to commit_loopback_range)
    for x in repo.git.log(n=commit_lookback_range, format='%H%n%s%n%b{}'.format(delimiter)).split(delimiter):
        x = x.strip()
        if not x:
            continue  # there's a blank one at the end

        lines = x.split('\n')
        hash = lines[0]
        message = '\n'.join(lines[1:])
        if hash in last_successful_commits:
            print("Reached last successful commit: {}".format(hash))
            break
        last_commit_messages.append({'hash': hash, 'message': message})

    # will be of the form 'Updating pom.xml for [[DEXREQ-123, DEXREQ-234, DEXREQ-456]]
    issue_keys = []

    # go through the commit messages from oldest to newest (so later reverts can remove earlier commits)
    for pair in reversed(last_commit_messages):
        hash = pair['hash']
        commit_message = pair['message'].encode('utf-8')

        new_keys = issue_key_extractor(commit_message)

        # this was probably a revert commit
        if b'revert' in commit_message.lower():
            print("Ignoring commit {} because it contains the string 'revert' -- removing keys: [{}]".format(hash, ", ".join(new_keys)))
            # remove the keys found in this commit
            for key in new_keys:
                try:
                    if key in issue_keys:
                        issue_keys.remove(key)
                except ValueError:
                    pass

            # find the commit hashes mentioned as being reverted
            reverted_commits = re.findall(b'This reverts commit ([a-f0-9]*).', commit_message)

            # loop up the commit messages for the reverted commits
            for rc in reverted_commits:
                try:
                    reverted_message = get_commit_message_for_commit(repo, rc)
                    # and remove those issue keys from our list
                    reverted_keys = issue_key_extractor(reverted_message)

                    print("\tReverted commit {} -- removing keys: [{}]".format(rc, ", ".join(reverted_keys)))
                    for key in reverted_keys:
                        if key in issue_keys:
                            issue_keys.remove(key)
                except GitCommandError:
                    print("Could not get commit message for what looked like a commit: {}".format(rc))
        else:
            print("Considering commit {} -- adding keys: [{}]\n{}".format(hash, ", ".join(new_keys), commit_message))
            issue_keys.extend(new_keys)

        print("==========")

    return issue_keys


def is_work_complete_for_tools(issue, tool_names):
    work_complete = True
    for tool in tool_names:
        # Determines if a given build_tool should report status back to the associated Jira issue
        if not util.is_tool_jira_reportable(tool):
            print("Ignoring {} for checking for work completion due to not a service team facing surface: {}".format(tool))
            continue

        custom_field_id = config.CUSTOM_FIELD_ID_FOR_TOOL[tool]
        custom_status = getattr(issue.fields, custom_field_id)

        # Skip any tools that have been labeled to skip its generation
        if config.BYPASS_CHECK_GENERATION_PREFIX + tool in issue.fields.labels:
            print("Ignoring {} for checking for work completion due to label: {}".format(tool, config.BYPASS_CHECK_GENERATION_PREFIX + tool))
            continue

        if not custom_status or not custom_status.value or not custom_status.value == config.CUSTOM_STATUS_DONE:
            work_complete = False
            break

    return work_complete


def default_issue_key_extractor(commit_message):
    try:
        commit_message = commit_message.decode('utf-8')
    except AttributeError:
        pass
    return util.parse_issue_keys_from_specific_commit_message(commit_message, "{} [[".format(config.GENERATION_COMMIT_MESSAGE_PREFIX))


# - this script will run on commits to the preview or master branch of the different SDKs or the CLI
# - for every trigger (merged commit), this job will look at recent commit messages to find if
#   there are any preview or public DEXREQ tickets that need to be updated to SDK Status: Done
if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Update SDK statuses to Done in DEXREQ tickets.')
    parser.add_argument('--build-id',
                        required=True,
                        help='The team city build id for the build that is running this script.  This is used to update the relevant JIRA tickets with links to the team city build')
    parser.add_argument('--tool',
                        default=config.CLI_NAME,
                        help='The tool for which to generate the build. Accepted values: {}'.format(', '.join(config.TOOL_NAMES)))
    parser.add_argument('--full-version',
                        help='The full version of this SDK / CLI that was produced.  This is used to generate the final comment with a link to the artifacts.')
    parser.add_argument('--short-version',
                        required=False,
                        help='The short version of this SDK / CLI that was produced.  This is used to generate the final comment with a link to the artifacts. Only used by Java (e.g. full version "1.2.44-preview1-20180806.212454-12", short version "1.2.44-preview1-SNAPSHOT")')
    parser.add_argument('--artifactory-repo',
                        required=False,
                        help='The artifactory repository where the artifact was published.  This is used to generate the final comment with a link to the artifacts. Only used by Java (e.g. "opc-public-sdk-snapshot-maven-local" or "opc-public-sdk-release-maven-local")')
    parser.add_argument('--issue',
                        action='append',
                        help='By default, we will process DEXREQ tickets found in recent commits. This allows you to specify specific DEXREQ issue(s) act on: --issue DEXREQ-123.')
    parser.add_argument('--allow-transition-overall-issue-to-done',
                        default=False,
                        action='store_true',
                        help='Allows transitioning overall issue status to "Done" if all CLI / SDK Status fields are "Done"')
    parser.add_argument('--allow-transition-overall-issue-to-deploy',
                        default=False,
                        action="store_true",
                        help='Allows transitioning the overall issue status to "To Deploy" if all CLI / SDK Status fields are "Done"')
    parser.add_argument('--dry-run',
                        default=False,
                        action='store_true',
                        help='Perform a dry-run')
    parser.add_argument('--build-conf-name',
                        required=True,
                        help="TeamCity build configuration name.")
    # Use this query and see if you get the right results:
    # https://teamcity.oci.oraclecorp.com/httpAuth/app/rest/builds/?locator=buildType:id:{build_conf_name),status:SUCCESS,branch:{teamcity_branch},count:100
    parser.add_argument('--teamcity-branch',
                        required=False,
                        help="Branch for TeamCity 'last successful build' query; examples: 'preview', 'master', '<default>")
    parser.add_argument('--build-type',
                        default=config.BUILD_TYPE_INDIVIDUAL_PREVIEW,
                        help='The build type to use, can be one of the following: {}'.format(', '.join(config.VALID_BUILD_TYPES)))

    args = parser.parse_args()
    build_id = args.build_id
    tool_name = args.tool
    full_version = args.full_version
    short_version = args.short_version
    artifactory_repo = args.artifactory_repo
    whitelisted_issues = args.issue
    allow_transition_overall_issue_to_done = args.allow_transition_overall_issue_to_done
    allow_transition_overall_issue_to_deploy = args.allow_transition_overall_issue_to_deploy
    teamcity_buildconf_name = args.build_conf_name
    teamcity_branch = args.teamcity_branch
    build_type = args.build_type
    config.IS_DRY_RUN = args.dry_run

    if TC_BASIC_AUTH and teamcity_buildconf_name:
        # get the commits of the last successful build
        last_successful_build_id = get_last_successful_build_id(teamcity_buildconf_name, teamcity_branch)
        last_successful_commits = get_commits_of_build(last_successful_build_id)
        commit_lookback_range = 1000
    else:
        commit_lookback_range = 10
        last_successful_commits = []
        print("Using lookback range {}, since TEAMCITY_USERNAME, TEAMCITY_PASSWORD, and --build-conf-name are not all set.".format(commit_lookback_range))

    # get current branches
    current_branches = {}
    for name, repo in zip(config.REPO_NAMES_FOR_TOOL[tool_name], config.REPOS_FOR_TOOL[tool_name]):
        b = repo.git.branch()
        print('Branches in repo {}:\n{}'.format(repo, str(b)))
        if not b:
            current_branches[name] = []
            continue
        current_branches[name] = [branch.strip()[2:] for branch in b.split('\n') if branch.startswith('* ')][0]

    print('Current branches per tool: {}'.format(str(current_branches)))

    branch = current_branches[config.REPO_NAMES_FOR_TOOL[tool_name][-1]]

    repo = config.REPOS_FOR_TOOL[tool_name][-1]

    issue_keys = get_issue_keys_from_repo(repo, last_successful_commits, commit_lookback_range, default_issue_key_extractor)
    print("Found issue keys: {}".format(", ".join(issue_keys)))

    if issue_keys:
        for issue_key in issue_keys:
            if whitelisted_issues and issue_key not in whitelisted_issues:
                print('Skipping processing ticket {} because it was not included in --issue filter'.format(issue_key))
                continue

            issue = util.get_dexreq_issue(issue_key)

            # don't try to mark fields on an old ticket type because the fields won't be available
            if issue.fields.issuetype.id != config.PREVIEW_ISSUE_TYPE_ID and issue.fields.issuetype.id != config.PUBLIC_ISSUE_TYPE_ID:
                print('Skipping issue {} of old ticket type: {}'.format(issue_key, issue.fields.issuetype.name))
                continue

            # update issue custom status for tool_name to 'Done'
            status = getattr(issue.fields, config.CUSTOM_FIELD_ID_FOR_TOOL[tool_name])
            if issue.fields and issue.fields.status and issue.fields.status.name == config.STATUS_DONE:
                print('Skipping marking ticket {issue_key} as Done for tool: {tool} because overall ticket status is Done'.format(issue_key=issue_key, tool=tool_name))
            elif status and status.value and status.value == config.CUSTOM_STATUS_DONE:
                # if the ticket is already marked done then ignore it
                print('Skipping re-marking ticket {issue_key} as Done for tool: {tool}'.format(issue_key=issue_key, tool=tool_name))
            else:
                # only mark the ticket SDK status as done if we are able to add a comment with links to the artifact
                comment_template = config.BULK_GENERATION_COMPLETE_COMMENT_TEMPLATE_FOR_TOOL.get(tool_name)
                if not comment_template:
                    print('Skipping marking ticket {issue_key} as Done for tool: {tool} because no comment template is defined'.format(issue_key=issue_key, tool=tool_name))
                    continue

                if not full_version:
                    print('Skipping marking ticket {issue_key} as Done for tool: {tool} because --full-version was not specified'.format(issue_key=issue_key, tool=tool_name))
                    continue

                util.add_jira_comment(
                    issue_key,
                    comment_template.format(
                        full_version=full_version,
                        short_version=short_version,
                        artifactory_repo=artifactory_repo
                    )
                )

                util.transition_issue_per_tool_status(util.JIRA_CLIENT(), issue, config.CUSTOM_STATUS_DONE, tool_name)

                if is_work_complete_for_tools(issue, util.get_jira_reportable_tool_names()):
                    if allow_transition_overall_issue_to_done:
                        util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_DONE)
                    elif allow_transition_overall_issue_to_deploy:
                        util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_TO_DEPLOY)
    else:
        print('Did not find any issue keys in commit message (text between double brackets: "[[issue-key]]"). Not updating any JIRA issues.')
