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
}

// NewAuditClientWithConfigurationProvider Creates a new default Audit client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAuditClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AuditClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = AuditClient{BaseClient: baseClient}
	region, err := configProvider.Region()
	if err != nil {
		return
	}

	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "audit", string(region))
	client.BasePath = "20160918"
	return
}

// GetConfiguration Get the configuration
func (client AuditClient) GetConfiguration(ctx context.Context, request GetConfigurationRequest) (response GetConfigurationResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/configuration", request)
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

// GetWorkRequest Gets details on a specified work request. The workRequestId is returned in opc-workrequest-id
// header for any asynchronous operation on Identity control plane service.
func (client AuditClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/workRequests/{workRequestId}", request)
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
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodGet, "/auditEvents", request)
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
	httpRequest, err := common.MakeDefaultHttpRequestWithTaggedStruct(http.MethodPut, "/configuration", request)
	if err != nil {
		return
	}

	_, err = client.Call(ctx, &httpRequest)
	return
}
