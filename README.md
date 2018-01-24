# Oracle Cloud Infrastructure Golang SDK ***Preview***
[![wercker status](https://app.wercker.com/status/09bc4818e7b1d70b04285331a9bdbc41/s/master "wercker status")](https://app.wercker.com/project/byKey/09bc4818e7b1d70b04285331a9bdbc41)

This is a ***preview*** of the official Go SDK for Oracle Cloud Infrastructure. During this stage of development, we need to push breaking changes <--- come back to this
> ***WARNING:***: To avoid breaking changes please consider using the [Go depenendency management tool](https://github.com/golang/dep), or vendoring this SDK.

## Dependencies
- Install [Golang](https://golang.org/dl/)
- Install [GNU Make](https://www.gnu.org/software/make/), using the package manager or binary distribution tool appropiate for your platform

Additionally, if you are planning to  build or develop the SDK you need to install the following dependencies:
- Install [github.com/stretchr/testify](https://github.com/stretchr/testify)
```sh
go get github.com/stretchr/testify
```


- Install [go lint](https://github.com/golang/lint)
```
go get -u github.com/golang/lint/golint
```
 


## Installing
Use the following command to install this SDK.

```
go get -u github.com/oracle/oci-go-sdk
```
Alternatively you can git clone this repo

## Working with the Go SDK
To start working with the Go SDK you import the service package, create a client, and then use that client to make calls

### Configuring 
To configure the Go SDK with your credentials, create the settings file: `$HOME/.oci/config`.
Add the following to the settings file

 ```
 [DEFAULT]
 user=[user ocid]
 fingerprint=[fingerprint]
 key_file=[path to pem file containing private key]
 tenancy=[tenancy ocid]
 region=[region for your tenancy]
 ```
Call the `common.DefaultConfigProvider()` function as follows:

 ```go
 // Import necessary packages
 import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/identity" // Identity or any other service you wish to call
)
 
 //...
 
configProvider := common.DefaultConfigProvider()
```

 To configure the SDK programmatically, implement the `ConfigurationProvider` interface
 ```go
// ConfigurationProvider wraps information about the account owner
type ConfigurationProvider interface {
	KeyProvider
	TenancyOCID() (string, error)
	UserOCID() (string, error)
	KeyFingerprint() (string, error)
	Region() (string, error)
}
```

### Making a request
To make a request, create a client for the service and then call a function from that client.

- *Creating a client*: All packages provide a function to create clients. It is of the form `New<ServiceName>ClientWithConfigurationProvider`,
such as `NewVirtualNetworkClientWithConfigurationProvider` or `NewIdentityClientWithConfigurationProvider`. To create a new client, 
pass a struct that conforms to the `ConfigurationProvider` interface, or use the `DefaultConfigProvider()` function in the common package.

For example: 
```go
config := common.DefaultConfigProvider()
client, err := identity.NewIdentityClientWithConfigurationProvider(config)
if err != nil { 
     panic(err)
}
```

- *Making calls*: After successfully creating client, you can now make calls to the service. Generally all functions associated with an operation
accept [`context.Context`](https://golang.org/pkg/context/) and a struct that wraps all input parameters. The functions then return a response struct
that contains the desired data, and an error struct that describes the error if an error ocurrs.

For example:
```go
response, err := client.GetGroup(context.Background(), identity.GetGroupRequest{GroupId:id})
if err != nil {
	//Something happened
	panic(err)
}
//Process the data in response struct
fmt.Println("Group's name is:", response.Name)
```

## Organization of the SDK
The `oci-go-sdk` contains the following:
- **Service packages**: All packages except `common` and any other package found inside `cmd`. These packages represent 
the Oracle Cloud Infrastructure services supported by the Go SDK. Each package represents a service. 
These packages include methods to interact with the service, structs that model 
input and output parameters and a client struct that acts as receiver for the above methods. All code in these packages is machine generated.

- **Common package**: Found in the `common` directory. The common package provides supporting functions and structs used by service packages, 
including: HTTP request/response (de)serialization, request signing, json parsing, pointer to reference and other helper functions. Most of the functions
in this package are meant are meant to be used by the service packages.

- **cmd**: Internal tools used by the `oci-go-sdk`

## Examples

## Documentation
See [godocs site]() for the Go SDK documentation


## Building and testing
Building is provided by the make file at the root of the project. To build the project execute

```
make build
```

To run the tests:
```
make test
```

## Questions, feedback or getting help
* [Stack Overflow](https://stackoverflow.com/): Please use the [oracle-cloud-infrastructure](https://stackoverflow.com/questions/tagged/oracle-cloud-infrastructure) and [oci-ruby-sdk](https://stackoverflow.com/questions/tagged/oci-ruby-sdk) tags in your post
* [Developer Tools section](https://community.oracle.com/community/cloud_computing/bare-metal/content?filterID=contentstatus%5Bpublished%5D~category%5Bdeveloper-tools%5D&filterID=contentstatus%5Bpublished%5D~objecttype~objecttype%5Bthread%5D) of the Oracle Cloud forums
* [My Oracle Support](https://support.oracle.com)


## Contributing
See [CONTRIBUTING](/CONTRIBUTING.md) for details

## License
Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.

This SDK and sample is dual licensed under the Universal Permissive License 1.0 and the Apache License 2.0.

See [LICENSE](/LICENSE.txt) for more details.

