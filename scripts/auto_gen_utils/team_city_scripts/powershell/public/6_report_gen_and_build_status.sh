set -e
set -x

pwd
ls -la

# activate venv from step 1 to run below python script
ls -la ./.sdk-venv
. ./.sdk-venv/bin/activate

cd oci-powershell-modules
branch=`git branch|grep "^\*"|cut -c3-`
if [[ ${branch} == *"bulk"* ]]; then
  build_type="bulk_pending_merge_public"
else
  build_type="individual_public"
fi

cd ..
# commit changes from generation and build for powershell
# this step will run no matter what happened before it so it can report success / failure to the JIRA tickets
cd autogen
python ./3_report_generation_status.py --build-id $BUILD_ID --tool PowerShell --build-type  ${build_type}
cd ..

ls -la ./oci-powershell-modules
