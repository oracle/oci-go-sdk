// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v48/common"
	"github.com/oracle/oci-go-sdk/v48/common/auth"
	"net/http"
)

//TransferApplianceEntitlementClient a client for TransferApplianceEntitlement
type TransferApplianceEntitlementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewTransferApplianceEntitlementClientWithConfigurationProvider Creates a new default TransferApplianceEntitlement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewTransferApplianceEntitlementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client TransferApplianceEntitlementClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newTransferApplianceEntitlementClientFromBaseClient(baseClient, provider)
}

// NewTransferApplianceEntitlementClientWithOboToken Creates a new default TransferApplianceEntitlement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewTransferApplianceEntitlementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client TransferApplianceEntitlementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newTransferApplianceEntitlementClientFromBaseClient(baseClient, configProvider)
}

func newTransferApplianceEntitlementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client TransferApplianceEntitlementClient, err error) {
	client = TransferApplianceEntitlementClient{BaseClient: baseClient}
	client.BasePath = "20171001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *TransferApplianceEntitlementClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dts", "https://datatransfer.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *TransferApplianceEntitlementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *TransferApplianceEntitlementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateTransferApplianceEntitlement Create the Entitlement to use a Transfer Appliance. It requires some offline process of review and signatures before request is granted.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/CreateTransferApplianceEntitlement.go.html to see an example of how to use CreateTransferApplianceEntitlement API.
func (client TransferApplianceEntitlementClient) CreateTransferApplianceEntitlement(ctx context.Context, request CreateTransferApplianceEntitlementRequest) (response CreateTransferApplianceEntitlementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTransferApplianceEntitlement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTransferApplianceEntitlementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTransferApplianceEntitlementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTransferApplianceEntitlementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTransferApplianceEntitlementResponse")
	}
	return
}

// createTransferApplianceEntitlement implements the OCIOperation interface (enables retrying operations)
func (client TransferApplianceEntitlementClient) createTransferApplianceEntitlement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transferApplianceEntitlement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTransferApplianceEntitlementResponse
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

// GetTransferApplianceEntitlement Describes the Transfer Appliance Entitlement in detail
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/GetTransferApplianceEntitlement.go.html to see an example of how to use GetTransferApplianceEntitlement API.
func (client TransferApplianceEntitlementClient) GetTransferApplianceEntitlement(ctx context.Context, request GetTransferApplianceEntitlementRequest) (response GetTransferApplianceEntitlementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getTransferApplianceEntitlement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTransferApplianceEntitlementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTransferApplianceEntitlementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTransferApplianceEntitlementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTransferApplianceEntitlementResponse")
	}
	return
}

// getTransferApplianceEntitlement implements the OCIOperation interface (enables retrying operations)
func (client TransferApplianceEntitlementClient) getTransferApplianceEntitlement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferApplianceEntitlement/{id}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTransferApplianceEntitlementResponse
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

// ListTransferApplianceEntitlement Lists Transfer Transfer Appliance Entitlement
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/ListTransferApplianceEntitlement.go.html to see an example of how to use ListTransferApplianceEntitlement API.
func (client TransferApplianceEntitlementClient) ListTransferApplianceEntitlement(ctx context.Context, request ListTransferApplianceEntitlementRequest) (response ListTransferApplianceEntitlementResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTransferApplianceEntitlement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTransferApplianceEntitlementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTransferApplianceEntitlementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTransferApplianceEntitlementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTransferApplianceEntitlementResponse")
	}
	return
}

// listTransferApplianceEntitlement implements the OCIOperation interface (enables retrying operations)
func (client TransferApplianceEntitlementClient) listTransferApplianceEntitlement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferApplianceEntitlement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTransferApplianceEntitlementResponse
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
