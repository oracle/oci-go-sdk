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

// TtsOracleSpeechSettings Settings to use for generating speech with a model from ORACLE family.
type TtsOracleSpeechSettings struct {

	// The format in which the input text has been supplied i.e., Text or SSML. The supported text types are:
	// - TEXT
	// - SSML
	TextType TtsOracleSpeechSettingsTextTypeEnum `mandatory:"false" json:"textType,omitempty"`

	// The sample rate of the generated audio. By default, the audio will be generated with speaker voice sample rate.
	SampleRateInHz *int `mandatory:"false" json:"sampleRateInHz"`

	// The format of audio in which the user wants the audio to be in. The supported output formats are:
	// - MP3
	// - OGG
	// - PCM
	// - JSON
	OutputFormat TtsOracleSpeechSettingsOutputFormatEnum `mandatory:"false" json:"outputFormat,omitempty"`

	// The kind of time stamp markings the user wants for the audio.
	// This property should be provided if outputFormat is json, otherwise it will be ignored.
	// null value (i.e. no value is not specified) indicates no speech marking.
	// The supported speech mark types are:
	// - SENTENCE
	// - WORD
	SpeechMarkTypes []TtsOracleSpeechSettingsSpeechMarkTypesEnum `mandatory:"false" json:"speechMarkTypes,omitempty"`
}

