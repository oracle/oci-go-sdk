// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateMigrationDetails While creating a migration, specify the source and the application that you want migrate.
// Each migration moves a single application from a specified source to a specified Oracle Cloud Infrastructure tenancy.
// If required, provide the credentials of the application administrator in the source environment.
// Application Migration uses this information to access the application, as well as discover application artifacts,
// such as the complete domain configuration along with data sources and other dependencies.
// You must also assign a name and provide a description for the migration.
// This helps you to identify the appropriate source environment when you have multiple sources defined.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateMigrationDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the source.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source.
	SourceId *string `mandatory:"true" json:"sourceId"`

	// Name of the application that you want to migrate from the source environment.
	ApplicationName *string `mandatory:"true" json:"applicationName"`

	DiscoveryDetails DiscoveryDetails `mandatory:"true" json:"discoveryDetails"`

	// User-friendly name of the application. This will be the name of the migrated application in Oracle Cloud Infrastructure.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the application that you are migrating.
	Description *string `mandatory:"false" json:"description"`

	// The pre-existing database type to be used in this migration. Currently, Application migration only supports Oracle Cloud
	// Infrastructure databases and this option is currently available only for `JAVA_CLOUD_SERVICE` and `WEBLOGIC_CLOUD_SERVICE` target instance types.
	PreCreatedTargetDatabaseType TargetDatabaseTypesEnum `mandatory:"false" json:"preCreatedTargetDatabaseType,omitempty"`

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

func (m CreateMigrationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMigrationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTargetDatabaseTypesEnum(string(m.PreCreatedTargetDatabaseType)); !ok && m.PreCreatedTargetDatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreCreatedTargetDatabaseType: %s. Supported values are: %s.", m.PreCreatedTargetDatabaseType, strings.Join(GetTargetDatabaseTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateMigrationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                  *string                           `json:"displayName"`
		Description                  *string                           `json:"description"`
		PreCreatedTargetDatabaseType TargetDatabaseTypesEnum           `json:"preCreatedTargetDatabaseType"`
		IsSelectiveMigration         *bool                             `json:"isSelectiveMigration"`
		ServiceConfig                map[string]ConfigurationField     `json:"serviceConfig"`
		ApplicationConfig            map[string]ConfigurationField     `json:"applicationConfig"`
		FreeformTags                 map[string]string                 `json:"freeformTags"`
		DefinedTags                  map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId                *string                           `json:"compartmentId"`
		SourceId                     *string                           `json:"sourceId"`
		ApplicationName              *string                           `json:"applicationName"`
		DiscoveryDetails             discoverydetails                  `json:"discoveryDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.PreCreatedTargetDatabaseType = model.PreCreatedTargetDatabaseType

	m.IsSelectiveMigration = model.IsSelectiveMigration

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
