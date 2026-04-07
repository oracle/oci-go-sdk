// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Limits Increase API
//
// Use the Limits Increase API to work with limit increase requests.
// For more information, see
// Working with Limit Increase Requests (https://docs.oracle.com/iaas/Content/General/service-limits/requests.htm).
//

package limitsincrease

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// LimitsIncreaseClient a client for LimitsIncrease
type LimitsIncreaseClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLimitsIncreaseClientWithConfigurationProvider Creates a new default LimitsIncrease client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLimitsIncreaseClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LimitsIncreaseClient, err error) {
	if enabled := common.CheckForEnabledServices("limitsincrease"); !enabled {
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
	return newLimitsIncreaseClientFromBaseClient(baseClient, provider)
}

// NewLimitsIncreaseClientWithOboToken Creates a new default LimitsIncrease client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewLimitsIncreaseClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LimitsIncreaseClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newLimitsIncreaseClientFromBaseClient(baseClient, configProvider)
}

func newLimitsIncreaseClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LimitsIncreaseClient, err error) {
	// LimitsIncrease service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("LimitsIncrease"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = LimitsIncreaseClient{BaseClient: baseClient}
	client.BasePath = "20251101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LimitsIncreaseClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("limitsincrease", "https://limits.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LimitsIncreaseClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *LimitsIncreaseClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelLimitsIncreaseItemRequest Withdraws the specified item from the limit increase request.
// For more information, see
// Withdrawing an Item from a Limit Increase Request (https://docs.oracle.com/iaas/Content/General/service-limits/withdraw-item-request.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limitsincrease/CancelLimitsIncreaseItemRequest.go.html to see an example of how to use CancelLimitsIncreaseItemRequest API.
// A default retry strategy applies to this operation CancelLimitsIncreaseItemRequest()
func (client LimitsIncreaseClient) CancelLimitsIncreaseItemRequest(ctx context.Context, request CancelLimitsIncreaseItemRequestRequest) (response CancelLimitsIncreaseItemRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelLimitsIncreaseItemRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelLimitsIncreaseItemRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelLimitsIncreaseItemRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelLimitsIncreaseItemRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelLimitsIncreaseItemRequestResponse")
	}
	return
}

// cancelLimitsIncreaseItemRequest implements the OCIOperation interface (enables retrying operations)
func (client LimitsIncreaseClient) cancelLimitsIncreaseItemRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/limitsIncreaseItemRequests/{limitsIncreaseItemRequestId}/actions/cancel", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelLimitsIncreaseItemRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "limitsIncrease", "CancelLimitsIncreaseItemRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "LimitsIncrease", "CancelLimitsIncreaseItemRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelLimitsIncreaseRequest Withdraws the specified limit increase request.
// For more information, see
// Withdrawing a Limit Increase Request (https://docs.oracle.com/iaas/Content/General/service-limits/withdraw-request.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limitsincrease/CancelLimitsIncreaseRequest.go.html to see an example of how to use CancelLimitsIncreaseRequest API.
// A default retry strategy applies to this operation CancelLimitsIncreaseRequest()
func (client LimitsIncreaseClient) CancelLimitsIncreaseRequest(ctx context.Context, request CancelLimitsIncreaseRequestRequest) (response CancelLimitsIncreaseRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelLimitsIncreaseRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelLimitsIncreaseRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelLimitsIncreaseRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelLimitsIncreaseRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelLimitsIncreaseRequestResponse")
	}
	return
}

// cancelLimitsIncreaseRequest implements the OCIOperation interface (enables retrying operations)
func (client LimitsIncreaseClient) cancelLimitsIncreaseRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/limitsIncreaseRequests/{limitsIncreaseRequestId}/actions/cancel", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelLimitsIncreaseRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "limitsIncrease", "CancelLimitsIncreaseRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "LimitsIncrease", "CancelLimitsIncreaseRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateLimitsIncreaseRequest Creates a limit increase request in the specified compartment.
// For more information, see
// Creating a Limit Increase Request (https://docs.oracle.com/iaas/Content/General/service-limits/create-request.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limitsincrease/CreateLimitsIncreaseRequest.go.html to see an example of how to use CreateLimitsIncreaseRequest API.
// A default retry strategy applies to this operation CreateLimitsIncreaseRequest()
func (client LimitsIncreaseClient) CreateLimitsIncreaseRequest(ctx context.Context, request CreateLimitsIncreaseRequestRequest) (response CreateLimitsIncreaseRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createLimitsIncreaseRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLimitsIncreaseRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLimitsIncreaseRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLimitsIncreaseRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLimitsIncreaseRequestResponse")
	}
	return
}

