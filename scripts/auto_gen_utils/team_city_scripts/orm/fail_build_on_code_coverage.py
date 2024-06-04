import argparse
import json
import os
import sys
import urllib3
import xml.etree.ElementTree as ET
import datetime

dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../..'))

from shared.bitbucket_utils import setup_bitbucket, get_pullrequest, make_general_comment  # noqa: ignore=F402
import shared.bitbucket_utils  # noqa: ignore=F402

CODECOV_PACKAGE_TEXT = """
{level} code coverage for {p} changed from {b} to {c}
"""

CODECOV_FAILED_HEADER_TEXT = """
build failed due to dropped code coverage. Here are the details.

"""

CODECOV_SUCCEEDED_HEADER_TEXT = """
code coverage numbers are acceptable. Here are the details.

"""

dry_run = False

ORM_CODECOV_BASELINE_URL = "https://objectstorage.us-phoenix-1.oraclecloud.com/p/MTnyrnb7DssIgX9qFoIZYoeCLScR5rD2eHWAmSyg_J8/n/dex-us-phoenix-1/b/codecov_baseline/o/orm_codecov_baseline.json"
local_json_file = "./orm_local_codecov.json"


class CoverageData:
    CONST_JACOCO_XML_PATH = "target/site/jacoco/jacoco.xml"
    CONST_INSTRUCTION_XPATH = "./counter[@type='INSTRUCTION']"
    CONST_BRANCH_XPATH = "./counter[@type='BRANCH']"
    CONST_DATE_XPATH = "./sessioninfo"
    CONST_DATE_FORMAT = "%Y-%m-%dT%H:%M:%Sz"

    def __init__(self, project_name, project_path):
        self.project_name = project_name
        self.project_path = project_path
        self.instructions_coverage = 0
        self.branch_coverage = 0
        utc = datetime.datetime.utcnow()
        self.date = utc.strftime(self.CONST_DATE_FORMAT)

    # given the report xml and xpath, return code coverage as percentage
    def __get_coverage(self, report_xml, xpath):
        missed = int(report_xml.getroot().find(xpath).get("missed"))
        covered = int(report_xml.getroot().find(xpath).get("covered"))
        total = missed + covered
        percentage = int((covered * 100) / total) if total != 0 else 0
        return percentage

    def extract_code_coverage(self):
        xml_path = os.path.join(self.project_path, self.CONST_JACOCO_XML_PATH)
        report_xml = ET.parse(xml_path)
        self.instructions_coverage = self.__get_coverage(report_xml, self.CONST_INSTRUCTION_XPATH)
        self.branch_coverage = self.__get_coverage(report_xml, self.CONST_INSTRUCTION_XPATH)
        self.date = datetime.datetime.fromtimestamp(int(report_xml.getroot().find(self.CONST_DATE_XPATH).get("start")) / 1000).strftime(self.CONST_DATE_FORMAT)


def get_codecov_baseline():
    if not dry_run:
        http = urllib3.PoolManager()
        response = http.request("GET", ORM_CODECOV_BASELINE_URL)
        return json.loads(response.data.decode("utf-8"))
    else:
        with open(local_json_file) as data_file:
            json_data = data_file.read()
            return json.loads(json_data)


def save_baseline(baseline):
    if not dry_run:
        http = urllib3.PoolManager()
        encoded_data = json.dumps(baseline).encode('utf-8')
        response = http.request(
            "PUT",
            ORM_CODECOV_BASELINE_URL,
            body=encoded_data,
            headers={'Content-Type': 'application/json'})
        return response.status
    else:
        file_name = os.path.join(local_json_file)
        with open(file_name, 'w') as f:
            f.write(json.dumps(baseline))
            f.flush()
            return 200


