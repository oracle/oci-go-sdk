from jira import JIRA
from config import PREVIEW_ISSUE_TYPE_ID
import datetime
import python_cli.jira_config as config
import python_cli.self_service_manual_change as manual_change
import sys
import re
from python_cli.exceptions import CliException

_JIRA_CLIENT = None

BOT_USERNAME = "[~gear-dexreq-automation]"
BOT_FULL_NAME = "DEXREQ Automation"
BOT_ID = "gear-dexreq-automation"
GENERATED_BRANCH_REGEX = re.compile(r"(?<=sourceBranch=refs%2Fheads%2F)[\s\S]+?(?=&title=)")
VALID_CHARS = "[^\d\w\s\[\]@\-\>]"  # noqa: W605
MANUAL_CHANGE_MATCH = "(\[{}(\w+)?]\s)([^\[]+)"  # noqa: W605
MANUAL_CHANGE_DELIMTER = "->"

OPEN_DESIGN_REVIEWS = 'project="Developer Experience" AND issuetype="Design Reviews" AND status not in (Done, Closed, Blocked)'
NO_MANUAL_CHANGES_REQUIRED = '.*(no).*(change).*(require).*'  # noqa: W605

NO_MANUAL_CHANGES_TICKET_CLOSED = """
[~{user}]
This Design Review ticket has been marked *Done* by DEXREQ Automation because "No changes required" was detected.

If this is a mistake or if you need further assistance, please assign the ticket to the CLI Support engineer.
"""

MANUAL_CHANGE_COMPLETED_TEMPLATE = """Manual Change Request Done

You can find your changes here:
https://bitbucket.oci.oraclecorp.com/projects/SDK/repos/python-cli/pull-requests?create&targetBranch=refs%2Fheads%2Fpreview&sourceBranch=refs%2Fheads%2F{}&targetRepoId=930

You can choose to create a PR, or use the branch, {}, as a base for any additional changes.

If you would like to add more changes on created branch by CLI bot, Follow below steps: 
1. Please refer to "Fork Model for Pull Requests" on how to utilize Bitbucket forks for development: https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=132774928#SDK/CLISelfServiceFrequentlyAskedQuestions-HowdoIdoapullrequestforSDK/CLIrepos?
2. Install Development env of OCI-CLI: https://confluence.oci.oraclecorp.com/display/DEX/Installing+OCI-CLI+using+Python3
3. Checkout to preview-DEX-<TICKET-NUMBER> branch after forking the repo
4. make changes
5. cd services/<your-service>
6. flake8 --ignore=F841,E501,W503 service/<Your Service Name>
7. run make docs inside service/<Your Service Name> to generate all docs
8. git add, commit and push changes to branch
9. create PR against preview branch

If you create a PR, post the PR link on this ticket after both builds have passed successfully
PR should be created against preview branch and approved by at least one of your teammates
Finally  assign the ticket to the CLI Support engineer
CLI support engineer will mark the ticket Done
Next step would be to create a Public DEXREQ ticket https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=43683000

If making additional manual changes, see the docs below:
Making CLI Manual changes: https://confluence.oci.oraclecorp.com/display/DEX/CLI+Manual+Code+Changes
Common manual change recipes: https://confluence.oci.oraclecorp.com/display/DEX/CLI+recipes+for+overriding+generated+code
Recommended CLI installation for Mac: https://confluence.oci.oraclecorp.com/display/DEX/Installing+OCI-CLI+using+Python3
"""

MANUAL_CHANGE_FAILED_TEMPLATE = """Manual Change Request Failed. {}

The manual changes will have to be done manually.
The latest status of the changes can be found here: https://bitbucket.oci.oraclecorp.com/projects/SDK/repos/python-cli/pull-requests?create&targetBranch=refs%2Fheads%2Fpreview&sourceBranch=refs%2Fheads%2F{}&targetRepoId=930
You can checkout this changes from this branch: {}

To make additional changes on top of this branch, or to make manual changes in general, see the docs below:
Making CLI Manual changes: https://confluence.oci.oraclecorp.com/display/DEX/CLI+Manual+Code+Changes
Common manual change recipes: https://confluence.oci.oraclecorp.com/display/DEX/CLI+recipes+for+overriding+generated+code
Recommended CLI installation for Mac: https://confluence.oci.oraclecorp.com/display/DEX/Installing+OCI-CLI+using+Python3
"""

