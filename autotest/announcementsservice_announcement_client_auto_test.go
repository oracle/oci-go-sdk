package autotest

import (
	"github.com/oracle/oci-go-sdk/announcementsservice"
	"github.com/oracle/oci-go-sdk/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createAnnouncementClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := announcementsservice.NewAnnouncementClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestAnnouncementClientGetAnnouncement(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("announcementsservice", "GetAnnouncement")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAnnouncement is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("announcementsservice", "Announcement", "GetAnnouncement", createAnnouncementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(announcementsservice.AnnouncementClient)

	body, err := testClient.getRequests("announcementsservice", "GetAnnouncement")
	assert.NoError(t, err)

	type GetAnnouncementRequestInfo struct {
		ContainerId string
		Request     announcementsservice.GetAnnouncementRequest
	}

	var requests []GetAnnouncementRequestInfo
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

			response, err := c.GetAnnouncement(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestAnnouncementClientGetAnnouncementUserStatus(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("announcementsservice", "GetAnnouncementUserStatus")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAnnouncementUserStatus is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("announcementsservice", "Announcement", "GetAnnouncementUserStatus", createAnnouncementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(announcementsservice.AnnouncementClient)

	body, err := testClient.getRequests("announcementsservice", "GetAnnouncementUserStatus")
	assert.NoError(t, err)

	type GetAnnouncementUserStatusRequestInfo struct {
		ContainerId string
		Request     announcementsservice.GetAnnouncementUserStatusRequest
	}

	var requests []GetAnnouncementUserStatusRequestInfo
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

			response, err := c.GetAnnouncementUserStatus(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestAnnouncementClientListAnnouncements(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("announcementsservice", "ListAnnouncements")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAnnouncements is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("announcementsservice", "Announcement", "ListAnnouncements", createAnnouncementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(announcementsservice.AnnouncementClient)

	body, err := testClient.getRequests("announcementsservice", "ListAnnouncements")
	assert.NoError(t, err)

	type ListAnnouncementsRequestInfo struct {
		ContainerId string
		Request     announcementsservice.ListAnnouncementsRequest
	}

	var requests []ListAnnouncementsRequestInfo
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
				r := req.(*announcementsservice.ListAnnouncementsRequest)
				return c.ListAnnouncements(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]announcementsservice.ListAnnouncementsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(announcementsservice.ListAnnouncementsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestAnnouncementClientUpdateAnnouncementUserStatus(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("announcementsservice", "UpdateAnnouncementUserStatus")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAnnouncementUserStatus is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("announcementsservice", "Announcement", "UpdateAnnouncementUserStatus", createAnnouncementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(announcementsservice.AnnouncementClient)

	body, err := testClient.getRequests("announcementsservice", "UpdateAnnouncementUserStatus")
	assert.NoError(t, err)

	type UpdateAnnouncementUserStatusRequestInfo struct {
		ContainerId string
		Request     announcementsservice.UpdateAnnouncementUserStatusRequest
	}

	var requests []UpdateAnnouncementUserStatusRequestInfo
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

			response, err := c.UpdateAnnouncementUserStatus(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
