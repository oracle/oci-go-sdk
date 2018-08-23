// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RestoreAutonomousDataWarehouseDetails Details to restore an Oracle Autonomous Data Warehouse.
type RestoreAutonomousDataWarehouseDetails struct {

	// The time to restore the database to.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m RestoreAutonomousDataWarehouseDetails) String() string {
	return common.PointerString(m)
}
