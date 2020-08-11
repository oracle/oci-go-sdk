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

// CreateTechSupportItemDetails Details about the issue that the technical support request relates to.
// **Caution:** Avoid using any confidential information when you supply string values using the API.
type CreateTechSupportItemDetails struct {
	Category *CreateCategoryDetails `mandatory:"false" json:"category"`

	SubCategory *CreateSubCategoryDetails `mandatory:"false" json:"subCategory"`

	IssueType *CreateIssueTypeDetails `mandatory:"false" json:"issueType"`

	// The display name of the item.
	Name *string `mandatory:"false" json:"name"`
}

//GetCategory returns Category
func (m CreateTechSupportItemDetails) GetCategory() *CreateCategoryDetails {
	return m.Category
}

//GetSubCategory returns SubCategory
func (m CreateTechSupportItemDetails) GetSubCategory() *CreateSubCategoryDetails {
	return m.SubCategory
}

//GetIssueType returns IssueType
func (m CreateTechSupportItemDetails) GetIssueType() *CreateIssueTypeDetails {
	return m.IssueType
}

//GetName returns Name
func (m CreateTechSupportItemDetails) GetName() *string {
	return m.Name
}

func (m CreateTechSupportItemDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateTechSupportItemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateTechSupportItemDetails CreateTechSupportItemDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateTechSupportItemDetails
	}{
		"tech",
		(MarshalTypeCreateTechSupportItemDetails)(m),
	}

	return json.Marshal(&s)
}
