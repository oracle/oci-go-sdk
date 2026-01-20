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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CohereCitationV2 A source reference or citation for a piece of content.
type CohereCitationV2 struct {

	// Start index of the cited snippet in the original source text.
	Start *int `mandatory:"false" json:"start"`

	// End index of the cited snippet in the original source text.
	End *int `mandatory:"false" json:"end"`

	// Text snippet that is being cited.
	Text *string `mandatory:"false" json:"text"`

	// Index of the content block in which this citation appears.
	ContentIndex *int `mandatory:"false" json:"contentIndex"`

	// The type of citation indicating what part of the response it is for.
	Type CohereCitationV2TypeEnum `mandatory:"false" json:"type,omitempty"`

	// List of source objects for this citation.
	Sources []CohereCitationSourceV2 `mandatory:"false" json:"sources"`
}

func (m CohereCitationV2) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CohereCitationV2) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCohereCitationV2TypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCohereCitationV2TypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CohereCitationV2TypeEnum Enum with underlying type: string
type CohereCitationV2TypeEnum string

// Set of constants representing the allowable values for CohereCitationV2TypeEnum
const (
	CohereCitationV2TypeTextContent     CohereCitationV2TypeEnum = "TEXT_CONTENT"
	CohereCitationV2TypeThinkingContent CohereCitationV2TypeEnum = "THINKING_CONTENT"
	CohereCitationV2TypePlan            CohereCitationV2TypeEnum = "PLAN"
)

var mappingCohereCitationV2TypeEnum = map[string]CohereCitationV2TypeEnum{
	"TEXT_CONTENT":     CohereCitationV2TypeTextContent,
	"THINKING_CONTENT": CohereCitationV2TypeThinkingContent,
	"PLAN":             CohereCitationV2TypePlan,
}

var mappingCohereCitationV2TypeEnumLowerCase = map[string]CohereCitationV2TypeEnum{
	"text_content":     CohereCitationV2TypeTextContent,
	"thinking_content": CohereCitationV2TypeThinkingContent,
	"plan":             CohereCitationV2TypePlan,
}

// GetCohereCitationV2TypeEnumValues Enumerates the set of values for CohereCitationV2TypeEnum
func GetCohereCitationV2TypeEnumValues() []CohereCitationV2TypeEnum {
	values := make([]CohereCitationV2TypeEnum, 0)
	for _, v := range mappingCohereCitationV2TypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCohereCitationV2TypeEnumStringValues Enumerates the set of values in String for CohereCitationV2TypeEnum
func GetCohereCitationV2TypeEnumStringValues() []string {
	return []string{
		"TEXT_CONTENT",
		"THINKING_CONTENT",
		"PLAN",
	}
}

// GetMappingCohereCitationV2TypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCohereCitationV2TypeEnum(val string) (CohereCitationV2TypeEnum, bool) {
	enum, ok := mappingCohereCitationV2TypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
