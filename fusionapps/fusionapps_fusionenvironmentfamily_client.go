// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/Identity/fusion-applications/home.htm).
//

package fusionapps

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//FusionEnvironmentFamilyClient a client for FusionEnvironmentFamily
type FusionEnvironmentFamilyClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewFusionEnvironmentFamilyClientWithConfigurationProvider Creates a new default FusionEnvironmentFamily client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewFusionEnvironmentFamilyClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client FusionEnvironmentFamilyClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newFusionEnvironmentFamilyClientFromBaseClient(baseClient, provider)
}

// NewFusionEnvironmentFamilyClientWithOboToken Creates a new default FusionEnvironmentFamily client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewFusionEnvironmentFamilyClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client FusionEnvironmentFamilyClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newFusionEnvironmentFamilyClientFromBaseClient(baseClient, configProvider)
}

func newFusionEnvironmentFamilyClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client FusionEnvironmentFamilyClient, err error) {
	// FusionEnvironmentFamily service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("FusionEnvironmentFamily"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = FusionEnvironmentFamilyClient{BaseClient: baseClient}
	client.BasePath = "20211201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *FusionEnvironmentFamilyClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("fusionapps", "https://fusionapps.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *FusionEnvironmentFamilyClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("Invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *FusionEnvironmentFamilyClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeFusionEnvironmentFamilyCompartment Moves a FusionEnvironmentFamily into a different compartment. When provided, If-Match is checked against ETag
// values of the resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ChangeFusionEnvironmentFamilyCompartment.go.html to see an example of how to use ChangeFusionEnvironmentFamilyCompartment API.
func (client FusionEnvironmentFamilyClient) ChangeFusionEnvironmentFamilyCompartment(ctx context.Context, request ChangeFusionEnvironmentFamilyCompartmentRequest) (response ChangeFusionEnvironmentFamilyCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeFusionEnvironmentFamilyCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeFusionEnvironmentFamilyCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeFusionEnvironmentFamilyCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeFusionEnvironmentFamilyCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeFusionEnvironmentFamilyCompartmentResponse")
	}
	return
}

// changeFusionEnvironmentFamilyCompartment implements the OCIOperation interface (enables retrying operations)
func (client FusionEnvironmentFamilyClient) changeFusionEnvironmentFamilyCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironmentFamilies/{fusionEnvironmentFamilyId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeFusionEnvironmentFamilyCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "FusionEnvironmentFamily", "ChangeFusionEnvironmentFamilyCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateFusionEnvironmentFamily Creates a new FusionEnvironmentFamily.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateFusionEnvironmentFamily.go.html to see an example of how to use CreateFusionEnvironmentFamily API.
func (client FusionEnvironmentFamilyClient) CreateFusionEnvironmentFamily(ctx context.Context, request CreateFusionEnvironmentFamilyRequest) (response CreateFusionEnvironmentFamilyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createFusionEnvironmentFamily, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFusionEnvironmentFamilyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFusionEnvironmentFamilyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFusionEnvironmentFamilyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFusionEnvironmentFamilyResponse")
	}
	return
}

// createFusionEnvironmentFamily implements the OCIOperation interface (enables retrying operations)
func (client FusionEnvironmentFamilyClient) createFusionEnvironmentFamily(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironmentFamilies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFusionEnvironmentFamilyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "FusionEnvironmentFamily", "CreateFusionEnvironmentFamily", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFusionEnvironmentFamily Deletes a FusionEnvironmentFamily resource by identifier
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/DeleteFusionEnvironmentFamily.go.html to see an example of how to use DeleteFusionEnvironmentFamily API.
func (client FusionEnvironmentFamilyClient) DeleteFusionEnvironmentFamily(ctx context.Context, request DeleteFusionEnvironmentFamilyRequest) (response DeleteFusionEnvironmentFamilyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFusionEnvironmentFamily, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFusionEnvironmentFamilyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFusionEnvironmentFamilyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFusionEnvironmentFamilyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFusionEnvironmentFamilyResponse")
	}
	return
}

