// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/identity"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	//"time"
)

var (
	rootTestCompartmentID = "ocidv1:tenancy:oc1:phx:1460406592660:aaaaaaaab4faofrfkxecohhjuivjq262pu"
	validUserID           = "ocid1.user.oc1..aaaaaaaav6gsclr6pd4yjqengmriylyck55lvon5ujjnhkok5gyxii34lvra"
	validCompartmetnID    = "ocid1.compartment.oc1..aaaaaaaa5dvrjzvfn3rub24nczhih3zb3a673b6tmbvpng3j5apobtxshlma"
	validGroupID          = "ocid1.group.oc1..aaaaaaaaf5lt5afvucqrtzvinpuhfq4rlhmdxqevcdqe7rv6d3vl5ytdqoyq"
)

// Group operations
func TestIdentityClient_CreateGroup(t *testing.T) {
	c := identity.NewClient()
	request := identity.CreateGroupRequest{}
	request.CompartmentID = rootTestCompartmentID
	request.Name = "GoSDK2_someGroup"
	request.Description = "GoSDK_someGroupDesc"
	r, err := c.CreateGroup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_GetGroup(t *testing.T) {
	c := identity.NewClient()
	testGroupOcid := "ocid1.group.oc1..aaaaaaaau5n6f3v7pv2bnqvfgbkxf2dryhbhxpbdshztbfvcyl6ovsgm5gma"
	request := identity.GetGroupRequest{GroupID: testGroupOcid}
	r, err := c.GetGroup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListGroups(t *testing.T) {
	c := identity.NewClient()
	request := identity.ListGroupsRequest{CompartmentID: rootTestCompartmentID}
	r, err := c.ListGroups(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	assert.NotEmpty(t, r.Items)
	assert.NotEmpty(t, r.OpcNextPage)
	return
}

func TestIdentityClient_DeleteGroup(t *testing.T) {
	c := identity.NewClient()
	cr, err := c.CreateGroup(context.Background(), identity.CreateGroupRequest{identity.CreateGroupDetails{
		CompartmentID: rootTestCompartmentID,
		Name:          "GoSDK2TestingVolatile",
		Description:   "GOSDK2TestingVol",
	}, ""})

	if err != nil {
		assert.Fail(t, "Creation of group should have succeeded")
		return
	}
	gid := cr.Group.ID
	assert.NotEmpty(t, gid)

	request := identity.DeleteGroupRequest{GroupID: gid}
	err = c.DeleteGroup(context.Background(), request)
	assert.NoError(t, err)
	return
}

// Compartment operations
//Can not delete compartments right now! Careful!
/*
func TestIdentityClient_CreateCompartment(t *testing.T) {
	c:= identity.NewClient()
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
func TestIdentityClient_UpdateCompartment(t *testing.T) {
	c := identity.NewClient()
	request := identity.UpdateCompartmentRequest{UpdateCompartmentDetails: identity.UpdateCompartmentDetails{
		Name:        "GOSDK2_Test",
		Description: "GOSDK2 description",
	},
		CompartmentID: validCompartmetnID,
	}
	r, err := c.UpdateCompartment(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_GetCompartment(t *testing.T) {
	c := identity.NewClient()
	request := identity.GetCompartmentRequest{CompartmentID: rootTestCompartmentID}
	r, err := c.GetCompartment(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

//User Operations
func TestIdentityClient_ListUsers(t *testing.T) {
	c := identity.NewClient()
	request := identity.ListUsersRequest{CompartmentID: rootTestCompartmentID}
	r, err := c.ListUsers(context.Background(), request)
	assert.NotEmpty(t, r.Items, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UserCRUD(t *testing.T) {
	c := identity.NewClient()
	request := identity.CreateUserRequest{}
	request.CompartmentID = rootTestCompartmentID
	request.Name = "GolangSDK2_testUser"
	request.Description = "Test user for golagn sdk2"
	resCreate, err := c.CreateUser(context.Background(), request)
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
	rUpdate.Description = "This is a new description"
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
	c := identity.NewClient()
	//add
	requestAdd := identity.AddUserToGroupRequest{}
	requestAdd.UserID = validUserID
	requestAdd.GroupID = validGroupID
	r, err := c.AddUserToGroup(context.Background(), requestAdd)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)

	//remove
	requestRemove := identity.RemoveUserFromGroupRequest{UserGroupMembershipID: r.UserGroupMembership.ID}
	err = c.RemoveUserFromGroup(context.Background(), requestRemove)
	assert.NoError(t, err)
	return
}

//Policy Operations see DEX-1945
//func TestIdentityClient_PolicyCRUD(t *testing.T) {
//	//Create
//	client := identity.NewClient()
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

//SecretKey operations
func TestIdentityClient_SecretKeyCRUD(t *testing.T) {
	c := identity.NewClient()
	request := identity.CreateCustomerSecretKeyRequest{}
	request.UserID = validUserID
	request.DisplayName = "GolangSDK2TestSecretKey"
	resCreate, err := c.CreateCustomerSecretKey(context.Background(), request)
	assert.NotEmpty(t, resCreate, fmt.Sprint(resCreate))
	assert.NoError(t, err)

	//Update
	rUpdate := identity.UpdateCustomerSecretKeyRequest{}
	rUpdate.CustomerSecretKeyID = resCreate.ID
	rUpdate.UserID = validUserID
	rUpdate.DisplayName = "This is a new description"
	resUpdate, err := c.UpdateCustomerSecretKey(context.Background(), rUpdate)
	assert.NotEmpty(t, resUpdate, fmt.Sprint(resUpdate))
	assert.NoError(t, err)

	//remove
	rDelete := identity.DeleteCustomerSecretKeyRequest{}
	rDelete.CustomerSecretKeyID = resCreate.ID
	rDelete.UserID = validUserID
	err = c.DeleteCustomerSecretKey(context.Background(), rDelete)
	assert.NoError(t, err)

	return
}

func TestIdentityClient_ListCustomerSecretKeys(t *testing.T) {
	c := identity.NewClient()
	request := identity.ListCustomerSecretKeysRequest{}
	request.UserID = validUserID
	r, err := c.ListCustomerSecretKeys(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

//Apikeys
func TestIdentityClient_ApiKeyCRUD(t *testing.T) {
	userID := ""
	c := identity.NewClient()
	request := identity.UploadApiKeyRequest{}
	request.UserID = userID
	request.Key = "some key"
	resCreate, err := c.UploadApiKey(context.Background(), request)
	assert.NotEmpty(t, resCreate, fmt.Sprint(resCreate))
	assert.NoError(t, err)

	//remove
	rDelete := identity.DeleteApiKeyRequest{}
	rDelete.Fingerprint = resCreate.Fingerprint
	rDelete.UserID = userID
	err = c.DeleteApiKey(context.Background(), rDelete)
	assert.NoError(t, err)

	return
}
func TestIdentityClient_ListApiKeys(t *testing.T) {
	c := identity.NewClient()
	request := identity.ListApiKeysRequest{}
	request.UserID = validUserID
	r, err := c.ListApiKeys(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_CreateIdentityProvider(t *testing.T) {
	c := identity.NewClient()
	request := identity.CreateIdentityProviderRequest{}
	r, err := c.CreateIdentityProvider(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_CreateIdpGroupMapping(t *testing.T) {
	c := identity.NewClient()
	request := identity.CreateIdpGroupMappingRequest{}
	r, err := c.CreateIdpGroupMapping(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_CreateOrResetUIPassword(t *testing.T) {
	c := identity.NewClient()
	request := identity.CreateOrResetUIPasswordRequest{}
	r, err := c.CreateOrResetUIPassword(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_CreateRegionSubscription(t *testing.T) {
	c := identity.NewClient()
	request := identity.CreateRegionSubscriptionRequest{}
	r, err := c.CreateRegionSubscription(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_CreateSwiftPassword(t *testing.T) {
	c := identity.NewClient()
	request := identity.CreateSwiftPasswordRequest{}
	r, err := c.CreateSwiftPassword(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_DeleteApiKey(t *testing.T) {
}

func TestIdentityClient_DeleteIdentityProvider(t *testing.T) {
	c := identity.NewClient()
	request := identity.DeleteIdentityProviderRequest{}
	err := c.DeleteIdentityProvider(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestIdentityClient_DeleteIdpGroupMapping(t *testing.T) {
	c := identity.NewClient()
	request := identity.DeleteIdpGroupMappingRequest{}
	err := c.DeleteIdpGroupMapping(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestIdentityClient_DeleteSwiftPassword(t *testing.T) {
	c := identity.NewClient()
	request := identity.DeleteSwiftPasswordRequest{}
	err := c.DeleteSwiftPassword(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestIdentityClient_GetIdentityProvider(t *testing.T) {
	c := identity.NewClient()
	request := identity.GetIdentityProviderRequest{}
	r, err := c.GetIdentityProvider(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_GetIdpGroupMapping(t *testing.T) {
	c := identity.NewClient()
	request := identity.GetIdpGroupMappingRequest{}
	r, err := c.GetIdpGroupMapping(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_GetTenancy(t *testing.T) {
	c := identity.NewClient()
	request := identity.GetTenancyRequest{}
	r, err := c.GetTenancy(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_GetUserGroupMembership(t *testing.T) {
	c := identity.NewClient()
	request := identity.GetUserGroupMembershipRequest{}
	r, err := c.GetUserGroupMembership(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListAvailabilityDomains(t *testing.T) {
	c := identity.NewClient()
	request := identity.ListAvailabilityDomainsRequest{}
	r, err := c.ListAvailabilityDomains(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListCompartments(t *testing.T) {
	c := identity.NewClient()
	request := identity.ListCompartmentsRequest{}
	r, err := c.ListCompartments(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListIdentityProviders(t *testing.T) {
	c := identity.NewClient()
	request := identity.ListIdentityProvidersRequest{}
	r, err := c.ListIdentityProviders(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListIdpGroupMappings(t *testing.T) {
	c := identity.NewClient()
	request := identity.ListIdpGroupMappingsRequest{}
	r, err := c.ListIdpGroupMappings(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListPolicies(t *testing.T) {
	c := identity.NewClient()
	request := identity.ListPoliciesRequest{}
	r, err := c.ListPolicies(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListRegionSubscriptions(t *testing.T) {
	c := identity.NewClient()
	request := identity.ListRegionSubscriptionsRequest{}
	r, err := c.ListRegionSubscriptions(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListRegions(t *testing.T) {
	c := identity.NewClient()
	r, err := c.ListRegions(context.Background())
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListSwiftPasswords(t *testing.T) {
	c := identity.NewClient()
	request := identity.ListSwiftPasswordsRequest{}
	r, err := c.ListSwiftPasswords(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListUserGroupMemberships(t *testing.T) {
	c := identity.NewClient()
	request := identity.ListUserGroupMembershipsRequest{}
	r, err := c.ListUserGroupMemberships(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}


func TestIdentityClient_UpdateGroup(t *testing.T) {
	c := identity.NewClient()
	request := identity.UpdateGroupRequest{}
	r, err := c.UpdateGroup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UpdateIdentityProvider(t *testing.T) {
	c := identity.NewClient()
	request := identity.UpdateIdentityProviderRequest{}
	r, err := c.UpdateIdentityProvider(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UpdateIdpGroupMapping(t *testing.T) {
	c := identity.NewClient()
	request := identity.UpdateIdpGroupMappingRequest{}
	r, err := c.UpdateIdpGroupMapping(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UpdateSwiftPassword(t *testing.T) {
	c := identity.NewClient()
	request := identity.UpdateSwiftPasswordRequest{}
	r, err := c.UpdateSwiftPassword(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UpdateUserState(t *testing.T) {
	c := identity.NewClient()
	request := identity.UpdateUserStateRequest{}
	r, err := c.UpdateUserState(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UploadApiKey(t *testing.T) {
	c := identity.NewClient()
	request := identity.UploadApiKeyRequest{}
	r, err := c.UploadApiKey(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}
