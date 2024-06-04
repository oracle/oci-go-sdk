import jira_client_util
import re
import json
import constants
import os
import csv
from itertools import takewhile


def snake_to_camel(snake_case):
    """

    :param snake_case:
    :return: camel case converted string
    """
    if '_' not in snake_case:
        return snake_case
    temp = snake_case.split('_')
    ret = ''.join(ele.title() for ele in temp)
    camel_case = str(ret)
    return camel_case


def get_queue_name_from_oci_cli_cmd_name(service_name):
    """
    fetches the jira-sd name for given oci cli command
    in the service-jira-phonebook-mapping.csv file.
    :param cmd: key in auto-route-jira-sd/auto_suggest.json
    :return: correct jira-sd queue name
    """
    with open('{}/{}'.format(constants.PATH_TO_CSV, constants.OCI_CLI_CMD_JIRA_SD)) as json_file:
        data = json.load(json_file)
        print(service_name)
        if service_name in data:
            return data[service_name]
        else:
            return None


def get_queue_name_from_api_tool_name(api_tool_name):
    """
    fetches the jira-sd name for given api_tool_name
    in the service-jira-phonebook-mapping.csv file.
    :param api_tool_name: as mentioned in service-jira-phonebook-mapping.csv
    :return: correct jira-sd queue name
    """
    jira_csv_mapping_file = constants.PATH_TO_CSV + constants.JIRA_MAP_CSV
    reader = csv.DictReader(open(jira_csv_mapping_file))
    for row in reader:
        if row['API/Tool'] == api_tool_name:
            print("Corresponding jira-sd: " + row["Jira-sd"])
            return row["Jira-sd"]
    return None


def get_suggestion_string_from_oci_cli_command(msg):
    try:
        for line in msg.split("\n"):
            if "oci " in line:
                line.lstrip().rstrip()
                cmd_list = line.split()
                # As command can be in format -'command = oci os ns get' ,so finding position of oci keyword
                oci_position = cmd_list.index('oci')
                # if 'oci' is present
                if oci_position != -1:
                    actual_cmd_list = cmd_list[oci_position:]
                    # remove all authentication params from command
                    for param in constants.AUTHENTICATION_PARAMS:
                        if param in actual_cmd_list:
                            param_idx = actual_cmd_list.index(param)
                            # ex- --auth abc
                            del actual_cmd_list[param_idx: param_idx + 2]
                if actual_cmd_list[0].lstrip().rstrip() == 'oci' and len(actual_cmd_list) > 2:
                    service_name = actual_cmd_list[1].lstrip().rstrip()
                    jira_sd = get_queue_name_from_oci_cli_cmd_name(service_name)
                    if jira_sd is not None:
                        suggestion = constants.JIRA_SD_MISROUTE_MESSAGE + "Queue: " + jira_sd.upper()
                    else:
                        suggestion = constants.JIRA_SD_MISROUTE_MESSAGE + "Service Team: " + service_name
                    print("suggestion string: " + suggestion)
                    return suggestion
                return None
    except Exception:
        print("Unable to parse jira-sd comment")
        return None


def get_suggestion_string_from_service_error_code(target_service='', operation_name='', error_code='',message=''):
    """
    If error_code is 401 then, ask customer to follow troubleshooting guide, still
    if the issue persists, then move it to service team.
    For other cases, (non 401 code), fetch appropriate jira-sd queue & suggest moving
    the issue in it.
    :param target_service: as per the service error
    :param operation_name: as per the service error
    :param error_code: as per the service error
    :return: The suggestion for the customer to be posted on jira-sd.
    """
    suggestion = ""
    # if 401 add suggestion string
    if error_code == "401" or 'status: 401' in message:
        suggestion = constants.SUGGESTION_MESSAGE

    # if operation_name and target service is present
    if operation_name != '' and target_service != '':
        operation_name = snake_to_camel(operation_name)
        api_tool_name = target_service.lower() + '_' + operation_name
        print(api_tool_name)
        jira_sd = get_queue_name_from_api_tool_name(api_tool_name)
        if jira_sd is not None:
            suggestion = suggestion + constants.JIRA_SD_MISROUTE_MESSAGE + "Queue: " + \
                         jira_sd.upper()
        else:
            suggestion = suggestion + constants.JIRA_SD_MISROUTE_MESSAGE + "Service: " + \
                                    target_service

    else:
        res = get_suggestion_string_from_oci_cli_command(message)
        # if command is provided suggest to move to right queue.
        if res is not None:
            suggestion = suggestion + res
        # if command is not provided in description,ask for it.
        else:
            suggestion = suggestion + constants.AUTOMATION_SEARCH_STRING + constants.JIRA_SD_ALTERNATE_MESSAGE
    print("suggestion string: " + suggestion)
    return suggestion


