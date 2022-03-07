// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// CreateUserDetails Details about creation of user.
type CreateUserDetails struct {

	// The OCID of the tenancy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// First name of the user.
	FirstName *string `mandatory:"true" json:"firstName"`

	// Last name of the user.
	LastName *string `mandatory:"true" json:"lastName"`

	// Country of the user.
	Country *string `mandatory:"true" json:"country"`

	// CSI to be associated to the user.
	Csi *string `mandatory:"true" json:"csi"`

	// Contact number of the user.
	Phone *string `mandatory:"true" json:"phone"`

	// Timezone of the user.
	Timezone *string `mandatory:"true" json:"timezone"`

	// Organization of the user.
	OrganizationName *string `mandatory:"true" json:"organizationName"`
}

func (m CreateUserDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateUserDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
