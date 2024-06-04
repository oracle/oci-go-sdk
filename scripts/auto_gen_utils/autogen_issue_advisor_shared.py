import datetime
import pytz
import config
import os
import parse
import re
import sys
from dotmap import DotMap
import util
from recordclass import recordclass

import shared.bitbucket_utils

IS_VERBOSE = False
PROCESS_COMMENTS_BY_ANYONE = False

QUIET_TIME_MINUTES = 5

value = os.environ.get('ALLOWED_PROCESSING_TIME_IN_HOURS')
if value:
    ALLOWED_PROCESSING_TIME_IN_HOURS = int(value)
else:
    ALLOWED_PROCESSING_TIME_IN_HOURS = 4

PACIFIC_TIME_ZONE = pytz.timezone("America/Los_Angeles")

DEXREQ_AUTOMATION_NAME = "DEXREQ Automation"
TICKET_STATE_ADVISORY_TEXT = "Ticket state summary:"

BRANCH_LINK_TYPE = "branch"
BUILD_LINK_TYPE = "build"
PR_LINK_TYPE = "diff"  # (pull request)

DEFAULT_JIRA_ISSUE_FIELDS = ['summary', 'description', 'labels', 'comment', 'status', 'reporter', 'issuetype', 'created']
CUSTOM_JIRA_ISSUE_FIELDS = [
    config.CUSTOM_FIELD_ID_ARTIFACT_ID,
    config.CUSTOM_FIELD_ID_GROUP_ID,
    config.CUSTOM_FIELD_ID_ARTIFACT_VERSION,
    config.CUSTOM_FIELD_ID_SPEC_LOCATION_IN_ARTIFACT,
    config.CUSTOM_FIELD_ID_SPEC_FRIENDLY_NAME,
    config.CUSTOM_FIELD_ID_SERVICE_SUBDOMAIN,
    config.CUSTOM_FIELD_ID_FEATURE_IDS,
    config.CUSTOM_FIELD_ID_JAVA_SDK_STATUS,
    config.CUSTOM_FIELD_ID_PYTHON_SDK_STATUS,
    config.CUSTOM_FIELD_ID_TYPESCRIPT_SDK_STATUS,
    config.CUSTOM_FIELD_ID_DOTNET_SDK_STATUS,
    config.CUSTOM_FIELD_ID_RUBY_SDK_STATUS,
    config.CUSTOM_FIELD_ID_GO_SDK_STATUS,
    config.CUSTOM_FIELD_ID_CLI_STATUS,
    config.CUSTOM_FIELD_ID_POWERSHELL_STATUS,
    config.CUSTOM_FIELD_ID_TEST_DATA_STATUS,
    config.CUSTOM_FIELD_ID_LEGACY_JAVA_SDK_STATUS,
    config.CUSTOM_FIELD_ID_SDK_CLI_GA_DATE,
    config.CUSTOM_FIELD_ID_PREVIEW_ISSUE,
    config.CUSTOM_FIELD_ID_CHANGELOG,
    config.CUSTOM_FIELD_ID_ACKNOWLEDGE_RESPONSIBILITIES,
    config.CUSTOM_FIELD_ID_UDX_TICKET,
    config.CUSTOM_FIELD_ID_FEATURE_API_IS_PUBLICLY_AVAILABLE,
    config.CUSTOM_FIELD_ID_CLI_CHANGELOG
]


ERROR_CHOOSE_PIPELINE_STATE = "ERROR_CHOOSE_PIPELINE_STATE"

ERROR_STATES = [
    ERROR_CHOOSE_PIPELINE_STATE
]

PIPELINE_NAMES = [config.PREVIEW_ISSUE_TYPE_NAME, config.PUBLIC_ISSUE_TYPE_NAME]


ServiceTeamMasterPrs = recordclass('ServiceTeamMasterPrs', 'merged approved approved_could_be_bypassed opened opened_could_be_bypassed missing bypassed')
BitbucketBuilds = recordclass('BitbucketBuilds', 'failed_master_pr_builds successful_master_pr_builds in_progress_master_pr_builds')
BitbucketBuildChecks = recordclass('BitbucketBuildChecks', 'all_required_prs_have_successful_builds all_required_prs_have_no_failed_builds')


