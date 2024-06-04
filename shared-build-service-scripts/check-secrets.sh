#!/bin/bash

set -e

SHARED_SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

if [ -z ${BITBUCKET_READONLY_PRIVATEKEY+x} ]; then
    echo "BITBUCKET_READONLY_PRIVATEKEY is unset"
else
    echo "BITBUCKET_READONLY_PRIVATEKEY is set, md5:"
    echo ${BITBUCKET_READONLY_PRIVATEKEY} | md5sum
fi

if [ -z ${BITBUCKET_READONLY_PASSPHRASE+x} ]; then
    echo "BITBUCKET_READONLY_PASSPHRASE is unset"
else
    echo "BITBUCKET_READONLY_PASSPHRASE is set, md5:"
    echo ${BITBUCKET_READONLY_PASSPHRASE} | md5sum
fi

if [ -z ${BITBUCKET_READWRITE_PRIVATEKEY+x} ]; then
    echo "BITBUCKET_READWRITE_PRIVATEKEY is unset"
else
    echo "BITBUCKET_READWRITE_PRIVATEKEY is set, md5:"
    echo ${BITBUCKET_READWRITE_PRIVATEKEY} | md5sum
fi

if [ -z ${BITBUCKET_READWRITE_PASSPHRASE+x} ]; then
    echo "BITBUCKET_READWRITE_PASSPHRASE is unset"
else
    echo "BITBUCKET_READWRITE_PASSPHRASE is set, md5:"
    echo ${BITBUCKET_READWRITE_PASSPHRASE} | md5sum
fi

if [ -z ${BITBUCKET_API_USERNAME+x} ]; then
    echo "BITBUCKET_API_USERNAME is unset"
else
    echo "BITBUCKET_API_USERNAME is set, md5:"
    echo ${BITBUCKET_API_USERNAME} | md5sum
fi

if [ -z ${BITBUCKET_API_PASSWORD+x} ]; then
    echo "BITBUCKET_API_PASSWORD is unset"
else
    echo "BITBUCKET_API_PASSWORD is set, md5:"
    echo ${BITBUCKET_API_PASSWORD} | md5sum
fi

if [ -z ${JIRA_USERNAME+x} ]; then
    echo "JIRA_USERNAME is unset"
else
    echo "JIRA_USERNAME is set, md5:"
    echo ${JIRA_USERNAME} | md5sum
fi

if [ -z ${JIRA_PASSWORD+x} ]; then
    echo "JIRA_PASSWORD is unset"
else
    echo "JIRA_PASSWORD is set, md5:"
    echo ${JIRA_PASSWORD} | md5sum
fi

git version
type -a git
