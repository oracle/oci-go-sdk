// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// MetricTypeEnum Enum with underlying type: string
type MetricTypeEnum string

// Set of constants representing the allowable values for MetricTypeEnum
const (
	MetricTypeOcpuHours     MetricTypeEnum = "OCPU_HOURS"
	MetricTypeInstanceHours MetricTypeEnum = "INSTANCE_HOURS"
	MetricTypeCoreHours     MetricTypeEnum = "CORE_HOURS"
	MetricTypeAmount        MetricTypeEnum = "AMOUNT"
	MetricTypeEach          MetricTypeEnum = "EACH"
)

var mappingMetricTypeEnum = map[string]MetricTypeEnum{
	"OCPU_HOURS":     MetricTypeOcpuHours,
	"INSTANCE_HOURS": MetricTypeInstanceHours,
	"CORE_HOURS":     MetricTypeCoreHours,
	"AMOUNT":         MetricTypeAmount,
	"EACH":           MetricTypeEach,
}

var mappingMetricTypeEnumLowerCase = map[string]MetricTypeEnum{
	"ocpu_hours":     MetricTypeOcpuHours,
	"instance_hours": MetricTypeInstanceHours,
	"core_hours":     MetricTypeCoreHours,
	"amount":         MetricTypeAmount,
	"each":           MetricTypeEach,
}

// GetMetricTypeEnumValues Enumerates the set of values for MetricTypeEnum
func GetMetricTypeEnumValues() []MetricTypeEnum {
	values := make([]MetricTypeEnum, 0)
	for _, v := range mappingMetricTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMetricTypeEnumStringValues Enumerates the set of values in String for MetricTypeEnum
func GetMetricTypeEnumStringValues() []string {
	return []string{
		"OCPU_HOURS",
		"INSTANCE_HOURS",
		"CORE_HOURS",
		"AMOUNT",
		"EACH",
	}
}

// GetMappingMetricTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetricTypeEnum(val string) (MetricTypeEnum, bool) {
	enum, ok := mappingMetricTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
