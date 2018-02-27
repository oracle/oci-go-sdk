// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
//
// Example code for Core Services API
//

package example

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/example/helpers"
)

// replace following variables with your instance info
const (
	// this is used by ExampleCreateImageDetails_Polymorphic
	objectStorageURIWtihImage = "[The Object Storage URL for the image which will be used to create an image.]"
	vcnDisplayName            = "OCI-GOSDK-Sample-VCN"
	subnetDisplayName1        = "OCI-GOSDK-Sample-Subnet1"
	subnetDisplayName2        = "OCI-GOSDK-Sample-Subnet2"
)

// ExampleLaunchInstance does create an instance
// NOTE: launch instance will create a new instance and VCN. please make sure delete the instance
// after execute this sample code, otherwise, you will be charged for the running instance
func ExampleLaunchInstance() {
	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.LogIfError(err)

	// create the launch instance request
	request := core.LaunchInstanceRequest{}
	request.CompartmentId = helpers.CompartmentID()
	request.DisplayName = common.String("OCI-Sample-Instance")
	request.AvailabilityDomain = helpers.AvailabilityDomain()

	// create a subnet or get the one already created
	subnet := CreateOrGetSubnet()
	fmt.Println("Subnet created")
	request.SubnetId = subnet.Id
	helpers.AddResource(*subnet.Id, "ExampleLaunchInstance_Subnet")
	helpers.AddResource(*subnet.VcnId, "ExampleLaunchInstance_VCN")

	// get a image
	image := listImages()[0]
	fmt.Println("list images")
	request.ImageId = image.Id

	// get all the shapes and filter the list by compatibility with the image
	shapes := listShapes(request.ImageId)
	fmt.Println("list shapes")
	request.Shape = shapes[0].Shape

	createResp, err := c.LaunchInstance(context.Background(), request)
	helpers.LogIfError(err)
	fmt.Println("instance created")
	helpers.AddResource(*createResp.Id, "ExampleLaunchInstance_Instance")

	// get new created instance
	getInstance := func() (interface{}, error) {
		c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
		if clerr != nil {
			return nil, clerr
		}

		request := core.GetInstanceRequest{
			InstanceId: createResp.Instance.Id,
		}

		readResp, err := c.GetInstance(context.Background(), request)

		if err != nil {
			return nil, clerr
		}

		return readResp, err
	}

	// wait for instance lifecyle become running
	helpers.LogIfError(
		helpers.RetryUntilTrueOrError(
			getInstance,
			helpers.CheckLifecycleState(string(core.InstanceLifecycleStateRunning)),
			time.Tick(10*time.Second),
			time.After((5 * time.Minute))))

	// Output:
	// VCN created
	// Subnet created
	// list images
	// list shapes
	// instance created
}

// ExampleCreateImageDetails_Polymorphic creates a boot disk image for the specified instance or
// imports an exported image from the Oracle Cloud Infrastructure Object Storage service.
func ExampleCreateImageDetails_Polymorphic() {
	request := core.CreateImageRequest{}
	request.CompartmentId = helpers.CompartmentID()

	// you can import an image based on the Object Storage URL 'core.ImageSourceViaObjectStorageUriDetails'
	// or based on the namespace, bucket name and object name 'core.ImageSourceViaObjectStorageTupleDetails'
	// following example shows how to import image from object storage uri, you can use another one:
	// request.ImageSourceDetails = core.ImageSourceViaObjectStorageTupleDetails
	sourceDetails := core.ImageSourceViaObjectStorageUriDetails{}
	sourceDetails.SourceUri = common.String(objectStorageURIWtihImage)

	request.ImageSourceDetails = sourceDetails

	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.LogIfError(err)

	_, err = c.CreateImage(context.Background(), request)
	helpers.LogIfError(err)
	fmt.Println("image created")

	// Output:
	// image created
}

// ExampleListShapes_Pagination demostrate how to use page parameter
func ExampleListShapes_Pagination() {
	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.LogIfError(err)

	request := core.ListShapesRequest{
		CompartmentId: helpers.CompartmentID(),
	}

	// to show how pagination works, reduce number of items to return in a paginated "List" call
	request.Limit = common.Int(2)

	listShapesFunc := func(request core.ListShapesRequest) (core.ListShapesResponse, error) {
		return c.ListShapes(context.Background(), request)
	}

	for r, err := listShapesFunc(request); ; r, err = listShapesFunc(request) {
		helpers.LogIfError(err)

		fmt.Printf("list shapes returns: %v", r.Items)

		if r.OpcNextPage != nil {
			// if there are more items in next page, fetch items from next page
			request.Page = r.OpcNextPage
		} else {
			// no more result, break the loop
			break
		}
	}

	fmt.Println("list shapes completed")

	// Output:
	// list shapes completed
}

