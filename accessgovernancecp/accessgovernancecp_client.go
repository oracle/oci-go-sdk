// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Access Governance APIs
//
// Use the Oracle Access Governance API to create, view, and manage GovernanceInstances.
//

package accessgovernancecp

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// AccessGovernanceCPClient a client for AccessGovernanceCP
type AccessGovernanceCPClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAccessGovernanceCPClientWithConfigurationProvider Creates a new default AccessGovernanceCP client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAccessGovernanceCPClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AccessGovernanceCPClient, err error) {
	if enabled := common.CheckForEnabledServices("accessgovernancecp"); !enabled {
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
	return newAccessGovernanceCPClientFromBaseClient(baseClient, provider)
}

// NewAccessGovernanceCPClientWithOboToken Creates a new default AccessGovernanceCP client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewAccessGovernanceCPClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client AccessGovernanceCPClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newAccessGovernanceCPClientFromBaseClient(baseClient, configProvider)
}

func newAccessGovernanceCPClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client AccessGovernanceCPClient, err error) {
	// AccessGovernanceCP service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("AccessGovernanceCP"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = AccessGovernanceCPClient{BaseClient: baseClient}
	client.BasePath = "20220518"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AccessGovernanceCPClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("accessgovernancecp", "https://cp-prod.access-governance.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AccessGovernanceCPClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *AccessGovernanceCPClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeGovernanceInstanceCompartment Moves a GovernanceInstance resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/accessgovernancecp/ChangeGovernanceInstanceCompartment.go.html to see an example of how to use ChangeGovernanceInstanceCompartment API.
// A default retry strategy applies to this operation ChangeGovernanceInstanceCompartment()
func (client AccessGovernanceCPClient) ChangeGovernanceInstanceCompartment(ctx context.Context, request ChangeGovernanceInstanceCompartmentRequest) (response ChangeGovernanceInstanceCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeGovernanceInstanceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeGovernanceInstanceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeGovernanceInstanceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeGovernanceInstanceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeGovernanceInstanceCompartmentResponse")
	}
	return
}

// changeGovernanceInstanceCompartment implements the OCIOperation interface (enables retrying operations)
func (client AccessGovernanceCPClient) changeGovernanceInstanceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/governanceInstances/{governanceInstanceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeGovernanceInstanceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/access-governance-cp/20220518/GovernanceInstance/ChangeGovernanceInstanceCompartment"
		err = common.PostProcessServiceError(err, "AccessGovernanceCP", "ChangeGovernanceInstanceCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateGovernanceInstance Creates a new GovernanceInstance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/accessgovernancecp/CreateGovernanceInstance.go.html to see an example of how to use CreateGovernanceInstance API.
// A default retry strategy applies to this operation CreateGovernanceInstance()
func (client AccessGovernanceCPClient) CreateGovernanceInstance(ctx context.Context, request CreateGovernanceInstanceRequest) (response CreateGovernanceInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createGovernanceInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateGovernanceInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateGovernanceInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateGovernanceInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateGovernanceInstanceResponse")
	}
	return
}

// createGovernanceInstance implements the OCIOperation interface (enables retrying operations)
func (client AccessGovernanceCPClient) createGovernanceInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/governanceInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateGovernanceInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/access-governance-cp/20220518/GovernanceInstance/CreateGovernanceInstance"
		err = common.PostProcessServiceError(err, "AccessGovernanceCP", "CreateGovernanceInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteGovernanceInstance Deletes an existing GovernanceInstance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/accessgovernancecp/DeleteGovernanceInstance.go.html to see an example of how to use DeleteGovernanceInstance API.
// A default retry strategy applies to this operation DeleteGovernanceInstance()
func (client AccessGovernanceCPClient) DeleteGovernanceInstance(ctx context.Context, request DeleteGovernanceInstanceRequest) (response DeleteGovernanceInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteGovernanceInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteGovernanceInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteGovernanceInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteGovernanceInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteGovernanceInstanceResponse")
	}
	return
}

// deleteGovernanceInstance implements the OCIOperation interface (enables retrying operations)
func (client AccessGovernanceCPClient) deleteGovernanceInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/governanceInstances/{governanceInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteGovernanceInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/access-governance-cp/20220518/GovernanceInstance/DeleteGovernanceInstance"
		err = common.PostProcessServiceError(err, "AccessGovernanceCP", "DeleteGovernanceInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetGovernanceInstance Gets a GovernanceInstance by OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/accessgovernancecp/GetGovernanceInstance.go.html to see an example of how to use GetGovernanceInstance API.
// A default retry strategy applies to this operation GetGovernanceInstance()
func (client AccessGovernanceCPClient) GetGovernanceInstance(ctx context.Context, request GetGovernanceInstanceRequest) (response GetGovernanceInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGovernanceInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGovernanceInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGovernanceInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGovernanceInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGovernanceInstanceResponse")
	}
	return
}

// getGovernanceInstance implements the OCIOperation interface (enables retrying operations)
func (client AccessGovernanceCPClient) getGovernanceInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/governanceInstances/{governanceInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGovernanceInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/access-governance-cp/20220518/GovernanceInstance/GetGovernanceInstance"
		err = common.PostProcessServiceError(err, "AccessGovernanceCP", "GetGovernanceInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetGovernanceInstanceConfiguration Gets the tenancy-wide configuration for GovernanceInstances
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/accessgovernancecp/GetGovernanceInstanceConfiguration.go.html to see an example of how to use GetGovernanceInstanceConfiguration API.
// A default retry strategy applies to this operation GetGovernanceInstanceConfiguration()
func (client AccessGovernanceCPClient) GetGovernanceInstanceConfiguration(ctx context.Context, request GetGovernanceInstanceConfigurationRequest) (response GetGovernanceInstanceConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGovernanceInstanceConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGovernanceInstanceConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGovernanceInstanceConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGovernanceInstanceConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGovernanceInstanceConfigurationResponse")
	}
	return
}

// getGovernanceInstanceConfiguration implements the OCIOperation interface (enables retrying operations)
func (client AccessGovernanceCPClient) getGovernanceInstanceConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/governanceInstances/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGovernanceInstanceConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/access-governance-cp/20220518/GovernanceInstanceConfiguration/GetGovernanceInstanceConfiguration"
		err = common.PostProcessServiceError(err, "AccessGovernanceCP", "GetGovernanceInstanceConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListGovernanceInstances Returns a list of Governance Instances.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/accessgovernancecp/ListGovernanceInstances.go.html to see an example of how to use ListGovernanceInstances API.
// A default retry strategy applies to this operation ListGovernanceInstances()
func (client AccessGovernanceCPClient) ListGovernanceInstances(ctx context.Context, request ListGovernanceInstancesRequest) (response ListGovernanceInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGovernanceInstances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGovernanceInstancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGovernanceInstancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGovernanceInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGovernanceInstancesResponse")
	}
	return
}

// listGovernanceInstances implements the OCIOperation interface (enables retrying operations)
func (client AccessGovernanceCPClient) listGovernanceInstances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/governanceInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListGovernanceInstancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/access-governance-cp/20220518/GovernanceInstanceCollection/ListGovernanceInstances"
		err = common.PostProcessServiceError(err, "AccessGovernanceCP", "ListGovernanceInstances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateGovernanceInstance Updates the GovernanceInstance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/accessgovernancecp/UpdateGovernanceInstance.go.html to see an example of how to use UpdateGovernanceInstance API.
// A default retry strategy applies to this operation UpdateGovernanceInstance()
func (client AccessGovernanceCPClient) UpdateGovernanceInstance(ctx context.Context, request UpdateGovernanceInstanceRequest) (response UpdateGovernanceInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateGovernanceInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateGovernanceInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateGovernanceInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateGovernanceInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateGovernanceInstanceResponse")
	}
	return
}

// updateGovernanceInstance implements the OCIOperation interface (enables retrying operations)
func (client AccessGovernanceCPClient) updateGovernanceInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/governanceInstances/{governanceInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateGovernanceInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/access-governance-cp/20220518/GovernanceInstance/UpdateGovernanceInstance"
		err = common.PostProcessServiceError(err, "AccessGovernanceCP", "UpdateGovernanceInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateGovernanceInstanceConfiguration Updates the tenancy-wide configuration for GovernanceInstances
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/accessgovernancecp/UpdateGovernanceInstanceConfiguration.go.html to see an example of how to use UpdateGovernanceInstanceConfiguration API.
// A default retry strategy applies to this operation UpdateGovernanceInstanceConfiguration()
func (client AccessGovernanceCPClient) UpdateGovernanceInstanceConfiguration(ctx context.Context, request UpdateGovernanceInstanceConfigurationRequest) (response UpdateGovernanceInstanceConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateGovernanceInstanceConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateGovernanceInstanceConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateGovernanceInstanceConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateGovernanceInstanceConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateGovernanceInstanceConfigurationResponse")
	}
	return
}

// updateGovernanceInstanceConfiguration implements the OCIOperation interface (enables retrying operations)
func (client AccessGovernanceCPClient) updateGovernanceInstanceConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/governanceInstances/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateGovernanceInstanceConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/access-governance-cp/20220518/GovernanceInstance/UpdateGovernanceInstanceConfiguration"
		err = common.PostProcessServiceError(err, "AccessGovernanceCP", "UpdateGovernanceInstanceConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
