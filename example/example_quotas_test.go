// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for demonstrating how Quotas can be managed using the OCI Golang SDK
// This example will perform the following operations sequentially-
// - Create a Quota
// - Get the created Quota
// - List all Quotas
// - Update the previously created Quota
// - Delete this Quota

// Description of common parameters
// compartmentId	: The OCID of the compartment where Quotas will reside (this has to be the root compartment)
// name			    : Name of the Quota
// description		: Description for the Quota
// statements		: An array of Quota statements written in the declarative language

package example

// Import necessary packages
import (
	"context"                             // To supply to the Quotas client while making requests
	"fmt"                                 // To print to the console
	"github.com/oracle/oci-go-sdk/common" // For common OCI types
	"github.com/oracle/oci-go-sdk/limits" // For types and methods corresponding to Limits
)

// Creates a new Quota with the details given in createQuotaDetails
func createQuota(client limits.QuotasClient, createQuotaDetails limits.CreateQuotaDetails) limits.CreateQuotaResponse {
	var response limits.CreateQuotaResponse
	response, err := client.CreateQuota(context.Background(), limits.CreateQuotaRequest{CreateQuotaDetails: createQuotaDetails})
	if err != nil {
		panic(err)
	}
	return response
}

// Gets the Quota corresponding to given quotaId
func getQuota(client limits.QuotasClient, quotaId string) limits.GetQuotaResponse {
	var response limits.GetQuotaResponse
	quotaIdStr := common.String(quotaId)
	response, err := client.GetQuota(context.Background(), limits.GetQuotaRequest{QuotaId: quotaIdStr})
	if err != nil {
		panic(err)
	}
	return response
}

// Lists Quotas under the Compartment corresponding to given compartmentId
func listQuotas(client limits.QuotasClient, compartmentId string) limits.ListQuotasResponse {
	var response limits.ListQuotasResponse
	compartmentIdStr := common.String(compartmentId)
	response, err := client.ListQuotas(context.Background(), limits.ListQuotasRequest{CompartmentId: compartmentIdStr})
	if err != nil {
		panic(err)
	}
	return response
}

// Updates the Quota corresponding to given quotaId with values given in updateQuotaDetails
func updateQuota(client limits.QuotasClient, quotaId string, updateQuotaDetails limits.UpdateQuotaDetails) limits.UpdateQuotaResponse {
	var response limits.UpdateQuotaResponse
	quotaIdStr := common.String(quotaId)
	response, err := client.UpdateQuota(context.Background(), limits.UpdateQuotaRequest{QuotaId: quotaIdStr, UpdateQuotaDetails: updateQuotaDetails})
	if err != nil {
		panic(err)
	}
	return response
}

// Deletes the Quota corresponding to given quotaId
func deleteQuota(client limits.QuotasClient, quotaId string) limits.DeleteQuotaResponse {
	var response limits.DeleteQuotaResponse
	quotaIdStr := common.String(quotaId)
	response, err := client.DeleteQuota(context.Background(), limits.DeleteQuotaRequest{QuotaId: quotaIdStr})
	if err != nil {
		panic(err)
	}
	return response
}

// ExampleQuotas runs an example demonstrating the use of OCI Golang SDK for managing Quotas
func ExampleQuotas() {

	// Initialize default config provider
	configProvider := common.DefaultConfigProvider()
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		panic(err)
	}

	// Initialize sample inputs
	compartmentId, _ := configProvider.TenancyOCID()
	name := "MyQuota"
	description := "This is a sample Quota"
	newDescription := "This is an updated Quota"
	statements := []string{"Zero test-family quota 'test-quota-1' in tenancy"}

	// Initialize Quotas client
	client, err := limits.NewQuotasClientWithConfigurationProvider(configProvider)
	if err != nil {
		panic(err)
	}

	// Create Quota
	fmt.Println("Creating Quota")
	createQuotaDetails := limits.CreateQuotaDetails{CompartmentId: &compartmentId, Name: &name, Description: &description, Statements: statements}
	createResponse := createQuota(client, createQuotaDetails)

	// Get Quota
	fmt.Println("Getting Quota")
	getQuota(client, *createResponse.Quota.Id)

	// List Quotas
	fmt.Println("Listing Quotas")
	listResponse := listQuotas(client, compartmentId)

	// Update Quota
	fmt.Println("Updating Quota")
	quotaId := *listResponse.Items[0].Id
	updateQuotaDetails := limits.UpdateQuotaDetails{Description: &newDescription}
	updateQuota(client, quotaId, updateQuotaDetails)

	// Delete Quota
	fmt.Println("Deleting Quota")
	deleteQuota(client, quotaId)

	fmt.Println("ExampleQuotas completed")

	// Output:
	// Creating Quota
	// Getting Quota
	// Listing Quotas
	// Updating Quota
	// Deleting Quota
	// ExampleQuotas completed
}
