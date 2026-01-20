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

// EmbedTextDetails Details for the request to embed texts.
type EmbedTextDetails struct {

	// Provide a list of strings or one base64 encoded image with `input_type` setting to `IMAGE`. If text embedding, each string can be words, a phrase, or a paragraph. The maximum length of each string entry in the list is 512 tokens.
	Inputs []string `mandatory:"true" json:"inputs"`

	ServingMode ServingMode `mandatory:"true" json:"servingMode"`

	// The OCID of compartment in which to call the Generative AI service to create text embeddings.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Whether or not to include the original inputs in the response. Results are index-based.
	IsEcho *bool `mandatory:"false" json:"isEcho"`

	// Specifies the types of embeddings you want to get back. Supports list of enums. Supported values :float, int8, uint8, binary, ubinary, base64. If nothing is passed default will be considered as float.
	EmbeddingTypes []EmbedTextDetailsEmbeddingTypesEnum `mandatory:"false" json:"embeddingTypes,omitempty"`

	// The number of dimensions of the output embedding. This is only available for embed-v4 and newer models. Possible values are 256, 512, 1024, and 1536.
	OutputDimensions *int `mandatory:"false" json:"outputDimensions"`

	// For an input that's longer than the maximum token length, specifies which part of the input text will be truncated.
	Truncate EmbedTextDetailsTruncateEnum `mandatory:"false" json:"truncate,omitempty"`

	// Specifies the input type.
	InputType EmbedTextDetailsInputTypeEnum `mandatory:"false" json:"inputType,omitempty"`
}

