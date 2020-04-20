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

// SourceApplication An application running in the source environment that is available for export.
type SourceApplication struct {

	// The name of the application
	Name *string `mandatory:"false" json:"name"`

	// The type of application
	Type MigrationTypesEnum `mandatory:"false" json:"type,omitempty"`

	// Unique identifier (OCID) for the Source to which the application belongs
	SourceId *string `mandatory:"false" json:"sourceId"`

	// The version of the application server
	Version *string `mandatory:"false" json:"version"`

	// The current application running state
	State *string `mandatory:"false" json:"state"`
}

func (m SourceApplication) String() string {
	return common.PointerString(m)
}
