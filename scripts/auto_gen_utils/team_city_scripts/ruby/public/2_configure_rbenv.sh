#!/bin/bash

# NOTE: The pom.xml build step (#3) in TC needs an environment variable configured as below
# for the build step to pick up the correct version of ruby with its installed gems from
# this script:
#    env.PATH => /home/teamcity/.rbenv/shims:%env.PATH%

RUBY_VERSION_TO_INSTALL=2.7.3
if [ ! -z "$1" ]; then
    RUBY_VERSION_TO_INSTALL=$1
fi

if [ ! -d "$HOME/.rbenv" ]; then
    echo Setting up rbenv and ruby-build to install the Ruby SDK locally

  	# Set up rbenv
  	git clone https://github.com/rbenv/rbenv.git ~/.rbenv
  	export PATH="$HOME/.rbenv/bin:$PATH"
  	eval "$(rbenv init -)"

  	# Install ruby-build so we can install Rubies
  	mkdir -p "$(rbenv root)"/plugins
  	git clone https://github.com/rbenv/ruby-build.git "$(rbenv root)"/plugins/ruby-build
fi

if [ -d "$HOME/.rbenv" ]; then
    export PATH="$HOME/.rbenv/bin:$PATH"
    eval "$(rbenv init -)"
fi

# In latest OL7.9 image, install Ruby will fail due to DTrace, temporarly disable it
if [ -z $(rbenv versions | grep "$RUBY_VERSION_TO_INSTALL") ]; then
    echo "Installing Ruby Version $RUBY_VERSION_TO_INSTALL"
    RUBY_CONFIGURE_OPTS="--disable-dtrace" rbenv install -s $RUBY_VERSION_TO_INSTALL
fi

if [ ! -f $HOME/.rbenv/version ] || [ $(cat $HOME/.rbenv/version) != $RUBY_VERSION_TO_INSTALL ]; then
    echo "Setting $RUBY_VERSION_TO_INSTALL as local Ruby version"
    rbenv local $RUBY_VERSION_TO_INSTALL
fi

echo "Ruby version"
which ruby
ruby --version

# Change to the ruby-sdk and install its dev dependencies
cd ruby-sdk
echo "Installing dependencies for $(pwd)"

# Now install the ruby dependencies
echo "Installing Ruby Dependencies"
gem install bundler
rbenv rehash
bundle install

# Verify gems
gem list

# Verify PATH
echo "Path: $PATH"

cd ..