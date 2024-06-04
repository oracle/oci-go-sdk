#!/bin/bash

# Usage:
#
# Read from stdin:
#
# echo "Hello world!" | buildServiceScripts/make-pr-comment.sh
#
#
# Read from file:
#
# buildServiceScripts/make-pr-comment.sh myFile.txt
#
# The script JSON-escapes the input.
#
# Uses username and password from environment variables:
# - BITBUCKET_API_USERNAME
# - BITBUCKET_API_PASSWORD
#
# Uses Bitbucket project and repo from environment variables:
# - BITBUCKET_PROJECT
# - BITBUCKET_REPO
#
# If those are not set, then the script will attempt to extract
# them from BLD_VCS_ROOT or BLD_VSC_ROOT.
#
# Only runs if in the context of a pull-request build.
# Extracts the pull request id from BLD_BRANCH.

set -e

# from https://stackoverflow.com/questions/10053678/escaping-characters-in-bash-for-json/77930217#77930217
function escape_json_string() {
  local input=$1
  for ((i = 0; i < ${#input}; i++)); do
    local char="${input:i:1}"
    local escaped="${char}"
    case "${char}" in
      $'"' ) escaped="\\\"";;
      $'\\') escaped="\\\\";;
      *)
        if [[ "${char}" < $'\x20' ]]; then
          case "${char}" in 
            $'\b') escaped="\\b";;
            $'\f') escaped="\\f";;
            $'\n') escaped="\\n";;
            $'\r') escaped="\\r";;
            $'\t') escaped="\\t";;
            *) escaped=$(printf "\u%04X" "'${char}")
          esac
        fi;;
    esac
    echo -n "${escaped}"
  done
}

SHARED_SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

if [ $# -lt 1 ]; then
    # read from STDIN
    INPUT=$(</dev/stdin)
    TEXT=`escape_json_string "$INPUT"`
else
    # read from file provided as first argument
    tempfile=`mktemp`
    while IFS='' read -r line
    do
        part=`escape_json_string "$line"`
        printf '%s\\n' "$part" >> $tempfile
    done < "$1"
    TEXT=$(<$tempfile)
    rm $tempfile
fi

if ! echo "${BLD_BRANCH}" | grep "^PR-[1-9][0-9]*$" > /dev/null; then
    echo "Not running as a pull-request build."
    exit 0
fi

PR_ID=`echo ${BLD_BRANCH} | sed "s/^PR-//"`
echo "Running as pull-request build for pull-request ${PR_ID}."

if [ -z ${BITBUCKET_API_USERNAME+x} ]; then
    echo "BITBUCKET_API_USERNAME is unset"
    exit 1
fi

if [ -z ${BITBUCKET_API_PASSWORD+x} ]; then
    echo "BITBUCKET_API_PASSWORD is unset"
    exit 1
fi

if [ ! -z ${BLD_VCS_ROOT+x} ]; then
    vcsroot="${BLD_VCS_ROOT}"
elif [ ! -z ${BLD_VSC_ROOT+x} ]; then
    # Build service team has misspelled the variable
    vcsroot="${BLD_VSC_ROOT}"
fi

echo "VCS root: ${vcsroot}"

if [ -z ${BITBUCKET_PROJECT+x} ]; then
    echo "BITBUCKET_PROJECT is unset, extracting from vcsroot ${vcsroot}"

    # ssh://git@bitbucket.oci.oraclecorp.com:7999/sdk/bmc-sdk-swagger-validator.git
    # should become "SDK"

    if ! echo "${vcsroot}" | grep "^ssh://git@bitbucket.oci.oraclecorp.com:7999/" > /dev/null; then 
        echo "vcsroot does not start with 'ssh://git@bitbucket.oci.oraclecorp.com:7999/'"
        exit 1
    fi
    BITBUCKET_PROJECT=`echo ${vcsroot} | sed 's_^ssh://[^:]*:7999/\([^/]*\).*$_\1_' | tr '[a-z]' '[A-Z]'`
    echo "Using BITBUCKET_PROJECT=${BITBUCKET_PROJECT}"
fi
if [ -z ${BITBUCKET_REPO+x} ]; then
    echo "BITBUCKET_REPO is unset, extracting from vcsroot ${vcsroot}"

    # ssh://git@bitbucket.oci.oraclecorp.com:7999/sdk/bmc-sdk-swagger-validator.git
    # should become "bmc-sdk-swagger-validator"

    if ! echo "${vcsroot}" | grep "^ssh://git@bitbucket.oci.oraclecorp.com:7999/" > /dev/null; then 
        echo "vcsroot does not start with 'ssh://git@bitbucket.oci.oraclecorp.com:7999/'"
        exit 1
    fi
    BITBUCKET_REPO=`echo ${vcsroot} | sed 's_^ssh://[^:]*:7999/[^/]*/\(.*\)\.git$_\1_'`
    echo "Using BITBUCKET_REPO=${BITBUCKET_REPO}"
fi

url="https://bitbucket.oci.oraclecorp.com/rest/api/1.0/projects/${BITBUCKET_PROJECT}/repos/${BITBUCKET_REPO}/pull-requests/${PR_ID}/comments"

curl -k -u ${BITBUCKET_API_USERNAME}:${BITBUCKET_API_PASSWORD} \
    ${url} \
    -X POST \
    -H "X-Atlassian-Token: no-check" \
    -H "Content-Type: application/json" \
    -d "{\"text\":\"${TEXT}\"}"
