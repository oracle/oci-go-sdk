#API Review Ticket Readiness Validation Check

## Introduction
The API Review Ticket Readiness Validation Check automation aims to reduce some of the manual load on the API Review Board
TPM's by doing an initial checks on an API Review Ticket raised for review.

The automation currrently find all tickets that are in `Needs triage` state and runs the following checks:-
1. There should be a UDX ticket linked to the ticket via the issue links field.
2. The API review PR link to review the spec changes is present.
3. The API Review PR is in `OPEN` state and does not have any merge conflicts.
4. The API Review PR has been peer-reviewed by at-least 1 member of the team requesting a review.
5. The API Review PR has a valid spec validation output file included in the PR.
6. The API Review PR should have at least one file with `cond.yaml` in it's name in the PR to make sure the requestor has a spec file to review attached.
7. The Spec Validation Output included should use a Spec Validator tool version that is at most 3 minor versions behind the latest version.
8. The Spec Validation Output shouldn't have any `Errors` or `Warnings`.
9. If there is an increase in either the `Supresses errors` or `Supressed warnings`, then, a corresponding spec validation config yaml file should also be present.


### Automation Check Fail Scenario
If the ticket fails any of the above checks we assign the ticket back to reporter and move its status to `More Information Needed`
and add a comment for the reporter with all issues detected for them to review and fix.

### Automation Check Pass Scenario
If the ticket passes all the above checks, then the automation assigns it to the API Review board TPM for triaging. 
The automation also adds a comment on the ticket to specify all the checks done that can be referenced by the TPM to cross check the results.

PS: The automation adds a label `APIRB-Ready-Check-Pass` to all tickets that pass the check to make sure it doesn't re-pick
already tested ticket again.

#### Some Caveats for Pass scenario
- As with all automation, there can be some scenarios where the tool doesn't work as intended. As such, the TPM, or the 
  requestor may request to bypass the automation check in such scenarios by adding the label `ByPass-APIRB-Ready-Check`
  to the ticket.
- There might be cases where the automation tool may not be able to do some checks. In such scenarios, the automation tool
  will add the reason in comments for the TPM and add a label `manual-check-required` to identify such tickets.
- For internal tickets, where there is no UDX assigned, the service team members can bypass this check by adding the
  label `Bypass-APIRB-Ready-Check-UDX` to the APIRB Jira ticket.


## Configuring the TC job without modifying the script
With changing requirements there might be cases where some modification may be required in order to keep the script behaving
per the current need. The script supports modifying the script values by setting the following environment variables:-
- `GATE_CHECK_PASS_LABEL` The label that will be added to tickets when the automation passes all checks.
- `BYPASS_GATE_CHECK_LABEL` The label that the TPM will use to bypass the automation check
- `MANUAL_CHECK_LABEL` The label that will be added to tickets when additional manual checks may be required.
- `QUERY_TEMPLATE` The JQL Query used to check which tickets would be considered for automation check.
- `DEFAULT_SPEC_VALIDATOR_TOOL_VERSION` Fallback version of the Spec Validator Tool version
- `VALIDATOR_TOOL_VERSION_CUTOFF` Default minor version cutoff allowed for the Tool
- `API_REVIEW_PR_PREFIX` The Prefix used to provide API Review PR links to the automation 
- `DEFAULT_TICKET_TRIAGE_ASSIGNEE` The TPM username who will be assigned this ticket when automation check passes.
- `BYPASS_UDX_CHECK_LABEL` The label that the users can use to bypass the UDX check

## Running the script locally
You can run the script via the following command:-
```
python ./team_city_scripts/api_review/gate_check_api_review_tickets.py
```
If you wish to run the script in dry run mode, please use:-
```
python ./team_city_scripts/api_review/gate_check_api_review_tickets.py --dry-run
```
PS: You will need to be connected to Oracle Network and set the following environment variables at minimum for this to work and :-
- JIRA_USERNAME
- JIRA_PASSWORD
- BITBUCKET_USERNAME
- BITBUCKET_PASSWORD
