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

# run the job to pull JIRAs from the queue and update POM
# this job expects to be run from ./autogen
cd auto-gen
ls -la
pwd

pip ${PIP_TIMEOUT_PARAMETER} install --trusted-host artifactory.oci.oraclecorp.com --index-url https://artifactory.oci.oraclecorp.com/api/pypi/global-dev-pypi/simple -r requirements.txt

# This should be defined within the TC job configuration: export BUILD_ID=%teamcity.build.id%
RUBY_VERSION=$(grep "VERSION" ../ruby-sdk/lib/oci/version.rb | cut -d "'" -f2 )
FULL_VERSION=$RUBY_VERSION.$BUILD_NUMBER
python ./5_mark_preview_tickets_done_post_merge.py \
    --build-id $BUILD_ID \
    --tool RubySDK \
    --full-version $FULL_VERSION \
    --build-conf-name Sdk_RubySdk_BuildGemAndDocs \
    --allow-transition-overall-issue-to-deploy

# back out into root directory
cd ..
