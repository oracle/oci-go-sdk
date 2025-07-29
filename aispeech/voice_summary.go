// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// VoiceSummary Details related to the voice available for the given language code and voice type.
type VoiceSummary struct {

	// Unique Id of the voice.
	VoiceId *string `mandatory:"true" json:"voiceId"`

	// A user-friendly display name of the language for the user.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Gender of the speaker.
	Gender VoiceSummaryGenderEnum `mandatory:"true" json:"gender"`

	// The sample rate of the speaker in Hertz.
	SampleRateInHertz *int `mandatory:"true" json:"sampleRateInHertz"`

	// The number of words the speaker can narrate per minute. It signifies the speed of the speech produced by the speaker.
	WordsPerMinute *int `mandatory:"true" json:"wordsPerMinute"`

	// A small description of the voice like its language and voice type.
	Description *string `mandatory:"false" json:"description"`

	// Models the particular speaker is aligned to.
	SupportedModels []string `mandatory:"false" json:"supportedModels"`

	// An abbreviated notation of region to which the language and accent of the speaker belongs to.
	LanguageCode *string `mandatory:"false" json:"languageCode"`

	// A description of region to which the language and accent of the speaker belongs to.
	LanguageDescription *string `mandatory:"false" json:"languageDescription"`

	// Whether this voice id is default voice used for inference.
	IsDefaultVoice *bool `mandatory:"false" json:"isDefaultVoice"`
}

func (m VoiceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VoiceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVoiceSummaryGenderEnum(string(m.Gender)); !ok && m.Gender != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Gender: %s. Supported values are: %s.", m.Gender, strings.Join(GetVoiceSummaryGenderEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VoiceSummaryGenderEnum Enum with underlying type: string
type VoiceSummaryGenderEnum string

// Set of constants representing the allowable values for VoiceSummaryGenderEnum
const (
	VoiceSummaryGenderMale   VoiceSummaryGenderEnum = "MALE"
	VoiceSummaryGenderFemale VoiceSummaryGenderEnum = "FEMALE"
)

var mappingVoiceSummaryGenderEnum = map[string]VoiceSummaryGenderEnum{
	"MALE":   VoiceSummaryGenderMale,
	"FEMALE": VoiceSummaryGenderFemale,
}

var mappingVoiceSummaryGenderEnumLowerCase = map[string]VoiceSummaryGenderEnum{
	"male":   VoiceSummaryGenderMale,
	"female": VoiceSummaryGenderFemale,
}

// GetVoiceSummaryGenderEnumValues Enumerates the set of values for VoiceSummaryGenderEnum
func GetVoiceSummaryGenderEnumValues() []VoiceSummaryGenderEnum {
	values := make([]VoiceSummaryGenderEnum, 0)
	for _, v := range mappingVoiceSummaryGenderEnum {
		values = append(values, v)
	}
	return values
}

// GetVoiceSummaryGenderEnumStringValues Enumerates the set of values in String for VoiceSummaryGenderEnum
func GetVoiceSummaryGenderEnumStringValues() []string {
	return []string{
		"MALE",
		"FEMALE",
	}
}

// GetMappingVoiceSummaryGenderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVoiceSummaryGenderEnum(val string) (VoiceSummaryGenderEnum, bool) {
	enum, ok := mappingVoiceSummaryGenderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
