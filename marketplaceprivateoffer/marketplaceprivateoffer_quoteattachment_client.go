// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePrivateOffer API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplaceprivateoffer

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// QuoteAttachmentClient a client for QuoteAttachment
type QuoteAttachmentClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewQuoteAttachmentClientWithConfigurationProvider Creates a new default QuoteAttachment client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewQuoteAttachmentClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client QuoteAttachmentClient, err error) {
	if enabled := common.CheckForEnabledServices("marketplaceprivateoffer"); !enabled {
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
	return newQuoteAttachmentClientFromBaseClient(baseClient, provider)
}

// NewQuoteAttachmentClientWithOboToken Creates a new default QuoteAttachment client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewQuoteAttachmentClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client QuoteAttachmentClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newQuoteAttachmentClientFromBaseClient(baseClient, configProvider)
}

func newQuoteAttachmentClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client QuoteAttachmentClient, err error) {
	// QuoteAttachment service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("QuoteAttachment"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = QuoteAttachmentClient{BaseClient: baseClient}
	client.BasePath = "20220901"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *QuoteAttachmentClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("marketplaceprivateoffer", "https://private-offer.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *QuoteAttachmentClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *QuoteAttachmentClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateQuoteAttachment Creates a new offer quote attachment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplaceprivateoffer/CreateQuoteAttachment.go.html to see an example of how to use CreateQuoteAttachment API.
// A default retry strategy applies to this operation CreateQuoteAttachment()
func (client QuoteAttachmentClient) CreateQuoteAttachment(ctx context.Context, request CreateQuoteAttachmentRequest) (response CreateQuoteAttachmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createQuoteAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateQuoteAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateQuoteAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateQuoteAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateQuoteAttachmentResponse")
	}
	return
}

// createQuoteAttachment implements the OCIOperation interface (enables retrying operations)
func (client QuoteAttachmentClient) createQuoteAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/offerQuotes/{offerQuoteId}/quoteAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateQuoteAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "quoteAttachment", "CreateQuoteAttachment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "QuoteAttachment", "CreateQuoteAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteQuoteAttachment Deletes an offer quote attachment resource by identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplaceprivateoffer/DeleteQuoteAttachment.go.html to see an example of how to use DeleteQuoteAttachment API.
// A default retry strategy applies to this operation DeleteQuoteAttachment()
func (client QuoteAttachmentClient) DeleteQuoteAttachment(ctx context.Context, request DeleteQuoteAttachmentRequest) (response DeleteQuoteAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteQuoteAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteQuoteAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteQuoteAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteQuoteAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteQuoteAttachmentResponse")
	}
	return
}

// deleteQuoteAttachment implements the OCIOperation interface (enables retrying operations)
func (client QuoteAttachmentClient) deleteQuoteAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/offerQuotes/{offerQuoteId}/quoteAttachments/{quoteAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteQuoteAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "quoteAttachment", "DeleteQuoteAttachment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "QuoteAttachment", "DeleteQuoteAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetQuoteAttachment Gets an offer quote attachment by identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplaceprivateoffer/GetQuoteAttachment.go.html to see an example of how to use GetQuoteAttachment API.
// A default retry strategy applies to this operation GetQuoteAttachment()
func (client QuoteAttachmentClient) GetQuoteAttachment(ctx context.Context, request GetQuoteAttachmentRequest) (response GetQuoteAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getQuoteAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetQuoteAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetQuoteAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetQuoteAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetQuoteAttachmentResponse")
	}
	return
}

// getQuoteAttachment implements the OCIOperation interface (enables retrying operations)
func (client QuoteAttachmentClient) getQuoteAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/offerQuotes/{offerQuoteId}/quoteAttachments/{quoteAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetQuoteAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "quoteAttachment", "GetQuoteAttachment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "QuoteAttachment", "GetQuoteAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetQuoteAttachmentContent Gets an offer quote attachment content by identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplaceprivateoffer/GetQuoteAttachmentContent.go.html to see an example of how to use GetQuoteAttachmentContent API.
// A default retry strategy applies to this operation GetQuoteAttachmentContent()
func (client QuoteAttachmentClient) GetQuoteAttachmentContent(ctx context.Context, request GetQuoteAttachmentContentRequest) (response GetQuoteAttachmentContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getQuoteAttachmentContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetQuoteAttachmentContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetQuoteAttachmentContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetQuoteAttachmentContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetQuoteAttachmentContentResponse")
	}
	return
}

// getQuoteAttachmentContent implements the OCIOperation interface (enables retrying operations)
func (client QuoteAttachmentClient) getQuoteAttachmentContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/offerQuotes/{offerQuoteId}/quoteAttachments/{quoteAttachmentId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetQuoteAttachmentContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "quoteAttachment", "GetQuoteAttachmentContent")
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "QuoteAttachment", "GetQuoteAttachmentContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListQuoteAttachments Returns a list of offer quote attachments. Requires either a reseller compartment ID or an ISV compartment ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplaceprivateoffer/ListQuoteAttachments.go.html to see an example of how to use ListQuoteAttachments API.
// A default retry strategy applies to this operation ListQuoteAttachments()
func (client QuoteAttachmentClient) ListQuoteAttachments(ctx context.Context, request ListQuoteAttachmentsRequest) (response ListQuoteAttachmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listQuoteAttachments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListQuoteAttachmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListQuoteAttachmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListQuoteAttachmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListQuoteAttachmentsResponse")
	}
	return
}

// listQuoteAttachments implements the OCIOperation interface (enables retrying operations)
func (client QuoteAttachmentClient) listQuoteAttachments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/offerQuotes/{offerQuoteId}/quoteAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListQuoteAttachmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "quoteAttachment", "ListQuoteAttachments")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "QuoteAttachment", "ListQuoteAttachments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
