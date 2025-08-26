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

// RealtimeParameters Parameters to be sent to the realtime speech service over a websocket connection.
type RealtimeParameters struct {

	// Audio encoding to use
	// - audio/raw;rate=16000
	// - audio/raw;rate=8000
	// - audio/raw;rate=8000;codec=mulaw
	// - audio/raw;rate=8000;codec=alaw
	Encoding *string `mandatory:"false" json:"encoding"`

	// Toggle for ack messages.
	IsAckEnabled *bool `mandatory:"false" json:"isAckEnabled"`

	// Silence threshold for Realtime Speech partial results in milliseconds.
	// Currently supported only for Oracle model.
	PartialSilenceThresholdInMs *int `mandatory:"false" json:"partialSilenceThresholdInMs"`

	// Silence threshold for Realtime Speech final results in milliseconds.
	// Currently supported only for Oracle model.
	FinalSilenceThresholdInMs *int `mandatory:"false" json:"finalSilenceThresholdInMs"`

	// When enabled sets the amount of confidence required for latest tokens before returning them as part of a new partial result
	// Currently supported only for Oracle model.
	StabilizePartialResults RealtimeParametersStabilizePartialResultsEnum `mandatory:"false" json:"stabilizePartialResults,omitempty"`

	// Select a model to use for generating transcriptions. Currently supported models are:
	// - ORACLE
	// - WHISPER
	ModelType *string `mandatory:"false" json:"modelType"`

	// Model Domain.
	ModelDomain RealtimeParametersModelDomainEnum `mandatory:"false" json:"modelDomain,omitempty"`

	//
	// Oracle model supported language codes are locale specific.
	// Locale value as per given in [https://datatracker.ietf.org/doc/html/rfc5646]
	// - en-US: English - United States (default)
	// - es-ES: Spanish - Spain
	// - pt-BR: Portuguese - Brazil
	// - en-GB: English - Great Britain
	// - en-AU: English - Australia
	// - en-IN: English - India
	// - hi-IN: Hindi - India
	// - fr-FR: French - France
	// - de-DE: German - Germany
	// - it-IT: Italian - Italy
	// Whisper model supported language codes are locale agnostic
	// - auto: Auto-detect language
	// - af: Afrikaans
	// - am: Amharic
	// - ar: Arabic
	// - as: Assamese
	// - az: Azerbaijani
	// - ba: Bashkir
	// - be: Belarusian
	// - bg: Bulgarian
	// - bn: Bengali
	// - bo: Tibetan
	// - br: Breton
	// - bs: Bosnian
	// - ca: Catalan
	// - cs: Czech
	// - cy: Welsh
	// - da: Danish
	// - de: German
	// - el: Greek
	// - en: English (default)
	// - es: Spanish
	// - et: Estonian
	// - eu: Basque
	// - fa: Persian
	// - fi: Finnish
	// - fo: Faroese
	// - fr: French
	// - gl: Galician
	// - gu: Gujarati
	// - ha: Hausa
	// - haw: Hawaiian
	// - he: Hebrew
	// - hi: Hindi
	// - hr: Croatian
	// - ht: Haitian Creole
	// - hu: Hungarian
	// - hy: Armenian
	// - id: Indonesian
	// - is: Icelandic
	// - it: Italian
	// - ja: Japanese
	// - jv: Javanese
	// - ka: Georgian
	// - kk: Kazakh
	// - km: Khmer
	// - kn: Kannada
	// - ko: Korean
	// - la: Latin
	// - lb: Luxembourgish
	// - ln: Lingala
	// - lo: Lao
	// - lt: Lithuanian
	// - lv: Latvian
	// - mg: Malagasy
	// - mi: Maori
	// - mk: Macedonian
	// - ml: Malayalam
	// - mn: Mongolian
	// - mr: Marathi
	// - ms: Malay
	// - mt: Maltese
	// - my: Burmese
	// - ne: Nepali
	// - nl: Dutch
	// - nn: Norwegian Nynorsk
	// - no: Norwegian
	// - oc: Occitan
	// - pa: Punjabi
	// - pl: Polish
	// - ps: Pashto
	// - pt: Portuguese
	// - ro: Romanian
	// - ru: Russian
	// - sa: Sanskrit
	// - sd: Sindhi
	// - si: Sinhala
	// - sk: Slovak
	// - sl: Slovenian
	// - sn: Shona
	// - so: Somali
	// - sq: Albanian
	// - sr: Serbian
	// - su: Sundanese
	// - sv: Swedish
	// - sw: Swahili
	// - ta: Tamil
	// - te: Telugu
	// - tg: Tajik
	// - th: Thai
	// - tk: Turkmen
	// - tl: Tagalog
	// - tr: Turkish
	// - tt: Tatar
	// - uk: Ukrainian
	// - ur: Urdu
	// - uz: Uzbek
	// - vi: Vietnamese
	// - yi: Yiddish
	// - yo: Yoruba
	// - zh: Chinese
	LanguageCode *string `mandatory:"false" json:"languageCode"`

	// If set to true, the service will not fail connection attempt if it encounters any issues that prevent the loading of all specified user customizations. Any invalid customizations will simply be ignored and connection will continue being established with the default base model and any remaining valid customizations.
	// If set to false, if the service is unable to load any of the specified customizations, an error detailing why will be returned and the session will end.
	// Currently supported only for Oracle model.
	ShouldIgnoreInvalidCustomizations *bool `mandatory:"false" json:"shouldIgnoreInvalidCustomizations"`

	// Array of customization objects.
	// Currently supported only for Oracle model.
	Customizations []CustomizationInference `mandatory:"false" json:"customizations"`

	// Configure punctuations in the generated transcriptions. Disabled by default.
	// - NONE: No punctuation in the transcription response
	// - SPOKEN: Punctuations in response only when verbally spoken
	// - AUTO: Automatic punctuation in the response, spoken punctuations are disabled
	// Spoken punctuation is currently supported only for the Oracle model in the Medical domain.
	Punctuation RealtimeParametersPunctuationEnum `mandatory:"false" json:"punctuation,omitempty"`
}

