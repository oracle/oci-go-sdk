// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ShippingAddress The representation of ShippingAddress
type ShippingAddress struct {
	Addressee *string `mandatory:"true" json:"addressee"`

	Address1 *string `mandatory:"true" json:"address1"`

	CityOrLocality *string `mandatory:"true" json:"cityOrLocality"`

	StateOrRegion *string `mandatory:"true" json:"stateOrRegion"`

	Zipcode *string `mandatory:"true" json:"zipcode"`

	Country *string `mandatory:"true" json:"country"`

	CareOf *string `mandatory:"false" json:"careOf"`

	Address2 *string `mandatory:"false" json:"address2"`

	Address3 *string `mandatory:"false" json:"address3"`

	Address4 *string `mandatory:"false" json:"address4"`

	PhoneNumber *string `mandatory:"false" json:"phoneNumber"`

	Email *string `mandatory:"false" json:"email"`
}

func (m ShippingAddress) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShippingAddress) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
