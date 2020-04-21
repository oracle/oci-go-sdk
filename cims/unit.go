// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

// UnitEnum Enum with underlying type: string
type UnitEnum string

// Set of constants representing the allowable values for UnitEnum
const (
	UnitCount UnitEnum = "COUNT"
	UnitGb    UnitEnum = "GB"
	UnitNone  UnitEnum = "NONE"
)

var mappingUnit = map[string]UnitEnum{
	"COUNT": UnitCount,
	"GB":    UnitGb,
	"NONE":  UnitNone,
}

// GetUnitEnumValues Enumerates the set of values for UnitEnum
func GetUnitEnumValues() []UnitEnum {
	values := make([]UnitEnum, 0)
	for _, v := range mappingUnit {
		values = append(values, v)
	}
	return values
}
