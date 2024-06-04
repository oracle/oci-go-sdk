import argparse
import click
import datetime
import json
import os
import re
import requests
import shutil
import tempfile
import textwrap
import traceback
import zipfile

import util
import config

import shared.version_utils
from shared.buildsvc_tc_compatibility import build_log_link

from jira import JIRAError

DEFAULT_JIRA_ISSUE_FIELDS = ['summary', 'description', 'status', 'labels']
CUSTOM_JIRA_ISSUE_FIELDS = [
    config.CUSTOM_FIELD_ID_ARTIFACT_ID,
    config.CUSTOM_FIELD_ID_GROUP_ID,
    config.CUSTOM_FIELD_ID_ARTIFACT_VERSION,
    config.CUSTOM_FIELD_ID_SPEC_LOCATION_IN_ARTIFACT,
    config.CUSTOM_FIELD_ID_SPEC_FRIENDLY_NAME,
    config.CUSTOM_FIELD_ID_SERVICE_SUBDOMAIN,
    config.CUSTOM_FIELD_ID_FEATURE_IDS,
    config.CUSTOM_FIELD_ID_UDX_TICKET,
    config.CUSTOM_FIELD_ID_JAVA_SDK_STATUS,
    config.CUSTOM_FIELD_ID_PYTHON_SDK_STATUS,
    config.CUSTOM_FIELD_ID_RUBY_SDK_STATUS,
    config.CUSTOM_FIELD_ID_GO_SDK_STATUS,
    config.CUSTOM_FIELD_ID_TYPESCRIPT_SDK_STATUS,
    config.CUSTOM_FIELD_ID_DOTNET_SDK_STATUS,
    config.CUSTOM_FIELD_ID_CLI_STATUS,
    config.CUSTOM_FIELD_ID_POWERSHELL_STATUS,
    config.CUSTOM_FIELD_ID_TEST_DATA_STATUS,
    config.CUSTOM_FIELD_ID_LEGACY_JAVA_SDK_STATUS
]

CUSTOM_JIRA_ISSUE_FIELDS_PER_BUILD_TYPE = {
    config.BUILD_TYPE_INDIVIDUAL_PREVIEW: [],
    config.BUILD_TYPE_BULK_PENDING_MERGE_PREVIEW: [],
    config.BUILD_TYPE_INDIVIDUAL_PUBLIC: [
        config.CUSTOM_FIELD_ID_PREVIEW_ISSUE,
        config.CUSTOM_FIELD_ID_SDK_CLI_GA_DATE,
        config.CUSTOM_FIELD_ID_CHANGELOG,
        config.CUSTOM_FIELD_ID_ACKNOWLEDGE_RESPONSIBILITIES
    ],
    config.BUILD_TYPE_BULK_PENDING_MERGE_PUBLIC: [
        config.CUSTOM_FIELD_ID_PREVIEW_ISSUE,
        config.CUSTOM_FIELD_ID_SDK_CLI_GA_DATE,
        config.CUSTOM_FIELD_ID_CHANGELOG,
        config.CUSTOM_FIELD_ID_ACKNOWLEDGE_RESPONSIBILITIES
    ]
}

FEATURE_ID_DIR = 'featureIds'
CODEGEN_CONFIG_DIR = 'codegenConfig'
CONDITIONAL_PRE_PROCESSOR_DIR = 'enabledGroups'

FEATURE_ID_FILES_FOR_TOOL = {
    config.JAVA_SDK_NAME:      [os.path.join(config.JAVA_SDK_REPO_RELATIVE_LOCATION, "bmc-codegen", FEATURE_ID_DIR)],
    config.PYTHON_SDK_NAME:    [os.path.join(config.PYTHON_SDK_REPO_RELATIVE_LOCATION, CODEGEN_CONFIG_DIR, FEATURE_ID_DIR)],
    config.CLI_NAME:           [
        os.path.join(config.PYTHON_SDK_REPO_RELATIVE_LOCATION, CODEGEN_CONFIG_DIR, FEATURE_ID_DIR),
        os.path.join(config.CLI_REPO_RELATIVE_LOCATION, CODEGEN_CONFIG_DIR, FEATURE_ID_DIR)],
    config.RUBY_SDK_NAME:      [os.path.join(config.RUBY_SDK_REPO_RELATIVE_LOCATION, "codegen", CODEGEN_CONFIG_DIR, FEATURE_ID_DIR)],
    config.GO_SDK_NAME:        [os.path.join(config.GO_SDK_REPO_RELATIVE_LOCATION, CODEGEN_CONFIG_DIR, FEATURE_ID_DIR)],
    config.TEST_DATA_GEN_NAME: [os.path.join(config.TEST_DATA_GEN_REPO_RELATIVE_LOCATION, "codegen", CODEGEN_CONFIG_DIR, FEATURE_ID_DIR)],
    config.TYPESCRIPT_SDK_NAME: [os.path.join(config.TYPESCRIPT_SDK_REPO_RELATIVE_LOCATION, "codegen", FEATURE_ID_DIR)],
    config.DOTNET_SDK_NAME:    [os.path.join(config.DOTNET_SDK_REPO_RELATIVE_LOCATION, "Codegen", CODEGEN_CONFIG_DIR, FEATURE_ID_DIR)],
    config.POWERSHELL_NAME:    [
        os.path.join(config.POWERSHELL_REPO_RELATIVE_LOCATION, "codegen", CODEGEN_CONFIG_DIR, FEATURE_ID_DIR),
        os.path.join(config.DOTNET_SDK_REPO_RELATIVE_LOCATION, "Codegen", CODEGEN_CONFIG_DIR, FEATURE_ID_DIR)
    ],
    config.LEGACY_JAVA_SDK_NAME:    [os.path.join(config.LEGACY_JAVA_SDK_REPO_RELATIVE_LOCATION, "bmc-codegen", FEATURE_ID_DIR)]
}

ENABLED_GROUPS_FILES_FOR_TOOL = {
    config.JAVA_SDK_NAME:      [os.path.join(config.JAVA_SDK_REPO_RELATIVE_LOCATION, "bmc-codegen", CONDITIONAL_PRE_PROCESSOR_DIR)],
    config.PYTHON_SDK_NAME:    [os.path.join(config.PYTHON_SDK_REPO_RELATIVE_LOCATION, CODEGEN_CONFIG_DIR, CONDITIONAL_PRE_PROCESSOR_DIR)],
    config.CLI_NAME:           [
        os.path.join(config.PYTHON_SDK_REPO_RELATIVE_LOCATION, CODEGEN_CONFIG_DIR, CONDITIONAL_PRE_PROCESSOR_DIR),
        os.path.join(config.CLI_REPO_RELATIVE_LOCATION, CODEGEN_CONFIG_DIR, CONDITIONAL_PRE_PROCESSOR_DIR)],
    config.RUBY_SDK_NAME:      [os.path.join(config.RUBY_SDK_REPO_RELATIVE_LOCATION, "codegen", CODEGEN_CONFIG_DIR, CONDITIONAL_PRE_PROCESSOR_DIR)],
    config.GO_SDK_NAME:        [os.path.join(config.GO_SDK_REPO_RELATIVE_LOCATION, CODEGEN_CONFIG_DIR, CONDITIONAL_PRE_PROCESSOR_DIR)],
    config.TEST_DATA_GEN_NAME: [os.path.join(config.TEST_DATA_GEN_REPO_RELATIVE_LOCATION, "codegen", CODEGEN_CONFIG_DIR, CONDITIONAL_PRE_PROCESSOR_DIR)],
    config.TYPESCRIPT_SDK_NAME: [os.path.join(config.TYPESCRIPT_SDK_REPO_RELATIVE_LOCATION, "codegen",CONDITIONAL_PRE_PROCESSOR_DIR)],
    config.DOTNET_SDK_NAME:    [os.path.join(config.DOTNET_SDK_REPO_RELATIVE_LOCATION, "Codegen", CODEGEN_CONFIG_DIR, CONDITIONAL_PRE_PROCESSOR_DIR)],
    config.POWERSHELL_NAME:    [
        os.path.join(config.POWERSHELL_REPO_RELATIVE_LOCATION, "codegen", CODEGEN_CONFIG_DIR, CONDITIONAL_PRE_PROCESSOR_DIR),
        os.path.join(config.DOTNET_SDK_REPO_RELATIVE_LOCATION, "Codegen", CODEGEN_CONFIG_DIR, CONDITIONAL_PRE_PROCESSOR_DIR)
    ],
    config.LEGACY_JAVA_SDK_NAME:    [os.path.join(config.LEGACY_JAVA_SDK_REPO_RELATIVE_LOCATION, "bmc-codegen", CONDITIONAL_PRE_PROCESSOR_DIR)]
}


