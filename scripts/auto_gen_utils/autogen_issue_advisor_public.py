import datetime
import pytz
import six
import sys

import config
import util
import parse

import shared.bitbucket_utils

from autogen_issue_advisor_shared import PACIFIC_TIME_ZONE
from autogen_issue_advisor_shared import get_successful_pull_requests_text
from autogen_issue_advisor_shared import printv
from autogen_issue_advisor_shared import PR_LINK_TYPE
from autogen_issue_advisor_shared import get_failed_links_text
from autogen_issue_advisor_shared import DEFAULT_JIRA_ISSUE_FIELDS, CUSTOM_JIRA_ISSUE_FIELDS
from autogen_issue_advisor_shared import handle_transition_for_processing
from autogen_issue_advisor_shared import execute_appropriate_handler
from autogen_issue_advisor_shared import execute_appropriate_transition_handler
from autogen_issue_advisor_shared import handle_transition_for_spec_review_pr
from autogen_issue_advisor_shared import process_all_spec_change_prs
from autogen_issue_advisor_shared import get_spec_review_pr
from autogen_issue_advisor_shared import BitbucketBuilds
from autogen_issue_advisor_shared import BitbucketBuildChecks
from autogen_issue_advisor_shared import ServiceTeamMasterPrs
from autogen_issue_advisor_shared import find_next_matching_date, process_date_override

# Format: comma-separated, with each part being either "+YYYY-MM-DD", "-YYYY-MM-DD", or "=YYYY-MM-DD".
# Starting with "=" sets the base date of the one-week cadence. If not set, it's 2022-06-14.
# Example "=2018-10-11,+2018-10-24,-2018-10-25" means
# "do releases every week, starting with 2018-10-11, but don't do one on October 25th, do it on October 24th instead"
# Release dates are usually on Tuesday.
PUBLIC_RELEASE_DATE_OVERRIDE = None
PUBLIC_RELEASE_DATE_CADENCE_START_DATE = "2022-08-09"
PUBLIC_RELEASE_DATE_CADENCE_IN_DAYS = 7
PUBLIC_RELEASE_DATE_GA_NUMBER = 218  # GA 218 on 2022-08-09

# Format: comma-separated, with each part being either "+YYYY-MM-DD", "-YYYY-MM-DD", or "=YYYY-MM-DD".
# Starting with "=" sets the base date of the one-week cadence. If not set, it's 2022-06-07.
# Example "=2018-10-05,+2018-10-18,-2018-10-19" means
# "do RR cut-off every week, starting with 2018-10-05, but don't do one on October 19th, do it on October 18th instead"
# Cut-off dates are usually Monday (EOD), 8 days before the corresponding release date.
RELEASE_REQUESTED_CUT_OFF_DATE_OVERRIDE = None
RELEASE_REQUESTED_CUT_OFF_DATE_CADENCE_START_DATE = "2022-08-01"
RELEASE_REQUESTED_CUT_OFF_DATE_CADENCE_IN_DAYS = 7

# See https://confluence.oci.oraclecorp.com/plugins/gliffy/viewer.action?inline=false&pageId=60737311&attachmentId=64632788&name=DEXREQ%20Public%20Ticket%20Lifecycle&ceoid=64632759&key=~michross&lastPage=%2Fpages%2Fviewpage.action%3FpageId%3D60737311&imageUrl=%2Fdownload%2Fattachments%2F64632759%2FDEXREQ%2520Public%2520Ticket%2520Lifecycle.png%3Fversion%3D13%26modificationDate%3D1547155991054%26api%3Dv2&gonUrl=%2Fdownload%2Fattachments%2F64632759%2FDEXREQ%2520Public%2520Ticket%2520Lifecycle%3Fversion%3D13%26modificationDate%3D1547155990982%26api%3Dv2

PUBLIC_NON_ACTIONABLE_STATE = "PUBLIC_NON_ACTIONABLE_STATE"
PUBLIC_PROCESSING_REQUESTED_STATE = "PUBLIC_PROCESSING_REQUESTED_STATE"
PUBLIC_PROCESSING_STATE = "PUBLIC_PROCESSING_STATE"
PUBLIC_SERVICE_TEAM_FAILURE_INVESTIGATION_STATE = "PUBLIC_SERVICE_TEAM_FAILURE_INVESTIGATION_STATE"
PUBLIC_DEX_SUPPORT_REQUIRED_STATE = "PUBLIC_DEX_SUPPORT_REQUIRED_STATE"
PUBLIC_SERVICE_TEAM_WORK_REQUIRED_STATE = "PUBLIC_SERVICE_TEAM_WORK_REQUIRED_STATE"
PUBLIC_RELEASE_REQUESTED_STATE = "PUBLIC_RELEASE_REQUESTED_STATE"
PUBLIC_RELEASE_APPROVED_STATE = "PUBLIC_RELEASE_APPROVED_STATE"
PUBLIC_PROCESSING_BULK_STATE = "PUBLIC_PROCESSING_BULK_STATE"
PUBLIC_DEX_BULK_REVIEW_STATE = "PUBLIC_DEX_BULK_REVIEW_STATE"
PUBLIC_TO_DEPLOY_STATE = "PUBLIC_TO_DEPLOY_STATE"
PUBLIC_DONE_STATE = "PUBLIC_DONE_STATE"

JIRA_URL_TEMPLATE = "https://jira.oci.oraclecorp.com/browse/{ticket}"

PUBLIC_STATES = [
    PUBLIC_NON_ACTIONABLE_STATE,
    PUBLIC_PROCESSING_REQUESTED_STATE,
    PUBLIC_PROCESSING_STATE,
    PUBLIC_SERVICE_TEAM_FAILURE_INVESTIGATION_STATE,
    PUBLIC_DEX_SUPPORT_REQUIRED_STATE,
    PUBLIC_SERVICE_TEAM_WORK_REQUIRED_STATE,
    PUBLIC_RELEASE_REQUESTED_STATE,
    PUBLIC_RELEASE_APPROVED_STATE,
    PUBLIC_PROCESSING_BULK_STATE,
    PUBLIC_DEX_BULK_REVIEW_STATE,
    PUBLIC_TO_DEPLOY_STATE,
    PUBLIC_DONE_STATE,
]


NON_ACTIONABLE_BECAUSE_OF_BACKLOG_TEMPLATE = """\
[~{reporter}], the ticket status is set to '{backlog_status}'. For processing to begin, please set the ticket status to '{requested_status}'.
"""


def for_non_actionable(issue_key, issue, summary):
    text = is_requested_ga_date_valid_comment(summary)

    if summary.jira.status == config.STATUS_BACKLOG:

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

    # Check if provided GA date is a release date. Leave a comment if not
    text = is_requested_ga_date_valid_comment(summary)

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


def get_service_team_master_prs(issue, summary):
    prs_per_tool = summary.pull_requests.on_service_team.master.prs_per_tool

    master_prs = ServiceTeamMasterPrs(
        merged=[],
        approved=[],
        approved_could_be_bypassed=[],
        opened=[],
        opened_could_be_bypassed=[],
        missing=[],
        bypassed=[]
    )

    if util.is_cli_pr_required(issue):
        master_prs.missing = [config.CLI_NAME]

    # Bypassed tools are not missing
    for tool_name in master_prs.missing:
        if config.BYPASS_CHECK_PR_MASTER_PREFIX + tool_name in summary.checks.bypass:
            master_prs.bypassed.append(tool_name)

    for tool_name in master_prs.bypassed:
        if tool_name in master_prs.missing:
            master_prs.missing.remove(tool_name)

    for tool_name, pr_by_status in six.iteritems(prs_per_tool):
        for pr_and_url in pr_by_status.open:
            if config.BYPASS_CHECK_PR_MASTER_PREFIX + tool_name in summary.checks.bypass:
                master_prs.opened_could_be_bypassed.append(pr_and_url)
            else:
                master_prs.opened.append(pr_and_url)
            if tool_name in master_prs.missing:
                master_prs.missing.remove(tool_name)
        for pr_and_url in pr_by_status.approved:
            if config.BYPASS_CHECK_PR_MASTER_PREFIX + tool_name in summary.checks.bypass:
                master_prs.approved_could_be_bypassed.append(pr_and_url)
            else:
                master_prs.approved.append(pr_and_url)
            if tool_name in master_prs.missing:
                master_prs.missing.remove(tool_name)
        for pr_and_url in pr_by_status.merged:
            master_prs.merged.append(pr_and_url)
            if tool_name in master_prs.missing:
                master_prs.missing.remove(tool_name)

    # Bypassed tools are not opened
    for pr_and_url in master_prs.opened:
        if config.BYPASS_CHECK_PR_MASTER_PREFIX + pr_and_url.tool_name in summary.checks.bypass:
            master_prs.bypassed.append(tool_name)

    for tool_name in master_prs.bypassed:
        master_prs.opened = [pr_and_url for pr_and_url in master_prs.opened if pr_and_url.tool_name != tool_name]

    return master_prs


