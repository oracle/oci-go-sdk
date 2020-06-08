// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for creating a Service Gateway

package example

import (
	"context"
	"fmt"
	"os"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/example/helpers"
)

func ExampleCreateServiceGateway() {
	displayName := "OCI-GOSDK-CreateServiceGateway-Example" // displayName for created VCN and ServiceGateway
	compartmentID := os.Getenv("OCI_COMPARTMENT_ID")        // OCI_COMPARTMENT_ID env variable must be defined

	// initialize VirtualNetworkClient
	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)
	ctx := context.Background()

	// create VCN
	createVcnRequest := core.CreateVcnRequest{
		RequestMetadata: helpers.GetRequestMetadataWithDefaultRetryPolicy(),
	}
	createVcnRequest.CompartmentId = common.String(compartmentID)
	createVcnRequest.DisplayName = common.String(displayName)
	createVcnRequest.CidrBlock = common.String("10.0.0.0/16")
	createVcnResponse, err := client.CreateVcn(ctx, createVcnRequest)
	helpers.FatalIfError(err)

	// create ServiceGateway
	createServiceGatewayRequest := core.CreateServiceGatewayRequest{
		RequestMetadata: helpers.GetRequestMetadataWithDefaultRetryPolicy(),
	}
	createServiceGatewayRequest.CompartmentId = common.String(compartmentID)
	createServiceGatewayRequest.DisplayName = common.String(displayName)
	createServiceGatewayRequest.VcnId = createVcnResponse.Id
	createServiceGatewayRequest.Services = []core.ServiceIdRequestDetails{}
	_, err = client.CreateServiceGateway(ctx, createServiceGatewayRequest)
	helpers.FatalIfError(err)

	fmt.Println("ServiceGateway created")
	// Output:
	// ServiceGateway created
}
