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
	PartialSilenceThresholdInMs *int `mandatory:"false" json:"partialSilenceThresholdInMs"`

	// Silence threshold for Realtime Speech final results in milliseconds.
	FinalSilenceThresholdInMs *int `mandatory:"false" json:"finalSilenceThresholdInMs"`

	// When enabled sets the amount of confidence required for latest tokens before returning them as part of a new partial result
	StabilizePartialResults RealtimeParametersStabilizePartialResultsEnum `mandatory:"false" json:"stabilizePartialResults,omitempty"`

	// Model Domain.
	ModelDomain RealtimeParametersModelDomainEnum `mandatory:"false" json:"modelDomain,omitempty"`

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
	LanguageCode *string `mandatory:"false" json:"languageCode"`

	// If set to true, the service will not fail connection attempt if it encounters any issues that prevent the loading of all specified user customizations. Any invalid customizations will simply be ignored and connection will continue being established with the default base model and any remaining valid customizations.
	// If set to false,  if the service is unable to load any of the specified customizations, an error detailing why will be returned and the session will end.
	ShouldIgnoreInvalidCustomizations *bool `mandatory:"false" json:"shouldIgnoreInvalidCustomizations"`

	// Array of customization objects.
	Customizations []CustomizationInference `mandatory:"false" json:"customizations"`
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
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
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
