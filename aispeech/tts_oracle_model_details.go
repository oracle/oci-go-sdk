// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// TtsOracleModelDetails Model specific properties to use with Oracle model for TTS.
type TtsOracleModelDetails interface {
}

type ttsoraclemodeldetails struct {
	JsonData  []byte
	ModelName string `json:"modelName"`
}

// UnmarshalJSON unmarshals json
func (m *ttsoraclemodeldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerttsoraclemodeldetails ttsoraclemodeldetails
	s := struct {
		Model Unmarshalerttsoraclemodeldetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelName = s.Model.ModelName

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *ttsoraclemodeldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelName {
	case "TTS_2_NATURAL":
		mm := TtsOracleTts2NaturalModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TTS_1_STANDARD":
		mm := TtsOracleTts1StandardModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TtsOracleModelDetails: %s.", m.ModelName)
		return *m, nil
	}
}

func (m ttsoraclemodeldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ttsoraclemodeldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TtsOracleModelDetailsModelNameEnum Enum with underlying type: string
type TtsOracleModelDetailsModelNameEnum string

// Set of constants representing the allowable values for TtsOracleModelDetailsModelNameEnum
const (
	TtsOracleModelDetailsModelName1Standard TtsOracleModelDetailsModelNameEnum = "TTS_1_STANDARD"
	TtsOracleModelDetailsModelName2Natural  TtsOracleModelDetailsModelNameEnum = "TTS_2_NATURAL"
)

var mappingTtsOracleModelDetailsModelNameEnum = map[string]TtsOracleModelDetailsModelNameEnum{
	"TTS_1_STANDARD": TtsOracleModelDetailsModelName1Standard,
	"TTS_2_NATURAL":  TtsOracleModelDetailsModelName2Natural,
}

var mappingTtsOracleModelDetailsModelNameEnumLowerCase = map[string]TtsOracleModelDetailsModelNameEnum{
	"tts_1_standard": TtsOracleModelDetailsModelName1Standard,
	"tts_2_natural":  TtsOracleModelDetailsModelName2Natural,
}

// GetTtsOracleModelDetailsModelNameEnumValues Enumerates the set of values for TtsOracleModelDetailsModelNameEnum
func GetTtsOracleModelDetailsModelNameEnumValues() []TtsOracleModelDetailsModelNameEnum {
	values := make([]TtsOracleModelDetailsModelNameEnum, 0)
	for _, v := range mappingTtsOracleModelDetailsModelNameEnum {
		values = append(values, v)
	}
	return values
}

// GetTtsOracleModelDetailsModelNameEnumStringValues Enumerates the set of values in String for TtsOracleModelDetailsModelNameEnum
func GetTtsOracleModelDetailsModelNameEnumStringValues() []string {
	return []string{
		"TTS_1_STANDARD",
		"TTS_2_NATURAL",
	}
}

// GetMappingTtsOracleModelDetailsModelNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTtsOracleModelDetailsModelNameEnum(val string) (TtsOracleModelDetailsModelNameEnum, bool) {
	enum, ok := mappingTtsOracleModelDetailsModelNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
