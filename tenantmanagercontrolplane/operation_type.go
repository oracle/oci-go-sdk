// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// TenantManager API
//
// A description of the TenantManager API
//

package tenantmanagercontrolplane

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateSenderInvitation    OperationTypeEnum = "CREATE_SENDER_INVITATION"
	OperationTypeAcceptRecipientInvitation OperationTypeEnum = "ACCEPT_RECIPIENT_INVITATION"
	OperationTypeCancelSenderInvitation    OperationTypeEnum = "CANCEL_SENDER_INVITATION"
)

var mappingOperationType = map[string]OperationTypeEnum{
	"CREATE_SENDER_INVITATION":    OperationTypeCreateSenderInvitation,
	"ACCEPT_RECIPIENT_INVITATION": OperationTypeAcceptRecipientInvitation,
	"CANCEL_SENDER_INVITATION":    OperationTypeCancelSenderInvitation,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationType {
		values = append(values, v)
	}
	return values
}
