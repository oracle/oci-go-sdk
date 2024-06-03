set -x
set -e

# configure git for this commit
git config --global user.email "$GIT_USER_EMAIL"
git config --global user.name "$GIT_USER_NAME"

# activate venv from step 1 to run below python script
ls -la ./.sdk-venv
. ./.sdk-venv/bin/activate

# get the branch we're on
cd oci-typescript-sdk
branch=`git branch|grep "^\*"|cut -c3-`
if [[ ${branch} == *"bulk"* ]]; then
  build_type="bulk_pending_merge_preview"
else
  build_type="individual_preview"
fi
cd ..

# commit changes from generation and build
cd autogen
ls -la
python ./4_on_generation_complete.py --build-id $BUILD_ID --tool TypescriptSDK --build-type ${build_type}
cd ..
