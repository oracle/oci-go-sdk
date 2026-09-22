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

// UsageBasedPricingPlan Usage-based SaaS pricing plan where all billing is driven by dimensions.
type UsageBasedPricingPlan struct {

	// List of pricing rates provider by publisher.
	Rates []PricingRate `mandatory:"true" json:"rates"`

	// The plan name.
	Name *string `mandatory:"true" json:"name"`

	// The plan description.
	PlanDescription *string `mandatory:"true" json:"planDescription"`

	// The usage dimensions that define billable usage for this plan.
	Dimensions []UsageDimension `mandatory:"true" json:"dimensions"`

	// Unique identifier of the pricing plan.
	PricingPlanKey *string `mandatory:"false" json:"pricingPlanKey"`

	// Additional metadata key/value pairs for the usage-based pricing.
	ExtendedMetadata map[string]string `mandatory:"false" json:"extendedMetadata"`

	// The plan billing frequency.
	BillingFrequency UsageBasedPricingPlanBillingFrequencyEnum `mandatory:"true" json:"billingFrequency"`

	// The plan duration.
	PlanDuration UsageBasedPricingPlanPlanDurationEnum `mandatory:"false" json:"planDuration,omitempty"`
}

// GetRates returns Rates
func (m UsageBasedPricingPlan) GetRates() []PricingRate {
	return m.Rates
}

func (m UsageBasedPricingPlan) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UsageBasedPricingPlan) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUsageBasedPricingPlanBillingFrequencyEnum(string(m.BillingFrequency)); !ok && m.BillingFrequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BillingFrequency: %s. Supported values are: %s.", m.BillingFrequency, strings.Join(GetUsageBasedPricingPlanBillingFrequencyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUsageBasedPricingPlanPlanDurationEnum(string(m.PlanDuration)); !ok && m.PlanDuration != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanDuration: %s. Supported values are: %s.", m.PlanDuration, strings.Join(GetUsageBasedPricingPlanPlanDurationEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UsageBasedPricingPlan) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUsageBasedPricingPlan UsageBasedPricingPlan
	s := struct {
		DiscriminatorParam string `json:"planType"`
		MarshalTypeUsageBasedPricingPlan
	}{
		"USAGE_BASED",
		(MarshalTypeUsageBasedPricingPlan)(m),
	}

	return json.Marshal(&s)
}

// UsageBasedPricingPlanBillingFrequencyEnum Enum with underlying type: string
type UsageBasedPricingPlanBillingFrequencyEnum string

// Set of constants representing the allowable values for UsageBasedPricingPlanBillingFrequencyEnum
const (
	UsageBasedPricingPlanBillingFrequencyMonthly    UsageBasedPricingPlanBillingFrequencyEnum = "MONTHLY"
	UsageBasedPricingPlanBillingFrequencyQuarterly  UsageBasedPricingPlanBillingFrequencyEnum = "QUARTERLY"
	UsageBasedPricingPlanBillingFrequencySemiAnnual UsageBasedPricingPlanBillingFrequencyEnum = "SEMI_ANNUAL"
	UsageBasedPricingPlanBillingFrequencyAnnual     UsageBasedPricingPlanBillingFrequencyEnum = "ANNUAL"
	UsageBasedPricingPlanBillingFrequencyBiennial   UsageBasedPricingPlanBillingFrequencyEnum = "BIENNIAL"
	UsageBasedPricingPlanBillingFrequencyTriennial  UsageBasedPricingPlanBillingFrequencyEnum = "TRIENNIAL"
)

var mappingUsageBasedPricingPlanBillingFrequencyEnum = map[string]UsageBasedPricingPlanBillingFrequencyEnum{
	"MONTHLY":     UsageBasedPricingPlanBillingFrequencyMonthly,
	"QUARTERLY":   UsageBasedPricingPlanBillingFrequencyQuarterly,
	"SEMI_ANNUAL": UsageBasedPricingPlanBillingFrequencySemiAnnual,
	"ANNUAL":      UsageBasedPricingPlanBillingFrequencyAnnual,
	"BIENNIAL":    UsageBasedPricingPlanBillingFrequencyBiennial,
	"TRIENNIAL":   UsageBasedPricingPlanBillingFrequencyTriennial,
}

var mappingUsageBasedPricingPlanBillingFrequencyEnumLowerCase = map[string]UsageBasedPricingPlanBillingFrequencyEnum{
	"monthly":     UsageBasedPricingPlanBillingFrequencyMonthly,
	"quarterly":   UsageBasedPricingPlanBillingFrequencyQuarterly,
	"semi_annual": UsageBasedPricingPlanBillingFrequencySemiAnnual,
	"annual":      UsageBasedPricingPlanBillingFrequencyAnnual,
	"biennial":    UsageBasedPricingPlanBillingFrequencyBiennial,
	"triennial":   UsageBasedPricingPlanBillingFrequencyTriennial,
}

