package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/resourcesearch"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createResourceSearchClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := resourcesearch.NewResourceSearchClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="rqs_engg_team_us_grp@oracle.com" jiraProject="RQS" opsJiraProject="RQS"
func TestResourceSearchClientGetResourceType(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcesearch", "GetResourceType")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetResourceType is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcesearch", "ResourceSearch", "GetResourceType", createResourceSearchClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcesearch.ResourceSearchClient)

	body, err := testClient.getRequests("resourcesearch", "GetResourceType")
	assert.NoError(t, err)

	type GetResourceTypeRequestInfo struct {
		ContainerId string
		Request     resourcesearch.GetResourceTypeRequest
	}

	var requests []GetResourceTypeRequestInfo
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

			response, err := c.GetResourceType(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="rqs_engg_team_us_grp@oracle.com" jiraProject="RQS" opsJiraProject="RQS"
func TestResourceSearchClientListResourceTypes(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcesearch", "ListResourceTypes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListResourceTypes is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcesearch", "ResourceSearch", "ListResourceTypes", createResourceSearchClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcesearch.ResourceSearchClient)

	body, err := testClient.getRequests("resourcesearch", "ListResourceTypes")
	assert.NoError(t, err)

	type ListResourceTypesRequestInfo struct {
		ContainerId string
		Request     resourcesearch.ListResourceTypesRequest
	}

	var requests []ListResourceTypesRequestInfo
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
				r := req.(*resourcesearch.ListResourceTypesRequest)
				return c.ListResourceTypes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcesearch.ListResourceTypesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcesearch.ListResourceTypesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="rqs_engg_team_us_grp@oracle.com" jiraProject="RQS" opsJiraProject="RQS"
func TestResourceSearchClientSearchResources(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("resourcesearch", "SearchResources")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SearchResources is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("resourcesearch", "ResourceSearch", "SearchResources", createResourceSearchClientWithProvider)
	assert.NoError(t, err)
	c := cc.(resourcesearch.ResourceSearchClient)

	body, err := testClient.getRequests("resourcesearch", "SearchResources")
	assert.NoError(t, err)

	type SearchResourcesRequestInfo struct {
		ContainerId string
		Request     resourcesearch.SearchResourcesRequest
	}

	var requests []SearchResourcesRequestInfo
	var pr []map[string]interface{}
	err = json.Unmarshal([]byte(body), &pr)
	assert.NoError(t, err)
	requests = make([]SearchResourcesRequestInfo, len(pr))
	polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
	polymorphicRequestInfo["SearchDetails"] =
		PolymorphicRequestUnmarshallingInfo{
			DiscriminatorName: "type",
			DiscriminatorValuesAndTypes: map[string]interface{}{
				"Structured": &resourcesearch.StructuredSearchDetails{},
				"FreeText":   &resourcesearch.FreeTextSearchDetails{},
			},
		}

	for i, ppr := range pr {
		conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
	}

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*resourcesearch.SearchResourcesRequest)
				return c.SearchResources(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]resourcesearch.SearchResourcesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(resourcesearch.SearchResourcesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
