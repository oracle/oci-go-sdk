// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// MarketplaceAdministratorClient a client for MarketplaceAdministrator
type MarketplaceAdministratorClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMarketplaceAdministratorClientWithConfigurationProvider Creates a new default MarketplaceAdministrator client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMarketplaceAdministratorClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MarketplaceAdministratorClient, err error) {
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
	return newMarketplaceAdministratorClientFromBaseClient(baseClient, provider)
}

// NewMarketplaceAdministratorClientWithOboToken Creates a new default MarketplaceAdministrator client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewMarketplaceAdministratorClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client MarketplaceAdministratorClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newMarketplaceAdministratorClientFromBaseClient(baseClient, configProvider)
}

func newMarketplaceAdministratorClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client MarketplaceAdministratorClient, err error) {
	// MarketplaceAdministrator service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("MarketplaceAdministrator"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = MarketplaceAdministratorClient{BaseClient: baseClient}
	client.BasePath = "20241201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *MarketplaceAdministratorClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("marketplacepublisher", "https://marketplace-publisher.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *MarketplaceAdministratorClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *MarketplaceAdministratorClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeListingRevisionStatusToNew Change the Listing Revision status to New
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ChangeListingRevisionStatusToNew.go.html to see an example of how to use ChangeListingRevisionStatusToNew API.
// A default retry strategy applies to this operation ChangeListingRevisionStatusToNew()
func (client MarketplaceAdministratorClient) ChangeListingRevisionStatusToNew(ctx context.Context, request ChangeListingRevisionStatusToNewRequest) (response ChangeListingRevisionStatusToNewResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeListingRevisionStatusToNew, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeListingRevisionStatusToNewResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeListingRevisionStatusToNewResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeListingRevisionStatusToNewResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeListingRevisionStatusToNewResponse")
	}
	return
}

// changeListingRevisionStatusToNew implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) changeListingRevisionStatusToNew(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/adminListingRevisions/{listingRevisionId}/status/new", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeListingRevisionStatusToNewResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "ChangeListingRevisionStatusToNew")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminListingRevision/ChangeListingRevisionStatusToNew"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "ChangeListingRevisionStatusToNew", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &adminlistingrevision{})
	return response, err
}

// CreateMarket Creates a new Market.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/CreateMarket.go.html to see an example of how to use CreateMarket API.
// A default retry strategy applies to this operation CreateMarket()
func (client MarketplaceAdministratorClient) CreateMarket(ctx context.Context, request CreateMarketRequest) (response CreateMarketResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMarket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMarketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMarketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMarketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMarketResponse")
	}
	return
}

// createMarket implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) createMarket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/adminMarkets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMarketResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "CreateMarket")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/Market/CreateMarket"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "CreateMarket", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAdminArtifact Gets the specified artifact's information.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetAdminArtifact.go.html to see an example of how to use GetAdminArtifact API.
// A default retry strategy applies to this operation GetAdminArtifact()
func (client MarketplaceAdministratorClient) GetAdminArtifact(ctx context.Context, request GetAdminArtifactRequest) (response GetAdminArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAdminArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAdminArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAdminArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAdminArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAdminArtifactResponse")
	}
	return
}

// getAdminArtifact implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) getAdminArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/adminArtifacts/{artifactId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAdminArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "GetAdminArtifact")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminArtifact/GetAdminArtifact"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "GetAdminArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &adminartifact{})
	return response, err
}

// GetAdminListing Returns details of the Listing.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetAdminListing.go.html to see an example of how to use GetAdminListing API.
// A default retry strategy applies to this operation GetAdminListing()
func (client MarketplaceAdministratorClient) GetAdminListing(ctx context.Context, request GetAdminListingRequest) (response GetAdminListingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAdminListing, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAdminListingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAdminListingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAdminListingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAdminListingResponse")
	}
	return
}

// getAdminListing implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) getAdminListing(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/adminListings/{listingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAdminListingResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "GetAdminListing")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminListing/GetAdminListing"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "GetAdminListing", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAdminListingRevision Returns details of the Listing revision.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetAdminListingRevision.go.html to see an example of how to use GetAdminListingRevision API.
// A default retry strategy applies to this operation GetAdminListingRevision()
func (client MarketplaceAdministratorClient) GetAdminListingRevision(ctx context.Context, request GetAdminListingRevisionRequest) (response GetAdminListingRevisionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAdminListingRevision, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAdminListingRevisionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAdminListingRevisionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAdminListingRevisionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAdminListingRevisionResponse")
	}
	return
}

// getAdminListingRevision implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) getAdminListingRevision(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/adminListingRevisions/{listingRevisionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAdminListingRevisionResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "GetAdminListingRevision")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminListingRevision/GetAdminListingRevision"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "GetAdminListingRevision", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &adminlistingrevision{})
	return response, err
}

