// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// User Details about the user.
type User struct {

	// A unique identifier for the user.
	Key *string `mandatory:"true" json:"key"`

	// The user's first name.
	FirstName *string `mandatory:"false" json:"firstName"`

	// The user's last name.
	LastName *string `mandatory:"false" json:"lastName"`

	// The country of the user.
	Country *string `mandatory:"false" json:"country"`

	// The CSI associated with the user.
	Csi *string `mandatory:"false" json:"csi"`

	// The user's contact phone number.
	Phone *string `mandatory:"false" json:"phone"`

	// The timezone of the user.
	Timezone *string `mandatory:"false" json:"timezone"`

	// The company that the user belongs to.
	OrganizationName *string `mandatory:"false" json:"organizationName"`

	// The OCID of the tenancy.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The email of the contact person.
	ContactEmail *string `mandatory:"false" json:"contactEmail"`

	// The kind of support request (type of support request).
	// For information about `ACCOUNT` support requests, see
	// Creating a Billing Support Request (https://docs.oracle.com/iaas/Content/GSG/support/create-incident-billing.htm).
	// For information about `LIMIT` support requests, see
	// Creating a Service Limit Increase Request (https://docs.oracle.com/iaas/Content/GSG/support/create-incident-limit.htm).
	// For information about `TECH` support requests, see
	// Creating a Technical Support Request (https://docs.oracle.com/iaas/Content/GSG/support/create-incident-technical.htm).
	ProblemType ProblemTypeEnum `mandatory:"false" json:"problemType,omitempty"`
}

func (m User) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m User) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingProblemTypeEnum(string(m.ProblemType)); !ok && m.ProblemType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProblemType: %s. Supported values are: %s.", m.ProblemType, strings.Join(GetProblemTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
