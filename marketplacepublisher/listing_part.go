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
	MetricType ListingPartMetricTypeEnum `mandatory:"true" json:"metricType"`

	// rate allocation, these are calculated based on rate information at listing revision.
	RateAllocation *float32 `mandatory:"true" json:"rateAllocation"`

	// Identifies whether part has Gov SKU.
	HasGovSku *bool `mandatory:"true" json:"hasGovSku"`

	// List of meters associated with the part.
	Meters []ListingMeter `mandatory:"true" json:"meters"`
}

func (m ListingPart) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ListingPart) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListingPartMetricTypeEnum(string(m.MetricType)); !ok && m.MetricType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricType: %s. Supported values are: %s.", m.MetricType, strings.Join(GetListingPartMetricTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListingPartMetricTypeEnum Enum with underlying type: string
type ListingPartMetricTypeEnum string

// Set of constants representing the allowable values for ListingPartMetricTypeEnum
const (
	ListingPartMetricTypeOcpuHours     ListingPartMetricTypeEnum = "OCPU_HOURS"
	ListingPartMetricTypeInstanceHours ListingPartMetricTypeEnum = "INSTANCE_HOURS"
	ListingPartMetricTypeCoreHours     ListingPartMetricTypeEnum = "CORE_HOURS"
)

var mappingListingPartMetricTypeEnum = map[string]ListingPartMetricTypeEnum{
	"OCPU_HOURS":     ListingPartMetricTypeOcpuHours,
	"INSTANCE_HOURS": ListingPartMetricTypeInstanceHours,
	"CORE_HOURS":     ListingPartMetricTypeCoreHours,
}

var mappingListingPartMetricTypeEnumLowerCase = map[string]ListingPartMetricTypeEnum{
	"ocpu_hours":     ListingPartMetricTypeOcpuHours,
	"instance_hours": ListingPartMetricTypeInstanceHours,
	"core_hours":     ListingPartMetricTypeCoreHours,
}

// GetListingPartMetricTypeEnumValues Enumerates the set of values for ListingPartMetricTypeEnum
func GetListingPartMetricTypeEnumValues() []ListingPartMetricTypeEnum {
	values := make([]ListingPartMetricTypeEnum, 0)
	for _, v := range mappingListingPartMetricTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListingPartMetricTypeEnumStringValues Enumerates the set of values in String for ListingPartMetricTypeEnum
func GetListingPartMetricTypeEnumStringValues() []string {
	return []string{
		"OCPU_HOURS",
		"INSTANCE_HOURS",
		"CORE_HOURS",
	}
}

// GetMappingListingPartMetricTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingPartMetricTypeEnum(val string) (ListingPartMetricTypeEnum, bool) {
	enum, ok := mappingListingPartMetricTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