# Sanity check
def variations_sanity_check(variations, variations_quick, name):
    for v in variations:
        found = False
        for q in variations_quick:
            if q in v:
                found = True
                break
        if not found:
            raise Exception("One of the {0} didn't contain any of the {0}_QUICK".format(name))


# If none of the *_QUICK variations is present, won't ever try the full regex
BUILD_FAIL_JIRA_MESSAGE_TEMPLATE_VARIATIONS_QUICK = ["Building the change for failed for tool:", "Building the change failed for tool:"]
BUILD_FAIL_JIRA_MESSAGE_TEMPLATE_VARIATIONS = [config.BUILD_FAIL_JIRA_MESSAGE_TEMPLATE,
"""Building the change for failed for tool: {tool_name} (repos: {repos}) {build_description}.

{links}

If necessary, you can use {branch_text} as a starting point to fix the build (e.g. if you made a breaking change in preview and you now have to change tests or samples).""",
"""Building the change failed for tool: {tool_name} (repos: {repos}) {build_description}.

{links}

If necessary, you can use {branch_text} as a starting point to fix the build (e.g. if you made a breaking change in preview and you now have to change tests or samples).

Information on the failure is in the following TeamCity city [build log|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId={build_id}&tab=buildLog]""",
"""Building the change failed for tool: {tool_name} (repos: {repos}) {build_description}.

{links}

If necessary, you can use {branch_text} as a starting point to fix the build (e.g. if you made a breaking change in preview and you now have to change tests or samples).

Information about the failure can be found in the {build_log_link}."""
]  # noqa: E124
variations_sanity_check(BUILD_FAIL_JIRA_MESSAGE_TEMPLATE_VARIATIONS, BUILD_FAIL_JIRA_MESSAGE_TEMPLATE_VARIATIONS_QUICK, "BUILD_FAIL_JIRA_MESSAGE_TEMPLATE_VARIATIONS")

BUILD_PASS_JIRA_MESSAGE_TEMPLATE_VARIATIONS_QUICK = ["Completed generating tool:"]
BUILD_PASS_JIRA_MESSAGE_TEMPLATE_VARIATIONS = [config.BUILD_PASS_JIRA_MESSAGE_TEMPLATE,
"""Completed generating tool: {tool_name} (repos: {repos}) {build_description}.

{links}

Artifacts can be found [here|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId={build_id}&tab=artifacts] (for the next 10 days).

For TeamCity access see the [self-service FAQ|https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=132774928#SDK/CLISelfServiceFrequentlyAskedQuestions-HowdoIgetaccesstoTeamCitylogsandartifacts?].""",
"""Completed generating tool: {tool_name} (repos: {repos}) {build_description}.

{links}

Artifacts can be found {build_artifacts_link} (for the next 10 days).

For build log and artifact access see the [self-service FAQ|https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=132774928#SDK/CLISelfServiceFrequentlyAskedQuestions-HowdoIgetaccesstologsandartifacts?].""",
]  # noqa: E124
variations_sanity_check(BUILD_PASS_JIRA_MESSAGE_TEMPLATE_VARIATIONS, BUILD_PASS_JIRA_MESSAGE_TEMPLATE_VARIATIONS_QUICK, "BUILD_PASS_JIRA_MESSAGE_TEMPLATE_VARIATIONS")

