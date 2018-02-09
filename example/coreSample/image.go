// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package coreSample

import (
	"context"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/example"
)

// ListImages lists the available images in the specified compartment.
func ListImages(compartmentID *string) []core.Image {
	request := core.ListImagesRequest{
		CompartmentId: compartmentID,
	}

	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	example.LogIfError(err)

	r, err := c.ListImages(context.Background(), request)
	example.LogIfError(err)

	return r.Items
}
