package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/ons"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createNotificationControlPlaneClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := ons.NewNotificationControlPlaneClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestNotificationControlPlaneClientCreateTopic(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ons", "CreateTopic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateTopic is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ons", "NotificationControlPlane", "CreateTopic", createNotificationControlPlaneClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ons.NotificationControlPlaneClient)

	body, err := testClient.getRequests("ons", "CreateTopic")
	assert.NoError(t, err)

	type CreateTopicRequestInfo struct {
		ContainerId string
		Request     ons.CreateTopicRequest
	}

	var requests []CreateTopicRequestInfo
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

			response, err := c.CreateTopic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestNotificationControlPlaneClientDeleteTopic(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ons", "DeleteTopic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteTopic is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ons", "NotificationControlPlane", "DeleteTopic", createNotificationControlPlaneClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ons.NotificationControlPlaneClient)

	body, err := testClient.getRequests("ons", "DeleteTopic")
	assert.NoError(t, err)

	type DeleteTopicRequestInfo struct {
		ContainerId string
		Request     ons.DeleteTopicRequest
	}

	var requests []DeleteTopicRequestInfo
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

			response, err := c.DeleteTopic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestNotificationControlPlaneClientGetTopic(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ons", "GetTopic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetTopic is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ons", "NotificationControlPlane", "GetTopic", createNotificationControlPlaneClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ons.NotificationControlPlaneClient)

	body, err := testClient.getRequests("ons", "GetTopic")
	assert.NoError(t, err)

	type GetTopicRequestInfo struct {
		ContainerId string
		Request     ons.GetTopicRequest
	}

	var requests []GetTopicRequestInfo
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

			response, err := c.GetTopic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestNotificationControlPlaneClientListTopics(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ons", "ListTopics")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListTopics is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ons", "NotificationControlPlane", "ListTopics", createNotificationControlPlaneClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ons.NotificationControlPlaneClient)

	body, err := testClient.getRequests("ons", "ListTopics")
	assert.NoError(t, err)

	type ListTopicsRequestInfo struct {
		ContainerId string
		Request     ons.ListTopicsRequest
	}

	var requests []ListTopicsRequestInfo
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
				r := req.(*ons.ListTopicsRequest)
				return c.ListTopics(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]ons.ListTopicsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(ons.ListTopicsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestNotificationControlPlaneClientUpdateTopic(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ons", "UpdateTopic")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateTopic is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ons", "NotificationControlPlane", "UpdateTopic", createNotificationControlPlaneClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ons.NotificationControlPlaneClient)

	body, err := testClient.getRequests("ons", "UpdateTopic")
	assert.NoError(t, err)

	type UpdateTopicRequestInfo struct {
		ContainerId string
		Request     ons.UpdateTopicRequest
	}

	var requests []UpdateTopicRequestInfo
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

			response, err := c.UpdateTopic(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
