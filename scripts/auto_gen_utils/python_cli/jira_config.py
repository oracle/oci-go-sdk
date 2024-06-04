import os
import urllib3

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

# JIRA REST API v2 documentation:
#  https://docs.atlassian.com/software/jira/docs/api/REST/7.8.0/
JIRA_SERVER = 'https://jira.oci.oraclecorp.com'
JIRA_SERVER_REST_API_VERSION = 2
JIRA_PROJECT = 'Developer Experience'
JIRA_PROJECT_KEY = 'DEX'
JIRA_USERNAME = "kerlee"

USERNAME = os.environ.get('JIRA_USERNAME')
PASSWORD = os.environ.get('JIRA_PASSWORD')
JSESSIONID = os.environ.get('JSESSIONID')

JIRA_OPTIONS = {
    'server': JIRA_SERVER,
    'rest_api_version': JIRA_SERVER_REST_API_VERSION,
    'verify': False
}

# Also used for Bitbucket
JIRA_BASIC_AUTH = (USERNAME, PASSWORD)

STATUS_DONE = "Done"
STATUS_CLOSED = "Closed"
STATUS_WITHDRAWN = "Withdrawn"
CLI_SELF_SERVICE_LABEL = "SelfServeManual"
CLI_MANUAL_CHANGES_LABEL = "ManualCLIChange"
CUSTOM_FIELD_ID_SDK_CLI_GA_DATE = 'customfield_13448'

DEXREQ_TERMINAL_STATES = [STATUS_DONE, STATUS_WITHDRAWN, STATUS_CLOSED]

CLI_DESIGN_REVIEW_TERMINAL_STATES = [STATUS_DONE, STATUS_CLOSED]
