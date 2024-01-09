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

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateLockbox                OperationTypeEnum = "CREATE_LOCKBOX"
	OperationTypeUpdateLockbox                OperationTypeEnum = "UPDATE_LOCKBOX"
	OperationTypeDeleteLockbox                OperationTypeEnum = "DELETE_LOCKBOX"
	OperationTypeMoveLockbox                  OperationTypeEnum = "MOVE_LOCKBOX"
	OperationTypeCreateAccessRequest          OperationTypeEnum = "CREATE_ACCESS_REQUEST"
	OperationTypeApproveAccessRequest         OperationTypeEnum = "APPROVE_ACCESS_REQUEST"
	OperationTypeRevokeAccessRequest          OperationTypeEnum = "REVOKE_ACCESS_REQUEST"
	OperationTypeCreateApprovalTemplate       OperationTypeEnum = "CREATE_APPROVAL_TEMPLATE"
	OperationTypeMoveApprovalTemplate         OperationTypeEnum = "MOVE_APPROVAL_TEMPLATE"
	OperationTypeUpdateApprovalTemplate       OperationTypeEnum = "UPDATE_APPROVAL_TEMPLATE"
	OperationTypeDeleteApprovalTemplate       OperationTypeEnum = "DELETE_APPROVAL_TEMPLATE"
	OperationTypeCreatePartner                OperationTypeEnum = "CREATE_PARTNER"
	OperationTypeRemindAccessRequest          OperationTypeEnum = "REMIND_ACCESS_REQUEST"
	OperationTypeCreateAccesscontextattribute OperationTypeEnum = "CREATE_ACCESSCONTEXTATTRIBUTE"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_LOCKBOX":                OperationTypeCreateLockbox,
	"UPDATE_LOCKBOX":                OperationTypeUpdateLockbox,
	"DELETE_LOCKBOX":                OperationTypeDeleteLockbox,
	"MOVE_LOCKBOX":                  OperationTypeMoveLockbox,
	"CREATE_ACCESS_REQUEST":         OperationTypeCreateAccessRequest,
	"APPROVE_ACCESS_REQUEST":        OperationTypeApproveAccessRequest,
	"REVOKE_ACCESS_REQUEST":         OperationTypeRevokeAccessRequest,
	"CREATE_APPROVAL_TEMPLATE":      OperationTypeCreateApprovalTemplate,
	"MOVE_APPROVAL_TEMPLATE":        OperationTypeMoveApprovalTemplate,
	"UPDATE_APPROVAL_TEMPLATE":      OperationTypeUpdateApprovalTemplate,
	"DELETE_APPROVAL_TEMPLATE":      OperationTypeDeleteApprovalTemplate,
	"CREATE_PARTNER":                OperationTypeCreatePartner,
	"REMIND_ACCESS_REQUEST":         OperationTypeRemindAccessRequest,
	"CREATE_ACCESSCONTEXTATTRIBUTE": OperationTypeCreateAccesscontextattribute,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_lockbox":                OperationTypeCreateLockbox,
	"update_lockbox":                OperationTypeUpdateLockbox,
	"delete_lockbox":                OperationTypeDeleteLockbox,
	"move_lockbox":                  OperationTypeMoveLockbox,
	"create_access_request":         OperationTypeCreateAccessRequest,
	"approve_access_request":        OperationTypeApproveAccessRequest,
	"revoke_access_request":         OperationTypeRevokeAccessRequest,
	"create_approval_template":      OperationTypeCreateApprovalTemplate,
	"move_approval_template":        OperationTypeMoveApprovalTemplate,
	"update_approval_template":      OperationTypeUpdateApprovalTemplate,
	"delete_approval_template":      OperationTypeDeleteApprovalTemplate,
	"create_partner":                OperationTypeCreatePartner,
	"remind_access_request":         OperationTypeRemindAccessRequest,
	"create_accesscontextattribute": OperationTypeCreateAccesscontextattribute,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_LOCKBOX",
		"UPDATE_LOCKBOX",
		"DELETE_LOCKBOX",
		"MOVE_LOCKBOX",
		"CREATE_ACCESS_REQUEST",
		"APPROVE_ACCESS_REQUEST",
		"REVOKE_ACCESS_REQUEST",
		"CREATE_APPROVAL_TEMPLATE",
		"MOVE_APPROVAL_TEMPLATE",
		"UPDATE_APPROVAL_TEMPLATE",
		"DELETE_APPROVAL_TEMPLATE",
		"CREATE_PARTNER",
		"REMIND_ACCESS_REQUEST",
		"CREATE_ACCESSCONTEXTATTRIBUTE",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
