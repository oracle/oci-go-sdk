from __future__ import print_function
import argparse
import os
import ssl
import sys
import urllib3
from glob import glob
import ocits_shared
from ocits_shared import HELP_URL, TC_URL, JACOCO_XML_PATH, get_dexreq_tickets_from_text, printv, parse_xml, extract_package_codecov_data_from_reportxml, get_class_tag_breakdown_from_ops, get_class_from_test_name, get_package_names_from_description, get_determined_text

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../..'))

from shared.bitbucket_utils import setup_bitbucket, get_pullrequest, make_general_comment  # noqa: ignore=F402
import shared.bitbucket_utils  # noqa: ignore=F402
import util  # noqa: ignore=F402
import config  # noqa: ignore=F402

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
if (not os.environ.get('PYTHONHTTPSVERIFY', '') and getattr(ssl, '_create_unverified_context', None)):
    ssl._create_default_https_context = ssl._create_unverified_context

#
# Warning messages
#

NO_TESTS_RUN = """
No tests could be run. The automation attempted to run matching tests, but none were found.

{author_text}Please make sure that you have specified the correct 'Issue Routing Tag' in the referenced DEXREQ tickets (or in the description of this pull request). Then restart the [TeamCity build]({tc_link}) by asking an SDK/CLI team member on #oci_public_sdks to click on the 'Run' button.

For more information, see [Pull Request Validation Builds for the Testing Service]({help_url}).
"""

NO_TESTS_MATCHED = """
The automation tried to run tests, but no tests matched the 'Issue Routing Tag'{issue_routing_tag_text}{test_classes_text}

{author_text}Please make sure that you have specified the correct 'Issue Routing Tag' in the referenced DEXREQ tickets (or in the description of this pull request). Then restart the [TeamCity build]({tc_link}) by asking an SDK/CLI team member on #oci_public_sdks to click on the 'Run' button.

For more information, see [Pull Request Validation Builds for the Testing Service]({help_url}).
"""

PARTIAL_TEXT_TESTS_SKIPPED = """


The following {skipped_tests_sop} were skipped because they were not whitelisted:

{skipped}
"""

ALL_TESTS_SKIPPED = """
{summary_text}

{author_text}No tests were run, because no tests were whitelisted. Please make sure your tests are whitelisted.{skipped_text}

More details are available in the log of the [TeamCity build]({tc_link}).

For more information, see [Pull Request Validation Builds for the Testing Service]({help_url}).
"""

TESTS_SUCCESSFUL = """
{summary_text}

All tests were successful. The following {successful_tests_sop} succeeded:

{successful}{skipped_text}

More details are available in the log of the [TeamCity build]({tc_link}).

For more information, see [Pull Request Validation Builds for the Testing Service]({help_url}).{codecov}
"""

PARTIAL_TEXT_TESTS_SUCCEEDED = """


The following {successful_tests_sop} succeeded:
{successful}
"""

TESTS_FAILED = """
{summary_text}

{author_text}Please make sure the tests pass. The following {failed_tests_sop} failed:
{failed}{succeeded_text}{skipped_text}

More details are available in the log of the [TeamCity build]({tc_link}).

For more information, see [Pull Request Validation Builds for the Testing Service]({help_url}).{codecov}
"""

BASELINE_CODECOV_LINK = """

[Baseline code coverage report]({baseline_report_url})
"""

CODECOV_PACKAGE_TEXT = """
##### `{package}`:
- Instruction coverage = {instr}
- Branch coverage = {branch}
- Line coverage = {line} {baseline_codecov_link}


"""

CODECOV_TEXT = """


###### Code Coverage Report:

{packagetext}

Full report available at [code coverage report](https://teamcity.oci.oraclecorp.com/repository/download/Sdk_OCI_Testing_Service_BuildPullRequest/{build_id}:id/jacocoreport/index.html).
"""

CODECOV_FAILURE_TEXT = """
"Build failed to meet the code coverage criteria (currently set to {p} percent.)"
"""


def get_summary_text(failed, skipped, successful):
    if not failed and not skipped and not successful:
        return "No tests were run."

    total_count = len(failed) + len(skipped) + len(successful)

    text = "The {test_tests_sop} finished. ".format(
        test_tests_sop="test" if total_count == 1 else "tests")
    previous = False

    if failed:
        text = text + "{count} {test_tests_sop} failed".format(
            count=len(failed),
            test_tests_sop="test" if len(failed) == 1 else "tests")
        previous = True
    if skipped:
        text = text + "{comma}{count} {test_tests_sop} were skipped".format(
            comma=", " if previous else "",
            count=len(skipped),
            test_tests_sop="test" if len(skipped) == 1 else "tests")
        previous = True
    if successful:
        text = text + "{comma}{count} {test_tests_sop} were successful".format(
            comma=", " if previous else "",
            count=len(successful),
            test_tests_sop="test" if len(successful) == 1 else "tests")
        previous = True

    text = text + "."

    return text


