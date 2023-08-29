// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// RoverBundleClient a client for RoverBundle
type RoverBundleClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewRoverBundleClientWithConfigurationProvider Creates a new default RoverBundle client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewRoverBundleClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client RoverBundleClient, err error) {
	if enabled := common.CheckForEnabledServices("rover"); !enabled {
		return client, fmt.Errorf("the Alloy configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local alloy_config file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newRoverBundleClientFromBaseClient(baseClient, provider)
}

// NewRoverBundleClientWithOboToken Creates a new default RoverBundle client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewRoverBundleClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client RoverBundleClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newRoverBundleClientFromBaseClient(baseClient, configProvider)
}

func newRoverBundleClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client RoverBundleClient, err error) {
	// RoverBundle service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("RoverBundle"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = RoverBundleClient{BaseClient: baseClient}
	client.BasePath = "20201210"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *RoverBundleClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("rover", "https://rover.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *RoverBundleClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *RoverBundleClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListRoverClusterRoverBundleRequests List all the roverBundleRequests for a given roverClusterId.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/ListRoverClusterRoverBundleRequests.go.html to see an example of how to use ListRoverClusterRoverBundleRequests API.
// A default retry strategy applies to this operation ListRoverClusterRoverBundleRequests()
func (client RoverBundleClient) ListRoverClusterRoverBundleRequests(ctx context.Context, request ListRoverClusterRoverBundleRequestsRequest) (response ListRoverClusterRoverBundleRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRoverClusterRoverBundleRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRoverClusterRoverBundleRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRoverClusterRoverBundleRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRoverClusterRoverBundleRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRoverClusterRoverBundleRequestsResponse")
	}
	return
}

// listRoverClusterRoverBundleRequests implements the OCIOperation interface (enables retrying operations)
func (client RoverBundleClient) listRoverClusterRoverBundleRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverClusters/{roverClusterId}/roverBundleRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRoverClusterRoverBundleRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverCluster/ListRoverClusterRoverBundleRequests"
		err = common.PostProcessServiceError(err, "RoverBundle", "ListRoverClusterRoverBundleRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRoverNodeRoverBundleRequests List all the roverBundleRequests for a given roverNodeId.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/ListRoverNodeRoverBundleRequests.go.html to see an example of how to use ListRoverNodeRoverBundleRequests API.
// A default retry strategy applies to this operation ListRoverNodeRoverBundleRequests()
func (client RoverBundleClient) ListRoverNodeRoverBundleRequests(ctx context.Context, request ListRoverNodeRoverBundleRequestsRequest) (response ListRoverNodeRoverBundleRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRoverNodeRoverBundleRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRoverNodeRoverBundleRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRoverNodeRoverBundleRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRoverNodeRoverBundleRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRoverNodeRoverBundleRequestsResponse")
	}
	return
}

// listRoverNodeRoverBundleRequests implements the OCIOperation interface (enables retrying operations)
func (client RoverBundleClient) listRoverNodeRoverBundleRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverNodes/{roverNodeId}/roverBundleRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRoverNodeRoverBundleRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNode/ListRoverNodeRoverBundleRequests"
		err = common.PostProcessServiceError(err, "RoverBundle", "ListRoverNodeRoverBundleRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestBundleRoverCluster Request to get rover bundle to the bucket in customer's tenancy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/RequestBundleRoverCluster.go.html to see an example of how to use RequestBundleRoverCluster API.
// A default retry strategy applies to this operation RequestBundleRoverCluster()
func (client RoverBundleClient) RequestBundleRoverCluster(ctx context.Context, request RequestBundleRoverClusterRequest) (response RequestBundleRoverClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestBundleRoverCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestBundleRoverClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestBundleRoverClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestBundleRoverClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestBundleRoverClusterResponse")
	}
	return
}

// requestBundleRoverCluster implements the OCIOperation interface (enables retrying operations)
func (client RoverBundleClient) requestBundleRoverCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverClusters/{roverClusterId}/actions/requestRoverBundle", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestBundleRoverClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverCluster/RequestBundleRoverCluster"
		err = common.PostProcessServiceError(err, "RoverBundle", "RequestBundleRoverCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestBundleRoverNode Request to get rover bundle to the bucket in customer's tenancy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/RequestBundleRoverNode.go.html to see an example of how to use RequestBundleRoverNode API.
// A default retry strategy applies to this operation RequestBundleRoverNode()
func (client RoverBundleClient) RequestBundleRoverNode(ctx context.Context, request RequestBundleRoverNodeRequest) (response RequestBundleRoverNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestBundleRoverNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestBundleRoverNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestBundleRoverNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestBundleRoverNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestBundleRoverNodeResponse")
	}
	return
}

// requestBundleRoverNode implements the OCIOperation interface (enables retrying operations)
func (client RoverBundleClient) requestBundleRoverNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverNodes/{roverNodeId}/actions/requestRoverBundle", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestBundleRoverNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNode/RequestBundleRoverNode"
		err = common.PostProcessServiceError(err, "RoverBundle", "RequestBundleRoverNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetrieveAvailableBundleVersionsRoverCluster Retrieve the latest available rover bundle version that can be upgraded to based on current bundle version.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/RetrieveAvailableBundleVersionsRoverCluster.go.html to see an example of how to use RetrieveAvailableBundleVersionsRoverCluster API.
// A default retry strategy applies to this operation RetrieveAvailableBundleVersionsRoverCluster()
func (client RoverBundleClient) RetrieveAvailableBundleVersionsRoverCluster(ctx context.Context, request RetrieveAvailableBundleVersionsRoverClusterRequest) (response RetrieveAvailableBundleVersionsRoverClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.retrieveAvailableBundleVersionsRoverCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetrieveAvailableBundleVersionsRoverClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetrieveAvailableBundleVersionsRoverClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetrieveAvailableBundleVersionsRoverClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetrieveAvailableBundleVersionsRoverClusterResponse")
	}
	return
}

// retrieveAvailableBundleVersionsRoverCluster implements the OCIOperation interface (enables retrying operations)
func (client RoverBundleClient) retrieveAvailableBundleVersionsRoverCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverClusters/{roverClusterId}/actions/retrieveAvailableRoverBundleVersion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetrieveAvailableBundleVersionsRoverClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverCluster/RetrieveAvailableBundleVersionsRoverCluster"
		err = common.PostProcessServiceError(err, "RoverBundle", "RetrieveAvailableBundleVersionsRoverCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetrieveAvailableBundleVersionsRoverNode Retrieve the latest available rover bundle version that can be upgraded to based on current bundle version.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/RetrieveAvailableBundleVersionsRoverNode.go.html to see an example of how to use RetrieveAvailableBundleVersionsRoverNode API.
// A default retry strategy applies to this operation RetrieveAvailableBundleVersionsRoverNode()
func (client RoverBundleClient) RetrieveAvailableBundleVersionsRoverNode(ctx context.Context, request RetrieveAvailableBundleVersionsRoverNodeRequest) (response RetrieveAvailableBundleVersionsRoverNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.retrieveAvailableBundleVersionsRoverNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetrieveAvailableBundleVersionsRoverNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetrieveAvailableBundleVersionsRoverNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetrieveAvailableBundleVersionsRoverNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetrieveAvailableBundleVersionsRoverNodeResponse")
	}
	return
}

// retrieveAvailableBundleVersionsRoverNode implements the OCIOperation interface (enables retrying operations)
func (client RoverBundleClient) retrieveAvailableBundleVersionsRoverNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverNodes/{roverNodeId}/actions/retrieveAvailableRoverBundleVersion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetrieveAvailableBundleVersionsRoverNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNode/RetrieveAvailableBundleVersionsRoverNode"
		err = common.PostProcessServiceError(err, "RoverBundle", "RetrieveAvailableBundleVersionsRoverNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetrieveBundleStatusRoverCluster Retrieve the status and progress of a rover bundle copy request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/RetrieveBundleStatusRoverCluster.go.html to see an example of how to use RetrieveBundleStatusRoverCluster API.
// A default retry strategy applies to this operation RetrieveBundleStatusRoverCluster()
func (client RoverBundleClient) RetrieveBundleStatusRoverCluster(ctx context.Context, request RetrieveBundleStatusRoverClusterRequest) (response RetrieveBundleStatusRoverClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.retrieveBundleStatusRoverCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetrieveBundleStatusRoverClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetrieveBundleStatusRoverClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetrieveBundleStatusRoverClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetrieveBundleStatusRoverClusterResponse")
	}
	return
}

// retrieveBundleStatusRoverCluster implements the OCIOperation interface (enables retrying operations)
func (client RoverBundleClient) retrieveBundleStatusRoverCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverClusters/{roverClusterId}/actions/retrieveRoverBundleStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetrieveBundleStatusRoverClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverCluster/RetrieveBundleStatusRoverCluster"
		err = common.PostProcessServiceError(err, "RoverBundle", "RetrieveBundleStatusRoverCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetrieveBundleStatusRoverNode Retrieve the status and progress of a rover bundle copy request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/RetrieveBundleStatusRoverNode.go.html to see an example of how to use RetrieveBundleStatusRoverNode API.
// A default retry strategy applies to this operation RetrieveBundleStatusRoverNode()
func (client RoverBundleClient) RetrieveBundleStatusRoverNode(ctx context.Context, request RetrieveBundleStatusRoverNodeRequest) (response RetrieveBundleStatusRoverNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.retrieveBundleStatusRoverNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetrieveBundleStatusRoverNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetrieveBundleStatusRoverNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetrieveBundleStatusRoverNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetrieveBundleStatusRoverNodeResponse")
	}
	return
}

// retrieveBundleStatusRoverNode implements the OCIOperation interface (enables retrying operations)
func (client RoverBundleClient) retrieveBundleStatusRoverNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverNodes/{roverNodeId}/actions/retrieveRoverBundleStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetrieveBundleStatusRoverNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNode/RetrieveBundleStatusRoverNode"
		err = common.PostProcessServiceError(err, "RoverBundle", "RetrieveBundleStatusRoverNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
