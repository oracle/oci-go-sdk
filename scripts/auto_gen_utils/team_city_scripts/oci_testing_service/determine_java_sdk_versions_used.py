from __future__ import print_function
import argparse
import os
import ssl
import sys
import urllib3
import xml.etree.ElementTree as ET
from glob import glob

import ocits_shared
from ocits_shared import parse_xml

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../..'))

import util  # noqa: ignore=F402
import config  # noqa: ignore=F402

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
if (not os.environ.get('PYTHONHTTPSVERIFY', '') and getattr(ssl, '_create_unverified_context', None)): 
    ssl._create_default_https_context = ssl._create_unverified_context


ns = {"ns":"http://maven.apache.org/POM/4.0.0"}

# allow default namespace for output, dont print ns0: prefixes everywhere
ET.register_namespace('',"http://maven.apache.org/POM/4.0.0")


def get_group_artifact_versions_from_pom_file(path):
    dependencies = []

    group_id = "com.oracle.oci.sdk"

    pom_files = [y for x in os.walk(path) for y in glob(os.path.join(x[0], 'pom.xml'))]
    for ldr_path in pom_files:
        pom = parse_xml(ldr_path)
        # Find all the places where <dependency><codegen.artifactory.groupId> and <properties><codegen.artifactory.artifactId> both match
        xpath = './/ns:dependency[ns:groupId="{}"]'.format(group_id)
        dependency_nodes = pom.findall(xpath, ns)
        for node in dependency_nodes:
            artifact_id = None
            version = None

            for child in node:
                tag = child.tag.replace('{{{}}}'.format(ns['ns']), "")
                if tag == "artifactId":
                    artifact_id = child.text
                elif tag == "version":
                    version = child.text

            if artifact_id and version:
                dependencies.append((group_id, artifact_id, version))

    return dependencies


#
# Parameters variable set up
#
parser = argparse.ArgumentParser(description='Determine the Java SDK versions used.')
parser.add_argument('--oci-testing-service-path', required=True, help="Path to the root directory of the OCI Testing Service")
parser.add_argument('-o', '--output', required=False, help='Output file')
parser.add_argument('-v', '--verbose', action='store_true', default=False, required=False, help='Verbose output')

args = parser.parse_args()

if args.verbose:
    ocits_shared.verbose = True

dependencies = get_group_artifact_versions_from_pom_file(args.oci_testing_service_path)

dependencies_str = "\n".join("{}:{}:{}".format(g, a, v) for g, a, v in dependencies)
print(dependencies_str)

if args.output:
    f = open(args.output, "w")
    f.write(dependencies_str)
