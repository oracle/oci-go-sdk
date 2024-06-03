set -e
set -x

pwd
ls -la

# activate venv from step 1 to run below python script
ls -la ./.sdk-venv
. ./.sdk-venv/bin/activate

cd oci-dotnet-sdk
if [[ `git branch|grep "\*"|grep "bulk"` ]]; then
  echo "Found bulk public"
  export BUILD_TYPE_ARG="--build-type bulk_pending_merge_public"
else
  export BUILD_TYPE_ARG="--build-type individual_public"
fi
cd ..

# commit changes from generation and build for typescript-sdk
# this step will run no matter what happened before it so it can report success / failure to the JIRA tickets
cd autogen
python ./3_report_generation_status.py --build-id $BUILD_ID --tool DotNetSDK $BUILD_TYPE_ARG
cd ..

ls -la ./oci-dotnet-sdk
