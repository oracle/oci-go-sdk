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

func createComputeClientWithProvider(p common.ConfigurationProvider, testConfig TestingConfig) (interface{}, error) {
    
    client, err := core.NewComputeClientWithConfigurationProvider(p)
    if testConfig.Endpoint != "" {
        client.Host = testConfig.Endpoint
    }else {
        client.SetRegion(testConfig.Region)
    }
   return client, err
}



// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientAttachBootVolume(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "AttachBootVolume")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("AttachBootVolume is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "AttachBootVolume")
    assert.NoError(t, err)

    type AttachBootVolumeRequestInfo struct {
        ContainerId string
        Request core.AttachBootVolumeRequest
    }

    var requests []AttachBootVolumeRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.AttachBootVolume(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientAttachVnic(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "AttachVnic")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("AttachVnic is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "AttachVnic")
    assert.NoError(t, err)

    type AttachVnicRequestInfo struct {
        ContainerId string
        Request core.AttachVnicRequest
    }

    var requests []AttachVnicRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.AttachVnic(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientAttachVolume(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "AttachVolume")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("AttachVolume is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "AttachVolume")
    assert.NoError(t, err)

    type AttachVolumeRequestInfo struct {
        ContainerId string
        Request core.AttachVolumeRequest
    }

    var requests []AttachVolumeRequestInfo
    var pr []map[string]interface{}
    err = json.Unmarshal([]byte(body), &pr)
    assert.NoError(t, err)
    requests = make([]AttachVolumeRequestInfo, len(pr))
    polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
    polymorphicRequestInfo["AttachVolumeDetails"] = 
        PolymorphicRequestUnmarshallingInfo{
        DiscriminatorName : "type",
        DiscriminatorValuesAndTypes: map[string]interface{}{ 
            "iscsi": &core.AttachIScsiVolumeDetails{},
            "paravirtualized": &core.AttachParavirtualizedVolumeDetails{},
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

            response, err := c.AttachVolume(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientCaptureConsoleHistory(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "CaptureConsoleHistory")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CaptureConsoleHistory is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "CaptureConsoleHistory")
    assert.NoError(t, err)

    type CaptureConsoleHistoryRequestInfo struct {
        ContainerId string
        Request core.CaptureConsoleHistoryRequest
    }

    var requests []CaptureConsoleHistoryRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CaptureConsoleHistory(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientCreateAppCatalogSubscription(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "CreateAppCatalogSubscription")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateAppCatalogSubscription is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "CreateAppCatalogSubscription")
    assert.NoError(t, err)

    type CreateAppCatalogSubscriptionRequestInfo struct {
        ContainerId string
        Request core.CreateAppCatalogSubscriptionRequest
    }

    var requests []CreateAppCatalogSubscriptionRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateAppCatalogSubscription(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientCreateImage(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "CreateImage")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateImage is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "CreateImage")
    assert.NoError(t, err)

    type CreateImageRequestInfo struct {
        ContainerId string
        Request core.CreateImageRequest
    }

    var requests []CreateImageRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateImage(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientCreateInstanceConsoleConnection(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "CreateInstanceConsoleConnection")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("CreateInstanceConsoleConnection is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "CreateInstanceConsoleConnection")
    assert.NoError(t, err)

    type CreateInstanceConsoleConnectionRequestInfo struct {
        ContainerId string
        Request core.CreateInstanceConsoleConnectionRequest
    }

    var requests []CreateInstanceConsoleConnectionRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.CreateInstanceConsoleConnection(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientDeleteAppCatalogSubscription(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "DeleteAppCatalogSubscription")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteAppCatalogSubscription is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "DeleteAppCatalogSubscription")
    assert.NoError(t, err)

    type DeleteAppCatalogSubscriptionRequestInfo struct {
        ContainerId string
        Request core.DeleteAppCatalogSubscriptionRequest
    }

    var requests []DeleteAppCatalogSubscriptionRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteAppCatalogSubscription(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientDeleteConsoleHistory(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "DeleteConsoleHistory")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteConsoleHistory is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "DeleteConsoleHistory")
    assert.NoError(t, err)

    type DeleteConsoleHistoryRequestInfo struct {
        ContainerId string
        Request core.DeleteConsoleHistoryRequest
    }

    var requests []DeleteConsoleHistoryRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteConsoleHistory(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientDeleteImage(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "DeleteImage")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteImage is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "DeleteImage")
    assert.NoError(t, err)

    type DeleteImageRequestInfo struct {
        ContainerId string
        Request core.DeleteImageRequest
    }

    var requests []DeleteImageRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteImage(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientDeleteInstanceConsoleConnection(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "DeleteInstanceConsoleConnection")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DeleteInstanceConsoleConnection is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "DeleteInstanceConsoleConnection")
    assert.NoError(t, err)

    type DeleteInstanceConsoleConnectionRequestInfo struct {
        ContainerId string
        Request core.DeleteInstanceConsoleConnectionRequest
    }

    var requests []DeleteInstanceConsoleConnectionRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DeleteInstanceConsoleConnection(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientDetachBootVolume(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "DetachBootVolume")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DetachBootVolume is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "DetachBootVolume")
    assert.NoError(t, err)

    type DetachBootVolumeRequestInfo struct {
        ContainerId string
        Request core.DetachBootVolumeRequest
    }

    var requests []DetachBootVolumeRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DetachBootVolume(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientDetachVnic(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "DetachVnic")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DetachVnic is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "DetachVnic")
    assert.NoError(t, err)

    type DetachVnicRequestInfo struct {
        ContainerId string
        Request core.DetachVnicRequest
    }

    var requests []DetachVnicRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DetachVnic(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientDetachVolume(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "DetachVolume")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("DetachVolume is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "DetachVolume")
    assert.NoError(t, err)

    type DetachVolumeRequestInfo struct {
        ContainerId string
        Request core.DetachVolumeRequest
    }

    var requests []DetachVolumeRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.DetachVolume(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientExportImage(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ExportImage")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ExportImage is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "ExportImage")
    assert.NoError(t, err)

    type ExportImageRequestInfo struct {
        ContainerId string
        Request core.ExportImageRequest
    }

    var requests []ExportImageRequestInfo
    var pr []map[string]interface{}
    err = json.Unmarshal([]byte(body), &pr)
    assert.NoError(t, err)
    requests = make([]ExportImageRequestInfo, len(pr))
    polymorphicRequestInfo := map[string]PolymorphicRequestUnmarshallingInfo{}
    polymorphicRequestInfo["ExportImageDetails"] = 
        PolymorphicRequestUnmarshallingInfo{
        DiscriminatorName : "destinationType",
        DiscriminatorValuesAndTypes: map[string]interface{}{ 
            "objectStorageUri": &core.ExportImageViaObjectStorageUriDetails{},
            "objectStorageTuple": &core.ExportImageViaObjectStorageTupleDetails{},
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

            response, err := c.ExportImage(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientGetAppCatalogListing(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "GetAppCatalogListing")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetAppCatalogListing is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "GetAppCatalogListing")
    assert.NoError(t, err)

    type GetAppCatalogListingRequestInfo struct {
        ContainerId string
        Request core.GetAppCatalogListingRequest
    }

    var requests []GetAppCatalogListingRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetAppCatalogListing(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientGetAppCatalogListingAgreements(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "GetAppCatalogListingAgreements")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetAppCatalogListingAgreements is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "GetAppCatalogListingAgreements")
    assert.NoError(t, err)

    type GetAppCatalogListingAgreementsRequestInfo struct {
        ContainerId string
        Request core.GetAppCatalogListingAgreementsRequest
    }

    var requests []GetAppCatalogListingAgreementsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetAppCatalogListingAgreements(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientGetAppCatalogListingResourceVersion(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "GetAppCatalogListingResourceVersion")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetAppCatalogListingResourceVersion is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "GetAppCatalogListingResourceVersion")
    assert.NoError(t, err)

    type GetAppCatalogListingResourceVersionRequestInfo struct {
        ContainerId string
        Request core.GetAppCatalogListingResourceVersionRequest
    }

    var requests []GetAppCatalogListingResourceVersionRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetAppCatalogListingResourceVersion(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientGetBootVolumeAttachment(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "GetBootVolumeAttachment")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetBootVolumeAttachment is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "GetBootVolumeAttachment")
    assert.NoError(t, err)

    type GetBootVolumeAttachmentRequestInfo struct {
        ContainerId string
        Request core.GetBootVolumeAttachmentRequest
    }

    var requests []GetBootVolumeAttachmentRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetBootVolumeAttachment(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientGetConsoleHistory(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "GetConsoleHistory")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetConsoleHistory is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "GetConsoleHistory")
    assert.NoError(t, err)

    type GetConsoleHistoryRequestInfo struct {
        ContainerId string
        Request core.GetConsoleHistoryRequest
    }

    var requests []GetConsoleHistoryRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetConsoleHistory(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientGetConsoleHistoryContent(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "GetConsoleHistoryContent")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetConsoleHistoryContent is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "GetConsoleHistoryContent")
    assert.NoError(t, err)

    type GetConsoleHistoryContentRequestInfo struct {
        ContainerId string
        Request core.GetConsoleHistoryContentRequest
    }

    var requests []GetConsoleHistoryContentRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetConsoleHistoryContent(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientGetImage(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "GetImage")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetImage is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "GetImage")
    assert.NoError(t, err)

    type GetImageRequestInfo struct {
        ContainerId string
        Request core.GetImageRequest
    }

    var requests []GetImageRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetImage(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientGetInstance(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "GetInstance")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetInstance is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "GetInstance")
    assert.NoError(t, err)

    type GetInstanceRequestInfo struct {
        ContainerId string
        Request core.GetInstanceRequest
    }

    var requests []GetInstanceRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetInstance(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientGetInstanceConsoleConnection(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "GetInstanceConsoleConnection")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetInstanceConsoleConnection is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "GetInstanceConsoleConnection")
    assert.NoError(t, err)

    type GetInstanceConsoleConnectionRequestInfo struct {
        ContainerId string
        Request core.GetInstanceConsoleConnectionRequest
    }

    var requests []GetInstanceConsoleConnectionRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetInstanceConsoleConnection(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientGetVnicAttachment(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "GetVnicAttachment")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetVnicAttachment is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "GetVnicAttachment")
    assert.NoError(t, err)

    type GetVnicAttachmentRequestInfo struct {
        ContainerId string
        Request core.GetVnicAttachmentRequest
    }

    var requests []GetVnicAttachmentRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetVnicAttachment(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientGetVolumeAttachment(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "GetVolumeAttachment")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetVolumeAttachment is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "GetVolumeAttachment")
    assert.NoError(t, err)

    type GetVolumeAttachmentRequestInfo struct {
        ContainerId string
        Request core.GetVolumeAttachmentRequest
    }

    var requests []GetVolumeAttachmentRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetVolumeAttachment(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientGetWindowsInstanceInitialCredentials(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "GetWindowsInstanceInitialCredentials")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("GetWindowsInstanceInitialCredentials is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "GetWindowsInstanceInitialCredentials")
    assert.NoError(t, err)

    type GetWindowsInstanceInitialCredentialsRequestInfo struct {
        ContainerId string
        Request core.GetWindowsInstanceInitialCredentialsRequest
    }

    var requests []GetWindowsInstanceInitialCredentialsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.GetWindowsInstanceInitialCredentials(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientInstanceAction(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "InstanceAction")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("InstanceAction is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "InstanceAction")
    assert.NoError(t, err)

    type InstanceActionRequestInfo struct {
        ContainerId string
        Request core.InstanceActionRequest
    }

    var requests []InstanceActionRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.InstanceAction(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientLaunchInstance(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "LaunchInstance")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("LaunchInstance is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "LaunchInstance")
    assert.NoError(t, err)

    type LaunchInstanceRequestInfo struct {
        ContainerId string
        Request core.LaunchInstanceRequest
    }

    var requests []LaunchInstanceRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.LaunchInstance(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientListAppCatalogListingResourceVersions(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ListAppCatalogListingResourceVersions")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListAppCatalogListingResourceVersions is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "Compute", "ListAppCatalogListingResourceVersions", createComputeClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeClient)

    body, err := testClient.getRequests("core", "ListAppCatalogListingResourceVersions")
    assert.NoError(t, err)

    type ListAppCatalogListingResourceVersionsRequestInfo struct {
        ContainerId string
        Request core.ListAppCatalogListingResourceVersionsRequest
    }

    var requests []ListAppCatalogListingResourceVersionsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*core.ListAppCatalogListingResourceVersionsRequest)
                return c.ListAppCatalogListingResourceVersions(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]core.ListAppCatalogListingResourceVersionsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(core.ListAppCatalogListingResourceVersionsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientListAppCatalogListings(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ListAppCatalogListings")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListAppCatalogListings is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "Compute", "ListAppCatalogListings", createComputeClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeClient)

    body, err := testClient.getRequests("core", "ListAppCatalogListings")
    assert.NoError(t, err)

    type ListAppCatalogListingsRequestInfo struct {
        ContainerId string
        Request core.ListAppCatalogListingsRequest
    }

    var requests []ListAppCatalogListingsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*core.ListAppCatalogListingsRequest)
                return c.ListAppCatalogListings(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]core.ListAppCatalogListingsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(core.ListAppCatalogListingsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientListAppCatalogSubscriptions(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ListAppCatalogSubscriptions")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListAppCatalogSubscriptions is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "Compute", "ListAppCatalogSubscriptions", createComputeClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeClient)

    body, err := testClient.getRequests("core", "ListAppCatalogSubscriptions")
    assert.NoError(t, err)

    type ListAppCatalogSubscriptionsRequestInfo struct {
        ContainerId string
        Request core.ListAppCatalogSubscriptionsRequest
    }

    var requests []ListAppCatalogSubscriptionsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*core.ListAppCatalogSubscriptionsRequest)
                return c.ListAppCatalogSubscriptions(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]core.ListAppCatalogSubscriptionsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(core.ListAppCatalogSubscriptionsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientListBootVolumeAttachments(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ListBootVolumeAttachments")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListBootVolumeAttachments is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "Compute", "ListBootVolumeAttachments", createComputeClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeClient)

    body, err := testClient.getRequests("core", "ListBootVolumeAttachments")
    assert.NoError(t, err)

    type ListBootVolumeAttachmentsRequestInfo struct {
        ContainerId string
        Request core.ListBootVolumeAttachmentsRequest
    }

    var requests []ListBootVolumeAttachmentsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*core.ListBootVolumeAttachmentsRequest)
                return c.ListBootVolumeAttachments(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]core.ListBootVolumeAttachmentsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(core.ListBootVolumeAttachmentsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientListConsoleHistories(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ListConsoleHistories")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListConsoleHistories is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "Compute", "ListConsoleHistories", createComputeClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeClient)

    body, err := testClient.getRequests("core", "ListConsoleHistories")
    assert.NoError(t, err)

    type ListConsoleHistoriesRequestInfo struct {
        ContainerId string
        Request core.ListConsoleHistoriesRequest
    }

    var requests []ListConsoleHistoriesRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*core.ListConsoleHistoriesRequest)
                return c.ListConsoleHistories(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]core.ListConsoleHistoriesResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(core.ListConsoleHistoriesResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientListImages(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ListImages")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListImages is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "Compute", "ListImages", createComputeClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeClient)

    body, err := testClient.getRequests("core", "ListImages")
    assert.NoError(t, err)

    type ListImagesRequestInfo struct {
        ContainerId string
        Request core.ListImagesRequest
    }

    var requests []ListImagesRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*core.ListImagesRequest)
                return c.ListImages(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]core.ListImagesResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(core.ListImagesResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientListInstanceConsoleConnections(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ListInstanceConsoleConnections")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListInstanceConsoleConnections is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "Compute", "ListInstanceConsoleConnections", createComputeClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeClient)

    body, err := testClient.getRequests("core", "ListInstanceConsoleConnections")
    assert.NoError(t, err)

    type ListInstanceConsoleConnectionsRequestInfo struct {
        ContainerId string
        Request core.ListInstanceConsoleConnectionsRequest
    }

    var requests []ListInstanceConsoleConnectionsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*core.ListInstanceConsoleConnectionsRequest)
                return c.ListInstanceConsoleConnections(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]core.ListInstanceConsoleConnectionsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(core.ListInstanceConsoleConnectionsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientListInstanceDevices(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ListInstanceDevices")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListInstanceDevices is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "Compute", "ListInstanceDevices", createComputeClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeClient)

    body, err := testClient.getRequests("core", "ListInstanceDevices")
    assert.NoError(t, err)

    type ListInstanceDevicesRequestInfo struct {
        ContainerId string
        Request core.ListInstanceDevicesRequest
    }

    var requests []ListInstanceDevicesRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*core.ListInstanceDevicesRequest)
                return c.ListInstanceDevices(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]core.ListInstanceDevicesResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(core.ListInstanceDevicesResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientListInstances(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ListInstances")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListInstances is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "Compute", "ListInstances", createComputeClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeClient)

    body, err := testClient.getRequests("core", "ListInstances")
    assert.NoError(t, err)

    type ListInstancesRequestInfo struct {
        ContainerId string
        Request core.ListInstancesRequest
    }

    var requests []ListInstancesRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*core.ListInstancesRequest)
                return c.ListInstances(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]core.ListInstancesResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(core.ListInstancesResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientListShapes(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ListShapes")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListShapes is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "Compute", "ListShapes", createComputeClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeClient)

    body, err := testClient.getRequests("core", "ListShapes")
    assert.NoError(t, err)

    type ListShapesRequestInfo struct {
        ContainerId string
        Request core.ListShapesRequest
    }

    var requests []ListShapesRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*core.ListShapesRequest)
                return c.ListShapes(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]core.ListShapesResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(core.ListShapesResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientListVnicAttachments(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ListVnicAttachments")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListVnicAttachments is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "Compute", "ListVnicAttachments", createComputeClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeClient)

    body, err := testClient.getRequests("core", "ListVnicAttachments")
    assert.NoError(t, err)

    type ListVnicAttachmentsRequestInfo struct {
        ContainerId string
        Request core.ListVnicAttachmentsRequest
    }

    var requests []ListVnicAttachmentsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*core.ListVnicAttachmentsRequest)
                return c.ListVnicAttachments(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]core.ListVnicAttachmentsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(core.ListVnicAttachmentsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientListVolumeAttachments(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "ListVolumeAttachments")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("ListVolumeAttachments is not enabled by the testing service")
    }

    cc, err := testClient.createClientForOperation("core", "Compute", "ListVolumeAttachments", createComputeClientWithProvider)
    assert.NoError(t, err)
    c := cc.(core.ComputeClient)

    body, err := testClient.getRequests("core", "ListVolumeAttachments")
    assert.NoError(t, err)

    type ListVolumeAttachmentsRequestInfo struct {
        ContainerId string
        Request core.ListVolumeAttachmentsRequest
    }

    var requests []ListVolumeAttachmentsRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)


    var retryPolicy *common.RetryPolicy
    for i, request := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            request.Request.RequestMetadata.RetryPolicy =  retryPolicy
            listFn := func(req common.OCIRequest) (common.OCIResponse, error) {
                r := req.(*core.ListVolumeAttachmentsRequest)
                return c.ListVolumeAttachments(context.Background(), *r)
            }

            listResponses, err := testClient.generateListResponses(&request.Request, listFn)
            typedListResponses := make([]core.ListVolumeAttachmentsResponse, len(listResponses))
            for i, lr := range listResponses {
                typedListResponses[i] = lr.(core.ListVolumeAttachmentsResponse)
            }

            message, err := testClient.validateResult(request.ContainerId, request.Request, typedListResponses, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientTerminateInstance(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "TerminateInstance")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("TerminateInstance is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "TerminateInstance")
    assert.NoError(t, err)

    type TerminateInstanceRequestInfo struct {
        ContainerId string
        Request core.TerminateInstanceRequest
    }

    var requests []TerminateInstanceRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.TerminateInstance(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientUpdateConsoleHistory(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "UpdateConsoleHistory")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("UpdateConsoleHistory is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "UpdateConsoleHistory")
    assert.NoError(t, err)

    type UpdateConsoleHistoryRequestInfo struct {
        ContainerId string
        Request core.UpdateConsoleHistoryRequest
    }

    var requests []UpdateConsoleHistoryRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.UpdateConsoleHistory(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientUpdateImage(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "UpdateImage")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("UpdateImage is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "UpdateImage")
    assert.NoError(t, err)

    type UpdateImageRequestInfo struct {
        ContainerId string
        Request core.UpdateImageRequest
    }

    var requests []UpdateImageRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.UpdateImage(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}

// IssueRoutingInfo tag="default" email="sic_block_storage_us_grp@oracle.com" jiraProject="BLOCK" opsJiraProject="BS"
func TestComputeClientUpdateInstance(t *testing.T) {
    enabled, err := testClient.isApiEnabled("core", "UpdateInstance")
    assert.NoError(t, err)
    if !enabled {
        t.Skip("UpdateInstance is not enabled by the testing service")
    }
    c, err := core.NewComputeClientWithConfigurationProvider(testConfig.ConfigurationProvider)
    assert.NoError(t, err)

    body, err := testClient.getRequests("core", "UpdateInstance")
    assert.NoError(t, err)

    type UpdateInstanceRequestInfo struct {
        ContainerId string
        Request core.UpdateInstanceRequest
    }

    var requests []UpdateInstanceRequestInfo
    err = json.Unmarshal([]byte(body), &requests)
    assert.NoError(t, err)

    var retryPolicy  *common.RetryPolicy
    for i, req := range requests {
        t.Run(fmt.Sprintf("request:%v", i), func(t *testing.T) {
            retryPolicy = retryPolicyForTests()
            req.Request.RequestMetadata.RetryPolicy =  retryPolicy

            response, err := c.UpdateInstance(context.Background(), req.Request)
            message, err := testClient.validateResult(req.ContainerId, req.Request, response, err)
            assert.NoError(t, err)
            assert.Empty(t, message, message)
        })
    }
}
