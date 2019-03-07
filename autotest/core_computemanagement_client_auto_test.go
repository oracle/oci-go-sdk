package autotest

import (
    "github.com/oracle/oci-go-sdk/core"
    "github.com/oracle/oci-go-sdk/common"

    "context"
    "encoding/json"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func createComputeManagementClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {
    
    client, err := core.NewComputeManagementClientWithConfigurationProvider(p)
    if testConfig.Endpoint != "" {
        client.Host = testConfig.Endpoint
    }else {
        client.SetRegion(testConfig.Region)
    }
   return client, err
}



// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientAttachLoadBalancer(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "AttachLoadBalancer")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("AttachLoadBalancer is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "ComputeManagement", "AttachLoadBalancer", createComputeManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeManagementClient)

	body, err := testClient.getRequests("core", "AttachLoadBalancer")
	assert.NoError(t, err)

	type AttachLoadBalancerRequestInfo struct {
		ContainerId string
		Request     core.AttachLoadBalancerRequest
	}

	var requests []AttachLoadBalancerRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.AttachLoadBalancer(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientCreateInstanceConfiguration(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "CreateInstanceConfiguration")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateInstanceConfiguration is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "ComputeManagement", "CreateInstanceConfiguration", createComputeManagementClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeManagementClient)

    body, err := testClient.getRequests("core", "CreateInstanceConfiguration")
    assert.NoError(t, err)

    type CreateInstanceConfigurationRequestInfo struct {
        ContainerId string
        Request core.CreateInstanceConfigurationRequest
    }

    var requests []CreateInstanceConfigurationRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateInstanceConfiguration(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientCreateInstancePool(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "CreateInstancePool")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateInstancePool is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "ComputeManagement", "CreateInstancePool", createComputeManagementClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeManagementClient)

    body, err := testClient.getRequests("core", "CreateInstancePool")
    assert.NoError(t, err)

    type CreateInstancePoolRequestInfo struct {
        ContainerId string
        Request core.CreateInstancePoolRequest
    }

    var requests []CreateInstancePoolRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateInstancePool(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientDeleteInstanceConfiguration(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "DeleteInstanceConfiguration")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteInstanceConfiguration is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "ComputeManagement", "DeleteInstanceConfiguration", createComputeManagementClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeManagementClient)

    body, err := testClient.getRequests("core", "DeleteInstanceConfiguration")
    assert.NoError(t, err)

    type DeleteInstanceConfigurationRequestInfo struct {
        ContainerId string
        Request core.DeleteInstanceConfigurationRequest
    }

    var requests []DeleteInstanceConfigurationRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteInstanceConfiguration(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientDetachLoadBalancer(t *testing.T) {
	enabled, err := testClient.isApiEnabled("core", "DetachLoadBalancer")
	assert.NoError(t, err)
	if !enabled {
		t.Skip("DetachLoadBalancer is not enabled by the testing service")
	}

	cc, err := testClient.createClientForOperation("core", "ComputeManagement", "DetachLoadBalancer", createComputeManagementClientWithProvider)
	assert.NoError(t, err)
	c := cc.(core.ComputeManagementClient)

	body, err := testClient.getRequests("core", "DetachLoadBalancer")
	assert.NoError(t, err)

	type DetachLoadBalancerRequestInfo struct {
		ContainerId string
		Request     core.DetachLoadBalancerRequest
	}

	var requests []DetachLoadBalancerRequestInfo
	err = json.Unmarshal([]byte(body), &requests)
	assert.NoError(t, err)

	var retryPolicy *common.RetryPolicy
	for i, req := range requests {
		t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
			retryPolicy = retryPolicyForTests()
			req.Request.RequestMetadata.RetryPolicy = retryPolicy

			response, err := c.DetachLoadBalancer(context.Background(), req.Request)
			message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
			assert.NoError(t, err)
			assert.Empty(t, message, message)
		})
	}
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientGetInstanceConfiguration(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "GetInstanceConfiguration")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetInstanceConfiguration is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "ComputeManagement", "GetInstanceConfiguration", createComputeManagementClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeManagementClient)

    body, err := testClient.getRequests("core", "GetInstanceConfiguration")
    assert.NoError(t, err)

    type GetInstanceConfigurationRequestInfo struct {
        ContainerId string
        Request core.GetInstanceConfigurationRequest
    }

    var requests []GetInstanceConfigurationRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetInstanceConfiguration(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientGetInstancePool(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "GetInstancePool")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetInstancePool is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "ComputeManagement", "GetInstancePool", createComputeManagementClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeManagementClient)

    body, err := testClient.getRequests("core", "GetInstancePool")
    assert.NoError(t, err)

    type GetInstancePoolRequestInfo struct {
        ContainerId string
        Request core.GetInstancePoolRequest
    }

    var requests []GetInstancePoolRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetInstancePool(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientLaunchInstanceConfiguration(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "LaunchInstanceConfiguration")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("LaunchInstanceConfiguration is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "ComputeManagement", "LaunchInstanceConfiguration", createComputeManagementClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeManagementClient)

    body, err := testClient.getRequests("core", "LaunchInstanceConfiguration")
    assert.NoError(t, err)

    type LaunchInstanceConfigurationRequestInfo struct {
        ContainerId string
        Request core.LaunchInstanceConfigurationRequest
    }

    var requests []LaunchInstanceConfigurationRequestInfo
    var pr []map[string]interface{}
    err = json.Unmarshal([]byte(body), &pr)
    assert.NoError(t, err)
    requests = make([]LaunchInstanceConfigurationRequestInfo, len(pr))
    polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
    polymorphicRequestInfo["InstanceConfigurationInstanceDetails"] = 
        PolymorphicRequestUnmarshallingInfo{
        DiscriminatorName : "instanceType",
        DiscriminatorValuesAndTypes: map[string]interface{}{ 
            "compute": &core.ComputeInstanceDetails{},
        },
    }

    for i, ppr := range pr {
        conditionalStructCopy(ppr, &requests[i], polymorphicRequestInfo, testClient.Log)
    }

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.LaunchInstanceConfiguration(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientListInstanceConfigurations(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ListInstanceConfigurations")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListInstanceConfigurations is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "ComputeManagement", "ListInstanceConfigurations", createComputeManagementClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeManagementClient)

    body, err := testClient.getRequests("core", "ListInstanceConfigurations")
    assert.NoError(t, err)

    type ListInstanceConfigurationsRequestInfo struct {
        ContainerId string
        Request core.ListInstanceConfigurationsRequest
    }

    var requests []ListInstanceConfigurationsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*core.ListInstanceConfigurationsRequest)
                return c.ListInstanceConfigurations(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]core.ListInstanceConfigurationsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(core.ListInstanceConfigurationsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientListInstancePoolInstances(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ListInstancePoolInstances")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListInstancePoolInstances is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "ComputeManagement", "ListInstancePoolInstances", createComputeManagementClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeManagementClient)

    body, err := testClient.getRequests("core", "ListInstancePoolInstances")
    assert.NoError(t, err)

    type ListInstancePoolInstancesRequestInfo struct {
        ContainerId string
        Request core.ListInstancePoolInstancesRequest
    }

    var requests []ListInstancePoolInstancesRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*core.ListInstancePoolInstancesRequest)
                return c.ListInstancePoolInstances(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]core.ListInstancePoolInstancesResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(core.ListInstancePoolInstancesResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientListInstancePools(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ListInstancePools")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListInstancePools is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "ComputeManagement", "ListInstancePools", createComputeManagementClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeManagementClient)

    body, err := testClient.getRequests("core", "ListInstancePools")
    assert.NoError(t, err)

    type ListInstancePoolsRequestInfo struct {
        ContainerId string
        Request core.ListInstancePoolsRequest
    }

    var requests []ListInstancePoolsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*core.ListInstancePoolsRequest)
                return c.ListInstancePools(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]core.ListInstancePoolsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(core.ListInstancePoolsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientResetInstancePool(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ResetInstancePool")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ResetInstancePool is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "ComputeManagement", "ResetInstancePool", createComputeManagementClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeManagementClient)

    body, err := testClient.getRequests("core", "ResetInstancePool")
    assert.NoError(t, err)

    type ResetInstancePoolRequestInfo struct {
        ContainerId string
        Request core.ResetInstancePoolRequest
    }

    var requests []ResetInstancePoolRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.ResetInstancePool(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientSoftresetInstancePool(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "SoftresetInstancePool")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("SoftresetInstancePool is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "ComputeManagement", "SoftresetInstancePool", createComputeManagementClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeManagementClient)

    body, err := testClient.getRequests("core", "SoftresetInstancePool")
    assert.NoError(t, err)

    type SoftresetInstancePoolRequestInfo struct {
        ContainerId string
        Request core.SoftresetInstancePoolRequest
    }

    var requests []SoftresetInstancePoolRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.SoftresetInstancePool(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientStartInstancePool(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "StartInstancePool")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("StartInstancePool is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "ComputeManagement", "StartInstancePool", createComputeManagementClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeManagementClient)

    body, err := testClient.getRequests("core", "StartInstancePool")
    assert.NoError(t, err)

    type StartInstancePoolRequestInfo struct {
        ContainerId string
        Request core.StartInstancePoolRequest
    }

    var requests []StartInstancePoolRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.StartInstancePool(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientStopInstancePool(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "StopInstancePool")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("StopInstancePool is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "ComputeManagement", "StopInstancePool", createComputeManagementClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeManagementClient)

    body, err := testClient.getRequests("core", "StopInstancePool")
    assert.NoError(t, err)

    type StopInstancePoolRequestInfo struct {
        ContainerId string
        Request core.StopInstancePoolRequest
    }

    var requests []StopInstancePoolRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.StopInstancePool(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientTerminateInstancePool(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "TerminateInstancePool")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("TerminateInstancePool is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "ComputeManagement", "TerminateInstancePool", createComputeManagementClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeManagementClient)

    body, err := testClient.getRequests("core", "TerminateInstancePool")
    assert.NoError(t, err)

    type TerminateInstancePoolRequestInfo struct {
        ContainerId string
        Request core.TerminateInstancePoolRequest
    }

    var requests []TerminateInstancePoolRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.TerminateInstancePool(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientUpdateInstanceConfiguration(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "UpdateInstanceConfiguration")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("UpdateInstanceConfiguration is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "ComputeManagement", "UpdateInstanceConfiguration", createComputeManagementClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeManagementClient)

    body, err := testClient.getRequests("core", "UpdateInstanceConfiguration")
    assert.NoError(t, err)

    type UpdateInstanceConfigurationRequestInfo struct {
        ContainerId string
        Request core.UpdateInstanceConfigurationRequest
    }

    var requests []UpdateInstanceConfigurationRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.UpdateInstanceConfiguration(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeManagementClientUpdateInstancePool(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "UpdateInstancePool")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("UpdateInstancePool is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "ComputeManagement", "UpdateInstancePool", createComputeManagementClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeManagementClient)

    body, err := testClient.getRequests("core", "UpdateInstancePool")
    assert.NoError(t, err)

    type UpdateInstancePoolRequestInfo struct {
        ContainerId string
        Request core.UpdateInstancePoolRequest
    }

    var requests []UpdateInstancePoolRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.UpdateInstancePool(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}
