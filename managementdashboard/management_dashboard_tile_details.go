// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ManagementDashboard API
//
// Management Dashboard micro-service provides a set of CRUD, import, export, and compartment related APIs (such as change compartment)   to support dashboard and saved search metadata preservation.  These APIs are mainly for client UIs, for various UI activities such as get list of all saved searches in a compartment, create a dashboard, open a saved search, etc.  Use export to retrieve  dashboards and their saved searches, then edit the Json if necessary (for example change compartmentIds), then import the result to  destination dashboard service.
// APIs validate all required properties to ensure properties are present and have correct type and values.
//
//

package managementdashboard

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ManagementDashboardTileDetails Properties of dashboard tile representing a saved search.
type ManagementDashboardTileDetails struct {

	// Display name for saved search.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Id of saved search.
	SavedSearchId *string `mandatory:"true" json:"savedSearchId"`

	// Row, Y position
	Row *int `mandatory:"true" json:"row"`

	// Column, X position
	Column *int `mandatory:"true" json:"column"`

	// Height position
	Height *int `mandatory:"true" json:"height"`

	// Width position
	Width *int `mandatory:"true" json:"width"`

	// Json for internationalization.
	Nls *interface{} `mandatory:"true" json:"nls"`

	// Json to contain options for UI.
	UiConfig *interface{} `mandatory:"true" json:"uiConfig"`

	// Array of Json to contain options for source of data.
	DataConfig []interface{} `mandatory:"true" json:"dataConfig"`

	// State of saved search.
	State ManagementDashboardTileDetailsStateEnum `mandatory:"true" json:"state"`

	// Drill down configuration
	DrilldownConfig *interface{} `mandatory:"true" json:"drilldownConfig"`
}

func (m ManagementDashboardTileDetails) String() string {
	return common.PointerString(m)
}

// ManagementDashboardTileDetailsStateEnum Enum with underlying type: string
type ManagementDashboardTileDetailsStateEnum string

// Set of constants representing the allowable values for ManagementDashboardTileDetailsStateEnum
const (
	ManagementDashboardTileDetailsStateDeleted      ManagementDashboardTileDetailsStateEnum = "DELETED"
	ManagementDashboardTileDetailsStateUnauthorized ManagementDashboardTileDetailsStateEnum = "UNAUTHORIZED"
	ManagementDashboardTileDetailsStateDefault      ManagementDashboardTileDetailsStateEnum = "DEFAULT"
)

var mappingManagementDashboardTileDetailsState = map[string]ManagementDashboardTileDetailsStateEnum{
	"DELETED":      ManagementDashboardTileDetailsStateDeleted,
	"UNAUTHORIZED": ManagementDashboardTileDetailsStateUnauthorized,
	"DEFAULT":      ManagementDashboardTileDetailsStateDefault,
}

// GetManagementDashboardTileDetailsStateEnumValues Enumerates the set of values for ManagementDashboardTileDetailsStateEnum
func GetManagementDashboardTileDetailsStateEnumValues() []ManagementDashboardTileDetailsStateEnum {
	values := make([]ManagementDashboardTileDetailsStateEnum, 0)
	for _, v := range mappingManagementDashboardTileDetailsState {
		values = append(values, v)
	}
	return values
}