func (m TtsOracleSpeechSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TtsOracleSpeechSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTtsOracleSpeechSettingsTextTypeEnum(string(m.TextType)); !ok && m.TextType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TextType: %s. Supported values are: %s.", m.TextType, strings.Join(GetTtsOracleSpeechSettingsTextTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTtsOracleSpeechSettingsOutputFormatEnum(string(m.OutputFormat)); !ok && m.OutputFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OutputFormat: %s. Supported values are: %s.", m.OutputFormat, strings.Join(GetTtsOracleSpeechSettingsOutputFormatEnumStringValues(), ",")))
	}
	for _, val := range m.SpeechMarkTypes {
		if _, ok := GetMappingTtsOracleSpeechSettingsSpeechMarkTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SpeechMarkTypes: %s. Supported values are: %s.", val, strings.Join(GetTtsOracleSpeechSettingsSpeechMarkTypesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TtsOracleSpeechSettingsTextTypeEnum Enum with underlying type: string
type TtsOracleSpeechSettingsTextTypeEnum string

// Set of constants representing the allowable values for TtsOracleSpeechSettingsTextTypeEnum
const (
	TtsOracleSpeechSettingsTextTypeText TtsOracleSpeechSettingsTextTypeEnum = "TEXT"
	TtsOracleSpeechSettingsTextTypeSsml TtsOracleSpeechSettingsTextTypeEnum = "SSML"
)

var mappingTtsOracleSpeechSettingsTextTypeEnum = map[string]TtsOracleSpeechSettingsTextTypeEnum{
	"TEXT": TtsOracleSpeechSettingsTextTypeText,
	"SSML": TtsOracleSpeechSettingsTextTypeSsml,
}

var mappingTtsOracleSpeechSettingsTextTypeEnumLowerCase = map[string]TtsOracleSpeechSettingsTextTypeEnum{
	"text": TtsOracleSpeechSettingsTextTypeText,
	"ssml": TtsOracleSpeechSettingsTextTypeSsml,
}

// GetTtsOracleSpeechSettingsTextTypeEnumValues Enumerates the set of values for TtsOracleSpeechSettingsTextTypeEnum
func GetTtsOracleSpeechSettingsTextTypeEnumValues() []TtsOracleSpeechSettingsTextTypeEnum {
	values := make([]TtsOracleSpeechSettingsTextTypeEnum, 0)
	for _, v := range mappingTtsOracleSpeechSettingsTextTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTtsOracleSpeechSettingsTextTypeEnumStringValues Enumerates the set of values in String for TtsOracleSpeechSettingsTextTypeEnum
func GetTtsOracleSpeechSettingsTextTypeEnumStringValues() []string {
	return []string{
		"TEXT",
		"SSML",
	}
}

// GetMappingTtsOracleSpeechSettingsTextTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTtsOracleSpeechSettingsTextTypeEnum(val string) (TtsOracleSpeechSettingsTextTypeEnum, bool) {
	enum, ok := mappingTtsOracleSpeechSettingsTextTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TtsOracleSpeechSettingsOutputFormatEnum Enum with underlying type: string
type TtsOracleSpeechSettingsOutputFormatEnum string

// Set of constants representing the allowable values for TtsOracleSpeechSettingsOutputFormatEnum
const (
	TtsOracleSpeechSettingsOutputFormatMp3  TtsOracleSpeechSettingsOutputFormatEnum = "MP3"
	TtsOracleSpeechSettingsOutputFormatOgg  TtsOracleSpeechSettingsOutputFormatEnum = "OGG"
	TtsOracleSpeechSettingsOutputFormatPcm  TtsOracleSpeechSettingsOutputFormatEnum = "PCM"
	TtsOracleSpeechSettingsOutputFormatJson TtsOracleSpeechSettingsOutputFormatEnum = "JSON"
)

var mappingTtsOracleSpeechSettingsOutputFormatEnum = map[string]TtsOracleSpeechSettingsOutputFormatEnum{
	"MP3":  TtsOracleSpeechSettingsOutputFormatMp3,
	"OGG":  TtsOracleSpeechSettingsOutputFormatOgg,
	"PCM":  TtsOracleSpeechSettingsOutputFormatPcm,
	"JSON": TtsOracleSpeechSettingsOutputFormatJson,
}

var mappingTtsOracleSpeechSettingsOutputFormatEnumLowerCase = map[string]TtsOracleSpeechSettingsOutputFormatEnum{
	"mp3":  TtsOracleSpeechSettingsOutputFormatMp3,
	"ogg":  TtsOracleSpeechSettingsOutputFormatOgg,
	"pcm":  TtsOracleSpeechSettingsOutputFormatPcm,
	"json": TtsOracleSpeechSettingsOutputFormatJson,
}

// GetTtsOracleSpeechSettingsOutputFormatEnumValues Enumerates the set of values for TtsOracleSpeechSettingsOutputFormatEnum
func GetTtsOracleSpeechSettingsOutputFormatEnumValues() []TtsOracleSpeechSettingsOutputFormatEnum {
	values := make([]TtsOracleSpeechSettingsOutputFormatEnum, 0)
	for _, v := range mappingTtsOracleSpeechSettingsOutputFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetTtsOracleSpeechSettingsOutputFormatEnumStringValues Enumerates the set of values in String for TtsOracleSpeechSettingsOutputFormatEnum
func GetTtsOracleSpeechSettingsOutputFormatEnumStringValues() []string {
	return []string{
		"MP3",
		"OGG",
		"PCM",
		"JSON",
	}
}

// GetMappingTtsOracleSpeechSettingsOutputFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTtsOracleSpeechSettingsOutputFormatEnum(val string) (TtsOracleSpeechSettingsOutputFormatEnum, bool) {
	enum, ok := mappingTtsOracleSpeechSettingsOutputFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TtsOracleSpeechSettingsSpeechMarkTypesEnum Enum with underlying type: string
type TtsOracleSpeechSettingsSpeechMarkTypesEnum string

// Set of constants representing the allowable values for TtsOracleSpeechSettingsSpeechMarkTypesEnum
const (
	TtsOracleSpeechSettingsSpeechMarkTypesSentence TtsOracleSpeechSettingsSpeechMarkTypesEnum = "SENTENCE"
	TtsOracleSpeechSettingsSpeechMarkTypesWord     TtsOracleSpeechSettingsSpeechMarkTypesEnum = "WORD"
)

var mappingTtsOracleSpeechSettingsSpeechMarkTypesEnum = map[string]TtsOracleSpeechSettingsSpeechMarkTypesEnum{
	"SENTENCE": TtsOracleSpeechSettingsSpeechMarkTypesSentence,
	"WORD":     TtsOracleSpeechSettingsSpeechMarkTypesWord,
}

var mappingTtsOracleSpeechSettingsSpeechMarkTypesEnumLowerCase = map[string]TtsOracleSpeechSettingsSpeechMarkTypesEnum{
	"sentence": TtsOracleSpeechSettingsSpeechMarkTypesSentence,
	"word":     TtsOracleSpeechSettingsSpeechMarkTypesWord,
}

// GetTtsOracleSpeechSettingsSpeechMarkTypesEnumValues Enumerates the set of values for TtsOracleSpeechSettingsSpeechMarkTypesEnum
func GetTtsOracleSpeechSettingsSpeechMarkTypesEnumValues() []TtsOracleSpeechSettingsSpeechMarkTypesEnum {
	values := make([]TtsOracleSpeechSettingsSpeechMarkTypesEnum, 0)
	for _, v := range mappingTtsOracleSpeechSettingsSpeechMarkTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetTtsOracleSpeechSettingsSpeechMarkTypesEnumStringValues Enumerates the set of values in String for TtsOracleSpeechSettingsSpeechMarkTypesEnum
func GetTtsOracleSpeechSettingsSpeechMarkTypesEnumStringValues() []string {
	return []string{
		"SENTENCE",
		"WORD",
	}
}

// GetMappingTtsOracleSpeechSettingsSpeechMarkTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTtsOracleSpeechSettingsSpeechMarkTypesEnum(val string) (TtsOracleSpeechSettingsSpeechMarkTypesEnum, bool) {
	enum, ok := mappingTtsOracleSpeechSettingsSpeechMarkTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