# What we call these things in the JIRA ticket
RESULT_FIELD_TO_JIRA_FIELD_NAME = {
    'subdomain': 'Service Subdomain',
    'spec_name': 'Service Friendly Name',
    'group_id': 'Spec Artifact Group Id',
    'artifact_id': 'Spec Artifact Id',
    'version': 'Spec Artifact Version',
    'relative_spec_path': 'Spec Location in Artifact'
}


TOOL_ARGUMENT_ALL = 'ALL'


BRANCH_TIMESTAMP = datetime.datetime.now().strftime('%Y-%m-%d-%H-%M-%S')


def limit_query_to_issue_keys(query, issues, for_any_status):
    if issues:
        issues_query = ' OR '.join(["key = '{}'".format(i) for i in issues])
        if for_any_status:
            query = issues_query
        else:
            query = '({}) AND ({})'.format(issues_query, query)
    return query


# returns True if processing of this issue should be skipped
def report_errors_and_warnings(issue, errors_map, warnings_map, tool_names):
    errors = errors_map[issue.key] if issue.key in errors_map else []
    warnings = warnings_map[issue.key] if issue.key in warnings_map else []

    if warnings and not errors:
        # comment on issue
        util.add_jira_comment(
            issue.key,
            """Several non-blocking issues in the ticket were found and had to be fixed:

            {warnings}

            The full build log can be found {build_log_link}.

            For TeamCity access see the [self-service FAQ|https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=132774928#SDK/CLISelfServiceFrequentlyAskedQuestions-HowdoIgetaccesstoTeamCitylogsandartifacts?].

            There is no need to re-start this ticket, but please keep these problems in mind when filing future tickets.""".format(
                warnings='\n'.join(['*- {}*'.format(warning) for warning in warnings]),
                build_id=build_id,
                build_log_link=build_log_link(build_id)
            ),
            comment_type=config.COMMENT_TYPE_INFO
        )
        return False  # continue
    if errors:
        if warnings:
            # comment on issue
            util.add_jira_comment(
                issue.key,
                """The job failed due to the following errors in the ticket:

                {errors}

                There were also other issues that had to be fixed:

                {warnings}

                The full build log can be found {build_log_link}.

                For TeamCity access see the [self-service FAQ|https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=132774928#SDK/CLISelfServiceFrequentlyAskedQuestions-HowdoIgetaccesstoTeamCitylogsandartifacts?].

                Please fix the referenced input parameters and and update status to '{processing_requested_state}' to re-run.""".format(
                    errors='\n'.join(['*- {}*'.format(error) for error in errors]),
                    warnings='\n'.join(['*- {}*'.format(warning) for warning in warnings]),
                    processing_requested_state=config.STATUS_PROCESSING_REQUESTED,
                    build_log_link=build_log_link(build_id)
                ),
                comment_type=config.COMMENT_TYPE_ERROR
            )
        else:
            # comment on issue
            util.add_jira_comment(
                issue.key,
                """The job failed due to the following errors in the ticket:

                {errors}

                The full build log can be found {build_log_link}.

                For TeamCity access see the [self-service FAQ|https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=132774928#SDK/CLISelfServiceFrequentlyAskedQuestions-HowdoIgetaccesstoTeamCitylogsandartifacts?].

                Please fix the referenced input parameters and update status to '{processing_requested_state}' to re-run.""".format(
                    errors='\n'.join(['*- {}*'.format(error) for error in errors]),
                    processing_requested_state=config.STATUS_PROCESSING_REQUESTED,
                    build_log_link=build_log_link(build_id)
                ),
                comment_type=config.COMMENT_TYPE_ERROR
            )

        if tool_names:
            for tool_name in tool_names:
                if tool_name in config.CUSTOM_FIELD_ID_FOR_TOOL and config.CUSTOM_FIELD_ID_FOR_TOOL[tool_name] and tool_name in config.CUSTOM_FIELD_NAME_FOR_TOOL:
                    # Only transition if there is a field in the Jira item for it
                    util.transition_issue_per_tool_status(util.JIRA_CLIENT(), issue, config.CUSTOM_STATUS_FAILURE, tool_name)

            # Reload after the transition
            issue = util.get_dexreq_issue(issue.key,
                fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS_PER_BUILD_TYPE[build_type]))

        # if an issue is already in 'DEX Support Required' based on failure for another tool, we do not want to overwrite that
        util.transition_issue_overall_status_if_not_in_status(util.JIRA_CLIENT(), issue, desired_status=config.STATUS_SERVICE_TEAM_FAILURE_INVESTIGATION, blacklisted_status=config.STATUS_DEX_SUPPORT_REQUIRED)
        return True  # skip
    return False  # continue


def generate_individual_to_do_requests(all_issues, build_id, base_branch, tool_name, allow_individual_tool_generation, build_type, for_any_status):
    warnings = {}  # key: issue_key, value: list
    errors = {}  # key: issue_key, value: list
    branches = {}

    if build_type == config.BUILD_TYPE_INDIVIDUAL_PUBLIC:
        build_description = "public release"
        branch_prefix = config.INDIVIDUAL_PUBLIC_BRANCH_PREFIX
        spec_generation_type = config.PUBLIC_SPEC_GENERATION_TYPE
    else:
        build_description = "preview"
        branch_prefix = config.INDIVIDUAL_PREVIEW_BRANCH_PREFIX
        spec_generation_type = config.PREVIEW_SPEC_GENERATION_TYPE

    is_tool_jira_reportable = util.is_tool_jira_reportable(tool_name)

    for issue in all_issues:
        custom_field_id_for_tool = config.CUSTOM_FIELD_ID_FOR_TOOL[tool_name]
        if custom_field_id_for_tool:
            custom_status_for_tool = getattr(issue.fields, custom_field_id_for_tool)
            if not for_any_status and allow_individual_tool_generation and issue.fields and (custom_status_for_tool is None or custom_status_for_tool.value != config.CUSTOM_STATUS_TODO):
                print('Skipping generating tool: {} for issue: {} because allow_individual_tool_generation = True and this tool is not set to To Do'.format(tool_name, issue.key))
                continue

        print("==============================================================================")
        print('Generating {} {} for {} - {}'.format(build_description, tool_name, issue.key, issue.fields.summary))
        print("==============================================================================")

        if is_tool_jira_reportable:
            util.transition_issue_per_tool_status(util.JIRA_CLIENT(), issue, config.CUSTOM_STATUS_PROCESSING, tool_name)
            # Remove a possible BACKWARD_INCOMPATIBLE_CHANGES_LABEL
            # Since we're restarting the generation, there may not be any more backward incompatible changes in this version
            # If so, we'll find out later again in warn_about_backward_incompatible_changes.py run as part of the Java SDK generation
            if config.BACKWARD_INCOMPATIBLE_CHANGES_LABEL in issue.fields.labels:
                new_labels = list(issue.fields.labels)
                new_labels.remove(config.BACKWARD_INCOMPATIBLE_CHANGES_LABEL)
                issue.update(fields={"labels": new_labels})

        # Reload after the transition
        issue = util.get_dexreq_issue(issue.key,
            fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS_PER_BUILD_TYPE[build_type]))

        try:
            # start by making sure we are on the base branch with no other changes present
            checkout_sdk_and_cli_branches(base_branch, tool_name)
            for repo in config.REPOS_FOR_TOOL[tool_name]:
                repo.git.reset('HEAD','--hard')
                repo.git.clean('-f')

            feature_ids = getattr(issue.fields, config.CUSTOM_FIELD_ID_FEATURE_IDS)
            if feature_ids:
                feature_ids = [feature_id.strip().lower() for feature_id in feature_ids.split(',')]
                print('Found the following feature IDs:')
                print(feature_ids)
                for feature_id_file in FEATURE_ID_FILES_FOR_TOOL[tool_name]:
                    util.update_feature_id_file(feature_id_file, feature_ids, issue.key)

                for enabled_groups_file in ENABLED_GROUPS_FILES_FOR_TOOL[tool_name]:
                    util.update_pre_processor_file(enabled_groups_file, feature_ids, issue.key)

            add_or_update_spec_params = convert_issue_to_script_params(issue)
            print('Parameters for add or update script:')
            pretty_print(add_or_update_spec_params)

            try:
                result = invoke_add_or_update_spec(util.JIRA_CLIENT(), issue, add_or_update_spec_params, tool_name, spec_generation_type)
                warn_on_unexpected_changes(issue, add_or_update_spec_params, tool_name, build_id, result, warnings)
            except click.exceptions.MissingParameter as e:
                print('ERROR: {}'.format(str(e)))
                add_error(errors, issue.key,
                    """The job failed due to a missing required parameter. {exception}.

                    The full build log can be found {build_log_link}.

                    For TeamCity access see the [self-service FAQ|https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=132774928#SDK/CLISelfServiceFrequentlyAskedQuestions-HowdoIgetaccesstoTeamCitylogsandartifacts?].

                    Please add the missing parameter and update status to '{processing_requested_state}' to re-run.""".format(
                        exception=str(e),
                        build_id=build_id,
                        processing_requested_state=config.STATUS_PROCESSING_REQUESTED,
                        build_log_link=build_log_link(build_id)
                    )
                )
                if is_tool_jira_reportable:
                    util.transition_issue_per_tool_status(util.JIRA_CLIENT(), issue, config.CUSTOM_STATUS_FAILURE, tool_name)

                    # Reload after the transition
                    issue = util.get_dexreq_issue(issue.key,
                        fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS_PER_BUILD_TYPE[build_type]))

                    # if an issue is already in 'DEX Support Required' based on failure for another tool, we do not want to overwrite that
                    util.transition_issue_overall_status_if_not_in_status(util.JIRA_CLIENT(),
                                                                          issue,
                                                                          desired_status=config.STATUS_SERVICE_TEAM_FAILURE_INVESTIGATION,
                                                                          blacklisted_status=config.STATUS_DEX_SUPPORT_REQUIRED)
                continue

            commit_message = '{commit_prefix} [[{issue_key}]]: {issue_summary}'.format(
                commit_prefix=config.POM_UPDATE_COMMIT_MESSAGE_PREFIX,
                issue_key=issue.key,
                issue_summary=issue.fields.summary
            )

            # Include the tool name in the branch so we can differentiate between Python SDK branches created for the Python SDK
            # preview vs Python SDK branches created as a depenency of the CLI preview
            prefix = '{}-{}-{}'.format(branch_prefix, tool_name, issue.key)

            branch_name = generate_time_stamped_branch_name(prefix)
            branches[issue.key] = branch_name
            for repo in config.REPOS_FOR_TOOL[tool_name]:
                git = repo.git
                git.checkout(b=branch_name)
                git.add(A=True)
                message = commit_message
                if 'nothing to commit' in git.status():
                    message = "{} (no change in spec version or feature ids)".format(message)
                print(message)
                git.commit("-m", message, "--allow-empty")
                if config.IS_DRY_RUN:
                    print('DRY-RUN: not pushing to branch {}'.format(branch_name))
                else:
                    git.push('-u','origin','HEAD')

        except Exception as e:
            print('EXCEPTION: {}'.format(str(e)))
            status_field = config.CUSTOM_FIELD_NAME_FOR_TOOL[tool_name] if util.is_tool_jira_reportable(tool_name) else "N/A"
            add_error(errors, issue.key,
                config.SELF_SERVICE_BUG_TEMPLATE.format(
                    exception=str(e),
                    custom_status_field=status_field,
                    build_log_link=build_log_link(build_id)
                )
            )
            if is_tool_jira_reportable:
                util.transition_issue_per_tool_status(util.JIRA_CLIENT(), issue, config.CUSTOM_STATUS_FAILURE, tool_name)
                util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_DEX_SUPPORT_REQUIRED)

                # Reload after the transition
                issue = util.get_dexreq_issue(issue.key,
                    fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS_PER_BUILD_TYPE[build_type]))

            traceback.print_exc()
            continue

    return warnings, errors, branches


