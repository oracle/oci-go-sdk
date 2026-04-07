// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Oracle Cloud Infrastructure Go SDK

package example

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

// This example file shows how to use Workload Identity Federation in Oracle Cloud.
// A token is obtained from a registered token issuer, which is exchanged for a valid
// User Principal Session Token (UPST). This UPST can then be used to authenticate
// to OCI services. A valid Configuration Provider is used to simplify the process
// and share credentials among service clients.

// Examples for token exchange grant type require Identity Propagation Trust
//  to be configured on desired OCI Identity Domain. Documentation for setting up a
//  Identity Propagation Trust can be found at the following location:
//  https://docs.oracle.com/en-us/iaas/Content/Identity/api-getstarted/json_web_token_exchange.htm
//
// Examples use the following environment variables:
//
//  Domain Variables (Required):
//    OCI_DOMAIN_URL - The URL for the Identity Domain issuing
//        User Principal Session Tokens (UPST) (ex. https://idcs-xxxx.identity.oraclecloud.com)
//    OCI_CLIENT_ID - Client ID of the OAuth client application
//    OCI_CLIENT_SECRET - Client secret of the OAuth client application
//    OCI_REGION - A valid OCI region to query
//
//  OCI Identity Client variables (Required):
//    OCI_ROOT_COMPARTMENT_ID - Root OCID of the OCI tenancy
//
//  Token Exchange from Configuration Provider From Token Variables:
//    YOUR_TOKEN - A valid token issued by a provider registered in Identity Propagation Trust Configuration
//
//  Token Exchange from Issuer Variables:
//    ISSUER_ENDPOINT - Token issuer endpoint for authorization server
//    ISSUER_ID - Client ID for client credentials flow
//    ISSUER_SECRET - Client secret for client credentials flow

// ExampleTokenExchangeConfigurationProviderFromToken demonstrates exchanging a token for
// a UPST and calling the Identity client
func ExampleTokenExchangeConfigurationProviderFromToken() {
	// YOUR_TOKEN MUST be a valid token issued by a registered Identity Propagation Trust
	token := os.Getenv("YOUR_TOKEN")
	builder := auth.TokenExchangeBuilder{
		DomainUrl:          os.Getenv("OCI_DOMAIN_URL"),
		ClientId:           os.Getenv("OCI_CLIENT_ID"),
		ClientSecret:       os.Getenv("OCI_CLIENT_SECRET"),
		Region:             os.Getenv("OCI_REGION"),
		RequestedTokenType: "urn:oci:token-type:oci-upst",
		SubjectTokenType:   "jwt",
		PublicKey:          os.Getenv("PUBLIC_KEY"),
	}
	provider, err := auth.TokenExchangeConfigurationProviderFromToken(token, builder)
	helpers.FatalIfError(err)

	tenancyID := os.Getenv("OCI_ROOT_COMPARTMENT_ID")
	request := identity.ListAvailabilityDomainsRequest{
		CompartmentId: &tenancyID,
	}

	client, err := identity.NewIdentityClientWithConfigurationProvider(provider)
	helpers.FatalIfError(err)

	r, err := client.ListAvailabilityDomains(context.Background(), request)
	helpers.FatalIfError(err)

	log.Printf("List of availability domains: %v\n", r.Items)
	fmt.Println("Done")

	// Output:
	// Done
}

// ExampleTokenExchangeConfigurationProviderFromIssuer demonstrates using a function to
// get and refresh UPSTs by calling the token issuer
func ExampleTokenExchangeConfigurationProviderFromIssuer() {

	// ExampleIssuer defined below that complies with the TokenIssuer interface. This
	// is what will issue the tokens we exchange for OCI UPSTs.
	issuer := &ExampleIssuer{
		IssuerEndpoint: os.Getenv("ISSUER_ENDPOINT"),
		ClientId:       os.Getenv("ISSUER_ID"),
		ClientSecret:   os.Getenv("ISSUER_SECRET"),
		Method:         http.MethodPost,
		HttpClient:     &http.Client{},
	}

	builder := auth.TokenExchangeBuilder{
		DomainUrl:          os.Getenv("OCI_DOMAIN_URL"),
		ClientId:           os.Getenv("OCI_CLIENT_ID"),
		ClientSecret:       os.Getenv("OCI_CLIENT_SECRET"),
		Region:             os.Getenv("OCI_REGION"),
		RequestedTokenType: "urn:oci:token-type:oci-upst",
		SubjectTokenType:   "jwt",
		PublicKey:          os.Getenv("PUBLIC_KEY"),
	}

	// The provider consumes the TokenIssuer and exchanges the issued token(s) for UPST(s)
	provider, err := auth.TokenExchangeConfigurationProviderFromIssuer(issuer, builder)
	helpers.FatalIfError(err)

	tenancyID := os.Getenv("OCI_ROOT_COMPARTMENT_ID")
	request := identity.ListAvailabilityDomainsRequest{
		CompartmentId: &tenancyID,
	}

	client, err := identity.NewIdentityClientWithConfigurationProvider(provider)
	helpers.FatalIfError(err)

	r, err := client.ListAvailabilityDomains(context.Background(), request)
	helpers.FatalIfError(err)

	log.Printf("List of availability domains: %v\n", r.Items)
	fmt.Println("Done")

	// Output:
	// Done
}

