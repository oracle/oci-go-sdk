// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

type CreateDbHomeWithDbSystemIdFromBackupDetails struct {

	// The OCID of the DB System.
	DbSystemID *string `mandatory:"true" json:"dbSystemId,omitempty"`

	Database *CreateDatabaseFromBackupDetails `mandatory:"true" json:"database,omitempty"`

	// The user-provided name of the database home.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// Source of database:
	//   NONE for creating a new database
	//   DB_BACKUP for creating a new database by restoring a backup
	Source CreateDbHomeWithDbSystemIdFromBackupDetailsSourceEnum `mandatory:"false" json:"source,omitempty"`
}

func (model CreateDbHomeWithDbSystemIdFromBackupDetails) String() string {
	return common.PointerString(model)
}

type CreateDbHomeWithDbSystemIdFromBackupDetailsSourceEnum string

const (
	CREATE_DB_HOME_WITH_DB_SYSTEM_ID_FROM_BACKUP_DETAILS_SOURCE_NONE      CreateDbHomeWithDbSystemIdFromBackupDetailsSourceEnum = "NONE"
	CREATE_DB_HOME_WITH_DB_SYSTEM_ID_FROM_BACKUP_DETAILS_SOURCE_DB_BACKUP CreateDbHomeWithDbSystemIdFromBackupDetailsSourceEnum = "DB_BACKUP"
	CREATE_DB_HOME_WITH_DB_SYSTEM_ID_FROM_BACKUP_DETAILS_SOURCE_UNKNOWN   CreateDbHomeWithDbSystemIdFromBackupDetailsSourceEnum = "UNKNOWN"
)

var mapping_createdbhomewithdbsystemidfrombackupdetails_source = map[string]CreateDbHomeWithDbSystemIdFromBackupDetailsSourceEnum{
	"NONE":      CREATE_DB_HOME_WITH_DB_SYSTEM_ID_FROM_BACKUP_DETAILS_SOURCE_NONE,
	"DB_BACKUP": CREATE_DB_HOME_WITH_DB_SYSTEM_ID_FROM_BACKUP_DETAILS_SOURCE_DB_BACKUP,
	"UNKNOWN":   CREATE_DB_HOME_WITH_DB_SYSTEM_ID_FROM_BACKUP_DETAILS_SOURCE_UNKNOWN,
}

func GetCreateDbHomeWithDbSystemIdFromBackupDetailsSourceEnumValues() []CreateDbHomeWithDbSystemIdFromBackupDetailsSourceEnum {
	values := make([]CreateDbHomeWithDbSystemIdFromBackupDetailsSourceEnum, 0)
	for _, v := range mapping_createdbhomewithdbsystemidfrombackupdetails_source {
		if v != CREATE_DB_HOME_WITH_DB_SYSTEM_ID_FROM_BACKUP_DETAILS_SOURCE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
