SDK Regions Updater
===================

Description
-----------
This package contains scripts to add new realms/regions to SDK's. The tools support the following OCI SDK's:

* .NET SDK
* Go SDK
* Python SDK
* Ruby SDK
* Java SDK (not implemented)
* Typescript SDK (not implemented)

Installing dependencies
------------------------
To install dependencies for this project, from the root of the repo execute the following command:

    `pip install --trusted-host=artifactory.oci.oraclecorp.com -i https://artifactory.oci.oraclecorp.com/api/pypi/global-release-pypi/simple -U skclient`
    `pip install -r requirements.txt`

Running scripts locally
-------------------------
Before running the scripts, an environment variables must be set for each of SDK's:

* DotNetSDK_path=/path/to/oci-dotnet-sdk
* GoSDK_path=/path/to/go/src/github.com/oracle/oci-go-sdk
* PythonSDK_path=/path/to/python-sdk
* RubySDK_path=/path/to/ruby-sdk
* JavaSDK_path=/path/to/java-sdk
* TypescriptSDK_path=/path/to/oci-typescript-sdk

To update regions for one specific SDK, execute the following command from the root of the repo:

    `python -m sdk_regions_updater.sdk_regions_updater --SDK <sdk_name>`

For example, to update regions for .NET SDK, use command:

    `python -m sdk_regions_updater.sdk_regions_updater --SDK DotNetSDK`

In order to update regions for all supported SDK's, specify "All":

    `python -m sdk_regions_updater.sdk_regions_updater --SDK All`

Running the script without providing --SDK argument will show the usage information, including a list of allowed values for --SDK argument:

    `python -m sdk_regions_updater.sdk_regions_updater`