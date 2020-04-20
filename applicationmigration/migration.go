// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Migration An application being migrated from a source environment to OCI.
type Migration struct {

	// Unique identifier (OCID) for the application
	Id *string `mandatory:"false" json:"id"`

	// Unique idenfifier (OCID) for the compartment where the Source is located.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Human-readable name of the migration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the migration.
	Description *string `mandatory:"false" json:"description"`

	// The date and time at which the migration was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Unique identifier (OCID) of the application source.
	SourceId *string `mandatory:"false" json:"sourceId"`

	// Name of the application being migrated from the source.
	ApplicationName *string `mandatory:"false" json:"applicationName"`

	// The type of application being migrated.
	ApplicationType MigrationTypesEnum `mandatory:"false" json:"applicationType,omitempty"`

	// Configuration required to migrate the application. In addition to the key and value, additional fields are provided to describe type type and purpose of each field. Only the value for each key is required when passing configuration to the CreateMigration operation.
	ServiceConfig map[string]ConfigurationField `mandatory:"false" json:"serviceConfig"`

	// Configuration required to migrate the application. In addition to the key and value, additional fields are provided to describe type type and purpose of each field. Only the value for each key is required when passing configuration to the CreateMigration operation.
	ApplicationConfig map[string]ConfigurationField `mandatory:"false" json:"applicationConfig"`

	// The current state of the Migration
	LifecycleState MigrationLifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Details about the current lifecycle state
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the overall Migration process
	MigrationState MigrationStatesEnum `mandatory:"false" json:"migrationState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Migration) String() string {
	return common.PointerString(m)
}
