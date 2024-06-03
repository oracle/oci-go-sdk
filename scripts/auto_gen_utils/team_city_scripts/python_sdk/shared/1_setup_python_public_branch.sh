set -e
set -x

echo Creating venv to install sdk locally
. /opt/odo/tox_sic/venv/bin/activate
virtualenv .sdk-venv
. .sdk-venv/bin/activate

pip ${PIP_TIMEOUT_PARAMETER} install -U pip

## Install AUTOGEN requirements##
pip ${PIP_TIMEOUT_PARAMETER} install --trusted-host artifactory.oci.oraclecorp.com --index-url https://artifactory.oci.oraclecorp.com/api/pypi/global-dev-pypi/simple -r autogen/requirements.txt

## Install SDK requirements ##
cd python-sdk
git pull
git fetch
pip ${PIP_TIMEOUT_PARAMETER} install --trusted-host artifactory.oci.oraclecorp.com --index-url https://artifactory.oci.oraclecorp.com/api/pypi/global-dev-pypi/simple -r requirements.txt
pip ${PIP_TIMEOUT_PARAMETER} install --trusted-host artifactory.oci.oraclecorp.com --index-url https://artifactory.oci.oraclecorp.com/api/pypi/global-dev-pypi/simple -r requirements-internal.txt
cd ..

# run the setup job for step 2 in ./autogen
cd autogen
ls -la
pwd

# checks out CLI branch with same name as SDK branch that triggered this build
python ./2_pre_generation_set_up.py --build-id $BUILD_ID --tool PythonSDK

# back out into root directory
cd ..
