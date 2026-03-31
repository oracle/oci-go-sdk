// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service NL2SQL API
//
// A description of the ReferenceService API. in progress
//

package generativeaidata

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ListEnrichmentJobsClient a client for ListEnrichmentJobs
type ListEnrichmentJobsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewListEnrichmentJobsClientWithConfigurationProvider Creates a new default ListEnrichmentJobs client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewListEnrichmentJobsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ListEnrichmentJobsClient, err error) {
	if enabled := common.CheckForEnabledServices("generativeaidata"); !enabled {
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
	return newListEnrichmentJobsClientFromBaseClient(baseClient, provider)
}

// NewListEnrichmentJobsClientWithOboToken Creates a new default ListEnrichmentJobs client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewListEnrichmentJobsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ListEnrichmentJobsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newListEnrichmentJobsClientFromBaseClient(baseClient, configProvider)
}

func newListEnrichmentJobsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ListEnrichmentJobsClient, err error) {
	// ListEnrichmentJobs service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ListEnrichmentJobs"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ListEnrichmentJobsClient{BaseClient: baseClient}
	client.BasePath = "20260325"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ListEnrichmentJobsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("generativeaidata", "https://inference.generativeai.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ListEnrichmentJobsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ListEnrichmentJobsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListEnrichmentJobs Lists EnrichmentJobs in a semantic store.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaidata/ListEnrichmentJobs.go.html to see an example of how to use ListEnrichmentJobs API.
// A default retry strategy applies to this operation ListEnrichmentJobs()
func (client ListEnrichmentJobsClient) ListEnrichmentJobs(ctx context.Context, request ListEnrichmentJobsRequest) (response ListEnrichmentJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEnrichmentJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEnrichmentJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEnrichmentJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEnrichmentJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEnrichmentJobsResponse")
	}
	return
}

// listEnrichmentJobs implements the OCIOperation interface (enables retrying operations)
func (client ListEnrichmentJobsClient) listEnrichmentJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/semanticStores/{semanticStoreId}/enrichmentJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEnrichmentJobsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ListEnrichmentJobs", "ListEnrichmentJobs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
