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

// CohereChatResponseV2 The response to the chat conversation.
type CohereChatResponseV2 struct {

	// Unique identifier for the generated reply
	Id *string `mandatory:"true" json:"id"`

	Message *CohereAssistantMessageV2 `mandatory:"true" json:"message"`

	// The log probabilities of the generated tokens.
	LogProbabilities []LogProbability `mandatory:"false" json:"logProbabilities"`

	// If there is an error during the streaming scenario, then the `errorMessage` parameter contains details for the error.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	Usage *Usage `mandatory:"false" json:"usage"`

	// Why the generation stopped.
	FinishReason CohereChatResponseV2FinishReasonEnum `mandatory:"true" json:"finishReason"`
}

func (m CohereChatResponseV2) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CohereChatResponseV2) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCohereChatResponseV2FinishReasonEnum(string(m.FinishReason)); !ok && m.FinishReason != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FinishReason: %s. Supported values are: %s.", m.FinishReason, strings.Join(GetCohereChatResponseV2FinishReasonEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CohereChatResponseV2) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCohereChatResponseV2 CohereChatResponseV2
	s := struct {
		DiscriminatorParam string `json:"apiFormat"`
		MarshalTypeCohereChatResponseV2
	}{
		"COHEREV2",
		(MarshalTypeCohereChatResponseV2)(m),
	}

	return json.Marshal(&s)
}

// CohereChatResponseV2FinishReasonEnum Enum with underlying type: string
type CohereChatResponseV2FinishReasonEnum string

// Set of constants representing the allowable values for CohereChatResponseV2FinishReasonEnum
const (
	CohereChatResponseV2FinishReasonComplete     CohereChatResponseV2FinishReasonEnum = "COMPLETE"
	CohereChatResponseV2FinishReasonStopSequence CohereChatResponseV2FinishReasonEnum = "STOP_SEQUENCE"
	CohereChatResponseV2FinishReasonMaxTokens    CohereChatResponseV2FinishReasonEnum = "MAX_TOKENS"
	CohereChatResponseV2FinishReasonToolCall     CohereChatResponseV2FinishReasonEnum = "TOOL_CALL"
	CohereChatResponseV2FinishReasonError        CohereChatResponseV2FinishReasonEnum = "ERROR"
)

var mappingCohereChatResponseV2FinishReasonEnum = map[string]CohereChatResponseV2FinishReasonEnum{
	"COMPLETE":      CohereChatResponseV2FinishReasonComplete,
	"STOP_SEQUENCE": CohereChatResponseV2FinishReasonStopSequence,
	"MAX_TOKENS":    CohereChatResponseV2FinishReasonMaxTokens,
	"TOOL_CALL":     CohereChatResponseV2FinishReasonToolCall,
	"ERROR":         CohereChatResponseV2FinishReasonError,
}

var mappingCohereChatResponseV2FinishReasonEnumLowerCase = map[string]CohereChatResponseV2FinishReasonEnum{
	"complete":      CohereChatResponseV2FinishReasonComplete,
	"stop_sequence": CohereChatResponseV2FinishReasonStopSequence,
	"max_tokens":    CohereChatResponseV2FinishReasonMaxTokens,
	"tool_call":     CohereChatResponseV2FinishReasonToolCall,
	"error":         CohereChatResponseV2FinishReasonError,
}

// GetCohereChatResponseV2FinishReasonEnumValues Enumerates the set of values for CohereChatResponseV2FinishReasonEnum
func GetCohereChatResponseV2FinishReasonEnumValues() []CohereChatResponseV2FinishReasonEnum {
	values := make([]CohereChatResponseV2FinishReasonEnum, 0)
	for _, v := range mappingCohereChatResponseV2FinishReasonEnum {
		values = append(values, v)
	}
	return values
}

// GetCohereChatResponseV2FinishReasonEnumStringValues Enumerates the set of values in String for CohereChatResponseV2FinishReasonEnum
func GetCohereChatResponseV2FinishReasonEnumStringValues() []string {
	return []string{
		"COMPLETE",
		"STOP_SEQUENCE",
		"MAX_TOKENS",
		"TOOL_CALL",
		"ERROR",
	}
}

// GetMappingCohereChatResponseV2FinishReasonEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCohereChatResponseV2FinishReasonEnum(val string) (CohereChatResponseV2FinishReasonEnum, bool) {
	enum, ok := mappingCohereChatResponseV2FinishReasonEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
