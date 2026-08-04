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

// OfferQuoteClient a client for OfferQuote
type OfferQuoteClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOfferQuoteClientWithConfigurationProvider Creates a new default OfferQuote client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOfferQuoteClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OfferQuoteClient, err error) {
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
	return newOfferQuoteClientFromBaseClient(baseClient, provider)
}

// NewOfferQuoteClientWithOboToken Creates a new default OfferQuote client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOfferQuoteClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OfferQuoteClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOfferQuoteClientFromBaseClient(baseClient, configProvider)
}

func newOfferQuoteClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OfferQuoteClient, err error) {
	// OfferQuote service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OfferQuote"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OfferQuoteClient{BaseClient: baseClient}
	client.BasePath = "20220901"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OfferQuoteClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("marketplaceprivateoffer", "https://private-offer.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OfferQuoteClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OfferQuoteClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateOfferQuote Creates a new offer quote.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplaceprivateoffer/CreateOfferQuote.go.html to see an example of how to use CreateOfferQuote API.
// A default retry strategy applies to this operation CreateOfferQuote()
func (client OfferQuoteClient) CreateOfferQuote(ctx context.Context, request CreateOfferQuoteRequest) (response CreateOfferQuoteResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOfferQuote, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOfferQuoteResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOfferQuoteResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOfferQuoteResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOfferQuoteResponse")
	}
	return
}

// createOfferQuote implements the OCIOperation interface (enables retrying operations)
func (client OfferQuoteClient) createOfferQuote(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/offerQuotes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOfferQuoteResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "offerQuote", "CreateOfferQuote")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OfferQuote", "CreateOfferQuote", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOfferQuote Deletes an offer quote resource by identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplaceprivateoffer/DeleteOfferQuote.go.html to see an example of how to use DeleteOfferQuote API.
// A default retry strategy applies to this operation DeleteOfferQuote()
func (client OfferQuoteClient) DeleteOfferQuote(ctx context.Context, request DeleteOfferQuoteRequest) (response DeleteOfferQuoteResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOfferQuote, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOfferQuoteResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOfferQuoteResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOfferQuoteResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOfferQuoteResponse")
	}
	return
}

// deleteOfferQuote implements the OCIOperation interface (enables retrying operations)
func (client OfferQuoteClient) deleteOfferQuote(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/offerQuotes/{offerQuoteId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOfferQuoteResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "offerQuote", "DeleteOfferQuote")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OfferQuote", "DeleteOfferQuote", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOfferQuote Gets an offer quote by identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplaceprivateoffer/GetOfferQuote.go.html to see an example of how to use GetOfferQuote API.
// A default retry strategy applies to this operation GetOfferQuote()
func (client OfferQuoteClient) GetOfferQuote(ctx context.Context, request GetOfferQuoteRequest) (response GetOfferQuoteResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOfferQuote, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOfferQuoteResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOfferQuoteResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOfferQuoteResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOfferQuoteResponse")
	}
	return
}

// getOfferQuote implements the OCIOperation interface (enables retrying operations)
func (client OfferQuoteClient) getOfferQuote(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/offerQuotes/{offerQuoteId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOfferQuoteResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "offerQuote", "GetOfferQuote")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OfferQuote", "GetOfferQuote", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOfferQuoteInternalDetail Gets an offer quote internal details by identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplaceprivateoffer/GetOfferQuoteInternalDetail.go.html to see an example of how to use GetOfferQuoteInternalDetail API.
// A default retry strategy applies to this operation GetOfferQuoteInternalDetail()
func (client OfferQuoteClient) GetOfferQuoteInternalDetail(ctx context.Context, request GetOfferQuoteInternalDetailRequest) (response GetOfferQuoteInternalDetailResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOfferQuoteInternalDetail, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOfferQuoteInternalDetailResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOfferQuoteInternalDetailResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOfferQuoteInternalDetailResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOfferQuoteInternalDetailResponse")
	}
	return
}

// getOfferQuoteInternalDetail implements the OCIOperation interface (enables retrying operations)
func (client OfferQuoteClient) getOfferQuoteInternalDetail(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/offerQuotes/{offerQuoteId}/internalDetails", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOfferQuoteInternalDetailResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "offerQuote", "GetOfferQuoteInternalDetail")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OfferQuote", "GetOfferQuoteInternalDetail", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOfferQuotes Returns a list of offer quotes. Requires either a reseller compartment ID or an ISV compartment ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplaceprivateoffer/ListOfferQuotes.go.html to see an example of how to use ListOfferQuotes API.
// A default retry strategy applies to this operation ListOfferQuotes()
func (client OfferQuoteClient) ListOfferQuotes(ctx context.Context, request ListOfferQuotesRequest) (response ListOfferQuotesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOfferQuotes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOfferQuotesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOfferQuotesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOfferQuotesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOfferQuotesResponse")
	}
	return
}

