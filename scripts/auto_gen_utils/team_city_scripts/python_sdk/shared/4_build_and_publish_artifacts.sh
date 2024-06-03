set -e
set -x

# ls -la ./.sdk-venv
# . ./.sdk-venv/bin/activate

# Set up python3 pyenv(v3.8.6)
echo "Setup python environment"
curl -L https://raw.githubusercontent.com/yyuu/pyenv-installer/master/bin/pyenv-installer | bash
export PATH="$HOME/.pyenv/bin:$PATH"
eval "$(pyenv init -)"
eval "$(pyenv init --path)"
eval "$(pyenv virtualenv-init -)"
export PYTHON_3_VERSION=3.6.5
pyenv install $PYTHON_3_VERSION -s
pyenv shell $PYTHON_3_VERSION

echo "Python Version"
python --version

pip install -U pip
pip ${PIP_TIMEOUT_PARAMETER} install -e ./python-sdk

# cat ~/.ssh/config

# configure git for this commit
git config --global user.email "$GIT_USER_EMAIL"
git config --global user.name "$GIT_USER_NAME"

cd python-sdk
# the below commands are copied from:
# Sdk_PythonSdk_PreviewTestsDocsWheel
# at some point we should consolidate Python SDK Preview Build to use build.sh

SDK_VERSION=$(tail -1 src/oci/version.py | cut -d '"' -f2)
DEV_VERSION=$SDK_VERSION.$BUILD_NUMBER
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
# pip install sphinx --timeout 120
# pip ${PIP_TIMEOUT_PARAMETER} install sphinx_rtd_theme
pip ${PIP_TIMEOUT_PARAMETER} install -r docs/requirements.txt
make docs
mkdir -p dist/oci-python-sdk-docs-$SDK_VERSION/
cp -r docs/_build/html/* dist/oci-python-sdk-docs-$SDK_VERSION/

echo Running Tests

if [ $TEST_ENABLE = "false" ]; then
  echo "TESTS HAVE BEEN DISABLED."
else
  pip ${PIP_TIMEOUT_PARAMETER} install tox
  tox -e flake8,py36
fi

pip install wheel
echo Building Wheel
make build

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

pushd dist/oci-python-sdk-docs-$SDK_VERSION
zip -r ../oci-python-sdk-docs-$SDK_VERSION.zip .
popd
cp dist/oci-python-sdk-docs-$SDK_VERSION.zip dist/dev/oci-python-sdk-docs-$DEV_VERSION.zip


echo Contents of dist folder
ls -la dist

# the build script creates a virtualenv inside this folder which we need to remove or it will be checked in
# commenting this out since we are not invoking build.sh
# rm -rf ./.sdk-venv

# AFTER building the wheel, reset src/oci/version.py back to regular version, we don't want to check in TC version
git checkout -- src/oci/version.py

cd ..

# write DEV_VERSION to a text file so next step can use it
echo $DEV_VERSION >> ~/.DEV_VERSION

# DEV_VERSION should contain the version string for this build of the Python SDK
# we need this in the next step to know which version the CLI should depend on
echo "Dev version: $DEV_VERSION"
