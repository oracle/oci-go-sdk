set -e
set -x

pwd
ls -la

# activate venv from step 1 to run below python script
ls -la ./.sdk-venv
. ./.sdk-venv/bin/activate

cd src/github.com/oracle/oci-go-sdk
if [[ `git branch|grep "\*"|grep "bulk"` ]]; then
  echo "Found bulk preview"
  export BUILD_TYPE_ARG="--build-type bulk_pending_merge_public"
  export PUSH_SPEC_DIFF_ARG=""
  export PUSH_SPEC_DIFF_UNPROTECTED_ARG=""
else
  export BUILD_TYPE_ARG="--build-type individual_public"
fi
cd ../../../..

if [ -n "$PUSH_SPEC_DIFF_UNPROTECTED_ARG" ]; then
  set +e
  diff=`diff -r specs-without-any-changes specs-without-conditional-groups/ | grep -v "^Only in specs-without-conditional-groups/"`
  diff_lines=`diff -r specs-without-any-changes specs-without-conditional-groups/ | grep -v "^Only in specs-without-conditional-groups/" | wc -l`
  if [ "$diff_lines" -gt "0" ]; then
    echo "WARNING: Detected changes in the spec that were not protected by conditional groups." >&2
    echo "$diff"
  else
    echo "Did not detect any changes that were not protected by conditional groups."
    export PUSH_SPEC_DIFF_UNPROTECTED_ARG=""
  fi
  set -e
fi

# commit changes from generation and build for go-sdk
# this step will run no matter what happened before it so it can report success / failure to the JIRA tickets
cd autogen
python ./3_report_generation_status.py --build-id $BUILD_ID --tool GoSDK $BUILD_TYPE_ARG $PUSH_SPEC_DIFF_ARG $PUSH_SPEC_DIFF_UNPROTECTED_ARG
cd ..

ls -la ./src/github.com/oracle/oci-go-sdk
