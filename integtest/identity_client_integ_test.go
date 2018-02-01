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

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/identity"
	"github.com/stretchr/testify/assert"
	"time"
	"reflect"
)

// Group operations CRUD
func TestIdentityClient_GroupCRUD(t *testing.T) {
	// test should not fail if a previous run failed to clean up
	groupName := getUniqueName("Group_CRUD")
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := identity.CreateGroupRequest{}
	request.CompartmentId = common.String(getTenancyID())
	request.Name = common.String(groupName)
	request.Description = common.String("GoSDK_someGroupDesc")
	r, err := c.CreateGroup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	failIfError(t, err)

	// if we've successfully created a group during testing, make sure that we delete it
	defer func() {
		//Delete
		rDel := identity.DeleteGroupRequest{GroupId: r.Id}
		err = c.DeleteGroup(context.Background(), rDel)
		assert.NoError(t, err)
	}()

	// validate group lifecycle state enum value after create
	assert.Equal(t, identity.GroupLifecycleStateActive, r.Group.LifecycleState)

	//Get
	rRead := identity.GetGroupRequest{GroupId: r.Id}
	resRead, err := c.GetGroup(context.Background(), rRead)
	assert.NotEmpty(t, r, fmt.Sprint(resRead.Id))
	failIfError(t, err)

	// validate group lifecycle state enum value after read
	assert.Equal(t, identity.GroupLifecycleStateActive, resRead.LifecycleState)

	//Update
	rUpdate := identity.UpdateGroupRequest{GroupId: r.Id}
	rUpdate.Description = common.String("New description")
	resUpdate, err := c.UpdateGroup(context.Background(), rUpdate)
	failIfError(t, err)
	assert.NotNil(t, resUpdate.Id)
}

func TestIdentityClient_ListGroups(t *testing.T) {

	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := identity.ListGroupsRequest{CompartmentId: common.String(getTenancyID())}
	r, err := c.ListGroups(context.Background(), request)
	failIfError(t, err)

	items := r.Items

	nextRequest := identity.ListGroupsRequest{CompartmentId: request.CompartmentId}
	nextRequest.Page = r.OpcNextPage

	for nextRequest.Page != nil {
		if r, err = c.ListGroups(context.Background(), nextRequest); err == nil {
			items = append(items, r.Items...)
			nextRequest.Page = r.OpcNextPage
		} else {
			failIfError(t, err)
			break
		}
	}

	assert.NotEmpty(t, items)

	return
}

// Compartment operations
//Can not delete compartments right now! Careful!

func TestIdentityClient_CreateCompartment(t *testing.T) {
	t.Skip("Compartment Creation cannot be undone. Remove Skip and manual execute test to verify")
	c, cfgErr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, cfgErr)

	request:= identity.CreateCompartmentRequest{CreateCompartmentDetails:identity.CreateCompartmentDetails{
		Name: common.String("Compartment_Test"),
		Description: common.String("Go SDK Comparment Test"),
		CompartmentId: common.String(getTenancyID()),
	}}

	r, err:= c.CreateCompartment(context.Background(), request)
	verifyResponseIsValid(t, r, err)
	assert.NotEmpty(t, r.Id)
	assert.Equal(t, request.Name, r.Name)
	assert.Equal(t, request.Description, r.Description)
	return
}


//Comparment RU
func TestIdentityClient_UpdateCompartment(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	//Update
	request := identity.UpdateCompartmentRequest{UpdateCompartmentDetails: identity.UpdateCompartmentDetails{
		Name:        common.String(GoSDK2_Test_Prefix + "UpdComp"),
		Description: common.String("GOSDK2 description2"),
	},
		CompartmentId: common.String(getCompartmentID()),
	}
	r, err := c.UpdateCompartment(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r, fmt.Sprint(r))

	rRead := identity.GetCompartmentRequest{CompartmentId: common.String(getTenancyID())}
	resRead, err := c.GetCompartment(context.Background(), rRead)
	failIfError(t, err)
	assert.NotEmpty(t, r, fmt.Sprint(resRead))
	return
}