def get_bitbucket_builds(approved_master_prs):
    builds = BitbucketBuilds(
        failed_master_pr_builds={},
        successful_master_pr_builds={},
        in_progress_master_pr_builds={}
    )

    for tool in util.get_jira_reportable_tool_names():
        builds.failed_master_pr_builds[tool] = []
        builds.successful_master_pr_builds[tool] = []
        builds.in_progress_master_pr_builds[tool] = []

    for pr_and_url in approved_master_prs:
        printv("Checking for successful {} builds in {}".format(pr_and_url.tool_name, pr_and_url.url))
        pr = pr_and_url.pr
        build_status = shared.bitbucket_utils.get_bitbucket_build_status_for_pr(pr)
        build_status['pr'] = pr_and_url.url
        printv(build_status)
        for build in build_status['values']:
            build_state = build['state']
            printv("\t{} - {}".format(build_state, build['url']))
            if build_state == "SUCCESSFUL":
                builds.successful_master_pr_builds[pr_and_url.tool_name].append(build_status)
            elif build_state == "FAILED":
                builds.failed_master_pr_builds[pr_and_url.tool_name].append(build_status)
            elif build_state == "INPROGRESS":
                builds.in_progress_master_pr_builds[pr_and_url.tool_name].append(build_status)

    return builds


def check_bitbuckets_builds(summary, builds, master_prs):
    tools_to_check = []
    tools_to_check.extend(master_prs.missing)
    tools_to_check.extend([pr_and_url.tool_name for pr_and_url in master_prs.approved])
    tools_to_check.extend([pr_and_url.tool_name for pr_and_url in master_prs.opened])

    checks = BitbucketBuildChecks(
        all_required_prs_have_successful_builds=True,
        all_required_prs_have_no_failed_builds=True,
    )

    printv("Looking for builds for tools {}".format(tools_to_check))

    for tool_name in tools_to_check:
        if config.BYPASS_CHECK_PR_MASTER_PREFIX + tool_name in summary.checks.bypass:
            # Bypassed -- if we don't need a PR at all, we don't need passing builds
            continue
        if config.BYPASS_CHECK_PR_BUILD_MASTER_PREFIX + tool_name in summary.checks.bypass:
            # Bypassed
            continue
        if not builds.successful_master_pr_builds[tool_name]:
            # No successful builds
            printv("successful_master_pr_builds[{}] == false".format(tool_name))
            checks.all_required_prs_have_successful_builds = False
            continue

        if builds.failed_master_pr_builds[tool_name]:
            # Some builds failed
            checks.all_required_prs_have_no_failed_builds = False

    return checks


def zero_or_more_string(text):
    return text


zero_or_more_string.pattern = r".*"
PR_REQUEST_LINK_TEMPLATE = 'https://bitbucket.oci.oraclecorp.com/projects/SDK/repos/{repo}/compare?targetBranch=refs%2Fheads%2F{target_branch}&sourceBranch=refs%2Fheads%2F{generated_branch}&title={title}&description={description}&targetRepoId={target_repo_id:z}'


def get_cherry_pick_target_text(summary, tool_name):
    last_build = summary.builds.last[tool_name]
    if last_build:
        if last_build.successful and last_build.generation_successful and last_build.build_successful:
            if last_build.links:
                for l in last_build.links:
                    if l.link_type == PR_LINK_TYPE:
                        url = l.url

                        result = parse.parse(PR_REQUEST_LINK_TEMPLATE, url, {"z": zero_or_more_string})

                        if result and result["generated_branch"] and result["repo"]:
                            return " into this branch: [{generated_branch}|https://bitbucket.oci.oraclecorp.com/projects/SDK/repos/{repo}/browse?at=refs%2Fheads%2F{generated_branch}]".format(
                                generated_branch=result["generated_branch"],
                                repo=result["repo"])

    return " to master"


SERVICE_TEAM_WORK_REQUIRED_STATE_TEMPLATE = """\
The SDKs and the CLI have been generated. It is now the job of the service team to review the output, as well as cherry-pick the recordings for this feature from the preview branch of the CLI into the generated public CLI branch linked below.

[~{reporter}], please do the following:

1. If you had manual changes to the CLI code, [cherry-pick|https://confluence.oci.oraclecorp.com/pages/viewpage.action?spaceKey=DEX&title=How+to+cherry-pick+changes+from+preview+branch+to+master+branch] the CLI changes from [preview|https://bitbucket.oci.oraclecorp.com/projects/SDK/repos/python-cli/browse?at=refs%2Fheads%2Fpreview]{cli_branch_text}. Create a pull request from your CLI branch into the CLI master branch. You don't have to create a pull request if you didn't have manual changes. It is expected that the service team would have tested the generated commands in the preview generation and design review process.

2. Please update CLI ChangeLog Entry field. More info about CLI change log can be found here: [CLI ChangeLog|https://confluence.oci.oraclecorp.com/pages/viewpage.action?spaceKey=DEX&title=Directives+for+writing+CHANGELOG+for+OCI+CLI] 

3. {review_text}
"""


REVIEW_GENERATED_CODE_TEXT = """\
Check that the generated SDKs contain (1) your entire change, and (2) nothing unwanted. If you are satisfied, then set the '{ready_status}' status for this ticket. If you find that something is missing from your change or there is something that should not be there, then please revise your spec. Generate a new spec artifact, then update the spec version in this ticket and set the ticket status back to '{requested_status}'."""


REVIEW_SPEC_DIFF_TEXT = """\
Review the [spec change pull request|{pr_url}] to make sure it contains (1) your entire change, and (2) nothing unwanted. If you are satisfied, then approve the spec pull request and set the '{ready_status}' status for this ticket. If you find that something is missing from your change or there is something that should not be there, then please set the spec pull request to 'Needs Work' and revise your spec. Generate a new spec artifact, then update the spec version in this ticket and set the ticket status back to '{requested_status}'.

Spec change pull request to review: {pr_url}"""


def get_public_review_text(issue_key, reporter):
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
        return REVIEW_SPEC_DIFF_TEXT.format(ready_status=config.STATUS_RELEASE_REQUESTED,
            requested_status=config.STATUS_PROCESSING_REQUESTED,
            pr_url=dexreq_pr_url), dexreq_pr, dexreq_pr_url

    return REVIEW_GENERATED_CODE_TEXT.format(ready_status=config.STATUS_RELEASE_REQUESTED,
        requested_status=config.STATUS_PROCESSING_REQUESTED), None, None


SERVICE_TEAM_WORK_REQUIRED_STATE_DATE_TEMPLATE = """\
According to the '{field}' field of this ticket, you are targeting a public SDK/CLI release of {ga_date}. If that is not correct, please update the '{field}' field with the right release date. In order to release on that date, as per [Self-Service Testing and Development Calendar|https://confluence.oci.oraclecorp.com/display/DEX/Self-Service+Testing+and+Development#Self-ServiceTestingandDevelopment-Calendar], you need complete the steps above and have the ticket in '{requested_state}'' by end of day {release_cutoff_date} (Seattle time)."""

SERVICE_TEAM_WORK_REQUIRED_STATE_APPROVED_PRS_TEMPLATE = """\
The following pull {request_requests_sop} to master {is_are} currently open for this ticket, and since {it_they} {has_have} passed validation and {has_have} and at least one approval{optional_each}, {it_they} {is_are} ready for the SDK/CLI team's review once the ticket transitions to '{status}':

{links}"""

SERVICE_TEAM_WORK_REQUIRED_STATE_BUILD_PROBLEM_PRS_TEMPLATE = """\
The following pull {request_requests_sop} to master {is_are} currently open for this ticket, but either {doesnt_dont} have{a_opt} successful validation {build_builds_sop} yet, or {its_their} validation {build_builds_sop} failed:

{links}"""


