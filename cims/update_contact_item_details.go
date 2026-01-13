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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateContactItemDetails Details for updating the support request activity.
type UpdateContactItemDetails struct {

	// Email id of the additional contact to be added to the support request.
	Contact *string `mandatory:"false" json:"contact"`

	// The type of activity occurring.
	// `NOTES` is the activity associated to attachments.
	// `PROBLEM_DESCRIPTION` is the activity associated to customer problem description.
	// `UPDATE` is the activity associated to adding comments.
	// `CLOSE` is the activity associated to closing the support request.
	// `REOPEN` is the activity associated to reopening the support request.
	// `ADD_CONTACT` is the activity associated to adding additional contact to the support request.
	ActivityType UpdateContactItemDetailsActivityTypeEnum `mandatory:"false" json:"activityType,omitempty"`
}

func (m UpdateContactItemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateContactItemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateContactItemDetailsActivityTypeEnum(string(m.ActivityType)); !ok && m.ActivityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActivityType: %s. Supported values are: %s.", m.ActivityType, strings.Join(GetUpdateContactItemDetailsActivityTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateContactItemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateContactItemDetails UpdateContactItemDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateContactItemDetails
	}{
		"contact",
		(MarshalTypeUpdateContactItemDetails)(m),
	}

	return json.Marshal(&s)
}

// UpdateContactItemDetailsActivityTypeEnum Enum with underlying type: string
type UpdateContactItemDetailsActivityTypeEnum string

// Set of constants representing the allowable values for UpdateContactItemDetailsActivityTypeEnum
const (
	UpdateContactItemDetailsActivityTypeNotes              UpdateContactItemDetailsActivityTypeEnum = "NOTES"
	UpdateContactItemDetailsActivityTypeProblemDescription UpdateContactItemDetailsActivityTypeEnum = "PROBLEM_DESCRIPTION"
	UpdateContactItemDetailsActivityTypeUpdate             UpdateContactItemDetailsActivityTypeEnum = "UPDATE"
	UpdateContactItemDetailsActivityTypeClose              UpdateContactItemDetailsActivityTypeEnum = "CLOSE"
	UpdateContactItemDetailsActivityTypeReopen             UpdateContactItemDetailsActivityTypeEnum = "REOPEN"
	UpdateContactItemDetailsActivityTypeAddContact         UpdateContactItemDetailsActivityTypeEnum = "ADD_CONTACT"
)

var mappingUpdateContactItemDetailsActivityTypeEnum = map[string]UpdateContactItemDetailsActivityTypeEnum{
	"NOTES":               UpdateContactItemDetailsActivityTypeNotes,
	"PROBLEM_DESCRIPTION": UpdateContactItemDetailsActivityTypeProblemDescription,
	"UPDATE":              UpdateContactItemDetailsActivityTypeUpdate,
	"CLOSE":               UpdateContactItemDetailsActivityTypeClose,
	"REOPEN":              UpdateContactItemDetailsActivityTypeReopen,
	"ADD_CONTACT":         UpdateContactItemDetailsActivityTypeAddContact,
}

var mappingUpdateContactItemDetailsActivityTypeEnumLowerCase = map[string]UpdateContactItemDetailsActivityTypeEnum{
	"notes":               UpdateContactItemDetailsActivityTypeNotes,
	"problem_description": UpdateContactItemDetailsActivityTypeProblemDescription,
	"update":              UpdateContactItemDetailsActivityTypeUpdate,
	"close":               UpdateContactItemDetailsActivityTypeClose,
	"reopen":              UpdateContactItemDetailsActivityTypeReopen,
	"add_contact":         UpdateContactItemDetailsActivityTypeAddContact,
}

// GetUpdateContactItemDetailsActivityTypeEnumValues Enumerates the set of values for UpdateContactItemDetailsActivityTypeEnum
func GetUpdateContactItemDetailsActivityTypeEnumValues() []UpdateContactItemDetailsActivityTypeEnum {
	values := make([]UpdateContactItemDetailsActivityTypeEnum, 0)
	for _, v := range mappingUpdateContactItemDetailsActivityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateContactItemDetailsActivityTypeEnumStringValues Enumerates the set of values in String for UpdateContactItemDetailsActivityTypeEnum
func GetUpdateContactItemDetailsActivityTypeEnumStringValues() []string {
	return []string{
		"NOTES",
		"PROBLEM_DESCRIPTION",
		"UPDATE",
		"CLOSE",
		"REOPEN",
		"ADD_CONTACT",
	}
}

// GetMappingUpdateContactItemDetailsActivityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateContactItemDetailsActivityTypeEnum(val string) (UpdateContactItemDetailsActivityTypeEnum, bool) {
	enum, ok := mappingUpdateContactItemDetailsActivityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
