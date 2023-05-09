// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Identity Domains Dynamic Resource Group API
// This script provides examples of CRUD (create, read, update, delete) operations for Identity Domain User API using Go SDK.
// This script will:
//
//    * Read user configuration and IdentityDomainEndpoint
//    * Construct IdentityDomainsClient
//    * Create a DynamicResourceGroup in the provided domain
//    * Delete the DynamicResourceGroup when all operations are done
//    * Get the DynamicResourceGroup by ID
//    * Find the DynamicResourceGroup by displayName using LIST method
//    * Find the DynamicResourceGroup by displayName using POST method
//    * Replace the DynamicResourceGroup using PUT method
//    * Update the DynamicResourceGroup using PATCH method
//
//  This script takes the following values from environment variables
//
//    * DOMAIN_ENDPOINT    - the domain endpoint
// 	  	    To find Domain URL, navigate to Identity > Domains in OCI console, choose relevant domain,
// 			then in the overview page, find "Domain URL" under "Domain Information", click "Copy".
// 			Should look like: domain_endpoint=https://idcs-...identity.oraclecloud.com.
//          Run this test with the copied endpoint, example command:
//          ```
// 			    DOMAIN_ENDPOINT=https://idcs-...identity.oraclecloud.com bash -c \
//              'go test -v -run ^TestIdentityDomainsDynamicResourceGroup$ github.com/oracle/oci-go-sdk/v65/example'
//          ```
// 			Further reading: https://docs.oracle.com/en-us/iaas/Content/Identity/domains/to-view-identity-domains.htm#view-identity-domains
// 			OCI config docs: https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm
//

package example

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/identitydomains"
)

func TestIdentityDomainsDynamicResourceGroup(t *testing.T) {
	// Parse environment variables to get DomainEndpoint
	identityDomainsParseEnvironmentVariables()

	// read user configuration and create a new IdentityDomainsClient instance
	config := common.DefaultConfigProvider()
	client, _ := identitydomains.NewIdentityDomainsClientWithConfigurationProvider(config, DomainEndpoint)

	dynamicResourceGroup := createDynamicResourceGroup(client)
	defer deleteDynamicResourceGroup(client, dynamicResourceGroup)
	getDynamicResourceGroupById(client, *dynamicResourceGroup.Id)
	findDynamicResourceGroupByDisplayName(client, *dynamicResourceGroup.DisplayName)
	findDynamicResourceGroupByDisplayNameUsingPost(client, *dynamicResourceGroup.DisplayName)
	replaceDynamicResourceGroup(client, dynamicResourceGroup)
	updateDynamicResourceGroup(client, dynamicResourceGroup)
}

func createDynamicResourceGroup(client identitydomains.IdentityDomainsClient) identitydomains.DynamicResourceGroup {
	displayName := fmt.Sprintf("go_sdk_displayName_%d", time.Now().Unix())
	description := "go sdk description"
	matchingRule := "Any {Any {instance.id = 'ocid1.instance.oc1.iad..exampleuniqueid1', instance.compartment.id = 'ocid1.compartment.oc1..exampleuniqueid2'}}"
	schemas := []string{"urn:ietf:params:scim:schemas:oracle:idcs:DynamicResourceGroup"}
	req := identitydomains.CreateDynamicResourceGroupRequest{
		DynamicResourceGroup: identitydomains.DynamicResourceGroup{
			DisplayName:  &displayName,
			Description:  &description,
			MatchingRule: &matchingRule,
			Schemas:      schemas,
		},
		AttributeSets: []identitydomains.AttributeSetsEnum{identitydomains.AttributeSetsAll},
	}
	res, err := client.CreateDynamicResourceGroup(context.Background(), req)
	if err != nil {
		panic(fmt.Errorf("Failed creating dynamic resource group: %v", err))
	}
	if res.HTTPResponse().StatusCode != 201 {
		panic("create dynamic resource group response != 201")
	}
	dynamicResourceGroup := res.DynamicResourceGroup
	printDynamicResourceGroup("Created dynamic resource group:", dynamicResourceGroup)
	return dynamicResourceGroup
}

func deleteDynamicResourceGroup(client identitydomains.IdentityDomainsClient, dynamicResourceGroup identitydomains.DynamicResourceGroup) {
	req := identitydomains.DeleteDynamicResourceGroupRequest{DynamicResourceGroupId: dynamicResourceGroup.Id}
	res, err := client.DeleteDynamicResourceGroup(context.Background(), req)
	if err != nil {
		panic(fmt.Errorf("Failed deleting dynamic resource group: %v", err))
	}
	if res.HTTPResponse().StatusCode != 204 {
		panic("delete dynamic resource group response != 204")
	}
	printDynamicResourceGroup("Deleted dynamic resource group:", dynamicResourceGroup)
}

