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

func createKmsCryptoClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {
	client, err := keymanagement.NewKmsCryptoClientWithConfigurationProvider(p, testConfig.Endpoint)
	return client, err
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestKmsCryptoClientDecrypt(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "Decrypt")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("Decrypt is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsCrypto", "Decrypt", createKmsCryptoClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsCryptoClient)

	body, err := testClient.getRequests("keymanagement", "Decrypt")
	assert.NoError(t, err)

	type DecryptRequestInfo struct {
		ContainerId string
		Request     keymanagement.DecryptRequest
	}

	var requests []DecryptRequestInfo
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

			response, err := c.Decrypt(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestKmsCryptoClientEncrypt(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "Encrypt")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("Encrypt is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsCrypto", "Encrypt", createKmsCryptoClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsCryptoClient)

	body, err := testClient.getRequests("keymanagement", "Encrypt")
	assert.NoError(t, err)

	type EncryptRequestInfo struct {
		ContainerId string
		Request     keymanagement.EncryptRequest
	}

	var requests []EncryptRequestInfo
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

			response, err := c.Encrypt(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestKmsCryptoClientGenerateDataEncryptionKey(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("keymanagement", "GenerateDataEncryptionKey")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GenerateDataEncryptionKey is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("keymanagement", "KmsCrypto", "GenerateDataEncryptionKey", createKmsCryptoClientWithProvider)
	assert.NoError(t, err)
	c := cc.(keymanagement.KmsCryptoClient)

	body, err := testClient.getRequests("keymanagement", "GenerateDataEncryptionKey")
	assert.NoError(t, err)

	type GenerateDataEncryptionKeyRequestInfo struct {
		ContainerId string
		Request     keymanagement.GenerateDataEncryptionKeyRequest
	}

	var requests []GenerateDataEncryptionKeyRequestInfo
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

			response, err := c.GenerateDataEncryptionKey(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
