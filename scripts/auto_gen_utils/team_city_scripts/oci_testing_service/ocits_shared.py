import json
import os
import re
import sys
import urllib3
import six
import xml.etree.ElementTree as ET
from functools import reduce

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../..'))

from shared.bitbucket_utils import make_general_comment  # noqa: ignore=F402
import util  # noqa: ignore=F402
import config  # noqa: ignore=F402

global verbose
verbose = False

global dry_run
dry_run = False


NO_DEXREQ_MARKER = "NO-DEXREQ"


def printv(str):
    global verbose
    if verbose:
        print(str)


DEFAULT_JIRA_ISSUE_FIELDS = ['summary', 'description', 'labels', 'comment', 'status', 'reporter']
CUSTOM_JIRA_ISSUE_FIELDS = [
    config.CUSTOM_FIELD_ID_ARTIFACT_ID,
    config.CUSTOM_FIELD_ID_GROUP_ID,
    config.CUSTOM_FIELD_ID_ARTIFACT_VERSION,
    config.CUSTOM_FIELD_ID_SPEC_LOCATION_IN_ARTIFACT,
    config.CUSTOM_FIELD_ID_SPEC_FRIENDLY_NAME,
    config.CUSTOM_FIELD_ID_SERVICE_SUBDOMAIN,
    config.CUSTOM_FIELD_ID_FEATURE_IDS,
    config.CUSTOM_FIELD_ID_JAVA_SDK_STATUS,
    config.CUSTOM_FIELD_ID_PYTHON_SDK_STATUS,
    config.CUSTOM_FIELD_ID_RUBY_SDK_STATUS,
    config.CUSTOM_FIELD_ID_GO_SDK_STATUS,
    config.CUSTOM_FIELD_ID_CLI_STATUS,
    config.CUSTOM_FIELD_ID_SDK_CLI_GA_DATE,
    config.CUSTOM_FIELD_ID_ISSUE_ROUTING_TAG
]

HELP_URL = "https://confluence.oci.oraclecorp.com/display/DEX/Pull+Request+Validation+Builds+for+the+Testing+Service"
CODECOV_BASELINE_URL_MASTER = "https://objectstorage.us-phoenix-1.oraclecloud.com/p/5mzix52OhxjnITDKe5bdefXAwxOLpEUEta1czeh_aK4/n/dex-us-phoenix-1/b/codecov_baseline/o/codecov_master.json"
CODECOV_BASELINE_URL_PREVIEW = "https://objectstorage.us-phoenix-1.oraclecloud.com/p/y4R2h_AwvDBBG0avDuy8ZilmHWQU8MrGP2GYadWP91Y/n/dex-us-phoenix-1/b/codecov_baseline/o/codecov_preview.json"

JACOCO_XML_PATH = "{java_sdk_path}/bmc-integtests/target/jacoco.xml"

NO_DEXREQ_TICKET_REFERENCED = """
Not running the OCI testing service tests.

The pull request title or description didn't reference any DEXREQ tickets, nor did it mention `{no_dexreq_marker}`.

{author_text}Please edit the title or description to include a DEXREQ issue key, e.g. `DEXREQ-259`. If you are fixing a test or making other changes not related to a DEXREQ ticket, then include `{no_dexreq_marker}` in the ticket.

Once you have made those changes, restart the [TeamCity build]({tc_link}) by asking an SDK/CLI team member on #oci_public_sdks to click on the 'Run' button.

Example: `DEXREQ-259: {pr_title}`

Example: `{no_dexreq_marker}: {pr_title}`

For more information, see [Pull Request Validation Builds for the Testing Service]({help_url}).
"""

NO_TESTS_FOUND_BECAUSE_NO_GROUP_ARTIFACT_ID = """
Not running the OCI testing service tests.

The referenced DEXREQ {ticket_sop} did not have 'Spec Group Id' and 'Spec Artifact Id' set.

{author_text}Please edit the referenced DEXREQ {ticket_sop} ({dexreq_tickets}) and fill out the 'Spec Group Id' and 'Spec Artifact Id' fields. Then restart the [TeamCity build]({tc_link}) by asking an SDK/CLI team member on #oci_public_sdks to click on the 'Run' button.

For more information, see [Pull Request Validation Builds for the Testing Service]({help_url}).
"""


