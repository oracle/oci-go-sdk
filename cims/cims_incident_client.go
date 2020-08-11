// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//IncidentClient a client for Incident
type IncidentClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewIncidentClientWithConfigurationProvider Creates a new default Incident client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewIncidentClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client IncidentClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newIncidentClientFromBaseClient(baseClient, configProvider)
}

// NewIncidentClientWithOboToken Creates a new default Incident client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewIncidentClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client IncidentClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newIncidentClientFromBaseClient(baseClient, configProvider)
}

func newIncidentClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client IncidentClient, err error) {
	client = IncidentClient{BaseClient: baseClient}
	client.BasePath = "20181231"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *IncidentClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("cims", "https://incidentmanagement.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *IncidentClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *IncidentClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateIncident Enables the customer to create an support ticket.
func (client IncidentClient) CreateIncident(ctx context.Context, request CreateIncidentRequest) (response CreateIncidentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createIncident, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateIncidentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateIncidentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateIncidentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateIncidentResponse")
	}
	return
}

// createIncident implements the OCIOperation interface (enables retrying operations)
func (client IncidentClient) createIncident(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/v2/incidents")
	if err != nil {
		return nil, err
	}

	var response CreateIncidentResponse
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

// GetIncident Gets the details of the support ticket.
func (client IncidentClient) GetIncident(ctx context.Context, request GetIncidentRequest) (response GetIncidentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getIncident, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetIncidentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetIncidentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetIncidentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetIncidentResponse")
	}
	return
}

// getIncident implements the OCIOperation interface (enables retrying operations)
func (client IncidentClient) getIncident(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/v2/incidents/{incidentKey}")
	if err != nil {
		return nil, err
	}

	var response GetIncidentResponse
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

// GetStatus Gets the status of the service.
func (client IncidentClient) GetStatus(ctx context.Context, request GetStatusRequest) (response GetStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStatusResponse")
	}
	return
}

// getStatus implements the OCIOperation interface (enables retrying operations)
func (client IncidentClient) getStatus(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/v2/incidents/status/{source}")
	if err != nil {
		return nil, err
	}

	var response GetStatusResponse
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

// ListIncidentResourceTypes During support ticket creation, returns the list of all possible products that Oracle Cloud Infrastructure supports.
func (client IncidentClient) ListIncidentResourceTypes(ctx context.Context, request ListIncidentResourceTypesRequest) (response ListIncidentResourceTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listIncidentResourceTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListIncidentResourceTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListIncidentResourceTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIncidentResourceTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIncidentResourceTypesResponse")
	}
	return
}

// listIncidentResourceTypes implements the OCIOperation interface (enables retrying operations)
func (client IncidentClient) listIncidentResourceTypes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/v2/incidents/incidentResourceTypes")
	if err != nil {
		return nil, err
	}

	var response ListIncidentResourceTypesResponse
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

// ListIncidents Returns the list of support tickets raised by the tenancy.
func (client IncidentClient) ListIncidents(ctx context.Context, request ListIncidentsRequest) (response ListIncidentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listIncidents, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListIncidentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListIncidentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIncidentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIncidentsResponse")
	}
	return
}

// listIncidents implements the OCIOperation interface (enables retrying operations)
func (client IncidentClient) listIncidents(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/v2/incidents")
	if err != nil {
		return nil, err
	}

	var response ListIncidentsResponse
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

// UpdateIncident Updates the specified support ticket's information.
func (client IncidentClient) UpdateIncident(ctx context.Context, request UpdateIncidentRequest) (response UpdateIncidentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateIncident, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateIncidentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateIncidentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateIncidentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateIncidentResponse")
	}
	return
}

// updateIncident implements the OCIOperation interface (enables retrying operations)
func (client IncidentClient) updateIncident(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/v2/incidents/{incidentKey}")
	if err != nil {
		return nil, err
	}

	var response UpdateIncidentResponse
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

// ValidateUser Checks whether the requested user is valid.
func (client IncidentClient) ValidateUser(ctx context.Context, request ValidateUserRequest) (response ValidateUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.validateUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateUserResponse")
	}
	return
}

// validateUser implements the OCIOperation interface (enables retrying operations)
func (client IncidentClient) validateUser(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/v2/incidents/user/validate")
	if err != nil {
		return nil, err
	}

	var response ValidateUserResponse
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
