ENV_RELEASE_DATE = 'release_date'
ENV_SLACK_WEBHOOK = 'slack_webhook'
ENV_VERSION = 'version'

JIRA_DATETIME_FORMAT = '%Y-%m-%dT%H:%M:%S.000+0000'

PROJECT = 'SDK'
BULK_PUBLIC_PR_TITLE = 'Auto Generated Bulk Public'
PR_DESCRIPTION = 'description'
DEXREQ_TICKET_PATTERN = '(DEXREQ-[\d]+)'  # noqa: W605
DEXREQ_TICKET_LABELS = 'labels'

SPEC_REPO = 'dexreq'
SPEC_PR_LINK = '(https://bitbucket.oci.oraclecorp.com/projects/SDK/repos/dexreq/pull-requests/.*)\]'  # noqa: W605
TITLE_FORMAT = '(.*)(API)*'  # noqa: W605

CHANGELOG_TITLE_PATTERN = '[^\w]*changelog[^\w]*'  # noqa: W605
CLI_COMMAND_PATTERN = '[-*`\s]*(oci[^`]*).*'  # noqa: W605
CLI_COMMAND_FORMAT = '  * ``{command}``'
DESCRIPTION_PATTERN = '[^\w[]*'  # noqa: W605
SDK_DESCRIPTION_FORMAT = '* {text}'
CLI_DESCRIPTION_FORMAT = '\n* {text}\n'

SDK_SECTION = 'SDK'
CLI_SECTION = 'CLI'
CLI_MANUAL_CHANGE_SECTION = 'CLI manual change'
JIRA_DOMAIN = 'https://jira.oci.oraclecorp.com/browse/'

PR_SUMMARY_TEMPLATE = '>{count} issue{plural} found for {section}: {issue_list}\n'
SDK_CHANGELOG_REPORT = '>SDK Change Log {mentions}\n{summary}{changelog_entries}\n'
CLI_CHANGELOG_HEADER_TEMPLATE = '{version} - {release_date}\n--------------------\n'
CLI_CHANGELOG_REPORT = '>CLI Change Log {mentions}\n{summary}\n{cli_header}{cli}'
CHANGELOG_ENTRY = '\n\n<{link}|{dexreq_issue}> by {service}\n{content}'

SDK_MENTIONS = ['<@WFABDC812>']
CLI_MENTIONS = ['<@U03KTDB0T8V>']
