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

// LlmInferenceRequest The base class for the inference requests.
type LlmInferenceRequest interface {
}

type llminferencerequest struct {
	JsonData    []byte
	RuntimeType string `json:"runtimeType"`
}

// UnmarshalJSON unmarshals json
func (m *llminferencerequest) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerllminferencerequest llminferencerequest
	s := struct {
		Model Unmarshalerllminferencerequest
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RuntimeType = s.Model.RuntimeType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *llminferencerequest) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RuntimeType {
	case "LLAMA":
		mm := LlamaLlmInferenceRequest{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COHERE":
		mm := CohereLlmInferenceRequest{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for LlmInferenceRequest: %s.", m.RuntimeType)
		return *m, nil
	}
}

func (m llminferencerequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m llminferencerequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LlmInferenceRequestRuntimeTypeEnum Enum with underlying type: string
type LlmInferenceRequestRuntimeTypeEnum string

// Set of constants representing the allowable values for LlmInferenceRequestRuntimeTypeEnum
const (
	LlmInferenceRequestRuntimeTypeCohere LlmInferenceRequestRuntimeTypeEnum = "COHERE"
	LlmInferenceRequestRuntimeTypeLlama  LlmInferenceRequestRuntimeTypeEnum = "LLAMA"
)

var mappingLlmInferenceRequestRuntimeTypeEnum = map[string]LlmInferenceRequestRuntimeTypeEnum{
	"COHERE": LlmInferenceRequestRuntimeTypeCohere,
	"LLAMA":  LlmInferenceRequestRuntimeTypeLlama,
}

var mappingLlmInferenceRequestRuntimeTypeEnumLowerCase = map[string]LlmInferenceRequestRuntimeTypeEnum{
	"cohere": LlmInferenceRequestRuntimeTypeCohere,
	"llama":  LlmInferenceRequestRuntimeTypeLlama,
}

// GetLlmInferenceRequestRuntimeTypeEnumValues Enumerates the set of values for LlmInferenceRequestRuntimeTypeEnum
func GetLlmInferenceRequestRuntimeTypeEnumValues() []LlmInferenceRequestRuntimeTypeEnum {
	values := make([]LlmInferenceRequestRuntimeTypeEnum, 0)
	for _, v := range mappingLlmInferenceRequestRuntimeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLlmInferenceRequestRuntimeTypeEnumStringValues Enumerates the set of values in String for LlmInferenceRequestRuntimeTypeEnum
func GetLlmInferenceRequestRuntimeTypeEnumStringValues() []string {
	return []string{
		"COHERE",
		"LLAMA",
	}
}

// GetMappingLlmInferenceRequestRuntimeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLlmInferenceRequestRuntimeTypeEnum(val string) (LlmInferenceRequestRuntimeTypeEnum, bool) {
	enum, ok := mappingLlmInferenceRequestRuntimeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
