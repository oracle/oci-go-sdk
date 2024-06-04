from __future__ import print_function
import argparse
import os
import urllib3
import re
import ssl
import traceback

import shared.bitbucket_utils  # noqa: ignore=F402

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
verbose = False
if (not os.environ.get('PYTHONHTTPSVERIFY', '') and getattr(ssl, '_create_unverified_context', None)): 
    ssl._create_default_https_context = ssl._create_unverified_context


def printv(str):
    global verbose
    if verbose:
        print(str)


def get_changed_files_for_commit(commit_hash, commit_path):
    try:
        pr_diff = shared.bitbucket_utils.get_commit_diff("SDK", "java-sdk", commit_hash, commit_path)
        json = pr_diff.json()
        printv(json)
        if json['truncated']:
            printv("Diff for {} is truncated".format(commit_hash))
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
        printv("Failed to get diff for {}".format(commit_hash))
        printv("type error: {}".format(str(e)))
        printv(traceback.format_exc())
        return None


#
# Parameters variable set up
#
parser = argparse.ArgumentParser(description='Determine the changed codegen projects to be used for generation')
parser.add_argument('--commit', required=True, help="The commit hash with the codegen pom.xml changes")
parser.add_argument('--commit-path', required=False, help="Subpath for the commit diff, e.g. 'bmc-codegen'")
parser.add_argument('--changed-modules-output-file', required=False, help="If provided, the changed modoules will be written to this file")
parser.add_argument('-v', '--verbose', action='store_true', default=False, required=False, help='Verbose output')

args = parser.parse_args()
shared.bitbucket_utils.setup_bitbucket(args)

if args.verbose:
    verbose = True

changed_files = get_changed_files_for_commit(args.commit, args.commit_path)

should_generate_everything = False
truncated = False

changed_modules = []

if changed_files:
    printv("{} changed files".format(len(changed_files)))

    for file in changed_files:
        printv(file)

        if file.lower().startswith("bmc-codegen/bmc-") and file.lower().endswith("pom.xml"):
            m = re.search(r'^.*bmc-([^/]*)', file)
            if m:
                module_name = m.group(0)
                if module_name not in changed_modules:
                    changed_modules.append(module_name)
                    printv("codegen pom file change in {}".format(file))
            else:
                printv("couldn't extract codegen pom module for file change in {}, generating everything".format(file))
                should_generate_everything = True
else:
    printv("Truncated response from Bitbucket, generating everything")
    truncated = True

changed_modules_output = ""
if not truncated and not should_generate_everything and changed_modules:
    changed_modules_output = "--projects {}".format(",".join(changed_modules))
    print(changed_modules_output)
else:
    printv("Not writing individual changed modules, truncated? {}, should generate everything? {}, changed_modules? {}".format(
        truncated, should_generate_everything, changed_modules))

if args.changed_modules_output_file:
    with open(args.changed_modules_output_file, 'w') as writer:
        writer.write(changed_modules_output)
