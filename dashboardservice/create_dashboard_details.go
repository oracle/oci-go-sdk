// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dashboards API
//
// Use the Oracle Cloud Infrastructure Dashboards service API to manage dashboards in the Console.
// Dashboards provide an organized and customizable view of resources and their metrics in the Console.
// For more information, see Dashboards (https://docs.cloud.oracle.com/Content/Dashboards/home.htm).
// **Important:** Resources for the Dashboards service are created in the tenacy's home region.
// Although it is possible to create dashboard and dashboard group resources in regions other than the home region,
// you won't be able to view those resources in the Console.
// Therefore, creating resources outside of the home region is not recommended.
//

package dashboardservice

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateDashboardDetails The base schema for creating a dashboard.
// Derived schemas have configurations and widgets specific to the  `schemaVersion`.
type CreateDashboardDetails interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dashboard group that the dashbaord is associated with.
	GetDashboardGroupId() *string

	// A user-friendly name for the dashboard. Does not have to be unique, and it can be changed. Avoid entering confidential information.
	// Leading and trailing spaces and the following special characters are not allowed: <>()=/'"&\
	GetDisplayName() *string

	// A short description of the dashboard. It can be changed. Avoid entering confidential information.
	// The following special characters are not allowed: <>()=/'"&\
	GetDescription() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createdashboarddetails struct {
	JsonData         []byte
	DashboardGroupId *string                           `mandatory:"true" json:"dashboardGroupId"`
	DisplayName      *string                           `mandatory:"false" json:"displayName"`
	Description      *string                           `mandatory:"false" json:"description"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SchemaVersion    string                            `json:"schemaVersion"`
}

// UnmarshalJSON unmarshals json
func (m *createdashboarddetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedashboarddetails createdashboarddetails
	s := struct {
		Model Unmarshalercreatedashboarddetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DashboardGroupId = s.Model.DashboardGroupId
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SchemaVersion = s.Model.SchemaVersion

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdashboarddetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SchemaVersion {
	case "V1":
		mm := CreateV1DashboardDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDashboardGroupId returns DashboardGroupId
func (m createdashboarddetails) GetDashboardGroupId() *string {
	return m.DashboardGroupId
}

//GetDisplayName returns DisplayName
func (m createdashboarddetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m createdashboarddetails) GetDescription() *string {
	return m.Description
}

//GetFreeformTags returns FreeformTags
func (m createdashboarddetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createdashboarddetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createdashboarddetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createdashboarddetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
