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

// CreateDataTransferMediumDetails Data Transfer Medium details for the Migration. If not specified, it will default to Database Link. Only one type
// of medium details can be specified.
type CreateDataTransferMediumDetails struct {
	DatabaseLinkDetails *CreateDatabaseLinkDetails `mandatory:"false" json:"databaseLinkDetails"`

	ObjectStorageDetails *CreateObjectStoreBucket `mandatory:"false" json:"objectStorageDetails"`
}

func (m CreateDataTransferMediumDetails) String() string {
	return common.PointerString(m)
}
