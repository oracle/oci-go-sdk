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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Dashboard The base schema for a dashboard.
// Derived schemas have configurations and widgets specific to the  `schemaVersion`.
type Dashboard interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dashboard resource.
	GetId() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dashboard group that the dashboard belongs to.
	GetDashboardGroupId() *string

	// A user-friendly name for the dashboard. Does not have to be unique, and it can be changed. Avoid entering confidential information.
	// Leading and trailing spaces and the following special characters are not allowed: <>()=/'"&\
	GetDisplayName() *string

	// A short description of the dashboard. It can be changed. Avoid entering confidential information.
	// The following special characters are not allowed: <>()=/'"&\
	GetDescription() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the dashboard. A dashboard is always in the same compartment as its dashboard group.
	GetCompartmentId() *string

	// The date and time the dashboard was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	//  Example: `2016-08-25T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	// The date and time the dashboard was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	GetTimeUpdated() *common.SDKTime

	// The current state of the dashboard.
	GetLifecycleState() DashboardLifecycleStateEnum

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type dashboard struct {
	JsonData         []byte
	Id               *string                           `mandatory:"true" json:"id"`
	DashboardGroupId *string                           `mandatory:"true" json:"dashboardGroupId"`
	DisplayName      *string                           `mandatory:"true" json:"displayName"`
	Description      *string                           `mandatory:"true" json:"description"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	TimeCreated      *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated      *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	LifecycleState   DashboardLifecycleStateEnum       `mandatory:"true" json:"lifecycleState"`
	FreeformTags     map[string]string                 `mandatory:"true" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	SchemaVersion    string                            `json:"schemaVersion"`
}

// UnmarshalJSON unmarshals json
func (m *dashboard) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdashboard dashboard
	s := struct {
		Model Unmarshalerdashboard
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DashboardGroupId = s.Model.DashboardGroupId
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.CompartmentId = s.Model.CompartmentId
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.SchemaVersion = s.Model.SchemaVersion

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dashboard) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SchemaVersion {
	case "V1":
		mm := V1Dashboard{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m dashboard) GetId() *string {
	return m.Id
}

//GetDashboardGroupId returns DashboardGroupId
func (m dashboard) GetDashboardGroupId() *string {
	return m.DashboardGroupId
}

//GetDisplayName returns DisplayName
func (m dashboard) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m dashboard) GetDescription() *string {
	return m.Description
}

//GetCompartmentId returns CompartmentId
func (m dashboard) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m dashboard) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m dashboard) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m dashboard) GetLifecycleState() DashboardLifecycleStateEnum {
	return m.LifecycleState
}

//GetFreeformTags returns FreeformTags
func (m dashboard) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m dashboard) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m dashboard) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m dashboard) String() string {
	return common.PointerString(m)
}

// DashboardLifecycleStateEnum Enum with underlying type: string
type DashboardLifecycleStateEnum string

// Set of constants representing the allowable values for DashboardLifecycleStateEnum
const (
	DashboardLifecycleStateCreating DashboardLifecycleStateEnum = "CREATING"
	DashboardLifecycleStateUpdating DashboardLifecycleStateEnum = "UPDATING"
	DashboardLifecycleStateActive   DashboardLifecycleStateEnum = "ACTIVE"
	DashboardLifecycleStateDeleting DashboardLifecycleStateEnum = "DELETING"
	DashboardLifecycleStateDeleted  DashboardLifecycleStateEnum = "DELETED"
	DashboardLifecycleStateFailed   DashboardLifecycleStateEnum = "FAILED"
)

var mappingDashboardLifecycleState = map[string]DashboardLifecycleStateEnum{
	"CREATING": DashboardLifecycleStateCreating,
	"UPDATING": DashboardLifecycleStateUpdating,
	"ACTIVE":   DashboardLifecycleStateActive,
	"DELETING": DashboardLifecycleStateDeleting,
	"DELETED":  DashboardLifecycleStateDeleted,
	"FAILED":   DashboardLifecycleStateFailed,
}

// GetDashboardLifecycleStateEnumValues Enumerates the set of values for DashboardLifecycleStateEnum
func GetDashboardLifecycleStateEnumValues() []DashboardLifecycleStateEnum {
	values := make([]DashboardLifecycleStateEnum, 0)
	for _, v := range mappingDashboardLifecycleState {
		values = append(values, v)
	}
	return values
}

// DashboardSchemaVersionEnum Enum with underlying type: string
type DashboardSchemaVersionEnum string

// Set of constants representing the allowable values for DashboardSchemaVersionEnum
const (
	DashboardSchemaVersionV1 DashboardSchemaVersionEnum = "V1"
)

var mappingDashboardSchemaVersion = map[string]DashboardSchemaVersionEnum{
	"V1": DashboardSchemaVersionV1,
}

// GetDashboardSchemaVersionEnumValues Enumerates the set of values for DashboardSchemaVersionEnum
func GetDashboardSchemaVersionEnumValues() []DashboardSchemaVersionEnum {
	values := make([]DashboardSchemaVersionEnum, 0)
	for _, v := range mappingDashboardSchemaVersion {
		values = append(values, v)
	}
	return values
}
