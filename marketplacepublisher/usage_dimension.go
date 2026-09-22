// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePublisherService API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UsageDimension Usage dimension for usage-based or hybrid SaaS pricing.
type UsageDimension struct {

	// The dimension name.
	DimensionName *string `mandatory:"true" json:"dimensionName"`

	// The dimension description.
	DimensionDescription *string `mandatory:"true" json:"dimensionDescription"`

	// The metric type used to measure usage for the dimension.
	MetricType MetricTypeEnum `mandatory:"true" json:"metricType"`

	// Billing frequency used for the dimension.
	DimensionBillingFrequency UsageDimensionDimensionBillingFrequencyEnum `mandatory:"true" json:"dimensionBillingFrequency"`

	// List of pricing rates provider by publisher for the dimension.
	Rates []PricingRate `mandatory:"true" json:"rates"`

	// Server-generated unique identifier for the usage dimension. If omitted in the request, it is generated automatically.
	DimensionKey *string `mandatory:"false" json:"dimensionKey"`

	// Included quantity before overage rates apply.
	IncludedQuantity *float32 `mandatory:"false" json:"includedQuantity"`

	// Additional metadata key/value pairs for the dimension.
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`
}

func (m UsageDimension) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UsageDimension) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMetricTypeEnum(string(m.MetricType)); !ok && m.MetricType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricType: %s. Supported values are: %s.", m.MetricType, strings.Join(GetMetricTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUsageDimensionDimensionBillingFrequencyEnum(string(m.DimensionBillingFrequency)); !ok && m.DimensionBillingFrequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DimensionBillingFrequency: %s. Supported values are: %s.", m.DimensionBillingFrequency, strings.Join(GetUsageDimensionDimensionBillingFrequencyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UsageDimensionDimensionBillingFrequencyEnum Enum with underlying type: string
type UsageDimensionDimensionBillingFrequencyEnum string

// Set of constants representing the allowable values for UsageDimensionDimensionBillingFrequencyEnum
const (
	UsageDimensionDimensionBillingFrequencyMonthly    UsageDimensionDimensionBillingFrequencyEnum = "MONTHLY"
	UsageDimensionDimensionBillingFrequencyQuarterly  UsageDimensionDimensionBillingFrequencyEnum = "QUARTERLY"
	UsageDimensionDimensionBillingFrequencySemiAnnual UsageDimensionDimensionBillingFrequencyEnum = "SEMI_ANNUAL"
	UsageDimensionDimensionBillingFrequencyAnnual     UsageDimensionDimensionBillingFrequencyEnum = "ANNUAL"
	UsageDimensionDimensionBillingFrequencyBiennial   UsageDimensionDimensionBillingFrequencyEnum = "BIENNIAL"
	UsageDimensionDimensionBillingFrequencyTriennial  UsageDimensionDimensionBillingFrequencyEnum = "TRIENNIAL"
)

var mappingUsageDimensionDimensionBillingFrequencyEnum = map[string]UsageDimensionDimensionBillingFrequencyEnum{
	"MONTHLY":     UsageDimensionDimensionBillingFrequencyMonthly,
	"QUARTERLY":   UsageDimensionDimensionBillingFrequencyQuarterly,
	"SEMI_ANNUAL": UsageDimensionDimensionBillingFrequencySemiAnnual,
	"ANNUAL":      UsageDimensionDimensionBillingFrequencyAnnual,
	"BIENNIAL":    UsageDimensionDimensionBillingFrequencyBiennial,
	"TRIENNIAL":   UsageDimensionDimensionBillingFrequencyTriennial,
}

var mappingUsageDimensionDimensionBillingFrequencyEnumLowerCase = map[string]UsageDimensionDimensionBillingFrequencyEnum{
	"monthly":     UsageDimensionDimensionBillingFrequencyMonthly,
	"quarterly":   UsageDimensionDimensionBillingFrequencyQuarterly,
	"semi_annual": UsageDimensionDimensionBillingFrequencySemiAnnual,
	"annual":      UsageDimensionDimensionBillingFrequencyAnnual,
	"biennial":    UsageDimensionDimensionBillingFrequencyBiennial,
	"triennial":   UsageDimensionDimensionBillingFrequencyTriennial,
}

// GetUsageDimensionDimensionBillingFrequencyEnumValues Enumerates the set of values for UsageDimensionDimensionBillingFrequencyEnum
func GetUsageDimensionDimensionBillingFrequencyEnumValues() []UsageDimensionDimensionBillingFrequencyEnum {
	values := make([]UsageDimensionDimensionBillingFrequencyEnum, 0)
	for _, v := range mappingUsageDimensionDimensionBillingFrequencyEnum {
		values = append(values, v)
	}
	return values
}

// GetUsageDimensionDimensionBillingFrequencyEnumStringValues Enumerates the set of values in String for UsageDimensionDimensionBillingFrequencyEnum
func GetUsageDimensionDimensionBillingFrequencyEnumStringValues() []string {
	return []string{
		"MONTHLY",
		"QUARTERLY",
		"SEMI_ANNUAL",
		"ANNUAL",
		"BIENNIAL",
		"TRIENNIAL",
	}
}

// GetMappingUsageDimensionDimensionBillingFrequencyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUsageDimensionDimensionBillingFrequencyEnum(val string) (UsageDimensionDimensionBillingFrequencyEnum, bool) {
	enum, ok := mappingUsageDimensionDimensionBillingFrequencyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
