// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// TransferPackageClient a client for TransferPackage
type TransferPackageClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewTransferPackageClientWithConfigurationProvider Creates a new default TransferPackage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewTransferPackageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client TransferPackageClient, err error) {
	if enabled := common.CheckForEnabledServices("dts"); !enabled {
		return client, fmt.Errorf("the Developer Tool configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local developer-tool-configuration.json file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newTransferPackageClientFromBaseClient(baseClient, provider)
}

// NewTransferPackageClientWithOboToken Creates a new default TransferPackage client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewTransferPackageClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client TransferPackageClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newTransferPackageClientFromBaseClient(baseClient, configProvider)
}

func newTransferPackageClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client TransferPackageClient, err error) {
	// TransferPackage service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("TransferPackage"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = TransferPackageClient{BaseClient: baseClient}
	client.BasePath = "20171001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *TransferPackageClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dts", "https://datatransfer.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *TransferPackageClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *TransferPackageClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AttachDevicesToTransferPackage Attaches Devices to a Transfer Package
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/AttachDevicesToTransferPackage.go.html to see an example of how to use AttachDevicesToTransferPackage API.
func (client TransferPackageClient) AttachDevicesToTransferPackage(ctx context.Context, request AttachDevicesToTransferPackageRequest) (response AttachDevicesToTransferPackageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.attachDevicesToTransferPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachDevicesToTransferPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachDevicesToTransferPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachDevicesToTransferPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachDevicesToTransferPackageResponse")
	}
	return
}

// attachDevicesToTransferPackage implements the OCIOperation interface (enables retrying operations)
func (client TransferPackageClient) attachDevicesToTransferPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transferJobs/{id}/transferPackages/{transferPackageLabel}/actions/attachDevices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachDevicesToTransferPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "TransferPackage", "AttachDevicesToTransferPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTransferPackage Create a new Transfer Package
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/CreateTransferPackage.go.html to see an example of how to use CreateTransferPackage API.
func (client TransferPackageClient) CreateTransferPackage(ctx context.Context, request CreateTransferPackageRequest) (response CreateTransferPackageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTransferPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTransferPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTransferPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTransferPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTransferPackageResponse")
	}
	return
}

// createTransferPackage implements the OCIOperation interface (enables retrying operations)
func (client TransferPackageClient) createTransferPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transferJobs/{id}/transferPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTransferPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "TransferPackage", "CreateTransferPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTransferPackage deletes a transfer Package
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/DeleteTransferPackage.go.html to see an example of how to use DeleteTransferPackage API.
func (client TransferPackageClient) DeleteTransferPackage(ctx context.Context, request DeleteTransferPackageRequest) (response DeleteTransferPackageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteTransferPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTransferPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTransferPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTransferPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTransferPackageResponse")
	}
	return
}

// deleteTransferPackage implements the OCIOperation interface (enables retrying operations)
func (client TransferPackageClient) deleteTransferPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/transferJobs/{id}/transferPackages/{transferPackageLabel}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTransferPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "TransferPackage", "DeleteTransferPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachDevicesFromTransferPackage Detaches Devices from a Transfer Package
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/DetachDevicesFromTransferPackage.go.html to see an example of how to use DetachDevicesFromTransferPackage API.
func (client TransferPackageClient) DetachDevicesFromTransferPackage(ctx context.Context, request DetachDevicesFromTransferPackageRequest) (response DetachDevicesFromTransferPackageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.detachDevicesFromTransferPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachDevicesFromTransferPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachDevicesFromTransferPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachDevicesFromTransferPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachDevicesFromTransferPackageResponse")
	}
	return
}

// detachDevicesFromTransferPackage implements the OCIOperation interface (enables retrying operations)
func (client TransferPackageClient) detachDevicesFromTransferPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transferJobs/{id}/transferPackages/{transferPackageLabel}/actions/detachDevices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachDevicesFromTransferPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "TransferPackage", "DetachDevicesFromTransferPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTransferPackage Describes a transfer package in detail
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/GetTransferPackage.go.html to see an example of how to use GetTransferPackage API.
func (client TransferPackageClient) GetTransferPackage(ctx context.Context, request GetTransferPackageRequest) (response GetTransferPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTransferPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTransferPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTransferPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTransferPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTransferPackageResponse")
	}
	return
}

// getTransferPackage implements the OCIOperation interface (enables retrying operations)
func (client TransferPackageClient) getTransferPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferJobs/{id}/transferPackages/{transferPackageLabel}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTransferPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "TransferPackage", "GetTransferPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTransferPackages Lists Transfer Packages associated with a transferJob
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/ListTransferPackages.go.html to see an example of how to use ListTransferPackages API.
func (client TransferPackageClient) ListTransferPackages(ctx context.Context, request ListTransferPackagesRequest) (response ListTransferPackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTransferPackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTransferPackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTransferPackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTransferPackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTransferPackagesResponse")
	}
	return
}

// listTransferPackages implements the OCIOperation interface (enables retrying operations)
func (client TransferPackageClient) listTransferPackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferJobs/{id}/transferPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTransferPackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "TransferPackage", "ListTransferPackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTransferPackage Updates a Transfer Package
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/UpdateTransferPackage.go.html to see an example of how to use UpdateTransferPackage API.
func (client TransferPackageClient) UpdateTransferPackage(ctx context.Context, request UpdateTransferPackageRequest) (response UpdateTransferPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTransferPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTransferPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTransferPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTransferPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTransferPackageResponse")
	}
	return
}

// updateTransferPackage implements the OCIOperation interface (enables retrying operations)
func (client TransferPackageClient) updateTransferPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/transferJobs/{id}/transferPackages/{transferPackageLabel}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTransferPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "TransferPackage", "UpdateTransferPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
