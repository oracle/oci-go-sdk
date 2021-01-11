// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// InstanceAgentService API
//
// Instance Agent Service API
//

package computeinstanceagent

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v32/common"
	"github.com/oracle/oci-go-sdk/v32/common/auth"
	"net/http"
)

//ComputeInstanceAgentClient a client for ComputeInstanceAgent
type ComputeInstanceAgentClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewComputeInstanceAgentClientWithConfigurationProvider Creates a new default ComputeInstanceAgent client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewComputeInstanceAgentClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ComputeInstanceAgentClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newComputeInstanceAgentClientFromBaseClient(baseClient, provider)
}

// NewComputeInstanceAgentClientWithOboToken Creates a new default ComputeInstanceAgent client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewComputeInstanceAgentClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ComputeInstanceAgentClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newComputeInstanceAgentClientFromBaseClient(baseClient, configProvider)
}

func newComputeInstanceAgentClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ComputeInstanceAgentClient, err error) {
	client = ComputeInstanceAgentClient{BaseClient: baseClient}
	client.BasePath = "20180530"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ComputeInstanceAgentClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("computeinstanceagent", "https://iaas.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ComputeInstanceAgentClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ComputeInstanceAgentClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelInstanceAgentCommand Cancel a command. Cancel is best effort attempt. If the commmand has already completed it will skip cancel.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computeinstanceagent/CancelInstanceAgentCommand.go.html to see an example of how to use CancelInstanceAgentCommand API.
func (client ComputeInstanceAgentClient) CancelInstanceAgentCommand(ctx context.Context, request CancelInstanceAgentCommandRequest) (response CancelInstanceAgentCommandResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelInstanceAgentCommand, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelInstanceAgentCommandResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelInstanceAgentCommandResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelInstanceAgentCommandResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelInstanceAgentCommandResponse")
	}
	return
}

// cancelInstanceAgentCommand implements the OCIOperation interface (enables retrying operations)
func (client ComputeInstanceAgentClient) cancelInstanceAgentCommand(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/instanceAgentCommands/{instanceAgentCommandId}")
	if err != nil {
		return nil, err
	}

	var response CancelInstanceAgentCommandResponse
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

// CreateInstanceAgentCommand Create command for one or more managed instances
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computeinstanceagent/CreateInstanceAgentCommand.go.html to see an example of how to use CreateInstanceAgentCommand API.
func (client ComputeInstanceAgentClient) CreateInstanceAgentCommand(ctx context.Context, request CreateInstanceAgentCommandRequest) (response CreateInstanceAgentCommandResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createInstanceAgentCommand, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateInstanceAgentCommandResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateInstanceAgentCommandResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateInstanceAgentCommandResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateInstanceAgentCommandResponse")
	}
	return
}

// createInstanceAgentCommand implements the OCIOperation interface (enables retrying operations)
func (client ComputeInstanceAgentClient) createInstanceAgentCommand(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/instanceAgentCommands")
	if err != nil {
		return nil, err
	}

	var response CreateInstanceAgentCommandResponse
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

// GetInstanceAgentCommand Gets information about the specified instance agent commandId.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computeinstanceagent/GetInstanceAgentCommand.go.html to see an example of how to use GetInstanceAgentCommand API.
func (client ComputeInstanceAgentClient) GetInstanceAgentCommand(ctx context.Context, request GetInstanceAgentCommandRequest) (response GetInstanceAgentCommandResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInstanceAgentCommand, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInstanceAgentCommandResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInstanceAgentCommandResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInstanceAgentCommandResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInstanceAgentCommandResponse")
	}
	return
}

// getInstanceAgentCommand implements the OCIOperation interface (enables retrying operations)
func (client ComputeInstanceAgentClient) getInstanceAgentCommand(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instanceAgentCommands/{instanceAgentCommandId}")
	if err != nil {
		return nil, err
	}

	var response GetInstanceAgentCommandResponse
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

// GetInstanceAgentCommandExecution Gets information about the status of specified instance agent commandId for the given instanceId.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computeinstanceagent/GetInstanceAgentCommandExecution.go.html to see an example of how to use GetInstanceAgentCommandExecution API.
func (client ComputeInstanceAgentClient) GetInstanceAgentCommandExecution(ctx context.Context, request GetInstanceAgentCommandExecutionRequest) (response GetInstanceAgentCommandExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInstanceAgentCommandExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInstanceAgentCommandExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInstanceAgentCommandExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInstanceAgentCommandExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInstanceAgentCommandExecutionResponse")
	}
	return
}

// getInstanceAgentCommandExecution implements the OCIOperation interface (enables retrying operations)
func (client ComputeInstanceAgentClient) getInstanceAgentCommandExecution(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instanceAgentCommands/{instanceAgentCommandId}/status")
	if err != nil {
		return nil, err
	}

	var response GetInstanceAgentCommandExecutionResponse
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

// ListInstanceAgentCommandExecutions List all executions of a command, i.e return command execution results from all targeted instances batch by
// batch.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computeinstanceagent/ListInstanceAgentCommandExecutions.go.html to see an example of how to use ListInstanceAgentCommandExecutions API.
func (client ComputeInstanceAgentClient) ListInstanceAgentCommandExecutions(ctx context.Context, request ListInstanceAgentCommandExecutionsRequest) (response ListInstanceAgentCommandExecutionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInstanceAgentCommandExecutions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInstanceAgentCommandExecutionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInstanceAgentCommandExecutionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInstanceAgentCommandExecutionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInstanceAgentCommandExecutionsResponse")
	}
	return
}

// listInstanceAgentCommandExecutions implements the OCIOperation interface (enables retrying operations)
func (client ComputeInstanceAgentClient) listInstanceAgentCommandExecutions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instanceAgentCommandExecutions")
	if err != nil {
		return nil, err
	}

	var response ListInstanceAgentCommandExecutionsResponse
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

// ListInstanceAgentCommands List Instance agent commands issued with the specified filter.
// Additonally you can filter commands sent to a particular InstanceId
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computeinstanceagent/ListInstanceAgentCommands.go.html to see an example of how to use ListInstanceAgentCommands API.
func (client ComputeInstanceAgentClient) ListInstanceAgentCommands(ctx context.Context, request ListInstanceAgentCommandsRequest) (response ListInstanceAgentCommandsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInstanceAgentCommands, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInstanceAgentCommandsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInstanceAgentCommandsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInstanceAgentCommandsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInstanceAgentCommandsResponse")
	}
	return
}

// listInstanceAgentCommands implements the OCIOperation interface (enables retrying operations)
func (client ComputeInstanceAgentClient) listInstanceAgentCommands(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instanceAgentCommands")
	if err != nil {
		return nil, err
	}

	var response ListInstanceAgentCommandsResponse
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