PREVIEW_NOT_DONE_TEMPLATE = """Manual Change Request Rejected.
Your manual changes cannot be generated until your preview ticket is *Done*.
Please request another manual change after your ticket is *Done*.
"""

NEW_DESIGN_REVIEW_TEMPLATE = """[~{user}]
Please review the generated commands above.
If no changes are required, please comment "No changes required" and assign the ticket to the CLI Support engineer.
If no changes were generated, please confirm that it is intended.

Otherwise, please follow these links if you would like to manually change generated CLI commands.

To make manual changes using our CLI Self Service:
    - Please make sure your Preview DEXREQ ticket is in Done status before making a manual change. Do NOT change it yourself.
    - Follow https://confluence.oci.oraclecorp.com/display/DEX/CLI+Self+Service+Guide and leave a comment using the appropriate template
    - Our CLI bot will read "[~gear-dexreq-automation] Manual Changes Requested" in comments and process the comment and create a branch automatically
    - Only last comment will be used to create a branch
    - Follow the comment on the ticket after branch is created

To make additional changes on top of this branch, or to make general manual changes, see the docs below:
Making CLI Manual changes: https://confluence.oci.oraclecorp.com/display/DEX/CLI+Manual+Code+Changes
Common manual change recipes: https://confluence.oci.oraclecorp.com/display/DEX/CLI+recipes+for+overriding+generated+code
Recommended CLI installation for Mac: https://confluence.oci.oraclecorp.com/display/DEX/Installing+OCI-CLI+using+Python3
"""

UPCOMING_GA_DATE_TEMPLATE = """[~{user}]
Your GA date is coming up in {days} days.
This Design Review ticket must be done a week before your GA date.
To ensure your feature is ready for release, please ensure all work for this ticket is done and this ticket is closed.
"""

RENAME_ROOT_GROUP = "RENAME ROOT GROUP"
RENAME_COMMAND = "RENAME COMMAND"
REMOVE_COMMAND = "REMOVE COMMAND"
MOVE_COMMAND = "MOVE COMMAND"
MOVE_GROUP = "MOVE GROUP"
RENAME_PARAMETER = "RENAME PARAMETER"
REMOVE_PARAMETER = "REMOVE PARAMETER"
FLATTEN_PARAMETER = "FLATTEN PARAMETER"

PREVIEW_DONE_STATE = "Done"

SUPPORTED_MANUAL_CHANGES = [RENAME_COMMAND, REMOVE_COMMAND, MOVE_COMMAND, MOVE_GROUP, RENAME_PARAMETER, REMOVE_PARAMETER]

RENAME_ROOT_GROUP_EXCEPTION = """Failed in renaming the root group {line} , and error is  {e}.Check if the {line} is present in the latest preview oci-cli  branch"""

RENAME_COMMAND_EXCEPTION = """"Failed in renaming the commands {line} , and error is  {e}. Check if the {line} is present in the latest preview oci-cli  branch"""

MOVE_GROUP_OR_COMMANDS_EXCEPTION = """Failed in moving the commands {line} , and error is  {e}. Check if both old and new groups are present in the latest preview oci-cli  branch for this line {line}. CLI bot cannot create a new group, it can only move commands to existing groups"""

RENAME_PARAMETER_EXCEPTION = """Failed in renaming the parameter the commands and error is  {e}"""

raised_exceptions = {}


def JIRA_CLIENT():
    global _JIRA_CLIENT

    if _JIRA_CLIENT:
        return _JIRA_CLIENT

    # attempt to log in using user name and password if present, if not use config.JSESSIONID
    if config.USERNAME and config.PASSWORD:
        print('Building JIRA client with username / password auth')
        _JIRA_CLIENT = JIRA(config.JIRA_OPTIONS, basic_auth=config.JIRA_BASIC_AUTH)
    elif config.JSESSIONID:
        print('Building JIRA client with cookie based auth')
        cookie_options = dict(config.JIRA_OPTIONS)
        cookie_options['cookies'] = {
            'JSESSIONID': config.JSESSIONID
        }

        _JIRA_CLIENT = JIRA(cookie_options)
    else:
        sys.exit('Could not authenticate with JIRA server. Must specify environment variables for either config.JSESSIONID or JIRA_USERNAME and JIRA_PASSWORD.')

    return _JIRA_CLIENT


def is_design_ticket_in_non_terminal_state(issue):
    return issue and issue.fields.status and issue.fields.status.name not in config.CLI_DESIGN_REVIEW_TERMINAL_STATES


