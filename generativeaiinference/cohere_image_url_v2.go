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

// CohereImageUrlV2 Provide a base64 encoded image or an image uri if it's supported.
type CohereImageUrlV2 struct {

	// URL of an image. Can be either a base64 data URI or a web URL (type can be either image_url or text).
	Url *string `mandatory:"true" json:"url"`

	// Controls the level of detail in image processing. "auto" is the default and lets the system choose, "low" is faster but less detailed, and "high" preserves maximum detail. You can save tokens and speed up responses by using detail "low".
	Detail CohereImageUrlV2DetailEnum `mandatory:"false" json:"detail,omitempty"`
}

func (m CohereImageUrlV2) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CohereImageUrlV2) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCohereImageUrlV2DetailEnum(string(m.Detail)); !ok && m.Detail != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Detail: %s. Supported values are: %s.", m.Detail, strings.Join(GetCohereImageUrlV2DetailEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CohereImageUrlV2DetailEnum Enum with underlying type: string
type CohereImageUrlV2DetailEnum string

// Set of constants representing the allowable values for CohereImageUrlV2DetailEnum
const (
	CohereImageUrlV2DetailAuto CohereImageUrlV2DetailEnum = "AUTO"
	CohereImageUrlV2DetailHigh CohereImageUrlV2DetailEnum = "HIGH"
	CohereImageUrlV2DetailLow  CohereImageUrlV2DetailEnum = "LOW"
)

var mappingCohereImageUrlV2DetailEnum = map[string]CohereImageUrlV2DetailEnum{
	"AUTO": CohereImageUrlV2DetailAuto,
	"HIGH": CohereImageUrlV2DetailHigh,
	"LOW":  CohereImageUrlV2DetailLow,
}

var mappingCohereImageUrlV2DetailEnumLowerCase = map[string]CohereImageUrlV2DetailEnum{
	"auto": CohereImageUrlV2DetailAuto,
	"high": CohereImageUrlV2DetailHigh,
	"low":  CohereImageUrlV2DetailLow,
}

// GetCohereImageUrlV2DetailEnumValues Enumerates the set of values for CohereImageUrlV2DetailEnum
func GetCohereImageUrlV2DetailEnumValues() []CohereImageUrlV2DetailEnum {
	values := make([]CohereImageUrlV2DetailEnum, 0)
	for _, v := range mappingCohereImageUrlV2DetailEnum {
		values = append(values, v)
	}
	return values
}

// GetCohereImageUrlV2DetailEnumStringValues Enumerates the set of values in String for CohereImageUrlV2DetailEnum
func GetCohereImageUrlV2DetailEnumStringValues() []string {
	return []string{
		"AUTO",
		"HIGH",
		"LOW",
	}
}

// GetMappingCohereImageUrlV2DetailEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCohereImageUrlV2DetailEnum(val string) (CohereImageUrlV2DetailEnum, bool) {
	enum, ok := mappingCohereImageUrlV2DetailEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