//User Operations
func TestIdentityClient_ListUsers(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := identity.ListUsersRequest{CompartmentId: common.String(getTenancyID())}
	r, err := c.ListUsers(context.Background(), request)
	failIfError(t, err)

	items := r.Items

	nextRequest := identity.ListUsersRequest{CompartmentId: request.CompartmentId}
	nextRequest.Page = r.OpcNextPage
	for nextRequest.Page != nil {
		if r, err = c.ListUsers(context.Background(), nextRequest); err == nil {
			items = append(items, r.Items...)
			nextRequest.Page = r.OpcNextPage
		} else {
			failIfError(t, err)
			break
		}
	}

	// Add verification for items array
	assert.NotEmpty(t, items)

	return
}

func TestIdentityClient_UserCRUD(t *testing.T) {
	// test should not fail if a previous run failed to clean up
	userName := getUniqueName("User_")
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := identity.CreateUserRequest{}
	request.CompartmentId = common.String(getTenancyID())
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
		rDelete.UserId = resCreate.Id
		err = c.DeleteUser(context.Background(), rDelete)
		assert.NoError(t, err)
	}()

	// validate user lifecycle state enum value after read
	assert.Equal(t, identity.UserLifecycleStateActive, resCreate.LifecycleState)

	//Read
	rRead := identity.GetUserRequest{UserId: resCreate.Id}
	resRead, err := c.GetUser(context.Background(), rRead)
	assert.NotEmpty(t, resRead, fmt.Sprint(resRead))
	assert.NoError(t, err)

	// validate user lifecycle state enum value after read
	assert.Equal(t, identity.UserLifecycleStateActive, resRead.LifecycleState)

	//Update
	rUpdate := identity.UpdateUserRequest{}
	rUpdate.UserId = resCreate.Id
	rUpdate.Description = common.String("This is a new description")
	resUpdate, err := c.UpdateUser(context.Background(), rUpdate)
	assert.NotEmpty(t, resUpdate, fmt.Sprint(resUpdate))
	assert.NoError(t, err)

	return
}

//User-Group operations
func TestIdentityClient_AddUserToGroup(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	// for robustness, create a user and group to use for this test. delete it at the end
	reqAddUser := identity.CreateUserRequest{}
	reqAddUser.CompartmentId = common.String(getTenancyID())
	reqAddUser.Name = common.String(getUniqueName("AUTG_User"))
	reqAddUser.Description = common.String("AddUserToGroup Test User")
	rspAddUser, err1 := c.CreateUser(context.Background(), reqAddUser)

	failIfError(t, err1)

	defer func() {
		// Delete the user
		reqUserDelete := identity.DeleteUserRequest{UserId: rspAddUser.Id}
		delUserErr := c.DeleteUser(context.Background(), reqUserDelete)
		assert.NoError(t, delUserErr)
	}()

	reqAddGroup := identity.CreateGroupRequest{}
	reqAddGroup.CompartmentId = common.String(getTenancyID())
	reqAddGroup.Name = common.String(getUniqueName("AUTG_Group_"))
	reqAddGroup.Description = common.String("AddUserToGroup Test Group")
	rspAddGroup, err2 := c.CreateGroup(context.Background(), reqAddGroup)

	failIfError(t, err2)

	defer func() {
		// Delete the group
		reqGroupDelete := identity.DeleteGroupRequest{GroupId: rspAddGroup.Id}
		delGrpErr := c.DeleteGroup(context.Background(), reqGroupDelete)
		assert.NoError(t, delGrpErr)
	}()

	//add
	reqAdd := identity.AddUserToGroupRequest{}
	reqAdd.UserId = rspAddUser.Id
	reqAdd.GroupId = rspAddGroup.Id
	rspAdd, err := c.AddUserToGroup(context.Background(), reqAdd)
	failIfError(t, err)
	assert.NotEmpty(t, rspAdd, fmt.Sprint(rspAdd))

	defer func() {
		//remove
		requestRemove := identity.RemoveUserFromGroupRequest{UserGroupMembershipId: rspAdd.UserGroupMembership.Id}
		err = c.RemoveUserFromGroup(context.Background(), requestRemove)
		failIfError(t, err)
	}()

	// validate user membership lifecycle state enum value after create
	assert.Equal(t, identity.UserGroupMembershipLifecycleStateActive, rspAdd.LifecycleState)

	// Read
	reqRead := identity.GetUserGroupMembershipRequest{}
	reqRead.UserGroupMembershipId = rspAdd.Id
	rspRead, readErr := c.GetUserGroupMembership(context.Background(), reqRead)
	verifyResponseIsValid(t, rspRead, readErr)
	assert.Equal(t, rspAdd.Id, rspRead.Id)
	return

}

