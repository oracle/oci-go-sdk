package core

import (
	"context"
	"log"

	"github.com/oracle/oci-go-sdk/example/helper"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/example/identity"
)

// LaunchInstance does create an instance
func LaunchInstance() core.Instance {
	// create a compartment, otherwise return the id of existing compartment
	compartmentID, err := identity.CreateCompartment()
	helper.LogIfError(err)

	request := core.LaunchInstanceRequest{}
	request.CompartmentId = compartmentID
	request.DisplayName = common.String("OCI_GOSDK_Sample_Instance")

	// get the availability domain
	availabilityDomains := identity.ListAvailabilityDomains()
	request.AvailabilityDomain = availabilityDomains[0].Name

	// create a subnet
	subnet := CreateSubnet()
	request.SubnetId = subnet.Id

	// get the image
	image := ListImages()[0]
	request.ImageId = image.Id

	// get all the shapes and filter the list by compatibility with the image
	shapes := listShapes(compartmentID, request.ImageId)
	request.Shape = shapes[0].Shape

	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helper.LogIfError(err)

	r, err := c.LaunchInstance(context.Background(), request)
	helper.LogIfError(err)

	return r.Instance
}

// ListShapes Lists the shapes that can be used to launch an instance within the specified compartment.
func listShapes(compartmentID *string, imageID *string) []core.Shape {
	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helper.LogIfError(err)

	request := core.ListShapesRequest{
		CompartmentId: compartmentID,
		ImageId:       imageID,
	}

	r, err := c.ListShapes(context.Background(), request)
	helper.LogIfError(err)

	if r.Items == nil || len(r.Items) == 0 {
		log.Fatalln("Invalid response from ListShapes")
	}

	return r.Items
}
