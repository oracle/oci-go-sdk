// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

type CreateDbHomeWithDbSystemIdDetails struct {

	// The OCID of the DB System.
	DbSystemID *string `mandatory:"true" json:"dbSystemId,omitempty"`

	Database *CreateDatabaseDetails `mandatory:"true" json:"database,omitempty"`

	// A valid Oracle database version. To get a list of supported versions, use the ListDbVersions operation.
	DbVersion *string `mandatory:"true" json:"dbVersion,omitempty"`

	// The user-provided name of the database home.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// Source of database:
	//   NONE for creating a new database
	//   DB_BACKUP for creating a new database by restoring a backup
	Source CreateDbHomeWithDbSystemIdDetailsSourceEnum `mandatory:"false" json:"source,omitempty"`
}

func (model CreateDbHomeWithDbSystemIdDetails) String() string {
	return common.PointerString(model)
}

type CreateDbHomeWithDbSystemIdDetailsSourceEnum string

const (
	CREATE_DB_HOME_WITH_DB_SYSTEM_ID_DETAILS_SOURCE_NONE      CreateDbHomeWithDbSystemIdDetailsSourceEnum = "NONE"
	CREATE_DB_HOME_WITH_DB_SYSTEM_ID_DETAILS_SOURCE_DB_BACKUP CreateDbHomeWithDbSystemIdDetailsSourceEnum = "DB_BACKUP"
	CREATE_DB_HOME_WITH_DB_SYSTEM_ID_DETAILS_SOURCE_UNKNOWN   CreateDbHomeWithDbSystemIdDetailsSourceEnum = "UNKNOWN"
)

var mapping_createdbhomewithdbsystemiddetails_source = map[string]CreateDbHomeWithDbSystemIdDetailsSourceEnum{
	"NONE":      CREATE_DB_HOME_WITH_DB_SYSTEM_ID_DETAILS_SOURCE_NONE,
	"DB_BACKUP": CREATE_DB_HOME_WITH_DB_SYSTEM_ID_DETAILS_SOURCE_DB_BACKUP,
	"UNKNOWN":   CREATE_DB_HOME_WITH_DB_SYSTEM_ID_DETAILS_SOURCE_UNKNOWN,
}

func GetCreateDbHomeWithDbSystemIdDetailsSourceEnumValues() []CreateDbHomeWithDbSystemIdDetailsSourceEnum {
	values := make([]CreateDbHomeWithDbSystemIdDetailsSourceEnum, 0)
	for _, v := range mapping_createdbhomewithdbsystemiddetails_source {
		if v != CREATE_DB_HOME_WITH_DB_SYSTEM_ID_DETAILS_SOURCE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
