Auto Generation Utils
========================

Note
-----

Every time you clone this repo, please set up the pre-commit hook, see (#Pre-Check-in-Validation): `git config core.hooksPath hooks`


Description
------------

This package contains scripts related to our SDK/CLI self-service pipeline. Most of these scripts run as part of TeamCity builds.


### Self-service Pipeline

TODO: update for Build service

Runs every hour:
- Preview: https://teamcity.oci.oraclecorp.com/viewType.html?buildTypeId=Sdk_SelfService_PublicV2_ProcessJiraAutoPublicQueue
- Public: https://teamcity.oci.oraclecorp.com/viewType.html?buildTypeId=Sdk_SelfService_Preview_ProcessJiraAutoPreviewQueue

```
1_process_preview_jira_queue.py
```

If there were tickets in 'Processing Requested', the above job will create new branches for the SDKs and the CLI and update the `pom.xml` file with the new data from the DEXREQ tickets (e.g. a spec artifact version change).

The scripts that actually modify the `pom.xml` files are

```
add_or_update_scripts/*.py
```

This Git change will trigger generation builds for each of the SDKs and the CLI:

TODO: update for Build service

- Preview: https://teamcity.oci.oraclecorp.com/project.html?projectId=Sdk_SelfService_Preview_AutoPreviewBuilds&tab=projectOverview
- Public: https://teamcity.oci.oraclecorp.com/project.html?projectId=Sdk_SelfService_PublicV2_AutoPublicBuilds&tab=projectOverview

These builds will run, among other things, these scripts:

```
2_pre_generation_set_up.py
3_report_generation_status.py
4_on_generation_complete.py
5_mark_preview_tickets_done_post_merge.py
```

The different SDKs have additional scripts that may be part of these generation builds in the `team_city_scripts` directory:

```
team_city_scripts/
├── cli
├── go
├── java
├── python_sdk
└── ruby
```


### Advisor

The "advisor" runs periodically, looks at DEXREQ JIRA tickets that have changed, and comments on their state. If there are actionable state changes, it also moves the tickets forward (or backward, if illegal manual changes were made) in the pipeline.

TODO: update for Build service

Runs every 10 minutes: https://teamcity.oci.oraclecorp.com/viewType.html?buildTypeId=Sdk_SelfService_RunTicketAdvisorV2

```
autogen_issue_advisor.py
autogen_issue_advisor_preview.py
autogen_issue_advisor_public.py
autogen_issue_advisor_shared.py
```


### Branch Cleanup

Runs daily and cleans up old branches, and merges/declines old spec diff PRs.

TODO: update for Build service

- https://teamcity.oci.oraclecorp.com/viewType.html?buildTypeId=Sdk_SelfService_Preview_CleanUpAutoPreviewBranches&tab=buildTypeHistoryList&branch_Sdk_SelfService_Preview=__all_branches__

```
clean_auto_branches.py
```


### Other Scripts for TeamCity Builds

There are several other scripts that run on TeamCity, e.g. as part of on-PR builds or triggered by timers.

```
team_city_scripts/
├── api_review
├── github_issues
├── oci_testing_service
├── orm
└── spec_validator
```


### Libraries/Utilities

There are a number of utility/library files:

- `shared/bitbucket_utils.py` -- interfacing with Bitbucket PRs and repos
- `shared/version_utils.py` -- dealing with Maven versions and checking if they are valid and increasing
- `config.py` -- mostly names, repos, etc. for the different SDKs and the CLI
- `util.py` -- interfacing with JIRA, among other things


Further reading
----------------

- [Self-service Testing and Development](https://confluence.oci.oraclecorp.com/display/DEX/Self-Service+Testing+and+Development)
- [Requesting a Public SDK](https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=43683000)
- [Requesting a Preview SDK](https://confluence.oci.oraclecorp.com/display/DEX/Requesting+a+preview+SDK+CLI)
- [Self-service Pipeline Limitations](https://confluence.oci.oraclecorp.com/display/DEX/Self-service+Pipeline+Limitations)
- [Preview/Public Pipeline Ticket Advisor](https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=60740381)
- [SDK Runbook - Exceptions in Self-Service Pipeline](https://confluence.oci.oraclecorp.com/display/DEX/SDK+Runbook+-+Exceptions+in+Self-Service+Pipeline)
- [SDK Runbook - Running the Pipeline on Just One SDK or CLI](https://confluence.oci.oraclecorp.com/display/DEX/SDK+Runbook+-+Running+the+Pipeline+on+Just+One+SDK+or+CLI)
- [Scripts to Add or Update Specs](https://confluence.oci.oraclecorp.com/display/DEX/Scripts+to+Add+or+Update+Specs)
- [Self Service Preview Generation](https://confluence.oci.oraclecorp.com/display/DEX/Self+Service+Preview+Generation)
- [Pull Request Validation Builds for the Testing Service](https://confluence.oci.oraclecorp.com/display/DEX/Pull+Request+Validation+Builds+for+the+Testing+Service)
- [OCI Testing Service Tests Run In PRs](https://confluence.oci.oraclecorp.com/display/~mricken/OCI+Testing+Service+Tests+Run+In+PRs)


Pre Check-in Validation
------------------------
Run the verify script before checking in which performs some basic validation:

`./verify.sh`

We recommend adding this as a pre-commit hook which can be done by running the following command:
`git config core.hooksPath hooks`


Installing dependencies
------------------------
To install dependencies for this project, execute the following command:

`pip install -r requirements.txt`


Running scripts locally
-------------------------
In order to run the scripts locally you will need a JIRA token set in the JSESSIONID environment variable.

To get this token, run the following script:

`python get_jira_access_token.py`

This will print out a token which you should set in your environment like so:

`export JSESSIONID={TOKEN HERE}`

If you get an error like the following, you haven't set JSESSIONID in your environment.  `config.JSESSIONID` is the name
of the Team City variable, but when running locally use JSESSIONID.

    Could not authenticate with JIRA server. Must specify environment variables for either config.JSESSIONID or JIRA_USERNAME and JIRA_PASSWORD.

Example of running 1_process_preview_jira_queue.p locally

    python 1_process_preview_jira_queue.py --dry-run --issue DEXREQ-354 --tool PythonSDK --allow-individual-tool-generation --for-any-status --build-type individual_public --base-branch master --build-id 645

Some scripts also need JIRA access. For that, please set the following environment variables:

    JIRA_USERNAME=...
    JIRA_PASSWORD=...

Some scripts also need Bitbucket access. For that, please set the following environment variables:

    BITBUCKET_USERNAME=...
    BITBUCKET_PASSWORD=...

If those aren't set, the scripts will fall back to using JIRA_USERNAME and JIRA_PASSWORD, for backward-compatibility reasons, even though the accounts are not linked anymore, and the credentials may differ.


Testing Jira bot changes locally
--------------------------------
- Use zsh shell for the following steps.
- Set the following env variables `PYTHONPATH`,`PYTHON_CLI_DIR` as shown below.

- Activate venv and install all requirements for `python-cli` and `auto-gen-utils` repo

- Use python 3.7 or higher version

Example directory structure :
```
dir1
├── cli-bot-env <venv folder>         
├── auto-gen-utils <repo folder> 
├── python-cli <repo folder> 
```
**Downloading python3.7 using pyenv**
```
zsh
echo "Creating a python virtual environment"
curl -L https://raw.githubusercontent.com/yyuu/pyenv-installer/master/bin/pyenv-installer | bash
export PATH="$HOME/.pyenv/bin:$PATH"
eval "$(pyenv init -)"
eval "$(pyenv init --path)"
eval "$(pyenv virtualenv-init -)"
# pyenv update
pyenv install 3.7.0 -s
pyenv shell 3.7.0

```
**Creating virtual venv**
```

python3.7 -m venv cli-bot-env
source  cli-bot-env/bin/activate

cd auto-gen-utils
export PYTHONPATH=$(pwd):$PYTHONPATH
pip3 install --pre --trusted-host artifactory.oci.oraclecorp.com --index-url https://artifactory.oci.oraclecorp.com/api/pypi/global-dev-pypi/simple -r requirements.txt
cd ..

```
**User needs to git clone python-cli (under dir1, refer above diagram)**
\
**Now setup python-cli development env as follows**
```

cd python-cli
export PYTHONPATH=$(pwd):$PYTHONPATH
export PYTHON_CLI_DIR=$(pwd)

../auto-gen-utils/python_cli/install_python_cli_local.sh

```

Refer the below example for testing a command against a dummy DEX issue. 
Note: Please always follow the first step to clear any local changes,commits and to check out preview in python-cli repo before running the script every time.
```
cd ../python-cli
git restore . && git clean -fd && git checkout preview && git branch -D preview-DEXTEST-1234
cd ..

cd auto-gen-utils
export issue="DEXTEST-1234"
export command="[~gear-dexreq-automation] Manual Changes Requested
[Rename Parameter]

oci bds instance install-patch --patch-version -> --patch-testing-param"
python3 python_cli/generate_local_changes.py $issue $command
```
