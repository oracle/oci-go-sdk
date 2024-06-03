import argparse
import textwrap
import datetime
import sys
import traceback
import time
import json
import six
from string import maketrans
from dotmap import DotMap
from jira.resources import Status

import config
import util

from config import PREVIEW_ISSUE_TYPE_ID, PREVIEW_ISSUE_TYPE_NAME
from config import PUBLIC_ISSUE_TYPE_ID, PUBLIC_ISSUE_TYPE_NAME

import shared.bitbucket_utils

import autogen_issue_advisor_shared

from autogen_issue_advisor_shared import printv
from autogen_issue_advisor_shared import check_should_update
from autogen_issue_advisor_shared import ERROR_CHOOSE_PIPELINE_STATE
from autogen_issue_advisor_shared import QUIET_TIME_MINUTES
from autogen_issue_advisor_shared import DEFAULT_JIRA_ISSUE_FIELDS, CUSTOM_JIRA_ISSUE_FIELDS
from autogen_issue_advisor_shared import DEXREQ_AUTOMATION_NAME
from autogen_issue_advisor_shared import TICKET_STATE_ADVISORY_TEXT
from autogen_issue_advisor_shared import ERROR_STATES
from autogen_issue_advisor_shared import PIPELINE_NAMES
from autogen_issue_advisor_shared import process_last_builds

import autogen_issue_advisor_preview
from autogen_issue_advisor_preview import get_preview_state
from autogen_issue_advisor_preview import PREVIEW_STATES
from autogen_issue_advisor_preview import advise_on_preview_issue
from autogen_issue_advisor_preview import handle_automated_preview_transitions

import autogen_issue_advisor_public
from autogen_issue_advisor_public import get_public_state
from autogen_issue_advisor_public import PUBLIC_STATES
from autogen_issue_advisor_public import advise_on_public_issue
from autogen_issue_advisor_public import handle_automated_public_transitions
from create_cli_design_review_ticket import get_cli_design_review_issues_for_udx
from create_cli_design_review_ticket import is_design_ticket_in_non_terminal_state
from dexreq_migration import process_preview_jira_queue


# Spot testing:
#
# I've tested this is through spot testing: Picking out a DEXREQ ticket from one of the tables, then seeing if it gives the right output.
#
# python autogen_issue_advisor.py --dry-run --issue DEXREQ-123
#
#
# Testing against all issues:
#
# And then I've run it against all issues, regardless of timestamp etc., and made sure it doesn't throw an exception for any of them.
#
# python autogen_issue_advisor.py --dry-run --force
#
# JIRA testing:
#
# I've tested the JIRA portion (without --dry-run) using the "spot testing" procedure.

ignore_wrong_pipeline = False
IGNORE_CHANGES_AFTER = None


def process_changelog(issue):
    # Simplify the changelog
    changelog = issue.changelog
    changelog_list = []
    for history in changelog.histories:
        history_record = DotMap()
        history_record.author = str(history.author)
        history_record.created = history.created
        items_list = []
        for item in history.items:
            item_record = DotMap()
            item_record.field = str(item.field)
            item_record.old = item.fromString
            item_record.new = item.toString
            items_list.append(item_record)
        history_record.changed_items = items_list
        changelog_list.append(history_record)
    return changelog_list


# returns statuses, all, any
def process_statuses(issue):
    sdk_statuses = {}
    for tool_name, jira_field_id in util.get_jira_custom_field_ids_for_tool().items():
        if util.is_tool_jira_reportable(tool_name):
            status = getattr(issue.fields, jira_field_id)
            sdk_statuses[tool_name] = str(status)

    all_sdks = {
        config.CUSTOM_STATUS_TODO: True,
        config.CUSTOM_STATUS_PROCESSING: True,
        config.CUSTOM_STATUS_FAILURE: True,
        config.CUSTOM_STATUS_SUCCESS: True,
        config.CUSTOM_STATUS_DONE: True,
        "None": True
    }
    any_sdks = {
        config.CUSTOM_STATUS_TODO: False,
        config.CUSTOM_STATUS_PROCESSING: False,
        config.CUSTOM_STATUS_FAILURE: False,
        config.CUSTOM_STATUS_SUCCESS: False,
        config.CUSTOM_STATUS_DONE: False,
        "None": False
    }
    for k,status in sdk_statuses.items():
        any_sdks[status] = True
        for status_type in all_sdks.keys():
            if not status_type == status:
                all_sdks[status_type] = False

    return sdk_statuses, all_sdks, any_sdks


