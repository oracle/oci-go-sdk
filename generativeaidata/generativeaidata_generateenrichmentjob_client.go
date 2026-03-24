// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Nl2sql API
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

// GenerateEnrichmentJobClient a client for GenerateEnrichmentJob
type GenerateEnrichmentJobClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewGenerateEnrichmentJobClientWithConfigurationProvider Creates a new default GenerateEnrichmentJob client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewGenerateEnrichmentJobClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client GenerateEnrichmentJobClient, err error) {
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
	return newGenerateEnrichmentJobClientFromBaseClient(baseClient, provider)
}

// NewGenerateEnrichmentJobClientWithOboToken Creates a new default GenerateEnrichmentJob client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewGenerateEnrichmentJobClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client GenerateEnrichmentJobClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newGenerateEnrichmentJobClientFromBaseClient(baseClient, configProvider)
}

func newGenerateEnrichmentJobClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client GenerateEnrichmentJobClient, err error) {
	// GenerateEnrichmentJob service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("GenerateEnrichmentJob"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = GenerateEnrichmentJobClient{BaseClient: baseClient}
	client.BasePath = "20260325"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *GenerateEnrichmentJobClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("generativeaidata", "https://inference.generativeai.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *GenerateEnrichmentJobClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *GenerateEnrichmentJobClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GenerateEnrichmentJob Creates a new asynchronous EnrichmentJob.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaidata/GenerateEnrichmentJob.go.html to see an example of how to use GenerateEnrichmentJob API.
// A default retry strategy applies to this operation GenerateEnrichmentJob()
func (client GenerateEnrichmentJobClient) GenerateEnrichmentJob(ctx context.Context, request GenerateEnrichmentJobRequest) (response GenerateEnrichmentJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateEnrichmentJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateEnrichmentJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateEnrichmentJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateEnrichmentJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateEnrichmentJobResponse")
	}
	return
}

// generateEnrichmentJob implements the OCIOperation interface (enables retrying operations)
func (client GenerateEnrichmentJobClient) generateEnrichmentJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20260325/semanticStores/{semanticStoreId}/actions/enrich", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateEnrichmentJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerateEnrichmentJob", "GenerateEnrichmentJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
