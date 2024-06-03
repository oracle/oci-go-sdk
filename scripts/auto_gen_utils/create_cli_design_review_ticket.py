import util
import config
import argparse
import sys
import re

from jira import JIRAError
from shared.buildsvc_tc_compatibility import build_log_link

MAX_SUB_COMMAND_LEN = 25
MAX_PARAM_LEN = 25
IS_ACTION_ITEM_OF_LINK = 'is action item of'
DEPENDS_ON_LINK = 'depends on'
REQUIRED_PARAM_TO_OPTIONAL_WARNING_TEMPLATE = 'Required param {} has been made optional.'
OPTIONAL_PARAM_TO_REQUIRED_WARNING_TEMPLATE = 'Optional param {} has been made required.'
PARAM_TOO_LONG_WARNING_TEMPLATE = 'Param {} has more than 25 characters. Please consider to shorten it if possible.'
SUB_COMMAND_TOO_LONG_WARNING_TEMPLATE = 'Sub-Command {} has more than 25 characters. Please consider to shorten it if possible.'
REQUIRED_PARAM_DELETED_OR_RENAME_WARNING = 'Required param {} is deleted/renamed.'
OPTIONAL_PARAM_DELETED_OR_RENAME_WARNING = 'Optional param {} is deleted/renamed.'
DESIGN_REVIEW_ISSUE_SUCCESS_MESSAGE_TEMPLATE = 'The following CLI Design Review ticket has been opened and needs your attention -- https://jira.oci.oraclecorp.com/browse/{}'
CLI_DESIGN_FAILURE_TEMPLATE = """The job failed to create/update CLI Design review ticket. {exception}.

The full build log can be found {build_log_link}.

For build log and artifact access see the [self-service FAQ|https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=132774928#SDK/CLISelfServiceFrequentlyAskedQuestions-HowdoIgetaccesstologsandartifacts?].
"""

MODIFIED_COMMANDS_SUMMARY_TEMPLATE = """

{{code:bash|title=Modified Existing Commands|borderStyle=solid}}{modified_commands}{{code}}
"""

REMOVED_COMMANDS_SUMMARY_TEMPLATE = """

{{code:bash|title=Removed Commands|borderStyle=solid}}{removed_commands}{{code}}
"""

ADDED_COMMANDS_SUMMARY_TEMPLATE = """

{{code:bash|title=Newly Added Commands|borderStyle=solid}}{added_commands}{{code}}
"""

GENERATED_PR_TEMPLATE = """
Generated code changes: {pr_link}
"""

CONFLICTING_PARAMETER_TEMPLATE = """
*Parameter conflict detected:* Your generated commands conflict with internal CLI parameters. Please make manual changes to the parameters that end with *-parameterconflict*.

"""

