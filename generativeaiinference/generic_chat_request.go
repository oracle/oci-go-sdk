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

// GenericChatRequest Details for the chat request.
type GenericChatRequest struct {

	// The series of messages in a chat request. Includes the previous messages in a conversation. Each message includes a role (`USER` or the `CHATBOT`) and content.
	Messages []Message `mandatory:"false" json:"messages"`

	// Whether to stream back partial progress. If set to true, as tokens become available, they are sent as data-only server-sent events.
	IsStream *bool `mandatory:"false" json:"isStream"`

	// The number of of generated texts that will be returned.
	NumGenerations *int `mandatory:"false" json:"numGenerations"`

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

	// The maximum number of tokens that can be generated per output sequence. The token count of your prompt plus `maxTokens` must not exceed the model's context length.
	// Not setting a value for maxTokens results in the possible use of model's full context length.
	MaxTokens *int `mandatory:"false" json:"maxTokens"`

	// Modifies the likelihood of specified tokens that appear in the completion.
	// Example: '{"6395": 2, "8134": 1, "21943": 0.5, "5923": -100}'
	LogitBias *interface{} `mandatory:"false" json:"logitBias"`
}

func (m GenericChatRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenericChatRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
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
		Messages         []message    `json:"messages"`
		IsStream         *bool        `json:"isStream"`
		NumGenerations   *int         `json:"numGenerations"`
		IsEcho           *bool        `json:"isEcho"`
		TopK             *int         `json:"topK"`
		TopP             *float64     `json:"topP"`
		Temperature      *float64     `json:"temperature"`
		FrequencyPenalty *float64     `json:"frequencyPenalty"`
		PresencePenalty  *float64     `json:"presencePenalty"`
		Stop             []string     `json:"stop"`
		LogProbs         *int         `json:"logProbs"`
		MaxTokens        *int         `json:"maxTokens"`
		LogitBias        *interface{} `json:"logitBias"`
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
	m.IsStream = model.IsStream

	m.NumGenerations = model.NumGenerations

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

	m.LogitBias = model.LogitBias

	return
}
