package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/monitoring"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createMonitoringClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {

	client, err := monitoring.NewMonitoringClientWithConfigurationProvider(p)
	if testConfig.Endpoint != "" {
		client.Host = testConfig.Endpoint
	} else {
		client.SetRegion(testConfig.Region)
	}
	return client, err
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientCreateAlarm(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("monitoring", "CreateAlarm")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateAlarm is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("monitoring", "Monitoring", "CreateAlarm", createMonitoringClientWithProvider)
	assert.NoError(t, err)
	c := cc.(monitoring.MonitoringClient)

	body, err := testClient.getRequests("monitoring", "CreateAlarm")
	assert.NoError(t, err)

	type CreateAlarmRequestInfo struct {
		ContainerId string
		Request     monitoring.CreateAlarmRequest
	}

	var requests []CreateAlarmRequestInfo
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

			response, err := c.CreateAlarm(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientDeleteAlarm(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("monitoring", "DeleteAlarm")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteAlarm is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("monitoring", "Monitoring", "DeleteAlarm", createMonitoringClientWithProvider)
	assert.NoError(t, err)
	c := cc.(monitoring.MonitoringClient)

	body, err := testClient.getRequests("monitoring", "DeleteAlarm")
	assert.NoError(t, err)

	type DeleteAlarmRequestInfo struct {
		ContainerId string
		Request     monitoring.DeleteAlarmRequest
	}

	var requests []DeleteAlarmRequestInfo
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

			response, err := c.DeleteAlarm(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientGetAlarm(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("monitoring", "GetAlarm")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAlarm is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("monitoring", "Monitoring", "GetAlarm", createMonitoringClientWithProvider)
	assert.NoError(t, err)
	c := cc.(monitoring.MonitoringClient)

	body, err := testClient.getRequests("monitoring", "GetAlarm")
	assert.NoError(t, err)

	type GetAlarmRequestInfo struct {
		ContainerId string
		Request     monitoring.GetAlarmRequest
	}

	var requests []GetAlarmRequestInfo
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

			response, err := c.GetAlarm(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientGetAlarmHistory(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("monitoring", "GetAlarmHistory")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetAlarmHistory is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("monitoring", "Monitoring", "GetAlarmHistory", createMonitoringClientWithProvider)
	assert.NoError(t, err)
	c := cc.(monitoring.MonitoringClient)

	body, err := testClient.getRequests("monitoring", "GetAlarmHistory")
	assert.NoError(t, err)

	type GetAlarmHistoryRequestInfo struct {
		ContainerId string
		Request     monitoring.GetAlarmHistoryRequest
	}

	var requests []GetAlarmHistoryRequestInfo
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
				r := req.(*monitoring.GetAlarmHistoryRequest)
				return c.GetAlarmHistory(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]monitoring.GetAlarmHistoryResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(monitoring.GetAlarmHistoryResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientListAlarms(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("monitoring", "ListAlarms")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAlarms is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("monitoring", "Monitoring", "ListAlarms", createMonitoringClientWithProvider)
	assert.NoError(t, err)
	c := cc.(monitoring.MonitoringClient)

	body, err := testClient.getRequests("monitoring", "ListAlarms")
	assert.NoError(t, err)

	type ListAlarmsRequestInfo struct {
		ContainerId string
		Request     monitoring.ListAlarmsRequest
	}

	var requests []ListAlarmsRequestInfo
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
				r := req.(*monitoring.ListAlarmsRequest)
				return c.ListAlarms(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]monitoring.ListAlarmsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(monitoring.ListAlarmsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientListAlarmsStatus(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("monitoring", "ListAlarmsStatus")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListAlarmsStatus is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("monitoring", "Monitoring", "ListAlarmsStatus", createMonitoringClientWithProvider)
	assert.NoError(t, err)
	c := cc.(monitoring.MonitoringClient)

	body, err := testClient.getRequests("monitoring", "ListAlarmsStatus")
	assert.NoError(t, err)

	type ListAlarmsStatusRequestInfo struct {
		ContainerId string
		Request     monitoring.ListAlarmsStatusRequest
	}

	var requests []ListAlarmsStatusRequestInfo
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
				r := req.(*monitoring.ListAlarmsStatusRequest)
				return c.ListAlarmsStatus(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]monitoring.ListAlarmsStatusResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(monitoring.ListAlarmsStatusResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientListMetrics(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("monitoring", "ListMetrics")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListMetrics is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("monitoring", "Monitoring", "ListMetrics", createMonitoringClientWithProvider)
	assert.NoError(t, err)
	c := cc.(monitoring.MonitoringClient)

	body, err := testClient.getRequests("monitoring", "ListMetrics")
	assert.NoError(t, err)

	type ListMetricsRequestInfo struct {
		ContainerId string
		Request     monitoring.ListMetricsRequest
	}

	var requests []ListMetricsRequestInfo
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
				r := req.(*monitoring.ListMetricsRequest)
				return c.ListMetrics(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]monitoring.ListMetricsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(monitoring.ListMetricsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientPostMetricData(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("monitoring", "PostMetricData")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("PostMetricData is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("monitoring", "Monitoring", "PostMetricData", createMonitoringClientWithProvider)
	assert.NoError(t, err)
	c := cc.(monitoring.MonitoringClient)

	body, err := testClient.getRequests("monitoring", "PostMetricData")
	assert.NoError(t, err)

	type PostMetricDataRequestInfo struct {
		ContainerId string
		Request     monitoring.PostMetricDataRequest
	}

	var requests []PostMetricDataRequestInfo
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

			response, err := c.PostMetricData(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientRemoveAlarmSuppression(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("monitoring", "RemoveAlarmSuppression")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("RemoveAlarmSuppression is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("monitoring", "Monitoring", "RemoveAlarmSuppression", createMonitoringClientWithProvider)
	assert.NoError(t, err)
	c := cc.(monitoring.MonitoringClient)

	body, err := testClient.getRequests("monitoring", "RemoveAlarmSuppression")
	assert.NoError(t, err)

	type RemoveAlarmSuppressionRequestInfo struct {
		ContainerId string
		Request     monitoring.RemoveAlarmSuppressionRequest
	}

	var requests []RemoveAlarmSuppressionRequestInfo
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

			response, err := c.RemoveAlarmSuppression(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientSummarizeMetricsData(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("monitoring", "SummarizeMetricsData")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("SummarizeMetricsData is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("monitoring", "Monitoring", "SummarizeMetricsData", createMonitoringClientWithProvider)
	assert.NoError(t, err)
	c := cc.(monitoring.MonitoringClient)

	body, err := testClient.getRequests("monitoring", "SummarizeMetricsData")
	assert.NoError(t, err)

	type SummarizeMetricsDataRequestInfo struct {
		ContainerId string
		Request     monitoring.SummarizeMetricsDataRequest
	}

	var requests []SummarizeMetricsDataRequestInfo
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

			response, err := c.SummarizeMetricsData(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="pic_ion_dev_grp@oracle.com" jiraProject="https://jira.oci.oraclecorp.com/projects/TEL" opsJiraProject="https://jira-sd.mc1.oracleiaas.com/projects/TEL"
func TestMonitoringClientUpdateAlarm(t *testing.T) {
	defer failTestOnPanic(t)

	enabled, err := testClient.isApiEnabled("monitoring", "UpdateAlarm")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateAlarm is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("monitoring", "Monitoring", "UpdateAlarm", createMonitoringClientWithProvider)
	assert.NoError(t, err)
	c := cc.(monitoring.MonitoringClient)

	body, err := testClient.getRequests("monitoring", "UpdateAlarm")
	assert.NoError(t, err)

	type UpdateAlarmRequestInfo struct {
		ContainerId string
		Request     monitoring.UpdateAlarmRequest
	}

	var requests []UpdateAlarmRequestInfo
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

			response, err := c.UpdateAlarm(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
