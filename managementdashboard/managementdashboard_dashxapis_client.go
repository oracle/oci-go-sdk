// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ManagementDashboard API
//
// Management Dashboard micro-service provides a set of CRUD, import, export, and compartment related APIs (such as change compartment)   to support dashboard and saved search metadata preservation.  These APIs are mainly for client UIs, for various UI activities such as get list of all saved searches in a compartment, create a dashboard, open a saved search, etc.  Use export to retrieve  dashboards and their saved searches, then edit the Json if necessary (for example change compartmentIds), then import the result to  destination dashboard service.
// APIs validate all required properties to ensure properties are present and have correct type and values.
//
//

package managementdashboard

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//DashxApisClient a client for DashxApis
type DashxApisClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDashxApisClientWithConfigurationProvider Creates a new default DashxApis client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDashxApisClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DashxApisClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newDashxApisClientFromBaseClient(baseClient, configProvider)
}

// NewDashxApisClientWithOboToken Creates a new default DashxApis client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewDashxApisClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DashxApisClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newDashxApisClientFromBaseClient(baseClient, configProvider)
}

func newDashxApisClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DashxApisClient, err error) {
	client = DashxApisClient{BaseClient: baseClient}
	client.BasePath = "20200901"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DashxApisClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("managementdashboard", "https://managementdashboards.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DashxApisClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DashxApisClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeManagementDashboardsCompartment Move the dashboard from existing compartment to a new compartment.
func (client DashxApisClient) ChangeManagementDashboardsCompartment(ctx context.Context, request ChangeManagementDashboardsCompartmentRequest) (response ChangeManagementDashboardsCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeManagementDashboardsCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeManagementDashboardsCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeManagementDashboardsCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeManagementDashboardsCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeManagementDashboardsCompartmentResponse")
	}
	return
}

// changeManagementDashboardsCompartment implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) changeManagementDashboardsCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managementDashboards/{managementDashboardId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeManagementDashboardsCompartmentResponse
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

// ChangeManagementSavedSearchesCompartment Move the saved search from existing compartment to a new compartment.
func (client DashxApisClient) ChangeManagementSavedSearchesCompartment(ctx context.Context, request ChangeManagementSavedSearchesCompartmentRequest) (response ChangeManagementSavedSearchesCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeManagementSavedSearchesCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeManagementSavedSearchesCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeManagementSavedSearchesCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeManagementSavedSearchesCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeManagementSavedSearchesCompartmentResponse")
	}
	return
}

// changeManagementSavedSearchesCompartment implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) changeManagementSavedSearchesCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managementSavedSearches/{managementSavedSearchId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeManagementSavedSearchesCompartmentResponse
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

// CreateManagementDashboard Creates a new dashboard.  Limit for number of saved searches in a dashboard is 20.
func (client DashxApisClient) CreateManagementDashboard(ctx context.Context, request CreateManagementDashboardRequest) (response CreateManagementDashboardResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createManagementDashboard, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateManagementDashboardResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateManagementDashboardResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateManagementDashboardResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateManagementDashboardResponse")
	}
	return
}

// createManagementDashboard implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) createManagementDashboard(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managementDashboards")
	if err != nil {
		return nil, err
	}

	var response CreateManagementDashboardResponse
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

// CreateManagementSavedSearch Creates a new saved search.
func (client DashxApisClient) CreateManagementSavedSearch(ctx context.Context, request CreateManagementSavedSearchRequest) (response CreateManagementSavedSearchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createManagementSavedSearch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateManagementSavedSearchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateManagementSavedSearchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateManagementSavedSearchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateManagementSavedSearchResponse")
	}
	return
}

// createManagementSavedSearch implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) createManagementSavedSearch(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managementSavedSearches")
	if err != nil {
		return nil, err
	}

	var response CreateManagementSavedSearchResponse
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

// DeleteManagementDashboard Deletes a Dashboard by id.
func (client DashxApisClient) DeleteManagementDashboard(ctx context.Context, request DeleteManagementDashboardRequest) (response DeleteManagementDashboardResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteManagementDashboard, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteManagementDashboardResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteManagementDashboardResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteManagementDashboardResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteManagementDashboardResponse")
	}
	return
}

// deleteManagementDashboard implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) deleteManagementDashboard(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managementDashboards/{managementDashboardId}")
	if err != nil {
		return nil, err
	}

	var response DeleteManagementDashboardResponse
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

// DeleteManagementSavedSearch Deletes a saved search by Id
func (client DashxApisClient) DeleteManagementSavedSearch(ctx context.Context, request DeleteManagementSavedSearchRequest) (response DeleteManagementSavedSearchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteManagementSavedSearch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteManagementSavedSearchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteManagementSavedSearchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteManagementSavedSearchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteManagementSavedSearchResponse")
	}
	return
}

