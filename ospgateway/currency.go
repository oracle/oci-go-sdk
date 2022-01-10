// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Billing Center Gateway API
//
// This site describes all the Rest endpoints of Billing Center Gateway.
//

package ospgateway

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// Currency Currency details model
type Currency struct {

	// Currency code
	CurrencyCode *string `mandatory:"false" json:"currencyCode"`

	// Currency symbol
	CurrencySymbol *string `mandatory:"false" json:"currencySymbol"`

	// Name of the currency
	Name *string `mandatory:"false" json:"name"`

	// USD conversion rate of the currency
	UsdConversion *float32 `mandatory:"false" json:"usdConversion"`

	// Round decimal point
	RoundDecimalPoint *float32 `mandatory:"false" json:"roundDecimalPoint"`
}

func (m Currency) String() string {
	return common.PointerString(m)
}
