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

// CohereToolV2 A definition of tool (function).
type CohereToolV2 struct {

	// The name of the tool to be called. Valid names contain only the characters a-z, A-Z, 0-9, _ and must not begin with a digit.
	Type CohereToolV2TypeEnum `mandatory:"true" json:"type"`

	Function *Function `mandatory:"true" json:"function"`
}

func (m CohereToolV2) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CohereToolV2) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCohereToolV2TypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCohereToolV2TypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CohereToolV2TypeEnum Enum with underlying type: string
type CohereToolV2TypeEnum string

// Set of constants representing the allowable values for CohereToolV2TypeEnum
const (
	CohereToolV2TypeFunction CohereToolV2TypeEnum = "FUNCTION"
)

var mappingCohereToolV2TypeEnum = map[string]CohereToolV2TypeEnum{
	"FUNCTION": CohereToolV2TypeFunction,
}

var mappingCohereToolV2TypeEnumLowerCase = map[string]CohereToolV2TypeEnum{
	"function": CohereToolV2TypeFunction,
}

// GetCohereToolV2TypeEnumValues Enumerates the set of values for CohereToolV2TypeEnum
func GetCohereToolV2TypeEnumValues() []CohereToolV2TypeEnum {
	values := make([]CohereToolV2TypeEnum, 0)
	for _, v := range mappingCohereToolV2TypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCohereToolV2TypeEnumStringValues Enumerates the set of values in String for CohereToolV2TypeEnum
func GetCohereToolV2TypeEnumStringValues() []string {
	return []string{
		"FUNCTION",
	}
}

// GetMappingCohereToolV2TypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCohereToolV2TypeEnum(val string) (CohereToolV2TypeEnum, bool) {
	enum, ok := mappingCohereToolV2TypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
