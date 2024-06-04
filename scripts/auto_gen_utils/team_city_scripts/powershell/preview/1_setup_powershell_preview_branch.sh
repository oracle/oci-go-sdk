set -e
set -x

echo Creating venv to install sdk locally
. /opt/odo/tox_sic/venv/bin/activate
virtualenv .sdk-venv
. .sdk-venv/bin/activate

# must disable StrictHostKeyChecking so that we don't get an interactive
# prompt later asking to confirm the host key
# If `set -e`, must disable "fail on non-zero exit code" using `set +e`
# because ssh returns 255
set +e
ssh -o StrictHostKeyChecking=no git@bitbucket.oci.oraclecorp.com -p 7999
set -e


## AUTOGEN ##
cd autogen
pip ${PIP_TIMEOUT_PARAMETER} install --trusted-host artifactory.oci.oraclecorp.com --index-url https://artifactory.oci.oraclecorp.com/api/pypi/global-dev-pypi/simple -r requirements.txt --quiet

# run the setup job for step 2 in ./autogen
ls -la
pwd

python ./2_pre_generation_set_up.py --build-id $BUILD_ID --tool PowerShell

# back out into root directory
cd ..