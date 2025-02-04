// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Inference API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service inference API to access your custom model endpoints, or to try the out-of-the-box models to /EN/generative-ai-inference/latest/ChatResult/Chat, /EN/generative-ai-inference/latest/GenerateTextResult/GenerateText, /EN/generative-ai-inference/latest/SummarizeTextResult/SummarizeText, and /EN/generative-ai-inference/latest/EmbedTextResult/EmbedText.
// To use a Generative AI custom model for inference, you must first create an endpoint for that model. Use the /EN/generative-ai/latest/ to /EN/generative-ai/latest/Model/ by fine-tuning an out-of-the-box model, or a previous version of a custom model, using your own data. Fine-tune the custom model on a /EN/generative-ai/latest/DedicatedAiCluster/. Then, create a /EN/generative-ai/latest/DedicatedAiCluster/ with an Endpoint to host your custom model. For resource management in the Generative AI service, use the /EN/generative-ai/latest/.
// To learn more about the service, see the Generative AI documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeaiinference

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ToolChoice The tool choice for a tool.
type ToolChoice interface {
}

type toolchoice struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *toolchoice) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertoolchoice toolchoice
	s := struct {
		Model Unmarshalertoolchoice
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *toolchoice) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "FUNCTION":
		mm := ToolChoiceFunction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := ToolChoiceNone{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AUTO":
		mm := ToolChoiceAuto{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REQUIRED":
		mm := ToolChoiceRequired{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ToolChoice: %s.", m.Type)
		return *m, nil
	}
}

func (m toolchoice) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m toolchoice) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ToolChoiceTypeEnum Enum with underlying type: string
type ToolChoiceTypeEnum string

// Set of constants representing the allowable values for ToolChoiceTypeEnum
const (
	ToolChoiceTypeNone     ToolChoiceTypeEnum = "NONE"
	ToolChoiceTypeAuto     ToolChoiceTypeEnum = "AUTO"
	ToolChoiceTypeRequired ToolChoiceTypeEnum = "REQUIRED"
	ToolChoiceTypeFunction ToolChoiceTypeEnum = "FUNCTION"
)

var mappingToolChoiceTypeEnum = map[string]ToolChoiceTypeEnum{
	"NONE":     ToolChoiceTypeNone,
	"AUTO":     ToolChoiceTypeAuto,
	"REQUIRED": ToolChoiceTypeRequired,
	"FUNCTION": ToolChoiceTypeFunction,
}

var mappingToolChoiceTypeEnumLowerCase = map[string]ToolChoiceTypeEnum{
	"none":     ToolChoiceTypeNone,
	"auto":     ToolChoiceTypeAuto,
	"required": ToolChoiceTypeRequired,
	"function": ToolChoiceTypeFunction,
}

// GetToolChoiceTypeEnumValues Enumerates the set of values for ToolChoiceTypeEnum
func GetToolChoiceTypeEnumValues() []ToolChoiceTypeEnum {
	values := make([]ToolChoiceTypeEnum, 0)
	for _, v := range mappingToolChoiceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetToolChoiceTypeEnumStringValues Enumerates the set of values in String for ToolChoiceTypeEnum
func GetToolChoiceTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"AUTO",
		"REQUIRED",
		"FUNCTION",
	}
}

// GetMappingToolChoiceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingToolChoiceTypeEnum(val string) (ToolChoiceTypeEnum, bool) {
	enum, ok := mappingToolChoiceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
