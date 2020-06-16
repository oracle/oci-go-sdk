// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//TransferJobClient a client for TransferJob
type TransferJobClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewTransferJobClientWithConfigurationProvider Creates a new default TransferJob client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewTransferJobClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client TransferJobClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newTransferJobClientFromBaseClient(baseClient, configProvider)
}

// NewTransferJobClientWithOboToken Creates a new default TransferJob client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewTransferJobClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client TransferJobClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newTransferJobClientFromBaseClient(baseClient, configProvider)
}

func newTransferJobClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client TransferJobClient, err error) {
	client = TransferJobClient{BaseClient: baseClient}
	client.BasePath = "20171001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *TransferJobClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dts", "https://datatransfer.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *TransferJobClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *TransferJobClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeTransferJobCompartment Moves a TransferJob into a different compartment.
func (client TransferJobClient) ChangeTransferJobCompartment(ctx context.Context, request ChangeTransferJobCompartmentRequest) (response ChangeTransferJobCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeTransferJobCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeTransferJobCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeTransferJobCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeTransferJobCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeTransferJobCompartmentResponse")
	}
	return
}

// changeTransferJobCompartment implements the OCIOperation interface (enables retrying operations)
func (client TransferJobClient) changeTransferJobCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transferJobs/{transferJobId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeTransferJobCompartmentResponse
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

// CreateTransferJob Create a new Transfer Job that corresponds with customer's logical dataset e.g. a DB or a filesystem.
func (client TransferJobClient) CreateTransferJob(ctx context.Context, request CreateTransferJobRequest) (response CreateTransferJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTransferJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTransferJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTransferJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTransferJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTransferJobResponse")
	}
	return
}

// createTransferJob implements the OCIOperation interface (enables retrying operations)
func (client TransferJobClient) createTransferJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transferJobs")
	if err != nil {
		return nil, err
	}

	var response CreateTransferJobResponse
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

// DeleteTransferJob deletes a transfer job
func (client TransferJobClient) DeleteTransferJob(ctx context.Context, request DeleteTransferJobRequest) (response DeleteTransferJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTransferJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTransferJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTransferJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTransferJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTransferJobResponse")
	}
	return
}

// deleteTransferJob implements the OCIOperation interface (enables retrying operations)
func (client TransferJobClient) deleteTransferJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/transferJobs/{id}")
	if err != nil {
		return nil, err
	}

	var response DeleteTransferJobResponse
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

// GetTransferJob Describes a transfer job in detail
func (client TransferJobClient) GetTransferJob(ctx context.Context, request GetTransferJobRequest) (response GetTransferJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTransferJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTransferJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTransferJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTransferJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTransferJobResponse")
	}
	return
}

// getTransferJob implements the OCIOperation interface (enables retrying operations)
func (client TransferJobClient) getTransferJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferJobs/{id}")
	if err != nil {
		return nil, err
	}

	var response GetTransferJobResponse
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

// ListTransferJobs Lists Transfer Jobs in a given compartment
func (client TransferJobClient) ListTransferJobs(ctx context.Context, request ListTransferJobsRequest) (response ListTransferJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTransferJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTransferJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTransferJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTransferJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTransferJobsResponse")
	}
	return
}

// listTransferJobs implements the OCIOperation interface (enables retrying operations)
func (client TransferJobClient) listTransferJobs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferJobs")
	if err != nil {
		return nil, err
	}

	var response ListTransferJobsResponse
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

// UpdateTransferJob Updates a Transfer Job that corresponds with customer's logical dataset e.g. a DB or a filesystem.
func (client TransferJobClient) UpdateTransferJob(ctx context.Context, request UpdateTransferJobRequest) (response UpdateTransferJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTransferJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTransferJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTransferJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTransferJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTransferJobResponse")
	}
	return
}

// updateTransferJob implements the OCIOperation interface (enables retrying operations)
func (client TransferJobClient) updateTransferJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/transferJobs/{id}")
	if err != nil {
		return nil, err
	}

	var response UpdateTransferJobResponse
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
