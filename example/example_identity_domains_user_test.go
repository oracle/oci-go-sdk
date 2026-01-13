// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Identity Domains User API
// This script provides examples of CRUD (create, read, update, delete) operations for Identity Domain User API using Go SDK.
// This script will:
//
//    * Read user configuration and IdentityDomainEndpoint
//    * Construct IdentityDomainsClient
//    * Create a user in the provided domain
//    * Delete the user when all operations are done
//    * Get the user by ID
//    * Find the user by displayName using LIST method
//    * Find the user by displayName using POST method
//    * Replace the user using PUT method
//    * Update the user using PATCH method
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
//              'go test -v -run ^TestIdentityDomainsUser$ github.com/oracle/oci-go-sdk/v65/example'
//          ```
// 			Further reading: https://docs.oracle.com/en-us/iaas/Content/Identity/domains/to-view-identity-domains.htm#view-identity-domains
// 			OCI config docs: https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm
//

package example

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/identitydomains"
)

var (
	DomainEndpoint string
)

func TestIdentityDomainsUser(t *testing.T) {
	// Parse environment variables to get DomainEndpoint
	identityDomainsParseEnvironmentVariables()

	// read user configuration and create a new IdentityDomainsClient instance
	config := common.DefaultConfigProvider()
	client, _ := identitydomains.NewIdentityDomainsClientWithConfigurationProvider(config, DomainEndpoint)

	user := createUser(client)
	defer deleteUser(client, user)
	getUserByID(client, *user.Id)
	findUserByDisplayName(client, *user.DisplayName)
	findUserByDisplayNameUsingPost(client, *user.DisplayName)
	replaceUser(client, user)
	updateUser(client, user)
}

func createUser(client identitydomains.IdentityDomainsClient) identitydomains.User {
	now := time.Now().Unix()
	displayName := fmt.Sprintf("go_sdk_user_displayName_%d", now)
	userName := fmt.Sprintf("go_sdk_user_userName_%d@example.com", now)
	familyName := fmt.Sprintf("go_sdk_user_familyName_%d", now)
	primaryEmail := true
	emails := []identitydomains.UserEmails{{
		Value:   &userName,
		Type:    identitydomains.UserEmailsTypeWork,
		Primary: &primaryEmail,
	}}
	schemas := []string{"urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"}
	req := identitydomains.CreateUserRequest{
		User: identitydomains.User{
			DisplayName: &displayName,
			UserName:    &userName,
			Name:        &identitydomains.UserName{FamilyName: &familyName},
			Emails:      emails,
			Schemas:     schemas,
		},
		AttributeSets: []identitydomains.AttributeSetsEnum{identitydomains.AttributeSetsAll},
	}

	res, err := client.CreateUser(context.Background(), req)
	if err != nil {
		redText := "\033[31m"
		regularText := "\033[0;39m"
		panic(fmt.Errorf("%sFailed creating user%s: %v", redText, regularText, err))
	}
	if res.HTTPResponse().StatusCode != 201 {
		panic("create user response != 201")
	}
	user := res.User
	printUser("Created user: ", user)
	return user
}

func deleteUser(client identitydomains.IdentityDomainsClient, user identitydomains.User) {
	req := identitydomains.DeleteUserRequest{UserId: user.Id}
	res, err := client.DeleteUser(context.Background(), req)
	if err != nil {
		panic(err)
	}
	if res.HTTPResponse().StatusCode != 204 {
		panic("delete user response != 204")
	}
	printUser("Deleted user: ", user)
}

func getUserByID(client identitydomains.IdentityDomainsClient, id string) {
	attributes := "displayName,userName"
	req := identitydomains.GetUserRequest{UserId: &id, Attributes: &attributes}
	res, err := client.GetUser(context.Background(), req)
	if err != nil {
		panic(err)
	}
	if res.HTTPResponse().StatusCode != 200 {
		panic("get user response != 200")
	}
	printUser("Found user by id:", res.User)
}

