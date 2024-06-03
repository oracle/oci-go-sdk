# -*- coding: utf-8 -*-
from __future__ import print_function
import argparse
import os
import ssl
import sys
import urllib3
from git import Repo

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../..'))

from shared.bitbucket_utils import setup_bitbucket, make_general_comment  # noqa: ignore=F402
import shared.bitbucket_utils  # noqa: ignore=F402
from shared.buildsvc_tc_compatibility import build_log_link


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
parser = argparse.ArgumentParser(description='Warn if source formatting is incorrect (really, warn if there are changed files in this git repository).')
parser.add_argument('--build-id', required=False, help="The TeamCity build id for the build that is running this script. This is used to update the relevant Bitbucket PRs with links to the TeamCity build")
parser.add_argument('--build-branch', required=True, help="The value of the teamcity.build.branch variable")
parser.add_argument('--java-sdk-dir', required=False, help='Directory of the Java SDK')
parser.add_argument('-v', '--verbose', action='store_true', default=False, required=False, help='Verbose output')
parser.add_argument('--dry-run', action='store_true', default=False, required=False, help='Dry-run, do not post comment to Bitbucket')

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

pr_id = None
try:
    # If the teamcity.build.branch variable is just something like "1234", then this is a 
    # validation build for pull request "1234"
    pr_id = int(args.build_branch)
except ValueError:
    print("Not a pull request validation build. Not warning.")
    sys.exit(0)

repo = Repo.init(java_sdk_dir)
git_status = repo.git.status()
git_diff = repo.git.diff('--name-only', '--ignore-submodules')

current_branch = [branch.strip()[2:] for branch in repo.git.branch().split('\n') if branch.startswith('* ')][0]
printv("source branch: {}".format(current_branch))

printv("git diff:\n{}".format(git_diff))

printv("git status: '{}'".format(git_status))

could_fix = False
if 'nothing to commit' not in git_status and "nothing added to commit" not in git_status:
    # Try to fix it

    try:
        repo.git.add("*.java")
        repo.git.add("*.properties")
        repo.git.commit("-m", "Committing source formatting changes. Automatically performed, see build log:\n\n{build_log_link}".format(
            build_log_link=build_log_link(args.build_id, text="build log")))
        if dry_run:
            print('DRY-RUN: not pushing to branch {}'.format(current_branch))
        else:
            repo.git.push('-u', 'origin', current_branch)

        could_fix = True
        text = "There were source formatting problems, but we automatically committed the changes to your branch. Note that this may kick off another validation build. Changes made:\n\n" + git_diff + "\n\nMore information may also available in the {build_log_link}.".format(build_log_link=build_log_link(build_id, text="build log"))
    except Exception as e:
        print('EXCEPTION: {}'.format(str(e)))
        print('Failed to push source formatting changes.')

    if not could_fix:  
        text = "There are source formatting problems in the files below. Run `mvn process-sources -Pdev`.\n\n" + git_diff + "\n\nDon't want to do this by hand anymore? Give the 'DEXREQ Automation' user write access to your repo ([similar instructions here](https://confluence.oci.oraclecorp.com/pages/viewpage.action?spaceKey=DEX&title=Pull+Request+Validation+Builds+for+the+Testing+Service#PullRequestValidationBuildsfortheTestingService-WhenUsingYourOwnFork))."

        if args.build_id:
            text = text + "\n\nPlease use the information above to diagnose the problem. More information may also available in the {build_log_link}.".format(
                build_log_link=build_log_link(args.build_id, text="build log"))

    print(text)

    if not dry_run:
        make_general_comment("SDK", "java-sdk", pr_id, text)
    else:
        print("DRY-RUN: Not making Bitbucket comment.")
else:
    print("Nothing to warn about!")
    could_fix = True

if not could_fix:
    sys.exit(1)
