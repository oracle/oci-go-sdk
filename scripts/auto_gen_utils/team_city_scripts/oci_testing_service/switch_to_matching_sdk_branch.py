from __future__ import print_function
import argparse
import os
import ssl
import sys
import urllib3
import re
from git import Repo, GitCommandError

import ocits_shared
from ocits_shared import HELP_URL, TC_URL, get_dexreq_tickets, printv, get_jira_issue, get_pull_requests_for_issue, filter_pull_requests, get_master_javasdk_pr_url

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../..'))

from shared.bitbucket_utils import setup_bitbucket, get_pullrequest, get_pr_source_branch, make_general_comment, get_pr_source_clone_ssh_url, get_repo_permissions_url  # noqa: ignore=F402
import shared.bitbucket_utils  # noqa: ignore=F402
import util  # noqa: ignore=F402
import config  # noqa: ignore=F402

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
if (not os.environ.get('PYTHONHTTPSVERIFY', '') and getattr(ssl, '_create_unverified_context', None)):
    ssl._create_default_https_context = ssl._create_unverified_context


#
# Warning messages
#

MATCHING_SDK_BRANCH_NOT_FOUND = """
Not running the OCI testing service tests.

For pull requests against the `master` branch, there must be a pull request of the {sdk_name} targeting its `master` branch, and that pull request has to be referenced from the public DEXREQ ticket. Unfortunately, no pull request of the {sdk_name} targeting `master` was found in the DEXREQ {ticket_sop} {tickets}.

{author_text}Make sure that you reference the same public DEXREQ ticket in both the OCI Testing Service pull request and the {sdk_name} pull request.

After you have made sure the {sdk_name} pull request is referenced from the DEXREQ {ticket_sop} {tickets}, restart the [TeamCity build]({tc_link}) by asking an SDK/CLI team member on #oci_public_sdks to click on the 'Run' button.

For more information, see [Pull Request Validation Builds for the Testing Service]({help_url}).
"""

COULD_NOT_DETERMINE_CLONE_URL = """
Not running the OCI testing service tests.

Could not determine the source repository's clone URL for the {sdk_name} pull request: {sdk_pr}.

{author_text}Please make sure you have given the 'DEXREQ Automation' user [read access to your {sdk_name} repository]({permissions_url}), and that you have added the ['Key for OCI testing service on-PR build access']({help_url}#PullRequestValidationBuildsfortheTestingService-WhenUsingYourOwnFork) to your {sdk_name} repository's access keys.

**Note: The read access and the access key have to be added to {sdk_name} repository, not the OCI testing service repository!**

After you have made sure access rights and keys are correct, restart the [TeamCity build]({tc_link}) by asking an SDK/CLI team member on #oci_public_sdks to click on the 'Run' button.

For more information, see the [When Using Your Own Fork]({help_url}#PullRequestValidationBuildsfortheTestingService-WhenUsingYourOwnFork) section of [Pull Request Validation Builds for the Testing Service]({help_url}#PullRequestValidationBuildsfortheTestingService-WhenUsingYourOwnFork).
"""

COULD_NOT_FETCH_REPO = """
Not running the OCI testing service tests.

Could not fetch the source repository for the {sdk_name} pull request: {sdk_pr}.

{author_text}Please make sure you have added the ['Key for OCI testing service on-PR build access']({help_url}#PullRequestValidationBuildsfortheTestingService-WhenUsingYourOwnFork) to your {sdk_name} repository's access keys.

**Note: The access key has to be added to {sdk_name} repository, not the OCI testing service repository!**

After you have made sure access rights and keys are correct, restart the [TeamCity build]({tc_link}) by asking an SDK/CLI team member on #oci_public_sdks to click on the 'Run' button.

For more information, see the [When Using Your Own Fork]({help_url}#PullRequestValidationBuildsfortheTestingService-WhenUsingYourOwnFork) section of [Pull Request Validation Builds for the Testing Service]({help_url}#PullRequestValidationBuildsfortheTestingService-WhenUsingYourOwnFork).
"""

USING_SDK_PR = """
Since this change is targeting the `master` branch, the build job is using another {sdk_name} pull request to build the {sdk_name}.

The pull request chosen to provide the code for the {sdk_name} build is:

{sdk_pr}

If this is not the right pull request, please see [My master testing service PR picks the wrong master Java SDK branch]({help_url}#PullRequestValidationBuildsfortheTestingService-MymastertestingservicePRpicksthewrongmasterJavaSDKbranch).

This build is now progressing in the [TeamCity build]({tc_link}). For more information, see the [When Using Your Own Fork]({help_url}#PullRequestValidationBuildsfortheTestingService-WhenUsingYourOwnFork) section of [Pull Request Validation Builds for the Testing Service]({help_url}#PullRequestValidationBuildsfortheTestingService-WhenUsingYourOwnFork).
"""

