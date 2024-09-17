// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TtsConfiguration Speech configuration for TTS API.
type TtsConfiguration interface {
}

type ttsconfiguration struct {
	JsonData    []byte
	ModelFamily string `json:"modelFamily"`
}

// UnmarshalJSON unmarshals json
func (m *ttsconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerttsconfiguration ttsconfiguration
	s := struct {
		Model Unmarshalerttsconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelFamily = s.Model.ModelFamily

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *ttsconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelFamily {
	case "ORACLE":
		mm := TtsOracleConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TtsConfiguration: %s.", m.ModelFamily)
		return *m, nil
	}
}

func (m ttsconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ttsconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TtsConfigurationModelFamilyEnum Enum with underlying type: string
type TtsConfigurationModelFamilyEnum string

// Set of constants representing the allowable values for TtsConfigurationModelFamilyEnum
const (
	TtsConfigurationModelFamilyOracle TtsConfigurationModelFamilyEnum = "ORACLE"
)

var mappingTtsConfigurationModelFamilyEnum = map[string]TtsConfigurationModelFamilyEnum{
	"ORACLE": TtsConfigurationModelFamilyOracle,
}

var mappingTtsConfigurationModelFamilyEnumLowerCase = map[string]TtsConfigurationModelFamilyEnum{
	"oracle": TtsConfigurationModelFamilyOracle,
}

// GetTtsConfigurationModelFamilyEnumValues Enumerates the set of values for TtsConfigurationModelFamilyEnum
func GetTtsConfigurationModelFamilyEnumValues() []TtsConfigurationModelFamilyEnum {
	values := make([]TtsConfigurationModelFamilyEnum, 0)
	for _, v := range mappingTtsConfigurationModelFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetTtsConfigurationModelFamilyEnumStringValues Enumerates the set of values in String for TtsConfigurationModelFamilyEnum
func GetTtsConfigurationModelFamilyEnumStringValues() []string {
	return []string{
		"ORACLE",
	}
}

// GetMappingTtsConfigurationModelFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTtsConfigurationModelFamilyEnum(val string) (TtsConfigurationModelFamilyEnum, bool) {
	enum, ok := mappingTtsConfigurationModelFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