func TestIdentityClient_ListUserGroupMemberships(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := identity.ListUserGroupMembershipsRequest{}
	request.UserId = common.String(getUserID())
	request.CompartmentId = common.String(getTenancyID())
	r, err := c.ListUserGroupMemberships(context.Background(), request)
	failIfError(t, err)

	items := r.Items

	nextRequest := identity.ListUserGroupMembershipsRequest{CompartmentId: request.CompartmentId, UserId: request.UserId}
	nextRequest.Page = r.OpcNextPage

	for nextRequest.Page != nil {
		if r, err := c.ListUserGroupMemberships(context.Background(), nextRequest); err == nil {
			items = append(items, r.Items...)
			nextRequest.Page = r.OpcNextPage
		} else {
			failIfError(t, err)
			break
		}
	}

	assert.NotEmpty(t, items)

	return
}

func TestIdentityClient_CreateOrResetUIPassword(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	//create the user
	u, err := createTestUser(c)
	failIfError(t, err)
	defer func() {
		failIfError(t, deleteTestUser(c, u.Id))
	}()

	request := identity.CreateOrResetUIPasswordRequest{}
	request.UserId = u.Id
	rspCreate, err := c.CreateOrResetUIPassword(context.Background(), request)
	failIfError(t, err)
	verifyResponseIsValid(t, rspCreate, err)

	assert.NotEmpty(t, rspCreate.OpcRequestId)
	assert.Equal(t, u.Id, rspCreate.UserId)
	assert.NotEmpty(t, rspCreate.Password)
	assert.Equal(t, identity.UiPasswordLifecycleStateActive, rspCreate.LifecycleState)

	// make the request again and ensure that we get a different password
	rspReset, err := c.CreateOrResetUIPassword(context.Background(), request)
	failIfError(t, err)
	verifyResponseIsValid(t, rspReset, err)

	assert.Equal(t, rspCreate.UserId, rspReset.UserId)
	assert.NotEqual(t, rspCreate.Password, rspReset.Password)
	assert.Equal(t, identity.UiPasswordLifecycleStateActive, rspCreate.LifecycleState)

	return
}

func TestIdentityClient_SwiftPasswordCRUD(t *testing.T) {

	createDesc := "Go SDK Test Swift Password - CREATED"
	updateDesc := "Go SDK Test Swift Password - UPDATED"
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	usr, usrErr := createTestUser(c)
	failIfError(t, usrErr)

	defer deleteTestUser(c, usr.Id)

	// Create Swift Password
	addReq := identity.CreateSwiftPasswordRequest{UserId: usr.Id}
	addReq.Description = &createDesc
	rspPwd, err := c.CreateSwiftPassword(context.Background(), addReq)
	verifyResponseIsValid(t, rspPwd, err)

	//Delete Swift Password
	defer func() {
		delReq := identity.DeleteSwiftPasswordRequest{}
		delReq.UserId = usr.Id
		delReq.SwiftPasswordId = rspPwd.Id
	}()

	assert.NotEmpty(t, rspPwd.Id)
	assert.Equal(t, usr.Id, rspPwd.UserId)
	assert.NotEmpty(t, rspPwd.Password)
	assert.Equal(t, identity.SwiftPasswordLifecycleStateActive, rspPwd.LifecycleState)
	assert.Equal(t, createDesc, *rspPwd.Description)

	// Update Swift Password
	updReq := identity.UpdateSwiftPasswordRequest{UserId: usr.Id}
	updReq.SwiftPasswordId = rspPwd.Id
	updReq.Description = &updateDesc
	updRsp, err := c.UpdateSwiftPassword(context.Background(), updReq)
	verifyResponseIsValid(t, updRsp, err)

	assert.NotEqual(t, rspPwd.Password, updRsp.Password)
	assert.Equal(t, identity.SwiftPasswordLifecycleStateActive, updRsp.LifecycleState)
	assert.Equal(t, updateDesc, *updRsp.Description)

	//assert.NotEmpty(t, updRsp.ExpiresOn)

	return
}

