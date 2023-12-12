// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePublisherService API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplacepublisher

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// MarketplacePublisherClient a client for MarketplacePublisher
type MarketplacePublisherClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMarketplacePublisherClientWithConfigurationProvider Creates a new default MarketplacePublisher client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMarketplacePublisherClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MarketplacePublisherClient, err error) {
	if enabled := common.CheckForEnabledServices("marketplacepublisher"); !enabled {
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
	return newMarketplacePublisherClientFromBaseClient(baseClient, provider)
}

// NewMarketplacePublisherClientWithOboToken Creates a new default MarketplacePublisher client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewMarketplacePublisherClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client MarketplacePublisherClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newMarketplacePublisherClientFromBaseClient(baseClient, configProvider)
}

func newMarketplacePublisherClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client MarketplacePublisherClient, err error) {
	// MarketplacePublisher service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("MarketplacePublisher"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = MarketplacePublisherClient{BaseClient: baseClient}
	client.BasePath = "20220901"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *MarketplacePublisherClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("marketplacepublisher", "https://marketplace-publisher.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *MarketplacePublisherClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *MarketplacePublisherClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ActivateTermVersion Mark the Term Version identified by the id as active
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ActivateTermVersion.go.html to see an example of how to use ActivateTermVersion API.
// A default retry strategy applies to this operation ActivateTermVersion()
func (client MarketplacePublisherClient) ActivateTermVersion(ctx context.Context, request ActivateTermVersionRequest) (response ActivateTermVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.activateTermVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateTermVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateTermVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateTermVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateTermVersionResponse")
	}
	return
}

// activateTermVersion implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) activateTermVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/termVersions/{termVersionId}/actions/Activate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ActivateTermVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/TermVersion/ActivateTermVersion"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ActivateTermVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelWorkRequest Cancels the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
// A default retry strategy applies to this operation CancelWorkRequest()
func (client MarketplacePublisherClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelWorkRequestResponse")
	}
	return
}

// cancelWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CascadingDeleteListing Cascade delete the listing and its subresources.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/CascadingDeleteListing.go.html to see an example of how to use CascadingDeleteListing API.
// A default retry strategy applies to this operation CascadingDeleteListing()
func (client MarketplacePublisherClient) CascadingDeleteListing(ctx context.Context, request CascadingDeleteListingRequest) (response CascadingDeleteListingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cascadingDeleteListing, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CascadingDeleteListingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CascadingDeleteListingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CascadingDeleteListingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CascadingDeleteListingResponse")
	}
	return
}

// cascadingDeleteListing implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) cascadingDeleteListing(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listings/{listingId}/actions/cascadingDelete", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CascadingDeleteListingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Listing/CascadingDeleteListing"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "CascadingDeleteListing", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CascadingDeleteListingRevision Cascade delete listing revision and its subresources.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/CascadingDeleteListingRevision.go.html to see an example of how to use CascadingDeleteListingRevision API.
// A default retry strategy applies to this operation CascadingDeleteListingRevision()
func (client MarketplacePublisherClient) CascadingDeleteListingRevision(ctx context.Context, request CascadingDeleteListingRevisionRequest) (response CascadingDeleteListingRevisionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cascadingDeleteListingRevision, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CascadingDeleteListingRevisionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CascadingDeleteListingRevisionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CascadingDeleteListingRevisionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CascadingDeleteListingRevisionResponse")
	}
	return
}

// cascadingDeleteListingRevision implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) cascadingDeleteListingRevision(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listingRevisions/{listingRevisionId}/actions/cascadingDelete", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CascadingDeleteListingRevisionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevision/CascadingDeleteListingRevision"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "CascadingDeleteListingRevision", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeArtifactCompartment Moves the specified artifact to the specified compartment within the same tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ChangeArtifactCompartment.go.html to see an example of how to use ChangeArtifactCompartment API.
// A default retry strategy applies to this operation ChangeArtifactCompartment()
func (client MarketplacePublisherClient) ChangeArtifactCompartment(ctx context.Context, request ChangeArtifactCompartmentRequest) (response ChangeArtifactCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeArtifactCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeArtifactCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeArtifactCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeArtifactCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeArtifactCompartmentResponse")
	}
	return
}

// changeArtifactCompartment implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) changeArtifactCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/artifacts/{artifactId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeArtifactCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Artifact/ChangeArtifactCompartment"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ChangeArtifactCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeListingCompartment Moves a listing from one compartment to another
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ChangeListingCompartment.go.html to see an example of how to use ChangeListingCompartment API.
// A default retry strategy applies to this operation ChangeListingCompartment()
func (client MarketplacePublisherClient) ChangeListingCompartment(ctx context.Context, request ChangeListingCompartmentRequest) (response ChangeListingCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeListingCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeListingCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeListingCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeListingCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeListingCompartmentResponse")
	}
	return
}

// changeListingCompartment implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) changeListingCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listings/{listingId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeListingCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Listing/ChangeListingCompartment"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ChangeListingCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeListingRevisionToNewStatus Updates the Listing Revision to New status
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ChangeListingRevisionToNewStatus.go.html to see an example of how to use ChangeListingRevisionToNewStatus API.
// A default retry strategy applies to this operation ChangeListingRevisionToNewStatus()
func (client MarketplacePublisherClient) ChangeListingRevisionToNewStatus(ctx context.Context, request ChangeListingRevisionToNewStatusRequest) (response ChangeListingRevisionToNewStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeListingRevisionToNewStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeListingRevisionToNewStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeListingRevisionToNewStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeListingRevisionToNewStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeListingRevisionToNewStatusResponse")
	}
	return
}

