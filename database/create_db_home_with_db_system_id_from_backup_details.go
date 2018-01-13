// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

type CreateDbHomeWithDbSystemIdFromBackupDetails struct {

	// The OCID of the DB System.
	DbSystemID *string `mandatory:"true" json:"dbSystemId,omitempty"`

	Database *CreateDatabaseFromBackupDetails `mandatory:"true" json:"database,omitempty"`

	// The user-provided name of the database home.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

//GetDbSystemID returns DbSystemID
func (m CreateDbHomeWithDbSystemIdFromBackupDetails) GetDbSystemID() *string {
	return m.DbSystemID
}

//GetDisplayName returns DisplayName
func (m CreateDbHomeWithDbSystemIdFromBackupDetails) GetDisplayName() *string {
	return m.DisplayName
}

func (m CreateDbHomeWithDbSystemIdFromBackupDetails) String() string {
	return common.PointerString(m)
}

//MarshalJSON marshals to json representation
func (m CreateDbHomeWithDbSystemIdFromBackupDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbHomeWithDbSystemIdFromBackupDetails CreateDbHomeWithDbSystemIdFromBackupDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDbHomeWithDbSystemIdFromBackupDetails
	}{
		"DB_BACKUP",
		(MarshalTypeCreateDbHomeWithDbSystemIdFromBackupDetails)(m),
	}

	return json.Marshal(&s)
}
