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

// CustomerSelectedBundleSku A bundle SKU whose quantity is selected by the customer within the configured range.
// For FLEX offers, the maxQuantity should be equal to minQuantity if provided.
type CustomerSelectedBundleSku struct {

	// The bundled SKU part number (prefaced with "B"). For example, the "E4 MEMORY" SKU will be bundled with the main "E4 OCPUs" SKU.
	Sku *string `mandatory:"true" json:"sku"`

	// The minimum quantity the customer can select for this bundled SKU.
	// For FLEX offers, minQuantity and maxQuantity must be equal if maxQuantity provided.
	MinQuantity *int64 `mandatory:"true" json:"minQuantity"`

	// The SKU display name is derived by CCD. If CCD cannot derive a unique value, it uses the caller-provided value.
	SkuDisplayName *string `mandatory:"false" json:"skuDisplayName"`

	// The SKU billing metric is derived by CCD. If CCD cannot derive a unique value, it uses the caller-provided value.
	SkuMetric *string `mandatory:"false" json:"skuMetric"`

	// List price of a SKU. Only set this property if the product or SKU is private. Tier pricing SKUs aren't supported.
	ListPrice *string `mandatory:"false" json:"listPrice"`

	// The currency code of `listPrice` associated with the product or SKU. Only set this property if the product or SKU is private.
	// Tier pricing SKUs aren't supported.
	CurrencyCode *string `mandatory:"false" json:"currencyCode"`

	// The maximum quantity the customer can select for this bundled SKU.
	// For FLEX offers must be equal to minQuantity if provided.
	MaxQuantity *int64 `mandatory:"false" json:"maxQuantity"`
}

// GetSku returns Sku
func (m CustomerSelectedBundleSku) GetSku() *string {
	return m.Sku
}

// GetSkuDisplayName returns SkuDisplayName
func (m CustomerSelectedBundleSku) GetSkuDisplayName() *string {
	return m.SkuDisplayName
}

// GetSkuMetric returns SkuMetric
func (m CustomerSelectedBundleSku) GetSkuMetric() *string {
	return m.SkuMetric
}

// GetListPrice returns ListPrice
func (m CustomerSelectedBundleSku) GetListPrice() *string {
	return m.ListPrice
}

// GetCurrencyCode returns CurrencyCode
func (m CustomerSelectedBundleSku) GetCurrencyCode() *string {
	return m.CurrencyCode
}

func (m CustomerSelectedBundleSku) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomerSelectedBundleSku) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CustomerSelectedBundleSku) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCustomerSelectedBundleSku CustomerSelectedBundleSku
	s := struct {
		DiscriminatorParam string `json:"quantityRule"`
		MarshalTypeCustomerSelectedBundleSku
	}{
		"CUSTOMER_SELECTED",
		(MarshalTypeCustomerSelectedBundleSku)(m),
	}

	return json.Marshal(&s)
}
