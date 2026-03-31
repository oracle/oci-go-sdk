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

// JobOutputLocationEnum Enum with underlying type: string
type JobOutputLocationEnum string

// Set of constants representing the allowable values for JobOutputLocationEnum
const (
	JobOutputLocationInline JobOutputLocationEnum = "INLINE"
)

var mappingJobOutputLocationEnum = map[string]JobOutputLocationEnum{
	"INLINE": JobOutputLocationInline,
}

var mappingJobOutputLocationEnumLowerCase = map[string]JobOutputLocationEnum{
	"inline": JobOutputLocationInline,
}

// GetJobOutputLocationEnumValues Enumerates the set of values for JobOutputLocationEnum
func GetJobOutputLocationEnumValues() []JobOutputLocationEnum {
	values := make([]JobOutputLocationEnum, 0)
	for _, v := range mappingJobOutputLocationEnum {
		values = append(values, v)
	}
	return values
}

// GetJobOutputLocationEnumStringValues Enumerates the set of values in String for JobOutputLocationEnum
func GetJobOutputLocationEnumStringValues() []string {
	return []string{
		"INLINE",
	}
}

// GetMappingJobOutputLocationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobOutputLocationEnum(val string) (JobOutputLocationEnum, bool) {
	enum, ok := mappingJobOutputLocationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
