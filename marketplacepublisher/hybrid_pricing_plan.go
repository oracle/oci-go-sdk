// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// HybridPricingPlan Hybrid SaaS pricing plan with a base fee and dimension-based overage billing.
type HybridPricingPlan struct {

	// List of pricing rates provider by publisher.
	Rates []PricingRate `mandatory:"true" json:"rates"`

	// The plan name.
	Name *string `mandatory:"true" json:"name"`

	// The plan description.
	PlanDescription *string `mandatory:"true" json:"planDescription"`

	// The usage dimensions that define included usage and overage billing for this plan.
	Dimensions []UsageDimension `mandatory:"true" json:"dimensions"`

	// Unique identifier of the pricing plan.
	PricingPlanKey *string `mandatory:"false" json:"pricingPlanKey"`

	// Additional metadata key/value pairs for the hybrid pricing.
	ExtendedMetadata map[string]string `mandatory:"false" json:"extendedMetadata"`

	// The plan billing frequency.
	BillingFrequency HybridPricingPlanBillingFrequencyEnum `mandatory:"true" json:"billingFrequency"`

	// The plan duration.
	PlanDuration HybridPricingPlanPlanDurationEnum `mandatory:"false" json:"planDuration,omitempty"`
}

// GetRates returns Rates
func (m HybridPricingPlan) GetRates() []PricingRate {
	return m.Rates
}

func (m HybridPricingPlan) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HybridPricingPlan) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHybridPricingPlanBillingFrequencyEnum(string(m.BillingFrequency)); !ok && m.BillingFrequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BillingFrequency: %s. Supported values are: %s.", m.BillingFrequency, strings.Join(GetHybridPricingPlanBillingFrequencyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingHybridPricingPlanPlanDurationEnum(string(m.PlanDuration)); !ok && m.PlanDuration != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanDuration: %s. Supported values are: %s.", m.PlanDuration, strings.Join(GetHybridPricingPlanPlanDurationEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HybridPricingPlan) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHybridPricingPlan HybridPricingPlan
	s := struct {
		DiscriminatorParam string `json:"planType"`
		MarshalTypeHybridPricingPlan
	}{
		"HYBRID",
		(MarshalTypeHybridPricingPlan)(m),
	}

	return json.Marshal(&s)
}

// HybridPricingPlanBillingFrequencyEnum Enum with underlying type: string
type HybridPricingPlanBillingFrequencyEnum string

// Set of constants representing the allowable values for HybridPricingPlanBillingFrequencyEnum
const (
	HybridPricingPlanBillingFrequencyMonthly    HybridPricingPlanBillingFrequencyEnum = "MONTHLY"
	HybridPricingPlanBillingFrequencyQuarterly  HybridPricingPlanBillingFrequencyEnum = "QUARTERLY"
	HybridPricingPlanBillingFrequencySemiAnnual HybridPricingPlanBillingFrequencyEnum = "SEMI_ANNUAL"
	HybridPricingPlanBillingFrequencyAnnual     HybridPricingPlanBillingFrequencyEnum = "ANNUAL"
	HybridPricingPlanBillingFrequencyBiennial   HybridPricingPlanBillingFrequencyEnum = "BIENNIAL"
	HybridPricingPlanBillingFrequencyTriennial  HybridPricingPlanBillingFrequencyEnum = "TRIENNIAL"
)

var mappingHybridPricingPlanBillingFrequencyEnum = map[string]HybridPricingPlanBillingFrequencyEnum{
	"MONTHLY":     HybridPricingPlanBillingFrequencyMonthly,
	"QUARTERLY":   HybridPricingPlanBillingFrequencyQuarterly,
	"SEMI_ANNUAL": HybridPricingPlanBillingFrequencySemiAnnual,
	"ANNUAL":      HybridPricingPlanBillingFrequencyAnnual,
	"BIENNIAL":    HybridPricingPlanBillingFrequencyBiennial,
	"TRIENNIAL":   HybridPricingPlanBillingFrequencyTriennial,
}

var mappingHybridPricingPlanBillingFrequencyEnumLowerCase = map[string]HybridPricingPlanBillingFrequencyEnum{
	"monthly":     HybridPricingPlanBillingFrequencyMonthly,
	"quarterly":   HybridPricingPlanBillingFrequencyQuarterly,
	"semi_annual": HybridPricingPlanBillingFrequencySemiAnnual,
	"annual":      HybridPricingPlanBillingFrequencyAnnual,
	"biennial":    HybridPricingPlanBillingFrequencyBiennial,
	"triennial":   HybridPricingPlanBillingFrequencyTriennial,
}

