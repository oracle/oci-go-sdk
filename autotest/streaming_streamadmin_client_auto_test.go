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

func createStreamAdminClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := streaming.NewStreamAdminClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestStreamAdminClientCreateStream(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "CreateStream")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateStream is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "CreateStream", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "CreateStream")
	assert.NoError(t, err)

	type CreateStreamRequestInfo struct {
		ContainerId string
		Request     streaming.CreateStreamRequest
	}

	var requests []CreateStreamRequestInfo
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

			response, err := c.CreateStream(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestStreamAdminClientDeleteStream(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "DeleteStream")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteStream is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "DeleteStream", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "DeleteStream")
	assert.NoError(t, err)

	type DeleteStreamRequestInfo struct {
		ContainerId string
		Request     streaming.DeleteStreamRequest
	}

	var requests []DeleteStreamRequestInfo
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

			response, err := c.DeleteStream(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestStreamAdminClientGetStream(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "GetStream")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetStream is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "GetStream", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "GetStream")
	assert.NoError(t, err)

	type GetStreamRequestInfo struct {
		ContainerId string
		Request     streaming.GetStreamRequest
	}

	var requests []GetStreamRequestInfo
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

			response, err := c.GetStream(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestStreamAdminClientListStreams(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "ListStreams")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListStreams is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "ListStreams", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "ListStreams")
	assert.NoError(t, err)

	type ListStreamsRequestInfo struct {
		ContainerId string
		Request     streaming.ListStreamsRequest
	}

	var requests []ListStreamsRequestInfo
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
				r := req.(*streaming.ListStreamsRequest)
				return c.ListStreams(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]streaming.ListStreamsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(streaming.ListStreamsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestStreamAdminClientUpdateStream(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("streaming", "UpdateStream")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateStream is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("streaming", "StreamAdmin", "UpdateStream", createStreamAdminClientWithProvider)
	assert.NoError(t, err)
	c := cc.(streaming.StreamAdminClient)

	body, err := testClient.getRequests("streaming", "UpdateStream")
	assert.NoError(t, err)

	type UpdateStreamRequestInfo struct {
		ContainerId string
		Request     streaming.UpdateStreamRequest
	}

	var requests []UpdateStreamRequestInfo
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

			response, err := c.UpdateStream(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
