# Oracle Cloud Infrastructure Golang SDK ***Preview***
[![wercker status](https://app.wercker.com/status/09bc4818e7b1d70b04285331a9bdbc41/s/master "wercker status")](https://app.wercker.com/project/byKey/09bc4818e7b1d70b04285331a9bdbc41)

This is a ***preview*** of the official go sdk for Oracle Cloud Infrastructure
> ***WARNING:***: To avoid breaking changes please consider using [go dep](https://github.com/golang/dep), or vendoring this sdk.
Although we make a conscious effort to not push breaking changes, sometimes they are needed. This is particularly true at the current stage 
of development.

## Dependencies
- Install [Golang](https://golang.org/dl/)
- Install [make](https://www.gnu.org/software/make/), through your favorite package manager or binary 
distribution of choice

Additionally, if you are planning to  build or develop the sdk you need the following dependencies
- Install [github.com/stretchr/testify](https://github.com/stretchr/testify)
```sh
go get github.com/stretchr/testify
```


- Install [go lint](https://github.com/golang/lint)
```
go get -u github.com/golang/lint/golint
```
 


## Installing
Simply clone this repo into your go sdk you can use the following command

```
git clone git@github.com:oracle/oci-go-sdk.git  $GOPATH/src/github.com/oracle
```

## Working with the SDK
Generally speaking, you can start working with the sdk by importing the service package, creating a client, 
and making calls with client.

### Configuring 
You can configure the sdk with your credentials by creating a settings file in:
 `$HOME/.oci/config` that looks like the following:
 ```
 [DEFAULT]
 user=[user ocid]
 fingerprint=[fingerprint]
 key_file=[path to pem file containing private key]
 tenancy=[tenancy ocid]
 region=[region for your tenancy]
 ```
 and calling the `common.DefaultConfigProvider()` function like so:
 ```go
 // Do not forget to import the necessary packages
 import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/identity" // Identity or any other service you wish to call
)
 
 //...
 
configProvider := common.DefaultConfigProvider()
```
 You can also configure the sdk programmatically by implementing the `ConfigurationProvider` interface
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
Making request can be achieved by: creating a client for the service that you wish to work with, followed by calling
the a function in the above client 
- *Creating a client*: all packages provide a function to create clients. It is of the form `NewXXXClientWithConfigurationProvider`.
You can create a new client by passing a struct that conforms to the `ConfigurationProvider` interface.
Here is a quick example that shows how to create a client
```go
config := common.DefaultConfigProvider()
client, err := identity.NewIdentityClientWithConfigurationProvider(config)
if err != nil { 
     panic(err)
}
```

- *Making calls*: after successfully creating client, you can now make calls to the service. Generally all operations
functions take in  a [`context.Context`](https://golang.org/pkg/context/) and a struct that wraps all input parameters. The return is usually a response struct
containing the desired data,  and an error struct describing the error if there was one, eg:
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
The `oci-go-sdk` can be broken down into:
- **service packages**: all packages except `common` and any other package found inside `cmd`. They represent Oracle Cloud Infrastructure services supported by the go sdk. 
Each package represents a service. These packages include methods to interact with the service, structs that model 
input and output parameters and a client struct that acts as receiver for the above methods. All code in these packages
is machine generated.

- **common package**: found in the `common` directory. The common package provides supporting functions and structs used by service packages, 
including: http request/response (de)serialization, request signing, json parsing, pointer to reference and other helper functions. Thought
there are some functions in this package that are meant for user consumption, most of them are meant to be used by the service packages.

- **cmd**: internal tools used by the `oci-go-sdk`

## Examples

## Documentation
You can find the documetation of this sdk in the [godocs site]()


## Building and testing
Building is provided by the make file at the root of the project. To build the project execute

```
make build
```

To execute the tests:
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

