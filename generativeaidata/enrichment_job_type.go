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

// EnrichmentJobTypeEnum Enum with underlying type: string
type EnrichmentJobTypeEnum string

// Set of constants representing the allowable values for EnrichmentJobTypeEnum
const (
	EnrichmentJobTypeFullBuild    EnrichmentJobTypeEnum = "FULL_BUILD"
	EnrichmentJobTypePartialBuild EnrichmentJobTypeEnum = "PARTIAL_BUILD"
	EnrichmentJobTypeDeltaRefresh EnrichmentJobTypeEnum = "DELTA_REFRESH"
)

var mappingEnrichmentJobTypeEnum = map[string]EnrichmentJobTypeEnum{
	"FULL_BUILD":    EnrichmentJobTypeFullBuild,
	"PARTIAL_BUILD": EnrichmentJobTypePartialBuild,
	"DELTA_REFRESH": EnrichmentJobTypeDeltaRefresh,
}

var mappingEnrichmentJobTypeEnumLowerCase = map[string]EnrichmentJobTypeEnum{
	"full_build":    EnrichmentJobTypeFullBuild,
	"partial_build": EnrichmentJobTypePartialBuild,
	"delta_refresh": EnrichmentJobTypeDeltaRefresh,
}

// GetEnrichmentJobTypeEnumValues Enumerates the set of values for EnrichmentJobTypeEnum
func GetEnrichmentJobTypeEnumValues() []EnrichmentJobTypeEnum {
	values := make([]EnrichmentJobTypeEnum, 0)
	for _, v := range mappingEnrichmentJobTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEnrichmentJobTypeEnumStringValues Enumerates the set of values in String for EnrichmentJobTypeEnum
func GetEnrichmentJobTypeEnumStringValues() []string {
	return []string{
		"FULL_BUILD",
		"PARTIAL_BUILD",
		"DELTA_REFRESH",
	}
}

// GetMappingEnrichmentJobTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnrichmentJobTypeEnum(val string) (EnrichmentJobTypeEnum, bool) {
	enum, ok := mappingEnrichmentJobTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
