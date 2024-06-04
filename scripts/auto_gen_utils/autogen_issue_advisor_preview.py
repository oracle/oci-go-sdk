import datetime
import pytz
import sys

import config
import util

from autogen_issue_advisor_shared import PACIFIC_TIME_ZONE
from autogen_issue_advisor_shared import get_successful_pull_requests_text
from autogen_issue_advisor_shared import get_failed_links_text
from autogen_issue_advisor_shared import handle_transition_for_processing
from autogen_issue_advisor_shared import execute_appropriate_transition_handler
from autogen_issue_advisor_shared import handle_transition_for_spec_review_pr
from autogen_issue_advisor_shared import get_spec_review_pr
from autogen_issue_advisor_shared import process_all_spec_change_prs
from autogen_issue_advisor_shared import process_date_override, find_next_matching_date

# Format: comma-separated, with each part being either "+YYYY-MM-DD" or "-YYYY-MM-DD"
# Example "+2018-10-04,-2018-10-05" means "do a bulk preview on October 4th (Thursday), but not on October 5th (Friday)"
BULK_PREVIEW_DATE_OVERRIDE = None

# Format: comma-separated, with each part being either "+YYYY-MM-DD", "-YYYY-MM-DD", or "=YYYY-MM-DD".
# Starting with "=" sets the base date of the one-week cadence. If not set, it's 2019-01-08.
# Example "=2018-10-11,+2018-10-24,-2018-10-25" means
# "do releases every week, starting with 2018-10-11, but don't do one on October 25th, do it on October 24th instead"
PREVIEW_RELEASE_DATE_OVERRIDE = None
PREVIEW_RELEASE_DATE_CADENCE_START_DATE = "2022-06-08"
PREVIEW_RELEASE_DATE_CADENCE_IN_DAYS = 7


PREVIEW_NON_ACTIONABLE_STATE = "PREVIEW_NON_ACTIONABLE_STATE"
PREVIEW_PROCESSING_REQUESTED_STATE = "PREVIEW_PROCESSING_REQUESTED_STATE"
PREVIEW_PROCESSING_STATE = "PREVIEW_PROCESSING_STATE"
PREVIEW_SERVICE_TEAM_FAILURE_INVESTIGATION_STATE = "PREVIEW_SERVICE_TEAM_FAILURE_INVESTIGATION_STATE"
PREVIEW_DEX_SUPPORT_REQUIRED_STATE = "PREVIEW_DEX_SUPPORT_REQUIRED_STATE"
PREVIEW_SERVICE_TEAM_REVIEW_REQUIRED_STATE = "PREVIEW_SERVICE_TEAM_REVIEW_REQUIRED_STATE"
PREVIEW_READY_FOR_PREVIEW_STATE = "PREVIEW_READY_FOR_PREVIEW_STATE"
PREVIEW_PROCESSING_BULK_STATE = "PREVIEW_PROCESSING_BULK_STATE"
PREVIEW_DEX_BULK_REVIEW_STATE = "PREVIEW_DEX_BULK_REVIEW_STATE"
PREVIEW_DONE_STATE = "PREVIEW_DONE_STATE"

PREVIEW_STATES = [
    PREVIEW_NON_ACTIONABLE_STATE,
    PREVIEW_PROCESSING_REQUESTED_STATE,
    PREVIEW_PROCESSING_STATE,
    PREVIEW_SERVICE_TEAM_FAILURE_INVESTIGATION_STATE,
    PREVIEW_DEX_SUPPORT_REQUIRED_STATE,
    PREVIEW_SERVICE_TEAM_REVIEW_REQUIRED_STATE,
    PREVIEW_READY_FOR_PREVIEW_STATE,
    PREVIEW_PROCESSING_BULK_STATE,
    PREVIEW_DEX_BULK_REVIEW_STATE,
    PREVIEW_DONE_STATE,
]


NON_ACTIONABLE_BECAUSE_OF_BACKLOG_TEMPLATE = """\
[~{reporter}], the ticket status is set to '{backlog_status}'. For processing to begin, please set the ticket status to '{requested_status}'.
"""


