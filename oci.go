
/*
Package oci This is a preview of the official go sdk for Oracle Cloud Infrastructure

	To avoid breaking changes please consider using https://github.com/golang/dep, or vendoring this sdk.
	Although we make a conscious effort to not push breaking changes, sometimes they are needed. This is particularly true at the current stage
	of development.

Working with the SDK

Generally speaking, you can start working with the sdk by importing the service package, creating a client,
and making calls with client.

Configuring

You can configure the sdk with your credentials by creating a settings file in: $HOME/.oci/config that looks like the following:
	 [DEFAULT]
	 user=[user ocid]
	 fingerprint=[fingerprint]
	 key_file=[path to pem file containing private key]
	 tenancy=[tenancy ocid]
	 region=[region for your tenancy]

and calling the common.DefaultConfigProvider() function like so:

	 // Do not forget to import the necessary packages
	 import (
		"github.com/oracle/oci-go-sdk/common"
		"github.com/oracle/oci-go-sdk/identity" // Identity or any other service you wish to call
	)

	configProvider := common.DefaultConfigProvider()

You can also configure the sdk programmatically by implementing the `ConfigurationProvider` interface
	// ConfigurationProvider wraps information about the account owner
	type ConfigurationProvider interface {
		KeyProvider
		TenancyOCID() (string, error)
		UserOCID() (string, error)
		KeyFingerprint() (string, error)
		Region() (string, error)
	}


Making a request

Making request can be achieved by: creating a client for the service that you wish to work with, followed by calling
the a function in the above client
- Creating a client: all packages provide a function to create clients. It is of the form `NewXXXClientWithConfigurationProvider`.
You can create a new client by passing a struct that conforms to the `ConfigurationProvider` interface.
Here is a quick example that shows how to create a client
	config := common.DefaultConfigProvider()
	client, err := identity.NewIdentityClientWithConfigurationProvider(config)
	if err != nil {
		 panic(err)
	}

- Making calls: after successfully creating client, you can now make calls to the service. Generally all operations
functions take in  a context.Context (https://golang.org/pkg/context/) and a struct that wraps all input parameters. The return is usually a response struct
containing the desired data,  and an error struct describing the error if there was one, eg:
	response, err := client.GetGroup(context.Background(), identity.GetGroupRequest{GroupId:id})
	if err != nil {
		//Something happened
		panic(err)
	}
	//Process the data in response struct
	fmt.Println("Group's name is:", response.Name)

Organization of the SDK

The oci-go-sdk is broken down into:

- service packages: all packages except `common` and any other package found inside `cmd`. They represent Oracle Cloud Infrastructure services supported by the go sdk.
Each package represents a service. These packages include methods to interact with the service, structs that model
input and output parameters and a client struct that acts as receiver for the above methods. All code in these packages
is machine generated.

- common package: found in the `common` directory. The common package provides supporting functions and structs used by service packages,
including: http request/response (de)serialization, request signing, json parsing, pointer to reference and other helper functions. Thought
there are some functions in this package that are meant for user consumption, most of them are meant to be used by the service packages.

- cmd: internal tools used by the `oci-go-sdk`
 */
package oci

//go:generate go run cmd/genver/main.go cmd/genver/version_template.go --output common/version.go