def get_test_xml_files(java_sdk_path):
    junit_xml_dir = os.path.join(java_sdk_path, "bmc-integtests/target/junit-xml")
    printv("Looking for TEST-*.xml files in {}".format(junit_xml_dir))
    return [y for x in os.walk(junit_xml_dir) for y in glob(os.path.join(x[0], 'TEST-*.xml'))]


def process_junit_xml(successful, failed, pom):
    name = pom.getroot().get('name').encode('utf8')
    errors = int(pom.getroot().get('errors').encode('utf8'))
    failures = int(pom.getroot().get('failures').encode('utf8'))

    print("errors: {} failures: {} -- {}".format(errors, failures, name))

    if errors + failures == 0:
        successful.append(name)
        return True

    # Something failed
    short_message = None
    message = None
    exception_type = None
    for test_case in pom.getroot().iter("testcase"):
        for error in test_case.iter("error"):
            message_node = error.get('message')
            if message_node:
                short_message = error.get('message').encode('utf8')
            else:
                short_message = "message attribute missing"
                print("message attribute missing")
            exception_type = error.get('type').encode('utf8')
            message = error.text.encode('utf8')

    print("{} -- {}: {}".format(name, exception_type, short_message))
    failed.append((name, errors, failures, short_message, message, exception_type))

    return False


def get_just_stack_trace(message):
    text = ""
    skip = True

    for line in message.split("\n"):
        if text and not skip:
            text = text + "\n"

        if line.startswith("\tat "):
            skip = False

        if not skip:
            text = text + line

    return text.strip()


def simplify_stack_trace(stack_trace):
    text = ""
    skip = False

    for line in stack_trace.split("\n"):
        if text and not skip:
            text = text + "\n"

        if line.startswith("\tat org.glassfish.jersey.server.model.internal."):
            skip = True

        if line.startswith("Caused by"):
            skip = False

        if not skip:
            text = text + line

    return text


def get_failed_text(failed):
    text = ""
    for t, e, f, sm, m, et in failed:
        if text:
            text = text + "\n"

        # Ignore unimportant stack trace in short message
        sm_ignore_pos = sm.find("\n\tat sun.reflect.GeneratedMethodAccessor54.invoke")
        if sm_ignore_pos:
            sm = sm[:sm_ignore_pos]

        sm = simplify_stack_trace(sm)
        m = simplify_stack_trace(m)

        sm = sm.replace(get_just_stack_trace(m), "")

        text = text + "- `{}`: {}\n\nException:\n```{}```\n".format(t, sm, m)

    return text


# get the persisted previous build codecov numbers from object storage for comparison
# get the baseline code coverage from branch, get the current code coverage from xml report
# compare and report
def get_codecov_text(target_branch, ops_file, code_coverages):
    if not ops_file:
        return ""
    # if target branch is not provided, assume no codecov reporting needed for backard compatibility
    if not target_branch:
        return ""
    # baseline_json = get_codecov_baseline(target_branch)
    xml_file_path = JACOCO_XML_PATH.format(java_sdk_path=args.java_sdk_path)
    report_xml = parse_xml(xml_file_path)
    package_text = ""
    codecov_map = get_class_tag_breakdown_from_ops(ops_file)
    service_data = extract_package_codecov_data_from_reportxml(report_xml, ops_file)
    # baseline_data = extract_package_codecov_data_from_baseline(baseline_json, ops_file)
    for tp in codecov_map:
        class_name = get_class_from_test_name(tp[0])
        service_item = list(filter(lambda class_tag: class_tag["class"] == class_name and class_tag["tag"] == tp[1], service_data))
        # baseline_item = list(filter(lambda class_tag: class_tag["class"] == class_name and class_tag["tag"] == tp[1], baseline_data)) 
        # disable baseline comparison for now
        package_text += get_codecov_package_text(service_item[0], None, "class: {c}, tag: {t}".format(c=class_name, t=tp[1]), code_coverages) 
    return CODECOV_TEXT.format(packagetext=package_text, build_id=args.build_id)


def get_codecov_lineitem_text(covered, missed):
    baseline = covered + missed
    current = (covered * 100 / baseline) if baseline != 0 else 0
    return str(current) + "% "


