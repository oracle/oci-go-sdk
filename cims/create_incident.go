// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests.
// For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm).
// **Note**: Before you can create service requests with this API,
// you need to have an Oracle Single Sign On (SSO) account,
// and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateIncident Details gathered during the creation of the support ticket.
type CreateIncident struct {

	// The OCID of the tenancy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	Ticket *CreateTicketDetails `mandatory:"true" json:"ticket"`

	// The kind of support ticket (type of support request).
	// For information about `ACCOUNT` support tickets, see
	// Creating a Billing Support Request (https://docs.cloud.oracle.com/iaas/Content/GSG/support/create-incident-billing.htm).
	// For information about `LIMIT` support tickets, see
	// Creating a Service Limit Increase Request (https://docs.cloud.oracle.com/iaas/Content/GSG/support/create-incident-limit.htm).
	// For information about `TECH` support tickets, see
	// Creating a Technical Support Request (https://docs.cloud.oracle.com/iaas/Content/GSG/support/create-incident-technical.htm).
	ProblemType ProblemTypeEnum `mandatory:"true" json:"problemType"`

	// The Customer Support Identifier (CSI) number associated with the support account.
	// The CSI is required for technical support tickets and optional for limits and billing tickets.
	Csi *string `mandatory:"false" json:"csi"`

	// The list of contacts.
	Contacts []Contact `mandatory:"false" json:"contacts"`

	// The incident referrer. This value is often the URL that the customer used when creating the support ticket.
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
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
