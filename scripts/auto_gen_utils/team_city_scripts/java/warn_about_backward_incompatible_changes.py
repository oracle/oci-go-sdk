from __future__ import print_function
import argparse
import os
import re
import shutil
import ssl
import sys
import urllib3
import getpass
import xml.etree.ElementTree as ET
from filecmp import dircmp
from glob import glob
from git import Repo

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../..'))

from shared.bitbucket_utils import setup_bitbucket, get_pullrequest, make_general_comment, clone_target_branch  # noqa: ignore=F402
from add_or_update_scripts.add_or_update_spec_utils import parse_pom  # noqa: ignore=F402
import shared.bitbucket_utils  # noqa: ignore=F402
import util  # noqa:E402
import config  # noqa:E402
from shared.buildsvc_tc_compatibility import build_log_link


KEEP_TEMP_FILES = True

CUSTOM_JIRA_ISSUE_FIELDS = [config.CUSTOM_FIELD_ID_GROUP_ID, config.CUSTOM_FIELD_ID_ARTIFACT_ID]

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


def append_diff_file(diff_files, file):
    if not file.endswith('.bak') and not file.endswith('.iml'):
        diff_files.append(file)


def get_diff_files(dcmp, diff_files=[]):
    for name in dcmp.diff_files:
        append_diff_file(diff_files, os.path.join(dcmp.left, name))
    for name in dcmp.left_only:
        append_diff_file(diff_files, os.path.join(dcmp.left, name))
    for name in dcmp.right_only:
        append_diff_file(diff_files, os.path.join(dcmp.left, name))
    for sub_dcmp in dcmp.subdirs.values():
        get_diff_files(sub_dcmp, diff_files)

    return diff_files


def get_changed_files(java_sdk_dir, pr_id, base_dir):
    should_clean_up = False
    if not base_dir:
        should_clean_up = not KEEP_TEMP_FILES
        base_dir = clone_target_branch(pr_id, "java-sdk")
        if not base_dir:
            return None

    dcmp = dircmp(java_sdk_dir, base_dir, ['RCS', 'CVS', 'tags', '.git', 'target',
        'pom.xml.versionsBackup', '.DS_Store', '.idea'])
    files = get_diff_files(dcmp)

    prefix = os.path.join(java_sdk_dir, "")

    files = [f[len(prefix):] if f.startswith(prefix) else f for f in files]

    printv("Files that differ:")
    for f in files:
        printv(f)

    if should_clean_up:
        shutil.rmtree(base_dir, ignore_errors=True, onerror=None)

    return files


def find_files_with_filter(dir, filter):
    """
    Search a path defined by dir for files with a name that matches filter

    @param str dir:
        Path to search

    @param str filter:
        filter for files to match on.  Example 'clirr-*.txt'
    """
    files = [os.path.abspath(y) for x in os.walk(dir) for y in glob(os.path.join(x[0], filter))]

    return files


def find_clirr_files(dir):
    clirr_files = find_files_with_filter(dir, 'clirr-*.txt')

    return clirr_files


def find_pom_files(dir):
    pom_files = find_files_with_filter(dir, "pom.xml")

    return pom_files


def get_project(dir, group_id, artifact_id):
    """
    This function searches all the pom files contained in 'dir' the group_id
    and artifact_id
    """

    ns = {"ns":"http://maven.apache.org/POM/4.0.0"}

    # allow default namespace for output, dont print ns0: prefixes everywhere
    ET.register_namespace('',"http://maven.apache.org/POM/4.0.0")

    group_id_xpath = ".//ns:properties//ns:codegen.artifactory.groupId"
    artifact_id_xpath = ".//ns:properties//ns:codegen.artifactory.artifactId"

    pom_files = find_pom_files(dir)

    file_name = None
    project = None

    for pom_file in pom_files:
        try:
            pom = parse_pom(pom_file)
            property = pom.findall(group_id_xpath, ns)[0]
            pom_group_id = property.text
            property = pom.findall(artifact_id_xpath, ns)[0]
            pom_artifact_id = property.text
        except IndexError:
            # An IndexError means findall did not return anything
            # for the path, which means this is not the pom file
            # we want.
            continue

        if group_id == pom_group_id and artifact_id == pom_artifact_id:
            file_name = pom_file
            break

    # If file_name is not None it should be in the form of
    # "/folders/java-sdk/bmc-codegen/{project}-codegen/pom.xml" where {project}
    # is the part that we want.
    if file_name:
        head, tail = os.path.split(file_name)
        if head.endswith("-codegen"):
            path_parts = head.split(os.sep)
            project = path_parts[-1][:-len("-codegen")]

    return project


#
# Parameters variable set up
#
parser = argparse.ArgumentParser(description='Warn if there are clirr errors in packages that are changed.  JIRA_USERNAME, JIRA_PASSWORD, BITBUCKET_USERNAME, AND BITBUCKET_PASSWORD are expected env vars.')
parser.add_argument('--build-id', required=False, help="The TeamCity build id for the build that is running this script. This is used to update the relevant Bitbucket PRs with links to the TeamCity build")
parser.add_argument('--build-branch', required=True, help="The value of the teamcity.build.branch variable")
parser.add_argument('--username', required=True, help='LDAP username ("firstname.lastname@oracle.com"; within TeamCity, use "%%system.teamcity.auth.userId%%")')
parser.add_argument('--password', required=False, help='LDAP password (within TeamCity, use "%%system.teamcity.auth.password%%")')
parser.add_argument('--java-sdk-dir', required=False, help='Directory of the Java SDK')
parser.add_argument('-v', '--verbose', action='store_true', default=False, required=False, help='Verbose output')
parser.add_argument('--base-dir', required=False, help='The directory that has the Git checkout of the target branch')
parser.add_argument('--keep-temp-files', action='store_true', default=False, required=False, help='Keep temporary files')
parser.add_argument('--dry-run', action='store_true', default=False, required=False, help='Dry-run, do not post comment to Bitbucket')
parser.add_argument('--build-type', choices=['DEXREQ', 'PR'], default='PR', help='Build type for the TeamCity job')

