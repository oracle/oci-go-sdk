from collections import defaultdict
from datetime import date, timedelta
from slack_sdk.webhook import WebhookClient

import argparse
import re
import sys
import os
import yaml

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../../..'))
sys.path.append(os.path.join(dir_path, '../../'))

from set_changed_service_env_variable import get_service_name_from_issue # noqa: ignore=F402
from shared import bitbucket_utils # noqa: ignore=F402

import config # noqa: ignore=F402
import constants # noqa: ignore=F402
import util # noqa: ignore=F402


def get_tool_name(branch):
    words = re.findall('(\w*)', branch)  # noqa: W605
    tool_name = set(words) & set(config.TOOL_NAMES)
    return tool_name.pop() if len(tool_name) > 0 else config.CLI_NAME


def get_monday_date():
    today = date.today()
    monday = today - timedelta(days=today.weekday())
    return monday.strftime(constants.JIRA_DATETIME_FORMAT)


def get_release_date():
    release = os.environ.get(constants.ENV_RELEASE_DATE)
    today = date.today()
    
    try:
        release = date.fromisoformat(release)
        
        if release > today:
            return release
        
    except (TypeError, ValueError):
        pass
    
    return today + timedelta(days=8 - today.weekday())


def get_issues_from_newest_bulk_public_pr(tool_name, cut_off_date):
    pr = get_bulk_public_pr(tool_name, cut_off_date)
    return re.findall(constants.DEXREQ_TICKET_PATTERN, pr[constants.PR_DESCRIPTION])


def get_bulk_public_pr(tool_name, cut_off_date):
    pr = bitbucket_utils.get_newest_pullrequest_with_string_after(
        constants.PROJECT, config.REPO_NAMES_FOR_TOOL[tool_name][-1], constants.BULK_PUBLIC_PR_TITLE, cut_off_date
    )
    
    if pr:
        return pr
    
    sys.exit('No PRs found for {} matching the title \'{}\' after {}'.format(
        config.REPO_NAMES_FOR_TOOL[tool_name][-1],
        constants.BULK_PUBLIC_PR_TITLE,
        cut_off_date.split('T')[0]
    ))


def get_changelog_entries(issues):
    output = defaultdict(str)
    
    for issue in issues:
        jira_issue = get_jira_issue(issue)
        service = get_service_name(jira_issue)
        sdk_entry, cli_entry = get_raw_changelog_entries(jira_issue)
        
        if not is_manual_change_ticket(jira_issue):
            cli_issues.append(issue)
        
        output[constants.SDK_SECTION] += format_entry(issue, service, sdk_entry)
        output[constants.CLI_SECTION] += format_entry(issue, service, cli_entry, True)

    return output


def get_jira_issue(issue):
    return jira.issue(issue, fields='{},{},{},{},comment'.format(
        config.CUSTOM_FIELD_ID_PREVIEW_ISSUE,
        config.CUSTOM_FIELD_ID_CHANGELOG,
        config.CUSTOM_FIELD_ID_CLI_CHANGELOG,
        constants.DEXREQ_TICKET_LABELS
    ))


def get_service_name(jira_issue):
    try:
        pr = bitbucket_utils.get_pullrequest_from_url(get_spec_pr_link(jira_issue)).json()
        return get_title_from_yaml_path(pr['id'], pr['fromRef']['latestCommit'])
    except Exception:
        return get_service_name_from_ticket(jira_issue)


def get_spec_pr_link(jira_issue):
    for comment in reversed(jira_issue.fields.comment.comments):
        match = re.search(constants.SPEC_PR_LINK, comment.body)
        if match:
            return match[1]


def get_title_from_yaml_path(id, commit_id):
    file = bitbucket_utils.get_file_content_from_commit_id_and_path(
        constants.PROJECT,
        constants.SPEC_REPO,
        get_yaml_path(id),
        commit_id
    ).content
    y = yaml.safe_load(file)
    return y['info']['title'].lower().rstrip(' api').title()


def get_yaml_path(id):
    pr_diff = bitbucket_utils.get_pullrequest_diff(constants.PROJECT, constants.SPEC_REPO, id).json()
    for diff in pr_diff['diffs']:
        destination = diff['destination']['toString']
        if destination.endswith('.yaml'):
            return destination


def get_service_name_from_ticket(jira_issue):
    preview_issue = getattr(jira_issue.fields, config.CUSTOM_FIELD_ID_PREVIEW_ISSUE)

    preview = jira.issue(preview_issue, fields=config.CUSTOM_FIELD_ID_SPEC_FRIENDLY_NAME)
    service = getattr(preview.fields, config.CUSTOM_FIELD_ID_SPEC_FRIENDLY_NAME)
    service_from_path = get_service_name_from_issue(jira_issue.id)

    if service == service_from_path:
        return service

    if service:
        return '{} from {}'.format(service, service_from_path)

    return service_from_path