// changeListingRevisionToNewStatus implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) changeListingRevisionToNewStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listingRevisions/{listingRevisionId}/actions/changeToNewStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeListingRevisionToNewStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevision/ChangeListingRevisionToNewStatus"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ChangeListingRevisionToNewStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeTermCompartment Moves a term from one compartment to another
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ChangeTermCompartment.go.html to see an example of how to use ChangeTermCompartment API.
// A default retry strategy applies to this operation ChangeTermCompartment()
func (client MarketplacePublisherClient) ChangeTermCompartment(ctx context.Context, request ChangeTermCompartmentRequest) (response ChangeTermCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeTermCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeTermCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeTermCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeTermCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeTermCompartmentResponse")
	}
	return
}

// changeTermCompartment implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) changeTermCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/terms/{termId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeTermCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Term/ChangeTermCompartment"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ChangeTermCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CloneListingRevision Clone the published/withdrawn Listing Revision identified by the identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/CloneListingRevision.go.html to see an example of how to use CloneListingRevision API.
// A default retry strategy applies to this operation CloneListingRevision()
func (client MarketplacePublisherClient) CloneListingRevision(ctx context.Context, request CloneListingRevisionRequest) (response CloneListingRevisionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cloneListingRevision, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CloneListingRevisionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CloneListingRevisionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CloneListingRevisionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CloneListingRevisionResponse")
	}
	return
}

// cloneListingRevision implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) cloneListingRevision(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listingRevisions/{listingRevisionId}/actions/clone", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CloneListingRevisionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevision/CloneListingRevision"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "CloneListingRevision", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateArtifact Creates a new artifact in your compartment.
// You must specify your compartment ID in the request object.
// You must also specify a *name* for the artifact(although it can be an empty string). It does not
// have to be unique, and you can change it anytime with UpdateArtifact..
// You must also specify a *artifactType* for the artifact. Allowed values are CONTAINER_IMAGE and HELM_CHART
// You must also provide the container or helm chart registry details for the corresponding images.
// Oracle container registry details (Registry/Concepts/registryoverview.htm).
// After you send your request, the new object's `status` will temporarily be IN_PROGRESS and `lifecycleState` will be CREATING.
// Before using the object, first make sure its `lifecycleState` has changed to ACTIVE and the status has changed to ‘AVAILABLE’ for the new Artifact.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/CreateArtifact.go.html to see an example of how to use CreateArtifact API.
// A default retry strategy applies to this operation CreateArtifact()
func (client MarketplacePublisherClient) CreateArtifact(ctx context.Context, request CreateArtifactRequest) (response CreateArtifactResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateArtifactResponse")
	}
	return
}

// createArtifact implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) createArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/artifacts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Artifact/CreateArtifact"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "CreateArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateListing Creates a new listing in your compartment.
// You must specify your compartment ID in the request object.
// You must also specify a *name* for the listing and cannot be updated later.
// You must also specify a *packageType* for the listing. Allowed values are CONTAINER_IMAGE and HELM_CHART
// After you send your request, the new object's `lifecycleState` will be CREATING.
// Before using the object, first make sure its `lifecycleState` has changed to ACTIVE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/CreateListing.go.html to see an example of how to use CreateListing API.
// A default retry strategy applies to this operation CreateListing()
func (client MarketplacePublisherClient) CreateListing(ctx context.Context, request CreateListingRequest) (response CreateListingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createListing, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateListingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateListingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateListingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateListingResponse")
	}
	return
}

// createListing implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) createListing(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateListingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Listing/CreateListing"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "CreateListing", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateListingRevision Creates a new Listing Revision.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/CreateListingRevision.go.html to see an example of how to use CreateListingRevision API.
// A default retry strategy applies to this operation CreateListingRevision()
func (client MarketplacePublisherClient) CreateListingRevision(ctx context.Context, request CreateListingRevisionRequest) (response CreateListingRevisionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createListingRevision, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateListingRevisionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateListingRevisionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateListingRevisionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateListingRevisionResponse")
	}
	return
}

// createListingRevision implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) createListingRevision(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listingRevisions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateListingRevisionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevision/CreateListingRevision"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "CreateListingRevision", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateListingRevisionAttachment Creates a new listing revision attachment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/CreateListingRevisionAttachment.go.html to see an example of how to use CreateListingRevisionAttachment API.
// A default retry strategy applies to this operation CreateListingRevisionAttachment()
func (client MarketplacePublisherClient) CreateListingRevisionAttachment(ctx context.Context, request CreateListingRevisionAttachmentRequest) (response CreateListingRevisionAttachmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createListingRevisionAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateListingRevisionAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateListingRevisionAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateListingRevisionAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateListingRevisionAttachmentResponse")
	}
	return
}

// createListingRevisionAttachment implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) createListingRevisionAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listingRevisionAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateListingRevisionAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionAttachment/CreateListingRevisionAttachment"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "CreateListingRevisionAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &listingrevisionattachment{})
	return response, err
}

// CreateListingRevisionNote Creates a new Listing Revision Note.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/CreateListingRevisionNote.go.html to see an example of how to use CreateListingRevisionNote API.
// A default retry strategy applies to this operation CreateListingRevisionNote()
func (client MarketplacePublisherClient) CreateListingRevisionNote(ctx context.Context, request CreateListingRevisionNoteRequest) (response CreateListingRevisionNoteResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createListingRevisionNote, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateListingRevisionNoteResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateListingRevisionNoteResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateListingRevisionNoteResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateListingRevisionNoteResponse")
	}
	return
}

// createListingRevisionNote implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) createListingRevisionNote(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listingRevisionNotes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateListingRevisionNoteResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionNote/CreateListingRevisionNote"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "CreateListingRevisionNote", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateListingRevisionPackage Creates a new Listing Revision Package.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/CreateListingRevisionPackage.go.html to see an example of how to use CreateListingRevisionPackage API.
// A default retry strategy applies to this operation CreateListingRevisionPackage()
func (client MarketplacePublisherClient) CreateListingRevisionPackage(ctx context.Context, request CreateListingRevisionPackageRequest) (response CreateListingRevisionPackageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createListingRevisionPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateListingRevisionPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateListingRevisionPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateListingRevisionPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateListingRevisionPackageResponse")
	}
	return
}