def process_pipeline(issue):
    pipeline = None
    ticket_type_id = issue.fields.issuetype.id
    if ticket_type_id == PREVIEW_ISSUE_TYPE_ID:
        pipeline = PREVIEW_ISSUE_TYPE_NAME
    if ticket_type_id == PUBLIC_ISSUE_TYPE_ID:
        pipeline = PUBLIC_ISSUE_TYPE_NAME

    return pipeline


def process_comments(issue):
    comments = []
    if "comment" in issue.raw["fields"] and "comments" in issue.raw["fields"]["comment"]:
        for entry in issue.raw["fields"]["comment"]["comments"]:
            item = DotMap()
            item.author = entry["author"]["displayName"]
            item.created = entry["created"]
            item.url = entry["self"]
            item.text = entry["body"]
            comments.append(item)

    return comments


def process_dates(issue, summary):
    if summary.jira.changelog:
        summary.dates.last_changelog = summary.jira.changelog[-1]
        summary.dates.last.created = summary.dates.last_changelog.created
        summary.dates.last.author = summary.dates.last_changelog.author
    if summary.jira.comments:
        summary.dates.last_comment = summary.jira.comments[-1]
        summary.dates.last.created = summary.dates.last_comment.created
        summary.dates.last.author = summary.dates.last_comment.author

    if summary.dates.last_changelog.created and summary.dates.last_comment.created:
        # both set
        if summary.dates.last_changelog.created > summary.dates.last_comment.created:
            summary.dates.last.created = summary.dates.last_changelog.created
            summary.dates.last.author = summary.dates.last_changelog.author
        else:
            summary.dates.last.created = summary.dates.last_comment.created
            summary.dates.last.author = summary.dates.last_comment.author

    last_issue_advisory = None
    for comment in summary.jira.comments:
        if (autogen_issue_advisor_shared.PROCESS_COMMENTS_BY_ANYONE or comment.author == DEXREQ_AUTOMATION_NAME) and comment.text and TICKET_STATE_ADVISORY_TEXT in comment.text.split("\n")[0]:
            # occurred in first line of comment
            last_issue_advisory = comment

    if last_issue_advisory:
        summary.dates.last_issue_advisory = last_issue_advisory


def process_reporter(issue):
    reporter_record = getattr(issue.fields, "reporter")
    reporter = getattr(reporter_record, "key")
    return reporter


def process_bypass_labels(issue, summary):
    bypass_labels = []

    for l in issue.fields.labels:
        if l in config.BYPASS_LABELS:
            bypass_labels.append(l.encode('utf-8'))

    summary.checks.bypass = bypass_labels

    return bypass_labels


def accept_non_bulk_prs(pr):
    name = None
    if 'name' in pr:
        name = pr['name']
    source_branch = None
    if 'source' in pr and 'branch' in pr['source']:
        source_branch = pr['source']['branch']

    if not name or not source_branch:
        return True

    if name.startswith("Auto Generated Bulk Preview for") and source_branch.startswith(config.GENERATION_BRANCH_PREFIX + "-" + config.BULK_PREVIEW_BRANCH_PREFIX + "-"):
        printv("Rejecting PR because it is a bulk preview PR: {}".format(pr['url']))
        return False

    return True


def process_pull_requests(issue, summary):
    tool_names = util.get_jira_reportable_tool_names()

    # Note: the 'last_update' (which consists of the PR changes and the validation builds)
    # does not currently take 'ignore_changes_after' into account.
    printv("Getting PR status from JIRA for issue {}".format(issue.key), flush=True)
    t0 = time.time()
    pr_status = util.get_pr_status_for_tools(util.JIRA_CLIENT(), issue, tool_names, target_branch_filter='master')
    t1 = time.time()
    printv("JIRA returned PR status for issue {}, took {:.3f} seconds".format(issue.key, t1 - t0), flush=True)

    # TODO: replace this with results from just looking at 'dexreq' and 'python-cli' repos
    summary.pull_requests.last_update = pr_status.last_update
    summary.dates.last_pr_change = pr_status.last_update

    on_service_team = pr_status
    on_service_team.tools = DotMap(on_service_team.tools)

    summary.pull_requests.on_service_team.master = on_service_team

    return on_service_team


