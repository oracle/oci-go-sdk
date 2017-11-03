// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package integtest

import (
	"context"
	"fmt"
	"testing"

	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/identity"
	"github.com/stretchr/testify/assert"
)

var (
	rootTestCompartmentID = "ocidv1:tenancy:oc1:phx:1460406592660:aaaaaaaab4faofrfkxecohhjuivjq262pu"
	validUserID           = "ocid1.user.oc1..aaaaaaaav6gsclr6pd4yjqengmriylyck55lvon5ujjnhkok5gyxii34lvra"
	validCompartmetnID    = "ocid1.compartment.oc1..aaaaaaaa5dvrjzvfn3rub24nczhih3zb3a673b6tmbvpng3j5apobtxshlma"
	validGroupID          = "ocid1.group.oc1..aaaaaaaayvxomawkk23wkp32cgdufufgqvx62qanmbn6vs3lv65xuc42r5sq"
	region                = common.REGION_PHX
)

func panicIfError(t *testing.T, err error) {
	if err != nil {
		t.Fail()
		panic(err)
	}
}

// Group operations CRUD
func TestIdentityClient_GroupCRUD(t *testing.T) {
	c := identity.NewClientForRegion(region)
	request := identity.CreateGroupRequest{}
	request.CompartmentID = &rootTestCompartmentID
	request.Name = common.String("GoSDK2_someGroup")
	request.Description = common.String("GoSDK_someGroupDesc")
	r, err := c.CreateGroup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	panicIfError(t, err)

	//Get
	rRead := identity.GetGroupRequest{GroupID: r.ID}
	resRead, err := c.GetGroup(context.Background(), rRead)
	assert.NotEmpty(t, r, fmt.Sprint(resRead.ID))
	panicIfError(t, err)

	//Update
	rUpdate := identity.UpdateGroupRequest{GroupID: r.ID}
	rUpdate.Description = common.String("New description")
	resUpdate, err := c.UpdateGroup(context.Background(), rUpdate)
	panicIfError(t, err)
	assert.NotNil(t, resUpdate.ID)

	//Delete
	rDel := identity.DeleteGroupRequest{GroupID: r.ID}
	err = c.DeleteGroup(context.Background(), rDel)
	assert.NoError(t, err)
}

