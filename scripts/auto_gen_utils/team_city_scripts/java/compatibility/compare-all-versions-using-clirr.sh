#!/bin/bash

# Syntax: $0 [--startWithVersion <version>] [--javaSdkDir <java-sdk-root-dir>] <maven-metadata-xml-file>
#
# --startWithVersion <version>
#       Skips versions older than <version>, meaning the first version çompared is <version>
#       to the version just preceding it.
#       If not provided, will start comparing the oldest version to the 2nd oldest version.
#
# --javaSdkDir <java-sdk-root-dir>
#       The directory where the GitHub version of the OCI Java SDK has been checked out
#       If not provided, uses the current directory.
#

# If you are running this on macOS `brew install gsed`.

if uname -a | grep -i "Darwin" > /dev/null; then
    if ! which gsed > /dev/null; then
        echo "If you are running this on macOS, please install gsed: 'brew install gsed'."
        exit 1
    fi
fi    

set -e
#set -x

# Some operating systems (including macOS) do not have
# the correct `sed` version that supports `-i<backupSuffix>` (without space).
# There, it is called `gsed`. The if statement below makes the script
# sed/gsed-agnostic.
if [ -z ${SED+x} ]; then
    if which gsed > /dev/null; then
        SED=`which gsed`
    else
        SED=`which sed`
    fi
fi

java_sdk_root_dir="."

while [[ $# -gt 1 ]]; do
    key="$1"

    case $key in
        --startWithVersion)
        start_with_version="$2"
        shift # past argument
        shift # past value
        ;;
        --javaSdkDir)
        java_sdk_root_dir="$2"
        shift # past argument
        shift # past value
        ;;
        *)    # unknown option
        echo "Unknown option: $key"
        echo "Syntax: $0 [--startWithVersion <version>] [--javaSdkDir <java-sdk-root-dir>] <maven-metadata-xml-file>"
        exit 1
        ;;
    esac
done

if [ $# -lt 1 ]; then
    echo "Syntax: $0 [--startWithVersion <version>] [--javaSdkDir <java-sdk-root-dir>] <maven-metadata-xml-file>"
    exit 1
fi
if [ $# -gt 1 ]; then
    echo "Syntax: $0 [--startWithVersion <version>] [--javaSdkDir <java-sdk-root-dir>] <maven-metadata-xml-file>"
    exit 1
fi

clirr="clirr:2.9.2-oracle-SNAPSHOT"
maven_metadata_xml_file="$1"

tempfileprefix=`basename $0`
# filtered_maven_metadata_xml_file=$(mktemp -t ${tempfileprefix})
filtered_maven_metadata_xml_file="${maven_metadata_xml_file}.filtered"

# skip the first line, it's the latest version: $SED '1,1d'
grep -o "<version>.*</version>" ${maven_metadata_xml_file} \
    | $SED 's_<version>\([^<]*\)</version>_\1_' \
    | $SED '1,1d' \
    | grep -v -i "beta" \
    | grep -v -i "experimental" \
    | grep -v -i "snapshot" \
    | grep -v -i "preview" \
    | grep -v -i "401stream" \
    > ${filtered_maven_metadata_xml_file}

temp_file=$(mktemp -t "${tempfileprefix}-temp")

run_dir=`pwd`
cd $java_sdk_root_dir

for version in `cat ${filtered_maven_metadata_xml_file}`; do
    if [ -z ${previous_version+x} ]; then
        # previous version unset
        previous_version="${version}"
        version_range_beginning="${version}"
        continue
    fi
    if [ -n "${start_with_version}" ]; then
        # start_with_version set, skip until ${version} is ${start_with_version}
        if [[ "${version}" != "${start_with_version}" ]]; then
            # not the desired version yet
            echo "Skipping ${version}"
            previous_version="${version}"
            version_range_beginning="${version}"
            continue
        fi
        echo "Starting with version ${start_with_version}"
        unset start_with_version
    fi

    echo "##################################################"
    echo "##################################################"
    echo "Comparing ${previous_version} to ${version}"

    git_tag="v${version}"
    echo -e "\tChecking out git tag ${git_tag}..."
    if ! git checkout "${git_tag}" > ${temp_file} 2>&1; then
        echo -e "ERROR: \tgit checkout ${git_tag} failed"
        cat ${temp_file}
        echo -e "\tCompatible version range: ${version_range_beginning} to ${previous_version}"
        unset previous_version
        continue
    fi

    git_date=`git log -n 1 --pretty=format:"%cd"`
    echo -e "\tVersion ${version} created ${git_date}"

    echo -e "\tRunning clirr..."
    clirr_output_file="${run_dir}/clirr-${version}.txt"
    if ! mvn ${clirr}:check \
        -f bmc-common/pom.xml \
        -DcomparisonVersion="${previous_version}" \
        -DtextOutputFile="${clirr_output_file}" \
        -DfailOnError="false" \
        -DfailOnWarning="false" \
        > ${temp_file} 2>&1; then
        echo "mvn clirr:check failed"
        cat ${temp_file}
        rm ${temp_file}
        exit 1
    fi
    echo -e "\tGenerated ${clirr_output_file}"
    if [ -s ${clirr_output_file} ]; then
        echo -e "\tClirr detected changes between ${previous_version} and ${version}"
        echo -e "\tCompatible version range: ${version_range_beginning} to ${previous_version}"
        version_range_beginning="${version}"
    fi

    previous_version="${version}"
done

echo -e "\tCompatible version range: ${version_range_beginning} to ${version}"

cd $run_dir

rm ${temp_file}