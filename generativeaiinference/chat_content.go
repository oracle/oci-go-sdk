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

// ChatContent The base class for the chat content.
type ChatContent interface {
}

type chatcontent struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *chatcontent) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerchatcontent chatcontent
	s := struct {
		Model Unmarshalerchatcontent
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *chatcontent) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "TEXT":
		mm := TextContent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ChatContent: %s.", m.Type)
		return *m, nil
	}
}

func (m chatcontent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m chatcontent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ChatContentTypeEnum Enum with underlying type: string
type ChatContentTypeEnum string

// Set of constants representing the allowable values for ChatContentTypeEnum
const (
	ChatContentTypeText ChatContentTypeEnum = "TEXT"
)

var mappingChatContentTypeEnum = map[string]ChatContentTypeEnum{
	"TEXT": ChatContentTypeText,
}

var mappingChatContentTypeEnumLowerCase = map[string]ChatContentTypeEnum{
	"text": ChatContentTypeText,
}

// GetChatContentTypeEnumValues Enumerates the set of values for ChatContentTypeEnum
func GetChatContentTypeEnumValues() []ChatContentTypeEnum {
	values := make([]ChatContentTypeEnum, 0)
	for _, v := range mappingChatContentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetChatContentTypeEnumStringValues Enumerates the set of values in String for ChatContentTypeEnum
func GetChatContentTypeEnumStringValues() []string {
	return []string{
		"TEXT",
	}
}

// GetMappingChatContentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChatContentTypeEnum(val string) (ChatContentTypeEnum, bool) {
	enum, ok := mappingChatContentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
