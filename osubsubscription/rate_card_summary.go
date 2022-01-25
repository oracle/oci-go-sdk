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

// RateCardSummary Rate Card Summary
type RateCardSummary struct {
	Product *Product `mandatory:"true" json:"product"`

	// Rate card net unit price
	NetUnitPrice *string `mandatory:"true" json:"netUnitPrice"`

	// Rate card overage price
	OveragePrice *string `mandatory:"true" json:"overagePrice"`

	// Rate card start date
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Rate card end date
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// Rate card discretionary discount percentage
	DiscretionaryDiscountPercentage *string `mandatory:"false" json:"discretionaryDiscountPercentage"`

	// Rate card price tier flag
	IsTier *bool `mandatory:"false" json:"isTier"`

	Currency *Currency `mandatory:"false" json:"currency"`

	// List of tiered rate card prices
	RateCardTiers []RateCardTier `mandatory:"false" json:"rateCardTiers"`
}

func (m RateCardSummary) String() string {
	return common.PointerString(m)
}
