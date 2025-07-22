// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePublisherService API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplacepublisher

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MeteredPricingPlan Metered pricing plan.
type MeteredPricingPlan struct {

	// List of pricing rates provider by publisher.
	Rates []PricingRate `mandatory:"true" json:"rates"`

	// Additional metadata key/value pairs for the metering pricing.
	ExtendedMetadata map[string]string `mandatory:"false" json:"extendedMetadata"`

	// The listing's pricing plan name.
	Name MeteredPricingPlanNameEnum `mandatory:"true" json:"name"`
}

// GetRates returns Rates
func (m MeteredPricingPlan) GetRates() []PricingRate {
	return m.Rates
}

func (m MeteredPricingPlan) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MeteredPricingPlan) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMeteredPricingPlanNameEnum(string(m.Name)); !ok && m.Name != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Name: %s. Supported values are: %s.", m.Name, strings.Join(GetMeteredPricingPlanNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MeteredPricingPlan) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMeteredPricingPlan MeteredPricingPlan
	s := struct {
		DiscriminatorParam string `json:"planType"`
		MarshalTypeMeteredPricingPlan
	}{
		"METERED",
		(MarshalTypeMeteredPricingPlan)(m),
	}

	return json.Marshal(&s)
}

// MeteredPricingPlanNameEnum Enum with underlying type: string
type MeteredPricingPlanNameEnum string

// Set of constants representing the allowable values for MeteredPricingPlanNameEnum
const (
	MeteredPricingPlanNameInstance                 MeteredPricingPlanNameEnum = "PER_INSTANCE"
	MeteredPricingPlanNameOcpuMinBillingHrs        MeteredPricingPlanNameEnum = "PER_OCPU_MIN_BILLING_HRS"
	MeteredPricingPlanNameOcpuLinear               MeteredPricingPlanNameEnum = "PER_OCPU_LINEAR"
	MeteredPricingPlanNameInstanceMonthlyInclusive MeteredPricingPlanNameEnum = "PER_INSTANCE_MONTHLY_INCLUSIVE"
)

var mappingMeteredPricingPlanNameEnum = map[string]MeteredPricingPlanNameEnum{
	"PER_INSTANCE":                   MeteredPricingPlanNameInstance,
	"PER_OCPU_MIN_BILLING_HRS":       MeteredPricingPlanNameOcpuMinBillingHrs,
	"PER_OCPU_LINEAR":                MeteredPricingPlanNameOcpuLinear,
	"PER_INSTANCE_MONTHLY_INCLUSIVE": MeteredPricingPlanNameInstanceMonthlyInclusive,
}

var mappingMeteredPricingPlanNameEnumLowerCase = map[string]MeteredPricingPlanNameEnum{
	"per_instance":                   MeteredPricingPlanNameInstance,
	"per_ocpu_min_billing_hrs":       MeteredPricingPlanNameOcpuMinBillingHrs,
	"per_ocpu_linear":                MeteredPricingPlanNameOcpuLinear,
	"per_instance_monthly_inclusive": MeteredPricingPlanNameInstanceMonthlyInclusive,
}

// GetMeteredPricingPlanNameEnumValues Enumerates the set of values for MeteredPricingPlanNameEnum
func GetMeteredPricingPlanNameEnumValues() []MeteredPricingPlanNameEnum {
	values := make([]MeteredPricingPlanNameEnum, 0)
	for _, v := range mappingMeteredPricingPlanNameEnum {
		values = append(values, v)
	}
	return values
}

// GetMeteredPricingPlanNameEnumStringValues Enumerates the set of values in String for MeteredPricingPlanNameEnum
func GetMeteredPricingPlanNameEnumStringValues() []string {
	return []string{
		"PER_INSTANCE",
		"PER_OCPU_MIN_BILLING_HRS",
		"PER_OCPU_LINEAR",
		"PER_INSTANCE_MONTHLY_INCLUSIVE",
	}
}

// GetMappingMeteredPricingPlanNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMeteredPricingPlanNameEnum(val string) (MeteredPricingPlanNameEnum, bool) {
	enum, ok := mappingMeteredPricingPlanNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
