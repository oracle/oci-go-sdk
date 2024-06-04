import argparse
import os
import xml.etree.ElementTree as ET
import sys
from packaging import version

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../..'))

from shared.bitbucket_utils import setup_bitbucket, make_general_comment, clone_target_branch  # noqa: ignore=F402
import shared.bitbucket_utils  # noqa: ignore=F402

ns = {"ns":"http://maven.apache.org/POM/4.0.0"}


def get_pom_version(sdk_dir):
    pom_path = os.path.join(sdk_dir, 'pom.xml')
    pom = ET.parse(pom_path)
    xpath = './ns:version'
    return pom.find(xpath, ns).text


def truncate_pom_version(pom_version):
    dash_pos = pom_version.find("-")
    if dash_pos >= 0:
        pom_version = pom_version[:dash_pos]
    return pom_version


#
# Parameters variable set up
#
parser = argparse.ArgumentParser(description='Warn if the pom version does not match the target')
parser.add_argument('--build-branch', required=False, help="The value of the teamcity.build.branch variable")
parser.add_argument('--username', required=False, help='LDAP username ("firstname.lastname@oracle.com"; within TeamCity, use "%%system.teamcity.auth.userId%%")')
parser.add_argument('--password', required=False, help='LDAP password (within TeamCity, use "%%system.teamcity.auth.password%%")')
parser.add_argument('--source_branch_root_dir', required=False, help='Root directory of the source branch')
parser.add_argument('--target_branch_root_dir', required=False, help='Root directory of the target branch')
parser.add_argument('-v', '--verbose', action='store_true', default=False, required=False, help='Verbose output')
parser.add_argument('--dry-run', action='store_true', default=False, required=False, help='Dry-run, do not post comment to Bitbucket')

verbose = False
args = parser.parse_args()
setup_bitbucket(args)

if args.build_branch:
    pr_id = None
    try:
        # If the teamcity.build.branch variable is just something like "1234", then this is a
        # validation build for pull request "1234"
        pr_id = int(args.build_branch)
    except ValueError:
        print("Not a pull request validation build.")
        sys.exit(0)

dry_run = True if args.dry_run else False

if args.verbose:
    verbose = True
    shared.bitbucket_utils.verbose = True

source_branch_root_dir = os.path.abspath(args.source_branch_root_dir) if args.source_branch_root_dir else os.getcwd()
target_branch_root_dir = os.path.abspath(args.target_branch_root_dir) if args.target_branch_root_dir else clone_target_branch(pr_id, "java-sdk")

source_pom_version = get_pom_version(source_branch_root_dir)
trunc_source_pom_version = source_pom_version
target_pom_version = get_pom_version(target_branch_root_dir)
trunc_target_pom_version = target_pom_version
if verbose:
    print("Source branch pom version is {source_pom_version}".format(source_pom_version=source_pom_version))
    print("Target branch pom version is {target_pom_version}".format(target_pom_version=target_pom_version))


# Ensure that pom versions are not mismatched between public and preview.
if '-preview1' in source_pom_version and '-preview1' not in target_pom_version:
    text = "Source branch pom version ({source_pom_version}) is for preview while the target branch version ({target_pom_version}) is not.".format(source_pom_version=source_pom_version, target_pom_version=target_pom_version)
    if verbose:
        print(text)
    if not dry_run:
        make_general_comment("SDK", "java-sdk", pr_id, text)
    sys.exit(1)


if '-preview1' in target_pom_version and '-preview1' not in source_pom_version:
    text = "Source branch pom version ({source_pom_version}) is not for preview while the target branch version ({target_pom_version}) is.".format(source_pom_version=source_pom_version, target_pom_version=target_pom_version)
    if verbose:
        print(text)
    if not dry_run:
        make_general_comment("SDK", "java-sdk", pr_id, text)
    sys.exit(1)

# Truncate any trailing text after the version numbers for comparison
trunc_source_pom_version = truncate_pom_version(source_pom_version)
trunc_target_pom_version = truncate_pom_version(target_pom_version)

if verbose:
    print("Truncated source branch pom version is {trunc_source_pom_version}".format(trunc_source_pom_version=trunc_source_pom_version))
    print("Truncated target branch pom version is {trunc_target_pom_version}".format(trunc_target_pom_version=trunc_target_pom_version))


if version.parse(trunc_source_pom_version) < version.parse(trunc_target_pom_version):
    text = "The pom version of the source branch ({source_pom_version}) is out of date with the target branch ({target_pom_version}). Please re-fetch from the remote and rebase your changes on top of the target branch.".format(source_pom_version=source_pom_version, target_pom_version=target_pom_version)
    if verbose:
        print(text)
    if not dry_run:
        make_general_comment("SDK", "java-sdk", pr_id, text)
    sys.exit(1)
elif verbose:
    print("The source and target branch pom versions are in sync.")
