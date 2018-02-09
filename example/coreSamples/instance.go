// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package coreSamples

import (
	"context"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/example"
)

// LaunchInstance does create an instance
func LaunchInstance(availabilityDomain *string, compartmentID *string, displayName *string) {
	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	example.LogIfError(err)

	// create the launch instance request
	request := core.LaunchInstanceRequest{}
	request.CompartmentId = compartmentID
	request.DisplayName = displayName
	request.AvailabilityDomain = availabilityDomain

	// create a virtual network
	vnc := CreateVcn(compartmentID)

	// create a subnet
	subnet := CreateSubnet(vnc, availabilityDomain, compartmentID)
	request.SubnetId = subnet.Id

	// get the image
	image := ListImages(compartmentID)[0]
	request.ImageId = image.Id

	// get all the shapes and filter the list by compatibility with the image
	shapes := ListShapes(compartmentID, request.ImageId)
	request.Shape = shapes[0].Shape

	_, err = c.LaunchInstance(context.Background(), request)
	example.LogIfError(err)

	return
}
