from __future__ import print_function
import argparse
import os
import requests
import urllib3
import re
import ssl
import sys
import traceback
from xml.etree import ElementTree
import getpass

import shared.bitbucket_utils  # noqa: ignore=F402

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
verbose = False
if (not os.environ.get('PYTHONHTTPSVERIFY', '') and getattr(ssl, '_create_unverified_context', None)): 
    ssl._create_default_https_context = ssl._create_unverified_context


def printv(str):
    global verbose
    if verbose:
        print(str)


def setup_auth():
    global args, username, password, auth
    username = args.username
    if args.password is not None:
        password = args.password
    else:
        password = getpass.getpass("LDAP password:")

    auth = (username, password)


def get_changed_files_for_change(change_id):
    url = "https://teamcity.oci.oraclecorp.com/app/rest/changes/id:{}".format(change_id)
    r = requests.get(url, verify=False, auth=auth)

    root = ElementTree.fromstring(r.content)
    files_node = root.find("files")

    files = []

    for child in files_node:
        if child.tag.lower() == "file":
            files.append(child.attrib['file'])

    return files


def get_commit_hashes(pr_id):
    pr_commits = shared.bitbucket_utils.get_pullrequest_commits("SDK", "legacy-java-sdk", pr_id)

    commit_hashes = []

    for change in pr_commits.json()['values']:
        commit_hashes.append(change['id'])
        printv("Commit hash {}".format(change['id']))

    return commit_hashes


def get_changed_files(build_id, pr_id):
    commit_hashes = None
    if pr_id:
        commit_hashes = get_commit_hashes(pr_id)

    url = "https://teamcity.oci.oraclecorp.com/app/rest/changes?locator=build:(id:{})".format(build_id)
    r = requests.get(url, verify=False, auth=auth)

    root = ElementTree.fromstring(r.content)

    files = []

    for child in root:
        if child.tag.lower() == "change":
            if commit_hashes and child.attrib['version'] not in commit_hashes:
                continue
            change_id = child.attrib['id']
            files.extend(get_changed_files_for_change(change_id))

    return files


def get_changed_files_for_pr(pr_id):
    try:
        pr_diff = shared.bitbucket_utils.get_pullrequest_diff("SDK", "legacy-java-sdk", pr_id)
        json = pr_diff.json()
        if json['truncated']:
            printv("Diff for {} is truncated".format(pr_id))
            return None

        files = []

        for diff in json['diffs']:
            if 'destination' in diff:
                diff_destination = diff['destination']
                if 'toString' in diff_destination:
                    diff_file = diff_destination['toString']
                    printv("Diff file {}".format(diff_file))
                    files.append(diff_file)

        return files
    except Exception as e:
        printv("Failed to get diff for {}".format(pr_id))
        printv("type error: {}".format(str(e)))
        printv(traceback.format_exc())
        return None


#
# Parameters variable set up
#
parser = argparse.ArgumentParser(description='Determine the recommended build profile for the change being built: "dev" or "quick"')
parser.add_argument('--build-id', required=False, help='TeamCity build id')
parser.add_argument('--username', required=False, help='LDAP username ("firstname.lastname@oracle.com"; within TeamCity, use "%%system.teamcity.auth.userId%%")')
parser.add_argument('--password', required=False, help='LDAP password (within TeamCity, use "%%system.teamcity.auth.password%%")')
parser.add_argument('--build-branch', required=False, help="The value of the teamcity.build.branch variable")
parser.add_argument('--changed-modules-output-file', required=False, help="If provided, the changed modoules will be written to this file")
parser.add_argument('-v', '--verbose', action='store_true', default=False, required=False, help='Verbose output')

args = parser.parse_args()
shared.bitbucket_utils.setup_bitbucket(args)

if args.verbose:
    verbose = True

