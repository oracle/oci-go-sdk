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

func createKmsManagementClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {
	client, err := keymanagement.NewKmsManagementClientWithConfigurationProvider(p, testConfig.Endpoint)
	return client, err
}

// IssueRoutingInfo tag="default" email="sparta_kms_us_grp@oracle.com" jiraProject="KMS" opsJiraProject="KMS"
func TestKmsManagementClientCreateKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "CreateKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsManagement", "CreateKey", createKmsManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsManagementClient)

	body, err := testClient.getRequests("keymanagement", "CreateKey")
	assert.NoError(t, err)

	type CreateKeyRequestInfo struct {
		ContainerId string
		Request     keymanagement.CreateKeyRequest
	}

	var requests []CreateKeyRequestInfo
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

			response, err := c.CreateKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sparta_kms_us_grp@oracle.com" jiraProject="KMS" opsJiraProject="KMS"
func TestKmsManagementClientCreateKeyVersion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "CreateKeyVersion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateKeyVersion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsManagement", "CreateKeyVersion", createKmsManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsManagementClient)

	body, err := testClient.getRequests("keymanagement", "CreateKeyVersion")
	assert.NoError(t, err)

	type CreateKeyVersionRequestInfo struct {
		ContainerId string
		Request     keymanagement.CreateKeyVersionRequest
	}

	var requests []CreateKeyVersionRequestInfo
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

			response, err := c.CreateKeyVersion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sparta_kms_us_grp@oracle.com" jiraProject="KMS" opsJiraProject="KMS"
func TestKmsManagementClientDisableKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "DisableKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DisableKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsManagement", "DisableKey", createKmsManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsManagementClient)

	body, err := testClient.getRequests("keymanagement", "DisableKey")
	assert.NoError(t, err)

	type DisableKeyRequestInfo struct {
		ContainerId string
		Request     keymanagement.DisableKeyRequest
	}

	var requests []DisableKeyRequestInfo
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

			response, err := c.DisableKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sparta_kms_us_grp@oracle.com" jiraProject="KMS" opsJiraProject="KMS"
func TestKmsManagementClientEnableKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "EnableKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("EnableKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsManagement", "EnableKey", createKmsManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsManagementClient)

	body, err := testClient.getRequests("keymanagement", "EnableKey")
	assert.NoError(t, err)

	type EnableKeyRequestInfo struct {
		ContainerId string
		Request     keymanagement.EnableKeyRequest
	}

	var requests []EnableKeyRequestInfo
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

			response, err := c.EnableKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sparta_kms_us_grp@oracle.com" jiraProject="KMS" opsJiraProject="KMS"
func TestKmsManagementClientGetKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "GetKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsManagement", "GetKey", createKmsManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsManagementClient)

	body, err := testClient.getRequests("keymanagement", "GetKey")
	assert.NoError(t, err)

	type GetKeyRequestInfo struct {
		ContainerId string
		Request     keymanagement.GetKeyRequest
	}

	var requests []GetKeyRequestInfo
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

			response, err := c.GetKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sparta_kms_us_grp@oracle.com" jiraProject="KMS" opsJiraProject="KMS"
func TestKmsManagementClientGetKeyVersion(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "GetKeyVersion")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetKeyVersion is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsManagement", "GetKeyVersion", createKmsManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsManagementClient)

	body, err := testClient.getRequests("keymanagement", "GetKeyVersion")
	assert.NoError(t, err)

	type GetKeyVersionRequestInfo struct {
		ContainerId string
		Request     keymanagement.GetKeyVersionRequest
	}

	var requests []GetKeyVersionRequestInfo
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

			response, err := c.GetKeyVersion(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sparta_kms_us_grp@oracle.com" jiraProject="KMS" opsJiraProject="KMS"
func TestKmsManagementClientListKeyVersions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "ListKeyVersions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListKeyVersions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsManagement", "ListKeyVersions", createKmsManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsManagementClient)

	body, err := testClient.getRequests("keymanagement", "ListKeyVersions")
	assert.NoError(t, err)

	type ListKeyVersionsRequestInfo struct {
		ContainerId string
		Request     keymanagement.ListKeyVersionsRequest
	}

	var requests []ListKeyVersionsRequestInfo
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
				r := req.(*keymanagement.ListKeyVersionsRequest)
				return c.ListKeyVersions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]keymanagement.ListKeyVersionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(keymanagement.ListKeyVersionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sparta_kms_us_grp@oracle.com" jiraProject="KMS" opsJiraProject="KMS"
func TestKmsManagementClientListKeys(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "ListKeys")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListKeys is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsManagement", "ListKeys", createKmsManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsManagementClient)

	body, err := testClient.getRequests("keymanagement", "ListKeys")
	assert.NoError(t, err)

	type ListKeysRequestInfo struct {
		ContainerId string
		Request     keymanagement.ListKeysRequest
	}

	var requests []ListKeysRequestInfo
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
				r := req.(*keymanagement.ListKeysRequest)
				return c.ListKeys(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]keymanagement.ListKeysResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(keymanagement.ListKeysResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sparta_kms_us_grp@oracle.com" jiraProject="KMS" opsJiraProject="KMS"
func TestKmsManagementClientUpdateKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "UpdateKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsManagement", "UpdateKey", createKmsManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsManagementClient)

	body, err := testClient.getRequests("keymanagement", "UpdateKey")
	assert.NoError(t, err)

	type UpdateKeyRequestInfo struct {
		ContainerId string
		Request     keymanagement.UpdateKeyRequest
	}

	var requests []UpdateKeyRequestInfo
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

			response, err := c.UpdateKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
