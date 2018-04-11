// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package integtest

import (
	"context"
	"fmt"

	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/identity"
	"github.com/stretchr/testify/assert"
)

// Group operations CRUD
func TestIdentityClient_GroupCRUD(t *testing.T) {
	// test should not fail if a previous run failed to clean up
	freeformTags := createOrGetFreeformTags(t)
	groupName := getUniqueName("Group_CRUD")
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := identity.CreateGroupRequest{}
	request.CompartmentId = common.String(getTenancyID())
	request.Name = common.String(groupName)
	request.Description = common.String("GoSDK_someGroupDesc")
	request.FreeformTags = freeformTags
	request.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	r, err := c.CreateGroup(context.Background(), request)
	verifyResponseIsValid(t, r, err)
	failIfError(t, err)

	// if we've successfully created a group during testing, make sure that we delete it
	defer func() {
		//Delete
		rDel := identity.DeleteGroupRequest{
			GroupId: r.Id,
			RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
		}
		resDelete, err := c.DeleteGroup(context.Background(), rDel)
		verifyResponseIsValid(t, resDelete, err)
	}()

	//Get with polling
	pollUntilActive := func(r common.OCIOperationResponse) bool {
		if converted, ok := r.Response.(identity.GetGroupResponse); ok {
			return converted.LifecycleState != identity.GroupLifecycleStateActive
		}
		return true
	}
	pollingGetRequest := identity.GetGroupRequest{
		GroupId: r.Id,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: getExponentialBackoffRetryPolicy(uint(10), pollUntilActive),
		},
	}
	// validate lifecycle state enum value after create
	responseAfterPolling, errAfterPolling := c.GetGroup(context.Background(), pollingGetRequest)
	verifyResponseIsValid(t, responseAfterPolling, errAfterPolling)
	assert.Equal(t, identity.GroupLifecycleStateActive, responseAfterPolling.Group.LifecycleState)
	assert.Equal(t, freeformTags, responseAfterPolling.FreeformTags)
	failIfError(t, errAfterPolling)

	//Update
	rUpdate := identity.UpdateGroupRequest{GroupId: r.Id}
	rUpdate.Description = common.String("New description")
	rUpdate.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	resUpdate, err := c.UpdateGroup(context.Background(), rUpdate)
	verifyResponseIsValid(t, resUpdate, err)
	failIfError(t, err)
	assert.NotNil(t, resUpdate.Id)
}

type fakeDispatcher struct {
	Reg   string
	Valid bool
}

func (f *fakeDispatcher) Do(r *http.Request) (*http.Response, error) {
	f.Valid = strings.Contains(r.URL.Host, f.Reg)
	return nil, fmt.Errorf("Fake dispatcher")
}

