from __future__ import print_function
import argparse
import os
import ssl
import sys
import urllib3
import ocits_shared
from ocits_shared import printv, HELP_URL, TC_URL

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

FAILURE = """
{author_text}{message}

{if_log_text}More information is {if_log_also}available in the [TeamCity build log]({tc_link}).

Once you have fixed the problem, you can restart the [TeamCity build]({tc_link}) by asking an SDK/CLI team member on #oci_public_sdks to click on the 'Run' button. That is not necessary if the fix required a code change to the OCI Testing Service. In that case, a new validation build will start automatically.

For more information, see [Pull Request Validation Builds for the Testing Service]({help_url}).{log}
"""


#
# Parameters variable set up
#
parser = argparse.ArgumentParser(description='Report testing service failure.')
parser.add_argument('--build-id', required=True, help="The TeamCity build id for the build that is running this script. This is used to update the relevant Bitbucket PRs with links to the TeamCity build")
parser.add_argument('--build-branch', required=True, help="The value of the teamcity.build.branch variable")
parser.add_argument('--log-input', required=False, help="Path to the log file for diagnostic output")
parser.add_argument('--message', required=True, help="Message to include in the PR comment")
parser.add_argument('-v', '--verbose', action='store_true', default=False, required=False, help='Verbose output')
parser.add_argument('--dry-run', action='store_true', default=False, required=False, help='Dry-run, do not post to PR')

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
    print("Not a pull request validation build.")
    sys.exit(2)

pr = get_pullrequest("SDK", "oci-testing-service", pr_id)
printv(pr.text)

author_text = ""
json = pr.json()
if json['author'] and json['author']['user'] and json['author']['user']['name']:
    author_text = "@{name}: ".format(name=json['author']['user']['name'])

log = ""
if args.log_input and os.path.exists(args.log_input):
    file = open(args.log_input, "r")
    log = file.read()
    log = "\n\n```\n" + log + "```"

tc_link = TC_URL.format(build_id=args.build_id)

text = FAILURE.format(
    message=args.message,
    tc_link=tc_link,
    help_url=HELP_URL,
    log=log,
    if_log_text="Please use the information below to diagnose the problem. " if log else "",
    if_log_also="also " if log else "",
    author_text=author_text)

if args.dry_run:
    print("DRY-RUN: {}".format(text))
else:
    make_general_comment("SDK", "oci-testing-service", pr_id, text)

# Don't fail the build, let that happen in the TC script
sys.exit(0)
