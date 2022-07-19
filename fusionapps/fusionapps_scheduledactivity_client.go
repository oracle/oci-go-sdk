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

//ScheduledActivityClient a client for ScheduledActivity
type ScheduledActivityClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewScheduledActivityClientWithConfigurationProvider Creates a new default ScheduledActivity client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewScheduledActivityClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ScheduledActivityClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newScheduledActivityClientFromBaseClient(baseClient, provider)
}

// NewScheduledActivityClientWithOboToken Creates a new default ScheduledActivity client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewScheduledActivityClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ScheduledActivityClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newScheduledActivityClientFromBaseClient(baseClient, configProvider)
}

func newScheduledActivityClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ScheduledActivityClient, err error) {
	// ScheduledActivity service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ScheduledActivity"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ScheduledActivityClient{BaseClient: baseClient}
	client.BasePath = "20211201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ScheduledActivityClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("fusionapps", "https://fusionapps.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ScheduledActivityClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ScheduledActivityClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetScheduledActivity Gets a ScheduledActivity by identifier
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetScheduledActivity.go.html to see an example of how to use GetScheduledActivity API.
func (client ScheduledActivityClient) GetScheduledActivity(ctx context.Context, request GetScheduledActivityRequest) (response GetScheduledActivityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getScheduledActivity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetScheduledActivityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetScheduledActivityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetScheduledActivityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetScheduledActivityResponse")
	}
	return
}

// getScheduledActivity implements the OCIOperation interface (enables retrying operations)
func (client ScheduledActivityClient) getScheduledActivity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/scheduledActivities/{scheduledActivityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetScheduledActivityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ScheduledActivity", "GetScheduledActivity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListScheduledActivities Returns a list of ScheduledActivities.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListScheduledActivities.go.html to see an example of how to use ListScheduledActivities API.
func (client ScheduledActivityClient) ListScheduledActivities(ctx context.Context, request ListScheduledActivitiesRequest) (response ListScheduledActivitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listScheduledActivities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListScheduledActivitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListScheduledActivitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListScheduledActivitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListScheduledActivitiesResponse")
	}
	return
}

// listScheduledActivities implements the OCIOperation interface (enables retrying operations)
func (client ScheduledActivityClient) listScheduledActivities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/scheduledActivities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListScheduledActivitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ScheduledActivity", "ListScheduledActivities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
