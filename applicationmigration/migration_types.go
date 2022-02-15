// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration API
//
// Application Migration simplifies the migration of applications from Oracle Cloud Infrastructure Classic to Oracle Cloud Infrastructure.
// You can use Application Migration API to migrate applications, such as Oracle Java Cloud Service, SOA Cloud Service, and Integration Classic
// instances, to Oracle Cloud Infrastructure. For more information, see
// Overview of Application Migration (https://docs.cloud.oracle.com/iaas/application-migration/appmigrationoverview.htm).
//

package applicationmigration

import (
	"strings"
)

// MigrationTypesEnum Enum with underlying type: string
type MigrationTypesEnum string

// Set of constants representing the allowable values for MigrationTypesEnum
const (
	MigrationTypesJcs   MigrationTypesEnum = "JCS"
	MigrationTypesSoacs MigrationTypesEnum = "SOACS"
	MigrationTypesOic   MigrationTypesEnum = "OIC"
	MigrationTypesOac   MigrationTypesEnum = "OAC"
	MigrationTypesIcs   MigrationTypesEnum = "ICS"
	MigrationTypesPcs   MigrationTypesEnum = "PCS"
)

var mappingMigrationTypesEnum = map[string]MigrationTypesEnum{
	"JCS":   MigrationTypesJcs,
	"SOACS": MigrationTypesSoacs,
	"OIC":   MigrationTypesOic,
	"OAC":   MigrationTypesOac,
	"ICS":   MigrationTypesIcs,
	"PCS":   MigrationTypesPcs,
}

// GetMigrationTypesEnumValues Enumerates the set of values for MigrationTypesEnum
func GetMigrationTypesEnumValues() []MigrationTypesEnum {
	values := make([]MigrationTypesEnum, 0)
	for _, v := range mappingMigrationTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetMigrationTypesEnumStringValues Enumerates the set of values in String for MigrationTypesEnum
func GetMigrationTypesEnumStringValues() []string {
	return []string{
		"JCS",
		"SOACS",
		"OIC",
		"OAC",
		"ICS",
		"PCS",
	}
}

// GetMappingMigrationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMigrationTypesEnum(val string) (MigrationTypesEnum, bool) {
	mappingMigrationTypesEnumIgnoreCase := make(map[string]MigrationTypesEnum)
	for k, v := range mappingMigrationTypesEnum {
		mappingMigrationTypesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingMigrationTypesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
