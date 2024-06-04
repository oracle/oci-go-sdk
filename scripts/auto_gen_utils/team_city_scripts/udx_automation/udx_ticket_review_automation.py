import os
import sys
import argparse
import logging
import datetime

dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../..'))

import util  # noqa: ignore=F402
import config  # noqa: ignore=F402


JIRA_AUTOMATION_BOT = 'jira-automation-bot'
UDX_AUTOMATION_MAINTAINER = os.environ.get('UDX_AUTOMATION_MAINTAINER') or 'anurggar'
ALLOWED_GA_WEEKDAYS = ['tuesday', 'wednesday']
BYPASS_UDX_AUTOMATION_LABEL = 'Bypass-UDX-Automation'
MANUAL_CHECK_LABEL = 'UDX-Manually-Created'
EXEMPTIONS_NEEDED_WHEN_NO_PUBLIC_API_CHANGES = ['SDK/CLI', 'Terraform', 'UX/Console', 'Docs']
UDX_AUTOMATION_LABEL = 'UDX-Intake-Automation'

IN_DESIGN_COMMENT = \
    '''
    [~{}],
    We have moved your ticket to *In Design* phase, however, please complete the steps mentioned below in order to proceed:-

    {}
    '''

GA_DATE_NOT_IN_ALLOWED_DAYS_COMMENT = \
    f'* The GA Date for this ticket is not a {"/".join(ALLOWED_GA_WEEKDAYS)}, if SDK/CLI or Terraform are impacted ' \
    f'you must *update the target GA date* on the top level ORM to these days in order to align with ' \
    f'the dev tools release schedule.'


GA_DATE_NOT_A_WEEKDAY_COMMENT = \
    f"* You must *update the target GA date* on the top level ORM to a week day in order to align with the dev tools release schedule"

VPAT_DOCUMENTATION_COMMENT = \
    f'* Please answer the VPAT related questions asked in the UDX ticket description. As per UDX updated GA requirement feature teams must provide the VPAT at the time of GA. Refer the [UDX Self Service On boarding|https://confluence.oci.oraclecorp.com/display/UDX/UDX+Self-Service+Onboarding+-+VPAT+Documentation] for additional details.'

EXCEPTION_OCCURRED_COMMENT = \
    '''
    [~{}],
    An error occurred in the UDX intake automation while processing this ticket. Please check the builds logs to rectify this issue.
    '''

MISSING_ROOT_ORM_COMMENT = \
    '''
    [~{}],
    Have you created an ORM ticket for this? UDX tickets should only be generated through the ORM process. Manually generated and cloned UDX tickets can't be processed and will be cancelled unless granted exception by the UDX team. See this wiki ([https://confluence.oci.oraclecorp.com/display/OCIRM/Getting+Started+with+Release+Management]) for info on creating an ORM ticket which will in turn create the UDX ticket
    '''

NO_CUSTOMER_FACING_CHANGES_ROOT_ORM_COMMENT = \
    '''
    [~{}],
    Have you created an ORM ticket with *Customer facing = No* ?, UDX tickets should only be generated through the ORM process with *Customer facing= Yes*. This UDX tickets can't be processed and will be cancelled. See this wiki ([https://confluence.oci.oraclecorp.com/display/OCIRM/Getting+Started+with+Release+Management]) for info on creating an ORM ticket which will in turn create the UDX ticket
    '''

MISSING_EXCEPTIONS_COMMENT = \
    '''
    [~{}],
    This ticket was submitted with "No API change to Public-Facing Endpoints" selected but no UDX Surfaces have been marked as Exempt.

    Please clarify which of these customer facing surfaces are impacted by this change:
    * Console UI
    * SDK/CLI
    * Terraform
    * Technical Content (Docs)

    Let us know which surfaces are impacted ASAP so that we can Un-Block this ticket and move it to the next step.
    In the future you can mark the UDX surfaces affected as described in the second bullet of the UDX Feature Release Checklist ([https://confluence.oci.oraclecorp.com/display/UDX/UDX+Feature+Release+Checklist])
    '''

TICKETS_IN_TRIAGE_QUERY = f'project = UDX AND status in ("Needs Triage") AND labels not in ({BYPASS_UDX_AUTOMATION_LABEL}, {MANUAL_CHECK_LABEL})' \
                          f'ORDER BY status ASC, cf[12197] ASC, cf[12196] DESC, cf[11140] ASC, key ASC, reporter ' \
                          f'DESC, updated DESC, priority DESC'