// createLimitsIncreaseRequest implements the OCIOperation interface (enables retrying operations)
func (client LimitsIncreaseClient) createLimitsIncreaseRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/limitsIncreaseRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateLimitsIncreaseRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "limitsIncrease", "CreateLimitsIncreaseRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "LimitsIncrease", "CreateLimitsIncreaseRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteLimitsIncreaseRequest Deletes the specified limit increase request.
// For more information, see
// Deleting a Limit Increase Request (https://docs.oracle.com/iaas/Content/General/service-limits/delete-request.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limitsincrease/DeleteLimitsIncreaseRequest.go.html to see an example of how to use DeleteLimitsIncreaseRequest API.
// A default retry strategy applies to this operation DeleteLimitsIncreaseRequest()
func (client LimitsIncreaseClient) DeleteLimitsIncreaseRequest(ctx context.Context, request DeleteLimitsIncreaseRequestRequest) (response DeleteLimitsIncreaseRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLimitsIncreaseRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLimitsIncreaseRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLimitsIncreaseRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLimitsIncreaseRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLimitsIncreaseRequestResponse")
	}
	return
}

// deleteLimitsIncreaseRequest implements the OCIOperation interface (enables retrying operations)
func (client LimitsIncreaseClient) deleteLimitsIncreaseRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/limitsIncreaseRequests/{limitsIncreaseRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteLimitsIncreaseRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "limitsIncrease", "DeleteLimitsIncreaseRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "LimitsIncrease", "DeleteLimitsIncreaseRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLimitsIncreaseItemRequest Gets the specified item from the limit increase request.
// For more information, see
// Getting Details for a Limit Increase Request (https://docs.oracle.com/iaas/Content/General/service-limits/get-request.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limitsincrease/GetLimitsIncreaseItemRequest.go.html to see an example of how to use GetLimitsIncreaseItemRequest API.
// A default retry strategy applies to this operation GetLimitsIncreaseItemRequest()
func (client LimitsIncreaseClient) GetLimitsIncreaseItemRequest(ctx context.Context, request GetLimitsIncreaseItemRequestRequest) (response GetLimitsIncreaseItemRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLimitsIncreaseItemRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLimitsIncreaseItemRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLimitsIncreaseItemRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLimitsIncreaseItemRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLimitsIncreaseItemRequestResponse")
	}
	return
}

// getLimitsIncreaseItemRequest implements the OCIOperation interface (enables retrying operations)
func (client LimitsIncreaseClient) getLimitsIncreaseItemRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/limitsIncreaseItemRequests/{limitsIncreaseItemRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLimitsIncreaseItemRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "limitsIncrease", "GetLimitsIncreaseItemRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "LimitsIncrease", "GetLimitsIncreaseItemRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLimitsIncreaseRequest Gets the specified limit increase request.
// For more information, see
// Getting Details for a Limit Increase Request (https://docs.oracle.com/iaas/Content/General/service-limits/get-request.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limitsincrease/GetLimitsIncreaseRequest.go.html to see an example of how to use GetLimitsIncreaseRequest API.
// A default retry strategy applies to this operation GetLimitsIncreaseRequest()
func (client LimitsIncreaseClient) GetLimitsIncreaseRequest(ctx context.Context, request GetLimitsIncreaseRequestRequest) (response GetLimitsIncreaseRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLimitsIncreaseRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLimitsIncreaseRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLimitsIncreaseRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLimitsIncreaseRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLimitsIncreaseRequestResponse")
	}
	return
}

// getLimitsIncreaseRequest implements the OCIOperation interface (enables retrying operations)
func (client LimitsIncreaseClient) getLimitsIncreaseRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/limitsIncreaseRequests/{limitsIncreaseRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLimitsIncreaseRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "limitsIncrease", "GetLimitsIncreaseRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "LimitsIncrease", "GetLimitsIncreaseRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLimitsIncreaseItemRequests Lists items in the specified limit increase request.
// For more information, see
// Getting Details for a Limit Increase Request (https://docs.oracle.com/iaas/Content/General/service-limits/get-request.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limitsincrease/ListLimitsIncreaseItemRequests.go.html to see an example of how to use ListLimitsIncreaseItemRequests API.
// A default retry strategy applies to this operation ListLimitsIncreaseItemRequests()
func (client LimitsIncreaseClient) ListLimitsIncreaseItemRequests(ctx context.Context, request ListLimitsIncreaseItemRequestsRequest) (response ListLimitsIncreaseItemRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLimitsIncreaseItemRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLimitsIncreaseItemRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLimitsIncreaseItemRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLimitsIncreaseItemRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLimitsIncreaseItemRequestsResponse")
	}
	return
}