CLI_DESIGN_SUMMARY_TEMPLATE = """
Please review the generated commands posted in the comments. Work with product management to suggest or accept changes to the generated command structure. After reviewing the generated commands:
* Before you review this ticket, make sure preview DEXREQ ticket is marked done. We run weekly bulk-preview job on Wednesday which will mark preview DEXREQ ticket to done. Don't change the preview DEXREQ manually to done. 
* If you feel that generated commands changes are good to go - please leave a comment saying *No changes required* on this ticket and assign the ticket to the SDK/CLI on-call engineer.
https://ocean.ocs.oraclecloud.com/org/8723F92A522EB107E0530424000A6404/team/AB9CC5B3A525512FE0530202FD0ACA13/calendars
* If you want to modify the generated commands, follow these steps:
    - Follow https://confluence.oci.oraclecorp.com/display/DEX/CLI+Self+Service+Guide and leave a comment using the appropriate template
    - Our CLI bot will read "[~gear-dexreq-automation] Manual Changes Requested" in comment and process the comment and create a branch automatically
    - Only last comment will be used to create a branch
    - Follow the comment on the ticket after branch is created

h4. Command Review Checklist:
# Are commands too long/verbose? Can commands be shortened and yet convey the correct meaning?
# Do any commands need to be renamed/hidden due to polymorphic input types?
# Do you need to rename a root command (load-balancer -> lb)?
# Are any parameters better specified as a file instead of text input?
# Are any parameters complex and can be expanded into parameters for individual elements?
# Do you need to restructure generated commands to match the common structure (oci + service + resource + operation)?

More info on CLI design review process: https://confluence.oci.oraclecorp.com/pages/viewpage.action?spaceKey=DEX&title=CLI+Command+Structure+Design+Review
SDK/CLI Support engineer information can be found here: https://ocean.ocs.oraclecloud.com/org/8723F92A522EB107E0530424000A6404/team/AB9CC5B3A525512FE0530202FD0ACA13/calendars or reach out to *#oci_public_cli* slack group.

Making CLI Manual changes using our CLI Self Service BOT: https://confluence.oci.oraclecorp.com/display/DEX/CLI+Self+Service+Guide
Making CLI Manual changes: https://confluence.oci.oraclecorp.com/display/DEX/CLI+Manual+Code+Changes
Common manual change recipes: https://confluence.oci.oraclecorp.com/display/DEX/CLI+recipes+for+overriding+generated+code
Recommended CLI installation for Mac: https://confluence.oci.oraclecorp.com/display/DEX/Installing+OCI-CLI+using+Python3

Once the manual change process is complete, please ensure end to end testing of all generated or modified CLI commands.
After your changes are merged, install the generated [preview build|https://confluence.oci.oraclecorp.com/pages/viewpage.action?spaceKey=DEX&title=CLI+FAQs#CLIFAQs-WherecanIfindthelatestversionofthepreviewCLI?].
For setting up the development environment, [refer|https://confluence.oci.oraclecorp.com/display/DEX/OCI+CLI+Development+Setup#OCICLIDevelopmentSetup-1.Settingupthedevelopmentenvironment].
For running tests in your tenancy please follow the [steps mentioned|https://confluence.oci.oraclecorp.com/display/DEX/OCI+CLI+Development+Setup#OCICLIDevelopmentSetup-Runningtestsagainstanothertenancy(RecommendedforServiceTeam)]
Test all the commands and param which are added, modified or deleted after the design review process.

*NEW FEATURE FOR AUTOMATING CLI MANUAL CHANGE REQUESTS*
*======================================================*
Service teams can now request CLI manual changes by updating their Spec. Check this link for more details https://confluence.oci.oraclecorp.com/display/DEX/CLI+Spec+Customisation%3A+User+Guide. The bot will auto detect and suggest if your changes are applicable for this automation.
"""

CLI_NO_CHANGES_TO_THE_COMMAND_STRUCTURE = """
{issue} is not introducing any new/modifying existing commands. This can only happen in the following cases:

* Updating documentation for the existing operations.
* Changes to request models which are complex type inputs for a CLI Command.
* Changes to response models.

In case if you are expecting that your DEXREQ ticket should have a generated new commands or modify existing commands, Please reach out to *#oci_public_cli* slack group.
"""

TRY_CLI_SPEC_CUSTOMIZATION_SUGGESTION_TEMPLATE = """

*NEW FEATURE FOR AUTOMATING CLI MANUAL CHANGE REQUESTS*
*======================================================*
Below generated APIs have been detected as applicable for CLI change requests through spec customisation. Service team can now request CLI changes by updating their Spec and re-processing DEXREQ ticket. Check this link for more details https://confluence.oci.oraclecorp.com/display/DEX/CLI+Spec+Customisation%3A+User+Guide. 

"""

NEW_DETECTED_API_SUMMARY_TEMPLATE = """

{{code:bash|borderStyle=solid}}{added_commands}{{code}}
"""


# This method parses command_summary full_json_input.txt file to construct a map with key as full
# command string and value as another map where complex param is the key and its structure is the value.
# i.e content_map[oci iam tag-namespace create][freeform-tags] ={"tagKey1": "tagValue1","tagKey2": "tagValue2"}
def generate_cli_dictionary_for_json_input(file_contents):
    content_map = {}
    for match in re.finditer(r'oci([^\n]+)\s*([\w-]+)\s*:\s*(.*?)(?=\s*oci|$)', file_contents, re.S):
        if match.group(1) in content_map:
            content_map[match.group(1)][match.group(2)] = match.group(3)
        else:
            content_map[match.group(1)] = {match.group(2): match.group(3)}

    return content_map


