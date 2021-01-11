// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
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

// TargetDatabaseTypesEnum Enum with underlying type: string
type TargetDatabaseTypesEnum string

// Set of constants representing the allowable values for TargetDatabaseTypesEnum
const (
	TargetDatabaseTypesDatabaseSystem TargetDatabaseTypesEnum = "DATABASE_SYSTEM"
	TargetDatabaseTypesNotSet         TargetDatabaseTypesEnum = "NOT_SET"
)

var mappingTargetDatabaseTypes = map[string]TargetDatabaseTypesEnum{
	"DATABASE_SYSTEM": TargetDatabaseTypesDatabaseSystem,
	"NOT_SET":         TargetDatabaseTypesNotSet,
}

// GetTargetDatabaseTypesEnumValues Enumerates the set of values for TargetDatabaseTypesEnum
func GetTargetDatabaseTypesEnumValues() []TargetDatabaseTypesEnum {
	values := make([]TargetDatabaseTypesEnum, 0)
	for _, v := range mappingTargetDatabaseTypes {
		values = append(values, v)
	}
	return values
}