def ignore_changes_after(issue):
    if not IGNORE_CHANGES_AFTER:
        return issue

    if "comment" in issue.raw["fields"] and "comments" in issue.raw["fields"]["comment"]:
        to_be_removed = []
        for entry in issue.raw["fields"]["comment"]["comments"]:
            datetime = entry["created"]
            if datetime > IGNORE_CHANGES_AFTER:
                to_be_removed.append(entry)

        for entry in to_be_removed:
            if entry in issue.raw["fields"]["comment"]["comments"]:
                issue.raw["fields"]["comment"]["comments"].remove(entry)
                printv("Ignoring {}".format(entry["self"]))

    already_changed_items = []

    if issue.changelog and issue.changelog.histories:
        to_be_removed = []
        for entry in issue.changelog.histories:
            datetime = entry.created
            if datetime > IGNORE_CHANGES_AFTER:
                to_be_removed.append(entry)
                printv("Ignoring changelog history change at {}".format(entry.created))
                for item in entry.items:
                    if item.field not in already_changed_items:
                        # This is the first time we're seeing a field change after changes
                        # are being ignored; set the field back to the old value
                        field = str(item.field)
                        already_changed_items.append(item.field)
                        if item.field == "status":
                            setattr(issue.fields, "status", Status(None, None, raw={'name': item.fromString}))
                            issue.raw['fields']["status"] = Status(None, None, raw={'name': item.fromString})
                            printv("Fixing {} to '{}'".format(field, str(item.fromString)))
                        if item.field == "labels":
                            setattr(issue.fields, "labels", item.fromString)
                            issue.raw['fields']["labels"] = item.fromString
                            printv("Fixing {} to '{}'".format(field, str(item.fromString)))
                        if item.field in config.CUSTOM_FIELD_NAME_FOR_ID.values():
                            for k, v in config.CUSTOM_FIELD_NAME_FOR_ID.items():
                                if item.field == v:
                                    setattr(issue.fields, k, item.fromString)
                                    issue.raw['fields'][k] = item.fromString
                                    printv("Fixing {} to '{}'".format(field, str(item.fromString)))
                                    break

        for entry in to_be_removed:
            if entry in issue.changelog.histories:
                issue.changelog.histories.remove(entry)

    return issue


# Returns summary, issue
def summarize_issue(issue_key=None):
    dexreq_public_errors = []
    dexreq_public_warnings = []

    printv("Querying JIRA for issue {}".format(issue_key), flush=True)
    t0 = time.time()
    issue = util.get_dexreq_issue(issue_key, fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS), expand=["changelog"],
        errors=dexreq_public_errors, warnings=dexreq_public_warnings)
    t1 = time.time()
    printv("JIRA returned issue {}, took {:.3f} seconds".format(issue_key, t1 - t0), flush=True)

    ignore_changes_after(issue)

    summary = DotMap()

    summary.errors = []
    summary.errors.extend(dexreq_public_errors)
    summary.errors.extend("Warning: {}".format(w) for w in dexreq_public_warnings)

    pipeline = process_pipeline(issue)
    summary.pipeline = pipeline

    summary.jira.reporter = process_reporter(issue)
    summary.jira.changelog = process_changelog(issue)
    summary.jira.comments = process_comments(issue)

    process_dates(issue, summary)

    t0 = time.time()
    process_last_builds(issue, summary)
    t1 = time.time()
    printv("process_last_builds took {:.3f} seconds".format(t1 - t0))

    process_bypass_labels(issue, summary)

    sdk_statuses, all_sdks, any_sdks = process_statuses(issue)

    # TODO: collect latest branches for each SDK

    summary.sdks.statuses = sdk_statuses
    summary.sdks.any = any_sdks
    summary.sdks.all = all_sdks

    summary.jira.status = str(getattr(issue.fields, "status"))

    value = getattr(issue.fields, config.CUSTOM_FIELD_ID_PREVIEW_ISSUE)
    if value:
        summary.preview_issues = ",".join(util.get_dexreq_issue_keys(str(value)))

    for field_name in issue.raw['fields']:
        field_value = issue.raw['fields'][field_name]
        if field_name in config.CUSTOM_FIELD_NAME_FOR_ID:
            field_name = config.CUSTOM_FIELD_NAME_FOR_ID[field_name]

        # this deals with the custom status fields, which look like this in the raw:
        # {u'self': u'https://jira.oci.oraclecorp.com/rest/api/2/customFieldOption/14502', u'id': u'14502', u'value': u'Pending Merge'}
        if type(field_value) is dict and "self" in field_value and "id" in field_value and ("value" in field_value or "name" in field_value):
            field_value = field_value["value"] if "value" in field_value else field_value["name"]

        encoded_field_name = field_name.encode('utf-8').lower().translate(maketrans(' /', '__'), '()-').replace('_&_', '_and_').replace('&', '_and_')
        if field_name not in ['comment'] and encoded_field_name not in summary.jira:
            summary.jira[encoded_field_name] = field_value

    process_pull_requests(issue, summary)

    if pipeline is None:
        state = ERROR_CHOOSE_PIPELINE_STATE
    elif pipeline is PREVIEW_ISSUE_TYPE_NAME:
        state = get_preview_state(issue, summary, sdk_statuses, all_sdks, any_sdks)
    elif pipeline is PUBLIC_ISSUE_TYPE_NAME:
        state = get_public_state(issue, summary, sdk_statuses, all_sdks, any_sdks)
    else:
        state = None

    summary.state = state

    if summary.pipeline == config.PUBLIC_ISSUE_TYPE_NAME:
        if update_summary_with_pending_cli_design_review_tickets(issue, summary):
            issue = util.get_dexreq_issue(issue.key, fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS + process_preview_jira_queue.CUSTOM_JIRA_ISSUE_FIELDS_PER_BUILD_TYPE[config.BUILD_TYPE_INDIVIDUAL_PUBLIC]))

    return summary, issue


