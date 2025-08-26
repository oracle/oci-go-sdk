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

// SaaSPricingPlan SaaS pricing plan.
type SaaSPricingPlan struct {

	// List of pricing rates provider by publisher.
	Rates []PricingRate `mandatory:"true" json:"rates"`

	// The plan name.
	Name *string `mandatory:"true" json:"name"`

	// The plan description.
	PlanDescription *string `mandatory:"true" json:"planDescription"`

	// Additional metadata key/value pairs for the saas pricing.
	ExtendedMetadata map[string]string `mandatory:"false" json:"extendedMetadata"`

	// The plan billing frequency.
	BillingFrequency SaaSPricingPlanBillingFrequencyEnum `mandatory:"true" json:"billingFrequency"`
}

// GetRates returns Rates
func (m SaaSPricingPlan) GetRates() []PricingRate {
	return m.Rates
}

func (m SaaSPricingPlan) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SaaSPricingPlan) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSaaSPricingPlanBillingFrequencyEnum(string(m.BillingFrequency)); !ok && m.BillingFrequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BillingFrequency: %s. Supported values are: %s.", m.BillingFrequency, strings.Join(GetSaaSPricingPlanBillingFrequencyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SaaSPricingPlan) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSaaSPricingPlan SaaSPricingPlan
	s := struct {
		DiscriminatorParam string `json:"planType"`
		MarshalTypeSaaSPricingPlan
	}{
		"FIXED",
		(MarshalTypeSaaSPricingPlan)(m),
	}

	return json.Marshal(&s)
}

// SaaSPricingPlanBillingFrequencyEnum Enum with underlying type: string
type SaaSPricingPlanBillingFrequencyEnum string

// Set of constants representing the allowable values for SaaSPricingPlanBillingFrequencyEnum
const (
	SaaSPricingPlanBillingFrequencyMonthly   SaaSPricingPlanBillingFrequencyEnum = "MONTHLY"
	SaaSPricingPlanBillingFrequencyQuarterly SaaSPricingPlanBillingFrequencyEnum = "QUARTERLY"
	SaaSPricingPlanBillingFrequencyYearly    SaaSPricingPlanBillingFrequencyEnum = "YEARLY"
)

var mappingSaaSPricingPlanBillingFrequencyEnum = map[string]SaaSPricingPlanBillingFrequencyEnum{
	"MONTHLY":   SaaSPricingPlanBillingFrequencyMonthly,
	"QUARTERLY": SaaSPricingPlanBillingFrequencyQuarterly,
	"YEARLY":    SaaSPricingPlanBillingFrequencyYearly,
}

var mappingSaaSPricingPlanBillingFrequencyEnumLowerCase = map[string]SaaSPricingPlanBillingFrequencyEnum{
	"monthly":   SaaSPricingPlanBillingFrequencyMonthly,
	"quarterly": SaaSPricingPlanBillingFrequencyQuarterly,
	"yearly":    SaaSPricingPlanBillingFrequencyYearly,
}

// GetSaaSPricingPlanBillingFrequencyEnumValues Enumerates the set of values for SaaSPricingPlanBillingFrequencyEnum
func GetSaaSPricingPlanBillingFrequencyEnumValues() []SaaSPricingPlanBillingFrequencyEnum {
	values := make([]SaaSPricingPlanBillingFrequencyEnum, 0)
	for _, v := range mappingSaaSPricingPlanBillingFrequencyEnum {
		values = append(values, v)
	}
	return values
}

// GetSaaSPricingPlanBillingFrequencyEnumStringValues Enumerates the set of values in String for SaaSPricingPlanBillingFrequencyEnum
func GetSaaSPricingPlanBillingFrequencyEnumStringValues() []string {
	return []string{
		"MONTHLY",
		"QUARTERLY",
		"YEARLY",
	}
}

// GetMappingSaaSPricingPlanBillingFrequencyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSaaSPricingPlanBillingFrequencyEnum(val string) (SaaSPricingPlanBillingFrequencyEnum, bool) {
	enum, ok := mappingSaaSPricingPlanBillingFrequencyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