func (m RealtimeParameters) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RealtimeParameters) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRealtimeParametersStabilizePartialResultsEnum(string(m.StabilizePartialResults)); !ok && m.StabilizePartialResults != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StabilizePartialResults: %s. Supported values are: %s.", m.StabilizePartialResults, strings.Join(GetRealtimeParametersStabilizePartialResultsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRealtimeParametersModelDomainEnum(string(m.ModelDomain)); !ok && m.ModelDomain != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelDomain: %s. Supported values are: %s.", m.ModelDomain, strings.Join(GetRealtimeParametersModelDomainEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRealtimeParametersPunctuationEnum(string(m.Punctuation)); !ok && m.Punctuation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Punctuation: %s. Supported values are: %s.", m.Punctuation, strings.Join(GetRealtimeParametersPunctuationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RealtimeParametersStabilizePartialResultsEnum Enum with underlying type: string
type RealtimeParametersStabilizePartialResultsEnum string

// Set of constants representing the allowable values for RealtimeParametersStabilizePartialResultsEnum
const (
	RealtimeParametersStabilizePartialResultsNone   RealtimeParametersStabilizePartialResultsEnum = "NONE"
	RealtimeParametersStabilizePartialResultsLow    RealtimeParametersStabilizePartialResultsEnum = "LOW"
	RealtimeParametersStabilizePartialResultsMedium RealtimeParametersStabilizePartialResultsEnum = "MEDIUM"
	RealtimeParametersStabilizePartialResultsHigh   RealtimeParametersStabilizePartialResultsEnum = "HIGH"
)

var mappingRealtimeParametersStabilizePartialResultsEnum = map[string]RealtimeParametersStabilizePartialResultsEnum{
	"NONE":   RealtimeParametersStabilizePartialResultsNone,
	"LOW":    RealtimeParametersStabilizePartialResultsLow,
	"MEDIUM": RealtimeParametersStabilizePartialResultsMedium,
	"HIGH":   RealtimeParametersStabilizePartialResultsHigh,
}

var mappingRealtimeParametersStabilizePartialResultsEnumLowerCase = map[string]RealtimeParametersStabilizePartialResultsEnum{
	"none":   RealtimeParametersStabilizePartialResultsNone,
	"low":    RealtimeParametersStabilizePartialResultsLow,
	"medium": RealtimeParametersStabilizePartialResultsMedium,
	"high":   RealtimeParametersStabilizePartialResultsHigh,
}

// GetRealtimeParametersStabilizePartialResultsEnumValues Enumerates the set of values for RealtimeParametersStabilizePartialResultsEnum
func GetRealtimeParametersStabilizePartialResultsEnumValues() []RealtimeParametersStabilizePartialResultsEnum {
	values := make([]RealtimeParametersStabilizePartialResultsEnum, 0)
	for _, v := range mappingRealtimeParametersStabilizePartialResultsEnum {
		values = append(values, v)
	}
	return values
}

// GetRealtimeParametersStabilizePartialResultsEnumStringValues Enumerates the set of values in String for RealtimeParametersStabilizePartialResultsEnum
func GetRealtimeParametersStabilizePartialResultsEnumStringValues() []string {
	return []string{
		"NONE",
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

// GetMappingRealtimeParametersStabilizePartialResultsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRealtimeParametersStabilizePartialResultsEnum(val string) (RealtimeParametersStabilizePartialResultsEnum, bool) {
	enum, ok := mappingRealtimeParametersStabilizePartialResultsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RealtimeParametersModelDomainEnum Enum with underlying type: string
type RealtimeParametersModelDomainEnum string

// Set of constants representing the allowable values for RealtimeParametersModelDomainEnum
const (
	RealtimeParametersModelDomainGeneric RealtimeParametersModelDomainEnum = "GENERIC"
	RealtimeParametersModelDomainMedical RealtimeParametersModelDomainEnum = "MEDICAL"
)

var mappingRealtimeParametersModelDomainEnum = map[string]RealtimeParametersModelDomainEnum{
	"GENERIC": RealtimeParametersModelDomainGeneric,
	"MEDICAL": RealtimeParametersModelDomainMedical,
}

var mappingRealtimeParametersModelDomainEnumLowerCase = map[string]RealtimeParametersModelDomainEnum{
	"generic": RealtimeParametersModelDomainGeneric,
	"medical": RealtimeParametersModelDomainMedical,
}

// GetRealtimeParametersModelDomainEnumValues Enumerates the set of values for RealtimeParametersModelDomainEnum
func GetRealtimeParametersModelDomainEnumValues() []RealtimeParametersModelDomainEnum {
	values := make([]RealtimeParametersModelDomainEnum, 0)
	for _, v := range mappingRealtimeParametersModelDomainEnum {
		values = append(values, v)
	}
	return values
}

// GetRealtimeParametersModelDomainEnumStringValues Enumerates the set of values in String for RealtimeParametersModelDomainEnum
func GetRealtimeParametersModelDomainEnumStringValues() []string {
	return []string{
		"GENERIC",
		"MEDICAL",
	}
}

// GetMappingRealtimeParametersModelDomainEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRealtimeParametersModelDomainEnum(val string) (RealtimeParametersModelDomainEnum, bool) {
	enum, ok := mappingRealtimeParametersModelDomainEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RealtimeParametersPunctuationEnum Enum with underlying type: string
type RealtimeParametersPunctuationEnum string

// Set of constants representing the allowable values for RealtimeParametersPunctuationEnum
const (
	RealtimeParametersPunctuationNone   RealtimeParametersPunctuationEnum = "NONE"
	RealtimeParametersPunctuationSpoken RealtimeParametersPunctuationEnum = "SPOKEN"
	RealtimeParametersPunctuationAuto   RealtimeParametersPunctuationEnum = "AUTO"
)

var mappingRealtimeParametersPunctuationEnum = map[string]RealtimeParametersPunctuationEnum{
	"NONE":   RealtimeParametersPunctuationNone,
	"SPOKEN": RealtimeParametersPunctuationSpoken,
	"AUTO":   RealtimeParametersPunctuationAuto,
}

var mappingRealtimeParametersPunctuationEnumLowerCase = map[string]RealtimeParametersPunctuationEnum{
	"none":   RealtimeParametersPunctuationNone,
	"spoken": RealtimeParametersPunctuationSpoken,
	"auto":   RealtimeParametersPunctuationAuto,
}

// GetRealtimeParametersPunctuationEnumValues Enumerates the set of values for RealtimeParametersPunctuationEnum
func GetRealtimeParametersPunctuationEnumValues() []RealtimeParametersPunctuationEnum {
	values := make([]RealtimeParametersPunctuationEnum, 0)
	for _, v := range mappingRealtimeParametersPunctuationEnum {
		values = append(values, v)
	}
	return values
}

// GetRealtimeParametersPunctuationEnumStringValues Enumerates the set of values in String for RealtimeParametersPunctuationEnum
func GetRealtimeParametersPunctuationEnumStringValues() []string {
	return []string{
		"NONE",
		"SPOKEN",
		"AUTO",
	}
}

// GetMappingRealtimeParametersPunctuationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRealtimeParametersPunctuationEnum(val string) (RealtimeParametersPunctuationEnum, bool) {
	enum, ok := mappingRealtimeParametersPunctuationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