// GetHybridPricingPlanBillingFrequencyEnumValues Enumerates the set of values for HybridPricingPlanBillingFrequencyEnum
func GetHybridPricingPlanBillingFrequencyEnumValues() []HybridPricingPlanBillingFrequencyEnum {
	values := make([]HybridPricingPlanBillingFrequencyEnum, 0)
	for _, v := range mappingHybridPricingPlanBillingFrequencyEnum {
		values = append(values, v)
	}
	return values
}

// GetHybridPricingPlanBillingFrequencyEnumStringValues Enumerates the set of values in String for HybridPricingPlanBillingFrequencyEnum
func GetHybridPricingPlanBillingFrequencyEnumStringValues() []string {
	return []string{
		"MONTHLY",
		"QUARTERLY",
		"SEMI_ANNUAL",
		"ANNUAL",
		"BIENNIAL",
		"TRIENNIAL",
	}
}

// GetMappingHybridPricingPlanBillingFrequencyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHybridPricingPlanBillingFrequencyEnum(val string) (HybridPricingPlanBillingFrequencyEnum, bool) {
	enum, ok := mappingHybridPricingPlanBillingFrequencyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// HybridPricingPlanPlanDurationEnum Enum with underlying type: string
type HybridPricingPlanPlanDurationEnum string

// Set of constants representing the allowable values for HybridPricingPlanPlanDurationEnum
const (
	HybridPricingPlanPlanDurationMonthly    HybridPricingPlanPlanDurationEnum = "MONTHLY"
	HybridPricingPlanPlanDurationQuarterly  HybridPricingPlanPlanDurationEnum = "QUARTERLY"
	HybridPricingPlanPlanDurationSemiAnnual HybridPricingPlanPlanDurationEnum = "SEMI_ANNUAL"
	HybridPricingPlanPlanDurationAnnual     HybridPricingPlanPlanDurationEnum = "ANNUAL"
	HybridPricingPlanPlanDurationBiennial   HybridPricingPlanPlanDurationEnum = "BIENNIAL"
	HybridPricingPlanPlanDurationTriennial  HybridPricingPlanPlanDurationEnum = "TRIENNIAL"
)

var mappingHybridPricingPlanPlanDurationEnum = map[string]HybridPricingPlanPlanDurationEnum{
	"MONTHLY":     HybridPricingPlanPlanDurationMonthly,
	"QUARTERLY":   HybridPricingPlanPlanDurationQuarterly,
	"SEMI_ANNUAL": HybridPricingPlanPlanDurationSemiAnnual,
	"ANNUAL":      HybridPricingPlanPlanDurationAnnual,
	"BIENNIAL":    HybridPricingPlanPlanDurationBiennial,
	"TRIENNIAL":   HybridPricingPlanPlanDurationTriennial,
}

var mappingHybridPricingPlanPlanDurationEnumLowerCase = map[string]HybridPricingPlanPlanDurationEnum{
	"monthly":     HybridPricingPlanPlanDurationMonthly,
	"quarterly":   HybridPricingPlanPlanDurationQuarterly,
	"semi_annual": HybridPricingPlanPlanDurationSemiAnnual,
	"annual":      HybridPricingPlanPlanDurationAnnual,
	"biennial":    HybridPricingPlanPlanDurationBiennial,
	"triennial":   HybridPricingPlanPlanDurationTriennial,
}

// GetHybridPricingPlanPlanDurationEnumValues Enumerates the set of values for HybridPricingPlanPlanDurationEnum
func GetHybridPricingPlanPlanDurationEnumValues() []HybridPricingPlanPlanDurationEnum {
	values := make([]HybridPricingPlanPlanDurationEnum, 0)
	for _, v := range mappingHybridPricingPlanPlanDurationEnum {
		values = append(values, v)
	}
	return values
}

// GetHybridPricingPlanPlanDurationEnumStringValues Enumerates the set of values in String for HybridPricingPlanPlanDurationEnum
func GetHybridPricingPlanPlanDurationEnumStringValues() []string {
	return []string{
		"MONTHLY",
		"QUARTERLY",
		"SEMI_ANNUAL",
		"ANNUAL",
		"BIENNIAL",
		"TRIENNIAL",
	}
}

// GetMappingHybridPricingPlanPlanDurationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHybridPricingPlanPlanDurationEnum(val string) (HybridPricingPlanPlanDurationEnum, bool) {
	enum, ok := mappingHybridPricingPlanPlanDurationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
