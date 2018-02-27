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

const (
	vcnDisplayName     = "OCI-GOSDK-Sample-VCN"
	subnetDisplayName1 = "OCI-GOSDK-Sample-Subnet1"
	subnetDisplayName2 = "OCI-GOSDK-Sample-Subnet2"

	// replace following variables with your instance info
	// this is used by ExampleCreateImageDetails_Polymorphic
	objectStorageURIWtihImage = "[The Object Storage URL for the image which will be used to create an image.]"
)

// ExampleLaunchInstance does create an instance
// NOTE: launch instance will create a new instance and VCN. please make sure delete the instance
// after execute this sample code, otherwise, you will be charged for the running instance
func ExampleLaunchInstance() {
	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.LogIfError(err)
	ctx := context.Background()

	// create the launch instance request
	request := core.LaunchInstanceRequest{}
	request.CompartmentId = helpers.CompartmentID()
	request.DisplayName = common.String("OCI-Sample-Instance")
	request.AvailabilityDomain = helpers.AvailabilityDomain()

	// create a subnet or get the one already created
	subnet := CreateOrGetSubnet()
	fmt.Println("subnet created")
	request.SubnetId = subnet.Id

	// get a image
	image := listImages(ctx, c)[0]
	fmt.Println("list images")
	request.ImageId = image.Id

	// get all the shapes and filter the list by compatibility with the image
	shapes := listShapes(ctx, c, request.ImageId)
	fmt.Println("list shapes")
	request.Shape = shapes[0].Shape

	createResp, err := c.LaunchInstance(ctx, request)
	helpers.LogIfError(err)
	fmt.Println("instance created")

	// get new created instance
	getInstance := func() (interface{}, error) {
		request := core.GetInstanceRequest{
			InstanceId: createResp.Instance.Id,
		}

		readResp, err := c.GetInstance(ctx, request)

		if err != nil {
			return nil, err
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

	defer func() {
		terminateInstance(ctx, c, createResp.Id)

		client, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
		helpers.LogIfError(clerr)

		vcnID := subnet.VcnId
		deleteSubnet(ctx, client, subnet.Id)
		deleteVcn(ctx, client, vcnID)
	}()

	// Output:
	// subnet created
	// list images
	// list shapes
	// instance created
	// terminating instance
	// instance terminated
	// deleteing subnet
	// subnet deleted
	// deleteing VCN
	// VCN deleted
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
}

// CreateOrGetVcn either creates a new Virtual Cloud Network (VCN) or get the one already exist
func CreateOrGetVcn() core.Vcn {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.LogIfError(clerr)
	ctx := context.Background()

	vcnItems := listVcns(ctx, c)

	for _, element := range vcnItems {
		if *element.DisplayName == vcnDisplayName {
			// VCN already created, return it
			return element
		}
	}

	// create a new VCN
	request := core.CreateVcnRequest{}
	request.CidrBlock = common.String("10.0.0.0/16")
	request.CompartmentId = helpers.CompartmentID()
	request.DisplayName = common.String(vcnDisplayName)
	request.DnsLabel = common.String("vcndns")

	r, err := c.CreateVcn(ctx, request)
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
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.LogIfError(clerr)
	ctx := context.Background()

	subnets := listSubnets(ctx, c)

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
	request := core.CreateSubnetRequest{}
	request.AvailabilityDomain = availableDomain
	request.CompartmentId = helpers.CompartmentID()
	request.CidrBlock = cidrBlock
	request.DisplayName = displayName
	request.DnsLabel = dnsLabel

	vcn := CreateOrGetVcn()
	request.VcnId = vcn.Id

	r, err := c.CreateSubnet(ctx, request)
	helpers.LogIfError(err)

	getSubnet := func() (interface{}, error) {
		getReq := core.GetSubnetRequest{
			SubnetId: r.Id,
		}

		getResp, err := c.GetSubnet(ctx, getReq)

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

	getResp, err := c.GetSecurityList(ctx, getReq)
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

	_, err = c.UpdateSecurityList(ctx, updateReq)
	helpers.LogIfError(err)

	return r.Subnet
}

func listVcns(ctx context.Context, c core.VirtualNetworkClient) []core.Vcn {
	request := core.ListVcnsRequest{
		CompartmentId: helpers.CompartmentID(),
	}

	r, err := c.ListVcns(ctx, request)
	helpers.LogIfError(err)
	return r.Items
}

func listSubnets(ctx context.Context, c core.VirtualNetworkClient) []core.Subnet {
	vcn := CreateOrGetVcn()

	request := core.ListSubnetsRequest{
		CompartmentId: helpers.CompartmentID(),
		VcnId:         vcn.Id,
	}

	r, err := c.ListSubnets(ctx, request)
	helpers.LogIfError(err)
	return r.Items
}

// ListImages lists the available images in the specified compartment.
func listImages(ctx context.Context, c core.ComputeClient) []core.Image {
	request := core.ListImagesRequest{
		CompartmentId: helpers.CompartmentID(),
	}

	r, err := c.ListImages(ctx, request)
	helpers.LogIfError(err)

	return r.Items
}

// ListShapes Lists the shapes that can be used to launch an instance within the specified compartment.
func listShapes(ctx context.Context, c core.ComputeClient, imageID *string) []core.Shape {
	request := core.ListShapesRequest{
		CompartmentId: helpers.CompartmentID(),
		ImageId:       imageID,
	}

	r, err := c.ListShapes(ctx, request)
	helpers.LogIfError(err)

	if r.Items == nil || len(r.Items) == 0 {
		log.Fatalln("Invalid response from ListShapes")
	}

	return r.Items
}

func terminateInstance(ctx context.Context, c core.ComputeClient, id *string) {
	request := core.TerminateInstanceRequest{
		InstanceId: id,
	}

	_, err := c.TerminateInstance(ctx, request)
	helpers.LogIfError(err)

	fmt.Println("terminating instance")

	// get new created instance
	getInstance := func() (interface{}, error) {
		request := core.GetInstanceRequest{
			InstanceId: id,
		}

		readResp, err := c.GetInstance(ctx, request)

		if err != nil {
			if readResp.RawResponse.StatusCode == 404 {
				// cannot find resources which means it's been deleted
				return core.Instance{LifecycleState: core.InstanceLifecycleStateTerminated}, nil
			}
			return nil, err
		}

		return readResp, err
	}

	// wait for instance lifecyle become terminated
	helpers.LogIfError(
		helpers.RetryUntilTrueOrError(
			getInstance,
			helpers.CheckLifecycleState(string(core.InstanceLifecycleStateTerminated)),
			time.Tick(10*time.Second),
			time.After((5 * time.Minute))))

	fmt.Println("instance terminated")

}

func deleteVcn(ctx context.Context, c core.VirtualNetworkClient, id *string) {
	request := core.DeleteVcnRequest{
		VcnId: id,
	}

	fmt.Println("deleteing VCN")
	_, err := c.DeleteVcn(ctx, request)
	helpers.LogIfError(err)

	getVcn := func() (interface{}, error) {
		getReq := core.GetVcnRequest{
			VcnId: id,
		}

		getResp, err := c.GetVcn(ctx, getReq)

		if err != nil {
			if getResp.RawResponse.StatusCode == 404 {
				// resource cannot found which means it's been deleted in this case
				return core.Vcn{LifecycleState: core.VcnLifecycleStateTerminated}, nil
			}

			return nil, err
		}

		return getResp, nil
	}

	// wait for lifecyle become terminated
	helpers.LogIfError(
		helpers.RetryUntilTrueOrError(
			getVcn,
			helpers.CheckLifecycleState(string(core.VcnLifecycleStateTerminated)),
			time.Tick(10*time.Second),
			time.After((5 * time.Minute))))

	fmt.Println("VCN deleted")
}

func deleteSubnet(ctx context.Context, c core.VirtualNetworkClient, id *string) {
	request := core.DeleteSubnetRequest{
		SubnetId: id,
	}

	_, err := c.DeleteSubnet(context.Background(), request)
	helpers.LogIfError(err)

	fmt.Println("deleteing subnet")

	getSubnet := func() (interface{}, error) {
		getReq := core.GetSubnetRequest{
			SubnetId: id,
		}

		getResp, err := c.GetSubnet(ctx, getReq)

		if err != nil {
			if getResp.RawResponse.StatusCode == 404 {
				// resource cannot found which means it's been deleted in this case
				return core.Subnet{LifecycleState: core.SubnetLifecycleStateTerminated}, nil
			}

			return nil, err
		}

		return getResp, nil
	}

	// wait for lifecyle become terminated
	helpers.LogIfError(
		helpers.RetryUntilTrueOrError(
			getSubnet,
			helpers.CheckLifecycleState(string(core.SubnetLifecycleStateTerminated)),
			time.Tick(10*time.Second),
			time.After((5 * time.Minute))))

	fmt.Println("subnet deleted")
}
