// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Threat Intelligence API
//
// Use the Threat Intelligence API to search for information about known threat indicators, including suspicious IP addresses, domain names, and other digital fingerprints. Threat Intelligence is a managed database of curated threat intelligence that comes from first party Oracle security insights, open source feeds, and vendor-procured data. For more information, see the Threat Intelligence documentation (https://docs.oracle.com/iaas/Content/threat-intel/home.htm).
//

package threatintelligence

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ThreatintelClient a client for Threatintel
type ThreatintelClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewThreatintelClientWithConfigurationProvider Creates a new default Threatintel client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewThreatintelClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ThreatintelClient, err error) {
	if enabled := common.CheckForEnabledServices("threatintelligence"); !enabled {
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
	return newThreatintelClientFromBaseClient(baseClient, provider)
}

// NewThreatintelClientWithOboToken Creates a new default Threatintel client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewThreatintelClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ThreatintelClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newThreatintelClientFromBaseClient(baseClient, configProvider)
}

func newThreatintelClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ThreatintelClient, err error) {
	// Threatintel service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Threatintel"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ThreatintelClient{BaseClient: baseClient}
	client.BasePath = "20220901"
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
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *ThreatintelClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetIndicator Get detailed information about a threat indicator with a given identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/threatintelligence/GetIndicator.go.html to see an example of how to use GetIndicator API.
// A default retry strategy applies to this operation GetIndicator()
func (client ThreatintelClient) GetIndicator(ctx context.Context, request GetIndicatorRequest) (response GetIndicatorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/threat-intel/20220901/Indicator/GetIndicator"
		err = common.PostProcessServiceError(err, "Threatintel", "GetIndicator", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListIndicatorCounts Get the current count of each threat indicator type. Indicator counts can be sorted in ascending or descending order.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/threatintelligence/ListIndicatorCounts.go.html to see an example of how to use ListIndicatorCounts API.
// A default retry strategy applies to this operation ListIndicatorCounts()
func (client ThreatintelClient) ListIndicatorCounts(ctx context.Context, request ListIndicatorCountsRequest) (response ListIndicatorCountsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/threat-intel/20220901/IndicatorCountCollection/ListIndicatorCounts"
		err = common.PostProcessServiceError(err, "Threatintel", "ListIndicatorCounts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListIndicators Get a list of threat indicator summaries based on the search criteria.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/threatintelligence/ListIndicators.go.html to see an example of how to use ListIndicators API.
// A default retry strategy applies to this operation ListIndicators()
func (client ThreatintelClient) ListIndicators(ctx context.Context, request ListIndicatorsRequest) (response ListIndicatorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/threat-intel/20220901/IndicatorSummaryCollection/ListIndicators"
		err = common.PostProcessServiceError(err, "Threatintel", "ListIndicators", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListThreatTypes Gets a list of threat types that are available to use as parameters when querying indicators.
// The list is sorted by threat type name according to the sort order query param.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/threatintelligence/ListThreatTypes.go.html to see an example of how to use ListThreatTypes API.
// A default retry strategy applies to this operation ListThreatTypes()
func (client ThreatintelClient) ListThreatTypes(ctx context.Context, request ListThreatTypesRequest) (response ListThreatTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/threat-intel/20220901/ThreatTypesCollection/ListThreatTypes"
		err = common.PostProcessServiceError(err, "Threatintel", "ListThreatTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeIndicators Get indicator summaries based on advanced search criteria.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/threatintelligence/SummarizeIndicators.go.html to see an example of how to use SummarizeIndicators API.
// A default retry strategy applies to this operation SummarizeIndicators()
func (client ThreatintelClient) SummarizeIndicators(ctx context.Context, request SummarizeIndicatorsRequest) (response SummarizeIndicatorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeIndicators, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeIndicatorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeIndicatorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeIndicatorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeIndicatorsResponse")
	}
	return
}

// summarizeIndicators implements the OCIOperation interface (enables retrying operations)
func (client ThreatintelClient) summarizeIndicators(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/indicators/actions/summarize", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeIndicatorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/threat-intel/20220901/Indicator/SummarizeIndicators"
		err = common.PostProcessServiceError(err, "Threatintel", "SummarizeIndicators", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