#
# Extracting DEXREQ tickets from text
#

def get_dexreq_tickets_from_text(text):
    tickets = re.findall(r"DEXREQ-\d+", text)

    unique = reduce(lambda l, x: l + [x] if x not in l else l, tickets, [])

    if not unique:
        printv("No DEXREQ tickets referenced in pull request title or description!")
    else:
        printv(unique)

    no_dexreq_marker = False
    if not unique and NO_DEXREQ_MARKER in text:
        # Only check if we didn't find any DEXREQ tickets
        printv("Found {} marker in pull request.".format(NO_DEXREQ_MARKER))
        no_dexreq_marker = True

    return unique, no_dexreq_marker


def get_dexreq_tickets(pr, tc_link):
    json = pr.json()

    if 'title' in json:
        title = json['title']
    else:
        title = ""
    if 'description' in json:
        description = json['description']
    else:
        description = ""

    tickets, no_dexreq_marker = get_dexreq_tickets_from_text(title + "\n" + description)

    if not tickets and not no_dexreq_marker:
        author_text = ""
        json = pr.json()
        if json['author'] and json['author']['user'] and json['author']['user']['name']:
            author_text = "@{name}: ".format(name=json['author']['user']['name'])

        text = NO_DEXREQ_TICKET_REFERENCED.format(
            tc_link=tc_link,
            pr_title=title,
            help_url=HELP_URL,
            author_text=author_text,
            no_dexreq_marker=NO_DEXREQ_MARKER)

        if not dry_run:
            make_general_comment("SDK", "oci-testing-service", json['id'], text)
        else:
            print("DRY-RUN: Not making BitBucket comment\n{}".format(text))

        # Don't fail the build
        sys.exit(0)

    return tickets, no_dexreq_marker


#
# JIRA
#


def get_dev_status_info_for_issue(issue):
    jira_internal_session = util.JIRA_CLIENT()._session

    issue_dev_status_url = config.JIRA_DEV_STATUS_REST_API_URL_FORMAT.format(issue.id)
    return json.loads(jira_internal_session.get(issue_dev_status_url).content)


def get_pull_requests_for_issue(issue):
    dev_status_info = get_dev_status_info_for_issue(issue)
    if dev_status_info['errors']:
        raise ValueError('There was an error retrieving pull request information for {}. Error(s): {}'.format(
            issue.key,
            json.dumps(dev_status_info['errors'])
        ))

    return dev_status_info['detail'][0]['pullRequests']


def get_jira_issue(issue_key):
    return util.get_dexreq_issue(issue_key, fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS), expand=["changelog"])


def filter_pull_requests(pull_requests, project, branch):
    filtered = []
    for x in pull_requests:
        if x['destination']['branch'] != branch:
            continue
        if x['destination']['repository']['avatarDescription'] != "SDK":
            continue
        if x['destination']['repository']['name'] != project:
            continue
        if x['status'] == "DECLINED":
            printv("Ignoring pull request {}, it has been DECLINED".format(x['id']))
            continue

        filtered.append(x)

    return filtered


def get_group_and_artifact_ids_from_jira(issue_key):
    issue = util.get_dexreq_issue(issue_key, fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS), expand=["changelog"])

    group_id = getattr(issue.fields, config.CUSTOM_FIELD_ID_GROUP_ID)
    if not group_id:
        printv("No group id in JIRA item '{}'".format(issue_key))

    artifact_id = getattr(issue.fields, config.CUSTOM_FIELD_ID_ARTIFACT_ID)
    if not artifact_id:
        printv("No artifact id in JIRA item '{}'".format(issue_key))

    return group_id, artifact_id


