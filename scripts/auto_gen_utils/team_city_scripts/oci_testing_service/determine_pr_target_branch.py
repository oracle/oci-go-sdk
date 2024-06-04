from __future__ import print_function
import argparse
import os
import ssl
import sys
import urllib3

import ocits_shared
from ocits_shared import printv

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../..'))

from shared.bitbucket_utils import setup_bitbucket, get_pullrequest, get_pr_target_branch  # noqa: ignore=F402
import shared.bitbucket_utils  # noqa: ignore=F402

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
if (not os.environ.get('PYTHONHTTPSVERIFY', '') and getattr(ssl, '_create_unverified_context', None)): 
    ssl._create_default_https_context = ssl._create_unverified_context


#
# Parameters variable set up
#
parser = argparse.ArgumentParser(description='Determine the target branch of the pull request.')
parser.add_argument('--build-branch', required=True, help="The value of the teamcity.build.branch variable")
parser.add_argument('-v', '--verbose', action='store_true', default=False, required=False, help='Verbose output')

args = parser.parse_args()
setup_bitbucket(args)

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
    sys.exit(0)

pr = get_pullrequest("SDK", "oci-testing-service", pr_id)
printv(pr.text)

target_branch = get_pr_target_branch(pr)
print(target_branch)
