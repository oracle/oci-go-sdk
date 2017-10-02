// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

type Identity interface {

	// Adds the specified user to the specified group and returns a `UserGroupMembership` object with its own OCID.
	// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
	// object, first make sure its `lifecycleState` has changed to ACTIVE.
	AddUserToGroup(request AddUserToGroupRequest) (response AddUserToGroupResponse, err error)

	// Creates a new compartment in your tenancy.
	// **Important:** Compartments cannot be renamed or deleted.
	// You must specify your tenancy's OCID as the compartment ID in the request object. Remember that the tenancy
	// is simply the root compartment. For information about OCIDs, see
	// [Resource Identifiers](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/General/Concepts/identifiers.htm).
	// You must also specify a *name* for the compartment, which must be unique across all compartments in
	// your tenancy and cannot be changed. You can use this name or the OCID when writing policies that apply
	// to the compartment. For more information about policies, see
	// [How Policies Work](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/Identity/Concepts/policies.htm).
	// You must also specify a *description* for the compartment (although it can be an empty string). It does
	// not have to be unique, and you can change it anytime with
	// UpdateCompartment.
	// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
	// object, first make sure its `lifecycleState` has changed to ACTIVE.
	CreateCompartment(request CreateCompartmentRequest) (response CreateCompartmentResponse, err error)

	// Creates a new group in your tenancy.
	// You must specify your tenancy's OCID as the compartment ID in the request object (remember that the tenancy
	// is simply the root compartment). Notice that IAM resources (users, groups, compartments, and some policies)
	// reside within the tenancy itself, unlike cloud resources such as compute instances, which typically
	// reside within compartments inside the tenancy. For information about OCIDs, see
	// [Resource Identifiers](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/General/Concepts/identifiers.htm).
	// You must also specify a *name* for the group, which must be unique across all groups in your tenancy and
	// cannot be changed. You can use this name or the OCID when writing policies that apply to the group. For more
	// information about policies, see [How Policies Work](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/Identity/Concepts/policies.htm).
	// You must also specify a *description* for the group (although it can be an empty string). It does not
	// have to be unique, and you can change it anytime with UpdateGroup.
	// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
	// object, first make sure its `lifecycleState` has changed to ACTIVE.
	// After creating the group, you need to put users in it and write policies for it.
	// See AddUserToGroup and
	// CreatePolicy.
	CreateGroup(request CreateGroupRequest) (response CreateGroupResponse, err error)

	// Creates a new identity provider in your tenancy. For more information, see
	// [Identity Providers and Federation](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/Identity/Concepts/federation.htm).
	// You must specify your tenancy's OCID as the compartment ID in the request object.
	// Remember that the tenancy is simply the root compartment. For information about
	// OCIDs, see [Resource Identifiers](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/General/Concepts/identifiers.htm).
	// You must also specify a *name* for the `IdentityProvider`, which must be unique
	// across all `IdentityProvider` objects in your tenancy and cannot be changed.
	// You must also specify a *description* for the `IdentityProvider` (although
	// it can be an empty string). It does not have to be unique, and you can change
	// it anytime with
	// UpdateIdentityProvider.
	// After you send your request, the new object's `lifecycleState` will temporarily
	// be CREATING. Before using the object, first make sure its `lifecycleState` has
	// changed to ACTIVE.
	CreateIdentityProvider(request CreateIdentityProviderRequest) (response CreateIdentityProviderResponse, err error)

	// Creates a single mapping between an IdP group and an IAM Service
	// Group.
	CreateIdpGroupMapping(request CreateIdpGroupMappingRequest) (response CreateIdpGroupMappingResponse, err error)

	// Creates a new Console one-time password for the specified user. For more information about user
	// credentials, see [User Credentials](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/Identity/Concepts/usercredentials.htm).
	// Use this operation after creating a new user, or if a user forgets their password. The new one-time
	// password is returned to you in the response, and you must securely deliver it to the user. They'll
	// be prompted to change this password the next time they sign in to the Console. If they don't change
	// it within 7 days, the password will expire and you'll need to create a new one-time password for the
	// user.
	// **Note:** The user's Console login is the unique name you specified when you created the user
	// (see CreateUser).
	CreateOrResetUIPassword(request CreateOrResetUIPasswordRequest) (response CreateOrResetUIPasswordResponse, err error)

	// Creates a new policy in the specified compartment (either the tenancy or another of your compartments).
	// If you're new to policies, see [Getting Started with Policies](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/Identity/Concepts/policygetstarted.htm).
	// You must specify a *name* for the policy, which must be unique across all policies in your tenancy
	// and cannot be changed.
	// You must also specify a *description* for the policy (although it can be an empty string). It does not
	// have to be unique, and you can change it anytime with UpdatePolicy.
	// You must specify one or more policy statements in the statements array. For information about writing
	// policies, see [How Policies Work](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/Identity/Concepts/policies.htm) and
	// [Common Policies](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/Identity/Concepts/commonpolicies.htm).
	// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
	// object, first make sure its `lifecycleState` has changed to ACTIVE.
	// New policies take effect typically within 10 seconds.
	CreatePolicy(request CreatePolicyRequest) (response CreatePolicyResponse, err error)

	// Creates a subscription to a region for a tenancy.
	CreateRegionSubscription(request CreateRegionSubscriptionRequest) (response CreateRegionSubscriptionResponse, err error)

	// Creates a new Swift password for the specified user. For information about what Swift passwords are for, see
	// [Managing User Credentials](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/Identity/Tasks/managingcredentials.htm).
	// You must specify a *description* for the Swift password (although it can be an empty string). It does not
	// have to be unique, and you can change it anytime with
	// UpdateSwiftPassword.
	// Every user has permission to create a Swift password for *their own user ID*. An administrator in your organization
	// does not need to write a policy to give users this ability. To compare, administrators who have permission to the
	// tenancy can use this operation to create a Swift password for any user, including themselves.
	CreateSwiftPassword(request CreateSwiftPasswordRequest) (response CreateSwiftPasswordResponse, err error)

	// Creates a new user in your tenancy. For conceptual information about users, your tenancy, and other
	// IAM Service components, see [Overview of the IAM Service](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/Identity/Concepts/overview.htm).
	// You must specify your tenancy's OCID as the compartment ID in the request object (remember that the
	// tenancy is simply the root compartment). Notice that IAM resources (users, groups, compartments, and
	// some policies) reside within the tenancy itself, unlike cloud resources such as compute instances,
	// which typically reside within compartments inside the tenancy. For information about OCIDs, see
	// [Resource Identifiers](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/General/Concepts/identifiers.htm).
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
	// [Required Keys and OCIDs](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/API/Concepts/apisigningkey.htm) and also
	// UploadApiKey).
	// **Important:** Make sure to inform the new user which compartment(s) they have access to.
	CreateUser(request CreateUserRequest) (response CreateUserResponse, err error)

	// Deletes the specified API signing key for the specified user.
	// Every user has permission to use this operation to delete a key for *their own user ID*. An
	// administrator in your organization does not need to write a policy to give users this ability.
	// To compare, administrators who have permission to the tenancy can use this operation to delete
	// a key for any user, including themselves.
	DeleteApiKey(request DeleteApiKeyRequest) (err error)

	// Deletes the specified group. The group must be empty.
	DeleteGroup(request DeleteGroupRequest) (err error)

	// Deletes the specified identity provider. The identity provider must not have
	// any group mappings (see IdpGroupMapping).
	DeleteIdentityProvider(request DeleteIdentityProviderRequest) (err error)

	// Deletes the specified group mapping.
	DeleteIdpGroupMapping(request DeleteIdpGroupMappingRequest) (err error)

	// Deletes the specified policy. The deletion takes effect typically within 10 seconds.
	DeletePolicy(request DeletePolicyRequest) (err error)

	// Deletes the specified Swift password for the specified user.
	DeleteSwiftPassword(request DeleteSwiftPasswordRequest) (err error)

	// Deletes the specified user. The user must not be in any groups.
	DeleteUser(request DeleteUserRequest) (err error)

	// Gets the specified compartment's information.
	// This operation does not return a list of all the resources inside the compartment. There is no single
	// API operation that does that. Compartments can contain multiple types of resources (instances, block
	// storage volumes, etc.). To find out what's in a compartment, you must call the "List" operation for
	// each resource type and specify the compartment's OCID as a query parameter in the request. For example,
	// call the ListInstances operation in the Cloud Compute
	// Service or the ListVolumes operation in Cloud Block Storage.
	GetCompartment(request GetCompartmentRequest) (response GetCompartmentResponse, err error)

	// Gets the specified group's information.
	// This operation does not return a list of all the users in the group. To do that, use
	// ListUserGroupMemberships and
	// provide the group's OCID as a query parameter in the request.
	GetGroup(request GetGroupRequest) (response GetGroupResponse, err error)

	// Gets the specified identity provider's information.
	GetIdentityProvider(request GetIdentityProviderRequest) (response GetIdentityProviderResponse, err error)

	// Gets the specified group mapping.
	GetIdpGroupMapping(request GetIdpGroupMappingRequest) (response GetIdpGroupMappingResponse, err error)

	// Gets the specified policy's information.
	GetPolicy(request GetPolicyRequest) (response GetPolicyResponse, err error)

	// Get the specified tenancy's information.
	GetTenancy(request GetTenancyRequest) (response GetTenancyResponse, err error)

	// Gets the specified user's information.
	GetUser(request GetUserRequest) (response GetUserResponse, err error)

	// Gets the specified UserGroupMembership's information.
	GetUserGroupMembership(request GetUserGroupMembershipRequest) (response GetUserGroupMembershipResponse, err error)

	// Lists the API signing keys for the specified user. A user can have a maximum of three keys.
	// Every user has permission to use this API call for *their own user ID*.  An administrator in your
	// organization does not need to write a policy to give users this ability.
	ListApiKeys(request ListApiKeysRequest) (response ListApiKeysResponse, err error)

	// Lists the Availability Domains in your tenancy. Specify the OCID of either the tenancy or another
	// of your compartments as the value for the compartment ID (remember that the tenancy is simply the root compartment).
	// See [Where to Get the Tenancy's OCID and User's OCID](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/API/Concepts/apisigningkey.htm#five).
	ListAvailabilityDomains(request ListAvailabilityDomainsRequest) (response ListAvailabilityDomainsResponse, err error)

	// Lists the compartments in your tenancy. You must specify your tenancy's OCID as the value
	// for the compartment ID (remember that the tenancy is simply the root compartment).
	// See [Where to Get the Tenancy's OCID and User's OCID](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/API/Concepts/apisigningkey.htm#five).
	ListCompartments(request ListCompartmentsRequest) (response ListCompartmentsResponse, err error)

	// Lists the groups in your tenancy. You must specify your tenancy's OCID as the value for
	// the compartment ID (remember that the tenancy is simply the root compartment).
	// See [Where to Get the Tenancy's OCID and User's OCID](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/API/Concepts/apisigningkey.htm#five).
	ListGroups(request ListGroupsRequest) (response ListGroupsResponse, err error)

	// Lists all the identity providers in your tenancy. You must specify the identity provider type (e.g., `SAML2` for
	// identity providers using the SAML2.0 protocol). You must specify your tenancy's OCID as the value for the
	// compartment ID (remember that the tenancy is simply the root compartment).
	// See [Where to Get the Tenancy's OCID and User's OCID](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/API/Concepts/apisigningkey.htm#five).
	ListIdentityProviders(request ListIdentityProvidersRequest) (response ListIdentityProvidersResponse, err error)

	// Lists the group mappings for the specified identity provider.
	ListIdpGroupMappings(request ListIdpGroupMappingsRequest) (response ListIdpGroupMappingsResponse, err error)

	// Lists the policies in the specified compartment (either the tenancy or another of your compartments).
	// See [Where to Get the Tenancy's OCID and User's OCID](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/API/Concepts/apisigningkey.htm#five).
	// To determine which policies apply to a particular group or compartment, you must view the individual
	// statements inside all your policies. There isn't a way to automatically obtain that information via the API.
	ListPolicies(request ListPoliciesRequest) (response ListPoliciesResponse, err error)

	// Lists the region subscriptions for the specified tenancy.
	ListRegionSubscriptions(request ListRegionSubscriptionsRequest) (response ListRegionSubscriptionsResponse, err error)

	// Lists all the regions offered by Oracle Bare Metal Cloud Services.
	ListRegions() (response ListRegionsResponse, err error)

	// Lists the Swift passwords for the specified user. The returned object contains the password's OCID, but not
	// the password itself. The actual password is returned only upon creation.
	ListSwiftPasswords(request ListSwiftPasswordsRequest) (response ListSwiftPasswordsResponse, err error)

	// Lists the `UserGroupMembership` objects in your tenancy. You must specify your tenancy's OCID
	// as the value for the compartment ID
	// (see [Where to Get the Tenancy's OCID and User's OCID](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/API/Concepts/apisigningkey.htm#five)).
	// You must also then filter the list in one of these ways:
	// - You can limit the results to just the memberships for a given user by specifying a `userId`.
	// - Similarly, you can limit the results to just the memberships for a given group by specifying a `groupId`.
	// - You can set both the `userId` and `groupId` to determine if the specified user is in the specified group.
	// If the answer is no, the response is an empty list.
	ListUserGroupMemberships(request ListUserGroupMembershipsRequest) (response ListUserGroupMembershipsResponse, err error)

	// Lists the users in your tenancy. You must specify your tenancy's OCID as the value for the
	// compartment ID (remember that the tenancy is simply the root compartment).
	// See [Where to Get the Tenancy's OCID and User's OCID](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/API/Concepts/apisigningkey.htm#five).
	ListUsers(request ListUsersRequest) (response ListUsersResponse, err error)

	// Removes a user from a group by deleting the corresponding `UserGroupMembership`.
	RemoveUserFromGroup(request RemoveUserFromGroupRequest) (err error)

	// Updates the specified compartment's description.
	UpdateCompartment(request UpdateCompartmentRequest) (response UpdateCompartmentResponse, err error)

	// Updates the specified group.
	UpdateGroup(request UpdateGroupRequest) (response UpdateGroupResponse, err error)

	// Updates the specified identity provider.
	UpdateIdentityProvider(request UpdateIdentityProviderRequest) (response UpdateIdentityProviderResponse, err error)

	// Updates the specified group mapping.
	UpdateIdpGroupMapping(request UpdateIdpGroupMappingRequest) (response UpdateIdpGroupMappingResponse, err error)

	// Updates the specified policy. You can update the description or the policy statements themselves.
	// Policy changes take effect typically within 10 seconds.
	UpdatePolicy(request UpdatePolicyRequest) (response UpdatePolicyResponse, err error)

	// Updates the specified Swift password's description.
	UpdateSwiftPassword(request UpdateSwiftPasswordRequest) (response UpdateSwiftPasswordResponse, err error)

	// Updates the description of the specified user.
	UpdateUser(request UpdateUserRequest) (response UpdateUserResponse, err error)

	// Updates the state of the specified user.
	UpdateUserState(request UpdateUserStateRequest) (response UpdateUserStateResponse, err error)

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
	UploadApiKey(request UploadApiKeyRequest) (response UploadApiKeyResponse, err error)
}
