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

// CohereMessageV2 A message that represents a single message dialog.
type CohereMessageV2 interface {

	// Contents of the chat message.
	GetContent() []CohereContentV2
}

type coheremessagev2 struct {
	JsonData []byte
	Content  json.RawMessage `mandatory:"true" json:"content"`
	Role     string          `json:"role"`
}

// UnmarshalJSON unmarshals json
func (m *coheremessagev2) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercoheremessagev2 coheremessagev2
	s := struct {
		Model Unmarshalercoheremessagev2
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Content = s.Model.Content
	m.Role = s.Model.Role

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *coheremessagev2) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Role {
	case "SYSTEM":
		mm := CohereSystemMessageV2{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ASSISTANT":
		mm := CohereAssistantMessageV2{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "USER":
		mm := CohereUserMessageV2{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TOOL":
		mm := CohereToolMessageV2{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CohereMessageV2: %s.", m.Role)
		return *m, nil
	}
}

// GetContent returns Content
func (m coheremessagev2) GetContent() json.RawMessage {
	return m.Content
}

func (m coheremessagev2) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m coheremessagev2) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CohereMessageV2RoleEnum Enum with underlying type: string
type CohereMessageV2RoleEnum string

// Set of constants representing the allowable values for CohereMessageV2RoleEnum
const (
	CohereMessageV2RoleAssistant CohereMessageV2RoleEnum = "ASSISTANT"
	CohereMessageV2RoleUser      CohereMessageV2RoleEnum = "USER"
	CohereMessageV2RoleSystem    CohereMessageV2RoleEnum = "SYSTEM"
	CohereMessageV2RoleTool      CohereMessageV2RoleEnum = "TOOL"
)

var mappingCohereMessageV2RoleEnum = map[string]CohereMessageV2RoleEnum{
	"ASSISTANT": CohereMessageV2RoleAssistant,
	"USER":      CohereMessageV2RoleUser,
	"SYSTEM":    CohereMessageV2RoleSystem,
	"TOOL":      CohereMessageV2RoleTool,
}

var mappingCohereMessageV2RoleEnumLowerCase = map[string]CohereMessageV2RoleEnum{
	"assistant": CohereMessageV2RoleAssistant,
	"user":      CohereMessageV2RoleUser,
	"system":    CohereMessageV2RoleSystem,
	"tool":      CohereMessageV2RoleTool,
}

// GetCohereMessageV2RoleEnumValues Enumerates the set of values for CohereMessageV2RoleEnum
func GetCohereMessageV2RoleEnumValues() []CohereMessageV2RoleEnum {
	values := make([]CohereMessageV2RoleEnum, 0)
	for _, v := range mappingCohereMessageV2RoleEnum {
		values = append(values, v)
	}
	return values
}

// GetCohereMessageV2RoleEnumStringValues Enumerates the set of values in String for CohereMessageV2RoleEnum
func GetCohereMessageV2RoleEnumStringValues() []string {
	return []string{
		"ASSISTANT",
		"USER",
		"SYSTEM",
		"TOOL",
	}
}

// GetMappingCohereMessageV2RoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCohereMessageV2RoleEnum(val string) (CohereMessageV2RoleEnum, bool) {
	enum, ok := mappingCohereMessageV2RoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
