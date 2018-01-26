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

	// get all the shapes from current compartment
	shapes := listShapes(compartmentID)

	availabilityDomains := identity.ListAvailabilityDomains()
	request := core.LaunchInstanceRequest{}
	request.AvailabilityDomain = availabilityDomains[0].Name
	request.CompartmentId = compartmentID
	request.Shape = shapes[0].Shape
	request.DisplayName = common.String("OCI_GOSDK_Sample_Instance")

	subnet := CreateSubnet()
	request.SubnetId = subnet.Id

	image := ListImages()[0]
	request.ImageId = image.Id

	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helper.LogIfError(err)

	r, err := c.LaunchInstance(context.Background(), request)
	helper.LogIfError(err)

	return r.Instance
}

func listShapes(compartmentID *string) []core.Shape {
	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helper.LogIfError(err)

	request := core.ListShapesRequest{
		CompartmentId: compartmentID,
	}

	r, err := c.ListShapes(context.Background(), request)
	helper.LogIfError(err)

	if r.Items == nil || len(r.Items) == 0 {
		log.Fatalln("Invalid response from ListShapes")
	}

	return r.Items
}
