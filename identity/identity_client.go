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
	common.BaseClient
}

func NewClient() (client IdentityClient) {
	client = IdentityClient{BaseClient: common.NewClient()}
	client.ServiceName = "identity"
	client.ApiVersion = "20160918"
	return
}

// Adds the specified user to the specified group and returns a `UserGroupMembership` object with its own OCID.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
func AddUserToGroup(client common.Client, request AddUserToGroupRequest) (response AddUserToGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/userGroupMemberships/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Creates a new compartment in your tenancy.
// **Important:** Compartments cannot be renamed or deleted.
// You must specify your tenancy's OCID as the compartment ID in the request object. Remember that the tenancy
// is simply the root compartment. For information about OCIDs, see
// [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You must also specify a *name* for the compartment, which must be unique across all compartments in
// your tenancy and cannot be changed. You can use this name or the OCID when writing policies that apply
// to the compartment. For more information about policies, see
// [How Policies Work]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policies.htm).
// You must also specify a *description* for the compartment (although it can be an empty string). It does
// not have to be unique, and you can change it anytime with
// UpdateCompartment.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
func CreateCompartment(client common.Client, request CreateCompartmentRequest) (response CreateCompartmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/compartments/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Creates a new group in your tenancy.
// You must specify your tenancy's OCID as the compartment ID in the request object (remember that the tenancy
// is simply the root compartment). Notice that IAM resources (users, groups, compartments, and some policies)
// reside within the tenancy itself, unlike cloud resources such as compute instances, which typically
// reside within compartments inside the tenancy. For information about OCIDs, see
// [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You must also specify a *name* for the group, which must be unique across all groups in your tenancy and
// cannot be changed. You can use this name or the OCID when writing policies that apply to the group. For more
// information about policies, see [How Policies Work]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policies.htm).
// You must also specify a *description* for the group (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with UpdateGroup.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
// After creating the group, you need to put users in it and write policies for it.
// See AddUserToGroup and
// CreatePolicy.
func CreateGroup(client common.Client, request CreateGroupRequest) (response CreateGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/groups/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Creates a new identity provider in your tenancy. For more information, see
// [Identity Providers and Federation]({{DOC_SERVER_URL}}/Content/Identity/Concepts/federation.htm).
// You must specify your tenancy's OCID as the compartment ID in the request object.
// Remember that the tenancy is simply the root compartment. For information about
// OCIDs, see [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You must also specify a *name* for the `IdentityProvider`, which must be unique
// across all `IdentityProvider` objects in your tenancy and cannot be changed.
// You must also specify a *description* for the `IdentityProvider` (although
// it can be an empty string). It does not have to be unique, and you can change
// it anytime with
// UpdateIdentityProvider.
// After you send your request, the new object's `lifecycleState` will temporarily
// be CREATING. Before using the object, first make sure its `lifecycleState` has
// changed to ACTIVE.
func CreateIdentityProvider(client common.Client, request CreateIdentityProviderRequest) (response CreateIdentityProviderResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/identityProviders/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Creates a single mapping between an IdP group and an IAM Service
// Group.
func CreateIdpGroupMapping(client common.Client, request CreateIdpGroupMappingRequest) (response CreateIdpGroupMappingResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/identityProviders/{identityProviderId}/groupMappings/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Creates a new Console one-time password for the specified user. For more information about user
// credentials, see [User Credentials]({{DOC_SERVER_URL}}/Content/Identity/Concepts/usercredentials.htm).
// Use this operation after creating a new user, or if a user forgets their password. The new one-time
// password is returned to you in the response, and you must securely deliver it to the user. They'll
// be prompted to change this password the next time they sign in to the Console. If they don't change
// it within 7 days, the password will expire and you'll need to create a new one-time password for the
// user.
// **Note:** The user's Console login is the unique name you specified when you created the user
// (see CreateUser).
func CreateOrResetUIPassword(client common.Client, request CreateOrResetUIPasswordRequest) (response CreateOrResetUIPasswordResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/users/{userId}/uiPassword", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Creates a new policy in the specified compartment (either the tenancy or another of your compartments).
// If you're new to policies, see [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
// You must specify a *name* for the policy, which must be unique across all policies in your tenancy
// and cannot be changed.
// You must also specify a *description* for the policy (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with UpdatePolicy.
// You must specify one or more policy statements in the statements array. For information about writing
// policies, see [How Policies Work]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policies.htm) and
// [Common Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/commonpolicies.htm).
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
// New policies take effect typically within 10 seconds.
func CreatePolicy(client common.Client, request CreatePolicyRequest) (response CreatePolicyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/policies/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Creates a subscription to a region for a tenancy.
func CreateRegionSubscription(client common.Client, request CreateRegionSubscriptionRequest) (response CreateRegionSubscriptionResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/tenancies/{tenancyId}/regionSubscriptions", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Creates a new Swift password for the specified user. For information about what Swift passwords are for, see
// [Managing User Credentials]({{DOC_SERVER_URL}}/Content/Identity/Tasks/managingcredentials.htm).
// You must specify a *description* for the Swift password (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with
// UpdateSwiftPassword.
// Every user has permission to create a Swift password for *their own user ID*. An administrator in your organization
// does not need to write a policy to give users this ability. To compare, administrators who have permission to the
// tenancy can use this operation to create a Swift password for any user, including themselves.
func CreateSwiftPassword(client common.Client, request CreateSwiftPasswordRequest) (response CreateSwiftPasswordResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/users/{userId}/swiftPasswords/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Creates a new user in your tenancy. For conceptual information about users, your tenancy, and other
// IAM Service components, see [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm).
// You must specify your tenancy's OCID as the compartment ID in the request object (remember that the
// tenancy is simply the root compartment). Notice that IAM resources (users, groups, compartments, and
// some policies) reside within the tenancy itself, unlike cloud resources such as compute instances,
// which typically reside within compartments inside the tenancy. For information about OCIDs, see
// [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You must also specify a *name* for the user, which must be unique across all users in your tenancy
// and cannot be changed. Allowed characters: No spaces. Only letters, numerals, hyphens, periods,
// underscores, +, and @. If you specify a name that's already in use, you'll get a 409 error.
// This name will be the user's login to the Console. You might want to pick a
// name that your company's own identity system (e.g., Active Directory, LDAP, etc.) already uses.
// If you delete a user and then create a new user with the same name, they'll be considered different
// users because they have different OCIDs.
// You must also specify a *description* for the user (although it can be an empty string).
// It does not have to be unique, and you can change it anytime with
// UpdateUser. You can use the field to provide the user's
// full name, a description, a nickname, or other information to generally identify the user.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before
// using the object, first make sure its `lifecycleState` has changed to ACTIVE.
// A new user has no permissions until you place the user in one or more groups (see
// AddUserToGroup). If the user needs to
// access the Console, you need to provide the user a password (see
// CreateOrResetUIPassword).
// If the user needs to access the Oracle Bare Metal Cloud Services REST API, you need to upload a
// public API signing key for that user (see
// [Required Keys and OCIDs]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm) and also
// UploadApiKey).
// **Important:** Make sure to inform the new user which compartment(s) they have access to.
func CreateUser(client common.Client, request CreateUserRequest) (response CreateUserResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/users/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Deletes the specified API signing key for the specified user.
// Every user has permission to use this operation to delete a key for *their own user ID*. An
// administrator in your organization does not need to write a policy to give users this ability.
// To compare, administrators who have permission to the tenancy can use this operation to delete
// a key for any user, including themselves.
func DeleteApiKey(client common.Client, request DeleteApiKeyRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/users/{userId}/apiKeys/{fingerprint}", request)
	if err != nil {
		return
	}

	_, err = client.Call(httpRequest)
	if err != nil {
		return
	}

	return
}

// Deletes the specified group. The group must be empty.
func DeleteGroup(client common.Client, request DeleteGroupRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/groups/{groupId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(httpRequest)
	if err != nil {
		return
	}

	return
}

// Deletes the specified identity provider. The identity provider must not have
// any group mappings (see IdpGroupMapping).
func DeleteIdentityProvider(client common.Client, request DeleteIdentityProviderRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/identityProviders/{identityProviderId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(httpRequest)
	if err != nil {
		return
	}

	return
}

// Deletes the specified group mapping.
func DeleteIdpGroupMapping(client common.Client, request DeleteIdpGroupMappingRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(httpRequest)
	if err != nil {
		return
	}

	return
}

// Deletes the specified policy. The deletion takes effect typically within 10 seconds.
func DeletePolicy(client common.Client, request DeletePolicyRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/policies/{policyId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(httpRequest)
	if err != nil {
		return
	}

	return
}

// Deletes the specified Swift password for the specified user.
func DeleteSwiftPassword(client common.Client, request DeleteSwiftPasswordRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/users/{userId}/swiftPasswords/{swiftPasswordId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(httpRequest)
	if err != nil {
		return
	}

	return
}

// Deletes the specified user. The user must not be in any groups.
func DeleteUser(client common.Client, request DeleteUserRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/users/{userId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(httpRequest)
	if err != nil {
		return
	}

	return
}

// Gets the specified compartment's information.
// This operation does not return a list of all the resources inside the compartment. There is no single
// API operation that does that. Compartments can contain multiple types of resources (instances, block
// storage volumes, etc.). To find out what's in a compartment, you must call the "List" operation for
// each resource type and specify the compartment's OCID as a query parameter in the request. For example,
// call the ListInstances operation in the Cloud Compute
// Service or the ListVolumes operation in Cloud Block Storage.
func GetCompartment(client common.Client, request GetCompartmentRequest) (response GetCompartmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/compartments/{compartmentId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Gets the specified group's information.
// This operation does not return a list of all the users in the group. To do that, use
// ListUserGroupMemberships and
// provide the group's OCID as a query parameter in the request.
func GetGroup(client common.Client, request GetGroupRequest) (response GetGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/groups/{groupId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Gets the specified identity provider's information.
func GetIdentityProvider(client common.Client, request GetIdentityProviderRequest) (response GetIdentityProviderResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/identityProviders/{identityProviderId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Gets the specified group mapping.
func GetIdpGroupMapping(client common.Client, request GetIdpGroupMappingRequest) (response GetIdpGroupMappingResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Gets the specified policy's information.
func GetPolicy(client common.Client, request GetPolicyRequest) (response GetPolicyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/policies/{policyId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Get the specified tenancy's information.
func GetTenancy(client common.Client, request GetTenancyRequest) (response GetTenancyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/tenancies/{tenancyId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Gets the specified user's information.
func GetUser(client common.Client, request GetUserRequest) (response GetUserResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/users/{userId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Gets the specified UserGroupMembership's information.
func GetUserGroupMembership(client common.Client, request GetUserGroupMembershipRequest) (response GetUserGroupMembershipResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/userGroupMemberships/{userGroupMembershipId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Lists the API signing keys for the specified user. A user can have a maximum of three keys.
// Every user has permission to use this API call for *their own user ID*.  An administrator in your
// organization does not need to write a policy to give users this ability.
func ListApiKeys(client common.Client, request ListApiKeysRequest) (response ListApiKeysResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/users/{userId}/apiKeys/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Lists the Availability Domains in your tenancy. Specify the OCID of either the tenancy or another
// of your compartments as the value for the compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func ListAvailabilityDomains(client common.Client, request ListAvailabilityDomainsRequest) (response ListAvailabilityDomainsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/availabilityDomains/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Lists the compartments in your tenancy. You must specify your tenancy's OCID as the value
// for the compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func ListCompartments(client common.Client, request ListCompartmentsRequest) (response ListCompartmentsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/compartments/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Lists the groups in your tenancy. You must specify your tenancy's OCID as the value for
// the compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func ListGroups(client common.Client, request ListGroupsRequest) (response ListGroupsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/groups/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Lists all the identity providers in your tenancy. You must specify the identity provider type (e.g., `SAML2` for
// identity providers using the SAML2.0 protocol). You must specify your tenancy's OCID as the value for the
// compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func ListIdentityProviders(client common.Client, request ListIdentityProvidersRequest) (response ListIdentityProvidersResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/identityProviders/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Lists the group mappings for the specified identity provider.
func ListIdpGroupMappings(client common.Client, request ListIdpGroupMappingsRequest) (response ListIdpGroupMappingsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/identityProviders/{identityProviderId}/groupMappings/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Lists the policies in the specified compartment (either the tenancy or another of your compartments).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
// To determine which policies apply to a particular group or compartment, you must view the individual
// statements inside all your policies. There isn't a way to automatically obtain that information via the API.
func ListPolicies(client common.Client, request ListPoliciesRequest) (response ListPoliciesResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/policies/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Lists the region subscriptions for the specified tenancy.
func ListRegionSubscriptions(client common.Client, request ListRegionSubscriptionsRequest) (response ListRegionSubscriptionsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/tenancies/{tenancyId}/regionSubscriptions", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Lists all the regions offered by Oracle Bare Metal Cloud Services.
func ListRegions(client common.Client) (response ListRegionsResponse, err error) {
	httpRequest := common.MakeDefaultHttpRequest(http.MethodGet, "/regions")

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Lists the Swift passwords for the specified user. The returned object contains the password's OCID, but not
// the password itself. The actual password is returned only upon creation.
func ListSwiftPasswords(client common.Client, request ListSwiftPasswordsRequest) (response ListSwiftPasswordsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/users/{userId}/swiftPasswords/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Lists the `UserGroupMembership` objects in your tenancy. You must specify your tenancy's OCID
// as the value for the compartment ID
// (see [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five)).
// You must also then filter the list in one of these ways:
// - You can limit the results to just the memberships for a given user by specifying a `userId`.
// - Similarly, you can limit the results to just the memberships for a given group by specifying a `groupId`.
// - You can set both the `userId` and `groupId` to determine if the specified user is in the specified group.
// If the answer is no, the response is an empty list.
func ListUserGroupMemberships(client common.Client, request ListUserGroupMembershipsRequest) (response ListUserGroupMembershipsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/userGroupMemberships/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Lists the users in your tenancy. You must specify your tenancy's OCID as the value for the
// compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func ListUsers(client common.Client, request ListUsersRequest) (response ListUsersResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/users/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Removes a user from a group by deleting the corresponding `UserGroupMembership`.
func RemoveUserFromGroup(client common.Client, request RemoveUserFromGroupRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/userGroupMemberships/{userGroupMembershipId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(httpRequest)
	if err != nil {
		return
	}

	return
}

// Updates the specified compartment's description.
func UpdateCompartment(client common.Client, request UpdateCompartmentRequest) (response UpdateCompartmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/compartments/{compartmentId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Updates the specified group.
func UpdateGroup(client common.Client, request UpdateGroupRequest) (response UpdateGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/groups/{groupId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Updates the specified identity provider.
func UpdateIdentityProvider(client common.Client, request UpdateIdentityProviderRequest) (response UpdateIdentityProviderResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/identityProviders/{identityProviderId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Updates the specified group mapping.
func UpdateIdpGroupMapping(client common.Client, request UpdateIdpGroupMappingRequest) (response UpdateIdpGroupMappingResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Updates the specified policy. You can update the description or the policy statements themselves.
// Policy changes take effect typically within 10 seconds.
func UpdatePolicy(client common.Client, request UpdatePolicyRequest) (response UpdatePolicyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/policies/{policyId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Updates the specified Swift password's description.
func UpdateSwiftPassword(client common.Client, request UpdateSwiftPasswordRequest) (response UpdateSwiftPasswordResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/users/{userId}/swiftPasswords/{swiftPasswordId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Updates the description of the specified user.
func UpdateUser(client common.Client, request UpdateUserRequest) (response UpdateUserResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/users/{userId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Updates the state of the specified user.
func UpdateUserState(client common.Client, request UpdateUserStateRequest) (response UpdateUserStateResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/users/{userId}/state/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}

// Uploads an API signing key for the specified user.
// Every user has permission to use this operation to upload a key for *their own user ID*. An
// administrator in your organization does not need to write a policy to give users this ability.
// To compare, administrators who have permission to the tenancy can use this operation to upload a
// key for any user, including themselves.
// **Important:** Even though you have permission to upload an API key, you might not yet
// have permission to do much else. If you try calling an operation unrelated to your own credential
// management (e.g., `ListUsers`, `LaunchInstance`) and receive an "unauthorized" error,
// check with an administrator to confirm which IAM Service group(s) you're in and what access
// you have. Also confirm you're working in the correct compartment.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using
// the object, first make sure its `lifecycleState` has changed to ACTIVE.
func UploadApiKey(client common.Client, request UploadApiKeyRequest) (response UploadApiKeyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/users/{userId}/apiKeys/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(httpRequest)
	if err != nil {
		return
	}

	response.RawResponse = *httpResponse
	err = common.UnmarshalResponse(httpResponse, &response)
	if err != nil {
		return
	}
	return
}
