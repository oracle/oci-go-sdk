import jira_client_util
import constants
import os
from datetime import datetime, timedelta


if __name__ == "__main__":

    username = os.environ[constants.JIRA_SD_USERNAME]
    password = os.environ[constants.JIRA_SD_PASSWORD]
    if constants.PROJECT_ID_ENV_VAR in os.environ:
        constants.PROJECT_ID = os.environ[constants.PROJECT_ID_ENV_VAR]

    jira = jira_client_util.JiraClient(constants.JIRA_SERVER,
                                       constants.JIRA_SERVER_REST_API_VERSION,
                                       username=username,password=password)

    print("PROJECT ID = {}".format(constants.PROJECT_ID))
    # Calculate the date 3 days ago
    three_days_ago = (datetime.now() - timedelta(days=3)).strftime('%Y-%m-%d')
    pending_customer_issue_list = jira.get_issue_list(constants.PROJECT_ID,constants.CLOSE_TICKET_STATE, three_days_ago)

    # close long pending customer ticket
    for issue in pending_customer_issue_list:
        print("Processing {}".format(issue))
        # select resolution
        resolution = {'name': 'Closed'}
        resolution_field = {'resolution': resolution}
        # closing comment
        jira.post_comment(issue, constants.CLOSE_COMMENT)
        # mark as resolved
        jira.close_pending_customer_issue(issue, "^^", resolution_field)

