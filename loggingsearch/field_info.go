// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Search API
//
// Search for logs in your compartements / log groups / log objects.
//

package loggingsearch

import (
	"github.com/oracle/oci-go-sdk/common"
)

// FieldInfo Contains field schema information.
type FieldInfo struct {

	// Field name
	FieldName *string `mandatory:"true" json:"fieldName"`

	// Field type -
	// * `STRING`: A sequence of characters.
	// * `NUMBER`: Numeric type which can be integer or floating point.
	// * `BOOLEAN`: Either true or false.
	// * `ARRAY`: An ordered collection of values.
	FieldType FieldInfoFieldTypeEnum `mandatory:"true" json:"fieldType"`
}

func (m FieldInfo) String() string {
	return common.PointerString(m)
}

// FieldInfoFieldTypeEnum Enum with underlying type: string
type FieldInfoFieldTypeEnum string

// Set of constants representing the allowable values for FieldInfoFieldTypeEnum
const (
	FieldInfoFieldTypeString  FieldInfoFieldTypeEnum = "STRING"
	FieldInfoFieldTypeNumber  FieldInfoFieldTypeEnum = "NUMBER"
	FieldInfoFieldTypeBoolean FieldInfoFieldTypeEnum = "BOOLEAN"
	FieldInfoFieldTypeArray   FieldInfoFieldTypeEnum = "ARRAY"
)

var mappingFieldInfoFieldType = map[string]FieldInfoFieldTypeEnum{
	"STRING":  FieldInfoFieldTypeString,
	"NUMBER":  FieldInfoFieldTypeNumber,
	"BOOLEAN": FieldInfoFieldTypeBoolean,
	"ARRAY":   FieldInfoFieldTypeArray,
}

// GetFieldInfoFieldTypeEnumValues Enumerates the set of values for FieldInfoFieldTypeEnum
func GetFieldInfoFieldTypeEnumValues() []FieldInfoFieldTypeEnum {
	values := make([]FieldInfoFieldTypeEnum, 0)
	for _, v := range mappingFieldInfoFieldType {
		values = append(values, v)
	}
	return values
}
