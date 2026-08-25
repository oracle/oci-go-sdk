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

// EnrichmentJobClient a client for EnrichmentJob
type EnrichmentJobClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewEnrichmentJobClientWithConfigurationProvider Creates a new default EnrichmentJob client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewEnrichmentJobClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client EnrichmentJobClient, err error) {
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
	return newEnrichmentJobClientFromBaseClient(baseClient, provider)
}

// NewEnrichmentJobClientWithOboToken Creates a new default EnrichmentJob client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewEnrichmentJobClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client EnrichmentJobClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newEnrichmentJobClientFromBaseClient(baseClient, configProvider)
}

func newEnrichmentJobClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client EnrichmentJobClient, err error) {
	// EnrichmentJob service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("EnrichmentJob"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = EnrichmentJobClient{BaseClient: baseClient}
	client.BasePath = "20260325"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *EnrichmentJobClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("generativeaidata", "https://inference.generativeai.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *EnrichmentJobClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *EnrichmentJobClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelEnrichmentJob Cancels an EnrichmentJob resource by identifier.
// The returned job includes the modelId used for enrichment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaidata/CancelEnrichmentJob.go.html to see an example of how to use CancelEnrichmentJob API.
// A default retry strategy applies to this operation CancelEnrichmentJob()
func (client EnrichmentJobClient) CancelEnrichmentJob(ctx context.Context, request CancelEnrichmentJobRequest) (response CancelEnrichmentJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelEnrichmentJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelEnrichmentJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelEnrichmentJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelEnrichmentJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelEnrichmentJobResponse")
	}
	return
}

// cancelEnrichmentJob implements the OCIOperation interface (enables retrying operations)
func (client EnrichmentJobClient) cancelEnrichmentJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/semanticStores/{semanticStoreId}/enrichmentJobs/{enrichmentJobId}/actions/cancel", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelEnrichmentJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "enrichmentJob", "CancelEnrichmentJob")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-nl2sql/20260325/EnrichmentJob/CancelEnrichmentJob"
		err = common.PostProcessServiceError(err, "EnrichmentJob", "CancelEnrichmentJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateEnrichmentJob Creates a new asynchronous EnrichmentJob.
// The returned job includes the modelId used for enrichment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaidata/GenerateEnrichmentJob.go.html to see an example of how to use GenerateEnrichmentJob API.
// A default retry strategy applies to this operation GenerateEnrichmentJob()
func (client EnrichmentJobClient) GenerateEnrichmentJob(ctx context.Context, request GenerateEnrichmentJobRequest) (response GenerateEnrichmentJobResponse, err error) {
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
func (client EnrichmentJobClient) generateEnrichmentJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/semanticStores/{semanticStoreId}/actions/enrich", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateEnrichmentJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "enrichmentJob", "GenerateEnrichmentJob")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-nl2sql/20260325/EnrichmentJob/GenerateEnrichmentJob"
		err = common.PostProcessServiceError(err, "EnrichmentJob", "GenerateEnrichmentJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEnrichmentJob Gets an EnrichmentJob by identifier.
// The returned job includes the modelId used for enrichment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaidata/GetEnrichmentJob.go.html to see an example of how to use GetEnrichmentJob API.
// A default retry strategy applies to this operation GetEnrichmentJob()
func (client EnrichmentJobClient) GetEnrichmentJob(ctx context.Context, request GetEnrichmentJobRequest) (response GetEnrichmentJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEnrichmentJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEnrichmentJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEnrichmentJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEnrichmentJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEnrichmentJobResponse")
	}
	return
}

// getEnrichmentJob implements the OCIOperation interface (enables retrying operations)
func (client EnrichmentJobClient) getEnrichmentJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/semanticStores/{semanticStoreId}/enrichmentJobs/{enrichmentJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEnrichmentJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "enrichmentJob", "GetEnrichmentJob")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-nl2sql/20260325/EnrichmentJob/GetEnrichmentJob"
		err = common.PostProcessServiceError(err, "EnrichmentJob", "GetEnrichmentJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEnrichmentJobs Lists EnrichmentJobs in a semantic store.
// Each returned job summary includes the modelId used for enrichment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaidata/ListEnrichmentJobs.go.html to see an example of how to use ListEnrichmentJobs API.
// A default retry strategy applies to this operation ListEnrichmentJobs()
func (client EnrichmentJobClient) ListEnrichmentJobs(ctx context.Context, request ListEnrichmentJobsRequest) (response ListEnrichmentJobsResponse, err error) {
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
func (client EnrichmentJobClient) listEnrichmentJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/semanticStores/{semanticStoreId}/enrichmentJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEnrichmentJobsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "enrichmentJob", "ListEnrichmentJobs")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-nl2sql/20260325/EnrichmentJobCollection/ListEnrichmentJobs"
		err = common.PostProcessServiceError(err, "EnrichmentJob", "ListEnrichmentJobs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
