# - this script is run as the as the last step of a build step of a job triggered on commits to SDK/CLI branches
# of the form auto-preview-* or auto-public-*
# - at this point the pom.xml has been updated, and the new code has been generated
# - this script reports back to JIRA whether or not the code generation + build was successful and commits / pushes the generated code

import argparse
import util
import config
import os
from os.path import exists
import shutil
import re
import shared.bitbucket_utils
from shared.buildsvc_tc_compatibility import build_log_link

try:
    from urllib import quote
except ImportError:
    from urllib.parse import quote


IGNORED_TESTS_FILE_PATH = "../python-cli/changed_services.txt"
SPEC_PR_DESCRIPTION = """
Please review these spec changes, and if they represent (1) your entire change, and (2) nothing else, approve this pull request.{merge_ticket_action}

If something is wrong with these spec changes, please set this pull request to "Needs Work".{decline_ticket_action}

Note that merging or declining this spec change does not affect your service's spec in any way. It is used merely in the automation of the SDK/CLI self-service pipeline.
"""


# Branch suffix is "DEXREQ-673-2019-08-16-20-58-17"
def push_spec_diff(spec_dir, build_type, issue_key, branch_suffix, description, diff_branch_suffix="-diff", is_pr_enabled=True,
        commit_issue_summary_template="Updated spec for {issue_key}{description_text}",
        pr_title_template="Spec changes for {issue_key}{description_text}",
        pr_description=""):
    branch_prefix = util.get_branch_prefix_for_spec_diff(build_type)
    baseline_branch_name = "{}-{}".format(branch_prefix, branch_suffix)
    diff_branch_name = baseline_branch_name + diff_branch_suffix

    repo = config.DEXREQ_REPO
    git = repo.git
    git.fetch("origin")
    print("Switching to dexreq repo branch {}".format(baseline_branch_name))
    git.checkout(baseline_branch_name)
    print("Creating new dexreq repo branch {}".format(diff_branch_name))
    git.checkout(B=diff_branch_name)

    for filename in os.listdir(spec_dir):
        source = os.path.join(spec_dir, filename)
        destination = os.path.join(config.DEXREQ_DIFF_REPO_RELATIVE_LOCATION, filename)
        print("Copying {} -> {}".format(source, destination))
        shutil.rmtree(destination, True)
        shutil.copytree(source, destination, ignore=shutil.ignore_patterns('*.lineNumberMapping'))

    git.add(A=True)

    commit_message = '{commit_prefix} [[{issue_key}]]: {issue_summary}'.format(
        commit_prefix=config.SPEC_BASELINE_COMMIT_MESSAGE_PREFIX,
        issue_key=issue_key,
        issue_summary=commit_issue_summary_template.format(
            issue_key=issue_key,
            description_text="" if not description else (": " + description))
    )

    message = commit_message
    if 'nothing to commit' in git.status():
        message = "{} (no change)".format(message)
    print(message)
    git.commit("-m", message, "--allow-empty")
    if config.IS_DRY_RUN:
        print('DRY-RUN: not pushing to branch {}'.format(diff_branch_name))
    else:
        git.push('-u','origin','HEAD')

    pr_url = None
    title = pr_title_template.format(
        issue_key=issue_key,
        description_text="" if not description else (": " + description)
    )
    if is_pr_enabled:

        if build_type == config.BUILD_TYPE_INDIVIDUAL_PREVIEW:
            merge_ticket_action = " Then transition your DEXREQ ticket to 'Ready for Preview'."
            decline_ticket_action = " Please revise your spec and generate a new spec artifact. Then update the spec version in your DEXREQ ticket and set the ticket status back to 'Processing Requested'."
        else:
            merge_ticket_action = ""
            decline_ticket_action = ""

        pr_url = util.create_pull_request(config.DEXREQ_REPO_NAME, diff_branch_name, title,
            SPEC_PR_DESCRIPTION.format(merge_ticket_action=merge_ticket_action,
                decline_ticket_action=decline_ticket_action),
            "", config.DEXREQ_REPO_NAME, baseline_branch_name)

        print("Pull request created: {}".format(pr_url))

        pr_id = None
        if pr_url:
            m = re.search("^.*bitbucket.*/projects/([^/]*)/repos/([^/]*)/pull-requests/([0-9]*).*$", pr_url)
            if m:
                pr_id = m.group(3)

        issue = util.get_dexreq_issue(issue_key, fields=['created'])
        created_date = getattr(issue.fields, 'created')
        print("To get all spec diff PRs for {}, listing all PRs newer than {}".format(issue_key, created_date))

        # The spec diff PR can't be older than the DEXREQ ticket, so only search that far

        prs = shared.bitbucket_utils.get_all_pullrequest_with_string_after('SDK', config.DEXREQ_REPO_NAME, issue.key, created_date)

        for pr in prs:
            print("Spec change pr {} is {}".format(pr['id'], pr['state']))

            if str(pr['id']) == str(pr_id):
                print("\tThis is the PR that was just opened")
                continue

            if pr['state'] == config.PULL_REQUEST_STATUS_OPEN:
                # Decline other open PRs
                print("\tDeclining spec change PR {}".format(pr['id']))
                shared.bitbucket_utils.decline_pr("SDK", config.DEXREQ_REPO_NAME, pr['id'], pr['version'])
                shared.bitbucket_utils.make_general_comment("SDK", config.DEXREQ_REPO_NAME, pr['id'], "This PR has been declined in favor of a more recent spec diff PR: {}".format(pr_url))
    else:
        # Not technically a PR (yet)
        pr_url = 'https://bitbucket.oci.oraclecorp.com/projects/SDK/repos/{repo}/compare?targetBranch=refs%2Fheads%2F{target_branch}&sourceBranch=refs%2Fheads%2F{diff_branch}&title={title}&description={description}'.format(
            repo=config.DEXREQ_REPO_NAME,
            diff_branch=diff_branch_name,
            title=quote(title),
            description=quote(description),
            target_branch=baseline_branch_name)

        print('Link with spec diff: {}'.format(pr_url))

    return pr_url


