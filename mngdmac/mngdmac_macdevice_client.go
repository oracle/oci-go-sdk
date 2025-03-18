// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Managed Services for Mac API
//
// Use the OCI Managed Services for Mac API to create and manage orders for dedicated, partially-managed Mac servers hosted in an OCI Data Center. For more information, see the user guide documentation for the [OCI Managed Services for Mac]()
//

package mngdmac

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// MacDeviceClient a client for MacDevice
type MacDeviceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMacDeviceClientWithConfigurationProvider Creates a new default MacDevice client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMacDeviceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MacDeviceClient, err error) {
	if enabled := common.CheckForEnabledServices("mngdmac"); !enabled {
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
	return newMacDeviceClientFromBaseClient(baseClient, provider)
}

// NewMacDeviceClientWithOboToken Creates a new default MacDevice client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewMacDeviceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client MacDeviceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newMacDeviceClientFromBaseClient(baseClient, configProvider)
}

func newMacDeviceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client MacDeviceClient, err error) {
	// MacDevice service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("MacDevice"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = MacDeviceClient{BaseClient: baseClient}
	client.BasePath = "20250320"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *MacDeviceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("mngdmac", "https://mngdmac.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *MacDeviceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *MacDeviceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetMacDevice Gets information about a MacDevice.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mngdmac/GetMacDevice.go.html to see an example of how to use GetMacDevice API.
// A default retry strategy applies to this operation GetMacDevice()
func (client MacDeviceClient) GetMacDevice(ctx context.Context, request GetMacDeviceRequest) (response GetMacDeviceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMacDevice, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMacDeviceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMacDeviceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMacDeviceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMacDeviceResponse")
	}
	return
}

// getMacDevice implements the OCIOperation interface (enables retrying operations)
func (client MacDeviceClient) getMacDevice(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/macOrders/{macOrderId}/macDevices/{macDeviceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMacDeviceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "MacDevice", "GetMacDevice", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMacDevices Gets a list of MacDevices assigned to this order.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mngdmac/ListMacDevices.go.html to see an example of how to use ListMacDevices API.
// A default retry strategy applies to this operation ListMacDevices()
func (client MacDeviceClient) ListMacDevices(ctx context.Context, request ListMacDevicesRequest) (response ListMacDevicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMacDevices, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMacDevicesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMacDevicesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMacDevicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMacDevicesResponse")
	}
	return
}

// listMacDevices implements the OCIOperation interface (enables retrying operations)
func (client MacDeviceClient) listMacDevices(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/macOrders/{macOrderId}/macDevices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMacDevicesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "MacDevice", "ListMacDevices", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TerminateMacDevice Terminates a MacDevice.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mngdmac/TerminateMacDevice.go.html to see an example of how to use TerminateMacDevice API.
// A default retry strategy applies to this operation TerminateMacDevice()
func (client MacDeviceClient) TerminateMacDevice(ctx context.Context, request TerminateMacDeviceRequest) (response TerminateMacDeviceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.terminateMacDevice, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TerminateMacDeviceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TerminateMacDeviceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TerminateMacDeviceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TerminateMacDeviceResponse")
	}
	return
}

// terminateMacDevice implements the OCIOperation interface (enables retrying operations)
func (client MacDeviceClient) terminateMacDevice(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/macOrders/{macOrderId}/macDevices/{macDeviceId}/actions/terminate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TerminateMacDeviceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "MacDevice", "TerminateMacDevice", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
