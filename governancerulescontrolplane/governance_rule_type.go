// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// GovernanceRuleTypeEnum Enum with underlying type: string
type GovernanceRuleTypeEnum string

// Set of constants representing the allowable values for GovernanceRuleTypeEnum
const (
	GovernanceRuleTypeQuota          GovernanceRuleTypeEnum = "QUOTA"
	GovernanceRuleTypeTag            GovernanceRuleTypeEnum = "TAG"
	GovernanceRuleTypeAllowedRegions GovernanceRuleTypeEnum = "ALLOWED_REGIONS"
)

var mappingGovernanceRuleTypeEnum = map[string]GovernanceRuleTypeEnum{
	"QUOTA":           GovernanceRuleTypeQuota,
	"TAG":             GovernanceRuleTypeTag,
	"ALLOWED_REGIONS": GovernanceRuleTypeAllowedRegions,
}

var mappingGovernanceRuleTypeEnumLowerCase = map[string]GovernanceRuleTypeEnum{
	"quota":           GovernanceRuleTypeQuota,
	"tag":             GovernanceRuleTypeTag,
	"allowed_regions": GovernanceRuleTypeAllowedRegions,
}

// GetGovernanceRuleTypeEnumValues Enumerates the set of values for GovernanceRuleTypeEnum
func GetGovernanceRuleTypeEnumValues() []GovernanceRuleTypeEnum {
	values := make([]GovernanceRuleTypeEnum, 0)
	for _, v := range mappingGovernanceRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGovernanceRuleTypeEnumStringValues Enumerates the set of values in String for GovernanceRuleTypeEnum
func GetGovernanceRuleTypeEnumStringValues() []string {
	return []string{
		"QUOTA",
		"TAG",
		"ALLOWED_REGIONS",
	}
}

// GetMappingGovernanceRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGovernanceRuleTypeEnum(val string) (GovernanceRuleTypeEnum, bool) {
	enum, ok := mappingGovernanceRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
