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

// CreateAutonomousDatabaseBackupDetails Details to create an Oracle Autonomous Database backup.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateAutonomousDatabaseBackupDetails struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.
	AutonomousDatabaseId *string `mandatory:"true" json:"autonomousDatabaseId"`

	// The user-friendly name for the backup. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`
}

func (m CreateAutonomousDatabaseBackupDetails) String() string {
	return common.PointerString(m)
}
