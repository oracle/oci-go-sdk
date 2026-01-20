// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CohereChatRequestV2 Details for the chat request for Cohere models.
type CohereChatRequestV2 struct {

	// A list of chat messages in chronological order, representing a conversation between the user and the model.
	Messages []CohereMessageV2 `mandatory:"true" json:"messages"`

	// A list of relevant documents that the model can refer to for generating grounded responses to the user's requests.
	// Some example keys that you can add to the dictionary are "text", "author", and "date". Keep the total word count of the strings in the dictionary to 300 words or less.
	Documents []interface{} `mandatory:"false" json:"documents"`

	CitationOptions *CitationOptionsV2 `mandatory:"false" json:"citationOptions"`

	// A list of available tools (functions) that the model may suggest invoking before producing a text response.
	Tools []CohereToolV2 `mandatory:"false" json:"tools"`

	// When set to true, tool calls in the Assistant message will be forced to follow the tool definition strictly. Note:The first few requests with a new set of tools will take longer to process.
	IsStrictToolsEnabled *bool `mandatory:"false" json:"isStrictToolsEnabled"`

	// The log probabilities of the generated tokens will be included in the response.
	IsLogProbsEnabled *bool `mandatory:"false" json:"isLogProbsEnabled"`

	Thinking *CohereThinkingV2 `mandatory:"false" json:"thinking"`

	ResponseFormat CohereResponseFormat `mandatory:"false" json:"responseFormat"`

	// When set to true, the response contains only a list of generated search queries without the search results and the model will not respond to the user's message.
	IsSearchQueriesOnly *bool `mandatory:"false" json:"isSearchQueriesOnly"`

	StreamOptions *StreamOptions `mandatory:"false" json:"streamOptions"`

	// Whether to stream the partial progress of the model's response. When set to true, as tokens become available, they are sent as data-only server-sent events.
	IsStream *bool `mandatory:"false" json:"isStream"`

	// The maximum number of output tokens that the model will generate for the response. The token count of your prompt plus maxTokens must not exceed the model's context length. For on-demand inferencing, the response length is capped at 4,000 tokens for each run.
	MaxTokens *int `mandatory:"false" json:"maxTokens"`

	// A number that sets the randomness of the generated output. A lower temperature means less random generations.
	// Use lower numbers for tasks such as question answering or summarizing. High temperatures can generate hallucinations or factually incorrect information. Start with temperatures lower than 1.0 and increase the temperature for more creative outputs, as you regenerate the prompts to refine the outputs.
	Temperature *float64 `mandatory:"false" json:"temperature"`

	// A sampling method in which the model chooses the next token randomly from the top k most likely tokens. A higher value for k generates more random output, which makes the output text sound more natural. The default value for k is 0 which disables this method and considers all tokens. To set a number for the likely tokens, choose an integer between 1 and 500.
	// If also using top p, then the model considers only the top tokens whose probabilities add up to p percent and ignores the rest of the k tokens. For example, if k is 20 but only the probabilities of the top 10 add up to the value of p, then only the top 10 tokens are chosen.
	TopK *int `mandatory:"false" json:"topK"`

	// If set to a probability 0.0 < p < 1.0, it ensures that only the most likely tokens, with total probability mass of p, are considered for generation at each step.
	// To eliminate tokens with low likelihood, assign p a minimum percentage for the next token's likelihood. For example, when p is set to 0.75, the model eliminates the bottom 25 percent for the next token. Set to 1.0 to consider all tokens and set to 0 to disable. If both k and p are enabled, p acts after k.
	TopP *float64 `mandatory:"false" json:"topP"`

	// To reduce repetitiveness of generated tokens, this number penalizes new tokens based on their frequency in the generated text so far. Greater numbers encourage the model to use new tokens, while lower numbers encourage the model to repeat the tokens. Set to 0 to disable.
	FrequencyPenalty *float64 `mandatory:"false" json:"frequencyPenalty"`

	// To reduce repetitiveness of generated tokens, this number penalizes new tokens based on whether they've appeared in the generated text so far. Greater numbers encourage the model to use new tokens, while lower numbers encourage the model to repeat the tokens.
	// Similar to frequency penalty, a penalty is applied to previously present tokens, except that this penalty is applied equally to all tokens that have already appeared, regardless of how many times they've appeared. Set to 0 to disable.
	PresencePenalty *float64 `mandatory:"false" json:"presencePenalty"`

	// If specified, the backend will make a best effort to sample tokens deterministically, so that repeated requests with the same seed and parameters yield the same result. However, determinism cannot be fully guaranteed.
	Seed *int `mandatory:"false" json:"seed"`

	// Stop the model generation when it reaches a stop sequence defined in this parameter.
	StopSequences []string `mandatory:"false" json:"stopSequences"`

	// The priority of the request (lower means earlier handling; default 0 highest priority). Higher priority requests are handled first, and dropped last when the system is under load.
	Priority *int `mandatory:"false" json:"priority"`

	// When enabled, the user’s `message` will be sent to the model without any preprocessing.
	IsRawPrompting *bool `mandatory:"false" json:"isRawPrompting"`

	// Used to control whether or not the model will be forced to use a tool when answering. When REQUIRED is specified, the model will be forced to use at least one of the user-defined tools, and the tools parameter must be passed in the request. When NONE is specified, the model will be forced not to use one of the specified tools, and give a direct response. If tool_choice isn’t specified, then the model is free to choose whether to use the specified tools or not. Note:This parameter is only compatible with models Command-r7b and newer.
	ToolsChoice CohereChatRequestV2ToolsChoiceEnum `mandatory:"false" json:"toolsChoice,omitempty"`

	// Safety mode: Adds a safety instruction for the model to use when generating responses.
	// Contextual: (Default) Puts fewer constraints on the output. It maintains core protections by aiming to reject harmful or illegal suggestions, but it allows profanity and some toxic content, sexually explicit and violent content, and content that contains medical, financial, or legal information. Contextual mode is suited for entertainment, creative, or academic use.
	// Strict: Aims to avoid sensitive topics, such as violent or sexual acts and profanity. This mode aims to provide a safer experience by prohibiting responses or recommendations that it finds inappropriate. Strict mode is suited for corporate use, such as for corporate communications and customer service.
	// Off: No safety mode is applied.
	// Note: This parameter is only compatible with models cohere.command-r-08-2024, cohere.command-r-plus-08-2024 and Cohere models released after these models. See release dates (https://docs.oracle.com/iaas/Content/generative-ai/deprecating.htm).
	SafetyMode CohereChatRequestV2SafetyModeEnum `mandatory:"false" json:"safetyMode,omitempty"`
}

