set -x
set -e

# this is ONLY for master
# earlier in order to get the CLI artifact to work we updated the CLI setup.py to reference a specific SDK build
# we DO NOT want to check this in since the master version of setup.py cannot reference a specific internal build of the SDK
# so here before we commit the changes, we revert back to the original SDK version without the team city build number
SDK_VERSION=$(cat ~/.SDK_VERSION)
echo "SDK version: $SDK_VERSION"

cd python-cli

# activate python env
source scripts/common_build_functions.sh
f_activate_virtualenv

./scripts/replace_oci_version.py $SDK_VERSION
cd ..

# get the branch we're on
cd python-cli
branch=`git branch|grep "^\*"|cut -c3-`
if [[ ${branch} == *"bulk"* ]]; then
  build_type="bulk_pending_merge_public"
else
  build_type="individual_public"
fi
cd ..

# commit changes from generation and build.sh for python-cli and python-sdk
cd autogen
ls -la
# This should be defined within the TC job configuration: export BUILD_ID=%teamcity.build.id%
python ./4_on_generation_complete.py --build-id $BUILD_ID --build-type ${build_type}
cd ..