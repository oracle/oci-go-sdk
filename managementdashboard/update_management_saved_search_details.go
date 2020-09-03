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

// UpdateManagementSavedSearchDetails Properties of a saved search.  Saved search id must not be provided.
type UpdateManagementSavedSearchDetails struct {

	// Display name for saved search.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Id for application (LA, APM, etc.) that owners this saved search.  Each owner has a unique Id.
	ProviderId *string `mandatory:"false" json:"providerId"`

	// Version.
	ProviderVersion *string `mandatory:"false" json:"providerVersion"`

	// Name for application (LA, APM, etc.) that owners this saved search.
	ProviderName *string `mandatory:"false" json:"providerName"`

	// The ocid of the compartment that owns the saved search.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// String boolean ("true" or "false") to indicate Out Of the Box saved search.
	IsOobSavedSearch *bool `mandatory:"false" json:"isOobSavedSearch"`

	// Description.
	Description *string `mandatory:"false" json:"description"`

	// Json for internationalization.
	Nls *interface{} `mandatory:"false" json:"nls"`

	// How to show the saved search.
	Type SavedSearchTypesEnum `mandatory:"false" json:"type,omitempty"`

	// Json to contain options for UI.
	UiConfig *interface{} `mandatory:"false" json:"uiConfig"`

	// Array of Json to contain options for source of data.
	DataConfig []interface{} `mandatory:"false" json:"dataConfig"`

	// Screenshot.
	ScreenImage *string `mandatory:"false" json:"screenImage"`

	// Version of the metadata.
	MetadataVersion *string `mandatory:"false" json:"metadataVersion"`

	// Template.
	WidgetTemplate *string `mandatory:"false" json:"widgetTemplate"`

	// View Model
	WidgetVM *string `mandatory:"false" json:"widgetVM"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateManagementSavedSearchDetails) String() string {
	return common.PointerString(m)
}
