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

// CohereThinkingV2 Configuration for reasoning features.
type CohereThinkingV2 struct {

	// Reasoning is enabled by default for models that support it, but can be turned off by setting type = disbaled.
	Type CohereThinkingV2TypeEnum `mandatory:"true" json:"type"`

	// The maximum number of tokens the model can use for thinking, which must be set to a positive integer. The model will stop thinking if it reaches the thinking token budget and will proceed with the response.
	TokenBudget *int `mandatory:"false" json:"tokenBudget"`
}

func (m CohereThinkingV2) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CohereThinkingV2) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCohereThinkingV2TypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCohereThinkingV2TypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CohereThinkingV2TypeEnum Enum with underlying type: string
type CohereThinkingV2TypeEnum string

// Set of constants representing the allowable values for CohereThinkingV2TypeEnum
const (
	CohereThinkingV2TypeEnabled  CohereThinkingV2TypeEnum = "ENABLED"
	CohereThinkingV2TypeDisabled CohereThinkingV2TypeEnum = "DISABLED"
)

var mappingCohereThinkingV2TypeEnum = map[string]CohereThinkingV2TypeEnum{
	"ENABLED":  CohereThinkingV2TypeEnabled,
	"DISABLED": CohereThinkingV2TypeDisabled,
}

var mappingCohereThinkingV2TypeEnumLowerCase = map[string]CohereThinkingV2TypeEnum{
	"enabled":  CohereThinkingV2TypeEnabled,
	"disabled": CohereThinkingV2TypeDisabled,
}

// GetCohereThinkingV2TypeEnumValues Enumerates the set of values for CohereThinkingV2TypeEnum
func GetCohereThinkingV2TypeEnumValues() []CohereThinkingV2TypeEnum {
	values := make([]CohereThinkingV2TypeEnum, 0)
	for _, v := range mappingCohereThinkingV2TypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCohereThinkingV2TypeEnumStringValues Enumerates the set of values in String for CohereThinkingV2TypeEnum
func GetCohereThinkingV2TypeEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingCohereThinkingV2TypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCohereThinkingV2TypeEnum(val string) (CohereThinkingV2TypeEnum, bool) {
	enum, ok := mappingCohereThinkingV2TypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
