// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ModelDetails Possible model types
type ModelDetails interface {

	// supported language default value is en
	GetLanguageCode() *string
}

type modeldetails struct {
	JsonData     []byte
	LanguageCode *string `mandatory:"false" json:"languageCode"`
	ModelType    string  `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *modeldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermodeldetails modeldetails
	s := struct {
		Model Unmarshalermodeldetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.LanguageCode = s.Model.LanguageCode
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *modeldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "NAMED_ENTITY_RECOGNITION":
		mm := NamedEntityRecognitionModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TEXT_CLASSIFICATION":
		mm := TextClassificationModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetLanguageCode returns LanguageCode
func (m modeldetails) GetLanguageCode() *string {
	return m.LanguageCode
}

func (m modeldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m modeldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModelDetailsModelTypeEnum Enum with underlying type: string
type ModelDetailsModelTypeEnum string

// Set of constants representing the allowable values for ModelDetailsModelTypeEnum
const (
	ModelDetailsModelTypeNamedEntityRecognition ModelDetailsModelTypeEnum = "NAMED_ENTITY_RECOGNITION"
	ModelDetailsModelTypeTextClassification     ModelDetailsModelTypeEnum = "TEXT_CLASSIFICATION"
)

var mappingModelDetailsModelTypeEnum = map[string]ModelDetailsModelTypeEnum{
	"NAMED_ENTITY_RECOGNITION": ModelDetailsModelTypeNamedEntityRecognition,
	"TEXT_CLASSIFICATION":      ModelDetailsModelTypeTextClassification,
}

var mappingModelDetailsModelTypeEnumLowerCase = map[string]ModelDetailsModelTypeEnum{
	"named_entity_recognition": ModelDetailsModelTypeNamedEntityRecognition,
	"text_classification":      ModelDetailsModelTypeTextClassification,
}

// GetModelDetailsModelTypeEnumValues Enumerates the set of values for ModelDetailsModelTypeEnum
func GetModelDetailsModelTypeEnumValues() []ModelDetailsModelTypeEnum {
	values := make([]ModelDetailsModelTypeEnum, 0)
	for _, v := range mappingModelDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelDetailsModelTypeEnumStringValues Enumerates the set of values in String for ModelDetailsModelTypeEnum
func GetModelDetailsModelTypeEnumStringValues() []string {
	return []string{
		"NAMED_ENTITY_RECOGNITION",
		"TEXT_CLASSIFICATION",
	}
}

// GetMappingModelDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelDetailsModelTypeEnum(val string) (ModelDetailsModelTypeEnum, bool) {
	enum, ok := mappingModelDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
