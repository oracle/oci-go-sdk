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

// Group operations CRUD
func TestIdentityClient_GroupCRUD(t *testing.T) {
	// test should not fail if a previous run failed to clean up
	groupName := getUniqueName("Group_CRUD")
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.CreateGroupRequest{}
	request.CompartmentID = common.String(getTenancyID())
	request.Name = common.String(groupName)
	request.Description = common.String("GoSDK_someGroupDesc")
	r, err := c.CreateGroup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	panicIfError(t, err)

	// if we've successfully created a group during testing, make sure that we delete it
	defer func() {
		//Delete
		rDel := identity.DeleteGroupRequest{GroupID: r.ID}
		err = c.DeleteGroup(context.Background(), rDel)
		assert.NoError(t, err)
	}()

	// validate group lifecycle state enum value after create
	assert.Equal(t, r.Group.LifecycleState, identity.GROUP_LIFECYCLE_STATE_ACTIVE)

	//Get
	rRead := identity.GetGroupRequest{GroupID: r.ID}
	resRead, err := c.GetGroup(context.Background(), rRead)
	assert.NotEmpty(t, r, fmt.Sprint(resRead.ID))
	panicIfError(t, err)

	// validate group lifecycle state enum value after read
	assert.Equal(t, resRead.LifecycleState, identity.GROUP_LIFECYCLE_STATE_ACTIVE)

	//Update
	rUpdate := identity.UpdateGroupRequest{GroupID: r.ID}
	rUpdate.Description = common.String("New description")
	resUpdate, err := c.UpdateGroup(context.Background(), rUpdate)
	panicIfError(t, err)
	assert.NotNil(t, resUpdate.ID)
}

