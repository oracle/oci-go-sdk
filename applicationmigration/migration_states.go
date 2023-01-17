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

var mappingMigrationStatesEnum = map[string]MigrationStatesEnum{
	"DISCOVERING_APPLICATION": MigrationStatesDiscoveringApplication,
	"DISCOVERY_FAILED":        MigrationStatesDiscoveryFailed,
	"MISSING_CONFIG_VALUES":   MigrationStatesMissingConfigValues,
	"READY":                   MigrationStatesReady,
	"MIGRATING":               MigrationStatesMigrating,
	"MIGRATION_FAILED":        MigrationStatesMigrationFailed,
	"MIGRATION_SUCCEEDED":     MigrationStatesMigrationSucceeded,
}

var mappingMigrationStatesEnumLowerCase = map[string]MigrationStatesEnum{
	"discovering_application": MigrationStatesDiscoveringApplication,
	"discovery_failed":        MigrationStatesDiscoveryFailed,
	"missing_config_values":   MigrationStatesMissingConfigValues,
	"ready":                   MigrationStatesReady,
	"migrating":               MigrationStatesMigrating,
	"migration_failed":        MigrationStatesMigrationFailed,
	"migration_succeeded":     MigrationStatesMigrationSucceeded,
}

// GetMigrationStatesEnumValues Enumerates the set of values for MigrationStatesEnum
func GetMigrationStatesEnumValues() []MigrationStatesEnum {
	values := make([]MigrationStatesEnum, 0)
	for _, v := range mappingMigrationStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetMigrationStatesEnumStringValues Enumerates the set of values in String for MigrationStatesEnum
func GetMigrationStatesEnumStringValues() []string {
	return []string{
		"DISCOVERING_APPLICATION",
		"DISCOVERY_FAILED",
		"MISSING_CONFIG_VALUES",
		"READY",
		"MIGRATING",
		"MIGRATION_FAILED",
		"MIGRATION_SUCCEEDED",
	}
}

// GetMappingMigrationStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMigrationStatesEnum(val string) (MigrationStatesEnum, bool) {
	enum, ok := mappingMigrationStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
