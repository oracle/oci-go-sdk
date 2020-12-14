// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v31/common"
	"github.com/oracle/oci-go-sdk/v31/common/auth"
	"net/http"
)

//RoverNodeClient a client for RoverNode
type RoverNodeClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewRoverNodeClientWithConfigurationProvider Creates a new default RoverNode client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewRoverNodeClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client RoverNodeClient, err error) {
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
//  as well as reading the region
func NewRoverNodeClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client RoverNodeClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newRoverNodeClientFromBaseClient(baseClient, configProvider)
}

func newRoverNodeClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client RoverNodeClient, err error) {
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
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *RoverNodeClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeRoverNodeCompartment Moves a rover node into a different compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/ChangeRoverNodeCompartment.go.html to see an example of how to use ChangeRoverNodeCompartment API.
func (client RoverNodeClient) ChangeRoverNodeCompartment(ctx context.Context, request ChangeRoverNodeCompartmentRequest) (response ChangeRoverNodeCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client RoverNodeClient) changeRoverNodeCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverNodes/{roverNodeId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeRoverNodeCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateRoverNode Creates a new RoverNode.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/CreateRoverNode.go.html to see an example of how to use CreateRoverNode API.
func (client RoverNodeClient) CreateRoverNode(ctx context.Context, request CreateRoverNodeRequest) (response CreateRoverNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client RoverNodeClient) createRoverNode(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverNodes")
	if err != nil {
		return nil, err
	}

	var response CreateRoverNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteRoverNode Deletes a RoverNode resource by identifier
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/DeleteRoverNode.go.html to see an example of how to use DeleteRoverNode API.
func (client RoverNodeClient) DeleteRoverNode(ctx context.Context, request DeleteRoverNodeRequest) (response DeleteRoverNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
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
func (client RoverNodeClient) deleteRoverNode(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/roverNodes/{roverNodeId}")
	if err != nil {
		return nil, err
	}

	var response DeleteRoverNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRoverNode Gets a RoverNode by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/GetRoverNode.go.html to see an example of how to use GetRoverNode API.
func (client RoverNodeClient) GetRoverNode(ctx context.Context, request GetRoverNodeRequest) (response GetRoverNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client RoverNodeClient) getRoverNode(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverNodes/{roverNodeId}")
	if err != nil {
		return nil, err
	}

	var response GetRoverNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRoverNodeCertificate Get the certificate for a rover node
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/GetRoverNodeCertificate.go.html to see an example of how to use GetRoverNodeCertificate API.
func (client RoverNodeClient) GetRoverNodeCertificate(ctx context.Context, request GetRoverNodeCertificateRequest) (response GetRoverNodeCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client RoverNodeClient) getRoverNodeCertificate(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverNodes/{roverNodeId}/certificate")
	if err != nil {
		return nil, err
	}

	var response GetRoverNodeCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRoverNodeEncryptionKey Get the data encryption key for a rover node.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/GetRoverNodeEncryptionKey.go.html to see an example of how to use GetRoverNodeEncryptionKey API.
func (client RoverNodeClient) GetRoverNodeEncryptionKey(ctx context.Context, request GetRoverNodeEncryptionKeyRequest) (response GetRoverNodeEncryptionKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client RoverNodeClient) getRoverNodeEncryptionKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverNodes/{roverNodeId}/encryptionKey")
	if err != nil {
		return nil, err
	}

	var response GetRoverNodeEncryptionKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRoverNodeGetRpt Get the resource principal token for a rover node
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/GetRoverNodeGetRpt.go.html to see an example of how to use GetRoverNodeGetRpt API.
func (client RoverNodeClient) GetRoverNodeGetRpt(ctx context.Context, request GetRoverNodeGetRptRequest) (response GetRoverNodeGetRptResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client RoverNodeClient) getRoverNodeGetRpt(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverNodes/{roverNodeId}/getRpt")
	if err != nil {
		return nil, err
	}

	var response GetRoverNodeGetRptResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRoverNodes Returns a list of RoverNodes.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/ListRoverNodes.go.html to see an example of how to use ListRoverNodes API.
func (client RoverNodeClient) ListRoverNodes(ctx context.Context, request ListRoverNodesRequest) (response ListRoverNodesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client RoverNodeClient) listRoverNodes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverNodes")
	if err != nil {
		return nil, err
	}

	var response ListRoverNodesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RoverNodeActionSetKey Get the resource principal public key for a rover node
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/RoverNodeActionSetKey.go.html to see an example of how to use RoverNodeActionSetKey API.
func (client RoverNodeClient) RoverNodeActionSetKey(ctx context.Context, request RoverNodeActionSetKeyRequest) (response RoverNodeActionSetKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client RoverNodeClient) roverNodeActionSetKey(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverNodes/{roverNodeId}/actions/setKey")
	if err != nil {
		return nil, err
	}

	var response RoverNodeActionSetKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRoverNode Updates the RoverNode
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/UpdateRoverNode.go.html to see an example of how to use UpdateRoverNode API.
func (client RoverNodeClient) UpdateRoverNode(ctx context.Context, request UpdateRoverNodeRequest) (response UpdateRoverNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client RoverNodeClient) updateRoverNode(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/roverNodes/{roverNodeId}")
	if err != nil {
		return nil, err
	}

	var response UpdateRoverNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