func getDynamicResourceGroupById(client identitydomains.IdentityDomainsClient, id string) {
	attributes := "displayName,description,matchingRule"
	req := identitydomains.GetDynamicResourceGroupRequest{DynamicResourceGroupId: &id, Attributes: &attributes}
	res, err := client.GetDynamicResourceGroup(context.Background(), req)
	if err != nil {
		panic(fmt.Errorf("Failed getting dynamic resource group by id: %v", err))
	}
	if res.HTTPResponse().StatusCode != 200 {
		panic("get dynamic resource group response != 200")
	}
	printDynamicResourceGroup("Found dynamic resource group by id:", res.DynamicResourceGroup)
}

func findDynamicResourceGroupByDisplayName(client identitydomains.IdentityDomainsClient, displayName string) {
	attributes := "displayName,description,matchingRule"
	filter := fmt.Sprintf("displayName eq \"%s\"", displayName)
	count := 1
	startIndex := 1
	req := identitydomains.ListDynamicResourceGroupsRequest{Filter: &filter, Count: &count, StartIndex: &startIndex, Attributes: &attributes}
	res, err := client.ListDynamicResourceGroups(context.Background(), req)
	if err != nil {
		panic(fmt.Errorf("Failed finding dynamic resource group by display name: %v", err))
	}
	if res.HTTPResponse().StatusCode != 200 {
		panic("find dynamic resource group by display name response != 200")
	}
	printDynamicResourceGroup("Found dynamic resource group by displayName:", res.Resources[0])
}

func findDynamicResourceGroupByDisplayNameUsingPost(client identitydomains.IdentityDomainsClient, displayName string) {
	attributes := []string{"displayName", "description", "matchingRule"}
	filter := fmt.Sprintf("displayName eq \"%s\"", displayName)
	count := 1
	startIndex := 1
	schemas := []string{"urn:ietf:params:scim:api:messages:2.0:SearchRequest"}
	req := identitydomains.DynamicResourceGroupSearchRequest{Attributes: attributes, Count: &count, StartIndex: &startIndex, Filter: &filter, Schemas: schemas}
	res, err := client.SearchDynamicResourceGroups(context.Background(), identitydomains.SearchDynamicResourceGroupsRequest{DynamicResourceGroupSearchRequest: req})
	if err != nil {
		panic(fmt.Errorf("Failed finding dynamic resource group by display name using POST: %v", err))
	}
	if res.HTTPResponse().StatusCode != 200 {
		panic("find dynamic resource group by display name using POST response != 200")
	}
	printDynamicResourceGroup("Found dynamic resource group by displayName using POST:", res.Resources[0])
}

func replaceDynamicResourceGroup(client identitydomains.IdentityDomainsClient, dynamicResourceGroup identitydomains.DynamicResourceGroup) {
	newDynamicResourceGroup := dynamicResourceGroup
	*newDynamicResourceGroup.Description = "new description"
	newDynamicResourceGroup.IdcsCreatedBy = nil
	newDynamicResourceGroup.IdcsLastModifiedBy = nil
	req := identitydomains.PutDynamicResourceGroupRequest{DynamicResourceGroupId: dynamicResourceGroup.Id, DynamicResourceGroup: newDynamicResourceGroup}
	res, err := client.PutDynamicResourceGroup(context.Background(), req)
	if err != nil {
		panic(fmt.Errorf("Failed replacing dynamic resource group: %v", err))
	}
	if res.HTTPResponse().StatusCode != 200 {
		panic("replace dynamic resource group response != 200")
	}
	printDynamicResourceGroup("Replaced dynamic resource group (updated description):", res.DynamicResourceGroup)
}

func updateDynamicResourceGroup(client identitydomains.IdentityDomainsClient, dynamicResourceGroup identitydomains.DynamicResourceGroup) {
	path := "description"
	var value interface{} = "updated description"
	operations := []identitydomains.Operations{{Op: identitydomains.OperationsOpReplace, Path: &path, Value: &value}}
	schemas := []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"}
	patchBody := identitydomains.PatchOp{Operations: operations, Schemas: schemas}
	req := identitydomains.PatchDynamicResourceGroupRequest{DynamicResourceGroupId: dynamicResourceGroup.Id, PatchOp: patchBody}
	res, err := client.PatchDynamicResourceGroup(context.Background(), req)
	if err != nil {
		panic(fmt.Errorf("Failed updating dynamic resource group: %v", err))
	}
	if res.HTTPResponse().StatusCode != 200 {
		panic("update dynamic resource group response != 200")
	}
	printDynamicResourceGroup("Updated dynamic resource group description:", res.DynamicResourceGroup)
}

func printDynamicResourceGroup(msg string, drg identitydomains.DynamicResourceGroup) {
	log.Printf("-------------------------------------------------")
	log.Printf("%s", msg)
	if drg.DisplayName != nil {
		log.Printf("     Display name: %s", *drg.DisplayName)
	}
	if drg.Ocid != nil {
		log.Printf("     Ocid: %s", *drg.Ocid)
	}
	if drg.Description != nil {
		log.Printf("     Description: %s", *drg.Description)
	}
	if drg.MatchingRule != nil {
		log.Printf("     MatchingRule: %s", *drg.MatchingRule)
	}
	if drg.Schemas != nil {
		log.Printf("     Schemas: %v", drg.Schemas)
	}
	log.Printf(" ")
}
