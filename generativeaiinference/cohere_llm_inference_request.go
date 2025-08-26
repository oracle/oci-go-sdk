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

// CohereLlmInferenceRequest Details for the text generation request for Cohere models.
type CohereLlmInferenceRequest struct {

	// Represents the prompt to be completed. The trailing white spaces are trimmed before completion.
	Prompt *string `mandatory:"true" json:"prompt"`

	// Whether to stream back partial progress. If set, tokens are sent as data-only server-sent events as they become available.
	IsStream *bool `mandatory:"false" json:"isStream"`

	// The number of generated texts that will be returned.
	NumGenerations *int `mandatory:"false" json:"numGenerations"`

	// Whether or not to return the user prompt in the response. This option only applies to non-stream results.
	IsEcho *bool `mandatory:"false" json:"isEcho"`

	// The maximum number of tokens to predict for each response. Includes input plus output tokens.
	MaxTokens *int `mandatory:"false" json:"maxTokens"`

	// A number that sets the randomness of the generated output. A lower temperature means a less random generations.
	// Use lower numbers for tasks with a correct answer such as question answering or summarizing. High temperatures can generate hallucinations or factually incorrect information. Start with temperatures lower than 1.0 and increase the temperature for more creative outputs, as you regenerate the prompts to refine the outputs.
	Temperature *float64 `mandatory:"false" json:"temperature"`

	// An integer that sets up the model to use only the top k most likely tokens in the generated output. A higher k introduces more randomness into the output making the output text sound more natural. Default value is 0 which disables this method and considers all tokens. To set a number for the likely tokens, choose an integer between 1 and 500.
	// If also using top p, then the model considers only the top tokens whose probabilities add up to p percent and ignores the rest of the k tokens. For example, if k is 20, but the probabilities of the top 10 add up to .75, then only the top 10 tokens are chosen.
	TopK *int `mandatory:"false" json:"topK"`

	// If set to a probability 0.0 < p < 1.0, it ensures that only the most likely tokens, with total probability mass of p, are considered for generation at each step.
	// To eliminate tokens with low likelihood, assign p a minimum percentage for the next token's likelihood. For example, when p is set to 0.75, the model eliminates the bottom 25 percent for the next token. Set to 1.0 to consider all tokens and set to 0 to disable. If both k and p are enabled, p acts after k.
	TopP *float64 `mandatory:"false" json:"topP"`

	// To reduce repetitiveness of generated tokens, this number penalizes new tokens based on their frequency in the generated text so far. Greater numbers encourage the model to use new tokens, while lower numbers encourage the model to repeat the tokens. Set to 0 to disable.
	FrequencyPenalty *float64 `mandatory:"false" json:"frequencyPenalty"`

	// To reduce repetitiveness of generated tokens, this number penalizes new tokens based on whether they've appeared in the generated text so far. Greater numbers encourage the model to use new tokens, while lower numbers encourage the model to repeat the tokens.
	// Similar to frequency penalty, a penalty is applied to previously present tokens, except that this penalty is applied equally to all tokens that have already appeared, regardless of how many times they've appeared. Set to 0 to disable.
	PresencePenalty *float64 `mandatory:"false" json:"presencePenalty"`

	// The generated text is cut at the end of the earliest occurrence of this stop sequence. The generated text will include this stop sequence.
	StopSequences []string `mandatory:"false" json:"stopSequences"`

	// Specifies how and if the token likelihoods are returned with the response.
	ReturnLikelihoods CohereLlmInferenceRequestReturnLikelihoodsEnum `mandatory:"false" json:"returnLikelihoods,omitempty"`

	// For an input that's longer than the maximum token length, specifies which part of the input text will be truncated.
	Truncate CohereLlmInferenceRequestTruncateEnum `mandatory:"false" json:"truncate,omitempty"`
}

func (m CohereLlmInferenceRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CohereLlmInferenceRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCohereLlmInferenceRequestReturnLikelihoodsEnum(string(m.ReturnLikelihoods)); !ok && m.ReturnLikelihoods != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReturnLikelihoods: %s. Supported values are: %s.", m.ReturnLikelihoods, strings.Join(GetCohereLlmInferenceRequestReturnLikelihoodsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCohereLlmInferenceRequestTruncateEnum(string(m.Truncate)); !ok && m.Truncate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Truncate: %s. Supported values are: %s.", m.Truncate, strings.Join(GetCohereLlmInferenceRequestTruncateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CohereLlmInferenceRequest) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCohereLlmInferenceRequest CohereLlmInferenceRequest
	s := struct {
		DiscriminatorParam string `json:"runtimeType"`
		MarshalTypeCohereLlmInferenceRequest
	}{
		"COHERE",
		(MarshalTypeCohereLlmInferenceRequest)(m),
	}

	return json.Marshal(&s)
}

// CohereLlmInferenceRequestReturnLikelihoodsEnum Enum with underlying type: string
type CohereLlmInferenceRequestReturnLikelihoodsEnum string

