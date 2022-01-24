// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription API Billing Schedule
//
// OneSubscription API for Billing Schedule information
//

package osubbillingschedule

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Product Product description
type Product struct {

	// Indicates the associated AR Invoice Number
	PartNumber *string `mandatory:"true" json:"partNumber"`

	// Product name
	Name *string `mandatory:"true" json:"name"`
}

func (m Product) String() string {
	return common.PointerString(m)
}
