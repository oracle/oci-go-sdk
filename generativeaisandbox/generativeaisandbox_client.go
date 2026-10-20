// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GenAI Sandbox API
//
// API for managing sandbox resources and exchanging sandbox access tokens.
//

package generativeaisandbox

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// GenerativeAiSandboxClient a client for GenerativeAiSandbox
type GenerativeAiSandboxClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewGenerativeAiSandboxClientWithConfigurationProvider Creates a new default GenerativeAiSandbox client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewGenerativeAiSandboxClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client GenerativeAiSandboxClient, err error) {
	if enabled := common.CheckForEnabledServices("generativeaisandbox"); !enabled {
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
	return newGenerativeAiSandboxClientFromBaseClient(baseClient, provider)
}

// NewGenerativeAiSandboxClientWithOboToken Creates a new default GenerativeAiSandbox client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewGenerativeAiSandboxClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client GenerativeAiSandboxClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newGenerativeAiSandboxClientFromBaseClient(baseClient, configProvider)
}

func newGenerativeAiSandboxClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client GenerativeAiSandboxClient, err error) {
	// GenerativeAiSandbox service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("GenerativeAiSandbox"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = GenerativeAiSandboxClient{BaseClient: baseClient}
	client.BasePath = "20260422"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *GenerativeAiSandboxClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("generativeaisandbox", "https://inference.generativeai.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *GenerativeAiSandboxClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *GenerativeAiSandboxClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelSandboxCommand Requests cancellation of a running sandbox command.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/CancelSandboxCommand.go.html to see an example of how to use CancelSandboxCommand API.
// A default retry strategy applies to this operation CancelSandboxCommand()
func (client GenerativeAiSandboxClient) CancelSandboxCommand(ctx context.Context, request CancelSandboxCommandRequest) (response CancelSandboxCommandResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelSandboxCommand, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelSandboxCommandResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelSandboxCommandResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelSandboxCommandResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelSandboxCommandResponse")
	}
	return
}

// cancelSandboxCommand implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) cancelSandboxCommand(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sandboxes/{sandboxId}/commands/{commandId}/actions/cancel", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelSandboxCommandResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "CancelSandboxCommand")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "CancelSandboxCommand", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSandbox Creates a sandbox resource and returns its initial state.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/CreateSandbox.go.html to see an example of how to use CreateSandbox API.
// A default retry strategy applies to this operation CreateSandbox()
func (client GenerativeAiSandboxClient) CreateSandbox(ctx context.Context, request CreateSandboxRequest) (response CreateSandboxResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSandbox, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSandboxResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSandboxResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSandboxResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSandboxResponse")
	}
	return
}

// createSandbox implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) createSandbox(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sandboxes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSandboxResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "CreateSandbox")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "CreateSandbox", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSandboxSnapshot Creates a sandbox snapshot resource and returns its initial state.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/CreateSandboxSnapshot.go.html to see an example of how to use CreateSandboxSnapshot API.
// A default retry strategy applies to this operation CreateSandboxSnapshot()
func (client GenerativeAiSandboxClient) CreateSandboxSnapshot(ctx context.Context, request CreateSandboxSnapshotRequest) (response CreateSandboxSnapshotResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSandboxSnapshot, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSandboxSnapshotResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSandboxSnapshotResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSandboxSnapshotResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSandboxSnapshotResponse")
	}
	return
}

// createSandboxSnapshot implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) createSandboxSnapshot(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/snapshots", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSandboxSnapshotResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "CreateSandboxSnapshot")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "CreateSandboxSnapshot", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSandbox Deletes a sandbox.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/DeleteSandbox.go.html to see an example of how to use DeleteSandbox API.
// A default retry strategy applies to this operation DeleteSandbox()
func (client GenerativeAiSandboxClient) DeleteSandbox(ctx context.Context, request DeleteSandboxRequest) (response DeleteSandboxResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSandbox, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSandboxResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSandboxResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSandboxResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSandboxResponse")
	}
	return
}

// deleteSandbox implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) deleteSandbox(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/sandboxes/{sandboxId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSandboxResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "DeleteSandbox")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "DeleteSandbox", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSandboxSnapshot Deletes a sandbox snapshot.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/DeleteSandboxSnapshot.go.html to see an example of how to use DeleteSandboxSnapshot API.
// A default retry strategy applies to this operation DeleteSandboxSnapshot()
func (client GenerativeAiSandboxClient) DeleteSandboxSnapshot(ctx context.Context, request DeleteSandboxSnapshotRequest) (response DeleteSandboxSnapshotResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSandboxSnapshot, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSandboxSnapshotResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSandboxSnapshotResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSandboxSnapshotResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSandboxSnapshotResponse")
	}
	return
}