def get_pipeline(build_type):
    if build_type == config.BUILD_TYPE_INDIVIDUAL_PREVIEW or build_type == config.BUILD_TYPE_BULK_PENDING_MERGE_PREVIEW:
        return config.PREVIEW_ISSUE_TYPE_NAME
    elif build_type == config.BUILD_TYPE_BULK_PENDING_MERGE_PUBLIC or build_type == config.BUILD_TYPE_INDIVIDUAL_PUBLIC:
        return config.PUBLIC_ISSUE_TYPE_NAME
    else:
        raise ValueError("Unknown build type: ".format(build_type))


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Post generation status updates to DEXREQ tickets.')
    parser.add_argument('--build-id',
                        required=True,
                        help='The team city build id for the build that is running this script.  This is used to update the relevant JIRA tickets with links to the team city build')
    parser.add_argument('--tool',
                        default=config.CLI_NAME,
                        help='The tool for which to generate the build. Accepted values: {}'.format(', '.join(config.TOOL_NAMES)))
    parser.add_argument('--dry-run',
                        default=False,
                        action='store_true',
                        help='Perform a dry-run')
    parser.add_argument('--push-spec-diff',
                        help='Push the changed spec (after generation and pre-processing) from the specified directory into a branch in the SDK/dexreq repo.')
    parser.add_argument('--push-spec-diff-unprotected-by-conditional-groups',
                        help='Push the changed spec WITHOUT ADDING CONDITIONAL GROUPS (after pre-processing) from the specified directory into a branch in the SDK/dexreq repo.')
    parser.add_argument('--build-type',
                        default=config.BUILD_TYPE_INDIVIDUAL_PREVIEW,
                        help='The build type to use, can be one of the following: {}'.format(', '.join(config.VALID_BUILD_TYPES)))
    parser.add_argument('--optional-file-for-dexreq-ticket',
                        default=None,
                        help='An optional text file to be included in the DEXREQ ticket.')

    args = parser.parse_args()

    shared.bitbucket_utils.setup_bitbucket(args)

    build_id = args.build_id
    tool_name = args.tool
    build_type = args.build_type
    config.IS_DRY_RUN = args.dry_run
    shared.bitbucket_utils.dry_run = args.dry_run

    # this script runs after both generation and build have completed
    # it will run even 'If previous steps have failed' to ensure we can report failures
    generation_pass, build_pass = util.were_steps_successful(tool_name)

    # report to JIRA tasks whether or not jobs succeeded
    # if either failed, just give link to build log and say that generation failed for some reason
    # this will be hard for external users to investigate so we want to cover easy errors earlier in the process with explicit errors:
    #   - spec artifact / group / version doesnt exist in artifactory
    #   - invalid param set
    #   - relative spec path doesn't point at a spec (yaml file)

    failure_step = None
    if not generation_pass:
        failure_step = 'Generation'
    elif not build_pass:
        failure_step = 'Build'

    # get current branch of the first repo
    current_branch = [branch.strip()[2:] for branch in config.REPOS_FOR_TOOL[tool_name][-1].git.branch().split('\n') if branch.startswith('* ')][0]

    # TODO: parse last commit message to related DEXREQ issue
    last_commit_message = config.REPOS_FOR_TOOL[tool_name][-1].git.log(n=1, format='%s%n%b')

    # will be of the form 'Updating pom.xml for DEXREQ-123: Preview for RQS
    print('Found last commit: {}'.format(last_commit_message))

    issue_keys = util.parse_issue_keys_from_commit_message(last_commit_message)
    print("Issue keys found: '{}'".format(", ".join(issue_keys)))

    descriptions = {}
    for issue_key in issue_keys:
        m = re.search(r"Updating pom.xml for \[\[.*{}.*\]\]:(.*)".format(issue_key), last_commit_message)
        if m:
            descriptions[issue_key] = m.group(1).strip()

    if failure_step:
        if issue_keys:
            for issue_key in issue_keys:
                if config.IS_DRY_RUN:
                    print("DRY-RUN: not transitioning {} {} status to {}".format(issue_key, tool_name, config.CUSTOM_STATUS_FAILURE))
                else:
                    additional_comment = ""
                    if args.optional_file_for_dexreq_ticket and exists(args.optional_file_for_dexreq_ticket):
                        # additional comment
                        with open(args.optional_file_for_dexreq_ticket) as f:
                            optional_file_contents_for_dexreq_ticket = f.read()
                            additional_comment = "\n\nSpec validator output:\n\n{code}" + optional_file_contents_for_dexreq_ticket + "\n{code}"

                    util.add_jira_comment(
                        issue_key,
                        config.STEP_FAILED_MESSAGE_TEMPLATE.format(
                            failure_step=failure_step,
                            tool_name=tool_name,
                            repos=util.join(config.REPO_NAMES_FOR_TOOL[tool_name]),
                            build_log_link=build_log_link(build_id, "build log"),
                            dex_support_required_status=config.STATUS_DEX_SUPPORT_REQUIRED,
                            additional_comment=additional_comment
                        ),
                        comment_type=config.COMMENT_TYPE_ERROR
                    )

                    # if it makes sense we can assign to different states based on build failing
                    issue = util.get_dexreq_issue(issue_key)
                    util.transition_issue_per_tool_status(util.JIRA_CLIENT(), issue, config.CUSTOM_STATUS_FAILURE, tool_name)
                    # Only transition the overall status to SERVICE TEAM FAILURE INVESTIGATION for non bulk gen build types
                    if build_type not in config.BULK_BUILD_TYPES and util.is_tool_jira_reportable(tool_name):
                        # if an issue is already in 'DEX Support Required' based on failure for another tool, we do not want to overwrite that
                        util.transition_issue_overall_status_if_not_in_status(util.JIRA_CLIENT(), issue, desired_status=config.STATUS_SERVICE_TEAM_FAILURE_INVESTIGATION, blacklisted_status=config.STATUS_DEX_SUPPORT_REQUIRED)
        else:
            print('Did not find any issue keys in commit message (text between double brackets: "[[issue-key]]"). Not updating any JIRA issues.')

        print('{failure_step} for tool: {tool_name} (repos: {repos}) failed. See previous build step(s) for more information'.format(
            failure_step=failure_step,
            tool_name=tool_name,
            repos=", ".join(config.REPO_NAMES_FOR_TOOL[tool_name]))
        )
    else:
        print(config.GENERATION_AND_BUILD_SUCCESSFUL_TEMPLATE.format(
            tool_name=tool_name,
            repos=", ".join(config.REPO_NAMES_FOR_TOOL[tool_name]))
        )

        if args.push_spec_diff and issue_keys:
            for issue_key in issue_keys:
                m = re.search("^.*{}-(.*)$".format(tool_name), current_branch)
                if m:
                    branch_suffix = m.group(1)
                    push_spec_diff(args.push_spec_diff, build_type, issue_key, branch_suffix, descriptions[issue_key] or "")
                else:
                    print("Did not find '{}' and a timestamp in the current branch '{}', not pushing spec diff".format(tool_name, current_branch))

        if args.push_spec_diff_unprotected_by_conditional_groups and issue_keys:
            for issue_key in issue_keys:
                issue = util.get_dexreq_issue(issue_key)
                is_bypassed = config.BYPASS_CHECK_CHANGES_NOT_BEHIND_CONDITIONAL_GROUPS in issue.fields.labels
                m = re.search("^.*{}-(.*)$".format(tool_name), current_branch)
                if m:
                    branch_suffix = m.group(1)
                    link = push_spec_diff(args.push_spec_diff_unprotected_by_conditional_groups, build_type, issue_key, branch_suffix, descriptions[issue_key] or "",
                        diff_branch_suffix="-unprotected", is_pr_enabled=False,
                        commit_issue_summary_template="{bypassed_text}Spec changes not protected by conditional groups for {issue_key}{description_text}".format(
                            issue_key=issue_key, bypassed_text="(Bypassed) " if is_bypassed else "", description_text=descriptions[issue_key]),
                        pr_title_template="{bypassed_text}Spec changes not protected by conditional groups for {issue_key}{description_text}".format(
                            issue_key=issue_key, bypassed_text="(Bypassed) " if is_bypassed else "", description_text=descriptions[issue_key]),
                        pr_description="These changes may be unwanted 'passengers' and cannot be filtered out, except by changing the spec itself. All spec changes should be protected by conditional groups.{bypassed_text}".format(
                            bypassed_text="\n\nThese changes were accepted using a bypass label in the {issue_key} ticket.".format(issue_key=issue_key)))

                    if link:
                        if is_bypassed:
                            util.add_jira_comment(
                                issue_key,
                                config.BYPASSED_UNPROTECTED_CHANGES_MESSAGE_TEMPLATE.format(
                                    pipeline=get_pipeline(build_type),
                                    link=link
                                ),
                                comment_type=config.COMMENT_TYPE_INFO
                            )
                        else:
                            util.add_jira_comment(
                                issue_key,
                                config.UNPROTECTED_CHANGES_MESSAGE_TEMPLATE.format(
                                    pipeline=get_pipeline(build_type),
                                    link=link
                                ),
                                comment_type=config.COMMENT_TYPE_INFO
                            )
                            print("Adding '{}' label to: {}".format(config.CHANGES_NOT_BEHIND_CONDITIONAL_GROUPS_LABEL, issue.key))
                            issue.add_field_value('labels', config.CHANGES_NOT_BEHIND_CONDITIONAL_GROUPS_LABEL)
                else:
                    print("Did not find '{}' and a timestamp in the current branch '{}', not pushing spec diff for changes not protected by conditional groups".format(tool_name, current_branch))
    # This is important in the preview generation when there is a conflict between generated and extended code in Python
    # CLI. In this case we comment the extended code and ignore the affected service tests. these ignored tests are
    # saved in a temp file changed_services.txt in the Python CLI directory, We check here if the file exists and
    # had data in it, if yes then we add the CLI-ManualChangesRequired label to the ticket to exclude it from the
    # bulk preview.
    try:
        print('Checking if CLI-ManualChangesRequired label need to be added to the ticket.')
        print('Looking for the file ' + IGNORED_TESTS_FILE_PATH)
        f = open(IGNORED_TESTS_FILE_PATH, "r")
        affected_services = f.read()
        print('File found and the affected services are: ')
        print(affected_services)
        if not affected_services or not len(affected_services):
            print('There is no services affected by a conflict between Generated and Extended code in Python CLI.')
        elif issue_keys:
            for issue_key in issue_keys:
                if config.IS_DRY_RUN:
                    print("DRY-RUN: adding label {} to {}".format(config.CLI_REQUIRED_MANUAL_CHANGES_LABEL, issue_key))
                else:
                    issue = util.get_dexreq_issue(issue_key)
                    if config.CLI_REQUIRED_MANUAL_CHANGES_LABEL not in issue.fields.labels:
                        print("Adding CLI-ManualChangesRequired label to: " + issue.key)
                        issue.add_field_value('labels', config.CLI_REQUIRED_MANUAL_CHANGES_LABEL)
                        issue.update()
                        print("Added CLI-ManualChangesRequired label to: " + issue.key)

            # We delete the temp file since no need for it now
            print("Removing the temp file: " + IGNORED_TESTS_FILE_PATH)
            # We delete the temp file since no need for it now
            os.remove(IGNORED_TESTS_FILE_PATH)
        else:
            print('Did not find any issue keys in commit message (text between double brackets: "[[issue-key]]"). Not able to add CLI-ManualChangesRequired label to the JIRA issue.')
    except:  # noqa: ignore=E722
        print('Could not find the file ' + IGNORED_TESTS_FILE_PATH)
