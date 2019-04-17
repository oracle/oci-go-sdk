package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/streaming"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createStreamClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := streaming.NewStreamClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestStreamClientConsumerCommit(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "ConsumerCommit")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ConsumerCommit is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "Stream", "ConsumerCommit", createStreamClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamClient)

	body, err := testClient.getRequests("streaming", "ConsumerCommit")
	assert.NoError(t, err)

	type ConsumerCommitRequestInfo struct {
		ContainerId string
		Request     streaming.ConsumerCommitRequest
	}

	var requests []ConsumerCommitRequestInfo
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

			response, err := c.ConsumerCommit(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestStreamClientConsumerHeartbeat(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "ConsumerHeartbeat")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ConsumerHeartbeat is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "Stream", "ConsumerHeartbeat", createStreamClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamClient)

	body, err := testClient.getRequests("streaming", "ConsumerHeartbeat")
	assert.NoError(t, err)

	type ConsumerHeartbeatRequestInfo struct {
		ContainerId string
		Request     streaming.ConsumerHeartbeatRequest
	}

	var requests []ConsumerHeartbeatRequestInfo
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

			response, err := c.ConsumerHeartbeat(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestStreamClientCreateCursor(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "CreateCursor")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateCursor is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "Stream", "CreateCursor", createStreamClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamClient)

	body, err := testClient.getRequests("streaming", "CreateCursor")
	assert.NoError(t, err)

	type CreateCursorRequestInfo struct {
		ContainerId string
		Request     streaming.CreateCursorRequest
	}

	var requests []CreateCursorRequestInfo
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

			response, err := c.CreateCursor(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestStreamClientCreateGroupCursor(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "CreateGroupCursor")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateGroupCursor is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "Stream", "CreateGroupCursor", createStreamClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamClient)

	body, err := testClient.getRequests("streaming", "CreateGroupCursor")
	assert.NoError(t, err)

	type CreateGroupCursorRequestInfo struct {
		ContainerId string
		Request     streaming.CreateGroupCursorRequest
	}

	var requests []CreateGroupCursorRequestInfo
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

			response, err := c.CreateGroupCursor(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestStreamClientGetGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "GetGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "Stream", "GetGroup", createStreamClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamClient)

	body, err := testClient.getRequests("streaming", "GetGroup")
	assert.NoError(t, err)

	type GetGroupRequestInfo struct {
		ContainerId string
		Request     streaming.GetGroupRequest
	}

	var requests []GetGroupRequestInfo
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

			response, err := c.GetGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestStreamClientGetMessages(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "GetMessages")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetMessages is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "Stream", "GetMessages", createStreamClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamClient)

	body, err := testClient.getRequests("streaming", "GetMessages")
	assert.NoError(t, err)

	type GetMessagesRequestInfo struct {
		ContainerId string
		Request     streaming.GetMessagesRequest
	}

	var requests []GetMessagesRequestInfo
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

			response, err := c.GetMessages(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestStreamClientPutMessages(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "PutMessages")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PutMessages is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "Stream", "PutMessages", createStreamClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamClient)

	body, err := testClient.getRequests("streaming", "PutMessages")
	assert.NoError(t, err)

	type PutMessagesRequestInfo struct {
		ContainerId string
		Request     streaming.PutMessagesRequest
	}

	var requests []PutMessagesRequestInfo
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

			response, err := c.PutMessages(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestStreamClientUpdateGroup(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "UpdateGroup")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateGroup is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "Stream", "UpdateGroup", createStreamClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamClient)

	body, err := testClient.getRequests("streaming", "UpdateGroup")
	assert.NoError(t, err)

	type UpdateGroupRequestInfo struct {
		ContainerId string
		Request     streaming.UpdateGroupRequest
	}

	var requests []UpdateGroupRequestInfo
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

			response, err := c.UpdateGroup(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
