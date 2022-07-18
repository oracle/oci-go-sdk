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

//ServiceAttachmentClient a client for ServiceAttachment
type ServiceAttachmentClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewServiceAttachmentClientWithConfigurationProvider Creates a new default ServiceAttachment client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewServiceAttachmentClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ServiceAttachmentClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newServiceAttachmentClientFromBaseClient(baseClient, provider)
}

// NewServiceAttachmentClientWithOboToken Creates a new default ServiceAttachment client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewServiceAttachmentClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ServiceAttachmentClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newServiceAttachmentClientFromBaseClient(baseClient, configProvider)
}

func newServiceAttachmentClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ServiceAttachmentClient, err error) {
	// ServiceAttachment service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ServiceAttachment"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ServiceAttachmentClient{BaseClient: baseClient}
	client.BasePath = "20211201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ServiceAttachmentClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("fusionapps", "https://fusionapps.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ServiceAttachmentClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ServiceAttachmentClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetServiceAttachment Gets a Service Attachment by identifier
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetServiceAttachment.go.html to see an example of how to use GetServiceAttachment API.
func (client ServiceAttachmentClient) GetServiceAttachment(ctx context.Context, request GetServiceAttachmentRequest) (response GetServiceAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getServiceAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetServiceAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetServiceAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetServiceAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetServiceAttachmentResponse")
	}
	return
}

// getServiceAttachment implements the OCIOperation interface (enables retrying operations)
func (client ServiceAttachmentClient) getServiceAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/serviceAttachments/{serviceAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetServiceAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ServiceAttachment", "GetServiceAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListServiceAttachments Returns a list of service attachments.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListServiceAttachments.go.html to see an example of how to use ListServiceAttachments API.
func (client ServiceAttachmentClient) ListServiceAttachments(ctx context.Context, request ListServiceAttachmentsRequest) (response ListServiceAttachmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listServiceAttachments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListServiceAttachmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListServiceAttachmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListServiceAttachmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListServiceAttachmentsResponse")
	}
	return
}

// listServiceAttachments implements the OCIOperation interface (enables retrying operations)
func (client ServiceAttachmentClient) listServiceAttachments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/serviceAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListServiceAttachmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ServiceAttachment", "ListServiceAttachments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
