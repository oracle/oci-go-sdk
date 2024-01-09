// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Search API
//
// Search for logs in your compartments, log groups, and log objects.
//

package loggingsearch

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// LogSearchClient a client for LogSearch
type LogSearchClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLogSearchClientWithConfigurationProvider Creates a new default LogSearch client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLogSearchClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LogSearchClient, err error) {
	if enabled := common.CheckForEnabledServices("loggingsearch"); !enabled {
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
	return newLogSearchClientFromBaseClient(baseClient, provider)
}

// NewLogSearchClientWithOboToken Creates a new default LogSearch client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewLogSearchClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LogSearchClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newLogSearchClientFromBaseClient(baseClient, configProvider)
}

func newLogSearchClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LogSearchClient, err error) {
	// LogSearch service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("LogSearch"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = LogSearchClient{BaseClient: baseClient}
	client.BasePath = "20190909"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LogSearchClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("loggingsearch", "https://logging.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LogSearchClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *LogSearchClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// SearchLogs Submit a query to search logs.
// See Using the API (https://docs.cloud.oracle.com/Content/Logging/Concepts/using_the_api_searchlogs.htm) for SDK examples.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loggingsearch/SearchLogs.go.html to see an example of how to use SearchLogs API.
func (client LogSearchClient) SearchLogs(ctx context.Context, request SearchLogsRequest) (response SearchLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.searchLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchLogsResponse")
	}
	return
}

// searchLogs implements the OCIOperation interface (enables retrying operations)
func (client LogSearchClient) searchLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/logging-search/20190909/SearchResult/SearchLogs"
		err = common.PostProcessServiceError(err, "LogSearch", "SearchLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
