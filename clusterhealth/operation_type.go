// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ITV Control Plane - Diagnosis Store API
//
// Use the ITV Control Plane Diagnosis Store API to manage diagnosis.
//

package clusterhealth

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateDiagnosisStore OperationTypeEnum = "CREATE_DIAGNOSIS_STORE"
	OperationTypeUpdateDiagnosisStore OperationTypeEnum = "UPDATE_DIAGNOSIS_STORE"
	OperationTypeDeleteDiagnosisStore OperationTypeEnum = "DELETE_DIAGNOSIS_STORE"
	OperationTypeMoveDiagnosisStore   OperationTypeEnum = "MOVE_DIAGNOSIS_STORE"
	OperationTypeRequestDiagnosis     OperationTypeEnum = "REQUEST_DIAGNOSIS"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_DIAGNOSIS_STORE": OperationTypeCreateDiagnosisStore,
	"UPDATE_DIAGNOSIS_STORE": OperationTypeUpdateDiagnosisStore,
	"DELETE_DIAGNOSIS_STORE": OperationTypeDeleteDiagnosisStore,
	"MOVE_DIAGNOSIS_STORE":   OperationTypeMoveDiagnosisStore,
	"REQUEST_DIAGNOSIS":      OperationTypeRequestDiagnosis,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_diagnosis_store": OperationTypeCreateDiagnosisStore,
	"update_diagnosis_store": OperationTypeUpdateDiagnosisStore,
	"delete_diagnosis_store": OperationTypeDeleteDiagnosisStore,
	"move_diagnosis_store":   OperationTypeMoveDiagnosisStore,
	"request_diagnosis":      OperationTypeRequestDiagnosis,
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
		"CREATE_DIAGNOSIS_STORE",
		"UPDATE_DIAGNOSIS_STORE",
		"DELETE_DIAGNOSIS_STORE",
		"MOVE_DIAGNOSIS_STORE",
		"REQUEST_DIAGNOSIS",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
