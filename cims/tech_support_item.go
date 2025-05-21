// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// TechSupportItem Details about the TechSupportItem object.
type TechSupportItem struct {

	// Unique identifier for the item.
	ItemKey *string `mandatory:"false" json:"itemKey"`

	// The display name of the item. Avoid entering confidential information.
	Name *string `mandatory:"false" json:"name"`

	Category *Category `mandatory:"false" json:"category"`

	SubCategory *SubCategory `mandatory:"false" json:"subCategory"`

	IssueType *IssueType `mandatory:"false" json:"issueType"`
}

// GetItemKey returns ItemKey
func (m TechSupportItem) GetItemKey() *string {
	return m.ItemKey
}

// GetName returns Name
func (m TechSupportItem) GetName() *string {
	return m.Name
}

// GetCategory returns Category
func (m TechSupportItem) GetCategory() *Category {
	return m.Category
}

// GetSubCategory returns SubCategory
func (m TechSupportItem) GetSubCategory() *SubCategory {
	return m.SubCategory
}

// GetIssueType returns IssueType
func (m TechSupportItem) GetIssueType() *IssueType {
	return m.IssueType
}

func (m TechSupportItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TechSupportItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TechSupportItem) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTechSupportItem TechSupportItem
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTechSupportItem
	}{
		"tech",
		(MarshalTypeTechSupportItem)(m),
	}

	return json.Marshal(&s)
}
