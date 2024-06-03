from __future__ import print_function
import argparse
import os
import ssl
import sys
import urllib3
import ocits_shared
from ocits_shared import HELP_URL, TC_URL, printv, get_dexreq_tickets, get_issue_routing_info_tag, get_issue_routing_info_tag_from_description, get_package_names_from_description, get_limit_text, get_determined_text

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../..'))

from shared.bitbucket_utils import setup_bitbucket, get_pullrequest, get_pr_target_branch, make_general_comment  # noqa: ignore=F402
import shared.bitbucket_utils  # noqa: ignore=F402
import util  # noqa: ignore=F402
import config  # noqa: ignore=F402

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
if (not os.environ.get('PYTHONHTTPSVERIFY', '') and getattr(ssl, '_create_unverified_context', None)): 
    ssl._create_default_https_context = ssl._create_unverified_context


#
# Warning messages
#


NO_TAGS_FOUND_RUNNING_ALL = """
Could not find any issue routing tags, because even though DEXREQ tickets were referenced, none of them had the 'Issue Routing Tag' field filled out.

The pull request description also didn't contain an issue routing tag in the form of `[IssueRoutingInfo.tag=sometag]`.

This means the automation will run all tests belonging to the spec specified in the DEXREQ tickets. That may take a long time.

{limit_text} Then restart the [TeamCity build]({tc_link}) by asking an SDK/CLI team member on #oci_public_sdks to click on the 'Run' button.

Note: The 'Issue Routing Tag' is the same value you have chosen in your spec for the `x-obmcs-issue-routing-tag` annotation. Only tests for operations with matching issue routing tag will be run ([more info](https://confluence.oci.oraclecorp.com/display/DEX/Issue+Routing)).

About to use the OCI testing service to run all tests in the following {class_sop}:

{test_classes}

{determined_text}

The testing progress can be monitored in the log of the [TeamCity build]({tc_link}).

For more information, see [Pull Request Validation Builds for the Testing Service]({help_url}).
"""


NO_DEXREQ_NO_TAGS_FOUND_RUNNING_ALL = """
Could not find any issue routing tags, because no DEXREQ tickets were referenced. The pull request description also didn't contain an issue routing tag in the form of `[IssueRoutingInfo.tag=sometag]`.

This is not a problem, but it means the automation will run all tests belonging to the modules specified in the pull request description in the form of `[RunTestsForModule=xyz]`. That may take a long time.

{limit_text} Then restart the [TeamCity build]({tc_link}) by asking an SDK/CLI team member on #oci_public_sdks to click on the 'Run' button.

Note: The 'Issue Routing Tag' is the same value you have chosen in your spec for the `x-obmcs-issue-routing-tag` annotation. Only tests for operations with matching issue routing tag will be run ([more info](https://confluence.oci.oraclecorp.com/display/DEX/Issue+Routing)).

About to use the OCI testing service to run all tests in the following {class_sop}:

{test_classes}

{determined_text}

The testing progress can be monitored in the log of the [TeamCity build]({tc_link}).

For more information, see [Pull Request Validation Builds for the Testing Service]({help_url}).
"""


RUNNING_THE_FOLLOWING_TESTS = """
About to use the OCI testing service to run tests with matching 'Issue Routing Tag' {issue_routing_tags} in the following {class_sop}:

{test_classes}

{determined_text}The testing progress can be monitored in the log of the [TeamCity build]({tc_link}).

For more information, see [Pull Request Validation Builds for the Testing Service]({help_url}).
"""


#
# Parameters variable set up
#
parser = argparse.ArgumentParser(description='Determine the issue routing info tag.')
parser.add_argument('--build-id', required=True, help="The TeamCity build id for the build that is running this script. This is used to update the relevant Bitbucket PRs with links to the TeamCity build")
parser.add_argument('--build-branch', required=True, help="The value of the teamcity.build.branch variable")
parser.add_argument('-o', '--output', required=False, help='Output file')
parser.add_argument('-v', '--verbose', action='store_true', default=False, required=False, help='Verbose output')
parser.add_argument('--dry-run', action='store_true', default=False, required=False, help='Dry-run')