def update_summary_with_pending_cli_design_review_tickets(issue, summary):
    udx_issue_keys = util.get_udx_issue_keys(getattr(issue.fields, config.CUSTOM_FIELD_ID_UDX_TICKET))
    printv("Checking for CLI Design Review issues for: {}".format(", ".join(udx_issue_keys)))
    design_review_tickets = get_cli_design_review_issues_for_udx(udx_issue_keys)

    pending_design_review_tickets = get_pending_cli_design_review_tickets(design_review_tickets)
    summary.cli.pending_design_reviews = [i.key for i in pending_design_review_tickets]

    printv("Checking for manual changes:")
    requires_manual_changes = are_cli_manual_changes_required(design_review_tickets)

    if requires_manual_changes and config.CLI_REQUIRED_MANUAL_CHANGES_LABEL not in issue.fields.labels:
        printv("Adding CLI-ManualChangesRequired label to: " + issue.key)
        issue.add_field_value('labels', config.CLI_REQUIRED_MANUAL_CHANGES_LABEL)
        # True indicates issue needs to be refreshed.
        return True

    # No need to update the issue field
    return False


def get_pending_cli_design_review_tickets(design_review_tickets):
    pending_design_review_tickets = filter(lambda issue: is_design_ticket_in_non_terminal_state(issue), design_review_tickets)
    printv("Pending CLI Design Reviews: " + ', '.join(issue.key for issue in pending_design_review_tickets))

    return pending_design_review_tickets


def are_cli_manual_changes_required(design_review_tickets):
    for issue in design_review_tickets:
        for label in issue.fields.labels:
            if label in config.CLI_MANUAL_CHANGES_LABELS:
                printv('Found manual changes for CLI Design review issue: {}'.format(issue.key))
                return True
    return False


CHOOSE_PIPELINE_ERROR_TEMPLATE = """\
[~{reporter}], the ticket does not specify if this is for the preview pipeline or the public pipeline.

If this is for the preview pipeline (LA), choose the '{preview_type}' ticket type.

If this is for the public pipeline (GA), choose the '{public_type}' ticket type.

In order to change ticket types, select the 'Move' action in the 'More' menu of the JIRA issue.

([Preview wiki|https://confluence.oci.oraclecorp.com/display/DEX/Requesting+a+preview+SDK+CLI#RequestingapreviewSDKCLI-A)Opentheticket] and [public wiki|https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=43683000#RequestingapublicSDK/CLI-Step3:Createaticketinthe%22DEX-REQUEST%22JIRA])"""


def for_choose_pipeline_error(issue_key, summary):
    global ignore_wrong_pipeline
    
    text = ""

    if not summary.pipeline:
        if ignore_wrong_pipeline:
            print("Ignoring {}, even though it has not chosen a correct ticket type.".format(issue_key))
        else:
            text = CHOOSE_PIPELINE_ERROR_TEMPLATE.format(
                preview_type=config.PREVIEW_ISSUE_TYPE_NAME,
                public_type=config.PUBLIC_ISSUE_TYPE_NAME,
                reporter=summary.jira.reporter)

    return text


