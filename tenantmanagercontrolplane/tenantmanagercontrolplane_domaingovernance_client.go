// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v52/common"
	"github.com/oracle/oci-go-sdk/v52/common/auth"
	"net/http"
)

//DomainGovernanceClient a client for DomainGovernance
type DomainGovernanceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDomainGovernanceClientWithConfigurationProvider Creates a new default DomainGovernance client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDomainGovernanceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DomainGovernanceClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newDomainGovernanceClientFromBaseClient(baseClient, provider)
}

// NewDomainGovernanceClientWithOboToken Creates a new default DomainGovernance client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewDomainGovernanceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DomainGovernanceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDomainGovernanceClientFromBaseClient(baseClient, configProvider)
}

func newDomainGovernanceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DomainGovernanceClient, err error) {
	// DomainGovernance service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DomainGovernanceClient{BaseClient: baseClient}
	client.BasePath = "20200801"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DomainGovernanceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("tenantmanagercontrolplane", "https://organizations.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DomainGovernanceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DomainGovernanceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateDomainGovernance Adds domain governance to a claimed domain.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/CreateDomainGovernance.go.html to see an example of how to use CreateDomainGovernance API.
func (client DomainGovernanceClient) CreateDomainGovernance(ctx context.Context, request CreateDomainGovernanceRequest) (response CreateDomainGovernanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDomainGovernance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDomainGovernanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDomainGovernanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDomainGovernanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDomainGovernanceResponse")
	}
	return
}

// createDomainGovernance implements the OCIOperation interface (enables retrying operations)
func (client DomainGovernanceClient) createDomainGovernance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/domainGovernances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDomainGovernanceResponse
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

// DeleteDomainGovernance Removes domain governance from a claimed domain.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/DeleteDomainGovernance.go.html to see an example of how to use DeleteDomainGovernance API.
func (client DomainGovernanceClient) DeleteDomainGovernance(ctx context.Context, request DeleteDomainGovernanceRequest) (response DeleteDomainGovernanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDomainGovernance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDomainGovernanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDomainGovernanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDomainGovernanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDomainGovernanceResponse")
	}
	return
}

// deleteDomainGovernance implements the OCIOperation interface (enables retrying operations)
func (client DomainGovernanceClient) deleteDomainGovernance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/domainGovernances/{domainGovernanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDomainGovernanceResponse
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

// GetDomainGovernance Gets information about the domain governance entity.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/GetDomainGovernance.go.html to see an example of how to use GetDomainGovernance API.
func (client DomainGovernanceClient) GetDomainGovernance(ctx context.Context, request GetDomainGovernanceRequest) (response GetDomainGovernanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDomainGovernance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDomainGovernanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDomainGovernanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDomainGovernanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDomainGovernanceResponse")
	}
	return
}

// getDomainGovernance implements the OCIOperation interface (enables retrying operations)
func (client DomainGovernanceClient) getDomainGovernance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/domainGovernances/{domainGovernanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDomainGovernanceResponse
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

// ListDomainGovernances Return a (paginated) list of domain governance entities.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListDomainGovernances.go.html to see an example of how to use ListDomainGovernances API.
func (client DomainGovernanceClient) ListDomainGovernances(ctx context.Context, request ListDomainGovernancesRequest) (response ListDomainGovernancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDomainGovernances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDomainGovernancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDomainGovernancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDomainGovernancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDomainGovernancesResponse")
	}
	return
}

// listDomainGovernances implements the OCIOperation interface (enables retrying operations)
func (client DomainGovernanceClient) listDomainGovernances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/domainGovernances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDomainGovernancesResponse
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

// UpdateDomainGovernance Updates the domain governance entity.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/UpdateDomainGovernance.go.html to see an example of how to use UpdateDomainGovernance API.
func (client DomainGovernanceClient) UpdateDomainGovernance(ctx context.Context, request UpdateDomainGovernanceRequest) (response UpdateDomainGovernanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDomainGovernance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDomainGovernanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDomainGovernanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDomainGovernanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDomainGovernanceResponse")
	}
	return
}

// updateDomainGovernance implements the OCIOperation interface (enables retrying operations)
func (client DomainGovernanceClient) updateDomainGovernance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/domainGovernances/{domainGovernanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDomainGovernanceResponse
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
