// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// ShippingAddress Shipping address for rover devices.
type ShippingAddress struct {

	// Addressee in shipping address.
	Addressee *string `mandatory:"true" json:"addressee"`

	// Address line 1.
	Address1 *string `mandatory:"true" json:"address1"`

	// city or locality for shipping address.
	CityOrLocality *string `mandatory:"true" json:"cityOrLocality"`

	// state or region for shipping address.
	StateOrRegion *string `mandatory:"true" json:"stateOrRegion"`

	// zipcode for shipping address.
	Zipcode *string `mandatory:"true" json:"zipcode"`

	// country for shipping address.
	Country *string `mandatory:"true" json:"country"`

	// recepient phone number.
	PhoneNumber *string `mandatory:"true" json:"phoneNumber"`

	// CareOf for shipping address.
	CareOf *string `mandatory:"false" json:"careOf"`

	// Address line 2.
	Address2 *string `mandatory:"false" json:"address2"`

	// Address line 3.
	Address3 *string `mandatory:"false" json:"address3"`

	// Address line 4.
	Address4 *string `mandatory:"false" json:"address4"`

	// recepient email address.
	Email *string `mandatory:"false" json:"email"`
}

func (m ShippingAddress) String() string {
	return common.PointerString(m)
}
