#!/bin/bash

set -e
set -x

echo Creating venv to install sdk locally
. /opt/odo/tox_sic/venv/bin/activate
virtualenv .sdk-venv
. .sdk-venv/bin/activate

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

cd autogen
ls -la
pwd

pip ${PIP_TIMEOUT_PARAMETER} install --trusted-host artifactory.oci.oraclecorp.com --index-url https://artifactory.oci.oraclecorp.com/api/pypi/global-dev-pypi/simple -r requirements.txt

ls -la

ISSUE_FILTER_ARG=""
if [ ! -z "$ISSUE_FILTER" ]; then
  ISSUE_FILTER_ARG="--issue $ISSUE_FILTER"
fi

BULK_PREVIEW_DATE_OVERRIDES_ARG=""
if [ ! -z "$BULK_PREVIEW_DATE_OVERRIDES" ]; then
  BULK_PREVIEW_DATE_OVERRIDES_ARG="--bulk-preview-date-overrides="${BULK_PREVIEW_DATE_OVERRIDES}
fi

PUBLIC_RELEASE_REQUESTED_CUT_OFF_DATE_OVERRIDES_ARG=""
if [ ! -z "$PUBLIC_RELEASE_REQUESTED_CUT_OFF_DATE_OVERRIDES" ]; then
  PUBLIC_RELEASE_REQUESTED_CUT_OFF_DATE_OVERRIDES_ARG="--public-release-requested-cut-off-date-overrides="${PUBLIC_RELEASE_REQUESTED_CUT_OFF_DATE_OVERRIDES}
fi

PUBLIC_RELEASE_DATE_OVERRIDES_ARG=""
if [ ! -z "${PUBLIC_RELEASE_DATE_OVERRIDES}" ]; then
  PUBLIC_RELEASE_DATE_OVERRIDES_ARG="--public-release-date-overrides="${PUBLIC_RELEASE_DATE_OVERRIDES}
fi

python ./autogen_issue_advisor.py \
    $PIPELINE_ARG \
    $DRY_RUN_ARG \
    $DISABLE_DATE_CHECK_ARG \
    $FORCE_ARG \
    $VERBOSE_ARG \
    $ISSUE_FILTER_ARG \
    $BULK_PREVIEW_DATE_OVERRIDES_ARG \
    $PUBLIC_RELEASE_REQUESTED_CUT_OFF_DATE_OVERRIDES_ARG \
    $PUBLIC_RELEASE_DATE_OVERRIDES_ARG \
    $IGNORE_WRONG_PIPELINE_ARG 

# back out into root directory
cd ..