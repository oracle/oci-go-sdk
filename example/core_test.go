// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
//
// Example code for Core Services API
//

package core_test

import (
	"context"
	"fmt"
	"log"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/example/helper"
)

// replace following variables with your instance info
const (
	availabilityDomain        = "[The Availability Domain of the instance. Example: Uocm:PHX-AD-1]"
	compartmentID             = "[The OCID of the compartment.]"
	instanceDisplayName       = "[A user-friendly name. Does not have to be unique, and it's changeable.]"
	objectStorageURIWtihImage = "[The Object Storage URL for the image which will be used to create an image.]"
)

// ExampleLaunchInstance does create an instance
// NOTE: launch instance will create a new instance and VCN. please make sure delete the instance
// after execute this sample code, otherwise, you will be charged for the running instance
func ExampleLaunchInstance() {
	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helper.LogIfError(err)

	// create the launch instance request
	request := core.LaunchInstanceRequest{}
	request.CompartmentId = common.String(compartmentID)
	request.DisplayName = common.String(instanceDisplayName)
	request.AvailabilityDomain = common.String(availabilityDomain)

	// create a virtual network
	vcn := createVcn()

	// create a subnet
	subnet := createSubnet(vcn.Id)
	request.SubnetId = subnet.Id

	// get a image
	image := listImages()[0]
	request.ImageId = image.Id

	// get all the shapes and filter the list by compatibility with the image
	shapes := listShapes(request.ImageId)
	request.Shape = shapes[0].Shape

	_, err = c.LaunchInstance(context.Background(), request)
	helper.LogIfError(err)

	return
}

// ExampleCreateImageDetails_Polymorphic creates a boot disk image for the specified instance or
// imports an exported image from the Oracle Cloud Infrastructure Object Storage service.
func ExampleCreateImageDetails_Polymorphic() {
	request := core.CreateImageRequest{}
	request.CompartmentId = common.String(compartmentID)

	// you can import an image based on the Object Storage URL 'core.ImageSourceViaObjectStorageUriDetails'
	// or based on the namespace, bucket name and object name 'core.ImageSourceViaObjectStorageTupleDetails'
	// following example shows how to import image from object storage uri, you can use another one:
	// request.ImageSourceDetails = core.ImageSourceViaObjectStorageTupleDetails
	sourceDetails := core.ImageSourceViaObjectStorageUriDetails{}
	sourceDetails.SourceUri = common.String(objectStorageURIWtihImage)

	request.ImageSourceDetails = sourceDetails

	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	LogIfError(err)

	_, err = c.CreateImage(context.Background(), request)
	LogIfError(err)

	return
}

// ExampleListShapes_Pagination demostrate how to use page parameter
func ExampleListShapes_Pagination() {
	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	LogIfError(err)

	request := core.ListShapesRequest{
		CompartmentId: common.String(compartmentID),
	}

	// to show how pagination works, reduce number of items to return in a paginated "List" call
	request.Limit = common.Int(2)

	listShapesFunc := func(request core.ListShapesRequest) (core.ListShapesResponse, error) {
		return c.ListShapes(context.Background(), request)
	}

	for r, err := listShapesFunc(request); ; r, err = listShapesFunc(request) {
		LogIfError(err)

		fmt.Printf("list shapes returns: %v", r.Items)

		if r.OpcNextPage != nil {
			// if there are more items in next page, fetch items from next page
			request.Page = r.OpcNextPage
		} else {
			// no more result, break the loop
			break
		}
	}

	return
}

// CreateVcn creates a new Virtual Cloud Network (VCN)
func createVcn() core.Vcn {
	request := core.CreateVcnRequest{}
	request.CidrBlock = common.String("10.0.0.0/16") // The CIDR IP address block of the VCN.
	request.CompartmentId = common.String(compartmentID)
	request.DisplayName = common.String("OCI GOSDK VCN DisplayName")

	c, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	LogIfError(err)

	r, err := c.CreateVcn(context.Background(), request)
	LogIfError(err)
	return r.Vcn
}

// CreateSubnet creates a new subnet in the specified VCN
func createSubnet(vcnID *string) core.Subnet {
	request := core.CreateSubnetRequest{}
	request.CompartmentId = common.String(compartmentID)
	request.CidrBlock = common.String("10.0.0.0/16") // The CIDR IP address range of the subnet.
	request.VcnId = vcnID
	request.DisplayName = common.String("OCI GOSDK Subnet DisplayName")
	request.AvailabilityDomain = common.String(availabilityDomain)

	c, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	LogIfError(err)

	r, err := c.CreateSubnet(context.Background(), request)
	LogIfError(err)
	return r.Subnet
}

// ListImages lists the available images in the specified compartment.
func listImages() []core.Image {
	request := core.ListImagesRequest{
		CompartmentId: common.String(compartmentID),
	}

	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	LogIfError(err)

	r, err := c.ListImages(context.Background(), request)
	LogIfError(err)

	return r.Items
}

// ListShapes Lists the shapes that can be used to launch an instance within the specified compartment.
func listShapes(imageID *string) []core.Shape {
	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	LogIfError(err)

	request := core.ListShapesRequest{
		CompartmentId: common.String(compartmentID),
		ImageId:       imageID,
	}

	r, err := c.ListShapes(context.Background(), request)
	LogIfError(err)

	if r.Items == nil || len(r.Items) == 0 {
		log.Fatalln("Invalid response from ListShapes")
	}

	return r.Items
}
