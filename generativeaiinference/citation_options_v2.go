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

// CitationOptionsV2 Options for controlling citation generation in RAG flows.
type CitationOptionsV2 struct {

	// Dictates the approach taken to generating citations as part of the RAG flow. Defaults to "accurate".   - "ACCURATE": More precise citation generation.   - "FAST": Faster but may be less precise.   - "OFF": Disables citation generation.   Note: `command-r7b-12-2024` and `command-a-03-2025` only support "FAST" and "OFF".
	Mode CitationOptionsV2ModeEnum `mandatory:"false" json:"mode,omitempty"`
}

func (m CitationOptionsV2) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CitationOptionsV2) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCitationOptionsV2ModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetCitationOptionsV2ModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CitationOptionsV2ModeEnum Enum with underlying type: string
type CitationOptionsV2ModeEnum string

// Set of constants representing the allowable values for CitationOptionsV2ModeEnum
const (
	CitationOptionsV2ModeFast     CitationOptionsV2ModeEnum = "FAST"
	CitationOptionsV2ModeAccurate CitationOptionsV2ModeEnum = "ACCURATE"
	CitationOptionsV2ModeOff      CitationOptionsV2ModeEnum = "OFF"
)

var mappingCitationOptionsV2ModeEnum = map[string]CitationOptionsV2ModeEnum{
	"FAST":     CitationOptionsV2ModeFast,
	"ACCURATE": CitationOptionsV2ModeAccurate,
	"OFF":      CitationOptionsV2ModeOff,
}

var mappingCitationOptionsV2ModeEnumLowerCase = map[string]CitationOptionsV2ModeEnum{
	"fast":     CitationOptionsV2ModeFast,
	"accurate": CitationOptionsV2ModeAccurate,
	"off":      CitationOptionsV2ModeOff,
}

// GetCitationOptionsV2ModeEnumValues Enumerates the set of values for CitationOptionsV2ModeEnum
func GetCitationOptionsV2ModeEnumValues() []CitationOptionsV2ModeEnum {
	values := make([]CitationOptionsV2ModeEnum, 0)
	for _, v := range mappingCitationOptionsV2ModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCitationOptionsV2ModeEnumStringValues Enumerates the set of values in String for CitationOptionsV2ModeEnum
func GetCitationOptionsV2ModeEnumStringValues() []string {
	return []string{
		"FAST",
		"ACCURATE",
		"OFF",
	}
}

// GetMappingCitationOptionsV2ModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCitationOptionsV2ModeEnum(val string) (CitationOptionsV2ModeEnum, bool) {
	enum, ok := mappingCitationOptionsV2ModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
