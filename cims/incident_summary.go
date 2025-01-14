// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests.
// For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm).
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

// IncidentSummary Details about the support ticket.
type IncidentSummary struct {

	// Unique identifier of the incident.
	Key *string `mandatory:"true" json:"key"`

	// The kind of support ticket (type of support request).
	// For information about `ACCOUNT` support tickets, see
	// Creating a Billing Support Request (https://docs.cloud.oracle.com/iaas/Content/GSG/support/create-incident-billing.htm).
	// For information about `LIMIT` support tickets, see
	// Creating a Service Limit Increase Request (https://docs.cloud.oracle.com/iaas/Content/GSG/support/create-incident-limit.htm).
	// For information about `TECH` support tickets, see
	// Creating a Technical Support Request (https://docs.cloud.oracle.com/iaas/Content/GSG/support/create-incident-technical.htm).
	ProblemType ProblemTypeEnum `mandatory:"true" json:"problemType"`

	// The OCID of the tenancy.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	ContactList *ContactList `mandatory:"false" json:"contactList"`

	TenancyInformation *TenancyInformation `mandatory:"false" json:"tenancyInformation"`

	Ticket *Ticket `mandatory:"false" json:"ticket"`

	IncidentType *IncidentResourceType `mandatory:"false" json:"incidentType"`

	// Technical support type (`TECH`) only: The identifier of the support request's user group in My Oracle Cloud Support portal.
	UserGroupId *string `mandatory:"false" json:"userGroupId"`

	// Technical support type (`TECH`) only: Name of the support request's user group in My Oracle Cloud Support portal.
	UserGroupName *string `mandatory:"false" json:"userGroupName"`

	// Technical support type (`TECH`) only: The identifier of the support request's primary contact (`primaryContactPartyName`) in My Oracle Cloud Support portal.
	PrimaryContactPartyId *string `mandatory:"false" json:"primaryContactPartyId"`

	// Technical support type (`TECH`) only: The name of the support request's primary contact in My Oracle Cloud Support portal.
	PrimaryContactPartyName *string `mandatory:"false" json:"primaryContactPartyName"`

	// Technical support type (`TECH`) only: Allows update of the support request in My Oracle Cloud Support portal,
	// when the user has write permission to the support request's user group.
	IsWritePermitted *bool `mandatory:"false" json:"isWritePermitted"`

	// Technical support type (`TECH`) only: Message indicating the user group (`userGroupId`) that was auto-selected for a new support request. This message appears when no user group was specified in the create request for a new technical support request.
	WarnMessage *string `mandatory:"false" json:"warnMessage"`
}

func (m IncidentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IncidentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProblemTypeEnum(string(m.ProblemType)); !ok && m.ProblemType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProblemType: %s. Supported values are: %s.", m.ProblemType, strings.Join(GetProblemTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
