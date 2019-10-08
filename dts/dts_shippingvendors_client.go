// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//ShippingVendorsClient a client for ShippingVendors
type ShippingVendorsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewShippingVendorsClientWithConfigurationProvider Creates a new default ShippingVendors client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewShippingVendorsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ShippingVendorsClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = ShippingVendorsClient{BaseClient: baseClient}
	client.BasePath = "20171001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ShippingVendorsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dts", "https://datatransfer.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ShippingVendorsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ShippingVendorsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListShippingVendors Lists available shipping vendors for Transfer Package delivery
func (client ShippingVendorsClient) ListShippingVendors(ctx context.Context, request ListShippingVendorsRequest) (response ListShippingVendorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listShippingVendors, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListShippingVendorsResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListShippingVendorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListShippingVendorsResponse")
	}
	return
}

// listShippingVendors implements the OCIOperation interface (enables retrying operations)
func (client ShippingVendorsClient) listShippingVendors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/shippingVendors")
	if err != nil {
		return nil, err
	}

	var response ListShippingVendorsResponse
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
