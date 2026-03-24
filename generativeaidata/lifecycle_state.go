// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Nl2sql API
//
// A description of the ReferenceService API. in progress
//

package generativeaidata

import (
	"strings"
)

// LifecycleStateEnum Enum with underlying type: string
type LifecycleStateEnum string

// Set of constants representing the allowable values for LifecycleStateEnum
const (
	LifecycleStateAccepted   LifecycleStateEnum = "ACCEPTED"
	LifecycleStateInProgress LifecycleStateEnum = "IN_PROGRESS"
	LifecycleStateFailed     LifecycleStateEnum = "FAILED"
	LifecycleStateSucceeded  LifecycleStateEnum = "SUCCEEDED"
	LifecycleStateCanceling  LifecycleStateEnum = "CANCELING"
	LifecycleStateCanceled   LifecycleStateEnum = "CANCELED"
)

var mappingLifecycleStateEnum = map[string]LifecycleStateEnum{
	"ACCEPTED":    LifecycleStateAccepted,
	"IN_PROGRESS": LifecycleStateInProgress,
	"FAILED":      LifecycleStateFailed,
	"SUCCEEDED":   LifecycleStateSucceeded,
	"CANCELING":   LifecycleStateCanceling,
	"CANCELED":    LifecycleStateCanceled,
}

var mappingLifecycleStateEnumLowerCase = map[string]LifecycleStateEnum{
	"accepted":    LifecycleStateAccepted,
	"in_progress": LifecycleStateInProgress,
	"failed":      LifecycleStateFailed,
	"succeeded":   LifecycleStateSucceeded,
	"canceling":   LifecycleStateCanceling,
	"canceled":    LifecycleStateCanceled,
}

// GetLifecycleStateEnumValues Enumerates the set of values for LifecycleStateEnum
func GetLifecycleStateEnumValues() []LifecycleStateEnum {
	values := make([]LifecycleStateEnum, 0)
	for _, v := range mappingLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLifecycleStateEnumStringValues Enumerates the set of values in String for LifecycleStateEnum
func GetLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLifecycleStateEnum(val string) (LifecycleStateEnum, bool) {
	enum, ok := mappingLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
