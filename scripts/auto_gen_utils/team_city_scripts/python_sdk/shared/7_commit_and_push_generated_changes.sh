set -x
set -e

# activate venv from step 1 to run below python script
ls -la ./.sdk-venv
. ./.sdk-venv/bin/activate

# commit changes from generation and build.sh for python-cli and python-sdk
# this step will run no matter what happened before it so it can report success / failure to the JIRA tickets
cd autogen
ls -la
python ./4_on_generation_complete.py --build-id $BUILD_ID --tool PythonSDK  --build-type $BUILD_TYPE
cd ..