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

// CustomizationModelDetails Customization details.
type CustomizationModelDetails struct {

	// Customization Domain
	Domain CustomizationModelDetailsDomainEnum `mandatory:"false" json:"domain,omitempty"`

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

func (m CustomizationModelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomizationModelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCustomizationModelDetailsDomainEnum(string(m.Domain)); !ok && m.Domain != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Domain: %s. Supported values are: %s.", m.Domain, strings.Join(GetCustomizationModelDetailsDomainEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CustomizationModelDetailsDomainEnum Enum with underlying type: string
type CustomizationModelDetailsDomainEnum string

// Set of constants representing the allowable values for CustomizationModelDetailsDomainEnum
const (
	CustomizationModelDetailsDomainGeneric CustomizationModelDetailsDomainEnum = "GENERIC"
	CustomizationModelDetailsDomainMedical CustomizationModelDetailsDomainEnum = "MEDICAL"
)

var mappingCustomizationModelDetailsDomainEnum = map[string]CustomizationModelDetailsDomainEnum{
	"GENERIC": CustomizationModelDetailsDomainGeneric,
	"MEDICAL": CustomizationModelDetailsDomainMedical,
}

var mappingCustomizationModelDetailsDomainEnumLowerCase = map[string]CustomizationModelDetailsDomainEnum{
	"generic": CustomizationModelDetailsDomainGeneric,
	"medical": CustomizationModelDetailsDomainMedical,
}

// GetCustomizationModelDetailsDomainEnumValues Enumerates the set of values for CustomizationModelDetailsDomainEnum
func GetCustomizationModelDetailsDomainEnumValues() []CustomizationModelDetailsDomainEnum {
	values := make([]CustomizationModelDetailsDomainEnum, 0)
	for _, v := range mappingCustomizationModelDetailsDomainEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomizationModelDetailsDomainEnumStringValues Enumerates the set of values in String for CustomizationModelDetailsDomainEnum
func GetCustomizationModelDetailsDomainEnumStringValues() []string {
	return []string{
		"GENERIC",
		"MEDICAL",
	}
}

// GetMappingCustomizationModelDetailsDomainEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomizationModelDetailsDomainEnum(val string) (CustomizationModelDetailsDomainEnum, bool) {
	enum, ok := mappingCustomizationModelDetailsDomainEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
