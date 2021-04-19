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
	"github.com/oracle/oci-go-sdk/v40/common"
)

// SourceApplicationSummary The properties that define an application, that is running in the source environment and which can be migrated to Oracle
// Cloud Infrastructure.
type SourceApplicationSummary struct {

	// The name of the application.
	Name *string `mandatory:"false" json:"name"`

	// The type of the application.
	Type MigrationTypesEnum `mandatory:"false" json:"type,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source to which the application belongs.
	SourceId *string `mandatory:"false" json:"sourceId"`

	// The version of the application.
	Version *string `mandatory:"false" json:"version"`

	// The current state of the application.
	State *string `mandatory:"false" json:"state"`
}

func (m SourceApplicationSummary) String() string {
	return common.PointerString(m)
}
