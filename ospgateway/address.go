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

// Address Address details model
type Address struct {

	// Name of the contact person
	ContactName *string `mandatory:"false" json:"contactName"`

	// Name of the customer company
	CompanyName *string `mandatory:"false" json:"companyName"`

	// Address line 1
	AddressLine1 *string `mandatory:"false" json:"addressLine1"`

	// Address line 2
	AddressLine2 *string `mandatory:"false" json:"addressLine2"`

	// Address line 3
	AddressLine3 *string `mandatory:"false" json:"addressLine3"`

	// Address line 4
	AddressLine4 *string `mandatory:"false" json:"addressLine4"`

	// Street name
	StreetName *string `mandatory:"false" json:"streetName"`

	// House no
	StreetNumber *string `mandatory:"false" json:"streetNumber"`

	// Name of the city
	City *string `mandatory:"false" json:"city"`

	Country *Country `mandatory:"false" json:"country"`

	// County name
	County *string `mandatory:"false" json:"county"`

	// Name of the state
	State *string `mandatory:"false" json:"state"`

	// ZIP no
	PostalCode *string `mandatory:"false" json:"postalCode"`

	// Name of the province
	Province *string `mandatory:"false" json:"province"`
}

func (m Address) String() string {
	return common.PointerString(m)
}
