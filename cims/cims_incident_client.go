// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests.
// For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm).
// **Note**: Before you can create service requests with this API,
// complete user registration at My Oracle Cloud Support
// and then ask your tenancy administrator to provide you authorization for the related user groups.
//

package cims

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// IncidentClient a client for Incident
type IncidentClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewIncidentClientWithConfigurationProvider Creates a new default Incident client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewIncidentClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client IncidentClient, err error) {
	if enabled := common.CheckForEnabledServices("cims"); !enabled {
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
	return newIncidentClientFromBaseClient(baseClient, provider)
}

// NewIncidentClientWithOboToken Creates a new default Incident client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewIncidentClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client IncidentClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newIncidentClientFromBaseClient(baseClient, configProvider)
}

func newIncidentClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client IncidentClient, err error) {
	// Incident service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Incident"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = IncidentClient{BaseClient: baseClient}
	client.BasePath = "20181231"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *IncidentClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("cims", "https://incidentmanagement.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *IncidentClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *IncidentClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateIncident Creates a support ticket in the specified tenancy.
// For more information, see Creating Support Requests (https://docs.cloud.oracle.com/iaas/Content/GSG/support/create-incident.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cims/CreateIncident.go.html to see an example of how to use CreateIncident API.
func (client IncidentClient) CreateIncident(ctx context.Context, request CreateIncidentRequest) (response CreateIncidentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
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
func (client IncidentClient) createIncident(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/v2/incidents", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateIncidentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/incidentmanagement/20181231/Incident/CreateIncident"
		err = common.PostProcessServiceError(err, "Incident", "CreateIncident", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetIncident Gets the specified support ticket.
// For more information, see Getting Details for a Support Request (https://docs.cloud.oracle.com/iaas/Content/GSG/support/get-incident.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cims/GetIncident.go.html to see an example of how to use GetIncident API.
func (client IncidentClient) GetIncident(ctx context.Context, request GetIncidentRequest) (response GetIncidentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
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
func (client IncidentClient) getIncident(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/v2/incidents/{incidentKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetIncidentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/incidentmanagement/20181231/Incident/GetIncident"
		err = common.PostProcessServiceError(err, "Incident", "GetIncident", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListIncidentResourceTypes Depending on the selected `productType`, either
// lists available products (service groups, services, service categories, and subcategories) for technical support tickets or
// lists limits and current usage for limit increase tickets.
// This operation is called during creation of technical support and limit increase tickets.
// For more information about listing products, see
// Listing Products for Support Requests (https://docs.cloud.oracle.com/iaas/Content/GSG/support/list-incident-resource-types-taxonomy.htm).
// For more information about listing limits, see
// Listing Limits for Service Limit Increase Requests (https://docs.cloud.oracle.com/iaas/Content/GSG/support/list-incident-resource-types-limit.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cims/ListIncidentResourceTypes.go.html to see an example of how to use ListIncidentResourceTypes API.
func (client IncidentClient) ListIncidentResourceTypes(ctx context.Context, request ListIncidentResourceTypesRequest) (response ListIncidentResourceTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
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
func (client IncidentClient) listIncidentResourceTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/v2/incidents/incidentResourceTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListIncidentResourceTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/incidentmanagement/20181231/IncidentResourceType/ListIncidentResourceTypes"
		err = common.PostProcessServiceError(err, "Incident", "ListIncidentResourceTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListIncidents Lists support tickets for the specified tenancy.
// For more information, see Listing Support Requests (https://docs.cloud.oracle.com/iaas/Content/GSG/support/list-incidents.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cims/ListIncidents.go.html to see an example of how to use ListIncidents API.
func (client IncidentClient) ListIncidents(ctx context.Context, request ListIncidentsRequest) (response ListIncidentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
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
func (client IncidentClient) listIncidents(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/v2/incidents", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListIncidentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/incidentmanagement/20181231/IncidentSummary/ListIncidents"
		err = common.PostProcessServiceError(err, "Incident", "ListIncidents", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateIncident Updates the specified support ticket.
// For more information, see Updating Support Requests (https://docs.cloud.oracle.com/iaas/Content/GSG/support/update-incident.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cims/UpdateIncident.go.html to see an example of how to use UpdateIncident API.
func (client IncidentClient) UpdateIncident(ctx context.Context, request UpdateIncidentRequest) (response UpdateIncidentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
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
func (client IncidentClient) updateIncident(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/v2/incidents/{incidentKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateIncidentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/incidentmanagement/20181231/UpdateIncident/UpdateIncident"
		err = common.PostProcessServiceError(err, "Incident", "UpdateIncident", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateUser Checks whether the requested user is valid.
// For more information, see Validating a User (https://docs.cloud.oracle.com/iaas/Content/GSG/support/validate-user.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cims/ValidateUser.go.html to see an example of how to use ValidateUser API.
func (client IncidentClient) ValidateUser(ctx context.Context, request ValidateUserRequest) (response ValidateUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
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
func (client IncidentClient) validateUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/v2/incidents/user/validate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/incidentmanagement/20181231/ValidationResponse/ValidateUser"
		err = common.PostProcessServiceError(err, "Incident", "ValidateUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
