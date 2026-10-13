// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Continuous Consumption Discounts API
//
// Use the Continuous Consumption Discounts API to manage continuous consumption discount operations.
//

package ccdoffers

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BundleSku A polymorphic bundle SKU. The `quantityRule` discriminator identifies how the quantity for this bundled SKU is determined.
// Every bundled SKU in an offer must use the same `quantityRule`; don't mix `FIXED_RATIO` and `CUSTOMER_SELECTED` bundled
// SKUs in one offer.
type BundleSku interface {

	// The bundled SKU part number (prefaced with "B"). For example, the "E4 MEMORY" SKU will be bundled with the main "E4 OCPUs" SKU.
	GetSku() *string

	// The SKU display name is derived by CCD. If CCD cannot derive a unique value, it uses the caller-provided value.
	GetSkuDisplayName() *string

	// The SKU billing metric is derived by CCD. If CCD cannot derive a unique value, it uses the caller-provided value.
	GetSkuMetric() *string

	// List price of a SKU. Only set this property if the product or SKU is private. Tier pricing SKUs aren't supported.
	GetListPrice() *string

	// The currency code of `listPrice` associated with the product or SKU. Only set this property if the product or SKU is private.
	// Tier pricing SKUs aren't supported.
	GetCurrencyCode() *string
}

type bundlesku struct {
	JsonData       []byte
	SkuDisplayName *string `mandatory:"false" json:"skuDisplayName"`
	SkuMetric      *string `mandatory:"false" json:"skuMetric"`
	ListPrice      *string `mandatory:"false" json:"listPrice"`
	CurrencyCode   *string `mandatory:"false" json:"currencyCode"`
	Sku            *string `mandatory:"true" json:"sku"`
	QuantityRule   string  `json:"quantityRule"`
}

// UnmarshalJSON unmarshals json
func (m *bundlesku) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbundlesku bundlesku
	s := struct {
		Model Unmarshalerbundlesku
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Sku = s.Model.Sku
	m.SkuDisplayName = s.Model.SkuDisplayName
	m.SkuMetric = s.Model.SkuMetric
	m.ListPrice = s.Model.ListPrice
	m.CurrencyCode = s.Model.CurrencyCode
	m.QuantityRule = s.Model.QuantityRule

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *bundlesku) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.QuantityRule {
	case "FIXED_RATIO":
		mm := FixedRatioBundleSku{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOMER_SELECTED":
		mm := CustomerSelectedBundleSku{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for BundleSku: %s.", m.QuantityRule)
		return *m, nil
	}
}

// GetSkuDisplayName returns SkuDisplayName
func (m bundlesku) GetSkuDisplayName() *string {
	return m.SkuDisplayName
}

// GetSkuMetric returns SkuMetric
func (m bundlesku) GetSkuMetric() *string {
	return m.SkuMetric
}

// GetListPrice returns ListPrice
func (m bundlesku) GetListPrice() *string {
	return m.ListPrice
}

// GetCurrencyCode returns CurrencyCode
func (m bundlesku) GetCurrencyCode() *string {
	return m.CurrencyCode
}

// GetSku returns Sku
func (m bundlesku) GetSku() *string {
	return m.Sku
}

func (m bundlesku) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m bundlesku) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BundleSkuQuantityRuleEnum Enum with underlying type: string
type BundleSkuQuantityRuleEnum string

// Set of constants representing the allowable values for BundleSkuQuantityRuleEnum
const (
	BundleSkuQuantityRuleFixedRatio       BundleSkuQuantityRuleEnum = "FIXED_RATIO"
	BundleSkuQuantityRuleCustomerSelected BundleSkuQuantityRuleEnum = "CUSTOMER_SELECTED"
)

var mappingBundleSkuQuantityRuleEnum = map[string]BundleSkuQuantityRuleEnum{
	"FIXED_RATIO":       BundleSkuQuantityRuleFixedRatio,
	"CUSTOMER_SELECTED": BundleSkuQuantityRuleCustomerSelected,
}

var mappingBundleSkuQuantityRuleEnumLowerCase = map[string]BundleSkuQuantityRuleEnum{
	"fixed_ratio":       BundleSkuQuantityRuleFixedRatio,
	"customer_selected": BundleSkuQuantityRuleCustomerSelected,
}

// GetBundleSkuQuantityRuleEnumValues Enumerates the set of values for BundleSkuQuantityRuleEnum
func GetBundleSkuQuantityRuleEnumValues() []BundleSkuQuantityRuleEnum {
	values := make([]BundleSkuQuantityRuleEnum, 0)
	for _, v := range mappingBundleSkuQuantityRuleEnum {
		values = append(values, v)
	}
	return values
}

// GetBundleSkuQuantityRuleEnumStringValues Enumerates the set of values in String for BundleSkuQuantityRuleEnum
func GetBundleSkuQuantityRuleEnumStringValues() []string {
	return []string{
		"FIXED_RATIO",
		"CUSTOMER_SELECTED",
	}
}

// GetMappingBundleSkuQuantityRuleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBundleSkuQuantityRuleEnum(val string) (BundleSkuQuantityRuleEnum, bool) {
	enum, ok := mappingBundleSkuQuantityRuleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