func TestIdentityClient_ListSwiftPasswords(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	usr, usrErr := createTestUser(c)
	failIfError(t, usrErr)
	defer deleteTestUser(c, usr.Id)

	pwdReq := identity.CreateSwiftPasswordRequest{UserId: usr.Id}
	pwdReq.Description = common.String("Test Swift Password 1")
	pwdRsp1, err := c.CreateSwiftPassword(context.Background(), pwdReq)
	verifyResponseIsValid(t, pwdRsp1, err)

	pwdReq.Description = common.String("Test Swift Password 2")
	pwdRsp2, err := c.CreateSwiftPassword(context.Background(), pwdReq)
	verifyResponseIsValid(t, pwdRsp2, err)

	request := identity.ListSwiftPasswordsRequest{UserId: usr.Id}
	r, err := c.ListSwiftPasswords(context.Background(), request)
	verifyResponseIsValid(t, r, err)

	assert.Equal(t, 2, len(r.Items))
	assert.NotEqual(t, r.Items[0].Id, r.Items[1].Id)
	assert.NotEqual(t, r.Items[0].Description, r.Items[1].Description)

	return
}

//Policy Operations see DEX-1945
func TestIdentityClient_PolicyCRUD(t *testing.T) {
	t.Skip("Policy Operations issue. See DEX-1945 ")

	//Create
	//client := identity.NewIdentityClientForRegion(getRegion())
	client, cfgErr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, cfgErr)

	createRequest := identity.CreatePolicyRequest{}
	createRequest.CompartmentId = common.String(getTenancyID())
	createRequest.Name = common.String("goSDK2Policy2")
	createRequest.Description = common.String("some policy")
	createRequest.Statements = []string{"Allow group goSDK2CreateGroup read all-resources on compartment egineztest"}
	createRequest.VersionDate = &common.SDKTime{time.Now()}
	createResponse, err := client.CreatePolicy(context.Background(), createRequest)
	verifyResponseIsValid(t, createResponse, err)

	defer func() {
		// Delete
		request := identity.DeletePolicyRequest{PolicyId:createResponse.Id}
		err = client.DeletePolicy(context.Background(), request)
		assert.NoError(t, err)
	}()

	//Read
	readRequest := identity.GetPolicyRequest{}
	readRequest.PolicyId = createResponse.Id
	readResponse, err := client.GetPolicy(context.Background(), readRequest)
	verifyResponseIsValid(t, readResponse, err)

	//Update

	updateRequest := identity.UpdatePolicyRequest{}
	updateRequest.PolicyId = createResponse.Id
	updateRequest.Description = common.String("new description")
	updateResponse, err := client.UpdatePolicy(context.Background(), updateRequest)
	assert.NotEmpty(t, updateResponse, fmt.Sprint(updateResponse))
	assert.NoError(t, err)

	return
}

func TestIdentityClient_ListPolicies(t *testing.T) {
	c, cfgErr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, cfgErr)
	listRequest := identity.ListPoliciesRequest{}
	listRequest.CompartmentId = common.String(getTenancyID())
	listResponse, err := c.ListPolicies(context.Background(), listRequest)
	failIfError(t, err)

	items := listResponse.Items

	nextRequest := identity.ListPoliciesRequest{CompartmentId: listRequest.CompartmentId}
	nextRequest.Page = listResponse.OpcNextPage

	for nextRequest.Page != nil {
		if r, err := c.ListPolicies(context.Background(), nextRequest); err == nil {
			items = append(items, r.Items...)
			nextRequest.Page = r.OpcNextPage
		} else {
			failIfError(t, err)
		}
	}

	assert.NotEmpty(t, items)

	return
}


//SecretKey operations
func TestIdentityClient_SecretKeyCRUD(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := identity.CreateCustomerSecretKeyRequest{}
	request.UserId = common.String(getUserID())
	request.DisplayName = common.String("GolangSDK2TestSecretKey")
	resCreate, err := c.CreateCustomerSecretKey(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resCreate, fmt.Sprint(resCreate))
	assert.NoError(t, err)

	defer func() {
		//remove
		rDelete := identity.DeleteCustomerSecretKeyRequest{}
		rDelete.CustomerSecretKeyId = resCreate.Id
		rDelete.UserId = common.String(getUserID())
		err = c.DeleteCustomerSecretKey(context.Background(), rDelete)
		failIfError(t, err)
	}()

	// validate user membership lifecycle state enum value after create
	assert.Equal(t, identity.CustomerSecretKeyLifecycleStateActive, resCreate.LifecycleState)

	//Update
	rUpdate := identity.UpdateCustomerSecretKeyRequest{}
	rUpdate.CustomerSecretKeyId = resCreate.Id
	rUpdate.UserId = common.String(getUserID())
	rUpdate.DisplayName = common.String("This is a new description")
	resUpdate, err := c.UpdateCustomerSecretKey(context.Background(), rUpdate)
	assert.NotEmpty(t, resUpdate, fmt.Sprint(resUpdate))
	failIfError(t, err)

	return
}

