// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Inference API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service inference API to access your custom model endpoints, or to try the out-of-the-box models to /EN/generative-ai-inference/latest/ChatResult/Chat, /EN/generative-ai-inference/latest/GenerateTextResult/GenerateText, /EN/generative-ai-inference/latest/SummarizeTextResult/SummarizeText, and /EN/generative-ai-inference/latest/EmbedTextResult/EmbedText.
// To use a Generative AI custom model for inference, you must first create an endpoint for that model. Use the /EN/generative-ai/latest/ to /EN/generative-ai/latest/Model/ by fine-tuning an out-of-the-box model, or a previous version of a custom model, using your own data. Fine-tune the custom model on a /EN/generative-ai/latest/DedicatedAiCluster/. Then, create a /EN/generative-ai/latest/DedicatedAiCluster/ with an Endpoint to host your custom model. For resource management in the Generative AI service, use the /EN/generative-ai/latest/.
// To learn more about the service, see the Generative AI documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeaiinference

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CohereChatRequest Details for the chat request for Cohere models.
type CohereChatRequest struct {

	// The text that the user inputs for the model to respond to.
	Message *string `mandatory:"true" json:"message"`

	// The list of previous messages between the user and the model. The chat history gives the model context for responding to the user's inputs.
	ChatHistory []CohereMessage `mandatory:"false" json:"chatHistory"`

	// A list of relevant documents that the model can refer to for generating grounded responses to the user's requests.
	// Some example keys that you can add to the dictionary are "text", "author", and "date". Keep the total word count of the strings in the dictionary to 300 words or less.
	// Example:
	// `[
	//   { "title": "Tall penguins", "snippet": "Emperor penguins are the tallest." },
	//   { "title": "Penguin habitats", "snippet": "Emperor penguins only live in Antarctica." }
	// ]`
	Documents []interface{} `mandatory:"false" json:"documents"`

	ResponseFormat CohereResponseFormat `mandatory:"false" json:"responseFormat"`

	// When set to true, the response contains only a list of generated search queries without the search results and the model will not respond to the user's message.
	IsSearchQueriesOnly *bool `mandatory:"false" json:"isSearchQueriesOnly"`

	// If specified, the default Cohere preamble is replaced with the provided preamble. A preamble is an initial guideline message that can change the model's overall chat behavior and conversation style. Default preambles vary for different models.
	// Example: `You are a travel advisor. Answer with a pirate tone.`
	PreambleOverride *string `mandatory:"false" json:"preambleOverride"`

	// Whether to stream the partial progress of the model's response. When set to true, as tokens become available, they are sent as data-only server-sent events.
	IsStream *bool `mandatory:"false" json:"isStream"`

	// The maximum number of output tokens that the model will generate for the response.
	MaxTokens *int `mandatory:"false" json:"maxTokens"`

	// The maximum number of input tokens to send to the model. If not specified, max_input_tokens is the model's context length limit minus a small buffer.
	MaxInputTokens *int `mandatory:"false" json:"maxInputTokens"`

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

	// If specified, the backend will make a best effort to sample tokens deterministically, such that repeated requests with the same seed and parameters should return the same result. However, determinism cannot be totally guaranteed.
	Seed *int `mandatory:"false" json:"seed"`

	// Returns the full prompt that was sent to the model when True.
	IsEcho *bool `mandatory:"false" json:"isEcho"`

	// A list of available tools (functions) that the model may suggest invoking before producing a text response.
	Tools []CohereTool `mandatory:"false" json:"tools"`

	// A list of results from invoking tools recommended by the model in the previous chat turn.
	ToolResults []CohereToolResult `mandatory:"false" json:"toolResults"`

	// When enabled, the model will issue (potentially multiple) tool calls in a single step, before it receives the tool responses and directly answers the user's original message.
	IsForceSingleStep *bool `mandatory:"false" json:"isForceSingleStep"`

	// Stop the model generation when it reaches a stop sequence defined in this parameter.
	StopSequences []string `mandatory:"false" json:"stopSequences"`

	// When enabled, the user’s `message` will be sent to the model without any preprocessing.
	IsRawPrompting *bool `mandatory:"false" json:"isRawPrompting"`

	// Defaults to OFF. Dictates how the prompt will be constructed. With `promptTruncation` set to AUTO_PRESERVE_ORDER, some elements from `chatHistory` and `documents` will be dropped to construct a prompt that fits within the model's context length limit. During this process the order of the documents and chat history will be preserved. With `prompt_truncation` set to OFF, no elements will be dropped.
	PromptTruncation CohereChatRequestPromptTruncationEnum `mandatory:"false" json:"promptTruncation,omitempty"`

	// When FAST is selected, citations are generated at the same time as the text output and the request will be completed sooner. May result in less accurate citations.
	CitationQuality CohereChatRequestCitationQualityEnum `mandatory:"false" json:"citationQuality,omitempty"`
}

