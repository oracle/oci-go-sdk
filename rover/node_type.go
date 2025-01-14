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

// NodeTypeEnum Enum with underlying type: string
type NodeTypeEnum string

// Set of constants representing the allowable values for NodeTypeEnum
const (
	NodeTypeStandalone NodeTypeEnum = "STANDALONE"
	NodeTypeClustered  NodeTypeEnum = "CLUSTERED"
	NodeTypeStation    NodeTypeEnum = "STATION"
)

var mappingNodeTypeEnum = map[string]NodeTypeEnum{
	"STANDALONE": NodeTypeStandalone,
	"CLUSTERED":  NodeTypeClustered,
	"STATION":    NodeTypeStation,
}

var mappingNodeTypeEnumLowerCase = map[string]NodeTypeEnum{
	"standalone": NodeTypeStandalone,
	"clustered":  NodeTypeClustered,
	"station":    NodeTypeStation,
}

// GetNodeTypeEnumValues Enumerates the set of values for NodeTypeEnum
func GetNodeTypeEnumValues() []NodeTypeEnum {
	values := make([]NodeTypeEnum, 0)
	for _, v := range mappingNodeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNodeTypeEnumStringValues Enumerates the set of values in String for NodeTypeEnum
func GetNodeTypeEnumStringValues() []string {
	return []string{
		"STANDALONE",
		"CLUSTERED",
		"STATION",
	}
}

// GetMappingNodeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodeTypeEnum(val string) (NodeTypeEnum, bool) {
	enum, ok := mappingNodeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
