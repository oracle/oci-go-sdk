// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Inference API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service inference API to access your custom model endpoints, or to try the out-of-the-box models to GenerateText, SummarizeText, and EmbedText.
// To use a Generative AI custom model for inference, you must first create an endpoint for that model. Use the Generative AI service management API (https://docs.cloud.oracle.com/#/en/generative-ai/latest/) to Model by fine-tuning an out-of-the-box model, or a previous version of a custom model, using your own data. Fine-tune the custom model on a  DedicatedAiCluster. Then, create a DedicatedAiCluster with an Endpoint to host your custom model. For resource management in the Generative AI service, use the Generative AI service management API (https://docs.cloud.oracle.com/#/en/generative-ai/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeaiinference

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// GenerativeAiInferenceClient a client for GenerativeAiInference
type GenerativeAiInferenceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewGenerativeAiInferenceClientWithConfigurationProvider Creates a new default GenerativeAiInference client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewGenerativeAiInferenceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client GenerativeAiInferenceClient, err error) {
	if enabled := common.CheckForEnabledServices("generativeaiinference"); !enabled {
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
	return newGenerativeAiInferenceClientFromBaseClient(baseClient, provider)
}

// NewGenerativeAiInferenceClientWithOboToken Creates a new default GenerativeAiInference client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewGenerativeAiInferenceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client GenerativeAiInferenceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newGenerativeAiInferenceClientFromBaseClient(baseClient, configProvider)
}

func newGenerativeAiInferenceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client GenerativeAiInferenceClient, err error) {
	// GenerativeAiInference service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("GenerativeAiInference"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = GenerativeAiInferenceClient{BaseClient: baseClient}
	client.BasePath = "20231130"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *GenerativeAiInferenceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("generativeaiinference", "https://inference.generativeai.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *GenerativeAiInferenceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *GenerativeAiInferenceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// EmbedText Produces embeddings for the inputs.
// An embedding is numeric representation of a piece of text. This text can be a phrase, a sentence, or one or more paragraphs. The Generative AI embedding model transforms each phrase, sentence, or paragraph that you input, into an array with 1024 numbers. You can use these embeddings for finding similarity in your input text such as finding phrases that are similar in context or category. Embeddings are mostly used for semantic searches where the search function focuses on the meaning of the text that it's searching through rather than finding results based on keywords.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiinference/EmbedText.go.html to see an example of how to use EmbedText API.
// A default retry strategy applies to this operation EmbedText()
func (client GenerativeAiInferenceClient) EmbedText(ctx context.Context, request EmbedTextRequest) (response EmbedTextResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.embedText, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EmbedTextResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EmbedTextResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EmbedTextResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EmbedTextResponse")
	}
	return
}

// embedText implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiInferenceClient) embedText(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/embedText", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EmbedTextResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-inference/20231130/EmbedTextResult/EmbedText"
		err = common.PostProcessServiceError(err, "GenerativeAiInference", "EmbedText", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateText Generates a text response based on the user prompt.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiinference/GenerateText.go.html to see an example of how to use GenerateText API.
// A default retry strategy applies to this operation GenerateText()
func (client GenerativeAiInferenceClient) GenerateText(ctx context.Context, request GenerateTextRequest) (response GenerateTextResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateText, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateTextResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateTextResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateTextResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateTextResponse")
	}
	return
}

// generateText implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiInferenceClient) generateText(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/generateText", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateTextResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-inference/20231130/GenerateTextResult/GenerateText"
		err = common.PostProcessServiceError(err, "GenerativeAiInference", "GenerateText", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeText Summarizes the input text.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiinference/SummarizeText.go.html to see an example of how to use SummarizeText API.
// A default retry strategy applies to this operation SummarizeText()
func (client GenerativeAiInferenceClient) SummarizeText(ctx context.Context, request SummarizeTextRequest) (response SummarizeTextResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.summarizeText, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeTextResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeTextResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeTextResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeTextResponse")
	}
	return
}

// summarizeText implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiInferenceClient) summarizeText(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/summarizeText", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeTextResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-inference/20231130/SummarizeTextResult/SummarizeText"
		err = common.PostProcessServiceError(err, "GenerativeAiInference", "SummarizeText", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