def handle_automated_transitions(issue_key, issue, summary):
    text = None
    transitioned = False
    if summary.pipeline == config.PREVIEW_ISSUE_TYPE_NAME:
        transitioned, text = handle_automated_preview_transitions(issue_key, issue, summary)
    elif summary.pipeline == config.PUBLIC_ISSUE_TYPE_NAME:
        transitioned, text = handle_automated_public_transitions(issue_key, issue, summary)
    if transitioned:
        printv("Automated transition:\n{}".format(text))
    else:
        printv("No automated transition.{}".format((" Text:\n" + text) if text else ""))

    return transitioned, text


class CustomEncoder(json.JSONEncoder):
    def default(self, obj):
        if (isinstance(obj, autogen_issue_advisor_shared.ServiceTeamMasterPrs) or  # noqa: ignore=W504
                isinstance(obj, autogen_issue_advisor_shared.BitbucketBuilds) or  # noqa: ignore=W504
                isinstance(obj, autogen_issue_advisor_shared.BitbucketBuildChecks) or  # noqa: ignore=W504
                isinstance(obj, util.PrStatusForTools) or  # noqa: ignore=W504
                isinstance(obj, util.PrsPerTool) or  # noqa: ignore=W504
                isinstance(obj, util.PrAndUrl)):  # noqa: ignore=W504
            return dict(obj.__dict__)
        return json.JSONEncoder.default(self, obj)


# Returns text, summary
def advise_on_issue(issue_key, force):
    summary, issue = summarize_issue(issue_key=issue_key)

    # Flush, so we make sure the output of the summarize_issue function is already there
    # NOTE: This is to help debug for DEX-6382
    sys.stdout.flush()

    printv("summary:")
    if autogen_issue_advisor_shared.IS_VERBOSE:
        print(json.dumps(summary, cls=CustomEncoder, sort_keys=True, indent=2))

    should_update = check_should_update(summary)

    transitioned, transition_text = handle_automated_transitions(issue_key, issue, summary)

    if transitioned:
        # Always update after a transition
        should_update = True

        # Get the latest post-transition summary
        new_summary, issue = summarize_issue(issue_key=issue_key)
        new_summary.transition_from_state = summary.state
        summary = new_summary

    if not should_update:
        print("Should not update")
        if not force:
            return None, summary
        else:
            print("Forcing update...")

    text = None

    if summary.state == ERROR_CHOOSE_PIPELINE_STATE:
        text = for_choose_pipeline_error(issue_key, summary)
    elif summary.pipeline == config.PREVIEW_ISSUE_TYPE_NAME:
        text = advise_on_preview_issue(issue_key, issue, summary)
    else:
        text = advise_on_public_issue(issue_key, issue, summary)

    if transition_text and text:
        text = transition_text + "\n" + text
    elif transition_text:
        text = transition_text

    if text:
        text = "{}\n\n{}".format(TICKET_STATE_ADVISORY_TEXT, text)

    return text, summary


def initialize_state_statistics(only_pipeline):
    state_statistics = {}
    if not only_pipeline or only_pipeline == PREVIEW_ISSUE_TYPE_NAME:
        for state in PREVIEW_STATES:
            state_statistics[state] = 0
            state_statistics["TRANSITION_FROM_{}".format(state)] = 0
    if not only_pipeline or only_pipeline == PUBLIC_ISSUE_TYPE_NAME:
        for state in PUBLIC_STATES:
            state_statistics[state] = 0
            state_statistics["TRANSITION_FROM_{}".format(state)] = 0
    for state in ERROR_STATES:
        state_statistics[state] = 0

    return state_statistics


def print_state_statistics(state_statistics, states):
    for state in states:
        if state in state_statistics:
            count = state_statistics[state]
            print("##teamcity[buildStatisticValue key='{}_count' value='{}']".format(state, count))


