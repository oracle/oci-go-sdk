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

//TransferDeviceClient a client for TransferDevice
type TransferDeviceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewTransferDeviceClientWithConfigurationProvider Creates a new default TransferDevice client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewTransferDeviceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client TransferDeviceClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = TransferDeviceClient{BaseClient: baseClient}
	client.BasePath = "20171001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *TransferDeviceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dts", "https://datatransfer.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *TransferDeviceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *TransferDeviceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateTransferDevice Create a new Transfer Device
func (client TransferDeviceClient) CreateTransferDevice(ctx context.Context, request CreateTransferDeviceRequest) (response CreateTransferDeviceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTransferDevice, policy)
	if err != nil {
		if ociResponse != nil {
			response = CreateTransferDeviceResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTransferDeviceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTransferDeviceResponse")
	}
	return
}

// createTransferDevice implements the OCIOperation interface (enables retrying operations)
func (client TransferDeviceClient) createTransferDevice(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transferJobs/{id}/transferDevices")
	if err != nil {
		return nil, err
	}

	var response CreateTransferDeviceResponse
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

// DeleteTransferDevice deletes a transfer Device
func (client TransferDeviceClient) DeleteTransferDevice(ctx context.Context, request DeleteTransferDeviceRequest) (response DeleteTransferDeviceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTransferDevice, policy)
	if err != nil {
		if ociResponse != nil {
			response = DeleteTransferDeviceResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTransferDeviceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTransferDeviceResponse")
	}
	return
}

// deleteTransferDevice implements the OCIOperation interface (enables retrying operations)
func (client TransferDeviceClient) deleteTransferDevice(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/transferJobs/{id}/transferDevices/{transferDeviceLabel}")
	if err != nil {
		return nil, err
	}

	var response DeleteTransferDeviceResponse
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

// GetTransferDevice Describes a transfer package in detail
func (client TransferDeviceClient) GetTransferDevice(ctx context.Context, request GetTransferDeviceRequest) (response GetTransferDeviceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTransferDevice, policy)
	if err != nil {
		if ociResponse != nil {
			response = GetTransferDeviceResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTransferDeviceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTransferDeviceResponse")
	}
	return
}

// getTransferDevice implements the OCIOperation interface (enables retrying operations)
func (client TransferDeviceClient) getTransferDevice(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferJobs/{id}/transferDevices/{transferDeviceLabel}")
	if err != nil {
		return nil, err
	}

	var response GetTransferDeviceResponse
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

// ListTransferDevices Lists Transfer Devices associated with a transferJob
func (client TransferDeviceClient) ListTransferDevices(ctx context.Context, request ListTransferDevicesRequest) (response ListTransferDevicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTransferDevices, policy)
	if err != nil {
		if ociResponse != nil {
			response = ListTransferDevicesResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTransferDevicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTransferDevicesResponse")
	}
	return
}

// listTransferDevices implements the OCIOperation interface (enables retrying operations)
func (client TransferDeviceClient) listTransferDevices(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferJobs/{id}/transferDevices")
	if err != nil {
		return nil, err
	}

	var response ListTransferDevicesResponse
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

// UpdateTransferDevice Updates a Transfer Device
func (client TransferDeviceClient) UpdateTransferDevice(ctx context.Context, request UpdateTransferDeviceRequest) (response UpdateTransferDeviceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTransferDevice, policy)
	if err != nil {
		if ociResponse != nil {
			response = UpdateTransferDeviceResponse{RawResponse: ociResponse.HTTPResponse()}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTransferDeviceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTransferDeviceResponse")
	}
	return
}

// updateTransferDevice implements the OCIOperation interface (enables retrying operations)
func (client TransferDeviceClient) updateTransferDevice(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/transferJobs/{id}/transferDevices/{transferDeviceLabel}")
	if err != nil {
		return nil, err
	}

	var response UpdateTransferDeviceResponse
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