STEP_FAILED_MESSAGE_TEMPLATE_VARIATIONS_QUICK = ["If it is unclear how to resolve the issue, you can set",]
STEP_FAILED_MESSAGE_TEMPLATE_VARIATIONS = [config.STEP_FAILED_MESSAGE_TEMPLATE,
"""{failure_step} for tool: {tool_name} (repos: {repos}) failed.
                        
Information on the failure is in the following TeamCity city [build log|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId={build_id}&tab=buildLog]
                        
If it is unclear how to resolve the issue, you can set {custom_status_field} to 'Manual Attention Required to request help from the SDK / CLI team.""",
"""{failure_step} for tool: {tool_name} (repos: {repos}) failed.

Information on the failure is in the following TeamCity city [build log|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId={build_id}&tab=buildLog]

If it is unclear how to resolve the issue, you can set {custom_status_field} to 'Manual Attention Required to request help from the SDK / CLI team.""",
"""{failure_step} for tool: {tool_name} (repos: {repos}) failed.

Information on the failure is in the following TeamCity city [build log|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId={build_id}&tab=buildLog]

If it is unclear how to resolve the issue, you can set the status to '{dex_support_required_status}' to request help from the SDK / CLI team.""",
"""{failure_step} for tool: {tool_name} (repos: {repos}) failed.

Information on the failure is in the following TeamCity city [build log|https://teamcity.oci.oraclecorp.com/viewLog.html?buildId={build_id}&tab=buildLog]

If it is unclear how to resolve the issue, you can set the status to '{dex_support_required_status}' to request help from the SDK / CLI team.{additional_comment}""",
"""{failure_step} for tool: {tool_name} (repos: {repos}) failed.

Information about the failure can be found in the {build_log_link}.

If it is unclear how to resolve the issue, you can set the status to '{dex_support_required_status}' to request help from the SDK / CLI team.{additional_comment}""",
]  # noqa: E124
variations_sanity_check(STEP_FAILED_MESSAGE_TEMPLATE_VARIATIONS, STEP_FAILED_MESSAGE_TEMPLATE_VARIATIONS_QUICK, "STEP_FAILED_MESSAGE_TEMPLATE_VARIATIONS")

TRANSITION_TO_DEX_SUPPORT_REQUESTED_TEMPLATE = """\
Automatically transitioning status to '{status}' since it has been stuck in processing for too long.
"""


def printv(s, flush=False):
    if IS_VERBOSE:
        print(s)
        if flush:
            # Flush, so we make sure the output of the issue key is already visible
            # NOTE: This is to help debug for DEX-6382
            sys.stdout.flush()


def check_should_update(summary):
    should_update = True

    last_comment_was_advisory = False
    # Do not update if the last change is the advisory
    if summary.dates.last and summary.dates.last_issue_advisory:
        if summary.dates.last.created == summary.dates.last_issue_advisory.created:
            # advisory is the last comment
            last_comment_was_advisory = True
            printv("Previous advisory is the last comment.")

    # Do not update if nothing has changed since the last advisory
    if summary.dates.last_issue_advisory and summary.dates.last_issue_advisory.created:
        last_advisory = summary.dates.last_issue_advisory.created
        printv("Previous advisory at {}".format(last_advisory))
        last_change = None
        if summary.dates.last_changelog and summary.dates.last_changelog.created:
            printv("Last JIRA change at  {}".format(summary.dates.last_changelog.created))
            last_change = summary.dates.last_changelog.created
        if summary.dates.last_pr_change:
            printv("Last PR change at    {}".format(summary.dates.last_pr_change))
            if not last_change or last_change < summary.dates.last_pr_change:
                printv("PR change was after JIRA change")
                last_change = summary.dates.last_pr_change
            else:
                printv("JIRA change was after PR change")
                if last_comment_was_advisory:
                    printv("Last comment is advisory, should not update")
                    should_update = False
        else:
            printv("No PR change, JIRA change is the only change")
            if last_comment_was_advisory:
                printv("Last comment is advisory, should not update")
                should_update = False

        if last_change:
            # Let's see if something was changed after the last advisory
            printv("Last change at       {}".format(last_change))
            if last_advisory < last_change:
                printv("Previous advisory is from before the last change.")
            else:
                should_update = False
                printv("Nothing has changed since the last advisory.")

        # Check if any Design Review Tickets have been closed
        if summary.dates.last_comment and summary.dates.last_comment.body:
            if "CLI Design Review" in summary.dates.last_comment.body:
                last_design_review_comment = summary.dates.last_comment.body.split("CLI Design Review", 1)[1]
                if last_design_review_comment.count("DEX-") != len(summary.cli.pending_design_reviews):
                    should_update = True

    # Do not update if the last change to the fields was less than QUIET_TIME_MINUTES ago.
    # A user may still be editing the ticket. We don't want to run on something that may
    # be inconsistent
    # TODO: we might want to have separate limits for update by automation and update by user
    if summary.dates.last_changelog and summary.dates.last_changelog.created:
        last_change = summary.dates.last_changelog.created
        printv("Last change at {}".format(last_change))
        now = datetime.datetime.utcnow()
        cut_off = now + datetime.timedelta(minutes=-QUIET_TIME_MINUTES)
        cut_off_timestamp = cut_off.isoformat("T")
        printv("Cut-off is at  {}".format(cut_off_timestamp))
        if last_change < cut_off_timestamp:
            printv("Quiet time has passed.")
        else:
            printv("Still in quiet time, not making advisory.")
            should_update = False

    return should_update


