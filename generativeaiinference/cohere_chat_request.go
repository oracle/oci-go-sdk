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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CohereChatRequest Details for the chat request for Cohere models.
type CohereChatRequest struct {

	// Text input for the model to respond to.
	Message *string `mandatory:"true" json:"message"`

	// A list of previous messages between the user and the model, meant to give the model conversational context for responding to the user's message.
	ChatHistory []CohereMessage `mandatory:"false" json:"chatHistory"`

	// list of relevant documents that the model can cite to generate a more accurate reply.
	// Some suggested keys are "text", "author", and "date". For better generation quality, it is
	// recommended to keep the total word count of the strings in the dictionary to under 300
	// words.
	Documents []interface{} `mandatory:"false" json:"documents"`

	// When true, the response will only contain a list of generated search queries, but no search will take place, and no reply from the model to the user's message will be generated.
	IsSearchQueriesOnly *bool `mandatory:"false" json:"isSearchQueriesOnly"`

	// When specified, the default Cohere preamble will be replaced with the provided one. Preambles are a part of the prompt used to adjust the model's overall behavior and conversation style. Default preambles vary for different models.
	PreambleOverride *string `mandatory:"false" json:"preambleOverride"`

	// Whether to stream back partial progress. If set, tokens are sent as data-only server-sent events as they become available.
	IsStream *bool `mandatory:"false" json:"isStream"`

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
}

func (m CohereChatRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CohereChatRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

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