// createListingRevisionPackage implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) createListingRevisionPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listingRevisionPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateListingRevisionPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionPackage/CreateListingRevisionPackage"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "CreateListingRevisionPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &listingrevisionpackage{})
	return response, err
}

// CreateTerm Creates a new Term.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/CreateTerm.go.html to see an example of how to use CreateTerm API.
// A default retry strategy applies to this operation CreateTerm()
func (client MarketplacePublisherClient) CreateTerm(ctx context.Context, request CreateTermRequest) (response CreateTermResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTerm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTermResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTermResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTermResponse")
	}
	return
}

// createTerm implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) createTerm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/terms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTermResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Term/CreateTerm"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "CreateTerm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTermVersion Creates a new Term Version.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/CreateTermVersion.go.html to see an example of how to use CreateTermVersion API.
// A default retry strategy applies to this operation CreateTermVersion()
func (client MarketplacePublisherClient) CreateTermVersion(ctx context.Context, request CreateTermVersionRequest) (response CreateTermVersionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTermVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTermVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTermVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTermVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTermVersionResponse")
	}
	return
}

// createTermVersion implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) createTermVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/termVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTermVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/TermVersion/CreateTermVersion"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "CreateTermVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteArtifact Deletes the specified artifact.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/DeleteArtifact.go.html to see an example of how to use DeleteArtifact API.
// A default retry strategy applies to this operation DeleteArtifact()
func (client MarketplacePublisherClient) DeleteArtifact(ctx context.Context, request DeleteArtifactRequest) (response DeleteArtifactResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteArtifactResponse")
	}
	return
}

// deleteArtifact implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) deleteArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/artifacts/{artifactId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Artifact/DeleteArtifact"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "DeleteArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteListing Deletes a listing by the identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/DeleteListing.go.html to see an example of how to use DeleteListing API.
// A default retry strategy applies to this operation DeleteListing()
func (client MarketplacePublisherClient) DeleteListing(ctx context.Context, request DeleteListingRequest) (response DeleteListingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteListing, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteListingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteListingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteListingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteListingResponse")
	}
	return
}

// deleteListing implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) deleteListing(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/listings/{listingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteListingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Listing/DeleteListing"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "DeleteListing", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteListingRevision Deletes a listing by the identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/DeleteListingRevision.go.html to see an example of how to use DeleteListingRevision API.
// A default retry strategy applies to this operation DeleteListingRevision()
func (client MarketplacePublisherClient) DeleteListingRevision(ctx context.Context, request DeleteListingRevisionRequest) (response DeleteListingRevisionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteListingRevision, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteListingRevisionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteListingRevisionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteListingRevisionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteListingRevisionResponse")
	}
	return
}

// deleteListingRevision implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) deleteListingRevision(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/listingRevisions/{listingRevisionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteListingRevisionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevision/DeleteListingRevision"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "DeleteListingRevision", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteListingRevisionAttachment Deletes a listing revision attachment by the identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/DeleteListingRevisionAttachment.go.html to see an example of how to use DeleteListingRevisionAttachment API.
// A default retry strategy applies to this operation DeleteListingRevisionAttachment()
func (client MarketplacePublisherClient) DeleteListingRevisionAttachment(ctx context.Context, request DeleteListingRevisionAttachmentRequest) (response DeleteListingRevisionAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteListingRevisionAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteListingRevisionAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteListingRevisionAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteListingRevisionAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteListingRevisionAttachmentResponse")
	}
	return
}

// deleteListingRevisionAttachment implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) deleteListingRevisionAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/listingRevisionAttachments/{listingRevisionAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteListingRevisionAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionAttachment/DeleteListingRevisionAttachment"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "DeleteListingRevisionAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteListingRevisionNote Deletes a listing revision note by the identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/DeleteListingRevisionNote.go.html to see an example of how to use DeleteListingRevisionNote API.
// A default retry strategy applies to this operation DeleteListingRevisionNote()
func (client MarketplacePublisherClient) DeleteListingRevisionNote(ctx context.Context, request DeleteListingRevisionNoteRequest) (response DeleteListingRevisionNoteResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteListingRevisionNote, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteListingRevisionNoteResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteListingRevisionNoteResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteListingRevisionNoteResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteListingRevisionNoteResponse")
	}
	return
}

// deleteListingRevisionNote implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) deleteListingRevisionNote(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/listingRevisionNotes/{listingRevisionNoteId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteListingRevisionNoteResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionNote/DeleteListingRevisionNote"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "DeleteListingRevisionNote", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteListingRevisionPackage Deletes a listing revision package by the identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/DeleteListingRevisionPackage.go.html to see an example of how to use DeleteListingRevisionPackage API.
// A default retry strategy applies to this operation DeleteListingRevisionPackage()
func (client MarketplacePublisherClient) DeleteListingRevisionPackage(ctx context.Context, request DeleteListingRevisionPackageRequest) (response DeleteListingRevisionPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteListingRevisionPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteListingRevisionPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteListingRevisionPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteListingRevisionPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteListingRevisionPackageResponse")
	}
	return
}

// deleteListingRevisionPackage implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) deleteListingRevisionPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/listingRevisionPackages/{listingRevisionPackageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteListingRevisionPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionPackage/DeleteListingRevisionPackage"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "DeleteListingRevisionPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTerm Deletes a Term by the identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/DeleteTerm.go.html to see an example of how to use DeleteTerm API.
// A default retry strategy applies to this operation DeleteTerm()
func (client MarketplacePublisherClient) DeleteTerm(ctx context.Context, request DeleteTermRequest) (response DeleteTermResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTerm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTermResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTermResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTermResponse")
	}
	return
}