# QUERY_TEMPLATE = 'issue = UDX-21538'
QUERY_TEMPLATE = os.environ.get('QUERY') or TICKETS_IN_TRIAGE_QUERY

TICKETS_IN_MORE_INFO = f'project = "User and Developer Experience" AND status = "More Information Needed" AND labels not in ({BYPASS_UDX_AUTOMATION_LABEL}, {MANUAL_CHECK_LABEL}) AND labels in ({UDX_AUTOMATION_LABEL})'

# create logger
logging.basicConfig()
logger = logging.getLogger('UDX_INTAKE_AUTOMATION')
logger.setLevel(logging.DEBUG)


def is_root_orm_customer_facing(root_orm):
    # customer facing field in root has cf_10720
    customer_facing = root_orm.fields.customfield_10720
    logger.info(f'The root orm ticket: {root_orm.key}, customer facing field: {customer_facing}')
    if customer_facing and 'Yes' in customer_facing.value:
        return True
    return False


def close_not_customer_facing_ticket(udx_ticket, root_orm):
    logger.info(f'Transitioning ticket {udx_ticket.key} to {config.STATUS_CLOSED} as root orm: {root_orm.key}'
                f' mentions there is no customer facing changes')
    if not config.IS_DRY_RUN:
        add_udx_automation_label(udx_ticket)
        ticket_creator = udx_ticket.fields.reporter.name if udx_ticket.fields.reporter else udx_ticket.fields.assignee.name
        util.JIRA_CLIENT().transition_issue(udx_ticket, config.STATUS_CLOSED, comment=NO_CUSTOMER_FACING_CHANGES_ROOT_ORM_COMMENT.format(ticket_creator, root_orm.key))


def update_manually_created_ticket(udx_ticket):
    logger.info(f'Updating ticket {udx_ticket.key} with missing Root ORM comment and adding manual check label')
    if not config.IS_DRY_RUN:
        add_udx_automation_label(udx_ticket)
        add_udx_manually_created_label(udx_ticket)
        # Add comment that root ORM is Missing.
        ticket_creator = udx_ticket.fields.reporter.name if udx_ticket.fields.reporter else udx_ticket.fields.assignee.name
        util.add_jira_comment(udx_ticket.key,comment=MISSING_ROOT_ORM_COMMENT.format(ticket_creator))


def transition_ticket_missing_exemptions_to_more_info_needed(udx_ticket):
    logger.info(f'Transitioning ticket: {udx_ticket.key} to status: {config.STATUS_MORE_INFORMATION_NEEDED} as ticket is missing exemptions')
    if not config.IS_DRY_RUN:
        add_udx_automation_label(udx_ticket)
        util.JIRA_CLIENT().transition_issue(udx_ticket, config.STATUS_MORE_INFORMATION_NEEDED)
        # Adding comment separately as the transition to More Information needed already has an auto-comment which
        # overrides the comment we provide in transition issue api
        util.add_jira_comment(udx_ticket.key, comment=MISSING_EXCEPTIONS_COMMENT.format(udx_ticket.fields.reporter.name))


def get_parent_tickets(udx_ticket):
    root_orm, wrapper_udx = None, None
    wrapper_udx_link_generator = (
        link.inwardIssue for link in udx_ticket.fields.issuelinks
        if hasattr(link, 'inwardIssue') and 'ORM' in link.inwardIssue.key
    )
    wrapper_udx = next(wrapper_udx_link_generator, None)
    if wrapper_udx:
        wrapper_udx = util.JIRA_CLIENT().issue(wrapper_udx.key)
        logger.debug(f'For ticket: {udx_ticket.key}, found wrapper udx ticket: {wrapper_udx.key}')
        root_orm_udx_link_generator = (
            link.outwardIssue for link in wrapper_udx.fields.issuelinks
            if hasattr(link, 'outwardIssue') and 'ORM' in link.outwardIssue.key
        )
        root_orm = next(root_orm_udx_link_generator, None)
        if root_orm:
            logger.debug(f'For wrapper UDX ticket: {wrapper_udx.key}, found Root ORM ticket: {root_orm.key}')
            root_orm = util.JIRA_CLIENT().issue(root_orm.key)
        else:
            logger.info(f'Root ORM ticket not found under wrapper UDX: {wrapper_udx.key}, setting wrapper udx as root ORM')
            root_orm = wrapper_udx
    else:
        logger.error(f'For ticket: {udx_ticket.key} is missing Wrapper UDX ORM ticket')
    return root_orm, wrapper_udx


