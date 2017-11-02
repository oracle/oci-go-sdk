// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"context"
	"fmt"
	"net/http"
)

type IdentityClient struct {
	common.BaseClient
}

//Create a new default Identity client for a given region
func NewClientForRegion(region common.Region) (client IdentityClient) {
	client = IdentityClient{BaseClient: common.NewClientForRegion(region)}

	client.Host = fmt.Sprintf(common.DefaultHostUrlTemplate, "identity", string(region))
	client.BasePath = "20160918"
	return
}

func closeIfValid(httpResponse *http.Response) {
	if httpResponse != nil && httpResponse.Body != nil {
		httpResponse.Body.Close()
	}
}

// Adds the specified user to the specified group and returns a `UserGroupMembership` object with its own OCID.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
func (client IdentityClient) AddUserToGroup(ctx context.Context, request AddUserToGroupRequest) (response AddUserToGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/userGroupMemberships/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Creates a new compartment in your tenancy.
// **Important:** Compartments cannot be deleted.
// You must specify your tenancy's OCID as the compartment ID in the request object. Remember that the tenancy
// is simply the root compartment. For information about OCIDs, see
// [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You must also specify a *name* for the compartment, which must be unique across all compartments in
// your tenancy. You can use this name or the OCID when writing policies that apply
// to the compartment. For more information about policies, see
// [How Policies Work]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policies.htm).
// You must also specify a *description* for the compartment (although it can be an empty string). It does
// not have to be unique, and you can change it anytime with
// UpdateCompartment.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
func (client IdentityClient) CreateCompartment(ctx context.Context, request CreateCompartmentRequest) (response CreateCompartmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/compartments/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Creates a new secret key for the specified user. Secret keys are used for authentication with the Object Storage Service's Amazon S3
// compatible API. For information, see
// [Managing User Credentials]({{DOC_SERVER_URL}}/Content/Identity/Tasks/managingcredentials.htm).
// You must specify a *description* for the secret key (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with
// UpdateCustomerSecretKey.
// Every user has permission to create a secret key for *their own user ID*. An administrator in your organization
// does not need to write a policy to give users this ability. To compare, administrators who have permission to the
// tenancy can use this operation to create a secret key for any user, including themselves.
func (client IdentityClient) CreateCustomerSecretKey(ctx context.Context, request CreateCustomerSecretKeyRequest) (response CreateCustomerSecretKeyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/users/{userId}/customerSecretKeys/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
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
func (client IdentityClient) CreateGroup(ctx context.Context, request CreateGroupRequest) (response CreateGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/groups/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
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
func (client IdentityClient) CreateIdentityProvider(ctx context.Context, request CreateIdentityProviderRequest) (response CreateIdentityProviderResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/identityProviders/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Creates a single mapping between an IdP group and an IAM Service
// Group.
func (client IdentityClient) CreateIdpGroupMapping(ctx context.Context, request CreateIdpGroupMappingRequest) (response CreateIdpGroupMappingResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/identityProviders/{identityProviderId}/groupMappings/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
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
func (client IdentityClient) CreateOrResetUIPassword(ctx context.Context, request CreateOrResetUIPasswordRequest) (response CreateOrResetUIPasswordResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/users/{userId}/uiPassword", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
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
func (client IdentityClient) CreatePolicy(ctx context.Context, request CreatePolicyRequest) (response CreatePolicyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/policies/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Creates a subscription to a region for a tenancy.
func (client IdentityClient) CreateRegionSubscription(ctx context.Context, request CreateRegionSubscriptionRequest) (response CreateRegionSubscriptionResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/tenancies/{tenancyId}/regionSubscriptions", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
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
func (client IdentityClient) CreateSwiftPassword(ctx context.Context, request CreateSwiftPasswordRequest) (response CreateSwiftPasswordResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/users/{userId}/swiftPasswords/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
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
// If the user needs to access the Oracle Cloud Infrastructure REST API, you need to upload a
// public API signing key for that user (see
// [Required Keys and OCIDs]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm) and also
// UploadApiKey).
// **Important:** Make sure to inform the new user which compartment(s) they have access to.
func (client IdentityClient) CreateUser(ctx context.Context, request CreateUserRequest) (response CreateUserResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/users/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Deletes the specified API signing key for the specified user.
// Every user has permission to use this operation to delete a key for *their own user ID*. An
// administrator in your organization does not need to write a policy to give users this ability.
// To compare, administrators who have permission to the tenancy can use this operation to delete
// a key for any user, including themselves.
func (client IdentityClient) DeleteApiKey(ctx context.Context, request DeleteApiKeyRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/users/{userId}/apiKeys/{fingerprint}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Deletes the specified secret key for the specified user.
func (client IdentityClient) DeleteCustomerSecretKey(ctx context.Context, request DeleteCustomerSecretKeyRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/users/{userId}/customerSecretKeys/{customerSecretKeyId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Deletes the specified group. The group must be empty.
func (client IdentityClient) DeleteGroup(ctx context.Context, request DeleteGroupRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/groups/{groupId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Deletes the specified identity provider. The identity provider must not have
// any group mappings (see IdpGroupMapping).
func (client IdentityClient) DeleteIdentityProvider(ctx context.Context, request DeleteIdentityProviderRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/identityProviders/{identityProviderId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Deletes the specified group mapping.
func (client IdentityClient) DeleteIdpGroupMapping(ctx context.Context, request DeleteIdpGroupMappingRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Deletes the specified policy. The deletion takes effect typically within 10 seconds.
func (client IdentityClient) DeletePolicy(ctx context.Context, request DeletePolicyRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/policies/{policyId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Deletes the specified Swift password for the specified user.
func (client IdentityClient) DeleteSwiftPassword(ctx context.Context, request DeleteSwiftPasswordRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/users/{userId}/swiftPasswords/{swiftPasswordId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Deletes the specified user. The user must not be in any groups.
func (client IdentityClient) DeleteUser(ctx context.Context, request DeleteUserRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/users/{userId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Gets the specified compartment's information.
// This operation does not return a list of all the resources inside the compartment. There is no single
// API operation that does that. Compartments can contain multiple types of resources (instances, block
// storage volumes, etc.). To find out what's in a compartment, you must call the "List" operation for
// each resource type and specify the compartment's OCID as a query parameter in the request. For example,
// call the ListInstances operation in the Cloud Compute
// Service or the ListVolumes operation in Cloud Block Storage.
func (client IdentityClient) GetCompartment(ctx context.Context, request GetCompartmentRequest) (response GetCompartmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/compartments/{compartmentId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets the specified group's information.
// This operation does not return a list of all the users in the group. To do that, use
// ListUserGroupMemberships and
// provide the group's OCID as a query parameter in the request.
func (client IdentityClient) GetGroup(ctx context.Context, request GetGroupRequest) (response GetGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/groups/{groupId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets the specified identity provider's information.
func (client IdentityClient) GetIdentityProvider(ctx context.Context, request GetIdentityProviderRequest) (response GetIdentityProviderResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/identityProviders/{identityProviderId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets the specified group mapping.
func (client IdentityClient) GetIdpGroupMapping(ctx context.Context, request GetIdpGroupMappingRequest) (response GetIdpGroupMappingResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets the specified policy's information.
func (client IdentityClient) GetPolicy(ctx context.Context, request GetPolicyRequest) (response GetPolicyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/policies/{policyId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Get the specified tenancy's information.
func (client IdentityClient) GetTenancy(ctx context.Context, request GetTenancyRequest) (response GetTenancyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/tenancies/{tenancyId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets the specified user's information.
func (client IdentityClient) GetUser(ctx context.Context, request GetUserRequest) (response GetUserResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/users/{userId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Gets the specified UserGroupMembership's information.
func (client IdentityClient) GetUserGroupMembership(ctx context.Context, request GetUserGroupMembershipRequest) (response GetUserGroupMembershipResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/userGroupMemberships/{userGroupMembershipId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the API signing keys for the specified user. A user can have a maximum of three keys.
// Every user has permission to use this API call for *their own user ID*.  An administrator in your
// organization does not need to write a policy to give users this ability.
func (client IdentityClient) ListApiKeys(ctx context.Context, request ListApiKeysRequest) (response ListApiKeysResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/users/{userId}/apiKeys/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the Availability Domains in your tenancy. Specify the OCID of either the tenancy or another
// of your compartments as the value for the compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListAvailabilityDomains(ctx context.Context, request ListAvailabilityDomainsRequest) (response ListAvailabilityDomainsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/availabilityDomains/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the compartments in your tenancy. You must specify your tenancy's OCID as the value
// for the compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListCompartments(ctx context.Context, request ListCompartmentsRequest) (response ListCompartmentsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/compartments/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the secret keys for the specified user. The returned object contains the secret key's OCID, but not
// the secret key itself. The actual secret key is returned only upon creation.
func (client IdentityClient) ListCustomerSecretKeys(ctx context.Context, request ListCustomerSecretKeysRequest) (response ListCustomerSecretKeysResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/users/{userId}/customerSecretKeys/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the groups in your tenancy. You must specify your tenancy's OCID as the value for
// the compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListGroups(ctx context.Context, request ListGroupsRequest) (response ListGroupsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/groups/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists all the identity providers in your tenancy. You must specify the identity provider type (e.g., `SAML2` for
// identity providers using the SAML2.0 protocol). You must specify your tenancy's OCID as the value for the
// compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListIdentityProviders(ctx context.Context, request ListIdentityProvidersRequest) (response ListIdentityProvidersResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/identityProviders/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the group mappings for the specified identity provider.
func (client IdentityClient) ListIdpGroupMappings(ctx context.Context, request ListIdpGroupMappingsRequest) (response ListIdpGroupMappingsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/identityProviders/{identityProviderId}/groupMappings/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the policies in the specified compartment (either the tenancy or another of your compartments).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
// To determine which policies apply to a particular group or compartment, you must view the individual
// statements inside all your policies. There isn't a way to automatically obtain that information via the API.
func (client IdentityClient) ListPolicies(ctx context.Context, request ListPoliciesRequest) (response ListPoliciesResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/policies/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the region subscriptions for the specified tenancy.
func (client IdentityClient) ListRegionSubscriptions(ctx context.Context, request ListRegionSubscriptionsRequest) (response ListRegionSubscriptionsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/tenancies/{tenancyId}/regionSubscriptions", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists all the regions offered by Oracle Cloud Infrastructure.
func (client IdentityClient) ListRegions(ctx context.Context) (response ListRegionsResponse, err error) {
	httpRequest := common.MakeDefaultHttpRequest(http.MethodGet, "/regions")

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the Swift passwords for the specified user. The returned object contains the password's OCID, but not
// the password itself. The actual password is returned only upon creation.
func (client IdentityClient) ListSwiftPasswords(ctx context.Context, request ListSwiftPasswordsRequest) (response ListSwiftPasswordsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/users/{userId}/swiftPasswords/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
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
func (client IdentityClient) ListUserGroupMemberships(ctx context.Context, request ListUserGroupMembershipsRequest) (response ListUserGroupMembershipsResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/userGroupMemberships/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Lists the users in your tenancy. You must specify your tenancy's OCID as the value for the
// compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListUsers(ctx context.Context, request ListUsersRequest) (response ListUsersResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/users/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Removes a user from a group by deleting the corresponding `UserGroupMembership`.
func (client IdentityClient) RemoveUserFromGroup(ctx context.Context, request RemoveUserFromGroupRequest) (err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodDelete, "/userGroupMemberships/{userGroupMembershipId}", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}

// Updates the specified compartment's description or name. You can't update the root compartment.
func (client IdentityClient) UpdateCompartment(ctx context.Context, request UpdateCompartmentRequest) (response UpdateCompartmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/compartments/{compartmentId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Updates the specified secret key's description.
func (client IdentityClient) UpdateCustomerSecretKey(ctx context.Context, request UpdateCustomerSecretKeyRequest) (response UpdateCustomerSecretKeyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/users/{userId}/customerSecretKeys/{customerSecretKeyId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Updates the specified group.
func (client IdentityClient) UpdateGroup(ctx context.Context, request UpdateGroupRequest) (response UpdateGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/groups/{groupId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Updates the specified identity provider.
func (client IdentityClient) UpdateIdentityProvider(ctx context.Context, request UpdateIdentityProviderRequest) (response UpdateIdentityProviderResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/identityProviders/{identityProviderId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Updates the specified group mapping.
func (client IdentityClient) UpdateIdpGroupMapping(ctx context.Context, request UpdateIdpGroupMappingRequest) (response UpdateIdpGroupMappingResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Updates the specified policy. You can update the description or the policy statements themselves.
// Policy changes take effect typically within 10 seconds.
func (client IdentityClient) UpdatePolicy(ctx context.Context, request UpdatePolicyRequest) (response UpdatePolicyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/policies/{policyId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Updates the specified Swift password's description.
func (client IdentityClient) UpdateSwiftPassword(ctx context.Context, request UpdateSwiftPasswordRequest) (response UpdateSwiftPasswordResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/users/{userId}/swiftPasswords/{swiftPasswordId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Updates the description of the specified user.
func (client IdentityClient) UpdateUser(ctx context.Context, request UpdateUserRequest) (response UpdateUserResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/users/{userId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// Updates the state of the specified user.
func (client IdentityClient) UpdateUserState(ctx context.Context, request UpdateUserStateRequest) (response UpdateUserStateResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/users/{userId}/state/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
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
func (client IdentityClient) UploadApiKey(ctx context.Context, request UploadApiKeyRequest) (response UploadApiKeyResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPost, "/users/{userId}/apiKeys/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer closeIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}
