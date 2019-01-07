package autotest

import (
    "github.com/oracle/oci-go-sdk/containerengine"
    "github.com/oracle/oci-go-sdk/common"

    "context"
    "encoding/json"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func createContainerEngineClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {
    
    client, err := containerengine.NewContainerEngineClientWithConfigurationProvider(p)
    if testConfig.Endpoint != "" {
        client.Host = testConfig.Endpoint
    }else {
        client.SetRegion(testConfig.Region)
    }
   return client, err
}



// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientCreateCluster(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "CreateCluster")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateCluster is not enabled by the testing service")
    }
    c, err := containerengine.NewContainerEngineClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("containerengine", "CreateCluster")
    assert.NoError(t, err)

    type CreateClusterRequestInfo struct {
        ContainerId string
        Request containerengine.CreateClusterRequest
    }

    var requests []CreateClusterRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateCluster(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientCreateKubeconfig(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "CreateKubeconfig")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateKubeconfig is not enabled by the testing service")
    }
    c, err := containerengine.NewContainerEngineClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("containerengine", "CreateKubeconfig")
    assert.NoError(t, err)

    type CreateKubeconfigRequestInfo struct {
        ContainerId string
        Request containerengine.CreateKubeconfigRequest
    }

    var requests []CreateKubeconfigRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateKubeconfig(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientCreateNodePool(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "CreateNodePool")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateNodePool is not enabled by the testing service")
    }
    c, err := containerengine.NewContainerEngineClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("containerengine", "CreateNodePool")
    assert.NoError(t, err)

    type CreateNodePoolRequestInfo struct {
        ContainerId string
        Request containerengine.CreateNodePoolRequest
    }

    var requests []CreateNodePoolRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateNodePool(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientDeleteCluster(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "DeleteCluster")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteCluster is not enabled by the testing service")
    }
    c, err := containerengine.NewContainerEngineClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("containerengine", "DeleteCluster")
    assert.NoError(t, err)

    type DeleteClusterRequestInfo struct {
        ContainerId string
        Request containerengine.DeleteClusterRequest
    }

    var requests []DeleteClusterRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteCluster(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientDeleteNodePool(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "DeleteNodePool")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteNodePool is not enabled by the testing service")
    }
    c, err := containerengine.NewContainerEngineClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("containerengine", "DeleteNodePool")
    assert.NoError(t, err)

    type DeleteNodePoolRequestInfo struct {
        ContainerId string
        Request containerengine.DeleteNodePoolRequest
    }

    var requests []DeleteNodePoolRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteNodePool(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientDeleteWorkRequest(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "DeleteWorkRequest")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteWorkRequest is not enabled by the testing service")
    }
    c, err := containerengine.NewContainerEngineClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("containerengine", "DeleteWorkRequest")
    assert.NoError(t, err)

    type DeleteWorkRequestRequestInfo struct {
        ContainerId string
        Request containerengine.DeleteWorkRequestRequest
    }

    var requests []DeleteWorkRequestRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteWorkRequest(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientGetCluster(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "GetCluster")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetCluster is not enabled by the testing service")
    }
    c, err := containerengine.NewContainerEngineClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("containerengine", "GetCluster")
    assert.NoError(t, err)

    type GetClusterRequestInfo struct {
        ContainerId string
        Request containerengine.GetClusterRequest
    }

    var requests []GetClusterRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetCluster(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientGetClusterOptions(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "GetClusterOptions")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetClusterOptions is not enabled by the testing service")
    }
    c, err := containerengine.NewContainerEngineClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("containerengine", "GetClusterOptions")
    assert.NoError(t, err)

    type GetClusterOptionsRequestInfo struct {
        ContainerId string
        Request containerengine.GetClusterOptionsRequest
    }

    var requests []GetClusterOptionsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetClusterOptions(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientGetNodePool(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "GetNodePool")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetNodePool is not enabled by the testing service")
    }
    c, err := containerengine.NewContainerEngineClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("containerengine", "GetNodePool")
    assert.NoError(t, err)

    type GetNodePoolRequestInfo struct {
        ContainerId string
        Request containerengine.GetNodePoolRequest
    }

    var requests []GetNodePoolRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetNodePool(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientGetNodePoolOptions(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "GetNodePoolOptions")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetNodePoolOptions is not enabled by the testing service")
    }
    c, err := containerengine.NewContainerEngineClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("containerengine", "GetNodePoolOptions")
    assert.NoError(t, err)

    type GetNodePoolOptionsRequestInfo struct {
        ContainerId string
        Request containerengine.GetNodePoolOptionsRequest
    }

    var requests []GetNodePoolOptionsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetNodePoolOptions(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientGetWorkRequest(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "GetWorkRequest")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetWorkRequest is not enabled by the testing service")
    }
    c, err := containerengine.NewContainerEngineClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("containerengine", "GetWorkRequest")
    assert.NoError(t, err)

    type GetWorkRequestRequestInfo struct {
        ContainerId string
        Request containerengine.GetWorkRequestRequest
    }

    var requests []GetWorkRequestRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetWorkRequest(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientListClusters(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "ListClusters")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListClusters is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "ListClusters", createContainerEngineClientWithProvider)
    assert.NoError(t, err)
    c := cc.(containerengine.ContainerEngineClient)

    body, err := testClient.getRequests("containerengine", "ListClusters")
    assert.NoError(t, err)

    type ListClustersRequestInfo struct {
        ContainerId string
        Request containerengine.ListClustersRequest
    }

    var requests []ListClustersRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*containerengine.ListClustersRequest)
                return c.ListClusters(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]containerengine.ListClustersResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(containerengine.ListClustersResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientListNodePools(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "ListNodePools")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListNodePools is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "ListNodePools", createContainerEngineClientWithProvider)
    assert.NoError(t, err)
    c := cc.(containerengine.ContainerEngineClient)

    body, err := testClient.getRequests("containerengine", "ListNodePools")
    assert.NoError(t, err)

    type ListNodePoolsRequestInfo struct {
        ContainerId string
        Request containerengine.ListNodePoolsRequest
    }

    var requests []ListNodePoolsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*containerengine.ListNodePoolsRequest)
                return c.ListNodePools(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]containerengine.ListNodePoolsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(containerengine.ListNodePoolsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientListWorkRequestErrors(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "ListWorkRequestErrors")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListWorkRequestErrors is not enabled by the testing service")
    }
    c, err := containerengine.NewContainerEngineClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("containerengine", "ListWorkRequestErrors")
    assert.NoError(t, err)

    type ListWorkRequestErrorsRequestInfo struct {
        ContainerId string
        Request containerengine.ListWorkRequestErrorsRequest
    }

    var requests []ListWorkRequestErrorsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.ListWorkRequestErrors(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientListWorkRequestLogs(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "ListWorkRequestLogs")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListWorkRequestLogs is not enabled by the testing service")
    }
    c, err := containerengine.NewContainerEngineClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("containerengine", "ListWorkRequestLogs")
    assert.NoError(t, err)

    type ListWorkRequestLogsRequestInfo struct {
        ContainerId string
        Request containerengine.ListWorkRequestLogsRequest
    }

    var requests []ListWorkRequestLogsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.ListWorkRequestLogs(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientListWorkRequests(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "ListWorkRequests")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListWorkRequests is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("containerengine", "ContainerEngine", "ListWorkRequests", createContainerEngineClientWithProvider)
    assert.NoError(t, err)
    c := cc.(containerengine.ContainerEngineClient)

    body, err := testClient.getRequests("containerengine", "ListWorkRequests")
    assert.NoError(t, err)

    type ListWorkRequestsRequestInfo struct {
        ContainerId string
        Request containerengine.ListWorkRequestsRequest
    }

    var requests []ListWorkRequestsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*containerengine.ListWorkRequestsRequest)
                return c.ListWorkRequests(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]containerengine.ListWorkRequestsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(containerengine.ListWorkRequestsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientUpdateCluster(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "UpdateCluster")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("UpdateCluster is not enabled by the testing service")
    }
    c, err := containerengine.NewContainerEngineClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("containerengine", "UpdateCluster")
    assert.NoError(t, err)

    type UpdateClusterRequestInfo struct {
        ContainerId string
        Request containerengine.UpdateClusterRequest
    }

    var requests []UpdateClusterRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.UpdateCluster(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="" email="" jiraProject="" opsJiraProject=""
func TestContainerEngineClientUpdateNodePool(t *testing.T) {
    enabled, err := testClient.isApiEnabled("containerengine", "UpdateNodePool")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("UpdateNodePool is not enabled by the testing service")
    }
    c, err := containerengine.NewContainerEngineClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("containerengine", "UpdateNodePool")
    assert.NoError(t, err)

    type UpdateNodePoolRequestInfo struct {
        ContainerId string
        Request containerengine.UpdateNodePoolRequest
    }

    var requests []UpdateNodePoolRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.UpdateNodePool(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}
