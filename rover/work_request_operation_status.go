// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// WorkRequestOperationStatusEnum Enum with underlying type: string
type WorkRequestOperationStatusEnum string

// Set of constants representing the allowable values for WorkRequestOperationStatusEnum
const (
	WorkRequestOperationStatusAccepted       WorkRequestOperationStatusEnum = "ACCEPTED"
	WorkRequestOperationStatusInProgress     WorkRequestOperationStatusEnum = "IN_PROGRESS"
	WorkRequestOperationStatusWaiting        WorkRequestOperationStatusEnum = "WAITING"
	WorkRequestOperationStatusNeedsAttention WorkRequestOperationStatusEnum = "NEEDS_ATTENTION"
	WorkRequestOperationStatusFailed         WorkRequestOperationStatusEnum = "FAILED"
	WorkRequestOperationStatusSucceeded      WorkRequestOperationStatusEnum = "SUCCEEDED"
	WorkRequestOperationStatusCanceling      WorkRequestOperationStatusEnum = "CANCELING"
	WorkRequestOperationStatusCanceled       WorkRequestOperationStatusEnum = "CANCELED"
)

var mappingWorkRequestOperationStatusEnum = map[string]WorkRequestOperationStatusEnum{
	"ACCEPTED":        WorkRequestOperationStatusAccepted,
	"IN_PROGRESS":     WorkRequestOperationStatusInProgress,
	"WAITING":         WorkRequestOperationStatusWaiting,
	"NEEDS_ATTENTION": WorkRequestOperationStatusNeedsAttention,
	"FAILED":          WorkRequestOperationStatusFailed,
	"SUCCEEDED":       WorkRequestOperationStatusSucceeded,
	"CANCELING":       WorkRequestOperationStatusCanceling,
	"CANCELED":        WorkRequestOperationStatusCanceled,
}

var mappingWorkRequestOperationStatusEnumLowerCase = map[string]WorkRequestOperationStatusEnum{
	"accepted":        WorkRequestOperationStatusAccepted,
	"in_progress":     WorkRequestOperationStatusInProgress,
	"waiting":         WorkRequestOperationStatusWaiting,
	"needs_attention": WorkRequestOperationStatusNeedsAttention,
	"failed":          WorkRequestOperationStatusFailed,
	"succeeded":       WorkRequestOperationStatusSucceeded,
	"canceling":       WorkRequestOperationStatusCanceling,
	"canceled":        WorkRequestOperationStatusCanceled,
}

// GetWorkRequestOperationStatusEnumValues Enumerates the set of values for WorkRequestOperationStatusEnum
func GetWorkRequestOperationStatusEnumValues() []WorkRequestOperationStatusEnum {
	values := make([]WorkRequestOperationStatusEnum, 0)
	for _, v := range mappingWorkRequestOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestOperationStatusEnumStringValues Enumerates the set of values in String for WorkRequestOperationStatusEnum
func GetWorkRequestOperationStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"NEEDS_ATTENTION",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingWorkRequestOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestOperationStatusEnum(val string) (WorkRequestOperationStatusEnum, bool) {
	enum, ok := mappingWorkRequestOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