// deleteManagementSavedSearch implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) deleteManagementSavedSearch(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managementSavedSearches/{managementSavedSearchId}")
	if err != nil {
		return nil, err
	}

	var response DeleteManagementSavedSearchResponse
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

// ExportDashboard Exports an array of dashboards and their saved searches.
func (client DashxApisClient) ExportDashboard(ctx context.Context, request ExportDashboardRequest) (response ExportDashboardResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.exportDashboard, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportDashboardResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportDashboardResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportDashboardResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportDashboardResponse")
	}
	return
}

// exportDashboard implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) exportDashboard(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementDashboards/actions/exportDashboard/{exportDashboardId}")
	if err != nil {
		return nil, err
	}

	var response ExportDashboardResponse
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

// GetManagementDashboard Get a Dashboard and its saved searches by id.  Deleted or unauthorized saved searches are marked by tile's state property.
func (client DashxApisClient) GetManagementDashboard(ctx context.Context, request GetManagementDashboardRequest) (response GetManagementDashboardResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.getManagementDashboard, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagementDashboardResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagementDashboardResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagementDashboardResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagementDashboardResponse")
	}
	return
}

// getManagementDashboard implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) getManagementDashboard(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementDashboards/{managementDashboardId}")
	if err != nil {
		return nil, err
	}

	var response GetManagementDashboardResponse
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

// GetManagementSavedSearch Get a saved search by Id.
func (client DashxApisClient) GetManagementSavedSearch(ctx context.Context, request GetManagementSavedSearchRequest) (response GetManagementSavedSearchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.getManagementSavedSearch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagementSavedSearchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagementSavedSearchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagementSavedSearchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagementSavedSearchResponse")
	}
	return
}

// getManagementSavedSearch implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) getManagementSavedSearch(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementSavedSearches/{managementSavedSearchId}")
	if err != nil {
		return nil, err
	}

	var response GetManagementSavedSearchResponse
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

// ImportDashboard Import an array of dashboards and their saved searches.
func (client DashxApisClient) ImportDashboard(ctx context.Context, request ImportDashboardRequest) (response ImportDashboardResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.importDashboard, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportDashboardResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportDashboardResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportDashboardResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportDashboardResponse")
	}
	return
}

// importDashboard implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) importDashboard(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managementDashboards/actions/importDashboard")
	if err != nil {
		return nil, err
	}

	var response ImportDashboardResponse
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

// ListManagementDashboards Gets list of dashboards and their saved searches for compartment with pagination.  Returned properties are a summary.
func (client DashxApisClient) ListManagementDashboards(ctx context.Context, request ListManagementDashboardsRequest) (response ListManagementDashboardsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagementDashboards, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagementDashboardsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagementDashboardsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagementDashboardsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagementDashboardsResponse")
	}
	return
}

// listManagementDashboards implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) listManagementDashboards(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementDashboards")
	if err != nil {
		return nil, err
	}

	var response ListManagementDashboardsResponse
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

// ListManagementSavedSearches Gets list of saved searches with pagination.  Returned properties are a summary.
func (client DashxApisClient) ListManagementSavedSearches(ctx context.Context, request ListManagementSavedSearchesRequest) (response ListManagementSavedSearchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagementSavedSearches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagementSavedSearchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagementSavedSearchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagementSavedSearchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagementSavedSearchesResponse")
	}
	return
}

// listManagementSavedSearches implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) listManagementSavedSearches(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementSavedSearches")
	if err != nil {
		return nil, err
	}

	var response ListManagementSavedSearchesResponse
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

// UpdateManagementDashboard Updates an existing dashboard identified by id path parameter.  Limit for number of saved searches in a dashboard is 20.
func (client DashxApisClient) UpdateManagementDashboard(ctx context.Context, request UpdateManagementDashboardRequest) (response UpdateManagementDashboardResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateManagementDashboard, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateManagementDashboardResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateManagementDashboardResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateManagementDashboardResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateManagementDashboardResponse")
	}
	return
}

// updateManagementDashboard implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) updateManagementDashboard(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managementDashboards/{managementDashboardId}")
	if err != nil {
		return nil, err
	}

	var response UpdateManagementDashboardResponse
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

// UpdateManagementSavedSearch Update an existing saved search.  Id cannot be updated.
func (client DashxApisClient) UpdateManagementSavedSearch(ctx context.Context, request UpdateManagementSavedSearchRequest) (response UpdateManagementSavedSearchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateManagementSavedSearch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateManagementSavedSearchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateManagementSavedSearchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateManagementSavedSearchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateManagementSavedSearchResponse")
	}
	return
}

// updateManagementSavedSearch implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) updateManagementSavedSearch(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managementSavedSearches/{managementSavedSearchId}")
	if err != nil {
		return nil, err
	}

	var response UpdateManagementSavedSearchResponse
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
