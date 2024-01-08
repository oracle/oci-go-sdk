// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// V1Dashboard A version 1 dashboard.
// The interpretation of the `config` and `widgets` fields depends on the runtime behavior of the Oracle Cloud Infrastructure Console.
// The sum of the `config` and `widget` fields JSON text representation cannot exceed 200 KB.
type V1Dashboard struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dashboard resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dashboard group that the dashboard belongs to.
	DashboardGroupId *string `mandatory:"true" json:"dashboardGroupId"`

	// A user-friendly name for the dashboard. Does not have to be unique, and it can be changed. Avoid entering confidential information.
	// Leading and trailing spaces and the following special characters are not allowed: <>()=/'"&\
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A short description of the dashboard. It can be changed. Avoid entering confidential information.
	// The following special characters are not allowed: <>()=/'"&\
	Description *string `mandatory:"true" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the dashboard. A dashboard is always in the same compartment as its dashboard group.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the dashboard was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	//  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the dashboard was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The visualization building blocks of the dashboard.
	Widgets []interface{} `mandatory:"true" json:"widgets"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The dashboard configuration. For example, the layout and widget placement.
	Config *interface{} `mandatory:"false" json:"config"`

	// The current state of the dashboard.
	LifecycleState DashboardLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m V1Dashboard) GetId() *string {
	return m.Id
}

// GetDashboardGroupId returns DashboardGroupId
func (m V1Dashboard) GetDashboardGroupId() *string {
	return m.DashboardGroupId
}

// GetDisplayName returns DisplayName
func (m V1Dashboard) GetDisplayName() *string {
	return m.DisplayName
}

// GetDescription returns Description
func (m V1Dashboard) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m V1Dashboard) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m V1Dashboard) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m V1Dashboard) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m V1Dashboard) GetLifecycleState() DashboardLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m V1Dashboard) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m V1Dashboard) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m V1Dashboard) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m V1Dashboard) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m V1Dashboard) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDashboardLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDashboardLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m V1Dashboard) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeV1Dashboard V1Dashboard
	s := struct {
		DiscriminatorParam string `json:"schemaVersion"`
		MarshalTypeV1Dashboard
	}{
		"V1",
		(MarshalTypeV1Dashboard)(m),
	}

	return json.Marshal(&s)
}