SERVICE_TEAM_WORK_REQUIRED_STATE_OPEN_PRS_TEMPLATE = """\
The following pull {request_requests_sop} to master {is_are} currently open for this ticket, but {has_have} not yet received approval from a service team member:

{links}"""


SERVICE_TEAM_WORK_REQUIRED_STATE_OPEN_PRS_COULD_BE_BYPASSED_TEMPLATE = """\
The following pull {request_requests_sop} to master {is_are} currently open for this ticket, but {has_have} not yet received approval from a service team member. {cap_It_They} could be bypassed because of {its_their} bypass label, but since you opened {this_these} pull {request_requests_sop}, we will review {it_them} after one of your team members has approved {it_them}. If you don't want to have {this_these} pull {request_requests_sop} merged, please decline {it_them}.

{links}"""


SERVICE_TEAM_WORK_REQUIRED_STATE_APPROVED_PRS_COULD_BE_BYPASSED_TEMPLATE = """\
The following pull {request_requests_sop} to master {is_are} currently open for this ticket, and {has_have} received approval from a service team member. {cap_It_They} could be bypassed because of {its_their} bypass label, but since you opened {this_these} pull {request_requests_sop}, we will review {it_them} now. If you don't want to have {this_these} pull {request_requests_sop} merged, please decline {it_them}.

{links}"""


SERVICE_TEAM_WORK_REQUIRED_STATE_MISSING_PRS_TEMPLATE = """\
Still missing {is_an_are} open pull {request_requests_sop} to master for:

{missing}"""


SERVICE_TEAM_WORK_REQUIRED_STATE_BYPASSED_PRS_TEMPLATE = """\
The following {tool_tools} {does_do} not require an open pull request to master because of {its_their} bypass {label_labels}:

{bypassed}"""

CLI_DESIGN_REVIEW_REQUIRED_TEMPLATE = """\
The following CLI Design Review {ticket_tickets_sop} currently open for this ticket, please work with CLI Support to resolve the ticket.

{links}"""

CLI_CHANGELOG_REQUIRED_TEXT = """\
CLI ChangeLog Entry field is empty. Please update the ticket with a valid [CLI ChangeLog|https://confluence.oci.oraclecorp.com/pages/viewpage.action?spaceKey=DEX&title=Directives+for+writing+CHANGELOG+for+OCI+CLI].
"""


def parse_date(date):
    try:
        return datetime.datetime.strptime(date, '%Y-%m-%d').date()
    except ValueError:
        pass
    try:
        return datetime.datetime.strptime(date, '%m/%d/%Y').date()
    except ValueError:
        pass
    try:
        return datetime.datetime.strptime(date, '%d/%b/%Y').date()
    except ValueError:
        pass
    try:
        return datetime.datetime.strptime(date, '%b/%d/%Y').date()
    except ValueError:
        pass

    # 2-digit years
    try:
        return datetime.datetime.strptime(date, '%y-%m-%d').date()
    except ValueError:
        pass
    try:
        return datetime.datetime.strptime(date, '%m/%d/%y').date()
    except ValueError:
        pass
    try:
        return datetime.datetime.strptime(date, '%d/%b/%y').date()
    except ValueError:
        pass
    try:
        return datetime.datetime.strptime(date, '%b/%d/%y').date()
    except ValueError:
        pass

    raise ValueError('SDK/CLI GA Date "{}" was not in format "YYYY-MM-DD" (preferred), "MM/DD/YYYY", "DD/MON/YYY", or "MON/DD/YYYY"'.format(date))


FEATURE_API_IS_PUBLICLY_AVAILABLE_ACKNOWLEDGED = """\


You have acknowledged that the feature API is publicly available and not behind a whitelist, meaning once the SDKs and the CLI are publicly released, it is immediately ready to receive traffic from all customers.
"""

FEATURE_API_IS_PUBLICLY_AVAILABLE_NOT_ACKNOWLEDGED = """\


You have not acknowledged that the feature API is publicly available and not behind a whitelist, meaning once the SDKs and the CLI are publicly released, it is immediately ready to receive traffic from all customers. We do not release features that are not publicly available or that are behind whitelists.

[~{reporter}], this release cannot proceed until you set the value of the the '{field}' field to 'Yes'.
"""

FEATURE_API_IS_PUBLICLY_AVAILABLE_NOT_ACKNOWLEDGED_SKIPPED = """\


You have not acknowledged that the feature API is publicly available and not behind a whitelist by setting the '{field}' field to 'Yes', meaning once the SDKs and the CLI are publicly released, it is immediately ready to receive traffic from all customers. We do not release features that are not publicly available or that are behind whitelists.

We are skipping this check, though, because of the '{label}' label and will proceed with the release anyway once all other requirements have been met.
"""