def generate_bulk_pending_merge(spec_generation_type, all_issues, build_id, base_branch, tool_name, allow_individual_tool_generation, for_any_status):
    warnings = {}  # key: issue_key, value: list
    errors = {}  # key: issue_key, value: list
    branches = {}

    # start by making sure we are on the base branch with no other changes present
    checkout_sdk_and_cli_branches(base_branch, tool_name)
    for repo in config.REPOS_FOR_TOOL[tool_name]:
        repo.git.reset('HEAD','--hard')
        repo.git.clean('-f')

    is_tool_jira_reportable = util.is_tool_jira_reportable(tool_name)

    successful_issues = []
    for issue in all_issues:
        custom_field_id_for_tool = config.CUSTOM_FIELD_ID_FOR_TOOL[tool_name]
        if custom_field_id_for_tool:
            custom_status_for_tool = getattr(issue.fields, custom_field_id_for_tool)
            if allow_individual_tool_generation and not for_any_status and issue.fields and (custom_status_for_tool is None or custom_status_for_tool.value != config.CUSTOM_STATUS_TODO):
                print('Skipping generating tool: {} for issue: {} because allow_individual_tool_generation = True and this tool is not set to To Do'.format(tool_name, issue.key))
                continue

        if tool_name == config.CLI_NAME:
            if util.is_cli_pr_required(issue):
                print('Skipping generating tool: {} for issue: {} as it requires manual changes'.format(tool_name, issue.key))
                continue

        print("==============================================================================")
        print('Generating {} {} for {} - {}'.format(spec_generation_type, tool_name, issue.key, issue.fields.summary))
        print("==============================================================================")

        if is_tool_jira_reportable:
            util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_PROCESSING_BULK)
            util.transition_issue_per_tool_status(util.JIRA_CLIENT(), issue, config.CUSTOM_STATUS_PROCESSING, tool_name)

        # Reload after the transition
        issue = util.get_dexreq_issue(issue.key,
            fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS_PER_BUILD_TYPE[build_type]))

        try:
            feature_ids = getattr(issue.fields, config.CUSTOM_FIELD_ID_FEATURE_IDS)
            if feature_ids:
                feature_ids = [feature_id.strip().lower() for feature_id in feature_ids.split(',')]
                print('Found the following feature IDs:')
                print(feature_ids)
                for feature_id_file in FEATURE_ID_FILES_FOR_TOOL[tool_name]:
                    util.update_feature_id_file(feature_id_file, feature_ids, issue.key)

                for enabled_groups_file in ENABLED_GROUPS_FILES_FOR_TOOL[tool_name]:
                    util.update_pre_processor_file(enabled_groups_file, feature_ids, issue.key)

            add_or_update_spec_params = convert_issue_to_script_params(issue)
            print('Parameters for add or update script:')
            pretty_print(add_or_update_spec_params)

            try:
                result = invoke_add_or_update_spec(util.JIRA_CLIENT(), issue, add_or_update_spec_params, tool_name, spec_generation_type)
                warn_on_unexpected_changes(issue, add_or_update_spec_params, tool_name, build_id, result, warnings)
            except click.exceptions.MissingParameter as e:
                add_error(errors, issue.key,
                    """The job failed due to a missing required parameter. {exception}.

                    The full build log can be found {build_log_link}.

                    For TeamCity access see the [self-service FAQ|https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=132774928#SDK/CLISelfServiceFrequentlyAskedQuestions-HowdoIgetaccesstoTeamCitylogsandartifacts?].

                    Please add the missing parameter and update status to {processing_requested_state} to re-run.""".format(
                        exception=str(e),
                        build_id=build_id,
                        processing_requested_state=config.STATUS_PROCESSING_REQUESTED,
                        build_log_link=build_log_link(build_id)
                    )
                )

                print('ERROR: {}'.format(str(e)))
                if config.IS_DRY_RUN:
                    print("DRY-RUN: not transitioning {} to {}".format(issue, config.CUSTOM_STATUS_FAILURE))
                elif not is_tool_jira_reportable:
                    print("Not transitioning {} to {} as tool {} is not reported to Jira".format(issue,
                                                                                                 config.CUSTOM_STATUS_FAILURE,
                                                                                                 tool_name))
                else:
                    util.transition_issue_per_tool_status(util.JIRA_CLIENT(), issue, config.CUSTOM_STATUS_FAILURE, tool_name)
                    # if an issue is already in 'DEX Support Required' based on failure for another tool, we do not want to overwrite that
                    util.transition_issue_overall_status_if_not_in_status(util.JIRA_CLIENT(), issue, desired_status=config.STATUS_SERVICE_TEAM_FAILURE_INVESTIGATION, blacklisted_status=config.STATUS_DEX_SUPPORT_REQUIRED)

                    # Reload after the transition
                    issue = util.get_dexreq_issue(issue.key,
                        fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS_PER_BUILD_TYPE[build_type]))
                continue
        except Exception as e:
            status_field = config.CUSTOM_FIELD_NAME_FOR_TOOL[tool_name] if util.is_tool_jira_reportable(tool_name) else "N/A"
            add_error(errors, issue.key,
                config.SELF_SERVICE_BUG_TEMPLATE.format(
                    exception=str(e),
                    custom_status_field=status_field,
                    build_log_link=build_log_link(build_id)
                )
            )

            if config.IS_DRY_RUN:
                print("DRY-RUN: not transitioning {} to {}".format(issue, config.CUSTOM_STATUS_FAILURE))
            elif not is_tool_jira_reportable:
                print("Not transitioning {} to {} as tool {} is not reported to Jira".format(issue,
                                                                                             config.CUSTOM_STATUS_FAILURE,
                                                                                             tool_name))
            else:
                # Mark the individual SDK status values to the corect state regardless if it's bypassed/ignored.
                util.transition_issue_per_tool_status(util.JIRA_CLIENT(), issue, config.CUSTOM_STATUS_FAILURE, tool_name)
                # Only transition the overal ticket status to DEX SUPPORT REQUIRED if the SDK isn't bypassed/ignored.
                if not config.BYPASS_CHECK_GENERATION_PREFIX + tool_name in issue.fields.labels:
                    util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_DEX_SUPPORT_REQUIRED)

                # Reload after the transition
                issue = util.get_dexreq_issue(issue.key,
                    fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS_PER_BUILD_TYPE[build_type]))
            traceback.print_exc()
            continue

        successful_issues.append(issue)

    # only commit changes if some of the issues were successful
    if len(successful_issues) > 0:
        print('Committing changes for issues that successfully generated...')
        successful_issue_keys = ', '.join([si.key for si in successful_issues])
        commit_message = '{commit_prefix} [[{issue_keys}]]'.format(
            commit_prefix=config.POM_UPDATE_COMMIT_MESSAGE_PREFIX,
            issue_keys=successful_issue_keys,
        )

        # Include the tool name in the branch so we can differentiate between Python SDK branches created for the Python SDK
        # preview/master vs Python SDK branches created as a depenency of the CLI preview/master
        branch_prefix_to_use = config.BULK_PUBLIC_BRANCH_PREFIX \
            if spec_generation_type == config.PUBLIC_SPEC_GENERATION_TYPE else config.BULK_PREVIEW_BRANCH_PREFIX
        prefix = '{}-{}'.format(branch_prefix_to_use, tool_name)
        branch_name = generate_time_stamped_branch_name(prefix)

        for issue_key in successful_issue_keys:
            branches[issue_key] = branch_name

        for repo in config.REPOS_FOR_TOOL[tool_name]:
            git = repo.git
            git.checkout(b=branch_name)
            git.add(A=True)
            message = commit_message
            if 'nothing to commit' in git.status():
                message = "{} (no change in spec version or feature ids)".format(message)
            print(message)
            git.commit("-m", message, "--allow-empty")
            if config.IS_DRY_RUN:
                print('DRY-RUN: not pushing to branch {}'.format(branch_name))
            else:
                git.push('-u','origin','HEAD')

    return warnings, errors, branches