def for_non_actionable(issue_key, issue, summary):
    text = ""

    if summary.jira.status == config.STATUS_BACKLOG:
        if text:
            text += "\n\n"

        text = text + NON_ACTIONABLE_BECAUSE_OF_BACKLOG_TEMPLATE.format(
            reporter=summary.jira.reporter,
            backlog_status=config.STATUS_BACKLOG,
            requested_status=config.STATUS_PROCESSING_REQUESTED)

    return text


SERVICE_TEAM_FAILURE_INVESTIGATION_TEMPLATE = """\
[~{reporter}], the status of the ticket is set to '{ticket_status}', because generation of {an_sdk_several_sdks} has failed, as indicated by the SDK status {field_fields} set to '{sdk_state}'. For {this_sdk_these_sdks}, please look into why the generation failed by examining the build log.

If you determine that something was wrong in your spec, please change your spec and generate a new spec artifact. Then update the spec version in this ticket and set the ticket status back to '{requested_status}'.

If you think the error is unrelated to your change, or if you need help figuring this out, set the ticket status to '{dex_support_status}', and our on-call engineer will look at it within a day.

The failed {sdk_is_sdks_are}:

{fields}{failed_builds_text}

([wiki|https://confluence.oci.oraclecorp.com/display/DEX/Requesting+a+preview+SDK+CLI#RequestingapreviewSDKCLI-Step4:Monitorticketstatusforautomatedupdatesandtakeactionifnecessary])"""


def for_service_team_failure_investigation(issue_key, issue, summary):
    text = ""

    failed = [(language, value) for language, value in summary.sdks.statuses.items() if value == config.CUSTOM_STATUS_FAILURE]
    fields = "\n".join(["- {}".format(language) for language, value in failed])

    failed_builds_text = get_failed_links_text(summary) or ""

    text = SERVICE_TEAM_FAILURE_INVESTIGATION_TEMPLATE.format(
        reporter=summary.jira.reporter,
        fields=fields,
        ticket_status=config.STATUS_SERVICE_TEAM_FAILURE_INVESTIGATION,
        sdk_state=config.CUSTOM_STATUS_FAILURE,
        dex_support_status=config.STATUS_DEX_SUPPORT_REQUIRED,
        requested_status=config.STATUS_PROCESSING_REQUESTED,
        an_sdk_several_sdks="an SDK" if len(failed) == 1 else "several SDKs",
        field_fields="field" if len(failed) == 1 else "fields",
        this_sdk_these_sdks="this SDK" if len(failed) == 1 else "these SDKs",
        sdk_is_sdks_are="SDK is" if len(failed) == 1 else "SDKs are",
        failed_builds_text=failed_builds_text)

    return text


SERVICE_TEAM_REVIEW_REQUIRED_TEMPLATE = """\
The preview SDKs have been generated. The service team has to review the output and approve it before the ticket can move on.

[~{reporter}], {review_text}{pr_text}
([wiki|https://confluence.oci.oraclecorp.com/display/DEX/Requesting+a+preview+SDK+CLI#RequestingapreviewSDKCLI-Step5:Verifygeneratedchanges])"""


REVIEW_GENERATED_CODE_TEXT = """\
check that the generated SDKs contain (1) your entire change, and (2) nothing unwanted.

If you are satisfied, then set the '{ready_status}' status for this ticket.

If you find that something is missing from your change or there is something that should not be there, then please revise your spec. Generate a new spec artifact, then update the spec version in this ticket and set the ticket status back to '{requested_status}'."""


REVIEW_SPEC_DIFF_TEXT = """\
review the [spec change pull request|{pr_url}] to make sure it contains (1) your entire change, and (2) nothing unwanted.

If you are satisfied, then approve the spec pull request and set the '{ready_status}' status for this ticket.

If you find that something is missing from your change or there is something that should not be there, then please set the spec pull request to "Needs Work" and revise your spec. Generate a new spec artifact, then update the spec version in this ticket and set the ticket status back to '{requested_status}'.

Spec change pull request to review: {pr_url}"""


