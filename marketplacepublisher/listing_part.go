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

// ListingPart A listing SKUs and meter information attached by marketplace admin.
type ListingPart struct {

	// The SKU/part name.
	Sku *string `mandatory:"true" json:"sku"`

	// The part's metric.
	MetricType MetricTypeEnum `mandatory:"true" json:"metricType"`

	// rate allocation, these are calculated based on rate information at listing revision.
	RateAllocation *float32 `mandatory:"true" json:"rateAllocation"`

	// Identifies whether part has Gov SKU.
	HasGovSku *bool `mandatory:"true" json:"hasGovSku"`

	// List of meters associated with the part.
	Meters []ListingMeter `mandatory:"true" json:"meters"`

	// Unique identifier of the pricing plan.
	PricingPlanKey *string `mandatory:"false" json:"pricingPlanKey"`

	// The billing model for SaaS paid listing parts.
	BillingModel ListingPartBillingModelEnum `mandatory:"false" json:"billingModel,omitempty"`
}

func (m ListingPart) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ListingPart) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMetricTypeEnum(string(m.MetricType)); !ok && m.MetricType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricType: %s. Supported values are: %s.", m.MetricType, strings.Join(GetMetricTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingListingPartBillingModelEnum(string(m.BillingModel)); !ok && m.BillingModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BillingModel: %s. Supported values are: %s.", m.BillingModel, strings.Join(GetListingPartBillingModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListingPartBillingModelEnum Enum with underlying type: string
type ListingPartBillingModelEnum string

// Set of constants representing the allowable values for ListingPartBillingModelEnum
const (
	ListingPartBillingModelFlatRate   ListingPartBillingModelEnum = "FLAT_RATE"
	ListingPartBillingModelUsageBased ListingPartBillingModelEnum = "USAGE_BASED"
)

var mappingListingPartBillingModelEnum = map[string]ListingPartBillingModelEnum{
	"FLAT_RATE":   ListingPartBillingModelFlatRate,
	"USAGE_BASED": ListingPartBillingModelUsageBased,
}

var mappingListingPartBillingModelEnumLowerCase = map[string]ListingPartBillingModelEnum{
	"flat_rate":   ListingPartBillingModelFlatRate,
	"usage_based": ListingPartBillingModelUsageBased,
}

// GetListingPartBillingModelEnumValues Enumerates the set of values for ListingPartBillingModelEnum
func GetListingPartBillingModelEnumValues() []ListingPartBillingModelEnum {
	values := make([]ListingPartBillingModelEnum, 0)
	for _, v := range mappingListingPartBillingModelEnum {
		values = append(values, v)
	}
	return values
}

// GetListingPartBillingModelEnumStringValues Enumerates the set of values in String for ListingPartBillingModelEnum
func GetListingPartBillingModelEnumStringValues() []string {
	return []string{
		"FLAT_RATE",
		"USAGE_BASED",
	}
}

// GetMappingListingPartBillingModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingPartBillingModelEnum(val string) (ListingPartBillingModelEnum, bool) {
	enum, ok := mappingListingPartBillingModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
