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

// RoverClusterClient a client for RoverCluster
type RoverClusterClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewRoverClusterClientWithConfigurationProvider Creates a new default RoverCluster client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewRoverClusterClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client RoverClusterClient, err error) {
	if enabled := common.CheckForEnabledServices("rover"); !enabled {
		return client, fmt.Errorf("the Developer Tool configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local developer-tool-configuration.json file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newRoverClusterClientFromBaseClient(baseClient, provider)
}

// NewRoverClusterClientWithOboToken Creates a new default RoverCluster client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewRoverClusterClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client RoverClusterClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newRoverClusterClientFromBaseClient(baseClient, configProvider)
}

func newRoverClusterClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client RoverClusterClient, err error) {
	// RoverCluster service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("RoverCluster"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = RoverClusterClient{BaseClient: baseClient}
	client.BasePath = "20201210"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *RoverClusterClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("rover", "https://rover.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *RoverClusterClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *RoverClusterClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeRoverClusterCompartment Moves a cluster into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/ChangeRoverClusterCompartment.go.html to see an example of how to use ChangeRoverClusterCompartment API.
// A default retry strategy applies to this operation ChangeRoverClusterCompartment()
func (client RoverClusterClient) ChangeRoverClusterCompartment(ctx context.Context, request ChangeRoverClusterCompartmentRequest) (response ChangeRoverClusterCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeRoverClusterCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeRoverClusterCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeRoverClusterCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeRoverClusterCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeRoverClusterCompartmentResponse")
	}
	return
}

// changeRoverClusterCompartment implements the OCIOperation interface (enables retrying operations)
func (client RoverClusterClient) changeRoverClusterCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverClusters/{roverClusterId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeRoverClusterCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverCluster/ChangeRoverClusterCompartment"
		err = common.PostProcessServiceError(err, "RoverCluster", "ChangeRoverClusterCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateRoverCluster Creates a new RoverCluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/CreateRoverCluster.go.html to see an example of how to use CreateRoverCluster API.
// A default retry strategy applies to this operation CreateRoverCluster()
func (client RoverClusterClient) CreateRoverCluster(ctx context.Context, request CreateRoverClusterRequest) (response CreateRoverClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createRoverCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateRoverClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateRoverClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRoverClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRoverClusterResponse")
	}
	return
}

// createRoverCluster implements the OCIOperation interface (enables retrying operations)
func (client RoverClusterClient) createRoverCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateRoverClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverCluster/CreateRoverCluster"
		err = common.PostProcessServiceError(err, "RoverCluster", "CreateRoverCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteRoverCluster Deletes a RoverCluster resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/DeleteRoverCluster.go.html to see an example of how to use DeleteRoverCluster API.
// A default retry strategy applies to this operation DeleteRoverCluster()
func (client RoverClusterClient) DeleteRoverCluster(ctx context.Context, request DeleteRoverClusterRequest) (response DeleteRoverClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteRoverCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRoverClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRoverClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRoverClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRoverClusterResponse")
	}
	return
}

// deleteRoverCluster implements the OCIOperation interface (enables retrying operations)
func (client RoverClusterClient) deleteRoverCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/roverClusters/{roverClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteRoverClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverCluster/DeleteRoverCluster"
		err = common.PostProcessServiceError(err, "RoverCluster", "DeleteRoverCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRoverCluster Gets a RoverCluster by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/GetRoverCluster.go.html to see an example of how to use GetRoverCluster API.
// A default retry strategy applies to this operation GetRoverCluster()
func (client RoverClusterClient) GetRoverCluster(ctx context.Context, request GetRoverClusterRequest) (response GetRoverClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRoverCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRoverClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRoverClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRoverClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRoverClusterResponse")
	}
	return
}

// getRoverCluster implements the OCIOperation interface (enables retrying operations)
func (client RoverClusterClient) getRoverCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverClusters/{roverClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRoverClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverCluster/GetRoverCluster"
		err = common.PostProcessServiceError(err, "RoverCluster", "GetRoverCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRoverClusterCertificate Get the certificate for a rover cluster
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/GetRoverClusterCertificate.go.html to see an example of how to use GetRoverClusterCertificate API.
// A default retry strategy applies to this operation GetRoverClusterCertificate()
func (client RoverClusterClient) GetRoverClusterCertificate(ctx context.Context, request GetRoverClusterCertificateRequest) (response GetRoverClusterCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRoverClusterCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRoverClusterCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRoverClusterCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRoverClusterCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRoverClusterCertificateResponse")
	}
	return
}

// getRoverClusterCertificate implements the OCIOperation interface (enables retrying operations)
func (client RoverClusterClient) getRoverClusterCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverClusters/{roverClusterId}/certificate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRoverClusterCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverClusterCertificate/GetRoverClusterCertificate"
		err = common.PostProcessServiceError(err, "RoverCluster", "GetRoverClusterCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRoverClusters Returns a list of RoverClusters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/ListRoverClusters.go.html to see an example of how to use ListRoverClusters API.
// A default retry strategy applies to this operation ListRoverClusters()
func (client RoverClusterClient) ListRoverClusters(ctx context.Context, request ListRoverClustersRequest) (response ListRoverClustersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRoverClusters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRoverClustersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRoverClustersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRoverClustersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRoverClustersResponse")
	}
	return
}

// listRoverClusters implements the OCIOperation interface (enables retrying operations)
func (client RoverClusterClient) listRoverClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRoverClustersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverCluster/ListRoverClusters"
		err = common.PostProcessServiceError(err, "RoverCluster", "ListRoverClusters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestAdditionalNodes Submit additional nodes request for a roverCluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/RequestAdditionalNodes.go.html to see an example of how to use RequestAdditionalNodes API.
// A default retry strategy applies to this operation RequestAdditionalNodes()
func (client RoverClusterClient) RequestAdditionalNodes(ctx context.Context, request RequestAdditionalNodesRequest) (response RequestAdditionalNodesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.requestAdditionalNodes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestAdditionalNodesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestAdditionalNodesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestAdditionalNodesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestAdditionalNodesResponse")
	}
	return
}

// requestAdditionalNodes implements the OCIOperation interface (enables retrying operations)
func (client RoverClusterClient) requestAdditionalNodes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverClusters/{roverClusterId}/actions/requestAdditionalNodes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestAdditionalNodesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverCluster/RequestAdditionalNodes"
		err = common.PostProcessServiceError(err, "RoverCluster", "RequestAdditionalNodes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRoverCluster Updates the RoverCluster
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/UpdateRoverCluster.go.html to see an example of how to use UpdateRoverCluster API.
// A default retry strategy applies to this operation UpdateRoverCluster()
func (client RoverClusterClient) UpdateRoverCluster(ctx context.Context, request UpdateRoverClusterRequest) (response UpdateRoverClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRoverCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRoverClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRoverClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRoverClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRoverClusterResponse")
	}
	return
}

// updateRoverCluster implements the OCIOperation interface (enables retrying operations)
func (client RoverClusterClient) updateRoverCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/roverClusters/{roverClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateRoverClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverCluster/UpdateRoverCluster"
		err = common.PostProcessServiceError(err, "RoverCluster", "UpdateRoverCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
