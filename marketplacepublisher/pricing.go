// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// Pricing the pricing information for the offer
type Pricing struct {

	// The currency supported for a given Offer, in the format specified by ISO-4217
	CurrencyType *string `mandatory:"false" json:"currencyType"`

	// The total amount an Offer costs
	TotalAmount *int64 `mandatory:"false" json:"totalAmount"`

	// The frequency at which the customer is billed for the Offer
	BillingCycle PricingBillingCycleEnum `mandatory:"false" json:"billingCycle,omitempty"`
}

func (m Pricing) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Pricing) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPricingBillingCycleEnum(string(m.BillingCycle)); !ok && m.BillingCycle != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BillingCycle: %s. Supported values are: %s.", m.BillingCycle, strings.Join(GetPricingBillingCycleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PricingBillingCycleEnum Enum with underlying type: string
type PricingBillingCycleEnum string

// Set of constants representing the allowable values for PricingBillingCycleEnum
const (
	PricingBillingCycleOneTime PricingBillingCycleEnum = "ONE_TIME"
)

var mappingPricingBillingCycleEnum = map[string]PricingBillingCycleEnum{
	"ONE_TIME": PricingBillingCycleOneTime,
}

var mappingPricingBillingCycleEnumLowerCase = map[string]PricingBillingCycleEnum{
	"one_time": PricingBillingCycleOneTime,
}

// GetPricingBillingCycleEnumValues Enumerates the set of values for PricingBillingCycleEnum
func GetPricingBillingCycleEnumValues() []PricingBillingCycleEnum {
	values := make([]PricingBillingCycleEnum, 0)
	for _, v := range mappingPricingBillingCycleEnum {
		values = append(values, v)
	}
	return values
}

// GetPricingBillingCycleEnumStringValues Enumerates the set of values in String for PricingBillingCycleEnum
func GetPricingBillingCycleEnumStringValues() []string {
	return []string{
		"ONE_TIME",
	}
}

// GetMappingPricingBillingCycleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPricingBillingCycleEnum(val string) (PricingBillingCycleEnum, bool) {
	enum, ok := mappingPricingBillingCycleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
