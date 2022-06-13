// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GovernanceRulesControlPlane API
//
// A description of the GovernanceRulesControlPlane API
//

package governancerulescontrolplane

import (
	"strings"
)

// GovernanceRuleLifecycleStateEnum Enum with underlying type: string
type GovernanceRuleLifecycleStateEnum string

// Set of constants representing the allowable values for GovernanceRuleLifecycleStateEnum
const (
	GovernanceRuleLifecycleStateActive  GovernanceRuleLifecycleStateEnum = "ACTIVE"
	GovernanceRuleLifecycleStateDeleted GovernanceRuleLifecycleStateEnum = "DELETED"
)

var mappingGovernanceRuleLifecycleStateEnum = map[string]GovernanceRuleLifecycleStateEnum{
	"ACTIVE":  GovernanceRuleLifecycleStateActive,
	"DELETED": GovernanceRuleLifecycleStateDeleted,
}

var mappingGovernanceRuleLifecycleStateEnumLowerCase = map[string]GovernanceRuleLifecycleStateEnum{
	"active":  GovernanceRuleLifecycleStateActive,
	"deleted": GovernanceRuleLifecycleStateDeleted,
}

// GetGovernanceRuleLifecycleStateEnumValues Enumerates the set of values for GovernanceRuleLifecycleStateEnum
func GetGovernanceRuleLifecycleStateEnumValues() []GovernanceRuleLifecycleStateEnum {
	values := make([]GovernanceRuleLifecycleStateEnum, 0)
	for _, v := range mappingGovernanceRuleLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetGovernanceRuleLifecycleStateEnumStringValues Enumerates the set of values in String for GovernanceRuleLifecycleStateEnum
func GetGovernanceRuleLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingGovernanceRuleLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGovernanceRuleLifecycleStateEnum(val string) (GovernanceRuleLifecycleStateEnum, bool) {
	enum, ok := mappingGovernanceRuleLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
