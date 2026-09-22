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

// AdministratorPublisherClient a client for AdministratorPublisher
type AdministratorPublisherClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAdministratorPublisherClientWithConfigurationProvider Creates a new default AdministratorPublisher client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAdministratorPublisherClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AdministratorPublisherClient, err error) {
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
	return newAdministratorPublisherClientFromBaseClient(baseClient, provider)
}

// NewAdministratorPublisherClientWithOboToken Creates a new default AdministratorPublisher client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewAdministratorPublisherClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client AdministratorPublisherClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newAdministratorPublisherClientFromBaseClient(baseClient, configProvider)
}

func newAdministratorPublisherClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client AdministratorPublisherClient, err error) {
	// AdministratorPublisher service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("AdministratorPublisher"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = AdministratorPublisherClient{BaseClient: baseClient}
	client.BasePath = "20241201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AdministratorPublisherClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("marketplacepublisher", "https://marketplace-publisher.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AdministratorPublisherClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *AdministratorPublisherClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetAdminPublisher Gets a Publisher by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplacepublisher/GetAdminPublisher.go.html to see an example of how to use GetAdminPublisher API.
// A default retry strategy applies to this operation GetAdminPublisher()
func (client AdministratorPublisherClient) GetAdminPublisher(ctx context.Context, request GetAdminPublisherRequest) (response GetAdminPublisherResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAdminPublisher, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAdminPublisherResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAdminPublisherResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAdminPublisherResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAdminPublisherResponse")
	}
	return
}

// getAdminPublisher implements the OCIOperation interface (enables retrying operations)
func (client AdministratorPublisherClient) getAdminPublisher(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/adminPublishers/{publisherId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAdminPublisherResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "administratorPublisher", "GetAdminPublisher")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/publisher/20241201/AdminPublisher/GetAdminPublisher"
		err = common.PostProcessServiceError(err, "AdministratorPublisher", "GetAdminPublisher", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
