package core

import (
	"context"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/example/helper"
	"github.com/oracle/oci-go-sdk/example/identity"
)

// ListImages Lists the available images in the specified compartment.
func ListImages() []core.Image {
	// create a compartment, otherwise return the id of existing compartment
	compartmentID, err := identity.CreateCompartment()
	helper.LogIfError(err)

	request := core.ListImagesRequest{
		CompartmentId: compartmentID,
	}

	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helper.LogIfError(err)

	r, err := c.ListImages(context.Background(), request)
	helper.LogIfError(err)

	return r.Items
}
