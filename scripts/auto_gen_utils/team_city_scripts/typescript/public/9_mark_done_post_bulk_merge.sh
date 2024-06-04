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


# run the job to pull JIRAs from the queue and update POM
# this job expects to be run from ./autogen
cd auto-gen
ls -la
pwd

pip ${PIP_TIMEOUT_PARAMETER} install --trusted-host artifactory.oci.oraclecorp.com --index-url https://artifactory.oci.oraclecorp.com/api/pypi/global-dev-pypi/simple -r requirements.txt

# This should be defined within the TC job configuration: export BUILD_ID=%teamcity.build.id%
FULL_VERSION=`cat ../oci-typescript-sdk/package_version`
python ./5_mark_preview_tickets_done_post_merge.py --build-id $BUILD_ID  --allow-transition-overall-issue-to-deploy --tool TypescriptSDK --full-version $FULL_VERSION --build-conf-name Sdk_TypeScriptSDK_BuildSdkMaster

# back out into root directory
cd ..