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

// CreateDbHomeWithDbSystemIdDetails.
type CreateDbHomeWithDbSystemIdDetails struct {

	// The OCID of the DB System.
	DbSystemID *string `mandatory:"true" json:"dbSystemId,omitempty"`

	Database *CreateDatabaseDetails `mandatory:"true" json:"database,omitempty"`

	// A valid Oracle database version. To get a list of supported versions, use the ListDbVersions operation.
	DbVersion *string `mandatory:"true" json:"dbVersion,omitempty"`

	// The user-provided name of the database home.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

func (model CreateDbHomeWithDbSystemIdDetails) GetDbSystemID() *string {
	return model.DbSystemID
}
func (model CreateDbHomeWithDbSystemIdDetails) GetDisplayName() *string {
	return model.DisplayName
}

func (m CreateDbHomeWithDbSystemIdDetails) String() string {
	return common.PointerString(m)
}

func (m CreateDbHomeWithDbSystemIdDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbHomeWithDbSystemIdDetails CreateDbHomeWithDbSystemIdDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDbHomeWithDbSystemIdDetails
	}{
		"NONE",
		(MarshalTypeCreateDbHomeWithDbSystemIdDetails)(m),
	}

	return json.Marshal(&s)
}