SUCCESSFUL_PULL_REQUESTS_TEMPLATE = """\
These are the most recent successful links to generated code. You can use them to examine the diff of your change:

{links}"""


# Primary repo for tool when occurring in PR links
PRIMARY_REPO_FOR_LINKS = {
    config.JAVA_SDK_NAME:   "java-sdk",
    config.PYTHON_SDK_NAME: "python-sdk",
    config.CLI_NAME:        "python-cli",
    config.RUBY_SDK_NAME:   "ruby-sdk-autogen-fork",
    config.GO_SDK_NAME:     "oci-go-sdk",
    config.TYPESCRIPT_SDK_NAME: "oci-typescript-sdk-autogen-fork",
    config.DOTNET_SDK_NAME:     "oci-dotnet-sdk-autogen-fork",
    config.POWERSHELL_NAME:      "oci-powershell-modules-autogen-fork",
    config.TEST_DATA_GEN_NAME:     "sdk-client-test-data-autogen-fork",
    config.LEGACY_JAVA_SDK_NAME:    "legacy-java-sdk"
}


def get_pr_link_text(link, tool):
    url = link.url

    target_repo = None
    result = re.search("repos/([^/]*)/pull-requests?", url)
    if result:
        if not PRIMARY_REPO_FOR_LINKS[tool] == result.group(1):
            target_repo = result.group(1)

    if target_repo:
        return "- [{target_repo} {link_type} from the {tool} build|{link}]\n".format(tool=tool, link_type=link.link_type, link=url, target_repo=target_repo)
    else:
        return "- [{tool} {link_type}|{link}]\n".format(tool=tool, link_type=link.link_type, link=url)


def get_successful_pull_requests_text(summary):
    successful_links = {}
    for tool, build in summary.builds.last.items():
        if build.successful:
            successful_links[tool] = build.links

    if successful_links:
        text = ""
        for tool, links in successful_links.items():
            for link in links:
                if link.link_type == PR_LINK_TYPE:
                    text = text + get_pr_link_text(link, tool)

        return SUCCESSFUL_PULL_REQUESTS_TEMPLATE.format(links=text)
    else:
        return None


def process_last_builds_failure_step(issue, summary, comment, last_builds, all_builds):
    # If none of the *_QUICK variations is present, won't ever try the full regex
    found = False
    for q in STEP_FAILED_MESSAGE_TEMPLATE_VARIATIONS_QUICK:
        if q in comment.text:
            found = True
            break
    if not found:
        return False

    # Try all templates we've used for the "{failure_step} for tool ... failed" message
    for template in STEP_FAILED_MESSAGE_TEMPLATE_VARIATIONS:
        result = parse.search(template, comment.text)
        # If this template worked and the "Generation" step failed, process it
        # (if the "Build" step failed, we process a different message in process_last_builds_build_fail below)
        if result and "Generation" in result["failure_step"]:
            tool_name = result["tool_name"]
            if "build_id" in result:
                build_url = "https://teamcity.oci.oraclecorp.com/viewLog.html?buildId={build_id}&tab=buildLog".format(build_id=result["build_id"])
            elif "build_log_link" in result:
                build_url = result["build_log_link"]
            else:
                build_url = None

            # comments are sorted by time, so later builds overwrite newer ones
            build = DotMap()
            build.successful = False
            build.generation_successful = False
            build.build_successful = False
            build.tool_name = tool_name
            build.created = comment.created
            build.repos = result["repos"]

            # process links
            build.links = [DotMap({
                "url": build_url,
                "link_type": BUILD_LINK_TYPE
            })]

            last_builds[tool_name] = build
            if not all_builds[tool_name]:
                all_builds[tool_name] = []
            all_builds[tool_name].append(build)

            return True  # If we have found one, the other variations don't matter anymore

    return False