// deleteFusionEnvironmentFamily implements the OCIOperation interface (enables retrying operations)
func (client FusionEnvironmentFamilyClient) deleteFusionEnvironmentFamily(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fusionEnvironmentFamilies/{fusionEnvironmentFamilyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFusionEnvironmentFamilyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "FusionEnvironmentFamily", "DeleteFusionEnvironmentFamily", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFusionEnvironmentFamily Retrieves a fusion environment family identified by its OCID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetFusionEnvironmentFamily.go.html to see an example of how to use GetFusionEnvironmentFamily API.
func (client FusionEnvironmentFamilyClient) GetFusionEnvironmentFamily(ctx context.Context, request GetFusionEnvironmentFamilyRequest) (response GetFusionEnvironmentFamilyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFusionEnvironmentFamily, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFusionEnvironmentFamilyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFusionEnvironmentFamilyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFusionEnvironmentFamilyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFusionEnvironmentFamilyResponse")
	}
	return
}

// getFusionEnvironmentFamily implements the OCIOperation interface (enables retrying operations)
func (client FusionEnvironmentFamilyClient) getFusionEnvironmentFamily(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironmentFamilies/{fusionEnvironmentFamilyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFusionEnvironmentFamilyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "FusionEnvironmentFamily", "GetFusionEnvironmentFamily", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFusionEnvironmentFamilyLimitsAndUsage Gets the number of environments (usage) of each type in the fusion environment family, as well as the limit that's allowed to be created based on the group's associated subscriptions.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetFusionEnvironmentFamilyLimitsAndUsage.go.html to see an example of how to use GetFusionEnvironmentFamilyLimitsAndUsage API.
func (client FusionEnvironmentFamilyClient) GetFusionEnvironmentFamilyLimitsAndUsage(ctx context.Context, request GetFusionEnvironmentFamilyLimitsAndUsageRequest) (response GetFusionEnvironmentFamilyLimitsAndUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFusionEnvironmentFamilyLimitsAndUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFusionEnvironmentFamilyLimitsAndUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFusionEnvironmentFamilyLimitsAndUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFusionEnvironmentFamilyLimitsAndUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFusionEnvironmentFamilyLimitsAndUsageResponse")
	}
	return
}

// getFusionEnvironmentFamilyLimitsAndUsage implements the OCIOperation interface (enables retrying operations)
func (client FusionEnvironmentFamilyClient) getFusionEnvironmentFamilyLimitsAndUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironmentFamilies/{fusionEnvironmentFamilyId}/limitsAndUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFusionEnvironmentFamilyLimitsAndUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "FusionEnvironmentFamily", "GetFusionEnvironmentFamilyLimitsAndUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFusionEnvironmentFamilySubscriptionDetail Gets the subscription details of an fusion environment family.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetFusionEnvironmentFamilySubscriptionDetail.go.html to see an example of how to use GetFusionEnvironmentFamilySubscriptionDetail API.
func (client FusionEnvironmentFamilyClient) GetFusionEnvironmentFamilySubscriptionDetail(ctx context.Context, request GetFusionEnvironmentFamilySubscriptionDetailRequest) (response GetFusionEnvironmentFamilySubscriptionDetailResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFusionEnvironmentFamilySubscriptionDetail, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFusionEnvironmentFamilySubscriptionDetailResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFusionEnvironmentFamilySubscriptionDetailResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFusionEnvironmentFamilySubscriptionDetailResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFusionEnvironmentFamilySubscriptionDetailResponse")
	}
	return
}

// getFusionEnvironmentFamilySubscriptionDetail implements the OCIOperation interface (enables retrying operations)
func (client FusionEnvironmentFamilyClient) getFusionEnvironmentFamilySubscriptionDetail(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironmentFamilies/{fusionEnvironmentFamilyId}/subscriptionDetails", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFusionEnvironmentFamilySubscriptionDetailResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "FusionEnvironmentFamily", "GetFusionEnvironmentFamilySubscriptionDetail", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFusionEnvironmentFamilies Returns a list of FusionEnvironmentFamilies.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListFusionEnvironmentFamilies.go.html to see an example of how to use ListFusionEnvironmentFamilies API.
func (client FusionEnvironmentFamilyClient) ListFusionEnvironmentFamilies(ctx context.Context, request ListFusionEnvironmentFamiliesRequest) (response ListFusionEnvironmentFamiliesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFusionEnvironmentFamilies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFusionEnvironmentFamiliesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFusionEnvironmentFamiliesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFusionEnvironmentFamiliesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFusionEnvironmentFamiliesResponse")
	}
	return
}

// listFusionEnvironmentFamilies implements the OCIOperation interface (enables retrying operations)
func (client FusionEnvironmentFamilyClient) listFusionEnvironmentFamilies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironmentFamilies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFusionEnvironmentFamiliesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "FusionEnvironmentFamily", "ListFusionEnvironmentFamilies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFusionEnvironmentFamily Updates the FusionEnvironmentFamily
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/UpdateFusionEnvironmentFamily.go.html to see an example of how to use UpdateFusionEnvironmentFamily API.
func (client FusionEnvironmentFamilyClient) UpdateFusionEnvironmentFamily(ctx context.Context, request UpdateFusionEnvironmentFamilyRequest) (response UpdateFusionEnvironmentFamilyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFusionEnvironmentFamily, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFusionEnvironmentFamilyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFusionEnvironmentFamilyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFusionEnvironmentFamilyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFusionEnvironmentFamilyResponse")
	}
	return
}

// updateFusionEnvironmentFamily implements the OCIOperation interface (enables retrying operations)
func (client FusionEnvironmentFamilyClient) updateFusionEnvironmentFamily(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fusionEnvironmentFamilies/{fusionEnvironmentFamilyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFusionEnvironmentFamilyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "FusionEnvironmentFamily", "UpdateFusionEnvironmentFamily", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
