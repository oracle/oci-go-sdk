// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"strings"
)

// EnclosureTypeEnum Enum with underlying type: string
type EnclosureTypeEnum string

// Set of constants representing the allowable values for EnclosureTypeEnum
const (
	EnclosureTypeRuggadized    EnclosureTypeEnum = "RUGGADIZED"
	EnclosureTypeNonRuggadized EnclosureTypeEnum = "NON_RUGGADIZED"
)

var mappingEnclosureTypeEnum = map[string]EnclosureTypeEnum{
	"RUGGADIZED":     EnclosureTypeRuggadized,
	"NON_RUGGADIZED": EnclosureTypeNonRuggadized,
}

var mappingEnclosureTypeEnumLowerCase = map[string]EnclosureTypeEnum{
	"ruggadized":     EnclosureTypeRuggadized,
	"non_ruggadized": EnclosureTypeNonRuggadized,
}

// GetEnclosureTypeEnumValues Enumerates the set of values for EnclosureTypeEnum
func GetEnclosureTypeEnumValues() []EnclosureTypeEnum {
	values := make([]EnclosureTypeEnum, 0)
	for _, v := range mappingEnclosureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEnclosureTypeEnumStringValues Enumerates the set of values in String for EnclosureTypeEnum
func GetEnclosureTypeEnumStringValues() []string {
	return []string{
		"RUGGADIZED",
		"NON_RUGGADIZED",
	}
}

// GetMappingEnclosureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnclosureTypeEnum(val string) (EnclosureTypeEnum, bool) {
	enum, ok := mappingEnclosureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