def add_error(errors, issue_key, message):
    print(message)
    errors_for_issue = []
    if issue_key in errors:
        errors_for_issue = errors[issue_key]
    else:
        errors[issue_key] = errors_for_issue
    errors_for_issue.append(message)


# Returns a singleton list [{field_id: 'field_id', new_value: 'new_value', field_name: 'field_name'}] triple if anything was changed
def validate_issue_properties(issue, warnings, field_id, converters):
    field_name = config.CUSTOM_FIELD_NAME_FOR_ID[field_id] or field_id
    original_value = getattr(issue.fields, field_id)
    if original_value:
        converted_value = original_value
        for c in converters:
            converted_value = c(converted_value)

        if converted_value != original_value:
            setattr(issue.fields, field_id, converted_value.encode('utf-8'))
            add_error(warnings, issue.key, "The field '{}' had formatting issues.".format(field_name))
            return [{"field_name": field_name, "field_id": field_id, "new_value": converted_value}]  # changed

        return []  # not changed
    else:
        return []  # not changed


def check_for_internal_ticket(issue):
    preview_issue = util.get_dexreq_issue(getattr(issue.fields, config.CUSTOM_FIELD_ID_PREVIEW_ISSUE))
    preview_summary = getattr(preview_issue.fields, "summary")
    preview_labels = getattr(preview_issue.fields, "labels")
    summary = getattr(issue.fields, "summary")
    feature_ids = getattr(issue.fields, config.CUSTOM_FIELD_ID_FEATURE_IDS)
    labels = getattr(issue.fields, "labels")
    if labels:
        for label in labels:
            if "internal" in label.lower():
                return True
    if preview_labels:
        for label in preview_labels:
            if "internal" in label.lower():
                return True
    if feature_ids:
        for feature_id in feature_ids:
            if "internal" in feature_id.lower():
                return True
    if summary and "internal" in summary.lower():
        return True
    if preview_summary and "internal" in preview_summary.lower():
        return True

    return False


def validate_helper_whitespace(original_value):
    return original_value.strip()


def validate_helper_upper(original_value):
    return original_value.upper()


def validate_helper_dexreq_jira(original_value):
    issues = util.get_dexreq_issue_keys(original_value)
    return ", ".join(issues)


def validate_helper_udx_jira(original_value):
    issues = util.get_jira_issue_keys(original_value, "UDX")
    return ", ".join(issues)


def validate_issue_whitespace(issue, warnings, field_id):
    return validate_issue_properties(issue, warnings, field_id, [validate_helper_whitespace])


def validate_issue_whitespace_upper(issue, warnings, field_id):
    return validate_issue_properties(issue, warnings, field_id, [validate_helper_whitespace, validate_helper_upper])


def validate_issue_whitespace_upper_dexreq_jira(issue, warnings, field_id):
    return validate_issue_properties(issue, warnings, field_id, [validate_helper_whitespace, validate_helper_upper, validate_helper_dexreq_jira])


def validate_issue_whitespace_upper_udx_jira(issue, warnings, field_id):
    return validate_issue_properties(issue, warnings, field_id, [validate_helper_whitespace, validate_helper_upper, validate_helper_udx_jira])


# Returns True if "name" contains dashes, periods, or whitespaces; else returns False
def check_disallowed_separators_in_name(name):
    disallowed_separators = ["-", ".", " "]
    for separator in disallowed_separators:
        if separator in name:
            return True
    return False


# Returns true if "name" is in camel case; else returns False
def check_camel_casing(name):
    pattern = '[a-z]+[A-Z]'
    if re.search(pattern, name) and "_" not in name:
        return True
    return False


