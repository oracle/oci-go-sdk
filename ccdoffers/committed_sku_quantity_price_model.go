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

// CommittedSkuQuantityPriceModel The committed SKU, quantity, and its price information model.
type CommittedSkuQuantityPriceModel struct {

	// The SKU part number (prefaced with "B") associated with the committed CCD.
	Sku *string `mandatory:"false" json:"sku"`

	// The canonical SKU display name derived from the associated offer.
	SkuDisplayName *string `mandatory:"false" json:"skuDisplayName"`

	// The canonical billing metric derived from the associated offer.
	SkuMetric *string `mandatory:"false" json:"skuMetric"`

	// True if the SKU is the primary SKU of a CCD offer, otheriwse, false.
	IsMainSku *bool `mandatory:"false" json:"isMainSku"`

	// The SKU quantity that's getting committed.
	Quantity *int64 `mandatory:"false" json:"quantity"`

	// The pricing model associated with the currency code.
	CurrencyCodeLocalization []CurrencyCodeLocalization `mandatory:"false" json:"currencyCodeLocalization"`
}

func (m CommittedSkuQuantityPriceModel) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CommittedSkuQuantityPriceModel) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
