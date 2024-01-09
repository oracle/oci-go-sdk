// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Managed Access API
//
// Use the Managed Access API to approve access requests, create and manage templates, and manage resource approval settings. For more information, see Managed Access Overview (https://docs.oracle.com/iaas/Content/managed-access/home.htm).
// Use the table of contents and search tool to explore the Managed Access API.
//

package lockbox

import (
	"strings"
)

// PersonaLevelEnum Enum with underlying type: string
type PersonaLevelEnum string

// Set of constants representing the allowable values for PersonaLevelEnum
const (
	PersonaLevelLevel1   PersonaLevelEnum = "LEVEL1"
	PersonaLevelLevel2   PersonaLevelEnum = "LEVEL2"
	PersonaLevelLevel3   PersonaLevelEnum = "LEVEL3"
	PersonaLevelAdmin    PersonaLevelEnum = "ADMIN"
	PersonaLevelOperator PersonaLevelEnum = "OPERATOR"
)

var mappingPersonaLevelEnum = map[string]PersonaLevelEnum{
	"LEVEL1":   PersonaLevelLevel1,
	"LEVEL2":   PersonaLevelLevel2,
	"LEVEL3":   PersonaLevelLevel3,
	"ADMIN":    PersonaLevelAdmin,
	"OPERATOR": PersonaLevelOperator,
}

var mappingPersonaLevelEnumLowerCase = map[string]PersonaLevelEnum{
	"level1":   PersonaLevelLevel1,
	"level2":   PersonaLevelLevel2,
	"level3":   PersonaLevelLevel3,
	"admin":    PersonaLevelAdmin,
	"operator": PersonaLevelOperator,
}

// GetPersonaLevelEnumValues Enumerates the set of values for PersonaLevelEnum
func GetPersonaLevelEnumValues() []PersonaLevelEnum {
	values := make([]PersonaLevelEnum, 0)
	for _, v := range mappingPersonaLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetPersonaLevelEnumStringValues Enumerates the set of values in String for PersonaLevelEnum
func GetPersonaLevelEnumStringValues() []string {
	return []string{
		"LEVEL1",
		"LEVEL2",
		"LEVEL3",
		"ADMIN",
		"OPERATOR",
	}
}

// GetMappingPersonaLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPersonaLevelEnum(val string) (PersonaLevelEnum, bool) {
	enum, ok := mappingPersonaLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
