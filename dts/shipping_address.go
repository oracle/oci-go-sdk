// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ShippingAddress The representation of ShippingAddress
type ShippingAddress struct {
	Addressee *string `mandatory:"false" json:"addressee"`

	CareOf *string `mandatory:"false" json:"careOf"`

	Address1 *string `mandatory:"false" json:"address1"`

	Address2 *string `mandatory:"false" json:"address2"`

	Address3 *string `mandatory:"false" json:"address3"`

	Address4 *string `mandatory:"false" json:"address4"`

	CityOrLocality *string `mandatory:"false" json:"cityOrLocality"`

	StateOrRegion *string `mandatory:"false" json:"stateOrRegion"`

	Zipcode *string `mandatory:"false" json:"zipcode"`

	Country *string `mandatory:"false" json:"country"`

	PhoneNumber *string `mandatory:"false" json:"phoneNumber"`

	Email *string `mandatory:"false" json:"email"`
}

func (m ShippingAddress) String() string {
	return common.PointerString(m)
}
