// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Continuous Consumption Discounts API
//
// Use the Continuous Consumption Discounts API to manage continuous consumption discount operations.
//

package ccdoffers

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// CcdCommitmentsClient a client for CcdCommitments
type CcdCommitmentsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewCcdCommitmentsClientWithConfigurationProvider Creates a new default CcdCommitments client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewCcdCommitmentsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client CcdCommitmentsClient, err error) {
	if enabled := common.CheckForEnabledServices("ccdoffers"); !enabled {
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
	return newCcdCommitmentsClientFromBaseClient(baseClient, provider)
}

// NewCcdCommitmentsClientWithOboToken Creates a new default CcdCommitments client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewCcdCommitmentsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client CcdCommitmentsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newCcdCommitmentsClientFromBaseClient(baseClient, configProvider)
}

func newCcdCommitmentsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client CcdCommitmentsClient, err error) {
	// CcdCommitments service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("CcdCommitments"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = CcdCommitmentsClient{BaseClient: baseClient}
	client.BasePath = "20240601"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *CcdCommitmentsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ccdoffers", "https://ccdOffers.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *CcdCommitmentsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *CcdCommitmentsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateCcdCommitment Request to create a new CCD commitment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ccdoffers/CreateCcdCommitment.go.html to see an example of how to use CreateCcdCommitment API.
// A default retry strategy applies to this operation CreateCcdCommitment()
func (client CcdCommitmentsClient) CreateCcdCommitment(ctx context.Context, request CreateCcdCommitmentRequest) (response CreateCcdCommitmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCcdCommitment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCcdCommitmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCcdCommitmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCcdCommitmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCcdCommitmentResponse")
	}
	return
}

// createCcdCommitment implements the OCIOperation interface (enables retrying operations)
func (client CcdCommitmentsClient) createCcdCommitment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/ccdCommitments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCcdCommitmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "ccdCommitments", "CreateCcdCommitment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "CcdCommitments", "CreateCcdCommitment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCcdCommitment Terminates the CCD commitment. You can't terminate an already canceled or ended commitment. If the CCD commitment has already started based on `timeCcdCommitmentStarts`, then you can only terminate within a week of its effective `timeCcdCommitmentStarts` time.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ccdoffers/DeleteCcdCommitment.go.html to see an example of how to use DeleteCcdCommitment API.
// A default retry strategy applies to this operation DeleteCcdCommitment()
func (client CcdCommitmentsClient) DeleteCcdCommitment(ctx context.Context, request DeleteCcdCommitmentRequest) (response DeleteCcdCommitmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCcdCommitment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCcdCommitmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCcdCommitmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCcdCommitmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCcdCommitmentResponse")
	}
	return
}

// deleteCcdCommitment implements the OCIOperation interface (enables retrying operations)
func (client CcdCommitmentsClient) deleteCcdCommitment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/ccdCommitments/{ccdCommitmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCcdCommitmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "ccdCommitments", "DeleteCcdCommitment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "CcdCommitments", "DeleteCcdCommitment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCcdCommitment Gets the CCD commitment using the CCD commitment ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ccdoffers/GetCcdCommitment.go.html to see an example of how to use GetCcdCommitment API.
// A default retry strategy applies to this operation GetCcdCommitment()
func (client CcdCommitmentsClient) GetCcdCommitment(ctx context.Context, request GetCcdCommitmentRequest) (response GetCcdCommitmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCcdCommitment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCcdCommitmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCcdCommitmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCcdCommitmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCcdCommitmentResponse")
	}
	return
}

// getCcdCommitment implements the OCIOperation interface (enables retrying operations)
func (client CcdCommitmentsClient) getCcdCommitment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ccdCommitments/{ccdCommitmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCcdCommitmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "ccdCommitments", "GetCcdCommitment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "CcdCommitments", "GetCcdCommitment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCcdCommitments Lists the CCD commitments.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ccdoffers/ListCcdCommitments.go.html to see an example of how to use ListCcdCommitments API.
// A default retry strategy applies to this operation ListCcdCommitments()
func (client CcdCommitmentsClient) ListCcdCommitments(ctx context.Context, request ListCcdCommitmentsRequest) (response ListCcdCommitmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCcdCommitments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCcdCommitmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCcdCommitmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCcdCommitmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCcdCommitmentsResponse")
	}
	return
}

// listCcdCommitments implements the OCIOperation interface (enables retrying operations)
func (client CcdCommitmentsClient) listCcdCommitments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ccdCommitments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCcdCommitmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "ccdCommitments", "ListCcdCommitments")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "CcdCommitments", "ListCcdCommitments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCcdCommitment Updates the free-form tags, defined tags, or both on a CCD commitment. Commercial commitment attributes are immutable.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ccdoffers/UpdateCcdCommitment.go.html to see an example of how to use UpdateCcdCommitment API.
// A default retry strategy applies to this operation UpdateCcdCommitment()
func (client CcdCommitmentsClient) UpdateCcdCommitment(ctx context.Context, request UpdateCcdCommitmentRequest) (response UpdateCcdCommitmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCcdCommitment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCcdCommitmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCcdCommitmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCcdCommitmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCcdCommitmentResponse")
	}
	return
}

// updateCcdCommitment implements the OCIOperation interface (enables retrying operations)
func (client CcdCommitmentsClient) updateCcdCommitment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/ccdCommitments/{ccdCommitmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCcdCommitmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "ccdCommitments", "UpdateCcdCommitment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "CcdCommitments", "UpdateCcdCommitment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
