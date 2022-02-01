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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DashboardGroup The base schema for a dashboard group.
type DashboardGroup struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dashboard group.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name for the dashboard. Does not have to be unique, and it can be changed. Avoid entering confidential information.
	// Leading and trailing spaces and the following special characters are not allowed: <>()=/'"&\
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A short description of the dashboard group. It can be changed. Avoid entering confidential information.
	// The following special characters are not allowed: <>()=/'"&\
	Description *string `mandatory:"true" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the dashboard group.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the dashboard group was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the dashboard group was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the `DashboardGroup` resource.
	LifecycleState DashboardGroupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DashboardGroup) String() string {
	return common.PointerString(m)
}

// DashboardGroupLifecycleStateEnum Enum with underlying type: string
type DashboardGroupLifecycleStateEnum string

// Set of constants representing the allowable values for DashboardGroupLifecycleStateEnum
const (
	DashboardGroupLifecycleStateCreating DashboardGroupLifecycleStateEnum = "CREATING"
	DashboardGroupLifecycleStateUpdating DashboardGroupLifecycleStateEnum = "UPDATING"
	DashboardGroupLifecycleStateActive   DashboardGroupLifecycleStateEnum = "ACTIVE"
	DashboardGroupLifecycleStateDeleting DashboardGroupLifecycleStateEnum = "DELETING"
	DashboardGroupLifecycleStateDeleted  DashboardGroupLifecycleStateEnum = "DELETED"
	DashboardGroupLifecycleStateFailed   DashboardGroupLifecycleStateEnum = "FAILED"
)

var mappingDashboardGroupLifecycleState = map[string]DashboardGroupLifecycleStateEnum{
	"CREATING": DashboardGroupLifecycleStateCreating,
	"UPDATING": DashboardGroupLifecycleStateUpdating,
	"ACTIVE":   DashboardGroupLifecycleStateActive,
	"DELETING": DashboardGroupLifecycleStateDeleting,
	"DELETED":  DashboardGroupLifecycleStateDeleted,
	"FAILED":   DashboardGroupLifecycleStateFailed,
}

// GetDashboardGroupLifecycleStateEnumValues Enumerates the set of values for DashboardGroupLifecycleStateEnum
func GetDashboardGroupLifecycleStateEnumValues() []DashboardGroupLifecycleStateEnum {
	values := make([]DashboardGroupLifecycleStateEnum, 0)
	for _, v := range mappingDashboardGroupLifecycleState {
		values = append(values, v)
	}
	return values
}
