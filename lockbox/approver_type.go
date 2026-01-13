// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ApproverTypeEnum Enum with underlying type: string
type ApproverTypeEnum string

// Set of constants representing the allowable values for ApproverTypeEnum
const (
	ApproverTypeGroup ApproverTypeEnum = "GROUP"
	ApproverTypeUser  ApproverTypeEnum = "USER"
)

var mappingApproverTypeEnum = map[string]ApproverTypeEnum{
	"GROUP": ApproverTypeGroup,
	"USER":  ApproverTypeUser,
}

var mappingApproverTypeEnumLowerCase = map[string]ApproverTypeEnum{
	"group": ApproverTypeGroup,
	"user":  ApproverTypeUser,
}

// GetApproverTypeEnumValues Enumerates the set of values for ApproverTypeEnum
func GetApproverTypeEnumValues() []ApproverTypeEnum {
	values := make([]ApproverTypeEnum, 0)
	for _, v := range mappingApproverTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetApproverTypeEnumStringValues Enumerates the set of values in String for ApproverTypeEnum
func GetApproverTypeEnumStringValues() []string {
	return []string{
		"GROUP",
		"USER",
	}
}

// GetMappingApproverTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApproverTypeEnum(val string) (ApproverTypeEnum, bool) {
	enum, ok := mappingApproverTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
