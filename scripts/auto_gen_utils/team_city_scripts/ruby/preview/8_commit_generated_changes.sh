#!/bin/bash

set -x
set -e

# activate venv from step 1 to run below python script
ls -la ./.sdk-venv
. ./.sdk-venv/bin/activate

# get the branch we're on
cd ruby-sdk
branch=`git branch|grep "^\*"|cut -c3-`
if [[ ${branch} == *"bulk"* ]]; then
  build_type="bulk_pending_merge_preview"
else
  build_type="individual_preview"
fi
cd ..

# commit changes from generation and build.sh for ruby-sdk
# this step will run no matter what happened before it so it can report success / failure to the JIRA tickets
cd autogen
ls -la
python ./4_on_generation_complete.py --build-id $BUILD_ID --tool RubySDK --build-type ${build_type}
cd ..