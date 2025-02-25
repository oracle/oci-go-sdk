// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dashboards API
//
// Use the Oracle Cloud Infrastructure Dashboards service API to manage dashboards in the Console.
// Dashboards provide an organized and customizable view of resources and their metrics in the Console.
// For more information, see Dashboards (https://docs.oracle.com/iaas/Content/Dashboards/home.htm).
// **Important:** Resources for the Dashboards service are created in the tenacy's home region.
// Although it is possible to create dashboard and dashboard group resources in regions other than the home region,
// you won't be able to view those resources in the Console.
// Therefore, creating resources outside of the home region is not recommended.
//

package dashboardservice

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DashboardClient a client for Dashboard
type DashboardClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDashboardClientWithConfigurationProvider Creates a new default Dashboard client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDashboardClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DashboardClient, err error) {
	if enabled := common.CheckForEnabledServices("dashboardservice"); !enabled {
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
	return newDashboardClientFromBaseClient(baseClient, provider)
}

// NewDashboardClientWithOboToken Creates a new default Dashboard client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDashboardClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DashboardClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDashboardClientFromBaseClient(baseClient, configProvider)
}

func newDashboardClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DashboardClient, err error) {
	// Dashboard service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Dashboard"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DashboardClient{BaseClient: baseClient}
	client.BasePath = "20210731"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DashboardClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dashboardservice", "https://dashboard.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DashboardClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DashboardClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeDashboardGroup Moves a Dashboard resource from one dashboardGroup identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dashboardservice/ChangeDashboardGroup.go.html to see an example of how to use ChangeDashboardGroup API.
func (client DashboardClient) ChangeDashboardGroup(ctx context.Context, request ChangeDashboardGroupRequest) (response ChangeDashboardGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeDashboardGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDashboardGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDashboardGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDashboardGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDashboardGroupResponse")
	}
	return
}

// changeDashboardGroup implements the OCIOperation interface (enables retrying operations)
func (client DashboardClient) changeDashboardGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dashboards/{dashboardId}/actions/changeDashboardGroup", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDashboardGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dashboard/20210731/Dashboard/ChangeDashboardGroup"
		err = common.PostProcessServiceError(err, "Dashboard", "ChangeDashboardGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDashboard Creates a new dashboard in the dashboard group's compartment using the details provided in request body.
// **Caution:** Resources for the Dashboard service are created in the tenacy's home region.
// Although it is possible to create dashboard resource in regions other than the home region,
// you won't be able to view those resources in the Console.
// Therefore, creating resources outside of the home region is not recommended.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dashboardservice/CreateDashboard.go.html to see an example of how to use CreateDashboard API.
func (client DashboardClient) CreateDashboard(ctx context.Context, request CreateDashboardRequest) (response CreateDashboardResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDashboard, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDashboardResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDashboardResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDashboardResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDashboardResponse")
	}
	return
}

// createDashboard implements the OCIOperation interface (enables retrying operations)
func (client DashboardClient) createDashboard(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dashboards", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDashboardResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dashboard/20210731/Dashboard/CreateDashboard"
		err = common.PostProcessServiceError(err, "Dashboard", "CreateDashboard", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &dashboard{})
	return response, err
}

// DeleteDashboard Deletes the specified dashboard. Uses the dashboard's OCID to determine which dashboard to delete.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dashboardservice/DeleteDashboard.go.html to see an example of how to use DeleteDashboard API.
func (client DashboardClient) DeleteDashboard(ctx context.Context, request DeleteDashboardRequest) (response DeleteDashboardResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDashboard, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDashboardResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDashboardResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDashboardResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDashboardResponse")
	}
	return
}

// deleteDashboard implements the OCIOperation interface (enables retrying operations)
func (client DashboardClient) deleteDashboard(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dashboards/{dashboardId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDashboardResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dashboard/20210731/Dashboard/DeleteDashboard"
		err = common.PostProcessServiceError(err, "Dashboard", "DeleteDashboard", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDashboard Gets the specified dashboard's information. Uses the dashboard's OCID to determine which dashboard to retrieve.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dashboardservice/GetDashboard.go.html to see an example of how to use GetDashboard API.
// A default retry strategy applies to this operation GetDashboard()
func (client DashboardClient) GetDashboard(ctx context.Context, request GetDashboardRequest) (response GetDashboardResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDashboard, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDashboardResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDashboardResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDashboardResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDashboardResponse")
	}
	return
}

// getDashboard implements the OCIOperation interface (enables retrying operations)
func (client DashboardClient) getDashboard(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dashboards/{dashboardId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDashboardResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dashboard/20210731/Dashboard/GetDashboard"
		err = common.PostProcessServiceError(err, "Dashboard", "GetDashboard", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &dashboard{})
	return response, err
}

// ListDashboards Returns a list of dashboards with a specific dashboard group ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dashboardservice/ListDashboards.go.html to see an example of how to use ListDashboards API.
// A default retry strategy applies to this operation ListDashboards()
func (client DashboardClient) ListDashboards(ctx context.Context, request ListDashboardsRequest) (response ListDashboardsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDashboards, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDashboardsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDashboardsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDashboardsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDashboardsResponse")
	}
	return
}

// listDashboards implements the OCIOperation interface (enables retrying operations)
func (client DashboardClient) listDashboards(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dashboards", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDashboardsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dashboard/20210731/DashboardCollection/ListDashboards"
		err = common.PostProcessServiceError(err, "Dashboard", "ListDashboards", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDashboard Updates the specified dashboard. Uses the dashboard's OCID to determine which dashboard to update.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dashboardservice/UpdateDashboard.go.html to see an example of how to use UpdateDashboard API.
func (client DashboardClient) UpdateDashboard(ctx context.Context, request UpdateDashboardRequest) (response UpdateDashboardResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDashboard, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDashboardResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDashboardResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDashboardResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDashboardResponse")
	}
	return
}

// updateDashboard implements the OCIOperation interface (enables retrying operations)
func (client DashboardClient) updateDashboard(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dashboards/{dashboardId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDashboardResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dashboard/20210731/Dashboard/UpdateDashboard"
		err = common.PostProcessServiceError(err, "Dashboard", "UpdateDashboard", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &dashboard{})
	return response, err
}