func (m EmbedTextDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmbedTextDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.EmbeddingTypes {
		if _, ok := GetMappingEmbedTextDetailsEmbeddingTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EmbeddingTypes: %s. Supported values are: %s.", val, strings.Join(GetEmbedTextDetailsEmbeddingTypesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingEmbedTextDetailsTruncateEnum(string(m.Truncate)); !ok && m.Truncate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Truncate: %s. Supported values are: %s.", m.Truncate, strings.Join(GetEmbedTextDetailsTruncateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEmbedTextDetailsInputTypeEnum(string(m.InputType)); !ok && m.InputType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InputType: %s. Supported values are: %s.", m.InputType, strings.Join(GetEmbedTextDetailsInputTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *EmbedTextDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IsEcho           *bool                                `json:"isEcho"`
		EmbeddingTypes   []EmbedTextDetailsEmbeddingTypesEnum `json:"embeddingTypes"`
		OutputDimensions *int                                 `json:"outputDimensions"`
		Truncate         EmbedTextDetailsTruncateEnum         `json:"truncate"`
		InputType        EmbedTextDetailsInputTypeEnum        `json:"inputType"`
		Inputs           []string                             `json:"inputs"`
		ServingMode      servingmode                          `json:"servingMode"`
		CompartmentId    *string                              `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.IsEcho = model.IsEcho

	m.EmbeddingTypes = make([]EmbedTextDetailsEmbeddingTypesEnum, len(model.EmbeddingTypes))
	copy(m.EmbeddingTypes, model.EmbeddingTypes)
	m.OutputDimensions = model.OutputDimensions

	m.Truncate = model.Truncate

	m.InputType = model.InputType

	m.Inputs = make([]string, len(model.Inputs))
	copy(m.Inputs, model.Inputs)
	nn, e = model.ServingMode.UnmarshalPolymorphicJSON(model.ServingMode.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ServingMode = nn.(ServingMode)
	} else {
		m.ServingMode = nil
	}

	m.CompartmentId = model.CompartmentId

	return
}

// EmbedTextDetailsEmbeddingTypesEnum Enum with underlying type: string
type EmbedTextDetailsEmbeddingTypesEnum string

// Set of constants representing the allowable values for EmbedTextDetailsEmbeddingTypesEnum
const (
	EmbedTextDetailsEmbeddingTypesFloat   EmbedTextDetailsEmbeddingTypesEnum = "float"
	EmbedTextDetailsEmbeddingTypesInt8    EmbedTextDetailsEmbeddingTypesEnum = "int8"
	EmbedTextDetailsEmbeddingTypesUint8   EmbedTextDetailsEmbeddingTypesEnum = "uint8"
	EmbedTextDetailsEmbeddingTypesBinary  EmbedTextDetailsEmbeddingTypesEnum = "binary"
	EmbedTextDetailsEmbeddingTypesUbinary EmbedTextDetailsEmbeddingTypesEnum = "ubinary"
	EmbedTextDetailsEmbeddingTypesBase64  EmbedTextDetailsEmbeddingTypesEnum = "base64"
)

var mappingEmbedTextDetailsEmbeddingTypesEnum = map[string]EmbedTextDetailsEmbeddingTypesEnum{
	"float":   EmbedTextDetailsEmbeddingTypesFloat,
	"int8":    EmbedTextDetailsEmbeddingTypesInt8,
	"uint8":   EmbedTextDetailsEmbeddingTypesUint8,
	"binary":  EmbedTextDetailsEmbeddingTypesBinary,
	"ubinary": EmbedTextDetailsEmbeddingTypesUbinary,
	"base64":  EmbedTextDetailsEmbeddingTypesBase64,
}

var mappingEmbedTextDetailsEmbeddingTypesEnumLowerCase = map[string]EmbedTextDetailsEmbeddingTypesEnum{
	"float":   EmbedTextDetailsEmbeddingTypesFloat,
	"int8":    EmbedTextDetailsEmbeddingTypesInt8,
	"uint8":   EmbedTextDetailsEmbeddingTypesUint8,
	"binary":  EmbedTextDetailsEmbeddingTypesBinary,
	"ubinary": EmbedTextDetailsEmbeddingTypesUbinary,
	"base64":  EmbedTextDetailsEmbeddingTypesBase64,
}

// GetEmbedTextDetailsEmbeddingTypesEnumValues Enumerates the set of values for EmbedTextDetailsEmbeddingTypesEnum
func GetEmbedTextDetailsEmbeddingTypesEnumValues() []EmbedTextDetailsEmbeddingTypesEnum {
	values := make([]EmbedTextDetailsEmbeddingTypesEnum, 0)
	for _, v := range mappingEmbedTextDetailsEmbeddingTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetEmbedTextDetailsEmbeddingTypesEnumStringValues Enumerates the set of values in String for EmbedTextDetailsEmbeddingTypesEnum
func GetEmbedTextDetailsEmbeddingTypesEnumStringValues() []string {
	return []string{
		"float",
		"int8",
		"uint8",
		"binary",
		"ubinary",
		"base64",
	}
}

// GetMappingEmbedTextDetailsEmbeddingTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmbedTextDetailsEmbeddingTypesEnum(val string) (EmbedTextDetailsEmbeddingTypesEnum, bool) {
	enum, ok := mappingEmbedTextDetailsEmbeddingTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// EmbedTextDetailsTruncateEnum Enum with underlying type: string
type EmbedTextDetailsTruncateEnum string

// Set of constants representing the allowable values for EmbedTextDetailsTruncateEnum
const (
	EmbedTextDetailsTruncateNone  EmbedTextDetailsTruncateEnum = "NONE"
	EmbedTextDetailsTruncateStart EmbedTextDetailsTruncateEnum = "START"
	EmbedTextDetailsTruncateEnd   EmbedTextDetailsTruncateEnum = "END"
)

var mappingEmbedTextDetailsTruncateEnum = map[string]EmbedTextDetailsTruncateEnum{
	"NONE":  EmbedTextDetailsTruncateNone,
	"START": EmbedTextDetailsTruncateStart,
	"END":   EmbedTextDetailsTruncateEnd,
}

var mappingEmbedTextDetailsTruncateEnumLowerCase = map[string]EmbedTextDetailsTruncateEnum{
	"none":  EmbedTextDetailsTruncateNone,
	"start": EmbedTextDetailsTruncateStart,
	"end":   EmbedTextDetailsTruncateEnd,
}

// GetEmbedTextDetailsTruncateEnumValues Enumerates the set of values for EmbedTextDetailsTruncateEnum
func GetEmbedTextDetailsTruncateEnumValues() []EmbedTextDetailsTruncateEnum {
	values := make([]EmbedTextDetailsTruncateEnum, 0)
	for _, v := range mappingEmbedTextDetailsTruncateEnum {
		values = append(values, v)
	}
	return values
}

// GetEmbedTextDetailsTruncateEnumStringValues Enumerates the set of values in String for EmbedTextDetailsTruncateEnum
func GetEmbedTextDetailsTruncateEnumStringValues() []string {
	return []string{
		"NONE",
		"START",
		"END",
	}
}

// GetMappingEmbedTextDetailsTruncateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmbedTextDetailsTruncateEnum(val string) (EmbedTextDetailsTruncateEnum, bool) {
	enum, ok := mappingEmbedTextDetailsTruncateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// EmbedTextDetailsInputTypeEnum Enum with underlying type: string
type EmbedTextDetailsInputTypeEnum string

// Set of constants representing the allowable values for EmbedTextDetailsInputTypeEnum
const (
	EmbedTextDetailsInputTypeSearchDocument EmbedTextDetailsInputTypeEnum = "SEARCH_DOCUMENT"
	EmbedTextDetailsInputTypeSearchQuery    EmbedTextDetailsInputTypeEnum = "SEARCH_QUERY"
	EmbedTextDetailsInputTypeClassification EmbedTextDetailsInputTypeEnum = "CLASSIFICATION"
	EmbedTextDetailsInputTypeClustering     EmbedTextDetailsInputTypeEnum = "CLUSTERING"
	EmbedTextDetailsInputTypeImage          EmbedTextDetailsInputTypeEnum = "IMAGE"
)

var mappingEmbedTextDetailsInputTypeEnum = map[string]EmbedTextDetailsInputTypeEnum{
	"SEARCH_DOCUMENT": EmbedTextDetailsInputTypeSearchDocument,
	"SEARCH_QUERY":    EmbedTextDetailsInputTypeSearchQuery,
	"CLASSIFICATION":  EmbedTextDetailsInputTypeClassification,
	"CLUSTERING":      EmbedTextDetailsInputTypeClustering,
	"IMAGE":           EmbedTextDetailsInputTypeImage,
}

var mappingEmbedTextDetailsInputTypeEnumLowerCase = map[string]EmbedTextDetailsInputTypeEnum{
	"search_document": EmbedTextDetailsInputTypeSearchDocument,
	"search_query":    EmbedTextDetailsInputTypeSearchQuery,
	"classification":  EmbedTextDetailsInputTypeClassification,
	"clustering":      EmbedTextDetailsInputTypeClustering,
	"image":           EmbedTextDetailsInputTypeImage,
}

// GetEmbedTextDetailsInputTypeEnumValues Enumerates the set of values for EmbedTextDetailsInputTypeEnum
func GetEmbedTextDetailsInputTypeEnumValues() []EmbedTextDetailsInputTypeEnum {
	values := make([]EmbedTextDetailsInputTypeEnum, 0)
	for _, v := range mappingEmbedTextDetailsInputTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEmbedTextDetailsInputTypeEnumStringValues Enumerates the set of values in String for EmbedTextDetailsInputTypeEnum
func GetEmbedTextDetailsInputTypeEnumStringValues() []string {
	return []string{
		"SEARCH_DOCUMENT",
		"SEARCH_QUERY",
		"CLASSIFICATION",
		"CLUSTERING",
		"IMAGE",
	}
}

// GetMappingEmbedTextDetailsInputTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmbedTextDetailsInputTypeEnum(val string) (EmbedTextDetailsInputTypeEnum, bool) {
	enum, ok := mappingEmbedTextDetailsInputTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
