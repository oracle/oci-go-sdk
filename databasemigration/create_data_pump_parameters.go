// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// REST API for Zero Downtime Migration (Oracle Database Migration Service --ODMS-- as customer-facing service name)
//
// Provides users the ability to perform Zero Downtime migration operations
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v39/common"
)

// CreateDataPumpParameters Optional parameters for Datapump Export and Import. Refer to https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/ODMS_DATAPUMP.html#GUID-62324358-2F26-4A94-B69F-1075D53FA96D__BABDECJE
type CreateDataPumpParameters struct {

	// False to force datapump worker process to run on one instance.
	IsCluster *bool `mandatory:"false" json:"isCluster"`

	// Estimate size of dumps that will be generated.
	Estimate DataPumpEstimateEnum `mandatory:"false" json:"estimate,omitempty"`

	// IMPORT: Specifies the action to be performed when data is loaded into a preexisting table.
	TableExistsAction DataPumpTableExistsActionEnum `mandatory:"false" json:"tableExistsAction,omitempty"`

	// Exclude paratemers for export and import.
	ExcludeParameters []DataPumpExcludeParametersEnum `mandatory:"false" json:"excludeParameters"`

	// Maximum number of worker processes that can be used for a Datapump Import job.
	// For an Autonomous Database, ODMS will automatically query its CPU core count and set this property.
	ImportParallelismDegree *int `mandatory:"false" json:"importParallelismDegree"`

	// Maximum number of worker processes that can be used for a Datapump Export job.
	ExportParallelismDegree *int `mandatory:"false" json:"exportParallelismDegree"`
}

func (m CreateDataPumpParameters) String() string {
	return common.PointerString(m)
}
