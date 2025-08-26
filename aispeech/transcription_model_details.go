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

// TranscriptionModelDetails Model details.
type TranscriptionModelDetails struct {

	// Select a model to use for generating transcriptions. Currently supported models are:
	// - ORACLE
	// - WHISPER_MEDIUM
	// - WHISPER_LARGE_V2 (upon service request)
	ModelType *string `mandatory:"false" json:"modelType"`

	// Domain for input files.
	Domain TranscriptionModelDetailsDomainEnum `mandatory:"false" json:"domain,omitempty"`

	//
	// Oracle supported language codes are (Oracle models are locale specific).
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
	// Whisper supported language codes are (Whisper models are locale agnostic).
	// - auto: Auto-detect language
	// - af: Afrikaans
	// - ar: Arabic
	// - az: Azerbaijani
	// - be: Belarusian
	// - bg: Bulgarian
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
	// - fa: Persian
	// - fi: Finnish
	// - fr: French
	// - gl: Galician
	// - he: Hebrew
	// - hi: Hindi
	// - hr: Croatian
	// - hu: Hungarian
	// - hy: Armenian
	// - id: Indonesian
	// - is: Icelandic
	// - it: Italian
	// - ja: Japanese
	// - kk: Kazakh
	// - kn: Kannada
	// - ko: Korean
	// - lt: Lithuanian
	// - lv: Latvian
	// - mi: Maori
	// - mk: Macedonian
	// - mr: Marathi
	// - ms: Malay
	// - ne: Nepali
	// - nl: Dutch
	// - no: Norwegian
	// - pl: Polish
	// - pt: Portuguese
	// - ro: Romanian
	// - ru: Russian
	// - sk: Slovak
	// - sl: Slovenian
	// - sr: Serbian
	// - sv: Swedish
	// - sw: Swahili
	// - ta: Tamil
	// - th: Thai
	// - tl: Tagalog
	// - tr: Turkish
	// - uk: Ukrainian
	// - ur: Urdu
	// - vi: Vietnamese
	// - zh: Chinese
	LanguageCode TranscriptionModelDetailsLanguageCodeEnum `mandatory:"false" json:"languageCode,omitempty"`

	TranscriptionSettings *TranscriptionSettings `mandatory:"false" json:"transcriptionSettings"`
}

