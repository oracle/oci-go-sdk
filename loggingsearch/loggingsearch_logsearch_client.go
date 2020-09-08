// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Search API
//
// Search for logs in your compartements / log groups / log objects.
//

package loggingsearch

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//LogSearchClient a client for LogSearch
type LogSearchClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLogSearchClientWithConfigurationProvider Creates a new default LogSearch client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLogSearchClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LogSearchClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newLogSearchClientFromBaseClient(baseClient, configProvider)
}

// NewLogSearchClientWithOboToken Creates a new default LogSearch client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewLogSearchClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LogSearchClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newLogSearchClientFromBaseClient(baseClient, configProvider)
}

func newLogSearchClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LogSearchClient, err error) {
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
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *LogSearchClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// SearchLogs Submit a query to search logs.
func (client LogSearchClient) SearchLogs(ctx context.Context, request SearchLogsRequest) (response SearchLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client LogSearchClient) searchLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/search")
	if err != nil {
		return nil, err
	}

	var response SearchLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