def get_raw_changelog_entries(jira_issue):
    return getattr(jira_issue.fields, config.CUSTOM_FIELD_ID_CHANGELOG), getattr(jira_issue.fields, config.CUSTOM_FIELD_ID_CLI_CHANGELOG)


def is_manual_change_ticket(jira_issue):
    labels = getattr(jira_issue.fields, constants.DEXREQ_TICKET_LABELS)
    return any(label in config.CLI_PR_REQUIRED_LABELS for label in labels)


def format_entry(issue, service, entry, is_cli=False):
    return constants.CHANGELOG_ENTRY.format(
        link=constants.JIRA_DOMAIN + issue,
        dexreq_issue=issue,
        service=service,
        content=parse(entry, is_cli)
    )


def parse(entry, is_cli):
    cleaned_entry = previous_line = ''
    for line in entry.splitlines():
        line = line.strip()
        
        # exclude empty lines and lines that only say "changelog"
        if len(line) == 0 or re.match(constants.CHANGELOG_TITLE_PATTERN, line.lower()):
            continue
        
        match = re.search(constants.DESCRIPTION_PATTERN, line)
        line = line[match.end():]
        
        # exclude lines with no words, except if it's a heading underline
        if len(line) == 0:
            if is_cli and len(previous_line) > 0 and not previous_line.startswith('~'):
                line = '~' * len(previous_line)
            else:
                continue
        
        # try to format the line if it has more than 1 word
        if len(line.split()) > 1:
            command = re.fullmatch(constants.CLI_COMMAND_PATTERN, line.lower())
            if command:
                line = constants.CLI_COMMAND_FORMAT.format(command=command[1])
            else:
                line = constants.CLI_DESCRIPTION_FORMAT.format(text=line) if is_cli else \
                    constants.SDK_DESCRIPTION_FORMAT.format(text=line)
        
        cleaned_entry += line + '\n'
        previous_line = line
    
    return cleaned_entry


def compile_reports(entries, issues, version, release_date):
    manual_change_issues = list(filter(lambda issue: issue not in cli_issues, issues))
    
    sdk_report = constants.SDK_CHANGELOG_REPORT.format(
        mentions=' '.join(constants.SDK_MENTIONS),
        summary=fill_summary_template(issues, constants.SDK_SECTION),
        changelog_entries=entries.get(constants.SDK_SECTION)
    )
    
    cli_summary = fill_summary_template(cli_issues, constants.CLI_SECTION) + \
        fill_summary_template(manual_change_issues, constants.CLI_MANUAL_CHANGE_SECTION)
    header = constants.CLI_CHANGELOG_HEADER_TEMPLATE.format(version=version, release_date=release_date)
    
    cli_report = constants.CLI_CHANGELOG_REPORT.format(
        mentions=' '.join(constants.CLI_MENTIONS),
        summary=cli_summary,
        cli_header=header,
        cli=entries.get(constants.CLI_SECTION)
    )
    
    return [sdk_report, cli_report]
    

def fill_summary_template(issue_list, section):
    return constants.PR_SUMMARY_TEMPLATE.format(
        count=len(issue_list),
        plural='' if len(issue_list) == 1 else 's',
        section=section,
        issue_list=issue_list
    )


def post_reports_on_slack(reports):
    url = os.environ.get(constants.ENV_SLACK_WEBHOOK)
    webhook = WebhookClient(url)
    for report in reports:
        webhook.send(text=report)


if __name__ == '__main__':
    """
    This script will be used to gather CHANGELOG entries from DEXREQ tickets included in SDK and CLI's
    bulk public PRs.
    """
    parser = argparse.ArgumentParser(description='Changelog entry collection post bulk public PR generation.')
    parser.add_argument('--branch',
                        required=True,
                        help='The branch that triggered this job')
    parser.add_argument('--tool',
                        default=config.CLI_NAME,
                        help='The tool for which to generate the preview. Accepted values: {}'.format(', '.join(config.TOOL_NAMES)))
    
    args = parser.parse_args()
    
    global cli_issues
    tool_name = get_tool_name(args.branch)
    version = os.environ.get(constants.ENV_VERSION)
    release_date = get_release_date()
    cli_issues = []
    
    bitbucket_utils.setup_bitbucket(args)
    jira = util.JIRA_CLIENT()
    
    issues = get_issues_from_newest_bulk_public_pr(tool_name, get_monday_date())
    changelog_entries = get_changelog_entries(issues)
    reports = compile_reports(changelog_entries, issues, version, release_date)
    post_reports_on_slack(reports)