args = parser.parse_args()
setup_bitbucket(args)

ocits_shared.dry_run = args.dry_run

if args.verbose:
    ocits_shared.verbose = True
    shared.bitbucket_utils.verbose = True

pr_id = None
try:
    # If the teamcity.build.branch variable is just something like "1234", then this is a 
    # validation build for pull request "1234"
    pr_id = int(args.build_branch)
except ValueError:
    # Only print this when using verbose, since we want the output be the target branch.
    printv("Not a pull request validation build.")
    sys.exit(1)

test_classes = []
if "TEST_CLASSES" in os.environ:
    test_classes = os.environ.get("TEST_CLASSES").split(",")
else:
    print("TEST_CLASSES environment variable has to be set (comma-separated list of Java class names)")
    sys.exit(1)

if not test_classes:
    # Empty -- "NO-DEXREQ" was set and no [RunTestsForModule=xyz] was set
    # This os ok.
    sys.exit(0)

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

tc_link = TC_URL.format(build_id=args.build_id)
tickets, no_dexreq_marker = get_dexreq_tickets(pr, tc_link)
package_names_from_description = get_package_names_from_description(description)

if no_dexreq_marker and not package_names_from_description:
    # If the user had "NO-DEXREQ" in the PR and didn't manually add modules using [RunTestsForModule=xyz], no tests will run
    # No point in looking for an issue routing tag.
    sys.exit(0)

tags_set = set([])
for issue_key in tickets:
    tags = get_issue_routing_info_tag(issue_key)

    if tags:
        for tag in tags:
            tags_set.add(tag)

tags = get_issue_routing_info_tag_from_description(description)
for tag in tags:
    tags_set.add(tag)

limit_text = get_limit_text(tickets, package_names_from_description)
determined_text = get_determined_text(tickets, package_names_from_description, test_classes)

if not tags_set and not no_dexreq_marker:
    printv("No issue routing info tags found!")

    text = NO_TAGS_FOUND_RUNNING_ALL.format(
        tc_link=tc_link,
        dexreq_tickets=", ".join(tickets),
        ticket_sop="ticket" if len(tickets) == 1 else "tickets",
        test_classes="\n".join("- `{}`".format(c) for c in test_classes),
        class_sop="class" if len(test_classes) == 1 else "classes",
        issue_routing_tags=", ".join("`{}`".format(t) for t in tags_set),
        limit_text=limit_text,
        determined_text=determined_text,
        help_url=HELP_URL)
elif not tags_set and no_dexreq_marker:
    printv("No issue routing info tags found, but NO-DEXREQ marker exists!")

    text = NO_DEXREQ_NO_TAGS_FOUND_RUNNING_ALL.format(
        tc_link=tc_link,
        dexreq_tickets=", ".join(tickets),
        ticket_sop="ticket" if len(tickets) == 1 else "tickets",
        test_classes="\n".join("- `{}`".format(c) for c in test_classes),
        class_sop="class" if len(test_classes) == 1 else "classes",
        issue_routing_tags=", ".join("`{}`".format(t) for t in tags_set),
        limit_text=limit_text,
        determined_text=determined_text,
        help_url=HELP_URL)
else:
    tags_str = ",".join(str(s) for s in tags_set)
    print(tags_str)

    if args.output:
        f = open(args.output, "w")
        f.write(tags_str)

    text = RUNNING_THE_FOLLOWING_TESTS.format(
        tc_link=tc_link,
        dexreq_tickets=", ".join(tickets),
        ticket_sop="ticket" if len(tickets) == 1 else "tickets",
        test_classes="\n".join("- `{}`".format(c) for c in test_classes),
        class_sop="class" if len(test_classes) == 1 else "classes",
        this_class_was_these_classes_were="This class was" if len(test_classes) == 1 else "These classes were",
        issue_routing_tags=", ".join("`{}`".format(t) for t in tags_set),
        help_url=HELP_URL,
        determined_text=determined_text)

if not ocits_shared.dry_run:
    make_general_comment("SDK", "oci-testing-service", pr_id, text)
else:
    print("DRY-RUN: Not making BitBucket comment\n{}".format(text))
