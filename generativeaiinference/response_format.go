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

// ResponseFormat An object specifying the format that the model must output.
type ResponseFormat interface {
}

type responseformat struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *responseformat) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerresponseformat responseformat
	s := struct {
		Model Unmarshalerresponseformat
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *responseformat) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "TEXT":
		mm := TextResponseFormat{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JSON_OBJECT":
		mm := JsonObjectResponseFormat{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JSON_SCHEMA":
		mm := JsonSchemaResponseFormat{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ResponseFormat: %s.", m.Type)
		return *m, nil
	}
}

func (m responseformat) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m responseformat) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResponseFormatTypeEnum Enum with underlying type: string
type ResponseFormatTypeEnum string

// Set of constants representing the allowable values for ResponseFormatTypeEnum
const (
	ResponseFormatTypeText       ResponseFormatTypeEnum = "TEXT"
	ResponseFormatTypeJsonObject ResponseFormatTypeEnum = "JSON_OBJECT"
	ResponseFormatTypeJsonSchema ResponseFormatTypeEnum = "JSON_SCHEMA"
)

var mappingResponseFormatTypeEnum = map[string]ResponseFormatTypeEnum{
	"TEXT":        ResponseFormatTypeText,
	"JSON_OBJECT": ResponseFormatTypeJsonObject,
	"JSON_SCHEMA": ResponseFormatTypeJsonSchema,
}

var mappingResponseFormatTypeEnumLowerCase = map[string]ResponseFormatTypeEnum{
	"text":        ResponseFormatTypeText,
	"json_object": ResponseFormatTypeJsonObject,
	"json_schema": ResponseFormatTypeJsonSchema,
}

// GetResponseFormatTypeEnumValues Enumerates the set of values for ResponseFormatTypeEnum
func GetResponseFormatTypeEnumValues() []ResponseFormatTypeEnum {
	values := make([]ResponseFormatTypeEnum, 0)
	for _, v := range mappingResponseFormatTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResponseFormatTypeEnumStringValues Enumerates the set of values in String for ResponseFormatTypeEnum
func GetResponseFormatTypeEnumStringValues() []string {
	return []string{
		"TEXT",
		"JSON_OBJECT",
		"JSON_SCHEMA",
	}
}

// GetMappingResponseFormatTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResponseFormatTypeEnum(val string) (ResponseFormatTypeEnum, bool) {
	enum, ok := mappingResponseFormatTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