def get_open_design_review_tickets():
    # Results default to 50, need to set to a higher arbitrary number to return all results
    design_reviews = JIRA_CLIENT().search_issues(OPEN_DESIGN_REVIEWS, maxResults=1000)
    return design_reviews


def add_jira_comment(issue_key, comment):
    JIRA_CLIENT().add_comment(issue_key, comment)


def get_jira_issue(issue_key):
    return JIRA_CLIENT().issue(issue_key)


# Returns "preview" if the Preview SDK/CLI ticket is closed otherwise None
def get_branch_with_changes(design_review):
    if check_if_preview_is_done(design_review):
        return "preview"
    else:
        add_jira_comment(design_review, PREVIEW_NOT_DONE_TEMPLATE)
        return None


# Checks if the preview ticket associated with the design review ticket is closed
def check_if_preview_is_done(design_review):
    if hasattr(design_review.fields, 'issuelinks'):
        for link in design_review.fields.issuelinks:
            if hasattr(link, 'inwardIssue'):
                issue = JIRA_CLIENT().issue(link.inwardIssue.key, fields='description, summary, status, issuetype')
                ticket_type_id = issue.fields.issuetype.id
                if ticket_type_id and ticket_type_id == PREVIEW_ISSUE_TYPE_ID:
                    print('Found Preview SDK issue: {}'.format(issue.key))
                    if PREVIEW_DONE_STATE in issue.fields.status.name:
                        return True
    return False


# Checks if the ticket is a Preview SDK ticket
def is_issue_summary_matches_preview_sdk(issue_summary):
    return issue_summary and 'Preview '.lower() in issue_summary.lower()


# Returns None if manual change is not requested, else returns with requested manual change comment
def check_comments_for_manual_change(issue):
    for comment in reversed(issue.fields.comment.comments):
        # This check is true if a manual change has already been executed
        if "MANUAL CHANGE REQUEST" in str(comment.body).upper() and BOT_ID in comment.author.name:
            return None
        # Checks for a NEW manual change request
        elif BOT_USERNAME in comment.body and "MANUAL CHANGES REQUESTED" in str(comment.body).upper()\
                and BOT_ID not in comment.author.name:
            print("Manual change request for {} found!".format(issue.key))
            return comment.body
    return None


def is_new_design_review(issue):
    if len(issue.fields.comment.comments) <= 1:
        print("Found new Design Review ticket: {}".format(issue.key))
        return True
    else:
        new_ticket = True
        for comment in issue.fields.comment.comments:
            if "Generated code changes:" not in comment.body:
                new_ticket = False
        return new_ticket


def no_changes_required(issue, not_by=[]):
    for comment in reversed(issue.fields.comment.comments):
        if 'fail' in comment.body or comment.author.name in not_by: 
            break
        if comment.author.name not in not_by + [BOT_ID] and re.match(NO_MANUAL_CHANGES_REQUIRED, comment.body.lower()):
            return True
    return False


# Checks if issue has been updated/modified in the last 7 days
def not_updated_in_seven_days(issue):
    comments = issue.fields.comment.comments
    if len(comments) > 0:
        last_comment = comments[-1]
        comment_datetime = convert_date_to_datetime(last_comment.created[:10])
        today = datetime.datetime.now()
        delta = today - comment_datetime
        if delta.days > 7:
            print("{} not updated in more than 7 days.".format(issue.key))
            return True
    return False


# Returns the associated UDX ticket from the Design Review ticket
def get_udx_ticket_from_design_review(design_review):
    if hasattr(design_review.fields, 'issuelinks'):
        for link in design_review.fields.issuelinks:
            if hasattr(link, 'inwardIssue'):
                issue = JIRA_CLIENT().issue(link.inwardIssue.key, fields='description, summary, status, issuelinks')
                if "UDX" in issue.key:
                    print('Found UDX ticket: {}'.format(issue.key))
                    return issue
    return None


# Returns the associated Public SDK/CLI ticket from the UDX ticket
def get_public_ticket_from_udx(udx_ticket):
    if hasattr(udx_ticket.fields, 'issuelinks'):
        for link in udx_ticket.fields.issuelinks:
            if hasattr(link, 'outwardIssue'):
                issue = JIRA_CLIENT().issue(link.outwardIssue.key)
                if is_issue_summary_matches_public_sdk(issue.fields.summary):
                    print('Found Public SDK issue: {}'.format(issue.key))
                    if issue.fields.status.name not in config.DEXREQ_TERMINAL_STATES:
                        return issue
    return None


