// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Audit API
//
// API for the Audit Service. You can use this API for queries, but not bulk-export operations.
//

package audit

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//AuditClient a client for Audit
type AuditClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAuditClientWithConfigurationProvider Creates a new default Audit client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAuditClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AuditClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = AuditClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AuditClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "audit", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AuditClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.config = &configProvider
	client.SetRegion(region)
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *AuditClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetConfiguration Get the configuration
func (client AuditClient) GetConfiguration(ctx context.Context, request GetConfigurationRequest) (response GetConfigurationResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getConfiguration, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetConfigurationResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client AuditClient) getConfiguration(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/configuration")
	if err != nil {
		return nil, err
	}

	var response GetConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListEvents Returns all audit events for the specified compartment that were processed within the specified time range.
func (client AuditClient) ListEvents(ctx context.Context, request ListEventsRequest) (response ListEventsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listEvents, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListEventsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client AuditClient) listEvents(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/auditEvents")
	if err != nil {
		return nil, err
	}

	var response ListEventsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateConfiguration Update the configuration
func (client AuditClient) UpdateConfiguration(ctx context.Context, request UpdateConfigurationRequest) (response UpdateConfigurationResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateConfiguration, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateConfigurationResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client AuditClient) updateConfiguration(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/configuration")
	if err != nil {
		return nil, err
	}

	var response UpdateConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}
