// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// ResponderModeTypesEnum Enum with underlying type: string
type ResponderModeTypesEnum string

// Set of constants representing the allowable values for ResponderModeTypesEnum
const (
	ResponderModeTypesAutoaction ResponderModeTypesEnum = "AUTOACTION"
	ResponderModeTypesUseraction ResponderModeTypesEnum = "USERACTION"
)

var mappingResponderModeTypesEnum = map[string]ResponderModeTypesEnum{
	"AUTOACTION": ResponderModeTypesAutoaction,
	"USERACTION": ResponderModeTypesUseraction,
}

var mappingResponderModeTypesEnumLowerCase = map[string]ResponderModeTypesEnum{
	"autoaction": ResponderModeTypesAutoaction,
	"useraction": ResponderModeTypesUseraction,
}

// GetResponderModeTypesEnumValues Enumerates the set of values for ResponderModeTypesEnum
func GetResponderModeTypesEnumValues() []ResponderModeTypesEnum {
	values := make([]ResponderModeTypesEnum, 0)
	for _, v := range mappingResponderModeTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetResponderModeTypesEnumStringValues Enumerates the set of values in String for ResponderModeTypesEnum
func GetResponderModeTypesEnumStringValues() []string {
	return []string{
		"AUTOACTION",
		"USERACTION",
	}
}

// GetMappingResponderModeTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResponderModeTypesEnum(val string) (ResponderModeTypesEnum, bool) {
	enum, ok := mappingResponderModeTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
