set -e
set -x

pwd
ls -la

cd python-cli
# activate python env
source scripts/common_build_functions.sh
f_activate_virtualenv

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
# This should be defined within the TC job configuration: export BUILD_ID=%teamcity.build.id%
python ./3_report_generation_status.py --build-id $BUILD_ID $BUILD_TYPE_ARG
cd ..

ls -la ./python-cli
ls -la ./python-sdk