def get_group_and_artifact_ids(tickets, tc_link, pr):
    group_artifact_ids_set = set([])
    problems = []
    for issue_key in tickets:
        group_id, artifact_id = get_group_and_artifact_ids_from_jira(issue_key)

        if group_id and artifact_id:
            group_artifact_ids_set.add((issue_key, group_id, artifact_id))
        else:
            problems.append(issue_key)

    printv("Ticket/group/artifact ids:\n{}".format("\n".join("\t{}/{}/{}".format(a,b,c) for a, b, c in group_artifact_ids_set)))

    if problems:
        print("Tickets without group/artifact ids found!")

        author_text = ""
        json = pr.json()
        if json['author'] and json['author']['user'] and json['author']['user']['name']:
            author_text = "@{name}: ".format(name=json['author']['user']['name'])

        text = NO_TESTS_FOUND_BECAUSE_NO_GROUP_ARTIFACT_ID.format(
            tc_link=tc_link,
            dexreq_tickets=", ".join(problems),
            ticket_sop="ticket" if len(problems) == 1 else "tickets",
            ticket_has_have="has" if len(problems) == 1 else "have",
            help_url=HELP_URL,
            author_text=author_text)
        make_general_comment("SDK", "oci-testing-service", json['id'], text)

        # Don't fail the build
        sys.exit(0)

    return group_artifact_ids_set


def get_issue_routing_info_tag(issue_key):
    issue = util.get_dexreq_issue(issue_key, fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS), expand=["changelog"])

    issue_routing_tag = getattr(issue.fields, config.CUSTOM_FIELD_ID_ISSUE_ROUTING_TAG)

    if not issue_routing_tag:
        printv("No issue routing tag in JIRA item '{}'".format(issue_key))
        unique = []
    else:
        tags = [x.strip() for x in issue_routing_tag.split(",")]

        unique = reduce(lambda l, x: l + [x] if x not in l else l, tags, [])

    return unique


def get_issue_routing_info_tag_from_description(description):
    tags = re.findall("\\[IssueRoutingInfo.tag=([^]]+)\\]", description)

    unique = reduce(lambda l, x: l + [x] if x not in l else l, tags, [])

    return unique


def get_package_names_from_description(description):
    tags = re.findall("\\[RunTestsForModule=([^]]+)\\]", description)

    unique = reduce(lambda l, x: l + [x] if x not in l else l, tags, [])

    return unique


def get_master_javasdk_pr_url(description):
    results = re.findall("\\[master.pr.javasdk=([^]]+)\\]", description)

    links = reduce(lambda l, x: l + [x] if x not in l else l, results, [])

    if len(links) == 0:
        return None
    elif len(links) > 1:
        raise ValueError("More than one Java SDK master PR URL found")

    return links[0]


class CommentedTreeBuilder(ET.XMLTreeBuilder):
    def __init__(self, html=0, target=None):
        ET.XMLTreeBuilder.__init__(self, html, target)
        self._parser.CommentHandler = self.handle_comment

    def handle_comment(self, data):
        self._target.start(ET.Comment, {})
        self._target.data(data)
        self._target.end(ET.Comment)


def parse_xml(file_name):
    return ET.parse(file_name, parser=ET.XMLParser(target=CommentedTreeBuilder()))


# given the service name and the operation name (which is the name of the test),
# we return xpath that maps to the code cov report xml entry for this method
def get_method_xpath(class_name, operation):
    method_xpath = ".//class[@name='{c}']//method[@name='{m}']".format(c=class_name, m=operation)
    return method_xpath


# given a service name (package), list of operations, aggregate the code cov values for a level (such as INSTRUCTION, LINE)
# and category ( missed or covered)
def get_value_for_codecov(xml, class_name, operations, level, category):
    value = 0
    for operation in operations:
        counter_xpath_format = "{m}/counter[@type='{lvl}']".format(m=get_method_xpath(class_name, operation), lvl=level)
        cat_val = get_value_for_path_if_exists(xml, counter_xpath_format, category)
        value += cat_val
    return value


def get_codecov_baseline_url(target_branch):
    if target_branch == "master":
        return CODECOV_BASELINE_URL_MASTER
    elif target_branch == "preview":
        return CODECOV_BASELINE_URL_PREVIEW
    else:
        printv("invalid target branch: " + target_branch)
        sys.exit(2)


