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

// AccountItem Details about the AccountItem object.
type AccountItem struct {

	// Unique identifier for the item.
	ItemKey *string `mandatory:"true" json:"itemKey"`

	// The display name of the item. Avoid entering confidential information.
	Name *string `mandatory:"false" json:"name"`

	Category *Category `mandatory:"false" json:"category"`

	SubCategory *SubCategory `mandatory:"false" json:"subCategory"`

	IssueType *IssueType `mandatory:"false" json:"issueType"`
}

// GetItemKey returns ItemKey
func (m AccountItem) GetItemKey() *string {
	return m.ItemKey
}

// GetName returns Name
func (m AccountItem) GetName() *string {
	return m.Name
}

// GetCategory returns Category
func (m AccountItem) GetCategory() *Category {
	return m.Category
}

// GetSubCategory returns SubCategory
func (m AccountItem) GetSubCategory() *SubCategory {
	return m.SubCategory
}

// GetIssueType returns IssueType
func (m AccountItem) GetIssueType() *IssueType {
	return m.IssueType
}

func (m AccountItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AccountItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AccountItem) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAccountItem AccountItem
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAccountItem
	}{
		"account",
		(MarshalTypeAccountItem)(m),
	}

	return json.Marshal(&s)
}
