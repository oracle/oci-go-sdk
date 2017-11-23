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
	failIfError(t, err)

	// if we've successfully created a group during testing, make sure that we delete it
	defer func() {
		//Delete
		rDel := identity.DeleteGroupRequest{GroupID: r.ID}
		err = c.DeleteGroup(context.Background(), rDel)
		assert.NoError(t, err)
	}()

	// validate group lifecycle state enum value after create
	assert.Equal(t, identity.GROUP_LIFECYCLE_STATE_ACTIVE, r.Group.LifecycleState)

	//Get
	rRead := identity.GetGroupRequest{GroupID: r.ID}
	resRead, err := c.GetGroup(context.Background(), rRead)
	assert.NotEmpty(t, r, fmt.Sprint(resRead.ID))
	failIfError(t, err)

	// validate group lifecycle state enum value after read
	assert.Equal(t, identity.GROUP_LIFECYCLE_STATE_ACTIVE, resRead.LifecycleState)

	//Update
	rUpdate := identity.UpdateGroupRequest{GroupID: r.ID}
	rUpdate.Description = common.String("New description")
	resUpdate, err := c.UpdateGroup(context.Background(), rUpdate)
	failIfError(t, err)
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
	failIfError(t, err)
	assert.NotEmpty(t, r, fmt.Sprint(r))

	rRead := identity.GetCompartmentRequest{CompartmentID: common.String(getTenancyID())}
	resRead, err := c.GetCompartment(context.Background(), rRead)
	failIfError(t, err)
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
	assert.Equal(t, identity.USER_LIFECYCLE_STATE_ACTIVE, resCreate.LifecycleState)

	//Read
	rRead := identity.GetUserRequest{UserID: resCreate.ID}
	resRead, err := c.GetUser(context.Background(), rRead)
	assert.NotEmpty(t, resRead, fmt.Sprint(resRead))
	assert.NoError(t, err)

	// validate user lifecycle state enum value after read
	assert.Equal(t, identity.USER_LIFECYCLE_STATE_ACTIVE, resRead.LifecycleState)

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

	// for robustness, create a user and group to use for this test. delete it at the end
	reqAddUser := identity.CreateUserRequest{}
	reqAddUser.CompartmentID = common.String(getTenancyID())
	reqAddUser.Name = common.String(getUniqueName("AUTG_User"))
	reqAddUser.Description = common.String("AddUserToGroup Test User")
	rspAddUser, err1 := c.CreateUser(context.Background(), reqAddUser)

	failIfError(t, err1)

	defer func() {
		// Delete the user
		reqUserDelete := identity.DeleteUserRequest{UserID: rspAddUser.ID}
		delUserErr := c.DeleteUser(context.Background(), reqUserDelete)
		assert.NoError(t, delUserErr)
	}()

	reqAddGroup := identity.CreateGroupRequest{}
	reqAddGroup.CompartmentID = common.String(getTenancyID())
	reqAddGroup.Name = common.String(getUniqueName("AUTG_Group_"))
	reqAddGroup.Description = common.String("AddUserToGroup Test Group")
	rspAddGroup, err2 := c.CreateGroup(context.Background(), reqAddGroup)

	failIfError(t, err2)

	defer func() {
		// Delete the group
		reqGroupDelete := identity.DeleteGroupRequest{GroupID: rspAddGroup.ID}
		delGrpErr := c.DeleteGroup(context.Background(), reqGroupDelete)
		assert.NoError(t, delGrpErr)
	}()

	//add
	reqAdd := identity.AddUserToGroupRequest{}
	reqAdd.UserID = rspAddUser.ID
	reqAdd.GroupID = rspAddGroup.ID
	rspAdd, err := c.AddUserToGroup(context.Background(), reqAdd)
	failIfError(t, err)
	assert.NotEmpty(t, rspAdd, fmt.Sprint(rspAdd))

	defer func() {
		//remove
		requestRemove := identity.RemoveUserFromGroupRequest{UserGroupMembershipID: rspAdd.UserGroupMembership.ID}
		err = c.RemoveUserFromGroup(context.Background(), requestRemove)
		failIfError(t, err)
	}()

	// validate user membership lifecycle state enum value after create
	assert.Equal(t, identity.USER_GROUP_MEMBERSHIP_LIFECYCLE_STATE_ACTIVE, rspAdd.LifecycleState)

	// Read
	reqRead := identity.GetUserGroupMembershipRequest{}
	reqRead.UserGroupMembershipID = rspAdd.ID
	rspRead, readErr := c.GetUserGroupMembership(context.Background(), reqRead)
	verifyResponseIsValid(t, rspRead, readErr)
	assert.Equal(t, rspAdd.ID, rspRead.ID)
	return

}