// ExampleIssuer satisfies the TokenIssuer interface. Defining one's own struct will
// allow for storage of as much or as little data as needed to get tokens from the
// authorization server. This allows for flexibility in defining how to retrieve
// tokens and any required data or logic.
type ExampleIssuer struct {
	IssuerEndpoint string       // Token endpoint of example authorization server
	ClientId       string       // Required for client credentials flow
	ClientSecret   string       // Required for client credentials flow
	Method         string       // Expected HTTP method for request
	HttpClient     *http.Client // Client to use for request
}

// GetToken allows the ExampleIssuer to satisfy the TokenIssuer interface. This receiver
// method will be called to get and refresh tokens. Anything can be substituted in this
// method as long as it returns a valid token in string form and an error.
func (e *ExampleIssuer) GetToken() (string, error) {

	// Required values for example request as defined by authorization server
	data := url.Values{
		"grant_type": {"client_credentials"}, // Client credentials flow
		"scope":      {"token_exchange"},     // Custom scope
	}

	request, err := http.NewRequest(e.Method, e.IssuerEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	// Basic authentication requires base64 encoding
	authHeader := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(
		e.ClientId+":"+e.ClientSecret)))

	// Headers required by example authorization server
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", authHeader)

	response, err := e.HttpClient.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-200 status code returned: %v", response.StatusCode)
	}

	var body map[string]interface{}
	if err = json.NewDecoder(response.Body).Decode(&body); err != nil {
		return "", fmt.Errorf("unable to unmarshal response: %s", err)
	}

	token, ok := body["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("unable to retrieve access token")
	}

	return token, nil
}

func ExampleTokenExchangeConfigurationProviderFromToken_toRPST() {
	// YOUR_TOKEN MUST be a valid token issued by a registered Identity Propagation Trust
	token := os.Getenv("YOUR_TOKEN")

	builder := auth.TokenExchangeBuilder{
		DomainUrl:          os.Getenv("OCI_DOMAIN_URL"),
		ClientId:           os.Getenv("OCI_CLIENT_ID"),
		ClientSecret:       os.Getenv("OCI_CLIENT_SECRET"),
		Region:             os.Getenv("OCI_REGION"),
		RequestedTokenType: "urn:oci:token-type:oci-rpst",
		ResType:            os.Getenv("RES_TYPE"),
		SubjectTokenType:   "jwt",
		PublicKey:          os.Getenv("PUBLIC_KEY"),
		RpstExp:            os.Getenv("RPST_EXP"),
	}
	provider, err := auth.TokenExchangeConfigurationProviderFromToken(token, builder)
	log.Printf("Provider details: %v", provider)
	helpers.FatalIfError(err)
	fmt.Println("Done")

	// Output:
	// Done
}

func ExampleTokenExchangeConfigurationProviderFromToken_toRPST1() {
	// YOUR_TOKEN MUST be a valid token issued by a registered Identity Propagation Trust
	token := os.Getenv("YOUR_TOKEN")
	instancePrincipalProvider, err := auth.InstancePrincipalConfigurationProvider()

	builder := auth.TokenExchangeBuilder{
		DomainUrl:                 os.Getenv("OCI_DOMAIN_URL"),
		Region:                    os.Getenv("OCI_REGION"),
		RequestedTokenType:        "urn:oci:token-type:oci-rpst",
		ResType:                   os.Getenv("RES_TYPE"),
		SubjectTokenType:          "jwt",
		PublicKey:                 os.Getenv("PUBLIC_KEY"),
		InstancePrincipalProvider: instancePrincipalProvider,
	}
	_, err = auth.TokenExchangeConfigurationProviderFromToken(token, builder)
	helpers.FatalIfError(err)
	fmt.Println("Done")

	// Output:
	// Done
}