def process_last_builds_build_failed(issue, summary, comment, last_builds, all_builds):
    # If none of the *_QUICK variations is present, won't ever try the full regex
    found = False
    for q in BUILD_FAIL_JIRA_MESSAGE_TEMPLATE_VARIATIONS_QUICK:
        if q in comment.text:
            found = True
            break
    if not found:
        return False

    # Try all templates we've used for the "Building the change for failed for tool..." message
    for template in BUILD_FAIL_JIRA_MESSAGE_TEMPLATE_VARIATIONS:
        result = parse.search(template, comment.text)
        # If this template worked, process it
        if result:
            tool_name = result["tool_name"]
            if "build_id" in result:
                build_url = "https://teamcity.oci.oraclecorp.com/viewLog.html?buildId={build_id}&tab=buildLog".format(build_id=result["build_id"])
            elif "build_log_link" in result:
                build_url = result["build_log_link"]
            else:
                build_url = None

            # comments are sorted by time, so later builds overwrite newer ones
            build = DotMap()
            build.successful = False
            build.generation_successful = True
            build.build_successful = False
            build.tool_name = tool_name
            build.created = comment.created
            build.repos = result["repos"]

            # process links
            links = re.findall(r"\|([^]]*)\]", result["links"], re.MULTILINE)
            build.links = []
            for url in links:
                build.links.append(DotMap({"url": url, "link_type": BRANCH_LINK_TYPE}))

            if "build_id" in result:
                build.links.append(DotMap({
                    "url": build_url,
                    "link_type": BUILD_LINK_TYPE
                }))

            last_builds[tool_name] = build
            if not all_builds[tool_name]:
                all_builds[tool_name] = []
            all_builds[tool_name].append(build)

            return True  # If we have found one, the other variations don't matter anymore

    return False


def process_last_builds_build_passed(issue, summary, comment, last_builds, all_builds):
    # If none of the *_QUICK variations is present, won't ever try the full regex
    found = False
    for q in BUILD_PASS_JIRA_MESSAGE_TEMPLATE_VARIATIONS_QUICK:
        if q in comment.text:
            found = True
            break
    if not found:
        return False

    # Try all templates we've used for the "Completed generating tool..." message
    for template in BUILD_PASS_JIRA_MESSAGE_TEMPLATE_VARIATIONS:
        result = parse.search(template, comment.text)
        # If this template worked, process it
        if result:
            tool_name = result["tool_name"]

            # comments are sorted by time, so later builds overwrite newer ones
            build = DotMap()
            build.successful = True
            build.generation_successful = True
            build.build_successful = True
            build.tool_name = tool_name
            build.created = comment.created
            build.repos = result["repos"]

            # process links
            links = re.findall(r"\|([^]]*)\]", result["links"], re.MULTILINE)
            build.links = []
            for url in links:
                build.links.append(DotMap({"url": url, "link_type": PR_LINK_TYPE}))

            last_builds[tool_name] = build
            if not all_builds[tool_name]:
                all_builds[tool_name] = []
            all_builds[tool_name].append(build)

            return True  # If we have found one, the other variations don't matter anymore

    return False


def process_last_builds(issue, summary):
    last_builds = DotMap()
    all_builds = DotMap()
    for comment in summary.jira.comments:
        # Unify line endings
        comment.text = comment.text.replace('\r\n', '\n').replace('\r', '\n')
        if PROCESS_COMMENTS_BY_ANYONE or comment.author == DEXREQ_AUTOMATION_NAME:
            found = process_last_builds_failure_step(issue, summary, comment, last_builds, all_builds)
            if found:
                # If it's a failure step, it can't be one of the other ones
                continue

            found = process_last_builds_build_failed(issue, summary, comment, last_builds, all_builds)   
            if found:
                # If it's a build failed message, it can't be one of the other ones
                continue

            found = process_last_builds_build_passed(issue, summary, comment, last_builds, all_builds)
            if found:
                # If it's a build passed message, it can't be one of the other ones
                continue

    summary.builds.last = last_builds
    summary.builds.all = all_builds


