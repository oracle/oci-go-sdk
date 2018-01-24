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
	err = client.SetConfigurationProvider(configProvider)
	return
}

// SetConfigurationProvider sets the configuration provider, returns an error if is not valid
func (client *AuditClient) SetConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	region, err := configProvider.Region()
	if err != nil {
		return err
	}
	client.config = &configProvider
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "audit", string(region))
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *AuditClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetConfiguration Get the configuration
func (client AuditClient) GetConfiguration(ctx context.Context, request GetConfigurationRequest) (response GetConfigurationResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/configuration", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// ListEvents Returns all audit events for the specified compartment that were processed within the specified time range.
func (client AuditClient) ListEvents(ctx context.Context, request ListEventsRequest) (response ListEventsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/auditEvents", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// UpdateConfiguration Update the configuration
func (client AuditClient) UpdateConfiguration(ctx context.Context, request UpdateConfigurationRequest) (err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/configuration", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}
