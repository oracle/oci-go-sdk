#!/bin/bash

## This script recieves
## 1. A setting for the private key
## 2. A setting for the make file name
## 3. A setting for make file target
## 4. A setting to run the make targer w or without debug

set -e ; set -x ;
echo $GOPATH
type -a go
type -a golint
type -a gofmt

# Inject the key into the key file
set +x 
cp ./go_sdk_test_user_key.pem $TF_VAR_private_key_path
set -x

#Run the test
# cd $GOPATH/src/github.com/oracle/oci-go-sdk/
cd ${SOURCE_DIR}
go mod tidy