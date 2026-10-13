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

// SecondarySkus A secondary SKU associated with a CCD commitment. Its quantity behavior is defined by the corresponding bundled SKU
// on the CCD offer.
type SecondarySkus struct {

	// The secondary SKU part number (prefaced with "B"). For example, the "E4 MEMORY" SKU will be the secondary SKU with the main "E4 OCPUs" SKU.
	// Must also be the same as `CcdOffers's` `bundleSkus` `sku` value.
	Sku *string `mandatory:"true" json:"sku"`

	// The canonical SKU display name derived from the associated offer. This property remains accepted in create and update
	// requests for compatibility, but CCD uses the offer-derived value. Avoid entering confidential information.
	SkuDisplayName *string `mandatory:"true" json:"skuDisplayName"`

	// The canonical billing metric derived from the associated offer. This property remains accepted in create and update
	// requests for compatibility, but CCD uses the offer-derived value.
	SkuMetric *string `mandatory:"false" json:"skuMetric"`

	// The resolved hourly commitment quantity for this secondary SKU. In create or update requests, this property is required
	// when the corresponding offer component uses `CUSTOMER_SELECTED`, and the value must be within the offer component's
	// minimum and maximum. For FLEX offers with `CUSTOMER_SELECTED`, this quantity must equal the offer component's `minQuantity`
	// and, if defined, `maxQuantity`. For `FIXED_RATIO`, this property may be omitted because CCD derives it from the primary
	// quantity and the offer-defined ratio. If an existing client supplies a quantity for `FIXED_RATIO`, the offer-defined ratio is authoritative.
	Quantity *int64 `mandatory:"false" json:"quantity"`
}

func (m SecondarySkus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecondarySkus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