// GetAdminListingRevisionAttachment Get the details of the specified listing revision attachment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetAdminListingRevisionAttachment.go.html to see an example of how to use GetAdminListingRevisionAttachment API.
// A default retry strategy applies to this operation GetAdminListingRevisionAttachment()
func (client MarketplaceAdministratorClient) GetAdminListingRevisionAttachment(ctx context.Context, request GetAdminListingRevisionAttachmentRequest) (response GetAdminListingRevisionAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAdminListingRevisionAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAdminListingRevisionAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAdminListingRevisionAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAdminListingRevisionAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAdminListingRevisionAttachmentResponse")
	}
	return
}

// getAdminListingRevisionAttachment implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) getAdminListingRevisionAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/adminListingRevisionAttachments/{listingRevisionAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAdminListingRevisionAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "GetAdminListingRevisionAttachment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminListingRevisionAttachment/GetAdminListingRevisionAttachment"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "GetAdminListingRevisionAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &adminlistingrevisionattachment{})
	return response, err
}

// GetAdminListingRevisionPackage Get the details of the specified version of a package.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetAdminListingRevisionPackage.go.html to see an example of how to use GetAdminListingRevisionPackage API.
// A default retry strategy applies to this operation GetAdminListingRevisionPackage()
func (client MarketplaceAdministratorClient) GetAdminListingRevisionPackage(ctx context.Context, request GetAdminListingRevisionPackageRequest) (response GetAdminListingRevisionPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAdminListingRevisionPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAdminListingRevisionPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAdminListingRevisionPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAdminListingRevisionPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAdminListingRevisionPackageResponse")
	}
	return
}

// getAdminListingRevisionPackage implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) getAdminListingRevisionPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/adminListingRevisionPackages/{listingRevisionPackageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAdminListingRevisionPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "GetAdminListingRevisionPackage")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminListingRevisionPackage/GetAdminListingRevisionPackage"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "GetAdminListingRevisionPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &adminlistingrevisionpackage{})
	return response, err
}

// GetAdminTerm Gets a Term by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetAdminTerm.go.html to see an example of how to use GetAdminTerm API.
// A default retry strategy applies to this operation GetAdminTerm()
func (client MarketplaceAdministratorClient) GetAdminTerm(ctx context.Context, request GetAdminTermRequest) (response GetAdminTermResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAdminTerm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAdminTermResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAdminTermResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAdminTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAdminTermResponse")
	}
	return
}

// getAdminTerm implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) getAdminTerm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/adminTerms/{termId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAdminTermResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "GetAdminTerm")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminTerm/GetAdminTerm"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "GetAdminTerm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAdminTermVersion Gets a Term Version by the identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetAdminTermVersion.go.html to see an example of how to use GetAdminTermVersion API.
// A default retry strategy applies to this operation GetAdminTermVersion()
func (client MarketplaceAdministratorClient) GetAdminTermVersion(ctx context.Context, request GetAdminTermVersionRequest) (response GetAdminTermVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAdminTermVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAdminTermVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAdminTermVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAdminTermVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAdminTermVersionResponse")
	}
	return
}

// getAdminTermVersion implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) getAdminTermVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/adminTermVersions/{termVersionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAdminTermVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "GetAdminTermVersion")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminTermVersion/GetAdminTermVersion"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "GetAdminTermVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAdminArtifacts Lists the artifacts in your compartment. You must specify your compartment's OCID as the value for
// the compartment ID.
// For information about OCIDs, see Resource Identifiers (Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListAdminArtifacts.go.html to see an example of how to use ListAdminArtifacts API.
// A default retry strategy applies to this operation ListAdminArtifacts()
func (client MarketplaceAdministratorClient) ListAdminArtifacts(ctx context.Context, request ListAdminArtifactsRequest) (response ListAdminArtifactsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAdminArtifacts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAdminArtifactsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAdminArtifactsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAdminArtifactsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAdminArtifactsResponse")
	}
	return
}

// listAdminArtifacts implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) listAdminArtifacts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/adminArtifacts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAdminArtifactsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "ListAdminArtifacts")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminArtifactCollection/ListAdminArtifacts"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "ListAdminArtifacts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAdminListingRevisionAttachments Gets the list of attachments for a listing revision
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListAdminListingRevisionAttachments.go.html to see an example of how to use ListAdminListingRevisionAttachments API.
// A default retry strategy applies to this operation ListAdminListingRevisionAttachments()
func (client MarketplaceAdministratorClient) ListAdminListingRevisionAttachments(ctx context.Context, request ListAdminListingRevisionAttachmentsRequest) (response ListAdminListingRevisionAttachmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAdminListingRevisionAttachments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAdminListingRevisionAttachmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAdminListingRevisionAttachmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAdminListingRevisionAttachmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAdminListingRevisionAttachmentsResponse")
	}
	return
}