def get_preview_review_text(issue_key, reporter):
    dexreq_pr_url = None
    dexreq_pr = get_spec_review_pr(issue_key)

    if dexreq_pr:
        hrefs = util.deep_get(dexreq_pr, 'links.self')
        if hrefs:
            dexreq_pr_url = util.deep_get(hrefs[0], 'href')

    if dexreq_pr_url:
        try:
            dexreq_pr_id = dexreq_pr['id']
            util.add_reviewer_to_pull_request(dexreq_pr_id, config.DEXREQ_REPO_NAME, [reporter])
        except:  # noqa: ignore=E722
            print("Could not add '{}' as reviewer".format(reporter))
            print('Error: {}. {}, line: {}'.format(sys.exc_info()[0],
                                                 sys.exc_info()[1],
                                                 sys.exc_info()[2].tb_lineno))
        return REVIEW_SPEC_DIFF_TEXT.format(ready_status=config.STATUS_READY_FOR_PREVIEW,
            requested_status=config.STATUS_PROCESSING_REQUESTED,
            pr_url=dexreq_pr_url), dexreq_pr, dexreq_pr_url

    return REVIEW_GENERATED_CODE_TEXT.format(ready_status=config.STATUS_READY_FOR_PREVIEW,
        requested_status=config.STATUS_PROCESSING_REQUESTED), None, None


def for_service_team_review_required(issue_key, issue, summary):
    review_text, dexreq_pr, dexreq_pr_url = get_preview_review_text(issue_key, summary.jira.reporter)
    pr_text = get_successful_pull_requests_text(summary)
    if pr_text:
        if dexreq_pr_url:
            pr_text = "\n\nThe links below allow you to examine the effect of your change on the generated source. All you have to do is review the above [spec change pull request|{pr_url}]. The SDK/CLI team will do all the work to include the change in the weekly bulk preview.\n\n".format(pr_url=dexreq_pr_url) + pr_text
        else:
            pr_text = "\n\nThe links below allow you to examine the effects of your change on the generated source. Just look at the diff and verify that the generated SDKs contain the change you expect. The SDK/CLI team will do all the work to include the change in the weekly bulk preview.\n\n" + pr_text
    else:
        pr_text = "\n"

    text = SERVICE_TEAM_REVIEW_REQUIRED_TEMPLATE.format(
        reporter=summary.jira.reporter,
        pr_text=pr_text,
        review_text=review_text)

    return text


PROCESSING_TEMPLATE = """\
The ticket is being processed. Please wait for the automation to produce results.
"""

PARTIALLY_PROCESSED_TEMPLATE = """\
The ticket is being processed, and some SDKs have already been generated.

Already done:
{fields}"""


def for_processing(issue_key, issue, summary):
    fields = "\n".join(["- {}".format(language) for language, value in summary.sdks.statuses.items()
        if value in [config.CUSTOM_STATUS_FAILURE, config.CUSTOM_STATUS_SUCCESS, config.CUSTOM_STATUS_DONE]])
    if fields:    
        return PARTIALLY_PROCESSED_TEMPLATE.format(fields=fields, reporter=summary.jira.reporter)
    else:
        return PROCESSING_TEMPLATE.format(reporter=summary.jira.reporter)


PROCESSING_BULK_TEMPLATE = """\
The ticket is being processed for the bulk preview. Please wait for the automation to produce results.
"""

PARTIALLY_PROCESSED_BULK_TEMPLATE = """\
The ticket is being processed for the bulk preview, and some SDKs have already been generated.

Already done:
{fields}"""


def for_processing_bulk(issue_key, issue, summary):
    fields = "\n".join(["- {}".format(language) for language, value in summary.sdks.statuses.items()
        if value in [config.CUSTOM_STATUS_FAILURE, config.CUSTOM_STATUS_SUCCESS, config.CUSTOM_STATUS_DONE]])
    if fields:    
        return PARTIALLY_PROCESSED_BULK_TEMPLATE.format(fields=fields, reporter=summary.jira.reporter)
    else:
        return PROCESSING_BULK_TEMPLATE.format(reporter=summary.jira.reporter)


def for_processing_requested(issue_key, issue, summary):
    return "This ticket is ready for automated processing. Please wait for the automation to start."


