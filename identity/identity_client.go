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
	"net/url"
)

type IdentityClient struct {
}

func (client *IdentityClient) makeHttpCall(r http.Request) {
	common.Debugln("Calling downstream service")
}

func (client *IdentityClient) AddUserToGroup(request AddUserToGroupRequest) (response AddUserToGroupResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/userGroupMemberships/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) CreateCompartment(request CreateCompartmentRequest) (response CreateCompartmentResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/compartments/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) CreateGroup(request CreateGroupRequest) (response CreateGroupResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/groups/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) CreateIdentityProvider(request CreateIdentityProviderRequest) (response CreateIdentityProviderResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/identityProviders/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) CreateIdpGroupMapping(request CreateIdpGroupMappingRequest) (response CreateIdpGroupMappingResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/identityProviders/{identityProviderId}/groupMappings/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) CreateOrResetUIPassword(request CreateOrResetUIPasswordRequest) (response CreateOrResetUIPasswordResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/users/{userId}/uiPassword"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) CreatePolicy(request CreatePolicyRequest) (response CreatePolicyResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/policies/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) CreateRegionSubscription(request CreateRegionSubscriptionRequest) (response CreateRegionSubscriptionResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/tenancies/{tenancyId}/regionSubscriptions"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) CreateSwiftPassword(request CreateSwiftPasswordRequest) (response CreateSwiftPasswordResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/users/{userId}/swiftPasswords/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) CreateUser(request CreateUserRequest) (response CreateUserResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/users/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) DeleteApiKey(request DeleteApiKeyRequest) (err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/users/{userId}/apiKeys/{fingerprint}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) DeleteGroup(request DeleteGroupRequest) (err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/groups/{groupId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) DeleteIdentityProvider(request DeleteIdentityProviderRequest) (err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/identityProviders/{identityProviderId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) DeleteIdpGroupMapping(request DeleteIdpGroupMappingRequest) (err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/identityProviders/{identityProviderId}/groupMappings/{mappingId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) DeletePolicy(request DeletePolicyRequest) (err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/policies/{policyId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) DeleteSwiftPassword(request DeleteSwiftPasswordRequest) (err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/users/{userId}/swiftPasswords/{swiftPasswordId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) DeleteUser(request DeleteUserRequest) (err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/users/{userId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) GetCompartment(request GetCompartmentRequest) (response GetCompartmentResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/compartments/{compartmentId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) GetGroup(request GetGroupRequest) (response GetGroupResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/groups/{groupId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) GetIdentityProvider(request GetIdentityProviderRequest) (response GetIdentityProviderResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/identityProviders/{identityProviderId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) GetIdpGroupMapping(request GetIdpGroupMappingRequest) (response GetIdpGroupMappingResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/identityProviders/{identityProviderId}/groupMappings/{mappingId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) GetPolicy(request GetPolicyRequest) (response GetPolicyResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/policies/{policyId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) GetTenancy(request GetTenancyRequest) (response GetTenancyResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/tenancies/{tenancyId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) GetUser(request GetUserRequest) (response GetUserResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/users/{userId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) GetUserGroupMembership(request GetUserGroupMembershipRequest) (response GetUserGroupMembershipResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/userGroupMemberships/{userGroupMembershipId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) ListApiKeys(request ListApiKeysRequest) (response ListApiKeysResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/users/{userId}/apiKeys/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) ListAvailabilityDomains(request ListAvailabilityDomainsRequest) (response ListAvailabilityDomainsResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/availabilityDomains/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) ListCompartments(request ListCompartmentsRequest) (response ListCompartmentsResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/compartments/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) ListGroups(request ListGroupsRequest) (response ListGroupsResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/groups/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) ListIdentityProviders(request ListIdentityProvidersRequest) (response ListIdentityProvidersResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/identityProviders/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) ListIdpGroupMappings(request ListIdpGroupMappingsRequest) (response ListIdpGroupMappingsResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/identityProviders/{identityProviderId}/groupMappings/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) ListPolicies(request ListPoliciesRequest) (response ListPoliciesResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/policies/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) ListRegionSubscriptions(request ListRegionSubscriptionsRequest) (response ListRegionSubscriptionsResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/tenancies/{tenancyId}/regionSubscriptions"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) ListRegions() (response ListRegionsResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) ListSwiftPasswords(request ListSwiftPasswordsRequest) (response ListSwiftPasswordsResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/users/{userId}/swiftPasswords/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) ListUserGroupMemberships(request ListUserGroupMembershipsRequest) (response ListUserGroupMembershipsResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/userGroupMemberships/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) ListUsers(request ListUsersRequest) (response ListUsersResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/users/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) RemoveUserFromGroup(request RemoveUserFromGroupRequest) (err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/userGroupMemberships/{userGroupMembershipId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) UpdateCompartment(request UpdateCompartmentRequest) (response UpdateCompartmentResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/compartments/{compartmentId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) UpdateGroup(request UpdateGroupRequest) (response UpdateGroupResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/groups/{groupId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) UpdateIdentityProvider(request UpdateIdentityProviderRequest) (response UpdateIdentityProviderResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/identityProviders/{identityProviderId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) UpdateIdpGroupMapping(request UpdateIdpGroupMappingRequest) (response UpdateIdpGroupMappingResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/identityProviders/{identityProviderId}/groupMappings/{mappingId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) UpdatePolicy(request UpdatePolicyRequest) (response UpdatePolicyResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/policies/{policyId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) UpdateSwiftPassword(request UpdateSwiftPasswordRequest) (response UpdateSwiftPasswordResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/users/{userId}/swiftPasswords/{swiftPasswordId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) UpdateUser(request UpdateUserRequest) (response UpdateUserResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/users/{userId}"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) UpdateUserState(request UpdateUserStateRequest) (response UpdateUserStateResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/users/{userId}/state/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}

func (client *IdentityClient) UploadApiKey(request UploadApiKeyRequest) (response UploadApiKeyResponse, err error) {
	httpRequest := http.Request{URL: &url.URL{}}
	httpRequest.URL.Path = "/users/{userId}/apiKeys/"
	err = common.HttpRequestMarshaller(request, &httpRequest)
	if err != nil {
		return
	}
	client.makeHttpCall(httpRequest)
	return
}
