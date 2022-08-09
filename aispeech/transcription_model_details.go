// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

	// Domain for input files.
	Domain TranscriptionModelDetailsDomainEnum `mandatory:"false" json:"domain,omitempty"`

	// Locale value as per given in [https://datatracker.ietf.org/doc/html/rfc5646].
	// - en-US: English - United States
	// - es-ES: Spanish - Spain
	// - pt-BR: Portuguese - Brazil
	// - en-GB: English - Great Britain
	// - en-AU: English - Australia
	// - en-IN: English - India
	// - hi-IN: Hindi - India
	// - fr-FR: French - France
	// - de-DE: German - Germany
	// - it-IT: Italian - Italy
	LanguageCode TranscriptionModelDetailsLanguageCodeEnum `mandatory:"false" json:"languageCode,omitempty"`
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
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
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
	}
}

// GetMappingTranscriptionModelDetailsLanguageCodeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTranscriptionModelDetailsLanguageCodeEnum(val string) (TranscriptionModelDetailsLanguageCodeEnum, bool) {
	enum, ok := mappingTranscriptionModelDetailsLanguageCodeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
