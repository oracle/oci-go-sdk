// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Managed Services for Mac API
//
// Use the OCI Managed Services for Mac API to create and manage orders for dedicated, partially-managed Mac servers hosted in an OCI Data Center. For more information, see the user guide documentation for the [OCI Managed Services for Mac]()
//

package mngdmac

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateMacOrder  OperationTypeEnum = "CREATE_MAC_ORDER"
	OperationTypeUpdateMacOrder  OperationTypeEnum = "UPDATE_MAC_ORDER"
	OperationTypeCancelMacOrder  OperationTypeEnum = "CANCEL_MAC_ORDER"
	OperationTypeDeleteMacOrder  OperationTypeEnum = "DELETE_MAC_ORDER"
	OperationTypeMoveMacOrder    OperationTypeEnum = "MOVE_MAC_ORDER"
	OperationTypeDeleteMacDevice OperationTypeEnum = "DELETE_MAC_DEVICE"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_MAC_ORDER":  OperationTypeCreateMacOrder,
	"UPDATE_MAC_ORDER":  OperationTypeUpdateMacOrder,
	"CANCEL_MAC_ORDER":  OperationTypeCancelMacOrder,
	"DELETE_MAC_ORDER":  OperationTypeDeleteMacOrder,
	"MOVE_MAC_ORDER":    OperationTypeMoveMacOrder,
	"DELETE_MAC_DEVICE": OperationTypeDeleteMacDevice,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_mac_order":  OperationTypeCreateMacOrder,
	"update_mac_order":  OperationTypeUpdateMacOrder,
	"cancel_mac_order":  OperationTypeCancelMacOrder,
	"delete_mac_order":  OperationTypeDeleteMacOrder,
	"move_mac_order":    OperationTypeMoveMacOrder,
	"delete_mac_device": OperationTypeDeleteMacDevice,
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
		"CREATE_MAC_ORDER",
		"UPDATE_MAC_ORDER",
		"CANCEL_MAC_ORDER",
		"DELETE_MAC_ORDER",
		"MOVE_MAC_ORDER",
		"DELETE_MAC_DEVICE",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