def update_udx_service_queue(udx_ticket, root_orm):
    udx_ticket_service_queue = udx_ticket.fields.customfield_13254
    if udx_ticket_service_queue is None:
        logger.debug(f'For ticket: {udx_ticket.key}, Service Queue is missing.')
        root_orm_service_queue = root_orm.fields.customfield_13254
        if root_orm_service_queue:
            logger.debug(f'For ticket: {udx_ticket.key}, Adding Service Queue: {root_orm_service_queue.name} from '
                         f'root ORM ticket: {root_orm.key}')
            if not config.IS_DRY_RUN:
                update_service_queue_dict = {
                    'customfield_13254': {
                        'key': root_orm_service_queue.key
                    }
                }
                udx_ticket.update(fields=update_service_queue_dict)
        else:
            logger.warning(f'Service Queue not found in either UDX ticket: {udx_ticket.key}, or its root ORM '
                           f'ticket: {root_orm.key}')


def update_udx_reporter_if_missing(udx_ticket, root_orm):
    if udx_ticket.fields.reporter and udx_ticket.fields.reporter.name == JIRA_AUTOMATION_BOT:
        reporter = root_orm.fields.reporter
        logger.debug(f'For ticket: {udx_ticket.key}, Reporter is missing. Adding reporter: {reporter.displayName} with '
                     f'user name: {reporter.name} from root ORM ticket: {root_orm.key}')

        if not config.IS_DRY_RUN:
            update_reporter_dict = {
                'reporter': {
                    'name': reporter.name
                }
            }
            udx_ticket.update(fields=update_reporter_dict)


def mark_udx_ticket_with_no_exemptions(udx_ticket):
    logger.debug(f'For ticket: {udx_ticket.key}, Setting Surfaces Exempt from Feature Impact = NO EXEMPTIONS')
    if not config.IS_DRY_RUN:
        update_surface_exemption_dict = {
            'customfield_13594': [{'id': '16504'}]
        }
        udx_ticket.update(fields=update_surface_exemption_dict)


def udx_ticket_has_public_api_changes(udx_ticket, root_orm):
    has_public_api_changes = root_orm.fields.customfield_13419
    if has_public_api_changes and has_public_api_changes.value and 'yes' in has_public_api_changes.value.lower():
        logger.debug(f'root orm ticket: {root_orm.key} has public facing API changes')
        logger.debug(f'Setting has public API changes for ticket: {udx_ticket.key} to "Yes"')
        if not config.IS_DRY_RUN:
            api_changes_update_dict = {
                'customfield_13419': {'id': '15842'}
            }
            udx_ticket.update(fields=api_changes_update_dict)
        return True
    else:
        logger.debug(f'Root ORM ticket: {root_orm.key} does not have public facing API changes')
        logger.debug(f'Setting has public API changes for ticket: {udx_ticket.key} to "No"')
        if not config.IS_DRY_RUN:
            api_changes_update_dict = {
                'customfield_13419': {'id': '15843'}
            }
            udx_ticket.update(fields=api_changes_update_dict)
        return False


def udx_ticket_missing_exemptions(udx_ticket):
    surface_exemptions = udx_ticket.fields.customfield_13594
    if surface_exemptions:
        logger.debug(f'For udx ticket: {udx_ticket.key}, the following surfaces are marked for exemption: {[e.value for e in surface_exemptions]}')
        # If no exemptions return True
        if any(e.value == "No Exemptions" for e in surface_exemptions):
            return True
        # If SDK/CLI or Terraform exemptions are present
        if any(e.value in EXEMPTIONS_NEEDED_WHEN_NO_PUBLIC_API_CHANGES for e in surface_exemptions):
            return False
    return True


def get_tc_ticket_linked_to_udx(udx_ticket):
    tc_link_generator = (
        link.outwardIssue for link in udx_ticket.fields.issuelinks
        if hasattr(link, 'outwardIssue') and 'TC' in link.outwardIssue.key
    )
    return next(tc_link_generator, None)


