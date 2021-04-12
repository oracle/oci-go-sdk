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

// CreateDataPumpSettings Optional settings for Datapump Export and Import jobs
type CreateDataPumpSettings struct {

	// DataPump job mode.
	// Refer to docs.oracle.com/en/database/oracle/oracle-database/19/arpls/ODMS_DATAPUMP.html#GUID-92C2CB46-8BC9-414D-B62E-79CD788C1E62__BABBDEHD
	JobMode DataPumpJobModeEnum `mandatory:"false" json:"jobMode,omitempty"`

	DataPumpParameters *CreateDataPumpParameters `mandatory:"false" json:"dataPumpParameters"`

	// Defines remapping to be applied to objects as they are processed.
	// Refer to https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/ODMS_DATAPUMP.html#GUID-0FC32790-91E6-4781-87A3-229DE024CB3D.
	MetadataRemaps []MetadataRemap `mandatory:"false" json:"metadataRemaps"`

	ExportDirectoryObject *CreateDirectoryObject `mandatory:"false" json:"exportDirectoryObject"`

	ImportDirectoryObject *CreateDirectoryObject `mandatory:"false" json:"importDirectoryObject"`
}

func (m CreateDataPumpSettings) String() string {
	return common.PointerString(m)
}
