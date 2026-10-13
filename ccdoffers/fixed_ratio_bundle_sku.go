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

// FixedRatioBundleSku A bundle SKU whose quantity is derived from the primary SKU quantity using a fixed ratio.
type FixedRatioBundleSku struct {

	// The bundled SKU part number (prefaced with "B"). For example, the "E4 MEMORY" SKU will be bundled with the main "E4 OCPUs" SKU.
	Sku *string `mandatory:"true" json:"sku"`

	// The primary-to-bundled SKU quantity ratio.
	Ratio *string `mandatory:"true" json:"ratio"`

	// The SKU display name is derived by CCD. If CCD cannot derive a unique value, it uses the caller-provided value.
	SkuDisplayName *string `mandatory:"false" json:"skuDisplayName"`

	// The SKU billing metric is derived by CCD. If CCD cannot derive a unique value, it uses the caller-provided value.
	SkuMetric *string `mandatory:"false" json:"skuMetric"`

	// List price of a SKU. Only set this property if the product or SKU is private. Tier pricing SKUs aren't supported.
	ListPrice *string `mandatory:"false" json:"listPrice"`

	// The currency code of `listPrice` associated with the product or SKU. Only set this property if the product or SKU is private.
	// Tier pricing SKUs aren't supported.
	CurrencyCode *string `mandatory:"false" json:"currencyCode"`
}

// GetSku returns Sku
func (m FixedRatioBundleSku) GetSku() *string {
	return m.Sku
}

// GetSkuDisplayName returns SkuDisplayName
func (m FixedRatioBundleSku) GetSkuDisplayName() *string {
	return m.SkuDisplayName
}

// GetSkuMetric returns SkuMetric
func (m FixedRatioBundleSku) GetSkuMetric() *string {
	return m.SkuMetric
}

// GetListPrice returns ListPrice
func (m FixedRatioBundleSku) GetListPrice() *string {
	return m.ListPrice
}

// GetCurrencyCode returns CurrencyCode
func (m FixedRatioBundleSku) GetCurrencyCode() *string {
	return m.CurrencyCode
}

func (m FixedRatioBundleSku) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FixedRatioBundleSku) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FixedRatioBundleSku) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFixedRatioBundleSku FixedRatioBundleSku
	s := struct {
		DiscriminatorParam string `json:"quantityRule"`
		MarshalTypeFixedRatioBundleSku
	}{
		"FIXED_RATIO",
		(MarshalTypeFixedRatioBundleSku)(m),
	}

	return json.Marshal(&s)
}
