// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests.
// For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm).
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

// Item Details about the item object.
type Item interface {

	// Unique identifier for the item.
	GetItemKey() *string

	// The display name of the item. Avoid entering confidential information.
	GetName() *string

	GetCategory() *Category

	GetSubCategory() *SubCategory

	GetIssueType() *IssueType
}

type item struct {
	JsonData    []byte
	ItemKey     *string      `mandatory:"false" json:"itemKey"`
	Name        *string      `mandatory:"false" json:"name"`
	Category    *Category    `mandatory:"false" json:"category"`
	SubCategory *SubCategory `mandatory:"false" json:"subCategory"`
	IssueType   *IssueType   `mandatory:"false" json:"issueType"`
	Type        string       `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *item) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleritem item
	s := struct {
		Model Unmarshaleritem
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ItemKey = s.Model.ItemKey
	m.Name = s.Model.Name
	m.Category = s.Model.Category
	m.SubCategory = s.Model.SubCategory
	m.IssueType = s.Model.IssueType
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *item) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "account":
		mm := AccountItem{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "limit":
		mm := LimitItem{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "tech":
		mm := TechSupportItem{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "activity":
		mm := ActivityItem{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Item: %s.", m.Type)
		return *m, nil
	}
}

// GetItemKey returns ItemKey
func (m item) GetItemKey() *string {
	return m.ItemKey
}

// GetName returns Name
func (m item) GetName() *string {
	return m.Name
}

// GetCategory returns Category
func (m item) GetCategory() *Category {
	return m.Category
}

// GetSubCategory returns SubCategory
func (m item) GetSubCategory() *SubCategory {
	return m.SubCategory
}

// GetIssueType returns IssueType
func (m item) GetIssueType() *IssueType {
	return m.IssueType
}

func (m item) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m item) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
