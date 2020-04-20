// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

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

var mappingMigrationTypes = map[string]MigrationTypesEnum{
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
	for _, v := range mappingMigrationTypes {
		values = append(values, v)
	}
	return values
}
