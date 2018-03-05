// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PathMatchType The representation of PathMatchType
type PathMatchType struct {

	// The type of matching for the associated path route.
	MatchType PathMatchTypeMatchTypeEnum `mandatory:"true" json:"matchType"`
}

func (m PathMatchType) String() string {
	return common.PointerString(m)
}

// PathMatchTypeMatchTypeEnum Enum with underlying type: string
type PathMatchTypeMatchTypeEnum string

// Set of constants representing the allowable values for PathMatchTypeMatchType
const (
	PathMatchTypeMatchTypeExactMatch              PathMatchTypeMatchTypeEnum = "EXACT_MATCH"
	PathMatchTypeMatchTypeForceLongestPrefixMatch PathMatchTypeMatchTypeEnum = "FORCE_LONGEST_PREFIX_MATCH"
	PathMatchTypeMatchTypePrefixMatch             PathMatchTypeMatchTypeEnum = "PREFIX_MATCH"
	PathMatchTypeMatchTypeSuffixMatch             PathMatchTypeMatchTypeEnum = "SUFFIX_MATCH"
)

var mappingPathMatchTypeMatchType = map[string]PathMatchTypeMatchTypeEnum{
	"EXACT_MATCH":                PathMatchTypeMatchTypeExactMatch,
	"FORCE_LONGEST_PREFIX_MATCH": PathMatchTypeMatchTypeForceLongestPrefixMatch,
	"PREFIX_MATCH":               PathMatchTypeMatchTypePrefixMatch,
	"SUFFIX_MATCH":               PathMatchTypeMatchTypeSuffixMatch,
}

// GetPathMatchTypeMatchTypeEnumValues Enumerates the set of values for PathMatchTypeMatchType
func GetPathMatchTypeMatchTypeEnumValues() []PathMatchTypeMatchTypeEnum {
	values := make([]PathMatchTypeMatchTypeEnum, 0)
	for _, v := range mappingPathMatchTypeMatchType {
		values = append(values, v)
	}
	return values
}