// listOfferQuotes implements the OCIOperation interface (enables retrying operations)
func (client OfferQuoteClient) listOfferQuotes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/offerQuotes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOfferQuotesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "offerQuote", "ListOfferQuotes")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OfferQuote", "ListOfferQuotes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RespondToOfferQuote ISV responds to an offer quote for reseller to review.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplaceprivateoffer/RespondToOfferQuote.go.html to see an example of how to use RespondToOfferQuote API.
// A default retry strategy applies to this operation RespondToOfferQuote()
func (client OfferQuoteClient) RespondToOfferQuote(ctx context.Context, request RespondToOfferQuoteRequest) (response RespondToOfferQuoteResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.respondToOfferQuote, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RespondToOfferQuoteResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RespondToOfferQuoteResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RespondToOfferQuoteResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RespondToOfferQuoteResponse")
	}
	return
}

// respondToOfferQuote implements the OCIOperation interface (enables retrying operations)
func (client OfferQuoteClient) respondToOfferQuote(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/offerQuotes/{offerQuoteId}/actions/respondToOfferQuote", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RespondToOfferQuoteResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "offerQuote", "RespondToOfferQuote")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OfferQuote", "RespondToOfferQuote", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SendOfferQuote Sends an offer quote to be reviewed and updated by the ISV. Validation will be run on the offer quote first to verify it is valid and contains all required fields.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplaceprivateoffer/SendOfferQuote.go.html to see an example of how to use SendOfferQuote API.
// A default retry strategy applies to this operation SendOfferQuote()
func (client OfferQuoteClient) SendOfferQuote(ctx context.Context, request SendOfferQuoteRequest) (response SendOfferQuoteResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.sendOfferQuote, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SendOfferQuoteResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SendOfferQuoteResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SendOfferQuoteResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SendOfferQuoteResponse")
	}
	return
}

// sendOfferQuote implements the OCIOperation interface (enables retrying operations)
func (client OfferQuoteClient) sendOfferQuote(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/offerQuotes/{offerQuoteId}/actions/sendOfferQuote", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SendOfferQuoteResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "offerQuote", "SendOfferQuote")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OfferQuote", "SendOfferQuote", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOfferQuote Updates the offer quote.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplaceprivateoffer/UpdateOfferQuote.go.html to see an example of how to use UpdateOfferQuote API.
// A default retry strategy applies to this operation UpdateOfferQuote()
func (client OfferQuoteClient) UpdateOfferQuote(ctx context.Context, request UpdateOfferQuoteRequest) (response UpdateOfferQuoteResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOfferQuote, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOfferQuoteResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOfferQuoteResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOfferQuoteResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOfferQuoteResponse")
	}
	return
}

// updateOfferQuote implements the OCIOperation interface (enables retrying operations)
func (client OfferQuoteClient) updateOfferQuote(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/offerQuotes/{offerQuoteId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOfferQuoteResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "offerQuote", "UpdateOfferQuote")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OfferQuote", "UpdateOfferQuote", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// WithdrawOfferQuote Withdraws an offer quote and transitions to previous state. Offer quotes can only be withdrawn by reseller before the ISV response and withdrawn by the ISV before reseller acceptance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplaceprivateoffer/WithdrawOfferQuote.go.html to see an example of how to use WithdrawOfferQuote API.
// A default retry strategy applies to this operation WithdrawOfferQuote()
func (client OfferQuoteClient) WithdrawOfferQuote(ctx context.Context, request WithdrawOfferQuoteRequest) (response WithdrawOfferQuoteResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.withdrawOfferQuote, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = WithdrawOfferQuoteResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = WithdrawOfferQuoteResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(WithdrawOfferQuoteResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into WithdrawOfferQuoteResponse")
	}
	return
}

// withdrawOfferQuote implements the OCIOperation interface (enables retrying operations)
func (client OfferQuoteClient) withdrawOfferQuote(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/offerQuotes/{offerQuoteId}/actions/withdrawOfferQuote", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response WithdrawOfferQuoteResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "offerQuote", "WithdrawOfferQuote")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OfferQuote", "WithdrawOfferQuote", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