# This method parses command_summary file to construct a map
# with key as full command string and CliCommandSummary object as value.
def generate_cli_dictionary(file_contents):
    contents = file_contents.split("\n")
    content_map = {}
    i = 0
    contents.remove("")
    while i < len(contents):
        if i % 4 == 0 and i + 3 < len(contents) and contents[i].startswith('oci'):
            content_map[contents[i]] = CliCommandSummary(
                contents[i],
                contents[i + 1],
                contents[i + 2],
                contents[i + 3])
            i += 4
    return content_map


# Detects if a required param has been modified to optional or vice versa and removes the
# common param in both the commands
def detect_change_in_required_or_optional_params(params_one, params_two, warning_str, modified_params):
    warnings = []
    for param in params_one[:]:
        if param in params_two:
            warnings.append(warning_str.format(param))
            params_one.remove(param)
            params_two.remove(param)
            modified_params.append(param)

    return params_one, params_two, warnings


def subtract_lists(list1, list2):
    return [element for element in list1 if element not in list2]


def get_modified_command(modified_cmd, original_cmd):
    modified_params = []
    # Get required and optional params
    m_required_params = modified_cmd.get_required_params()
    m_optional_params = modified_cmd.get_optional_params()
    o_required_params = original_cmd.get_required_params()
    o_optional_params = original_cmd.get_optional_params()

    # Filter unchanged params
    m_unique_required_params = subtract_lists(m_required_params, o_required_params)
    o_unique_required_params = subtract_lists(o_required_params, m_required_params)
    m_unique_optional_params = subtract_lists(m_optional_params, o_optional_params)
    o_unique_optional_params = subtract_lists(o_optional_params, m_optional_params)

    # Detect if any of the required params are dropped to optional params
    o_unique_required_params, m_unique_optional_params, required_change_warnings = detect_change_in_required_or_optional_params(
        o_unique_required_params, m_unique_optional_params, REQUIRED_PARAM_TO_OPTIONAL_WARNING_TEMPLATE, modified_params)

    # Detect if any of optional params are moved to required params
    m_unique_required_params, o_unique_optional_params, optional_change_warnings = detect_change_in_required_or_optional_params(
        m_unique_required_params, o_unique_optional_params, OPTIONAL_PARAM_TO_REQUIRED_WARNING_TEMPLATE, modified_params)

    command_summary_warnings = required_change_warnings + optional_change_warnings
    # Add newly added required params
    modified_params.extend(m_unique_required_params)
    # Add newly added optional params
    modified_params.extend(m_unique_optional_params)

    # Renamed/Deleted required params.
    for param in o_unique_required_params:
        command_summary_warnings.append(REQUIRED_PARAM_DELETED_OR_RENAME_WARNING.format(param))
    # Renamed/Deleted optional params.
    for param in o_unique_optional_params:
        command_summary_warnings.append(OPTIONAL_PARAM_DELETED_OR_RENAME_WARNING.format(param))
    modified_cli_command_summary = ModifiedCliCommandSummary(modified_cmd)
    modified_cli_command_summary.modified_params = modified_params
    modified_cli_command_summary.warnings = command_summary_warnings

    return modified_cli_command_summary


