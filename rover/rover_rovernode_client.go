// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// RoverNodeClient a client for RoverNode
type RoverNodeClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewRoverNodeClientWithConfigurationProvider Creates a new default RoverNode client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewRoverNodeClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client RoverNodeClient, err error) {
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
	return newRoverNodeClientFromBaseClient(baseClient, provider)
}

// NewRoverNodeClientWithOboToken Creates a new default RoverNode client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewRoverNodeClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client RoverNodeClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newRoverNodeClientFromBaseClient(baseClient, configProvider)
}

func newRoverNodeClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client RoverNodeClient, err error) {
	// RoverNode service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("RoverNode"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = RoverNodeClient{BaseClient: baseClient}
	client.BasePath = "20201210"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *RoverNodeClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("rover", "https://rover.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *RoverNodeClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *RoverNodeClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeRoverNodeCompartment Moves a rover node into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/ChangeRoverNodeCompartment.go.html to see an example of how to use ChangeRoverNodeCompartment API.
// A default retry strategy applies to this operation ChangeRoverNodeCompartment()
func (client RoverNodeClient) ChangeRoverNodeCompartment(ctx context.Context, request ChangeRoverNodeCompartmentRequest) (response ChangeRoverNodeCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeRoverNodeCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeRoverNodeCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeRoverNodeCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeRoverNodeCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeRoverNodeCompartmentResponse")
	}
	return
}

// changeRoverNodeCompartment implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) changeRoverNodeCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverNodes/{roverNodeId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeRoverNodeCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNode/ChangeRoverNodeCompartment"
		err = common.PostProcessServiceError(err, "RoverNode", "ChangeRoverNodeCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateRoverNode Creates a new RoverNode.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/CreateRoverNode.go.html to see an example of how to use CreateRoverNode API.
// A default retry strategy applies to this operation CreateRoverNode()
func (client RoverNodeClient) CreateRoverNode(ctx context.Context, request CreateRoverNodeRequest) (response CreateRoverNodeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createRoverNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateRoverNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateRoverNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRoverNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRoverNodeResponse")
	}
	return
}

// createRoverNode implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) createRoverNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverNodes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateRoverNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNode/CreateRoverNode"
		err = common.PostProcessServiceError(err, "RoverNode", "CreateRoverNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteRoverNode Deletes a RoverNode resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/DeleteRoverNode.go.html to see an example of how to use DeleteRoverNode API.
// A default retry strategy applies to this operation DeleteRoverNode()
func (client RoverNodeClient) DeleteRoverNode(ctx context.Context, request DeleteRoverNodeRequest) (response DeleteRoverNodeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteRoverNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRoverNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRoverNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRoverNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRoverNodeResponse")
	}
	return
}

// deleteRoverNode implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) deleteRoverNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/roverNodes/{roverNodeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteRoverNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNode/DeleteRoverNode"
		err = common.PostProcessServiceError(err, "RoverNode", "DeleteRoverNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRoverNode Gets a RoverNode by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/GetRoverNode.go.html to see an example of how to use GetRoverNode API.
// A default retry strategy applies to this operation GetRoverNode()
func (client RoverNodeClient) GetRoverNode(ctx context.Context, request GetRoverNodeRequest) (response GetRoverNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRoverNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRoverNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRoverNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRoverNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRoverNodeResponse")
	}
	return
}

// getRoverNode implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) getRoverNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverNodes/{roverNodeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRoverNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNode/GetRoverNode"
		err = common.PostProcessServiceError(err, "RoverNode", "GetRoverNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRoverNodeCertificate Get the certificate for a rover node
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/GetRoverNodeCertificate.go.html to see an example of how to use GetRoverNodeCertificate API.
// A default retry strategy applies to this operation GetRoverNodeCertificate()
func (client RoverNodeClient) GetRoverNodeCertificate(ctx context.Context, request GetRoverNodeCertificateRequest) (response GetRoverNodeCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRoverNodeCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRoverNodeCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRoverNodeCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRoverNodeCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRoverNodeCertificateResponse")
	}
	return
}

// getRoverNodeCertificate implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) getRoverNodeCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverNodes/{roverNodeId}/certificate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRoverNodeCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNodeCertificate/GetRoverNodeCertificate"
		err = common.PostProcessServiceError(err, "RoverNode", "GetRoverNodeCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRoverNodeEncryptionKey Get the data encryption key for a rover node.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/GetRoverNodeEncryptionKey.go.html to see an example of how to use GetRoverNodeEncryptionKey API.
// A default retry strategy applies to this operation GetRoverNodeEncryptionKey()
func (client RoverNodeClient) GetRoverNodeEncryptionKey(ctx context.Context, request GetRoverNodeEncryptionKeyRequest) (response GetRoverNodeEncryptionKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRoverNodeEncryptionKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRoverNodeEncryptionKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRoverNodeEncryptionKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRoverNodeEncryptionKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRoverNodeEncryptionKeyResponse")
	}
	return
}

// getRoverNodeEncryptionKey implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) getRoverNodeEncryptionKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverNodes/{roverNodeId}/encryptionKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRoverNodeEncryptionKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNodeEncryptionKey/GetRoverNodeEncryptionKey"
		err = common.PostProcessServiceError(err, "RoverNode", "GetRoverNodeEncryptionKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRoverNodeGetRpt Get the resource principal token for a rover node
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/GetRoverNodeGetRpt.go.html to see an example of how to use GetRoverNodeGetRpt API.
// A default retry strategy applies to this operation GetRoverNodeGetRpt()
func (client RoverNodeClient) GetRoverNodeGetRpt(ctx context.Context, request GetRoverNodeGetRptRequest) (response GetRoverNodeGetRptResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRoverNodeGetRpt, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRoverNodeGetRptResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRoverNodeGetRptResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRoverNodeGetRptResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRoverNodeGetRptResponse")
	}
	return
}

// getRoverNodeGetRpt implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) getRoverNodeGetRpt(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverNodes/{roverNodeId}/getRpt", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRoverNodeGetRptResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNodeGetRpt/GetRoverNodeGetRpt"
		err = common.PostProcessServiceError(err, "RoverNode", "GetRoverNodeGetRpt", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRoverNodes Returns a list of RoverNodes.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/ListRoverNodes.go.html to see an example of how to use ListRoverNodes API.
// A default retry strategy applies to this operation ListRoverNodes()
func (client RoverNodeClient) ListRoverNodes(ctx context.Context, request ListRoverNodesRequest) (response ListRoverNodesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRoverNodes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRoverNodesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRoverNodesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRoverNodesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRoverNodesResponse")
	}
	return
}

// listRoverNodes implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) listRoverNodes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverNodes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRoverNodesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNode/ListRoverNodes"
		err = common.PostProcessServiceError(err, "RoverNode", "ListRoverNodes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RoverNodeActionRetrieveCaBundle Retrieve Ca Bundle for a rover node
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/RoverNodeActionRetrieveCaBundle.go.html to see an example of how to use RoverNodeActionRetrieveCaBundle API.
// A default retry strategy applies to this operation RoverNodeActionRetrieveCaBundle()
func (client RoverNodeClient) RoverNodeActionRetrieveCaBundle(ctx context.Context, request RoverNodeActionRetrieveCaBundleRequest) (response RoverNodeActionRetrieveCaBundleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.roverNodeActionRetrieveCaBundle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RoverNodeActionRetrieveCaBundleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RoverNodeActionRetrieveCaBundleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RoverNodeActionRetrieveCaBundleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RoverNodeActionRetrieveCaBundleResponse")
	}
	return
}

// roverNodeActionRetrieveCaBundle implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) roverNodeActionRetrieveCaBundle(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverNodes/{roverNodeId}/actions/retrieveCaBundle", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RoverNodeActionRetrieveCaBundleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNode/RoverNodeActionRetrieveCaBundle"
		err = common.PostProcessServiceError(err, "RoverNode", "RoverNodeActionRetrieveCaBundle", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RoverNodeActionSetKey Get the resource principal public key for a rover node
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/RoverNodeActionSetKey.go.html to see an example of how to use RoverNodeActionSetKey API.
// A default retry strategy applies to this operation RoverNodeActionSetKey()
func (client RoverNodeClient) RoverNodeActionSetKey(ctx context.Context, request RoverNodeActionSetKeyRequest) (response RoverNodeActionSetKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.roverNodeActionSetKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RoverNodeActionSetKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RoverNodeActionSetKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RoverNodeActionSetKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RoverNodeActionSetKeyResponse")
	}
	return
}

// roverNodeActionSetKey implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) roverNodeActionSetKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverNodes/{roverNodeId}/actions/setKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RoverNodeActionSetKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNodeSetKey/RoverNodeActionSetKey"
		err = common.PostProcessServiceError(err, "RoverNode", "RoverNodeActionSetKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RoverNodeGenerateCertificate Request to generate certificate for a roverNode.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/RoverNodeGenerateCertificate.go.html to see an example of how to use RoverNodeGenerateCertificate API.
// A default retry strategy applies to this operation RoverNodeGenerateCertificate()
func (client RoverNodeClient) RoverNodeGenerateCertificate(ctx context.Context, request RoverNodeGenerateCertificateRequest) (response RoverNodeGenerateCertificateResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.roverNodeGenerateCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RoverNodeGenerateCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RoverNodeGenerateCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RoverNodeGenerateCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RoverNodeGenerateCertificateResponse")
	}
	return
}

// roverNodeGenerateCertificate implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) roverNodeGenerateCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverNodes/{roverNodeId}/actions/generateCertificate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RoverNodeGenerateCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNode/RoverNodeGenerateCertificate"
		err = common.PostProcessServiceError(err, "RoverNode", "RoverNodeGenerateCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RoverNodeRenewCertificate Request to renew certificate for a roverNode.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/RoverNodeRenewCertificate.go.html to see an example of how to use RoverNodeRenewCertificate API.
// A default retry strategy applies to this operation RoverNodeRenewCertificate()
func (client RoverNodeClient) RoverNodeRenewCertificate(ctx context.Context, request RoverNodeRenewCertificateRequest) (response RoverNodeRenewCertificateResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.roverNodeRenewCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RoverNodeRenewCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RoverNodeRenewCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RoverNodeRenewCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RoverNodeRenewCertificateResponse")
	}
	return
}

// roverNodeRenewCertificate implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) roverNodeRenewCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverNodes/{roverNodeId}/actions/renewCertificate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RoverNodeRenewCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNode/RoverNodeRenewCertificate"
		err = common.PostProcessServiceError(err, "RoverNode", "RoverNodeRenewCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RoverNodeReplaceCertificateAuthority Request to replace certificate authority for a roverNode.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/RoverNodeReplaceCertificateAuthority.go.html to see an example of how to use RoverNodeReplaceCertificateAuthority API.
// A default retry strategy applies to this operation RoverNodeReplaceCertificateAuthority()
func (client RoverNodeClient) RoverNodeReplaceCertificateAuthority(ctx context.Context, request RoverNodeReplaceCertificateAuthorityRequest) (response RoverNodeReplaceCertificateAuthorityResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.roverNodeReplaceCertificateAuthority, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RoverNodeReplaceCertificateAuthorityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RoverNodeReplaceCertificateAuthorityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RoverNodeReplaceCertificateAuthorityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RoverNodeReplaceCertificateAuthorityResponse")
	}
	return
}

// roverNodeReplaceCertificateAuthority implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) roverNodeReplaceCertificateAuthority(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverNodes/{roverNodeId}/actions/replaceCertificateAuthority", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RoverNodeReplaceCertificateAuthorityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNode/RoverNodeReplaceCertificateAuthority"
		err = common.PostProcessServiceError(err, "RoverNode", "RoverNodeReplaceCertificateAuthority", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RoverNodeRetrieveLeafCertificate Retrieve the leaf certificate info for a rover node
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/RoverNodeRetrieveLeafCertificate.go.html to see an example of how to use RoverNodeRetrieveLeafCertificate API.
// A default retry strategy applies to this operation RoverNodeRetrieveLeafCertificate()
func (client RoverNodeClient) RoverNodeRetrieveLeafCertificate(ctx context.Context, request RoverNodeRetrieveLeafCertificateRequest) (response RoverNodeRetrieveLeafCertificateResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.roverNodeRetrieveLeafCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RoverNodeRetrieveLeafCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RoverNodeRetrieveLeafCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RoverNodeRetrieveLeafCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RoverNodeRetrieveLeafCertificateResponse")
	}
	return
}

// roverNodeRetrieveLeafCertificate implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) roverNodeRetrieveLeafCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverNodes/{roverNodeId}/actions/retrieveLeafCertificate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RoverNodeRetrieveLeafCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNode/RoverNodeRetrieveLeafCertificate"
		err = common.PostProcessServiceError(err, "RoverNode", "RoverNodeRetrieveLeafCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRoverNode Updates the RoverNode
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/UpdateRoverNode.go.html to see an example of how to use UpdateRoverNode API.
// A default retry strategy applies to this operation UpdateRoverNode()
func (client RoverNodeClient) UpdateRoverNode(ctx context.Context, request UpdateRoverNodeRequest) (response UpdateRoverNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRoverNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRoverNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRoverNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRoverNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRoverNodeResponse")
	}
	return
}

// updateRoverNode implements the OCIOperation interface (enables retrying operations)
func (client RoverNodeClient) updateRoverNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/roverNodes/{roverNodeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateRoverNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/rover/20201210/RoverNode/UpdateRoverNode"
		err = common.PostProcessServiceError(err, "RoverNode", "UpdateRoverNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