pr_id = None
if args.build_branch:
    try:
        # If the teamcity.build.branch variable is just something like "1234", then this is a
        # validation build for pull request "1234"
        pr_id = int(args.build_branch)
    except ValueError:
        print("Not a pull request validation build.")
        sys.exit(0)

build_id = args.build_id

if pr_id:
    changed_files = get_changed_files_for_pr(pr_id)
else:   
    if not args.build_id or not args.username:
        print("--build-id and --username required if --build-branch is not set")
        sys.exit(1)

    # Fall back to TeamCity diff
    setup_auth()
    changed_files = get_changed_files(build_id, pr_id)

if not changed_files:
    # No changes, let's rebuild everything to be safe.
    # This can happen when a job failed, and then we hit "Run" again
    pom_file_change = True
else:
    pom_file_change = False

hand_written_directories = [
    "bmc-common/",
    "bmc-smoketests/",
    "bmc-addons/",
    "bmc-circuitbreaker/",
    "bmc-hand-written/",
    "bmc-examples/",
    "bmc-encryption/",
    "bmc-objectstorage/bmc-objectstorage-extensions/"
    "bmc-objectstorage/bmc-objectstorage-generated/src/main/java/com/oracle/bmc/objectstorage/internal/http/ObjectMetadataInterceptor.java",
    "bmc-streaming/src/main/java/com/oracle/bmc/streaming/StreamClientBuilder.java",
    "bmc-streaming/src/main/java/com/oracle/bmc/streaming/AbstractStreamBasedClientBuilder.java",
    "bmc-streaming/src/main/java/com/oracle/bmc/streaming/StreamAsyncClientBuilder.java",
    "bmc-keymanagement/src/main/java/com/oracle/bmc/keymanagement/KmsManagementClientBuilder.java",
    "bmc-keymanagement/src/main/java/com/oracle/bmc/keymanagement/KmsManagementAsyncClientBuilder.java",
    "bmc-keymanagement/src/main/java/com/oracle/bmc/keymanagement/KmsCryptoAsyncClientBuilder.java",
    "bmc-keymanagement/src/main/java/com/oracle/bmc/keymanagement/AbstractVaultBasedClientBuilder.java",
    "bmc-keymanagement/src/main/java/com/oracle/bmc/keymanagement/AbstractKmsCryptoClientBuilder.java",
    "bmc-keymanagement/src/main/java/com/oracle/bmc/keymanagement/KmsCryptoClientBuilder.java",
    "bmc-keymanagement/src/main/java/com/oracle/bmc/keymanagement/AbstractKmsManagementClientBuilder.java",
    "bmc-shaded/"
    "bmc-shaded-smoketests/"
]
hand_written_change = False
truncated = False

root_module_changed = False
changed_modules = []

if changed_files:
    printv("{} changed files".format(len(changed_files)))

    for file in changed_files:
        printv(file)

        if file.lower().endswith("pom.xml"):
            if not file.lower().startswith("bmc-codegen/bmc-"):
                pom_file_change = True
                printv("pom file change in {}".format(file))
        for d in hand_written_directories:
            if file.lower().startswith(d):
                hand_written_change = True
                printv("Change in hand-written directory in {}".format(file))

        m = re.search(r'^.*bmc-([^/]*)', file)
        if m:
            module_name = m.group(0)
            if module_name not in changed_modules:
                changed_modules.append(module_name)
        else:
            root_module_changed = True
else:
    printv("Truncated response from Bitbucket, building everything")
    truncated = True

if pom_file_change or hand_written_change or truncated:
    print('dev')
else:
    print('quick')

if args.changed_modules_output_file:
    with open(args.changed_modules_output_file, 'w') as writer:
        if not truncated and not root_module_changed and changed_modules:
            changed_modules_output = "--projects {}".format(",".join(changed_modules))
            printv(changed_modules_output)
            writer.write(changed_modules_output)
        else:
            printv("Not writing individual changed modules, truncated? {}, root module changed? {}, changed_modules? {}".format(
                truncated, root_module_changed, changed_modules))
