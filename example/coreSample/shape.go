// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package coreSample

import (
	"context"
	"fmt"
	"log"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/example"
)

// ListShapesWithPagination demostrate how to use page parameter
func ListShapesWithPagination(compartmentID *string) {
	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	example.LogIfError(err)

	request := core.ListShapesRequest{
		CompartmentId: compartmentID,
	}

	// to show how pagination works, reduce number of items to return in a paginated "List" call
	request.Limit = common.Int(2)

	listShapesFunc := func(request core.ListShapesRequest) (core.ListShapesResponse, error) {
		return c.ListShapes(context.Background(), request)
	}

	for r, err := listShapesFunc(request); ; r, err = listShapesFunc(request) {
		example.LogIfError(err)

		fmt.Printf("list shapes result: %v", r.Items)

		if r.OpcNextPage != nil {
			// if there are more items in next page, fetch items from next page
			request.Page = r.OpcNextPage
		} else {
			// no next page, break the loop
			break
		}
	}

	return
}

// ListShapes Lists the shapes that can be used to launch an instance within the specified compartment.
func ListShapes(compartmentID *string, imageID *string) []core.Shape {
	c, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	example.LogIfError(err)

	request := core.ListShapesRequest{
		CompartmentId: compartmentID,
		ImageId:       imageID,
	}

	r, err := c.ListShapes(context.Background(), request)
	example.LogIfError(err)

	if r.Items == nil || len(r.Items) == 0 {
		log.Fatalln("Invalid response from ListShapes")
	}

	return r.Items
}
