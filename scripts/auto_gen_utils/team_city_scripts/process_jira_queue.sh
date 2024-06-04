#!/bin/bash

set -e
set -x

echo Creating venv to install sdk locally
. /opt/odo/tox_sic/venv/bin/activate
virtualenv .sdk-venv
. .sdk-venv/bin/activate

# Upgrade python version to Python=3.6.5
echo "Setup python environment"
curl -L https://raw.githubusercontent.com/yyuu/pyenv-installer/master/bin/pyenv-installer | bash
export PATH="$HOME/.pyenv/bin:$PATH"
eval "$(pyenv init -)"
eval "$(pyenv init --path)"
eval "$(pyenv virtualenv-init -)"
export PYTHON_3_VERSION=3.6.5
pyenv install $PYTHON_3_VERSION -s
pyenv shell $PYTHON_3_VERSION
echo "Python Version"
python --version
pip ${PIP_TIMEOUT_PARAMETER} install -U pip

# must disable StrictHostKeyChecking so that we don't get an interactive
# prompt later asking to confirm the host key
# Must disable -e (fail on non-zero exit code) because ssh returns 255
set +e
ssh -o StrictHostKeyChecking=no git@bitbucket.oci.oraclecorp.com -p 7999
set -e

# Old way of doing that:
# ls -la ~/.ssh
#
# cat ~/.ssh/config
# 
# printf "\n\nHost * \n  StrictHostKeyChecking no\n" >> ~/.ssh/config
# 
# cat ~/.ssh/config

# configure git for this commit
git config --global user.email "$GIT_USER_EMAIL"
git config --global user.name "$GIT_USER_NAME"

echo "## CLI ##"
cd python-cli
git fetch $GIT_OPTS origin $BASE_BRANCH
git checkout $BASE_BRANCH
if [ $? -ne 0 ]; then
  echo "Failed to check out base branch: $BASE_BRANCH.  Exiting script."
  exit 1
fi

pip ${PIP_TIMEOUT_PARAMETER} install --use-deprecated=legacy-resolver --trusted-host artifactory.oci.oraclecorp.com --index-url https://artifactory.oci.oraclecorp.com/api/pypi/global-dev-pypi/simple -r requirements.txt
pip ${PIP_TIMEOUT_PARAMETER} install --use-deprecated=legacy-resolver --trusted-host artifactory.oci.oraclecorp.com --index-url https://artifactory.oci.oraclecorp.com/api/pypi/global-dev-pypi/simple -r requirements-internal.txt

git diff --color | cat
cd ..

echo "## Python SDK ##"
cd python-sdk
git fetch $GIT_OPTS origin $BASE_BRANCH
git checkout $BASE_BRANCH
if [ $? -ne 0 ]; then
  echo "Failed to check out base branch: $BASE_BRANCH.  Exiting script."
  exit 1
fi

pip ${PIP_TIMEOUT_PARAMETER} install --trusted-host artifactory.oci.oraclecorp.com --index-url https://artifactory.oci.oraclecorp.com/api/pypi/global-dev-pypi/simple -r requirements.txt
pip ${PIP_TIMEOUT_PARAMETER} install --trusted-host artifactory.oci.oraclecorp.com --index-url https://artifactory.oci.oraclecorp.com/api/pypi/global-dev-pypi/simple -r requirements-internal.txt

git diff --color | cat
cd ..

echo "## Java SDK ##"
cd java-sdk
git fetch $GIT_OPTS origin $BASE_BRANCH
git checkout $BASE_BRANCH
if [ $? -ne 0 ]; then
  echo "Failed to check out base branch: $BASE_BRANCH.  Exiting script."
  exit 1
fi

git diff --color | cat
cd ..

echo "## GO SDK ##"
cd src/github.com/oracle/oci-go-sdk
git fetch $GIT_OPTS origin $BASE_BRANCH
git checkout $TEST_BRANCH
if [ $? -ne 0 ]; then
  echo "Failed to check out base branch: $BASE_BRANCH.  Exiting script."
  exit 1
fi

git diff --color | cat
cd ../../../..

echo "## Ruby SDK ##"
cd ruby-sdk
git fetch $GIT_OPTS origin $BASE_BRANCH
git checkout $BASE_BRANCH
if [ $? -ne 0 ]; then
  echo "Failed to check out base branch: $BASE_BRANCH.  Exiting script."
  exit 1
fi

git diff --color | cat
cd ..

echo "## Typescript SDK ##"
cd oci-typescript-sdk
git fetch $GIT_OPTS origin $BASE_BRANCH
git checkout $BASE_BRANCH
if [ $? -ne 0 ]; then
    echo "Failed to check out base branch: $BASE_BRANCH.  Exiting script."
    exit 1