// deleteSandboxSnapshot implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) deleteSandboxSnapshot(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/snapshots/{snapshotId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSandboxSnapshotResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "DeleteSandboxSnapshot")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "DeleteSandboxSnapshot", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExtendSandbox Extends the duration of a sandbox.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/ExtendSandbox.go.html to see an example of how to use ExtendSandbox API.
// A default retry strategy applies to this operation ExtendSandbox()
func (client GenerativeAiSandboxClient) ExtendSandbox(ctx context.Context, request ExtendSandboxRequest) (response ExtendSandboxResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.extendSandbox, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExtendSandboxResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExtendSandboxResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExtendSandboxResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExtendSandboxResponse")
	}
	return
}

// extendSandbox implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) extendSandbox(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sandboxes/{sandboxId}/actions/extend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExtendSandboxResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "ExtendSandbox")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "ExtendSandbox", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSandbox Gets information about a sandbox.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/GetSandbox.go.html to see an example of how to use GetSandbox API.
// A default retry strategy applies to this operation GetSandbox()
func (client GenerativeAiSandboxClient) GetSandbox(ctx context.Context, request GetSandboxRequest) (response GetSandboxResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSandbox, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSandboxResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSandboxResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSandboxResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSandboxResponse")
	}
	return
}

// getSandbox implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) getSandbox(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sandboxes/{sandboxId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSandboxResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "GetSandbox")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "GetSandbox", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSandboxCommand Gets command metadata for a sandbox command.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/GetSandboxCommand.go.html to see an example of how to use GetSandboxCommand API.
// A default retry strategy applies to this operation GetSandboxCommand()
func (client GenerativeAiSandboxClient) GetSandboxCommand(ctx context.Context, request GetSandboxCommandRequest) (response GetSandboxCommandResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSandboxCommand, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSandboxCommandResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSandboxCommandResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSandboxCommandResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSandboxCommandResponse")
	}
	return
}

// getSandboxCommand implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) getSandboxCommand(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sandboxes/{sandboxId}/commands/{commandId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSandboxCommandResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "GetSandboxCommand")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "GetSandboxCommand", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSandboxCommandOutput Gets captured output for a sandbox command starting from the requested event sequence.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/GetSandboxCommandOutput.go.html to see an example of how to use GetSandboxCommandOutput API.
// A default retry strategy applies to this operation GetSandboxCommandOutput()
func (client GenerativeAiSandboxClient) GetSandboxCommandOutput(ctx context.Context, request GetSandboxCommandOutputRequest) (response GetSandboxCommandOutputResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSandboxCommandOutput, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSandboxCommandOutputResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSandboxCommandOutputResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSandboxCommandOutputResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSandboxCommandOutputResponse")
	}
	return
}

// getSandboxCommandOutput implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) getSandboxCommandOutput(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sandboxes/{sandboxId}/commands/{commandId}/output", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSandboxCommandOutputResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "GetSandboxCommandOutput")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "GetSandboxCommandOutput", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSandboxSnapshot Gets information about a sandbox snapshot.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/GetSandboxSnapshot.go.html to see an example of how to use GetSandboxSnapshot API.
// A default retry strategy applies to this operation GetSandboxSnapshot()
func (client GenerativeAiSandboxClient) GetSandboxSnapshot(ctx context.Context, request GetSandboxSnapshotRequest) (response GetSandboxSnapshotResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSandboxSnapshot, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSandboxSnapshotResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSandboxSnapshotResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSandboxSnapshotResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSandboxSnapshotResponse")
	}
	return
}

// getSandboxSnapshot implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) getSandboxSnapshot(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/snapshots/{snapshotId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSandboxSnapshotResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "GetSandboxSnapshot")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "GetSandboxSnapshot", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSandboxCommand Lists sandbox commands for a sandbox with optional pagination.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/ListSandboxCommand.go.html to see an example of how to use ListSandboxCommand API.
// A default retry strategy applies to this operation ListSandboxCommand()
func (client GenerativeAiSandboxClient) ListSandboxCommand(ctx context.Context, request ListSandboxCommandRequest) (response ListSandboxCommandResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSandboxCommand, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSandboxCommandResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSandboxCommandResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSandboxCommandResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSandboxCommandResponse")
	}
	return
}

