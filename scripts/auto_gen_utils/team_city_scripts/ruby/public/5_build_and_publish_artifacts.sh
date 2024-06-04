#!/bin/bash

set -e
set -x

# must disable StrictHostKeyChecking so that we don't get an interactive
# prompt later asking to confirm the host key
# Must disable -e (fail on non-zero exit code) because ssh returns 255
set +e
ssh -o StrictHostKeyChecking=no git@bitbucket.oci.oraclecorp.com -p 7999
set -e

# Old way of doing that:
# ls -la ~/.ssh
#
# cat ~/.ssh/config
# 
# printf "\n\nHost * \n  StrictHostKeyChecking no\n" >> ~/.ssh/config
# 
# cat ~/.ssh/config

# configure git for this commit
git config --global user.email "$GIT_USER_EMAIL"
git config --global user.name "$GIT_USER_NAME"

## Maven version from build used to name the dist zip archive ##
echo Maven Version $MAVEN_VERSION

## Verify rbenv setup ##
echo Ruby Version
ruby --version

echo Gem Version
gem --version

echo Gem env
gem env

# Install extra gems
echo Install typhoeus
gem install typhoeus -v 1.0.2 --user-install

echo Install inifile
gem install inifile -v 3.0.0 --user-install

echo Install minitest
gem install minitest -v 5.8.3 --user-install

echo Install rake
gem install rake -v 10.4.2 --user-install

echo Install bundle
gem install bundler -v 2.3.27 --user-install

echo Python version
python --version

cd ruby-sdk
bundle install

# This will copy the code to a specific region variant (r2) to process the links for the docs
cd scripts
python pre_process.py r2
echo Pre-processing complete

# Now build the gem for the r2 variant
cd ../variants/r2

echo Building gem
gem build oci.gemspec || { echo 'Failed to build oci gem (probably a test failure).' ; exit 1; }

version=$(grep "VERSION" lib/oci/version.rb | cut -d\' -f2)
find . -name \*.rb |xargs sed -i "s#https://docs\.cloud\.oracle\.com/en-us/iaas/tools/ruby-sdk-examples/latest/#https://docs\.cloud\.oracle\.com/en-us/iaas/tools/ruby-sdk-examples/$version/#g"

echo Buiding RDocs
ruby <<EOF
require 'yard'
require 'os'
files = File.open('../../.yarddoc-files').readlines.map(&:strip)
YARD::CLI::CommandParser.run('doc', *files, '-m', 'markdown')
EOF

echo Creating the Zip file
ls -la scripts/package*.sh
. scripts/package-oci.sh

echo Current directory
pwd

echo Contents of dist folder
ls -la dist

# Return to base directory from ruby-sdk/variants/r2
cd ../../..