def for_service_team_work_required(issue_key, issue, summary):
    review_text, dexreq_pr, dexreq_pr_url = get_public_review_text(issue_key, summary.jira.reporter)
    pr_text = get_successful_pull_requests_text(summary)
    if pr_text:
        if dexreq_pr_url:
            pr_text = "\nThe links below allow you to examine the effect of your change on the generated source. All you have to do is review the above [spec change pull request|{pr_url}]. The SDK/CLI team will do all the work to include the change in the weekly bulk public release.\n\n".format(pr_url=dexreq_pr_url) + pr_text
        else:
            pr_text = "\nThe links below allow you to examine the effects of your change on the generated source. Just look at the diff and verify that the generated SDKs contain the change you expect. The SDK/CLI team will do all the work to include the change in the weekly bulk public release.\n\n" + pr_text
    else:
        pr_text = "\n"

    cli_branch_text = get_cherry_pick_target_text(summary, config.CLI_NAME)

    # Check if provided GA date is a release date. Leave a comment if not
    text = is_requested_ga_date_valid_comment(summary)

    text += SERVICE_TEAM_WORK_REQUIRED_STATE_TEMPLATE.format(
        status=config.STATUS_RELEASE_REQUESTED,
        reporter=summary.jira.reporter,
        cli_branch_text=cli_branch_text,
        review_text=review_text)

    if pr_text:
        text = text + "\n\n" + pr_text

    if summary.jira.sdk_cli_ga_date:
        ga_date = parse_date(summary.jira.sdk_cli_ga_date)

        release_cutoff_date = get_previous_cut_off_date_before_public_release(issue_key, ga_date)

        text = text + "\n\n" + SERVICE_TEAM_WORK_REQUIRED_STATE_DATE_TEMPLATE.format(field=config.CUSTOM_FIELD_NAME_FOR_ID[config.CUSTOM_FIELD_ID_SDK_CLI_GA_DATE],
            ga_date=ga_date.strftime('%A, %d %B %Y'), release_cutoff_date=release_cutoff_date.strftime('%A, %d %B %Y'), requested_state=config.STATUS_RELEASE_REQUESTED)

    if summary.pull_requests.on_service_team.master.prs_per_tool:
        master_prs = get_service_team_master_prs(issue, summary)

        builds = get_bitbucket_builds(master_prs.approved)

        checks = check_bitbuckets_builds(summary, builds, master_prs)

        build_problem_pr = []
        build_problem_pr_bypassed = []
        build_problem_pr_unnecessary = []
        if not checks.all_required_prs_have_successful_builds or not checks.all_required_prs_have_no_failed_builds:
            printv("Checking which approved builds have build problems")
            for pr_and_url in master_prs.approved + master_prs.approved_could_be_bypassed:
                build_failed = False
                for b in builds.failed_master_pr_builds[pr_and_url.tool_name]:
                    if b['pr'] == pr_and_url.url:
                        build_failed = True
                        break

                build_succeeded = False
                for b in builds.successful_master_pr_builds[pr_and_url.tool_name]:
                    if b['pr'] == pr_and_url.url:
                        build_succeeded = True
                        break

                if build_failed or not build_succeeded:
                    if (config.BYPASS_CHECK_PR_BUILD_MASTER_PREFIX + pr_and_url.tool_name) in summary.checks.bypass:
                        build_problem_pr_bypassed.append(pr_and_url)
                    elif (config.BYPASS_CHECK_PR_MASTER_PREFIX + pr_and_url.tool_name) in summary.checks.bypass:
                        build_problem_pr_unnecessary.append(pr_and_url)
                    else:
                        build_problem_pr.append(pr_and_url)

        # Remove builds with a problem from the approved builds
        for pr_and_url in build_problem_pr:
            if pr_and_url in master_prs.approved:
                master_prs.approved.remove(pr_and_url)

        if master_prs.approved or build_problem_pr or master_prs.opened or master_prs.missing or master_prs.bypassed or master_prs.opened_could_be_bypassed or master_prs.approved_could_be_bypassed:
            text = text + "\n\n{panel:title=Master PR Status}"
            if master_prs.approved:
                text = text + "\n\n" + SERVICE_TEAM_WORK_REQUIRED_STATE_APPROVED_PRS_TEMPLATE.format(
                    status=config.STATUS_RELEASE_REQUESTED,
                    request_requests_sop="request" if len(master_prs.approved) == 1 else "requests",
                    is_are="is" if len(master_prs.approved) == 1 else "are",
                    has_have="has" if len(master_prs.approved) == 1 else "have",
                    it_they="it" if len(master_prs.approved) == 1 else "they",
                    optional_each="" if len(master_prs.approved) == 1 else " each",
                    links="\n".join(["- {}: {}{}".format(pr_and_url.tool_name, pr_and_url.url, " (failed, but the check was bypassed)" if pr_and_url in build_problem_pr_bypassed else "") for pr_and_url in master_prs.approved]),
                )

            if master_prs.approved_could_be_bypassed:
                text = text + "\n\n" + SERVICE_TEAM_WORK_REQUIRED_STATE_APPROVED_PRS_COULD_BE_BYPASSED_TEMPLATE.format(
                    request_requests_sop="request" if len(master_prs.approved_could_be_bypassed) == 1 else "requests",
                    is_are="is" if len(master_prs.approved_could_be_bypassed) == 1 else "are",
                    has_have="has" if len(master_prs.approved_could_be_bypassed) == 1 else "have",
                    cap_It_They="It" if len(master_prs.approved_could_be_bypassed) == 1 else "They",
                    this_these="this" if len(master_prs.approved_could_be_bypassed) == 1 else "these",
                    it_them="it" if len(master_prs.approved_could_be_bypassed) == 1 else "them",
                    its_their="its" if len(master_prs.approved_could_be_bypassed) == 1 else "their",
                    links="\n".join(["- {}: {}{}{}".format(pr_and_url.tool_name, pr_and_url.url,
                                    " (failed, but the check was bypassed)" if pr_and_url in build_problem_pr_bypassed else "",
                                    " (build failed)" if pr_and_url in build_problem_pr + build_problem_pr_unnecessary else "") for pr_and_url in master_prs.approved_could_be_bypassed]),
                )

            if build_problem_pr:
                text = text + "\n\n" + SERVICE_TEAM_WORK_REQUIRED_STATE_BUILD_PROBLEM_PRS_TEMPLATE.format(
                    request_requests_sop="request" if len(build_problem_pr) == 1 else "requests",
                    is_are="is" if len(build_problem_pr) == 1 else "are",
                    doesnt_dont="doesn't" if len(build_problem_pr) == 1 else "don't",
                    a_opt=" a" if len(build_problem_pr) == 1 else "",
                    build_builds_sop="build" if len(build_problem_pr) == 1 else "builds",
                    its_their="its" if len(build_problem_pr) == 1 else "their",
                    links="\n".join(["- {}: {}".format(pr_and_url.tool_name, pr_and_url.url) for pr_and_url in build_problem_pr]),
                )

            if master_prs.opened:
                text = text + "\n\n" + SERVICE_TEAM_WORK_REQUIRED_STATE_OPEN_PRS_TEMPLATE.format(
                    request_requests_sop="request" if len(master_prs.opened) == 1 else "requests",
                    is_are="is" if len(master_prs.opened) == 1 else "are",
                    has_have="has" if len(master_prs.opened) == 1 else "have",
                    links="\n".join(["- {}: {}".format(pr_and_url.tool_name, pr_and_url.url) for pr_and_url in master_prs.opened]),
                )

            if master_prs.opened_could_be_bypassed:
                text = text + "\n\n" + SERVICE_TEAM_WORK_REQUIRED_STATE_OPEN_PRS_COULD_BE_BYPASSED_TEMPLATE.format(
                    request_requests_sop="request" if len(master_prs.opened_could_be_bypassed) == 1 else "requests",
                    is_are="is" if len(master_prs.opened_could_be_bypassed) == 1 else "are",
                    has_have="has" if len(master_prs.opened_could_be_bypassed) == 1 else "have",
                    cap_It_They="It" if len(master_prs.opened_could_be_bypassed) == 1 else "They",
                    this_these="this" if len(master_prs.opened_could_be_bypassed) == 1 else "these",
                    it_them="it" if len(master_prs.opened_could_be_bypassed) == 1 else "them",
                    its_their="its" if len(master_prs.opened_could_be_bypassed) == 1 else "their",
                    links="\n".join(["- {}: {}".format(pr_and_url.tool_name, pr_and_url.url) for pr_and_url in master_prs.opened_could_be_bypassed]),
                )

            if master_prs.missing:
                text = text + "\n\n" + SERVICE_TEAM_WORK_REQUIRED_STATE_MISSING_PRS_TEMPLATE.format(
                    request_requests_sop="request" if len(master_prs.missing) == 1 else "requests",
                    is_an_are="is an" if len(master_prs.missing) == 1 else "are",
                    missing="\n".join(["- {}".format(tool_name) for tool_name in master_prs.missing])
                )

            if master_prs.bypassed:
                text = text + "\n\n" + SERVICE_TEAM_WORK_REQUIRED_STATE_BYPASSED_PRS_TEMPLATE.format(
                    tool_tools="tool" if len(master_prs.bypassed) == 1 else "tools",
                    does_do="does" if len(master_prs.bypassed) == 1 else "do",
                    its_their="its" if len(master_prs.bypassed) == 1 else "their",
                    label_labels="label" if len(master_prs.bypassed) == 1 else "labels",
                    bypassed="\n".join(["- {} (label {}{})".format(tool_name, config.BYPASS_CHECK_PR_MASTER_PREFIX, tool_name) for tool_name in master_prs.bypassed])
                )
            text = text + "{panel}"

    if summary.jira.feature_api_is_publicly_available_and_unwhitelisted_in_prod == "Yes":
        text = text + FEATURE_API_IS_PUBLICLY_AVAILABLE_ACKNOWLEDGED
    else:
        if config.BYPASS_CHECK_API_PUBLICLY_AVAILABLE in summary.checks.bypass:
            text = text + FEATURE_API_IS_PUBLICLY_AVAILABLE_NOT_ACKNOWLEDGED_SKIPPED.format(
                label=config.BYPASS_CHECK_API_PUBLICLY_AVAILABLE,
                field=config.CUSTOM_FIELD_NAME_FOR_ID[config.CUSTOM_FIELD_ID_FEATURE_API_IS_PUBLICLY_AVAILABLE])
        else:
            text = text + FEATURE_API_IS_PUBLICLY_AVAILABLE_NOT_ACKNOWLEDGED.format(reporter=summary.jira.reporter,
                    field=config.CUSTOM_FIELD_NAME_FOR_ID[config.CUSTOM_FIELD_ID_FEATURE_API_IS_PUBLICLY_AVAILABLE])

    if summary.cli.pending_design_reviews:
        if config.BYPASS_CLI_DESIGN_REVIEW_CHECK not in summary.checks.bypass:
            text = text + "\n\n" + CLI_DESIGN_REVIEW_REQUIRED_TEMPLATE.format(
                ticket_tickets_sop="ticket is" if len(summary.cli.pending_design_reviews) == 1 else "tickets are",
                links="\n".join(JIRA_URL_TEMPLATE.format(ticket=key) for key in summary.cli.pending_design_reviews)
            )
        else:
            printv("Bypassing CLI Design review check for: " + issue.key)
    else:
        printv("No pending CLI Design review for: " + issue.key)

    return text


PROCESSING_TEMPLATE = """\
The ticket is being processed. Please wait for the automation to produce results.
"""

PARTIALLY_PROCESSED_TEMPLATE = """\
The ticket is being processed, and some SDKs have already been generated.

Already done:
{fields}"""


