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

// LimitItem Reserved for future use.
type LimitItem struct {

	// Unique identifier for the item.
	ItemKey *string `mandatory:"true" json:"itemKey"`

	// The display name of the item.
	Name *string `mandatory:"false" json:"name"`

	Category *Category `mandatory:"false" json:"category"`

	SubCategory *SubCategory `mandatory:"false" json:"subCategory"`

	IssueType *IssueType `mandatory:"false" json:"issueType"`

	// The currently available limit of the resource.
	CurrentLimit *int `mandatory:"false" json:"currentLimit"`

	// The current usage of the resource.
	CurrentUsage *int `mandatory:"false" json:"currentUsage"`

	// The requested limit for the resource.
	RequestedLimit *int `mandatory:"false" json:"requestedLimit"`

	// The status of the request.
	LimitStatus LimitItemLimitStatusEnum `mandatory:"false" json:"limitStatus,omitempty"`
}

//GetItemKey returns ItemKey
func (m LimitItem) GetItemKey() *string {
	return m.ItemKey
}

//GetName returns Name
func (m LimitItem) GetName() *string {
	return m.Name
}

//GetCategory returns Category
func (m LimitItem) GetCategory() *Category {
	return m.Category
}

//GetSubCategory returns SubCategory
func (m LimitItem) GetSubCategory() *SubCategory {
	return m.SubCategory
}

//GetIssueType returns IssueType
func (m LimitItem) GetIssueType() *IssueType {
	return m.IssueType
}

func (m LimitItem) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m LimitItem) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLimitItem LimitItem
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeLimitItem
	}{
		"limit",
		(MarshalTypeLimitItem)(m),
	}

	return json.Marshal(&s)
}

// LimitItemLimitStatusEnum Enum with underlying type: string
type LimitItemLimitStatusEnum string

// Set of constants representing the allowable values for LimitItemLimitStatusEnum
const (
	LimitItemLimitStatusApproved          LimitItemLimitStatusEnum = "APPROVED"
	LimitItemLimitStatusPartiallyApproved LimitItemLimitStatusEnum = "PARTIALLY_APPROVED"
	LimitItemLimitStatusNotApproved       LimitItemLimitStatusEnum = "NOT_APPROVED"
)

var mappingLimitItemLimitStatus = map[string]LimitItemLimitStatusEnum{
	"APPROVED":           LimitItemLimitStatusApproved,
	"PARTIALLY_APPROVED": LimitItemLimitStatusPartiallyApproved,
	"NOT_APPROVED":       LimitItemLimitStatusNotApproved,
}

// GetLimitItemLimitStatusEnumValues Enumerates the set of values for LimitItemLimitStatusEnum
func GetLimitItemLimitStatusEnumValues() []LimitItemLimitStatusEnum {
	values := make([]LimitItemLimitStatusEnum, 0)
	for _, v := range mappingLimitItemLimitStatus {
		values = append(values, v)
	}
	return values
}
