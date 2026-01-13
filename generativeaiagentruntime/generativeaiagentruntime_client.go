// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Agents Client API
//
// OCI Generative AI Agents is a fully managed service that combines the power of large language models (LLMs) with an intelligent retrieval system to create contextually relevant answers by searching your knowledge base, making your AI applications smart and efficient.
// OCI Generative AI Agents supports several ways to onboard your data and then allows you and your customers to interact with your data using a chat interface or API.
// Use the Generative AI Agents Client API to create and manage client chat sessions. A session represents an interactive conversation initiated by a user through an API to engage with an agent. It involves a series of exchanges where the user sends queries or prompts, and the agent responds with relevant information, actions, or assistance based on the user's input. The session persists for the duration of the interaction, maintaining context and continuity to provide coherent and meaningful responses throughout the conversation.
// For creating and managing agents, knowledge bases, data sources, endpoints, and data ingestion jobs see the /EN/generative-ai-agents/latest/.
// To learn more about the service, see the Generative AI Agents documentation (https://docs.oracle.com/iaas/Content/generative-ai-agents/home.htm).
//

package generativeaiagentruntime

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// GenerativeAiAgentRuntimeClient a client for GenerativeAiAgentRuntime
type GenerativeAiAgentRuntimeClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewGenerativeAiAgentRuntimeClientWithConfigurationProvider Creates a new default GenerativeAiAgentRuntime client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewGenerativeAiAgentRuntimeClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client GenerativeAiAgentRuntimeClient, err error) {
	if enabled := common.CheckForEnabledServices("generativeaiagentruntime"); !enabled {
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
	return newGenerativeAiAgentRuntimeClientFromBaseClient(baseClient, provider)
}

// NewGenerativeAiAgentRuntimeClientWithOboToken Creates a new default GenerativeAiAgentRuntime client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewGenerativeAiAgentRuntimeClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client GenerativeAiAgentRuntimeClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newGenerativeAiAgentRuntimeClientFromBaseClient(baseClient, configProvider)
}

func newGenerativeAiAgentRuntimeClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client GenerativeAiAgentRuntimeClient, err error) {
	// GenerativeAiAgentRuntime service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("GenerativeAiAgentRuntime"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = GenerativeAiAgentRuntimeClient{BaseClient: baseClient}
	client.BasePath = "20240531"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *GenerativeAiAgentRuntimeClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("generativeaiagentruntime", "https://agent-runtime.generativeai.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *GenerativeAiAgentRuntimeClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *GenerativeAiAgentRuntimeClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// Chat Chat on endpoint with provided messages.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiagentruntime/Chat.go.html to see an example of how to use Chat API.
// A default retry strategy applies to this operation Chat()
func (client GenerativeAiAgentRuntimeClient) Chat(ctx context.Context, request ChatRequest) (response ChatResponse, err error) {
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
func (client GenerativeAiAgentRuntimeClient) chat(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/agentEndpoints/{agentEndpointId}/actions/chat", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChatResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-agents-client/20240531/AgentEndpoint/Chat"
		err = common.PostProcessServiceError(err, "GenerativeAiAgentRuntime", "Chat", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSession A session represents an interactive conversation initiated by a user through an API to engage with an  agent. It involves a series of exchanges where the user sends queries or prompts, and the agent responds with relevant information, actions, or assistance based on the user's input. The session persists for the duration of the interaction, maintaining context and continuity to provide coherent and meaningful responses throughout the conversation.Creates an agent session.
// Use this API to create an agent session.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiagentruntime/CreateSession.go.html to see an example of how to use CreateSession API.
// A default retry strategy applies to this operation CreateSession()
func (client GenerativeAiAgentRuntimeClient) CreateSession(ctx context.Context, request CreateSessionRequest) (response CreateSessionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSession, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSessionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSessionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSessionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSessionResponse")
	}
	return
}

// createSession implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiAgentRuntimeClient) createSession(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/agentEndpoints/{agentEndpointId}/sessions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSessionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-agents-client/20240531/Session/CreateSession"
		err = common.PostProcessServiceError(err, "GenerativeAiAgentRuntime", "CreateSession", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSession Delete a session and all its associated information.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiagentruntime/DeleteSession.go.html to see an example of how to use DeleteSession API.
// A default retry strategy applies to this operation DeleteSession()
func (client GenerativeAiAgentRuntimeClient) DeleteSession(ctx context.Context, request DeleteSessionRequest) (response DeleteSessionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSession, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSessionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSessionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSessionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSessionResponse")
	}
	return
}

// deleteSession implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiAgentRuntimeClient) deleteSession(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/agentEndpoints/{agentEndpointId}/sessions/{sessionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSessionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-agents-client/20240531/Session/DeleteSession"
		err = common.PostProcessServiceError(err, "GenerativeAiAgentRuntime", "DeleteSession", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSession Return the session resource identified by the session ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiagentruntime/GetSession.go.html to see an example of how to use GetSession API.
// A default retry strategy applies to this operation GetSession()
func (client GenerativeAiAgentRuntimeClient) GetSession(ctx context.Context, request GetSessionRequest) (response GetSessionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSession, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSessionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSessionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSessionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSessionResponse")
	}
	return
}

// getSession implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiAgentRuntimeClient) getSession(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/agentEndpoints/{agentEndpointId}/sessions/{sessionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSessionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-agents-client/20240531/Session/GetSession"
		err = common.PostProcessServiceError(err, "GenerativeAiAgentRuntime", "GetSession", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetrieveMetadata Returns metadata of provided knowledgeBase. Return available metadata with information of field names, their types, supported operations, and possible values.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiagentruntime/RetrieveMetadata.go.html to see an example of how to use RetrieveMetadata API.
// A default retry strategy applies to this operation RetrieveMetadata()
func (client GenerativeAiAgentRuntimeClient) RetrieveMetadata(ctx context.Context, request RetrieveMetadataRequest) (response RetrieveMetadataResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.retrieveMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetrieveMetadataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetrieveMetadataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetrieveMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetrieveMetadataResponse")
	}
	return
}

// retrieveMetadata implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiAgentRuntimeClient) retrieveMetadata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/knowledgeBases/{knowledgeBaseId}/actions/retrieveMetadata", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetrieveMetadataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-agents-client/20240531/KnowledgeBaseMetadataSummary/RetrieveMetadata"
		err = common.PostProcessServiceError(err, "GenerativeAiAgentRuntime", "RetrieveMetadata", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSession Update session metadata, including but not limited to description, tags.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaiagentruntime/UpdateSession.go.html to see an example of how to use UpdateSession API.
// A default retry strategy applies to this operation UpdateSession()
func (client GenerativeAiAgentRuntimeClient) UpdateSession(ctx context.Context, request UpdateSessionRequest) (response UpdateSessionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSession, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSessionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSessionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSessionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSessionResponse")
	}
	return
}

// updateSession implements the OCIOperation interface (enables retrying operations)
func (client GenerativeAiAgentRuntimeClient) updateSession(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/agentEndpoints/{agentEndpointId}/sessions/{sessionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSessionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/generative-ai-agents-client/20240531/Session/UpdateSession"
		err = common.PostProcessServiceError(err, "GenerativeAiAgentRuntime", "UpdateSession", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