def get_client_version(comment):
    version_line_pattern = r'.*client[-_\s]{0,1}version.*?:'
    quote_or_space = r'\\\'\" '
    quote_or_space_or_newline_or_comma = r',\'\"$ '
    for line in comment.splitlines():
        match_object = re.search(version_line_pattern, line.lower())
        if match_object:
            line = line[match_object.end():].lstrip(quote_or_space)
            until_word = list(
                takewhile(lambda x: x not in quote_or_space_or_newline_or_comma, line))
            return ''.join(until_word)
    return None


def parse_python_error(error_string):
    operation_pattern = "operation_name: (.*?),"
    service_pattern = "target_service: (.*?),"
    status_pattern = "status: (.*?),"
    operation = re.search(operation_pattern, error_string)
    service = re.search(service_pattern, error_string)
    status = re.search(status_pattern, error_string)
    if operation is not None:
        operation = operation.group(1)
    if service is not None:
        service = service.group(1)
    if status is not None:
        status = status.group(1)
    else:
        status = ""
    return {constants.ERROR_FIELD_OPERATION_NAME: operation,
            constants.ERROR_FIELD_TARGET_SERVICE: service,
            constants.ERROR_FIELD_STATUS: status}


def parse_java_error(error_string):
    operation_pattern = "Error returned by (.*) operation"
    service_pattern = "operation in (.*) service"
    operation = re.search(operation_pattern, error_string)
    service = re.search(service_pattern, error_string)
    if operation is not None:
        operation = operation.group(1).strip()
    if service is not None:
        service = service.group(1).strip()
    return {
        constants.ERROR_FIELD_OPERATION_NAME: operation,
        constants.ERROR_FIELD_TARGET_SERVICE: service}


def parse_go_error(error_string):
    operation_pattern = "Operation Name: (.*)"
    service_pattern = "Error returned by (.*) Service\."
    operation = re.search(operation_pattern, error_string)
    service = re.search(service_pattern, error_string)
    if operation is not None:
        operation = operation.group(1).strip()
    if service is not None:
        service = service.group(1).strip()
    return {constants.ERROR_FIELD_OPERATION_NAME: operation,
            constants.ERROR_FIELD_TARGET_SERVICE: service}


def parse_ts_error(error_string):
    operation_pattern = "operationName: (.*?),"
    service_pattern = "targetService: (.*?),"
    operation_name = re.search(operation_pattern, error_string).group(1)
    if operation_name is not None:
        operation_name_char_arr = list(operation_name)
        operation_name_char_arr[0] = operation_name_char_arr[0].upper()
        operation_name = ''.join(operation_name_char_arr)
        operation_name = operation_name.strip()
    service = re.search(service_pattern, error_string)
    if service is not None:
        service = service.group(1).strip()
    return {constants.ERROR_FIELD_OPERATION_NAME: operation_name,
            constants.ERROR_FIELD_TARGET_SERVICE: service}


def find_error_and_parse(error_string):
    error_dict = {}
    if all_keywords_in_error(constants.KEYWORDS, error_string.lower()):
        client_version = get_client_version(error_string)
        print(client_version)
        if client_version is not None:
            if "Oracle-JavaSDK" in client_version:
                error_dict = parse_java_error(error_string)
            elif "Oracle-GoSDK" in client_version or "Oracle-DotNetSDK" in client_version:
                error_dict = parse_go_error(error_string)
            elif "Oracle-TypeScriptSDK" in client_version:
                error_dict = parse_ts_error(error_string)
            elif "Oracle-PythonSDK" in client_version or "":
                error_dict = parse_python_error(error_string)
            if error_dict == {}:
                print("Service error is not provided by the reporter. No actions taken")
                return None
            key_list_err_dict = list(error_dict.keys())
            if error_dict[constants.ERROR_FIELD_OPERATION_NAME] in constants.EMPTY or \
                    error_dict[constants.ERROR_FIELD_TARGET_SERVICE] in constants.EMPTY:
                print("Could not find the operation_name or target_service")
                return None
            if constants.ERROR_FIELD_STATUS not in key_list_err_dict:
                error_dict[constants.ERROR_FIELD_STATUS] = ""
            return get_suggestion_string_from_service_error_code(
                error_dict[constants.ERROR_FIELD_TARGET_SERVICE],
                error_dict[constants.ERROR_FIELD_OPERATION_NAME],
                error_dict[constants.ERROR_FIELD_STATUS])
        else:
            print("Client version not provided  ")
    return None