func (m CohereChatRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CohereChatRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCohereChatRequestPromptTruncationEnum(string(m.PromptTruncation)); !ok && m.PromptTruncation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PromptTruncation: %s. Supported values are: %s.", m.PromptTruncation, strings.Join(GetCohereChatRequestPromptTruncationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCohereChatRequestCitationQualityEnum(string(m.CitationQuality)); !ok && m.CitationQuality != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CitationQuality: %s. Supported values are: %s.", m.CitationQuality, strings.Join(GetCohereChatRequestCitationQualityEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CohereChatRequest) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCohereChatRequest CohereChatRequest
	s := struct {
		DiscriminatorParam string `json:"apiFormat"`
		MarshalTypeCohereChatRequest
	}{
		"COHERE",
		(MarshalTypeCohereChatRequest)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CohereChatRequest) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ChatHistory         []coheremessage                       `json:"chatHistory"`
		Documents           []interface{}                         `json:"documents"`
		ResponseFormat      cohereresponseformat                  `json:"responseFormat"`
		IsSearchQueriesOnly *bool                                 `json:"isSearchQueriesOnly"`
		PreambleOverride    *string                               `json:"preambleOverride"`
		IsStream            *bool                                 `json:"isStream"`
		MaxTokens           *int                                  `json:"maxTokens"`
		MaxInputTokens      *int                                  `json:"maxInputTokens"`
		Temperature         *float64                              `json:"temperature"`
		TopK                *int                                  `json:"topK"`
		TopP                *float64                              `json:"topP"`
		PromptTruncation    CohereChatRequestPromptTruncationEnum `json:"promptTruncation"`
		FrequencyPenalty    *float64                              `json:"frequencyPenalty"`
		PresencePenalty     *float64                              `json:"presencePenalty"`
		Seed                *int                                  `json:"seed"`
		IsEcho              *bool                                 `json:"isEcho"`
		Tools               []CohereTool                          `json:"tools"`
		ToolResults         []CohereToolResult                    `json:"toolResults"`
		IsForceSingleStep   *bool                                 `json:"isForceSingleStep"`
		StopSequences       []string                              `json:"stopSequences"`
		IsRawPrompting      *bool                                 `json:"isRawPrompting"`
		CitationQuality     CohereChatRequestCitationQualityEnum  `json:"citationQuality"`
		Message             *string                               `json:"message"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ChatHistory = make([]CohereMessage, len(model.ChatHistory))
	for i, n := range model.ChatHistory {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ChatHistory[i] = nn.(CohereMessage)
		} else {
			m.ChatHistory[i] = nil
		}
	}
	m.Documents = make([]interface{}, len(model.Documents))
	copy(m.Documents, model.Documents)
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

	m.PreambleOverride = model.PreambleOverride

	m.IsStream = model.IsStream

	m.MaxTokens = model.MaxTokens

	m.MaxInputTokens = model.MaxInputTokens

	m.Temperature = model.Temperature

	m.TopK = model.TopK

	m.TopP = model.TopP

	m.PromptTruncation = model.PromptTruncation

	m.FrequencyPenalty = model.FrequencyPenalty

	m.PresencePenalty = model.PresencePenalty

	m.Seed = model.Seed

	m.IsEcho = model.IsEcho

	m.Tools = make([]CohereTool, len(model.Tools))
	copy(m.Tools, model.Tools)
	m.ToolResults = make([]CohereToolResult, len(model.ToolResults))
	copy(m.ToolResults, model.ToolResults)
	m.IsForceSingleStep = model.IsForceSingleStep

	m.StopSequences = make([]string, len(model.StopSequences))
	copy(m.StopSequences, model.StopSequences)
	m.IsRawPrompting = model.IsRawPrompting

	m.CitationQuality = model.CitationQuality

	m.Message = model.Message

	return
}

// CohereChatRequestPromptTruncationEnum Enum with underlying type: string
type CohereChatRequestPromptTruncationEnum string

// Set of constants representing the allowable values for CohereChatRequestPromptTruncationEnum
const (
	CohereChatRequestPromptTruncationOff               CohereChatRequestPromptTruncationEnum = "OFF"
	CohereChatRequestPromptTruncationAutoPreserveOrder CohereChatRequestPromptTruncationEnum = "AUTO_PRESERVE_ORDER"
)

var mappingCohereChatRequestPromptTruncationEnum = map[string]CohereChatRequestPromptTruncationEnum{
	"OFF":                 CohereChatRequestPromptTruncationOff,
	"AUTO_PRESERVE_ORDER": CohereChatRequestPromptTruncationAutoPreserveOrder,
}

var mappingCohereChatRequestPromptTruncationEnumLowerCase = map[string]CohereChatRequestPromptTruncationEnum{
	"off":                 CohereChatRequestPromptTruncationOff,
	"auto_preserve_order": CohereChatRequestPromptTruncationAutoPreserveOrder,
}

// GetCohereChatRequestPromptTruncationEnumValues Enumerates the set of values for CohereChatRequestPromptTruncationEnum
func GetCohereChatRequestPromptTruncationEnumValues() []CohereChatRequestPromptTruncationEnum {
	values := make([]CohereChatRequestPromptTruncationEnum, 0)
	for _, v := range mappingCohereChatRequestPromptTruncationEnum {
		values = append(values, v)
	}
	return values
}

// GetCohereChatRequestPromptTruncationEnumStringValues Enumerates the set of values in String for CohereChatRequestPromptTruncationEnum
func GetCohereChatRequestPromptTruncationEnumStringValues() []string {
	return []string{
		"OFF",
		"AUTO_PRESERVE_ORDER",
	}
}

// GetMappingCohereChatRequestPromptTruncationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCohereChatRequestPromptTruncationEnum(val string) (CohereChatRequestPromptTruncationEnum, bool) {
	enum, ok := mappingCohereChatRequestPromptTruncationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CohereChatRequestCitationQualityEnum Enum with underlying type: string
type CohereChatRequestCitationQualityEnum string

// Set of constants representing the allowable values for CohereChatRequestCitationQualityEnum
const (
	CohereChatRequestCitationQualityAccurate CohereChatRequestCitationQualityEnum = "ACCURATE"
	CohereChatRequestCitationQualityFast     CohereChatRequestCitationQualityEnum = "FAST"
)

var mappingCohereChatRequestCitationQualityEnum = map[string]CohereChatRequestCitationQualityEnum{
	"ACCURATE": CohereChatRequestCitationQualityAccurate,
	"FAST":     CohereChatRequestCitationQualityFast,
}

var mappingCohereChatRequestCitationQualityEnumLowerCase = map[string]CohereChatRequestCitationQualityEnum{
	"accurate": CohereChatRequestCitationQualityAccurate,
	"fast":     CohereChatRequestCitationQualityFast,
}

// GetCohereChatRequestCitationQualityEnumValues Enumerates the set of values for CohereChatRequestCitationQualityEnum
func GetCohereChatRequestCitationQualityEnumValues() []CohereChatRequestCitationQualityEnum {
	values := make([]CohereChatRequestCitationQualityEnum, 0)
	for _, v := range mappingCohereChatRequestCitationQualityEnum {
		values = append(values, v)
	}
	return values
}

// GetCohereChatRequestCitationQualityEnumStringValues Enumerates the set of values in String for CohereChatRequestCitationQualityEnum
func GetCohereChatRequestCitationQualityEnumStringValues() []string {
	return []string{
		"ACCURATE",
		"FAST",
	}
}

// GetMappingCohereChatRequestCitationQualityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCohereChatRequestCitationQualityEnum(val string) (CohereChatRequestCitationQualityEnum, bool) {
	enum, ok := mappingCohereChatRequestCitationQualityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
