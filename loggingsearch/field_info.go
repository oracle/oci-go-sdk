// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Search API
//
// Search for logs in your compartments, log groups, and log objects.
//

package loggingsearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FieldInfo Contains field schema information.
type FieldInfo struct {

	// Field name
	FieldName *string `mandatory:"true" json:"fieldName"`

	// Field type -
	// * `STRING`: A sequence of characters.
	// * `NUMBER`: Numeric type which can be an integer or floating point.
	// * `BOOLEAN`: Either true or false.
	// * `ARRAY`: An ordered collection of values.
	FieldType FieldInfoFieldTypeEnum `mandatory:"true" json:"fieldType"`
}

func (m FieldInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FieldInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFieldInfoFieldTypeEnum(string(m.FieldType)); !ok && m.FieldType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FieldType: %s. Supported values are: %s.", m.FieldType, strings.Join(GetFieldInfoFieldTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingFieldInfoFieldTypeEnum = map[string]FieldInfoFieldTypeEnum{
	"STRING":  FieldInfoFieldTypeString,
	"NUMBER":  FieldInfoFieldTypeNumber,
	"BOOLEAN": FieldInfoFieldTypeBoolean,
	"ARRAY":   FieldInfoFieldTypeArray,
}

var mappingFieldInfoFieldTypeEnumLowerCase = map[string]FieldInfoFieldTypeEnum{
	"string":  FieldInfoFieldTypeString,
	"number":  FieldInfoFieldTypeNumber,
	"boolean": FieldInfoFieldTypeBoolean,
	"array":   FieldInfoFieldTypeArray,
}

// GetFieldInfoFieldTypeEnumValues Enumerates the set of values for FieldInfoFieldTypeEnum
func GetFieldInfoFieldTypeEnumValues() []FieldInfoFieldTypeEnum {
	values := make([]FieldInfoFieldTypeEnum, 0)
	for _, v := range mappingFieldInfoFieldTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFieldInfoFieldTypeEnumStringValues Enumerates the set of values in String for FieldInfoFieldTypeEnum
func GetFieldInfoFieldTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"NUMBER",
		"BOOLEAN",
		"ARRAY",
	}
}

// GetMappingFieldInfoFieldTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFieldInfoFieldTypeEnum(val string) (FieldInfoFieldTypeEnum, bool) {
	enum, ok := mappingFieldInfoFieldTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
