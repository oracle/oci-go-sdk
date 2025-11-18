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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GenericChatRequest Details for the chat request.
type GenericChatRequest struct {

	// The series of messages in a chat request. Includes the previous messages in a conversation. Each message includes a role (`USER` or the `CHATBOT`) and content.
	Messages []Message `mandatory:"false" json:"messages"`

	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format, and querying for objects via API or the dashboard.
	// Keys are strings with a maximum length of 64 characters. Values are strings with a maximum length of 512 characters.
	Metadata *interface{} `mandatory:"false" json:"metadata"`

	// Whether to stream back partial progress. If set to true, as tokens become available, they are sent as data-only server-sent events.
	IsStream *bool `mandatory:"false" json:"isStream"`

	StreamOptions *StreamOptions `mandatory:"false" json:"streamOptions"`

	// The number of of generated texts that will be returned.
	NumGenerations *int `mandatory:"false" json:"numGenerations"`

	// If specified, the backend will make a best effort to sample tokens deterministically, so that repeated requests with the same seed and parameters yield the same result. However, determinism cannot be fully guaranteed.
	Seed *int `mandatory:"false" json:"seed"`

	// Whether to include the user prompt in the response. Applies only to non-stream results.
	IsEcho *bool `mandatory:"false" json:"isEcho"`

	// An integer that sets up the model to use only the top k most likely tokens in the generated output. A higher k introduces more randomness into the output making the output text sound more natural. Default value is -1 which means to consider all tokens. Setting to 0 disables this method and considers all tokens.
	// If also using top p, then the model considers only the top tokens whose probabilities add up to p percent and ignores the rest of the k tokens. For example, if k is 20, but the probabilities of the top 10 add up to .75, then only the top 10 tokens are chosen.
	TopK *int `mandatory:"false" json:"topK"`

	// If set to a probability 0.0 < p < 1.0, it ensures that only the most likely tokens, with total probability mass of p, are considered for generation at each step.
	// To eliminate tokens with low likelihood, assign p a minimum percentage for the next token's likelihood. For example, when p is set to 0.75, the model eliminates the bottom 25 percent for the next token. Set to 1 to consider all tokens and set to 0 to disable. If both k and p are enabled, p acts after k.
	TopP *float64 `mandatory:"false" json:"topP"`

	// A number that sets the randomness of the generated output. A lower temperature means a less random generations.
	// Use lower numbers for tasks with a correct answer such as question answering or summarizing. High temperatures can generate hallucinations or factually incorrect information. Start with temperatures lower than 1.0 and increase the temperature for more creative outputs, as you regenerate the prompts to refine the outputs.
	Temperature *float64 `mandatory:"false" json:"temperature"`

	// To reduce repetitiveness of generated tokens, this number penalizes new tokens based on their frequency in the generated text so far. Values > 0 encourage the model to use new tokens and values < 0 encourage the model to repeat tokens. Set to 0 to disable.
	FrequencyPenalty *float64 `mandatory:"false" json:"frequencyPenalty"`

	// To reduce repetitiveness of generated tokens, this number penalizes new tokens based on whether they've appeared in the generated text so far. Values > 0 encourage the model to use new tokens and values < 0 encourage the model to repeat tokens.
	// Similar to frequency penalty, a penalty is applied to previously present tokens, except that this penalty is applied equally to all tokens that have already appeared, regardless of how many times they've appeared. Set to 0 to disable.
	PresencePenalty *float64 `mandatory:"false" json:"presencePenalty"`

	// List of strings that stop the generation if they are generated for the response text. The returned output will not contain the stop strings.
	Stop []string `mandatory:"false" json:"stop"`

	// Includes the logarithmic probabilities for the most likely output tokens and the chosen tokens.
	// For example, if the log probability is 5, the API returns a list of the 5 most likely tokens. The API returns the log probability of the sampled token, so there might be up to logprobs+1 elements in the response.
	LogProbs *int `mandatory:"false" json:"logProbs"`

	// The maximum number of tokens that can be generated per output sequence. The token count of your prompt plus maxTokens must not exceed the model's context length. For on-demand inferencing, the response length is capped at 4,000 tokens for each run.
	MaxTokens *int `mandatory:"false" json:"maxTokens"`

	// An upper bound for the number of tokens that can be generated for a completion, including visible output tokens and reasoning tokens.
	MaxCompletionTokens *int `mandatory:"false" json:"maxCompletionTokens"`

	// Modifies the likelihood of specified tokens that appear in the completion.
	// Example: '{"6395": 2, "8134": 1, "21943": 0.5, "5923": -100}'
	LogitBias *interface{} `mandatory:"false" json:"logitBias"`

	Prediction Prediction `mandatory:"false" json:"prediction"`

	ResponseFormat ResponseFormat `mandatory:"false" json:"responseFormat"`

	ToolChoice ToolChoice `mandatory:"false" json:"toolChoice"`

	// Whether to enable parallel function calling during tool use.
	IsParallelToolCalls *bool `mandatory:"false" json:"isParallelToolCalls"`

	// A list of tools the model may call. Use this to provide a list of functions the model may generate JSON inputs for. A max of 128 functions are supported.
	Tools []ToolDefinition `mandatory:"false" json:"tools"`

	WebSearchOptions *WebSearchOptions `mandatory:"false" json:"webSearchOptions"`

	// Constrains effort on reasoning for reasoning models. Currently supported values are minimal, low, medium, and high. Reducing reasoning effort can result in faster responses and fewer tokens used on reasoning in a response.
	ReasoningEffort GenericChatRequestReasoningEffortEnum `mandatory:"false" json:"reasoningEffort,omitempty"`

	// Constrains the verbosity of the model's response. Lower values will result in more concise responses, while higher values will result in more verbose responses.
	Verbosity GenericChatRequestVerbosityEnum `mandatory:"false" json:"verbosity,omitempty"`

	// Specifies the processing type used for serving the request.
	ServiceTier GenericChatRequestServiceTierEnum `mandatory:"false" json:"serviceTier,omitempty"`
}