def search_in_comments(jira, jira_id):
    """
    Iterate through all comments in the ticket, and search the service error
    text in them.
    """
    # first, search in description
    description = jira.get_description(str(jira_id))
    if description is not None:
        description = description.replace(" ", ' ').replace('{', '').replace('}', '').\
            replace('\'', '').replace('"', '')
        res = find_error_and_parse(description)
        if res is not None:
            return res
        elif all_keywords_in_error(constants.SERVICE_ERROR_KEYWORDS, description):
            return get_suggestion_string_from_service_error_code(description)

    # Second, search in comments
    comment_list = jira.get_comment(str(jira_id))
    for comment in reversed(comment_list):
        error_string = comment.body.replace(" ", ' ').replace('{', '').replace('}', '').replace(
            '\'', '').replace('"', '')
        res = find_error_and_parse(error_string)
        if res is not None:
            return res
        # If customer didn't follow SR template or customer is using old oci-cli version, try to find CLI command and
        # find correct jira-sd queue
        elif all_keywords_in_error(constants.SERVICE_ERROR_KEYWORDS, comment.body) and jira.get_comment_author(
                comment) != constants.AUTHOR:
            print(
                "Customer may be using old version of OCI-CLI. Service Error found in jira-sd comment")
            return get_suggestion_string_from_service_error_code(error_string)
    return None


def all_keywords_in_error(keywords, error_string):
    for keyword in keywords:
        if keyword not in error_string:
            return False
    return True

def get_prefix_label(project_id):
    if project_id == constants.CLI_PROJECT_ID:
        return constants.CLI_LABEL_PREFIX
    else:
        return constants.SDK_LABEL_PREFIX

def add_jira_sd_labels(issue, labels, prefix = ''):
    current_labels = issue.fields.labels if issue.fields.labels else []
    for label_to_add in labels:
        if label_to_add not in current_labels:
            current_labels.append(prefix + label_to_add)
    issue.update(fields={'labels': current_labels})

def is_already_processed(jira, jira_id):
    comment_list = jira.get_comment(str(jira_id))
    # Iterate through all comments in the ticket, search for Automation comment or author
    for comment in comment_list:
        if constants.AUTOMATION_SEARCH_STRING in comment.body:
            print("This ticket {} is already processed by automation".format(jira_id))
            return True
        # This condition should work when we have a jira-sd bot as author
        if jira.get_comment_author(comment) == constants.THIS_BOT_AUTHOR:
            print("This ticket {} is already processed by automation".format(jira_id))
            return True
    return False


if __name__ == "__main__":

    username = os.environ[constants.JIRA_SD_USERNAME]
    password = os.environ[constants.JIRA_SD_PASSWORD]
    if constants.PROJECT_ID_ENV_VAR in os.environ:
        constants.PROJECT_ID = os.environ[constants.PROJECT_ID_ENV_VAR]

    jira = jira_client_util.JiraClient(constants.JIRA_SERVER,
                                       constants.JIRA_SERVER_REST_API_VERSION,
                                       username=username,password=password)

    print("PROJECT ID = {}".format(constants.PROJECT_ID))
    jira_issue_list = jira.get_issue_list(constants.PROJECT_ID, constants.JIRA_BOT_PROCESSING_STATES)
    for issue in jira_issue_list:
        print("Processing {}".format(issue))
        if is_already_processed(jira, issue):
            continue
        reply_comment = search_in_comments(jira, issue)
        if reply_comment is None:
            print("Could not find Jira-sd queue name for jira: {} ".format(issue))
            continue
        else:
            # Its redundant, to be removed after state transition workflow is stable.
            if constants.JIRA_SD_ALTERNATE_MESSAGE not in reply_comment:
                prefix  = get_prefix_label(constants.PROJECT_ID)
                add_jira_sd_labels(issue, constants.LABELS, prefix)
                # add common label
                add_jira_sd_labels(issue, constants.COMMON_LABEL)
            jira.post_comment(issue, reply_comment)
            # move the state to pending customer.
            jira.change_state(issue, constants.PENDING_CUSTOMER_STATE, "^^",
                              constants.PENDING_CUSTOMER_STATE_CHECK)