def get_preview_review_complete_date(issue_key, summary):
    # Find timestamp when "Ready for Preview" was last set
    preview_review_complete_date = None
    if summary.jira and summary.jira.changelog:
        for cl in summary.jira.changelog:
            for item in cl.changed_items:
                if item.field == "status" and item.new == config.STATUS_READY_FOR_PREVIEW:
                    # status was changed to config.STATUS_READY_FOR_PREVIEW
                    preview_review_complete_datetime = pytz.utc.localize(datetime.datetime.strptime(cl.created, '%Y-%m-%dT%H:%M:%S.%f+0000'))

                    # change this to Pacific time

                    preview_review_complete_datetime_pacific = preview_review_complete_datetime.astimezone(PACIFIC_TIME_ZONE)                     

                    preview_review_complete_date = preview_review_complete_datetime_pacific.date()

    return preview_review_complete_date


def get_next_bulk_preview_date(issue_key, summary):
    # Start by finding the "Ready for Preview" timestamp
    preview_review_complete_date = get_preview_review_complete_date(issue_key, summary)
    print("Status '{}' was set for issue '{}' on: {}".format(config.STATUS_READY_FOR_PREVIEW, issue_key, preview_review_complete_date))

    if not preview_review_complete_date:
        # Couldn't figure out when label was added, decline to give a date
        return None

    base_release_date, base_release_ga_number, release_overrides = \
        process_date_override(BULK_PREVIEW_DATE_OVERRIDE, PREVIEW_RELEASE_DATE_CADENCE_START_DATE, 140,
                              '--bulk-preview-date-overrides')

    bulk_preview_date = base_release_date

    # Find the next Wednesday after this date (if added on a Wednesday, go to next Wednesday)
    while True:
        bulk_preview_date = find_next_matching_date(bulk_preview_date, base_release_date,
                                                    release_overrides, PREVIEW_RELEASE_DATE_CADENCE_IN_DAYS,
                                                    "preview release")
        print("Finding next bulk preview date, preview_review_complete_date='{}', candidate='{}'".format(preview_review_complete_date, bulk_preview_date))

        if bulk_preview_date > preview_review_complete_date:
            break

        bulk_preview_date += datetime.timedelta(1)

    print("Next bulk preview for issue '{}' is scheduled for: {}".format(issue_key, bulk_preview_date))

    return bulk_preview_date


def for_ready_for_bulk(issue_key, issue, summary):
    text = "This ticket is ready for the next bulk preview SDK generation."

    date = get_next_bulk_preview_date(issue_key, summary)
    if date:
        preview_review_complete_date = get_preview_review_complete_date(issue_key, summary)
        text = text + "\n\nSince the '{}' status was set on {} (Seattle time), the next bulk preview generation that this change can be included in is scheduled for {}.".format(
            config.STATUS_READY_FOR_PREVIEW, preview_review_complete_date.strftime('%A, %d %B %Y'), date.strftime('%A, %d %B %Y'))

    text = text + "\n\n" + "If the ticket ends up going into '{}', please ignore that and do not change the ticket status - SDK/CLI team will handle it.".format(config.STATUS_SERVICE_TEAM_FAILURE_INVESTIGATION)

    text = text + "\n\n" + "([wiki|https://confluence.oci.oraclecorp.com/display/DEX/Requesting+a+preview+SDK+CLI#RequestingapreviewSDKCLI-RequestsreadytobeincludedinweeklypreviewSDK(ReadyforPreview)])"

    return text


DONE_TEMPLATE = """\
This ticket is done. All SDKs have been successfully generated and merged.

Next, start working on your samples and tests and get them into the preview branches.

See [Testing of feature support|https://confluence.oci.oraclecorp.com/display/DEX/Self-Service+Testing+and+Development] and [Code samples of feature support|https://confluence.oci.oraclecorp.com/display/DEX/Self-Service+Testing+and+Development].

Testing of generated CLI commands is the Service team's responsibility. For CLI testing, install the latest generated [preview build|https://confluence.oci.oraclecorp.com/pages/viewpage.action?spaceKey=DEX&title=CLI+FAQs#CLIFAQs-WherecanIfindthelatestversionofthepreviewCLI?]. Then to setup the development environment, [refer|https://confluence.oci.oraclecorp.com/display/DEX/OCI+CLI+Development+Setup#OCICLIDevelopmentSetup-1.Settingupthedevelopmentenvironment]. For running tests in your tenancy, please follow the [steps mentioned|https://confluence.oci.oraclecorp.com/display/DEX/OCI+CLI+Development+Setup#OCICLIDevelopmentSetup-Runningtestsagainstanothertenancy(RecommendedforServiceTeam)] Test all the commands/ param which are added, modified or deleted in this DEXREQ process. Please do end to end testing, which involves getting a response from service and validate. 
"""


