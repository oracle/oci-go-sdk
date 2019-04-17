package autotest

import (
	"github.com/oracle/oci-go-sdk/budget"
	"github.com/oracle/oci-go-sdk/common"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createBudgetClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := budget.NewBudgetClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="plat_compartments_us_grp@oracle.com" jiraProject="COMP" opsJiraProject="COMP"
func TestBudgetClientCreateAlertRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("budget", "CreateAlertRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAlertRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("budget", "Budget", "CreateAlertRule", createBudgetClientWithProvider)
	assert.NoError(t, err)
	c := cc.(budget.BudgetClient)

	body, err := testClient.getRequests("budget", "CreateAlertRule")
	assert.NoError(t, err)

	type CreateAlertRuleRequestInfo struct {
		ContainerId string
		Request     budget.CreateAlertRuleRequest
	}

	var requests []CreateAlertRuleRequestInfo
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

			response, err := c.CreateAlertRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="plat_compartments_us_grp@oracle.com" jiraProject="COMP" opsJiraProject="COMP"
func TestBudgetClientCreateBudget(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("budget", "CreateBudget")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateBudget is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("budget", "Budget", "CreateBudget", createBudgetClientWithProvider)
	assert.NoError(t, err)
	c := cc.(budget.BudgetClient)

	body, err := testClient.getRequests("budget", "CreateBudget")
	assert.NoError(t, err)

	type CreateBudgetRequestInfo struct {
		ContainerId string
		Request     budget.CreateBudgetRequest
	}

	var requests []CreateBudgetRequestInfo
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

			response, err := c.CreateBudget(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="plat_compartments_us_grp@oracle.com" jiraProject="COMP" opsJiraProject="COMP"
func TestBudgetClientDeleteAlertRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("budget", "DeleteAlertRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAlertRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("budget", "Budget", "DeleteAlertRule", createBudgetClientWithProvider)
	assert.NoError(t, err)
	c := cc.(budget.BudgetClient)

	body, err := testClient.getRequests("budget", "DeleteAlertRule")
	assert.NoError(t, err)

	type DeleteAlertRuleRequestInfo struct {
		ContainerId string
		Request     budget.DeleteAlertRuleRequest
	}

	var requests []DeleteAlertRuleRequestInfo
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

			response, err := c.DeleteAlertRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="plat_compartments_us_grp@oracle.com" jiraProject="COMP" opsJiraProject="COMP"
func TestBudgetClientDeleteBudget(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("budget", "DeleteBudget")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteBudget is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("budget", "Budget", "DeleteBudget", createBudgetClientWithProvider)
	assert.NoError(t, err)
	c := cc.(budget.BudgetClient)

	body, err := testClient.getRequests("budget", "DeleteBudget")
	assert.NoError(t, err)

	type DeleteBudgetRequestInfo struct {
		ContainerId string
		Request     budget.DeleteBudgetRequest
	}

	var requests []DeleteBudgetRequestInfo
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

			response, err := c.DeleteBudget(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="plat_compartments_us_grp@oracle.com" jiraProject="COMP" opsJiraProject="COMP"
func TestBudgetClientGetAlertRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("budget", "GetAlertRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAlertRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("budget", "Budget", "GetAlertRule", createBudgetClientWithProvider)
	assert.NoError(t, err)
	c := cc.(budget.BudgetClient)

	body, err := testClient.getRequests("budget", "GetAlertRule")
	assert.NoError(t, err)

	type GetAlertRuleRequestInfo struct {
		ContainerId string
		Request     budget.GetAlertRuleRequest
	}

	var requests []GetAlertRuleRequestInfo
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

			response, err := c.GetAlertRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="plat_compartments_us_grp@oracle.com" jiraProject="COMP" opsJiraProject="COMP"
func TestBudgetClientGetBudget(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("budget", "GetBudget")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBudget is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("budget", "Budget", "GetBudget", createBudgetClientWithProvider)
	assert.NoError(t, err)
	c := cc.(budget.BudgetClient)

	body, err := testClient.getRequests("budget", "GetBudget")
	assert.NoError(t, err)

	type GetBudgetRequestInfo struct {
		ContainerId string
		Request     budget.GetBudgetRequest
	}

	var requests []GetBudgetRequestInfo
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

			response, err := c.GetBudget(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="plat_compartments_us_grp@oracle.com" jiraProject="COMP" opsJiraProject="COMP"
func TestBudgetClientListAlertRules(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("budget", "ListAlertRules")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAlertRules is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("budget", "Budget", "ListAlertRules", createBudgetClientWithProvider)
	assert.NoError(t, err)
	c := cc.(budget.BudgetClient)

	body, err := testClient.getRequests("budget", "ListAlertRules")
	assert.NoError(t, err)

	type ListAlertRulesRequestInfo struct {
		ContainerId string
		Request     budget.ListAlertRulesRequest
	}

	var requests []ListAlertRulesRequestInfo
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
				r := req.(*budget.ListAlertRulesRequest)
				return c.ListAlertRules(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]budget.ListAlertRulesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(budget.ListAlertRulesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="plat_compartments_us_grp@oracle.com" jiraProject="COMP" opsJiraProject="COMP"
func TestBudgetClientListBudgets(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("budget", "ListBudgets")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBudgets is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("budget", "Budget", "ListBudgets", createBudgetClientWithProvider)
	assert.NoError(t, err)
	c := cc.(budget.BudgetClient)

	body, err := testClient.getRequests("budget", "ListBudgets")
	assert.NoError(t, err)

	type ListBudgetsRequestInfo struct {
		ContainerId string
		Request     budget.ListBudgetsRequest
	}

	var requests []ListBudgetsRequestInfo
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
				r := req.(*budget.ListBudgetsRequest)
				return c.ListBudgets(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]budget.ListBudgetsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(budget.ListBudgetsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="plat_compartments_us_grp@oracle.com" jiraProject="COMP" opsJiraProject="COMP"
func TestBudgetClientUpdateAlertRule(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("budget", "UpdateAlertRule")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAlertRule is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("budget", "Budget", "UpdateAlertRule", createBudgetClientWithProvider)
	assert.NoError(t, err)
	c := cc.(budget.BudgetClient)

	body, err := testClient.getRequests("budget", "UpdateAlertRule")
	assert.NoError(t, err)

	type UpdateAlertRuleRequestInfo struct {
		ContainerId string
		Request     budget.UpdateAlertRuleRequest
	}

	var requests []UpdateAlertRuleRequestInfo
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

			response, err := c.UpdateAlertRule(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="plat_compartments_us_grp@oracle.com" jiraProject="COMP" opsJiraProject="COMP"
func TestBudgetClientUpdateBudget(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("budget", "UpdateBudget")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateBudget is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("budget", "Budget", "UpdateBudget", createBudgetClientWithProvider)
	assert.NoError(t, err)
	c := cc.(budget.BudgetClient)

	body, err := testClient.getRequests("budget", "UpdateBudget")
	assert.NoError(t, err)

	type UpdateBudgetRequestInfo struct {
		ContainerId string
		Request     budget.UpdateBudgetRequest
	}

	var requests []UpdateBudgetRequestInfo
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

			response, err := c.UpdateBudget(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
