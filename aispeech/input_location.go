// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Speech API
//
// The OCI Speech Service harnesses the power of spoken language by allowing developers to easily convert file-based data containing human speech into highly accurate text transcriptions.
//

package aispeech

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v63/common"
	"strings"
)

// InputLocation The location of the input(s).
type InputLocation interface {
}

type inputlocation struct {
	JsonData     []byte
	LocationType string `json:"locationType"`
}

// UnmarshalJSON unmarshals json
func (m *inputlocation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinputlocation inputlocation
	s := struct {
		Model Unmarshalerinputlocation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.LocationType = s.Model.LocationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *inputlocation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.LocationType {
	case "OBJECT_LIST_FILE_INPUT_LOCATION":
		mm := ObjectListFileInputLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_LIST_INLINE_INPUT_LOCATION":
		mm := ObjectListInlineInputLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m inputlocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m inputlocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InputLocationLocationTypeEnum Enum with underlying type: string
type InputLocationLocationTypeEnum string

// Set of constants representing the allowable values for InputLocationLocationTypeEnum
const (
	InputLocationLocationTypeInlineInputLocation InputLocationLocationTypeEnum = "OBJECT_LIST_INLINE_INPUT_LOCATION"
	InputLocationLocationTypeFileInputLocation   InputLocationLocationTypeEnum = "OBJECT_LIST_FILE_INPUT_LOCATION"
)

var mappingInputLocationLocationTypeEnum = map[string]InputLocationLocationTypeEnum{
	"OBJECT_LIST_INLINE_INPUT_LOCATION": InputLocationLocationTypeInlineInputLocation,
	"OBJECT_LIST_FILE_INPUT_LOCATION":   InputLocationLocationTypeFileInputLocation,
}

var mappingInputLocationLocationTypeEnumLowerCase = map[string]InputLocationLocationTypeEnum{
	"object_list_inline_input_location": InputLocationLocationTypeInlineInputLocation,
	"object_list_file_input_location":   InputLocationLocationTypeFileInputLocation,
}

// GetInputLocationLocationTypeEnumValues Enumerates the set of values for InputLocationLocationTypeEnum
func GetInputLocationLocationTypeEnumValues() []InputLocationLocationTypeEnum {
	values := make([]InputLocationLocationTypeEnum, 0)
	for _, v := range mappingInputLocationLocationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInputLocationLocationTypeEnumStringValues Enumerates the set of values in String for InputLocationLocationTypeEnum
func GetInputLocationLocationTypeEnumStringValues() []string {
	return []string{
		"OBJECT_LIST_INLINE_INPUT_LOCATION",
		"OBJECT_LIST_FILE_INPUT_LOCATION",
	}
}

// GetMappingInputLocationLocationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInputLocationLocationTypeEnum(val string) (InputLocationLocationTypeEnum, bool) {
	enum, ok := mappingInputLocationLocationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