// CreateOrGetVcn either creates a new Virtual Cloud Network (VCN) or get the one already exist
func CreateOrGetVcn() core.Vcn {
	vcnItems := listVcns()

	for _, element := range vcnItems {
		if *element.DisplayName == vcnDisplayName {
			// VCN already created, return it
			return element
		}
	}

	// create a new VCN
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.LogIfError(clerr)

	request := core.CreateVcnRequest{}
	request.CidrBlock = common.String("10.0.0.0/16")
	request.CompartmentId = helpers.CompartmentID()
	request.DisplayName = common.String(vcnDisplayName)
	request.DnsLabel = common.String("vcndns")

	r, err := c.CreateVcn(context.Background(), request)
	helpers.LogIfError(err)
	return r.Vcn
}

// CreateSubnet creates a new subnet or get the one already exist
func CreateOrGetSubnet() core.Subnet {
	return CreateOrGetSubnetWithDetails(
		common.String(subnetDisplayName1),
		common.String("10.0.0.0/24"),
		common.String("subnetdns1"),
		helpers.AvailabilityDomain())
}

// CreateOrGetSubnetWithDetails either creates a new Virtual Cloud Network (VCN) or get the one already exist
// with detail info
func CreateOrGetSubnetWithDetails(displayName *string, cidrBlock *string, dnsLabel *string, availableDomain *string) core.Subnet {
	subnets := listSubnets()

	if displayName == nil {
		displayName = common.String(subnetDisplayName1)
	}

	// check if the subnet has already been created
	for _, element := range subnets {
		if *element.DisplayName == *displayName {
			// find the subnet, return it
			return element
		}
	}

	// create a new subnet
	vcn := CreateOrGetVcn()
	request := core.CreateSubnetRequest{}
	request.AvailabilityDomain = availableDomain
	request.CompartmentId = helpers.CompartmentID()
	request.CidrBlock = cidrBlock
	request.VcnId = vcn.Id
	request.DisplayName = displayName
	request.DnsLabel = dnsLabel

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.LogIfError(clerr)

	r, err := c.CreateSubnet(context.Background(), request)
	helpers.LogIfError(err)

	getSubnet := func() (interface{}, error) {
		getReq := core.GetSubnetRequest{
			SubnetId: r.Id,
		}

		getResp, err := c.GetSubnet(context.Background(), getReq)

		if err != nil {
			return nil, err
		}

		return getResp, nil
	}

	// wait for lifecyle become running
	helpers.LogIfError(
		helpers.RetryUntilTrueOrError(
			getSubnet,
			helpers.CheckLifecycleState(string(core.SubnetLifecycleStateAvailable)),
			time.Tick(10*time.Second),
			time.After((5 * time.Minute))))

	// update the security rules
	getReq := core.GetSecurityListRequest{
		SecurityListId: common.String(r.SecurityListIds[0]),
	}

	getResp, err := c.GetSecurityList(context.Background(), getReq)
	helpers.LogIfError(err)

	// this security rule allows remote control the instance
	portRange := core.PortRange{
		Max: common.Int(1521),
		Min: common.Int(1521),
	}

	newRules := append(getResp.IngressSecurityRules, core.IngressSecurityRule{
		Protocol: common.String("6"), // TCP
		Source:   common.String("0.0.0.0/0"),
		TcpOptions: &core.TcpOptions{
			DestinationPortRange: &portRange,
		},
	})

	updateReq := core.UpdateSecurityListRequest{
		SecurityListId: common.String(r.SecurityListIds[0]),
	}

	updateReq.IngressSecurityRules = newRules

	_, err = c.UpdateSecurityList(context.Background(), updateReq)
	helpers.LogIfError(err)

	return r.Subnet
}

func listVcns() []core.Vcn {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.LogIfError(clerr)
	request := core.ListVcnsRequest{
		CompartmentId: helpers.CompartmentID(),
	}

	r, err := c.ListVcns(context.Background(), request)
	helpers.LogIfError(err)
	return r.Items
}

func listSubnets() []core.Subnet {
	vcn := CreateOrGetVcn()

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.LogIfError(clerr)
	request := core.ListSubnetsRequest{
		CompartmentId: helpers.CompartmentID(),
		VcnId:         vcn.Id,
	}

	r, err := c.ListSubnets(context.Background(), request)
	helpers.LogIfError(err)
	return r.Items
}

// ListImages lists the available images in the specified compartment.
func listImages() []core.Image {
	request := core.ListImagesRequest{
		CompartmentId: helpers.CompartmentID(),
	}

	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.LogIfError(err)

	r, err := c.ListImages(context.Background(), request)
	helpers.LogIfError(err)

	return r.Items
}

// ListShapes Lists the shapes that can be used to launch an instance within the specified compartment.
func listShapes(imageID *string) []core.Shape {
	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.LogIfError(err)

	request := core.ListShapesRequest{
		CompartmentId: helpers.CompartmentID(),
		ImageId:       imageID,
	}

	r, err := c.ListShapes(context.Background(), request)
	helpers.LogIfError(err)

	if r.Items == nil || len(r.Items) == 0 {
		log.Fatalln("Invalid response from ListShapes")
	}

	return r.Items
}

func terminateInstance(id *string) {

}

func deleteVcn(id *string) {}

func deleteSubnet(id *string) {}
