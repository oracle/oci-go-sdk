// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Compartments Service API
// This script provides an example on how to move a compartment to a different compartment
// This script will
//
//   * create cp_source_GOSDK under tenancy
//   * create cp_target_GOSDK under tenancy
//   * move cp_source_GOSDK under cp_target_GOSDK

package example

import (
	"context"
	"fmt"
	"log"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/oracle/oci-go-sdk/identity"
)

// ExampleMoveCompartment Moves an active compartment under a different parent
func ExampleMoveCompartment() {
	c, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// The OCID of the tenancy containing the compartment.
	tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
	helpers.FatalIfError(err)

	ctx := context.Background()
	cpSource := createCompartment(ctx, c, common.String(tenancyID), common.String("cp_source_GOSDK"))
	cpTarget := createCompartment(ctx, c, common.String(tenancyID), common.String("cp_target_GOSDK"))

	moveDetail := identity.MoveCompartmentDetails{
		TargetCompartmentId: cpTarget,
	}

	moveRequest := identity.MoveCompartmentRequest{
		CompartmentId:          cpSource,
		MoveCompartmentDetails: moveDetail,
		RequestMetadata:        helpers.GetRequestMetadataWithDefaultRetryPolicy(),
	}

	resp, err := c.MoveCompartment(ctx, moveRequest)
	helpers.FatalIfError(err)
	log.Printf("move compartment with workrequest id: %s", *resp.OpcWorkRequestId)
	fmt.Println("move compartment request is accepted")

	// get cpSource new parent
	cpSourceNewParent := getCompartment(ctx, c, cpSource).CompartmentId
	cpSourceNewParentName := getCompartment(ctx, c, cpSourceNewParent).Name

	log.Printf("cp_source_GOSDK new parent is: %v", *cpSourceNewParentName)
	fmt.Println("move compartment completed")

	// Output:
	// move compartment request is accepted
	// move compartment completed
}

func createCompartment(ctx context.Context, client identity.IdentityClient, tenantId *string, compartmentName *string) *string {
	detail := identity.CreateCompartmentDetails{
		CompartmentId: tenantId,
		Name:          compartmentName,
		Description:   compartmentName,
	}
	request := identity.CreateCompartmentRequest{
		CreateCompartmentDetails: detail,
		RequestMetadata:          helpers.GetRequestMetadataWithDefaultRetryPolicy(),
	}

	resp, err := client.CreateCompartment(ctx, request)
	helpers.FatalIfError(err)

	return resp.Id
}

func getCompartment(ctx context.Context, client identity.IdentityClient, compartmentId *string) identity.Compartment {
	request := identity.GetCompartmentRequest{
		CompartmentId:   compartmentId,
		RequestMetadata: helpers.GetRequestMetadataWithDefaultRetryPolicy(),
	}

	resp, err := client.GetCompartment(ctx, request)
	helpers.FatalIfError(err)

	return resp.Compartment
}