def get_tc_ticket_from_root_orm(root_orm):
    wrapper_tc_link_generator = (
        link.inwardIssue for link in root_orm.fields.issuelinks
        if hasattr(link, 'inwardIssue') and hasattr(link.inwardIssue, 'fields') and 'Technical Content' in link.inwardIssue.fields.summary
    )
    wrapper_tc = next(wrapper_tc_link_generator, None)
    if wrapper_tc:
        # Get Fill Wrapper TC ticket via its key
        wrapper_tc = util.JIRA_CLIENT().issue(wrapper_tc.key)
        logger.debug(f'Found wrapper tc ticket: {wrapper_tc.key} via root orm: {root_orm.key}')

        # Find the TC ticket linked to it via its outward issue links
        tc_link_generator = (
            link.outwardIssue for link in wrapper_tc.fields.issuelinks
            if hasattr(link, 'outwardIssue') and 'TC' in link.outwardIssue.key
        )
        return next(tc_link_generator, None)
    else:
        logger.debug(f'For root ORM: {root_orm.key}, No wrapper TC ticket found!')


def link_tc_ticket_to_udx(udx_ticket, root_orm):
    tc_ticket = get_tc_ticket_linked_to_udx(udx_ticket)
    if tc_ticket is None:
        logger.debug(f'No TC ticket linked to udx ticket: {udx_ticket.key}, checking in root orm: {root_orm.key}')
        tc_ticket = get_tc_ticket_from_root_orm(root_orm)
        if tc_ticket is None:
            logger.info(f'No TC ticket found in root orm: {root_orm.key}')
        else:
            logger.info(f'Found TC ticket: {tc_ticket.key} via root orm: {root_orm.key}')
            if not config.IS_DRY_RUN:
                util.JIRA_CLIENT().create_issue_link(type='Required', inwardIssue=udx_ticket.key, outwardIssue=tc_ticket.key)
    else:
        logger.info(f'TC ticket: {tc_ticket.key} already linked to udx ticket: {udx_ticket.key}')


def is_ga_date_tuesday_or_wednesday(udx_ticket):
    ga_date = udx_ticket.fields.customfield_11140
    logger.debug(f'For UDX ticket:{udx_ticket.key}, found GA date: {ga_date}')
    year, month, day = (int(x) for x in ga_date.split('-'))
    day_name = datetime.date(year, month, day).strftime("%A")
    if day_name.lower() in ALLOWED_GA_WEEKDAYS:
        logger.debug(f'The {ga_date} falls on allowed GA days {ALLOWED_GA_WEEKDAYS}')
        return True
    logger.debug(f'The {ga_date} does not fall on allowed GA days {ALLOWED_GA_WEEKDAYS}')
    return False


def is_ga_date_a_weekday(udx_ticket):
    ga_date = udx_ticket.fields.customfield_11140
    logger.debug(f'For UDX ticket:{udx_ticket.key}, found GA date: {ga_date}')
    ga_date = datetime.datetime.strptime(ga_date, '%Y-%m-%d').date()
    if ga_date.weekday() < 5:
        return True
    else:
        return False


def add_label(udx_ticket, label):
    if not config.IS_DRY_RUN:
        labels = udx_ticket.fields.labels or []
        labels.append(label)
        udx_ticket.update(fields={"labels": labels})


def add_udx_automation_label(udx_ticket):
    logger.debug(f'Adding label: {UDX_AUTOMATION_LABEL}, to ticket: {udx_ticket.key}')
    add_label(udx_ticket, UDX_AUTOMATION_LABEL)


def add_udx_manually_created_label(udx_ticket):
    logger.debug(f'Adding label: {MANUAL_CHECK_LABEL}, to ticket: {udx_ticket.key}')
    add_label(udx_ticket, MANUAL_CHECK_LABEL)


