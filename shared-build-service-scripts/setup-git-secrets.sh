#!/bin/bash

# Usage:
#
# source setup-git-secrets.sh
#
# or
#
# . setup-git-secrets.sh
#
# Since the ssh-agent must run in order to use private keys
# with passphrase, it is not enough to just run this script.
# It has to be "sourced" (see https://linuxize.com/post/bash-source-command/).
# Otherwise, the ssh-agent stops running at the end of this script,
# and later ssh or git commands will fail.

set -e
set -x

SHARED_SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

if [ -z ${BITBUCKET_READONLY_PRIVATEKEY+x} ]; then
    echo "BITBUCKET_READONLY_PRIVATEKEY is unset"
    exit 1
else
    echo "BITBUCKET_READONLY_PRIVATEKEY is set"
fi
if [ -z ${BITBUCKET_READONLY_PASSPHRASE+x} ]; then
    echo "BITBUCKET_READONLY_PASSPHRASE is unset"
    exit 1
else
    echo "BITBUCKET_READONLY_PASSPHRASE is set"
fi

mkdir -p ~/.ssh

# Disable 'set -x', otherwise we log the contents of the private key
old_set_x=${-//[^x]/}
set +x
printf '%s' "${BITBUCKET_READONLY_PRIVATEKEY}" > ~/.ssh/id_rsa
if [[ -n "$old_set_x" ]]; then set -x; fi

chmod 600 ~/.ssh/*

ssh-keyscan -p 7999 -t rsa bitbucket.oci.oraclecorp.com >> ~/.ssh/known_hosts

if [ ! -z ${SSH_AGENT_PID+x} ] && ps -p ${SSH_AGENT_PID} > /dev/null
then
  	echo "ssh-agent is already running, pid ${SSH_AGENT_PID}"
else
	eval `ssh-agent -s`
  	echo "Started ssh-agent, pid ${SSH_AGENT_PID}"
fi

echo "Adding ~/.ssh/id_rsa with passphrase coming from BITBUCKET_READONLY_PASSPHRASE"
SSH_ASKPASS=${SHARED_SCRIPT_DIR}/ssh_give_pass.sh ssh-add ~/.ssh/id_rsa <<< "${BITBUCKET_READONLY_PASSPHRASE}"

echo "Added ~/.ssh/id_rsa"

ssh-add -L
