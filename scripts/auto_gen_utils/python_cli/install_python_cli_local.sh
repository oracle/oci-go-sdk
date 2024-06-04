#!/usr/bin/env bash

find . -name '*.pyc' -delete
pip install --pre --trusted-host artifactory.oci.oraclecorp.com --index-url https://artifactory.oci.oraclecorp.com/api/pypi/global-dev-pypi/simple -r requirements.txt
pip install --trusted-host=artifactory.oci.oraclecorp.com -e .