def for_done(issue_key, issue, summary):
    warnings = process_all_spec_change_prs(issue)

    text = DONE_TEMPLATE

    if warnings:
        text = text + "\n\nWarnings:\n" + warnings

    return text


PARTIAL_DEX_SUPPORT_REQUIRED_FAILED_SDKS_TEMPLATE = """\

{sdk_sdks} that {requires_require} the on-call's attention:
{fields}"""


DEX_SUPPORT_REQUIRED_TEMPLATE = """\
The service team has asked the SDK/CLI team for help. The on-call engineer will respond to this ticket once a day.
{failed_sdk_text}
"""


def for_dex_support_required(issue_key, issue, summary):
    failed = [(language, value) for language, value in summary.sdks.statuses.items() if value == config.CUSTOM_STATUS_FAILURE]

    failed_sdk_text = ""
    if failed:
        fields = "\n".join(["- {}".format(language) for language, value in failed])
        failed_sdk_text = PARTIAL_DEX_SUPPORT_REQUIRED_FAILED_SDKS_TEMPLATE.format(
            fields=fields,
            sdk_sdks="SDK" if len(failed) == 1 else "SDKs",
            requires_require="requires" if len(failed) == 1 else "require")

    return DEX_SUPPORT_REQUIRED_TEMPLATE.format(
        reporter=summary.jira.reporter,
        failed_sdk_text=failed_sdk_text)


DEX_BULK_REVIEW_REQUIRED_TEMPLATE = """\
The bulk preview SDKs have been generated. The SDK/CLI team's on-call engineer will review the pull requests and merge them.{pr_text}

([wiki|https://confluence.oci.oraclecorp.com/display/DEX/Requesting+a+preview+SDK+CLI#RequestingapreviewSDKCLI-Step5:Verifygeneratedchanges])"""


def for_dex_bulk_review(issue_key, issue, summary):
    # TODO: get PR links
    pr_text = None
    if pr_text:
        pr_text = "\n\nPull requests to review and merge:\n\n" + pr_text
    else:
        pr_text = ""

    text = DEX_BULK_REVIEW_REQUIRED_TEMPLATE.format(
        reporter=summary.jira.reporter,
        ready_status=config.STATUS_READY_FOR_PREVIEW,
        requested_status=config.STATUS_PROCESSING_REQUESTED,
        pr_text=pr_text)

    return text


def additional_checks(issue_key, issue, summary):
    return ""


PREVIEW_HANDLERS = {
    PREVIEW_NON_ACTIONABLE_STATE: for_non_actionable,
    PREVIEW_PROCESSING_REQUESTED_STATE: for_processing_requested,
    PREVIEW_PROCESSING_STATE: for_processing,
    PREVIEW_SERVICE_TEAM_FAILURE_INVESTIGATION_STATE: for_service_team_failure_investigation,
    PREVIEW_DEX_SUPPORT_REQUIRED_STATE: for_dex_support_required,
    PREVIEW_SERVICE_TEAM_REVIEW_REQUIRED_STATE: for_service_team_review_required,
    PREVIEW_READY_FOR_PREVIEW_STATE: for_ready_for_bulk,
    PREVIEW_PROCESSING_BULK_STATE: for_processing_bulk,
    PREVIEW_DEX_BULK_REVIEW_STATE: for_dex_bulk_review,
    PREVIEW_DONE_STATE: for_done
}


