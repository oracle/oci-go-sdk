// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v43/common"
	"github.com/oracle/oci-go-sdk/v43/common/auth"
	"net/http"
)

//OrdersClient a client for Orders
type OrdersClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOrdersClientWithConfigurationProvider Creates a new default Orders client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOrdersClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OrdersClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newOrdersClientFromBaseClient(baseClient, provider)
}

// NewOrdersClientWithOboToken Creates a new default Orders client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewOrdersClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OrdersClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOrdersClientFromBaseClient(baseClient, configProvider)
}

func newOrdersClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OrdersClient, err error) {
	client = OrdersClient{BaseClient: baseClient}
	client.BasePath = "20200801"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OrdersClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("tenantmanagercontrolplane", "https://organizations.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OrdersClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OrdersClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ActivateOrder Triggers an order activation workflow on behalf of the tenant given by compartment id in the body.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ActivateOrder.go.html to see an example of how to use ActivateOrder API.
func (client OrdersClient) ActivateOrder(ctx context.Context, request ActivateOrderRequest) (response ActivateOrderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.activateOrder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateOrderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateOrderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateOrderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateOrderResponse")
	}
	return
}

// activateOrder implements the OCIOperation interface (enables retrying operations)
func (client OrdersClient) activateOrder(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/orders/{activationToken}/actions/activate", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ActivateOrderResponse
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

// GetOrder Returns the Order Details given by the order id in the JWT
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/GetOrder.go.html to see an example of how to use GetOrder API.
func (client OrdersClient) GetOrder(ctx context.Context, request GetOrderRequest) (response GetOrderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOrder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOrderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOrderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOrderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOrderResponse")
	}
	return
}

// getOrder implements the OCIOperation interface (enables retrying operations)
func (client OrdersClient) getOrder(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/orders/{activationToken}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetOrderResponse
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
