// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription API Subscription, Commitment and and Rate Card Details
//
// Set of APIs that return the Subscription Details, Commitment and Effective Rate Card Details
//

package osubsubscription

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// RateCardTier Rate Card Tier details
type RateCardTier struct {

	// Rate card tier quantity range
	UpToQuantity *string `mandatory:"false" json:"upToQuantity"`

	// Rate card tier net unit price
	NetUnitPrice *string `mandatory:"false" json:"netUnitPrice"`

	// Rate card tier overage price
	OveragePrice *string `mandatory:"false" json:"overagePrice"`
}

func (m RateCardTier) String() string {
	return common.PointerString(m)
}