def get_codecov_lineitem_text_with_baseline(covered, missed, baseline_covered, baseline_missed):
    current = get_percentage(covered, missed)
    baseline = get_percentage(baseline_covered, baseline_missed)
    delta = ""
    if current == baseline:
        delta = "**(no change)**"
    elif current > baseline:
        delta = "**( +" + str(current - baseline) + "%)**"
    else:
        delta = "**( -" + str(baseline - current) + "%)**"
    return str(current) + "% " + delta


def get_percentage(covered, missed):
    denominator = covered + missed
    return int(covered * 100 / denominator) if denominator != 0 else 100


# we get current and baseline structure and we compare corresponding values
def get_codecov_package_text(current_data, baseline_data, package_label, code_coverages):
    current = current_data["data"]
    code_coverages.append(get_percentage(current["coveredInstructions"], current["missedInstructions"]))
    if baseline_data:
        baseline = baseline_data["data"]
        return CODECOV_PACKAGE_TEXT.format(
            package=package_label,
            instr=get_codecov_lineitem_text_with_baseline(current["coveredInstructions"], current["missedInstructions"], baseline["coveredInstructions"], baseline["missedInstructions"]),
            branch=get_codecov_lineitem_text_with_baseline(current["coveredBranches"], current["missedBranches"], baseline["coveredBranches"], baseline["missedBranches"]),
            line=get_codecov_lineitem_text_with_baseline(current["coveredLines"], current["missedLines"], baseline["coveredLines"], baseline["missedLines"]),
            baseline_codecov_link=BASELINE_CODECOV_LINK.format(baseline_report_url=baseline["url"]))
    else:
        return CODECOV_PACKAGE_TEXT.format(
            package=package_label,
            instr=get_codecov_lineitem_text(current["coveredInstructions"], current["missedInstructions"]),
            branch=get_codecov_lineitem_text(current["coveredBranches"], current["missedBranches"]),
            line=get_codecov_lineitem_text(current["coveredLines"], current["missedLines"]),
            baseline_codecov_link="")


#
# Parameters variable set up
#
parser = argparse.ArgumentParser(description='Report test results.')
parser.add_argument('--build-id', required=True, help="The TeamCity build id for the build that is running this script. This is used to update the relevant Bitbucket PRs with links to the TeamCity build")
parser.add_argument('--build-branch', required=True, help="The value of the teamcity.build.branch variable")
parser.add_argument('-t', '--target-branch', required=False, help="The target branch - master or preview")
parser.add_argument('--java-sdk-path', required=True, help="Path to the root directory of the Java SDK")
parser.add_argument('-v', '--verbose', action='store_true', default=False, required=False, help='Verbose output')
parser.add_argument('--dry-run', action='store_true', default=False, required=False, help='Dry-run, do not post to PR')
parser.add_argument('--codecov-operations-file', required=True, help='list of operations to consider for code coverage')
parser.add_argument('--expected-coverage-percentage', type=int, required=False, help='fail the PR if code coverage numbers drop')

args = parser.parse_args()
setup_bitbucket(args)

ocits_shared.dry_run = args.dry_run
shared.bitbucket_utils.dry_run = args.dry_run

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
    sys.exit(2)

pr = get_pullrequest("SDK", "oci-testing-service", pr_id)
printv(pr.text)
pr_json = pr.json()
tc_link = TC_URL.format(build_id=args.build_id)

if 'title' in pr_json:
    title = pr_json['title']
else:
    title = ""
if 'description' in pr_json:
    description = pr_json['description']
else:
    description = ""

author_text = ""
pr_json = pr.json()
if pr_json['author'] and pr_json['author']['user'] and pr_json['author']['user']['name']:
    author_text = "@{name}: ".format(name=pr_json['author']['user']['name'])

package_names_from_description = get_package_names_from_description(description)
tickets, no_dexreq_marker = get_dexreq_tickets_from_text(title + "\n" + description)

if no_dexreq_marker:
    sys.exit(0)

test_xml_files = get_test_xml_files(args.java_sdk_path)

if not test_xml_files:
    text = NO_TESTS_RUN.format(
        tc_link=tc_link,
        help_url=HELP_URL,
        author_text=author_text)

    make_general_comment("SDK", "oci-testing-service", pr_id, text)

    sys.exit(1)

successful = []
failed = []

for f in test_xml_files:
    print("processing " + f)
    test_xml = parse_xml(f)
    process_junit_xml(successful, failed, test_xml)

successful_text = "\n".join("- `{}`".format(t) for t in successful)