def query_all_issues(only_pipeline):
    # Query all unresolved issues, plus those that were set to 'Done' in the last 4 days (to make it easier to span weekends, in case automation fails)
    query = 'project = {JIRA_PROJECT} AND (resolution = Unresolved OR (status changed to Done AFTER "-4d"))'.format(
        JIRA_PROJECT=config.JIRA_PROJECT)

    if only_pipeline:
        query = query + ' AND issuetype = "{}"'.format(only_pipeline)

    # Since here, we're only interested in issue key and summary, let's not use
    # the util.search_dexreq_issues function that also looks at the linked preview
    # ticket in the case of public tickets
    all_issues = util.jira_search_issues(query, fields=['summary, created'])

    for issue in all_issues:
        try:
            print('{} - {}'.format(issue.key, issue.fields.summary.encode('utf-8')))
        except Exception as error:
            print('Problem with {}'.format(issue.key))
            exception_string = traceback.format_exc()
            print("Unexpected error: {}\n{}".format(type(error), exception_string))

    return all_issues


def post_comment_if_new(issue_key, summary, text, force, commented_issues):
    should_comment = True
    if not force and "{color:" + config.COMMENT_TYPE_TO_COLOR[config.COMMENT_TYPE_INFO] + "}" + text + "{color}" == summary.dates.last_issue_advisory.text:
        printv("Not commenting, because the last issue advisory on {} had the same text.".format(summary.dates.last_issue_advisory.created))
        should_comment = False

    if should_comment:
        # comment on issue
        util.add_jira_comment(
            issue_key,
            text,
            comment_type=config.COMMENT_TYPE_INFO
        )
        commented_issues.append(issue_key)

    return should_comment


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Ticket advisor (preview and public).')
    parser.add_argument('--issue',
                        action='append',
                        help='By default, we query JIRA. This allows you to specify a DEXREQ issue to process instead: --issue DEXREQ-123')
    parser.add_argument('--bulk-preview-date-overrides',
                        help='''Date overrides for bulk preview. Comma-separated list, with each part being either "+YYYY-MM-DD" or "-YYYY-MM-DD".
                        Example "+2018-10-04,-2018-10-05" means "do a bulk preview on October 4th (Thursday), but not on October 5th (Friday)"''')
    parser.add_argument('--public-release-date-overrides',
                        help='''Date overrides for public release. Comma-separated list, with each part being either "+YYYY-MM-DD", "-YYYY-MM-DD", or "=YYYY-MM-DD@GA".
                        Starting with "=" sets the base date of the two-week cadence, and "@GA" indicates that this should be the GA with the number "GA". If not set, it's {}.
                        Example "=2018-10-11@123,+2018-10-24,-2018-10-25" means "do releases every two weeks, starting with GA 123 on 2018-10-11, but don't do one on October 25th, do it on October 24th instead"'''.format(autogen_issue_advisor_public.PUBLIC_RELEASE_DATE_CADENCE_START_DATE))
    parser.add_argument('--public-release-requested-cut-off-date-overrides',
                        help='''Date overrides for public 'release requested' cut-off date. Comma-separated list, with each part being either "+YYYY-MM-DD", "-YYYY-MM-DD", or "=YYYY-MM-DD".
                        Starting with "=" sets the base date of the two-week cadence. If not set, it's {}.
                        Example "=2018-10-05,+2018-10-18,-2018-10-19" means "do RR cut-off every two weeks, starting with 2018-10-05, but don't do one on October 19th, do it on October 18th instead"'''.format(autogen_issue_advisor_public.RELEASE_REQUESTED_CUT_OFF_DATE_CADENCE_START_DATE))
    parser.add_argument('--dry-run',
                        default=False,
                        action='store_true',
                        help='Perform a dry-run')
    parser.add_argument('--force',
                        default=False,
                        action='store_true',
                        help='Force an update, even if the issue has not been quiet for {} minutes and there have been no updates since the last advisory.'.format(QUIET_TIME_MINUTES))
    parser.add_argument('--verbose',
                        default=False,
                        action='store_true',
                        help='Verbose logging')
    parser.add_argument('--pipeline',
                        help='Limit to pipeline, one of {}. Default is both.'.format(", ".join(['"{}"'.format(x) for x in PIPELINE_NAMES])))
    parser.add_argument('--show-ga-calendar',
                        default=False,
                        action='store_true',
                        help='''Show a GA calendar (number of future releases specified using --show-ga-calendar-count).''')
    parser.add_argument('--show-ga-calendar-count',
                        default=10,
                        help='''When showing a GA calendar, show the specified number of future releases.''')
    parser.add_argument('--show-preview-calendar',
                        default=False,
                        action='store_true',
                        help='''Show a Preview calendar (number of future preview releases specified using --show-preview-calendar-count).''')
    parser.add_argument('--show-preview-calendar-count',
                        default=10,
                        help='''When showing a preview calendar, show the specified number of future preview releases.''')
    parser.add_argument('--ignore-wrong-pipeline',
                        default=False,
                        action='store_true',
                        help='Ignore wrong pipeline (i.e. not "{}" or "{}" issue type)'.format(config.PREVIEW_ISSUE_TYPE_NAME, config.PUBLIC_ISSUE_TYPE_NAME))
    parser.add_argument('--process-comments-by-anyone',
                        default=False,
                        action='store_true',
                        help='Process comments by anyone, not just by "{}"'.format(DEXREQ_AUTOMATION_NAME))
    parser.add_argument('--ignore-changes-after',
                        help='Ignore changes after the specified datetime YYYY-MM-DDThh:mm:ss')
    parser.add_argument('--disable_date_check',
                        default=False,
                        action='store_true',
                        help='Check if the provided GA date is a public release date. Leave a warning if not')

    failed = []

    args = parser.parse_args()
    autogen_issue_advisor_shared.IS_VERBOSE = args.verbose
    util.IS_VERBOSE = args.verbose
    shared.bitbucket_utils.verbose = args.verbose

    autogen_issue_advisor_shared.PROCESS_COMMENTS_BY_ANYONE = args.process_comments_by_anyone

    config.IS_DRY_RUN = args.dry_run
    shared.bitbucket_utils.dry_run = args.dry_run

    config.DISABLE_COMMENT_INCORRECT_DATES = args.disable_date_check

    only_pipeline = args.pipeline
    force = args.force

    # date overrides
    cut_off_date_or = args.public_release_requested_cut_off_date_overrides
    release_date_or = args.public_release_date_overrides

    # If one of the overrides contains '=', the other should as well
    if release_date_or and cut_off_date_or:
        if ('=' in release_date_or and '=' not in cut_off_date_or) or ('=' not in release_date_or and '=' in cut_off_date_or):
            raise ValueError("\'=\' found in either public_release_requested_cut_off_date_overrides or public_release_date_overrides, but not both.")
        
    # Check that base dates are within 3 weeks of each other
    base_cut_off_date, ga_number, cut_off_overrides = autogen_issue_advisor_shared.process_date_override(cut_off_date_or,
        autogen_issue_advisor_public.RELEASE_REQUESTED_CUT_OFF_DATE_CADENCE_START_DATE, None, '--public-to-deploy-cut-off-date-overrides')
    base_release_date, base_release_ga_number, release_overrides = autogen_issue_advisor_shared.process_date_override(release_date_or,
        autogen_issue_advisor_public.PUBLIC_RELEASE_DATE_CADENCE_START_DATE, autogen_issue_advisor_public.PUBLIC_RELEASE_DATE_GA_NUMBER, '--public-release-date-overrides')
    
    delta = base_cut_off_date - base_release_date
    delta = abs(delta.days)

    if (delta > 21):
        raise ValueError("public_release_requested_cut_off_date_overrides and public_release_date_overrides should be within 3 weeks of each other") 

    autogen_issue_advisor_preview.BULK_PREVIEW_DATE_OVERRIDE = args.bulk_preview_date_overrides
    autogen_issue_advisor_public.PUBLIC_RELEASE_DATE_OVERRIDE = release_date_or
    autogen_issue_advisor_public.RELEASE_REQUESTED_CUT_OFF_DATE_OVERRIDE = cut_off_date_or

    IGNORE_CHANGES_AFTER = args.ignore_changes_after

    if args.show_ga_calendar:
        autogen_issue_advisor_public.show_ga_calendar(int(args.show_ga_calendar_count))
        sys.exit(0)

    if args.show_preview_calendar:
        autogen_issue_advisor_preview.show_preview_calendar(int(args.show_preview_calendar_count))
        sys.exit(0)

    shared.bitbucket_utils.setup_bitbucket(args)

    if args.pipeline and args.pipeline not in PIPELINE_NAMES:
        raise ValueError("Pipeline, if specified, must be one of: {}".format(", ".join(['"{}"'.format(x) for x in PIPELINE_NAMES])))

    if args.issue:
        if only_pipeline:
            raise ValueError("Cannot use --issue with --pipeline.")
        issues = [util.JIRA_CLIENT().issue(issue_key, fields=['summary, created']) for issue_key in args.issue]
    else:
        issues = query_all_issues(only_pipeline)

    # Find the oldest issue
    oldest_created_date = None
    oldest_ticket = None
    for issue in issues:
        created_date = getattr(issue.fields, 'created')
        if not oldest_created_date or created_date < oldest_created_date:
            oldest_created_date = created_date
            oldest_ticket = issue.key

    print("Oldest ticket was created {} ({})".format(oldest_created_date, oldest_ticket))

    # Prime the caches
    all_repo_names = [config.DEXREQ_REPO_NAME]
    for tool, repo_names in six.iteritems(config.REPO_NAMES_FOR_TOOL):
        all_repo_names.extend(repo_names)

    for repo_name in set(all_repo_names):
        printv("Priming PR cache for repo {}".format(repo_name))
        shared.bitbucket_utils.get_all_pullrequest_with_string_after('SDK', repo_name, '', oldest_created_date)

    if args.ignore_wrong_pipeline:
        ignore_wrong_pipeline = True

    print("##teamcity[testSuiteStarted name='autogen_issue_advisor']")

    state_statistics = initialize_state_statistics(only_pipeline)

    commented_issues = []
    for issue_key in [issue.key for issue in issues]:
        if config.should_ignore_issue(issue_key):
            print(textwrap.dedent("""\
            ========================================
            Issue: {} - being ignored per env var {}""").format(issue_key, config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME))
            print("========================================")
            continue

        print(textwrap.dedent("""\
        ========================================
        Issue: {}""").format(issue_key))
        issue_start = datetime.datetime.now()
        print("##teamcity[testStarted name='{}' captureStandardOutput='true']".format(issue_key))
        try:
            # Flush, so we make sure the output of the issue key is already visible
            # NOTE: This is to help debug for DEX-6382
            sys.stdout.flush()
            text, summary = advise_on_issue(issue_key, force)

            if summary.state:
                state_statistics[summary.state] += 1

            if summary.transition_from_state:
                printv("summary.transition_from_state: {}".format(summary.transition_from_state))
                state_statistics["TRANSITION_FROM_{}".format(summary.transition_from_state)] += 1

            if text:
                if summary.errors:
                    text = text + "\n\nErrors the advisor noticed:\n{}".format("\n".join("- {}".format(e) for e in summary.errors))

                post_comment_if_new(issue_key, summary, text, force, commented_issues)
        except Exception as error:
            exception_string = traceback.format_exc()
            print("Unexpected error: {}\n{}".format(type(error), exception_string))
            failed.append(issue_key)
            print("##teamcity[testFailed name='{}' message='{}' details='{}']".format(issue_key, type(error), exception_string))
        finally:
            issue_end = datetime.datetime.now()

            print("##teamcity[testFinished name='{}' duration='{}']".format(issue_key, int((issue_end - issue_start).total_seconds() * 1000)))

    print("##teamcity[testSuiteFinished name='autogen_issue_advisor']")
    print("========================================")

    print("##teamcity[buildStatisticValue key='{}' value='{}']".format("comment_count", len(commented_issues)))
    print("##teamcity[buildStatisticValue key='{}' value='{}']".format("error_count", len(failed)))
    transition_count = sum([state_statistics["TRANSITION_FROM_{}".format(state)] for state in PREVIEW_STATES + PUBLIC_STATES])
    print("##teamcity[buildStatisticValue key='{}' value='{}']".format("transition_count", transition_count))

    print_state_statistics(state_statistics, PREVIEW_STATES)
    print_state_statistics(state_statistics, PUBLIC_STATES)
    print_state_statistics(state_statistics, ERROR_STATES)
    print_state_statistics(state_statistics, ["TRANSITION_FROM_{}".format(state) for state in PREVIEW_STATES])
    print_state_statistics(state_statistics, ["TRANSITION_FROM_{}".format(state) for state in PUBLIC_STATES])

    if config.IS_DRY_RUN:
        print("DRY-RUN: Would have left {} comment(s)".format(len(commented_issues)))
    else:
        print("Left {} comment(s)".format(len(commented_issues)))

    if commented_issues:
        print("Commented on the following issues:\n{}".format("\n".join(commented_issues)))

    if failed:
        print("The following issues failed:\n{}".format("\n".join(failed)))
        sys.exit(1)
