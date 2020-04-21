// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

// OrchestrationVariableTypeEnumEnum Enum with underlying type: string
type OrchestrationVariableTypeEnumEnum string

// Set of constants representing the allowable values for OrchestrationVariableTypeEnumEnum
const (
	OrchestrationVariableTypeEnumString  OrchestrationVariableTypeEnumEnum = "STRING"
	OrchestrationVariableTypeEnumInteger OrchestrationVariableTypeEnumEnum = "INTEGER"
)

var mappingOrchestrationVariableTypeEnum = map[string]OrchestrationVariableTypeEnumEnum{
	"STRING":  OrchestrationVariableTypeEnumString,
	"INTEGER": OrchestrationVariableTypeEnumInteger,
}

// GetOrchestrationVariableTypeEnumEnumValues Enumerates the set of values for OrchestrationVariableTypeEnumEnum
func GetOrchestrationVariableTypeEnumEnumValues() []OrchestrationVariableTypeEnumEnum {
	values := make([]OrchestrationVariableTypeEnumEnum, 0)
	for _, v := range mappingOrchestrationVariableTypeEnum {
		values = append(values, v)
	}
	return values
}
