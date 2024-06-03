// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Inference API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service inference API to access your custom model endpoints, or to try the out-of-the-box models to Chat, GenerateText, SummarizeText, and EmbedText.
// To use a Generative AI custom model for inference, you must first create an endpoint for that model. Use the Generative AI service management API (https://docs.cloud.oracle.com/#/en/generative-ai/latest/) to Model by fine-tuning an out-of-the-box model, or a previous version of a custom model, using your own data. Fine-tune the custom model on a  DedicatedAiCluster. Then, create a DedicatedAiCluster with an Endpoint to host your custom model. For resource management in the Generative AI service, use the Generative AI service management API (https://docs.cloud.oracle.com/#/en/generative-ai/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeaiinference

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Message A message that represents a single chat dialog.
type Message interface {

	// Contents of the chat message.
	GetContent() []ChatContent
}

type message struct {
	JsonData []byte
	Content  json.RawMessage `mandatory:"false" json:"content"`
	Role     string          `json:"role"`
}

// UnmarshalJSON unmarshals json
func (m *message) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermessage message
	s := struct {
		Model Unmarshalermessage
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
func (m *message) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Role {
	case "SYSTEM":
		mm := SystemMessage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ASSISTANT":
		mm := AssistantMessage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "USER":
		mm := UserMessage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Message: %s.", m.Role)
		return *m, nil
	}
}

// GetContent returns Content
func (m message) GetContent() json.RawMessage {
	return m.Content
}

func (m message) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m message) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MessageRoleEnum Enum with underlying type: string
type MessageRoleEnum string

// Set of constants representing the allowable values for MessageRoleEnum
const (
	MessageRoleSystem    MessageRoleEnum = "SYSTEM"
	MessageRoleUser      MessageRoleEnum = "USER"
	MessageRoleAssistant MessageRoleEnum = "ASSISTANT"
)

var mappingMessageRoleEnum = map[string]MessageRoleEnum{
	"SYSTEM":    MessageRoleSystem,
	"USER":      MessageRoleUser,
	"ASSISTANT": MessageRoleAssistant,
}

var mappingMessageRoleEnumLowerCase = map[string]MessageRoleEnum{
	"system":    MessageRoleSystem,
	"user":      MessageRoleUser,
	"assistant": MessageRoleAssistant,
}

// GetMessageRoleEnumValues Enumerates the set of values for MessageRoleEnum
func GetMessageRoleEnumValues() []MessageRoleEnum {
	values := make([]MessageRoleEnum, 0)
	for _, v := range mappingMessageRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetMessageRoleEnumStringValues Enumerates the set of values in String for MessageRoleEnum
func GetMessageRoleEnumStringValues() []string {
	return []string{
		"SYSTEM",
		"USER",
		"ASSISTANT",
	}
}

// GetMappingMessageRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMessageRoleEnum(val string) (MessageRoleEnum, bool) {
	enum, ok := mappingMessageRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
