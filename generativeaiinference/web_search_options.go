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

// WebSearchOptions Options for performing a web search to augment the response.
type WebSearchOptions struct {

	// Specifies the size of the web search context.
	//   - HIGH: Most comprehensive context, highest cost, slower response.
	//   - MEDIUM: Balanced context, cost, and latency.
	//   - LOW: Least context, lowest cost, fastest response, but potentially lower answer quality.
	SearchContextSize WebSearchOptionsSearchContextSizeEnum `mandatory:"false" json:"searchContextSize,omitempty"`

	UserLocation *ApproximateLocation `mandatory:"false" json:"userLocation"`
}

func (m WebSearchOptions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WebSearchOptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingWebSearchOptionsSearchContextSizeEnum(string(m.SearchContextSize)); !ok && m.SearchContextSize != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SearchContextSize: %s. Supported values are: %s.", m.SearchContextSize, strings.Join(GetWebSearchOptionsSearchContextSizeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WebSearchOptionsSearchContextSizeEnum Enum with underlying type: string
type WebSearchOptionsSearchContextSizeEnum string

// Set of constants representing the allowable values for WebSearchOptionsSearchContextSizeEnum
const (
	WebSearchOptionsSearchContextSizeHigh   WebSearchOptionsSearchContextSizeEnum = "HIGH"
	WebSearchOptionsSearchContextSizeMedium WebSearchOptionsSearchContextSizeEnum = "MEDIUM"
	WebSearchOptionsSearchContextSizeLow    WebSearchOptionsSearchContextSizeEnum = "LOW"
)

var mappingWebSearchOptionsSearchContextSizeEnum = map[string]WebSearchOptionsSearchContextSizeEnum{
	"HIGH":   WebSearchOptionsSearchContextSizeHigh,
	"MEDIUM": WebSearchOptionsSearchContextSizeMedium,
	"LOW":    WebSearchOptionsSearchContextSizeLow,
}

var mappingWebSearchOptionsSearchContextSizeEnumLowerCase = map[string]WebSearchOptionsSearchContextSizeEnum{
	"high":   WebSearchOptionsSearchContextSizeHigh,
	"medium": WebSearchOptionsSearchContextSizeMedium,
	"low":    WebSearchOptionsSearchContextSizeLow,
}

// GetWebSearchOptionsSearchContextSizeEnumValues Enumerates the set of values for WebSearchOptionsSearchContextSizeEnum
func GetWebSearchOptionsSearchContextSizeEnumValues() []WebSearchOptionsSearchContextSizeEnum {
	values := make([]WebSearchOptionsSearchContextSizeEnum, 0)
	for _, v := range mappingWebSearchOptionsSearchContextSizeEnum {
		values = append(values, v)
	}
	return values
}

// GetWebSearchOptionsSearchContextSizeEnumStringValues Enumerates the set of values in String for WebSearchOptionsSearchContextSizeEnum
func GetWebSearchOptionsSearchContextSizeEnumStringValues() []string {
	return []string{
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

// GetMappingWebSearchOptionsSearchContextSizeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWebSearchOptionsSearchContextSizeEnum(val string) (WebSearchOptionsSearchContextSizeEnum, bool) {
	enum, ok := mappingWebSearchOptionsSearchContextSizeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
