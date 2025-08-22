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

// CreateIncident Details gathered during the creation of the support request.
type CreateIncident struct {

	// The OCID of the tenancy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	Ticket *CreateTicketDetails `mandatory:"true" json:"ticket"`

	// The kind of support request (type of support request).
	// For information about `ACCOUNT` support requests, see
	// Creating a Billing Support Request (https://docs.oracle.com/iaas/Content/GSG/support/create-incident-billing.htm).
	// For information about `LIMIT` support requests, see
	// Creating a Service Limit Increase Request (https://docs.oracle.com/iaas/Content/GSG/support/create-incident-limit.htm).
	// For information about `TECH` support requests, see
	// Creating a Technical Support Request (https://docs.oracle.com/iaas/Content/GSG/support/create-incident-technical.htm).
	ProblemType ProblemTypeEnum `mandatory:"true" json:"problemType"`

	// Deprecated. The Customer Support Identifier (CSI) number associated with the support account.
	// The CSI is optional for all support request types.
	Csi *string `mandatory:"false" json:"csi"`

	// Technical support type (`TECH`) only: Identifier of the user group to assign the new support request to.
	// To find identifiers of user groups that you have access to, run the
	// ValidateUser operation.
	// Note: The Customer User Administrator (CUA) can manage user groups by name using
	// My Oracle Cloud Support portal (https://support.oracle.com).
	UserGroupId *string `mandatory:"false" json:"userGroupId"`

	// The list of contacts.
	Contacts []Contact `mandatory:"false" json:"contacts"`

	// The incident referrer. This value is often the URL that the customer used when creating the support request.
	Referrer *string `mandatory:"false" json:"referrer"`
}

func (m CreateIncident) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateIncident) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProblemTypeEnum(string(m.ProblemType)); !ok && m.ProblemType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProblemType: %s. Supported values are: %s.", m.ProblemType, strings.Join(GetProblemTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
