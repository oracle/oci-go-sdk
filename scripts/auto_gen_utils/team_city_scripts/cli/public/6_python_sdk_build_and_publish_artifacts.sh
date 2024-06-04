set -e
set -x

touch sdk_build_start.txt

# activate python environment
cd python-cli
source scripts/common_build_functions.sh
f_activate_virtualenv
cd ..

pip ${PIP_TIMEOUT_PARAMETER} install -e ./python-sdk

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

cd python-sdk

# the below commands are copied from:
# Sdk_PythonSdk_PreviewTestsDocsWheel
# at some point we should consolidate Python SDK Preview Build to use build.sh

SDK_VERSION=$(tail -1 src/oci/version.py | cut -d '"' -f2)

# TODO: come up with a better versioning scheme for these artifacts
# for right now, we are using +0.{build number} so that it is always behind the +{build number} that
# the official python build wheel job publishes
DEV_VERSION=$SDK_VERSION+0.$BUILD_NUMBER.selfservice
echo SDK Version Number $SDK_VERSION
echo Build Version Number $DEV_VERSION

echo Rewriting version from $SDK_VERSION to $DEV_VERSION
# Replace the version with the DEV_VERSION (SDK_VERSION + Build Number) so that we can make
# referencing and declaring dependencies on preview CLIs more explicit
rm src/oci/version.py
cat <<EOF > src/oci/version.py
# coding: utf-8
# Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.

__version__ = "$DEV_VERSION"

EOF

# Echo out the version to confirm
cat src/oci/version.py

echo Building Docs
pip install sphinx --timeout 120
pip ${PIP_TIMEOUT_PARAMETER} install sphinx_rtd_theme

echo Running Tests

if [ "$TEST_ENABLE" = "false" ]; then
  echo "TESTS HAVE BEEN DISABLED."
else
  pip ${PIP_TIMEOUT_PARAMETER} install tox
  tox -e flake8,py27,py35
fi

echo Building Wheel
# Redirect STDOUT and STDERR to a file to avoid resource unavailable error in TeamCity jobs.
mkdir -p docs/_build/html
make build >> build_output.txt 2>&1

# Create a dev directory that will contain versions of the whl, zip, and docs meant for
# the dev pypi artifactory. Each artifact includes the build number in the version to avoid
# conflicts.

mkdir -p dist/dev/
if [ -f "dist/oci-$DEV_VERSION-py3-none-any.whl" ]; then
	cp dist/oci-$DEV_VERSION-py3-none-any.whl dist/dev/oci-$DEV_VERSION-py3-none-any.whl
else
	cp dist/oci-$DEV_VERSION-py2.py3-none-any.whl dist/dev/oci-$DEV_VERSION-py2.py3-none-any.whl
fi
cp dist/oci-python-sdk-$DEV_VERSION.zip dist/dev/oci-python-sdk-$DEV_VERSION.zip


echo Contents of dist folder
ls -la dist

# the build script creates a virtualenv inside this folder which we need to remove or it will be checked in
# commenting this out since we are not invoking build.sh
# rm -rf ./.sdk-venv

# AFTER building the wheel, reset src/oci/version.py back to regular version, we don't want to check in TC version
git checkout -- src/oci/version.py

# Delete build_output.txt.
rm build_output.txt

cd ..

# write DEV_VERSION to a text file so next step can use it
echo $DEV_VERSION >> ~/.DEV_VERSION

# write SDK_VERSION to a text file so next step can use it
echo $SDK_VERSION >> ~/.SDK_VERSION

# DEV_VERSION should contain the version string for this build of the Python SDK
# we need this in the next step to know which version the CLI should depend on
echo "Dev version: $DEV_VERSION"
echo "SDK version: $SDK_VERSION"