def get_codecov_baseline(target_branch):
    if not dry_run:
        http = urllib3.PoolManager()
        response = http.request("GET", get_codecov_baseline_url(target_branch))
        return json.loads(response.data.decode("utf-8"))
    else:
        with open("./{b}_local_codecov.json".format(b=target_branch)) as data_file:
            json_data = data_file.read()
            return json.loads(json_data)


def save_codecov_baseline(data, target_branch):
    if not dry_run:
        http = urllib3.PoolManager()
        encoded_data = json.dumps(data).encode('utf-8')
        response = http.request(
            "PUT",
            get_codecov_baseline_url(target_branch),
            body=encoded_data,
            headers={'Content-Type': 'application/json'})
        return response.status
    else:
        file_name = os.path.join(".", "{b}_local_codecov.json".format(b=target_branch))
        with open(file_name, 'w') as f:
            f.write(json.dumps(data))
            f.flush()


def get_value_for_path_if_exists(xml, package_xpath_format, category):
    value = 0
    levelNode = xml.getroot().find(package_xpath_format)
    if levelNode is not None:
        categoryAttr = levelNode.get(category)
        if categoryAttr is not None:
            value += int(categoryAttr)
        else:
            print("for path {p}, category {c} is missing, code coverage report needs investigation".format(p=package_xpath_format, c=category))
    else:
        print("path {p} is missing, skipping the values ".format(p=package_xpath_format))
    return value


# given the ops file which has json structure list of test class name, test method name, issue roting tag
# we return a dictionary that maps {service class name, tag} => list of service methods
# if the code coverage ops file doesn't exist, we return an empty map
def get_class_tag_breakdown_from_ops(ops_file):
    codecov_map = dict()
    if os.path.isfile(ops_file):
        with open(ops_file) as fp:
            json_data = fp.read()
            try:
                ops = json.loads(json_data)
                for entry in ops["Tests"]:
                    tp = (entry["Class"], entry["Tag"])
                    if tp not in codecov_map:
                        codecov_map[tp] = set()
                    codecov_map[tp].add(entry["Test"])
            except ValueError:
                print("get_class_tag_breakdown_from_ops:invalid code coverage operations file")
    return codecov_map


# given the ops file, we get project key associated with tuple (class, tag). If multiple project keys are found, we return None
# since in that case code coverage cannot be reported to a specific Jira project
def get_project_key_for_class_tag(ops_file, class_tag_pair):
    print(class_tag_pair)
    project = None
    if os.path.isfile(ops_file):
        with open(ops_file) as fp:
            json_data = fp.read()
            try:
                ops = json.loads(json_data)
                for entry in ops["Tests"]:
                    if entry["Class"] == class_tag_pair[0] and entry["Tag"] == class_tag_pair[1]:
                        if project is None:
                            project = entry["ProjectKey"]
                        elif project != entry["ProjectKey"]:
                            print("multiple project keys found for " + class_tag_pair)
                            print("cannot use code coverage data for cutting ticket")
                            return None
            except ValueError:
                print("get_project_key_for_class_tag: invalid code coverage operations file")
    return project


def get_operations_from_test_methods(test_methods):
    operations = []
    for method in test_methods:
        operations.append(method[:-(len("Test"))])
    return operations


# remove IntegrationAutoTest suffix, and concatenate Client for example
# com.oracle.bmc.email.EmailIntegrationAutoTest should become com/oracle/bmc/email/EmailClient
def get_class_from_test_name(test_class_name):
    return test_class_name[:-(len("IntegrationAutoTest"))].replace(".", "/") + "Client"


