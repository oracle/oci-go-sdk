// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ActivityItem Details about the ActivityItem object.
type ActivityItem struct {

	// Unique identifier for the item.
	ItemKey *string `mandatory:"true" json:"itemKey"`

	// The display name of the item.
	Name *string `mandatory:"false" json:"name"`

	Category *Category `mandatory:"false" json:"category"`

	SubCategory *SubCategory `mandatory:"false" json:"subCategory"`

	IssueType *IssueType `mandatory:"false" json:"issueType"`

	// Comments added with the activity on the support ticket.
	Comments *string `mandatory:"false" json:"comments"`

	// The time when the activity was created, in milliseconds since epoch time.
	TimeCreated *int `mandatory:"false" json:"timeCreated"`

	// The time when the activity was updated, in milliseconds since epoch time.
	TimeUpdated *int `mandatory:"false" json:"timeUpdated"`

	// The type of activity occuring on the support ticket.
	ActivityType ActivityItemActivityTypeEnum `mandatory:"false" json:"activityType,omitempty"`

	// The person who updates the activity on the support ticket.
	ActivityAuthor ActivityItemActivityAuthorEnum `mandatory:"false" json:"activityAuthor,omitempty"`
}

//GetItemKey returns ItemKey
func (m ActivityItem) GetItemKey() *string {
	return m.ItemKey
}

//GetName returns Name
func (m ActivityItem) GetName() *string {
	return m.Name
}

//GetCategory returns Category
func (m ActivityItem) GetCategory() *Category {
	return m.Category
}

//GetSubCategory returns SubCategory
func (m ActivityItem) GetSubCategory() *SubCategory {
	return m.SubCategory
}

//GetIssueType returns IssueType
func (m ActivityItem) GetIssueType() *IssueType {
	return m.IssueType
}

func (m ActivityItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ActivityItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingActivityItemActivityTypeEnum(string(m.ActivityType)); !ok && m.ActivityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActivityType: %s. Supported values are: %s.", m.ActivityType, strings.Join(GetActivityItemActivityTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingActivityItemActivityAuthorEnum(string(m.ActivityAuthor)); !ok && m.ActivityAuthor != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActivityAuthor: %s. Supported values are: %s.", m.ActivityAuthor, strings.Join(GetActivityItemActivityAuthorEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ActivityItem) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeActivityItem ActivityItem
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeActivityItem
	}{
		"activity",
		(MarshalTypeActivityItem)(m),
	}

	return json.Marshal(&s)
}

// ActivityItemActivityTypeEnum Enum with underlying type: string
type ActivityItemActivityTypeEnum string

// Set of constants representing the allowable values for ActivityItemActivityTypeEnum
const (
	ActivityItemActivityTypeNotes              ActivityItemActivityTypeEnum = "NOTES"
	ActivityItemActivityTypeProblemDescription ActivityItemActivityTypeEnum = "PROBLEM_DESCRIPTION"
	ActivityItemActivityTypeUpdate             ActivityItemActivityTypeEnum = "UPDATE"
	ActivityItemActivityTypeClose              ActivityItemActivityTypeEnum = "CLOSE"
)

var mappingActivityItemActivityTypeEnum = map[string]ActivityItemActivityTypeEnum{
	"NOTES":               ActivityItemActivityTypeNotes,
	"PROBLEM_DESCRIPTION": ActivityItemActivityTypeProblemDescription,
	"UPDATE":              ActivityItemActivityTypeUpdate,
	"CLOSE":               ActivityItemActivityTypeClose,
}

var mappingActivityItemActivityTypeEnumLowerCase = map[string]ActivityItemActivityTypeEnum{
	"notes":               ActivityItemActivityTypeNotes,
	"problem_description": ActivityItemActivityTypeProblemDescription,
	"update":              ActivityItemActivityTypeUpdate,
	"close":               ActivityItemActivityTypeClose,
}

// GetActivityItemActivityTypeEnumValues Enumerates the set of values for ActivityItemActivityTypeEnum
func GetActivityItemActivityTypeEnumValues() []ActivityItemActivityTypeEnum {
	values := make([]ActivityItemActivityTypeEnum, 0)
	for _, v := range mappingActivityItemActivityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetActivityItemActivityTypeEnumStringValues Enumerates the set of values in String for ActivityItemActivityTypeEnum
func GetActivityItemActivityTypeEnumStringValues() []string {
	return []string{
		"NOTES",
		"PROBLEM_DESCRIPTION",
		"UPDATE",
		"CLOSE",
	}
}

// GetMappingActivityItemActivityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActivityItemActivityTypeEnum(val string) (ActivityItemActivityTypeEnum, bool) {
	enum, ok := mappingActivityItemActivityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ActivityItemActivityAuthorEnum Enum with underlying type: string
type ActivityItemActivityAuthorEnum string

// Set of constants representing the allowable values for ActivityItemActivityAuthorEnum
const (
	ActivityItemActivityAuthorCustomer ActivityItemActivityAuthorEnum = "CUSTOMER"
	ActivityItemActivityAuthorOracle   ActivityItemActivityAuthorEnum = "ORACLE"
)

var mappingActivityItemActivityAuthorEnum = map[string]ActivityItemActivityAuthorEnum{
	"CUSTOMER": ActivityItemActivityAuthorCustomer,
	"ORACLE":   ActivityItemActivityAuthorOracle,
}

var mappingActivityItemActivityAuthorEnumLowerCase = map[string]ActivityItemActivityAuthorEnum{
	"customer": ActivityItemActivityAuthorCustomer,
	"oracle":   ActivityItemActivityAuthorOracle,
}

// GetActivityItemActivityAuthorEnumValues Enumerates the set of values for ActivityItemActivityAuthorEnum
func GetActivityItemActivityAuthorEnumValues() []ActivityItemActivityAuthorEnum {
	values := make([]ActivityItemActivityAuthorEnum, 0)
	for _, v := range mappingActivityItemActivityAuthorEnum {
		values = append(values, v)
	}
	return values
}

// GetActivityItemActivityAuthorEnumStringValues Enumerates the set of values in String for ActivityItemActivityAuthorEnum
func GetActivityItemActivityAuthorEnumStringValues() []string {
	return []string{
		"CUSTOMER",
		"ORACLE",
	}
}

// GetMappingActivityItemActivityAuthorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActivityItemActivityAuthorEnum(val string) (ActivityItemActivityAuthorEnum, bool) {
	enum, ok := mappingActivityItemActivityAuthorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