func TestIdentityClient_ListGroups(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.ListGroupsRequest{CompartmentID: common.String(getTenancyID())}
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
	c:= identity.NewIdentityClientForRegion(getRegion())
	request:= identity.CreateCompartmentRequest{CreateCompartmentDetails:identity.CreateCompartmentDetails{
		Name:"egTest2",
		Description:"egTest2_descp",
		CompartmentID:getTenancyID(),
	}}
	r, err:= identity.CreateCompartment(c, request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}
*/

//Comparment RU
func TestIdentityClient_UpdateCompartment(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	//Update
	request := identity.UpdateCompartmentRequest{UpdateCompartmentDetails: identity.UpdateCompartmentDetails{
		Name:        common.String(GoSDK2_Test_Prefix + "UpdComp"),
		Description: common.String("GOSDK2 description2"),
	},
		CompartmentID: common.String(getCompartmentID()),
	}
	r, err := c.UpdateCompartment(context.Background(), request)
	panicIfError(t, err)
	assert.NotEmpty(t, r, fmt.Sprint(r))

	rRead := identity.GetCompartmentRequest{CompartmentID: common.String(getTenancyID())}
	resRead, err := c.GetCompartment(context.Background(), rRead)
	panicIfError(t, err)
	assert.NotEmpty(t, r, fmt.Sprint(resRead))
	return
}

//User Operations
func TestIdentityClient_ListUsers(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.ListUsersRequest{CompartmentID: common.String(getTenancyID())}
	r, err := c.ListUsers(context.Background(), request)
	assert.NotEmpty(t, r.Items, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UserCRUD(t *testing.T) {
	// test should not fail if a previous run failed to clean up
	userName := getUniqueName("User_")
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.CreateUserRequest{}
	request.CompartmentID = common.String(getTenancyID())
	request.Name = common.String(userName)
	request.Description = common.String("Test user for golang sdk2")
	resCreate, err := c.CreateUser(context.Background(), request)
	fmt.Println(resCreate)
	assert.NotEmpty(t, resCreate, fmt.Sprint(resCreate))
	assert.NoError(t, err)

	// if we've successfully created a user during testing, make sure that we delete it
	defer func() {
		//remove
		rDelete := identity.DeleteUserRequest{}
		rDelete.UserID = resCreate.ID
		err = c.DeleteUser(context.Background(), rDelete)
		assert.NoError(t, err)
	}()

	// validate user lifecycle state enum value after read
	assert.Equal(t, resCreate.LifecycleState, identity.USER_LIFECYCLE_STATE_ACTIVE)

	//Read
	rRead := identity.GetUserRequest{UserID: resCreate.ID}
	resRead, err := c.GetUser(context.Background(), rRead)
	assert.NotEmpty(t, resRead, fmt.Sprint(resRead))
	assert.NoError(t, err)

	// validate user lifecycle state enum value after read
	assert.Equal(t, resRead.LifecycleState, identity.USER_LIFECYCLE_STATE_ACTIVE)

	//Update
	rUpdate := identity.UpdateUserRequest{}
	rUpdate.UserID = resCreate.ID
	rUpdate.Description = common.String("This is a new description")
	resUpdate, err := c.UpdateUser(context.Background(), rUpdate)
	assert.NotEmpty(t, resUpdate, fmt.Sprint(resUpdate))
	assert.NoError(t, err)

	return
}

//User-Group operations
func TestIdentityClient_AddUserToGroup(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	//add
	requestAdd := identity.AddUserToGroupRequest{}
	requestAdd.UserID = common.String(getUserID())
	requestAdd.GroupID = common.String(getGroupID())
	r, err := c.AddUserToGroup(context.Background(), requestAdd)
	panicIfError(t, err)
	assert.NotEmpty(t, r, fmt.Sprint(r))

	defer func() {
		//remove
		requestRemove := identity.RemoveUserFromGroupRequest{UserGroupMembershipID: r.UserGroupMembership.ID}
		err = c.RemoveUserFromGroup(context.Background(), requestRemove)
		panicIfError(t, err)
	}()

	// validate user membership lifecycle state enum value after create
	assert.Equal(t, r.LifecycleState, identity.USER_GROUP_MEMBERSHIP_LIFECYCLE_STATE_ACTIVE)

	return
}

//Policy Operations see DEX-1945
//func TestIdentityClient_PolicyCRUD(t *testing.T) {
//	//Create
//	client := identity.NewIdentityClientForRegion(getRegion())
//	/*
//	createRequest := identity.CreatePolicyRequest{}
//	createRequest.CompartmentID = getTenancyID()
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
//	defer func() {
//		// Delete
//		request := identity.DeletePolicyRequest{PolicyID:createResponseID}
//		err = client.DeletePolicy(context.Background(), request)
//		assert.NoError(t, err)
//	}()
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
//
//	return
//}
//

//SecretKey operations
func TestIdentityClient_SecretKeyCRUD(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.CreateCustomerSecretKeyRequest{}
	request.UserID = common.String(getUserID())
	request.DisplayName = common.String("GolangSDK2TestSecretKey")
	resCreate, err := c.CreateCustomerSecretKey(context.Background(), request)
	panicIfError(t, err)
	assert.NotEmpty(t, resCreate, fmt.Sprint(resCreate))
	assert.NoError(t, err)

	defer func() {
		//remove
		rDelete := identity.DeleteCustomerSecretKeyRequest{}
		rDelete.CustomerSecretKeyID = resCreate.ID
		rDelete.UserID = common.String(getUserID())
		err = c.DeleteCustomerSecretKey(context.Background(), rDelete)
		panicIfError(t, err)
	}()

	// validate user membership lifecycle state enum value after create
	assert.Equal(t, resCreate.LifecycleState, identity.CUSTOMER_SECRET_KEY_LIFECYCLE_STATE_ACTIVE)

	//Update
	rUpdate := identity.UpdateCustomerSecretKeyRequest{}
	rUpdate.CustomerSecretKeyID = resCreate.ID
	rUpdate.UserID = common.String(getUserID())
	rUpdate.DisplayName = common.String("This is a new description")
	resUpdate, err := c.UpdateCustomerSecretKey(context.Background(), rUpdate)
	assert.NotEmpty(t, resUpdate, fmt.Sprint(resUpdate))
	panicIfError(t, err)

	return
}

func TestIdentityClient_ListCustomerSecretKeys(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.ListCustomerSecretKeysRequest{}
	request.UserID = common.String(getUserID())
	r, err := c.ListCustomerSecretKeys(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	panicIfError(t, err)
	return
}

//Apikeys
func TestIdentityClient_ApiKeyCRUD(t *testing.T) {
	userID := ""
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.UploadApiKeyRequest{}
	request.UserID = common.String(userID)
	request.Key = common.String("some key")
	resCreate, err := c.UploadApiKey(context.Background(), request)
	assert.Error(t, err)
	ser, ok := common.IsServiceError(err)
	assert.True(t, ok)
	assert.NotEmpty(t, ser.GetMessage())

	defer func() {
		//remove
		rDelete := identity.DeleteApiKeyRequest{}
		rDelete.Fingerprint = resCreate.Fingerprint
		rDelete.UserID = common.String(userID)
		err = c.DeleteApiKey(context.Background(), rDelete)
		ser, ok = common.IsServiceError(err)
		assert.False(t, ok)
		assert.NotEmpty(t, err.Error())
	}()

	// TODO: [2017-Nov-07::shalka] presently LifecycleState isn't being set on ApiKey struct in the Response => merits
	//  further investigation
	// validate api key lifecycle state enum value after create
	// assert.Equal(t, resCreate.LifecycleState, identity.API_KEY_LIFECYCLE_STATE_ACTIVE)

	return
}
func TestIdentityClient_ListApiKeys(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.ListApiKeysRequest{}
	request.UserID = common.String(getUserID())
	r, err := c.ListApiKeys(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	panicIfError(t, err)
	return
}

func TestIdentityClient_IdentityProviderCRUD(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())

	//Create the Identity Provider Request
	rCreate := identity.CreateIdentityProviderRequest{}
	rCreate.CompartmentID = common.String(getCompartmentID())
	rCreate.Name = common.String(getUniqueName("Ident_Provider_"))
	rCreate.Description = common.String("CRUD Test Identity Provider")
	rCreate.ProductType = identity.CREATE_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_ADFS
	rCreate.Protocol = identity.CREATE_IDENTITY_PROVIDER_DETAILS_PROTOCOL_SAML2

	// Create
	rspCreate, err := c.CreateIdentityProvider(context.Background(), rCreate)

	// Validate response
	verifyResponseIsValid(t, rspCreate, err)

	// Verify requested values are correct
	assert.NotEmpty(t, rspCreate.ID)
	assert.NotEmpty(t, rspCreate.OpcRequestID)
	assert.Equal(t, rspCreate.CompartmentID, rCreate.CompartmentID)
	assert.Equal(t, rspCreate.Name, rCreate.Name)
	assert.Equal(t, rspCreate.Description, rCreate.Description)
	assert.Equal(t, rspCreate.ProductType, rCreate.ProductType)
	assert.Equal(t, rspCreate.Protocol, rCreate.Protocol)

	assert.Equal(t, rspCreate.LifecycleState, identity.API_KEY_LIFECYCLE_STATE_ACTIVE)

	defer func() {
		//remove
		rDelete := identity.DeleteIdentityProviderRequest{}
		rDelete.IdentityProviderID = rspCreate.ID
		err := c.DeleteIdentityProvider(context.Background(), rDelete)
		panicIfError(t, err)
	}()

	// Read
	rRead := identity.GetIdentityProviderRequest{}
	rRead.IdentityProviderID = rspCreate.ID
	rspRead, err := c.GetIdentityProvider(context.Background(), rRead)
	verifyResponseIsValid(t, rspRead, err)
	assert.Equal(t, rspRead.ID, rRead.IdentityProviderID)

	// Update
	rUpdate := identity.UpdateIdentityProviderRequest{}
	rUpdate.IdentityProviderID = rspRead.ID
	rUpdate.Protocol = identity.UPDATE_IDENTITY_PROVIDER_DETAILS_PROTOCOL_UNKNOWN
	rUpdate.Description = common.String("CRUD Test Identity Provider UPDATED")
	rspUpdate, err := c.UpdateIdentityProvider(context.Background(), rUpdate)

	verifyResponseIsValid(t, rspUpdate, err)
	assert.Equal(t, rspUpdate.Protocol, rUpdate.Protocol )
	assert.Equal(t, rspUpdate.Description, rUpdate.Description)

	return
}

func TestIdentityClient_ListIdentityProviders(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.ListIdentityProvidersRequest{}
	request.CompartmentID = common.String(getCompartmentID())
	request.Protocol = identity.LIST_IDENTITY_PROVIDERS_PROTOCOL_SAML2
	response, err := c.ListIdentityProviders(context.Background(), request)
	verifyResponseIsValid(t, response, err)
	return
}

// TODO
//func TestIdentityClient_CreateIdpGroupMapping(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.CreateIdpGroupMappingRequest{}
//	r, err := c.CreateIdpGroupMapping(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_CreateOrResetUIPassword(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.CreateOrResetUIPasswordRequest{}
//	r, err := c.CreateOrResetUIPassword(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_CreateRegionSubscription(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.CreateRegionSubscriptionRequest{}
//	r, err := c.CreateRegionSubscription(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_CreateSwiftPassword(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
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

//
//func TestIdentityClient_DeleteIdpGroupMapping(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.DeleteIdpGroupMappingRequest{}
//	err := c.DeleteIdpGroupMapping(context.Background(), request)
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_DeleteSwiftPassword(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.DeleteSwiftPasswordRequest{}
//	err := c.DeleteSwiftPassword(context.Background(), request)
//	assert.NoError(t, err)
//	return
//}
//

//
//func TestIdentityClient_GetIdpGroupMapping(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.GetIdpGroupMappingRequest{}
//	r, err := c.GetIdpGroupMapping(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_GetTenancy(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.GetTenancyRequest{}
//	r, err := c.GetTenancy(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_GetUserGroupMembership(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.GetUserGroupMembershipRequest{}
//	r, err := c.GetUserGroupMembership(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_ListAvailabilityDomains(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.ListAvailabilityDomainsRequest{}
//	r, err := c.ListAvailabilityDomains(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_ListCompartments(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.ListCompartmentsRequest{}
//	r, err := c.ListCompartments(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//

//
//func TestIdentityClient_ListIdpGroupMappings(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.ListIdpGroupMappingsRequest{}
//	r, err := c.ListIdpGroupMappings(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_ListPolicies(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.ListPoliciesRequest{}
//	r, err := c.ListPolicies(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_ListRegionSubscriptions(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.ListRegionSubscriptionsRequest{}
//	r, err := c.ListRegionSubscriptions(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_ListRegions(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	r, err := c.ListRegions(context.Background())
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_ListSwiftPasswords(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.ListSwiftPasswordsRequest{}
//	r, err := c.ListSwiftPasswords(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_ListUserGroupMemberships(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
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

//
//func TestIdentityClient_UpdateIdpGroupMapping(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.UpdateIdpGroupMappingRequest{}
//	r, err := c.UpdateIdpGroupMapping(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_UpdateSwiftPassword(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.UpdateSwiftPasswordRequest{}
//	r, err := c.UpdateSwiftPassword(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_UpdateUserState(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.UpdateUserStateRequest{}
//	r, err := c.UpdateUserState(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}
//
//func TestIdentityClient_UploadApiKey(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.UploadApiKeyRequest{}
//	r, err := c.UploadApiKey(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}

func TestBadHost(t *testing.T) {
	client := identity.NewIdentityClientForRegion(getRegion())
	client.Host = "badhostname"
	response, err := client.ListRegions(context.Background())
	assert.Nil(t, response.RawResponse)
	assert.Error(t, err)
}
