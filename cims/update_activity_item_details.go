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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateActivityItemDetails Details for updating the support ticket activity.
type UpdateActivityItemDetails struct {

	// Comments updated at the time that the activity occurs.
	Comments *string `mandatory:"false" json:"comments"`

	// The type of activity occurring.
	ActivityType UpdateActivityItemDetailsActivityTypeEnum `mandatory:"false" json:"activityType,omitempty"`
}

func (m UpdateActivityItemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateActivityItemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateActivityItemDetailsActivityTypeEnum(string(m.ActivityType)); !ok && m.ActivityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActivityType: %s. Supported values are: %s.", m.ActivityType, strings.Join(GetUpdateActivityItemDetailsActivityTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
	UpdateActivityItemDetailsActivityTypeReopen             UpdateActivityItemDetailsActivityTypeEnum = "REOPEN"
)

var mappingUpdateActivityItemDetailsActivityTypeEnum = map[string]UpdateActivityItemDetailsActivityTypeEnum{
	"NOTES":               UpdateActivityItemDetailsActivityTypeNotes,
	"PROBLEM_DESCRIPTION": UpdateActivityItemDetailsActivityTypeProblemDescription,
	"UPDATE":              UpdateActivityItemDetailsActivityTypeUpdate,
	"CLOSE":               UpdateActivityItemDetailsActivityTypeClose,
	"REOPEN":              UpdateActivityItemDetailsActivityTypeReopen,
}

var mappingUpdateActivityItemDetailsActivityTypeEnumLowerCase = map[string]UpdateActivityItemDetailsActivityTypeEnum{
	"notes":               UpdateActivityItemDetailsActivityTypeNotes,
	"problem_description": UpdateActivityItemDetailsActivityTypeProblemDescription,
	"update":              UpdateActivityItemDetailsActivityTypeUpdate,
	"close":               UpdateActivityItemDetailsActivityTypeClose,
	"reopen":              UpdateActivityItemDetailsActivityTypeReopen,
}

// GetUpdateActivityItemDetailsActivityTypeEnumValues Enumerates the set of values for UpdateActivityItemDetailsActivityTypeEnum
func GetUpdateActivityItemDetailsActivityTypeEnumValues() []UpdateActivityItemDetailsActivityTypeEnum {
	values := make([]UpdateActivityItemDetailsActivityTypeEnum, 0)
	for _, v := range mappingUpdateActivityItemDetailsActivityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateActivityItemDetailsActivityTypeEnumStringValues Enumerates the set of values in String for UpdateActivityItemDetailsActivityTypeEnum
func GetUpdateActivityItemDetailsActivityTypeEnumStringValues() []string {
	return []string{
		"NOTES",
		"PROBLEM_DESCRIPTION",
		"UPDATE",
		"CLOSE",
		"REOPEN",
	}
}

// GetMappingUpdateActivityItemDetailsActivityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateActivityItemDetailsActivityTypeEnum(val string) (UpdateActivityItemDetailsActivityTypeEnum, bool) {
	enum, ok := mappingUpdateActivityItemDetailsActivityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
