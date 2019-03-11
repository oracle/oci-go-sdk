package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/keymanagement"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createKmsVaultClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := keymanagement.NewKmsVaultClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestKmsVaultClientCancelVaultDeletion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "CancelVaultDeletion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CancelVaultDeletion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsVault", "CancelVaultDeletion", createKmsVaultClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsVaultClient)

	body, err := testClient.getRequests("keymanagement", "CancelVaultDeletion")
	assert.NoError(t, err)

	type CancelVaultDeletionRequestInfo struct {
		ContainerId string
		Request     keymanagement.CancelVaultDeletionRequest
	}

	var requests []CancelVaultDeletionRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CancelVaultDeletion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestKmsVaultClientCreateVault(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "CreateVault")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateVault is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsVault", "CreateVault", createKmsVaultClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsVaultClient)

	body, err := testClient.getRequests("keymanagement", "CreateVault")
	assert.NoError(t, err)

	type CreateVaultRequestInfo struct {
		ContainerId string
		Request     keymanagement.CreateVaultRequest
	}

	var requests []CreateVaultRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateVault(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestKmsVaultClientGetVault(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "GetVault")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetVault is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsVault", "GetVault", createKmsVaultClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsVaultClient)

	body, err := testClient.getRequests("keymanagement", "GetVault")
	assert.NoError(t, err)

	type GetVaultRequestInfo struct {
		ContainerId string
		Request     keymanagement.GetVaultRequest
	}

	var requests []GetVaultRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetVault(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestKmsVaultClientListVaults(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "ListVaults")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListVaults is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsVault", "ListVaults", createKmsVaultClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsVaultClient)

	body, err := testClient.getRequests("keymanagement", "ListVaults")
	assert.NoError(t, err)

	type ListVaultsRequestInfo struct {
		ContainerId string
		Request     keymanagement.ListVaultsRequest
	}

	var requests []ListVaultsRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*keymanagement.ListVaultsRequest)
				return c.ListVaults(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]keymanagement.ListVaultsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(keymanagement.ListVaultsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestKmsVaultClientScheduleVaultDeletion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "ScheduleVaultDeletion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ScheduleVaultDeletion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsVault", "ScheduleVaultDeletion", createKmsVaultClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsVaultClient)

	body, err := testClient.getRequests("keymanagement", "ScheduleVaultDeletion")
	assert.NoError(t, err)

	type ScheduleVaultDeletionRequestInfo struct {
		ContainerId string
		Request     keymanagement.ScheduleVaultDeletionRequest
	}

	var requests []ScheduleVaultDeletionRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ScheduleVaultDeletion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestKmsVaultClientUpdateVault(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "UpdateVault")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateVault is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsVault", "UpdateVault", createKmsVaultClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsVaultClient)

	body, err := testClient.getRequests("keymanagement", "UpdateVault")
	assert.NoError(t, err)

	type UpdateVaultRequestInfo struct {
		ContainerId string
		Request     keymanagement.UpdateVaultRequest
	}

	var requests []UpdateVaultRequestInfo
	var dataHolder []map[string]interface{}
	err = json.Unmarshal([]byte(body), &dataHolder)
	assert.NoError(t, err)
	err = unmarshalRequestInfo(dataHolder, &requests, testClient.Log)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateVault(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
