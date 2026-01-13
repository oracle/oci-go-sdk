// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Access Governance APIs
//
// Use the Oracle Access Governance API to create, view, and manage GovernanceInstances.
//

package accessgovernancecp

import (
	"strings"
)

// InstanceLifecycleStateEnum Enum with underlying type: string
type InstanceLifecycleStateEnum string

// Set of constants representing the allowable values for InstanceLifecycleStateEnum
const (
	InstanceLifecycleStateCreating       InstanceLifecycleStateEnum = "CREATING"
	InstanceLifecycleStateActive         InstanceLifecycleStateEnum = "ACTIVE"
	InstanceLifecycleStateDeleting       InstanceLifecycleStateEnum = "DELETING"
	InstanceLifecycleStateDeleted        InstanceLifecycleStateEnum = "DELETED"
	InstanceLifecycleStateNeedsAttention InstanceLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingInstanceLifecycleStateEnum = map[string]InstanceLifecycleStateEnum{
	"CREATING":        InstanceLifecycleStateCreating,
	"ACTIVE":          InstanceLifecycleStateActive,
	"DELETING":        InstanceLifecycleStateDeleting,
	"DELETED":         InstanceLifecycleStateDeleted,
	"NEEDS_ATTENTION": InstanceLifecycleStateNeedsAttention,
}

var mappingInstanceLifecycleStateEnumLowerCase = map[string]InstanceLifecycleStateEnum{
	"creating":        InstanceLifecycleStateCreating,
	"active":          InstanceLifecycleStateActive,
	"deleting":        InstanceLifecycleStateDeleting,
	"deleted":         InstanceLifecycleStateDeleted,
	"needs_attention": InstanceLifecycleStateNeedsAttention,
}

// GetInstanceLifecycleStateEnumValues Enumerates the set of values for InstanceLifecycleStateEnum
func GetInstanceLifecycleStateEnumValues() []InstanceLifecycleStateEnum {
	values := make([]InstanceLifecycleStateEnum, 0)
	for _, v := range mappingInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for InstanceLifecycleStateEnum
func GetInstanceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceLifecycleStateEnum(val string) (InstanceLifecycleStateEnum, bool) {
	enum, ok := mappingInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
