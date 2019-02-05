package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/identity"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createIdentityClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := identity.NewIdentityClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientAddUserToGroup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "AddUserToGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AddUserToGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "AddUserToGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "AddUserToGroup")
	assert.NoError(t, err)

	type AddUserToGroupRequestInfo struct {
		ContainerId string
		Request     identity.AddUserToGroupRequest
	}

	var requests []AddUserToGroupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.AddUserToGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientCreateAuthToken(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "CreateAuthToken")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAuthToken is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateAuthToken", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateAuthToken")
	assert.NoError(t, err)

	type CreateAuthTokenRequestInfo struct {
		ContainerId string
		Request     identity.CreateAuthTokenRequest
	}

	var requests []CreateAuthTokenRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateAuthToken(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientCreateCompartment(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "CreateCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateCompartment", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateCompartment")
	assert.NoError(t, err)

	type CreateCompartmentRequestInfo struct {
		ContainerId string
		Request     identity.CreateCompartmentRequest
	}

	var requests []CreateCompartmentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientCreateCustomerSecretKey(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "CreateCustomerSecretKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateCustomerSecretKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateCustomerSecretKey", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateCustomerSecretKey")
	assert.NoError(t, err)

	type CreateCustomerSecretKeyRequestInfo struct {
		ContainerId string
		Request     identity.CreateCustomerSecretKeyRequest
	}

	var requests []CreateCustomerSecretKeyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateCustomerSecretKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientCreateDynamicGroup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "CreateDynamicGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateDynamicGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateDynamicGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateDynamicGroup")
	assert.NoError(t, err)

	type CreateDynamicGroupRequestInfo struct {
		ContainerId string
		Request     identity.CreateDynamicGroupRequest
	}

	var requests []CreateDynamicGroupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateDynamicGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientCreateGroup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "CreateGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateGroup")
	assert.NoError(t, err)

	type CreateGroupRequestInfo struct {
		ContainerId string
		Request     identity.CreateGroupRequest
	}

	var requests []CreateGroupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientCreateIdentityProvider(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "CreateIdentityProvider")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateIdentityProvider is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateIdentityProvider", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateIdentityProvider")
	assert.NoError(t, err)

	type CreateIdentityProviderRequestInfo struct {
		ContainerId string
		Request     identity.CreateIdentityProviderRequest
	}

	var requests []CreateIdentityProviderRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]CreateIdentityProviderRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["CreateIdentityProviderDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "protocol",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"SAML2": &identity.CreateSaml2IdentityProviderDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateIdentityProvider(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientCreateIdpGroupMapping(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "CreateIdpGroupMapping")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateIdpGroupMapping is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateIdpGroupMapping", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateIdpGroupMapping")
	assert.NoError(t, err)

	type CreateIdpGroupMappingRequestInfo struct {
		ContainerId string
		Request     identity.CreateIdpGroupMappingRequest
	}

	var requests []CreateIdpGroupMappingRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateIdpGroupMapping(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientCreateOrResetUIPassword(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "CreateOrResetUIPassword")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateOrResetUIPassword is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateOrResetUIPassword", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateOrResetUIPassword")
	assert.NoError(t, err)

	type CreateOrResetUIPasswordRequestInfo struct {
		ContainerId string
		Request     identity.CreateOrResetUIPasswordRequest
	}

	var requests []CreateOrResetUIPasswordRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateOrResetUIPassword(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientCreatePolicy(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "CreatePolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreatePolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreatePolicy", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreatePolicy")
	assert.NoError(t, err)

	type CreatePolicyRequestInfo struct {
		ContainerId string
		Request     identity.CreatePolicyRequest
	}

	var requests []CreatePolicyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreatePolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientCreateRegionSubscription(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "CreateRegionSubscription")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateRegionSubscription is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateRegionSubscription", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateRegionSubscription")
	assert.NoError(t, err)

	type CreateRegionSubscriptionRequestInfo struct {
		ContainerId string
		Request     identity.CreateRegionSubscriptionRequest
	}

	var requests []CreateRegionSubscriptionRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateRegionSubscription(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientCreateSmtpCredential(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "CreateSmtpCredential")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSmtpCredential is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateSmtpCredential", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateSmtpCredential")
	assert.NoError(t, err)

	type CreateSmtpCredentialRequestInfo struct {
		ContainerId string
		Request     identity.CreateSmtpCredentialRequest
	}

	var requests []CreateSmtpCredentialRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateSmtpCredential(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientCreateSwiftPassword(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "CreateSwiftPassword")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSwiftPassword is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateSwiftPassword", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateSwiftPassword")
	assert.NoError(t, err)

	type CreateSwiftPasswordRequestInfo struct {
		ContainerId string
		Request     identity.CreateSwiftPasswordRequest
	}

	var requests []CreateSwiftPasswordRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateSwiftPassword(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientCreateTag(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "CreateTag")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateTag is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateTag", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateTag")
	assert.NoError(t, err)

	type CreateTagRequestInfo struct {
		ContainerId string
		Request     identity.CreateTagRequest
	}

	var requests []CreateTagRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateTag(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientCreateTagNamespace(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "CreateTagNamespace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateTagNamespace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateTagNamespace", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateTagNamespace")
	assert.NoError(t, err)

	type CreateTagNamespaceRequestInfo struct {
		ContainerId string
		Request     identity.CreateTagNamespaceRequest
	}

	var requests []CreateTagNamespaceRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateTagNamespace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientCreateUser(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "CreateUser")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateUser is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "CreateUser", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "CreateUser")
	assert.NoError(t, err)

	type CreateUserRequestInfo struct {
		ContainerId string
		Request     identity.CreateUserRequest
	}

	var requests []CreateUserRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateUser(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientDeleteApiKey(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "DeleteApiKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteApiKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteApiKey", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteApiKey")
	assert.NoError(t, err)

	type DeleteApiKeyRequestInfo struct {
		ContainerId string
		Request     identity.DeleteApiKeyRequest
	}

	var requests []DeleteApiKeyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteApiKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientDeleteAuthToken(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "DeleteAuthToken")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAuthToken is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteAuthToken", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteAuthToken")
	assert.NoError(t, err)

	type DeleteAuthTokenRequestInfo struct {
		ContainerId string
		Request     identity.DeleteAuthTokenRequest
	}

	var requests []DeleteAuthTokenRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteAuthToken(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientDeleteCompartment(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "DeleteCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteCompartment", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteCompartment")
	assert.NoError(t, err)

	type DeleteCompartmentRequestInfo struct {
		ContainerId string
		Request     identity.DeleteCompartmentRequest
	}

	var requests []DeleteCompartmentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientDeleteCustomerSecretKey(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "DeleteCustomerSecretKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteCustomerSecretKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteCustomerSecretKey", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteCustomerSecretKey")
	assert.NoError(t, err)

	type DeleteCustomerSecretKeyRequestInfo struct {
		ContainerId string
		Request     identity.DeleteCustomerSecretKeyRequest
	}

	var requests []DeleteCustomerSecretKeyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteCustomerSecretKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientDeleteDynamicGroup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "DeleteDynamicGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteDynamicGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteDynamicGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteDynamicGroup")
	assert.NoError(t, err)

	type DeleteDynamicGroupRequestInfo struct {
		ContainerId string
		Request     identity.DeleteDynamicGroupRequest
	}

	var requests []DeleteDynamicGroupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteDynamicGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientDeleteGroup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "DeleteGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteGroup")
	assert.NoError(t, err)

	type DeleteGroupRequestInfo struct {
		ContainerId string
		Request     identity.DeleteGroupRequest
	}

	var requests []DeleteGroupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientDeleteIdentityProvider(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "DeleteIdentityProvider")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteIdentityProvider is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteIdentityProvider", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteIdentityProvider")
	assert.NoError(t, err)

	type DeleteIdentityProviderRequestInfo struct {
		ContainerId string
		Request     identity.DeleteIdentityProviderRequest
	}

	var requests []DeleteIdentityProviderRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteIdentityProvider(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientDeleteIdpGroupMapping(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "DeleteIdpGroupMapping")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteIdpGroupMapping is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteIdpGroupMapping", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteIdpGroupMapping")
	assert.NoError(t, err)

	type DeleteIdpGroupMappingRequestInfo struct {
		ContainerId string
		Request     identity.DeleteIdpGroupMappingRequest
	}

	var requests []DeleteIdpGroupMappingRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteIdpGroupMapping(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientDeletePolicy(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "DeletePolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeletePolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeletePolicy", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeletePolicy")
	assert.NoError(t, err)

	type DeletePolicyRequestInfo struct {
		ContainerId string
		Request     identity.DeletePolicyRequest
	}

	var requests []DeletePolicyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeletePolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientDeleteSmtpCredential(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "DeleteSmtpCredential")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSmtpCredential is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteSmtpCredential", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteSmtpCredential")
	assert.NoError(t, err)

	type DeleteSmtpCredentialRequestInfo struct {
		ContainerId string
		Request     identity.DeleteSmtpCredentialRequest
	}

	var requests []DeleteSmtpCredentialRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteSmtpCredential(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientDeleteSwiftPassword(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "DeleteSwiftPassword")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSwiftPassword is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteSwiftPassword", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteSwiftPassword")
	assert.NoError(t, err)

	type DeleteSwiftPasswordRequestInfo struct {
		ContainerId string
		Request     identity.DeleteSwiftPasswordRequest
	}

	var requests []DeleteSwiftPasswordRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteSwiftPassword(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientDeleteUser(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "DeleteUser")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteUser is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "DeleteUser", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "DeleteUser")
	assert.NoError(t, err)

	type DeleteUserRequestInfo struct {
		ContainerId string
		Request     identity.DeleteUserRequest
	}

	var requests []DeleteUserRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteUser(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientGetCompartment(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "GetCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetCompartment", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetCompartment")
	assert.NoError(t, err)

	type GetCompartmentRequestInfo struct {
		ContainerId string
		Request     identity.GetCompartmentRequest
	}

	var requests []GetCompartmentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientGetDynamicGroup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "GetDynamicGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetDynamicGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetDynamicGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetDynamicGroup")
	assert.NoError(t, err)

	type GetDynamicGroupRequestInfo struct {
		ContainerId string
		Request     identity.GetDynamicGroupRequest
	}

	var requests []GetDynamicGroupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetDynamicGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientGetGroup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "GetGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetGroup")
	assert.NoError(t, err)

	type GetGroupRequestInfo struct {
		ContainerId string
		Request     identity.GetGroupRequest
	}

	var requests []GetGroupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientGetIdentityProvider(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "GetIdentityProvider")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetIdentityProvider is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetIdentityProvider", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetIdentityProvider")
	assert.NoError(t, err)

	type GetIdentityProviderRequestInfo struct {
		ContainerId string
		Request     identity.GetIdentityProviderRequest
	}

	var requests []GetIdentityProviderRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetIdentityProvider(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientGetIdpGroupMapping(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "GetIdpGroupMapping")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetIdpGroupMapping is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetIdpGroupMapping", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetIdpGroupMapping")
	assert.NoError(t, err)

	type GetIdpGroupMappingRequestInfo struct {
		ContainerId string
		Request     identity.GetIdpGroupMappingRequest
	}

	var requests []GetIdpGroupMappingRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetIdpGroupMapping(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientGetPolicy(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "GetPolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetPolicy", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetPolicy")
	assert.NoError(t, err)

	type GetPolicyRequestInfo struct {
		ContainerId string
		Request     identity.GetPolicyRequest
	}

	var requests []GetPolicyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetPolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientGetTag(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "GetTag")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTag is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetTag", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetTag")
	assert.NoError(t, err)

	type GetTagRequestInfo struct {
		ContainerId string
		Request     identity.GetTagRequest
	}

	var requests []GetTagRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetTag(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientGetTagNamespace(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "GetTagNamespace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTagNamespace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetTagNamespace", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetTagNamespace")
	assert.NoError(t, err)

	type GetTagNamespaceRequestInfo struct {
		ContainerId string
		Request     identity.GetTagNamespaceRequest
	}

	var requests []GetTagNamespaceRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetTagNamespace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientGetTenancy(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "GetTenancy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTenancy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetTenancy", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetTenancy")
	assert.NoError(t, err)

	type GetTenancyRequestInfo struct {
		ContainerId string
		Request     identity.GetTenancyRequest
	}

	var requests []GetTenancyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetTenancy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientGetUser(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "GetUser")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetUser is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetUser", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetUser")
	assert.NoError(t, err)

	type GetUserRequestInfo struct {
		ContainerId string
		Request     identity.GetUserRequest
	}

	var requests []GetUserRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetUser(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientGetUserGroupMembership(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "GetUserGroupMembership")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetUserGroupMembership is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetUserGroupMembership", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetUserGroupMembership")
	assert.NoError(t, err)

	type GetUserGroupMembershipRequestInfo struct {
		ContainerId string
		Request     identity.GetUserGroupMembershipRequest
	}

	var requests []GetUserGroupMembershipRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetUserGroupMembership(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientGetWorkRequest(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "GetWorkRequest", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     identity.GetWorkRequestRequest
	}

	var requests []GetWorkRequestRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListApiKeys(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListApiKeys")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListApiKeys is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListApiKeys", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListApiKeys")
	assert.NoError(t, err)

	type ListApiKeysRequestInfo struct {
		ContainerId string
		Request     identity.ListApiKeysRequest
	}

	var requests []ListApiKeysRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ListApiKeys(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListAuthTokens(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListAuthTokens")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAuthTokens is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListAuthTokens", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListAuthTokens")
	assert.NoError(t, err)

	type ListAuthTokensRequestInfo struct {
		ContainerId string
		Request     identity.ListAuthTokensRequest
	}

	var requests []ListAuthTokensRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ListAuthTokens(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListAvailabilityDomains(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListAvailabilityDomains")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAvailabilityDomains is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListAvailabilityDomains", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListAvailabilityDomains")
	assert.NoError(t, err)

	type ListAvailabilityDomainsRequestInfo struct {
		ContainerId string
		Request     identity.ListAvailabilityDomainsRequest
	}

	var requests []ListAvailabilityDomainsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ListAvailabilityDomains(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListCompartments(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListCompartments")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCompartments is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListCompartments", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListCompartments")
	assert.NoError(t, err)

	type ListCompartmentsRequestInfo struct {
		ContainerId string
		Request     identity.ListCompartmentsRequest
	}

	var requests []ListCompartmentsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*identity.ListCompartmentsRequest)
				return c.ListCompartments(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListCompartmentsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListCompartmentsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListCostTrackingTags(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListCostTrackingTags")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCostTrackingTags is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListCostTrackingTags", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListCostTrackingTags")
	assert.NoError(t, err)

	type ListCostTrackingTagsRequestInfo struct {
		ContainerId string
		Request     identity.ListCostTrackingTagsRequest
	}

	var requests []ListCostTrackingTagsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*identity.ListCostTrackingTagsRequest)
				return c.ListCostTrackingTags(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListCostTrackingTagsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListCostTrackingTagsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListCustomerSecretKeys(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListCustomerSecretKeys")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCustomerSecretKeys is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListCustomerSecretKeys", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListCustomerSecretKeys")
	assert.NoError(t, err)

	type ListCustomerSecretKeysRequestInfo struct {
		ContainerId string
		Request     identity.ListCustomerSecretKeysRequest
	}

	var requests []ListCustomerSecretKeysRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ListCustomerSecretKeys(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListDynamicGroups(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListDynamicGroups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListDynamicGroups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListDynamicGroups", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListDynamicGroups")
	assert.NoError(t, err)

	type ListDynamicGroupsRequestInfo struct {
		ContainerId string
		Request     identity.ListDynamicGroupsRequest
	}

	var requests []ListDynamicGroupsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*identity.ListDynamicGroupsRequest)
				return c.ListDynamicGroups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListDynamicGroupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListDynamicGroupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListFaultDomains(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListFaultDomains")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListFaultDomains is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListFaultDomains", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListFaultDomains")
	assert.NoError(t, err)

	type ListFaultDomainsRequestInfo struct {
		ContainerId string
		Request     identity.ListFaultDomainsRequest
	}

	var requests []ListFaultDomainsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ListFaultDomains(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListGroups(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListGroups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListGroups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListGroups", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListGroups")
	assert.NoError(t, err)

	type ListGroupsRequestInfo struct {
		ContainerId string
		Request     identity.ListGroupsRequest
	}

	var requests []ListGroupsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*identity.ListGroupsRequest)
				return c.ListGroups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListGroupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListGroupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListIdentityProviderGroups(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListIdentityProviderGroups")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListIdentityProviderGroups is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListIdentityProviderGroups", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListIdentityProviderGroups")
	assert.NoError(t, err)

	type ListIdentityProviderGroupsRequestInfo struct {
		ContainerId string
		Request     identity.ListIdentityProviderGroupsRequest
	}

	var requests []ListIdentityProviderGroupsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*identity.ListIdentityProviderGroupsRequest)
				return c.ListIdentityProviderGroups(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListIdentityProviderGroupsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListIdentityProviderGroupsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListIdentityProviders(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListIdentityProviders")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListIdentityProviders is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListIdentityProviders", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListIdentityProviders")
	assert.NoError(t, err)

	type ListIdentityProvidersRequestInfo struct {
		ContainerId string
		Request     identity.ListIdentityProvidersRequest
	}

	var requests []ListIdentityProvidersRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*identity.ListIdentityProvidersRequest)
				return c.ListIdentityProviders(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListIdentityProvidersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListIdentityProvidersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListIdpGroupMappings(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListIdpGroupMappings")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListIdpGroupMappings is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListIdpGroupMappings", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListIdpGroupMappings")
	assert.NoError(t, err)

	type ListIdpGroupMappingsRequestInfo struct {
		ContainerId string
		Request     identity.ListIdpGroupMappingsRequest
	}

	var requests []ListIdpGroupMappingsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*identity.ListIdpGroupMappingsRequest)
				return c.ListIdpGroupMappings(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListIdpGroupMappingsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListIdpGroupMappingsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListPolicies(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListPolicies")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPolicies is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListPolicies", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListPolicies")
	assert.NoError(t, err)

	type ListPoliciesRequestInfo struct {
		ContainerId string
		Request     identity.ListPoliciesRequest
	}

	var requests []ListPoliciesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*identity.ListPoliciesRequest)
				return c.ListPolicies(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListPoliciesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListPoliciesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListRegionSubscriptions(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListRegionSubscriptions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRegionSubscriptions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListRegionSubscriptions", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListRegionSubscriptions")
	assert.NoError(t, err)

	type ListRegionSubscriptionsRequestInfo struct {
		ContainerId string
		Request     identity.ListRegionSubscriptionsRequest
	}

	var requests []ListRegionSubscriptionsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ListRegionSubscriptions(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListRegions(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListRegions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListRegions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListRegions", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListRegions")
	assert.NoError(t, err)

	type ListRegionsRequestInfo struct {
		ContainerId string
		Request     interface{}
	}

	var requests []ListRegionsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {

			response, err := c.ListRegions(context.Background())
			message, err := testClient.validateResult(req.ContainerId, nil, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListSmtpCredentials(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListSmtpCredentials")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSmtpCredentials is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListSmtpCredentials", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListSmtpCredentials")
	assert.NoError(t, err)

	type ListSmtpCredentialsRequestInfo struct {
		ContainerId string
		Request     identity.ListSmtpCredentialsRequest
	}

	var requests []ListSmtpCredentialsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ListSmtpCredentials(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListSwiftPasswords(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListSwiftPasswords")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSwiftPasswords is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListSwiftPasswords", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListSwiftPasswords")
	assert.NoError(t, err)

	type ListSwiftPasswordsRequestInfo struct {
		ContainerId string
		Request     identity.ListSwiftPasswordsRequest
	}

	var requests []ListSwiftPasswordsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ListSwiftPasswords(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListTagNamespaces(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListTagNamespaces")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTagNamespaces is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListTagNamespaces", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListTagNamespaces")
	assert.NoError(t, err)

	type ListTagNamespacesRequestInfo struct {
		ContainerId string
		Request     identity.ListTagNamespacesRequest
	}

	var requests []ListTagNamespacesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*identity.ListTagNamespacesRequest)
				return c.ListTagNamespaces(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListTagNamespacesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListTagNamespacesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListTags(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListTags")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTags is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListTags", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListTags")
	assert.NoError(t, err)

	type ListTagsRequestInfo struct {
		ContainerId string
		Request     identity.ListTagsRequest
	}

	var requests []ListTagsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*identity.ListTagsRequest)
				return c.ListTags(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListTagsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListTagsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListUserGroupMemberships(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListUserGroupMemberships")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListUserGroupMemberships is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListUserGroupMemberships", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListUserGroupMemberships")
	assert.NoError(t, err)

	type ListUserGroupMembershipsRequestInfo struct {
		ContainerId string
		Request     identity.ListUserGroupMembershipsRequest
	}

	var requests []ListUserGroupMembershipsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*identity.ListUserGroupMembershipsRequest)
				return c.ListUserGroupMemberships(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListUserGroupMembershipsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListUserGroupMembershipsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListUsers(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListUsers")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListUsers is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListUsers", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListUsers")
	assert.NoError(t, err)

	type ListUsersRequestInfo struct {
		ContainerId string
		Request     identity.ListUsersRequest
	}

	var requests []ListUsersRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*identity.ListUsersRequest)
				return c.ListUsers(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListUsersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListUsersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientListWorkRequests(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ListWorkRequests", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     identity.ListWorkRequestsRequest
	}

	var requests []ListWorkRequestsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*identity.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]identity.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(identity.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientRemoveUserFromGroup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "RemoveUserFromGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RemoveUserFromGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "RemoveUserFromGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "RemoveUserFromGroup")
	assert.NoError(t, err)

	type RemoveUserFromGroupRequestInfo struct {
		ContainerId string
		Request     identity.RemoveUserFromGroupRequest
	}

	var requests []RemoveUserFromGroupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.RemoveUserFromGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientResetIdpScimClient(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "ResetIdpScimClient")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ResetIdpScimClient is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "ResetIdpScimClient", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "ResetIdpScimClient")
	assert.NoError(t, err)

	type ResetIdpScimClientRequestInfo struct {
		ContainerId string
		Request     identity.ResetIdpScimClientRequest
	}

	var requests []ResetIdpScimClientRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ResetIdpScimClient(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientUpdateAuthToken(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "UpdateAuthToken")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAuthToken is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateAuthToken", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateAuthToken")
	assert.NoError(t, err)

	type UpdateAuthTokenRequestInfo struct {
		ContainerId string
		Request     identity.UpdateAuthTokenRequest
	}

	var requests []UpdateAuthTokenRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateAuthToken(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientUpdateCompartment(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "UpdateCompartment")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateCompartment is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateCompartment", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateCompartment")
	assert.NoError(t, err)

	type UpdateCompartmentRequestInfo struct {
		ContainerId string
		Request     identity.UpdateCompartmentRequest
	}

	var requests []UpdateCompartmentRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateCompartment(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientUpdateCustomerSecretKey(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "UpdateCustomerSecretKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateCustomerSecretKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateCustomerSecretKey", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateCustomerSecretKey")
	assert.NoError(t, err)

	type UpdateCustomerSecretKeyRequestInfo struct {
		ContainerId string
		Request     identity.UpdateCustomerSecretKeyRequest
	}

	var requests []UpdateCustomerSecretKeyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateCustomerSecretKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientUpdateDynamicGroup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "UpdateDynamicGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateDynamicGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateDynamicGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateDynamicGroup")
	assert.NoError(t, err)

	type UpdateDynamicGroupRequestInfo struct {
		ContainerId string
		Request     identity.UpdateDynamicGroupRequest
	}

	var requests []UpdateDynamicGroupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateDynamicGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientUpdateGroup(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "UpdateGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateGroup", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateGroup")
	assert.NoError(t, err)

	type UpdateGroupRequestInfo struct {
		ContainerId string
		Request     identity.UpdateGroupRequest
	}

	var requests []UpdateGroupRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientUpdateIdentityProvider(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "UpdateIdentityProvider")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateIdentityProvider is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateIdentityProvider", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateIdentityProvider")
	assert.NoError(t, err)

	type UpdateIdentityProviderRequestInfo struct {
		ContainerId string
		Request     identity.UpdateIdentityProviderRequest
	}

	var requests []UpdateIdentityProviderRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]UpdateIdentityProviderRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["UpdateIdentityProviderDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "protocol",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"SAML2": &identity.UpdateSaml2IdentityProviderDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateIdentityProvider(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientUpdateIdpGroupMapping(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "UpdateIdpGroupMapping")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateIdpGroupMapping is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateIdpGroupMapping", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateIdpGroupMapping")
	assert.NoError(t, err)

	type UpdateIdpGroupMappingRequestInfo struct {
		ContainerId string
		Request     identity.UpdateIdpGroupMappingRequest
	}

	var requests []UpdateIdpGroupMappingRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateIdpGroupMapping(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientUpdatePolicy(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "UpdatePolicy")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdatePolicy is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdatePolicy", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdatePolicy")
	assert.NoError(t, err)

	type UpdatePolicyRequestInfo struct {
		ContainerId string
		Request     identity.UpdatePolicyRequest
	}

	var requests []UpdatePolicyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdatePolicy(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientUpdateSmtpCredential(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "UpdateSmtpCredential")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSmtpCredential is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateSmtpCredential", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateSmtpCredential")
	assert.NoError(t, err)

	type UpdateSmtpCredentialRequestInfo struct {
		ContainerId string
		Request     identity.UpdateSmtpCredentialRequest
	}

	var requests []UpdateSmtpCredentialRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateSmtpCredential(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientUpdateSwiftPassword(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "UpdateSwiftPassword")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSwiftPassword is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateSwiftPassword", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateSwiftPassword")
	assert.NoError(t, err)

	type UpdateSwiftPasswordRequestInfo struct {
		ContainerId string
		Request     identity.UpdateSwiftPasswordRequest
	}

	var requests []UpdateSwiftPasswordRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateSwiftPassword(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientUpdateTag(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "UpdateTag")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTag is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateTag", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateTag")
	assert.NoError(t, err)

	type UpdateTagRequestInfo struct {
		ContainerId string
		Request     identity.UpdateTagRequest
	}

	var requests []UpdateTagRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateTag(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientUpdateTagNamespace(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "UpdateTagNamespace")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTagNamespace is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateTagNamespace", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateTagNamespace")
	assert.NoError(t, err)

	type UpdateTagNamespaceRequestInfo struct {
		ContainerId string
		Request     identity.UpdateTagNamespaceRequest
	}

	var requests []UpdateTagNamespaceRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateTagNamespace(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientUpdateUser(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "UpdateUser")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateUser is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateUser", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateUser")
	assert.NoError(t, err)

	type UpdateUserRequestInfo struct {
		ContainerId string
		Request     identity.UpdateUserRequest
	}

	var requests []UpdateUserRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateUser(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientUpdateUserCapabilities(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "UpdateUserCapabilities")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateUserCapabilities is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateUserCapabilities", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateUserCapabilities")
	assert.NoError(t, err)

	type UpdateUserCapabilitiesRequestInfo struct {
		ContainerId string
		Request     identity.UpdateUserCapabilitiesRequest
	}

	var requests []UpdateUserCapabilitiesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateUserCapabilities(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientUpdateUserState(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "UpdateUserState")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateUserState is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UpdateUserState", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UpdateUserState")
	assert.NoError(t, err)

	type UpdateUserStateRequestInfo struct {
		ContainerId string
		Request     identity.UpdateUserStateRequest
	}

	var requests []UpdateUserStateRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateUserState(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestIdentityClientUploadApiKey(t *testing.T) {
	enabled, err := testClient.isApiEnabled("identity", "UploadApiKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UploadApiKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("identity", "Identity", "UploadApiKey", createIdentityClientWithProvider)
	assert.NoError(t, err)
	c := cc.(identity.IdentityClient)

	body, err := testClient.getRequests("identity", "UploadApiKey")
	assert.NoError(t, err)

	type UploadApiKeyRequestInfo struct {
		ContainerId string
		Request     identity.UploadApiKeyRequest
	}

	var requests []UploadApiKeyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UploadApiKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
