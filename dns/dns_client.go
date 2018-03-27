// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Public DNS Service
//
// API for managing DNS zones, records, and policies.
//

package dns

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//DnsClient a client for Dns
type DnsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDnsClientWithConfigurationProvider Creates a new default Dns client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDnsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DnsClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = DnsClient{BaseClient: baseClient}
	client.BasePath = "20180115"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DnsClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "dns", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DnsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DnsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateZone Creates a new zone in the specified compartment. The `compartmentId`
// query parameter is required if the `Content-Type` header for the
// request is `text/dns`.
func (client DnsClient) CreateZone(ctx context.Context, request CreateZoneRequest) (response CreateZoneResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/zones", request)
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

// DeleteDomainRecords Deletes all records at the specified zone and domain.
func (client DnsClient) DeleteDomainRecords(ctx context.Context, request DeleteDomainRecordsRequest) (response DeleteDomainRecordsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/zones/{zoneNameOrId}/records/{domain}", request)
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

// DeleteRRSet Deletes all records in the specified RRSet.
func (client DnsClient) DeleteRRSet(ctx context.Context, request DeleteRRSetRequest) (response DeleteRRSetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/zones/{zoneNameOrId}/records/{domain}/{rtype}", request)
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

// DeleteZone Deletes the specified zone. A `204` response indicates that zone has been
// successfully deleted.
func (client DnsClient) DeleteZone(ctx context.Context, request DeleteZoneRequest) (response DeleteZoneResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/zones/{zoneNameOrId}", request)
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

// GetDomainRecords Gets a list of all records at the specified zone and domain.
// The results are sorted by `rtype` in alphabetical order by default. You
// can optionally filter and/or sort the results using the listed parameters.
func (client DnsClient) GetDomainRecords(ctx context.Context, request GetDomainRecordsRequest) (response GetDomainRecordsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/zones/{zoneNameOrId}/records/{domain}", request)
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

// GetRRSet Gets a list of all records in the specified RRSet. The results are
// sorted by `recordHash` by default.
func (client DnsClient) GetRRSet(ctx context.Context, request GetRRSetRequest) (response GetRRSetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/zones/{zoneNameOrId}/records/{domain}/{rtype}", request)
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

// GetZone Gets information about the specified zone, including its creation date,
// zone type, and serial.
func (client DnsClient) GetZone(ctx context.Context, request GetZoneRequest) (response GetZoneResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/zones/{zoneNameOrId}", request)
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

// GetZoneRecords Gets all records in the specified zone. The results are
// sorted by `domain` in alphabetical order by default. For more
// information about records, please see Resource Record (RR) TYPEs (https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4).
func (client DnsClient) GetZoneRecords(ctx context.Context, request GetZoneRecordsRequest) (response GetZoneRecordsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/zones/{zoneNameOrId}/records", request)
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

// ListZones Gets a list of all zones in the specified compartment. The collection
// can be filtered by name, time created, and zone type.
func (client DnsClient) ListZones(ctx context.Context, request ListZonesRequest) (response ListZonesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/zones", request)
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

// PatchDomainRecords Replaces records in the specified zone at a domain. You can update one record or all records for the specified zone depending on the changes provided in the request body. You can also add or remove records using this function.
func (client DnsClient) PatchDomainRecords(ctx context.Context, request PatchDomainRecordsRequest) (response PatchDomainRecordsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPatch, "/zones/{zoneNameOrId}/records/{domain}", request)
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

// PatchRRSet Updates records in the specified RRSet.
func (client DnsClient) PatchRRSet(ctx context.Context, request PatchRRSetRequest) (response PatchRRSetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPatch, "/zones/{zoneNameOrId}/records/{domain}/{rtype}", request)
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

// PatchZoneRecords Updates a collection of records in the specified zone. You can update
// one record or all records for the specified zone depending on the
// changes provided in the request body. You can also add or remove records
// using this function.
func (client DnsClient) PatchZoneRecords(ctx context.Context, request PatchZoneRecordsRequest) (response PatchZoneRecordsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPatch, "/zones/{zoneNameOrId}/records", request)
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

// UpdateDomainRecords Replaces records in the specified zone at a domain with the records
// specified in the request body. If a specified record does not exist,
// it will be created. If the record exists, then it will be updated to
// represent the record in the body of the request. If a record in the zone
// does not exist in the request body, the record will be removed from the
// zone.
func (client DnsClient) UpdateDomainRecords(ctx context.Context, request UpdateDomainRecordsRequest) (response UpdateDomainRecordsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/zones/{zoneNameOrId}/records/{domain}", request)
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

// UpdateRRSet Replaces records in the specified RRSet.
func (client DnsClient) UpdateRRSet(ctx context.Context, request UpdateRRSetRequest) (response UpdateRRSetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/zones/{zoneNameOrId}/records/{domain}/{rtype}", request)
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

// UpdateZone Updates the specified secondary zone with your new external master
// server information. For more information about secondary zone, see
// Manage DNS Service Zone (https://docs.us-phoenix-1.oraclecloud.com/Content/DNS/Tasks/managingdnszones.htm).
func (client DnsClient) UpdateZone(ctx context.Context, request UpdateZoneRequest) (response UpdateZoneResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/zones/{zoneNameOrId}", request)
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

// UpdateZoneRecords Replaces records in the specified zone with the records specified in the
// request body. If a specified record does not exist, it will be created.
// If the record exists, then it will be updated to represent the record in
// the body of the request. If a record in the zone does not exist in the
// request body, the record will be removed from the zone.
func (client DnsClient) UpdateZoneRecords(ctx context.Context, request UpdateZoneRecordsRequest) (response UpdateZoneRecordsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/zones/{zoneNameOrId}/records", request)
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