USING_MASTER = """
Building this change against the {sdk_name} in the `master` branch.

This change is targeting the `master` branch, the build job is using another {sdk_name} pull request to build the {sdk_name}. The pull request chosen to provide the code for the {sdk_name} build is:

{sdk_pr}

If this is not the right pull request, please see [My master testing service PR picks the wrong master Java SDK branch]({help_url}#PullRequestValidationBuildsfortheTestingService-MymastertestingservicePRpicksthewrongmasterJavaSDKbranch).

Since that {sdk_name} pull request has already been merged, we are building the testing service in this pull request against the {sdk_name} in the `master` branch.

This build is now progressing in the [TeamCity build]({tc_link}). For more information, see the [When Using Your Own Fork]({help_url}#PullRequestValidationBuildsfortheTestingService-WhenUsingYourOwnFork) section of [Pull Request Validation Builds for the Testing Service]({help_url}#PullRequestValidationBuildsfortheTestingService-WhenUsingYourOwnFork).
"""


# Target the master branch. The message can contain the {sdk_branch} placeholder.
# The program will exit inside this function.
def target_master(output_file, message):
    sdk_branch = "master"
    sdk_repo.git.fetch("origin")
    sdk_repo.git.checkout(sdk_branch)
    printv(message.format(sdk_branch=sdk_branch))

    print(sdk_branch)

    if output_file:
        with open(output_file, 'w') as f:
            f.write(sdk_branch)

    sys.exit(0)


#
# Parameters variable set up
#
parser = argparse.ArgumentParser(description='Determine the matching SDK branch, and switch to it if it exists.')
parser.add_argument('--build-id', required=True, help="The TeamCity build id for the build that is running this script. This is used to update the relevant Bitbucket PRs with links to the TeamCity build")
parser.add_argument('--build-branch', required=True, help="The value of the teamcity.build.branch variable")
parser.add_argument('--sdk-path', required=True, help="Path to the root directory of the SDK")
parser.add_argument('--sdk-name', default="Java SDK", required=False, help="Name of the SDK (if not specified, uses 'Java SDK')")
parser.add_argument('--sdk-project', default="java-sdk", required=False, help="Project for the SDK (if not specified, uses 'java-sdk')")
parser.add_argument('-o', '--output', required=False, help='Output file')
parser.add_argument('-v', '--verbose', action='store_true', default=False, required=False, help='Verbose output')
parser.add_argument('--dry-run', action='store_true', default=False, required=False, help='Dry-run')

args = parser.parse_args()

ocits_shared.dry_run = args.dry_run

if args.verbose:
    ocits_shared.verbose = True
    shared.bitbucket_utils.verbose = True

setup_bitbucket(args)

sdk_repo = Repo.init(args.sdk_path)

pr_id = None
try:
    # If the teamcity.build.branch variable is just something like "1234", then this is a
    # validation build for pull request "1234"
    pr_id = int(args.build_branch)
except ValueError:
    # Only print this when using verbose, since we want the output be the target branch.
    printv("Not a pull request validation build.")
    sys.exit(2)

tc_link = TC_URL.format(build_id=args.build_id)

pr = get_pullrequest("SDK", "oci-testing-service", pr_id)
printv(pr.text)

json = pr.json()

if 'title' in json:
    title = json['title']
else:
    title = ""
if 'description' in json:
    description = json['description']
else:
    description = ""

author_text = ""
json = pr.json()
if json['author'] and json['author']['user'] and json['author']['user']['name']:
    author_text = "@{name}: ".format(name=json['author']['user']['name'])

sdk_pr_url = get_master_javasdk_pr_url(description)

if sdk_pr_url:
    print("Using URL from PR description: {}".format(sdk_pr_url))
    m = re.search("^.*bitbucket.*/projects/([^/]*)/repos/([^/]*)/pull-requests/([0-9]*).*$", sdk_pr_url)
    if m:
        sdk_pr_id = m.group(3)
    else:
        raise ValueError("Unknown PR URL: {}".format(sdk_pr_url))
