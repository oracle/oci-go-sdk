// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// PropertySortByEnum Enum with underlying type: string
type PropertySortByEnum string

// Set of constants representing the allowable values for PropertySortByEnum
const (
	PropertySortByTimeCreated PropertySortByEnum = "timeCreated"
	PropertySortByDisplayName PropertySortByEnum = "displayName"
)

var mappingPropertySortByEnum = map[string]PropertySortByEnum{
	"timeCreated": PropertySortByTimeCreated,
	"displayName": PropertySortByDisplayName,
}

var mappingPropertySortByEnumLowerCase = map[string]PropertySortByEnum{
	"timecreated": PropertySortByTimeCreated,
	"displayname": PropertySortByDisplayName,
}

// GetPropertySortByEnumValues Enumerates the set of values for PropertySortByEnum
func GetPropertySortByEnumValues() []PropertySortByEnum {
	values := make([]PropertySortByEnum, 0)
	for _, v := range mappingPropertySortByEnum {
		values = append(values, v)
	}
	return values
}

// GetPropertySortByEnumStringValues Enumerates the set of values in String for PropertySortByEnum
func GetPropertySortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingPropertySortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPropertySortByEnum(val string) (PropertySortByEnum, bool) {
	enum, ok := mappingPropertySortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