# Checks if the Public SDK/CLI ticket is GA within a month
def is_ga_date_within_month(public_ticket):
    if hasattr(public_ticket.fields, config.CUSTOM_FIELD_ID_SDK_CLI_GA_DATE):
        ga_date = getattr(public_ticket.fields, config.CUSTOM_FIELD_ID_SDK_CLI_GA_DATE)
        if ga_date is None:
            return None
        ga_datetime = convert_date_to_datetime(ga_date)
        today = datetime.datetime.now()
        delta = ga_datetime - today
        if delta.days <= 28:
            return delta.days
    return None


def convert_date_to_datetime(ga_date):
    return datetime.datetime.strptime(ga_date, '%Y-%m-%d')


def is_issue_summary_matches_public_sdk(issue_summary):
    return issue_summary and 'Public '.lower() in issue_summary.lower()


# Returns the generated CLI branch from a preview SDK/CLI ticket
def get_generated_branch(issue):
    for comment in reversed(issue.fields.comment.comments):
        if "Generated code changes:" in comment.body:
            generated_branch = GENERATED_BRANCH_REGEX.search(comment.body)
            if generated_branch:
                generated_branch = generated_branch.group(0)
            print("Found generated branch: {}".format(generated_branch))
            return generated_branch
    return None


def execute_manual_change(request):
    request_upper = request.upper()
    # Cleans the request for formatting
    request_upper = re.sub(VALID_CHARS, '', request_upper)

    for operation_type in SUPPORTED_MANUAL_CHANGES:
        if operation_type in request_upper:
            try:
                body = re.compile(MANUAL_CHANGE_MATCH.format(operation_type)).search(request_upper)
                if body:
                    body = body.group(3)
                    body = body.lower()
                    MANUAL_CHANGE_FUNCTIONS[operation_type](body)
            except CliException as exception:
                print("enqueuing  raised exceptions for {} and the exception is {}".format(operation_type, exception.message))
                raised_exceptions[exception.line] = exception

    if raised_exceptions:
        print("found exceptions while executing manual changes")
        raise Exception("found exceptions while executing manual changes")

    print("Successfully executed manual change!")
    manual_change.run_make_docs(find_service(request))
    manual_change.run_make_docs(find_service(request))
    print("Successfully ran make docs")


def rename_root_group(request):
    request_upper = request.upper()
    # Cleans the request for formatting
    request_upper = re.sub(VALID_CHARS, '', request_upper)
    print(request_upper)

    if RENAME_ROOT_GROUP in request_upper:
        body = re.compile(MANUAL_CHANGE_MATCH.format(RENAME_ROOT_GROUP)).search(request_upper)
        if body:
            body = body.group(3)
            body = body.lower()
            MANUAL_CHANGE_FUNCTIONS[RENAME_ROOT_GROUP](body)
            print("Successfully executed renaming root group!")
    else:
        return


def find_service(request):
    for line in reversed(request.splitlines()):
        if 'oci' in line:
            values = line.split()
            for i, val in enumerate(values):
                if 'oci' in val:
                    return values[i + 1].strip()


def execute_rename_root_group(body):
    for line in body.splitlines():
        if not line.strip():
            continue
        request = line.split(MANUAL_CHANGE_DELIMTER)
        old_root_group = sanitize_command(request[0])
        new_root_group = sanitize_command(request[1])
        try:
            manual_change.rename_root_group(old_root_group[-1], new_root_group[-1])
        except Exception as e:
            raise CliException(RENAME_ROOT_GROUP_EXCEPTION.format(line, e), line)


def execute_rename_command(body):
    commands = []
    new_names = []
    for line in body.splitlines():
        if not line.strip():
            continue
        request = line.split(MANUAL_CHANGE_DELIMTER)
        old_command = sanitize_command(request[0])
        commands.append(old_command)
        new_command = sanitize_command(request[1])
        new_names.append(new_command[-1])
        try:
            manual_change.rename_commands(commands, new_names)
        except Exception as e:
            raise CliException(RENAME_COMMAND_EXCEPTION.format(line=line, e=e), line)
        commands = []
        new_names = []


def execute_remove_command(body):
    commands = []
    for line in body.splitlines():
        if not line.strip():
            continue
        request = line.split(MANUAL_CHANGE_DELIMTER)
        old_command = sanitize_command(request[0])
        commands.append(old_command)
        manual_change.remove_commands(commands)
        commands = []