// Set of constants representing the allowable values for CohereLlmInferenceRequestReturnLikelihoodsEnum
const (
	CohereLlmInferenceRequestReturnLikelihoodsNone       CohereLlmInferenceRequestReturnLikelihoodsEnum = "NONE"
	CohereLlmInferenceRequestReturnLikelihoodsAll        CohereLlmInferenceRequestReturnLikelihoodsEnum = "ALL"
	CohereLlmInferenceRequestReturnLikelihoodsGeneration CohereLlmInferenceRequestReturnLikelihoodsEnum = "GENERATION"
)

var mappingCohereLlmInferenceRequestReturnLikelihoodsEnum = map[string]CohereLlmInferenceRequestReturnLikelihoodsEnum{
	"NONE":       CohereLlmInferenceRequestReturnLikelihoodsNone,
	"ALL":        CohereLlmInferenceRequestReturnLikelihoodsAll,
	"GENERATION": CohereLlmInferenceRequestReturnLikelihoodsGeneration,
}

var mappingCohereLlmInferenceRequestReturnLikelihoodsEnumLowerCase = map[string]CohereLlmInferenceRequestReturnLikelihoodsEnum{
	"none":       CohereLlmInferenceRequestReturnLikelihoodsNone,
	"all":        CohereLlmInferenceRequestReturnLikelihoodsAll,
	"generation": CohereLlmInferenceRequestReturnLikelihoodsGeneration,
}

// GetCohereLlmInferenceRequestReturnLikelihoodsEnumValues Enumerates the set of values for CohereLlmInferenceRequestReturnLikelihoodsEnum
func GetCohereLlmInferenceRequestReturnLikelihoodsEnumValues() []CohereLlmInferenceRequestReturnLikelihoodsEnum {
	values := make([]CohereLlmInferenceRequestReturnLikelihoodsEnum, 0)
	for _, v := range mappingCohereLlmInferenceRequestReturnLikelihoodsEnum {
		values = append(values, v)
	}
	return values
}

// GetCohereLlmInferenceRequestReturnLikelihoodsEnumStringValues Enumerates the set of values in String for CohereLlmInferenceRequestReturnLikelihoodsEnum
func GetCohereLlmInferenceRequestReturnLikelihoodsEnumStringValues() []string {
	return []string{
		"NONE",
		"ALL",
		"GENERATION",
	}
}

// GetMappingCohereLlmInferenceRequestReturnLikelihoodsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCohereLlmInferenceRequestReturnLikelihoodsEnum(val string) (CohereLlmInferenceRequestReturnLikelihoodsEnum, bool) {
	enum, ok := mappingCohereLlmInferenceRequestReturnLikelihoodsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CohereLlmInferenceRequestTruncateEnum Enum with underlying type: string
type CohereLlmInferenceRequestTruncateEnum string

// Set of constants representing the allowable values for CohereLlmInferenceRequestTruncateEnum
const (
	CohereLlmInferenceRequestTruncateNone  CohereLlmInferenceRequestTruncateEnum = "NONE"
	CohereLlmInferenceRequestTruncateStart CohereLlmInferenceRequestTruncateEnum = "START"
	CohereLlmInferenceRequestTruncateEnd   CohereLlmInferenceRequestTruncateEnum = "END"
)

var mappingCohereLlmInferenceRequestTruncateEnum = map[string]CohereLlmInferenceRequestTruncateEnum{
	"NONE":  CohereLlmInferenceRequestTruncateNone,
	"START": CohereLlmInferenceRequestTruncateStart,
	"END":   CohereLlmInferenceRequestTruncateEnd,
}

var mappingCohereLlmInferenceRequestTruncateEnumLowerCase = map[string]CohereLlmInferenceRequestTruncateEnum{
	"none":  CohereLlmInferenceRequestTruncateNone,
	"start": CohereLlmInferenceRequestTruncateStart,
	"end":   CohereLlmInferenceRequestTruncateEnd,
}

// GetCohereLlmInferenceRequestTruncateEnumValues Enumerates the set of values for CohereLlmInferenceRequestTruncateEnum
func GetCohereLlmInferenceRequestTruncateEnumValues() []CohereLlmInferenceRequestTruncateEnum {
	values := make([]CohereLlmInferenceRequestTruncateEnum, 0)
	for _, v := range mappingCohereLlmInferenceRequestTruncateEnum {
		values = append(values, v)
	}
	return values
}

// GetCohereLlmInferenceRequestTruncateEnumStringValues Enumerates the set of values in String for CohereLlmInferenceRequestTruncateEnum
func GetCohereLlmInferenceRequestTruncateEnumStringValues() []string {
	return []string{
		"NONE",
		"START",
		"END",
	}
}

// GetMappingCohereLlmInferenceRequestTruncateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCohereLlmInferenceRequestTruncateEnum(val string) (CohereLlmInferenceRequestTruncateEnum, bool) {
	enum, ok := mappingCohereLlmInferenceRequestTruncateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