def validate_issue(issue, build_type, dexreq_public_errors, dexreq_public_warnings):
    errors = {}
    warnings = {}

    errors.update(dexreq_public_errors)
    warnings.update(dexreq_public_warnings)

    print('Checking for whitespace and other formatting issues...')

    # These apply both to preview tickets and public tickets
    # (though for public tickets, they will look/update the preview ticket)
    updates = []
    updates.extend(validate_issue_whitespace(issue, warnings, config.CUSTOM_FIELD_ID_ARTIFACT_ID))
    updates.extend(validate_issue_whitespace(issue, warnings, config.CUSTOM_FIELD_ID_ARTIFACT_VERSION))
    updates.extend(validate_issue_whitespace(issue, warnings, config.CUSTOM_FIELD_ID_GROUP_ID))
    updates.extend(validate_issue_whitespace(issue, warnings, config.CUSTOM_FIELD_ID_SPEC_LOCATION_IN_ARTIFACT))
    updates.extend(validate_issue_whitespace(issue, warnings, config.CUSTOM_FIELD_ID_SPEC_FRIENDLY_NAME))
    updates.extend(validate_issue_whitespace(issue, warnings, config.CUSTOM_FIELD_ID_SERVICE_SUBDOMAIN))
    updates.extend(validate_issue_whitespace(issue, warnings, config.CUSTOM_FIELD_ID_FEATURE_IDS))
    updates.extend(validate_issue_whitespace_upper_udx_jira(issue, warnings, config.CUSTOM_FIELD_ID_UDX_TICKET))

    if build_type == config.BUILD_TYPE_INDIVIDUAL_PUBLIC:
        updates.extend(validate_issue_whitespace_upper_dexreq_jira(issue, warnings, config.CUSTOM_FIELD_ID_PREVIEW_ISSUE))
        updates.extend(validate_issue_whitespace(issue, warnings, config.CUSTOM_FIELD_ID_CHANGELOG))
        updates.extend(validate_issue_whitespace(issue, warnings, config.CUSTOM_FIELD_ID_SDK_CLI_GA_DATE))

    if updates:
        # Make all updates at once
        try:
            error = util.update_issue_fields(util.JIRA_CLIENT(), issue, updates)
            if error:
                add_error(warnings, issue.key, "One or more of the fields '{}' had formatting issues, and we couldn't fix them (Error: {}).".format(", ".join(u['field_name'] for u in updates), error))
        except JIRAError as e:
            add_error(warnings, issue.key, "One or more of the fields '{}' had formatting issues, and we couldn't fix them (Error: {}).".format(", ".join(u['field_name'] for u in updates), str(e)))

        # Reload after making changes
        issue = util.get_dexreq_issue(issue.key,
            fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS_PER_BUILD_TYPE[build_type]))

    version = getattr(issue.fields, config.CUSTOM_FIELD_ID_ARTIFACT_VERSION)
    if not version:
        add_error(errors, issue.key, 'You did not provide an artifact version. This field is always mandatory.')

    artifact_id = getattr(issue.fields, config.CUSTOM_FIELD_ID_ARTIFACT_ID)
    if not artifact_id:
        add_error(errors, issue.key, 'You did not provide an artifact id. This field is always mandatory.')

    if config.BYPASS_CHECK_FEATURE_ID_LABEL not in issue.fields.labels:
        feature_ids = getattr(issue.fields, config.CUSTOM_FIELD_ID_FEATURE_IDS)
        if not feature_ids:
            add_error(errors, issue.key, 'You did not provide a feature id. This field is mandatory. Please provide a feature id. If you have an exceptional case, or if you are a new service that is launching its service for the first time in SDK/CLI, you can request an exception in the slack channel #oci_public_sdks')

    if config.BYPASS_CHECK_SPEC_EXTENSION_LABEL not in issue.fields.labels:
        spec_location = getattr(issue.fields, config.CUSTOM_FIELD_ID_SPEC_LOCATION_IN_ARTIFACT)
        if spec_location and (not spec_location.endswith('.cond.yaml')):
            add_error(errors, issue.key, 'The spec file provided is not a ".cond.yaml" file. Please provide a ".cond.yaml" file. If you have an exceptional case, you can request an exception in the slack channel #oci_public_sdks.')

    if config.BYPASS_CHECK_SERVICE_FRIENDLY_NAME_LABEL not in issue.fields.labels:
        service_friendly_name = getattr(issue.fields, config.CUSTOM_FIELD_ID_SPEC_FRIENDLY_NAME)
        confluence_doc_refer_msg = "Please check [here|https://confluence.oci.oraclecorp.com/pages/viewpage.action?spaceKey=DEX&title=Requesting+a+preview+SDK+CLI#RequestingapreviewSDKCLI-II)IfthisisthefirsttimeyouarerequestingapreviewSDK/CLIforthisspec] for clarity."
        if service_friendly_name and check_disallowed_separators_in_name(service_friendly_name):
            err_message = "Service friendly name provided by you has one/more dashes ('-'), periods ('.'), or whitespaces (' ') as word separators. This is not allowed. Please use underscores ('_') instead to separate words. "
            err_message += confluence_doc_refer_msg
            add_error(errors, issue.key, err_message)
        if service_friendly_name and check_camel_casing(service_friendly_name):
            err_message = "Service friendly name provided by you is in camel-case, which is not allowed. Please use snake-case instead ie use underscores ('_') instead to separate words. "
            err_message += confluence_doc_refer_msg
            add_error(errors, issue.key, err_message)
        if service_friendly_name and ("service" in service_friendly_name.lower()):
            err_message = "Service friendly name provided by you contains the keyword 'service'. This is not allowed. If you have an exceptional case, you can request an exception in the slack channel #oci_public_sdks."
            add_error(errors, issue.key, err_message)

    if config.BYPASS_CHECK_UDX_TICKET_LABEL not in issue.fields.labels:
        udx_ticket = getattr(issue.fields, config.CUSTOM_FIELD_ID_UDX_TICKET)
        if not udx_ticket:
            add_error(errors, issue.key, 'You did not provide a UDX ticket. This field is always mandatory.')
        else:
            udx_ticket_keys = util.get_udx_issue_keys(udx_ticket)
            if udx_ticket_keys:
                jira_client = util.JIRA_CLIENT()
                for u in udx_ticket_keys:
                    try:
                        udx_ticket_issue = jira_client.issue(u)
                        udx_ticket_status = str(getattr(udx_ticket_issue.fields, "status"))
                        if udx_ticket_status in config.UDX_TICKET_DISALLOWED_STATES:
                            add_error(errors, issue.key, 'The UDX ticket provided: {} is {}. It has to be an open ticket.'.format(u, udx_ticket_status))
                        jira_client.create_issue_link(config.UDX_TICKET_LINK_RELATIONSHIP, udx_ticket_issue, issue)
                    except JIRAError as e:
                        if e.status_code == 404:
                            add_error(errors, issue.key, 'The UDX ticket provided could not be resolved. Please ensure {} is a valid JIRA issue.'.format(u))
                        else:
                            add_error(errors, issue.key, 'The UDX ticket provided: {} could not be resolved. Error: {}.'.format(u, str(e)))
            else:
                add_error(errors, issue.key, 'The value provided \'{}\' is not valid for the \'UDX Ticket\' field.'.format(udx_ticket))

    if build_type == config.BUILD_TYPE_INDIVIDUAL_PUBLIC:
        sdk_cli_ga_date = getattr(issue.fields, config.CUSTOM_FIELD_ID_SDK_CLI_GA_DATE)
        if not sdk_cli_ga_date:
            add_error(errors, issue.key, 'You did not provide an SDK / CLI GA Date. This field is mandatory for public tickets.')

        acknowledge_responsibilities = getattr(issue.fields, config.CUSTOM_FIELD_ID_ACKNOWLEDGE_RESPONSIBILITIES)
        if not acknowledge_responsibilities:
            add_error(errors, issue.key, 'You did not check the box "Acknowledge Responsibilities". Please review the corresponding document and check the box.')

        changelog = getattr(issue.fields, config.CUSTOM_FIELD_ID_CHANGELOG)
        if not changelog:
            add_error(errors, issue.key, 'You did not provide a CHANGELOG entry". This field is mandatory for public tickets.')

        if config.BYPASS_CHECK_INTERNAL_TICKET_LABEL not in issue.fields.labels:
            if check_for_internal_ticket(issue):
                add_error(errors, issue.key, 'Public ticket failed as the preview ticket linked to it is marked as "internal", either in the summary, feature ID(s), and/or labels fields of the preview ticket. Internal features should not be added to the public OCI SDKs/CLI, as they\'re not meant to be consumed by external customers. If you have an exceptional case, you can request an exception in the slack channel #oci_public_sdks.')

    version_error = None

    if version:
        print('Checking for unacceptable versions...')
        version_error = shared.version_utils.is_version_not_acceptable(version)

    if version_error:
        add_error(errors, issue.key, version_error)

    if version and artifact_id and getattr(issue.fields, config.CUSTOM_FIELD_ID_GROUP_ID):
        print('Checking fully specified artifact...')
        # if we are fully specifying a spec, make sure it exists in artifactory
        full_version = getattr(issue.fields, config.CUSTOM_FIELD_ID_ARTIFACT_VERSION)
        ARTIFACTORY_URL_FORMAT = 'https://artifactory.oci.oraclecorp.com/libs-release/{group}/{artifact_id}/{version}/{artifact_id}-{version}.jar'
        artifact_url = ARTIFACTORY_URL_FORMAT.format(
            group=getattr(issue.fields, config.CUSTOM_FIELD_ID_GROUP_ID).replace('.', '/'),
            artifact_id=artifact_id,
            version=full_version
        )
        print('artifact_url: "{}"',format(artifact_url))

        snapshot_artifactory_url = None
        # we dont allow -SNAPSHOT versions but we do allow timestamped snapshots, which are in a separate location in artifactory so we need to check there
        # this regex attempts to remove the timestamp from a timed snapshot build, which we need in order to build the artifactory URL
        version_without_timed_snapshot = re.sub(r'-[0-9]{4}[0-1]{1}[0-9]{1}[0-3]{1}[0-9]{1}\.[0-2]{1}[0-9]{1}[0-5]{1}[0-9]{1}[0-5]{1}[0-9]{1}-[0-9]+$', '', full_version)
        print('Version: "{}"'.format(full_version))
        print('Version without timed snapshot: "{}"'.format(version_without_timed_snapshot))
        if version_without_timed_snapshot != full_version:
            SNAPSHOT_ARTIFACTORY_URL_FORMAT = 'https://artifactory.oci.oraclecorp.com/libs-snapshot/{group}/{artifact_id}/{version_without_timed_snapshot}-SNAPSHOT/{artifact_id}-{version}.jar'
            snapshot_artifactory_url = SNAPSHOT_ARTIFACTORY_URL_FORMAT.format(
                group=getattr(issue.fields, config.CUSTOM_FIELD_ID_GROUP_ID).replace('.', '/'),
                artifact_id=artifact_id,
                version=full_version,
                version_without_timed_snapshot=version_without_timed_snapshot
            )
            print('snapshot_artifactory_url: "{}"'.format(snapshot_artifactory_url))

        download_success = False
        try:
            print('attempting to download from artifact URL: {}'.format(artifact_url))
            response = requests.get(artifact_url, verify=False)
            download_success = response.status_code == 200
            print('download_success: {}'.format(download_success))

            if not download_success and snapshot_artifactory_url:
                # if that didn't work, try snapshot
                print('attempting to download from artifact URL: {}'.format(snapshot_artifactory_url))
                response = requests.get(snapshot_artifactory_url, verify=False)
                download_success = response.status_code == 200
                print('download_success: {}'.format(download_success))

            if not download_success:
                if snapshot_artifactory_url:
                    add_error(errors, issue.key,
                        'Failed to download specified artifact from {release_artifactory_url} or {snapshot_artifactory_url}. Please confirm that your group id and artifact id are correct.'.format(
                            release_artifactory_url=artifact_url,
                            snapshot_artifactory_url=snapshot_artifactory_url
                        )
                    )
                else:
                    add_error(errors, issue.key,
                        'Failed to download specified artifact from {release_artifactory_url}. Please confirm that your group id and artifact id are correct.'.format(
                            release_artifactory_url=artifact_url,
                            snapshot_artifactory_url=snapshot_artifactory_url
                        )
                    )
        except requests.exceptions.RequestException as e:
            print('Request error while attempting to verify artifact existence.  Allowing process to continue.' + str(e))

        # only attempt to look inside the artifact if it was downloaded successfully, and we have a relative spec path
        if getattr(issue.fields, config.CUSTOM_FIELD_ID_SPEC_LOCATION_IN_ARTIFACT):
            if download_success:
                if not os.path.exists("temp"):
                    os.makedirs("temp")

                f = tempfile.NamedTemporaryFile(delete=False, prefix="{issue}-{artifact_id}".format(issue=issue.key, artifact_id=artifact_id), suffix=".jar", dir="temp")
                f.write(response.content)
                f.close()
                print("File name: {}".format(f.name))

                with zipfile.ZipFile(f.name, "r") as zip_ref:
                    # extract to temp dir
                    temp_dir = tempfile.mkdtemp()
                    zip_ref.extractall(temp_dir)
                    spec_location = os.path.join(temp_dir, getattr(issue.fields, config.CUSTOM_FIELD_ID_SPEC_LOCATION_IN_ARTIFACT))
                    if not os.path.isfile(spec_location):
                        add_error(errors, issue.key, 'Spec file did not exist inside artifact. Please verify that a file exists inside your artifact at the location specified by "Spec Location in Artifact".')

    # This check causes the OraCache SDK generation to fail as the endpoint format is different. For more details see https://jira.oci.oraclecorp.com/browse/DEX-4083
    # if getattr(issue.fields, config.CUSTOM_FIELD_ID_SERVICE_SUBDOMAIN):
    #    full_domain = '{subdomain}.us-phoenix-1.oraclecloud.com'.format(subdomain=getattr(issue.fields, config.CUSTOM_FIELD_ID_SERVICE_SUBDOMAIN))
    #
    #    try:
    #        socket.gethostbyname(full_domain)
    #    except socket.gaierror:
    #        add_error(errors, issue.key, 'Failed to resolve service domain: {}. Please ensure the subdomain specified in the ticket is correct.'.format(full_domain))

    return errors, warnings, bool(updates)


