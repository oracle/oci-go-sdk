// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

// ProblemTypeEnum Enum with underlying type: string
type ProblemTypeEnum string

// Set of constants representing the allowable values for ProblemTypeEnum
const (
	ProblemTypeLimit       ProblemTypeEnum = "LIMIT"
	ProblemTypeLegacyLimit ProblemTypeEnum = "LEGACY_LIMIT"
	ProblemTypeTech        ProblemTypeEnum = "TECH"
	ProblemTypeAccount     ProblemTypeEnum = "ACCOUNT"
)

var mappingProblemType = map[string]ProblemTypeEnum{
	"LIMIT":        ProblemTypeLimit,
	"LEGACY_LIMIT": ProblemTypeLegacyLimit,
	"TECH":         ProblemTypeTech,
	"ACCOUNT":      ProblemTypeAccount,
}

// GetProblemTypeEnumValues Enumerates the set of values for ProblemTypeEnum
func GetProblemTypeEnumValues() []ProblemTypeEnum {
	values := make([]ProblemTypeEnum, 0)
	for _, v := range mappingProblemType {
		values = append(values, v)
	}
	return values
}
