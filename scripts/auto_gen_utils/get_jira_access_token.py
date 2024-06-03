from jira import JIRA

import config
import six
import sys
import getpass
import traceback


username = six.moves.input('JIRA username: ')
password = getpass.getpass('JIRA password: ')

client = JIRA(config.JIRA_OPTIONS, auth=(username, password))

session_id = client._session.cookies.get('JSESSIONID')

cookie_options = dict(config.JIRA_OPTIONS)
cookie_options['cookies'] = {
    'JSESSIONID': session_id
}

cookie_client = JIRA(cookie_options)

try:
    # make a test request to ensure this works
    cookie_client.server_info()
    cookie_client.search_issues(config.TODO_PREVIEW_CLI_TICKETS_JQL)
    print('JIRA session was valid')
except Exception as e:  # noqa:F841
    print('Error failed to authenticate with JIRA using session id')
    traceback.print_exc()
    sys.exit(1)

print('JIRA session ID: {}'.format(session_id))
