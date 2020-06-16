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

//ApplianceExportJobClient a client for ApplianceExportJob
type ApplianceExportJobClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewApplianceExportJobClientWithConfigurationProvider Creates a new default ApplianceExportJob client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewApplianceExportJobClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ApplianceExportJobClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newApplianceExportJobClientFromBaseClient(baseClient, configProvider)
}

// NewApplianceExportJobClientWithOboToken Creates a new default ApplianceExportJob client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewApplianceExportJobClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ApplianceExportJobClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newApplianceExportJobClientFromBaseClient(baseClient, configProvider)
}

func newApplianceExportJobClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ApplianceExportJobClient, err error) {
	client = ApplianceExportJobClient{BaseClient: baseClient}
	client.BasePath = "20171001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ApplianceExportJobClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dts", "https://datatransfer.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ApplianceExportJobClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ApplianceExportJobClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeApplianceExportJobCompartment Moves a ApplianceExportJob into a different compartment.
func (client ApplianceExportJobClient) ChangeApplianceExportJobCompartment(ctx context.Context, request ChangeApplianceExportJobCompartmentRequest) (response ChangeApplianceExportJobCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeApplianceExportJobCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeApplianceExportJobCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeApplianceExportJobCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeApplianceExportJobCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeApplianceExportJobCompartmentResponse")
	}
	return
}

// changeApplianceExportJobCompartment implements the OCIOperation interface (enables retrying operations)
func (client ApplianceExportJobClient) changeApplianceExportJobCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/applianceExportJobs/{applianceExportJobId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeApplianceExportJobCompartmentResponse
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

// CreateApplianceExportJob Creates a new Appliance Export Job that corresponds with customer's logical dataset
func (client ApplianceExportJobClient) CreateApplianceExportJob(ctx context.Context, request CreateApplianceExportJobRequest) (response CreateApplianceExportJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createApplianceExportJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateApplianceExportJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateApplianceExportJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateApplianceExportJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateApplianceExportJobResponse")
	}
	return
}

// createApplianceExportJob implements the OCIOperation interface (enables retrying operations)
func (client ApplianceExportJobClient) createApplianceExportJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/applianceExportJobs")
	if err != nil {
		return nil, err
	}

	var response CreateApplianceExportJobResponse
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

// DeleteApplianceExportJob deletes a Appliance Export Job
func (client ApplianceExportJobClient) DeleteApplianceExportJob(ctx context.Context, request DeleteApplianceExportJobRequest) (response DeleteApplianceExportJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteApplianceExportJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteApplianceExportJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteApplianceExportJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteApplianceExportJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteApplianceExportJobResponse")
	}
	return
}

// deleteApplianceExportJob implements the OCIOperation interface (enables retrying operations)
func (client ApplianceExportJobClient) deleteApplianceExportJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/applianceExportJobs/{applianceExportJobId}")
	if err != nil {
		return nil, err
	}

	var response DeleteApplianceExportJobResponse
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

// GetApplianceExportJob Describes a Appliance Export Job in detail
func (client ApplianceExportJobClient) GetApplianceExportJob(ctx context.Context, request GetApplianceExportJobRequest) (response GetApplianceExportJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getApplianceExportJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetApplianceExportJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetApplianceExportJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetApplianceExportJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetApplianceExportJobResponse")
	}
	return
}

// getApplianceExportJob implements the OCIOperation interface (enables retrying operations)
func (client ApplianceExportJobClient) getApplianceExportJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/applianceExportJobs/{applianceExportJobId}")
	if err != nil {
		return nil, err
	}

	var response GetApplianceExportJobResponse
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

// ListApplianceExportJobs Lists Appliance Export Jobs in a given compartment
func (client ApplianceExportJobClient) ListApplianceExportJobs(ctx context.Context, request ListApplianceExportJobsRequest) (response ListApplianceExportJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listApplianceExportJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListApplianceExportJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListApplianceExportJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListApplianceExportJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListApplianceExportJobsResponse")
	}
	return
}

// listApplianceExportJobs implements the OCIOperation interface (enables retrying operations)
func (client ApplianceExportJobClient) listApplianceExportJobs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/applianceExportJobs")
	if err != nil {
		return nil, err
	}

	var response ListApplianceExportJobsResponse
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

// UpdateApplianceExportJob Updates a Appliance Export Job that corresponds with customer's logical dataset.
func (client ApplianceExportJobClient) UpdateApplianceExportJob(ctx context.Context, request UpdateApplianceExportJobRequest) (response UpdateApplianceExportJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateApplianceExportJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateApplianceExportJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateApplianceExportJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateApplianceExportJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateApplianceExportJobResponse")
	}
	return
}

// updateApplianceExportJob implements the OCIOperation interface (enables retrying operations)
func (client ApplianceExportJobClient) updateApplianceExportJob(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/applianceExportJobs/{applianceExportJobId}")
	if err != nil {
		return nil, err
	}

	var response UpdateApplianceExportJobResponse
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
