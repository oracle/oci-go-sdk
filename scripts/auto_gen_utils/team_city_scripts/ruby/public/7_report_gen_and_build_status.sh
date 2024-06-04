#!/bin/bash

set -e
set -x

pwd
ls -la

# activate venv from step 1 to run below python script
ls -la ./.sdk-venv
. ./.sdk-venv/bin/activate

cd ruby-sdk
if [[ `git branch|grep "\*"|grep "bulk"` ]]; then
  echo "Found bulk preview"
  export BUILD_TYPE_ARG="--build-type bulk_pending_merge_public"
else
  export BUILD_TYPE_ARG="--build-type individual_public"
fi
cd ..

# commit changes from generation and build.sh for python-cli and python-sdk
# this step will run no matter what happened before it so it can report success / failure to the JIRA tickets
cd autogen
python ./3_report_generation_status.py --build-id $BUILD_ID --tool RubySDK $BUILD_TYPE_ARG
cd ..

ls -la ./ruby-sdk