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

func createNotificationDataPlaneClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := ons.NewNotificationDataPlaneClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestNotificationDataPlaneClientCreateSubscription(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ons", "CreateSubscription")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateSubscription is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ons", "NotificationDataPlane", "CreateSubscription", createNotificationDataPlaneClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ons.NotificationDataPlaneClient)

	body, err := testClient.getRequests("ons", "CreateSubscription")
	assert.NoError(t, err)

	type CreateSubscriptionRequestInfo struct {
		ContainerId string
		Request     ons.CreateSubscriptionRequest
	}

	var requests []CreateSubscriptionRequestInfo
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

			response, err := c.CreateSubscription(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestNotificationDataPlaneClientDeleteSubscription(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ons", "DeleteSubscription")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteSubscription is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ons", "NotificationDataPlane", "DeleteSubscription", createNotificationDataPlaneClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ons.NotificationDataPlaneClient)

	body, err := testClient.getRequests("ons", "DeleteSubscription")
	assert.NoError(t, err)

	type DeleteSubscriptionRequestInfo struct {
		ContainerId string
		Request     ons.DeleteSubscriptionRequest
	}

	var requests []DeleteSubscriptionRequestInfo
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

			response, err := c.DeleteSubscription(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestNotificationDataPlaneClientGetConfirmSubscription(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ons", "GetConfirmSubscription")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetConfirmSubscription is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ons", "NotificationDataPlane", "GetConfirmSubscription", createNotificationDataPlaneClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ons.NotificationDataPlaneClient)

	body, err := testClient.getRequests("ons", "GetConfirmSubscription")
	assert.NoError(t, err)

	type GetConfirmSubscriptionRequestInfo struct {
		ContainerId string
		Request     ons.GetConfirmSubscriptionRequest
	}

	var requests []GetConfirmSubscriptionRequestInfo
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

			response, err := c.GetConfirmSubscription(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestNotificationDataPlaneClientGetSubscription(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ons", "GetSubscription")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetSubscription is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ons", "NotificationDataPlane", "GetSubscription", createNotificationDataPlaneClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ons.NotificationDataPlaneClient)

	body, err := testClient.getRequests("ons", "GetSubscription")
	assert.NoError(t, err)

	type GetSubscriptionRequestInfo struct {
		ContainerId string
		Request     ons.GetSubscriptionRequest
	}

	var requests []GetSubscriptionRequestInfo
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

			response, err := c.GetSubscription(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestNotificationDataPlaneClientGetUnsubscription(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ons", "GetUnsubscription")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetUnsubscription is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ons", "NotificationDataPlane", "GetUnsubscription", createNotificationDataPlaneClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ons.NotificationDataPlaneClient)

	body, err := testClient.getRequests("ons", "GetUnsubscription")
	assert.NoError(t, err)

	type GetUnsubscriptionRequestInfo struct {
		ContainerId string
		Request     ons.GetUnsubscriptionRequest
	}

	var requests []GetUnsubscriptionRequestInfo
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

			response, err := c.GetUnsubscription(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestNotificationDataPlaneClientListSubscriptions(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ons", "ListSubscriptions")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListSubscriptions is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ons", "NotificationDataPlane", "ListSubscriptions", createNotificationDataPlaneClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ons.NotificationDataPlaneClient)

	body, err := testClient.getRequests("ons", "ListSubscriptions")
	assert.NoError(t, err)

	type ListSubscriptionsRequestInfo struct {
		ContainerId string
		Request     ons.ListSubscriptionsRequest
	}

	var requests []ListSubscriptionsRequestInfo
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
				r := req.(*ons.ListSubscriptionsRequest)
				return c.ListSubscriptions(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]ons.ListSubscriptionsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(ons.ListSubscriptionsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestNotificationDataPlaneClientPublishMessage(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ons", "PublishMessage")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PublishMessage is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ons", "NotificationDataPlane", "PublishMessage", createNotificationDataPlaneClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ons.NotificationDataPlaneClient)

	body, err := testClient.getRequests("ons", "PublishMessage")
	assert.NoError(t, err)

	type PublishMessageRequestInfo struct {
		ContainerId string
		Request     ons.PublishMessageRequest
	}

	var requests []PublishMessageRequestInfo
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

			response, err := c.PublishMessage(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestNotificationDataPlaneClientResendSubscriptionConfirmation(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ons", "ResendSubscriptionConfirmation")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ResendSubscriptionConfirmation is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ons", "NotificationDataPlane", "ResendSubscriptionConfirmation", createNotificationDataPlaneClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ons.NotificationDataPlaneClient)

	body, err := testClient.getRequests("ons", "ResendSubscriptionConfirmation")
	assert.NoError(t, err)

	type ResendSubscriptionConfirmationRequestInfo struct {
		ContainerId string
		Request     ons.ResendSubscriptionConfirmationRequest
	}

	var requests []ResendSubscriptionConfirmationRequestInfo
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

			response, err := c.ResendSubscriptionConfirmation(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestNotificationDataPlaneClientUpdateSubscription(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("ons", "UpdateSubscription")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateSubscription is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("ons", "NotificationDataPlane", "UpdateSubscription", createNotificationDataPlaneClientWithProvider)
	assert.NoError(t, err)
	c := cc.(ons.NotificationDataPlaneClient)

	body, err := testClient.getRequests("ons", "UpdateSubscription")
	assert.NoError(t, err)

	type UpdateSubscriptionRequestInfo struct {
		ContainerId string
		Request     ons.UpdateSubscriptionRequest
	}

	var requests []UpdateSubscriptionRequestInfo
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

			response, err := c.UpdateSubscription(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
