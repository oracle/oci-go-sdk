// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service NL2SQL API
//
// A description of the ReferenceService API. in progress
//

package generativeaidata

import (
	"strings"
)

// CompletionModeEnum Enum with underlying type: string
type CompletionModeEnum string

// Set of constants representing the allowable values for CompletionModeEnum
const (
	CompletionModeWaitForCompletion CompletionModeEnum = "WAIT_FOR_COMPLETION"
	CompletionModeBackgroundJob     CompletionModeEnum = "BACKGROUND_JOB"
)

var mappingCompletionModeEnum = map[string]CompletionModeEnum{
	"WAIT_FOR_COMPLETION": CompletionModeWaitForCompletion,
	"BACKGROUND_JOB":      CompletionModeBackgroundJob,
}

var mappingCompletionModeEnumLowerCase = map[string]CompletionModeEnum{
	"wait_for_completion": CompletionModeWaitForCompletion,
	"background_job":      CompletionModeBackgroundJob,
}

// GetCompletionModeEnumValues Enumerates the set of values for CompletionModeEnum
func GetCompletionModeEnumValues() []CompletionModeEnum {
	values := make([]CompletionModeEnum, 0)
	for _, v := range mappingCompletionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCompletionModeEnumStringValues Enumerates the set of values in String for CompletionModeEnum
func GetCompletionModeEnumStringValues() []string {
	return []string{
		"WAIT_FOR_COMPLETION",
		"BACKGROUND_JOB",
	}
}

// GetMappingCompletionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCompletionModeEnum(val string) (CompletionModeEnum, bool) {
	enum, ok := mappingCompletionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