else:
    tickets, no_dexreq_marker = get_dexreq_tickets(pr, tc_link)

    if no_dexreq_marker:
        target_master(args.output, "Since this is NO-DEXREQ, should target tip of '{sdk_branch}'")
    else:
        printv("Tickets: {}".format(", ".join(tickets)))

        sdk_pr_candidates = []

        for issue_key in tickets:
            issue = get_jira_issue(issue_key)
            pull_requests = get_pull_requests_for_issue(issue)
            master_sdk_pull_request = filter_pull_requests(pull_requests, args.sdk_project, "master")

            for pr in master_sdk_pull_request:
                sdk_pr_candidates.append((pr['id'].replace('#', '')))

        if not sdk_pr_candidates:
            print("Could not find {} PR targetting master in any of the PRs referenced in the DEXREQ tickets {}".format(args.sdk_name, ", ".join(tickets)))

            text = MATCHING_SDK_BRANCH_NOT_FOUND.format(
                tc_link=tc_link,
                sdk_name=args.sdk_name,
                tickets=", ".join(tickets),
                ticket_sop="ticket" if len(tickets) == 1 else "tickets",
                help_url=HELP_URL,
                author_text=author_text)

            if not ocits_shared.dry_run:
                make_general_comment("SDK", "oci-testing-service", pr_id, text)
            else:
                print("DRY-RUN: Not making BitBucket comment\n{}".format(text))

            # Don't fail the build
            sys.exit(1)

        latest_candidate = 0
        sdk_pr_url = None
        for candidate in sdk_pr_candidates:
            num = int(candidate)
            url = "https://bitbucket.oci.oraclecorp.com/projects/SDK/repos/{sdk_project}/pull-requests/{num}/overview".format(sdk_project=args.sdk_project, num=num)
            printv("Candidate PR: {}".format(url))
            if num > latest_candidate:
                latest_candidate = num
                sdk_pr_url = url

        sdk_pr_id = str(latest_candidate)
        printv("Choosing PR: {}".format(sdk_pr_id))

sdk_pr = get_pullrequest("SDK", args.sdk_project, sdk_pr_id)
printv(sdk_pr.text)

already_merged = sdk_pr.json()['state'] == 'MERGED'
if already_merged:
    already_merged = True
    text = USING_MASTER.format(
        tc_link=tc_link,
        sdk_name=args.sdk_name,
        sdk_pr=sdk_pr_url,
        help_url=HELP_URL)
else:
    text = USING_SDK_PR.format(
        tc_link=tc_link,
        sdk_name=args.sdk_name,
        sdk_pr=sdk_pr_url,
        help_url=HELP_URL)

make_general_comment("SDK", "oci-testing-service", pr_id, text)

if already_merged:
    target_master(args.output, "SDK PR has already been merged, should target tip of '{sdk_branch}'")

remote_clone_ssh_url = get_pr_source_clone_ssh_url(sdk_pr)
printv("Remote Clone SSH URL: {}".format(remote_clone_ssh_url))

if not remote_clone_ssh_url:
    print("Could not determine clone URL")

    text = COULD_NOT_DETERMINE_CLONE_URL.format(
        tc_link=tc_link,
        sdk_name=args.sdk_name,
        sdk_pr=sdk_pr_url,
        permissions_url=get_repo_permissions_url(sdk_pr),
        help_url=HELP_URL,
        author_text=author_text)

    if not ocits_shared.dry_run:
        make_general_comment("SDK", "oci-testing-service", pr_id, text)
    else:
        print("DRY-RUN: Not making BitBucket comment\n{}".format(text))

    # Don't fail the build
    sys.exit(1)

remote_sdk_branch = get_pr_source_branch(sdk_pr)
printv("Remote SDK branch: {}".format(remote_sdk_branch))

# git remote add fork ssh://git@bitbucket.oci.oraclecorp.com:7999/~XXX/java-sdk.git
sdk_repo.git.remote("add", "fork", remote_clone_ssh_url)

try:
    # git fetch fork
    sdk_repo.git.fetch("fork")
except GitCommandError:
    print("Could not fetch the remote repo")

    text = COULD_NOT_FETCH_REPO.format(
        tc_link=tc_link,
        sdk_name=args.sdk_name,
        sdk_pr=sdk_pr_url,
        help_url=HELP_URL,
        author_text=author_text)

    if not ocits_shared.dry_run:
        make_general_comment("SDK", "oci-testing-service", pr_id, text)
    else:
        print("DRY-RUN: Not making BitBucket comment\n{}".format(text))

    # Don't fail the build
    sys.exit(1)

# git branch fork_generated-auto-public-JavaSDK-DEXREQ-NNN-YYYY-MM-DD-HH-mm-ss fork/fork_generated-auto-public-JavaSDK-DEXREQ-NNN-YYYY-MM-DD-HH-mm-ss
sdk_branch = "fork_{}".format(remote_sdk_branch)
sdk_repo.git.branch(sdk_branch, "fork/{}".format(remote_sdk_branch))

# git checkout fork_generated-auto-public-JavaSDK-DEXREQ-248-2019-01-23-19-08-17
sdk_repo.git.checkout(sdk_branch)

print(sdk_branch)

if args.output:
    with open(args.output, 'w') as f:
        f.write(sdk_branch)
