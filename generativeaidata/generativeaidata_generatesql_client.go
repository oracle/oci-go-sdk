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

// GenerateSqlClient a client for GenerateSql
type GenerateSqlClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewGenerateSqlClientWithConfigurationProvider Creates a new default GenerateSql client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewGenerateSqlClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client GenerateSqlClient, err error) {
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
	return newGenerateSqlClientFromBaseClient(baseClient, provider)
}

// NewGenerateSqlClientWithOboToken Creates a new default GenerateSql client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewGenerateSqlClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client GenerateSqlClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newGenerateSqlClientFromBaseClient(baseClient, configProvider)
}

func newGenerateSqlClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client GenerateSqlClient, err error) {
	// GenerateSql service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("GenerateSql"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = GenerateSqlClient{BaseClient: baseClient}
	client.BasePath = "20260325"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *GenerateSqlClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("generativeaidata", "https://inference.generativeai.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *GenerateSqlClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *GenerateSqlClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GenerateSqlFromNl Generates a SQL query from a natural language input for the specified SemanticStore.
// This operation creates a GenerateSqlFromNlJob. The request can either be processed as a
// background job or wait for completion, depending on completionMode.
// If modelId is provided, the service validates it and uses it for schema-linking in table selection,
// SQL generation, and SQL refinement. If modelId is omitted, the service default model is used. The response
// includes the modelId used for generation.
// Clients should inspect lifecycleState in the response:
// - If lifecycleState is SUCCEEDED, SQL generation is complete and the result is available in jobOutput.
// - If lifecycleState is ACCEPTED or IN_PROGRESS, poll GetGenerateSqlFromNlJob using the returned job id.
// - If lifecycleState is FAILED or CANCELED, SQL generation did not complete. See lifecycleDetails for more information.
// The GetGenerateSqlFromNlJob endpoint is the source of truth for final job state and result availability.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaidata/GenerateSqlFromNl.go.html to see an example of how to use GenerateSqlFromNl API.
// A default retry strategy applies to this operation GenerateSqlFromNl()
func (client GenerateSqlClient) GenerateSqlFromNl(ctx context.Context, request GenerateSqlFromNlRequest) (response GenerateSqlFromNlResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateSqlFromNl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateSqlFromNlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateSqlFromNlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateSqlFromNlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateSqlFromNlResponse")
	}
	return
}

// generateSqlFromNl implements the OCIOperation interface (enables retrying operations)
func (client GenerateSqlClient) generateSqlFromNl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/semanticStores/{semanticStoreId}/actions/generateSqlFromNl", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateSqlFromNlResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generateSql", "GenerateSqlFromNl")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-nl2sql/20260325/GenerateSqlFromNlJob/GenerateSqlFromNl"
		err = common.PostProcessServiceError(err, "GenerateSql", "GenerateSqlFromNl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetGenerateSqlFromNlJob Retrieves the current state of a GenerateSqlFromNlJob.
// Clients should poll this endpoint until lifecycleState is SUCCEEDED or FAILED.
// When lifecycleState is SUCCEEDED, the result is available in jobOutput.
// The returned job includes the modelId used for generation.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaidata/GetGenerateSqlFromNlJob.go.html to see an example of how to use GetGenerateSqlFromNlJob API.
// A default retry strategy applies to this operation GetGenerateSqlFromNlJob()
func (client GenerateSqlClient) GetGenerateSqlFromNlJob(ctx context.Context, request GetGenerateSqlFromNlJobRequest) (response GetGenerateSqlFromNlJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGenerateSqlFromNlJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGenerateSqlFromNlJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGenerateSqlFromNlJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGenerateSqlFromNlJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGenerateSqlFromNlJobResponse")
	}
	return
}

// getGenerateSqlFromNlJob implements the OCIOperation interface (enables retrying operations)
func (client GenerateSqlClient) getGenerateSqlFromNlJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/semanticStores/{semanticStoreId}/generateSqlFromNlJobs/{generateSqlFromNlJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGenerateSqlFromNlJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generateSql", "GetGenerateSqlFromNlJob")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-nl2sql/20260325/GenerateSqlFromNlJob/GetGenerateSqlFromNlJob"
		err = common.PostProcessServiceError(err, "GenerateSql", "GetGenerateSqlFromNlJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
