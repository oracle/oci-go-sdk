#!/bin/bash

set -e
set -u

# Syntax: update-codegen-version.sh 1.49-SNAPSHOT path/to/sdks branch-prefix description
#
# Underneath the path/to/sdks, we should see the following:
# .
# ├── java-sdk
# ├── python-cli
# ├── python-sdk
# ├── ruby-sdk
# ├── oci-typescript-sdk
# └── src
#     └── github.com
#         └── oracle
#             └── oci-go-sdk

if [ $# -ne 4 ]; then
    echo "Illegal number of parameters: need 4 parameters"
    echo "Syntax: $0 <version> <path-to-sdks> <branch-prefix> <description>"
    exit 1
fi

version="$1"
rootpath="$2"
prefix="$3"
desc="$4"

cd $rootpath

if [ ! -f java-sdk/bmc-codegen/pom.xml ]; then
    echo "java-sdk/bmc-codegen/pom.xml not found"
fi
if [ ! -f ruby-sdk/pom.xml ]; then
    echo "ruby-sdk/pom.xml not found"
fi
if [ ! -f src/github.com/oracle/oci-go-sdk/pom.xml ]; then
    echo "src/github.com/oracle/oci-go-sdk/pom.xml not found"
fi
if [ ! -f python-sdk/pom.xml ]; then
    echo "python-sdk/pom.xml not found"
fi
if [ ! -f python-cli/pom.xml ]; then
    echo "python-cli/pom.xml not found"
fi
if [ ! -f oci-typescript-sdk/codegen/pom.xml ]; then
    echo "oci-typescript-sdk/codegen/pom.xml not found"
fi
if [ ! -f legacy-java-sdk/bmc-codegen/pom.xml ]; then
    echo "legacy-java-sdk/bmc-codegen/pom.xml not found"
fi

checkout_and_update () {
    source_branch="$1"

    git checkout "${source_branch}"
    git reset --hard
    git pull origin "${source_branch}"

    sed -i .bak -e "/bmc-sdk-swagger-maven-plugin/ {" -e "n; s/>[^<]*</>${version}</" -e "}" pom.xml
    rm pom.xml.bak
}

commit () {
    dest_branch="$1"
    description="$2"

    git status

    set +e # temporarily allow errors
    git branch -D "${dest_branch}"
    set -e
    git checkout -B "${dest_branch}"

    git add -A

    git commit -a -m "${description}" --allow-empty
}

echo "========================"
echo "===       JAVA       ==="
echo "========================"

preview_branch="${prefix}-java-preview"
master_branch="${prefix}-java-master"

cd java-sdk/bmc-codegen

git fetch

echo "========================"
echo "===   JAVA PREVIEW   ==="
echo "========================"

checkout_and_update preview
mvn clean process-sources -Pcodegen
commit "${preview_branch}" "${desc} (Java Preview)"
java_preview_pr=`git push origin +"${preview_branch}" 2>&1 | grep -A 1 "remote: Create pull request for" | tail -n 1`

echo "========================"
echo "===   JAVA MASTER    ==="
echo "========================"

checkout_and_update master
mvn clean process-sources -Pcodegen
commit "${master_branch}" "${desc} (Java Master)"
java_master_pr=`git push origin +"${master_branch}" 2>&1 | grep -A 1 "remote: Create pull request for" | tail -n 1`

cd ../..

echo "========================"
echo "===       RUBY       ==="
echo "========================"

preview_branch="${prefix}-ruby-preview"
master_branch="${prefix}-ruby-master"

cd ruby-sdk

git fetch

echo "========================"
echo "===   RUBY PREVIEW   ==="
echo "========================"

checkout_and_update preview
mvn clean install
commit "${preview_branch}" "${desc} (Ruby Preview)"
ruby_preview_pr=`git push origin +"${preview_branch}" 2>&1 | grep -A 1 "remote: Create pull request for" | tail -n 1`

echo "========================"
echo "===   RUBY MASTER    ==="
echo "========================"

checkout_and_update master
mvn clean install
commit "${master_branch}" "${desc} (Ruby Master)"
ruby_master_pr=`git push origin +"${master_branch}" 2>&1 | grep -A 1 "remote: Create pull request for"  | tail -n 1`

cd ..

echo "========================"
echo "===        GO        ==="
echo "========================"

preview_branch="${prefix}-go-preview"
master_branch="${prefix}-go-master"

export GOPATH="${PWD}"
cd src/github.com/oracle/oci-go-sdk

git fetch

echo "========================"
echo "===    GO PREVIEW    ==="
echo "========================"

checkout_and_update preview
make -f MakefileDevelopment.mk release
commit "${preview_branch}" "${desc} (Go Preview)"
go_preview_pr=`git push origin +"${preview_branch}" 2>&1 | grep -A 1 "remote: Create pull request for" | tail -n 1`


echo "========================"
echo "===     GO MASTER    ==="
echo "========================"

checkout_and_update master
make -f MakefileDevelopment.mk release
commit "${master_branch}" "${desc} (Go Master)"
go_master_pr=`git push origin +"${master_branch}" 2>&1 | grep -A 1 "remote: Create pull request for"  | tail -n 1`

cd ../../../..

echo "========================"
echo "===     PYTHON       ==="
echo "========================"

preview_branch="${prefix}-python-preview"
master_branch="${prefix}-python-master"

cd python-sdk

git fetch

echo "========================"
echo "===  PYTHON PREVIEW  ==="
echo "========================"

checkout_and_update preview
mvn clean install
commit "${preview_branch}" "${desc} (Python Preview)"
python_preview_pr=`git push origin +"${preview_branch}" 2>&1 | grep -A 1 "remote: Create pull request for" | tail -n 1`


echo "========================"
echo "===   PYTHON MASTER  ==="
echo "========================"

checkout_and_update master
mvn clean install
commit "${master_branch}" "${desc} (Python Master)"
python_master_pr=`git push origin +"${master_branch}" 2>&1 | grep -A 1 "remote: Create pull request for"  | tail -n 1`

cd ..

echo "========================"
echo "===       CLI        ==="
echo "========================"

preview_branch="${prefix}-cli-preview"
master_branch="${prefix}-cli-master"

cd python-cli

git fetch

echo "========================"
echo "===    CLI PREVIEW   ==="
echo "========================"

checkout_and_update preview
mvn clean install
commit "${preview_branch}" "${desc} (CLI Preview)"
cli_preview_pr=`git push origin +"${preview_branch}" 2>&1 | grep -A 1 "remote: Create pull request for" | tail -n 1`


echo "========================"
echo "===    CLI MASTER    ==="
echo "========================"

checkout_and_update master
mvn clean install
commit "${master_branch}" "${desc} (CLI Master)"
cli_master_pr=`git push origin +"${master_branch}" 2>&1 | grep -A 1 "remote: Create pull request for"  | tail -n 1`

cd ..

echo "========================"
echo "===  PULL REQUESTS   ==="
echo "========================"

echo "Java:"
echo "${java_preview_pr}"
echo "${java_master_pr}"
echo "Ruby:"
echo "${ruby_preview_pr}"
echo "${ruby_master_pr}"
echo "Go:"
echo "${go_preview_pr}"
echo "${go_master_pr}"
echo "Python:"
echo "${python_preview_pr}"
echo "${python_master_pr}"
echo "CLI:"
echo "${cli_preview_pr}"
echo "${cli_master_pr}"
