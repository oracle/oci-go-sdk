// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Threat Intelligence API
//
// Use the Threat Intelligence API to view indicators of compromise and related items. For more information, see Overview of Threat Intelligence (https://docs.cloud.oracle.com/Content/ThreatIntelligence/Concepts/threatintelligenceoverview.htm).
//

package threatintelligence

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"github.com/oracle/oci-go-sdk/v62/common/auth"
	"net/http"
)

//ThreatintelClient a client for Threatintel
type ThreatintelClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewThreatintelClientWithConfigurationProvider Creates a new default Threatintel client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewThreatintelClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ThreatintelClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newThreatintelClientFromBaseClient(baseClient, provider)
}

// NewThreatintelClientWithOboToken Creates a new default Threatintel client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewThreatintelClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ThreatintelClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newThreatintelClientFromBaseClient(baseClient, configProvider)
}

func newThreatintelClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ThreatintelClient, err error) {
	// Threatintel service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ThreatintelClient{BaseClient: baseClient}
	client.BasePath = "20210831"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ThreatintelClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("threatintelligence", "https://api-threatintel.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ThreatintelClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ThreatintelClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetIndicator Gets a detailed indicator by identifier
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/threatintelligence/GetIndicator.go.html to see an example of how to use GetIndicator API.
func (client ThreatintelClient) GetIndicator(ctx context.Context, request GetIndicatorRequest) (response GetIndicatorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getIndicator, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetIndicatorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetIndicatorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetIndicatorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetIndicatorResponse")
	}
	return
}

// getIndicator implements the OCIOperation interface (enables retrying operations)
func (client ThreatintelClient) getIndicator(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/indicators/{indicatorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetIndicatorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/threat-intel/20210831/Indicator/GetIndicator"
		err = common.PostProcessServiceError(err, "Threatintel", "GetIndicator", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListIndicatorCounts Get the current count of each indicator type.  Results can be sorted ASC or DESC by count.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/threatintelligence/ListIndicatorCounts.go.html to see an example of how to use ListIndicatorCounts API.
func (client ThreatintelClient) ListIndicatorCounts(ctx context.Context, request ListIndicatorCountsRequest) (response ListIndicatorCountsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listIndicatorCounts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListIndicatorCountsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListIndicatorCountsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIndicatorCountsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIndicatorCountsResponse")
	}
	return
}

// listIndicatorCounts implements the OCIOperation interface (enables retrying operations)
func (client ThreatintelClient) listIndicatorCounts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/indicatorCounts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListIndicatorCountsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/threat-intel/20210831/IndicatorCountCollection/ListIndicatorCounts"
		err = common.PostProcessServiceError(err, "Threatintel", "ListIndicatorCounts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListIndicators Returns a list of IndicatorSummary objects.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/threatintelligence/ListIndicators.go.html to see an example of how to use ListIndicators API.
func (client ThreatintelClient) ListIndicators(ctx context.Context, request ListIndicatorsRequest) (response ListIndicatorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listIndicators, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListIndicatorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListIndicatorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIndicatorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIndicatorsResponse")
	}
	return
}

// listIndicators implements the OCIOperation interface (enables retrying operations)
func (client ThreatintelClient) listIndicators(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/indicators", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListIndicatorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/threat-intel/20210831/IndicatorSummaryCollection/ListIndicators"
		err = common.PostProcessServiceError(err, "Threatintel", "ListIndicators", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListThreatTypes Gets a list of threat types that are available to use as parameters when querying indicators.
// This is sorted by threat type name according to the sort order query param.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/threatintelligence/ListThreatTypes.go.html to see an example of how to use ListThreatTypes API.
func (client ThreatintelClient) ListThreatTypes(ctx context.Context, request ListThreatTypesRequest) (response ListThreatTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listThreatTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListThreatTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListThreatTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListThreatTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListThreatTypesResponse")
	}
	return
}

// listThreatTypes implements the OCIOperation interface (enables retrying operations)
func (client ThreatintelClient) listThreatTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/threatTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListThreatTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/threat-intel/20210831/ThreatTypesCollection/ListThreatTypes"
		err = common.PostProcessServiceError(err, "Threatintel", "ListThreatTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
