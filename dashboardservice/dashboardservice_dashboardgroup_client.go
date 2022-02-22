// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dashboards API
//
// Use the Oracle Cloud Infrastructure Dashboards service API to manage dashboards in the Console.
// Dashboards provide an organized and customizable view of resources and their metrics in the Console.
// For more information, see Dashboards (https://docs.cloud.oracle.com/Content/Dashboards/home.htm).
// **Important:** Resources for the Dashboards service are created in the tenacy's home region.
// Although it is possible to create dashboard and dashboard group resources in regions other than the home region,
// you won't be able to view those resources in the Console.
// Therefore, creating resources outside of the home region is not recommended.
//

package dashboardservice

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"github.com/oracle/oci-go-sdk/v59/common/auth"
	"net/http"
)

//DashboardGroupClient a client for DashboardGroup
type DashboardGroupClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDashboardGroupClientWithConfigurationProvider Creates a new default DashboardGroup client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDashboardGroupClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DashboardGroupClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newDashboardGroupClientFromBaseClient(baseClient, provider)
}

// NewDashboardGroupClientWithOboToken Creates a new default DashboardGroup client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewDashboardGroupClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DashboardGroupClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDashboardGroupClientFromBaseClient(baseClient, configProvider)
}

func newDashboardGroupClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DashboardGroupClient, err error) {
	// DashboardGroup service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DashboardGroupClient{BaseClient: baseClient}
	client.BasePath = "20210731"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DashboardGroupClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dashboardservice", "https://dashboard.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DashboardGroupClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DashboardGroupClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateDashboardGroup Creates a new dashboard group using the details provided in request body.
// **Caution:** Resources for the Dashboard service are created in the tenacy's home region.
// Although it is possible to create dashboard group resource in regions other than the home region,
// you won't be able to view those resources in the Console.
// Therefore, creating resources outside of the home region is not recommended.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dashboardservice/CreateDashboardGroup.go.html to see an example of how to use CreateDashboardGroup API.
func (client DashboardGroupClient) CreateDashboardGroup(ctx context.Context, request CreateDashboardGroupRequest) (response CreateDashboardGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDashboardGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDashboardGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDashboardGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDashboardGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDashboardGroupResponse")
	}
	return
}

// createDashboardGroup implements the OCIOperation interface (enables retrying operations)
func (client DashboardGroupClient) createDashboardGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dashboardGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDashboardGroupResponse
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

// DeleteDashboardGroup Deletes the specified dashboard group. Uses the dashboard group's OCID to determine which dashboard group to delete.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dashboardservice/DeleteDashboardGroup.go.html to see an example of how to use DeleteDashboardGroup API.
func (client DashboardGroupClient) DeleteDashboardGroup(ctx context.Context, request DeleteDashboardGroupRequest) (response DeleteDashboardGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDashboardGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDashboardGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDashboardGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDashboardGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDashboardGroupResponse")
	}
	return
}

// deleteDashboardGroup implements the OCIOperation interface (enables retrying operations)
func (client DashboardGroupClient) deleteDashboardGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dashboardGroups/{dashboardGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDashboardGroupResponse
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

// GetDashboardGroup Gets the specified dashboard group's information. Uses the dashboard group's OCID to determine which dashboard to retrieve.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dashboardservice/GetDashboardGroup.go.html to see an example of how to use GetDashboardGroup API.
func (client DashboardGroupClient) GetDashboardGroup(ctx context.Context, request GetDashboardGroupRequest) (response GetDashboardGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDashboardGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDashboardGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDashboardGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDashboardGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDashboardGroupResponse")
	}
	return
}

// getDashboardGroup implements the OCIOperation interface (enables retrying operations)
func (client DashboardGroupClient) getDashboardGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dashboardGroups/{dashboardGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDashboardGroupResponse
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

// ListDashboardGroups Returns a list of dashboard groups with a specific compartment ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dashboardservice/ListDashboardGroups.go.html to see an example of how to use ListDashboardGroups API.
func (client DashboardGroupClient) ListDashboardGroups(ctx context.Context, request ListDashboardGroupsRequest) (response ListDashboardGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDashboardGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDashboardGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDashboardGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDashboardGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDashboardGroupsResponse")
	}
	return
}

// listDashboardGroups implements the OCIOperation interface (enables retrying operations)
func (client DashboardGroupClient) listDashboardGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dashboardGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDashboardGroupsResponse
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

// UpdateDashboardGroup Updates the specified dashboard group. Uses the dashboard group's OCID to determine which dashboard group to update.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dashboardservice/UpdateDashboardGroup.go.html to see an example of how to use UpdateDashboardGroup API.
func (client DashboardGroupClient) UpdateDashboardGroup(ctx context.Context, request UpdateDashboardGroupRequest) (response UpdateDashboardGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDashboardGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDashboardGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDashboardGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDashboardGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDashboardGroupResponse")
	}
	return
}

// updateDashboardGroup implements the OCIOperation interface (enables retrying operations)
func (client DashboardGroupClient) updateDashboardGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dashboardGroups/{dashboardGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDashboardGroupResponse
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