def get_preview_state(issue, summary, sdk_statuses, all_sdks, any_sdks):
    if summary.jira.status == config.STATUS_PROCESSING_REQUESTED:
        state = PREVIEW_PROCESSING_REQUESTED_STATE

    elif summary.jira.status == config.STATUS_PROCESSING:
        state = PREVIEW_PROCESSING_STATE

    elif summary.jira.status == config.STATUS_SERVICE_TEAM_FAILURE_INVESTIGATION:
        state = PREVIEW_SERVICE_TEAM_FAILURE_INVESTIGATION_STATE

    elif summary.jira.status == config.STATUS_DEX_SUPPORT_REQUIRED:
        state = PREVIEW_DEX_SUPPORT_REQUIRED_STATE

    elif summary.jira.status == config.STATUS_SERVICE_TEAM_REVIEW_REQUIRED:
        state = PREVIEW_SERVICE_TEAM_REVIEW_REQUIRED_STATE

    elif summary.jira.status == config.STATUS_READY_FOR_PREVIEW:
        state = PREVIEW_READY_FOR_PREVIEW_STATE

    elif summary.jira.status == config.STATUS_PROCESSING_BULK:
        state = PREVIEW_PROCESSING_BULK_STATE

    elif summary.jira.status == config.STATUS_DEX_BULK_REVIEW:
        state = PREVIEW_DEX_BULK_REVIEW_STATE

    elif summary.jira.status == config.STATUS_DONE:
        state = PREVIEW_DONE_STATE

    else:
        state = PREVIEW_NON_ACTIONABLE_STATE

    return state


def advise_on_preview_issue(issue_key, issue, summary):
    text = ""
    if summary.state in PREVIEW_HANDLERS and PREVIEW_HANDLERS[summary.state]:
        handler = PREVIEW_HANDLERS[summary.state]
        text = handler(issue_key, issue, summary)

    text = text + additional_checks(issue_key, issue, summary)

    return text


TRANSITION_TO_BACKLOG_SINCE_SPEC_PR_DECLINED = """\
Automatically transitioned status back to '{status}'. The [spec diff pull request|{pr_url}] was {action}. Please revise your spec and generate a new spec artifact. Then update the spec version in your DEXREQ ticket and set the ticket status back to '{processing_requested_state}'.
"""

TRANSITION_TO_READY_FOR_PREVIEW_SINCE_SPEC_PR_MERGED = """\
Automatically transitioned status to '{status}'. The [spec diff pull request|{pr_url}] was {action}.
"""


def handle_transition_for_service_team_review_required_preconditions_check(issue_key, issue, summary):
    return True, ""


def handle_transition_for_service_team_review_required(issue_key, issue, summary):
    return handle_transition_for_spec_review_pr(issue_key, issue, summary, config.STATUS_READY_FOR_PREVIEW, TRANSITION_TO_READY_FOR_PREVIEW_SINCE_SPEC_PR_MERGED,
        TRANSITION_TO_BACKLOG_SINCE_SPEC_PR_DECLINED, handle_transition_for_service_team_review_required_preconditions_check)


PREVIEW_TRANSITION_HANDLERS = {
    PREVIEW_PROCESSING_STATE: handle_transition_for_processing,
    PREVIEW_SERVICE_TEAM_REVIEW_REQUIRED_STATE: handle_transition_for_service_team_review_required,
    PREVIEW_PROCESSING_BULK_STATE: handle_transition_for_processing
}


def handle_automated_preview_transitions(issue_key, issue, summary):
    return execute_appropriate_transition_handler(PREVIEW_TRANSITION_HANDLERS, issue_key, issue, summary)


#
# Show Preview Calendar
#
def show_preview_calendar(count=10):
    base_release_date, base_release_ga_number, release_overrides = \
        process_date_override(BULK_PREVIEW_DATE_OVERRIDE, PREVIEW_RELEASE_DATE_CADENCE_START_DATE, 140,
                              '--bulk-preview-date-overrides')

    # since preview releases are independent of GA date , base_release_ga_number is not used anywhere.
    now_datetime_utc = pytz.utc.localize(datetime.datetime.utcnow())
    now_datetime = now_datetime_utc.astimezone(PACIFIC_TIME_ZONE)
    now = now_datetime.date()
    release_date = base_release_date

    print("| SDK/CLI Preview Release Date |")
    print("| ---------------------------- |")

    while count > 0:
        release_date = find_next_matching_date(release_date, base_release_date,
                                               release_overrides, PREVIEW_RELEASE_DATE_CADENCE_IN_DAYS,
                                               "preview release")

        if release_date >= now:
            print("| {} |".format(release_date))
            count -= 1

        release_date += datetime.timedelta(1)
