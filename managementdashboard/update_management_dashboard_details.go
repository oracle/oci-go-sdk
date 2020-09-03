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

// UpdateManagementDashboardDetails Properties for a dashboard.  Dashboard id must not be provided.
type UpdateManagementDashboardDetails struct {

	// Provider Id.
	ProviderId *string `mandatory:"false" json:"providerId"`

	// Provider name.
	ProviderName *string `mandatory:"false" json:"providerName"`

	// Provider version.
	ProviderVersion *string `mandatory:"false" json:"providerVersion"`

	// Dashboard tiles array.
	Tiles []ManagementDashboardTileDetails `mandatory:"false" json:"tiles"`

	// Display name for dashboard.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Dashboard's description.
	Description *string `mandatory:"false" json:"description"`

	// The ocid of the compartment that owns the dashboard.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// String boolean ("true" or "false").  OOB (Out of the Box) dashboards are only provided by Oracle.  They cannot be modified by non-Oracle.
	IsOobDashboard *bool `mandatory:"false" json:"isOobDashboard"`

	// String boolean ("true" or "false").  When false, dashboard is not shown in dashboard home.
	IsShowInHome *bool `mandatory:"false" json:"isShowInHome"`

	// Version of the metadata.
	MetadataVersion *string `mandatory:"false" json:"metadataVersion"`

	// String boolean ("true" or "false").  Whether to show the dashboard description.
	IsShowDescription *bool `mandatory:"false" json:"isShowDescription"`

	// Screen image.
	ScreenImage *string `mandatory:"false" json:"screenImage"`

	// Json for internationalization.
	Nls *interface{} `mandatory:"false" json:"nls"`

	// Json to contain options for UI.
	UiConfig *interface{} `mandatory:"false" json:"uiConfig"`

	// Array of Json to contain options for source of data.
	DataConfig []interface{} `mandatory:"false" json:"dataConfig"`

	// NORMAL meaning single dashboard, or SET meaning dashboard set.
	Type *string `mandatory:"false" json:"type"`

	// String boolean ("true" or "false").
	IsFavorite *bool `mandatory:"false" json:"isFavorite"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateManagementDashboardDetails) String() string {
	return common.PointerString(m)
}
