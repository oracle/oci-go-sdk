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
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// TranscriptionFilter Transcription Filter.
type TranscriptionFilter interface {
}

type transcriptionfilter struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *transcriptionfilter) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertranscriptionfilter transcriptionfilter
	s := struct {
		Model Unmarshalertranscriptionfilter
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *transcriptionfilter) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "PROFANITY":
		mm := ProfanityTranscriptionFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m transcriptionfilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m transcriptionfilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TranscriptionFilterTypeEnum Enum with underlying type: string
type TranscriptionFilterTypeEnum string

// Set of constants representing the allowable values for TranscriptionFilterTypeEnum
const (
	TranscriptionFilterTypeProfanity TranscriptionFilterTypeEnum = "PROFANITY"
)

var mappingTranscriptionFilterTypeEnum = map[string]TranscriptionFilterTypeEnum{
	"PROFANITY": TranscriptionFilterTypeProfanity,
}

var mappingTranscriptionFilterTypeEnumLowerCase = map[string]TranscriptionFilterTypeEnum{
	"profanity": TranscriptionFilterTypeProfanity,
}

// GetTranscriptionFilterTypeEnumValues Enumerates the set of values for TranscriptionFilterTypeEnum
func GetTranscriptionFilterTypeEnumValues() []TranscriptionFilterTypeEnum {
	values := make([]TranscriptionFilterTypeEnum, 0)
	for _, v := range mappingTranscriptionFilterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTranscriptionFilterTypeEnumStringValues Enumerates the set of values in String for TranscriptionFilterTypeEnum
func GetTranscriptionFilterTypeEnumStringValues() []string {
	return []string{
		"PROFANITY",
	}
}

// GetMappingTranscriptionFilterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTranscriptionFilterTypeEnum(val string) (TranscriptionFilterTypeEnum, bool) {
	enum, ok := mappingTranscriptionFilterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