// deleteTerm implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) deleteTerm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/terms/{termId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTermResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Term/DeleteTerm"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "DeleteTerm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTermVersion Deletes a Term by the identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/DeleteTermVersion.go.html to see an example of how to use DeleteTermVersion API.
// A default retry strategy applies to this operation DeleteTermVersion()
func (client MarketplacePublisherClient) DeleteTermVersion(ctx context.Context, request DeleteTermVersionRequest) (response DeleteTermVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTermVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTermVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTermVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTermVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTermVersionResponse")
	}
	return
}

// deleteTermVersion implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) deleteTermVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/termVersions/{termVersionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTermVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/TermVersion/DeleteTermVersion"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "DeleteTermVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetArtifact Gets the specified artifact's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetArtifact.go.html to see an example of how to use GetArtifact API.
// A default retry strategy applies to this operation GetArtifact()
func (client MarketplacePublisherClient) GetArtifact(ctx context.Context, request GetArtifactRequest) (response GetArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetArtifactResponse")
	}
	return
}

// getArtifact implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) getArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/artifacts/{artifactId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Artifact/GetArtifact"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "GetArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &artifact{})
	return response, err
}

// GetCategory Gets the specified category's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetCategory.go.html to see an example of how to use GetCategory API.
// A default retry strategy applies to this operation GetCategory()
func (client MarketplacePublisherClient) GetCategory(ctx context.Context, request GetCategoryRequest) (response GetCategoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCategory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCategoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCategoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCategoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCategoryResponse")
	}
	return
}

// getCategory implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) getCategory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/categories/{categoryCode}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCategoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Category/GetCategory"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "GetCategory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetListing Gets the details for a listing.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetListing.go.html to see an example of how to use GetListing API.
// A default retry strategy applies to this operation GetListing()
func (client MarketplacePublisherClient) GetListing(ctx context.Context, request GetListingRequest) (response GetListingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getListing, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetListingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetListingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetListingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetListingResponse")
	}
	return
}

// getListing implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) getListing(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listings/{listingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetListingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Listing/GetListing"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "GetListing", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetListingRevision Gets the details for a listing revision.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetListingRevision.go.html to see an example of how to use GetListingRevision API.
// A default retry strategy applies to this operation GetListingRevision()
func (client MarketplacePublisherClient) GetListingRevision(ctx context.Context, request GetListingRevisionRequest) (response GetListingRevisionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getListingRevision, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetListingRevisionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetListingRevisionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetListingRevisionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetListingRevisionResponse")
	}
	return
}

// getListingRevision implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) getListingRevision(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listingRevisions/{listingRevisionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetListingRevisionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevision/GetListingRevision"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "GetListingRevision", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetListingRevisionAttachment Get the details of the specified listing revision attachment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetListingRevisionAttachment.go.html to see an example of how to use GetListingRevisionAttachment API.
// A default retry strategy applies to this operation GetListingRevisionAttachment()
func (client MarketplacePublisherClient) GetListingRevisionAttachment(ctx context.Context, request GetListingRevisionAttachmentRequest) (response GetListingRevisionAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getListingRevisionAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetListingRevisionAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetListingRevisionAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetListingRevisionAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetListingRevisionAttachmentResponse")
	}
	return
}

// getListingRevisionAttachment implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) getListingRevisionAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listingRevisionAttachments/{listingRevisionAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetListingRevisionAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionAttachment/GetListingRevisionAttachment"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "GetListingRevisionAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &listingrevisionattachment{})
	return response, err
}

// GetListingRevisionNote Get note details by the identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetListingRevisionNote.go.html to see an example of how to use GetListingRevisionNote API.
// A default retry strategy applies to this operation GetListingRevisionNote()
func (client MarketplacePublisherClient) GetListingRevisionNote(ctx context.Context, request GetListingRevisionNoteRequest) (response GetListingRevisionNoteResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getListingRevisionNote, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetListingRevisionNoteResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetListingRevisionNoteResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetListingRevisionNoteResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetListingRevisionNoteResponse")
	}
	return
}

// getListingRevisionNote implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) getListingRevisionNote(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listingRevisionNotes/{listingRevisionNoteId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetListingRevisionNoteResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionNote/GetListingRevisionNote"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "GetListingRevisionNote", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetListingRevisionPackage Get the details of the specified version of a package.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetListingRevisionPackage.go.html to see an example of how to use GetListingRevisionPackage API.
// A default retry strategy applies to this operation GetListingRevisionPackage()
func (client MarketplacePublisherClient) GetListingRevisionPackage(ctx context.Context, request GetListingRevisionPackageRequest) (response GetListingRevisionPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getListingRevisionPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetListingRevisionPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetListingRevisionPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetListingRevisionPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetListingRevisionPackageResponse")
	}
	return
}

// getListingRevisionPackage implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) getListingRevisionPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listingRevisionPackages/{listingRevisionPackageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetListingRevisionPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionPackage/GetListingRevisionPackage"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "GetListingRevisionPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &listingrevisionpackage{})
	return response, err
}

// GetMarket Gets the specified market's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetMarket.go.html to see an example of how to use GetMarket API.
// A default retry strategy applies to this operation GetMarket()
func (client MarketplacePublisherClient) GetMarket(ctx context.Context, request GetMarketRequest) (response GetMarketResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMarket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMarketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMarketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMarketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMarketResponse")
	}
	return
}

// getMarket implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) getMarket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/markets/{marketCode}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMarketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Market/GetMarket"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "GetMarket", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetProduct Gets a Product by code identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetProduct.go.html to see an example of how to use GetProduct API.
// A default retry strategy applies to this operation GetProduct()
func (client MarketplacePublisherClient) GetProduct(ctx context.Context, request GetProductRequest) (response GetProductResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getProduct, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetProductResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetProductResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetProductResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetProductResponse")
	}
	return
}

