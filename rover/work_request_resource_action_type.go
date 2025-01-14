// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"strings"
)

// WorkRequestResourceActionTypeEnum Enum with underlying type: string
type WorkRequestResourceActionTypeEnum string

// Set of constants representing the allowable values for WorkRequestResourceActionTypeEnum
const (
	WorkRequestResourceActionTypeInProgress WorkRequestResourceActionTypeEnum = "IN_PROGRESS"
	WorkRequestResourceActionTypeFailed     WorkRequestResourceActionTypeEnum = "FAILED"
	WorkRequestResourceActionTypeCreated    WorkRequestResourceActionTypeEnum = "CREATED"
	WorkRequestResourceActionTypeUpdated    WorkRequestResourceActionTypeEnum = "UPDATED"
)

var mappingWorkRequestResourceActionTypeEnum = map[string]WorkRequestResourceActionTypeEnum{
	"IN_PROGRESS": WorkRequestResourceActionTypeInProgress,
	"FAILED":      WorkRequestResourceActionTypeFailed,
	"CREATED":     WorkRequestResourceActionTypeCreated,
	"UPDATED":     WorkRequestResourceActionTypeUpdated,
}

var mappingWorkRequestResourceActionTypeEnumLowerCase = map[string]WorkRequestResourceActionTypeEnum{
	"in_progress": WorkRequestResourceActionTypeInProgress,
	"failed":      WorkRequestResourceActionTypeFailed,
	"created":     WorkRequestResourceActionTypeCreated,
	"updated":     WorkRequestResourceActionTypeUpdated,
}

// GetWorkRequestResourceActionTypeEnumValues Enumerates the set of values for WorkRequestResourceActionTypeEnum
func GetWorkRequestResourceActionTypeEnumValues() []WorkRequestResourceActionTypeEnum {
	values := make([]WorkRequestResourceActionTypeEnum, 0)
	for _, v := range mappingWorkRequestResourceActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestResourceActionTypeEnumStringValues Enumerates the set of values in String for WorkRequestResourceActionTypeEnum
func GetWorkRequestResourceActionTypeEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"FAILED",
		"CREATED",
		"UPDATED",
	}
}

// GetMappingWorkRequestResourceActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestResourceActionTypeEnum(val string) (WorkRequestResourceActionTypeEnum, bool) {
	enum, ok := mappingWorkRequestResourceActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
