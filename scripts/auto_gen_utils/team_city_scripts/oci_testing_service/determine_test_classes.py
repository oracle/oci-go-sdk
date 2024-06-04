from __future__ import print_function
import argparse
import os
import ssl
import sys
import urllib3
import re
import xml.etree.ElementTree as ET
from glob import glob

import ocits_shared
from ocits_shared import HELP_URL, TC_URL, printv, get_dexreq_tickets, get_group_and_artifact_ids, parse_xml, get_package_names_from_description

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


ns = {"ns":"http://maven.apache.org/POM/4.0.0"}

# allow default namespace for output, dont print ns0: prefixes everywhere
ET.register_namespace('',"http://maven.apache.org/POM/4.0.0")


#
# Warning messages
#


NO_TESTS_FOUND_BECAUSE_UNKNOWN_SPEC = """
Not running the OCI testing service tests.

The referenced DEXREQ {ticket_sop} {dexreq_tickets} {ticket_has_have} 'Spec Group Id' and 'Spec Artifact Id' set ({ids_list}), but no spec with those ids was not found in the Java SDK.

{author_text}Please ensure that the 'Spec Group Id' and 'Spec Artifact Id' fields are filled out correctly, and that the SDK for {that_spec_those_specs} has been generated. Then restart the [TeamCity build]({tc_link}) by asking an SDK/CLI team member on #oci_public_sdks to click on the 'Run' button.

For more information, see [Pull Request Validation Builds for the Testing Service]({help_url}).
"""


NO_TESTS_FOUND_BECAUSE_NO_TEST_CLASSES = """
Not running the OCI testing service tests.

The referenced DEXREQ {ticket_sop} {dexreq_tickets} {ticket_has_have} 'Spec Group Id' and 'Spec Artifact Id' set, which resulted in the Java package {name_sop} {java_package_names}. Unfortunately, {this_package_these_packages} did not contain `*AutoTest.java` files.

{author_text}Please ensure that the 'Spec Group Id' and 'Spec Artifact Id' fields are filled out correctly, and that the SDK for that spec has been generated. Then restart the [TeamCity build]({tc_link}) by asking an SDK/CLI team member on #oci_public_sdks to click on the 'Run' button.

For more information, see [Pull Request Validation Builds for the Testing Service]({help_url}).
"""


NO_TESTS_FOUND_BUT_NO_DEXREQ = """
Not running the OCI testing service tests.

No test classes were found, because no DEXREQ tickets were referenced in the pull request title or description. But since the `{no_dexreq_marker}` marker was present in the description, this is not a problem; the tests just won't get run.

If you do want to run tests, you can specify which modules should have their tests run by adding annotations in the form of `[RunTestsForModule=xyz]` to the pull request description. You may also want to limit the tests that then get run by specifying the issue routing tag in the form of `[IssueRoutingInfo.tag=sometag]` in the pull request description,

For more information, see [Pull Request Validation Builds for the Testing Service]({help_url}).
"""


def get_package_name_from_pom_file(group_id, artifact_id, java_sdk_path):
    codegen_dir = os.path.join(java_sdk_path, "bmc-codegen")
    pom_files = [y for x in os.walk(codegen_dir) for y in glob(os.path.join(x[0], 'pom.xml'))]
    for ldr_path in pom_files:
        pom = parse_xml(ldr_path)
        # Find all the places where <properties><codegen.artifactory.groupId> and <properties><codegen.artifactory.artifactId> both match
        xpath = './/ns:properties[ns:codegen.artifactory.groupId="{}"][ns:codegen.artifactory.artifactId="{}"]'.format(group_id, artifact_id)
        properties = pom.findall(xpath, ns)
        if properties:
            codegen_artifact_id = pom.findall("./ns:artifactId", ns)[0].text
            m = re.match("oci-java-sdk-([^-]*)-codegen", codegen_artifact_id)
            if m:
                return m.group(1)

    return None


def get_package_names(group_artifact_ids_set, pr):
    global tc_link

    package_names_set = set([])
    problems = []
    for issue_key, group_id, artifact_id in group_artifact_ids_set:
        package_name = get_package_name_from_pom_file(group_id, artifact_id, args.java_sdk_path)

        if package_name:
            package_names_set.add(package_name)
        else:
            problems.append((issue_key, group_id, artifact_id))

    if problems:
        print("Had problems determining the package name for tickets: {}".format(", ".join(i for i, g, a in problems)))

        author_text = ""
        json = pr.json()
        if json['author'] and json['author']['user'] and json['author']['user']['name']:
            author_text = "@{name}: ".format(name=json['author']['user']['name'])

        text = NO_TESTS_FOUND_BECAUSE_UNKNOWN_SPEC.format(
            tc_link=tc_link,
            dexreq_tickets=", ".join(i for i, g, a in problems),
            ids_list=", ".join("`{}:{}`".format(g, a) for i, g, a in problems) + ("" if len(problems) == 1 else ", respectively"),
            ticket_sop="ticket" if len(problems) == 1 else "tickets",
            ticket_has_have="has" if len(problems) == 1 else "have",
            that_spec_those_specs="that spec" if len(problems) == 1 else "those specs",
            help_url=HELP_URL,
            author_text=author_text)

        if not ocits_shared.dry_run:
            make_general_comment("SDK", "oci-testing-service", pr_id, text)
        else:
            print("DRY-RUN: Not making BitBucket comment\n{}".format(text))

        # Don't fail the build
        sys.exit(1)

    printv("Package names: {}".format(", ".join(package_names_set)))

    return package_names_set