def intake_ticket(udx_ticket):
    try:
        logger.debug(f'Processing ticket: {udx_ticket.key}')
        root_orm, wrapper_udx = get_parent_tickets(udx_ticket)

        # 1. Close any ticket which are manually created
        if root_orm is None:
            update_manually_created_ticket(udx_ticket)
            return

        # 2.Close any tickets which don't have any Customer facing changes
        if not is_root_orm_customer_facing(root_orm):
            close_not_customer_facing_ticket(udx_ticket, root_orm)
            return

        # Start Intake Process
        # 3. If the UDX reporter is Unassigned then set it to the root ORM ticket reporter
        update_udx_reporter_if_missing(udx_ticket, root_orm)

        # 4. Update the UDX ticket Service Queue to match the value in Root ORM ticket
        update_udx_service_queue(udx_ticket, root_orm)

        # 5 Link TC tickets to UDX
        link_tc_ticket_to_udx(udx_ticket, root_orm)

        additional_in_design_comment = [VPAT_DOCUMENTATION_COMMENT]

        # 6 Check if the exemptions are present on the UDX ticket if there are no Public API changes
        if udx_ticket_has_public_api_changes(udx_ticket, root_orm):
            # 6.a If ticket has public facing api changes, set UDX Surfaces Exempt from Feature Impact = NO EXEMPTIONS
            mark_udx_ticket_with_no_exemptions(udx_ticket)

            # 6.a.1 Add comment if GA date is not a Tuesday/Wednesday
            if not is_ga_date_tuesday_or_wednesday(udx_ticket):
                additional_in_design_comment.append(GA_DATE_NOT_IN_ALLOWED_DAYS_COMMENT)
        else:
            # 6.b For ticket with no public facing API changes
            # 6.b.1 If no exemptions or no SDK/CLI, Terraform exemptions then transition to More information needed
            if udx_ticket_missing_exemptions(udx_ticket):
                transition_ticket_missing_exemptions_to_more_info_needed(udx_ticket)
                return
            else:
                # 6.b.2 If it has SDK/CLI, Terraform exemptions, check if GA date is a weekday
                if not is_ga_date_a_weekday(udx_ticket):
                    additional_in_design_comment.append(GA_DATE_NOT_A_WEEKDAY_COMMENT)

        # 7 Move the ticket to In-design
        if not config.IS_DRY_RUN:
            util.JIRA_CLIENT().transition_issue(udx_ticket, config.STATUS_IN_DESIGN)
            add_udx_automation_label(udx_ticket)
            util.add_jira_comment(udx_ticket.key, comment=IN_DESIGN_COMMENT.format(udx_ticket.fields.reporter.name, '\n'.join(additional_in_design_comment)))

    except Exception as e:
        logger.error(f'Failed to process ticket: {udx_ticket} due to exception: {e}')
        util.add_jira_comment(udx_ticket.key, EXCEPTION_OCCURRED_COMMENT.format(UDX_AUTOMATION_MAINTAINER), config.COMMENT_TYPE_INFO)


def process_more_info_needed_tickets(udx_ticket):
    if not udx_ticket_missing_exemptions(udx_ticket):
        logger.info(f'The ticket: {udx_ticket.key} has surfaces marked for exemptions')

        additional_in_design_comment = [VPAT_DOCUMENTATION_COMMENT]

        # Check if release date is weekday else add additional comment
        if not is_ga_date_a_weekday(udx_ticket):
            additional_in_design_comment.append(GA_DATE_NOT_A_WEEKDAY_COMMENT)

        if not config.IS_DRY_RUN:
            util.JIRA_CLIENT().transition_issue(udx_ticket, config.STATUS_NEEDS_TRIAGE)
            util.JIRA_CLIENT().transition_issue(udx_ticket, config.STATUS_IN_DESIGN)
            util.add_jira_comment(udx_ticket.key, comment=IN_DESIGN_COMMENT.format(udx_ticket.fields.reporter.name, '\n'.join(additional_in_design_comment)))


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Udx Intake Automation')
    parser.add_argument('--dry-run',
                        default=False,
                        action='store_true',
                        help='Perform a dry-run')

    args = parser.parse_args()
    config.IS_DRY_RUN = args.dry_run

    if config.IS_DRY_RUN:
        logger.info('Running in dry-run mode')

    # Process tickets in More Information needed
    tickets_in_more_info_needed = util.jira_search_issues(TICKETS_IN_MORE_INFO)
    if len(tickets_in_more_info_needed) == 0:
        logger.info(f'No ticket found in More Info needed state')
    else:
        for ticket in tickets_in_more_info_needed:
            process_more_info_needed_tickets(ticket)

    # Find all tickets to API Review ticket readiness check
    tickets_in_triage = util.jira_search_issues(QUERY_TEMPLATE)
    if len(tickets_in_triage) == 0:
        logger.info(f'No actionable tickets found for Query: {QUERY_TEMPLATE}')
    else:
        logger.info(f'Total Tickets found: {len(tickets_in_triage)}')
        for ticket in tickets_in_triage:
            intake_ticket(ticket)
