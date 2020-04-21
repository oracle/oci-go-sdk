// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateTicketDetails Details of Ticket created
type CreateTicketDetails struct {

	// Severity of the ticket. eg: HIGH, MEDIUM
	Severity CreateTicketDetailsSeverityEnum `mandatory:"true" json:"severity"`

	// Title of ticket
	Title *string `mandatory:"true" json:"title"`

	// Details of ticket
	Description *string `mandatory:"true" json:"description"`

	// List of resources
	ResourceList []CreateResourceDetails `mandatory:"false" json:"resourceList"`
}

func (m CreateTicketDetails) String() string {
	return common.PointerString(m)
}

// CreateTicketDetailsSeverityEnum Enum with underlying type: string
type CreateTicketDetailsSeverityEnum string

// Set of constants representing the allowable values for CreateTicketDetailsSeverityEnum
const (
	CreateTicketDetailsSeverityHighest CreateTicketDetailsSeverityEnum = "HIGHEST"
	CreateTicketDetailsSeverityHigh    CreateTicketDetailsSeverityEnum = "HIGH"
	CreateTicketDetailsSeverityMedium  CreateTicketDetailsSeverityEnum = "MEDIUM"
)

var mappingCreateTicketDetailsSeverity = map[string]CreateTicketDetailsSeverityEnum{
	"HIGHEST": CreateTicketDetailsSeverityHighest,
	"HIGH":    CreateTicketDetailsSeverityHigh,
	"MEDIUM":  CreateTicketDetailsSeverityMedium,
}

// GetCreateTicketDetailsSeverityEnumValues Enumerates the set of values for CreateTicketDetailsSeverityEnum
func GetCreateTicketDetailsSeverityEnumValues() []CreateTicketDetailsSeverityEnum {
	values := make([]CreateTicketDetailsSeverityEnum, 0)
	for _, v := range mappingCreateTicketDetailsSeverity {
		values = append(values, v)
	}
	return values
}
