// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"strings"
)

// SortOrderEnum Enum with underlying type: string
type SortOrderEnum string

// Set of constants representing the allowable values for SortOrderEnum
const (
	SortOrderAsc  SortOrderEnum = "ASC"
	SortOrderDesc SortOrderEnum = "DESC"
)

var mappingSortOrderEnum = map[string]SortOrderEnum{
	"ASC":  SortOrderAsc,
	"DESC": SortOrderDesc,
}

var mappingSortOrderEnumLowerCase = map[string]SortOrderEnum{
	"asc":  SortOrderAsc,
	"desc": SortOrderDesc,
}

// GetSortOrderEnumValues Enumerates the set of values for SortOrderEnum
func GetSortOrderEnumValues() []SortOrderEnum {
	values := make([]SortOrderEnum, 0)
	for _, v := range mappingSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSortOrderEnumStringValues Enumerates the set of values in String for SortOrderEnum
func GetSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSortOrderEnum(val string) (SortOrderEnum, bool) {
	enum, ok := mappingSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
