// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//IdentityClient a client for Identity
type IdentityClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewIdentityClientWithConfigurationProvider Creates a new default Identity client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewIdentityClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client IdentityClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = IdentityClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *IdentityClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "identity", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *IdentityClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.config = &configProvider
	client.SetRegion(region)
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *IdentityClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddUserToGroup Adds the specified user to the specified group and returns a `UserGroupMembership` object with its own OCID.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
func (client IdentityClient) AddUserToGroup(ctx context.Context, request AddUserToGroupRequest) (response AddUserToGroupResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.addUserToGroup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(AddUserToGroupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) addUserToGroup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/userGroupMemberships/")
	if err != nil {
		return nil, err
	}

	var response AddUserToGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateCompartment Creates a new compartment in your tenancy.
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
	ociResponse, e := client.Retry(ctx, request, client.createCompartment, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateCompartmentResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) createCompartment(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/compartments/")
	if err != nil {
		return nil, err
	}

	var response CreateCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateCustomerSecretKey Creates a new secret key for the specified user. Secret keys are used for authentication with the Object Storage Service's Amazon S3
// compatible API. For information, see
// [Managing User Credentials]({{DOC_SERVER_URL}}/Content/Identity/Tasks/managingcredentials.htm).
// You must specify a *description* for the secret key (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with
// UpdateCustomerSecretKey.
// Every user has permission to create a secret key for *their own user ID*. An administrator in your organization
// does not need to write a policy to give users this ability. To compare, administrators who have permission to the
// tenancy can use this operation to create a secret key for any user, including themselves.
func (client IdentityClient) CreateCustomerSecretKey(ctx context.Context, request CreateCustomerSecretKeyRequest) (response CreateCustomerSecretKeyResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createCustomerSecretKey, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateCustomerSecretKeyResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) createCustomerSecretKey(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/users/{userId}/customerSecretKeys/")
	if err != nil {
		return nil, err
	}

	var response CreateCustomerSecretKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateGroup Creates a new group in your tenancy.
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
	ociResponse, e := client.Retry(ctx, request, client.createGroup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateGroupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) createGroup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/groups/")
	if err != nil {
		return nil, err
	}

	var response CreateGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateIdentityProvider Creates a new identity provider in your tenancy. For more information, see
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
	ociResponse, e := client.Retry(ctx, request, client.createIdentityProvider, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateIdentityProviderResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) createIdentityProvider(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/identityProviders/")
	if err != nil {
		return nil, err
	}

	var response CreateIdentityProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &identityprovider{})
	return response, nil
}

// CreateIdpGroupMapping Creates a single mapping between an IdP group and an IAM Service
// Group.
func (client IdentityClient) CreateIdpGroupMapping(ctx context.Context, request CreateIdpGroupMappingRequest) (response CreateIdpGroupMappingResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createIdpGroupMapping, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateIdpGroupMappingResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) createIdpGroupMapping(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/identityProviders/{identityProviderId}/groupMappings/")
	if err != nil {
		return nil, err
	}

	var response CreateIdpGroupMappingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateOrResetUIPassword Creates a new Console one-time password for the specified user. For more information about user
// credentials, see [User Credentials]({{DOC_SERVER_URL}}/Content/Identity/Concepts/usercredentials.htm).
// Use this operation after creating a new user, or if a user forgets their password. The new one-time
// password is returned to you in the response, and you must securely deliver it to the user. They'll
// be prompted to change this password the next time they sign in to the Console. If they don't change
// it within 7 days, the password will expire and you'll need to create a new one-time password for the
// user.
// **Note:** The user's Console login is the unique name you specified when you created the user
// (see CreateUser).
func (client IdentityClient) CreateOrResetUIPassword(ctx context.Context, request CreateOrResetUIPasswordRequest) (response CreateOrResetUIPasswordResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createOrResetUIPassword, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateOrResetUIPasswordResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) createOrResetUIPassword(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/users/{userId}/uiPassword")
	if err != nil {
		return nil, err
	}

	var response CreateOrResetUIPasswordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreatePolicy Creates a new policy in the specified compartment (either the tenancy or another of your compartments).
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
	ociResponse, e := client.Retry(ctx, request, client.createPolicy, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreatePolicyResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) createPolicy(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/policies/")
	if err != nil {
		return nil, err
	}

	var response CreatePolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateRegionSubscription Creates a subscription to a region for a tenancy.
func (client IdentityClient) CreateRegionSubscription(ctx context.Context, request CreateRegionSubscriptionRequest) (response CreateRegionSubscriptionResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createRegionSubscription, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateRegionSubscriptionResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) createRegionSubscription(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/tenancies/{tenancyId}/regionSubscriptions")
	if err != nil {
		return nil, err
	}

	var response CreateRegionSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateSwiftPassword Creates a new Swift password for the specified user. For information about what Swift passwords are for, see
// [Managing User Credentials]({{DOC_SERVER_URL}}/Content/Identity/Tasks/managingcredentials.htm).
// You must specify a *description* for the Swift password (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with
// UpdateSwiftPassword.
// Every user has permission to create a Swift password for *their own user ID*. An administrator in your organization
// does not need to write a policy to give users this ability. To compare, administrators who have permission to the
// tenancy can use this operation to create a Swift password for any user, including themselves.
func (client IdentityClient) CreateSwiftPassword(ctx context.Context, request CreateSwiftPasswordRequest) (response CreateSwiftPasswordResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createSwiftPassword, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateSwiftPasswordResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) createSwiftPassword(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/users/{userId}/swiftPasswords/")
	if err != nil {
		return nil, err
	}

	var response CreateSwiftPasswordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateUser Creates a new user in your tenancy. For conceptual information about users, your tenancy, and other
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
	ociResponse, e := client.Retry(ctx, request, client.createUser, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateUserResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) createUser(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/users/")
	if err != nil {
		return nil, err
	}

	var response CreateUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteApiKey Deletes the specified API signing key for the specified user.
// Every user has permission to use this operation to delete a key for *their own user ID*. An
// administrator in your organization does not need to write a policy to give users this ability.
// To compare, administrators who have permission to the tenancy can use this operation to delete
// a key for any user, including themselves.
func (client IdentityClient) DeleteApiKey(ctx context.Context, request DeleteApiKeyRequest) (response DeleteApiKeyResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteApiKey, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteApiKeyResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) deleteApiKey(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/users/{userId}/apiKeys/{fingerprint}")
	if err != nil {
		return nil, err
	}

	var response DeleteApiKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteCustomerSecretKey Deletes the specified secret key for the specified user.
func (client IdentityClient) DeleteCustomerSecretKey(ctx context.Context, request DeleteCustomerSecretKeyRequest) (response DeleteCustomerSecretKeyResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteCustomerSecretKey, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteCustomerSecretKeyResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) deleteCustomerSecretKey(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/users/{userId}/customerSecretKeys/{customerSecretKeyId}")
	if err != nil {
		return nil, err
	}

	var response DeleteCustomerSecretKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteGroup Deletes the specified group. The group must be empty.
func (client IdentityClient) DeleteGroup(ctx context.Context, request DeleteGroupRequest) (response DeleteGroupResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteGroup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteGroupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) deleteGroup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/groups/{groupId}")
	if err != nil {
		return nil, err
	}

	var response DeleteGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteIdentityProvider Deletes the specified identity provider. The identity provider must not have
// any group mappings (see IdpGroupMapping).
func (client IdentityClient) DeleteIdentityProvider(ctx context.Context, request DeleteIdentityProviderRequest) (response DeleteIdentityProviderResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteIdentityProvider, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteIdentityProviderResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) deleteIdentityProvider(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/identityProviders/{identityProviderId}")
	if err != nil {
		return nil, err
	}

	var response DeleteIdentityProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteIdpGroupMapping Deletes the specified group mapping.
func (client IdentityClient) DeleteIdpGroupMapping(ctx context.Context, request DeleteIdpGroupMappingRequest) (response DeleteIdpGroupMappingResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteIdpGroupMapping, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteIdpGroupMappingResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) deleteIdpGroupMapping(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}")
	if err != nil {
		return nil, err
	}

	var response DeleteIdpGroupMappingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeletePolicy Deletes the specified policy. The deletion takes effect typically within 10 seconds.
func (client IdentityClient) DeletePolicy(ctx context.Context, request DeletePolicyRequest) (response DeletePolicyResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deletePolicy, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeletePolicyResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) deletePolicy(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/policies/{policyId}")
	if err != nil {
		return nil, err
	}

	var response DeletePolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteSwiftPassword Deletes the specified Swift password for the specified user.
func (client IdentityClient) DeleteSwiftPassword(ctx context.Context, request DeleteSwiftPasswordRequest) (response DeleteSwiftPasswordResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteSwiftPassword, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteSwiftPasswordResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) deleteSwiftPassword(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/users/{userId}/swiftPasswords/{swiftPasswordId}")
	if err != nil {
		return nil, err
	}

	var response DeleteSwiftPasswordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteUser Deletes the specified user. The user must not be in any groups.
func (client IdentityClient) DeleteUser(ctx context.Context, request DeleteUserRequest) (response DeleteUserResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteUser, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteUserResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) deleteUser(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/users/{userId}")
	if err != nil {
		return nil, err
	}

	var response DeleteUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetCompartment Gets the specified compartment's information.
// This operation does not return a list of all the resources inside the compartment. There is no single
// API operation that does that. Compartments can contain multiple types of resources (instances, block
// storage volumes, etc.). To find out what's in a compartment, you must call the "List" operation for
// each resource type and specify the compartment's OCID as a query parameter in the request. For example,
// call the ListInstances operation in the Cloud Compute
// Service or the ListVolumes operation in Cloud Block Storage.
func (client IdentityClient) GetCompartment(ctx context.Context, request GetCompartmentRequest) (response GetCompartmentResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getCompartment, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetCompartmentResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) getCompartment(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/compartments/{compartmentId}")
	if err != nil {
		return nil, err
	}

	var response GetCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetGroup Gets the specified group's information.
// This operation does not return a list of all the users in the group. To do that, use
// ListUserGroupMemberships and
// provide the group's OCID as a query parameter in the request.
func (client IdentityClient) GetGroup(ctx context.Context, request GetGroupRequest) (response GetGroupResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getGroup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetGroupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) getGroup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/groups/{groupId}")
	if err != nil {
		return nil, err
	}

	var response GetGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetIdentityProvider Gets the specified identity provider's information.
func (client IdentityClient) GetIdentityProvider(ctx context.Context, request GetIdentityProviderRequest) (response GetIdentityProviderResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getIdentityProvider, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetIdentityProviderResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) getIdentityProvider(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/identityProviders/{identityProviderId}")
	if err != nil {
		return nil, err
	}

	var response GetIdentityProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &identityprovider{})
	return response, nil
}

// GetIdpGroupMapping Gets the specified group mapping.
func (client IdentityClient) GetIdpGroupMapping(ctx context.Context, request GetIdpGroupMappingRequest) (response GetIdpGroupMappingResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getIdpGroupMapping, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetIdpGroupMappingResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) getIdpGroupMapping(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}")
	if err != nil {
		return nil, err
	}

	var response GetIdpGroupMappingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetPolicy Gets the specified policy's information.
func (client IdentityClient) GetPolicy(ctx context.Context, request GetPolicyRequest) (response GetPolicyResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getPolicy, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetPolicyResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) getPolicy(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/policies/{policyId}")
	if err != nil {
		return nil, err
	}

	var response GetPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetTenancy Get the specified tenancy's information.
func (client IdentityClient) GetTenancy(ctx context.Context, request GetTenancyRequest) (response GetTenancyResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getTenancy, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetTenancyResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) getTenancy(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/tenancies/{tenancyId}")
	if err != nil {
		return nil, err
	}

	var response GetTenancyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetUser Gets the specified user's information.
func (client IdentityClient) GetUser(ctx context.Context, request GetUserRequest) (response GetUserResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getUser, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetUserResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) getUser(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/users/{userId}")
	if err != nil {
		return nil, err
	}

	var response GetUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetUserGroupMembership Gets the specified UserGroupMembership's information.
func (client IdentityClient) GetUserGroupMembership(ctx context.Context, request GetUserGroupMembershipRequest) (response GetUserGroupMembershipResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getUserGroupMembership, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetUserGroupMembershipResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) getUserGroupMembership(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/userGroupMemberships/{userGroupMembershipId}")
	if err != nil {
		return nil, err
	}

	var response GetUserGroupMembershipResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListApiKeys Lists the API signing keys for the specified user. A user can have a maximum of three keys.
// Every user has permission to use this API call for *their own user ID*.  An administrator in your
// organization does not need to write a policy to give users this ability.
func (client IdentityClient) ListApiKeys(ctx context.Context, request ListApiKeysRequest) (response ListApiKeysResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listApiKeys, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListApiKeysResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) listApiKeys(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/users/{userId}/apiKeys/")
	if err != nil {
		return nil, err
	}

	var response ListApiKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListAvailabilityDomains Lists the Availability Domains in your tenancy. Specify the OCID of either the tenancy or another
// of your compartments as the value for the compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListAvailabilityDomains(ctx context.Context, request ListAvailabilityDomainsRequest) (response ListAvailabilityDomainsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listAvailabilityDomains, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListAvailabilityDomainsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) listAvailabilityDomains(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/availabilityDomains/")
	if err != nil {
		return nil, err
	}

	var response ListAvailabilityDomainsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListCompartments Lists the compartments in your tenancy. You must specify your tenancy's OCID as the value
// for the compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListCompartments(ctx context.Context, request ListCompartmentsRequest) (response ListCompartmentsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listCompartments, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListCompartmentsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) listCompartments(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/compartments/")
	if err != nil {
		return nil, err
	}

	var response ListCompartmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListCustomerSecretKeys Lists the secret keys for the specified user. The returned object contains the secret key's OCID, but not
// the secret key itself. The actual secret key is returned only upon creation.
func (client IdentityClient) ListCustomerSecretKeys(ctx context.Context, request ListCustomerSecretKeysRequest) (response ListCustomerSecretKeysResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listCustomerSecretKeys, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListCustomerSecretKeysResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) listCustomerSecretKeys(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/users/{userId}/customerSecretKeys/")
	if err != nil {
		return nil, err
	}

	var response ListCustomerSecretKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListFaultDomains Lists the Fault Domains in your tenancy. Specify the OCID of either the tenancy or another
// of your compartments as the value for the compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListFaultDomains(ctx context.Context, request ListFaultDomainsRequest) (response ListFaultDomainsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listFaultDomains, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListFaultDomainsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) listFaultDomains(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/faultDomains/")
	if err != nil {
		return nil, err
	}

	var response ListFaultDomainsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListGroups Lists the groups in your tenancy. You must specify your tenancy's OCID as the value for
// the compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListGroups(ctx context.Context, request ListGroupsRequest) (response ListGroupsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listGroups, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListGroupsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) listGroups(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/groups/")
	if err != nil {
		return nil, err
	}

	var response ListGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

//listidentityprovider allows to unmarshal list of polymorphic IdentityProvider
type listidentityprovider []identityprovider

//UnmarshalPolymorphicJSON unmarshals polymorphic json list of items
func (m *listidentityprovider) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	res := make([]IdentityProvider, len(*m))
	for i, v := range *m {
		nn, err := v.UnmarshalPolymorphicJSON(v.JsonData)
		if err != nil {
			return nil, err
		}
		res[i] = nn.(IdentityProvider)
	}
	return res, nil
}

// ListIdentityProviders Lists all the identity providers in your tenancy. You must specify the identity provider type (e.g., `SAML2` for
// identity providers using the SAML2.0 protocol). You must specify your tenancy's OCID as the value for the
// compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListIdentityProviders(ctx context.Context, request ListIdentityProvidersRequest) (response ListIdentityProvidersResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listIdentityProviders, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListIdentityProvidersResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) listIdentityProviders(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/identityProviders/")
	if err != nil {
		return nil, err
	}

	var response ListIdentityProvidersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &listidentityprovider{})
	return response, nil
}

// ListIdpGroupMappings Lists the group mappings for the specified identity provider.
func (client IdentityClient) ListIdpGroupMappings(ctx context.Context, request ListIdpGroupMappingsRequest) (response ListIdpGroupMappingsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listIdpGroupMappings, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListIdpGroupMappingsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) listIdpGroupMappings(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/identityProviders/{identityProviderId}/groupMappings/")
	if err != nil {
		return nil, err
	}

	var response ListIdpGroupMappingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListPolicies Lists the policies in the specified compartment (either the tenancy or another of your compartments).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
// To determine which policies apply to a particular group or compartment, you must view the individual
// statements inside all your policies. There isn't a way to automatically obtain that information via the API.
func (client IdentityClient) ListPolicies(ctx context.Context, request ListPoliciesRequest) (response ListPoliciesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listPolicies, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListPoliciesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) listPolicies(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/policies/")
	if err != nil {
		return nil, err
	}

	var response ListPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListRegionSubscriptions Lists the region subscriptions for the specified tenancy.
func (client IdentityClient) ListRegionSubscriptions(ctx context.Context, request ListRegionSubscriptionsRequest) (response ListRegionSubscriptionsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listRegionSubscriptions, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListRegionSubscriptionsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) listRegionSubscriptions(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/tenancies/{tenancyId}/regionSubscriptions")
	if err != nil {
		return nil, err
	}

	var response ListRegionSubscriptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListRegions Lists all the regions offered by Oracle Cloud Infrastructure.
func (client IdentityClient) ListRegions(ctx context.Context, request ListRegionsRequest) (response ListRegionsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listRegions, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListRegionsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) listRegions(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/regions")
	if err != nil {
		return nil, err
	}

	var response ListRegionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListSwiftPasswords Lists the Swift passwords for the specified user. The returned object contains the password's OCID, but not
// the password itself. The actual password is returned only upon creation.
func (client IdentityClient) ListSwiftPasswords(ctx context.Context, request ListSwiftPasswordsRequest) (response ListSwiftPasswordsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listSwiftPasswords, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListSwiftPasswordsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) listSwiftPasswords(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/users/{userId}/swiftPasswords/")
	if err != nil {
		return nil, err
	}

	var response ListSwiftPasswordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListUserGroupMemberships Lists the `UserGroupMembership` objects in your tenancy. You must specify your tenancy's OCID
// as the value for the compartment ID
// (see [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five)).
// You must also then filter the list in one of these ways:
// - You can limit the results to just the memberships for a given user by specifying a `userId`.
// - Similarly, you can limit the results to just the memberships for a given group by specifying a `groupId`.
// - You can set both the `userId` and `groupId` to determine if the specified user is in the specified group.
// If the answer is no, the response is an empty list.
func (client IdentityClient) ListUserGroupMemberships(ctx context.Context, request ListUserGroupMembershipsRequest) (response ListUserGroupMembershipsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listUserGroupMemberships, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListUserGroupMembershipsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) listUserGroupMemberships(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/userGroupMemberships/")
	if err != nil {
		return nil, err
	}

	var response ListUserGroupMembershipsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListUsers Lists the users in your tenancy. You must specify your tenancy's OCID as the value for the
// compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListUsers(ctx context.Context, request ListUsersRequest) (response ListUsersResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listUsers, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListUsersResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) listUsers(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/users/")
	if err != nil {
		return nil, err
	}

	var response ListUsersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// RemoveUserFromGroup Removes a user from a group by deleting the corresponding `UserGroupMembership`.
func (client IdentityClient) RemoveUserFromGroup(ctx context.Context, request RemoveUserFromGroupRequest) (response RemoveUserFromGroupResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.removeUserFromGroup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(RemoveUserFromGroupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) removeUserFromGroup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/userGroupMemberships/{userGroupMembershipId}")
	if err != nil {
		return nil, err
	}

	var response RemoveUserFromGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateCompartment Updates the specified compartment's description or name. You can't update the root compartment.
func (client IdentityClient) UpdateCompartment(ctx context.Context, request UpdateCompartmentRequest) (response UpdateCompartmentResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateCompartment, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateCompartmentResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) updateCompartment(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/compartments/{compartmentId}")
	if err != nil {
		return nil, err
	}

	var response UpdateCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateCustomerSecretKey Updates the specified secret key's description.
func (client IdentityClient) UpdateCustomerSecretKey(ctx context.Context, request UpdateCustomerSecretKeyRequest) (response UpdateCustomerSecretKeyResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateCustomerSecretKey, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateCustomerSecretKeyResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) updateCustomerSecretKey(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/users/{userId}/customerSecretKeys/{customerSecretKeyId}")
	if err != nil {
		return nil, err
	}

	var response UpdateCustomerSecretKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateGroup Updates the specified group.
func (client IdentityClient) UpdateGroup(ctx context.Context, request UpdateGroupRequest) (response UpdateGroupResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateGroup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateGroupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) updateGroup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/groups/{groupId}")
	if err != nil {
		return nil, err
	}

	var response UpdateGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateIdentityProvider Updates the specified identity provider.
func (client IdentityClient) UpdateIdentityProvider(ctx context.Context, request UpdateIdentityProviderRequest) (response UpdateIdentityProviderResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateIdentityProvider, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateIdentityProviderResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) updateIdentityProvider(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/identityProviders/{identityProviderId}")
	if err != nil {
		return nil, err
	}

	var response UpdateIdentityProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &identityprovider{})
	return response, nil
}

// UpdateIdpGroupMapping Updates the specified group mapping.
func (client IdentityClient) UpdateIdpGroupMapping(ctx context.Context, request UpdateIdpGroupMappingRequest) (response UpdateIdpGroupMappingResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateIdpGroupMapping, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateIdpGroupMappingResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) updateIdpGroupMapping(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}")
	if err != nil {
		return nil, err
	}

	var response UpdateIdpGroupMappingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdatePolicy Updates the specified policy. You can update the description or the policy statements themselves.
// Policy changes take effect typically within 10 seconds.
func (client IdentityClient) UpdatePolicy(ctx context.Context, request UpdatePolicyRequest) (response UpdatePolicyResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updatePolicy, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdatePolicyResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) updatePolicy(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/policies/{policyId}")
	if err != nil {
		return nil, err
	}

	var response UpdatePolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateSwiftPassword Updates the specified Swift password's description.
func (client IdentityClient) UpdateSwiftPassword(ctx context.Context, request UpdateSwiftPasswordRequest) (response UpdateSwiftPasswordResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateSwiftPassword, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateSwiftPasswordResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) updateSwiftPassword(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/users/{userId}/swiftPasswords/{swiftPasswordId}")
	if err != nil {
		return nil, err
	}

	var response UpdateSwiftPasswordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateUser Updates the description of the specified user.
func (client IdentityClient) UpdateUser(ctx context.Context, request UpdateUserRequest) (response UpdateUserResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateUser, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateUserResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) updateUser(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/users/{userId}")
	if err != nil {
		return nil, err
	}

	var response UpdateUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateUserState Updates the state of the specified user.
func (client IdentityClient) UpdateUserState(ctx context.Context, request UpdateUserStateRequest) (response UpdateUserStateResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateUserState, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateUserStateResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) updateUserState(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/users/{userId}/state/")
	if err != nil {
		return nil, err
	}

	var response UpdateUserStateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UploadApiKey Uploads an API signing key for the specified user.
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
	ociResponse, e := client.Retry(ctx, request, client.uploadApiKey, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UploadApiKeyResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client IdentityClient) uploadApiKey(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/users/{userId}/apiKeys/")
	if err != nil {
		return nil, err
	}

	var response UploadApiKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}
