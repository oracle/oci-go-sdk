// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package identitySample

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/example"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/identity"
)

// ListAvailabilityDomains Lists the Availability Domains in your tenancy.
// Specify the OCID of either the tenancy or another of your compartments as
// the value for the compartment ID (remember that the tenancy is simply the root compartment).
func ListAvailabilityDomains() {
	c, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	example.LogIfError(err)

	// The OCID of the tenancy containing the compartment.
	tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
	example.LogIfError(err)

	request := identity.ListAvailabilityDomainsRequest{
		CompartmentId: &tenancyID,
	}

	r, err := c.ListAvailabilityDomains(context.Background(), request)
	example.LogIfError(err)

	fmt.Printf("List of available domains: %v", r.Items)
	return
}