# Compares the local command summary file with the remote file to find
# differences in the command structure and to check if any new commands have
# added to the local command summary file.
def get_generated_commands(diff):
    modified_cmds = []
    added_cmds = []
    complex_modified_cmds = []
    # filter non command-summary changes

    if (diff.a_path and 'docs/command-summary' in diff.a_path and 'full_json_input.txt' not in diff.a_path) and \
            (diff.b_path and 'docs/command-summary' in diff.b_path and 'full_json_input.txt' not in diff.b_path):
        original_file = diff.a_blob.data_stream.read().decode('utf-8')
        original_dictionary = generate_cli_dictionary(original_file)
        modified_file = diff.b_blob.data_stream.read().decode('utf-8')
        modified_dictionary = generate_cli_dictionary(modified_file)

        for key in modified_dictionary:
            if original_dictionary.__contains__(key):
                if modified_dictionary[key] != original_dictionary[key]:
                    # change in command structure detected
                    modified_cmds.append(get_modified_command(modified_dictionary[key], original_dictionary[key]))
            else:
                added_cmds.append(modified_dictionary[key])
    if (diff.a_path and 'docs/command-summary' in diff.a_path and 'full_json_input.txt' in diff.a_path) and (
            diff.b_path and 'docs/command-summary' in diff.b_path and 'full_json_input.txt' in diff.b_path):
        original_file = diff.a_blob.data_stream.read().decode('utf-8')
        original_dictionary = generate_cli_dictionary_for_json_input(original_file)
        modified_file = diff.b_blob.data_stream.read().decode('utf-8')
        modified_dictionary = generate_cli_dictionary_for_json_input(modified_file)
        for command in modified_dictionary:
            if original_dictionary.__contains__(command):
                for complex_param in modified_dictionary[command]:
                    if complex_param in original_dictionary[command]:
                        if modified_dictionary[command][complex_param] != original_dictionary[command][complex_param]:
                            complex_modified_cmds.append((command, complex_param))
    return modified_cmds, added_cmds, complex_modified_cmds


# When a new path is defined in the oci-cli, This returns all the commands defined in the new path.
def get_newly_added_commands(diff):
    original_dictionary = {}
    # filter non command-summary changes
    if diff.b_path and 'docs/command-summary' in diff.b_path and 'full_json_input.txt' not in diff.b_path:
        added_file = diff.b_blob.data_stream.read().decode('utf-8')
        original_dictionary = generate_cli_dictionary(added_file)
    return original_dictionary


def get_removed_commands(diff):
    removed_cmds = []
    # filter non command-summary changes
    if (diff.a_path and 'docs/command-summary' in diff.a_path and 'full_json_input.txt' not in diff.a_path) and \
            (diff.b_path and 'docs/command-summary' in diff.b_path and 'full_json_input.txt' not in diff.b_path):
        original_file = diff.a_blob.data_stream.read().decode('utf-8')
        original_dictionary = generate_cli_dictionary(original_file)
        modified_file = diff.b_blob.data_stream.read().decode('utf-8')
        modified_dictionary = generate_cli_dictionary(modified_file)
        for key in original_dictionary:
            if not modified_dictionary.__contains__(key):
                removed_cmds.append(original_dictionary[key])
    return removed_cmds


def is_issue_summary_matches_cli_design(issue_summary):
    return issue_summary and 'CLI Design Review'.lower() in issue_summary.lower()


def create_design_ticket(dexreq_issue):
    print('Creating new CLI Design Review Issue')
    sprint_id = util.find_dex_tools_active_sprint_id()
    udx_issue_field = getattr(dexreq_issue.fields, config.CUSTOM_FIELD_ID_UDX_TICKET)
    udx_issue_keys = util.get_udx_issue_keys(udx_issue_field)
    related_issues = ', '.join([dexreq_issue.key] + udx_issue_keys) if udx_issue_keys else dexreq_issue.key
    fields = {
        'project': 'DEX',
        'summary': 'CLI Design Review for {}'.format(related_issues),
        'description': CLI_DESIGN_SUMMARY_TEMPLATE,
        'issuetype': {'name': 'Design Reviews'},
        'components': [{'name': 'CLI'}],
        'labels': ['CLIDesignReview'],
        'assignee': {'name': dexreq_issue.fields.creator.name},
        config.CUSTOM_FIELD_ID_SPRINT: sprint_id
    }

    design_ticket = util.JIRA_CLIENT().create_issue(fields)

    util.JIRA_CLIENT().create_issue_link(IS_ACTION_ITEM_OF_LINK, design_ticket, dexreq_issue)
    if udx_issue_keys:
        for u in udx_issue_keys:
            udx_issue = util.JIRA_CLIENT().issue(u)
            util.JIRA_CLIENT().create_issue_link(DEPENDS_ON_LINK, udx_issue, design_ticket)

    print('Issue {} has been created'.format(design_ticket.key))
    return design_ticket