# Filter out skipped tests (org.junit.AssumptionViolatedException) -- not whitelisted
skipped = []
for t, e, f, sm, m, et in failed:
    if et == 'org.junit.AssumptionViolatedException':
        # This is a "no tests matched the issue routing tag" error
        skipped.append((t, e, f, sm, m, et))

skipped_text = ""
if skipped:
    skipped_text = PARTIAL_TEXT_TESTS_SKIPPED.format(
        skipped_tests_sop="test" if len(skipped) == 1 else "tests",
        skipped="\n".join("- `{}`".format(t) for t, e, f, sm, m, et in skipped))

# Remove errors that were skipped
failed[:] = [x for x in failed if x not in skipped]

# Find errors due to no matching tests
no_matching_tests = []
for t, e, f, sm, m, et in failed:
    if t == 'initializationError(org.junit.runner.manipulation.Filter)' and sm.startswith('No tests found matching ProjectFilter'):
        # This is a "no tests matched the issue routing tag" error
        no_matching_tests.append((t, e, f, sm, m, et))

# Remove errors due to no matching tests from list of regular failures
failed[:] = [x for x in failed if x not in no_matching_tests]

if not failed and not successful and not skipped:
    # All failures were "no tests matched the issue routing tags" -- there were no actual test failures
    issue_routing_tag_text = ""
    if "ISSUE_ROUTING_TAG" in os.environ:
        issue_routing_tags = os.environ.get("ISSUE_ROUTING_TAG").split(",")
        issue_routing_tag_text = " {}".format(", ".join("`{}`".format(t) for t in issue_routing_tags))

    test_classes_text = ""
    if "TEST_CLASSES" in os.environ:
        test_classes = os.environ.get("TEST_CLASSES").split(",")
        determined_text = get_determined_text(tickets, package_names_from_description, test_classes)
        test_classes_text = " in the following {class_sop}:\n\n{test_classes}\n\n{determined_text}".format(
            class_sop="class" if len(test_classes) == 1 else "classes",
            test_classes="\n".join("- `{}`".format(c) for c in test_classes),
            determined_text=determined_text)

    text = NO_TESTS_MATCHED.format(
        tc_link=tc_link,
        issue_routing_tag_text=issue_routing_tag_text,
        test_classes_text=test_classes_text,
        help_url=HELP_URL,
        author_text=author_text)

    make_general_comment("SDK", "oci-testing-service", pr_id, text)

    sys.exit(1)
elif failed:
    failed_text = get_failed_text(failed)
    succeeded_text = ""
    if successful:
        succeeded_text = PARTIAL_TEXT_TESTS_SUCCEEDED.format(
            successful=successful_text,
            successful_tests_sop="test" if len(successful) == 1 else "tests")
    code_coverages = []
    text = TESTS_FAILED.format(
        tc_link=tc_link,
        failed=failed_text,
        failed_tests_sop="test" if len(failed) == 1 else "tests",
        succeeded_text=succeeded_text,
        skipped_text=skipped_text,
        help_url=HELP_URL,
        summary_text=get_summary_text(failed, skipped, successful),
        author_text=author_text,
        codecov=get_codecov_text(args.target_branch, args.codecov_operations_file, code_coverages))

    make_general_comment("SDK", "oci-testing-service", pr_id, text)

    sys.exit(1)
elif not failed and not successful and skipped:
    text = ALL_TESTS_SKIPPED.format(
        tc_link=tc_link,
        skipped_text=skipped_text,
        help_url=HELP_URL,
        summary_text=get_summary_text(failed, skipped, successful),
        author_text=author_text)

    make_general_comment("SDK", "oci-testing-service", pr_id, text)

    sys.exit(1)
else:
    code_coverages = []
    text = TESTS_SUCCESSFUL.format(
        tc_link=tc_link,
        successful=successful_text,
        successful_tests_sop="test" if len(successful) == 1 else "tests",
        skipped_text=skipped_text,
        help_url=HELP_URL,
        summary_text=get_summary_text(failed, skipped, successful),
        codecov=get_codecov_text(args.target_branch, args.codecov_operations_file, code_coverages))
    fail_on_code_coverage = False
    print(code_coverages)
    if args.expected_coverage_percentage is not None:
        for code_coverage in code_coverages:
            if code_coverage < args.expected_coverage_percentage:
                fail_on_code_coverage = True
                break
    make_general_comment("SDK", "oci-testing-service", pr_id, text)
    if fail_on_code_coverage:
        make_general_comment("SDK", "oci-testing-service", pr_id, CODECOV_FAILURE_TEXT.format(p=args.expected_coverage_percentage))
        sys.exit(1)