parser = argparse.ArgumentParser(description='fail build on code coverage')
parser.add_argument('--project-root', required=True, help="Path to the root directory of the ORM project")
parser.add_argument('--dry-run', default=False, action='store_true', required=False, help='Dry-run, do not save the data')
parser.add_argument('--build-branch', required=True, help="The value of the teamcity.build.branch variable")
parser.add_argument('--project', required=True, help="The project for which code coverage is collected")
parser.add_argument('--fail-on-code-coverage', default=False, type=lambda x: (str(x).lower() == 'true'), help="The project for which code coverage is collected")

args = parser.parse_args()
setup_bitbucket(args)
dry_run = args.dry_run
fail_on_code_coverage = args.fail_on_code_coverage
project_root = args.project_root
project = args.project

pr_id = None
try:
    # If the teamcity.build.branch variable is just something like "1234", then this is a
    # validation build for pull request "1234"
    pr_id = int(args.build_branch)
    pr = get_pullrequest("ORC", args.project, pr_id)
    print(pr.text)
except ValueError:
    # Only print this when using verbose, since we want the output be the target branch.
    print("Not a pull request validation build.")
    # sys.exit(0)

orm_project_dict = {
    "ambassador":"ambassador",
    "work-request-management":"work-request-management",
    "image-builder-tools":"image-builder-tools",
    "job-management":"job-management",
    "maestro":"maestro",
    "orm-commons":"orm-commons",
    "persistence-lib":"persistence-lib",
    "terraformer-pool-management":"terraformer-pool-management",
    "work-request-management":"work-request-management",
    "workflow":"workflow",
    "workflow-worker":"workflow-worker"
}

terraformer_dict = {
    "terraformer":"."
}

project_dict = None

if args.project == "orm":
    project_dict = orm_project_dict
elif args.project == "terraformer":
    project_dict = terraformer_dict
else:
    print("Not a valid project")
    sys.exit(2)

baseline = get_codecov_baseline()
codecov_dropped = False

for proj in project_dict:
    data = CoverageData(proj, os.path.join(project_root, project_dict[proj]))
    data.extract_code_coverage()
    print(data.instructions_coverage)
    if proj in baseline:
        package_text_instr = CODECOV_PACKAGE_TEXT.format(level="instructions", p=proj, b=baseline[proj]["instructions_coverage"], c=data.instructions_coverage)
        package_text_branch = CODECOV_PACKAGE_TEXT.format(level="branch", p=proj, b=baseline[proj]["branch_coverage"], c=data.branch_coverage)

        if(baseline[proj]["instructions_coverage"] > data.instructions_coverage):
            CODECOV_FAILED_HEADER_TEXT += package_text_instr
            codecov_dropped = True
        else:
            CODECOV_SUCCEEDED_HEADER_TEXT += package_text_instr
            print("updating instructions code coverage for {p} from {b} to {c}".format(p=proj, b=baseline[proj]["instructions_coverage"], c=data.instructions_coverage))
            baseline[proj]["instructions_coverage"] = data.instructions_coverage

        if(baseline[proj]["branch_coverage"] > data.branch_coverage):
            CODECOV_FAILED_HEADER_TEXT += package_text_branch
            codecov_dropped = True
        else:
            CODECOV_SUCCEEDED_HEADER_TEXT += package_text_branch
            print("updating instructions code coverage for {p} from {b} to {c}".format(p=proj, b=baseline[proj]["branch_coverage"], c=data.branch_coverage))
            baseline[proj]["branch_coverage"] = data.branch_coverage
    else:
        baseline[proj] = {}
        baseline[proj]["instructions_coverage"] = data.instructions_coverage
        baseline[proj]["branch_coverage"] = data.branch_coverage
    baseline[proj]["date"] = data.date

print(baseline)
codecov_message = CODECOV_FAILED_HEADER_TEXT if codecov_dropped else CODECOV_SUCCEEDED_HEADER_TEXT
val = save_baseline(baseline)
print(codecov_message)
if not dry_run:
    if pr_id:
        make_general_comment("ORC", args.project, pr_id, codecov_message)
if codecov_dropped and fail_on_code_coverage:
    sys.exit(1)
