// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center API
//
// OCI Control Center (OCC) service enables you to monitor the region-level cloud consumption and provides the region-level capacity data, in realms where OCC is available. Use the OCI Control Center (OCC) API to explore region-level capacity and utilization information about core services. For more information, see OCI Control Center (https://docs.cloud.oracle.com/iaas/Content/control-center/home.htm).
//

package ocicontrolcenter

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OccMetricsClient a client for OccMetrics
type OccMetricsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOccMetricsClientWithConfigurationProvider Creates a new default OccMetrics client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOccMetricsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OccMetricsClient, err error) {
	if enabled := common.CheckForEnabledServices("ocicontrolcenter"); !enabled {
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
	return newOccMetricsClientFromBaseClient(baseClient, provider)
}

// NewOccMetricsClientWithOboToken Creates a new default OccMetrics client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOccMetricsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OccMetricsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOccMetricsClientFromBaseClient(baseClient, configProvider)
}

func newOccMetricsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OccMetricsClient, err error) {
	// OccMetrics service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OccMetrics"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OccMetricsClient{BaseClient: baseClient}
	client.BasePath = "20230515"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OccMetricsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ocicontrolcenter", "https://control-center.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OccMetricsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OccMetricsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListMetricProperties Returns a list of available metrics for the given namespace. The results for metrics with dimensions includes list of all the associated dimensions. The results are sorted by the metricName and then by dimension in ascending alphabetical order. For a list of valid namespaces, see ListNamespaces.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocicontrolcenter/ListMetricProperties.go.html to see an example of how to use ListMetricProperties API.
// A default retry strategy applies to this operation ListMetricProperties()
func (client OccMetricsClient) ListMetricProperties(ctx context.Context, request ListMetricPropertiesRequest) (response ListMetricPropertiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMetricProperties, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMetricPropertiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMetricPropertiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMetricPropertiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMetricPropertiesResponse")
	}
	return
}

// listMetricProperties implements the OCIOperation interface (enables retrying operations)
func (client OccMetricsClient) listMetricProperties(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/metricProperties/{namespaceName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMetricPropertiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occ/20230515/MetricPropertyCollection/ListMetricProperties"
		err = common.PostProcessServiceError(err, "OccMetrics", "ListMetricProperties", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNamespaces List all the available source services called namespaces emitting metrics for this region. The namespaces are sorted in ascending alphabetical order.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocicontrolcenter/ListNamespaces.go.html to see an example of how to use ListNamespaces API.
// A default retry strategy applies to this operation ListNamespaces()
func (client OccMetricsClient) ListNamespaces(ctx context.Context, request ListNamespacesRequest) (response ListNamespacesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNamespaces, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNamespacesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNamespacesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNamespacesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNamespacesResponse")
	}
	return
}

// listNamespaces implements the OCIOperation interface (enables retrying operations)
func (client OccMetricsClient) listNamespaces(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNamespacesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occ/20230515/NamespaceCollection/ListNamespaces"
		err = common.PostProcessServiceError(err, "OccMetrics", "ListNamespaces", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedMetricData Returns the summarized data for the given metric from the given namespace.  The aggregation method depends on the metric.
// The metric data can be filtered by providing the dimension, startTime or endTime.  The metric
// data in the response is sorted by dimension in ascending order and then by sampleTime in ascending chronological order.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocicontrolcenter/RequestSummarizedMetricData.go.html to see an example of how to use RequestSummarizedMetricData API.
// A default retry strategy applies to this operation RequestSummarizedMetricData()
func (client OccMetricsClient) RequestSummarizedMetricData(ctx context.Context, request RequestSummarizedMetricDataRequest) (response RequestSummarizedMetricDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedMetricData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedMetricDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedMetricDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedMetricDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedMetricDataResponse")
	}
	return
}

// requestSummarizedMetricData implements the OCIOperation interface (enables retrying operations)
func (client OccMetricsClient) requestSummarizedMetricData(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/requestSummarizedMetricData", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedMetricDataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occ/20230515/SummarizedMetricDataCollection/RequestSummarizedMetricData"
		err = common.PostProcessServiceError(err, "OccMetrics", "RequestSummarizedMetricData", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