func TestIdentityClient_ListCustomerSecretKeys(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := identity.ListCustomerSecretKeysRequest{}
	request.UserId = common.String(getUserID())
	r, err := c.ListCustomerSecretKeys(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	failIfError(t, err)
	return
}

//Apikeys
func TestIdentityClient_ApiKeyCRUD(t *testing.T) {
	userId := ""
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := identity.UploadApiKeyRequest{}
	request.UserId = common.String(userId)
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
		rDelete.UserId = common.String(userId)
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
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := identity.ListApiKeysRequest{}
	request.UserId = common.String(getUserID())
	r, err := c.ListApiKeys(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	failIfError(t, err)
	return
}

func TestIdentityClient_IdentityProviderCRUD(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	// Create the Identity Provider Request
	rCreate := identity.CreateIdentityProviderRequest{}
	details := identity.CreateSaml2IdentityProviderDetails{}
	details.CompartmentId = common.String(getTenancyID())
	details.Name = common.String(getUniqueName("Ident_Provider_"))
	details.Description = common.String("CRUD Test Identity Provider")
	details.ProductType = identity.CreateIdentityProviderDetailsProductTypeAdfs
	details.Metadata = common.String(readSampleFederationMetadata(t))
	rCreate.CreateIdentityProviderDetails = details

	// Create
	rspCreate, createErr := c.CreateIdentityProvider(context.Background(), rCreate)

	failIfError(t, createErr)
	verifyResponseIsValid(t, rspCreate, createErr)

	defer func() {
		//remove
		fmt.Println("Deleting Identity Provider")
		if rspCreate.GetId() != nil {
			rDelete := identity.DeleteIdentityProviderRequest{}
			rDelete.IdentityProviderId = rspCreate.GetId()
			err := c.DeleteIdentityProvider(context.Background(), rDelete)
			failIfError(t, err)
		}
	}()

	// Verify requested values are correct
	assert.NotEmpty(t, rspCreate.GetId())
	assert.NotEmpty(t, rspCreate.OpcRequestId)
	assert.Equal(t, *rCreate.GetCompartmentId(), *rspCreate.GetCompartmentId())
	assert.NotEmpty(t, *rspCreate.GetName())

	// Read
	rRead := identity.GetIdentityProviderRequest{}
	rRead.IdentityProviderId = rspCreate.GetId()
	rspRead, readErr := c.GetIdentityProvider(context.Background(), rRead)
	failIfError(t, readErr)
	verifyResponseIsValid(t, rspRead, readErr)
	assert.Equal(t, *rRead.IdentityProviderId, *rspRead.GetId())

	// Update
	rUpdate := identity.UpdateIdentityProviderRequest{}
	updateDetails := identity.UpdateSaml2IdentityProviderDetails{}
	updateDetails.Description = common.String("New description")
	rUpdate.IdentityProviderId = rspCreate.GetId()
	rUpdate.UpdateIdentityProviderDetails = updateDetails
	rspUpdate, updErr := c.UpdateIdentityProvider(context.Background(), rUpdate)

	failIfError(t, updErr)
	verifyResponseIsValid(t, rspUpdate, updErr)
	assert.Equal(t, *rspCreate.GetId(), *rspUpdate.GetId())
	assert.Equal(t, "New description", *rspUpdate.GetDescription())

	// Create the group mapping
	u, uErr := createTestUser(c)
	failIfError(t, uErr)
	defer deleteTestUser(c, u.Id)

	g, gErr := createTestGroup(c)
	failIfError(t, gErr)
	defer deleteTestGroup(c, g.Id)

	reqCreateMapping := identity.CreateIdpGroupMappingRequest{}
	reqCreateMapping.IdentityProviderId = rspCreate.GetId()
	reqCreateMapping.GroupId = g.Id
	idpGrpName := *rspCreate.GetName()
	groupName := *g.Name
	reqCreateMapping.IdpGroupName = common.String(idpGrpName + "_TO_" + groupName)

	rspCreateMapping, createMapErr := c.CreateIdpGroupMapping(context.Background(), reqCreateMapping)
	failIfError(t, createMapErr)

	// Delete mapping
	defer func() {
		fmt.Println("Deleting Identity Provider Group Mapping")
		reqDelete := identity.DeleteIdpGroupMappingRequest{MappingId: rspCreateMapping.Id, IdentityProviderId: rspCreateMapping.IdpId}
		delErr := c.DeleteIdpGroupMapping(context.Background(), reqDelete)
		failIfError(t, delErr)
	}()

	verifyResponseIsValid(t, rspCreateMapping, createMapErr)

	assert.NotEmpty(t, *rspCreateMapping.Id)
	assert.NotEmpty(t, *rspCreateMapping.GroupId)
	assert.NotEmpty(t, *rspCreateMapping.OpcRequestId)
	assert.NotEmpty(t, *rspCreateMapping.IdpGroupName)
	assert.Equal(t, *rspCreate.GetId(), *rspCreateMapping.IdpId)
	assert.NotEmpty(t, rspCreateMapping.TimeCreated)
	assert.Equal(t, identity.IdpGroupMappingLifecycleStateActive, rspCreateMapping.LifecycleState)

	//Read group mapping
	reqReadMapping := identity.GetIdpGroupMappingRequest{IdentityProviderId: rspCreateMapping.IdpId, MappingId: rspCreateMapping.Id}
	rspReadMapping, readMapErr := c.GetIdpGroupMapping(context.Background(), reqReadMapping)
	verifyResponseIsValid(t, rspReadMapping, readMapErr)

	assert.Equal(t, rspCreateMapping.Id, rspReadMapping.Id)
	assert.Equal(t, rspCreateMapping.IdpId, rspReadMapping.IdpId)
	assert.Equal(t, identity.IdpGroupMappingLifecycleStateActive, rspReadMapping.LifecycleState)

	//update group mapping
	reqUpdMapping := identity.UpdateIdpGroupMappingRequest{}
	reqUpdMapping.MappingId = rspReadMapping.Id
	reqUpdMapping.IdentityProviderId = rspReadMapping.IdpId
	reqUpdMapping.GroupId = rspReadMapping.GroupId
	updatedName := *rspReadMapping.IdpGroupName + " - Updated"
	reqUpdMapping.IdpGroupName = common.String(updatedName)
	rspUpdMapping, updMapErr := c.UpdateIdpGroupMapping(context.Background(), reqUpdMapping)
	verifyResponseIsValid(t, rspUpdMapping, updMapErr)

	assert.NotEmpty(t, *rspUpdMapping.Id)
	assert.NotEmpty(t, *rspUpdMapping.GroupId)
	assert.NotEmpty(t, *rspUpdMapping.OpcRequestId)
	assert.Equal(t, *rspReadMapping.Id, *rspUpdMapping.Id)
	assert.Equal(t, *rspReadMapping.IdpId, *rspUpdMapping.IdpId)
	assert.NotEqual(t, *rspReadMapping.IdpGroupName, *rspUpdMapping.IdpGroupName)
	assert.NotEmpty(t, rspUpdMapping.TimeCreated)
	assert.Equal(t, identity.IdpGroupMappingLifecycleStateActive, rspUpdMapping.LifecycleState)

	return
}

func TestIdentityClient_ListIdentityProviders(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	//Create the Identity Provider Request
	rCreate := identity.CreateIdentityProviderRequest{}
	details := identity.CreateSaml2IdentityProviderDetails{}
	details.CompartmentId = common.String(getTenancyID())
	details.Name = common.String(getUniqueName("Ident_Provider_"))
	details.Description = common.String("CRUD Test Identity Provider")
	details.ProductType = identity.CreateIdentityProviderDetailsProductTypeIdcs
	details.Metadata = common.String(readSampleFederationMetadata(t))
	rCreate.CreateIdentityProviderDetails = details

	//Create
	rspCreate, createErr := c.CreateIdentityProvider(context.Background(), rCreate)
	failIfError(t, createErr)
	verifyResponseIsValid(t, rspCreate, createErr)

	//Delete
	deleteFn := func() {
		if rspCreate.GetId() != nil {
			rDelete := identity.DeleteIdentityProviderRequest{}
			rDelete.IdentityProviderId = rspCreate.GetId()
			err := c.DeleteIdentityProvider(context.Background(), rDelete)
			failIfError(t, err)
		}
	}
	defer deleteFn()

	rRead := identity.GetIdentityProviderRequest{}
	rRead.IdentityProviderId = rspCreate.GetId()
	rspRead, readErr := c.GetIdentityProvider(context.Background(), rRead)
	failIfError(t, readErr)
	verifyResponseIsValid(t, rspRead, readErr)
	assert.Equal(t, *rRead.IdentityProviderId, *rspRead.GetId())


	//Listing
	request := identity.ListIdentityProvidersRequest{}
	request.CompartmentId = common.String(getTenancyID())
	request.Protocol = identity.ListIdentityProvidersProtocolSaml2
	response, err := c.ListIdentityProviders(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, response.Items)
	presentAndEqual := false
	for _, val := range response.Items {
		if *val.GetId() == *rspCreate.GetId() {
			presentAndEqual = reflect.TypeOf(identity.Saml2IdentityProvider{}) ==  reflect.TypeOf(val)
		}
	}
	assert.True(t, presentAndEqual)

	items := response.Items

	nextRequest := identity.ListIdentityProvidersRequest{CompartmentId: request.CompartmentId,
		Protocol:identity.ListIdentityProvidersProtocolSaml2}
	nextRequest.Page = response.OpcNextPage

	for nextRequest.Page != nil {
		if r, err := c.ListIdentityProviders(context.Background(), nextRequest); err == nil {
			items = append(items, r.Items...)
			nextRequest.Page = r.OpcNextPage
		} else {
			failIfError(t, err)
			break
		}
	}

	return
}

func TestIdentityClient_GetTenancy(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := identity.GetTenancyRequest{TenancyId: common.String(getTenancyID())}
	r, err := c.GetTenancy(context.Background(), request)
	verifyResponseIsValid(t, r, err)

	assert.Equal(t, request.TenancyId, r.Id)
	assert.NotEmpty(t, r.OpcRequestId)

	return
}

func TestIdentityClient_ListAvailabilityDomains(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := identity.ListAvailabilityDomainsRequest{CompartmentId: common.String(getCompartmentID())}
	r, err := c.ListAvailabilityDomains(context.Background(), request)
	failIfError(t, err)

	items := r.Items

	assert.NotEmpty(t, items)

	return
}

func TestIdentityClient_ListCompartments(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := identity.ListCompartmentsRequest{CompartmentId: common.String(getTenancyID())}
	r, err := c.ListCompartments(context.Background(), request)
	failIfError(t, err)

	items := r.Items

	nextRequest := identity.ListCompartmentsRequest{CompartmentId: request.CompartmentId}
	nextRequest.Page = r.OpcNextPage

	for nextRequest.Page != nil {
		if r, err := c.ListCompartments(context.Background(), nextRequest); err == nil {
			items = append(items, r.Items...)
			nextRequest.Page = r.OpcNextPage
		} else {
			failIfError(t, err)
			break
		}
	}

	assert.NotEmpty(t, items)

	return
}

func TestIdentityClient_ListRegions(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	r, err := c.ListRegions(context.Background())
	verifyResponseIsValid(t, r, err)
	assert.NotEmpty(t, r.Items)
	return
}

func TestIdentityClient_UpdateUserState(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	usr, usrErr := createTestUser(c)
	failIfError(t, usrErr)
	defer deleteTestUser(c, usr.Id)

	request := identity.UpdateUserStateRequest{}
	request.UserId = usr.Id
	request.Blocked = common.Bool(false)

	r, err := c.UpdateUserState(context.Background(), request)
	verifyResponseIsValid(t, r, err)

	assert.Equal(t, identity.UserLifecycleStateActive, r.LifecycleState)
	return
}

// This test can only realistically be run once since once a region is subscribed to, there is no
// mechanism to unsubscribe to it. For now we will skip
func TestIdentityClient_CreateRegionSubscription(t *testing.T) {
	t.Skip("SKIPPING: Region Subscriptions cannot be undone.")
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := identity.CreateRegionSubscriptionRequest{}
	request.TenancyId = common.String(getTenancyID())
	request.RegionKey = common.String("FRA")
	r, err := c.CreateRegionSubscription(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListRegionSubscriptions(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := identity.ListRegionSubscriptionsRequest{TenancyId: common.String(getTenancyID())}
	r, err := c.ListRegionSubscriptions(context.Background(), request)
	failIfError(t, err)
	items := r.Items
	assert.NotEmpty(t, items)
	return
}

func TestBadHost(t *testing.T) {
	client, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	client.Host = "badhostname"
	response, err := client.ListRegions(context.Background())
	assert.Nil(t, response.RawResponse)
	assert.Error(t, err)
}