func TestIdentityClient_ListGroups(t *testing.T) {
	c := identity.NewClientForRegion(region)
	request := identity.ListGroupsRequest{CompartmentID: &rootTestCompartmentID}
	r, err := c.ListGroups(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	assert.NotEmpty(t, r.Items)
	assert.NotEmpty(t, r.OpcNextPage)

	return
}

// Compartment operations
//Can not delete compartments right now! Careful!
/*
func TestIdentityClient_CreateCompartment(t *testing.T) {
	c:= identity.NewClientForRegion(region)
	request:= identity.CreateCompartmentRequest{CreateCompartmentDetails:identity.CreateCompartmentDetails{
		Name:"egTest2",
		Description:"egTest2_descp",
		CompartmentID:rootTestCompartmentID,
	}}
	r, err:= identity.CreateCompartment(c, request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}
*/

//Comparment RU
func TestIdentityClient_UpdateCompartment(t *testing.T) {
	c := identity.NewClientForRegion(region)
	//Update
	request := identity.UpdateCompartmentRequest{UpdateCompartmentDetails: identity.UpdateCompartmentDetails{
		Name:        common.String("GOSDK2_Test"),
		Description: common.String("GOSDK2 description2"),
	},
		CompartmentID: common.String(validCompartmetnID),
	}
	r, err := c.UpdateCompartment(context.Background(), request)
	panicIfError(t, err)
	assert.NotEmpty(t, r, fmt.Sprint(r))

	rRead := identity.GetCompartmentRequest{CompartmentID: common.String(rootTestCompartmentID)}
	resRead, err := c.GetCompartment(context.Background(), rRead)
	panicIfError(t, err)
	assert.NotEmpty(t, r, fmt.Sprint(resRead))
	return
}

//User Operations
func TestIdentityClient_ListUsers(t *testing.T) {
	c := identity.NewClientForRegion(region)
	request := identity.ListUsersRequest{CompartmentID: common.String(rootTestCompartmentID)}
	r, err := c.ListUsers(context.Background(), request)
	assert.NotEmpty(t, r.Items, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UserCRUD(t *testing.T) {
	c := identity.NewClientForRegion(region)
	request := identity.CreateUserRequest{}
	request.CompartmentID = common.String(rootTestCompartmentID)
	request.Name = common.String("GolangSDK2_testUser")
	request.Description = common.String("Test user for golagn sdk2")
	resCreate, err := c.CreateUser(context.Background(), request)
	fmt.Println(resCreate)
	assert.NotEmpty(t, resCreate, fmt.Sprint(resCreate))
	assert.NoError(t, err)

	//Read
	rRead := identity.GetUserRequest{UserID: resCreate.ID}
	resRead, err := c.GetUser(context.Background(), rRead)
	assert.NotEmpty(t, resRead, fmt.Sprint(resRead))
	assert.NoError(t, err)

	//Update
	rUpdate := identity.UpdateUserRequest{}
	rUpdate.UserID = resCreate.ID
	rUpdate.Description = common.String("This is a new description")
	resUpdate, err := c.UpdateUser(context.Background(), rUpdate)
	assert.NotEmpty(t, resUpdate, fmt.Sprint(resUpdate))
	assert.NoError(t, err)

	//remove
	rDelete := identity.DeleteUserRequest{}
	rDelete.UserID = resCreate.ID
	err = c.DeleteUser(context.Background(), rDelete)
	assert.NoError(t, err)

	return
}

//User-Group operations
func TestIdentityClient_AddUserToGroup(t *testing.T) {
	c := identity.NewClientForRegion(region)
	//add
	requestAdd := identity.AddUserToGroupRequest{}
	requestAdd.UserID = common.String(validUserID)
	requestAdd.GroupID = common.String(validGroupID)
	r, err := c.AddUserToGroup(context.Background(), requestAdd)
	panicIfError(t, err)
	assert.NotEmpty(t, r, fmt.Sprint(r))

	//remove
	requestRemove := identity.RemoveUserFromGroupRequest{UserGroupMembershipID: r.UserGroupMembership.ID}
	err = c.RemoveUserFromGroup(context.Background(), requestRemove)
	panicIfError(t, err)
	return
}

//Policy Operations see DEX-1945
//func TestIdentityClient_PolicyCRUD(t *testing.T) {
//	//Create
//	client := identity.NewClientForRegion(region)
//	/*
//	createRequest := identity.CreatePolicyRequest{}
//	createRequest.CompartmentID = rootTestCompartmentID
//	createRequest.Name = "goSDK2Policy2"
//	createRequest.Description = "some policy"
//	createRequest.Statements = []string{"Allow group goSDK2CreateGroup read all-resources on compartment egineztest"}
//	createRequest.VersionDate = time.Now()
//	createResponse, err := client.CreatePolicy(context.Background(), createRequest)
//	assert.NotEmpty(t, createResponse, fmt.Sprint(createResponse))
//	assert.NoError(t, err)
//	*/
//	createResponseID :=  ""
//
//	//Read
//	readRequest := identity.GetPolicyRequest{}
//	readRequest.PolicyID = createResponseID
//	readResponse, err := client.GetPolicy(context.Background(), readRequest)
//	assert.NotEmpty(t, readResponse, fmt.Sprint(readResponse))
//	assert.NoError(t, err)
//
//	//Update
//	/*
//	updateRequest := identity.UpdatePolicyRequest{}
//	updateRequest.PolicyID = createResponseID
//	updateRequest.Description = "new descripption"
//	updateResponse, err := client.UpdatePolicy(context.Background(), updateRequest)
//	assert.NotEmpty(t, updateResponse, fmt.Sprint(updateResponse))
//	assert.NoError(t, err)
//	*/
//
//	request := identity.DeletePolicyRequest{PolicyID:createResponseID}
//	err = client.DeletePolicy(context.Background(), request)
//	assert.NoError(t, err)
//
//	return
//}
//

//SecretKey operations
func TestIdentityClient_SecretKeyCRUD(t *testing.T) {
	c := identity.NewClientForRegion(region)
	request := identity.CreateCustomerSecretKeyRequest{}
	request.UserID = common.String(validUserID)
	request.DisplayName = common.String("GolangSDK2TestSecretKey")
	resCreate, err := c.CreateCustomerSecretKey(context.Background(), request)
	panicIfError(t, err)
	assert.NotEmpty(t, resCreate, fmt.Sprint(resCreate))
	assert.NoError(t, err)

	//Update
	rUpdate := identity.UpdateCustomerSecretKeyRequest{}
	rUpdate.CustomerSecretKeyID = resCreate.ID
	rUpdate.UserID = common.String(validUserID)
	rUpdate.DisplayName = common.String("This is a new description")
	resUpdate, err := c.UpdateCustomerSecretKey(context.Background(), rUpdate)
	assert.NotEmpty(t, resUpdate, fmt.Sprint(resUpdate))
	panicIfError(t, err)

	//remove
	rDelete := identity.DeleteCustomerSecretKeyRequest{}
	rDelete.CustomerSecretKeyID = resCreate.ID
	rDelete.UserID = common.String(validUserID)
	err = c.DeleteCustomerSecretKey(context.Background(), rDelete)
	panicIfError(t, err)

	return
}

func TestIdentityClient_ListCustomerSecretKeys(t *testing.T) {
	c := identity.NewClientForRegion(region)
	request := identity.ListCustomerSecretKeysRequest{}
	request.UserID = common.String(validUserID)
	r, err := c.ListCustomerSecretKeys(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	panicIfError(t, err)
	return
}

//Apikeys
func TestIdentityClient_ApiKeyCRUD(t *testing.T) {
	userID := ""
	c := identity.NewClientForRegion(region)
	request := identity.UploadApiKeyRequest{}
	request.UserID = common.String(userID)
	request.Key = common.String("some key")
	resCreate, err := c.UploadApiKey(context.Background(), request)
	assert.Error(t, err)
	ser, ok := common.IsServiceError(err)
	assert.True(t, ok)
	assert.NotEmpty(t, ser.GetMessage())

	//remove
	rDelete := identity.DeleteApiKeyRequest{}
	rDelete.Fingerprint = resCreate.Fingerprint
	rDelete.UserID = common.String(userID)
	err = c.DeleteApiKey(context.Background(), rDelete)
	ser, ok = common.IsServiceError(err)
	assert.False(t, ok)
	assert.NotEmpty(t, err.Error())

	return
}
func TestIdentityClient_ListApiKeys(t *testing.T) {
	c := identity.NewClientForRegion(region)
	request := identity.ListApiKeysRequest{}
	request.UserID = common.String(validUserID)
	r, err := c.ListApiKeys(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	panicIfError(t, err)
	return
}

//TODO
//func TestIdentityClient_CreateIdentityProvider(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.CreateIdentityProviderRequest{}
//	r, err := c.CreateIdentityProvider(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_CreateIdpGroupMapping(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.CreateIdpGroupMappingRequest{}
//	r, err := c.CreateIdpGroupMapping(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_CreateOrResetUIPassword(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.CreateOrResetUIPasswordRequest{}
//	r, err := c.CreateOrResetUIPassword(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_CreateRegionSubscription(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.CreateRegionSubscriptionRequest{}
//	r, err := c.CreateRegionSubscription(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_CreateSwiftPassword(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.CreateSwiftPasswordRequest{}
//	r, err := c.CreateSwiftPassword(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_DeleteApiKey(t *testing.T) {
//}
//
//func TestIdentityClient_DeleteIdentityProvider(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.DeleteIdentityProviderRequest{}
//	err := c.DeleteIdentityProvider(context.Background(), request)
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_DeleteIdpGroupMapping(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.DeleteIdpGroupMappingRequest{}
//	err := c.DeleteIdpGroupMapping(context.Background(), request)
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_DeleteSwiftPassword(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.DeleteSwiftPasswordRequest{}
//	err := c.DeleteSwiftPassword(context.Background(), request)
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_GetIdentityProvider(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.GetIdentityProviderRequest{}
//	r, err := c.GetIdentityProvider(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_GetIdpGroupMapping(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.GetIdpGroupMappingRequest{}
//	r, err := c.GetIdpGroupMapping(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_GetTenancy(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.GetTenancyRequest{}
//	r, err := c.GetTenancy(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_GetUserGroupMembership(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.GetUserGroupMembershipRequest{}
//	r, err := c.GetUserGroupMembership(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_ListAvailabilityDomains(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.ListAvailabilityDomainsRequest{}
//	r, err := c.ListAvailabilityDomains(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_ListCompartments(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.ListCompartmentsRequest{}
//	r, err := c.ListCompartments(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_ListIdentityProviders(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.ListIdentityProvidersRequest{}
//	r, err := c.ListIdentityProviders(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_ListIdpGroupMappings(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.ListIdpGroupMappingsRequest{}
//	r, err := c.ListIdpGroupMappings(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_ListPolicies(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.ListPoliciesRequest{}
//	r, err := c.ListPolicies(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_ListRegionSubscriptions(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.ListRegionSubscriptionsRequest{}
//	r, err := c.ListRegionSubscriptions(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_ListRegions(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	r, err := c.ListRegions(context.Background())
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_ListSwiftPasswords(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.ListSwiftPasswordsRequest{}
//	r, err := c.ListSwiftPasswords(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_ListUserGroupMemberships(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.ListUserGroupMembershipsRequest{}
//	r, err := c.ListUserGroupMemberships(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_UpdateGroup(t *testing.T) {
//}
//
//func TestIdentityClient_UpdateIdentityProvider(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.UpdateIdentityProviderRequest{}
//	r, err := c.UpdateIdentityProvider(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_UpdateIdpGroupMapping(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.UpdateIdpGroupMappingRequest{}
//	r, err := c.UpdateIdpGroupMapping(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_UpdateSwiftPassword(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.UpdateSwiftPasswordRequest{}
//	r, err := c.UpdateSwiftPassword(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_UpdateUserState(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.UpdateUserStateRequest{}
//	r, err := c.UpdateUserState(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_UploadApiKey(t *testing.T) {
//	c := identity.NewClientForRegion(region)
//	request := identity.UploadApiKeyRequest{}
//	r, err := c.UploadApiKey(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}

func TestBadHost(t *testing.T) {
	client := identity.NewClientForRegion(region)
	client.Host = "badhostname"
	response, err := client.ListRegions(context.Background())
	assert.Nil(t, response.RawResponse)
	assert.Error(t, err)
}
