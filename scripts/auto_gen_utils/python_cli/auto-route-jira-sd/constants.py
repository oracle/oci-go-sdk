JIRA_SD_TOKEN_ENV_VAR = 'JIRA_SD_TOKEN'
KEYWORDS = ["operation", "service", "client", "version"]
AUTOMATION_SEARCH_STRING = "This is an auto-generated message from JIRA BOT."
JIRA_SD_MISROUTE_MESSAGE = "{} Error message indicates that this issue is happening at the " \
                           "service end. Please move this ticket to ".format(AUTOMATION_SEARCH_STRING)
JIRA_SD_ALTERNATE_MESSAGE = "Please provide more details (example-command used,version you are using). \n\n"
EMPTY = [None, '']
JIRA_SD_FOUND = 0
PENDING_CUSTOMER_STATE = "PENDING CUSTOMER"
PENDING_CUSTOMER_STATE_CHECK = "Pending Customer"
PATH_TO_CSV = "../python-cli/scripts/doc_gen/issue_routing/"
JIRA_MAP_CSV = "service_jira_phonebook_mapping.csv"
OCI_CLI_CMD_JIRA_SD = "oci_cli_cmd_jira_sd.json"
JIRA_SERVER = "https://jira-sd.mc1.oracleiaas.com"
JIRA_SERVER_REST_API_VERSION = 2
TOKEN = ""
PROJECT_ID = "27401"
AUTHOR = "Jira Automation Bot"
ERROR_FIELD_TARGET_SERVICE = "target_service"
ERROR_FIELD_OPERATION_NAME = "operation_name"
ERROR_FIELD_STATUS = "status"
SERVICE_ERROR_KEYWORDS = ["ServiceError:","opc-request-id"]
AUTHENTICATION_PARAMS = ["--profile", "--config-file", "--auth", "--region", "--endpoint", "--cert-bundle", "--auth-purpose", "--cli-rc-file"]
PROJECT_ID_ENV_VAR = "PROJECT_ID"
THIS_BOT_AUTHOR = "auto-suggest-jira-sd-sdk"
SUGGESTION_MESSAGE = AUTOMATION_SEARCH_STRING + " Please first troubleshoot using this doc " \
                             "https://docs.oracle.com/en-us/iaas/Content/API/References" \
                             "/apierrors.htm#apierrors_401 If issue still exists, please try the next " \
                             "suggestion. \n\n"
JIRA_SD_USERNAME="JIRA_SD_USERNAME"
JIRA_SD_PASSWORD="JIRA_SD_PASSWORD"
SDK_PROJECT_ID= "11001"
CLI_PROJECT_ID= "27401"
SDK_LABEL_PREFIX = "CD_SDK_"
CLI_LABEL_PREFIX = "CD_CLI_"
LABELS = ["Misrouted", "SERVICE_ERR_BOT"]
COMMON_LABEL=["DS_MISROUTED"]
CLOSE_COMMENT = AUTOMATION_SEARCH_STRING + " Closing the ticket as there is no reponse from the customer. " \
                                           "Please reopen the ticket if any further help is needed."
CLOSE_TICKET_STATE = ["Pending Customer"]
JIRA_BOT_PROCESSING_STATES = ["Pending Engineering", "Open", "In Progress"]