# returns CliDesignReview ticket associated with DexReq ticket.
# Creates a new CliDesignReview ticket if DexReq doesn't have any review ticket associated to it.
def get_cli_design_ticket(dexreq_issue, create_issue_if_absent=True):
    if hasattr(dexreq_issue.fields, 'issuelinks'):
        for link in dexreq_issue.fields.issuelinks:
            if hasattr(link, 'outwardIssue'):
                issue = util.JIRA_CLIENT().issue(link.outwardIssue.key,
                                                 fields='description, summary, status, issuetype, labels')
                if is_issue_summary_matches_cli_design(issue.fields.summary):
                    print('Found CLI Design review issue: {}'.format(issue.key))
                    return issue

    if create_issue_if_absent:
        return create_design_ticket(dexreq_issue)

    return None


def is_design_ticket_in_non_terminal_state(issue):
    return issue and issue.fields.status and issue.fields.status.name not in config.CLI_DESIGN_REVIEW_TERMINAL_STATES


# Returns pending design review tickets associated with an UDX ticket.
def get_cli_design_review_issues_for_udx(udx_issue_keys):
    design_tickets = []
    for udx_issue_key in udx_issue_keys:
        try:
            udx_issue = util.JIRA_CLIENT().issue(udx_issue_key)
            if hasattr(udx_issue.fields, 'issuelinks'):
                for link in udx_issue.fields.issuelinks:
                    if hasattr(link, 'outwardIssue'):
                        issue = util.JIRA_CLIENT().issue(link.outwardIssue.key, fields='description, summary, status, labels')
                        if is_issue_summary_matches_cli_design(issue.fields.summary):
                            print('Found CLI Design review issue: {}'.format(issue.key))
                            design_tickets.append(issue)
        except JIRAError as e:
            print("JIRAError while getting CLI design review issues for {}: {}".format(udx_issue_keys, str(e)))

    return design_tickets


def get_cli_command_str_with_warnings(cli_cmd, is_modified_command=False, modified_complex_params=[]):
    params = cli_cmd.get_all_params_str() if not is_modified_command else cli_cmd.get_modified_params_str()
    complete_command = cli_cmd.cmd_str + ' ' + params
    return_string = complete_command
    for complex_cmd, complex_param_name in modified_complex_params:
        if complex_cmd in complete_command:
            return_string += "\n## Complex param --" + complex_param_name + " has been updated."
    if cli_cmd.get_warnings_str():
        return_string += '\n' + cli_cmd.get_warnings_str()

    return return_string


def check_for_long_params(cli_cmds, is_modified_cmds=False):
    for cli_cmd in cli_cmds:
        params = cli_cmd.get_all_params() if not is_modified_cmds else cli_cmd.modified_params
        for param in params:
            if len(param) > MAX_PARAM_LEN:
                cli_cmd.add_warning(PARAM_TOO_LONG_WARNING_TEMPLATE.format(param))


def check_for_long_sub_commands(cli_cmds):
    for cli_cmd in cli_cmds:
        sub_commands = cli_cmd.cmd_str.split(' ')
        for sub_command in sub_commands:
            if len(sub_command) > MAX_SUB_COMMAND_LEN:
                cli_cmd.add_warning(SUB_COMMAND_TOO_LONG_WARNING_TEMPLATE.format(sub_command))


# This functions returns commands eligible for spec-customization in comment format.
# The commands which are new addition in the service are eligible for spec customisation.
def check_new_added_commands(newly_added_commands_string):
    if newly_added_commands_string.strip():
        NEW_DETECTED_API_COMMAND_FORMATTED_TEXT = NEW_DETECTED_API_SUMMARY_TEMPLATE.format(added_commands=newly_added_commands_string)
        return TRY_CLI_SPEC_CUSTOMIZATION_SUGGESTION_TEMPLATE + NEW_DETECTED_API_COMMAND_FORMATTED_TEXT
    return ""


