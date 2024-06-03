from __future__ import print_function
import argparse
import os
import ssl
import sys
import urllib3
import re
from git import Repo

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../..'))

from shared.bitbucket_utils import setup_bitbucket, make_general_comment, get_pullrequest, get_pr_source_branch  # noqa: ignore=F402
import shared.bitbucket_utils  # noqa: ignore=F402

KEEP_TEMP_FILES = True

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
verbose = False
if (not os.environ.get('PYTHONHTTPSVERIFY', '') and getattr(ssl, '_create_unverified_context', None)): 
    ssl._create_default_https_context = ssl._create_unverified_context


def printv(str):
    global verbose
    if verbose:
        print(str)


#
# Parameters variable set up
#
parser = argparse.ArgumentParser(description='Check out the source branch, as indicated in the PR. Warning: this is destructive to the contents in the Java SDK directory.')
parser.add_argument('--build-id', required=False, help="The TeamCity build id for the build that is running this script. This is used to update the relevant Bitbucket PRs with links to the TeamCity build")
parser.add_argument('--build-branch', required=True, help="The value of the teamcity.build.branch variable")
parser.add_argument('--java-sdk-dir', required=False, help='Directory of the Java SDK')
parser.add_argument('-v', '--verbose', action='store_true', default=False, required=False, help='Verbose output')

args = parser.parse_args()
setup_bitbucket(args)

if args.verbose:
    verbose = True
    shared.bitbucket_utils.verbose = True

if args.java_sdk_dir:
    java_sdk_dir = os.path.abspath(args.java_sdk_dir)
else:
    java_sdk_dir = os.getcwd()

pr_id = None
try:
    # If the teamcity.build.branch variable is just something like "1234", then this is a 
    # validation build for pull request "1234"
    pr_id = int(args.build_branch)
except ValueError:
    print("Not a pull request validation build. Not warning.")
    sys.exit(0)

pr = get_pullrequest("SDK", "java-sdk", pr_id)
printv(pr.text)

try:
    repo = Repo.init(java_sdk_dir)

    current_commit = None
    current_branch = [branch.strip()[2:] for branch in repo.git.branch().split('\n') if branch.startswith('* ')][0]
    printv("current branch: {}".format(current_branch))
    result = re.search(r'\(HEAD detached at ([^)]*)\)', current_branch)
    if not result:
        # this is what it looks like in Team City
        result = re.search(r'\(detached from ([^)]*)\)', current_branch)
    if result:
        current_commit = result.group(1)

    source_branch = get_pr_source_branch(pr)
    printv("source branch: {}".format(source_branch))

    repo.git.fetch("origin")
    repo.git.checkout(source_branch)

    if current_commit:
        repo.git.reset('--hard', current_commit)
        printv("resetting to commit: {}".format(current_commit))
except Exception as e:
    print('EXCEPTION: {}'.format(str(e)))
    print('Failed to change to source branch.')
