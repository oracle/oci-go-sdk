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

// CcdOffersClient a client for CcdOffers
type CcdOffersClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewCcdOffersClientWithConfigurationProvider Creates a new default CcdOffers client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewCcdOffersClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client CcdOffersClient, err error) {
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
	return newCcdOffersClientFromBaseClient(baseClient, provider)
}

// NewCcdOffersClientWithOboToken Creates a new default CcdOffers client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewCcdOffersClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client CcdOffersClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newCcdOffersClientFromBaseClient(baseClient, configProvider)
}

func newCcdOffersClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client CcdOffersClient, err error) {
	// CcdOffers service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("CcdOffers"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = CcdOffersClient{BaseClient: baseClient}
	client.BasePath = "20240601"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *CcdOffersClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ccdoffers", "https://ccdOffers.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *CcdOffersClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *CcdOffersClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetCcdOffer Get a particular CCD offer details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ccdoffers/GetCcdOffer.go.html to see an example of how to use GetCcdOffer API.
// A default retry strategy applies to this operation GetCcdOffer()
func (client CcdOffersClient) GetCcdOffer(ctx context.Context, request GetCcdOfferRequest) (response GetCcdOfferResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCcdOffer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCcdOfferResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCcdOfferResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCcdOfferResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCcdOfferResponse")
	}
	return
}

// getCcdOffer implements the OCIOperation interface (enables retrying operations)
func (client CcdOffersClient) getCcdOffer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ccdOffers/{ccdOfferId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCcdOfferResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "ccdOffers", "GetCcdOffer")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "CcdOffers", "GetCcdOffer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCcdOffers Lists the CCD offers for the requested subscription. A FLEX offer is no longer returned after it is either accepted,
// rejected or expired.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ccdoffers/ListCcdOffers.go.html to see an example of how to use ListCcdOffers API.
// A default retry strategy applies to this operation ListCcdOffers()
func (client CcdOffersClient) ListCcdOffers(ctx context.Context, request ListCcdOffersRequest) (response ListCcdOffersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCcdOffers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCcdOffersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCcdOffersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCcdOffersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCcdOffersResponse")
	}
	return
}

// listCcdOffers implements the OCIOperation interface (enables retrying operations)
func (client CcdOffersClient) listCcdOffers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ccdOffers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCcdOffersResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "ccdOffers", "ListCcdOffers")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "CcdOffers", "ListCcdOffers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
