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
    $FORCE_ARG \
    $VERBOSE_ARG \
    $ISSUE_FILTER_ARG \
    $BULK_PREVIEW_DATE_OVERRIDES_ARG \
    $PUBLIC_RELEASE_REQUESTED_CUT_OFF_DATE_OVERRIDES_ARG \
    $PUBLIC_RELEASE_DATE_OVERRIDES_ARG \
    --show-ga-calendar \
    --show-ga-calendar-count $SHOW_GA_CALENDAR_COUNT | tee ../ga-dates.txt


python ./autogen_issue_advisor.py \
    $PIPELINE_ARG \
    $DRY_RUN_ARG \
    $FORCE_ARG \
    $VERBOSE_ARG \
    $ISSUE_FILTER_ARG \
    $BULK_PREVIEW_DATE_OVERRIDES_ARG \
    --show-preview-calendar \
    --show-preview-calendar-count $SHOW_PREVIEW_CALENDAR_COUNT | tee ../preview-dates.txt

# back out into root directory
cd ..


# uploads ga-dates.txt
curl https://objectstorage.us-phoenix-1.oraclecloud.com/p/Lu2GYrJGOs6JCEibFpXmirYt1LAIXAWuiC6Cx8ps8KXBAxPeOJWmOb6yz0QNH3qS/n/dex-us-phoenix-1/b/generated_markdown/o/ga-dates.txt --upload-file ga-dates.txt

# uploads preview-dates.txt
curl https://objectstorage.us-phoenix-1.oraclecloud.com/p/UXNHi_BuJ58pHnv4Zucc6oE1WE_-VqeTp5ELkoN1Wqb_WCikNwb7BaMIM0BI6eJ-/n/dex-us-phoenix-1/b/generated_markdown/o/preview-dates.txt --upload-file preview-dates.txt
