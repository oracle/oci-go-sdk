// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example for Resource Search Service API
// Search for resources across your cloud infrastructure

package example

import (
	"context"
	"fmt"
	"log"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/resourcesearch"
)

func Example_ResourceSearch() {
	client, err := resourcesearch.NewResourceSearchClientWithConfigurationProvider(common.DefaultConfigProvider())
	ctx := context.Background()
	helpers.FatalIfError(err)

	// list resource types
	listReq := resourcesearch.ListResourceTypesRequest{}
	listResp, err := client.ListResourceTypes(ctx, listReq)
	fmt.Println("list resource types")

	for _, element := range listResp.Items {
		log.Printf("Resource: %s", *element.Name)
	}

	// get group type details
	getReq := resourcesearch.GetResourceTypeRequest{
		Name: common.String("Group"),
	}
	getResp, err := client.GetResourceType(context.Background(), getReq)
	helpers.FatalIfError(err)
	fmt.Println("get group type details")
	log.Printf("Resource type: %s", getResp.ResourceType)

	// search resource by freetext
	searchReq := resourcesearch.SearchResourcesRequest{
		SearchDetails: resourcesearch.FreeTextSearchDetails{
			Text: common.String("displayname"),
		},
	}

	freeSearchResp, err := client.SearchResources(context.Background(), searchReq)
	helpers.FatalIfError(err)
	fmt.Println("search resource by freetext")

	for _, element := range freeSearchResp.Items {
		log.Printf("Resource: %s", element)
	}

	searchReq.SearchDetails = resourcesearch.StructuredSearchDetails{
		MatchingContextType: resourcesearch.SearchDetailsMatchingContextTypeHighlights,
		Query:               common.String("query all resources"),
	}

	structureSearchResp, err := client.SearchResources(context.Background(), searchReq)
	helpers.FatalIfError(err)

	// search resource by structured query
	fmt.Println("search resource by structured query")

	for _, element := range structureSearchResp.Items {
		log.Printf("Resource: %s", element)
	}

	// Output:
	// list resource types
	// get group type details
	// search resource by freetext
	// search resource by structured query
}
