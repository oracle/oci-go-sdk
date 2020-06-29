// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateMigrationDetails An application being migrated from a source environment to OCI.
type CreateMigrationDetails struct {

	// Unique idenfifier (OCID) for the compartment where the Source is located.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Unique identifier (OCID) of the application source.
	SourceId *string `mandatory:"true" json:"sourceId"`

	// Name of the application being migrated from the source.
	ApplicationName *string `mandatory:"true" json:"applicationName"`

	DiscoveryDetails DiscoveryDetails `mandatory:"true" json:"discoveryDetails"`

	// Human-readable name of the application.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the application.
	Description *string `mandatory:"false" json:"description"`

	// Configuration required to migrate the application. In addition to the key and value, additional fields are provided to describe type type and purpose of each field. Only the value for each key is required when passing configuration to the CreateMigration operation.
	ServiceConfig map[string]ConfigurationField `mandatory:"false" json:"serviceConfig"`

	// Configuration required to migrate the application. In addition to the key and value, additional fields are provided to describe type type and purpose of each field. Only the value for each key is required when passing configuration to the CreateMigration operation.
	ApplicationConfig map[string]ConfigurationField `mandatory:"false" json:"applicationConfig"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateMigrationDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateMigrationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName       *string                           `json:"displayName"`
		Description       *string                           `json:"description"`
		ServiceConfig     map[string]ConfigurationField     `json:"serviceConfig"`
		ApplicationConfig map[string]ConfigurationField     `json:"applicationConfig"`
		FreeformTags      map[string]string                 `json:"freeformTags"`
		DefinedTags       map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId     *string                           `json:"compartmentId"`
		SourceId          *string                           `json:"sourceId"`
		ApplicationName   *string                           `json:"applicationName"`
		DiscoveryDetails  discoverydetails                  `json:"discoveryDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.ServiceConfig = model.ServiceConfig

	m.ApplicationConfig = model.ApplicationConfig

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.SourceId = model.SourceId

	m.ApplicationName = model.ApplicationName

	nn, e = model.DiscoveryDetails.UnmarshalPolymorphicJSON(model.DiscoveryDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DiscoveryDetails = nn.(DiscoveryDetails)
	} else {
		m.DiscoveryDetails = nil
	}

	return
}
