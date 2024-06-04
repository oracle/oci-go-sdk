#!/bin/bash

set -e
set -x

# Check if tag was set in the environment variable
if [ -z ${OCI_GO_SDK_VERSION_TAG+x} ]; then
    # No, point it at the master branch
    export OCI_GO_SDK_VERSION_TAG="master"
else
    # Check the format of the tag.
    if [[ ${OCI_GO_SDK_VERSION_TAG} =~ ^v[0-9][0-9]*\.[0-9][0-9]*\.[0-9][0-9]*$ ]]; then
        echo "Version tag: ${OCI_GO_SDK_VERSION_TAG}"
    else
        echo "Version tag: ${OCI_GO_SDK_VERSION_TAG} should look like 'v1.2.3'"
        exit 1
    fi
fi

# Kill godoc servers that might be running
set +e
killall godoc
set -e

version=$(echo ${OCI_GO_SDK_VERSION_TAG}|cut -d 'v' -f2)
find ../.. -name \*.go |xargs sed -i "s#https://docs\.cloud\.oracle\.com/en-us/iaas/tools/go-sdk-examples/latest/#https://docs\.cloud\.oracle\.com/en-us/iaas/tools/go-sdk-examples/${version}/#g"


# Check if port was set in the environment variable
if [ -z ${GODOC_PORT+x} ]; then
    # No, set it to port 6060
    export GODOC_PORT=6060
fi

echo "Running godoc on port ${GODOC_PORT}..."

godoc -http=:${GODOC_PORT} &
GODOC_PID=$!
echo "Web server PID is ${GODOC_PID}"

ATTEMPTS=24 # 240 seconds = 4 minutes max wait time for godoc server to have finished indexing
WAIT_TIME_SECONDS=10

n=0

# Temporarily turn "fail on error" off, because wget will report some 404s, but that's okay
set +e

until [ "$n" -ge $ATTEMPTS ]
do
    n=$((n+1))

    echo "Attempt $n of ${ATTEMPTS}:"
    echo "Waiting ${WAIT_TIME_SECONDS} seconds..."
    sleep ${WAIT_TIME_SECONDS}

    # Mirror the website, starting at /pkg/github.com/oracle/oci-go-sdk/
    # But only allow the /pkg/github.com/oracle/oci-go-sdk/ and the lib directories.
    if wget http://localhost:${GODOC_PORT}/pkg/github.com/oracle/oci-go-sdk/ ; then
        # This was successful
        n=0
        break
    fi
done

if [ "$n" -gt 0 ]; then
    echo "Failed to contact godoc webserver. Aborting..."
    exit 1
fi

echo "Successfully contacted godoc webserver. Mirroring..."

# Mirror the website, starting at /pkg/github.com/oracle/oci-go-sdk/
# But only allow the /pkg/github.com/oracle/oci-go-sdk/ and the lib directories.
wget -nv -m -k -erobots=off --no-host-directories --no-use-server-timestamps \
     --include-directories pkg/github.com/oracle/oci-go-sdk,lib \
     http://localhost:${GODOC_PORT}/pkg/github.com/oracle/oci-go-sdk/ \
    2>&1 | grep -v "Last-modified header missing"

# Turn "fail on error" back on
set -e


# Since we only have the directory for the oci-go-sdk, move the lib directory with
# the stylesheets and scripts into it. We replace the references below.
mv lib pkg/github.com/oracle/oci-go-sdk/

find . -type f -name \*.html -or -name \*.css

echo "Replacing links..."

function xargs_inplace_sed() {
    if [[ "$OSTYPE" == "darwin"* ]]; then
        xargs -n 10 sed -i '' "$@"
    else
        xargs -n 10 sed -i "$@"
    fi
}

# There are some links that point to localhost.
# Change them so they point to the public internet.
# (May not work on ONSR, but we can't mirror everything -- the entire Go documentation).
# We also remove the jquery stylesheets and scripts. They have 404s in them.
find . -type f -name \*.html -or -name \*.css \
    | xargs_inplace_sed \
        -e "s_http://localhost:${GODOC_PORT}/pkg/builtin_https://godoc.org/builtin_g" \
        -e "s_http://localhost:${GODOC_PORT}/pkg/_https://godoc.org/_g" \
        -e "s_http://localhost:${GODOC_PORT}/doc/_https://golang.org/_g" \
        -e "s_http://localhost:${GODOC_PORT}/search_https://golang.org/search_g" \
        -e "s_http://localhost:${GODOC_PORT}/blog/_https://blog.golang.org/_g" \
        -e "s_http://localhost:${GODOC_PORT}/help/_https://golang.org/help/_g" \
        -e "s_http://localhost:${GODOC_PORT}/project/_https://golang.org/project/_g" \
        -e "s_http://localhost:${GODOC_PORT}/LICENSE_https://golang.org/LICENSE_g" \
        -e "s_http://localhost:${GODOC_PORT}/src/github.com/oracle/oci-go-sdk/_https://github.com/oracle/oci-go-sdk/blob/${OCI_GO_SDK_VERSION_TAG}/_g" \
        -e 's_^<link.*jquery.*>$__' \
        -e 's_^<script.*jquery.*></script>$__' \
        -e "s_http://localhost:${GODOC_PORT}/_https://golang.org/_g" \
        -e "s_../../../../lib/godoc_lib/godoc_g"

# Change into the directory with the docs and zip them up at top level.
cd pkg/github.com/oracle/oci-go-sdk/
zip -r ../../../../oci-go-sdk-godoc.zip *
cd -

# Kill the godoc server.
kill ${GODOC_PID}
