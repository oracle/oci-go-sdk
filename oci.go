
/*
This is a preview of the official go sdk for Oracle Cloud Infrastructure

Installation

Refer to https://github.com/oracle/oci-go-sdk/blob/master/README.md#installing for installation instructions.

Configuration

Refer to https://github.com/oracle/oci-go-sdk/blob/master/README.md#configuring for configuration instructions.

Quickstart

The following example shows how to get started with the sdk

	import (
		"context"
		"fmt"

		"github.com/oracle/oci-go-sdk/common"
		"github.com/oracle/oci-go-sdk/identity"
	)

	func main() {
		c, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
		LogIfError(err)

		// The OCID of the tenancy containing the compartment.
		tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
		LogIfError(err)

		request := identity.ListAvailabilityDomainsRequest{
			CompartmentId: &tenancyID,
		}

		r, err := c.ListAvailabilityDomains(context.Background(), request)
		LogIfError(err)

		fmt.Printf("List of available domains: %v", r.Items)
		return
	}

More examples can be found in the oci-go-sdk repo: https://github.com/oracle/oci-go-sdk/tree/master/example

Optional fields in the SDK

Optional fields are represented with the `mandatory:"false"` tag on input structs. The sdk will omit all optional fields that are nil.
In the case of enum-type fields, the sdk will omit fields whose value is the empty string.

Signing custom requests

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

Polymorphic json request and responses

Some operations accept or return polymorphic json objects. The SDK models such objects as an interface. Further the SDK provides
structs that implement such interfaces. Thus for all operations that expect interfaces, pass the struct in the SDK that satisfies
such interface. For example:

	c, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		panic(err)
	}

	// The CreateIdentityProviderRequest. Takes a CreateIdentityProviderDetails interface as input
	rCreate := identity.CreateIdentityProviderRequest{}

	// The CreateSaml2IdentityProviderDetails implement such interface.
	details := identity.CreateSaml2IdentityProviderDetails{}
	details.CompartmentId = common.String(getTenancyID())
	details.Name = common.String("someName")
	//... more setup if needed
	// Use the above struct
	rCreate.CreateIdentityProviderDetails = details

	// Make the call
	rspCreate, createErr := c.CreateIdentityProvider(context.Background(), rCreate)

In the case of a response you can type assert the interface to the expected type. For example:

	rRead := identity.GetIdentityProviderRequest{}
	rRead.IdentityProviderId = common.String("aValidId")
	response, err := c.GetIdentityProvider(context.Background(), rRead)

	provider := response.IdentityProvider.(identity.Saml2IdentityProvider)


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
