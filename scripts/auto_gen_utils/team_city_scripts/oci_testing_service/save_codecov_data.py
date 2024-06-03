#
# we get called by the scheduled job that runs tests for a service.
# we locate the jacoco report xml, extract the code coverage numbers and save the information
# to a predefined object storage location
#
import argparse
import json
import re
import datetime
import ocits_shared
from ocits_shared import JACOCO_XML_PATH, parse_xml, get_codecov_baseline, save_codecov_baseline, extract_package_codecov_data_from_reportxml
from jira_ticket_reporter import JiraTicketReporter

CODECOV_REPORT_URL = "https://teamcity.oci.oraclecorp.com/repository/download/{build_config_id}/{build_id}:id/jacocoreport/index.html"
#
# xml_file_path - code coverage xml file expected at a specific location
# package - the name of the service such as email
# target_branch - the branch the tests were run against - currently master or preview
# we get the saved baseline data from object storage, we extract the current code coverage data
# from the report xml and we update the object storage with this information
# if build_config_id is provided, we use it as is to generate the URL for code coverage report else
# we try to guess the URL using build_config_name
#


def get_percent_tests_passed(ops_file):
    with open(ops_file) as fp:
        json_data = fp.read()
        ops = json.loads(json_data)
        summary = ops["Summary"]
        return int((summary["Tests run"] - (summary["Skipped"] + summary["Failures"])) * 100 / summary["Tests run"])


# apart from saving data to PAR, we also now file bug if code coverage drops
def save_codecov_data_from_reportxml(xml_file_path, package, ops_file, target_branch, build_config_name, build_config_id, build_id, expected_coverage_percentage, dry_run):
    report_xml = parse_xml(xml_file_path)
    baseline_json = get_codecov_baseline(target_branch)
    class_tag_list = extract_package_codecov_data_from_reportxml(report_xml, ops_file)
    print(class_tag_list)
    utc = datetime.datetime.utcnow()
    for class_tag in class_tag_list:
        current_data = class_tag["data"]
        current_data["url"] = get_codecov_report_url(build_config_name, build_config_id, build_id)
        current_data["timestamp"] = utc.strftime("%Y-%m-%dT%H:%M:%Sz")
        current_data["testsPassedPercentage"] = get_percent_tests_passed(ops_file)
        # project_key = get_project_key_for_class_tag(ops_file, (class_tag["testClass"], class_tag["tag"]))
        # https://jira.oci.oraclecorp.com/browse/APOLLO-2549 : Change Project key to cut ticket for Apollo
        log_jira_if_codecov_drops(class_tag, "APOLLO", expected_coverage_percentage, dry_run)
    baseline_json[package] = class_tag_list
    error_code = save_codecov_baseline(baseline_json, target_branch)
    print("error code for saving baseline: " + str(error_code))


# compare current instructions coverage percentage if below the specified bar, open/update jira ticket for the specified project key
def log_jira_if_codecov_drops(class_tag, project_key, expected_coverage_percentage, dry_run):
    total = (class_tag["data"]["coveredInstructions"] + class_tag["data"]["missedInstructions"])
    current_coverage = 100 * class_tag["data"]["coveredInstructions"] / total if total != 0 else 100
    if current_coverage < expected_coverage_percentage:
        reporter = JiraTicketReporter()
        package_label = "class: {c}, tag: {t}".format(c=class_tag["class"], t=class_tag["tag"])
        reporter.report_codecov_to_jira_ticket(project_key, class_tag["tag"], package_label, expected_coverage_percentage, current_coverage, class_tag["data"]["url"], dry_run)
    print("current coverage number {value}, expected: {baseline}".format(value=current_coverage, baseline=expected_coverage_percentage))


#
# try to construct build configuration id using build configuration name by following rules
# replace (string1) by _string1
# strip all non alphabet characters
#
def get_codecov_report_url(build_config_name, build_config_id, build_id):
    if not build_config_id:
        build_config_name = build_config_name.replace("(", "_")
        build_config_name = build_config_name.replace(")", "")
        regex = re.compile('[^a-zA-Z_]')
        build_config_id = "Sdk_OCI_Testing_Service_" + regex.sub('', build_config_name)

    return CODECOV_REPORT_URL.format(build_config_id=build_config_id, build_id=build_id)


parser = argparse.ArgumentParser(description='save test results')
parser.add_argument('--java-sdk-path', required=True, help="Path to the root directory of the Java SDK")
parser.add_argument('-p', '--package', required=True, help='service package name')
parser.add_argument('--expected-coverage-percentage', default=0, type=int, help='expected code coverage percentage, opens ticket if coverage is below this number')
parser.add_argument('--codecov-operations-file', required=True, help='list of operations to consider for code coverage')
parser.add_argument('-t', '--target-branch', required=True, help='target branch - preview or master')
parser.add_argument('--build-configuration-id', required=False, help="ID of the build configuration")
parser.add_argument('--build-configuration-name', required=True, help="Name of the build configuration")
parser.add_argument('--build-id', required=True, help="build id of this build")
parser.add_argument('--dry-run', default=False, action='store_true', required=False, help='Dry-run, do not save the data')

args = parser.parse_args()
ocits_shared.dry_run = args.dry_run
xml_file_path = JACOCO_XML_PATH.format(java_sdk_path=args.java_sdk_path)
save_codecov_data_from_reportxml(xml_file_path, args.package, args.codecov_operations_file, args.target_branch, args.build_configuration_name, args.build_configuration_id, args.build_id, args.expected_coverage_percentage, args.dry_run)
