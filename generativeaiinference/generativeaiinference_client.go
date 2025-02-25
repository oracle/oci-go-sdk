// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Inference API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service inference API to access your custom model endpoints, or to try the out-of-the-box models to /EN/generative-ai-inference/latest/ChatResult/Chat, /EN/generative-ai-inference/latest/GenerateTextResult/GenerateText, /EN/generative-ai-inference/latest/SummarizeTextResult/SummarizeText, and /EN/generative-ai-inference/latest/EmbedTextResult/EmbedText.
// To use a Generative AI custom model for inference, you must first create an endpoint for that model. Use the /EN/generative-ai/latest/ to /EN/generative-ai/latest/Model/ by fine-tuning an out-of-the-box model, or a previous version of a custom model, using your own data. Fine-tune the custom model on a /EN/generative-ai/latest/DedicatedAiCluster/. Then, create a /EN/generative-ai/latest/DedicatedAiCluster/ with an Endpoint to host your custom model. For resource management in the Generative AI service, use the /EN/generative-ai/latest/.
// To learn more about the service, see the Generative AI documentation (https://docs.oracle.com/iaas/Content/generative-ai/home.htm).
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

// Chat Creates a response for the given conversation.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiinference/Chat.go.html to see an example of how to use Chat API.
// A default retry strategy applies to this operation Chat()
func (client GenerativeAiInferenceClient) Chat(ctx context.Context, request ChatRequest) (response ChatResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.chat, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChatResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChatResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChatResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChatResponse")
	}
	return
}

// chat implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiInferenceClient) chat(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/chat", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChatResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-inference/20231130/ChatResult/Chat"
		err = common.PostProcessServiceError(err, "GenerativeAiInference", "Chat", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EmbedText Produces embeddings for the inputs.
// An embedding is numeric representation of a piece of text. This text can be a phrase, a sentence, or one or more paragraphs. The Generative AI embedding model transforms each phrase, sentence, or paragraph that you input, into an array with 1024 numbers. You can use these embeddings for finding similarity in your input text such as finding phrases that are similar in context or category. Embeddings are mostly used for semantic searches where the search function focuses on the meaning of the text that it's searching through rather than finding results based on keywords.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiinference/EmbedText.go.html to see an example of how to use EmbedText API.
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
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiinference/GenerateText.go.html to see an example of how to use GenerateText API.
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

// RerankText Reranks the text responses based on the input documents and a prompt.
// Rerank assigns an index and a relevance score to each document, indicating which document is most related to the prompt.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiinference/RerankText.go.html to see an example of how to use RerankText API.
// A default retry strategy applies to this operation RerankText()
func (client GenerativeAiInferenceClient) RerankText(ctx context.Context, request RerankTextRequest) (response RerankTextResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rerankText, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RerankTextResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RerankTextResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RerankTextResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RerankTextResponse")
	}
	return
}

// rerankText implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiInferenceClient) rerankText(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/rerankText", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RerankTextResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-inference/20231130/RerankTextResult/RerankText"
		err = common.PostProcessServiceError(err, "GenerativeAiInference", "RerankText", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeText Summarizes the input text.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiinference/SummarizeText.go.html to see an example of how to use SummarizeText API.
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
