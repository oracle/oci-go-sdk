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

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestKmsCryptoClientDecrypt(t *testing.T) {
	enabled, err := testClient.isApiEnabled("keymanagement", "Decrypt")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("Decrypt is not enabled by the testing service")
	}
	endpoint, err := testClient.getEndpointForService("keymanagement", "KmsCrypto")
	assert.NoError(t, err)
	c, err := keymanagement.NewKmsCryptoClientWithConfigurationProvider(testConfig.ConfigurationProvider, endpoint)
	assert.NoError(t, err)

	body, err := testClient.getRequests("keymanagement", "Decrypt")
	assert.NoError(t, err)

	type DecryptRequestInfo struct {
		ContainerId string
		Request     keymanagement.DecryptRequest
	}

	var requests []DecryptRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.Decrypt(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestKmsCryptoClientEncrypt(t *testing.T) {
	enabled, err := testClient.isApiEnabled("keymanagement", "Encrypt")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("Encrypt is not enabled by the testing service")
	}
	endpoint, err := testClient.getEndpointForService("keymanagement", "KmsCrypto")
	assert.NoError(t, err)
	c, err := keymanagement.NewKmsCryptoClientWithConfigurationProvider(testConfig.ConfigurationProvider, endpoint)
	assert.NoError(t, err)

	body, err := testClient.getRequests("keymanagement", "Encrypt")
	assert.NoError(t, err)

	type EncryptRequestInfo struct {
		ContainerId string
		Request     keymanagement.EncryptRequest
	}

	var requests []EncryptRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.Encrypt(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestKmsCryptoClientGenerateDataEncryptionKey(t *testing.T) {
	enabled, err := testClient.isApiEnabled("keymanagement", "GenerateDataEncryptionKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GenerateDataEncryptionKey is not enabled by the testing service")
	}
	endpoint, err := testClient.getEndpointForService("keymanagement", "KmsCrypto")
	assert.NoError(t, err)
	c, err := keymanagement.NewKmsCryptoClientWithConfigurationProvider(testConfig.ConfigurationProvider, endpoint)
	assert.NoError(t, err)

	body, err := testClient.getRequests("keymanagement", "GenerateDataEncryptionKey")
	assert.NoError(t, err)

	type GenerateDataEncryptionKeyRequestInfo struct {
		ContainerId string
		Request     keymanagement.GenerateDataEncryptionKeyRequest
	}

	var requests []GenerateDataEncryptionKeyRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GenerateDataEncryptionKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
