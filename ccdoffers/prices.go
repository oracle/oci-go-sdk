// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Continuous Consumption Discounts API
//
// Use the Continuous Consumption Discounts API to manage continuous consumption discount operations.
//

package ccdoffers

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Prices The SKU external listing price.
type Prices struct {

	// Pricing model.
	Model PricesModelEnum `mandatory:"false" json:"model,omitempty"`

	// Currency value.
	Value *string `mandatory:"false" json:"value"`

	// The minimum quantity for tier pricing (exclusive). Only applicable for products that support range-based pricing.
	RangeMin *string `mandatory:"false" json:"rangeMin"`

	// The maximum quantity for tier pricing (inclusive). Only applicable for products that support range-based pricing.
	RangeMax *string `mandatory:"false" json:"rangeMax"`

	// The `rangeMin` and `rangeMax` unit of measurement (for example, "GB" or "TB").
	RangeUnit *string `mandatory:"false" json:"rangeUnit"`
}

func (m Prices) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Prices) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPricesModelEnum(string(m.Model)); !ok && m.Model != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Model: %s. Supported values are: %s.", m.Model, strings.Join(GetPricesModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PricesModelEnum Enum with underlying type: string
type PricesModelEnum string

// Set of constants representing the allowable values for PricesModelEnum
const (
	PricesModelPayAsYouGo PricesModelEnum = "PAY_AS_YOU_GO"
)

var mappingPricesModelEnum = map[string]PricesModelEnum{
	"PAY_AS_YOU_GO": PricesModelPayAsYouGo,
}

var mappingPricesModelEnumLowerCase = map[string]PricesModelEnum{
	"pay_as_you_go": PricesModelPayAsYouGo,
}

// GetPricesModelEnumValues Enumerates the set of values for PricesModelEnum
func GetPricesModelEnumValues() []PricesModelEnum {
	values := make([]PricesModelEnum, 0)
	for _, v := range mappingPricesModelEnum {
		values = append(values, v)
	}
	return values
}

// GetPricesModelEnumStringValues Enumerates the set of values in String for PricesModelEnum
func GetPricesModelEnumStringValues() []string {
	return []string{
		"PAY_AS_YOU_GO",
	}
}

// GetMappingPricesModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPricesModelEnum(val string) (PricesModelEnum, bool) {
	enum, ok := mappingPricesModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
