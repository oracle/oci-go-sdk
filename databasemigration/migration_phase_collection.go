// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// REST API for Zero Downtime Migration (Oracle Database Migration Service --ODMS-- as customer-facing service name)
//
// Provides users the ability to perform Zero Downtime migration operations
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v40/common"
)

// MigrationPhaseCollection Results of a Migration Phase search. Contains a collection of valid ODMS Job Phases.
type MigrationPhaseCollection struct {

	// Items in collection.
	Items []MigrationPhaseSummary `mandatory:"true" json:"items"`
}

func (m MigrationPhaseCollection) String() string {
	return common.PointerString(m)
}
