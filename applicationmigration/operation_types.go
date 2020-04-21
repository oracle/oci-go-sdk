// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesCreateSource               OperationTypesEnum = "CREATE_SOURCE"
	OperationTypesUpdateSource               OperationTypesEnum = "UPDATE_SOURCE"
	OperationTypesDeleteSource               OperationTypesEnum = "DELETE_SOURCE"
	OperationTypesCreateMigration            OperationTypesEnum = "CREATE_MIGRATION"
	OperationTypesUpdateMigration            OperationTypesEnum = "UPDATE_MIGRATION"
	OperationTypesDeleteMigration            OperationTypesEnum = "DELETE_MIGRATION"
	OperationTypesAuthorizeSource            OperationTypesEnum = "AUTHORIZE_SOURCE"
	OperationTypesDiscoverApplication        OperationTypesEnum = "DISCOVER_APPLICATION"
	OperationTypesMigrateApplication         OperationTypesEnum = "MIGRATE_APPLICATION"
	OperationTypesChangeSourceCompartment    OperationTypesEnum = "CHANGE_SOURCE_COMPARTMENT"
	OperationTypesChangeMigrationCompartment OperationTypesEnum = "CHANGE_MIGRATION_COMPARTMENT"
)

var mappingOperationTypes = map[string]OperationTypesEnum{
	"CREATE_SOURCE":                OperationTypesCreateSource,
	"UPDATE_SOURCE":                OperationTypesUpdateSource,
	"DELETE_SOURCE":                OperationTypesDeleteSource,
	"CREATE_MIGRATION":             OperationTypesCreateMigration,
	"UPDATE_MIGRATION":             OperationTypesUpdateMigration,
	"DELETE_MIGRATION":             OperationTypesDeleteMigration,
	"AUTHORIZE_SOURCE":             OperationTypesAuthorizeSource,
	"DISCOVER_APPLICATION":         OperationTypesDiscoverApplication,
	"MIGRATE_APPLICATION":          OperationTypesMigrateApplication,
	"CHANGE_SOURCE_COMPARTMENT":    OperationTypesChangeSourceCompartment,
	"CHANGE_MIGRATION_COMPARTMENT": OperationTypesChangeMigrationCompartment,
}

// GetOperationTypesEnumValues Enumerates the set of values for OperationTypesEnum
func GetOperationTypesEnumValues() []OperationTypesEnum {
	values := make([]OperationTypesEnum, 0)
	for _, v := range mappingOperationTypes {
		values = append(values, v)
	}
	return values
}
