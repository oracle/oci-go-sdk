// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

type PathMatchType struct {

	// The type of matching for the associated path route.
	MatchType PathMatchTypeMatchTypeEnum `mandatory:"true" json:"matchType,omitempty"`
}

func (model PathMatchType) String() string {
	return common.PointerString(model)
}

type PathMatchTypeMatchTypeEnum string

const (
	PATH_MATCH_TYPE_MATCH_TYPE_EXACT_MATCH                PathMatchTypeMatchTypeEnum = "EXACT_MATCH"
	PATH_MATCH_TYPE_MATCH_TYPE_FORCE_LONGEST_PREFIX_MATCH PathMatchTypeMatchTypeEnum = "FORCE_LONGEST_PREFIX_MATCH"
	PATH_MATCH_TYPE_MATCH_TYPE_PREFIX_MATCH               PathMatchTypeMatchTypeEnum = "PREFIX_MATCH"
	PATH_MATCH_TYPE_MATCH_TYPE_SUFFIX_MATCH               PathMatchTypeMatchTypeEnum = "SUFFIX_MATCH"
	PATH_MATCH_TYPE_MATCH_TYPE_UNKNOWN                    PathMatchTypeMatchTypeEnum = "UNKNOWN"
)

var mapping_pathmatchtype_matchType = map[string]PathMatchTypeMatchTypeEnum{
	"EXACT_MATCH":                PATH_MATCH_TYPE_MATCH_TYPE_EXACT_MATCH,
	"FORCE_LONGEST_PREFIX_MATCH": PATH_MATCH_TYPE_MATCH_TYPE_FORCE_LONGEST_PREFIX_MATCH,
	"PREFIX_MATCH":               PATH_MATCH_TYPE_MATCH_TYPE_PREFIX_MATCH,
	"SUFFIX_MATCH":               PATH_MATCH_TYPE_MATCH_TYPE_SUFFIX_MATCH,
	"UNKNOWN":                    PATH_MATCH_TYPE_MATCH_TYPE_UNKNOWN,
}

func GetPathMatchTypeMatchTypeEnumValues() []PathMatchTypeMatchTypeEnum {
	values := make([]PathMatchTypeMatchTypeEnum, 0)
	for _, v := range mapping_pathmatchtype_matchType {
		if v != PATH_MATCH_TYPE_MATCH_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
