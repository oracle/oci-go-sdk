SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

echo "===== Setting up git secrets ====="
source ${SCRIPT_DIR}/../shared-build-service-scripts/setup-git-secrets.sh
echo "===== Done setting up git secrets ====="

git clone ssh://git@bitbucket.oci.oraclecorp.com:7999/sdk/oci-sdk-cli-keys.git -b master --single-branch --depth 1