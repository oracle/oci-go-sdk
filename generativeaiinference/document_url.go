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

// DocumentUrl The base64 encoded document data or a document uri if it's supported.
type DocumentUrl struct {

	// The base64 encoded document data or a document uri if it's supported.
	// Example for a pdf document:
	//   `{
	//     "type": "DOCUMENT",
	//     "documentUrl": {
	//       "url": "data:application/pdf;base64,<base64 encoded document content>"
	//     }
	//   }`
	// Example with a document uri:
	//   `{
	//     "type": "DOCUMENT",
	//     "videoUrl": {
	//       "url": "data:application/pdf;uri,<document uri>"
	//     }
	//   }`
	Url *string `mandatory:"true" json:"url"`

	// The default value is AUTO and only AUTO is supported. This option controls how to convert the base64 encoded document to tokens.
	Detail DocumentUrlDetailEnum `mandatory:"false" json:"detail,omitempty"`
}

func (m DocumentUrl) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DocumentUrl) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDocumentUrlDetailEnum(string(m.Detail)); !ok && m.Detail != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Detail: %s. Supported values are: %s.", m.Detail, strings.Join(GetDocumentUrlDetailEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DocumentUrlDetailEnum Enum with underlying type: string
type DocumentUrlDetailEnum string

// Set of constants representing the allowable values for DocumentUrlDetailEnum
const (
	DocumentUrlDetailAuto DocumentUrlDetailEnum = "AUTO"
	DocumentUrlDetailHigh DocumentUrlDetailEnum = "HIGH"
	DocumentUrlDetailLow  DocumentUrlDetailEnum = "LOW"
)

var mappingDocumentUrlDetailEnum = map[string]DocumentUrlDetailEnum{
	"AUTO": DocumentUrlDetailAuto,
	"HIGH": DocumentUrlDetailHigh,
	"LOW":  DocumentUrlDetailLow,
}

var mappingDocumentUrlDetailEnumLowerCase = map[string]DocumentUrlDetailEnum{
	"auto": DocumentUrlDetailAuto,
	"high": DocumentUrlDetailHigh,
	"low":  DocumentUrlDetailLow,
}

// GetDocumentUrlDetailEnumValues Enumerates the set of values for DocumentUrlDetailEnum
func GetDocumentUrlDetailEnumValues() []DocumentUrlDetailEnum {
	values := make([]DocumentUrlDetailEnum, 0)
	for _, v := range mappingDocumentUrlDetailEnum {
		values = append(values, v)
	}
	return values
}

// GetDocumentUrlDetailEnumStringValues Enumerates the set of values in String for DocumentUrlDetailEnum
func GetDocumentUrlDetailEnumStringValues() []string {
	return []string{
		"AUTO",
		"HIGH",
		"LOW",
	}
}

// GetMappingDocumentUrlDetailEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDocumentUrlDetailEnum(val string) (DocumentUrlDetailEnum, bool) {
	enum, ok := mappingDocumentUrlDetailEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