FAILED_BUILDS_TEMPLATE = """\


These are the most recent failed builds and branches:

{builds}"""


def get_failed_links_text(summary):
    failed_builds = {}
    for tool, build in summary.builds.last.items():
        if not build.successful:
            failed_builds[tool] = build.links

    if failed_builds:
        text = ""
        for tool, links in failed_builds.items():
            for link in links:
                if link.link_type == PR_LINK_TYPE:
                    text = text + get_pr_link_text(link, tool)
                if link.link_type == BRANCH_LINK_TYPE:
                    result = re.search("repos/([^/]*)/browse?", link.url)
                    tool_from_link = tool
                    if result:
                        for tn, rn in util.get_jira_reportable_repo_names_for_tool().items():
                            # The last of the repos is actually the one that belongs to this tool
                            # (see PythonCLI: [python-sdk, python-cli])
                            if rn[-1] == result.group(1):
                                tool_from_link = tn
                                break

                    text = text + "- [{tool_from_link} {link_type}|{link}]\n".format(tool_from_link=tool_from_link, link_type=link.link_type, link=link.url)
                else:
                    text = text + "- [{tool} {link_type}|{link}]\n".format(tool=tool, link_type=link.link_type, link=link.url)

        return FAILED_BUILDS_TEMPLATE.format(builds=text)
    else:
        return None


# Turn a list into a string, with " and " before the last element
def list_to_string(list, item_prefix=""):
    if len(list) == 0:
        return ""
    elif len(list) == 1:
        return "{}{}".format(item_prefix, list[0])
    elif len(list) == 2:
        return "{}{} and {}{}".format(item_prefix, list[0], item_prefix, list[1])
    else:
        return "{}, and {}".format(", ".join("{}{}".format(item_prefix, str(e)) for e in list[:-1]), "{}{}".format(item_prefix, list[-1]))


def handle_transition_for_processing(issue_key, issue, summary):
    transitioned = False
    text = None

    printv("Checking for transition from {} to {}...".format(summary.jira.status, config.STATUS_DEX_SUPPORT_REQUIRED))
    last_status_update = None
    for cl in reversed(summary.jira.changelog):
        is_status_change = False
        for ci in cl.changed_items:
            if ci.field == "status":
                # This is the last status change, after this, we stop looking
                is_status_change = True
                if cl.author == DEXREQ_AUTOMATION_NAME:
                    last_status_update = cl.created
                break
        if is_status_change:
            # We've reached the last status change, stop looking
            break

    if not last_status_update:
        printv("Ticket was not set to {} by DEXREQ_AUTOMATION. Skipping checking this ticket".format(summary.jira.status))
        return transitioned, text

    printv("Ticket status updated at {}".format(last_status_update))
    now = datetime.datetime.utcnow()
    cut_off = now + datetime.timedelta(hours=-ALLOWED_PROCESSING_TIME_IN_HOURS)
    cut_off_timestamp = cut_off.isoformat("T")
    printv("Cut-off is at {}".format(cut_off_timestamp))

    if last_status_update < cut_off_timestamp:
        transitioned = True
        printv("Ticket stuck in processing for too long. Setting the ticket status to DEX_SUPPORT_REQUIRED")
        util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_DEX_SUPPORT_REQUIRED)
        text = TRANSITION_TO_DEX_SUPPORT_REQUESTED_TEMPLATE.format(status=config.STATUS_DEX_SUPPORT_REQUIRED)
    else:
        printv("Auto-generation of DEX surfaces in progress. Will check again later.")

    return transitioned, text