// listAdminListingRevisionAttachments implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) listAdminListingRevisionAttachments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/adminListingRevisionAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAdminListingRevisionAttachmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "ListAdminListingRevisionAttachments")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminListingRevisionAttachmentCollection/ListAdminListingRevisionAttachments"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "ListAdminListingRevisionAttachments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAdminListingRevisionPackages Gets the list of packages for a listing revision.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListAdminListingRevisionPackages.go.html to see an example of how to use ListAdminListingRevisionPackages API.
// A default retry strategy applies to this operation ListAdminListingRevisionPackages()
func (client MarketplaceAdministratorClient) ListAdminListingRevisionPackages(ctx context.Context, request ListAdminListingRevisionPackagesRequest) (response ListAdminListingRevisionPackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAdminListingRevisionPackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAdminListingRevisionPackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAdminListingRevisionPackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAdminListingRevisionPackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAdminListingRevisionPackagesResponse")
	}
	return
}

// listAdminListingRevisionPackages implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) listAdminListingRevisionPackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/adminListingRevisionPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAdminListingRevisionPackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "ListAdminListingRevisionPackages")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminListingRevisionPackageCollection/ListAdminListingRevisionPackages"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "ListAdminListingRevisionPackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAdminListingRevisions Returns a list of all the Listing revisions.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListAdminListingRevisions.go.html to see an example of how to use ListAdminListingRevisions API.
// A default retry strategy applies to this operation ListAdminListingRevisions()
func (client MarketplaceAdministratorClient) ListAdminListingRevisions(ctx context.Context, request ListAdminListingRevisionsRequest) (response ListAdminListingRevisionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAdminListingRevisions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAdminListingRevisionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAdminListingRevisionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAdminListingRevisionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAdminListingRevisionsResponse")
	}
	return
}

// listAdminListingRevisions implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) listAdminListingRevisions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/adminListingRevisions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAdminListingRevisionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "ListAdminListingRevisions")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminListingRevisionCollection/ListAdminListingRevisions"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "ListAdminListingRevisions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAdminPublisherSkus Retrieve a list of publisher SKUs
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListAdminPublisherSkus.go.html to see an example of how to use ListAdminPublisherSkus API.
// A default retry strategy applies to this operation ListAdminPublisherSkus()
func (client MarketplaceAdministratorClient) ListAdminPublisherSkus(ctx context.Context, request ListAdminPublisherSkusRequest) (response ListAdminPublisherSkusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAdminPublisherSkus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAdminPublisherSkusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAdminPublisherSkusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAdminPublisherSkusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAdminPublisherSkusResponse")
	}
	return
}

// listAdminPublisherSkus implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) listAdminPublisherSkus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/adminPublisherSkus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAdminPublisherSkusResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "ListAdminPublisherSkus")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminPublisherSkuCollection/ListAdminPublisherSkus"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "ListAdminPublisherSkus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAdminTermVersions Returns a list of the publisher term versions.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListAdminTermVersions.go.html to see an example of how to use ListAdminTermVersions API.
// A default retry strategy applies to this operation ListAdminTermVersions()
func (client MarketplaceAdministratorClient) ListAdminTermVersions(ctx context.Context, request ListAdminTermVersionsRequest) (response ListAdminTermVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAdminTermVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAdminTermVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAdminTermVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAdminTermVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAdminTermVersionsResponse")
	}
	return
}

// listAdminTermVersions implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) listAdminTermVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/adminTermVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAdminTermVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "ListAdminTermVersions")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminTermVersionCollection/ListAdminTermVersions"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "ListAdminTermVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAdminTerms Returns a list of the publisher terms.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListAdminTerms.go.html to see an example of how to use ListAdminTerms API.
// A default retry strategy applies to this operation ListAdminTerms()
func (client MarketplaceAdministratorClient) ListAdminTerms(ctx context.Context, request ListAdminTermsRequest) (response ListAdminTermsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAdminTerms, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAdminTermsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAdminTermsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAdminTermsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAdminTermsResponse")
	}
	return
}

// listAdminTerms implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) listAdminTerms(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/adminTerms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAdminTermsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "ListAdminTerms")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminTermCollection/ListAdminTerms"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "ListAdminTerms", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAdminWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/ListAdminWorkRequests.go.html to see an example of how to use ListAdminWorkRequests API.
// A default retry strategy applies to this operation ListAdminWorkRequests()
func (client MarketplaceAdministratorClient) ListAdminWorkRequests(ctx context.Context, request ListAdminWorkRequestsRequest) (response ListAdminWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAdminWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAdminWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAdminWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAdminWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAdminWorkRequestsResponse")
	}
	return
}

// listAdminWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceAdministratorClient) listAdminWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/adminWorkRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAdminWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "marketplaceAdministrator", "ListAdminWorkRequests")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminWorkRequestSummary/ListAdminWorkRequests"
		err = common.PostProcessServiceError(err, "MarketplaceAdministrator", "ListAdminWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