args = parser.parse_args()
setup_bitbucket(args)

if args.dry_run:
    dry_run = True
else:
    dry_run = False

if args.verbose:
    verbose = True
    shared.bitbucket_utils.verbose = True

if args.java_sdk_dir:
    java_sdk_dir = os.path.abspath(args.java_sdk_dir)
else:
    java_sdk_dir = os.getcwd()

if args.keep_temp_files:
    KEEP_TEMP_FILES = True

clirr_files = find_clirr_files(java_sdk_dir)
if not clirr_files:
    print("No clirr*.txt files found. Nothing to warn about.")
    sys.exit(0)

comment_destination = "Bitbucket PR"
if args.build_type == "DEXREQ":
    comment_destination = "DEXREQ ticket"

pr_id = args.build_branch
if args.build_type == 'PR':
    try:
        # If the teamcity.build.branch variable is just something like "1234", then this is a
        # validation build for pull request "1234"
        pr_id = int(args.build_branch)
    except ValueError:
        print("Not a pull request validation build. Not warning.")
        sys.exit(0)
    setup_auth()

warn_about_clirr_files = []

changed_files = []
issue = None
if args.build_type == 'DEXREQ':
    # Get DEXREQ ticket from Git log
    repo = Repo(java_sdk_dir)
    last_commit_message = repo.git.log(n=1, format='%s%n%b')
    dexreq_issues = util.parse_issue_keys_from_commit_message(last_commit_message)

    # Get artifact id and group id from the first DEXREQ Ticket
    issue = None
    project = None
    if dexreq_issues and len(dexreq_issues) > 0:
        issue = util.get_dexreq_issue(dexreq_issues[0], fields=(CUSTOM_JIRA_ISSUE_FIELDS))

    if issue:
        group_id = getattr(issue.fields, config.CUSTOM_FIELD_ID_GROUP_ID)
        artifact_id = getattr(issue.fields, config.CUSTOM_FIELD_ID_ARTIFACT_ID)
        project = get_project(os.path.join(java_sdk_dir, "bmc-codegen"),
                            group_id,
                            artifact_id)

    if project:
        print("Adding {0} to changed files".format(project))
        changed_files.append(project)
    else:
        print("No project found, warning about all backward compatibilities!")

else:
    try:
        changed_files = get_changed_files(java_sdk_dir, pr_id, args.base_dir)
    except Exception as e:
        print(e)
        print("Could not get changed files, warning about all backward compatibilities!")

if not changed_files:
    # No changes, let's warn about everything to be safe.
    # This can happen when a job failed, and then we hit "Run" again
    warn_about_clirr_files = clirr_files
else:
    # Find the clirr files that correspond to changed files
    for clirr_file in clirr_files:
        printv("Checking {}".format(clirr_file))
        project_dir = os.path.abspath(os.path.join(clirr_file, os.pardir, os.pardir))

        prefix = os.path.join(java_sdk_dir, "")
        if project_dir.startswith(prefix):
            project_dir = project_dir[len(prefix):]

        # See if any changed files start with this
        for cf in changed_files:
            if cf.startswith(project_dir):
                printv("\tFound changed file in this module: {}".format(cf))
                warn_about_clirr_files.append(clirr_file)
                break

messages_for_modules = {}

for file in warn_about_clirr_files:
    printv("Emit errors from {}".format(file))

    m = re.match(r"^.*/clirr-(.*)\.txt$", file)
    if m:
        module_name = m.group(1)
    else:
        module_name = "Unknown module"

    with open(file, 'r') as content_file:
        content = content_file.read()
        if content and len(content) > 0:
            if module_name in messages_for_modules:
                text = messages_for_modules[module_name] + "\n"
            else:
                text = ""
            text = text + "\n" + content
            messages_for_modules[module_name] = text

if messages_for_modules:
    text = "Clirr detected backward incompatibilities possibly related to this change, compared to the latest release version of the OCI Java SDK:\n\n"

    if not changed_files:
        text = text + "(Note: Could not detect changed files; including all backward incompatibilities to be safe.)\n\n"

    for module, messages in messages_for_modules.items():
        text = text + "{}:\n{}\n\n".format(module, messages)

    if args.build_id:
        text = text + "\n\nPlease use the information above to diagnose the problem. More information may also available in the {build_log_link}.".format(
            build_log_link=build_log_link(args.build_id, text="build log"))

    if text:
        text = text.encode('utf8')

    print(text)

    if not dry_run:
        if args.build_type == "DEXREQ":
            if issue:
                print("Adding backwards incompatibility comment to {}".format(issue.key))
                util.add_jira_comment(issue.key, text)
                printv("Adding '{}' label to: {}".format(config.BACKWARD_INCOMPATIBLE_CHANGES_LABEL, issue.key))
                issue.add_field_value('labels', config.BACKWARD_INCOMPATIBLE_CHANGES_LABEL)
            else:
                print("No DEXREQ ticket identified. Cannot add comment")
        else:
            make_general_comment("SDK", "java-sdk", pr_id, text)
    else:
        print("DRY-RUN: Not adding comment to {}".format(comment_destination))
else:
    print("Nothing to warn about!")
