set -e
set -x

touch cli_build_start.txt

# must disable StrictHostKeyChecking so that we don't get an interactive
# prompt later asking to confirm the host key
# Must disable -e (fail on non-zero exit code) because ssh returns 255
set +e
ssh -o StrictHostKeyChecking=no git@bitbucket.oci.oraclecorp.com -p 7999
set -e

# Old way of doing that:
# ls -la ~/.ssh
#
# cat ~/.ssh/config
# 
# printf "\n\nHost * \n  StrictHostKeyChecking no\n" >> ~/.ssh/config
# 
# cat ~/.ssh/config

# configure git for this commit
git config --global user.email "$GIT_USER_EMAIL"
git config --global user.name "$GIT_USER_NAME"

DEV_VERSION=$(cat ~/.DEV_VERSION)
SDK_VERSION=$(cat ~/.SDK_VERSION)

cd python-cli
source scripts/common_build_functions.sh
f_activate_virtualenv
./scripts/replace_oci_version.py $DEV_VERSION

# sleep to allow time for Python SDK from last step to be published
sleep 300

# run regular python build to produce artifacts
# TEST_ENABLE is set to 'false' so tests are skipped
source scripts/build_preview.sh "individual_public"

# the build script creates a virtualenv inside this folder which we need to remove or it will be checked in
rm -rf ./.sdk-venv

cd ..

# DEV_VERSION should contain the version string for this build of the CLI
echo "Dev version: $DEV_VERSION"