fi
git diff --color | cat
cd ..

echo "## .NET SDK ##"
cd oci-dotnet-sdk
git fetch $GIT_OPTS origin $BASE_BRANCH
git checkout $BASE_BRANCH
if [ $? -ne 0 ]; then
    echo "Failed to check out base branch: $BASE_BRANCH.  Exiting script."
    exit 1
fi
git diff --color | cat
cd ..

echo "## PowerShell ##"
cd oci-powershell-modules
git fetch $GIT_OPTS origin $BASE_BRANCH
git checkout $BASE_BRANCH
if [ $? -ne 0 ]; then
    echo "Failed to check out base branch: $BASE_BRANCH.  Exiting script."
    exit 1
fi
git diff --color | cat
cd ..

echo "## Legacy Java SDK ##"
cd legacy-java-sdk
git fetch $GIT_OPTS origin $BASE_BRANCH
git checkout $BASE_BRANCH
if [ $? -ne 0 ]; then
  echo "Failed to check out base branch: $BASE_BRANCH.  Exiting script."
  exit 1
fi
git diff --color | cat
cd ..

echo "## pom.xml Updates ##"
# run the job to pull JIRAs from the queue and update POM
# this job expects to be run from ./autogen
cd autogen
ls -la
pwd

pip ${PIP_TIMEOUT_PARAMETER} install --trusted-host artifactory.oci.oraclecorp.com --index-url https://artifactory.oci.oraclecorp.com/api/pypi/global-dev-pypi/simple -r requirements.txt

ls -la

ALLOW_INDIVIDUAL_TOOL_GENERATION_ARG=""
if [ "$ALLOW_INDIVIDUAL_TOOL_GENERATION" == "true" ]; then
  ALLOW_INDIVIDUAL_TOOL_GENERATION_ARG="--allow-individual-tool-generation"
fi

ISSUE_FILTER_ARG=""
if [ ! -z "$ISSUE_FILTER" ]; then
  ISSUE_FILTER_ARG="--issue "$ISSUE_FILTER
fi


if [ ! -z "$LIST_OF_GO_TICKETS" ]; then
  source ./change_dexreq_to_release_approved.sh $LIST_OF_GO_TICKETS
fi



# updated script can process all tools at once
python ./1_process_preview_jira_queue.py --tool ALL --build-id $BUILD_ID --build-type $BUILD_TYPE --base-branch $BASE_BRANCH $ALLOW_INDIVIDUAL_TOOL_GENERATION_ARG $DRY_RUN_ARG $ISSUE_FILTER_ARG $PUSH_SPEC_BASELINE_ARG $VERBOSE_ARG

# # CLI
# python ./1_process_preview_jira_queue.py --build-id $BUILD_ID --build-type $BUILD_TYPE --base-branch $BASE_BRANCH $ALLOW_INDIVIDUAL_TOOL_GENERATION_ARG $DRY_RUN_ARG $ISSUE_FILTER_ARG $PUSH_SPEC_BASELINE_ARG
# # Java SDK
# python ./1_process_preview_jira_queue.py --build-id $BUILD_ID --build-type $BUILD_TYPE --tool JavaSDK --base-branch $BASE_BRANCH $ALLOW_INDIVIDUAL_TOOL_GENERATION_ARG $DRY_RUN_ARG $ISSUE_FILTER_ARG $PUSH_SPEC_BASELINE_ARG
# # Python SDK
# python ./1_process_preview_jira_queue.py --build-id $BUILD_ID --build-type $BUILD_TYPE --tool PythonSDK --base-branch $BASE_BRANCH $ALLOW_INDIVIDUAL_TOOL_GENERATION_ARG $DRY_RUN_ARG $ISSUE_FILTER_ARG $PUSH_SPEC_BASELINE_ARG
# # Go SDK
# python ./1_process_preview_jira_queue.py --build-id $BUILD_ID --build-type $BUILD_TYPE --tool GoSDK --base-branch $TEST_BRANCH $ALLOW_INDIVIDUAL_TOOL_GENERATION_ARG $DRY_RUN_ARG $ISSUE_FILTER_ARG $PUSH_SPEC_BASELINE_ARG
# # Ruby SDK
# python ./1_process_preview_jira_queue.py --build-id $BUILD_ID --build-type $BUILD_TYPE --tool RubySDK --base-branch $TEST_BRANCH $ALLOW_INDIVIDUAL_TOOL_GENERATION_ARG $DRY_RUN_ARG $ISSUE_FILTER_ARG $PUSH_SPEC_BASELINE_ARG

# back out into root directory
cd ..