// GetUsageBasedPricingPlanBillingFrequencyEnumValues Enumerates the set of values for UsageBasedPricingPlanBillingFrequencyEnum
func GetUsageBasedPricingPlanBillingFrequencyEnumValues() []UsageBasedPricingPlanBillingFrequencyEnum {
	values := make([]UsageBasedPricingPlanBillingFrequencyEnum, 0)
	for _, v := range mappingUsageBasedPricingPlanBillingFrequencyEnum {
		values = append(values, v)
	}
	return values
}

// GetUsageBasedPricingPlanBillingFrequencyEnumStringValues Enumerates the set of values in String for UsageBasedPricingPlanBillingFrequencyEnum
func GetUsageBasedPricingPlanBillingFrequencyEnumStringValues() []string {
	return []string{
		"MONTHLY",
		"QUARTERLY",
		"SEMI_ANNUAL",
		"ANNUAL",
		"BIENNIAL",
		"TRIENNIAL",
	}
}

// GetMappingUsageBasedPricingPlanBillingFrequencyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUsageBasedPricingPlanBillingFrequencyEnum(val string) (UsageBasedPricingPlanBillingFrequencyEnum, bool) {
	enum, ok := mappingUsageBasedPricingPlanBillingFrequencyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UsageBasedPricingPlanPlanDurationEnum Enum with underlying type: string
type UsageBasedPricingPlanPlanDurationEnum string

// Set of constants representing the allowable values for UsageBasedPricingPlanPlanDurationEnum
const (
	UsageBasedPricingPlanPlanDurationMonthly    UsageBasedPricingPlanPlanDurationEnum = "MONTHLY"
	UsageBasedPricingPlanPlanDurationQuarterly  UsageBasedPricingPlanPlanDurationEnum = "QUARTERLY"
	UsageBasedPricingPlanPlanDurationSemiAnnual UsageBasedPricingPlanPlanDurationEnum = "SEMI_ANNUAL"
	UsageBasedPricingPlanPlanDurationAnnual     UsageBasedPricingPlanPlanDurationEnum = "ANNUAL"
	UsageBasedPricingPlanPlanDurationBiennial   UsageBasedPricingPlanPlanDurationEnum = "BIENNIAL"
	UsageBasedPricingPlanPlanDurationTriennial  UsageBasedPricingPlanPlanDurationEnum = "TRIENNIAL"
)

var mappingUsageBasedPricingPlanPlanDurationEnum = map[string]UsageBasedPricingPlanPlanDurationEnum{
	"MONTHLY":     UsageBasedPricingPlanPlanDurationMonthly,
	"QUARTERLY":   UsageBasedPricingPlanPlanDurationQuarterly,
	"SEMI_ANNUAL": UsageBasedPricingPlanPlanDurationSemiAnnual,
	"ANNUAL":      UsageBasedPricingPlanPlanDurationAnnual,
	"BIENNIAL":    UsageBasedPricingPlanPlanDurationBiennial,
	"TRIENNIAL":   UsageBasedPricingPlanPlanDurationTriennial,
}

var mappingUsageBasedPricingPlanPlanDurationEnumLowerCase = map[string]UsageBasedPricingPlanPlanDurationEnum{
	"monthly":     UsageBasedPricingPlanPlanDurationMonthly,
	"quarterly":   UsageBasedPricingPlanPlanDurationQuarterly,
	"semi_annual": UsageBasedPricingPlanPlanDurationSemiAnnual,
	"annual":      UsageBasedPricingPlanPlanDurationAnnual,
	"biennial":    UsageBasedPricingPlanPlanDurationBiennial,
	"triennial":   UsageBasedPricingPlanPlanDurationTriennial,
}

// GetUsageBasedPricingPlanPlanDurationEnumValues Enumerates the set of values for UsageBasedPricingPlanPlanDurationEnum
func GetUsageBasedPricingPlanPlanDurationEnumValues() []UsageBasedPricingPlanPlanDurationEnum {
	values := make([]UsageBasedPricingPlanPlanDurationEnum, 0)
	for _, v := range mappingUsageBasedPricingPlanPlanDurationEnum {
		values = append(values, v)
	}
	return values
}

// GetUsageBasedPricingPlanPlanDurationEnumStringValues Enumerates the set of values in String for UsageBasedPricingPlanPlanDurationEnum
func GetUsageBasedPricingPlanPlanDurationEnumStringValues() []string {
	return []string{
		"MONTHLY",
		"QUARTERLY",
		"SEMI_ANNUAL",
		"ANNUAL",
		"BIENNIAL",
		"TRIENNIAL",
	}
}

// GetMappingUsageBasedPricingPlanPlanDurationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUsageBasedPricingPlanPlanDurationEnum(val string) (UsageBasedPricingPlanPlanDurationEnum, bool) {
	enum, ok := mappingUsageBasedPricingPlanPlanDurationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