// listSandboxCommand implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) listSandboxCommand(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sandboxes/{sandboxId}/commands", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSandboxCommandResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "ListSandboxCommand")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "ListSandboxCommand", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSandboxSnapshots Lists sandbox snapshot resources with optional filtering and pagination.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/ListSandboxSnapshots.go.html to see an example of how to use ListSandboxSnapshots API.
// A default retry strategy applies to this operation ListSandboxSnapshots()
func (client GenerativeAiSandboxClient) ListSandboxSnapshots(ctx context.Context, request ListSandboxSnapshotsRequest) (response ListSandboxSnapshotsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSandboxSnapshots, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSandboxSnapshotsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSandboxSnapshotsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSandboxSnapshotsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSandboxSnapshotsResponse")
	}
	return
}

// listSandboxSnapshots implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) listSandboxSnapshots(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/snapshots", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSandboxSnapshotsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "ListSandboxSnapshots")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "ListSandboxSnapshots", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSandboxes Lists sandbox resources with optional filtering and pagination. Specify exactly one of
// `projectId` or `compartmentId`. A compartment-scoped request returns sandboxes from
// Generative AI projects directly contained in that compartment; child compartments are
// not included.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/ListSandboxes.go.html to see an example of how to use ListSandboxes API.
// A default retry strategy applies to this operation ListSandboxes()
func (client GenerativeAiSandboxClient) ListSandboxes(ctx context.Context, request ListSandboxesRequest) (response ListSandboxesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSandboxes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSandboxesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSandboxesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSandboxesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSandboxesResponse")
	}
	return
}

// listSandboxes implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) listSandboxes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sandboxes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSandboxesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "ListSandboxes")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "ListSandboxes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ReadSandboxFile Reads a file from a sandbox.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/ReadSandboxFile.go.html to see an example of how to use ReadSandboxFile API.
// A default retry strategy applies to this operation ReadSandboxFile()
func (client GenerativeAiSandboxClient) ReadSandboxFile(ctx context.Context, request ReadSandboxFileRequest) (response ReadSandboxFileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.readSandboxFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReadSandboxFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReadSandboxFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReadSandboxFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReadSandboxFileResponse")
	}
	return
}

// readSandboxFile implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) readSandboxFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sandboxes/{sandboxId}/files/read", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ReadSandboxFileResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "ReadSandboxFile")
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "ReadSandboxFile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResumeSandbox Resumes a suspended sandbox by scaling its backing runtime back to one replica.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/ResumeSandbox.go.html to see an example of how to use ResumeSandbox API.
// A default retry strategy applies to this operation ResumeSandbox()
func (client GenerativeAiSandboxClient) ResumeSandbox(ctx context.Context, request ResumeSandboxRequest) (response ResumeSandboxResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.resumeSandbox, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResumeSandboxResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResumeSandboxResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResumeSandboxResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResumeSandboxResponse")
	}
	return
}

// resumeSandbox implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) resumeSandbox(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sandboxes/{sandboxId}/actions/resume", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResumeSandboxResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "ResumeSandbox")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "ResumeSandbox", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RunSandboxCommand Runs a command in a sandbox and returns the command identifier immediately.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/RunSandboxCommand.go.html to see an example of how to use RunSandboxCommand API.
// A default retry strategy applies to this operation RunSandboxCommand()
func (client GenerativeAiSandboxClient) RunSandboxCommand(ctx context.Context, request RunSandboxCommandRequest) (response RunSandboxCommandResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.runSandboxCommand, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RunSandboxCommandResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RunSandboxCommandResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RunSandboxCommandResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RunSandboxCommandResponse")
	}
	return
}

// runSandboxCommand implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) runSandboxCommand(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sandboxes/{sandboxId}/commands", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RunSandboxCommandResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "RunSandboxCommand")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "RunSandboxCommand", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RunSandboxCommandAndStream Runs a command in a sandbox and streams command events as newline-delimited JSON. Each line is a serialized SandboxCommandEvent.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/RunSandboxCommandAndStream.go.html to see an example of how to use RunSandboxCommandAndStream API.
// A default retry strategy applies to this operation RunSandboxCommandAndStream()
func (client GenerativeAiSandboxClient) RunSandboxCommandAndStream(ctx context.Context, request RunSandboxCommandAndStreamRequest) (response RunSandboxCommandAndStreamResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.runSandboxCommandAndStream, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RunSandboxCommandAndStreamResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RunSandboxCommandAndStreamResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RunSandboxCommandAndStreamResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RunSandboxCommandAndStreamResponse")
	}
	return
}