// listLimitsIncreaseItemRequests implements the OCIOperation interface (enables retrying operations)
func (client LimitsIncreaseClient) listLimitsIncreaseItemRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/limitsIncreaseItemRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLimitsIncreaseItemRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "limitsIncrease", "ListLimitsIncreaseItemRequests")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "LimitsIncrease", "ListLimitsIncreaseItemRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLimitsIncreaseQuestions Gets the fields for the specified service and limit name.
// Service name is required if limit name is provided.
// If limit name is not provided, returns all questions for the specified service.
// If service name is not provided, returns all available questions.
// For more information, see
// Creating a Limit Increase Request (https://docs.oracle.com/iaas/Content/General/service-limits/create-request.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limitsincrease/ListLimitsIncreaseQuestions.go.html to see an example of how to use ListLimitsIncreaseQuestions API.
// A default retry strategy applies to this operation ListLimitsIncreaseQuestions()
func (client LimitsIncreaseClient) ListLimitsIncreaseQuestions(ctx context.Context, request ListLimitsIncreaseQuestionsRequest) (response ListLimitsIncreaseQuestionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLimitsIncreaseQuestions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLimitsIncreaseQuestionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLimitsIncreaseQuestionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLimitsIncreaseQuestionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLimitsIncreaseQuestionsResponse")
	}
	return
}

// listLimitsIncreaseQuestions implements the OCIOperation interface (enables retrying operations)
func (client LimitsIncreaseClient) listLimitsIncreaseQuestions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/limitsIncreaseQuestionnaires", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLimitsIncreaseQuestionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "limitsIncrease", "ListLimitsIncreaseQuestions")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "LimitsIncrease", "ListLimitsIncreaseQuestions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLimitsIncreaseRequests Lists limit increase requests in the specified compartment.
// For more information, see
// Listing Limit Increase Requests (https://docs.oracle.com/iaas/Content/General/service-limits/list-requests.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limitsincrease/ListLimitsIncreaseRequests.go.html to see an example of how to use ListLimitsIncreaseRequests API.
// A default retry strategy applies to this operation ListLimitsIncreaseRequests()
func (client LimitsIncreaseClient) ListLimitsIncreaseRequests(ctx context.Context, request ListLimitsIncreaseRequestsRequest) (response ListLimitsIncreaseRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLimitsIncreaseRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLimitsIncreaseRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLimitsIncreaseRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLimitsIncreaseRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLimitsIncreaseRequestsResponse")
	}
	return
}

// listLimitsIncreaseRequests implements the OCIOperation interface (enables retrying operations)
func (client LimitsIncreaseClient) listLimitsIncreaseRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/limitsIncreaseRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLimitsIncreaseRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "limitsIncrease", "ListLimitsIncreaseRequests")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "LimitsIncrease", "ListLimitsIncreaseRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchLimitsIncreaseRequest Adds one or more comments to the specified limit increase request.
// For more information, see
// Adding Comments to a Limit Increase Request (https://docs.oracle.com/iaas/Content/General/service-limits/comment-request.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limitsincrease/PatchLimitsIncreaseRequest.go.html to see an example of how to use PatchLimitsIncreaseRequest API.
// A default retry strategy applies to this operation PatchLimitsIncreaseRequest()
func (client LimitsIncreaseClient) PatchLimitsIncreaseRequest(ctx context.Context, request PatchLimitsIncreaseRequestRequest) (response PatchLimitsIncreaseRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchLimitsIncreaseRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchLimitsIncreaseRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchLimitsIncreaseRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchLimitsIncreaseRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchLimitsIncreaseRequestResponse")
	}
	return
}

// patchLimitsIncreaseRequest implements the OCIOperation interface (enables retrying operations)
func (client LimitsIncreaseClient) patchLimitsIncreaseRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/limitsIncreaseRequests/{limitsIncreaseRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchLimitsIncreaseRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "limitsIncrease", "PatchLimitsIncreaseRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "LimitsIncrease", "PatchLimitsIncreaseRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateLimitsIncreaseRequest Updates the limit increase request.
// For more information, see
// Updating a Limit Increase Request (https://docs.oracle.com/iaas/Content/General/service-limits/update-request.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/limitsincrease/UpdateLimitsIncreaseRequest.go.html to see an example of how to use UpdateLimitsIncreaseRequest API.
// A default retry strategy applies to this operation UpdateLimitsIncreaseRequest()
func (client LimitsIncreaseClient) UpdateLimitsIncreaseRequest(ctx context.Context, request UpdateLimitsIncreaseRequestRequest) (response UpdateLimitsIncreaseRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLimitsIncreaseRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLimitsIncreaseRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLimitsIncreaseRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLimitsIncreaseRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLimitsIncreaseRequestResponse")
	}
	return
}

// updateLimitsIncreaseRequest implements the OCIOperation interface (enables retrying operations)
func (client LimitsIncreaseClient) updateLimitsIncreaseRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/limitsIncreaseRequests/{limitsIncreaseRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateLimitsIncreaseRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "limitsIncrease", "UpdateLimitsIncreaseRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "LimitsIncrease", "UpdateLimitsIncreaseRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
