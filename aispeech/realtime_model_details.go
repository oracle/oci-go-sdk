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

// RealtimeModelDetails Details of the ASR model used by the realtime speech service.
type RealtimeModelDetails struct {

	// Model Domain.
	Domain RealtimeModelDetailsDomainEnum `mandatory:"false" json:"domain,omitempty"`

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
}

func (m RealtimeModelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RealtimeModelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRealtimeModelDetailsDomainEnum(string(m.Domain)); !ok && m.Domain != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Domain: %s. Supported values are: %s.", m.Domain, strings.Join(GetRealtimeModelDetailsDomainEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RealtimeModelDetailsDomainEnum Enum with underlying type: string
type RealtimeModelDetailsDomainEnum string

// Set of constants representing the allowable values for RealtimeModelDetailsDomainEnum
const (
	RealtimeModelDetailsDomainGeneric RealtimeModelDetailsDomainEnum = "GENERIC"
	RealtimeModelDetailsDomainMedical RealtimeModelDetailsDomainEnum = "MEDICAL"
)

var mappingRealtimeModelDetailsDomainEnum = map[string]RealtimeModelDetailsDomainEnum{
	"GENERIC": RealtimeModelDetailsDomainGeneric,
	"MEDICAL": RealtimeModelDetailsDomainMedical,
}

var mappingRealtimeModelDetailsDomainEnumLowerCase = map[string]RealtimeModelDetailsDomainEnum{
	"generic": RealtimeModelDetailsDomainGeneric,
	"medical": RealtimeModelDetailsDomainMedical,
}

// GetRealtimeModelDetailsDomainEnumValues Enumerates the set of values for RealtimeModelDetailsDomainEnum
func GetRealtimeModelDetailsDomainEnumValues() []RealtimeModelDetailsDomainEnum {
	values := make([]RealtimeModelDetailsDomainEnum, 0)
	for _, v := range mappingRealtimeModelDetailsDomainEnum {
		values = append(values, v)
	}
	return values
}

// GetRealtimeModelDetailsDomainEnumStringValues Enumerates the set of values in String for RealtimeModelDetailsDomainEnum
func GetRealtimeModelDetailsDomainEnumStringValues() []string {
	return []string{
		"GENERIC",
		"MEDICAL",
	}
}

// GetMappingRealtimeModelDetailsDomainEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRealtimeModelDetailsDomainEnum(val string) (RealtimeModelDetailsDomainEnum, bool) {
	enum, ok := mappingRealtimeModelDetailsDomainEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
