// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v60/common"
	"github.com/oracle/oci-go-sdk/v60/common/auth"
	"net/http"
)

//RoverEntitlementClient a client for RoverEntitlement
type RoverEntitlementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewRoverEntitlementClientWithConfigurationProvider Creates a new default RoverEntitlement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewRoverEntitlementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client RoverEntitlementClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newRoverEntitlementClientFromBaseClient(baseClient, provider)
}

// NewRoverEntitlementClientWithOboToken Creates a new default RoverEntitlement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewRoverEntitlementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client RoverEntitlementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newRoverEntitlementClientFromBaseClient(baseClient, configProvider)
}

func newRoverEntitlementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client RoverEntitlementClient, err error) {
	// RoverEntitlement service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = RoverEntitlementClient{BaseClient: baseClient}
	client.BasePath = "20201210"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *RoverEntitlementClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("rover", "https://rover.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *RoverEntitlementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *RoverEntitlementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeRoverEntitlementCompartment Moves an entitlement into a different compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/ChangeRoverEntitlementCompartment.go.html to see an example of how to use ChangeRoverEntitlementCompartment API.
// A default retry strategy applies to this operation ChangeRoverEntitlementCompartment()
func (client RoverEntitlementClient) ChangeRoverEntitlementCompartment(ctx context.Context, request ChangeRoverEntitlementCompartmentRequest) (response ChangeRoverEntitlementCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeRoverEntitlementCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeRoverEntitlementCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeRoverEntitlementCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeRoverEntitlementCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeRoverEntitlementCompartmentResponse")
	}
	return
}

// changeRoverEntitlementCompartment implements the OCIOperation interface (enables retrying operations)
func (client RoverEntitlementClient) changeRoverEntitlementCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverEntitlements/{roverEntitlementId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeRoverEntitlementCompartmentResponse
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

// CreateRoverEntitlement Create the Entitlement to use a Rover Device. It requires some offline process of review and signatures before request is granted.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/CreateRoverEntitlement.go.html to see an example of how to use CreateRoverEntitlement API.
// A default retry strategy applies to this operation CreateRoverEntitlement()
func (client RoverEntitlementClient) CreateRoverEntitlement(ctx context.Context, request CreateRoverEntitlementRequest) (response CreateRoverEntitlementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createRoverEntitlement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateRoverEntitlementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateRoverEntitlementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRoverEntitlementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRoverEntitlementResponse")
	}
	return
}

// createRoverEntitlement implements the OCIOperation interface (enables retrying operations)
func (client RoverEntitlementClient) createRoverEntitlement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/roverEntitlements", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateRoverEntitlementResponse
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

// DeleteRoverEntitlement Deletes a rover entitlement
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/DeleteRoverEntitlement.go.html to see an example of how to use DeleteRoverEntitlement API.
// A default retry strategy applies to this operation DeleteRoverEntitlement()
func (client RoverEntitlementClient) DeleteRoverEntitlement(ctx context.Context, request DeleteRoverEntitlementRequest) (response DeleteRoverEntitlementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteRoverEntitlement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRoverEntitlementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRoverEntitlementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRoverEntitlementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRoverEntitlementResponse")
	}
	return
}

// deleteRoverEntitlement implements the OCIOperation interface (enables retrying operations)
func (client RoverEntitlementClient) deleteRoverEntitlement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/roverEntitlements/{roverEntitlementId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteRoverEntitlementResponse
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

// GetRoverEntitlement Describes the Rover Device Entitlement in detail
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/GetRoverEntitlement.go.html to see an example of how to use GetRoverEntitlement API.
// A default retry strategy applies to this operation GetRoverEntitlement()
func (client RoverEntitlementClient) GetRoverEntitlement(ctx context.Context, request GetRoverEntitlementRequest) (response GetRoverEntitlementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getRoverEntitlement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRoverEntitlementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRoverEntitlementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRoverEntitlementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRoverEntitlementResponse")
	}
	return
}

// getRoverEntitlement implements the OCIOperation interface (enables retrying operations)
func (client RoverEntitlementClient) getRoverEntitlement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverEntitlements/{roverEntitlementId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRoverEntitlementResponse
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

// ListRoverEntitlements Returns a list of RoverEntitlements.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/ListRoverEntitlements.go.html to see an example of how to use ListRoverEntitlements API.
// A default retry strategy applies to this operation ListRoverEntitlements()
func (client RoverEntitlementClient) ListRoverEntitlements(ctx context.Context, request ListRoverEntitlementsRequest) (response ListRoverEntitlementsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRoverEntitlements, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRoverEntitlementsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRoverEntitlementsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRoverEntitlementsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRoverEntitlementsResponse")
	}
	return
}

// listRoverEntitlements implements the OCIOperation interface (enables retrying operations)
func (client RoverEntitlementClient) listRoverEntitlements(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/roverEntitlements", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRoverEntitlementsResponse
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

// UpdateRoverEntitlement Updates the RoverEntitlement
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/rover/UpdateRoverEntitlement.go.html to see an example of how to use UpdateRoverEntitlement API.
// A default retry strategy applies to this operation UpdateRoverEntitlement()
func (client RoverEntitlementClient) UpdateRoverEntitlement(ctx context.Context, request UpdateRoverEntitlementRequest) (response UpdateRoverEntitlementResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRoverEntitlement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRoverEntitlementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRoverEntitlementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRoverEntitlementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRoverEntitlementResponse")
	}
	return
}

// updateRoverEntitlement implements the OCIOperation interface (enables retrying operations)
func (client RoverEntitlementClient) updateRoverEntitlement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/roverEntitlements/{roverEntitlementId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateRoverEntitlementResponse
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
