// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Speech API
//
// The OCI Speech Service harnesses the power of spoken language by allowing developers to easily convert file-based data containing human speech into highly accurate text transcriptions.
//

package aispeech

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//AIServiceSpeechClient a client for AIServiceSpeech
type AIServiceSpeechClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAIServiceSpeechClientWithConfigurationProvider Creates a new default AIServiceSpeech client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAIServiceSpeechClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AIServiceSpeechClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newAIServiceSpeechClientFromBaseClient(baseClient, provider)
}

// NewAIServiceSpeechClientWithOboToken Creates a new default AIServiceSpeech client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewAIServiceSpeechClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client AIServiceSpeechClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newAIServiceSpeechClientFromBaseClient(baseClient, configProvider)
}

func newAIServiceSpeechClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client AIServiceSpeechClient, err error) {
	// AIServiceSpeech service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("AIServiceSpeech"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = AIServiceSpeechClient{BaseClient: baseClient}
	client.BasePath = "20220101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AIServiceSpeechClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("aispeech", "https://speech.aiservice.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AIServiceSpeechClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *AIServiceSpeechClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelTranscriptionJob Canceling the job cancels all the tasks under it.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aispeech/CancelTranscriptionJob.go.html to see an example of how to use CancelTranscriptionJob API.
func (client AIServiceSpeechClient) CancelTranscriptionJob(ctx context.Context, request CancelTranscriptionJobRequest) (response CancelTranscriptionJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelTranscriptionJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelTranscriptionJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelTranscriptionJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelTranscriptionJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelTranscriptionJobResponse")
	}
	return
}

// cancelTranscriptionJob implements the OCIOperation interface (enables retrying operations)
func (client AIServiceSpeechClient) cancelTranscriptionJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transcriptionJobs/{transcriptionJobId}/actions/cancel", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelTranscriptionJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/speech/20220101/TranscriptionJob/CancelTranscriptionJob"
		err = common.PostProcessServiceError(err, "AIServiceSpeech", "CancelTranscriptionJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelTranscriptionTask Cancel Transcription Task
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aispeech/CancelTranscriptionTask.go.html to see an example of how to use CancelTranscriptionTask API.
func (client AIServiceSpeechClient) CancelTranscriptionTask(ctx context.Context, request CancelTranscriptionTaskRequest) (response CancelTranscriptionTaskResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelTranscriptionTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelTranscriptionTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelTranscriptionTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelTranscriptionTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelTranscriptionTaskResponse")
	}
	return
}

// cancelTranscriptionTask implements the OCIOperation interface (enables retrying operations)
func (client AIServiceSpeechClient) cancelTranscriptionTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transcriptionJobs/{transcriptionJobId}/transcriptionTasks/{transcriptionTaskId}/actions/cancel", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelTranscriptionTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/speech/20220101/TranscriptionTask/CancelTranscriptionTask"
		err = common.PostProcessServiceError(err, "AIServiceSpeech", "CancelTranscriptionTask", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeTranscriptionJobCompartment Moves a transcription Job resource into a different compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aispeech/ChangeTranscriptionJobCompartment.go.html to see an example of how to use ChangeTranscriptionJobCompartment API.
func (client AIServiceSpeechClient) ChangeTranscriptionJobCompartment(ctx context.Context, request ChangeTranscriptionJobCompartmentRequest) (response ChangeTranscriptionJobCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeTranscriptionJobCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeTranscriptionJobCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeTranscriptionJobCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeTranscriptionJobCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeTranscriptionJobCompartmentResponse")
	}
	return
}