def execute_move_group(body):
    for line in body.splitlines():
        if not line.strip():
            continue
        request = line.split(MANUAL_CHANGE_DELIMTER)
        old_group = sanitize_command(request[0])
        new_group = sanitize_command(request[1])
        try:
            manual_change.move_group(old_group, new_group)
        except Exception as e:
            raise CliException(MOVE_GROUP_OR_COMMANDS_EXCEPTION.format(line=line, e=e), line)


def execute_move_command(body):
    for line in body.splitlines():
        if not line.strip():
            continue
        request = line.split(MANUAL_CHANGE_DELIMTER)
        command = sanitize_command(request[0])
        new_group = sanitize_command(request[1])
        try:
            manual_change.move_command(command, new_group)
        except Exception as e:
            raise CliException(MOVE_GROUP_OR_COMMANDS_EXCEPTION.format(line=line, e=e), line)


def execute_rename_parameter(body):
    for line in body.splitlines():
        if not line.strip():
            continue
        if 'oci' in line and MANUAL_CHANGE_DELIMTER in line:
            rename_parameter([line])


def rename_parameter(body):
    command = []
    old_params = []
    new_params = []
    options = None
    for line in body:
        if not line.strip():
            continue
        if ":" in line:
            if options is None:
                options = {}
            option = line.split(':')
            options.update({option[0].strip(): option[1].strip('][').split(', ')})
            # options.update({option[0].strip(): ast.literal_eval(option[1].strip())})
        else:
            request = line.split(MANUAL_CHANGE_DELIMTER)
            old_command = sanitize_command(request[0])
            command = old_command[:find_parameter_index(old_command)]
            old_params = old_command[find_parameter_index(old_command):]
            old_params = [x.replace('--', '') for x in old_params]
            new_command = sanitize_command(request[1])
            new_params = new_command[find_parameter_index(new_command):]
            new_params = [x.replace('--', '') for x in new_params]
    print("Renaming Command")
    print(command)
    print(old_params)
    print(new_params)
    try:
        manual_change.rename_parameters(command, old_params, new_params, options)
    except Exception as e:
        raise CliException(RENAME_COMMAND_EXCEPTION.format(line=line, e=e), line)


def execute_remove_parameter(body):
    for line in body.splitlines():
        if not line.strip():
            continue
        if 'oci' in line:
            remove_parameter([line])


def remove_parameter(body):
    command = []
    params = []
    for line in body:
        if not line.strip():
            continue
        request = line.split(MANUAL_CHANGE_DELIMTER)
        old_command = sanitize_command(request[0])
        command = old_command[:find_parameter_index(old_command)]
        params = old_command[find_parameter_index(old_command):]
        params = [x.replace('--', '') for x in params]
    manual_change.remove_parameters(command, params)


# TODO
def execute_flatten_parameter(body):
    flatten_query = []
    query = False
    for line in body.splitlines():
        if not line.strip():
            continue
        if 'oci' in line and query:
            query = False
            flatten_parameter(flatten_query)
            flatten_query = [line]
        else:
            query = True
            flatten_query.append(line)
    flatten_parameter(flatten_query)


# TODO
def flatten_parameter(body):
    command = []
    params = []
    params_flattened = []
    params_options = {}
    for line in body:
        if not line.strip():
            continue
        if 'oci' in line:
            command = sanitize_command(line)
        else:
            option = line.split(':')  # noqa: F841
    manual_change.flatten_parameters(command, params, params_flattened, params_options)


# Removes leading/trailing whitespaces from command and converts into comma separated list
# oci audit list -> ['audit', 'list']
def sanitize_command(request):
    request = request.split()
    if "oci" == request[0]:
        request = request[1:]
    request = [x.strip() for x in request]
    return request


def find_parameter_index(command):
    for i, group in enumerate(command):
        if '--' in group:
            return i


MANUAL_CHANGE_FUNCTIONS = {RENAME_ROOT_GROUP: execute_rename_root_group,
                           RENAME_COMMAND: execute_rename_command,
                           MOVE_COMMAND: execute_move_command,
                           REMOVE_COMMAND: execute_remove_command,
                           MOVE_GROUP: execute_move_group,
                           RENAME_PARAMETER: execute_rename_parameter,
                           REMOVE_PARAMETER: execute_remove_parameter}


def handle_exceptions(issue_key):
    comment = "CLI bot failed to create branch based on your last comment. Please fix these errors  \n"
    for key, exception in raised_exceptions.items():
        comment += key + "\n " + exception.message
        comment += "\n**********\n"
    print(comment)
    add_jira_comment(issue_key, comment)
