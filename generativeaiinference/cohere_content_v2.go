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

// CohereContentV2 The base class for the chat content.
type CohereContentV2 interface {
}

type coherecontentv2 struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *coherecontentv2) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercoherecontentv2 coherecontentv2
	s := struct {
		Model Unmarshalercoherecontentv2
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *coherecontentv2) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "THINKING":
		mm := CohereThinkingContentV2{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXT":
		mm := CohereTextContentV2{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IMAGE_URL":
		mm := CohereImageContentV2{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DOCUMENT":
		mm := CohereDocumentContentV2{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CohereContentV2: %s.", m.Type)
		return *m, nil
	}
}

func (m coherecontentv2) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m coherecontentv2) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CohereContentV2TypeEnum Enum with underlying type: string
type CohereContentV2TypeEnum string

// Set of constants representing the allowable values for CohereContentV2TypeEnum
const (
	CohereContentV2TypeText     CohereContentV2TypeEnum = "TEXT"
	CohereContentV2TypeImageUrl CohereContentV2TypeEnum = "IMAGE_URL"
	CohereContentV2TypeDocument CohereContentV2TypeEnum = "DOCUMENT"
	CohereContentV2TypeThinking CohereContentV2TypeEnum = "THINKING"
)

var mappingCohereContentV2TypeEnum = map[string]CohereContentV2TypeEnum{
	"TEXT":      CohereContentV2TypeText,
	"IMAGE_URL": CohereContentV2TypeImageUrl,
	"DOCUMENT":  CohereContentV2TypeDocument,
	"THINKING":  CohereContentV2TypeThinking,
}

var mappingCohereContentV2TypeEnumLowerCase = map[string]CohereContentV2TypeEnum{
	"text":      CohereContentV2TypeText,
	"image_url": CohereContentV2TypeImageUrl,
	"document":  CohereContentV2TypeDocument,
	"thinking":  CohereContentV2TypeThinking,
}

// GetCohereContentV2TypeEnumValues Enumerates the set of values for CohereContentV2TypeEnum
func GetCohereContentV2TypeEnumValues() []CohereContentV2TypeEnum {
	values := make([]CohereContentV2TypeEnum, 0)
	for _, v := range mappingCohereContentV2TypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCohereContentV2TypeEnumStringValues Enumerates the set of values in String for CohereContentV2TypeEnum
func GetCohereContentV2TypeEnumStringValues() []string {
	return []string{
		"TEXT",
		"IMAGE_URL",
		"DOCUMENT",
		"THINKING",
	}
}

// GetMappingCohereContentV2TypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCohereContentV2TypeEnum(val string) (CohereContentV2TypeEnum, bool) {
	enum, ok := mappingCohereContentV2TypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