def convert_issue_to_script_params(issue):
    custom_field_id_to_param_name = {
        config.CUSTOM_FIELD_ID_ARTIFACT_ID: 'artifact_id',
        config.CUSTOM_FIELD_ID_GROUP_ID: 'group_id',
        config.CUSTOM_FIELD_ID_ARTIFACT_VERSION: 'version',
        config.CUSTOM_FIELD_ID_SPEC_FRIENDLY_NAME: 'spec_name',
        config.CUSTOM_FIELD_ID_SPEC_LOCATION_IN_ARTIFACT: 'relative_spec_path',
        config.CUSTOM_FIELD_ID_SERVICE_SUBDOMAIN: 'subdomain'
    }

    params = {}

    for field, param_name in custom_field_id_to_param_name.items():
        value = getattr(issue.fields, field)
        if value:
            params[param_name] = value

    return params


def invoke_add_or_update_spec(jira_client, issue, params, tool_name, spec_generation_type):
    # def add_or_update_spec(artifact_id, group_id, spec_name, relative_spec_path, version, spec_generation_type, add_sub_groups):
    # Example:
    # {
    #     "artifact_id": "maestro-spec",
    #     "endpoint": "https://orchestration.{domain}/20170630",
    #     "group_id": "com.oracle.pic.orchestration",
    #     "relative_spec_path": "api.yaml",
    #     "spec_generation_type": "PREVIEW",
    #     "spec_name": "orchestration",
    #     "version": "0.0.1-SNAPSHOT"
    # }
    #
    # python ./scripts/add_or_update_spec.py \
    # --group-id "$GROUP_ID" \
    # --artifact-id "$ARTIFACT_ID" \
    # --version "$VERSION" \
    # --spec-name "$SPEC_NAME" \
    # --relative-spec-path "$RELATIVE_SPEC_PATH" \
    # --endpoint "$ENDPOINT"

    result = None

    # for now, hard code spec generation type to preview since we are only using this automation for preview SDK / CLI
    params['spec_generation_type'] = spec_generation_type

    result = None

    if tool_name in [config.PYTHON_SDK_NAME, config.CLI_NAME]:
        # This runs for both Python SDK and Python CLI
        from add_or_update_scripts import python_sdk_add_or_update_spec
        # Update Python SDK pom.xml
        sdk_params = dict(params)
        sdk_params.update({
            'pom_location': os.path.join(config.PYTHON_SDK_REPO_RELATIVE_LOCATION, 'pom.xml'),
            'github_whitelist_location': os.path.join(config.PYTHON_SDK_REPO_RELATIVE_LOCATION, 'github.whitelist')
        })
        result = python_sdk_add_or_update_spec.add_or_update_spec(**sdk_params)
        # Note, we cannot return here, since CLI executes both the above block and the one below

    if tool_name == config.CLI_NAME:
        from add_or_update_scripts import cli_add_or_update_spec
        # Update CLI pom.xml
        cli_params = dict(params)
        cli_params.update({
            'pom_location': os.path.join(config.CLI_REPO_RELATIVE_LOCATION, 'services')
        })

        # strip out '--endpoint' because it is not a valid argument for the CLI script
        if 'endpoint' in cli_params:
            del cli_params['endpoint']

        if 'subdomain' in cli_params:
            del cli_params['subdomain']

        # invoke directly through python
        result = cli_add_or_update_spec.add_or_update_spec(**cli_params)

    if tool_name == config.JAVA_SDK_NAME:
        from add_or_update_scripts import java_sdk_add_or_update_spec
        # Update Java SDK pom.xml
        sdk_params = dict(params)
        sdk_params.update({
            'pom_location': os.path.join(config.JAVA_SDK_REPO_RELATIVE_LOCATION, 'pom.xml'),
            'github_whitelist_location': os.path.join(config.JAVA_SDK_REPO_RELATIVE_LOCATION, 'github.whitelist')
        })
        result = java_sdk_add_or_update_spec.add_or_update_spec(**sdk_params)

    if tool_name == config.GO_SDK_NAME:
        from add_or_update_scripts import go_sdk_add_or_update_spec
        # Update Go SDK pom.xml
        sdk_params = dict(params)
        sdk_params.update({
            'pom_location': os.path.join(config.GO_SDK_REPO_RELATIVE_LOCATION, 'pom.xml'),
            'github_whitelist_location': os.path.join(config.GO_SDK_REPO_RELATIVE_LOCATION, 'github.whitelist'),
            'makefile_location': os.path.join(config.GO_SDK_REPO_RELATIVE_LOCATION, 'Makefile')
        })
        result = go_sdk_add_or_update_spec.add_or_update_spec(**sdk_params)

    if tool_name == config.RUBY_SDK_NAME:
        from add_or_update_scripts.ruby_sdk_add_or_update_spec import RubySpecUpdater, RUBY_MODULE_LOCATION, \
            RUBY_POM_FILE_TEMPLATE, RUBY_SPEC_PARAMS_XML_PATH_DICT
        # Update Ruby SDK pom.xml
        sdk_params = dict(params)
        sdk_params.update({
            'pom_location': os.path.join(config.RUBY_SDK_REPO_RELATIVE_LOCATION, 'pom.xml'),
            # Note: github_whitelist isn't used in Ruby for preview
        })
        ruby_spec_updater = RubySpecUpdater(RUBY_MODULE_LOCATION, RUBY_POM_FILE_TEMPLATE, RUBY_SPEC_PARAMS_XML_PATH_DICT)
        result = ruby_spec_updater.add_or_update_spec(**sdk_params)

    if tool_name == config.TEST_DATA_GEN_NAME:
        from add_or_update_scripts.datagen_add_or_update_spec import TestDataGenSpecUpdater, \
            TEST_DATA_GEN_MODULE_LOCATION, TEST_DATA_GEN_POM_FILE_TEMPLATE, TEST_DATA_GEN_SPEC_PARAMS_XML_PATH_DICT
        # Update sdk-client-test-data pom.xml
        sdk_params = dict(params)
        sdk_params.update({
            'pom_location': os.path.join(config.TEST_DATA_GEN_REPO_RELATIVE_LOCATION, 'pom.xml'),
        })
        test_data_gen_spec_updater = TestDataGenSpecUpdater(TEST_DATA_GEN_MODULE_LOCATION,
                                                            TEST_DATA_GEN_POM_FILE_TEMPLATE,
                                                            TEST_DATA_GEN_SPEC_PARAMS_XML_PATH_DICT)
        result = test_data_gen_spec_updater.add_or_update_spec(**sdk_params)

    if tool_name == config.TYPESCRIPT_SDK_NAME:
        from add_or_update_scripts import typescript_sdk_add_or_update_spec
        # Update Typescript SDK pom.xml
        sdk_params = dict(params)
        sdk_params.update({
            'pom_location': os.path.join(config.TYPESCRIPT_SDK_REPO_RELATIVE_LOCATION, 'package_version'),
            'github_whitelist_location': os.path.join(config.TYPESCRIPT_SDK_REPO_RELATIVE_LOCATION, 'github.whitelist')
        })
        result = typescript_sdk_add_or_update_spec.add_or_update_spec(**sdk_params)

    if tool_name in [config.DOTNET_SDK_NAME, config.POWERSHELL_NAME]:
        from add_or_update_scripts import dotnet_sdk_add_or_update_spec
        # Update .NET SDK pom.xml
        sdk_params = dict(params)
        sdk_params.update({
            'pom_location': os.path.join(config.DOTNET_SDK_REPO_RELATIVE_LOCATION, 'pom.xml'),
            'github_whitelist_location': os.path.join(config.DOTNET_SDK_REPO_RELATIVE_LOCATION, 'github.whitelist')
        })
        result = dotnet_sdk_add_or_update_spec.add_or_update_spec(**sdk_params)

    if tool_name == config.POWERSHELL_NAME:
        from add_or_update_scripts import powershell_add_or_update_spec
        # Update .NET SDK pom.xml
        sdk_params = dict(params)
        sdk_params.update({
            'pom_location': os.path.join(config.POWERSHELL_REPO_RELATIVE_LOCATION, 'pom.xml'),
            'github_whitelist_location': os.path.join(config.POWERSHELL_REPO_RELATIVE_LOCATION, 'github.whitelist')
        })
        result = powershell_add_or_update_spec.add_or_update_spec(**sdk_params)

    if tool_name == config.LEGACY_JAVA_SDK_NAME:
        from add_or_update_scripts import legacy_java_sdk_add_or_update_spec
        # Update Legacy Java SDK pom.xml
        sdk_params = dict(params)
        sdk_params.update({
            'pom_location': os.path.join(config.LEGACY_JAVA_SDK_REPO_RELATIVE_LOCATION, 'pom.xml'),
            'github_whitelist_location': os.path.join(config.LEGACY_JAVA_SDK_REPO_RELATIVE_LOCATION, 'github.whitelist')
        })
        result = legacy_java_sdk_add_or_update_spec.add_or_update_spec(**sdk_params)

    if not result:
        raise ValueError("Unknown tool: '{}'".format(tool_name))

    return result