func TestIdentityClient_ListUserGroupMemberships(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.ListUserGroupMembershipsRequest{}
	request.UserID = common.String(getUserID())
	request.CompartmentID = common.String(getTenancyID())
	r, err := c.ListUserGroupMemberships(context.Background(), request)
	verifyResponseIsValid(t, r, err)
	assert.NotZero(t, len(r.Items))
	return
}

func TestIdentityClient_CreateOrResetUIPassword(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())

	//create the user
	u, err := createTestUser(c)
	failIfError(t, err)
	defer func() {
		failIfError(t, deleteTestUser(c, u.ID))
	}()

	request := identity.CreateOrResetUIPasswordRequest{}
	request.UserID = u.ID
	rspCreate, err := c.CreateOrResetUIPassword(context.Background(), request)
	failIfError(t, err)
	verifyResponseIsValid(t, rspCreate, err)

	assert.NotEmpty(t, rspCreate.OpcRequestID)
	assert.Equal(t, u.ID, rspCreate.UserID)
	assert.NotEmpty(t, rspCreate.Password)
	assert.Equal(t, identity.UI_PASSWORD_LIFECYCLE_STATE_ACTIVE, rspCreate.LifecycleState)

	// make the request again and ensure that we get a different password
	rspReset, err := c.CreateOrResetUIPassword(context.Background(), request)
	failIfError(t, err)
	verifyResponseIsValid(t, rspReset, err)

	assert.Equal(t, rspCreate.UserID, rspReset.UserID)
	assert.NotEqual(t, rspCreate.Password, rspReset.Password)
	assert.Equal(t, identity.UI_PASSWORD_LIFECYCLE_STATE_ACTIVE, rspCreate.LifecycleState)

	return
}

func TestIdentityClient_SwiftPasswordCRUD(t *testing.T) {

	createDesc := "Go SDK Test Swift Password - CREATED"
	updateDesc := "Go SDK Test Swift Password - UPDATED"
	c := identity.NewIdentityClientForRegion(getRegion())

	usr, usrErr := createTestUser(c)
	failIfError(t, usrErr)

	defer deleteTestUser(c, usr.ID)

	// Create Swift Password
	addReq := identity.CreateSwiftPasswordRequest{UserID: usr.ID}
	addReq.Description = &createDesc
	rspPwd, err := c.CreateSwiftPassword(context.Background(), addReq)
	verifyResponseIsValid(t, rspPwd, err)

	//Delete Swift Password
	defer func() {
		delReq := identity.DeleteSwiftPasswordRequest{}
		delReq.UserID = usr.ID
		delReq.SwiftPasswordID = rspPwd.ID
	}()

	assert.NotEmpty(t, rspPwd.ID)
	assert.Equal(t, usr.ID, rspPwd.UserID)
	assert.NotEmpty(t, rspPwd.Password)
	assert.Equal(t, identity.SWIFT_PASSWORD_LIFECYCLE_STATE_ACTIVE, rspPwd.LifecycleState)
	assert.Equal(t, createDesc, *rspPwd.Description)

	// Update Swift Password
	updReq := identity.UpdateSwiftPasswordRequest{UserID: usr.ID}
	updReq.SwiftPasswordID = rspPwd.ID
	updReq.Description = &updateDesc
	updRsp, err := c.UpdateSwiftPassword(context.Background(), updReq)
	verifyResponseIsValid(t, updRsp, err)

	assert.NotEqual(t, rspPwd.Password, updRsp.Password)
	assert.Equal(t, identity.SWIFT_PASSWORD_LIFECYCLE_STATE_ACTIVE, updRsp.LifecycleState)
	assert.Equal(t, updateDesc, *updRsp.Description)

	//assert.NotEmpty(t, updRsp.ExpiresOn)

	return
}

