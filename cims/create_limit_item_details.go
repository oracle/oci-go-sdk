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

// CreateLimitItemDetails Reserved for future use.
type CreateLimitItemDetails struct {
	Category *CreateCategoryDetails `mandatory:"false" json:"category"`

	SubCategory *CreateSubCategoryDetails `mandatory:"false" json:"subCategory"`

	IssueType *CreateIssueTypeDetails `mandatory:"false" json:"issueType"`

	// The display name of the item.
	Name *string `mandatory:"false" json:"name"`

	// The limit of the resource currently available.
	CurrentLimit *int `mandatory:"false" json:"currentLimit"`

	// The current usage of the resource.
	CurrentUsage *int `mandatory:"false" json:"currentUsage"`

	// Reserved for future use.
	RequestedLimit *int `mandatory:"false" json:"requestedLimit"`

	// The current status of the request.
	LimitStatus CreateLimitItemDetailsLimitStatusEnum `mandatory:"false" json:"limitStatus,omitempty"`
}

//GetCategory returns Category
func (m CreateLimitItemDetails) GetCategory() *CreateCategoryDetails {
	return m.Category
}

//GetSubCategory returns SubCategory
func (m CreateLimitItemDetails) GetSubCategory() *CreateSubCategoryDetails {
	return m.SubCategory
}

//GetIssueType returns IssueType
func (m CreateLimitItemDetails) GetIssueType() *CreateIssueTypeDetails {
	return m.IssueType
}

//GetName returns Name
func (m CreateLimitItemDetails) GetName() *string {
	return m.Name
}

func (m CreateLimitItemDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateLimitItemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateLimitItemDetails CreateLimitItemDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateLimitItemDetails
	}{
		"limit",
		(MarshalTypeCreateLimitItemDetails)(m),
	}

	return json.Marshal(&s)
}

// CreateLimitItemDetailsLimitStatusEnum Enum with underlying type: string
type CreateLimitItemDetailsLimitStatusEnum string

// Set of constants representing the allowable values for CreateLimitItemDetailsLimitStatusEnum
const (
	CreateLimitItemDetailsLimitStatusApproved          CreateLimitItemDetailsLimitStatusEnum = "APPROVED"
	CreateLimitItemDetailsLimitStatusPartiallyApproved CreateLimitItemDetailsLimitStatusEnum = "PARTIALLY_APPROVED"
	CreateLimitItemDetailsLimitStatusNotApproved       CreateLimitItemDetailsLimitStatusEnum = "NOT_APPROVED"
)

var mappingCreateLimitItemDetailsLimitStatus = map[string]CreateLimitItemDetailsLimitStatusEnum{
	"APPROVED":           CreateLimitItemDetailsLimitStatusApproved,
	"PARTIALLY_APPROVED": CreateLimitItemDetailsLimitStatusPartiallyApproved,
	"NOT_APPROVED":       CreateLimitItemDetailsLimitStatusNotApproved,
}

// GetCreateLimitItemDetailsLimitStatusEnumValues Enumerates the set of values for CreateLimitItemDetailsLimitStatusEnum
func GetCreateLimitItemDetailsLimitStatusEnumValues() []CreateLimitItemDetailsLimitStatusEnum {
	values := make([]CreateLimitItemDetailsLimitStatusEnum, 0)
	for _, v := range mappingCreateLimitItemDetailsLimitStatus {
		values = append(values, v)
	}
	return values
}