def for_processing(issue_key, issue, summary):

    # Check if provided GA date is a release date. Leave a comment if not
    text = is_requested_ga_date_valid_comment(summary)

    fields = "\n".join(["- {}".format(language) for language, value in summary.sdks.statuses.items()
        if value in [config.CUSTOM_STATUS_FAILURE, config.CUSTOM_STATUS_SUCCESS, config.CUSTOM_STATUS_DONE]])
    if fields:
        return text + PARTIALLY_PROCESSED_TEMPLATE.format(fields=fields, reporter=summary.jira.reporter)
    else:
        return text + PROCESSING_TEMPLATE.format(reporter=summary.jira.reporter)


PARTIALLY_MERGED_TEMPLATE = """\
The ticket is being processed, and some SDKs have already been merged.

Already merged:
{fields}"""


def for_processing_requested(issue_key, issue, summary):

    # Move ticket to Service Team Work Required and leave a comment if the provided GA date is incorrect. 
    comment = is_requested_ga_date_valid_comment(summary)
    if len(comment) > 0:
        util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_SERVICE_TEAM_WORK_REQUIRED)
        return comment

    return "This ticket is ready for automated processing. Please wait for the automation to start."


def get_release_requested_status_date(issue_key, summary):
    # Find timestamp when status "Release Requested" was added
    release_requested_datetime_pacific = None
    if summary.jira and summary.jira.changelog:
        for cl in summary.jira.changelog:
            for item in cl.changed_items:
                if item.field != "status":
                    continue

                # There was no old status, or "Release Requested" wasn't it
                was_not_signed_off = not item.old or config.STATUS_RELEASE_REQUESTED != item.old
                # AND
                # There is a new status, and "Release Requested" is it
                is_now_signed_off = item.new and config.STATUS_RELEASE_REQUESTED == item.new
                if was_not_signed_off and is_now_signed_off:
                    # config.STATUS_RELEASE_REQUESTED was added
                    release_requested_datetime = pytz.utc.localize(datetime.datetime.strptime(cl.created, '%Y-%m-%dT%H:%M:%S.%f+0000'))

                    # change this to Pacific time
                    release_requested_datetime_pacific = release_requested_datetime.astimezone(PACIFIC_TIME_ZONE)

    release_requested_date = None
    if release_requested_datetime_pacific:
        release_requested_date = release_requested_datetime_pacific.date()

    printv("get_release_requested_status_date returns {}, based on Seattle time {}".format(release_requested_date, release_requested_datetime_pacific))

    return release_requested_date


def find_previous_matching_date(start_date, base_date, overrides, cadence_in_days, date_name):
    date = start_date
    while True:
        printv("{}: {}".format(date_name, date))
        date_string = date.isoformat()
        if (date - base_date).days % cadence_in_days == 0:
            if date_string in overrides and not overrides[date_string]:
                # NOT on this day
                printv("Elected to not have regular {} scheduled on: {}".format(date_name, date_string))
                date -= datetime.timedelta(1)
                continue
            else:
                # either no override set for this date, or it's a positive override
                printv("Having regular {} scheduled on: {}".format(date_name, date_string))
                break
        if date_string in overrides and overrides[date_string]:
            # cut-off on this day
            printv("Elected to have special {} scheduled on: {}".format(date_name, date_string))
            break

        date -= datetime.timedelta(1)

    return date


def get_next_release_requested_cut_off_date(release_requested_date):
    base_cut_off_date, ga_number, cut_off_overrides = process_date_override(RELEASE_REQUESTED_CUT_OFF_DATE_OVERRIDE,
        RELEASE_REQUESTED_CUT_OFF_DATE_CADENCE_START_DATE, None, '--public-to-deploy-cut-off-date-overrides')

    # Find the Wednesday after this date (if added on a Wednesday, that's ok).
    # (We're requiring them to be in 'Release Requested' by end of Wednesday.)
    release_requested_cut_off_date = find_next_matching_date(release_requested_date,
        base_cut_off_date, cut_off_overrides, RELEASE_REQUESTED_CUT_OFF_DATE_CADENCE_IN_DAYS,
        "'{}' cut-off".format(config.STATUS_RELEASE_REQUESTED))

    return release_requested_cut_off_date


def get_next_public_release_date_after_cut_off(release_requested_cut_off_date):
    # Now release_requested_cut_off_date is the next 'Release Requested' cut-off date
    # Add two days, because we'll never have the release day be the same as the
    # cut-off day. And our cut-off is EOD Monday, but the matching release day is not the
    # next Tuesday, but the Tuesday after that.
    release_date = release_requested_cut_off_date + datetime.timedelta(2)

    base_release_date, base_release_ga_number, release_overrides = process_date_override(PUBLIC_RELEASE_DATE_OVERRIDE,
        PUBLIC_RELEASE_DATE_CADENCE_START_DATE, PUBLIC_RELEASE_DATE_GA_NUMBER, '--public-release-date-overrides')

    # Find the next release day after that cut-off.
    release_date = find_next_matching_date(release_date,
        base_release_date, release_overrides, PUBLIC_RELEASE_DATE_CADENCE_IN_DAYS,
        "public release")

    print("Next public release after {} is scheduled for: {}".format(release_requested_cut_off_date, release_date))

    return release_date


COMMENT_TEMPLATE_INCORRECT_DATES = """\

[~{reporter}] The provided SDK/CLI GA Date of {requested_date} is not a release date. The next closest release date would be {next_release_date}. Please update this field accordingly."""


# Determines the next release date after the given date, inclusive.
# If the requested date is a release date, it should return it
def get_next_public_release_date_after_ga(requested_ga_date):
    base_release_date, base_release_ga_number, release_overrides = process_date_override(PUBLIC_RELEASE_DATE_OVERRIDE,
    PUBLIC_RELEASE_DATE_CADENCE_START_DATE, PUBLIC_RELEASE_DATE_GA_NUMBER, '--public-release-date-overrides')

    release_date = find_next_matching_date(requested_ga_date,
        base_release_date, release_overrides, PUBLIC_RELEASE_DATE_CADENCE_IN_DAYS,
        "Release Date")

    return release_date


def is_requested_ga_date_valid_comment(summary):
    comment = ""
    if summary.jira.sdk_cli_ga_date:
        requested_ga_date = parse_date(summary.jira.sdk_cli_ga_date)
        release_after_req_ga = get_next_public_release_date_after_ga(requested_ga_date)

        if release_after_req_ga != requested_ga_date:
            comment = COMMENT_TEMPLATE_INCORRECT_DATES.format(reporter=summary.jira.reporter, requested_date=requested_ga_date, next_release_date=release_after_req_ga) + "\n\n"

    if config.DISABLE_COMMENT_INCORRECT_DATES:
        printv("Disabled commenting on incorrect GA dates. Not Leaving comment: \n" + comment)
        return ""

    return comment


def get_next_public_release_date_after_cut_off_for_issue(issue_key, summary):
    # Find timestamp when "Release Requested" was set
    release_requested_date = get_release_requested_status_date(issue_key, summary)
    print("Status '{}' was set on issue '{}' on: {}".format(config.STATUS_RELEASE_REQUESTED, issue_key, release_requested_date))

    if not release_requested_date:
        # Couldn't figure out when status was set, decline to give a date
        return None

    # Find the Monday after this date (if added on a Monday, that's ok).
    # (We're requiring them to be in 'Release Requested' by end of Monday.)
    release_requested_cut_off_date = get_next_release_requested_cut_off_date(release_requested_date)

    return get_next_public_release_date_after_cut_off(release_requested_cut_off_date)


def get_next_public_release_date():
    base_release_date, base_release_ga_number, release_overrides = process_date_override(PUBLIC_RELEASE_DATE_OVERRIDE,
        PUBLIC_RELEASE_DATE_CADENCE_START_DATE, PUBLIC_RELEASE_DATE_GA_NUMBER, '--public-release-date-overrides')

    now_datetime_utc = pytz.utc.localize(datetime.datetime.utcnow())
    now_datetime = now_datetime_utc.astimezone(PACIFIC_TIME_ZONE)
    now = now_datetime.date()

    release_date = find_next_matching_date(now,
        base_release_date, release_overrides, PUBLIC_RELEASE_DATE_CADENCE_IN_DAYS,
        "Release Date")

    return release_date


