// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// AccessRequestActionTypeEnum Enum with underlying type: string
type AccessRequestActionTypeEnum string

// Set of constants representing the allowable values for AccessRequestActionTypeEnum
const (
	AccessRequestActionTypeApprove AccessRequestActionTypeEnum = "APPROVE"
	AccessRequestActionTypeDeny    AccessRequestActionTypeEnum = "DENY"
	AccessRequestActionTypeRevoke  AccessRequestActionTypeEnum = "REVOKE"
	AccessRequestActionTypeCancel  AccessRequestActionTypeEnum = "CANCEL"
)

var mappingAccessRequestActionTypeEnum = map[string]AccessRequestActionTypeEnum{
	"APPROVE": AccessRequestActionTypeApprove,
	"DENY":    AccessRequestActionTypeDeny,
	"REVOKE":  AccessRequestActionTypeRevoke,
	"CANCEL":  AccessRequestActionTypeCancel,
}

var mappingAccessRequestActionTypeEnumLowerCase = map[string]AccessRequestActionTypeEnum{
	"approve": AccessRequestActionTypeApprove,
	"deny":    AccessRequestActionTypeDeny,
	"revoke":  AccessRequestActionTypeRevoke,
	"cancel":  AccessRequestActionTypeCancel,
}

// GetAccessRequestActionTypeEnumValues Enumerates the set of values for AccessRequestActionTypeEnum
func GetAccessRequestActionTypeEnumValues() []AccessRequestActionTypeEnum {
	values := make([]AccessRequestActionTypeEnum, 0)
	for _, v := range mappingAccessRequestActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAccessRequestActionTypeEnumStringValues Enumerates the set of values in String for AccessRequestActionTypeEnum
func GetAccessRequestActionTypeEnumStringValues() []string {
	return []string{
		"APPROVE",
		"DENY",
		"REVOKE",
		"CANCEL",
	}
}

// GetMappingAccessRequestActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccessRequestActionTypeEnum(val string) (AccessRequestActionTypeEnum, bool) {
	enum, ok := mappingAccessRequestActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
