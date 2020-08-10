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

// CreateItemDetails Details gathered during item creation.
// **Caution:** Avoid using any confidential information when you supply string values using the API.
type CreateItemDetails interface {
	GetCategory() *CreateCategoryDetails

	GetSubCategory() *CreateSubCategoryDetails

	GetIssueType() *CreateIssueTypeDetails

	// The display name of the item.
	GetName() *string
}

type createitemdetails struct {
	JsonData    []byte
	Category    *CreateCategoryDetails    `mandatory:"false" json:"category"`
	SubCategory *CreateSubCategoryDetails `mandatory:"false" json:"subCategory"`
	IssueType   *CreateIssueTypeDetails   `mandatory:"false" json:"issueType"`
	Name        *string                   `mandatory:"false" json:"name"`
	Type        string                    `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createitemdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateitemdetails createitemdetails
	s := struct {
		Model Unmarshalercreateitemdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Category = s.Model.Category
	m.SubCategory = s.Model.SubCategory
	m.IssueType = s.Model.IssueType
	m.Name = s.Model.Name
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createitemdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "tech":
		mm := CreateTechSupportItemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "limit":
		mm := CreateLimitItemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetCategory returns Category
func (m createitemdetails) GetCategory() *CreateCategoryDetails {
	return m.Category
}

//GetSubCategory returns SubCategory
func (m createitemdetails) GetSubCategory() *CreateSubCategoryDetails {
	return m.SubCategory
}

//GetIssueType returns IssueType
func (m createitemdetails) GetIssueType() *CreateIssueTypeDetails {
	return m.IssueType
}

//GetName returns Name
func (m createitemdetails) GetName() *string {
	return m.Name
}

func (m createitemdetails) String() string {
	return common.PointerString(m)
}