// getProduct implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) getProduct(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/products/{productCode}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetProductResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Product/GetProduct"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "GetProduct", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPublisher Gets a Publisher by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetPublisher.go.html to see an example of how to use GetPublisher API.
// A default retry strategy applies to this operation GetPublisher()
func (client MarketplacePublisherClient) GetPublisher(ctx context.Context, request GetPublisherRequest) (response GetPublisherResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPublisher, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPublisherResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPublisherResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPublisherResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPublisherResponse")
	}
	return
}

// getPublisher implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) getPublisher(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/publishers/{publisherId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPublisherResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Publisher/GetPublisher"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "GetPublisher", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTerm Gets a Term by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetTerm.go.html to see an example of how to use GetTerm API.
// A default retry strategy applies to this operation GetTerm()
func (client MarketplacePublisherClient) GetTerm(ctx context.Context, request GetTermRequest) (response GetTermResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTerm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTermResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTermResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTermResponse")
	}
	return
}

// getTerm implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) getTerm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/terms/{termId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTermResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Term/GetTerm"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "GetTerm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTermVersion Gets a Term Version by the identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetTermVersion.go.html to see an example of how to use GetTermVersion API.
// A default retry strategy applies to this operation GetTermVersion()
func (client MarketplacePublisherClient) GetTermVersion(ctx context.Context, request GetTermVersionRequest) (response GetTermVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTermVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTermVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTermVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTermVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTermVersionResponse")
	}
	return
}

// getTermVersion implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) getTermVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/termVersions/{termVersionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTermVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/TermVersion/GetTermVersion"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "GetTermVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets details of the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client MarketplacePublisherClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListArtifacts Lists the artifacts in your compartment. You must specify your compartment's OCID as the value for
// the compartment ID.
// For information about OCIDs, see Resource Identifiers (Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListArtifacts.go.html to see an example of how to use ListArtifacts API.
// A default retry strategy applies to this operation ListArtifacts()
func (client MarketplacePublisherClient) ListArtifacts(ctx context.Context, request ListArtifactsRequest) (response ListArtifactsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listArtifacts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListArtifactsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListArtifactsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListArtifactsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListArtifactsResponse")
	}
	return
}

// listArtifacts implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) listArtifacts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/artifacts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListArtifactsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ArtifactCollection/ListArtifacts"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ListArtifacts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCategories Lists the categories in your compartment. You must specify your compartment's OCID as the value for
// the compartment ID.
// For information about OCIDs, see Resource Identifiers (Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListCategories.go.html to see an example of how to use ListCategories API.
// A default retry strategy applies to this operation ListCategories()
func (client MarketplacePublisherClient) ListCategories(ctx context.Context, request ListCategoriesRequest) (response ListCategoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCategories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCategoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCategoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCategoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCategoriesResponse")
	}
	return
}

// listCategories implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) listCategories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/categories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCategoriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/CategoryCollection/ListCategories"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ListCategories", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListListingRevisionAttachments Gets the list of attachments for a listing revision in a compartment
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListListingRevisionAttachments.go.html to see an example of how to use ListListingRevisionAttachments API.
// A default retry strategy applies to this operation ListListingRevisionAttachments()
func (client MarketplacePublisherClient) ListListingRevisionAttachments(ctx context.Context, request ListListingRevisionAttachmentsRequest) (response ListListingRevisionAttachmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listListingRevisionAttachments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListListingRevisionAttachmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListListingRevisionAttachmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListListingRevisionAttachmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListListingRevisionAttachmentsResponse")
	}
	return
}

// listListingRevisionAttachments implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) listListingRevisionAttachments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listingRevisionAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListListingRevisionAttachmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionAttachmentCollection/ListListingRevisionAttachments"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ListListingRevisionAttachments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListListingRevisionNotes Gets the list of notes for a listing revision.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListListingRevisionNotes.go.html to see an example of how to use ListListingRevisionNotes API.
// A default retry strategy applies to this operation ListListingRevisionNotes()
func (client MarketplacePublisherClient) ListListingRevisionNotes(ctx context.Context, request ListListingRevisionNotesRequest) (response ListListingRevisionNotesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listListingRevisionNotes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListListingRevisionNotesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListListingRevisionNotesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListListingRevisionNotesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListListingRevisionNotesResponse")
	}
	return
}

// listListingRevisionNotes implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) listListingRevisionNotes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listingRevisionNotes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListListingRevisionNotesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionNoteCollection/ListListingRevisionNotes"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ListListingRevisionNotes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListListingRevisionPackages Gets the list of packages for a listing revision.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListListingRevisionPackages.go.html to see an example of how to use ListListingRevisionPackages API.
// A default retry strategy applies to this operation ListListingRevisionPackages()
func (client MarketplacePublisherClient) ListListingRevisionPackages(ctx context.Context, request ListListingRevisionPackagesRequest) (response ListListingRevisionPackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listListingRevisionPackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListListingRevisionPackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListListingRevisionPackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListListingRevisionPackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListListingRevisionPackagesResponse")
	}
	return
}

// listListingRevisionPackages implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) listListingRevisionPackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listingRevisionPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListListingRevisionPackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionPackageCollection/ListListingRevisionPackages"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ListListingRevisionPackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListListingRevisions Lists the list of listing revisions for a specific listing ID, compartment ID or listing revision status.
// You can specify your compartment's OCID as the value for the compartment ID.
// For information about OCIDs, see Resource Identifiers (Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListListingRevisions.go.html to see an example of how to use ListListingRevisions API.
// A default retry strategy applies to this operation ListListingRevisions()
func (client MarketplacePublisherClient) ListListingRevisions(ctx context.Context, request ListListingRevisionsRequest) (response ListListingRevisionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listListingRevisions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListListingRevisionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListListingRevisionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListListingRevisionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListListingRevisionsResponse")
	}
	return
}

