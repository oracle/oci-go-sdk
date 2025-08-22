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

// PricingPlan A pricing plan provided by the Publisher.
type PricingPlan interface {

	// List of pricing rates provider by publisher.
	GetRates() []PricingRate
}

type pricingplan struct {
	JsonData []byte
	Rates    []PricingRate `mandatory:"true" json:"rates"`
	PlanType string        `json:"planType"`
}

// UnmarshalJSON unmarshals json
func (m *pricingplan) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpricingplan pricingplan
	s := struct {
		Model Unmarshalerpricingplan
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Rates = s.Model.Rates
	m.PlanType = s.Model.PlanType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *pricingplan) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PlanType {
	case "METERED":
		mm := MeteredPricingPlan{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FIXED":
		mm := SaaSPricingPlan{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for PricingPlan: %s.", m.PlanType)
		return *m, nil
	}
}

// GetRates returns Rates
func (m pricingplan) GetRates() []PricingRate {
	return m.Rates
}

func (m pricingplan) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m pricingplan) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PricingPlanPlanTypeEnum Enum with underlying type: string
type PricingPlanPlanTypeEnum string

// Set of constants representing the allowable values for PricingPlanPlanTypeEnum
const (
	PricingPlanPlanTypeMetered PricingPlanPlanTypeEnum = "METERED"
	PricingPlanPlanTypeFixed   PricingPlanPlanTypeEnum = "FIXED"
)

var mappingPricingPlanPlanTypeEnum = map[string]PricingPlanPlanTypeEnum{
	"METERED": PricingPlanPlanTypeMetered,
	"FIXED":   PricingPlanPlanTypeFixed,
}

var mappingPricingPlanPlanTypeEnumLowerCase = map[string]PricingPlanPlanTypeEnum{
	"metered": PricingPlanPlanTypeMetered,
	"fixed":   PricingPlanPlanTypeFixed,
}

// GetPricingPlanPlanTypeEnumValues Enumerates the set of values for PricingPlanPlanTypeEnum
func GetPricingPlanPlanTypeEnumValues() []PricingPlanPlanTypeEnum {
	values := make([]PricingPlanPlanTypeEnum, 0)
	for _, v := range mappingPricingPlanPlanTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPricingPlanPlanTypeEnumStringValues Enumerates the set of values in String for PricingPlanPlanTypeEnum
func GetPricingPlanPlanTypeEnumStringValues() []string {
	return []string{
		"METERED",
		"FIXED",
	}
}

// GetMappingPricingPlanPlanTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPricingPlanPlanTypeEnum(val string) (PricingPlanPlanTypeEnum, bool) {
	enum, ok := mappingPricingPlanPlanTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
