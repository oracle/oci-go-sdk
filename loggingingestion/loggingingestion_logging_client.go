// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PublicLoggingDataplane API
//
// PublicLoggingDataplane API specification
//

package loggingingestion

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//LoggingClient a client for Logging
type LoggingClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLoggingClientWithConfigurationProvider Creates a new default Logging client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLoggingClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LoggingClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newLoggingClientFromBaseClient(baseClient, configProvider)
}

// NewLoggingClientWithOboToken Creates a new default Logging client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewLoggingClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LoggingClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newLoggingClientFromBaseClient(baseClient, configProvider)
}

func newLoggingClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LoggingClient, err error) {
	client = LoggingClient{BaseClient: baseClient}
	client.BasePath = "20200831"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LoggingClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("loggingingestion", "https://ingestion.logging.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LoggingClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *LoggingClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// PutLogs This Api allows ingesting logs associated with a logId. Success
// response implies the data has been accepted.
func (client LoggingClient) PutLogs(ctx context.Context, request PutLogsRequest) (response PutLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.putLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutLogsResponse")
	}
	return
}

// putLogs implements the OCIOperation interface (enables retrying operations)
func (client LoggingClient) putLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/logs/{logId}/actions/push")
	if err != nil {
		return nil, err
	}

	var response PutLogsResponse
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
