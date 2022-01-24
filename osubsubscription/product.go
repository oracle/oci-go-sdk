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

// Product Product description
type Product struct {

	// Product part numner
	PartNumber *string `mandatory:"true" json:"partNumber"`

	// Product name
	Name *string `mandatory:"true" json:"name"`

	// Unit of measure
	UnitOfMeasure *string `mandatory:"true" json:"unitOfMeasure"`

	// Metered service billing category
	BillingCategory *string `mandatory:"false" json:"billingCategory"`

	// Product category
	ProductCategory *string `mandatory:"false" json:"productCategory"`

	// Rate card part type of Product
	UcmRateCardPartType *string `mandatory:"false" json:"ucmRateCardPartType"`
}

func (m Product) String() string {
	return common.PointerString(m)
}
