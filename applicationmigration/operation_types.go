// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

var mappingOperationTypesEnum = map[string]OperationTypesEnum{
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

var mappingOperationTypesEnumLowerCase = map[string]OperationTypesEnum{
	"create_source":                OperationTypesCreateSource,
	"update_source":                OperationTypesUpdateSource,
	"delete_source":                OperationTypesDeleteSource,
	"create_migration":             OperationTypesCreateMigration,
	"update_migration":             OperationTypesUpdateMigration,
	"delete_migration":             OperationTypesDeleteMigration,
	"authorize_source":             OperationTypesAuthorizeSource,
	"discover_application":         OperationTypesDiscoverApplication,
	"migrate_application":          OperationTypesMigrateApplication,
	"change_source_compartment":    OperationTypesChangeSourceCompartment,
	"change_migration_compartment": OperationTypesChangeMigrationCompartment,
}

// GetOperationTypesEnumValues Enumerates the set of values for OperationTypesEnum
func GetOperationTypesEnumValues() []OperationTypesEnum {
	values := make([]OperationTypesEnum, 0)
	for _, v := range mappingOperationTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypesEnumStringValues Enumerates the set of values in String for OperationTypesEnum
func GetOperationTypesEnumStringValues() []string {
	return []string{
		"CREATE_SOURCE",
		"UPDATE_SOURCE",
		"DELETE_SOURCE",
		"CREATE_MIGRATION",
		"UPDATE_MIGRATION",
		"DELETE_MIGRATION",
		"AUTHORIZE_SOURCE",
		"DISCOVER_APPLICATION",
		"MIGRATE_APPLICATION",
		"CHANGE_SOURCE_COMPARTMENT",
		"CHANGE_MIGRATION_COMPARTMENT",
	}
}

// GetMappingOperationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypesEnum(val string) (OperationTypesEnum, bool) {
	enum, ok := mappingOperationTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
