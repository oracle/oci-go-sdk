// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Speech API
//
// The OCI Speech Service harnesses the power of spoken language by allowing developers to easily convert file-based data containing human speech into highly accurate text transcriptions.
//

package aispeech

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RealtimeMessageResultTranscriptionToken Individual transcription tokens.
type RealtimeMessageResultTranscriptionToken struct {

	// Transcription token.
	Token *string `mandatory:"true" json:"token"`

	// Start time in milliseconds for the transcription token.
	StartTimeInMs *int `mandatory:"true" json:"startTimeInMs"`

	// End time in milliseconds for the transcription token.
	EndTimeInMs *int `mandatory:"true" json:"endTimeInMs"`

	// Confidence score for the transcription token.
	Confidence *float32 `mandatory:"true" json:"confidence"`

	// Type of the transcription token.
	Type RealtimeMessageResultTranscriptionTokenTypeEnum `mandatory:"true" json:"type"`
}

func (m RealtimeMessageResultTranscriptionToken) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RealtimeMessageResultTranscriptionToken) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRealtimeMessageResultTranscriptionTokenTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetRealtimeMessageResultTranscriptionTokenTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RealtimeMessageResultTranscriptionTokenTypeEnum Enum with underlying type: string
type RealtimeMessageResultTranscriptionTokenTypeEnum string

// Set of constants representing the allowable values for RealtimeMessageResultTranscriptionTokenTypeEnum
const (
	RealtimeMessageResultTranscriptionTokenTypeWord        RealtimeMessageResultTranscriptionTokenTypeEnum = "WORD"
	RealtimeMessageResultTranscriptionTokenTypePunctuation RealtimeMessageResultTranscriptionTokenTypeEnum = "PUNCTUATION"
)

var mappingRealtimeMessageResultTranscriptionTokenTypeEnum = map[string]RealtimeMessageResultTranscriptionTokenTypeEnum{
	"WORD":        RealtimeMessageResultTranscriptionTokenTypeWord,
	"PUNCTUATION": RealtimeMessageResultTranscriptionTokenTypePunctuation,
}

var mappingRealtimeMessageResultTranscriptionTokenTypeEnumLowerCase = map[string]RealtimeMessageResultTranscriptionTokenTypeEnum{
	"word":        RealtimeMessageResultTranscriptionTokenTypeWord,
	"punctuation": RealtimeMessageResultTranscriptionTokenTypePunctuation,
}

// GetRealtimeMessageResultTranscriptionTokenTypeEnumValues Enumerates the set of values for RealtimeMessageResultTranscriptionTokenTypeEnum
func GetRealtimeMessageResultTranscriptionTokenTypeEnumValues() []RealtimeMessageResultTranscriptionTokenTypeEnum {
	values := make([]RealtimeMessageResultTranscriptionTokenTypeEnum, 0)
	for _, v := range mappingRealtimeMessageResultTranscriptionTokenTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRealtimeMessageResultTranscriptionTokenTypeEnumStringValues Enumerates the set of values in String for RealtimeMessageResultTranscriptionTokenTypeEnum
func GetRealtimeMessageResultTranscriptionTokenTypeEnumStringValues() []string {
	return []string{
		"WORD",
		"PUNCTUATION",
	}
}

// GetMappingRealtimeMessageResultTranscriptionTokenTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRealtimeMessageResultTranscriptionTokenTypeEnum(val string) (RealtimeMessageResultTranscriptionTokenTypeEnum, bool) {
	enum, ok := mappingRealtimeMessageResultTranscriptionTokenTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