func TestIdentityClient_ListSwiftPasswords(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())

	usr, usrErr := createTestUser(c)
	failIfError(t, usrErr)
	defer deleteTestUser(c, usr.ID)

	pwdReq := identity.CreateSwiftPasswordRequest{UserID: usr.ID}
	pwdReq.Description = common.String("Test Swift Password 1")
	pwdRsp1, err := c.CreateSwiftPassword(context.Background(), pwdReq)
	verifyResponseIsValid(t, pwdRsp1, err)

	pwdReq.Description = common.String("Test Swift Password 2")
	pwdRsp2, err := c.CreateSwiftPassword(context.Background(), pwdReq)
	verifyResponseIsValid(t, pwdRsp2, err)

	request := identity.ListSwiftPasswordsRequest{}
	request.UserID = usr.ID
	r, err := c.ListSwiftPasswords(context.Background(), request)
	verifyResponseIsValid(t, r, err)

	assert.Equal(t, 2, len(r.Items))
	assert.NotEqual(t, r.Items[0].ID, r.Items[1].ID)
	assert.NotEqual(t, r.Items[0].Description, r.Items[1].Description)

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
//func TestIdentityClient_ListPolicies(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.ListPoliciesRequest{}
//	r, err := c.ListPolicies(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
//}

//SecretKey operations
func TestIdentityClient_SecretKeyCRUD(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.CreateCustomerSecretKeyRequest{}
	request.UserID = common.String(getUserID())
	request.DisplayName = common.String("GolangSDK2TestSecretKey")
	resCreate, err := c.CreateCustomerSecretKey(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resCreate, fmt.Sprint(resCreate))
	assert.NoError(t, err)

	defer func() {
		//remove
		rDelete := identity.DeleteCustomerSecretKeyRequest{}
		rDelete.CustomerSecretKeyID = resCreate.ID
		rDelete.UserID = common.String(getUserID())
		err = c.DeleteCustomerSecretKey(context.Background(), rDelete)
		failIfError(t, err)
	}()

	// validate user membership lifecycle state enum value after create
	assert.Equal(t, identity.CUSTOMER_SECRET_KEY_LIFECYCLE_STATE_ACTIVE, resCreate.LifecycleState)

	//Update
	rUpdate := identity.UpdateCustomerSecretKeyRequest{}
	rUpdate.CustomerSecretKeyID = resCreate.ID
	rUpdate.UserID = common.String(getUserID())
	rUpdate.DisplayName = common.String("This is a new description")
	resUpdate, err := c.UpdateCustomerSecretKey(context.Background(), rUpdate)
	assert.NotEmpty(t, resUpdate, fmt.Sprint(resUpdate))
	failIfError(t, err)

	return
}

