// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration API
//
// Application Migration simplifies the migration of applications from Oracle Cloud Infrastructure Classic to Oracle Cloud Infrastructure.
// You can use Application Migration API to migrate applications, such as Oracle Java Cloud Service, SOA Cloud Service, and Integration Classic
// instances, to Oracle Cloud Infrastructure. For more information, see
// Overview of Application Migration (https://docs.cloud.oracle.com/iaas/application-migration/appmigrationoverview.htm).
//

package applicationmigration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// UpdateMigrationDetails Provide configuration information about the application in the target environment. Application Migration migrates the
// application to the target environment only after you provide this information. The information that you must provide varies
// depending on the type of application that you are migrating.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateMigrationDetails struct {

	// User-friendly name of the migration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the migration.
	Description *string `mandatory:"false" json:"description"`

	DiscoveryDetails DiscoveryDetails `mandatory:"false" json:"discoveryDetails"`

	// If set to `true`, Application Migration migrates the application resources selectively depending on the source.
	IsSelectiveMigration *bool `mandatory:"false" json:"isSelectiveMigration"`

	// Configuration required to migrate the application. In addition to the key and value, additional fields are provided
	// to describe type type and purpose of each field. Only the value for each key is required when passing configuration to the
	// CreateMigration operation.
	ServiceConfig map[string]ConfigurationField `mandatory:"false" json:"serviceConfig"`

	// Configuration required to migrate the application. In addition to the key and value, additional fields are provided
	// to describe type type and purpose of each field. Only the value for each key is required when passing configuration to the
	// CreateMigration operation.
	ApplicationConfig map[string]ConfigurationField `mandatory:"false" json:"applicationConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateMigrationDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateMigrationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string                           `json:"displayName"`
		Description          *string                           `json:"description"`
		DiscoveryDetails     discoverydetails                  `json:"discoveryDetails"`
		IsSelectiveMigration *bool                             `json:"isSelectiveMigration"`
		ServiceConfig        map[string]ConfigurationField     `json:"serviceConfig"`
		ApplicationConfig    map[string]ConfigurationField     `json:"applicationConfig"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	nn, e = model.DiscoveryDetails.UnmarshalPolymorphicJSON(model.DiscoveryDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DiscoveryDetails = nn.(DiscoveryDetails)
	} else {
		m.DiscoveryDetails = nil
	}

	m.IsSelectiveMigration = model.IsSelectiveMigration

	m.ServiceConfig = model.ServiceConfig

	m.ApplicationConfig = model.ApplicationConfig

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
