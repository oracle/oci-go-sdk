// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package integtest

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/identity"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testRegionForIdentity = common.REGION_PHX
)

func TestIdentityClient_AddUserToGroup(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.AddUserToGroupRequest{}
	r, err := c.AddUserToGroup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_CreateCompartment(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.CreateCompartmentRequest{}
	r, err := c.CreateCompartment(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_CreateCustomerSecretKey(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.CreateCustomerSecretKeyRequest{}
	r, err := c.CreateCustomerSecretKey(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_CreateGroup(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.CreateGroupRequest{}
	r, err := c.CreateGroup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_CreateIdentityProvider(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.CreateIdentityProviderRequest{}
	r, err := c.CreateIdentityProvider(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_CreateIdpGroupMapping(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.CreateIdpGroupMappingRequest{}
	r, err := c.CreateIdpGroupMapping(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_CreateOrResetUIPassword(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.CreateOrResetUIPasswordRequest{}
	r, err := c.CreateOrResetUIPassword(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_CreatePolicy(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.CreatePolicyRequest{}
	r, err := c.CreatePolicy(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_CreateRegionSubscription(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.CreateRegionSubscriptionRequest{}
	r, err := c.CreateRegionSubscription(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_CreateSwiftPassword(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.CreateSwiftPasswordRequest{}
	r, err := c.CreateSwiftPassword(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_CreateUser(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.CreateUserRequest{}
	r, err := c.CreateUser(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_DeleteApiKey(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.DeleteApiKeyRequest{}
	err := c.DeleteApiKey(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestIdentityClient_DeleteCustomerSecretKey(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.DeleteCustomerSecretKeyRequest{}
	err := c.DeleteCustomerSecretKey(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestIdentityClient_DeleteGroup(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.DeleteGroupRequest{}
	err := c.DeleteGroup(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestIdentityClient_DeleteIdentityProvider(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.DeleteIdentityProviderRequest{}
	err := c.DeleteIdentityProvider(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestIdentityClient_DeleteIdpGroupMapping(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.DeleteIdpGroupMappingRequest{}
	err := c.DeleteIdpGroupMapping(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestIdentityClient_DeletePolicy(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.DeletePolicyRequest{}
	err := c.DeletePolicy(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestIdentityClient_DeleteSwiftPassword(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.DeleteSwiftPasswordRequest{}
	err := c.DeleteSwiftPassword(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestIdentityClient_DeleteUser(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.DeleteUserRequest{}
	err := c.DeleteUser(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestIdentityClient_GetCompartment(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.GetCompartmentRequest{}
	r, err := c.GetCompartment(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_GetGroup(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.GetGroupRequest{}
	r, err := c.GetGroup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_GetIdentityProvider(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.GetIdentityProviderRequest{}
	r, err := c.GetIdentityProvider(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_GetIdpGroupMapping(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.GetIdpGroupMappingRequest{}
	r, err := c.GetIdpGroupMapping(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_GetPolicy(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.GetPolicyRequest{}
	r, err := c.GetPolicy(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_GetTenancy(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.GetTenancyRequest{}
	r, err := c.GetTenancy(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_GetUser(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.GetUserRequest{}
	r, err := c.GetUser(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_GetUserGroupMembership(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.GetUserGroupMembershipRequest{}
	r, err := c.GetUserGroupMembership(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListApiKeys(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.ListApiKeysRequest{}
	r, err := c.ListApiKeys(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListAvailabilityDomains(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.ListAvailabilityDomainsRequest{}
	r, err := c.ListAvailabilityDomains(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListCompartments(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.ListCompartmentsRequest{}
	r, err := c.ListCompartments(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListCustomerSecretKeys(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.ListCustomerSecretKeysRequest{}
	r, err := c.ListCustomerSecretKeys(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListFaultDomains(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.ListFaultDomainsRequest{}
	r, err := c.ListFaultDomains(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListGroups(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.ListGroupsRequest{}
	r, err := c.ListGroups(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListIdentityProviders(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.ListIdentityProvidersRequest{}
	r, err := c.ListIdentityProviders(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListIdpGroupMappings(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.ListIdpGroupMappingsRequest{}
	r, err := c.ListIdpGroupMappings(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListPolicies(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.ListPoliciesRequest{}
	r, err := c.ListPolicies(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListRegionSubscriptions(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.ListRegionSubscriptionsRequest{}
	r, err := c.ListRegionSubscriptions(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListRegions(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	r, err := c.ListRegions(context.Background())
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListSwiftPasswords(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.ListSwiftPasswordsRequest{}
	r, err := c.ListSwiftPasswords(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListUserGroupMemberships(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.ListUserGroupMembershipsRequest{}
	r, err := c.ListUserGroupMemberships(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_ListUsers(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.ListUsersRequest{}
	r, err := c.ListUsers(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_RemoveUserFromGroup(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.RemoveUserFromGroupRequest{}
	err := c.RemoveUserFromGroup(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UpdateCompartment(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.UpdateCompartmentRequest{}
	r, err := c.UpdateCompartment(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UpdateCustomerSecretKey(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.UpdateCustomerSecretKeyRequest{}
	r, err := c.UpdateCustomerSecretKey(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UpdateGroup(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.UpdateGroupRequest{}
	r, err := c.UpdateGroup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UpdateIdentityProvider(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.UpdateIdentityProviderRequest{}
	r, err := c.UpdateIdentityProvider(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UpdateIdpGroupMapping(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.UpdateIdpGroupMappingRequest{}
	r, err := c.UpdateIdpGroupMapping(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UpdatePolicy(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.UpdatePolicyRequest{}
	r, err := c.UpdatePolicy(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UpdateSwiftPassword(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.UpdateSwiftPasswordRequest{}
	r, err := c.UpdateSwiftPassword(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UpdateUser(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.UpdateUserRequest{}
	r, err := c.UpdateUser(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UpdateUserState(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.UpdateUserStateRequest{}
	r, err := c.UpdateUserState(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestIdentityClient_UploadApiKey(t *testing.T) {
	t.Skip("Not implemented")
	c := identity.NewIdentityClientForRegion(testRegionForIdentity)
	request := identity.UploadApiKeyRequest{}
	r, err := c.UploadApiKey(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}
