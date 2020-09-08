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

// ManagementDashboardForImportExportDetails Properties for a dashboard, including dashboard id, and saved searches, for import purposes.
type ManagementDashboardForImportExportDetails struct {

	// Dashboard Id. Must be providied if OOB, otherwise must not be provided.
	DashboardId *string `mandatory:"true" json:"dashboardId"`

	// Provider Id.
	ProviderId *string `mandatory:"true" json:"providerId"`

	// Provider name.
	ProviderName *string `mandatory:"true" json:"providerName"`

	// Provider version.
	ProviderVersion *string `mandatory:"true" json:"providerVersion"`

	// Dashboard tiles array
	Tiles []ManagementDashboardTileDetails `mandatory:"true" json:"tiles"`

	// Display name for dashboard.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Dashboard's description.
	Description *string `mandatory:"true" json:"description"`

	// The ocid of the compartment that owns the dashboard.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// String boolean ("true" or "false"). OOB (Out of the Box) dashboards are only provided by Oracle.  They cannot be modified by non-Oracle.
	IsOobDashboard *bool `mandatory:"true" json:"isOobDashboard"`

	// String boolean ("true" or "false").  When false, dashboard is not shown in dashboard home.
	IsShowInHome *bool `mandatory:"true" json:"isShowInHome"`

	// Version of the metadata.
	MetadataVersion *string `mandatory:"true" json:"metadataVersion"`

	// String boolean ("true" or "false").  When false, dashboard is not automatically refreshed in intervals.
	IsShowDescription *bool `mandatory:"true" json:"isShowDescription"`

	// screen image.
	ScreenImage *string `mandatory:"true" json:"screenImage"`

	// Json for internationalization.
	Nls *interface{} `mandatory:"true" json:"nls"`

	// Json to contain options for UI.
	UiConfig *interface{} `mandatory:"true" json:"uiConfig"`

	// Array of Json to contain options for source of data.
	DataConfig []interface{} `mandatory:"true" json:"dataConfig"`

	// NORMAL means single dashboard, SET means dashboard set.
	Type *string `mandatory:"true" json:"type"`

	// String boolean ("true" or "false").
	IsFavorite *bool `mandatory:"true" json:"isFavorite"`

	// Array of saved searches.
	SavedSearches []ManagementSavedSearchForImportDetails `mandatory:"true" json:"savedSearches"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ManagementDashboardForImportExportDetails) String() string {
	return common.PointerString(m)
}
