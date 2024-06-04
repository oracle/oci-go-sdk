from jira import JIRA


class JiraClient:
    def __init__(self, server, rest_api_version, username, password, content_type='application/json', verify_certi=True):
        jira_option = {
            'server': server,
            'rest_api_version': rest_api_version,
            'verify': verify_certi,
            'headers': {
                'Content-Type': content_type
            }
        }
        self.jira_instance = JIRA(jira_option,basic_auth=(username,password))

    def get_jira_issue(self, jira_id):
        return self.jira_instance.issue(jira_id)

    def get_description(self, jira_id):
        issue = self.get_jira_issue(jira_id)
        return issue.fields.description

    def post_comment(self, jira_id, comment):
        self.jira_instance.add_comment(self.get_jira_issue((jira_id)), comment)

    def get_comment(self, jira_id):
        issue = self.get_jira_issue(jira_id)
        return issue.fields.comment.comments

    def close_pending_customer_issue(self, jira_id, closing_comment, resolution_field):
        issue = self.get_jira_issue(jira_id)
        current_state = str(issue.fields.status)
        if current_state == 'Pending Customer':
            resolved_state = "Resolve Issue  "
            self.jira_instance.transition_issue(issue, transition="Start Progress")
            self.jira_instance.transition_issue(issue, transition=resolved_state, comment=closing_comment,
                                                fields=resolution_field)

    def change_state(self, jira_id, new_state, closing_comment, state_condition):
        issue = self.get_jira_issue(jira_id)
        current_state = str(issue.fields.status)
        if current_state == "Pending Engineering" or current_state == "Open":
            self.jira_instance.transition_issue(issue, transition="Start Progress")
        if current_state != state_condition:
            self.jira_instance.transition_issue(issue, transition=new_state, comment=closing_comment)
        else:
            print("State condition check failed. Probably issue is already in the new state.")

    def get_comment_author(self, comment):
        return str(comment.author.displayName)

    def get_issue_list(self, project_id, state_list=None, update_date=None):
        query = 'project={}'.format(project_id)
        if state_list:
            state_list_str = ", ".join(['"{}"'.format(status) for status in state_list])
            query += f' and status in ({state_list_str})'
        if update_date:
            query += f' and updated<"{update_date}"'
        issues = self.jira_instance.search_issues(query)
        return issues
