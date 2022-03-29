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
	"github.com/oracle/oci-go-sdk/v64/common"
	"strings"
)

// UpdateV1DashboardDetails Details for updating a version 1 dashboard.
// The interpretation of the `config` and `widgets` fields depends on the runtime behavior of the Oracle Cloud Infrastructure Console.
// The sum of the `config` and `widget` fields JSON text representation cannot exceed 200 KB.
type UpdateV1DashboardDetails struct {

	// A user-friendly name for the dashboard. Does not have to be unique, and it can be changed. Avoid entering confidential information.
	// Leading and trailing spaces and the following special characters are not allowed: <>()=/'"&\
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the dashboard. It can be changed. Avoid entering confidential information.
	// The following special characters are not allowed: <>()=/'"&\
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The layout and widget placement for the dashboard.
	Config *interface{} `mandatory:"false" json:"config"`

	// The basic visualization building blocks of a dashboard.
	Widgets []interface{} `mandatory:"false" json:"widgets"`
}

//GetDisplayName returns DisplayName
func (m UpdateV1DashboardDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m UpdateV1DashboardDetails) GetDescription() *string {
	return m.Description
}

//GetFreeformTags returns FreeformTags
func (m UpdateV1DashboardDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateV1DashboardDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateV1DashboardDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateV1DashboardDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateV1DashboardDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateV1DashboardDetails UpdateV1DashboardDetails
	s := struct {
		DiscriminatorParam string `json:"schemaVersion"`
		MarshalTypeUpdateV1DashboardDetails
	}{
		"V1",
		(MarshalTypeUpdateV1DashboardDetails)(m),
	}

	return json.Marshal(&s)
}
