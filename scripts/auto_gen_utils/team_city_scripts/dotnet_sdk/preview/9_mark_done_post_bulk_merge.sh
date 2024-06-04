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

# Generate version string for .NET SDK
cd ../oci-dotnet-sdk
./nuget-package.sh version
cd ../auto-gen
# $BUILD_ID should be defined within the TC job configuration: export BUILD_ID=%teamcity.build.id%
FULL_VERSION=`cat ../oci-dotnet-sdk/version.txt`.$BUILD_ID
python ./5_mark_preview_tickets_done_post_merge.py --build-id $BUILD_ID --tool DotNetSDK --full-version $FULL_VERSION --allow-transition-overall-issue-to-done --build-conf-name Sdk_DotNetSdk_BuildDotNetSdkPreview

# back out into root directory
cd ..