def warn_on_unexpected_changes(issue, params, tool_name, build_id, result, warnings):
    if not result:
        return

    if not result.existing:
        return

    print(result)

    if "subdomain" in result.changed:
        previous_subdomain = result.previous["subdomain"]
        new_subdomain = params["subdomain"]
        add_error(warnings, issue.key, "For the {}, the field '{}' was changed from '{}' to '{}'.".format(tool_name, RESULT_FIELD_TO_JIRA_FIELD_NAME["subdomain"], previous_subdomain, new_subdomain))
    if "relative_spec_path" in result.changed:
        previous_path = result.previous["relative_spec_path"]
        new_path = params["relative_spec_path"]
        add_error(warnings, issue.key, "For the {}, the field '{}' was changed from '{}' to '{}'.".format(tool_name, RESULT_FIELD_TO_JIRA_FIELD_NAME["relative_spec_path"], previous_path, new_path))

    for ignored in result.ignored:
        add_error(warnings, issue.key, "For the {}, the field '{}' was ignored. It cannot be changed using self-service.".format(tool_name, RESULT_FIELD_TO_JIRA_FIELD_NAME[ignored]))


def convert_underscore_name_to_param(name):
    return '--' + name.replace('_', '-')


def pretty_print(input):
    print(json.dumps(input, indent=4, sort_keys=True))


# Note: This uses the globally set BRANCH_TIMESTAMP; therefore, all branches created using one
#       execution of this script will use the same timestamp
def generate_time_stamped_branch_name(prefix):
    return '{prefix}-{timestamp}'.format(prefix=prefix, timestamp=BRANCH_TIMESTAMP)


def checkout_sdk_and_cli_branches(base_branch, tool_name):
    for repo in config.REPOS_FOR_TOOL[tool_name]:
        repo.git.checkout(base_branch)


# Branch suffix is "DEXREQ-673-2019-08-16-20-58-17"
def push_spec_baseline(spec_dir, build_type, issue, branch_suffix):
    branch_prefix = util.get_branch_prefix_for_spec_diff(build_type)
    branch_name = "{}-{}".format(branch_prefix, branch_suffix)

    repo = config.DEXREQ_REPO
    git = repo.git
    git.checkout(B=branch_name)

    for filename in os.listdir(spec_dir):
        source = os.path.join(spec_dir, filename)
        destination = os.path.join(config.DEXREQ_DIFF_REPO_RELATIVE_LOCATION, filename)
        print("Copying {} -> {}".format(source, destination))
        shutil.rmtree(destination, True)
        shutil.copytree(source, destination, ignore=shutil.ignore_patterns('*.lineNumberMapping'))

    git.add(A=True)

    commit_message = '{commit_prefix} [[{issue_key}]]: {issue_summary}'.format(
        commit_prefix=config.SPEC_BASELINE_COMMIT_MESSAGE_PREFIX,
        issue_key=issue.key,
        issue_summary=issue.fields.summary
    )

    message = commit_message
    if 'nothing to commit' in git.status():
        message = "{} (no change)".format(message)
    print(message)
    git.commit("-m", message, "--allow-empty")
    if config.IS_DRY_RUN:
        print('DRY-RUN: not pushing to branch {}'.format(branch_name))
    else:
        git.push('-u','origin','HEAD')


# default is used when the attribute is missing
# noneDefault is used when the attribute exists but returns None
def getAttr(o, attribute, default=None, noneDefault=None):
    Attr = getattr(o, attribute, default)
    return noneDefault if Attr is None else Attr


