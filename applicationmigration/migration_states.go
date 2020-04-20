// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

// MigrationStatesEnum Enum with underlying type: string
type MigrationStatesEnum string

// Set of constants representing the allowable values for MigrationStatesEnum
const (
	MigrationStatesDiscoveringApplication MigrationStatesEnum = "DISCOVERING_APPLICATION"
	MigrationStatesDiscoveryFailed        MigrationStatesEnum = "DISCOVERY_FAILED"
	MigrationStatesMissingConfigValues    MigrationStatesEnum = "MISSING_CONFIG_VALUES"
	MigrationStatesReady                  MigrationStatesEnum = "READY"
	MigrationStatesMigrating              MigrationStatesEnum = "MIGRATING"
	MigrationStatesMigrationFailed        MigrationStatesEnum = "MIGRATION_FAILED"
	MigrationStatesMigrationSucceeded     MigrationStatesEnum = "MIGRATION_SUCCEEDED"
)

var mappingMigrationStates = map[string]MigrationStatesEnum{
	"DISCOVERING_APPLICATION": MigrationStatesDiscoveringApplication,
	"DISCOVERY_FAILED":        MigrationStatesDiscoveryFailed,
	"MISSING_CONFIG_VALUES":   MigrationStatesMissingConfigValues,
	"READY":                   MigrationStatesReady,
	"MIGRATING":               MigrationStatesMigrating,
	"MIGRATION_FAILED":        MigrationStatesMigrationFailed,
	"MIGRATION_SUCCEEDED":     MigrationStatesMigrationSucceeded,
}

// GetMigrationStatesEnumValues Enumerates the set of values for MigrationStatesEnum
func GetMigrationStatesEnumValues() []MigrationStatesEnum {
	values := make([]MigrationStatesEnum, 0)
	for _, v := range mappingMigrationStates {
		values = append(values, v)
	}
	return values
}
