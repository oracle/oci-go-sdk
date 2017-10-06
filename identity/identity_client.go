// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"net/http"
)

type IdentityClient struct {
	common.SDKClient
}

func NewClient() (client IdentityClient) {
	client = IdentityClient{SDKClient: common.NewClient()}
	client.ServiceName = "identity"
	client.ApiVersion = "20160918"
	return
}

func NewClientWithInterceptor(interceptor common.RequestInterceptor) (client IdentityClient) {
	client = IdentityClient{SDKClient: common.NewClientWithInterceptor(interceptor)}
	client.ServiceName = "identity"
	client.ApiVersion = "20160918"
	return
}

func (client IdentityClient) AddUserToGroup(request AddUserToGroupRequest) (response AddUserToGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/userGroupMemberships/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) CreateCompartment(request CreateCompartmentRequest) (response CreateCompartmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/compartments/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) CreateGroup(request CreateGroupRequest) (response CreateGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/groups/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) CreateIdentityProvider(request CreateIdentityProviderRequest) (response CreateIdentityProviderResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/identityProviders/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) CreateIdpGroupMapping(request CreateIdpGroupMappingRequest) (response CreateIdpGroupMappingResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/identityProviders/{identityProviderId}/groupMappings/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) CreateOrResetUIPassword(request CreateOrResetUIPasswordRequest) (response CreateOrResetUIPasswordResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/users/{userId}/uiPassword", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) CreatePolicy(request CreatePolicyRequest) (response CreatePolicyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/policies/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) CreateRegionSubscription(request CreateRegionSubscriptionRequest) (response CreateRegionSubscriptionResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/tenancies/{tenancyId}/regionSubscriptions", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) CreateSwiftPassword(request CreateSwiftPasswordRequest) (response CreateSwiftPasswordResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/users/{userId}/swiftPasswords/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) CreateUser(request CreateUserRequest) (response CreateUserResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/users/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) DeleteApiKey(request DeleteApiKeyRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/users/{userId}/apiKeys/{fingerprint}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) DeleteGroup(request DeleteGroupRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/groups/{groupId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) DeleteIdentityProvider(request DeleteIdentityProviderRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/identityProviders/{identityProviderId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) DeleteIdpGroupMapping(request DeleteIdpGroupMappingRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) DeletePolicy(request DeletePolicyRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/policies/{policyId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) DeleteSwiftPassword(request DeleteSwiftPasswordRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/users/{userId}/swiftPasswords/{swiftPasswordId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) DeleteUser(request DeleteUserRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/users/{userId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) GetCompartment(request GetCompartmentRequest) (response GetCompartmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/compartments/{compartmentId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) GetGroup(request GetGroupRequest) (response GetGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/groups/{groupId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) GetIdentityProvider(request GetIdentityProviderRequest) (response GetIdentityProviderResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/identityProviders/{identityProviderId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) GetIdpGroupMapping(request GetIdpGroupMappingRequest) (response GetIdpGroupMappingResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) GetPolicy(request GetPolicyRequest) (response GetPolicyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/policies/{policyId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) GetTenancy(request GetTenancyRequest) (response GetTenancyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/tenancies/{tenancyId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) GetUser(request GetUserRequest) (response GetUserResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/users/{userId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) GetUserGroupMembership(request GetUserGroupMembershipRequest) (response GetUserGroupMembershipResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/userGroupMemberships/{userGroupMembershipId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) ListApiKeys(request ListApiKeysRequest) (response ListApiKeysResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/users/{userId}/apiKeys/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) ListAvailabilityDomains(request ListAvailabilityDomainsRequest) (response ListAvailabilityDomainsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/availabilityDomains/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) ListCompartments(request ListCompartmentsRequest) (response ListCompartmentsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/compartments/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) ListGroups(request ListGroupsRequest) (response ListGroupsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/groups/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) ListIdentityProviders(request ListIdentityProvidersRequest) (response ListIdentityProvidersResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/identityProviders/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) ListIdpGroupMappings(request ListIdpGroupMappingsRequest) (response ListIdpGroupMappingsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/identityProviders/{identityProviderId}/groupMappings/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) ListPolicies(request ListPoliciesRequest) (response ListPoliciesResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/policies/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) ListRegionSubscriptions(request ListRegionSubscriptionsRequest) (response ListRegionSubscriptionsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/tenancies/{tenancyId}/regionSubscriptions", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) ListRegions() (response ListRegionsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequest(http.MethodGet, "/regions")
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) ListSwiftPasswords(request ListSwiftPasswordsRequest) (response ListSwiftPasswordsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/users/{userId}/swiftPasswords/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) ListUserGroupMemberships(request ListUserGroupMembershipsRequest) (response ListUserGroupMembershipsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/userGroupMemberships/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) ListUsers(request ListUsersRequest) (response ListUsersResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/users/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) RemoveUserFromGroup(request RemoveUserFromGroupRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/userGroupMemberships/{userGroupMembershipId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) UpdateCompartment(request UpdateCompartmentRequest) (response UpdateCompartmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/compartments/{compartmentId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) UpdateGroup(request UpdateGroupRequest) (response UpdateGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/groups/{groupId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) UpdateIdentityProvider(request UpdateIdentityProviderRequest) (response UpdateIdentityProviderResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/identityProviders/{identityProviderId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) UpdateIdpGroupMapping(request UpdateIdpGroupMappingRequest) (response UpdateIdpGroupMappingResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) UpdatePolicy(request UpdatePolicyRequest) (response UpdatePolicyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/policies/{policyId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) UpdateSwiftPassword(request UpdateSwiftPasswordRequest) (response UpdateSwiftPasswordResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/users/{userId}/swiftPasswords/{swiftPasswordId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) UpdateUser(request UpdateUserRequest) (response UpdateUserResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/users/{userId}", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) UpdateUserState(request UpdateUserStateRequest) (response UpdateUserStateResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/users/{userId}/state/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}

func (client IdentityClient) UploadApiKey(request UploadApiKeyRequest) (response UploadApiKeyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/users/{userId}/apiKeys/", request)
	if err != nil {
		return
	}
	client.Call(httpRequest)
	return
}
