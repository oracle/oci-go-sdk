// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateActivityItemDetails Details for udpating the support ticket activity.
// **Caution:** Avoid using any confidential information when you supply string values using the API.
type UpdateActivityItemDetails struct {

	// Comments updated at the time that the activity occurs.
	Comments *string `mandatory:"false" json:"comments"`

	// The type of activity occurring.
	ActivityType UpdateActivityItemDetailsActivityTypeEnum `mandatory:"false" json:"activityType,omitempty"`
}

func (m UpdateActivityItemDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateActivityItemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateActivityItemDetails UpdateActivityItemDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateActivityItemDetails
	}{
		"activity",
		(MarshalTypeUpdateActivityItemDetails)(m),
	}

	return json.Marshal(&s)
}

// UpdateActivityItemDetailsActivityTypeEnum Enum with underlying type: string
type UpdateActivityItemDetailsActivityTypeEnum string

// Set of constants representing the allowable values for UpdateActivityItemDetailsActivityTypeEnum
const (
	UpdateActivityItemDetailsActivityTypeNotes              UpdateActivityItemDetailsActivityTypeEnum = "NOTES"
	UpdateActivityItemDetailsActivityTypeProblemDescription UpdateActivityItemDetailsActivityTypeEnum = "PROBLEM_DESCRIPTION"
	UpdateActivityItemDetailsActivityTypeUpdate             UpdateActivityItemDetailsActivityTypeEnum = "UPDATE"
	UpdateActivityItemDetailsActivityTypeClose              UpdateActivityItemDetailsActivityTypeEnum = "CLOSE"
)

var mappingUpdateActivityItemDetailsActivityType = map[string]UpdateActivityItemDetailsActivityTypeEnum{
	"NOTES":               UpdateActivityItemDetailsActivityTypeNotes,
	"PROBLEM_DESCRIPTION": UpdateActivityItemDetailsActivityTypeProblemDescription,
	"UPDATE":              UpdateActivityItemDetailsActivityTypeUpdate,
	"CLOSE":               UpdateActivityItemDetailsActivityTypeClose,
}

// GetUpdateActivityItemDetailsActivityTypeEnumValues Enumerates the set of values for UpdateActivityItemDetailsActivityTypeEnum
func GetUpdateActivityItemDetailsActivityTypeEnumValues() []UpdateActivityItemDetailsActivityTypeEnum {
	values := make([]UpdateActivityItemDetailsActivityTypeEnum, 0)
	for _, v := range mappingUpdateActivityItemDetailsActivityType {
		values = append(values, v)
	}
	return values
}
