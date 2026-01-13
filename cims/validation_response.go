// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests.
// For more information, see Getting Help and Contacting Support (https://docs.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm).
// **Note**: Before you can create service requests with this API,
// complete user registration at My Oracle Cloud Support
// and then ask your tenancy administrator to provide you authorization for the related user groups.
//

package cims

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ValidationResponse The validation response returned when checking whether the requested user is valid.
type ValidationResponse struct {

	// Boolean value that indicates whether the requested user is valid.
	IsValidUser *bool `mandatory:"false" json:"isValidUser"`

	// Array of user groups that the user has access to for creating technical support type (`TECH`) support requests.
	// Each user group is indicated by its identifier and name (`userGroupId` and `userGroupName`).
	// Note: The Customer User Administrator (CUA) can manage user groups by name using
	// My Oracle Cloud Support portal (https://support.oracle.com).
	WritePermittedUserGroupInfos []CmosUserGroupInfo `mandatory:"false" json:"writePermittedUserGroupInfos"`
}

func (m ValidationResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ValidationResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