func (m GenericChatRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenericChatRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGenericChatRequestReasoningEffortEnum(string(m.ReasoningEffort)); !ok && m.ReasoningEffort != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReasoningEffort: %s. Supported values are: %s.", m.ReasoningEffort, strings.Join(GetGenericChatRequestReasoningEffortEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGenericChatRequestVerbosityEnum(string(m.Verbosity)); !ok && m.Verbosity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Verbosity: %s. Supported values are: %s.", m.Verbosity, strings.Join(GetGenericChatRequestVerbosityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGenericChatRequestServiceTierEnum(string(m.ServiceTier)); !ok && m.ServiceTier != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceTier: %s. Supported values are: %s.", m.ServiceTier, strings.Join(GetGenericChatRequestServiceTierEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GenericChatRequest) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGenericChatRequest GenericChatRequest
	s := struct {
		DiscriminatorParam string `json:"apiFormat"`
		MarshalTypeGenericChatRequest
	}{
		"GENERIC",
		(MarshalTypeGenericChatRequest)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *GenericChatRequest) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Messages            []message                             `json:"messages"`
		ReasoningEffort     GenericChatRequestReasoningEffortEnum `json:"reasoningEffort"`
		Verbosity           GenericChatRequestVerbosityEnum       `json:"verbosity"`
		Metadata            *interface{}                          `json:"metadata"`
		IsStream            *bool                                 `json:"isStream"`
		StreamOptions       *StreamOptions                        `json:"streamOptions"`
		NumGenerations      *int                                  `json:"numGenerations"`
		Seed                *int                                  `json:"seed"`
		IsEcho              *bool                                 `json:"isEcho"`
		TopK                *int                                  `json:"topK"`
		TopP                *float64                              `json:"topP"`
		Temperature         *float64                              `json:"temperature"`
		FrequencyPenalty    *float64                              `json:"frequencyPenalty"`
		PresencePenalty     *float64                              `json:"presencePenalty"`
		Stop                []string                              `json:"stop"`
		LogProbs            *int                                  `json:"logProbs"`
		MaxTokens           *int                                  `json:"maxTokens"`
		MaxCompletionTokens *int                                  `json:"maxCompletionTokens"`
		LogitBias           *interface{}                          `json:"logitBias"`
		Prediction          prediction                            `json:"prediction"`
		ResponseFormat      responseformat                        `json:"responseFormat"`
		ToolChoice          toolchoice                            `json:"toolChoice"`
		IsParallelToolCalls *bool                                 `json:"isParallelToolCalls"`
		Tools               []tooldefinition                      `json:"tools"`
		WebSearchOptions    *WebSearchOptions                     `json:"webSearchOptions"`
		ServiceTier         GenericChatRequestServiceTierEnum     `json:"serviceTier"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Messages = make([]Message, len(model.Messages))
	for i, n := range model.Messages {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Messages[i] = nn.(Message)
		} else {
			m.Messages[i] = nil
		}
	}
	m.ReasoningEffort = model.ReasoningEffort

	m.Verbosity = model.Verbosity

	m.Metadata = model.Metadata

	m.IsStream = model.IsStream

	m.StreamOptions = model.StreamOptions

	m.NumGenerations = model.NumGenerations

	m.Seed = model.Seed

	m.IsEcho = model.IsEcho

	m.TopK = model.TopK

	m.TopP = model.TopP

	m.Temperature = model.Temperature

	m.FrequencyPenalty = model.FrequencyPenalty

	m.PresencePenalty = model.PresencePenalty

	m.Stop = make([]string, len(model.Stop))
	copy(m.Stop, model.Stop)
	m.LogProbs = model.LogProbs

	m.MaxTokens = model.MaxTokens

	m.MaxCompletionTokens = model.MaxCompletionTokens

	m.LogitBias = model.LogitBias

	nn, e = model.Prediction.UnmarshalPolymorphicJSON(model.Prediction.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Prediction = nn.(Prediction)
	} else {
		m.Prediction = nil
	}

	nn, e = model.ResponseFormat.UnmarshalPolymorphicJSON(model.ResponseFormat.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResponseFormat = nn.(ResponseFormat)
	} else {
		m.ResponseFormat = nil
	}

	nn, e = model.ToolChoice.UnmarshalPolymorphicJSON(model.ToolChoice.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ToolChoice = nn.(ToolChoice)
	} else {
		m.ToolChoice = nil
	}

	m.IsParallelToolCalls = model.IsParallelToolCalls

	m.Tools = make([]ToolDefinition, len(model.Tools))
	for i, n := range model.Tools {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Tools[i] = nn.(ToolDefinition)
		} else {
			m.Tools[i] = nil
		}
	}
	m.WebSearchOptions = model.WebSearchOptions

	m.ServiceTier = model.ServiceTier

	return
}

// GenericChatRequestReasoningEffortEnum Enum with underlying type: string
type GenericChatRequestReasoningEffortEnum string

// Set of constants representing the allowable values for GenericChatRequestReasoningEffortEnum
const (
	GenericChatRequestReasoningEffortMinimal GenericChatRequestReasoningEffortEnum = "MINIMAL"
	GenericChatRequestReasoningEffortLow     GenericChatRequestReasoningEffortEnum = "LOW"
	GenericChatRequestReasoningEffortMedium  GenericChatRequestReasoningEffortEnum = "MEDIUM"
	GenericChatRequestReasoningEffortHigh    GenericChatRequestReasoningEffortEnum = "HIGH"
)

var mappingGenericChatRequestReasoningEffortEnum = map[string]GenericChatRequestReasoningEffortEnum{
	"MINIMAL": GenericChatRequestReasoningEffortMinimal,
	"LOW":     GenericChatRequestReasoningEffortLow,
	"MEDIUM":  GenericChatRequestReasoningEffortMedium,
	"HIGH":    GenericChatRequestReasoningEffortHigh,
}

var mappingGenericChatRequestReasoningEffortEnumLowerCase = map[string]GenericChatRequestReasoningEffortEnum{
	"minimal": GenericChatRequestReasoningEffortMinimal,
	"low":     GenericChatRequestReasoningEffortLow,
	"medium":  GenericChatRequestReasoningEffortMedium,
	"high":    GenericChatRequestReasoningEffortHigh,
}

// GetGenericChatRequestReasoningEffortEnumValues Enumerates the set of values for GenericChatRequestReasoningEffortEnum
func GetGenericChatRequestReasoningEffortEnumValues() []GenericChatRequestReasoningEffortEnum {
	values := make([]GenericChatRequestReasoningEffortEnum, 0)
	for _, v := range mappingGenericChatRequestReasoningEffortEnum {
		values = append(values, v)
	}
	return values
}

// GetGenericChatRequestReasoningEffortEnumStringValues Enumerates the set of values in String for GenericChatRequestReasoningEffortEnum
func GetGenericChatRequestReasoningEffortEnumStringValues() []string {
	return []string{
		"MINIMAL",
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

// GetMappingGenericChatRequestReasoningEffortEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenericChatRequestReasoningEffortEnum(val string) (GenericChatRequestReasoningEffortEnum, bool) {
	enum, ok := mappingGenericChatRequestReasoningEffortEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GenericChatRequestVerbosityEnum Enum with underlying type: string
type GenericChatRequestVerbosityEnum string

// Set of constants representing the allowable values for GenericChatRequestVerbosityEnum
const (
	GenericChatRequestVerbosityLow    GenericChatRequestVerbosityEnum = "LOW"
	GenericChatRequestVerbosityMedium GenericChatRequestVerbosityEnum = "MEDIUM"
	GenericChatRequestVerbosityHigh   GenericChatRequestVerbosityEnum = "HIGH"
)

var mappingGenericChatRequestVerbosityEnum = map[string]GenericChatRequestVerbosityEnum{
	"LOW":    GenericChatRequestVerbosityLow,
	"MEDIUM": GenericChatRequestVerbosityMedium,
	"HIGH":   GenericChatRequestVerbosityHigh,
}

var mappingGenericChatRequestVerbosityEnumLowerCase = map[string]GenericChatRequestVerbosityEnum{
	"low":    GenericChatRequestVerbosityLow,
	"medium": GenericChatRequestVerbosityMedium,
	"high":   GenericChatRequestVerbosityHigh,
}

// GetGenericChatRequestVerbosityEnumValues Enumerates the set of values for GenericChatRequestVerbosityEnum
func GetGenericChatRequestVerbosityEnumValues() []GenericChatRequestVerbosityEnum {
	values := make([]GenericChatRequestVerbosityEnum, 0)
	for _, v := range mappingGenericChatRequestVerbosityEnum {
		values = append(values, v)
	}
	return values
}

// GetGenericChatRequestVerbosityEnumStringValues Enumerates the set of values in String for GenericChatRequestVerbosityEnum
func GetGenericChatRequestVerbosityEnumStringValues() []string {
	return []string{
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

// GetMappingGenericChatRequestVerbosityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenericChatRequestVerbosityEnum(val string) (GenericChatRequestVerbosityEnum, bool) {
	enum, ok := mappingGenericChatRequestVerbosityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GenericChatRequestServiceTierEnum Enum with underlying type: string
type GenericChatRequestServiceTierEnum string

// Set of constants representing the allowable values for GenericChatRequestServiceTierEnum
const (
	GenericChatRequestServiceTierAuto     GenericChatRequestServiceTierEnum = "AUTO"
	GenericChatRequestServiceTierDefault  GenericChatRequestServiceTierEnum = "DEFAULT"
	GenericChatRequestServiceTierPriority GenericChatRequestServiceTierEnum = "PRIORITY"
)

var mappingGenericChatRequestServiceTierEnum = map[string]GenericChatRequestServiceTierEnum{
	"AUTO":     GenericChatRequestServiceTierAuto,
	"DEFAULT":  GenericChatRequestServiceTierDefault,
	"PRIORITY": GenericChatRequestServiceTierPriority,
}

var mappingGenericChatRequestServiceTierEnumLowerCase = map[string]GenericChatRequestServiceTierEnum{
	"auto":     GenericChatRequestServiceTierAuto,
	"default":  GenericChatRequestServiceTierDefault,
	"priority": GenericChatRequestServiceTierPriority,
}

// GetGenericChatRequestServiceTierEnumValues Enumerates the set of values for GenericChatRequestServiceTierEnum
func GetGenericChatRequestServiceTierEnumValues() []GenericChatRequestServiceTierEnum {
	values := make([]GenericChatRequestServiceTierEnum, 0)
	for _, v := range mappingGenericChatRequestServiceTierEnum {
		values = append(values, v)
	}
	return values
}

// GetGenericChatRequestServiceTierEnumStringValues Enumerates the set of values in String for GenericChatRequestServiceTierEnum
func GetGenericChatRequestServiceTierEnumStringValues() []string {
	return []string{
		"AUTO",
		"DEFAULT",
		"PRIORITY",
	}
}

// GetMappingGenericChatRequestServiceTierEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenericChatRequestServiceTierEnum(val string) (GenericChatRequestServiceTierEnum, bool) {
	enum, ok := mappingGenericChatRequestServiceTierEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
