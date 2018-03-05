// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
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

//AuditClientData a client for Audit
type AuditClientData struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAuditClientWithConfigurationProvider Creates a new default Audit client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAuditClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (AuditClient, error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return nil, err
	}

	client := AuditClientData{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	return &client, err
}

// GetBaseClient get the BaseClient object of this client
func (client *AuditClientData) GetBaseClient() *common.BaseClient {
	return &client.BaseClient
}

// SetRegion overrides the region of this client.
func (client *AuditClientData) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "audit", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AuditClientData) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *AuditClientData) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetConfiguration Get the configuration
func (client AuditClientData) GetConfiguration(ctx context.Context, request GetConfigurationRequest) (response GetConfigurationResponse, err error) {
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
func (client AuditClientData) ListEvents(ctx context.Context, request ListEventsRequest) (response ListEventsResponse, err error) {
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
func (client AuditClientData) UpdateConfiguration(ctx context.Context, request UpdateConfigurationRequest) (response UpdateConfigurationResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/configuration", request)
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

// AuditClient is client interface for Audit
type AuditClient interface {
	GetBaseClient() *common.BaseClient
	ConfigurationProvider() *common.ConfigurationProvider
	GetConfiguration(ctx context.Context, request GetConfigurationRequest) (response GetConfigurationResponse, err error)
	ListEvents(ctx context.Context, request ListEventsRequest) (response ListEventsResponse, err error)
	UpdateConfiguration(ctx context.Context, request UpdateConfigurationRequest) (response UpdateConfigurationResponse, err error)
}