// listListingRevisions implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) listListingRevisions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listingRevisions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListListingRevisionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionCollection/ListListingRevisions"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ListListingRevisions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListListings Lists the listings in your compartment. You must specify your compartment's OCID as the value for
// the compartment ID.
// For information about OCIDs, see Resource Identifiers (Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListListings.go.html to see an example of how to use ListListings API.
// A default retry strategy applies to this operation ListListings()
func (client MarketplacePublisherClient) ListListings(ctx context.Context, request ListListingsRequest) (response ListListingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listListings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListListingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListListingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListListingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListListingsResponse")
	}
	return
}

// listListings implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) listListings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListListingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingCollection/ListListings"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ListListings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMarkets Lists the markets in your compartment. You must specify your compartment's OCID as the value for
// the compartment ID.
// For information about OCIDs, see Resource Identifiers (Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListMarkets.go.html to see an example of how to use ListMarkets API.
// A default retry strategy applies to this operation ListMarkets()
func (client MarketplacePublisherClient) ListMarkets(ctx context.Context, request ListMarketsRequest) (response ListMarketsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMarkets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMarketsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMarketsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMarketsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMarketsResponse")
	}
	return
}

// listMarkets implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) listMarkets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/markets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMarketsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/MarketCollection/ListMarkets"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ListMarkets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProducts Lists the products in your compartment. You must specify your compartment's OCID as the value for
// the compartment ID.
// For information about OCIDs, see Resource Identifiers (Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListProducts.go.html to see an example of how to use ListProducts API.
// A default retry strategy applies to this operation ListProducts()
func (client MarketplacePublisherClient) ListProducts(ctx context.Context, request ListProductsRequest) (response ListProductsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProducts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProductsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProductsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProductsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProductsResponse")
	}
	return
}

// listProducts implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) listProducts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/products", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProductsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ProductCollection/ListProducts"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ListProducts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPublishers Returns a list of publishers.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListPublishers.go.html to see an example of how to use ListPublishers API.
// A default retry strategy applies to this operation ListPublishers()
func (client MarketplacePublisherClient) ListPublishers(ctx context.Context, request ListPublishersRequest) (response ListPublishersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPublishers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPublishersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPublishersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPublishersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPublishersResponse")
	}
	return
}

// listPublishers implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) listPublishers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/publishers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPublishersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/PublisherCollection/ListPublishers"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ListPublishers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTermVersions Returns a list of the publisher term versions.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListTermVersions.go.html to see an example of how to use ListTermVersions API.
// A default retry strategy applies to this operation ListTermVersions()
func (client MarketplacePublisherClient) ListTermVersions(ctx context.Context, request ListTermVersionsRequest) (response ListTermVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTermVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTermVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTermVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTermVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTermVersionsResponse")
	}
	return
}

// listTermVersions implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) listTermVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/termVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTermVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/TermVersionCollection/ListTermVersions"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ListTermVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTerms Returns a list of the publisher terms.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListTerms.go.html to see an example of how to use ListTerms API.
// A default retry strategy applies to this operation ListTerms()
func (client MarketplacePublisherClient) ListTerms(ctx context.Context, request ListTermsRequest) (response ListTermsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTerms, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTermsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTermsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTermsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTermsResponse")
	}
	return
}

// listTerms implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) listTerms(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/terms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTermsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/TermCollection/ListTerms"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ListTerms", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Returns a (paginated) list of errors for the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client MarketplacePublisherClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Returns a (paginated) list of logs for the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client MarketplacePublisherClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client MarketplacePublisherClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MarkListingRevisionPackageAsDefault Mark the Listing Revision Package identified by the id and package version as default
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/MarkListingRevisionPackageAsDefault.go.html to see an example of how to use MarkListingRevisionPackageAsDefault API.
// A default retry strategy applies to this operation MarkListingRevisionPackageAsDefault()
func (client MarketplacePublisherClient) MarkListingRevisionPackageAsDefault(ctx context.Context, request MarkListingRevisionPackageAsDefaultRequest) (response MarkListingRevisionPackageAsDefaultResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.markListingRevisionPackageAsDefault, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MarkListingRevisionPackageAsDefaultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MarkListingRevisionPackageAsDefaultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MarkListingRevisionPackageAsDefaultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MarkListingRevisionPackageAsDefaultResponse")
	}
	return
}

// markListingRevisionPackageAsDefault implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) markListingRevisionPackageAsDefault(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listingRevisionPackages/{listingRevisionPackageId}/actions/markAsDefault", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response MarkListingRevisionPackageAsDefaultResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionPackage/MarkListingRevisionPackageAsDefault"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "MarkListingRevisionPackageAsDefault", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PublishListingRevision Publish the Listing revision identified by Identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/PublishListingRevision.go.html to see an example of how to use PublishListingRevision API.
// A default retry strategy applies to this operation PublishListingRevision()
func (client MarketplacePublisherClient) PublishListingRevision(ctx context.Context, request PublishListingRevisionRequest) (response PublishListingRevisionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.publishListingRevision, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PublishListingRevisionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PublishListingRevisionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PublishListingRevisionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PublishListingRevisionResponse")
	}
	return
}

// publishListingRevision implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) publishListingRevision(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listingRevisions/{listingRevisionId}/actions/publish", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PublishListingRevisionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevision/PublishListingRevision"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "PublishListingRevision", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PublishListingRevisionAsPrivate Updates the Listing Revision to PublishAsPrivate
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/PublishListingRevisionAsPrivate.go.html to see an example of how to use PublishListingRevisionAsPrivate API.
// A default retry strategy applies to this operation PublishListingRevisionAsPrivate()
func (client MarketplacePublisherClient) PublishListingRevisionAsPrivate(ctx context.Context, request PublishListingRevisionAsPrivateRequest) (response PublishListingRevisionAsPrivateResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.publishListingRevisionAsPrivate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PublishListingRevisionAsPrivateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PublishListingRevisionAsPrivateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PublishListingRevisionAsPrivateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PublishListingRevisionAsPrivateResponse")
	}
	return
}

