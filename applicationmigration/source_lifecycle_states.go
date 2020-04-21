// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

// SourceLifecycleStatesEnum Enum with underlying type: string
type SourceLifecycleStatesEnum string

// Set of constants representing the allowable values for SourceLifecycleStatesEnum
const (
	SourceLifecycleStatesCreating SourceLifecycleStatesEnum = "CREATING"
	SourceLifecycleStatesDeleting SourceLifecycleStatesEnum = "DELETING"
	SourceLifecycleStatesUpdating SourceLifecycleStatesEnum = "UPDATING"
	SourceLifecycleStatesActive   SourceLifecycleStatesEnum = "ACTIVE"
	SourceLifecycleStatesInactive SourceLifecycleStatesEnum = "INACTIVE"
	SourceLifecycleStatesDeleted  SourceLifecycleStatesEnum = "DELETED"
)

var mappingSourceLifecycleStates = map[string]SourceLifecycleStatesEnum{
	"CREATING": SourceLifecycleStatesCreating,
	"DELETING": SourceLifecycleStatesDeleting,
	"UPDATING": SourceLifecycleStatesUpdating,
	"ACTIVE":   SourceLifecycleStatesActive,
	"INACTIVE": SourceLifecycleStatesInactive,
	"DELETED":  SourceLifecycleStatesDeleted,
}

// GetSourceLifecycleStatesEnumValues Enumerates the set of values for SourceLifecycleStatesEnum
func GetSourceLifecycleStatesEnumValues() []SourceLifecycleStatesEnum {
	values := make([]SourceLifecycleStatesEnum, 0)
	for _, v := range mappingSourceLifecycleStates {
		values = append(values, v)
	}
	return values
}