def get_previous_cut_off_date_before_public_release(issue_key, ga_date):
    print("SDK/CLI GA Date field for issue '{}' is set to: '{}'".format(issue_key, ga_date))

    base_release_date, base_release_ga_number, release_overrides = process_date_override(PUBLIC_RELEASE_DATE_OVERRIDE,
        PUBLIC_RELEASE_DATE_CADENCE_START_DATE, PUBLIC_RELEASE_DATE_GA_NUMBER, '--public-release-date-overrides')
    base_cut_off_date, ga_number, cut_off_overrides = process_date_override(RELEASE_REQUESTED_CUT_OFF_DATE_OVERRIDE,
        RELEASE_REQUESTED_CUT_OFF_DATE_CADENCE_START_DATE, None, '--public-release-requested-cut-off-date-overrides')

    release_date = base_release_date
    release_requested_cut_off_date = base_cut_off_date
    ga_number = base_release_ga_number

    while release_date < ga_date:
        release_date = find_next_matching_date(release_date,
            base_release_date, release_overrides, PUBLIC_RELEASE_DATE_CADENCE_IN_DAYS,
            "public release")
        release_requested_cut_off_date = find_next_matching_date(release_requested_cut_off_date,
            base_cut_off_date, cut_off_overrides, RELEASE_REQUESTED_CUT_OFF_DATE_CADENCE_IN_DAYS,
            "public release requested cut-off")

        if release_date >= ga_date:
            break

        release_date += datetime.timedelta(1)
        release_requested_cut_off_date += datetime.timedelta(1)
        ga_number += 1

    release_requested_cut_off_date += datetime.timedelta(-1)

    print("{} cut-off day before public release on '{}' is scheduled for: '{}'".format(config.STATUS_RELEASE_REQUESTED, ga_date, release_requested_cut_off_date))

    return release_requested_cut_off_date


def get_next_release_requested_cut_off_date_for_ga_date(release_requested_date, requested_ga_date):
    while True:
        printv("Find cut-off date for {}".format(release_requested_date))
        cut_off_date = get_next_release_requested_cut_off_date(release_requested_date)
        printv("Cut-off date for {} is {}".format(release_requested_date, cut_off_date))
        printv("GA date for cut-off date {}".format(cut_off_date))
        ga_date = get_next_public_release_date_after_cut_off(cut_off_date)
        printv("GA date for cut-off date {} is {}".format(cut_off_date, ga_date))

        if ga_date >= requested_ga_date:
            break

        release_requested_date = cut_off_date + datetime.timedelta(1)

    printv("Next cut-off date for desired GA {} is {} (actual GA {})".format(requested_ga_date, cut_off_date, ga_date))

    return cut_off_date


RELEASE_REQUESTED_OPEN_PRS_TEMPLATE = """\
The following pull {request_requests_sop} to master {is_are} currently open for this ticket and should now be reviewed by the SDK/CLI team:

{links}"""

RELEASE_REQUESTED_MERGED_PRS_TEMPLATE = """\
The following pull {request_requests_sop} to master {has_have} already been merged:

{links}"""

RELEASE_REQUESTED_BYPASSED_PRS_TEMPLATE = """\
The following {tool_tools} {does_do} not require a merged pull request to master because of {its_their} bypass {label_labels}:

{links}"""

RELEASE_REQUESTED_TEMPLATE = """\
The service team has indicated that the change is ready for release. The SDK/CLI team will review any open pull requests.

The next step is the UDX Go/No-Go meeting, after which the SDK/CLI team will merge pull requests and initiate the bulk public release."""


RELEASE_REQUESTED_TEMPLATE_DATES_TEMPLATE = """\


The public release that the changes will be included in is scheduled for {release_date}. The '{status}' cut-off for this release date is 11:59 PM {cut_off_date} (Seattle time). The '{status}' status was set on {status_date} (Seattle time). This ticket meets the '{status}' cut-off. If corresponding UDX and ORM tickets are approved, then the ticket will be released as part of the release on {release_date}."""


RELEASE_REQUESTED_TEMPLATE_MISSED_CUT_OFF_WARNING_TEMPLATE = """\


*[~{reporter}], you have missed the '{status}' cut-off for releasing on {requested_release_date}*, which was the date you requested using the '{field}' field in this ticket. The '{status}' cut-off date for this ticket was 11:59 PM {release_cutoff_date} (Seattle time), however the '{status}' status was set on {status_date}. Hence, the changes associated with this ticket cannot be released as requested."""


RELEASE_REQUESTED_TEMPLATE_MISSED_CUT_OFF_WARNING_EXCEPTION_PROCESS_TEMPLATE = """\


If you believe you have reasons for an exception to the cut-off date, in the hopes of still releasing on {requested_release_date}, please urgently look at the [Exception Process for the SDK/CLI|https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=132774928#SDK/CLISelf-ServiceFrequentlyAskedQuestions-WhatistheexceptionprocessifIwanttoskipsomepartofSDK/CLIself-service,orGAwithoutSDK/CLI?]."""


def for_release_requested(issue_key, issue, summary):
    text = RELEASE_REQUESTED_TEMPLATE

    next_ga_date = get_next_public_release_date_after_cut_off_for_issue(issue_key, summary)

    if summary.jira.sdk_cli_ga_date:
        requested_ga_date = parse_date(summary.jira.sdk_cli_ga_date)
        release_cutoff_date = get_previous_cut_off_date_before_public_release(issue_key, requested_ga_date)
    else:
        # Just assume they want the next GA.
        requested_ga_date = next_ga_date

    ga_date_comment = is_requested_ga_date_valid_comment(summary)

    # Only leave these comments if requested date is a release date. Otherwise comments will not make sense.
    # If the requested date is not a release date, a different comment has already been left. 
    if len(ga_date_comment) > 0:
        text = ga_date_comment
    elif next_ga_date: 
        release_requested_date = get_release_requested_status_date(issue_key, summary)
        release_requested_cut_off_date = get_next_release_requested_cut_off_date_for_ga_date(release_requested_date, requested_ga_date)

        if next_ga_date > requested_ga_date:
            text = text + RELEASE_REQUESTED_TEMPLATE_MISSED_CUT_OFF_WARNING_TEMPLATE.format(
                reporter=summary.jira.reporter,
                field=config.CUSTOM_FIELD_NAME_FOR_ID[config.CUSTOM_FIELD_ID_SDK_CLI_GA_DATE],
                status=config.STATUS_RELEASE_REQUESTED,
                requested_release_date=requested_ga_date.strftime('%A, %d %B %Y'),
                release_cutoff_date=release_cutoff_date.strftime('%A, %d %B %Y'),
                status_date=release_requested_date.strftime('%A, %d %B %Y'))

        text = text + RELEASE_REQUESTED_TEMPLATE_DATES_TEMPLATE.format(
            status=config.STATUS_RELEASE_REQUESTED,
            status_date=release_requested_date.strftime('%A, %d %B %Y'),
            cut_off_date=release_requested_cut_off_date.strftime('%A, %d %B %Y'),
            release_date=(next_ga_date if next_ga_date > requested_ga_date else requested_ga_date).strftime('%A, %d %B %Y'))

        if next_ga_date > requested_ga_date:
            text = text + RELEASE_REQUESTED_TEMPLATE_MISSED_CUT_OFF_WARNING_EXCEPTION_PROCESS_TEMPLATE.format(
                requested_release_date=requested_ga_date.strftime('%A, %d %B %Y'))

    master_prs = get_service_team_master_prs(issue, summary)

    # treat all opened PRs the same
    master_prs.opened.extend(master_prs.opened_could_be_bypassed)
    master_prs.opened.extend(master_prs.approved)

    if master_prs.merged or master_prs.opened or master_prs.bypassed:
        text = text + "\n\n{panel:title=Master PR Status}"

        if master_prs.opened:
            text = text + "\n\n" + RELEASE_REQUESTED_OPEN_PRS_TEMPLATE.format(
                request_requests_sop="request" if len(master_prs.opened) == 1 else "requests",
                is_are="is" if len(master_prs.opened) == 1 else "are",
                links="\n".join(["- {}: {}".format(pr_and_url.tool_name, pr_and_url.url) for pr_and_url in master_prs.opened]),
            )

        if master_prs.merged:
            text = text + "\n\n" + RELEASE_REQUESTED_MERGED_PRS_TEMPLATE.format(
                request_requests_sop="request" if len(master_prs.merged) == 1 else "requests",
                has_have="has" if len(master_prs.merged) == 1 else "have",
                links="\n".join(["- {}: {}".format(pr_and_url.tool_name, pr_and_url.url) for pr_and_url in master_prs.merged]),
            )

        if master_prs.bypassed:
            text = text + "\n\n" + RELEASE_REQUESTED_BYPASSED_PRS_TEMPLATE.format(
                tool_tools="tool" if len(master_prs.bypassed) == 1 else "tools",
                does_do="does" if len(master_prs.bypassed) == 1 else "do",
                its_their="its" if len(master_prs.bypassed) == 1 else "their",
                label_labels="label" if len(master_prs.bypassed) == 1 else "labels",
                links="\n".join(["- {} (label {}{})".format(tool_name, config.BYPASS_CHECK_PR_MASTER_PREFIX, tool_name) for tool_name in master_prs.bypassed])
            )
        text = text + "{panel}"

    return text


