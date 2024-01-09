// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GovernanceRulesControlPlane API
//
// A description of the GovernanceRulesControlPlane API
//

package governancerulescontrolplane

import (
	"strings"
)

// CreationOptionEnum Enum with underlying type: string
type CreationOptionEnum string

// Set of constants representing the allowable values for CreationOptionEnum
const (
	CreationOptionTemplate CreationOptionEnum = "TEMPLATE"
	CreationOptionClone    CreationOptionEnum = "CLONE"
)

var mappingCreationOptionEnum = map[string]CreationOptionEnum{
	"TEMPLATE": CreationOptionTemplate,
	"CLONE":    CreationOptionClone,
}

var mappingCreationOptionEnumLowerCase = map[string]CreationOptionEnum{
	"template": CreationOptionTemplate,
	"clone":    CreationOptionClone,
}

// GetCreationOptionEnumValues Enumerates the set of values for CreationOptionEnum
func GetCreationOptionEnumValues() []CreationOptionEnum {
	values := make([]CreationOptionEnum, 0)
	for _, v := range mappingCreationOptionEnum {
		values = append(values, v)
	}
	return values
}

// GetCreationOptionEnumStringValues Enumerates the set of values in String for CreationOptionEnum
func GetCreationOptionEnumStringValues() []string {
	return []string{
		"TEMPLATE",
		"CLONE",
	}
}

// GetMappingCreationOptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreationOptionEnum(val string) (CreationOptionEnum, bool) {
	enum, ok := mappingCreationOptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
