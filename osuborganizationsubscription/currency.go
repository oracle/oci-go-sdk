// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription Gateway API Organization's Subscription
//
// API that returns data for the list of subscription ids returned from Organizations API
//

package osuborganizationsubscription

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Currency Currency details
type Currency struct {

	// Currency Code
	IsoCode *string `mandatory:"true" json:"isoCode"`

	// Currency name
	Name *string `mandatory:"false" json:"name"`

	// Standard Precision of the Currency
	StdPrecision *int64 `mandatory:"false" json:"stdPrecision"`
}

func (m Currency) String() string {
	return common.PointerString(m)
}
