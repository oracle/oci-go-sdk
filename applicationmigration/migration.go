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
	"github.com/oracle/oci-go-sdk/v38/common"
)

// Migration The properties that define a migration. A migration represents the end-to-end workflow of moving an application from a source
// environment to Oracle Cloud Infrastructure. Each migration moves a single application to Oracle Cloud Infrastructure.
// For more information, see Manage Migrations (https://docs.cloud.oracle.com/iaas/application-migration/manage_migrations.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator.
// If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type Migration struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the migration.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the migration.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// User-friendly name of the migration.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the migration.
	Description *string `mandatory:"false" json:"description"`

	// The date and time at which the migration was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source with which this migration is associated.
	SourceId *string `mandatory:"false" json:"sourceId"`

	// Name of the application which is being migrated. This is the name of the application in the source environment.
	ApplicationName *string `mandatory:"false" json:"applicationName"`

	// The type of application being migrated.
	ApplicationType MigrationTypesEnum `mandatory:"false" json:"applicationType,omitempty"`

	// The pre-existing database type to be used in this migration. Currently, Application migration only supports Oracle Cloud
	// Infrastrure databases and this option is currently available only for `JAVA_CLOUD_SERVICE` and `WEBLOGIC_CLOUD_SERVICE` target instance types.
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

	// The current state of the migration.
	LifecycleState MigrationLifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Details about the current lifecycle state of the migration.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the overall migration process.
	MigrationState MigrationStatesEnum `mandatory:"false" json:"migrationState,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Migration) String() string {
	return common.PointerString(m)
}
