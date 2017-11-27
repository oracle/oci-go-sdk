// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oci-go-sdk/common"
)

type CreateDatabaseFromBackupDetails struct {

	// A strong password for SYS, SYSTEM, PDB Admin and TDE Wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
	AdminPassword *string `mandatory:"true" json:"adminPassword,omitempty"`

	// The backup OCID.
	BackupID *string `mandatory:"true" json:"backupId,omitempty"`

	// The password to open the TDE wallet.
	BackupTDEPassword *string `mandatory:"true" json:"backupTDEPassword,omitempty"`
}

func (model CreateDatabaseFromBackupDetails) String() string {
	return common.PointerString(model)
}