def get_test_classes_for_package(package_name, java_sdk_path):
    # bmc-integtests/src
    # +-- test
    #     +-- java
    #         +-- com
    #             +-- oracle
    #                 +-- bmc
    #                     +-- <package_name>
    #                         +-- *AutoTest.java
    src_dir = os.path.join(java_sdk_path, "bmc-integtests/src/test/java")
    package_dir = os.path.join(src_dir, "com/oracle/bmc/{}".format(package_name))
    printv("Looking for *AutoTest.java in {}".format(package_dir))
    test_files = [y for x in os.walk(package_dir) for y in glob(os.path.join(x[0], '*AutoTest.java'))]

    printv("Test files:\n{}".format("\n".join("\t{}".format(x) for x in test_files)))

    test_classes = []
    for test_file in test_files:
        if test_file.startswith(src_dir + "/") and test_file.endswith(".java"):
            within_src_dir = test_file[len(src_dir + "/"):-len(".java")]
            test_classes.append(within_src_dir.replace('/', '.'))

    return test_classes


def get_test_classes(package_names_set, java_sdk_path):
    global tickets, tc_link

    test_classes_set = set([])

    for package_name in package_names_set:
        test_classes = get_test_classes_for_package(package_name, java_sdk_path)

        if test_classes:
            for test_class in test_classes:
                test_classes_set.add(test_class)

    return test_classes_set


def get_test_classes_when_tickets_optional(package_names_set, java_sdk_path, pr_id):
    test_classes_set = get_test_classes(package_names_set, java_sdk_path)
    if not test_classes_set:
        text = NO_TESTS_FOUND_BUT_NO_DEXREQ.format(
            no_dexreq_marker=ocits_shared.NO_DEXREQ_MARKER,
            help_url=HELP_URL)

        if not ocits_shared.dry_run:
            make_general_comment("SDK", "oci-testing-service", pr_id, text)
        else:
            print("DRY-RUN: Not making BitBucket comment\n{}".format(text))

    return test_classes_set


def get_test_classes_when_tickets_required(package_names_set, java_sdk_path, pd_id, pr, tickets, tc_link):
    group_artifact_ids_set = get_group_and_artifact_ids(tickets, tc_link, pr)
    package_names_set.update(get_package_names(group_artifact_ids_set, pr))
    test_classes_set = get_test_classes(package_names_set, java_sdk_path)

    if not test_classes_set:
        print("No test classes found!")

        author_text = ""
        json = pr.json()
        if json['author'] and json['author']['user'] and json['author']['user']['name']:
            author_text = "@{name}: ".format(name=json['author']['user']['name'])

        text = NO_TESTS_FOUND_BECAUSE_NO_TEST_CLASSES.format(
            tc_link=tc_link,
            dexreq_tickets=", ".join(tickets),
            ticket_sop="ticket" if len(tickets) == 1 else "tickets",
            ticket_has_have="has" if len(tickets) == 1 else "have",
            java_package_names=", ".join("`com.oracle.bmc.{}`".format(x) for x in package_names_set),
            name_sop="name" if len(package_names_set) == 1 else "names",
            this_package_these_packages="this package" if len(package_names_set) == 1 else "these packages",
            help_url=HELP_URL,
            author_text=author_text)

        if not ocits_shared.dry_run:
            make_general_comment("SDK", "oci-testing-service", pr_id, text)
        else:
            print("DRY-RUN: Not making BitBucket comment\n{}".format(text))

        # Fail the build
        sys.exit(1)

    return test_classes_set


#
# Parameters variable set up
#
parser = argparse.ArgumentParser(description='Determine the test classes to run.')
parser.add_argument('--build-id', required=True, help="The TeamCity build id for the build that is running this script. This is used to update the relevant Bitbucket PRs with links to the TeamCity build")
parser.add_argument('--build-branch', required=True, help="The value of the teamcity.build.branch variable")
parser.add_argument('--java-sdk-path', required=True, help="Path to the root directory of the Java SDK")
parser.add_argument('-o', '--output', required=False, help='Output file')
parser.add_argument('-p', '--packages-output', required=False, help='Output file for packages')
parser.add_argument('-v', '--verbose', action='store_true', default=False, required=False, help='Verbose output')
parser.add_argument('--dry-run', action='store_true', default=False, required=False, help='Dry-run')

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
tc_link = TC_URL.format(build_id=args.build_id)

package_names_set = set([])
test_classes_set = set([])

if 'description' in pr.json():
    description = pr.json()['description']
else:
    description = ""

# Start with the test classes from the PR description
package_names_set.update(get_package_names_from_description(description))

tickets, no_dexreq_marker = get_dexreq_tickets(pr, tc_link)
if no_dexreq_marker:
    test_classes_set = get_test_classes_when_tickets_optional(package_names_set, args.java_sdk_path, pr_id)
else:
    test_classes_set = get_test_classes_when_tickets_required(package_names_set, args.java_sdk_path, pr_id, pr, tickets, tc_link)

test_classes_str = ",".join(str(s) for s in test_classes_set)
print(test_classes_str)

if args.output:
    f = open(args.output, "w")
    f.write(test_classes_str)

if args.packages_output:
    f = open(args.packages_output, "w")
    f.write(",".join(str(s) for s in package_names_set))
    f.flush()
