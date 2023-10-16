// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePublisherService API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplacepublisher

import (
	"strings"
)

// ArtifactTypeEnumEnum Enum with underlying type: string
type ArtifactTypeEnumEnum string

// Set of constants representing the allowable values for ArtifactTypeEnumEnum
const (
	ArtifactTypeEnumContainerImage ArtifactTypeEnumEnum = "CONTAINER_IMAGE"
	ArtifactTypeEnumHelmChart      ArtifactTypeEnumEnum = "HELM_CHART"
)

var mappingArtifactTypeEnumEnum = map[string]ArtifactTypeEnumEnum{
	"CONTAINER_IMAGE": ArtifactTypeEnumContainerImage,
	"HELM_CHART":      ArtifactTypeEnumHelmChart,
}

var mappingArtifactTypeEnumEnumLowerCase = map[string]ArtifactTypeEnumEnum{
	"container_image": ArtifactTypeEnumContainerImage,
	"helm_chart":      ArtifactTypeEnumHelmChart,
}

// GetArtifactTypeEnumEnumValues Enumerates the set of values for ArtifactTypeEnumEnum
func GetArtifactTypeEnumEnumValues() []ArtifactTypeEnumEnum {
	values := make([]ArtifactTypeEnumEnum, 0)
	for _, v := range mappingArtifactTypeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetArtifactTypeEnumEnumStringValues Enumerates the set of values in String for ArtifactTypeEnumEnum
func GetArtifactTypeEnumEnumStringValues() []string {
	return []string{
		"CONTAINER_IMAGE",
		"HELM_CHART",
	}
}

// GetMappingArtifactTypeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArtifactTypeEnumEnum(val string) (ArtifactTypeEnumEnum, bool) {
	enum, ok := mappingArtifactTypeEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