// runSandboxCommandAndStream implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) runSandboxCommandAndStream(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sandboxes/{sandboxId}/commands/actions/runAndStream", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RunSandboxCommandAndStreamResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "RunSandboxCommandAndStream")
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "RunSandboxCommandAndStream", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RunSandboxCommandAndWait Runs a command in a sandbox, blocks until command execution finishes, and returns the final command metadata.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/RunSandboxCommandAndWait.go.html to see an example of how to use RunSandboxCommandAndWait API.
// A default retry strategy applies to this operation RunSandboxCommandAndWait()
func (client GenerativeAiSandboxClient) RunSandboxCommandAndWait(ctx context.Context, request RunSandboxCommandAndWaitRequest) (response RunSandboxCommandAndWaitResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.runSandboxCommandAndWait, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RunSandboxCommandAndWaitResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RunSandboxCommandAndWaitResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RunSandboxCommandAndWaitResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RunSandboxCommandAndWaitResponse")
	}
	return
}

// runSandboxCommandAndWait implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) runSandboxCommandAndWait(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sandboxes/{sandboxId}/commands/actions/runAndWait", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RunSandboxCommandAndWaitResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "RunSandboxCommandAndWait")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "RunSandboxCommandAndWait", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopSandbox Stops a running sandbox.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/StopSandbox.go.html to see an example of how to use StopSandbox API.
// A default retry strategy applies to this operation StopSandbox()
func (client GenerativeAiSandboxClient) StopSandbox(ctx context.Context, request StopSandboxRequest) (response StopSandboxResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopSandbox, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopSandboxResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopSandboxResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopSandboxResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopSandboxResponse")
	}
	return
}

// stopSandbox implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) stopSandbox(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sandboxes/{sandboxId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopSandboxResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "StopSandbox")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "StopSandbox", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StreamSandboxCommandEvents Streams command events for an existing sandbox command from `fromSeq` through `untilSeq` using Server-Sent Events. Each SSE `data` field contains a serialized SandboxCommandEvent.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/StreamSandboxCommandEvents.go.html to see an example of how to use StreamSandboxCommandEvents API.
// A default retry strategy applies to this operation StreamSandboxCommandEvents()
func (client GenerativeAiSandboxClient) StreamSandboxCommandEvents(ctx context.Context, request StreamSandboxCommandEventsRequest) (response StreamSandboxCommandEventsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.streamSandboxCommandEvents, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StreamSandboxCommandEventsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StreamSandboxCommandEventsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StreamSandboxCommandEventsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StreamSandboxCommandEventsResponse")
	}
	return
}

// streamSandboxCommandEvents implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) streamSandboxCommandEvents(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sandboxes/{sandboxId}/commands/{commandId}/events", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StreamSandboxCommandEventsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "StreamSandboxCommandEvents")
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "StreamSandboxCommandEvents", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SuspendSandbox Suspends a running sandbox by scaling its backing runtime to zero replicas while retaining the workspace block volume.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/SuspendSandbox.go.html to see an example of how to use SuspendSandbox API.
// A default retry strategy applies to this operation SuspendSandbox()
func (client GenerativeAiSandboxClient) SuspendSandbox(ctx context.Context, request SuspendSandboxRequest) (response SuspendSandboxResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.suspendSandbox, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SuspendSandboxResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SuspendSandboxResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SuspendSandboxResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SuspendSandboxResponse")
	}
	return
}

// suspendSandbox implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) suspendSandbox(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sandboxes/{sandboxId}/actions/suspend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SuspendSandboxResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "SuspendSandbox")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "SuspendSandbox", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSandbox Updates a sandbox.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/UpdateSandbox.go.html to see an example of how to use UpdateSandbox API.
// A default retry strategy applies to this operation UpdateSandbox()
func (client GenerativeAiSandboxClient) UpdateSandbox(ctx context.Context, request UpdateSandboxRequest) (response UpdateSandboxResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSandbox, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSandboxResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSandboxResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSandboxResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSandboxResponse")
	}
	return
}

// updateSandbox implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) updateSandbox(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/sandboxes/{sandboxId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSandboxResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "UpdateSandbox")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "UpdateSandbox", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// WriteSandboxFile Writes a file in a sandbox and returns the resulting file resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaisandbox/WriteSandboxFile.go.html to see an example of how to use WriteSandboxFile API.
// A default retry strategy applies to this operation WriteSandboxFile()
func (client GenerativeAiSandboxClient) WriteSandboxFile(ctx context.Context, request WriteSandboxFileRequest) (response WriteSandboxFileResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.writeSandboxFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = WriteSandboxFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = WriteSandboxFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(WriteSandboxFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into WriteSandboxFileResponse")
	}
	return
}

// writeSandboxFile implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiSandboxClient) writeSandboxFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sandboxes/{sandboxId}/files/write", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response WriteSandboxFileResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "generativeAiSandbox", "WriteSandboxFile")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GenerativeAiSandbox", "WriteSandboxFile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
