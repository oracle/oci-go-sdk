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

// MigrationLifecycleStatesEnum Enum with underlying type: string
type MigrationLifecycleStatesEnum string

// Set of constants representing the allowable values for MigrationLifecycleStatesEnum
const (
	MigrationLifecycleStatesCreating  MigrationLifecycleStatesEnum = "CREATING"
	MigrationLifecycleStatesActive    MigrationLifecycleStatesEnum = "ACTIVE"
	MigrationLifecycleStatesInactive  MigrationLifecycleStatesEnum = "INACTIVE"
	MigrationLifecycleStatesUpdating  MigrationLifecycleStatesEnum = "UPDATING"
	MigrationLifecycleStatesSucceeded MigrationLifecycleStatesEnum = "SUCCEEDED"
	MigrationLifecycleStatesDeleting  MigrationLifecycleStatesEnum = "DELETING"
	MigrationLifecycleStatesDeleted   MigrationLifecycleStatesEnum = "DELETED"
)

var mappingMigrationLifecycleStatesEnum = map[string]MigrationLifecycleStatesEnum{
	"CREATING":  MigrationLifecycleStatesCreating,
	"ACTIVE":    MigrationLifecycleStatesActive,
	"INACTIVE":  MigrationLifecycleStatesInactive,
	"UPDATING":  MigrationLifecycleStatesUpdating,
	"SUCCEEDED": MigrationLifecycleStatesSucceeded,
	"DELETING":  MigrationLifecycleStatesDeleting,
	"DELETED":   MigrationLifecycleStatesDeleted,
}

// GetMigrationLifecycleStatesEnumValues Enumerates the set of values for MigrationLifecycleStatesEnum
func GetMigrationLifecycleStatesEnumValues() []MigrationLifecycleStatesEnum {
	values := make([]MigrationLifecycleStatesEnum, 0)
	for _, v := range mappingMigrationLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetMigrationLifecycleStatesEnumStringValues Enumerates the set of values in String for MigrationLifecycleStatesEnum
func GetMigrationLifecycleStatesEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"SUCCEEDED",
		"DELETING",
		"DELETED",
	}
}