// publishListingRevisionAsPrivate implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) publishListingRevisionAsPrivate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listingRevisions/{listingRevisionId}/actions/publishAsPrivate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PublishListingRevisionAsPrivateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevision/PublishListingRevisionAsPrivate"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "PublishListingRevisionAsPrivate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PublishListingRevisionPackage Updates the Listing Revision Package to publish status
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/PublishListingRevisionPackage.go.html to see an example of how to use PublishListingRevisionPackage API.
// A default retry strategy applies to this operation PublishListingRevisionPackage()
func (client MarketplacePublisherClient) PublishListingRevisionPackage(ctx context.Context, request PublishListingRevisionPackageRequest) (response PublishListingRevisionPackageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.publishListingRevisionPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PublishListingRevisionPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PublishListingRevisionPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PublishListingRevisionPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PublishListingRevisionPackageResponse")
	}
	return
}

// publishListingRevisionPackage implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) publishListingRevisionPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listingRevisionPackages/{listingRevisionPackageId}/actions/publish", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PublishListingRevisionPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionPackage/PublishListingRevisionPackage"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "PublishListingRevisionPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SubmitListingRevisionForReview Update the Listing Revision identified by the id for review
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/SubmitListingRevisionForReview.go.html to see an example of how to use SubmitListingRevisionForReview API.
// A default retry strategy applies to this operation SubmitListingRevisionForReview()
func (client MarketplacePublisherClient) SubmitListingRevisionForReview(ctx context.Context, request SubmitListingRevisionForReviewRequest) (response SubmitListingRevisionForReviewResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.submitListingRevisionForReview, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SubmitListingRevisionForReviewResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SubmitListingRevisionForReviewResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SubmitListingRevisionForReviewResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SubmitListingRevisionForReviewResponse")
	}
	return
}

// submitListingRevisionForReview implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) submitListingRevisionForReview(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listingRevisions/{listingRevisionId}/actions/submitForReview", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SubmitListingRevisionForReviewResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevision/SubmitListingRevisionForReview"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "SubmitListingRevisionForReview", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UnPublishListingRevisionPackage Updates the Listing Revision Package to Unpublish status
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/UnPublishListingRevisionPackage.go.html to see an example of how to use UnPublishListingRevisionPackage API.
// A default retry strategy applies to this operation UnPublishListingRevisionPackage()
func (client MarketplacePublisherClient) UnPublishListingRevisionPackage(ctx context.Context, request UnPublishListingRevisionPackageRequest) (response UnPublishListingRevisionPackageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.unPublishListingRevisionPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UnPublishListingRevisionPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UnPublishListingRevisionPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UnPublishListingRevisionPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UnPublishListingRevisionPackageResponse")
	}
	return
}

// unPublishListingRevisionPackage implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) unPublishListingRevisionPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listingRevisionPackages/{listingRevisionPackageId}/actions/unPublish", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UnPublishListingRevisionPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionPackage/UnPublishListingRevisionPackage"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "UnPublishListingRevisionPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateArtifact Updates the specified artifact identified by the id.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/UpdateArtifact.go.html to see an example of how to use UpdateArtifact API.
// A default retry strategy applies to this operation UpdateArtifact()
func (client MarketplacePublisherClient) UpdateArtifact(ctx context.Context, request UpdateArtifactRequest) (response UpdateArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateArtifactResponse")
	}
	return
}

// updateArtifact implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) updateArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/artifacts/{artifactId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Artifact/UpdateArtifact"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "UpdateArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateListing Updates the specified Listing identified by the id.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/UpdateListing.go.html to see an example of how to use UpdateListing API.
// A default retry strategy applies to this operation UpdateListing()
func (client MarketplacePublisherClient) UpdateListing(ctx context.Context, request UpdateListingRequest) (response UpdateListingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateListing, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateListingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateListingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateListingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateListingResponse")
	}
	return
}

// updateListing implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) updateListing(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/listings/{listingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateListingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Listing/UpdateListing"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "UpdateListing", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateListingRevision Updates the Listing Revision
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/UpdateListingRevision.go.html to see an example of how to use UpdateListingRevision API.
// A default retry strategy applies to this operation UpdateListingRevision()
func (client MarketplacePublisherClient) UpdateListingRevision(ctx context.Context, request UpdateListingRevisionRequest) (response UpdateListingRevisionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateListingRevision, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateListingRevisionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateListingRevisionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateListingRevisionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateListingRevisionResponse")
	}
	return
}

// updateListingRevision implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) updateListingRevision(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/listingRevisions/{listingRevisionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateListingRevisionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevision/UpdateListingRevision"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "UpdateListingRevision", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateListingRevisionAttachment Updates the Listing Revision Attachment
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/UpdateListingRevisionAttachment.go.html to see an example of how to use UpdateListingRevisionAttachment API.
// A default retry strategy applies to this operation UpdateListingRevisionAttachment()
func (client MarketplacePublisherClient) UpdateListingRevisionAttachment(ctx context.Context, request UpdateListingRevisionAttachmentRequest) (response UpdateListingRevisionAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateListingRevisionAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateListingRevisionAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateListingRevisionAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateListingRevisionAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateListingRevisionAttachmentResponse")
	}
	return
}

// updateListingRevisionAttachment implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) updateListingRevisionAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/listingRevisionAttachments/{listingRevisionAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateListingRevisionAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionAttachment/UpdateListingRevisionAttachment"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "UpdateListingRevisionAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &listingrevisionattachment{})
	return response, err
}

// UpdateListingRevisionAttachmentContent Update a file to listing revision attachment
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/UpdateListingRevisionAttachmentContent.go.html to see an example of how to use UpdateListingRevisionAttachmentContent API.
// A default retry strategy applies to this operation UpdateListingRevisionAttachmentContent()
func (client MarketplacePublisherClient) UpdateListingRevisionAttachmentContent(ctx context.Context, request UpdateListingRevisionAttachmentContentRequest) (response UpdateListingRevisionAttachmentContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateListingRevisionAttachmentContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateListingRevisionAttachmentContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateListingRevisionAttachmentContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateListingRevisionAttachmentContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateListingRevisionAttachmentContentResponse")
	}
	return
}

