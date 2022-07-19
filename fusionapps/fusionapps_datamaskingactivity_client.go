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

//DataMaskingActivityClient a client for DataMaskingActivity
type DataMaskingActivityClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDataMaskingActivityClientWithConfigurationProvider Creates a new default DataMaskingActivity client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDataMaskingActivityClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DataMaskingActivityClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newDataMaskingActivityClientFromBaseClient(baseClient, provider)
}

// NewDataMaskingActivityClientWithOboToken Creates a new default DataMaskingActivity client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewDataMaskingActivityClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DataMaskingActivityClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDataMaskingActivityClientFromBaseClient(baseClient, configProvider)
}

func newDataMaskingActivityClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DataMaskingActivityClient, err error) {
	// DataMaskingActivity service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DataMaskingActivity"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DataMaskingActivityClient{BaseClient: baseClient}
	client.BasePath = "20211201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DataMaskingActivityClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("fusionapps", "https://fusionapps.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DataMaskingActivityClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DataMaskingActivityClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateDataMaskingActivity Creates a new DataMaskingActivity.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateDataMaskingActivity.go.html to see an example of how to use CreateDataMaskingActivity API.
func (client DataMaskingActivityClient) CreateDataMaskingActivity(ctx context.Context, request CreateDataMaskingActivityRequest) (response CreateDataMaskingActivityResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDataMaskingActivity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDataMaskingActivityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDataMaskingActivityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataMaskingActivityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataMaskingActivityResponse")
	}
	return
}

// createDataMaskingActivity implements the OCIOperation interface (enables retrying operations)
func (client DataMaskingActivityClient) createDataMaskingActivity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/dataMaskingActivities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDataMaskingActivityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DataMaskingActivity", "CreateDataMaskingActivity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDataMaskingActivity Gets a DataMaskingActivity by identifier
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetDataMaskingActivity.go.html to see an example of how to use GetDataMaskingActivity API.
func (client DataMaskingActivityClient) GetDataMaskingActivity(ctx context.Context, request GetDataMaskingActivityRequest) (response GetDataMaskingActivityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataMaskingActivity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataMaskingActivityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataMaskingActivityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataMaskingActivityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataMaskingActivityResponse")
	}
	return
}

// getDataMaskingActivity implements the OCIOperation interface (enables retrying operations)
func (client DataMaskingActivityClient) getDataMaskingActivity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/dataMaskingActivities/{dataMaskingActivityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDataMaskingActivityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DataMaskingActivity", "GetDataMaskingActivity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDataMaskingActivities Returns a list of DataMaskingActivities.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListDataMaskingActivities.go.html to see an example of how to use ListDataMaskingActivities API.
func (client DataMaskingActivityClient) ListDataMaskingActivities(ctx context.Context, request ListDataMaskingActivitiesRequest) (response ListDataMaskingActivitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataMaskingActivities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataMaskingActivitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataMaskingActivitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataMaskingActivitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataMaskingActivitiesResponse")
	}
	return
}

// listDataMaskingActivities implements the OCIOperation interface (enables retrying operations)
func (client DataMaskingActivityClient) listDataMaskingActivities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/dataMaskingActivities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDataMaskingActivitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DataMaskingActivity", "ListDataMaskingActivities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