# This uses tests/output/command-summary/* files to determine change in command structure.
# We compare local command-summary/* files with the remote preview branch.
def get_command_summary_changes():
    cli_repo = config.CLI_REPO
    head_commit = cli_repo.head.commit
    head_commit_origin = cli_repo.commit('origin/preview')
    diff = head_commit_origin.diff(head_commit, create_patch=True)

    modified_commands = []
    newly_added_commands = []
    modified_complex_params = []
    removed_commands = []

    for mdiff in diff.iter_change_type('M'):
        mc, nc, cmc = get_generated_commands(mdiff)
        modified_commands.extend(mc)
        newly_added_commands.extend(nc)
        modified_complex_params.extend(cmc)

    for adiff in diff.iter_change_type('A'):
        newly_added_commands.extend(list(get_newly_added_commands(adiff).values()))

    for ddiff in diff.iter_change_type('M'):
        removed_commands.extend(get_removed_commands(ddiff))

    check_for_long_params(modified_commands, True)
    check_for_long_params(newly_added_commands)
    check_for_long_sub_commands(newly_added_commands)

    modified_commands_string = '\n\n'.join(
        map(lambda cli_cmd: get_cli_command_str_with_warnings(cli_cmd, True, modified_complex_params), modified_commands)
    )

    newly_added_commands_string = '\n\n'.join(
        map(lambda cli_cmd: get_cli_command_str_with_warnings(cli_cmd), newly_added_commands)
    )

    removed_commands_string = '\n\n'.join(
        map(lambda cli_cmd: get_cli_command_str_with_warnings(cli_cmd), removed_commands)
    )

    modified_commands_comment = MODIFIED_COMMANDS_SUMMARY_TEMPLATE.format(modified_commands=modified_commands_string) if modified_commands_string else ""
    added_commands_comment = ADDED_COMMANDS_SUMMARY_TEMPLATE.format(added_commands=newly_added_commands_string) if newly_added_commands_string else ""
    removed_commands_comment = REMOVED_COMMANDS_SUMMARY_TEMPLATE.format(removed_commands=removed_commands_string) if removed_commands_string else ""
    conflicting_parameters_comment = check_for_conflicting_params(modified_commands_comment + added_commands_comment)

    print('Modified Commands: {}'.format(modified_commands_string))
    print('Newly Added Commands: {}'.format(newly_added_commands_string))
    print('Removed Commands: {}'.format(removed_commands_string))
    try_cli_spec_customization_comment = check_new_added_commands(newly_added_commands_string)
    return modified_commands_comment + added_commands_comment + removed_commands_comment + conflicting_parameters_comment + try_cli_spec_customization_comment


def check_for_conflicting_params(command_summary_changes):
    if "-parameterconflict" in command_summary_changes:
        return CONFLICTING_PARAMETER_TEMPLATE
    return ""


class CliCommandSummary:
    default_cli_params = ['--from-json', '--help', '--limit', '--page', '--page-size', '--sort-by',
                          '--sort-order', '--if-match', '--wait-for-state', '--wait-interval-seconds',
                          '--max-wait-seconds', '--lifecycle-state']

    def __init__(self, cmd_str, count, required_params, optional_params, warnings=None):
        self.cmd_str = cmd_str
        self.count = count
        self.required_params = required_params
        self.optional_params = optional_params
        if warnings is None:
            warnings = []
        self.warnings = warnings

    def get_required_params(self):
        return self.required_params.replace('Required Parameters:', '').strip().split(', ')

    def get_optional_params(self):
        optional_params = self.optional_params.replace('Optional Parameters:', '').strip().split(', ')
        # exclude default cli params in the optional params - we don't want to display them!
        optional_params = [param for param in optional_params if param not in CliCommandSummary.default_cli_params]
        return optional_params

    def get_all_params(self):
        required_params = self.get_required_params()
        optional_params = self.get_optional_params()
        return required_params + optional_params

    def get_all_params_str(self):
        return ', '.join(filter(None, self.get_all_params()))

    def add_warning(self, warning):
        self.warnings.append(warning)

    def get_warnings_str(self):
        return '\n'.join('## ' + warning for warning in self.warnings)

    def __eq__(self, other):
        if not isinstance(other, CliCommandSummary):
            return NotImplemented

        return self.cmd_str == other.cmd_str and \
            self.count == other.count and \
            self.required_params == other.required_params and \
            self.optional_params == other.optional_params


