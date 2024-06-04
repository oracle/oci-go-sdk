set -e
set -x

pwd
ls -la

# activate venv from step 1 to run below python script
ls -la ./.sdk-venv
. ./.sdk-venv/bin/activate

# commit changes from generation and build.sh for python-cli and python-sdk
# this step will run no matter what happened before it so it can report success / failure to the JIRA tickets
# $BUILD_TYPE_ARG is defined in the TC build step configuration
cd autogen
python ./3_report_generation_status.py --build-id $BUILD_ID --tool PythonSDK $BUILD_TYPE_ARG
cd ..

ls -la ./python-sdk