// updateListingRevisionAttachmentContent implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) updateListingRevisionAttachmentContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/listingRevisionAttachments/{listingRevisionAttachmentId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateListingRevisionAttachmentContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionAttachment/UpdateListingRevisionAttachmentContent"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "UpdateListingRevisionAttachmentContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &listingrevisionattachment{})
	return response, err
}

// UpdateListingRevisionIconContent Updates the Listing Revision
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/UpdateListingRevisionIconContent.go.html to see an example of how to use UpdateListingRevisionIconContent API.
// A default retry strategy applies to this operation UpdateListingRevisionIconContent()
func (client MarketplacePublisherClient) UpdateListingRevisionIconContent(ctx context.Context, request UpdateListingRevisionIconContentRequest) (response UpdateListingRevisionIconContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateListingRevisionIconContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateListingRevisionIconContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateListingRevisionIconContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateListingRevisionIconContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateListingRevisionIconContentResponse")
	}
	return
}

// updateListingRevisionIconContent implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) updateListingRevisionIconContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/listingRevisions/{listingRevisionId}/icon/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateListingRevisionIconContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevision/UpdateListingRevisionIconContent"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "UpdateListingRevisionIconContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateListingRevisionPackage Updates the Listing Revision Package.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/UpdateListingRevisionPackage.go.html to see an example of how to use UpdateListingRevisionPackage API.
// A default retry strategy applies to this operation UpdateListingRevisionPackage()
func (client MarketplacePublisherClient) UpdateListingRevisionPackage(ctx context.Context, request UpdateListingRevisionPackageRequest) (response UpdateListingRevisionPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateListingRevisionPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateListingRevisionPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateListingRevisionPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateListingRevisionPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateListingRevisionPackageResponse")
	}
	return
}

// updateListingRevisionPackage implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) updateListingRevisionPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/listingRevisionPackages/{listingRevisionPackageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateListingRevisionPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevisionPackage/UpdateListingRevisionPackage"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "UpdateListingRevisionPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &listingrevisionpackage{})
	return response, err
}

// UpdateTerm Updates the Term
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/UpdateTerm.go.html to see an example of how to use UpdateTerm API.
// A default retry strategy applies to this operation UpdateTerm()
func (client MarketplacePublisherClient) UpdateTerm(ctx context.Context, request UpdateTermRequest) (response UpdateTermResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTerm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTermResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTermResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTermResponse")
	}
	return
}

// updateTerm implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) updateTerm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/terms/{termId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTermResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Term/UpdateTerm"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "UpdateTerm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTermVersion Updates the Term Version
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/UpdateTermVersion.go.html to see an example of how to use UpdateTermVersion API.
// A default retry strategy applies to this operation UpdateTermVersion()
func (client MarketplacePublisherClient) UpdateTermVersion(ctx context.Context, request UpdateTermVersionRequest) (response UpdateTermVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTermVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTermVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTermVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTermVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTermVersionResponse")
	}
	return
}

// updateTermVersion implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) updateTermVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/termVersions/{termVersionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTermVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/TermVersion/UpdateTermVersion"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "UpdateTermVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTermVersionContent Updates the Term Version attachment
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/UpdateTermVersionContent.go.html to see an example of how to use UpdateTermVersionContent API.
// A default retry strategy applies to this operation UpdateTermVersionContent()
func (client MarketplacePublisherClient) UpdateTermVersionContent(ctx context.Context, request UpdateTermVersionContentRequest) (response UpdateTermVersionContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTermVersionContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTermVersionContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTermVersionContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTermVersionContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTermVersionContentResponse")
	}
	return
}

// updateTermVersionContent implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) updateTermVersionContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/termVersions/{termVersionId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTermVersionContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/TermVersion/UpdateTermVersionContent"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "UpdateTermVersionContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateAndPublishArtifact Validate and publish artifact.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ValidateAndPublishArtifact.go.html to see an example of how to use ValidateAndPublishArtifact API.
// A default retry strategy applies to this operation ValidateAndPublishArtifact()
func (client MarketplacePublisherClient) ValidateAndPublishArtifact(ctx context.Context, request ValidateAndPublishArtifactRequest) (response ValidateAndPublishArtifactResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.validateAndPublishArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateAndPublishArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateAndPublishArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateAndPublishArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateAndPublishArtifactResponse")
	}
	return
}

// validateAndPublishArtifact implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) validateAndPublishArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/artifacts/{artifactId}/actions/validateAndPublish", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateAndPublishArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/Artifact/ValidateAndPublishArtifact"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "ValidateAndPublishArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// WithdrawListingRevision Update the Listing Revision identified by the id as Withdraw/UnPublished.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/WithdrawListingRevision.go.html to see an example of how to use WithdrawListingRevision API.
// A default retry strategy applies to this operation WithdrawListingRevision()
func (client MarketplacePublisherClient) WithdrawListingRevision(ctx context.Context, request WithdrawListingRevisionRequest) (response WithdrawListingRevisionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.withdrawListingRevision, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = WithdrawListingRevisionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = WithdrawListingRevisionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(WithdrawListingRevisionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into WithdrawListingRevisionResponse")
	}
	return
}

// withdrawListingRevision implements the OCIOperation interface (enables retrying operations)
func (client MarketplacePublisherClient) withdrawListingRevision(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listingRevisions/{listingRevisionId}/actions/withdraw", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response WithdrawListingRevisionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20220901/ListingRevision/WithdrawListingRevision"
		err = common.PostProcessServiceError(err, "MarketplacePublisher", "WithdrawListingRevision", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