# return an array of jsons, each json will have structure something like {"class": classname, "tag": tagname, "data": {aggregated data}}
def extract_package_codecov_data_from_reportxml(xml, ops_file):
    codecov_map = get_class_tag_breakdown_from_ops(ops_file)
    service_data = []
    for tp, test_methods in six.iteritems(codecov_map):
        operations = get_operations_from_test_methods(test_methods)
        class_name = get_class_from_test_name(tp[0])
        current_data = {
            "missedInstructions": get_value_for_codecov(xml, class_name, operations, "INSTRUCTION", "missed"),
            "coveredInstructions": get_value_for_codecov(xml, class_name, operations, "INSTRUCTION","covered"),
            "missedBranches": get_value_for_codecov(xml, class_name, operations, "BRANCH", "missed"),
            "coveredBranches": get_value_for_codecov(xml, class_name, operations, "BRANCH", "covered"),
            "missedLines": get_value_for_codecov(xml, class_name, operations, "LINE", "missed"),
            "coveredLines": get_value_for_codecov(xml, class_name, operations, "LINE", "covered")
        }
        class_tag_data = {
            "class": class_name,
            "tag": tp[1],
            "data": current_data,
            "testClass": tp[0]
        }
        service_data.append(class_tag_data)
    return service_data


# return an array of jsons, each json will have structure something like {"class": classname, "tag": tagname, "data": {aggregated data}}
# from the baseline code coverage json. we match the classname/tag from ops_file to filter ther results
def extract_package_codecov_data_from_baseline(baseline_json, ops_file):
    codecov_map = get_class_tag_breakdown_from_ops(ops_file)
    codecov_data = []
    for tp in codecov_map:
        codecov_class_tag_data = list(filter(lambda class_tag: class_tag["class"] == get_class_from_test_name(tp[0]) and class_tag["tag"] == tp[1], sum(baseline_json.values(), [])))
        codecov_data.extend(codecov_class_tag_data)
    return codecov_data


# Given a list of DEXREQ tickets and a list of package names from the pull request description,
# generate the text that tells users how they can limit the tests that get run.
def get_limit_text(tickets, package_names_from_description):
    if tickets:
        return "If you want to limit the number of tests run, please edit the referenced DEXREQ {ticket_sop} ({dexreq_tickets}) and fill out the 'Issue Routing Tag' {field_sop}, or include an issue routing tag in the form of `[IssueRoutingInfo.tag=sometag]` in the description of the pull request.".format(
            dexreq_tickets=", ".join(tickets),
            ticket_sop="ticket" if len(tickets) == 1 else "tickets",
            field_sop="field" if len(tickets) == 1 else "fields")
    elif not tickets and package_names_from_description:
        return "If you want to limit the number of tests run, please include an issue routing tag in the form of `[IssueRoutingInfo.tag=sometag]` in the description of the pull request."
    else:
        return None


# Given a list of DEXREQ tickets and a list of package names from the pull request description,
# generate the text that tells users how the tests to be run were determined.
def get_determined_text(tickets, package_names_from_description, test_classes):
    if tickets and package_names_from_description:
        return "{this_class_was_these_classes_were} determined using the 'Spec Group Id' and 'Spec Artifact Id' values set in the referenced DEXREQ {ticket_sop} {dexreq_tickets} and the `[RunTestsForModule=xyz]` {annotation_sop} in the pull request description.".format(
            this_class_was_these_classes_were="This class was" if len(test_classes) == 1 else "These classes were",
            dexreq_tickets=", ".join(tickets),
            ticket_sop="ticket" if len(tickets) == 1 else "tickets",
            annotation_sop="annotation" if len(package_names_from_description) == 1 else "annotations")
    elif tickets and not package_names_from_description:
        return "{this_class_was_these_classes_were} determined using the 'Spec Group Id' and 'Spec Artifact Id' values set in the referenced DEXREQ {ticket_sop} {dexreq_tickets}.".format(
            this_class_was_these_classes_were="This class was" if len(test_classes) == 1 else "These classes were",
            dexreq_tickets=", ".join(tickets),
            ticket_sop="ticket" if len(tickets) == 1 else "tickets")
    elif not tickets and package_names_from_description:
        return "{this_class_was_these_classes_were} determined using the {annotations} {annotation_sop} in the pull request description.".format(
            this_class_was_these_classes_were="This class was" if len(test_classes) == 1 else "These classes were",
            dexreq_tickets=", ".join(tickets),
            ticket_sop="ticket" if len(tickets) == 1 else "tickets",
            annotations=", ".join("`[RunTestsForModule={}]`".format(p) for p in package_names_from_description),
            annotation_sop="annotation" if len(package_names_from_description) == 1 else "annotations")
    else:
        return None
