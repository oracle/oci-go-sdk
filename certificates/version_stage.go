// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Retrieval API
//
// API for retrieving certificates.
//

package certificates

import (
	"strings"
)

// VersionStageEnum Enum with underlying type: string
type VersionStageEnum string

// Set of constants representing the allowable values for VersionStageEnum
const (
	VersionStageCurrent           VersionStageEnum = "CURRENT"
	VersionStagePending           VersionStageEnum = "PENDING"
	VersionStagePendingActivation VersionStageEnum = "PENDING_ACTIVATION"
	VersionStageLatest            VersionStageEnum = "LATEST"
	VersionStagePrevious          VersionStageEnum = "PREVIOUS"
	VersionStageDeprecated        VersionStageEnum = "DEPRECATED"
	VersionStageFailed            VersionStageEnum = "FAILED"
)

var mappingVersionStageEnum = map[string]VersionStageEnum{
	"CURRENT":            VersionStageCurrent,
	"PENDING":            VersionStagePending,
	"PENDING_ACTIVATION": VersionStagePendingActivation,
	"LATEST":             VersionStageLatest,
	"PREVIOUS":           VersionStagePrevious,
	"DEPRECATED":         VersionStageDeprecated,
	"FAILED":             VersionStageFailed,
}

var mappingVersionStageEnumLowerCase = map[string]VersionStageEnum{
	"current":            VersionStageCurrent,
	"pending":            VersionStagePending,
	"pending_activation": VersionStagePendingActivation,
	"latest":             VersionStageLatest,
	"previous":           VersionStagePrevious,
	"deprecated":         VersionStageDeprecated,
	"failed":             VersionStageFailed,
}

// GetVersionStageEnumValues Enumerates the set of values for VersionStageEnum
func GetVersionStageEnumValues() []VersionStageEnum {
	values := make([]VersionStageEnum, 0)
	for _, v := range mappingVersionStageEnum {
		values = append(values, v)
	}
	return values
}

// GetVersionStageEnumStringValues Enumerates the set of values in String for VersionStageEnum
func GetVersionStageEnumStringValues() []string {
	return []string{
		"CURRENT",
		"PENDING",
		"PENDING_ACTIVATION",
		"LATEST",
		"PREVIOUS",
		"DEPRECATED",
		"FAILED",
	}
}

// GetMappingVersionStageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVersionStageEnum(val string) (VersionStageEnum, bool) {
	enum, ok := mappingVersionStageEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