class ModifiedCliCommandSummary(CliCommandSummary):
    def __init__(self, cli_command_summary, modified_params=[]):
        self.modified_params = modified_params
        CliCommandSummary.__init__(self, cli_command_summary.cmd_str, cli_command_summary.count, cli_command_summary.required_params, cli_command_summary.optional_params, cli_command_summary.warnings)

    def get_modified_params_str(self):
        return ', '.join(filter(None, self.modified_params))


def get_pull_request_link_for_generated_branch():
    try:
        with open('preview-pr.txt', 'r') as filehandle:
            pr_link = filehandle.read()
            print('PR link found: {}'.format(pr_link))
    except Exception as e:
        print(str(e))
        return ""
    return pr_link


if __name__ == "__main__":
    # Last step of CLI Preview build. This script will create a CLI Design review issue and
    # associates it with DEXREQ and UDX issues. CLI Design review issue will be automatically assigned to
    # the DEXREQ reporter.
    parser = argparse.ArgumentParser(description='CLI Design review issue creation!')
    parser.add_argument('--build-id',
                        required=True,
                        help='The team city build id for the build that is running this script.  This is used to update the relevant JIRA tickets with links to the team city build')
    parser.add_argument('--tool',
                        default=config.CLI_NAME,
                        help='The tool for which to generate the preview. Accepted values: {}'.format(', '.join(config.TOOL_NAMES)))
    parser.add_argument('--build-type',
                        default=config.BUILD_TYPE_INDIVIDUAL_PREVIEW,
                        help='The build type to use, can be one of the following: {}'.format(', '.join(config.VALID_BUILD_TYPES)))
    parser.add_argument('--dry-run',
                        default=False,
                        action='store_true',
                        help='Perform a dry-run')

    args = parser.parse_args()
    tool_name = args.tool
    build_id = args.build_id
    build_type = args.build_type
    config.IS_DRY_RUN = args.dry_run

    if build_type != config.BUILD_TYPE_INDIVIDUAL_PREVIEW or tool_name != config.CLI_NAME:
        print('CLI Design review issue generation not required for build type: {}'.format(build_type))
        sys.exit(0)

    generation_pass, build_pass = util.were_steps_successful(tool_name)
    if not (generation_pass and build_pass):
        print('Generation or Build did not pass, not proceeding.')
        sys.exit(0)

    last_commit_message = util.get_last_commit_message(tool_name)
    issue_keys = util.parse_issue_keys_from_commit_message(last_commit_message)
    if len(issue_keys) != 1:
        print('More than one DEXReq issues found {}. Not generating CLI Design review ticket', issue_keys)
        sys.exit(0)
    dexreq_issue = issue_keys[0]
    try:
        command_summary_changes_str = get_command_summary_changes()
        dex_issue = get_cli_design_ticket(util.get_dexreq_issue(dexreq_issue))
        pr_link = GENERATED_PR_TEMPLATE.format(pr_link=get_pull_request_link_for_generated_branch())
        if command_summary_changes_str:
            comment_str = command_summary_changes_str + pr_link
            # if ticket is in 'Done'/'Closed' state, change state to 'Needs Triage Status'
            if dex_issue.fields.status.name in [config.STATUS_DONE, config.STATUS_CLOSED]:
                util.transition_issue_overall_status(util.JIRA_CLIENT(), dex_issue, config.STATUS_NEEDS_TRIAGE_STATUS)
        else:
            comment_str = CLI_NO_CHANGES_TO_THE_COMMAND_STRUCTURE.format(issue=dexreq_issue) + pr_link

        print('Updating CLI Design Review issue {} with {} '.format(dex_issue.key, comment_str))
        util.add_jira_comment(dex_issue, comment_str)
        util.add_jira_comment(dexreq_issue, DESIGN_REVIEW_ISSUE_SUCCESS_MESSAGE_TEMPLATE.format(dex_issue.key))

    except Exception as e:
        issue = util.get_dexreq_issue(dexreq_issue)
        util.add_jira_comment(
            issue.key,
            CLI_DESIGN_FAILURE_TEMPLATE.format(
                exception=str(e),
                build_log_link=build_log_link(build_id)
            ),
            comment_type=config.COMMENT_TYPE_ERROR
        )