func findUserByDisplayName(client identitydomains.IdentityDomainsClient, displayName string) {
	attributes := "displayName,userName"
	filter := fmt.Sprintf("displayName eq \"%s\"", displayName)
	count := 1
	startIndex := 1
	req := identitydomains.ListUsersRequest{Filter: &filter, Count: &count, StartIndex: &startIndex, Attributes: &attributes}
	res, err := client.ListUsers(context.Background(), req)
	if err != nil {
		panic(err)
	}
	if res.HTTPResponse().StatusCode != 200 {
		panic("find user by display name response != 200")
	}
	printUser("Found user by displayName:", res.Resources[0])
}

func findUserByDisplayNameUsingPost(client identitydomains.IdentityDomainsClient, displayName string) {
	attributes := []string{"displayName", "userName"}
	filter := fmt.Sprintf("displayName eq \"%s\"", displayName)
	count := 1
	startIndex := 1
	schemas := []string{"urn:ietf:params:scim:api:messages:2.0:SearchRequest"}
	req := identitydomains.UserSearchRequest{Attributes: attributes, Count: &count, StartIndex: &startIndex, Filter: &filter, Schemas: schemas}
	res, err := client.SearchUsers(context.Background(), identitydomains.SearchUsersRequest{UserSearchRequest: req})
	if err != nil {
		panic(err)
	}
	if res.HTTPResponse().StatusCode != 200 {
		panic("find user by display name using POST response != 200")
	}
	printUser("Found user by displayName using POST:", res.Resources[0])
}

func replaceUser(client identitydomains.IdentityDomainsClient, user identitydomains.User) {
	newUser := user
	*newUser.DisplayName = "replaced display name"
	newUser.IdcsCreatedBy = nil
	newUser.IdcsLastModifiedBy = nil
	req := identitydomains.PutUserRequest{UserId: user.Id, User: newUser}
	res, err := client.PutUser(context.Background(), req)
	if err != nil {
		panic(err)
	}
	if res.HTTPResponse().StatusCode != 200 {
		panic("replace user response != 200")
	}
	printUser("Replaced user (replaced display name):", res.User)
}

func updateUser(client identitydomains.IdentityDomainsClient, user identitydomains.User) {
	path := "displayName"
	var value interface{} = "updated display name"
	operations := []identitydomains.Operations{{Op: identitydomains.OperationsOpReplace, Path: &path, Value: &value}}
	schemas := []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"}
	patchBody := identitydomains.PatchOp{Operations: operations, Schemas: schemas}
	req := identitydomains.PatchUserRequest{UserId: user.Id, PatchOp: patchBody}
	res, err := client.PatchUser(context.Background(), req)
	if err != nil {
		panic(err)
	}
	if res.HTTPResponse().StatusCode != 200 {
		panic("update user response != 200")
	}
	printUser("Updated user (updated display name):", res.User)
}

func printUser(msg string, u identitydomains.User) {
	log.Printf("-------------------------------------------------")
	log.Printf("%s", msg)
	if u.UserName != nil {
		log.Printf("     User name: %s", *u.UserName)
	}
	if u.DisplayName != nil {
		log.Printf("     Display name: %s", *u.DisplayName)
	}
	if u.Ocid != nil {
		log.Printf("     Ocid: %s", *u.Ocid)
	}
	if u.Description != nil {
		log.Printf("     Description: %s", *u.Description)
	}
	if u.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags != nil && u.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags.DefinedTags != nil {
		log.Printf("     Defined Tags: %s", u.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags.DefinedTags)
	}
	log.Printf(" ")
}

func identityDomainsParseEnvironmentVariables() {

	DomainEndpoint = os.Getenv("DOMAIN_ENDPOINT")

	if DomainEndpoint == "" {
		panic("Please specify DOMAIN_ENDPOINT in the env.")
	}

	log.Printf("DOMAIN_ENDPOINT: %s", DomainEndpoint)
}
