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

// SupportedCurrencySummary The model for the currencies supported by the Marketplace Publisher Service.
type SupportedCurrencySummary struct {

	// The currency code, in the format specified by ISO-4217.
	CurrencyCode *string `mandatory:"true" json:"currencyCode"`

	// Oracle exchange rate for the currency in USD.
	ExchangeRate *float32 `mandatory:"true" json:"exchangeRate"`

	// The current state for the currency.
	LifecycleState SupportedCurrencySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

func (m SupportedCurrencySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SupportedCurrencySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSupportedCurrencySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSupportedCurrencySummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SupportedCurrencySummaryLifecycleStateEnum Enum with underlying type: string
type SupportedCurrencySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for SupportedCurrencySummaryLifecycleStateEnum
const (
	SupportedCurrencySummaryLifecycleStateActive   SupportedCurrencySummaryLifecycleStateEnum = "ACTIVE"
	SupportedCurrencySummaryLifecycleStateInactive SupportedCurrencySummaryLifecycleStateEnum = "INACTIVE"
)

var mappingSupportedCurrencySummaryLifecycleStateEnum = map[string]SupportedCurrencySummaryLifecycleStateEnum{
	"ACTIVE":   SupportedCurrencySummaryLifecycleStateActive,
	"INACTIVE": SupportedCurrencySummaryLifecycleStateInactive,
}

var mappingSupportedCurrencySummaryLifecycleStateEnumLowerCase = map[string]SupportedCurrencySummaryLifecycleStateEnum{
	"active":   SupportedCurrencySummaryLifecycleStateActive,
	"inactive": SupportedCurrencySummaryLifecycleStateInactive,
}

// GetSupportedCurrencySummaryLifecycleStateEnumValues Enumerates the set of values for SupportedCurrencySummaryLifecycleStateEnum
func GetSupportedCurrencySummaryLifecycleStateEnumValues() []SupportedCurrencySummaryLifecycleStateEnum {
	values := make([]SupportedCurrencySummaryLifecycleStateEnum, 0)
	for _, v := range mappingSupportedCurrencySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSupportedCurrencySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for SupportedCurrencySummaryLifecycleStateEnum
func GetSupportedCurrencySummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingSupportedCurrencySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSupportedCurrencySummaryLifecycleStateEnum(val string) (SupportedCurrencySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingSupportedCurrencySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
