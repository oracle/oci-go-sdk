// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests.
// For more information, see Getting Help and Contacting Support (https://docs.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm).
// **Note**: Before you can create service requests with this API,
// complete user registration at My Oracle Cloud Support
// and then ask your tenancy administrator to provide you authorization for the related user groups.
//

package cims

import (
	"strings"
)

// ProblemTypeEnum Enum with underlying type: string
type ProblemTypeEnum string

// Set of constants representing the allowable values for ProblemTypeEnum
const (
	ProblemTypeLimit       ProblemTypeEnum = "LIMIT"
	ProblemTypeLegacyLimit ProblemTypeEnum = "LEGACY_LIMIT"
	ProblemTypeTech        ProblemTypeEnum = "TECH"
	ProblemTypeAccount     ProblemTypeEnum = "ACCOUNT"
	ProblemTypeTaxonomy    ProblemTypeEnum = "TAXONOMY"
)

var mappingProblemTypeEnum = map[string]ProblemTypeEnum{
	"LIMIT":        ProblemTypeLimit,
	"LEGACY_LIMIT": ProblemTypeLegacyLimit,
	"TECH":         ProblemTypeTech,
	"ACCOUNT":      ProblemTypeAccount,
	"TAXONOMY":     ProblemTypeTaxonomy,
}

var mappingProblemTypeEnumLowerCase = map[string]ProblemTypeEnum{
	"limit":        ProblemTypeLimit,
	"legacy_limit": ProblemTypeLegacyLimit,
	"tech":         ProblemTypeTech,
	"account":      ProblemTypeAccount,
	"taxonomy":     ProblemTypeTaxonomy,
}

// GetProblemTypeEnumValues Enumerates the set of values for ProblemTypeEnum
func GetProblemTypeEnumValues() []ProblemTypeEnum {
	values := make([]ProblemTypeEnum, 0)
	for _, v := range mappingProblemTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetProblemTypeEnumStringValues Enumerates the set of values in String for ProblemTypeEnum
func GetProblemTypeEnumStringValues() []string {
	return []string{
		"LIMIT",
		"LEGACY_LIMIT",
		"TECH",
		"ACCOUNT",
		"TAXONOMY",
	}
}

// GetMappingProblemTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProblemTypeEnum(val string) (ProblemTypeEnum, bool) {
	enum, ok := mappingProblemTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