#  This script requires two inputs:
#   - build_id: The team city build id for the build that is running this script.  This is used to update the relevant JIRA tickets with links to the team city build
#   - build_type: Value can be sneak_peek or preview (default will be preview)
#
#  For preview builds, we will aggregate all open preview tickets and include them in the same build
#  For sneak_peek builds, we will aggregate all open sneak peek tickets and include them in the same build
#
#  TODO: It likely isn't an issue that any given sneak peek may contain extra changes from other tickets but we may update this in the future to generate individual
#  builds with *only* those changes for each sneak peek ticket
#
#  The script assumes it is being invoked from the root directory of the python-cli repository and that there is a sibling directory 'python-sdk'
#  with the Python SDK in it
#
if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Process the DEXREQ JIRA queue (preview and public).')
    parser.add_argument('--build-id',
                        required=True,
                        help='The team city build id for the build that is running this script.  This is used to update the relevant JIRA tickets with links to the team city build')
    parser.add_argument('--build-type',
                        default=config.BUILD_TYPE_INDIVIDUAL_PREVIEW,
                        help='The build type to use, can be one of the following: {}'.format(', '.join(config.VALID_BUILD_TYPES)))
    parser.add_argument('--base-branch',
                        default='preview',
                        help='The base branch to start from')
    parser.add_argument('--tool',
                        default=config.CLI_NAME,
                        help='The tool for which to generate the preview. Accepted values: {}'.format(', '.join(config.TOOL_NAMES + [TOOL_ARGUMENT_ALL])))
    parser.add_argument('--dry-run',
                        default=False,
                        action='store_true',
                        help='Perform a dry-run')
    parser.add_argument('--allow-individual-tool-generation',
                        default=False,
                        action='store_true',
                        help=textwrap.dedent("""\
                            By default, we require ALL SDK / CLI statuses to be ready before running automation.  For one-off previews this
                            means all statuses must be in "To Do" and for bulk previews all statuses must be in "Pending Merge".  This flag allows
                            generating previews for all tools that are ready, and ignoring those which are not."""))
    parser.add_argument('--issue',
                        action='append',
                        help='By default, we query JIRA. This allows you to specify a DEXREQ issue to process instead: --issue DEXREQ-123')
    parser.add_argument('--for-any-status',
                        default=False,
                        action='store_true',
                        help='Ignore status fields (don\'t require them to be in "To Do"). Can only be used together with --issue.')
    parser.add_argument('--push-spec-baseline',
                        help='Push the baseline spec (after pre-processing, before generation) from the specified directory into a branch in the SDK/dexreq repo.')
    parser.add_argument('--verbose',
                        default=False,
                        action='store_true',
                        help='Verbose logging')

    args = parser.parse_args()

    util.IS_VERBOSE = args.verbose
    build_id = args.build_id
    build_type = args.build_type
    base_branch = args.base_branch
    allow_individual_tool_generation = args.allow_individual_tool_generation
    tool_name = args.tool
    config.IS_DRY_RUN = args.dry_run
    issues_filter = args.issue
    if args.for_any_status and not issues_filter:
        raise ValueError("--for-any-status can only be used with --issue")
    for_any_status = args.for_any_status

    if base_branch.lower() == "preview" and build_type in config.PUBLIC_BUILD_TYPES:
        raise ValueError("Used base branch '{}' with build type '{}'".format(base_branch, build_type))
    if base_branch.lower() == "master" and build_type in config.PREVIEW_BUILD_TYPES:
        raise ValueError("Used base branch '{}' with build type '{}'".format(base_branch, build_type))

    if tool_name not in config.TOOL_NAMES and not tool_name == TOOL_ARGUMENT_ALL:
        raise ValueError("Tool name must be one of: {}".format(', '.join(config.TOOL_NAMES + [TOOL_ARGUMENT_ALL])))

    if args.push_spec_baseline and (tool_name not in [TOOL_ARGUMENT_ALL, config.GO_SDK_NAME]):
        raise ValueError("--push-spec-baseline can only be used with --tool ALL or --tool GoSDK")

    tools_to_run = config.TOOL_NAMES if tool_name == TOOL_ARGUMENT_ALL else [tool_name]
    try:
        if build_type == config.BUILD_TYPE_BULK_PENDING_MERGE_PUBLIC:
            query = config.READY_FOR_PUBLIC_ANY_TOOLS_REQUESTED \
                if allow_individual_tool_generation else config.READY_FOR_PUBLIC_ALL_TOOLS
            generate_func = lambda all_issues: generate_bulk_pending_merge(config.PUBLIC_SPEC_GENERATION_TYPE,  # noqa:E731
                                                                           all_issues,
                                                                           build_id,
                                                                           base_branch,
                                                                           tool_name,
                                                                           allow_individual_tool_generation,
                                                                           for_any_status)
            processing_status = config.STATUS_PROCESSING_BULK
        elif build_type == config.BUILD_TYPE_BULK_PENDING_MERGE_PREVIEW:
            # in order to allow re-running bulk preview for a single tool we need to get all tickets in 'Ready for Preview' with any tool status set to 'To Do'
            # inside the generation logic, we will only build for tools that are explicitly set to 'To Do'
            query = config.READY_FOR_PREVIEW_ANY_TOOLS_REQUESTED \
                if allow_individual_tool_generation else config.READY_FOR_PREVIEW_ALL_TOOLS
            generate_func = lambda all_issues: generate_bulk_pending_merge(config.PREVIEW_SPEC_GENERATION_TYPE,  # noqa:E731
                                                                           all_issues,
                                                                           build_id,
                                                                           base_branch,
                                                                           tool_name,
                                                                           allow_individual_tool_generation,
                                                                           for_any_status)
            processing_status = config.STATUS_PROCESSING_BULK
        elif build_type == config.BUILD_TYPE_INDIVIDUAL_PREVIEW:
            query = config.TODO_PREVIEW_ALL_TOOLS
            generate_func = lambda all_issues: generate_individual_to_do_requests(all_issues,    # noqa:E731
                                                                                  build_id,
                                                                                  base_branch,
                                                                                  tool_name,
                                                                                  allow_individual_tool_generation,
                                                                                  build_type,
                                                                                  for_any_status)
            processing_status = config.STATUS_PROCESSING
        elif build_type == config.BUILD_TYPE_INDIVIDUAL_PUBLIC:
            query = config.TODO_PUBLIC_ALL_TOOLS
            generate_func = lambda all_issues: generate_individual_to_do_requests(all_issues,  # noqa:E731
                                                                                  build_id,
                                                                                  base_branch,
                                                                                  tool_name,
                                                                                  allow_individual_tool_generation,
                                                                                  build_type,
                                                                                  for_any_status)
            processing_status = config.STATUS_PROCESSING
        else:
            raise ValueError('Build type must be one of: {}'.format(', '.join(config.VALID_BUILD_TYPES)))

        query = limit_query_to_issue_keys(query, issues_filter, for_any_status)
        print("query = {}".format(query))

        dexreq_public_errors = {}
        dexreq_public_warnings = {}

        # TODO: actual loading of public ticket data from preview ticket not yet enabled
        all_issues = util.search_dexreq_issues(query, fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS_PER_BUILD_TYPE[build_type]),
            errors=dexreq_public_errors, warnings=dexreq_public_warnings)

        # sort issues so that latest version is applied last if there are multiple tickets for the same spec
        # if no version is set, default to 0 (it will not succeed without a version so it doesn't matter when it happens)
        # this is only necessary for bulk preview, but doesn't hurt to do for other build types
        all_issues.sort(key=lambda x: getAttr(x.fields, config.CUSTOM_FIELD_ID_ARTIFACT_VERSION, noneDefault="0"))

        print('Found the following issues to generate {} builds for:'.format(build_type))
        issues_to_be_processed = []
        for issue in all_issues:
            if config.should_ignore_issue(issue.key):
                print('{} - {} - being ignored per env var {}'.format(issue.key, issue.fields.summary, config.DEXREQ_IGNORED_ISSUES_ENV_VAR_NAME))
                continue

            # set all issues to processing now, if we do this as we generate for each tool it could overwrite meaningful statuses
            # for previous tools like 'DEX Support Required' or 'Service Team Failure Investigation'
            util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, processing_status)

            # Reload after the transition
            issue = util.get_dexreq_issue(issue.key,
                fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS_PER_BUILD_TYPE[build_type]))

            print('{} - {}'.format(issue.key, issue.fields.summary))

            validation_errors, validation_warnings, changed = validate_issue(issue, build_type, dexreq_public_errors, dexreq_public_warnings)
            if not report_errors_and_warnings(issue, validation_errors, validation_warnings, tools_to_run):
                if changed:
                    # Reload after making changes (but only reload if the issue is to be processed; if not, it's ok if the issue is out of date; we're never referring to it again)
                    issue = util.get_dexreq_issue(issue.key,
                        fields=(DEFAULT_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS + CUSTOM_JIRA_ISSUE_FIELDS_PER_BUILD_TYPE[build_type]))

                issues_to_be_processed.append(issue)

        generation_warnings = {}  # key: issue_key, value: list
        generation_errors = {}  # key: issue_key, value: list
        for tool_name in tools_to_run:
            checkout_sdk_and_cli_branches(base_branch, tool_name)
            per_sdk_generation_warnings, per_sdk_generation_errors, branches = generate_func(issues_to_be_processed)
            generation_warnings.update(per_sdk_generation_warnings)
            generation_errors.update(per_sdk_generation_errors)

            is_individual_build = build_type in [config.BUILD_TYPE_INDIVIDUAL_PREVIEW, config.BUILD_TYPE_INDIVIDUAL_PUBLIC]
            is_go_build = (tool_name == config.GO_SDK_NAME)
            if args.push_spec_baseline and is_individual_build and is_go_build:
                for issue in issues_to_be_processed:
                    print("Branches: {}".format(branches))
                    if issue.key in branches:
                        m = re.search("^.*{}-(.*)$".format(config.GO_SDK_NAME), branches[issue.key])
                        if m:
                            branch_suffix = m.group(1)
                            push_spec_baseline(args.push_spec_baseline, build_type, issue, branch_suffix)
                    else:
                        print("Could not push spec baseline, there was no {} branch for {}, probably (probably 'Skipping generating tool: GoSDK' above)".format(config.GO_SDK_NAME, issue.key))

        # these are warnings that are SDK specific
        # examples: "For the GoSDK, the field 'Spec Artifact Group Id' was ignored. It cannot be changed using self-service."
        if generation_warnings:
            for issue_key in generation_warnings:
                # comment on issue
                util.add_jira_comment(
                    issue_key,
                    """Please check that the following settings indeed reflect what you wanted to do:

                    {warnings}

                    The full build log can be found {build_log_link}.

                    For TeamCity access see the [self-service FAQ|https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=132774928#SDK/CLISelfServiceFrequentlyAskedQuestions-HowdoIgetaccesstoTeamCitylogsandartifacts?].

                    There is no need to re-start this ticket, unless you determine that some settings were incorrect.""".format(
                        warnings='\n'.join(['*- {}*'.format(warning) for warning in generation_warnings[issue_key]]),
                        build_log_link=build_log_link(build_id)
                    ),
                    comment_type=config.COMMENT_TYPE_INFO
                )

        # these are errors that are SDK specific
        # examples: Missing required parameter for given SDK add_or_update script, unexpected exception thrown in add_or_update script
        if generation_errors:
            for issue_key in generation_errors:
                for generation_error in generation_errors[issue_key]:
                    util.add_jira_comment(
                        issue_key,
                        generation_error,
                        comment_type=config.COMMENT_TYPE_ERROR
                    )

    except Exception as e:  # noqa:F841
        # TODO: report error to JIRA tasks
        raise
