from __future__ import print_function
import argparse
import os
import requests
import urllib3
import ssl
import hashlib
import re
import ntpath

GROUP_ID = "com.oracle.oci.sdk"
ARTIFACT_ID = "oci-java-sdk-dist"
ARTIFACT_TYPE = "zip"

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
verbose = False
if (not os.environ.get('PYTHONHTTPSVERIFY', '') and getattr(ssl, '_create_unverified_context', None)): 
    ssl._create_default_https_context = ssl._create_unverified_context


def printv(str):
    global verbose
    if verbose:
        print(str)


# Unfortunately, this needs to be authenticated
def get_artifact_sha256(auth, build_name, build_number, group_id, artifact_id, version, artifact_type):
    requested_module = "{}:{}:{}".format(group_id, artifact_id, version)
    url = "https://artifactory.oci.oraclecorp.com/api/build/{}/{}".format(build_name, build_number)
    printv("URL: {}".format(url))

    r = requests.get(url, verify=False, auth=auth)

    modules = r.json()['buildInfo']['modules']

    for module in modules:
        module_id = module['id']
        if module_id == requested_module:
            artifacts = module['artifacts']
            for artifact in artifacts:
                if artifact['type'] == artifact_type:
                    sha256 = artifact['sha256']
                    printv('Found artifact "{}", type "{}": sha256 is "{}"'.format(module_id, artifact_type, sha256))
                    return sha256

    printv('Not found: artifact "{}", type "{}"'.format(requested_module, artifact_type))

    return None


def get_file_sha1(artifact_file, block_size=65536):
    sha1 = hashlib.sha1()
    with open(artifact_file, 'rb') as f:
        for block in iter(lambda: f.read(block_size), b''):
            sha1.update(block)
    return sha1.hexdigest()


def get_snapshot_version_with_sha1(group_id, artifact_id, version, artifact_type, sha1):
    url = "https://artifactory.oci.oraclecorp.com/api/storage/opc-public-sdk-snapshot-maven-local/{}/{}/{}".format(group_id.replace('.','/'), artifact_id, version)
    r = requests.get(url, verify=False)
    children = r.json()['children']

    for child in reversed(children):
        if not child['folder']:
            candidate = child['uri'][1:]
            if (candidate.startswith(artifact_id) and candidate.endswith(artifact_type)):
                m = re.search(artifact_id + "-(.*)." + artifact_type, candidate)
                candidate_version = m.group(1)
                printv("Candidate version: {}".format(candidate_version))

                candidate_url = url + "/" + candidate
                r = requests.get(candidate_url, verify=False)

                candidate_sha1 = r.json()['checksums']['sha1']
                printv("\tCandidate sha1: {}".format(candidate_sha1))
                if candidate_sha1 == sha1:
                    return candidate_version

    return None


#
# Parameters variable set up
#
parser = argparse.ArgumentParser(description='Determine the full version of a Legacy Java SDK build just deployed to Artifactory')
parser.add_argument('--group-id', default=GROUP_ID, help='Group id of the artifact, default is {}'.format(GROUP_ID))
parser.add_argument('--artifact-id', default=ARTIFACT_ID, help='Artifact id of the artifact, default is {}'.format(ARTIFACT_ID))
parser.add_argument('--version', required=False, help='Version that was built (e.g. "1.2.3" or "1.2.3-SNAPSHOT")')
parser.add_argument('--artifact-type', default=ARTIFACT_TYPE, help='Artifact type of the artifact, default is {}'.format(ARTIFACT_TYPE))
parser.add_argument('--file', required=True, help='The built {}:{} artifact whose full version should be retrieved'.format(GROUP_ID, ARTIFACT_ID))
parser.add_argument('-v', '--verbose', action='store_true', default=False, required=False, help='Verbose output')

args = parser.parse_args()

if args.verbose:
    verbose = True

group_id = args.group_id
artifact_id = args.artifact_id
artifact_type = args.artifact_type
artifact_file = args.file

if args.version:
    version = args.version
else:
    file_name = ntpath.basename(artifact_file)
    if not file_name.startswith(artifact_id):
        print('File name does not start with "{}", cannot autodetect version'.format(artifact_id))
        exit(1)
    if not file_name.endswith(artifact_type):
        print('File name does not end with "{}", cannot autodetect version'.format(artifact_type))
        exit(1)
    m = re.search(artifact_id + "-(.*)." + artifact_type, file_name)
    version = m.group(1)
    printv('Autodetected version "{}"'.format(version))

if not version.upper().endswith("-SNAPSHOT"):
    # if it's not a snapshot, the full version is just the version that was built
    print(version)
    exit(0)

printv('Recognizing version "{}" as snapshot, need to determine full version'.format(version))

requested_sha1 = get_file_sha1(artifact_file)
printv('sha1 of file "{}" is: {}'.format(artifact_file, requested_sha1))

timed_snapshot = get_snapshot_version_with_sha1(group_id, artifact_id, version, artifact_type, requested_sha1)
print(timed_snapshot)
