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

// SavedSearchTypesEnum Enum with underlying type: string
type SavedSearchTypesEnum string

// Set of constants representing the allowable values for SavedSearchTypesEnum
const (
	SavedSearchTypesSearchShowInDashboard     SavedSearchTypesEnum = "SEARCH_SHOW_IN_DASHBOARD"
	SavedSearchTypesSearchDontShowInDashboard SavedSearchTypesEnum = "SEARCH_DONT_SHOW_IN_DASHBOARD"
	SavedSearchTypesWidgetShowInDashboard     SavedSearchTypesEnum = "WIDGET_SHOW_IN_DASHBOARD"
	SavedSearchTypesWidgetDontShowInDashboard SavedSearchTypesEnum = "WIDGET_DONT_SHOW_IN_DASHBOARD"
)

var mappingSavedSearchTypes = map[string]SavedSearchTypesEnum{
	"SEARCH_SHOW_IN_DASHBOARD":      SavedSearchTypesSearchShowInDashboard,
	"SEARCH_DONT_SHOW_IN_DASHBOARD": SavedSearchTypesSearchDontShowInDashboard,
	"WIDGET_SHOW_IN_DASHBOARD":      SavedSearchTypesWidgetShowInDashboard,
	"WIDGET_DONT_SHOW_IN_DASHBOARD": SavedSearchTypesWidgetDontShowInDashboard,
}

// GetSavedSearchTypesEnumValues Enumerates the set of values for SavedSearchTypesEnum
func GetSavedSearchTypesEnumValues() []SavedSearchTypesEnum {
	values := make([]SavedSearchTypesEnum, 0)
	for _, v := range mappingSavedSearchTypes {
		values = append(values, v)
	}
	return values
}