func TestIdentityClient_ListCustomerSecretKeys(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.ListCustomerSecretKeysRequest{}
	request.UserID = common.String(getUserID())
	r, err := c.ListCustomerSecretKeys(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	failIfError(t, err)
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
	// assert.Equal(t, identity.API_KEY_LIFECYCLE_STATE_ACTIVE, resCreate.LifecycleState)

	return
}
func TestIdentityClient_ListApiKeys(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.ListApiKeysRequest{}
	request.UserID = common.String(getUserID())
	r, err := c.ListApiKeys(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	failIfError(t, err)
	return
}

func TestIdentityClient_IdentityProviderCRUD(t *testing.T) {
	t.Skip("Blocked by DEX-2116")
	c := identity.NewIdentityClientForRegion(getRegion())

	//Create the Identity Provider Request
	rCreate := identity.CreateIdentityProviderRequest{}
	details := identity.CreateSaml2IdentityProviderDetails{}
	details.CompartmentID = common.String(getTenancyID())
	details.Name = common.String(getUniqueName("Ident_Provider_"))
	details.Description = common.String("CRUD Test Identity Provider")
	details.ProductType = identity.CREATE_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_ADFS
	details.Metadata = common.String(readSampleFederationMetadata(t))
	rCreate.CreateIdentityProviderDetails = details

	// Create
	rspCreate, err := c.CreateIdentityProvider(context.Background(), rCreate)
	failIfError(t, err)
	verifyResponseIsValid(t, rspCreate, err)

	// Verify requested values are correct
	assert.NotEmpty(t, rspCreate.GetID())
	assert.NotEmpty(t, rspCreate.OpcRequestID)
	assert.Equal(t, *rCreate.GetCompartmentID(), *rspCreate.GetCompartmentID())
	assert.NotEmpty(t, *rspCreate.GetName())

	defer func() {
		//remove
		rDelete := identity.DeleteIdentityProviderRequest{}
		rDelete.IdentityProviderID = rspCreate.GetID()
		err := c.DeleteIdentityProvider(context.Background(), rDelete)
		failIfError(t, err)
	}()

	// Read
	rRead := identity.GetIdentityProviderRequest{}
	rRead.IdentityProviderID = rspCreate.GetID()
	rspRead, err := c.GetIdentityProvider(context.Background(), rRead)
	failIfError(t, err)
	verifyResponseIsValid(t, rspRead, err)
	assert.Equal(t, *rRead.IdentityProviderID, *rspRead.GetID())

	// Update
	rUpdate := identity.UpdateIdentityProviderRequest{}
	updateDetails := identity.UpdateSaml2IdentityProviderDetails{}
	updateDetails.Description = common.String("New description")
	rUpdate.IdentityProviderID = rspCreate.GetID()
	rUpdate.UpdateIdentityProviderDetails = updateDetails
	rspUpdate, err := c.UpdateIdentityProvider(context.Background(), rUpdate)

	failIfError(t, err)
	verifyResponseIsValid(t, rspUpdate, err)
	assert.Equal(t, *rspCreate.GetID(), *rspUpdate.GetID())
	assert.Equal(t, "New description", *rspUpdate.GetDescription())

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

func TestIdentityClient_GetTenancy(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.GetTenancyRequest{TenancyID: common.String(getTenancyID())}
	r, err := c.GetTenancy(context.Background(), request)
	verifyResponseIsValid(t, r, err)

	assert.Equal(t, request.TenancyID, r.ID)
	assert.NotEmpty(t, r.OpcRequestID)

	return
}

func TestIdentityClient_ListAvailabilityDomains(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.ListAvailabilityDomainsRequest{CompartmentID: common.String(getCompartmentID())}
	r, err := c.ListAvailabilityDomains(context.Background(), request)
	verifyResponseIsValid(t, r, err)

	assert.NotEmpty(t, r.OpcRequestID)
	assert.NotZero(t, len(r.Items))
	return
}

func TestIdentityClient_ListCompartments(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.ListCompartmentsRequest{CompartmentID: common.String(getTenancyID())}
	r, err := c.ListCompartments(context.Background(), request)
	verifyResponseIsValid(t, r, err)

	assert.NotEmpty(t, r.OpcRequestID)
	assert.NotZero(t, len(r.Items))
	return
}

func TestIdentityClient_ListRegions(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	r, err := c.ListRegions(context.Background())
	verifyResponseIsValid(t, r, err)
	assert.NotZero(t, len(r.Items))
	return
}

func TestIdentityClient_UpdateUserState(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	usr, usrErr := createTestUser(c)
	failIfError(t, usrErr)
	defer deleteTestUser(c, usr.ID)

	request := identity.UpdateUserStateRequest{}
	request.UserID = usr.ID
	request.Blocked = common.Bool(true)

	r, err := c.UpdateUserState(context.Background(), request)
	verifyResponseIsValid(t, r, err)

	assert.Equal(t, 2, r.InactiveStatus)

	request.Blocked = common.Bool(false)
	r, err = c.UpdateUserState(context.Background(), request)
	verifyResponseIsValid(t, r, err)

	assert.Equal(t, identity.USER_LIFECYCLE_STATE_INACTIVE, r.LifecycleState)
	assert.Equal(t, 0, r.InactiveStatus)
	return
}

// This test can only realistically be run once since once a region is subscribed to, there is no
// mechanism to unsubscribe to it. For now we will skip
func TestIdentityClient_CreateRegionSubscription(t *testing.T) {
	t.Skip("SKIPPING: Region Subscriptions cannot be undone.")
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.CreateRegionSubscriptionRequest{}
	request.TenancyID = common.String(getTenancyID())
	request.RegionKey = common.String("FRA")
	r, err := c.CreateRegionSubscription(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListRegionSubscriptions(t *testing.T) {
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.ListRegionSubscriptionsRequest{TenancyID: common.String(getTenancyID())}
	r, err := c.ListRegionSubscriptions(context.Background(), request)
	failIfError(t, err)
	verifyResponseIsValid(t, r, err)
	assert.NotZero(t, len(r.Items))
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
//
//func TestIdentityClient_DeleteIdpGroupMapping(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.DeleteIdpGroupMappingRequest{}
//	err := c.DeleteIdpGroupMapping(context.Background(), request)
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
//func TestIdentityClient_ListIdpGroupMappings(t *testing.T) {
//	c := identity.NewIdentityClientForRegion(getRegion())
//	request := identity.ListIdpGroupMappingsRequest{}
//	r, err := c.ListIdpGroupMappings(context.Background(), request)
//	assert.NotEmpty(t, r, fmt.Sprint(r))
//	assert.NoError(t, err)
//	return
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

func TestBadHost(t *testing.T) {
	client := identity.NewIdentityClientForRegion(getRegion())
	client.Host = "badhostname"
	response, err := client.ListRegions(context.Background())
	assert.Nil(t, response.RawResponse)
	assert.Error(t, err)
}