def for_release_approved(issue_key, issue, summary):
    return "This ticket has been approved for release and is ready for the next bulk public SDK generation." + "\n\n" + "If the ticket ends up going into '{}', please ignore that and do not change the ticket status - SDK/CLI team will handle it.".format(config.STATUS_SERVICE_TEAM_FAILURE_INVESTIGATION)


PROCESSING_BULK_TEMPLATE = """\
The ticket is being processed for the bulk public build. Please wait for the automation to produce results.

([wiki|https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=43683000#RequestingapublicSDK/CLI-Step4:Notifications])"""

PARTIALLY_PROCESSED_BULK_TEMPLATE = """\
The ticket is being processed for the bulk public build, and some SDKs have already been generated.

Already done:
{fields}

([wiki|https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=43683000#RequestingapublicSDK/CLI-Step4:Notifications])"""


def for_processing_bulk(issue_key, issue, summary):
    fields = "\n".join(["- {}".format(language) for language, value in summary.sdks.statuses.items()
                        if value in [config.CUSTOM_STATUS_FAILURE, config.CUSTOM_STATUS_SUCCESS,
                                     config.CUSTOM_STATUS_DONE]])
    if fields:
        return PARTIALLY_PROCESSED_BULK_TEMPLATE.format(fields=fields)
    else:
        return PROCESSING_BULK_TEMPLATE


def for_dex_bulk_review(issue_key, issue, summary):
    return """\
The bulk public SDKs have been generated. The SDK/CLI team's on-call engineer will review the pull requests and merge them.

([wiki|https://confluence.oci.oraclecorp.com/display/DEX/Requesting+a+preview+SDK+CLI#RequestingapreviewSDKCLI-Step5:Verifygeneratedchanges])"""


TO_DEPLOY_TEMPLATE = """\
This feature has been successfully generated and merged for all SDKs/CLI. It is now pending public release."""

TO_DEPLOY_DATE_TEMPLATE = """\


The change will be publicly released on {release_date}."""


def for_to_deploy(issue_key, issue, summary):
    text = TO_DEPLOY_TEMPLATE

    date = get_next_public_release_date()
    if date:
        text = text + TO_DEPLOY_DATE_TEMPLATE.format(
            release_date=date.strftime('%A, %d %B %Y'))

    return text


def for_done(issue_key, issue, summary):
    warnings = process_all_spec_change_prs(issue)

    text = "This ticket is done. All SDKs have been successfully generated, merged, and publicly released."

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

    # Check if provided GA date is a release date. Leave a comment if not
    text = is_requested_ga_date_valid_comment(summary)

    failed = [(language, value) for language, value in summary.sdks.statuses.items() if value == config.CUSTOM_STATUS_FAILURE]

    failed_sdk_text = ""
    if failed:
        fields = "\n".join(["- {}".format(language) for language, value in failed])
        failed_sdk_text = PARTIAL_DEX_SUPPORT_REQUIRED_FAILED_SDKS_TEMPLATE.format(
            fields=fields,
            sdk_sdks="SDK" if len(failed) == 1 else "SDKs",
            requires_require="requires" if len(failed) == 1 else "require")

    return text + DEX_SUPPORT_REQUIRED_TEMPLATE.format(
        reporter=summary.jira.reporter,
        failed_sdk_text=failed_sdk_text)


def additional_checks(issue_key, issue, summary):
    text = ""

    return text


PUBLIC_HANDLERS = {
    PUBLIC_NON_ACTIONABLE_STATE: for_non_actionable,
    PUBLIC_PROCESSING_REQUESTED_STATE: for_processing_requested,
    PUBLIC_PROCESSING_STATE: for_processing,
    PUBLIC_SERVICE_TEAM_FAILURE_INVESTIGATION_STATE: for_service_team_failure_investigation,
    PUBLIC_DEX_SUPPORT_REQUIRED_STATE: for_dex_support_required,
    PUBLIC_SERVICE_TEAM_WORK_REQUIRED_STATE: for_service_team_work_required,
    PUBLIC_RELEASE_REQUESTED_STATE: for_release_requested,
    PUBLIC_RELEASE_APPROVED_STATE: for_release_approved,
    PUBLIC_PROCESSING_BULK_STATE: for_processing_bulk,
    PUBLIC_DEX_BULK_REVIEW_STATE: for_dex_bulk_review,
    PUBLIC_TO_DEPLOY_STATE: for_to_deploy,
    PUBLIC_DONE_STATE: for_done
}


def get_public_state(issue, summary, sdk_statuses, all_sdks, any_sdks):
    if summary.jira.status == config.STATUS_PROCESSING_REQUESTED:
        state = PUBLIC_PROCESSING_REQUESTED_STATE

    elif summary.jira.status == config.STATUS_PROCESSING:
        state = PUBLIC_PROCESSING_STATE

    elif summary.jira.status == config.STATUS_SERVICE_TEAM_FAILURE_INVESTIGATION:
        state = PUBLIC_SERVICE_TEAM_FAILURE_INVESTIGATION_STATE

    elif summary.jira.status == config.STATUS_DEX_SUPPORT_REQUIRED:
        state = PUBLIC_DEX_SUPPORT_REQUIRED_STATE

    elif summary.jira.status == config.STATUS_SERVICE_TEAM_WORK_REQUIRED:
        state = PUBLIC_SERVICE_TEAM_WORK_REQUIRED_STATE

    elif summary.jira.status == config.STATUS_RELEASE_REQUESTED:
        state = PUBLIC_RELEASE_REQUESTED_STATE

    elif summary.jira.status == config.STATUS_RELEASE_APPROVED:
        state = PUBLIC_RELEASE_APPROVED_STATE

    elif summary.jira.status == config.STATUS_PROCESSING_BULK:
        state = PUBLIC_PROCESSING_BULK_STATE

    elif summary.jira.status == config.STATUS_DEX_BULK_REVIEW:
        state = PUBLIC_DEX_BULK_REVIEW_STATE

    elif summary.jira.status == config.STATUS_TO_DEPLOY:
        state = PUBLIC_TO_DEPLOY_STATE

    elif summary.jira.status == config.STATUS_DONE:
        state = PUBLIC_DONE_STATE

    else:
        state = PUBLIC_NON_ACTIONABLE_STATE

    return state


def advise_on_public_issue(issue_key, issue, summary):
    text = execute_appropriate_handler(PUBLIC_HANDLERS, issue_key, issue, summary)

    text = text + additional_checks(issue_key, issue, summary)

    return text


#
# Automated transitions
#


TRANSITION_TO_SERVICE_TEAM_WORK_REQUIRED_TEMPLATE = """\
Automatically transitioning status back to '{status}'. The service team's work is not complete yet.{problem_text}
"""