func TestIdentityClient_OverrideRegion(t *testing.T) {
	c, _ := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	region := "newRegion"
	c.SetRegion(region)
	f := fakeDispatcher{Reg: region}
	// Avoid calling the service as we do no know if we have access to that region
	c.HTTPClient = &f
	rList := identity.ListGroupsRequest{
		CompartmentId: common.String(getTenancyID()),
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	c.ListGroups(context.Background(), rList)
	assert.True(t, f.Valid)
}

func TestIdentityClient_ListGroups(t *testing.T) {

	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := identity.ListGroupsRequest{
		CompartmentId: common.String(getTenancyID()),
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	r, err := c.ListGroups(context.Background(), request)
	verifyResponseIsValid(t, r, err)
	failIfError(t, err)

	items := r.Items

	nextRequest := identity.ListGroupsRequest{
		CompartmentId: request.CompartmentId,
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	nextRequest.Page = r.OpcNextPage

	for nextRequest.Page != nil {
		if r, err = c.ListGroups(context.Background(), nextRequest); err == nil {
			verifyResponseIsValid(t, r, err)
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
	c, cfgErr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, cfgErr)

	request := identity.CreateCompartmentRequest{
		CreateCompartmentDetails: identity.CreateCompartmentDetails{
			Name:          common.String("Compartment_Test"),
			Description:   common.String("Go SDK Comparment Test"),
			CompartmentId: common.String(getTenancyID()),
		},
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}

	r, err := c.CreateCompartment(context.Background(), request)
	verifyResponseIsValid(t, r, err)

	//Get with polling
	pollUntilActive := func(r common.OCIOperationResponse) bool {
		if converted, ok := r.Response.(identity.GetCompartmentResponse); ok {
			return converted.LifecycleState != identity.CompartmentLifecycleStateActive
		}
		return true
	}
	pollingGetRequest := identity.GetCompartmentRequest{
		CompartmentId: r.CompartmentId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: getExponentialBackoffRetryPolicy(uint(0), pollUntilActive),
		},
	}
	// validate lifecycle state enum value after create
	responseAfterPolling, errAfterPolling := c.GetCompartment(context.Background(), pollingGetRequest)
	verifyResponseIsValid(t, responseAfterPolling, errAfterPolling)
	assert.Equal(t, identity.CompartmentLifecycleStateActive, responseAfterPolling.LifecycleState)
	failIfError(t, errAfterPolling)

	assert.NotEmpty(t, r.Id)
	assert.Equal(t, request.Name, r.Name)
	assert.Equal(t, request.Description, r.Description)
	return
}

//Comparment RU
func TestIdentityClient_UpdateCompartment(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	//Update
	request := identity.UpdateCompartmentRequest{
		UpdateCompartmentDetails: identity.UpdateCompartmentDetails{
			// cannot use same name to update compartment, generate a random one via this function
			Name:        common.String(GoSDK2_Test_Prefix + "UpdComp" + getRandomString(10)),
			Description: common.String("GOSDK2 description2"),
		},
		CompartmentId: common.String(getCompartmentID()),
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	r, err := c.UpdateCompartment(context.Background(), request)
	verifyResponseIsValid(t, r, err)
	failIfError(t, err)

	rRead := identity.GetCompartmentRequest{
		CompartmentId: common.String(getTenancyID()),
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	resRead, err := c.GetCompartment(context.Background(), rRead)
	verifyResponseIsValid(t, resRead, err)
	failIfError(t, err)
	return
}

//User Operations
func TestIdentityClient_ListUsers(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := identity.ListUsersRequest{
		CompartmentId: common.String(getTenancyID()),
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	r, err := c.ListUsers(context.Background(), request)
	verifyResponseIsValid(t, r, err)
	failIfError(t, err)

	items := r.Items

	nextRequest := identity.ListUsersRequest{
		CompartmentId: request.CompartmentId,
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	nextRequest.Page = r.OpcNextPage
	for nextRequest.Page != nil {
		if r, err = c.ListUsers(context.Background(), nextRequest); err == nil {
			verifyResponseIsValid(t, r, err)
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
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	user := createTestUser(t, common.String(getUniqueName("User_")))
	assert.NotEmpty(t, user)
	assert.NotEmpty(t, user.Id)

	// if we've successfully created a user during testing, make sure that we delete it
	defer func() {
		//remove
		rDelete := identity.DeleteUserRequest{}
		rDelete.UserId = user.Id
		rDelete.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
		resDelete, err := c.DeleteUser(context.Background(), rDelete)
		verifyResponseIsValid(t, resDelete, err)
		assert.NotEmpty(t, resDelete.OpcRequestId)
	}()

	//Get with polling
	pollUntilActive := func(r common.OCIOperationResponse) bool {
		if converted, ok := r.Response.(identity.GetUserResponse); ok {
			return converted.LifecycleState != identity.UserLifecycleStateActive
		}
		return true
	}
	pollingGetRequest := identity.GetUserRequest{
		UserId: user.Id,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: getExponentialBackoffRetryPolicy(uint(0), pollUntilActive),
		},
	}
	// validate lifecycle state enum value after create
	responseAfterPolling, errAfterPolling := c.GetUser(context.Background(), pollingGetRequest)
	verifyResponseIsValid(t, responseAfterPolling, errAfterPolling)
	assert.Equal(t, identity.UserLifecycleStateActive, responseAfterPolling.LifecycleState)

	//Update
	rUpdate := identity.UpdateUserRequest{}
	rUpdate.UserId = user.Id
	rUpdate.Description = common.String("This is a new description")
	resUpdate, err := c.UpdateUser(context.Background(), rUpdate)
	verifyResponseIsValid(t, resUpdate, err)

	return
}

//User-Group operations
func TestIdentityClient_AddUserToGroup(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	// for robustness, create a user and group to use for this test. delete it at the end
	reqAddUser := identity.CreateUserRequest{}
	reqAddUser.CompartmentId = common.String(getTenancyID())
	reqAddUser.Name = common.String(getUniqueName("AUTG_User"))
	reqAddUser.Description = common.String("AddUserToGroup Test User")
	reqAddUser.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	rspAddUser, err1 := c.CreateUser(context.Background(), reqAddUser)
	verifyResponseIsValid(t, rspAddUser, err1)
	failIfError(t, err1)

	defer func() {
		// Delete the user
		reqUserDelete := identity.DeleteUserRequest{
			UserId: rspAddUser.Id,
			RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
		}
		delUserResponse, delUserErr := c.DeleteUser(context.Background(), reqUserDelete)
		verifyResponseIsValid(t, delUserResponse, delUserErr)
	}()

	reqAddGroup := identity.CreateGroupRequest{}
	reqAddGroup.CompartmentId = common.String(getTenancyID())
	reqAddGroup.Name = common.String(getUniqueName("AUTG_Group_"))
	reqAddGroup.Description = common.String("AddUserToGroup Test Group")
	reqAddGroup.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	rspAddGroup, err2 := c.CreateGroup(context.Background(), reqAddGroup)
	verifyResponseIsValid(t, rspAddGroup, err2)
	failIfError(t, err2)

	defer func() {
		// Delete the group
		reqGroupDelete := identity.DeleteGroupRequest{
			GroupId: rspAddGroup.Id,
			RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
		}
		delRes, delGrpErr := c.DeleteGroup(context.Background(), reqGroupDelete)
		verifyResponseIsValid(t, delRes, delGrpErr)
		assert.NotEmpty(t, delRes.OpcRequestId)
	}()

	//add
	reqAdd := identity.AddUserToGroupRequest{}
	reqAdd.UserId = rspAddUser.Id
	reqAdd.GroupId = rspAddGroup.Id
	reqAdd.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	rspAdd, err := c.AddUserToGroup(context.Background(), reqAdd)
	verifyResponseIsValid(t, rspAdd, err)
	failIfError(t, err)

	defer func() {
		//remove
		requestRemove := identity.RemoveUserFromGroupRequest{
			UserGroupMembershipId: rspAdd.UserGroupMembership.Id,
			RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
		}
		_, err = c.RemoveUserFromGroup(context.Background(), requestRemove)
		failIfError(t, err)
	}()

	//Get with polling
	pollUntilActive := func(r common.OCIOperationResponse) bool {
		if converted, ok := r.Response.(identity.GetUserGroupMembershipResponse); ok {
			return converted.LifecycleState != identity.UserGroupMembershipLifecycleStateActive
		}
		return true
	}
	pollingGetRequest := identity.GetUserGroupMembershipRequest{
		UserGroupMembershipId: rspAdd.Id,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: getExponentialBackoffRetryPolicy(uint(10), pollUntilActive),
		},
	}
	// validate lifecycle state enum value after create
	responseAfterPolling, errAfterPolling := c.GetUserGroupMembership(context.Background(), pollingGetRequest)
	verifyResponseIsValid(t, responseAfterPolling, errAfterPolling)
	assert.Equal(t, identity.UserGroupMembershipLifecycleStateActive, responseAfterPolling.LifecycleState)

	// Read
	assert.Equal(t, rspAdd.Id, responseAfterPolling.Id)

	return

}

func TestIdentityClient_ListUserGroupMemberships(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := identity.ListUserGroupMembershipsRequest{}
	request.UserId = common.String(getUserID())
	request.CompartmentId = common.String(getTenancyID())
	request.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	r, err := c.ListUserGroupMemberships(context.Background(), request)
	verifyResponseIsValid(t, r, err)
	failIfError(t, err)

	items := r.Items

	nextRequest := identity.ListUserGroupMembershipsRequest{
		CompartmentId: request.CompartmentId,
		UserId: request.UserId,
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	nextRequest.Page = r.OpcNextPage

	for nextRequest.Page != nil {
		if r, err := c.ListUserGroupMemberships(context.Background(), nextRequest); err == nil {
			verifyResponseIsValid(t, r, err)
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
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	// create or get the test user
	user := createOrGetUser(t)
	request := identity.CreateOrResetUIPasswordRequest{}
	request.UserId = user.Id
	request.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	rspCreate, err := c.CreateOrResetUIPassword(context.Background(), request)
	verifyResponseIsValid(t, rspCreate, err)
	failIfError(t, err)

	assert.NotEmpty(t, rspCreate.OpcRequestId)
	assert.Equal(t, user.Id, rspCreate.UserId)
	assert.NotEmpty(t, rspCreate.Password)

	assert.Equal(t, identity.UiPasswordLifecycleStateActive, rspCreate.LifecycleState)

	// make the request again and ensure that we get a different password
	rspReset, err := c.CreateOrResetUIPassword(context.Background(), request)
	verifyResponseIsValid(t, rspReset, err)
	failIfError(t, err)

	assert.Equal(t, rspCreate.UserId, rspReset.UserId)
	assert.NotEqual(t, rspCreate.Password, rspReset.Password)
	assert.Equal(t, identity.UiPasswordLifecycleStateActive, rspCreate.LifecycleState)

	return
}

func TestIdentityClient_SwiftPasswordCRUD(t *testing.T) {

	createDesc := "Go SDK Test Swift Password - CREATED"
	updateDesc := "Go SDK Test Swift Password - UPDATED"
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	usr := createOrGetUser(t)

	// Create Swift Password
	addReq := identity.CreateSwiftPasswordRequest{UserId: usr.Id}
	addReq.Description = &createDesc
	addReq.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	rspPwd, err := c.CreateSwiftPassword(context.Background(), addReq)
	verifyResponseIsValid(t, rspPwd, err)

	// Delete Swift Password
	defer func() {
		delReq := identity.DeleteSwiftPasswordRequest{}
		delReq.UserId = usr.Id
		delReq.SwiftPasswordId = rspPwd.Id
		delReq.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
		delRes, err := c.DeleteSwiftPassword(context.Background(), delReq)
		verifyResponseIsValid(t, delRes, err)
		failIfError(t, err)
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
	updReq.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	updRsp, err := c.UpdateSwiftPassword(context.Background(), updReq)
	verifyResponseIsValid(t, updRsp, err)

	assert.NotEqual(t, rspPwd.Password, updRsp.Password)
	assert.Equal(t, identity.SwiftPasswordLifecycleStateActive, updRsp.LifecycleState)
	assert.Equal(t, updateDesc, *updRsp.Description)

	//assert.NotEmpty(t, updRsp.ExpiresOn)

	return
}

func TestIdentityClient_ListSwiftPasswords(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	usr := createOrGetUser(t)

	request := identity.ListSwiftPasswordsRequest{
		UserId: usr.Id,
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	r, err := c.ListSwiftPasswords(context.Background(), request)
	verifyResponseIsValid(t, r, err)

	deletePwd := func(pwdId *string) {
		delReq := identity.DeleteSwiftPasswordRequest{}
		delReq.UserId = usr.Id
		delReq.SwiftPasswordId = pwdId
		delReq.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
		delRes, err := c.DeleteSwiftPassword(context.Background(), delReq)
		verifyResponseIsValid(t, delRes, err)
		failIfError(t, err)
	}

	// delete password if already there
	if len(r.Items) != 0 {
		for _, item := range r.Items {
			deletePwd(item.Id)
		}
	}

	pwdReq := identity.CreateSwiftPasswordRequest{
		UserId: usr.Id,
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	pwdReq.Description = common.String("Test Swift Password 1")
	pwdRsp1, err := c.CreateSwiftPassword(context.Background(), pwdReq)
	verifyResponseIsValid(t, pwdRsp1, err)

	pwdReq.Description = common.String("Test Swift Password 2")
	pwdRsp2, err := c.CreateSwiftPassword(context.Background(), pwdReq)
	verifyResponseIsValid(t, pwdRsp2, err)

	request = identity.ListSwiftPasswordsRequest{
		UserId: usr.Id,
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	r, err = c.ListSwiftPasswords(context.Background(), request)
	verifyResponseIsValid(t, r, err)
	assert.Equal(t, 2, len(r.Items))
	assert.NotEqual(t, r.Items[0].Id, r.Items[1].Id)
	assert.NotEqual(t, r.Items[0].Description, r.Items[1].Description)

	// Delete Swift Password just created
	defer func() {
		pwdIDs := []*string{pwdRsp1.Id, pwdRsp2.Id}
		for _, pwdID := range pwdIDs {
			deletePwd(pwdID)
		}
	}()

	return
}

//Policy Operations see DEX-1945
func TestIdentityClient_PolicyCRUD(t *testing.T) {
	t.Skip("Policy Operations issue. See DEX-1945 ")

	//Create
	//client := identity.NewIdentityClientForRegion(getRegion())
	client, cfgErr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, cfgErr)

	createRequest := identity.CreatePolicyRequest{}
	createRequest.CompartmentId = common.String(getTenancyID())
	createRequest.Name = common.String("goSDK2Policy2")
	createRequest.Description = common.String("some policy")
	createRequest.Statements = []string{"Allow group goSDK2CreateGroup read all-resources on compartment egineztest"}
	createRequest.VersionDate = &common.SDKTime{time.Now()}
	createRequest.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	createResponse, err := client.CreatePolicy(context.Background(), createRequest)
	verifyResponseIsValid(t, createResponse, err)

	defer func() {
		// Delete
		request := identity.DeletePolicyRequest{
			PolicyId: createResponse.Id,
			RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
		}
		delRes, err := client.DeletePolicy(context.Background(), request)
		verifyResponseIsValid(t, delRes, err)
		assert.NotEmpty(t, delRes.OpcRequestId)
	}()

	//Read
	readRequest := identity.GetPolicyRequest{}
	readRequest.PolicyId = createResponse.Id
	readRequest.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	readResponse, err := client.GetPolicy(context.Background(), readRequest)
	verifyResponseIsValid(t, readResponse, err)

	//Update

	updateRequest := identity.UpdatePolicyRequest{}
	updateRequest.PolicyId = createResponse.Id
	updateRequest.Description = common.String("new description")
	updateRequest.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	updateResponse, err := client.UpdatePolicy(context.Background(), updateRequest)
	verifyResponseIsValid(t, updateResponse, err)

	return
}

func TestIdentityClient_ListPolicies(t *testing.T) {
	c, cfgErr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, cfgErr)
	listRequest := identity.ListPoliciesRequest{}
	listRequest.CompartmentId = common.String(getTenancyID())
	listRequest.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	listResponse, err := c.ListPolicies(context.Background(), listRequest)
	verifyResponseIsValid(t, listResponse, err)
	failIfError(t, err)

	items := listResponse.Items

	nextRequest := identity.ListPoliciesRequest{
		CompartmentId: listRequest.CompartmentId,
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	nextRequest.Page = listResponse.OpcNextPage

	for nextRequest.Page != nil {
		if r, err := c.ListPolicies(context.Background(), nextRequest); err == nil {
			verifyResponseIsValid(t, r, err)
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
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := identity.CreateCustomerSecretKeyRequest{}
	request.UserId = common.String(getUserID())
	request.DisplayName = common.String("GolangSDK2TestSecretKey")
	request.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	resCreate, err := c.CreateCustomerSecretKey(context.Background(), request)
	verifyResponseIsValid(t, resCreate, err)
	failIfError(t, err)

	defer func() {
		//remove
		rDelete := identity.DeleteCustomerSecretKeyRequest{}
		rDelete.CustomerSecretKeyId = resCreate.Id
		rDelete.UserId = common.String(getUserID())
		rDelete.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
		delRes, err := c.DeleteCustomerSecretKey(context.Background(), rDelete)
		verifyResponseIsValid(t, delRes, err)
		failIfError(t, err)
		assert.NotEmpty(t, delRes.OpcRequestId)
	}()

	// validate user membership lifecycle state enum value after create
	assert.Equal(t, identity.CustomerSecretKeyLifecycleStateActive, resCreate.LifecycleState)

	//Update
	rUpdate := identity.UpdateCustomerSecretKeyRequest{}
	rUpdate.CustomerSecretKeyId = resCreate.Id
	rUpdate.UserId = common.String(getUserID())
	rUpdate.DisplayName = common.String("This is a new description")
	rUpdate.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	resUpdate, err := c.UpdateCustomerSecretKey(context.Background(), rUpdate)
	verifyResponseIsValid(t, resUpdate, err)
	failIfError(t, err)

	return
}

func TestIdentityClient_ListCustomerSecretKeys(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := identity.ListCustomerSecretKeysRequest{}
	request.UserId = common.String(getUserID())
	request.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	r, err := c.ListCustomerSecretKeys(context.Background(), request)
	verifyResponseIsValid(t, r, err)
	failIfError(t, err)
	return
}

//Apikeys
func TestIdentityClient_ApiKeyCRUD(t *testing.T) {
	userId := ""
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := identity.UploadApiKeyRequest{}
	request.UserId = common.String(userId)
	request.Key = common.String("some key")
	resCreate, err := c.UploadApiKey(context.Background(), request)
	assert.Error(t, err)

	defer func() {
		//remove
		rDelete := identity.DeleteApiKeyRequest{}
		rDelete.Fingerprint = resCreate.Fingerprint
		rDelete.UserId = common.String(userId)
		_, err := c.DeleteApiKey(context.Background(), rDelete)
		assert.Error(t, err)
	}()

	// TODO: [2017-Nov-07::shalka] presently LifecycleState isn't being set on ApiKey struct in the Response => merits
	//  further investigation
	// validate api key lifecycle state enum value after create
	// assert.Equal(t, identity.API_KEY_LIFECYCLE_STATE_ACTIVE, resCreate.LifecycleState)

	return
}
func TestIdentityClient_ListApiKeys(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := identity.ListApiKeysRequest{}
	request.UserId = common.String(getUserID())
	request.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	r, err := c.ListApiKeys(context.Background(), request)
	verifyResponseIsValid(t, r, err)
	failIfError(t, err)
	return
}

func TestIdentityClient_IdentityProviderCRUD(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
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
	rCreate.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()

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
			rDelete.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
			delRes, err := c.DeleteIdentityProvider(context.Background(), rDelete)
			verifyResponseIsValid(t, delRes, err)
			failIfError(t, err)
			assert.NotEmpty(t, delRes.OpcRequestId)
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
	rRead.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
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
	rUpdate.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	rspUpdate, updErr := c.UpdateIdentityProvider(context.Background(), rUpdate)

	failIfError(t, updErr)
	verifyResponseIsValid(t, rspUpdate, updErr)
	assert.Equal(t, *rspCreate.GetId(), *rspUpdate.GetId())
	assert.Equal(t, "New description", *rspUpdate.GetDescription())

	g := createOrGetTestGroup(t)

	reqCreateMapping := identity.CreateIdpGroupMappingRequest{}
	reqCreateMapping.IdentityProviderId = rspCreate.GetId()
	reqCreateMapping.GroupId = g.Id
	reqCreateMapping.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	idpGrpName := *rspCreate.GetName()
	groupName := *g.Name
	reqCreateMapping.IdpGroupName = common.String(idpGrpName + "_TO_" + groupName)

	rspCreateMapping, createMapErr := c.CreateIdpGroupMapping(context.Background(), reqCreateMapping)
	verifyResponseIsValid(t, rspCreateMapping, createMapErr)
	failIfError(t, createMapErr)

	// Delete mapping
	defer func() {
		fmt.Println("Deleting Identity Provider Group Mapping")
		reqDelete := identity.DeleteIdpGroupMappingRequest{
			MappingId: rspCreateMapping.Id,
			IdentityProviderId: rspCreateMapping.IdpId,
			RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
		}
		delRes, delErr := c.DeleteIdpGroupMapping(context.Background(), reqDelete)
		verifyResponseIsValid(t, delRes, delErr)
		failIfError(t, delErr)
	}()

	assert.NotEmpty(t, *rspCreateMapping.Id)
	assert.NotEmpty(t, *rspCreateMapping.GroupId)
	assert.NotEmpty(t, *rspCreateMapping.OpcRequestId)
	assert.NotEmpty(t, *rspCreateMapping.IdpGroupName)
	assert.Equal(t, *rspCreate.GetId(), *rspCreateMapping.IdpId)
	assert.NotEmpty(t, rspCreateMapping.TimeCreated)



	//Get with polling
	pollUntilActive := func(r common.OCIOperationResponse) bool {
		if converted, ok := r.Response.(identity.GetIdpGroupMappingResponse); ok {
			return converted.LifecycleState != identity.IdpGroupMappingLifecycleStateActive
		}
		return true
	}
	pollingGetRequest := identity.GetIdpGroupMappingRequest{
		IdentityProviderId: rspCreateMapping.IdpId,
		MappingId: rspCreateMapping.Id,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: getExponentialBackoffRetryPolicy(uint(10), pollUntilActive),
		},
	}
	// validate lifecycle state enum value after create
	responseAfterPolling, errAfterPolling := c.GetIdpGroupMapping(context.Background(), pollingGetRequest)
	verifyResponseIsValid(t, responseAfterPolling, errAfterPolling)
	assert.Equal(t, identity.IdpGroupMappingLifecycleStateActive, responseAfterPolling.LifecycleState)
	failIfError(t, errAfterPolling)

	assert.Equal(t, rspCreateMapping.Id, responseAfterPolling.Id)
	assert.Equal(t, rspCreateMapping.IdpId, responseAfterPolling.IdpId)

	//update group mapping
	reqUpdMapping := identity.UpdateIdpGroupMappingRequest{}
	reqUpdMapping.MappingId = responseAfterPolling.Id
	reqUpdMapping.IdentityProviderId = responseAfterPolling.IdpId
	reqUpdMapping.GroupId = responseAfterPolling.GroupId
	reqUpdMapping.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	updatedName := *responseAfterPolling.IdpGroupName + " - Updated"
	reqUpdMapping.IdpGroupName = common.String(updatedName)
	rspUpdMapping, updMapErr := c.UpdateIdpGroupMapping(context.Background(), reqUpdMapping)
	verifyResponseIsValid(t, rspUpdMapping, updMapErr)

	assert.NotEmpty(t, *rspUpdMapping.Id)
	assert.NotEmpty(t, *rspUpdMapping.GroupId)
	assert.NotEmpty(t, *rspUpdMapping.OpcRequestId)
	assert.Equal(t, *responseAfterPolling.Id, *rspUpdMapping.Id)
	assert.Equal(t, *responseAfterPolling.IdpId, *rspUpdMapping.IdpId)
	assert.NotEqual(t, *responseAfterPolling.IdpGroupName, *rspUpdMapping.IdpGroupName)
	assert.NotEmpty(t, rspUpdMapping.TimeCreated)
	assert.Equal(t, identity.IdpGroupMappingLifecycleStateActive, rspUpdMapping.LifecycleState)

	return
}

func TestIdentityClient_ListIdentityProviders(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
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
	rCreate.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()

	//Create
	rspCreate, createErr := c.CreateIdentityProvider(context.Background(), rCreate)
	failIfError(t, createErr)
	verifyResponseIsValid(t, rspCreate, createErr)

	//Delete
	deleteFn := func() {
		if rspCreate.GetId() != nil {
			rDelete := identity.DeleteIdentityProviderRequest{}
			rDelete.IdentityProviderId = rspCreate.GetId()
			rDelete.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
			delRes, err := c.DeleteIdentityProvider(context.Background(), rDelete)
			verifyResponseIsValid(t, delRes, err)
			failIfError(t, err)
		}
	}
	defer deleteFn()

	rRead := identity.GetIdentityProviderRequest{}
	rRead.IdentityProviderId = rspCreate.GetId()
	rRead.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	rspRead, readErr := c.GetIdentityProvider(context.Background(), rRead)
	failIfError(t, readErr)
	verifyResponseIsValid(t, rspRead, readErr)
	assert.Equal(t, *rRead.IdentityProviderId, *rspRead.GetId())

	//Listing
	request := identity.ListIdentityProvidersRequest{}
	request.CompartmentId = common.String(getTenancyID())
	request.Protocol = identity.ListIdentityProvidersProtocolSaml2
	request.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()
	response, err := c.ListIdentityProviders(context.Background(), request)
	verifyResponseIsValid(t, response, err)
	failIfError(t, err)
	assert.NotEmpty(t, response.Items)
	presentAndEqual := false
	for _, val := range response.Items {
		if *val.GetId() == *rspCreate.GetId() {
			presentAndEqual = reflect.TypeOf(identity.Saml2IdentityProvider{}) == reflect.TypeOf(val)
		}
	}
	assert.True(t, presentAndEqual)

	items := response.Items

	nextRequest := identity.ListIdentityProvidersRequest{
		CompartmentId: request.CompartmentId,
		Protocol: identity.ListIdentityProvidersProtocolSaml2,
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	nextRequest.Page = response.OpcNextPage

	for nextRequest.Page != nil {
		if r, err := c.ListIdentityProviders(context.Background(), nextRequest); err == nil {
			verifyResponseIsValid(t, r, err)
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
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := identity.GetTenancyRequest{
		TenancyId: common.String(getTenancyID()),
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	r, err := c.GetTenancy(context.Background(), request)
	verifyResponseIsValid(t, r, err)

	assert.Equal(t, request.TenancyId, r.Id)
	assert.NotEmpty(t, r.OpcRequestId)

	return
}

func TestIdentityClient_GetTenancy_EmptyTenancyID(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	// use empty string for tenancyID which is in path
	request := identity.GetTenancyRequest{TenancyId: common.String("")}
	r, err := c.GetTenancy(context.Background(), request)

	assert.Error(t, err)
	assert.Equal(t, "value cannot be empty for field TenancyId in path", err.Error())
	assert.Empty(t, r.OpcRequestId)
}

func TestIdentityClient_ListAvailabilityDomains(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := identity.ListAvailabilityDomainsRequest{
		CompartmentId: common.String(getCompartmentID()),
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	r, err := c.ListAvailabilityDomains(context.Background(), request)
	verifyResponseIsValid(t, r, err)
	failIfError(t, err)

	items := r.Items

	assert.NotEmpty(t, items)

	return
}

func TestIdentityClient_ListCompartments(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := identity.ListCompartmentsRequest{
		CompartmentId: common.String(getTenancyID()),
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	r, err := c.ListCompartments(context.Background(), request)
	verifyResponseIsValid(t, r, err)
	failIfError(t, err)

	items := r.Items

	nextRequest := identity.ListCompartmentsRequest{
		CompartmentId: request.CompartmentId,
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	nextRequest.Page = r.OpcNextPage

	for nextRequest.Page != nil {
		if r, err := c.ListCompartments(context.Background(), nextRequest); err == nil {
			verifyResponseIsValid(t, r, err)
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
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	r, err := c.ListRegions(context.Background())
	verifyResponseIsValid(t, r, err)
	assert.NotEmpty(t, r.Items)
	return
}

func TestIdentityClient_UpdateUserState(t *testing.T) {
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	usr := createOrGetUser(t)

	request := identity.UpdateUserStateRequest{}
	request.UserId = usr.Id
	request.Blocked = common.Bool(false)
	request.RequestMetadata = getRequestMetadataWithDefaultRetryPolicy()

	r, err := c.UpdateUserState(context.Background(), request)
	verifyResponseIsValid(t, r, err)

	assert.Equal(t, identity.UserLifecycleStateActive, r.LifecycleState)
	return
}

// This test can only realistically be run once since once a region is subscribed to, there is no
// mechanism to unsubscribe to it. For now we will skip
func TestIdentityClient_CreateRegionSubscription(t *testing.T) {
	t.Skip("SKIPPING: Region Subscriptions cannot be undone.")
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
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
	c, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	request := identity.ListRegionSubscriptionsRequest{
		TenancyId: common.String(getTenancyID()),
		RequestMetadata: getRequestMetadataWithDefaultRetryPolicy(),
	}
	r, err := c.ListRegionSubscriptions(context.Background(), request)
	verifyResponseIsValid(t, r, err)
	failIfError(t, err)
	items := r.Items
	assert.NotEmpty(t, items)
	return
}

func TestBadHost(t *testing.T) {
	client, clerr := identity.NewIdentityClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)
	client.Host = "badhostname"
	response, err := client.ListRegions(context.Background())
	assert.Nil(t, response.RawResponse)
	assert.Error(t, err)
}
