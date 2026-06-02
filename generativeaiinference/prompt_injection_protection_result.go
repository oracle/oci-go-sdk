// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Inference API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service inference API to access your custom model endpoints, or to try the out-of-the-box models to /EN/generative-ai-inference/latest/ChatResult/Chat, /EN/generative-ai-inference/latest/GenerateTextResult/GenerateText, /EN/generative-ai-inference/latest/SummarizeTextResult/SummarizeText, and /EN/generative-ai-inference/latest/EmbedTextResult/EmbedText.
// To use a Generative AI custom model for inference, you must first create an endpoint for that model. Use the /EN/generative-ai/latest/ to /EN/generative-ai/latest/Model/ by fine-tuning an out-of-the-box model, or a previous version of a custom model, using your own data. Fine-tune the custom model on a /EN/generative-ai/latest/DedicatedAiCluster/. Then, create a /EN/generative-ai/latest/DedicatedAiCluster/ with an Endpoint to host your custom model. For resource management in the Generative AI service, use the /EN/generative-ai/latest/.
// To learn more about the service, see the Generative AI documentation (https://docs.oracle.com/iaas/Content/generative-ai/home.htm).
// **Important:** The IP addresses behind each DNS endpoint might change over time. Always use the DNS hostname listed under the following **API Endpoints** section and avoid using hard-coded fixed IP addresses.
//

package generativeaiinference

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PromptInjectionProtectionResult The result of prompt injection protection.
type PromptInjectionProtectionResult struct {

	// The score indicating the likelihood of a prompt injection attack.
	Score *float64 `mandatory:"true" json:"score"`

	// The input modalities flagged by the prompt injection result. Present only when the request
	// is processed using a non-empty `multimodalInput`.
	FlaggedModalities []PromptInjectionProtectionResultFlaggedModalitiesEnum `mandatory:"false" json:"flaggedModalities,omitempty"`
}

func (m PromptInjectionProtectionResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PromptInjectionProtectionResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.FlaggedModalities {
		if _, ok := GetMappingPromptInjectionProtectionResultFlaggedModalitiesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FlaggedModalities: %s. Supported values are: %s.", val, strings.Join(GetPromptInjectionProtectionResultFlaggedModalitiesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PromptInjectionProtectionResultFlaggedModalitiesEnum Enum with underlying type: string
type PromptInjectionProtectionResultFlaggedModalitiesEnum string

// Set of constants representing the allowable values for PromptInjectionProtectionResultFlaggedModalitiesEnum
const (
	PromptInjectionProtectionResultFlaggedModalitiesText  PromptInjectionProtectionResultFlaggedModalitiesEnum = "TEXT"
	PromptInjectionProtectionResultFlaggedModalitiesImage PromptInjectionProtectionResultFlaggedModalitiesEnum = "IMAGE"
)

var mappingPromptInjectionProtectionResultFlaggedModalitiesEnum = map[string]PromptInjectionProtectionResultFlaggedModalitiesEnum{
	"TEXT":  PromptInjectionProtectionResultFlaggedModalitiesText,
	"IMAGE": PromptInjectionProtectionResultFlaggedModalitiesImage,
}

var mappingPromptInjectionProtectionResultFlaggedModalitiesEnumLowerCase = map[string]PromptInjectionProtectionResultFlaggedModalitiesEnum{
	"text":  PromptInjectionProtectionResultFlaggedModalitiesText,
	"image": PromptInjectionProtectionResultFlaggedModalitiesImage,
}

// GetPromptInjectionProtectionResultFlaggedModalitiesEnumValues Enumerates the set of values for PromptInjectionProtectionResultFlaggedModalitiesEnum
func GetPromptInjectionProtectionResultFlaggedModalitiesEnumValues() []PromptInjectionProtectionResultFlaggedModalitiesEnum {
	values := make([]PromptInjectionProtectionResultFlaggedModalitiesEnum, 0)
	for _, v := range mappingPromptInjectionProtectionResultFlaggedModalitiesEnum {
		values = append(values, v)
	}
	return values
}

// GetPromptInjectionProtectionResultFlaggedModalitiesEnumStringValues Enumerates the set of values in String for PromptInjectionProtectionResultFlaggedModalitiesEnum
func GetPromptInjectionProtectionResultFlaggedModalitiesEnumStringValues() []string {
	return []string{
		"TEXT",
		"IMAGE",
	}
}

// GetMappingPromptInjectionProtectionResultFlaggedModalitiesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPromptInjectionProtectionResultFlaggedModalitiesEnum(val string) (PromptInjectionProtectionResultFlaggedModalitiesEnum, bool) {
	enum, ok := mappingPromptInjectionProtectionResultFlaggedModalitiesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