func (m CohereChatRequestV2) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CohereChatRequestV2) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCohereChatRequestV2ToolsChoiceEnum(string(m.ToolsChoice)); !ok && m.ToolsChoice != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ToolsChoice: %s. Supported values are: %s.", m.ToolsChoice, strings.Join(GetCohereChatRequestV2ToolsChoiceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCohereChatRequestV2SafetyModeEnum(string(m.SafetyMode)); !ok && m.SafetyMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SafetyMode: %s. Supported values are: %s.", m.SafetyMode, strings.Join(GetCohereChatRequestV2SafetyModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CohereChatRequestV2) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCohereChatRequestV2 CohereChatRequestV2
	s := struct {
		DiscriminatorParam string `json:"apiFormat"`
		MarshalTypeCohereChatRequestV2
	}{
		"COHEREV2",
		(MarshalTypeCohereChatRequestV2)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CohereChatRequestV2) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Documents            []interface{}                      `json:"documents"`
		CitationOptions      *CitationOptionsV2                 `json:"citationOptions"`
		ToolsChoice          CohereChatRequestV2ToolsChoiceEnum `json:"toolsChoice"`
		Tools                []CohereToolV2                     `json:"tools"`
		IsStrictToolsEnabled *bool                              `json:"isStrictToolsEnabled"`
		IsLogProbsEnabled    *bool                              `json:"isLogProbsEnabled"`
		Thinking             *CohereThinkingV2                  `json:"thinking"`
		ResponseFormat       cohereresponseformat               `json:"responseFormat"`
		IsSearchQueriesOnly  *bool                              `json:"isSearchQueriesOnly"`
		StreamOptions        *StreamOptions                     `json:"streamOptions"`
		IsStream             *bool                              `json:"isStream"`
		MaxTokens            *int                               `json:"maxTokens"`
		Temperature          *float64                           `json:"temperature"`
		TopK                 *int                               `json:"topK"`
		TopP                 *float64                           `json:"topP"`
		FrequencyPenalty     *float64                           `json:"frequencyPenalty"`
		PresencePenalty      *float64                           `json:"presencePenalty"`
		Seed                 *int                               `json:"seed"`
		StopSequences        []string                           `json:"stopSequences"`
		Priority             *int                               `json:"priority"`
		IsRawPrompting       *bool                              `json:"isRawPrompting"`
		SafetyMode           CohereChatRequestV2SafetyModeEnum  `json:"safetyMode"`
		Messages             []coheremessagev2                  `json:"messages"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Documents = make([]interface{}, len(model.Documents))
	copy(m.Documents, model.Documents)
	m.CitationOptions = model.CitationOptions

	m.ToolsChoice = model.ToolsChoice

	m.Tools = make([]CohereToolV2, len(model.Tools))
	copy(m.Tools, model.Tools)
	m.IsStrictToolsEnabled = model.IsStrictToolsEnabled

	m.IsLogProbsEnabled = model.IsLogProbsEnabled

	m.Thinking = model.Thinking

	nn, e = model.ResponseFormat.UnmarshalPolymorphicJSON(model.ResponseFormat.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResponseFormat = nn.(CohereResponseFormat)
	} else {
		m.ResponseFormat = nil
	}

	m.IsSearchQueriesOnly = model.IsSearchQueriesOnly

	m.StreamOptions = model.StreamOptions

	m.IsStream = model.IsStream

	m.MaxTokens = model.MaxTokens

	m.Temperature = model.Temperature

	m.TopK = model.TopK

	m.TopP = model.TopP

	m.FrequencyPenalty = model.FrequencyPenalty

	m.PresencePenalty = model.PresencePenalty

	m.Seed = model.Seed

	m.StopSequences = make([]string, len(model.StopSequences))
	copy(m.StopSequences, model.StopSequences)
	m.Priority = model.Priority

	m.IsRawPrompting = model.IsRawPrompting

	m.SafetyMode = model.SafetyMode

	m.Messages = make([]CohereMessageV2, len(model.Messages))
	for i, n := range model.Messages {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Messages[i] = nn.(CohereMessageV2)
		} else {
			m.Messages[i] = nil
		}
	}
	return
}

// CohereChatRequestV2ToolsChoiceEnum Enum with underlying type: string
type CohereChatRequestV2ToolsChoiceEnum string

// Set of constants representing the allowable values for CohereChatRequestV2ToolsChoiceEnum
const (
	CohereChatRequestV2ToolsChoiceRequired CohereChatRequestV2ToolsChoiceEnum = "REQUIRED"
	CohereChatRequestV2ToolsChoiceNone     CohereChatRequestV2ToolsChoiceEnum = "NONE"
)

var mappingCohereChatRequestV2ToolsChoiceEnum = map[string]CohereChatRequestV2ToolsChoiceEnum{
	"REQUIRED": CohereChatRequestV2ToolsChoiceRequired,
	"NONE":     CohereChatRequestV2ToolsChoiceNone,
}

var mappingCohereChatRequestV2ToolsChoiceEnumLowerCase = map[string]CohereChatRequestV2ToolsChoiceEnum{
	"required": CohereChatRequestV2ToolsChoiceRequired,
	"none":     CohereChatRequestV2ToolsChoiceNone,
}

// GetCohereChatRequestV2ToolsChoiceEnumValues Enumerates the set of values for CohereChatRequestV2ToolsChoiceEnum
func GetCohereChatRequestV2ToolsChoiceEnumValues() []CohereChatRequestV2ToolsChoiceEnum {
	values := make([]CohereChatRequestV2ToolsChoiceEnum, 0)
	for _, v := range mappingCohereChatRequestV2ToolsChoiceEnum {
		values = append(values, v)
	}
	return values
}

// GetCohereChatRequestV2ToolsChoiceEnumStringValues Enumerates the set of values in String for CohereChatRequestV2ToolsChoiceEnum
func GetCohereChatRequestV2ToolsChoiceEnumStringValues() []string {
	return []string{
		"REQUIRED",
		"NONE",
	}
}

// GetMappingCohereChatRequestV2ToolsChoiceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCohereChatRequestV2ToolsChoiceEnum(val string) (CohereChatRequestV2ToolsChoiceEnum, bool) {
	enum, ok := mappingCohereChatRequestV2ToolsChoiceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CohereChatRequestV2SafetyModeEnum Enum with underlying type: string
type CohereChatRequestV2SafetyModeEnum string

// Set of constants representing the allowable values for CohereChatRequestV2SafetyModeEnum
const (
	CohereChatRequestV2SafetyModeContextual CohereChatRequestV2SafetyModeEnum = "CONTEXTUAL"
	CohereChatRequestV2SafetyModeStrict     CohereChatRequestV2SafetyModeEnum = "STRICT"
	CohereChatRequestV2SafetyModeOff        CohereChatRequestV2SafetyModeEnum = "OFF"
)

var mappingCohereChatRequestV2SafetyModeEnum = map[string]CohereChatRequestV2SafetyModeEnum{
	"CONTEXTUAL": CohereChatRequestV2SafetyModeContextual,
	"STRICT":     CohereChatRequestV2SafetyModeStrict,
	"OFF":        CohereChatRequestV2SafetyModeOff,
}

var mappingCohereChatRequestV2SafetyModeEnumLowerCase = map[string]CohereChatRequestV2SafetyModeEnum{
	"contextual": CohereChatRequestV2SafetyModeContextual,
	"strict":     CohereChatRequestV2SafetyModeStrict,
	"off":        CohereChatRequestV2SafetyModeOff,
}

// GetCohereChatRequestV2SafetyModeEnumValues Enumerates the set of values for CohereChatRequestV2SafetyModeEnum
func GetCohereChatRequestV2SafetyModeEnumValues() []CohereChatRequestV2SafetyModeEnum {
	values := make([]CohereChatRequestV2SafetyModeEnum, 0)
	for _, v := range mappingCohereChatRequestV2SafetyModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCohereChatRequestV2SafetyModeEnumStringValues Enumerates the set of values in String for CohereChatRequestV2SafetyModeEnum
func GetCohereChatRequestV2SafetyModeEnumStringValues() []string {
	return []string{
		"CONTEXTUAL",
		"STRICT",
		"OFF",
	}
}

// GetMappingCohereChatRequestV2SafetyModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCohereChatRequestV2SafetyModeEnum(val string) (CohereChatRequestV2SafetyModeEnum, bool) {
	enum, ok := mappingCohereChatRequestV2SafetyModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
