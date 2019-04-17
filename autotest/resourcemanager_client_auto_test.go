package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/resourcemanager"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createResourceManagerClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientCancelJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "CancelJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CancelJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "CancelJob", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "CancelJob")
	assert.NoError(t, err)

	type CancelJobRequestInfo struct {
		ContainerId string
		Request     resourcemanager.CancelJobRequest
	}

	var requests []CancelJobRequestInfo
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

			response, err := c.CancelJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientCreateJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "CreateJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "CreateJob", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "CreateJob")
	assert.NoError(t, err)

	type CreateJobRequestInfo struct {
		ContainerId string
		Request     resourcemanager.CreateJobRequest
	}

	var requests []CreateJobRequestInfo
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

			response, err := c.CreateJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientCreateStack(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "CreateStack")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateStack is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "CreateStack", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "CreateStack")
	assert.NoError(t, err)

	type CreateStackRequestInfo struct {
		ContainerId string
		Request     resourcemanager.CreateStackRequest
	}

	var requests []CreateStackRequestInfo
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

			response, err := c.CreateStack(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientDeleteStack(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "DeleteStack")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteStack is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "DeleteStack", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "DeleteStack")
	assert.NoError(t, err)

	type DeleteStackRequestInfo struct {
		ContainerId string
		Request     resourcemanager.DeleteStackRequest
	}

	var requests []DeleteStackRequestInfo
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

			response, err := c.DeleteStack(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetJob", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetJob")
	assert.NoError(t, err)

	type GetJobRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetJobRequest
	}

	var requests []GetJobRequestInfo
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

			response, err := c.GetJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetJobLogs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetJobLogs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJobLogs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetJobLogs", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetJobLogs")
	assert.NoError(t, err)

	type GetJobLogsRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetJobLogsRequest
	}

	var requests []GetJobLogsRequestInfo
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
				r := req.(*resourcemanager.GetJobLogsRequest)
				return c.GetJobLogs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcemanager.GetJobLogsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcemanager.GetJobLogsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetJobLogsContent(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetJobLogsContent")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJobLogsContent is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetJobLogsContent", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetJobLogsContent")
	assert.NoError(t, err)

	type GetJobLogsContentRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetJobLogsContentRequest
	}

	var requests []GetJobLogsContentRequestInfo
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

			response, err := c.GetJobLogsContent(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetJobTfConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetJobTfConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJobTfConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetJobTfConfig", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetJobTfConfig")
	assert.NoError(t, err)

	type GetJobTfConfigRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetJobTfConfigRequest
	}

	var requests []GetJobTfConfigRequestInfo
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

			response, err := c.GetJobTfConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetJobTfState(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetJobTfState")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetJobTfState is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetJobTfState", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetJobTfState")
	assert.NoError(t, err)

	type GetJobTfStateRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetJobTfStateRequest
	}

	var requests []GetJobTfStateRequestInfo
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

			response, err := c.GetJobTfState(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetStack(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetStack")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetStack is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetStack", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetStack")
	assert.NoError(t, err)

	type GetStackRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetStackRequest
	}

	var requests []GetStackRequestInfo
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

			response, err := c.GetStack(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientGetStackTfConfig(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "GetStackTfConfig")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetStackTfConfig is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "GetStackTfConfig", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "GetStackTfConfig")
	assert.NoError(t, err)

	type GetStackTfConfigRequestInfo struct {
		ContainerId string
		Request     resourcemanager.GetStackTfConfigRequest
	}

	var requests []GetStackTfConfigRequestInfo
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

			response, err := c.GetStackTfConfig(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientListJobs(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "ListJobs")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListJobs is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "ListJobs", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "ListJobs")
	assert.NoError(t, err)

	type ListJobsRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ListJobsRequest
	}

	var requests []ListJobsRequestInfo
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
				r := req.(*resourcemanager.ListJobsRequest)
				return c.ListJobs(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcemanager.ListJobsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcemanager.ListJobsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientListStacks(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "ListStacks")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListStacks is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "ListStacks", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "ListStacks")
	assert.NoError(t, err)

	type ListStacksRequestInfo struct {
		ContainerId string
		Request     resourcemanager.ListStacksRequest
	}

	var requests []ListStacksRequestInfo
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
				r := req.(*resourcemanager.ListStacksRequest)
				return c.ListStacks(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcemanager.ListStacksResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcemanager.ListStacksResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientUpdateJob(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "UpdateJob")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateJob is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "UpdateJob", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "UpdateJob")
	assert.NoError(t, err)

	type UpdateJobRequestInfo struct {
		ContainerId string
		Request     resourcemanager.UpdateJobRequest
	}

	var requests []UpdateJobRequestInfo
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

			response, err := c.UpdateJob(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="team_oci_orm_us_grp@oracle.com" jiraProject="ORCH" opsJiraProject="OS"
func TestResourceManagerClientUpdateStack(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcemanager", "UpdateStack")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateStack is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcemanager", "ResourceManager", "UpdateStack", createResourceManagerClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcemanager.ResourceManagerClient)

	body, err := testClient.getRequests("resourcemanager", "UpdateStack")
	assert.NoError(t, err)

	type UpdateStackRequestInfo struct {
		ContainerId string
		Request     resourcemanager.UpdateStackRequest
	}

	var requests []UpdateStackRequestInfo
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

			response, err := c.UpdateStack(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
