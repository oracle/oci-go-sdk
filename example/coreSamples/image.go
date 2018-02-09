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

// CreateImageDetailsPolymorphic creates a boot disk image for the specified instance or imports an
// exported image from the Oracle Cloud Infrastructure Object Storage service.
func CreateImageDetailsPolymorphic(compartmentID *string, imageSourceURI *string) {
	request := core.CreateImageRequest{}
	request.CompartmentId = compartmentID

	// you can import an image based on the Object Storage URL 'core.ImageSourceViaObjectStorageUriDetails'
	// or based on the namespace, bucket name and object name 'core.ImageSourceViaObjectStorageTupleDetails'
	// following example shows how to import image from object storage uri, you can use another one, i.e.
	// request.ImageSourceDetails = core.ImageSourceViaObjectStorageTupleDetails
	// this shows how polymorphic json works in SDK
	sourceDetails := core.ImageSourceViaObjectStorageUriDetails{}
	sourceDetails.SourceUri = imageSourceURI

	// request.ImageSourceDetails = core.ImageSourceViaObjectStorageTupleDetails
	request.ImageSourceDetails = sourceDetails

	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	example.LogIfError(err)

	_, err = c.CreateImage(context.Background(), request)
	example.LogIfError(err)

	return
}
