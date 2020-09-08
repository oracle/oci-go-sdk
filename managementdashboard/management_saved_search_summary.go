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

// ManagementSavedSearchSummary Summary of properties of a saved search.
type ManagementSavedSearchSummary struct {

	// id for saved search.  Must be provided if OOB, otherwise must not be provided.
	Id *string `mandatory:"true" json:"id"`

	// Display name for saved search.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// String boolean ("true" or "false") to indicate Out Of the Box saved search.
	IsOobSavedSearch *bool `mandatory:"true" json:"isOobSavedSearch"`

	// Id for application (LA, APM, etc.) that owners this saved search.  Each owner has a unique Id.
	ProviderId *string `mandatory:"true" json:"providerId"`

	// The ocid of the compartment that owns the saved search.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Description.
	Description *string `mandatory:"true" json:"description"`

	// Json for internationalization.
	Nls *interface{} `mandatory:"true" json:"nls"`

	// How to show the saved search.
	Type SavedSearchTypesEnum `mandatory:"true" json:"type"`

	// Json to contain options for UI.
	UiConfig *interface{} `mandatory:"true" json:"uiConfig"`

	// Array of Json to contain options for source of data.
	DataConfig []interface{} `mandatory:"true" json:"dataConfig"`

	// Created by which user.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// Updated by which user.
	UpdatedBy *string `mandatory:"true" json:"updatedBy"`

	// Screenshot.
	ScreenImage *string `mandatory:"true" json:"screenImage"`

	// Version of the metadata.
	MetadataVersion *string `mandatory:"true" json:"metadataVersion"`

	// Template.
	WidgetTemplate *string `mandatory:"true" json:"widgetTemplate"`

	// View Model
	WidgetVM *string `mandatory:"true" json:"widgetVM"`

	// Current state of saved search.
	LifecycleState LifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ManagementSavedSearchSummary) String() string {
	return common.PointerString(m)
}
