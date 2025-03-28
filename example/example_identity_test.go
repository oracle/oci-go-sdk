// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Identity and Access Management Service API

package example

import (
	"context"
	"fmt"
	"log"

	"net/http"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

// ExampleListAvailabilityDomains Lists the Availability Domains in your tenancy.
// Specify the OCID of either the tenancy or another of your compartments as
// the value for the compartment ID (remember that the tenancy is simply the root compartment).
func ExampleListAvailabilityDomains() {
	c, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// The OCID of the tenancy containing the compartment.
	tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
	helpers.FatalIfError(err)

	request := identity.ListAvailabilityDomainsRequest{
		CompartmentId: &tenancyID,
	}

	r, err := c.ListAvailabilityDomains(context.Background(), request)
	helpers.FatalIfError(err)

	log.Printf("list of available domains: %v", r.Items)
	fmt.Println("list available domains completed")

	// Output:
	// list available domains completed
}

// ExampleListGroupsWithCustomSignedHeader Lists groups by passing a custom signed header in the request
func ExampleListGroupsWithCustomSignedHeader() {
	provider := common.DefaultConfigProvider()
	c, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	//Bear in mind that services expect well known headers to be signed. Signing arbitrary headers
	//might lead to authentication errors
	customHeader := "opc-my-token"
	allHeaders := append(common.DefaultGenericHeaders(), customHeader)
	c.Signer = common.RequestSigner(provider, allHeaders, common.DefaultBodyHeaders())
	c.Interceptor = func(request *http.Request) error {
		request.Header.Add(customHeader, "customvalue")
		return nil
	}

	// The OCID of the tenancy containing the compartment.
	tenancyID, _ := provider.TenancyOCID()
	request := identity.ListGroupsRequest{
		CompartmentId: common.String(tenancyID),
	}
	r, err := c.ListGroups(context.Background(), request)
	helpers.FatalIfError(err)

	log.Printf("list groups completed: %v", r.Items)
	fmt.Println("list groups completed")

	// Output:
	// list groups completed
}
