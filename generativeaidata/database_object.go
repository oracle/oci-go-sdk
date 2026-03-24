// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Nl2sql API
//
// A description of the ReferenceService API. in progress
//

package generativeaidata

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseObject The type of database object
type DatabaseObject struct {

	// The type of the database object (e.g., table, view, procedure).
	Type DatabaseObjectTypeEnum `mandatory:"true" json:"type"`

	// The fully qualified name of the database object.
	Name *string `mandatory:"true" json:"name"`

	// An optional description of the DatabaseObject.
	Description *string `mandatory:"false" json:"description"`
}

func (m DatabaseObject) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseObject) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseObjectTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDatabaseObjectTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseObjectTypeEnum Enum with underlying type: string
type DatabaseObjectTypeEnum string

// Set of constants representing the allowable values for DatabaseObjectTypeEnum
const (
	DatabaseObjectTypeView       DatabaseObjectTypeEnum = "VIEW"
	DatabaseObjectTypeTable      DatabaseObjectTypeEnum = "TABLE"
	DatabaseObjectTypeAnnotation DatabaseObjectTypeEnum = "ANNOTATION"
	DatabaseObjectTypeComment    DatabaseObjectTypeEnum = "COMMENT"
)

var mappingDatabaseObjectTypeEnum = map[string]DatabaseObjectTypeEnum{
	"VIEW":       DatabaseObjectTypeView,
	"TABLE":      DatabaseObjectTypeTable,
	"ANNOTATION": DatabaseObjectTypeAnnotation,
	"COMMENT":    DatabaseObjectTypeComment,
}

var mappingDatabaseObjectTypeEnumLowerCase = map[string]DatabaseObjectTypeEnum{
	"view":       DatabaseObjectTypeView,
	"table":      DatabaseObjectTypeTable,
	"annotation": DatabaseObjectTypeAnnotation,
	"comment":    DatabaseObjectTypeComment,
}

// GetDatabaseObjectTypeEnumValues Enumerates the set of values for DatabaseObjectTypeEnum
func GetDatabaseObjectTypeEnumValues() []DatabaseObjectTypeEnum {
	values := make([]DatabaseObjectTypeEnum, 0)
	for _, v := range mappingDatabaseObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseObjectTypeEnumStringValues Enumerates the set of values in String for DatabaseObjectTypeEnum
func GetDatabaseObjectTypeEnumStringValues() []string {
	return []string{
		"VIEW",
		"TABLE",
		"ANNOTATION",
		"COMMENT",
	}
}

// GetMappingDatabaseObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseObjectTypeEnum(val string) (DatabaseObjectTypeEnum, bool) {
	enum, ok := mappingDatabaseObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