def get_spec_review_pr(issue_key):
    issue = util.get_dexreq_issue(issue_key, fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS))

    created_date = getattr(issue.fields, 'created')
    printv("To get spec diff PR, listing all PRs newer than {}".format(created_date))

    # The spec diff PR can't be older than the DEXREQ ticket, so only search that far
    return shared.bitbucket_utils.get_newest_pullrequest_with_string_after('SDK', config.DEXREQ_REPO_NAME, issue_key, created_date)


def handle_transition_for_spec_review_pr(issue_key, issue, summary, approved_target_state, approved_template, declined_template, check_approved_preconditions_fn):
    transitioned = False
    text = None

    dexreq_pr_url = None
    dexreq_pr = get_spec_review_pr(issue_key)

    if dexreq_pr:
        hrefs = util.deep_get(dexreq_pr, 'links.self')
        if hrefs:
            dexreq_pr_url = util.deep_get(hrefs[0], 'href')

    if dexreq_pr_url:
        rejected_text = "set to 'Needs Work'"
        accepted_text = "approved"
        is_rejected = False
        is_accepted = False

        for reviewer in dexreq_pr['reviewers']:
            if reviewer['status'] == 'APPROVED':
                is_accepted = True
            if reviewer['status'] == 'NEEDS_WORK':
                is_rejected = True

        if dexreq_pr['state'] == 'DECLINED':
            is_rejected = True
            rejected_text = "declined"

        if dexreq_pr['state'] == 'MERGED':
            is_accepted = True
            accepted_text = "merged"

        if is_rejected:
            transitioned = True
            summary.transition_from_state = summary.state
            util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_BACKLOG)

            # Refresh issue after transitionh
            issue = util.get_dexreq_issue(issue_key, fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS), expand=["changelog"])

            text = declined_template.format(status=config.STATUS_BACKLOG,
                pr_url=dexreq_pr_url,
                processing_requested_state=config.STATUS_PROCESSING_REQUESTED,
                action=rejected_text)

        elif is_accepted:
            are_preconditions_met, precondition_check_text = check_approved_preconditions_fn(issue_key, issue, summary)
            if are_preconditions_met:
                transitioned = True
                summary.transition_from_state = summary.state
                util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, approved_target_state)

                # Refresh issue after transitionh
                issue = util.get_dexreq_issue(issue_key, fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS), expand=["changelog"])

                text = approved_template.format(status=approved_target_state,
                    pr_url=dexreq_pr_url,
                    action=accepted_text)
            else:
                text = precondition_check_text

    return transitioned, text


def execute_appropriate_handler(handlers, issue_key, issue, summary):
    text = ""
    if summary.state in handlers and handlers[summary.state]:
        handler = handlers[summary.state]
        if handler:
            text = handler(issue_key, issue, summary)

    return text


def execute_appropriate_transition_handler(handlers, issue_key, issue, summary):
    transitioned = False
    text = ""
    if summary.state in handlers and handlers[summary.state]:
        handler = handlers[summary.state]
        if handler:
            transitioned, text = handler(issue_key, issue, summary)

    return transitioned, text


