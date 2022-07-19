// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.cloud.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams Media Streams (https://docs.cloud.oracle.com/iaas/Content/dms-mediastream/home.htm).
// Use the table of contents and search tool to explore the Media Flow API and Media Streams API.
//

package mediaservices

import (
	"strings"
)

// MediaWorkflowJobFactSortByEnum Enum with underlying type: string
type MediaWorkflowJobFactSortByEnum string

// Set of constants representing the allowable values for MediaWorkflowJobFactSortByEnum
const (
	MediaWorkflowJobFactSortByKey MediaWorkflowJobFactSortByEnum = "key"
)

var mappingMediaWorkflowJobFactSortByEnum = map[string]MediaWorkflowJobFactSortByEnum{
	"key": MediaWorkflowJobFactSortByKey,
}

var mappingMediaWorkflowJobFactSortByEnumLowerCase = map[string]MediaWorkflowJobFactSortByEnum{
	"key": MediaWorkflowJobFactSortByKey,
}

// GetMediaWorkflowJobFactSortByEnumValues Enumerates the set of values for MediaWorkflowJobFactSortByEnum
func GetMediaWorkflowJobFactSortByEnumValues() []MediaWorkflowJobFactSortByEnum {
	values := make([]MediaWorkflowJobFactSortByEnum, 0)
	for _, v := range mappingMediaWorkflowJobFactSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetMediaWorkflowJobFactSortByEnumStringValues Enumerates the set of values in String for MediaWorkflowJobFactSortByEnum
func GetMediaWorkflowJobFactSortByEnumStringValues() []string {
	return []string{
		"key",
	}
}

// GetMappingMediaWorkflowJobFactSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMediaWorkflowJobFactSortByEnum(val string) (MediaWorkflowJobFactSortByEnum, bool) {
	enum, ok := mappingMediaWorkflowJobFactSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}