def ready_for_release_requested_status(issue_key, issue, summary):
    printv("Checking for unauthorized transition from '{}' to '{}'...".format(config.STATUS_SERVICE_TEAM_WORK_REQUIRED, config.STATUS_RELEASE_REQUESTED))

    master_prs = get_service_team_master_prs(issue, summary)
    printv("master_prs: {}".format(master_prs))

    builds = get_bitbucket_builds(master_prs.approved)
    printv("builds: {}".format(builds))

    checks = check_bitbuckets_builds(summary, builds, master_prs)
    printv("checks: {}".format(checks))

    # No missing master PRs
    # No open PRs that aren't approved
    # No open PRs that could be bypassed (if they can be bypassed, why are they open?)
    is_master_ok = not master_prs.missing and not master_prs.opened and not master_prs.opened_could_be_bypassed

    # All PRs have successful builds
    # No PRs have failed builds
    are_builds_ok = checks.all_required_prs_have_successful_builds and checks.all_required_prs_have_no_failed_builds

    feature_api_publicly_available_ok = config.BYPASS_CHECK_API_PUBLICLY_AVAILABLE in summary.checks.bypass or summary.jira.feature_api_is_publicly_available_and_unwhitelisted_in_prod == "Yes"

    design_review_ok = config.BYPASS_CLI_DESIGN_REVIEW_CHECK in summary.checks.bypass or not summary.cli.pending_design_reviews

    cli_changelog_ok = summary.jira.cli_changelog_entry and summary.jira.cli_changelog_entry.strip() and len(summary.jira.cli_changelog_entry) > 0

    print("is_master_ok: {}, are_builds_ok: {}, feature_api_publicly_available_ok: {}, design_review_ok: {}, cli_changelog_ok: {}".format(is_master_ok, are_builds_ok, feature_api_publicly_available_ok, design_review_ok, cli_changelog_ok))

    problems = []
    if not is_master_ok:
        problems.append("The master pull requests are not done yet.")
    if not are_builds_ok:
        problems.append("The validation builds for the master pull requests are not all passing.")
    if not feature_api_publicly_available_ok:
        problems.append("The feature API is not publicly available yet.")
    if not design_review_ok:
        problems.append("The CLI design review is not done yet.")
    if not cli_changelog_ok:
        problems.append("CLI ChangeLog Entry field is empty.")

    return (is_master_ok and are_builds_ok and feature_api_publicly_available_ok and design_review_ok and cli_changelog_ok), "\n".join(problems)


def handle_transition_for_release_requested(issue_key, issue, summary):
    transitioned = False
    text = None

    ready, problem_text = ready_for_release_requested_status(issue_key, issue, summary)
    if not ready:
        transitioned = True
        summary.transition_from_state = summary.state
        util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_SERVICE_TEAM_WORK_REQUIRED)

        # Refresh issue after transitionh
        issue = util.get_dexreq_issue(issue_key, fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS), expand=["changelog"])

        text = TRANSITION_TO_SERVICE_TEAM_WORK_REQUIRED_TEMPLATE.format(
            status=config.STATUS_SERVICE_TEAM_WORK_REQUIRED,
            problem_text="" if not problem_text else "\n\n{}".format(problem_text))
    else:
        print("Issue {} can stay in status '{}' because service team work appears complete.".format(issue_key, config.STATUS_RELEASE_REQUESTED))

    return transitioned, text


TRANSITION_TO_RELEASE_APPROVED_UNDONE = """\
Automatically transitioned status back to '{status}'. Please do not update the status to '{release_approved_status}' unless you are a member of the SDK/CLI team.
"""


def handle_transition_for_release_approved(issue_key, issue, summary):
    transitioned = False
    text = None

    if config.BYPASS_CHECK_PREVENT_MANUAL_STATUS_CHANGES not in summary.checks.bypass:
        printv("Checking if transition to '{}' was done by unauthorized person".format(config.STATUS_RELEASE_APPROVED))
        unauthorized_change = False
        for cl in reversed(summary.jira.changelog):
            is_status_change = False
            for ci in cl.changed_items:
                if ci.field == "status":
                    # This is the last status change, after this, we stop looking
                    is_status_change = True
                    if ci.new == config.STATUS_RELEASE_APPROVED and cl.author not in config.APPROVED_DEX_TEAM_MEMBERS:
                        # This was an unauthorized manual change to "Release Approved'
                        unauthorized_change = ci.old
                    break
            if is_status_change:
                # We've reached the last status change, stop looking
                break
        if unauthorized_change:
            printv("Undoing the unauthorized change to '{}', since it was not done by the SDK/CLI team. Switching back to '{}'".format(
                config.STATUS_RELEASE_APPROVED,
                unauthorized_change))
            transitioned = True
            summary.transition_from_state = summary.state
            util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, unauthorized_change)

            # Refresh issue after transitionh
            issue = util.get_dexreq_issue(issue_key, fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS), expand=["changelog"])

            text = TRANSITION_TO_RELEASE_APPROVED_UNDONE.format(
                status=unauthorized_change,
                release_approved_status=config.STATUS_RELEASE_APPROVED)

    return transitioned, text


TRANSITION_TO_BACKLOG_SINCE_SPEC_PR_DECLINED = """\
Automatically transitioned status back to '{status}'. The [spec diff pull request|{pr_url}] was {action}. Please revise your spec and generate a new spec artifact. Since you rejected this spec change in the public pipeline, you have to restart the entire process, beginning with preview. Create a new preview DEXREQ ticket with the updated artifact version set that ticket's status to '{processing_requested_state}'.
"""

TRANSITION_TO_RELEASE_REQUESTED_SINCE_SPEC_PR_MERGED = """\
Automatically transitioned status to '{status}'. The [spec diff pull request|{pr_url}] was {action}.
"""


def handle_transition_for_service_team_work_required_preconditions_check(issue_key, issue, summary):
    ready, problem_text = ready_for_release_requested_status(issue_key, issue, summary)

    if ready:
        return True, ""

    return ready, "Not transitioning to '{status}' yet, the service team work is not done yet.{problem_text}\n\n".format(status=config.STATUS_RELEASE_REQUESTED,
        problem_text="" if not problem_text else "\n\n" + problem_text)


def handle_transition_for_service_team_work_required(issue_key, issue, summary):
    return handle_transition_for_spec_review_pr(issue_key, issue, summary, config.STATUS_RELEASE_REQUESTED, TRANSITION_TO_RELEASE_REQUESTED_SINCE_SPEC_PR_MERGED,
        TRANSITION_TO_BACKLOG_SINCE_SPEC_PR_DECLINED, handle_transition_for_service_team_work_required_preconditions_check)


PUBLIC_TRANSITION_HANDLERS = {
    PUBLIC_SERVICE_TEAM_WORK_REQUIRED_STATE: handle_transition_for_service_team_work_required,
    PUBLIC_RELEASE_REQUESTED_STATE: handle_transition_for_release_requested,
    PUBLIC_RELEASE_APPROVED_STATE: handle_transition_for_release_approved,
    PUBLIC_PROCESSING_STATE: handle_transition_for_processing
}


def handle_automated_public_transitions(issue_key, issue, summary):
    return execute_appropriate_transition_handler(PUBLIC_TRANSITION_HANDLERS, issue_key, issue, summary)


#
# GA calendar
#

def show_ga_calendar(count=10):
    base_release_date, base_release_ga_number, release_overrides = process_date_override(PUBLIC_RELEASE_DATE_OVERRIDE,
        PUBLIC_RELEASE_DATE_CADENCE_START_DATE, PUBLIC_RELEASE_DATE_GA_NUMBER, '--public-release-date-overrides')
    base_cut_off_date, ga_number, cut_off_overrides = process_date_override(RELEASE_REQUESTED_CUT_OFF_DATE_OVERRIDE,
        RELEASE_REQUESTED_CUT_OFF_DATE_CADENCE_START_DATE, None, '--public-release-requested-cut-off-date-overrides')

    now_datetime_utc = pytz.utc.localize(datetime.datetime.utcnow())
    now_datetime = now_datetime_utc.astimezone(PACIFIC_TIME_ZONE)
    now = now_datetime.date()
    release_date = base_release_date
    release_requested_cut_off_date = base_cut_off_date
    ga_number = base_release_ga_number

    print("| GA  | SDK/CLI Public Release Date | {} Cut-Off Date |".format(config.STATUS_RELEASE_REQUESTED))
    print("|-----|------------|------------|")

    while count > 0:
        release_date = find_next_matching_date(release_date,
            base_release_date, release_overrides, PUBLIC_RELEASE_DATE_CADENCE_IN_DAYS,
            "public release")
        release_requested_cut_off_date = find_next_matching_date(release_requested_cut_off_date,
            base_cut_off_date, cut_off_overrides, RELEASE_REQUESTED_CUT_OFF_DATE_CADENCE_IN_DAYS,
            "public release requested cut-off")

        if release_date >= now:
            print("| {} | {} | {} |".format(ga_number, release_date, release_requested_cut_off_date))
            count -= 1

        release_date += datetime.timedelta(1)
        release_requested_cut_off_date += datetime.timedelta(1)
        ga_number += 1