func (m TranscriptionModelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TranscriptionModelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTranscriptionModelDetailsDomainEnum(string(m.Domain)); !ok && m.Domain != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Domain: %s. Supported values are: %s.", m.Domain, strings.Join(GetTranscriptionModelDetailsDomainEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTranscriptionModelDetailsLanguageCodeEnum(string(m.LanguageCode)); !ok && m.LanguageCode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LanguageCode: %s. Supported values are: %s.", m.LanguageCode, strings.Join(GetTranscriptionModelDetailsLanguageCodeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TranscriptionModelDetailsDomainEnum Enum with underlying type: string
type TranscriptionModelDetailsDomainEnum string

// Set of constants representing the allowable values for TranscriptionModelDetailsDomainEnum
const (
	TranscriptionModelDetailsDomainGeneric TranscriptionModelDetailsDomainEnum = "GENERIC"
)

var mappingTranscriptionModelDetailsDomainEnum = map[string]TranscriptionModelDetailsDomainEnum{
	"GENERIC": TranscriptionModelDetailsDomainGeneric,
}

var mappingTranscriptionModelDetailsDomainEnumLowerCase = map[string]TranscriptionModelDetailsDomainEnum{
	"generic": TranscriptionModelDetailsDomainGeneric,
}

// GetTranscriptionModelDetailsDomainEnumValues Enumerates the set of values for TranscriptionModelDetailsDomainEnum
func GetTranscriptionModelDetailsDomainEnumValues() []TranscriptionModelDetailsDomainEnum {
	values := make([]TranscriptionModelDetailsDomainEnum, 0)
	for _, v := range mappingTranscriptionModelDetailsDomainEnum {
		values = append(values, v)
	}
	return values
}

// GetTranscriptionModelDetailsDomainEnumStringValues Enumerates the set of values in String for TranscriptionModelDetailsDomainEnum
func GetTranscriptionModelDetailsDomainEnumStringValues() []string {
	return []string{
		"GENERIC",
	}
}

// GetMappingTranscriptionModelDetailsDomainEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTranscriptionModelDetailsDomainEnum(val string) (TranscriptionModelDetailsDomainEnum, bool) {
	enum, ok := mappingTranscriptionModelDetailsDomainEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TranscriptionModelDetailsLanguageCodeEnum Enum with underlying type: string
type TranscriptionModelDetailsLanguageCodeEnum string

// Set of constants representing the allowable values for TranscriptionModelDetailsLanguageCodeEnum
const (
	TranscriptionModelDetailsLanguageCodeEnUs TranscriptionModelDetailsLanguageCodeEnum = "en-US"
	TranscriptionModelDetailsLanguageCodeEsEs TranscriptionModelDetailsLanguageCodeEnum = "es-ES"
	TranscriptionModelDetailsLanguageCodePtBr TranscriptionModelDetailsLanguageCodeEnum = "pt-BR"
	TranscriptionModelDetailsLanguageCodeEnGb TranscriptionModelDetailsLanguageCodeEnum = "en-GB"
	TranscriptionModelDetailsLanguageCodeEnAu TranscriptionModelDetailsLanguageCodeEnum = "en-AU"
	TranscriptionModelDetailsLanguageCodeEnIn TranscriptionModelDetailsLanguageCodeEnum = "en-IN"
	TranscriptionModelDetailsLanguageCodeHiIn TranscriptionModelDetailsLanguageCodeEnum = "hi-IN"
	TranscriptionModelDetailsLanguageCodeFrFr TranscriptionModelDetailsLanguageCodeEnum = "fr-FR"
	TranscriptionModelDetailsLanguageCodeDeDe TranscriptionModelDetailsLanguageCodeEnum = "de-DE"
	TranscriptionModelDetailsLanguageCodeItIt TranscriptionModelDetailsLanguageCodeEnum = "it-IT"
	TranscriptionModelDetailsLanguageCodeAuto TranscriptionModelDetailsLanguageCodeEnum = "auto"
	TranscriptionModelDetailsLanguageCodeAf   TranscriptionModelDetailsLanguageCodeEnum = "af"
	TranscriptionModelDetailsLanguageCodeAr   TranscriptionModelDetailsLanguageCodeEnum = "ar"
	TranscriptionModelDetailsLanguageCodeAz   TranscriptionModelDetailsLanguageCodeEnum = "az"
	TranscriptionModelDetailsLanguageCodeBe   TranscriptionModelDetailsLanguageCodeEnum = "be"
	TranscriptionModelDetailsLanguageCodeBg   TranscriptionModelDetailsLanguageCodeEnum = "bg"
	TranscriptionModelDetailsLanguageCodeBs   TranscriptionModelDetailsLanguageCodeEnum = "bs"
	TranscriptionModelDetailsLanguageCodeCa   TranscriptionModelDetailsLanguageCodeEnum = "ca"
	TranscriptionModelDetailsLanguageCodeCs   TranscriptionModelDetailsLanguageCodeEnum = "cs"
	TranscriptionModelDetailsLanguageCodeCy   TranscriptionModelDetailsLanguageCodeEnum = "cy"
	TranscriptionModelDetailsLanguageCodeDa   TranscriptionModelDetailsLanguageCodeEnum = "da"
	TranscriptionModelDetailsLanguageCodeDe   TranscriptionModelDetailsLanguageCodeEnum = "de"
	TranscriptionModelDetailsLanguageCodeEl   TranscriptionModelDetailsLanguageCodeEnum = "el"
	TranscriptionModelDetailsLanguageCodeEn   TranscriptionModelDetailsLanguageCodeEnum = "en"
	TranscriptionModelDetailsLanguageCodeEs   TranscriptionModelDetailsLanguageCodeEnum = "es"
	TranscriptionModelDetailsLanguageCodeEt   TranscriptionModelDetailsLanguageCodeEnum = "et"
	TranscriptionModelDetailsLanguageCodeFa   TranscriptionModelDetailsLanguageCodeEnum = "fa"
	TranscriptionModelDetailsLanguageCodeFi   TranscriptionModelDetailsLanguageCodeEnum = "fi"
	TranscriptionModelDetailsLanguageCodeFr   TranscriptionModelDetailsLanguageCodeEnum = "fr"
	TranscriptionModelDetailsLanguageCodeGl   TranscriptionModelDetailsLanguageCodeEnum = "gl"
	TranscriptionModelDetailsLanguageCodeHe   TranscriptionModelDetailsLanguageCodeEnum = "he"
	TranscriptionModelDetailsLanguageCodeHi   TranscriptionModelDetailsLanguageCodeEnum = "hi"
	TranscriptionModelDetailsLanguageCodeHr   TranscriptionModelDetailsLanguageCodeEnum = "hr"
	TranscriptionModelDetailsLanguageCodeHu   TranscriptionModelDetailsLanguageCodeEnum = "hu"
	TranscriptionModelDetailsLanguageCodeHy   TranscriptionModelDetailsLanguageCodeEnum = "hy"
	TranscriptionModelDetailsLanguageCodeId   TranscriptionModelDetailsLanguageCodeEnum = "id"
	TranscriptionModelDetailsLanguageCodeIs   TranscriptionModelDetailsLanguageCodeEnum = "is"
	TranscriptionModelDetailsLanguageCodeIt   TranscriptionModelDetailsLanguageCodeEnum = "it"
	TranscriptionModelDetailsLanguageCodeJa   TranscriptionModelDetailsLanguageCodeEnum = "ja"
	TranscriptionModelDetailsLanguageCodeKk   TranscriptionModelDetailsLanguageCodeEnum = "kk"
	TranscriptionModelDetailsLanguageCodeKn   TranscriptionModelDetailsLanguageCodeEnum = "kn"
	TranscriptionModelDetailsLanguageCodeKo   TranscriptionModelDetailsLanguageCodeEnum = "ko"
	TranscriptionModelDetailsLanguageCodeLt   TranscriptionModelDetailsLanguageCodeEnum = "lt"
	TranscriptionModelDetailsLanguageCodeLv   TranscriptionModelDetailsLanguageCodeEnum = "lv"
	TranscriptionModelDetailsLanguageCodeMi   TranscriptionModelDetailsLanguageCodeEnum = "mi"
	TranscriptionModelDetailsLanguageCodeMk   TranscriptionModelDetailsLanguageCodeEnum = "mk"
	TranscriptionModelDetailsLanguageCodeMr   TranscriptionModelDetailsLanguageCodeEnum = "mr"
	TranscriptionModelDetailsLanguageCodeMs   TranscriptionModelDetailsLanguageCodeEnum = "ms"
	TranscriptionModelDetailsLanguageCodeNe   TranscriptionModelDetailsLanguageCodeEnum = "ne"
	TranscriptionModelDetailsLanguageCodeNl   TranscriptionModelDetailsLanguageCodeEnum = "nl"
	TranscriptionModelDetailsLanguageCodeNo   TranscriptionModelDetailsLanguageCodeEnum = "no"
	TranscriptionModelDetailsLanguageCodePl   TranscriptionModelDetailsLanguageCodeEnum = "pl"
	TranscriptionModelDetailsLanguageCodePt   TranscriptionModelDetailsLanguageCodeEnum = "pt"
	TranscriptionModelDetailsLanguageCodeRo   TranscriptionModelDetailsLanguageCodeEnum = "ro"
	TranscriptionModelDetailsLanguageCodeRu   TranscriptionModelDetailsLanguageCodeEnum = "ru"
	TranscriptionModelDetailsLanguageCodeSk   TranscriptionModelDetailsLanguageCodeEnum = "sk"
	TranscriptionModelDetailsLanguageCodeSl   TranscriptionModelDetailsLanguageCodeEnum = "sl"
	TranscriptionModelDetailsLanguageCodeSr   TranscriptionModelDetailsLanguageCodeEnum = "sr"
	TranscriptionModelDetailsLanguageCodeSv   TranscriptionModelDetailsLanguageCodeEnum = "sv"
	TranscriptionModelDetailsLanguageCodeSw   TranscriptionModelDetailsLanguageCodeEnum = "sw"
	TranscriptionModelDetailsLanguageCodeTa   TranscriptionModelDetailsLanguageCodeEnum = "ta"
	TranscriptionModelDetailsLanguageCodeTh   TranscriptionModelDetailsLanguageCodeEnum = "th"
	TranscriptionModelDetailsLanguageCodeTl   TranscriptionModelDetailsLanguageCodeEnum = "tl"
	TranscriptionModelDetailsLanguageCodeTr   TranscriptionModelDetailsLanguageCodeEnum = "tr"
	TranscriptionModelDetailsLanguageCodeUk   TranscriptionModelDetailsLanguageCodeEnum = "uk"
	TranscriptionModelDetailsLanguageCodeUr   TranscriptionModelDetailsLanguageCodeEnum = "ur"
	TranscriptionModelDetailsLanguageCodeVi   TranscriptionModelDetailsLanguageCodeEnum = "vi"
	TranscriptionModelDetailsLanguageCodeZh   TranscriptionModelDetailsLanguageCodeEnum = "zh"
)

var mappingTranscriptionModelDetailsLanguageCodeEnum = map[string]TranscriptionModelDetailsLanguageCodeEnum{
	"en-US": TranscriptionModelDetailsLanguageCodeEnUs,
	"es-ES": TranscriptionModelDetailsLanguageCodeEsEs,
	"pt-BR": TranscriptionModelDetailsLanguageCodePtBr,
	"en-GB": TranscriptionModelDetailsLanguageCodeEnGb,
	"en-AU": TranscriptionModelDetailsLanguageCodeEnAu,
	"en-IN": TranscriptionModelDetailsLanguageCodeEnIn,
	"hi-IN": TranscriptionModelDetailsLanguageCodeHiIn,
	"fr-FR": TranscriptionModelDetailsLanguageCodeFrFr,
	"de-DE": TranscriptionModelDetailsLanguageCodeDeDe,
	"it-IT": TranscriptionModelDetailsLanguageCodeItIt,
	"auto":  TranscriptionModelDetailsLanguageCodeAuto,
	"af":    TranscriptionModelDetailsLanguageCodeAf,
	"ar":    TranscriptionModelDetailsLanguageCodeAr,
	"az":    TranscriptionModelDetailsLanguageCodeAz,
	"be":    TranscriptionModelDetailsLanguageCodeBe,
	"bg":    TranscriptionModelDetailsLanguageCodeBg,
	"bs":    TranscriptionModelDetailsLanguageCodeBs,
	"ca":    TranscriptionModelDetailsLanguageCodeCa,
	"cs":    TranscriptionModelDetailsLanguageCodeCs,
	"cy":    TranscriptionModelDetailsLanguageCodeCy,
	"da":    TranscriptionModelDetailsLanguageCodeDa,
	"de":    TranscriptionModelDetailsLanguageCodeDe,
	"el":    TranscriptionModelDetailsLanguageCodeEl,
	"en":    TranscriptionModelDetailsLanguageCodeEn,
	"es":    TranscriptionModelDetailsLanguageCodeEs,
	"et":    TranscriptionModelDetailsLanguageCodeEt,
	"fa":    TranscriptionModelDetailsLanguageCodeFa,
	"fi":    TranscriptionModelDetailsLanguageCodeFi,
	"fr":    TranscriptionModelDetailsLanguageCodeFr,
	"gl":    TranscriptionModelDetailsLanguageCodeGl,
	"he":    TranscriptionModelDetailsLanguageCodeHe,
	"hi":    TranscriptionModelDetailsLanguageCodeHi,
	"hr":    TranscriptionModelDetailsLanguageCodeHr,
	"hu":    TranscriptionModelDetailsLanguageCodeHu,
	"hy":    TranscriptionModelDetailsLanguageCodeHy,
	"id":    TranscriptionModelDetailsLanguageCodeId,
	"is":    TranscriptionModelDetailsLanguageCodeIs,
	"it":    TranscriptionModelDetailsLanguageCodeIt,
	"ja":    TranscriptionModelDetailsLanguageCodeJa,
	"kk":    TranscriptionModelDetailsLanguageCodeKk,
	"kn":    TranscriptionModelDetailsLanguageCodeKn,
	"ko":    TranscriptionModelDetailsLanguageCodeKo,
	"lt":    TranscriptionModelDetailsLanguageCodeLt,
	"lv":    TranscriptionModelDetailsLanguageCodeLv,
	"mi":    TranscriptionModelDetailsLanguageCodeMi,
	"mk":    TranscriptionModelDetailsLanguageCodeMk,
	"mr":    TranscriptionModelDetailsLanguageCodeMr,
	"ms":    TranscriptionModelDetailsLanguageCodeMs,
	"ne":    TranscriptionModelDetailsLanguageCodeNe,
	"nl":    TranscriptionModelDetailsLanguageCodeNl,
	"no":    TranscriptionModelDetailsLanguageCodeNo,
	"pl":    TranscriptionModelDetailsLanguageCodePl,
	"pt":    TranscriptionModelDetailsLanguageCodePt,
	"ro":    TranscriptionModelDetailsLanguageCodeRo,
	"ru":    TranscriptionModelDetailsLanguageCodeRu,
	"sk":    TranscriptionModelDetailsLanguageCodeSk,
	"sl":    TranscriptionModelDetailsLanguageCodeSl,
	"sr":    TranscriptionModelDetailsLanguageCodeSr,
	"sv":    TranscriptionModelDetailsLanguageCodeSv,
	"sw":    TranscriptionModelDetailsLanguageCodeSw,
	"ta":    TranscriptionModelDetailsLanguageCodeTa,
	"th":    TranscriptionModelDetailsLanguageCodeTh,
	"tl":    TranscriptionModelDetailsLanguageCodeTl,
	"tr":    TranscriptionModelDetailsLanguageCodeTr,
	"uk":    TranscriptionModelDetailsLanguageCodeUk,
	"ur":    TranscriptionModelDetailsLanguageCodeUr,
	"vi":    TranscriptionModelDetailsLanguageCodeVi,
	"zh":    TranscriptionModelDetailsLanguageCodeZh,
}

var mappingTranscriptionModelDetailsLanguageCodeEnumLowerCase = map[string]TranscriptionModelDetailsLanguageCodeEnum{
	"en-us": TranscriptionModelDetailsLanguageCodeEnUs,
	"es-es": TranscriptionModelDetailsLanguageCodeEsEs,
	"pt-br": TranscriptionModelDetailsLanguageCodePtBr,
	"en-gb": TranscriptionModelDetailsLanguageCodeEnGb,
	"en-au": TranscriptionModelDetailsLanguageCodeEnAu,
	"en-in": TranscriptionModelDetailsLanguageCodeEnIn,
	"hi-in": TranscriptionModelDetailsLanguageCodeHiIn,
	"fr-fr": TranscriptionModelDetailsLanguageCodeFrFr,
	"de-de": TranscriptionModelDetailsLanguageCodeDeDe,
	"it-it": TranscriptionModelDetailsLanguageCodeItIt,
	"auto":  TranscriptionModelDetailsLanguageCodeAuto,
	"af":    TranscriptionModelDetailsLanguageCodeAf,
	"ar":    TranscriptionModelDetailsLanguageCodeAr,
	"az":    TranscriptionModelDetailsLanguageCodeAz,
	"be":    TranscriptionModelDetailsLanguageCodeBe,
	"bg":    TranscriptionModelDetailsLanguageCodeBg,
	"bs":    TranscriptionModelDetailsLanguageCodeBs,
	"ca":    TranscriptionModelDetailsLanguageCodeCa,
	"cs":    TranscriptionModelDetailsLanguageCodeCs,
	"cy":    TranscriptionModelDetailsLanguageCodeCy,
	"da":    TranscriptionModelDetailsLanguageCodeDa,
	"de":    TranscriptionModelDetailsLanguageCodeDe,
	"el":    TranscriptionModelDetailsLanguageCodeEl,
	"en":    TranscriptionModelDetailsLanguageCodeEn,
	"es":    TranscriptionModelDetailsLanguageCodeEs,
	"et":    TranscriptionModelDetailsLanguageCodeEt,
	"fa":    TranscriptionModelDetailsLanguageCodeFa,
	"fi":    TranscriptionModelDetailsLanguageCodeFi,
	"fr":    TranscriptionModelDetailsLanguageCodeFr,
	"gl":    TranscriptionModelDetailsLanguageCodeGl,
	"he":    TranscriptionModelDetailsLanguageCodeHe,
	"hi":    TranscriptionModelDetailsLanguageCodeHi,
	"hr":    TranscriptionModelDetailsLanguageCodeHr,
	"hu":    TranscriptionModelDetailsLanguageCodeHu,
	"hy":    TranscriptionModelDetailsLanguageCodeHy,
	"id":    TranscriptionModelDetailsLanguageCodeId,
	"is":    TranscriptionModelDetailsLanguageCodeIs,
	"it":    TranscriptionModelDetailsLanguageCodeIt,
	"ja":    TranscriptionModelDetailsLanguageCodeJa,
	"kk":    TranscriptionModelDetailsLanguageCodeKk,
	"kn":    TranscriptionModelDetailsLanguageCodeKn,
	"ko":    TranscriptionModelDetailsLanguageCodeKo,
	"lt":    TranscriptionModelDetailsLanguageCodeLt,
	"lv":    TranscriptionModelDetailsLanguageCodeLv,
	"mi":    TranscriptionModelDetailsLanguageCodeMi,
	"mk":    TranscriptionModelDetailsLanguageCodeMk,
	"mr":    TranscriptionModelDetailsLanguageCodeMr,
	"ms":    TranscriptionModelDetailsLanguageCodeMs,
	"ne":    TranscriptionModelDetailsLanguageCodeNe,
	"nl":    TranscriptionModelDetailsLanguageCodeNl,
	"no":    TranscriptionModelDetailsLanguageCodeNo,
	"pl":    TranscriptionModelDetailsLanguageCodePl,
	"pt":    TranscriptionModelDetailsLanguageCodePt,
	"ro":    TranscriptionModelDetailsLanguageCodeRo,
	"ru":    TranscriptionModelDetailsLanguageCodeRu,
	"sk":    TranscriptionModelDetailsLanguageCodeSk,
	"sl":    TranscriptionModelDetailsLanguageCodeSl,
	"sr":    TranscriptionModelDetailsLanguageCodeSr,
	"sv":    TranscriptionModelDetailsLanguageCodeSv,
	"sw":    TranscriptionModelDetailsLanguageCodeSw,
	"ta":    TranscriptionModelDetailsLanguageCodeTa,
	"th":    TranscriptionModelDetailsLanguageCodeTh,
	"tl":    TranscriptionModelDetailsLanguageCodeTl,
	"tr":    TranscriptionModelDetailsLanguageCodeTr,
	"uk":    TranscriptionModelDetailsLanguageCodeUk,
	"ur":    TranscriptionModelDetailsLanguageCodeUr,
	"vi":    TranscriptionModelDetailsLanguageCodeVi,
	"zh":    TranscriptionModelDetailsLanguageCodeZh,
}

// GetTranscriptionModelDetailsLanguageCodeEnumValues Enumerates the set of values for TranscriptionModelDetailsLanguageCodeEnum
func GetTranscriptionModelDetailsLanguageCodeEnumValues() []TranscriptionModelDetailsLanguageCodeEnum {
	values := make([]TranscriptionModelDetailsLanguageCodeEnum, 0)
	for _, v := range mappingTranscriptionModelDetailsLanguageCodeEnum {
		values = append(values, v)
	}
	return values
}

// GetTranscriptionModelDetailsLanguageCodeEnumStringValues Enumerates the set of values in String for TranscriptionModelDetailsLanguageCodeEnum
func GetTranscriptionModelDetailsLanguageCodeEnumStringValues() []string {
	return []string{
		"en-US",
		"es-ES",
		"pt-BR",
		"en-GB",
		"en-AU",
		"en-IN",
		"hi-IN",
		"fr-FR",
		"de-DE",
		"it-IT",
		"auto",
		"af",
		"ar",
		"az",
		"be",
		"bg",
		"bs",
		"ca",
		"cs",
		"cy",
		"da",
		"de",
		"el",
		"en",
		"es",
		"et",
		"fa",
		"fi",
		"fr",
		"gl",
		"he",
		"hi",
		"hr",
		"hu",
		"hy",
		"id",
		"is",
		"it",
		"ja",
		"kk",
		"kn",
		"ko",
		"lt",
		"lv",
		"mi",
		"mk",
		"mr",
		"ms",
		"ne",
		"nl",
		"no",
		"pl",
		"pt",
		"ro",
		"ru",
		"sk",
		"sl",
		"sr",
		"sv",
		"sw",
		"ta",
		"th",
		"tl",
		"tr",
		"uk",
		"ur",
		"vi",
		"zh",
	}
}

// GetMappingTranscriptionModelDetailsLanguageCodeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTranscriptionModelDetailsLanguageCodeEnum(val string) (TranscriptionModelDetailsLanguageCodeEnum, bool) {
	enum, ok := mappingTranscriptionModelDetailsLanguageCodeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
