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

// TtsAudioConfig Use this schema to specify handling of audio response.
// If audioConfig is not provided, raw response is handed over for the user to handle.
type TtsAudioConfig interface {
}

type ttsaudioconfig struct {
	JsonData   []byte
	ConfigType string `json:"configType"`
}

// UnmarshalJSON unmarshals json
func (m *ttsaudioconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerttsaudioconfig ttsaudioconfig
	s := struct {
		Model Unmarshalerttsaudioconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ConfigType = s.Model.ConfigType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *ttsaudioconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigType {
	case "BASE_AUDIO_CONFIG":
		mm := TtsBaseAudioConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for TtsAudioConfig: %s.", m.ConfigType)
		return *m, nil
	}
}

func (m ttsaudioconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ttsaudioconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TtsAudioConfigConfigTypeEnum Enum with underlying type: string
type TtsAudioConfigConfigTypeEnum string

// Set of constants representing the allowable values for TtsAudioConfigConfigTypeEnum
const (
	TtsAudioConfigConfigTypeBaseAudioConfig TtsAudioConfigConfigTypeEnum = "BASE_AUDIO_CONFIG"
)

var mappingTtsAudioConfigConfigTypeEnum = map[string]TtsAudioConfigConfigTypeEnum{
	"BASE_AUDIO_CONFIG": TtsAudioConfigConfigTypeBaseAudioConfig,
}

var mappingTtsAudioConfigConfigTypeEnumLowerCase = map[string]TtsAudioConfigConfigTypeEnum{
	"base_audio_config": TtsAudioConfigConfigTypeBaseAudioConfig,
}

// GetTtsAudioConfigConfigTypeEnumValues Enumerates the set of values for TtsAudioConfigConfigTypeEnum
func GetTtsAudioConfigConfigTypeEnumValues() []TtsAudioConfigConfigTypeEnum {
	values := make([]TtsAudioConfigConfigTypeEnum, 0)
	for _, v := range mappingTtsAudioConfigConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTtsAudioConfigConfigTypeEnumStringValues Enumerates the set of values in String for TtsAudioConfigConfigTypeEnum
func GetTtsAudioConfigConfigTypeEnumStringValues() []string {
	return []string{
		"BASE_AUDIO_CONFIG",
	}
}

// GetMappingTtsAudioConfigConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTtsAudioConfigConfigTypeEnum(val string) (TtsAudioConfigConfigTypeEnum, bool) {
	enum, ok := mappingTtsAudioConfigConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
