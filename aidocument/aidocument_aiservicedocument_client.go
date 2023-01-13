// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//AIServiceDocumentClient a client for AIServiceDocument
type AIServiceDocumentClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAIServiceDocumentClientWithConfigurationProvider Creates a new default AIServiceDocument client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAIServiceDocumentClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AIServiceDocumentClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newAIServiceDocumentClientFromBaseClient(baseClient, provider)
}

// NewAIServiceDocumentClientWithOboToken Creates a new default AIServiceDocument client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewAIServiceDocumentClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client AIServiceDocumentClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newAIServiceDocumentClientFromBaseClient(baseClient, configProvider)
}

func newAIServiceDocumentClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client AIServiceDocumentClient, err error) {
	// AIServiceDocument service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("AIServiceDocument"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = AIServiceDocumentClient{BaseClient: baseClient}
	client.BasePath = "20221109"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AIServiceDocumentClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("aidocument", "https://document.aiservice.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AIServiceDocumentClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *AIServiceDocumentClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelProcessorJob Cancel a processor job.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aidocument/CancelProcessorJob.go.html to see an example of how to use CancelProcessorJob API.
func (client AIServiceDocumentClient) CancelProcessorJob(ctx context.Context, request CancelProcessorJobRequest) (response CancelProcessorJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelProcessorJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelProcessorJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelProcessorJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelProcessorJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelProcessorJobResponse")
	}
	return
}

// cancelProcessorJob implements the OCIOperation interface (enables retrying operations)
func (client AIServiceDocumentClient) cancelProcessorJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/processorJobs/{processorJobId}/actions/cancel", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelProcessorJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/document-understanding/20221109/ProcessorJob/CancelProcessorJob"
		err = common.PostProcessServiceError(err, "AIServiceDocument", "CancelProcessorJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateProcessorJob Create a processor job for document analysis.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aidocument/CreateProcessorJob.go.html to see an example of how to use CreateProcessorJob API.
// A default retry strategy applies to this operation CreateProcessorJob()
func (client AIServiceDocumentClient) CreateProcessorJob(ctx context.Context, request CreateProcessorJobRequest) (response CreateProcessorJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createProcessorJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateProcessorJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateProcessorJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateProcessorJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateProcessorJobResponse")
	}
	return
}

// createProcessorJob implements the OCIOperation interface (enables retrying operations)
func (client AIServiceDocumentClient) createProcessorJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/processorJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateProcessorJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/document-understanding/20221109/ProcessorJob/CreateProcessorJob"
		err = common.PostProcessServiceError(err, "AIServiceDocument", "CreateProcessorJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetProcessorJob Get the details of a processor job.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aidocument/GetProcessorJob.go.html to see an example of how to use GetProcessorJob API.
// A default retry strategy applies to this operation GetProcessorJob()
func (client AIServiceDocumentClient) GetProcessorJob(ctx context.Context, request GetProcessorJobRequest) (response GetProcessorJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getProcessorJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetProcessorJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetProcessorJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetProcessorJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetProcessorJobResponse")
	}
	return
}

// getProcessorJob implements the OCIOperation interface (enables retrying operations)
func (client AIServiceDocumentClient) getProcessorJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/processorJobs/{processorJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetProcessorJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/document-understanding/20221109/ProcessorJob/GetProcessorJob"
		err = common.PostProcessServiceError(err, "AIServiceDocument", "GetProcessorJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
