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

// LockboxAutoApprovalStateEnum Enum with underlying type: string
type LockboxAutoApprovalStateEnum string

// Set of constants representing the allowable values for LockboxAutoApprovalStateEnum
const (
	LockboxAutoApprovalStateEnabled  LockboxAutoApprovalStateEnum = "ENABLED"
	LockboxAutoApprovalStateDisabled LockboxAutoApprovalStateEnum = "DISABLED"
)

var mappingLockboxAutoApprovalStateEnum = map[string]LockboxAutoApprovalStateEnum{
	"ENABLED":  LockboxAutoApprovalStateEnabled,
	"DISABLED": LockboxAutoApprovalStateDisabled,
}

var mappingLockboxAutoApprovalStateEnumLowerCase = map[string]LockboxAutoApprovalStateEnum{
	"enabled":  LockboxAutoApprovalStateEnabled,
	"disabled": LockboxAutoApprovalStateDisabled,
}

// GetLockboxAutoApprovalStateEnumValues Enumerates the set of values for LockboxAutoApprovalStateEnum
func GetLockboxAutoApprovalStateEnumValues() []LockboxAutoApprovalStateEnum {
	values := make([]LockboxAutoApprovalStateEnum, 0)
	for _, v := range mappingLockboxAutoApprovalStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLockboxAutoApprovalStateEnumStringValues Enumerates the set of values in String for LockboxAutoApprovalStateEnum
func GetLockboxAutoApprovalStateEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingLockboxAutoApprovalStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLockboxAutoApprovalStateEnum(val string) (LockboxAutoApprovalStateEnum, bool) {
	enum, ok := mappingLockboxAutoApprovalStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
