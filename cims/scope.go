// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests.
// For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm).
// **Note**: Before you can create service requests with this API,
// complete user registration at My Oracle Cloud Support
// and then ask your tenancy administrator to provide you authorization for the related user groups.
//

package cims

import (
	"strings"
)

// ScopeEnum Enum with underlying type: string
type ScopeEnum string

// Set of constants representing the allowable values for ScopeEnum
const (
	ScopeAd      ScopeEnum = "AD"
	ScopeRegion  ScopeEnum = "REGION"
	ScopeTenancy ScopeEnum = "TENANCY"
	ScopeNone    ScopeEnum = "NONE"
)

var mappingScopeEnum = map[string]ScopeEnum{
	"AD":      ScopeAd,
	"REGION":  ScopeRegion,
	"TENANCY": ScopeTenancy,
	"NONE":    ScopeNone,
}

var mappingScopeEnumLowerCase = map[string]ScopeEnum{
	"ad":      ScopeAd,
	"region":  ScopeRegion,
	"tenancy": ScopeTenancy,
	"none":    ScopeNone,
}

// GetScopeEnumValues Enumerates the set of values for ScopeEnum
func GetScopeEnumValues() []ScopeEnum {
	values := make([]ScopeEnum, 0)
	for _, v := range mappingScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetScopeEnumStringValues Enumerates the set of values in String for ScopeEnum
func GetScopeEnumStringValues() []string {
	return []string{
		"AD",
		"REGION",
		"TENANCY",
		"NONE",
	}
}

// GetMappingScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScopeEnum(val string) (ScopeEnum, bool) {
	enum, ok := mappingScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
