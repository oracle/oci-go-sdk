
/*
Package oci This is a preview of the official go sdk for Oracle Cloud Infrastructure

	To avoid breaking changes please consider using https://github.com/golang/dep, or vendoring this sdk.
	Although we make a conscious effort to not push breaking changes, sometimes they are needed. This is particularly true at the current stage
	of development.

Installation

Use the following command to install this SDK:

	go get -u github.com/oracle/oci-go-sdk


Configuration

Configure the Go SDK with your credentials, see SDK and Tool Configuration https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/sdkconfig.htm for instructions.

Once you have configured the Go SDK with your credentials, call `common.DefaultConfigProvider()` function as follows:

	 // Import necessary packages
	 import (
		"github.com/oracle/oci-go-sdk/common"
		"github.com/oracle/oci-go-sdk/identity" // Identity or any other service you wish to make requests to
	)

	 //...

	configProvider := common.DefaultConfigProvider()

To configure the SDK programmatically, implement the `ConfigurationProvider` interface shown below:
	// ConfigurationProvider wraps information about the account owner
	type ConfigurationProvider interface {
		KeyProvider
		TenancyOCID() (string, error)
		UserOCID() (string, error)
		KeyFingerprint() (string, error)
		Region() (string, error)
	}

Or you can use the `NewRawConfigurationProvider` function that accepts all mandatory fields of the ConfigurationProvider interface
as input

	passphrase := "password"
	myConfigProvider := NewRawConfigurationProvider("tenencyOcid", "userOcid", "someRegion", "fingerprintOfKey", &passphrase)



Quickstart

Making request can be achieved by: creating a client for the service that you wish to work with, followed by calling
the a function in the above client

Creating a client: all packages provide a function to create clients. It is of the form `NewXXXClientWithConfigurationProvider`.
You can create a new client by passing a struct that conforms to the `ConfigurationProvider` interface.
Here is a quick example that shows how to create a client
	config := common.DefaultConfigProvider()
	client, err := identity.NewIdentityClientWithConfigurationProvider(config)
	if err != nil {
		 panic(err)
	}

Making calls: after successfully creating client, you can now make calls to the service. Generally all operations
functions take in  a context.Context (https://golang.org/pkg/context/) and a struct that wraps all input parameters. The return is usually a response struct
containing the desired data,  and an error struct describing the error if there was one, eg:
	response, err := client.GetGroup(context.Background(), identity.GetGroupRequest{GroupId:id})
	if err != nil {
		//Something happened
		panic(err)
	}
	//Process the data in response struct
	fmt.Println("Group's name is:", response.Name)

Most oci resources have a CRUD operations and usually take a mix of mandatory or optional resources. For example:

	compartmentID := "compartmentOCID"
	request := core.CreateVcnRequest{}
	request.CidrBlock = common.String("10.0.0.0/16") // The CIDR IP address block of the VCN.
	request.CompartmentId = common.String(compartmentID)

	c, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		panic(err)
 	}

	r, err := c.CreateVcn(context.Background(), request)
	if err != nil {
		panic(err)
 	}

	// Do some work
	// Now delete the VCN
	_, err = c.DeleteVcn(context.Background(), core.DeleteVcnRequest{VcnId: r.Id})

For a full example of these operations, head to the examples directory in the oci-go-sdk repo: https://github.com/oracle/oci-go-sdk/tree/master/example

Optional fields in the SDK

The sdk will omit all optional fields that are nil. In the case of enum-type fields, the sdk will omit fields whose value
is the empty string.

Signer and Non-standard calls

The oci-go-sdk exposes a stand-alone signer that can be used to sign custom requests. For example:

	var request http.Request
	request = ... // some custom request

	//Mandatory headers to be used in the sign process
	defaultGenericHeaders    = []string{"date", "(request-target)", "host"}

	//Optional headers
	optionalHeaders = []string{"content-length", "content-type", "x-content-sha256"}

	// A predicate that specifies when to use the optional signing headers
	optionalHeadersPredicate := func (r *http.Request) bool {
		return r.Method == http.MethodPost
	}

	//And a provider of cryptographic keys
	provider := common.DefaultConfigProvider()

	//Build the signer
	signer := common.RequestSigner(provider, defaultGenericHeaders, optionalHeaders, optionalHeadersPredicate)

	//Signs the request
	signer.Sign(&request)

For more information on the signing algorithm refer to:  https://tools.ietf.org/html/draft-cavage-http-signatures-08

Pagination

When you call a list operation (for example list_instances()) will retrieve a page of results. In order to retrieve more data,
you have to continue to make calls to the list operation, passing in the value of the most recent response's OpcNextPage attribute
as a parameter to the next list operation call. The file: https://github.com/oracle/oci-go-sdk/blob/master/example/example_core_test.go
has an example of a paginated call

Polymorphic json request and responses

Some request return polymorphic responses


Logging and Debugging

The oci-go-sdk has a built-in logging mechanism used internally. The internal logging logic is used to record the raw http
requests, responses and potential errors when (un)marshalling request and responses.

To expose debugging logs set the environment variable "OCI_GO_SDK_DEBUG" to "1", or some other non emtpy string.





Forward and Backwards Compatibility

Some response fields are enum-typed. In the future, individual services may return values not covered by existing enums
for that field. To address this possibility, every enum-type response field is a model as a type that supports any string.
Thus if service returns a value that is not recognized by your version of the SDK, then the response field will be set to this value.

 */
package oci

//go:generate go run cmd/genver/main.go cmd/genver/version_template.go --output common/version.go