// changeTranscriptionJobCompartment implements the OCIOperation interface (enables retrying operations)
func (client AIServiceSpeechClient) changeTranscriptionJobCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transcriptionJobs/{transcriptionJobId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeTranscriptionJobCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/speech/20220101/TranscriptionJob/ChangeTranscriptionJobCompartment"
		err = common.PostProcessServiceError(err, "AIServiceSpeech", "ChangeTranscriptionJobCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTranscriptionJob Creates a new Transcription Job.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aispeech/CreateTranscriptionJob.go.html to see an example of how to use CreateTranscriptionJob API.
func (client AIServiceSpeechClient) CreateTranscriptionJob(ctx context.Context, request CreateTranscriptionJobRequest) (response CreateTranscriptionJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTranscriptionJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTranscriptionJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTranscriptionJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTranscriptionJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTranscriptionJobResponse")
	}
	return
}

// createTranscriptionJob implements the OCIOperation interface (enables retrying operations)
func (client AIServiceSpeechClient) createTranscriptionJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/transcriptionJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTranscriptionJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/speech/20220101/TranscriptionJob/CreateTranscriptionJob"
		err = common.PostProcessServiceError(err, "AIServiceSpeech", "CreateTranscriptionJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTranscriptionJob Gets a Transcription Job by identifier
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aispeech/GetTranscriptionJob.go.html to see an example of how to use GetTranscriptionJob API.
func (client AIServiceSpeechClient) GetTranscriptionJob(ctx context.Context, request GetTranscriptionJobRequest) (response GetTranscriptionJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTranscriptionJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTranscriptionJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTranscriptionJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTranscriptionJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTranscriptionJobResponse")
	}
	return
}

// getTranscriptionJob implements the OCIOperation interface (enables retrying operations)
func (client AIServiceSpeechClient) getTranscriptionJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transcriptionJobs/{transcriptionJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTranscriptionJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/speech/20220101/TranscriptionJob/GetTranscriptionJob"
		err = common.PostProcessServiceError(err, "AIServiceSpeech", "GetTranscriptionJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTranscriptionTask Gets a Transcription Task by identifier
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aispeech/GetTranscriptionTask.go.html to see an example of how to use GetTranscriptionTask API.
func (client AIServiceSpeechClient) GetTranscriptionTask(ctx context.Context, request GetTranscriptionTaskRequest) (response GetTranscriptionTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTranscriptionTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTranscriptionTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTranscriptionTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTranscriptionTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTranscriptionTaskResponse")
	}
	return
}

// getTranscriptionTask implements the OCIOperation interface (enables retrying operations)
func (client AIServiceSpeechClient) getTranscriptionTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transcriptionJobs/{transcriptionJobId}/transcriptionTasks/{transcriptionTaskId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTranscriptionTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/speech/20220101/TranscriptionTask/GetTranscriptionTask"
		err = common.PostProcessServiceError(err, "AIServiceSpeech", "GetTranscriptionTask", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTranscriptionJobs Returns a list of Transcription Jobs.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aispeech/ListTranscriptionJobs.go.html to see an example of how to use ListTranscriptionJobs API.
func (client AIServiceSpeechClient) ListTranscriptionJobs(ctx context.Context, request ListTranscriptionJobsRequest) (response ListTranscriptionJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTranscriptionJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTranscriptionJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTranscriptionJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTranscriptionJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTranscriptionJobsResponse")
	}
	return
}

// listTranscriptionJobs implements the OCIOperation interface (enables retrying operations)
func (client AIServiceSpeechClient) listTranscriptionJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transcriptionJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTranscriptionJobsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/speech/20220101/TranscriptionJob/ListTranscriptionJobs"
		err = common.PostProcessServiceError(err, "AIServiceSpeech", "ListTranscriptionJobs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTranscriptionTasks Returns a list of Transcription Tasks.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aispeech/ListTranscriptionTasks.go.html to see an example of how to use ListTranscriptionTasks API.
func (client AIServiceSpeechClient) ListTranscriptionTasks(ctx context.Context, request ListTranscriptionTasksRequest) (response ListTranscriptionTasksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTranscriptionTasks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTranscriptionTasksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTranscriptionTasksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTranscriptionTasksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTranscriptionTasksResponse")
	}
	return
}

// listTranscriptionTasks implements the OCIOperation interface (enables retrying operations)
func (client AIServiceSpeechClient) listTranscriptionTasks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transcriptionJobs/{transcriptionJobId}/transcriptionTasks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTranscriptionTasksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/speech/20220101/TranscriptionTask/ListTranscriptionTasks"
		err = common.PostProcessServiceError(err, "AIServiceSpeech", "ListTranscriptionTasks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTranscriptionJob Updates the Transcription Job
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aispeech/UpdateTranscriptionJob.go.html to see an example of how to use UpdateTranscriptionJob API.
func (client AIServiceSpeechClient) UpdateTranscriptionJob(ctx context.Context, request UpdateTranscriptionJobRequest) (response UpdateTranscriptionJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTranscriptionJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTranscriptionJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTranscriptionJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTranscriptionJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTranscriptionJobResponse")
	}
	return
}

// updateTranscriptionJob implements the OCIOperation interface (enables retrying operations)
func (client AIServiceSpeechClient) updateTranscriptionJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/transcriptionJobs/{transcriptionJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTranscriptionJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/speech/20220101/TranscriptionJob/UpdateTranscriptionJob"
		err = common.PostProcessServiceError(err, "AIServiceSpeech", "UpdateTranscriptionJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
