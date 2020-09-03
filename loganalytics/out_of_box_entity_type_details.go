// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
)

// OutOfBoxEntityTypeDetails A Single Entity Type Definition
type OutOfBoxEntityTypeDetails struct {

	// Log analytics entity type name.
	Name *string `mandatory:"true" json:"name"`

	// Internal name for the log analytics entity type.
	InternalName *string `mandatory:"true" json:"internalName"`

	// Log analytics entity type category. Category will be used for grouping and filtering.
	Category *string `mandatory:"true" json:"category"`

	// Log analytics entity type group. Supported values: ClOUD, NON_CLOUD.
	CloudType OutOfBoxEntityTypeDetailsCloudTypeEnum `mandatory:"true" json:"cloudType"`

	// A Single Entity Type Property Definition
	Properties []EntityTypeProperty `mandatory:"false" json:"properties"`
}

func (m OutOfBoxEntityTypeDetails) String() string {
	return common.PointerString(m)
}

// OutOfBoxEntityTypeDetailsCloudTypeEnum Enum with underlying type: string
type OutOfBoxEntityTypeDetailsCloudTypeEnum string

// Set of constants representing the allowable values for OutOfBoxEntityTypeDetailsCloudTypeEnum
const (
	OutOfBoxEntityTypeDetailsCloudTypeCloud    OutOfBoxEntityTypeDetailsCloudTypeEnum = "CLOUD"
	OutOfBoxEntityTypeDetailsCloudTypeNonCloud OutOfBoxEntityTypeDetailsCloudTypeEnum = "NON_CLOUD"
)

var mappingOutOfBoxEntityTypeDetailsCloudType = map[string]OutOfBoxEntityTypeDetailsCloudTypeEnum{
	"CLOUD":     OutOfBoxEntityTypeDetailsCloudTypeCloud,
	"NON_CLOUD": OutOfBoxEntityTypeDetailsCloudTypeNonCloud,
}

// GetOutOfBoxEntityTypeDetailsCloudTypeEnumValues Enumerates the set of values for OutOfBoxEntityTypeDetailsCloudTypeEnum
func GetOutOfBoxEntityTypeDetailsCloudTypeEnumValues() []OutOfBoxEntityTypeDetailsCloudTypeEnum {
	values := make([]OutOfBoxEntityTypeDetailsCloudTypeEnum, 0)
	for _, v := range mappingOutOfBoxEntityTypeDetailsCloudType {
		values = append(values, v)
	}
	return values
}
