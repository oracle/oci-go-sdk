package autotest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/loadbalancer"

	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientCreateBackend(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "CreateBackend")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateBackend is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "CreateBackend")
	assert.NoError(t, err)

	type CreateBackendRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreateBackendRequest
	}

	var requests []CreateBackendRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateBackend(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientCreateBackendSet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "CreateBackendSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateBackendSet is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "CreateBackendSet")
	assert.NoError(t, err)

	type CreateBackendSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreateBackendSetRequest
	}

	var requests []CreateBackendSetRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateBackendSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientCreateCertificate(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "CreateCertificate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateCertificate is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "CreateCertificate")
	assert.NoError(t, err)

	type CreateCertificateRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreateCertificateRequest
	}

	var requests []CreateCertificateRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateCertificate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientCreateHostname(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "CreateHostname")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateHostname is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "CreateHostname")
	assert.NoError(t, err)

	type CreateHostnameRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreateHostnameRequest
	}

	var requests []CreateHostnameRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateHostname(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientCreateListener(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "CreateListener")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateListener is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "CreateListener")
	assert.NoError(t, err)

	type CreateListenerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreateListenerRequest
	}

	var requests []CreateListenerRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateListener(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientCreateLoadBalancer(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "CreateLoadBalancer")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreateLoadBalancer is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "CreateLoadBalancer")
	assert.NoError(t, err)

	type CreateLoadBalancerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreateLoadBalancerRequest
	}

	var requests []CreateLoadBalancerRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreateLoadBalancer(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientCreatePathRouteSet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "CreatePathRouteSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("CreatePathRouteSet is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "CreatePathRouteSet")
	assert.NoError(t, err)

	type CreatePathRouteSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.CreatePathRouteSetRequest
	}

	var requests []CreatePathRouteSetRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.CreatePathRouteSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientDeleteBackend(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "DeleteBackend")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteBackend is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "DeleteBackend")
	assert.NoError(t, err)

	type DeleteBackendRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeleteBackendRequest
	}

	var requests []DeleteBackendRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteBackend(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientDeleteBackendSet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "DeleteBackendSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteBackendSet is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "DeleteBackendSet")
	assert.NoError(t, err)

	type DeleteBackendSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeleteBackendSetRequest
	}

	var requests []DeleteBackendSetRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteBackendSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientDeleteCertificate(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "DeleteCertificate")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteCertificate is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "DeleteCertificate")
	assert.NoError(t, err)

	type DeleteCertificateRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeleteCertificateRequest
	}

	var requests []DeleteCertificateRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteCertificate(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientDeleteHostname(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "DeleteHostname")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteHostname is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "DeleteHostname")
	assert.NoError(t, err)

	type DeleteHostnameRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeleteHostnameRequest
	}

	var requests []DeleteHostnameRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteHostname(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientDeleteListener(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "DeleteListener")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteListener is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "DeleteListener")
	assert.NoError(t, err)

	type DeleteListenerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeleteListenerRequest
	}

	var requests []DeleteListenerRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteListener(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientDeleteLoadBalancer(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "DeleteLoadBalancer")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeleteLoadBalancer is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "DeleteLoadBalancer")
	assert.NoError(t, err)

	type DeleteLoadBalancerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeleteLoadBalancerRequest
	}

	var requests []DeleteLoadBalancerRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeleteLoadBalancer(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientDeletePathRouteSet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "DeletePathRouteSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DeletePathRouteSet is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "DeletePathRouteSet")
	assert.NoError(t, err)

	type DeletePathRouteSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.DeletePathRouteSetRequest
	}

	var requests []DeletePathRouteSetRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DeletePathRouteSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientGetBackend(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "GetBackend")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBackend is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "GetBackend")
	assert.NoError(t, err)

	type GetBackendRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetBackendRequest
	}

	var requests []GetBackendRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetBackend(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientGetBackendHealth(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "GetBackendHealth")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBackendHealth is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "GetBackendHealth")
	assert.NoError(t, err)

	type GetBackendHealthRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetBackendHealthRequest
	}

	var requests []GetBackendHealthRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetBackendHealth(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientGetBackendSet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "GetBackendSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBackendSet is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "GetBackendSet")
	assert.NoError(t, err)

	type GetBackendSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetBackendSetRequest
	}

	var requests []GetBackendSetRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetBackendSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientGetBackendSetHealth(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "GetBackendSetHealth")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetBackendSetHealth is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "GetBackendSetHealth")
	assert.NoError(t, err)

	type GetBackendSetHealthRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetBackendSetHealthRequest
	}

	var requests []GetBackendSetHealthRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetBackendSetHealth(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientGetHealthChecker(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "GetHealthChecker")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetHealthChecker is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "GetHealthChecker")
	assert.NoError(t, err)

	type GetHealthCheckerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetHealthCheckerRequest
	}

	var requests []GetHealthCheckerRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetHealthChecker(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientGetHostname(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "GetHostname")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetHostname is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "GetHostname")
	assert.NoError(t, err)

	type GetHostnameRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetHostnameRequest
	}

	var requests []GetHostnameRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetHostname(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientGetLoadBalancer(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "GetLoadBalancer")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLoadBalancer is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "GetLoadBalancer")
	assert.NoError(t, err)

	type GetLoadBalancerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetLoadBalancerRequest
	}

	var requests []GetLoadBalancerRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetLoadBalancer(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientGetLoadBalancerHealth(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "GetLoadBalancerHealth")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetLoadBalancerHealth is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "GetLoadBalancerHealth")
	assert.NoError(t, err)

	type GetLoadBalancerHealthRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetLoadBalancerHealthRequest
	}

	var requests []GetLoadBalancerHealthRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetLoadBalancerHealth(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientGetPathRouteSet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "GetPathRouteSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetPathRouteSet is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "GetPathRouteSet")
	assert.NoError(t, err)

	type GetPathRouteSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetPathRouteSetRequest
	}

	var requests []GetPathRouteSetRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetPathRouteSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientGetWorkRequest(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "GetWorkRequest")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("GetWorkRequest is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "GetWorkRequest")
	assert.NoError(t, err)

	type GetWorkRequestRequestInfo struct {
		ContainerId string
		Request     loadbalancer.GetWorkRequestRequest
	}

	var requests []GetWorkRequestRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.GetWorkRequest(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientListBackendSets(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "ListBackendSets")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBackendSets is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "ListBackendSets")
	assert.NoError(t, err)

	type ListBackendSetsRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListBackendSetsRequest
	}

	var requests []ListBackendSetsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ListBackendSets(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientListBackends(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "ListBackends")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListBackends is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "ListBackends")
	assert.NoError(t, err)

	type ListBackendsRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListBackendsRequest
	}

	var requests []ListBackendsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ListBackends(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientListCertificates(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "ListCertificates")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListCertificates is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "ListCertificates")
	assert.NoError(t, err)

	type ListCertificatesRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListCertificatesRequest
	}

	var requests []ListCertificatesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ListCertificates(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientListHostnames(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "ListHostnames")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListHostnames is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "ListHostnames")
	assert.NoError(t, err)

	type ListHostnamesRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListHostnamesRequest
	}

	var requests []ListHostnamesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ListHostnames(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientListLoadBalancerHealths(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "ListLoadBalancerHealths")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLoadBalancerHealths is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "ListLoadBalancerHealths")
	assert.NoError(t, err)

	type ListLoadBalancerHealthsRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListLoadBalancerHealthsRequest
	}

	var requests []ListLoadBalancerHealthsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*loadbalancer.ListLoadBalancerHealthsRequest)
				return c.ListLoadBalancerHealths(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loadbalancer.ListLoadBalancerHealthsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loadbalancer.ListLoadBalancerHealthsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientListLoadBalancers(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "ListLoadBalancers")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListLoadBalancers is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "ListLoadBalancers")
	assert.NoError(t, err)

	type ListLoadBalancersRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListLoadBalancersRequest
	}

	var requests []ListLoadBalancersRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*loadbalancer.ListLoadBalancersRequest)
				return c.ListLoadBalancers(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loadbalancer.ListLoadBalancersResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loadbalancer.ListLoadBalancersResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientListPathRouteSets(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "ListPathRouteSets")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPathRouteSets is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "ListPathRouteSets")
	assert.NoError(t, err)

	type ListPathRouteSetsRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListPathRouteSetsRequest
	}

	var requests []ListPathRouteSetsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.ListPathRouteSets(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientListPolicies(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "ListPolicies")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListPolicies is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "ListPolicies")
	assert.NoError(t, err)

	type ListPoliciesRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListPoliciesRequest
	}

	var requests []ListPoliciesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*loadbalancer.ListPoliciesRequest)
				return c.ListPolicies(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loadbalancer.ListPoliciesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loadbalancer.ListPoliciesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientListProtocols(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "ListProtocols")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListProtocols is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "ListProtocols")
	assert.NoError(t, err)

	type ListProtocolsRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListProtocolsRequest
	}

	var requests []ListProtocolsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*loadbalancer.ListProtocolsRequest)
				return c.ListProtocols(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loadbalancer.ListProtocolsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loadbalancer.ListProtocolsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientListShapes(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "ListShapes")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListShapes is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "ListShapes")
	assert.NoError(t, err)

	type ListShapesRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListShapesRequest
	}

	var requests []ListShapesRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*loadbalancer.ListShapesRequest)
				return c.ListShapes(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loadbalancer.ListShapesResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loadbalancer.ListShapesResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientListWorkRequests(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "ListWorkRequests")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("ListWorkRequests is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "ListWorkRequests")
	assert.NoError(t, err)

	type ListWorkRequestsRequestInfo struct {
		ContainerId string
		Request     loadbalancer.ListWorkRequestsRequest
	}

	var requests []ListWorkRequestsRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, request := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			request.Request.RequestMetadata.RetryPolicy = retryPolicy
			listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
				r := req.(*loadbalancer.ListWorkRequestsRequest)
				return c.ListWorkRequests(context.Background(), *r)
			}

			listResponses, err := testClient.generateListResponses(&request.Request, listFn)
			typedListResponses := make([]loadbalancer.ListWorkRequestsResponse, len(listResponses))
			for i, lr := range listResponses {
				typedListResponses[i] = lr.(loadbalancer.ListWorkRequestsResponse)
			}

			message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientUpdateBackend(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateBackend")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateBackend is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "UpdateBackend")
	assert.NoError(t, err)

	type UpdateBackendRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateBackendRequest
	}

	var requests []UpdateBackendRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateBackend(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientUpdateBackendSet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateBackendSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateBackendSet is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "UpdateBackendSet")
	assert.NoError(t, err)

	type UpdateBackendSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateBackendSetRequest
	}

	var requests []UpdateBackendSetRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateBackendSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientUpdateHealthChecker(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateHealthChecker")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateHealthChecker is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "UpdateHealthChecker")
	assert.NoError(t, err)

	type UpdateHealthCheckerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateHealthCheckerRequest
	}

	var requests []UpdateHealthCheckerRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateHealthChecker(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientUpdateHostname(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateHostname")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateHostname is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "UpdateHostname")
	assert.NoError(t, err)

	type UpdateHostnameRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateHostnameRequest
	}

	var requests []UpdateHostnameRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateHostname(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientUpdateListener(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateListener")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateListener is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "UpdateListener")
	assert.NoError(t, err)

	type UpdateListenerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateListenerRequest
	}

	var requests []UpdateListenerRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateListener(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientUpdateLoadBalancer(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdateLoadBalancer")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdateLoadBalancer is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "UpdateLoadBalancer")
	assert.NoError(t, err)

	type UpdateLoadBalancerRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdateLoadBalancerRequest
	}

	var requests []UpdateLoadBalancerRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdateLoadBalancer(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestLoadBalancerClientUpdatePathRouteSet(t *testing.T) {
	enabled, err := testClient.isApiEnabled("loadbalancer", "UpdatePathRouteSet")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("UpdatePathRouteSet is not enabled by the testing service")
	}
	c, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(testConfig.ConfigurationProvider)
	assert.NoError(t, err)

	body, err := testClient.getRequests("loadbalancer", "UpdatePathRouteSet")
	assert.NoError(t, err)

	type UpdatePathRouteSetRequestInfo struct {
		ContainerId string
		Request     loadbalancer.UpdatePathRouteSetRequest
	}

	var requests []UpdatePathRouteSetRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.UpdatePathRouteSet(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}