def process_all_spec_change_prs(issue):
    warnings = None
    created_date = getattr(issue.fields, 'created')
    printv("To get all spec diff PR, listing all PRs newer than {}".format(created_date))

    # The spec diff PR can't be older than the DEXREQ ticket, so only search that far

    prs = shared.bitbucket_utils.get_all_pullrequest_with_string_after('SDK', config.DEXREQ_REPO_NAME, issue.key, created_date)

    open_spec_change_prs = []
    merged_spec_change_prs = []
    most_recent_approved_spec_change_pr = None

    for pr in prs:
        print("Spec change pr {} is {}".format(pr['id'], pr['state']))

        if pr['state'] == config.PULL_REQUEST_STATUS_OPEN:
            open_spec_change_prs.append(pr)

            if not most_recent_approved_spec_change_pr or most_recent_approved_spec_change_pr['toRef']['id'] < pr['toRef']['id']:
                is_accepted = False
                is_rejected = False
                if pr['reviewers']:
                    for reviewer in pr['reviewers']:
                        if reviewer['status'] == 'APPROVED':
                            is_accepted = True
                        if reviewer['status'] == 'NEEDS_WORK':
                            is_rejected = True

                if is_accepted and not is_rejected:
                    most_recent_approved_spec_change_pr = pr
        elif pr['state'] == config.PULL_REQUEST_STATUS_MERGED:
            merged_spec_change_prs.append(pr)

    printv("Found {} open PRs and {} merged PRs. Most recent approved PR is {}".format(len(open_spec_change_prs), len(merged_spec_change_prs),
        most_recent_approved_spec_change_pr['id'] if most_recent_approved_spec_change_pr else "None"))

    if not merged_spec_change_prs and most_recent_approved_spec_change_pr:
        # Nothing has been merged yet, merge the most recent approved spec change PR
        printv("Merging spec change PR {}".format(most_recent_approved_spec_change_pr['id']))
        shared.bitbucket_utils.merge_pr("SDK", config.DEXREQ_REPO_NAME, most_recent_approved_spec_change_pr['id'], most_recent_approved_spec_change_pr['version'])

    if most_recent_approved_spec_change_pr:
        for pr in open_spec_change_prs:
            if pr == most_recent_approved_spec_change_pr:
                continue

            # Decline other open PRs
            printv("Declining spec change PR {}".format(pr['id']))
            shared.bitbucket_utils.decline_pr("SDK", config.DEXREQ_REPO_NAME, pr['id'], pr['version'])
    elif open_spec_change_prs:
        warnings = "Found open spec change PRs, but none of them were approved. Why not?"
        for pr in open_spec_change_prs:
            hrefs = util.deep_get(pr, 'links.self')
            if hrefs:
                warnings = warnings + "\n - [Spec diff PR {}|{}]".format(pr['id'], util.deep_get(hrefs[0], 'href'))
            else:
                warnings = warnings + "\n - Spec diff PR {}".format(pr['id'])
        print(warnings)

    return warnings


# returns base_date, base_date_ga_number overrides
def process_date_override(override_string, default_base_date, default_base_date_ga_number, parameter_name):
    base_date_string = None
    base_date_ga_number = None
    overrides = {}
    if override_string:
        parts = override_string.split(",")
        for p in parts:
            date = p[1:]
            if p.startswith("="):
                if base_date_string:
                    raise ValueError("Base date ('=YYYY-MM-DD') for {} set more than once: '{}' and '{}'".format(parameter_name, base_date_string, date))
                base_date_string = date
            else:
                overrides[date] = p.startswith("+")

    if not base_date_string:
        base_date_string = default_base_date

    if "@" in base_date_string:
        if not default_base_date_ga_number:
            raise ValueError("Cannot set GA number ('=YYYY-MM-DD@GA') for {}: '{}'".format(parameter_name, base_date_string))
        parts = base_date_string.split("@")
        if len(parts) != 2:
            raise ValueError("Should contain at most one '@' ('=YYYY-MM-DD' or '=YYYY-MM-DD@GA') for {}: '{}'".format(parameter_name, base_date_string))
        base_date_string = parts[0]
        base_date_ga_number = int(parts[1])

    if not base_date_ga_number:
        base_date_ga_number = default_base_date_ga_number

    base_date = datetime.datetime.strptime(base_date_string, "%Y-%m-%d").date()

    return base_date, base_date_ga_number, overrides


def find_next_matching_date(start_date, base_date, overrides, cadence_in_days, date_name):
    date = start_date
    while True:
        printv("{}: {}".format(date_name, date))
        date_string = date.isoformat()
        if (date - base_date).days % cadence_in_days == 0:
            if date_string in overrides and not overrides[date_string]:
                # NOT on this day
                printv("Elected to not have regular {} scheduled on: {}".format(date_name, date_string))
                date += datetime.timedelta(1)
                continue
            else:
                # either no override set for this date, or it's a positive override
                printv("Having regular {} scheduled on: {}".format(date_name, date_string))
                break
        if date_string in overrides and overrides[date_string]:
            # cut-off on this day
            printv("Elected to have special {} scheduled on: {}".format(date_name, date_string))
            break

        date += datetime.timedelta(1)

    return date
