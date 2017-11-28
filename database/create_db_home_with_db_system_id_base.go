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

type CreateDbHomeWithDbSystemIdBase struct {

	// The OCID of the DB System.
	DbSystemID *string `mandatory:"true" json:"dbSystemId,omitempty"`

	// The user-provided name of the database home.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// Source of database:
	//   NONE for creating a new database
	//   DB_BACKUP for creating a new database by restoring a backup
	Source CreateDbHomeWithDbSystemIdBaseSourceEnum `mandatory:"false" json:"source,omitempty"`
}

func (model CreateDbHomeWithDbSystemIdBase) String() string {
	return common.PointerString(model)
}

type CreateDbHomeWithDbSystemIdBaseSourceEnum string

const (
	CREATE_DB_HOME_WITH_DB_SYSTEM_ID_BASE_SOURCE_NONE      CreateDbHomeWithDbSystemIdBaseSourceEnum = "NONE"
	CREATE_DB_HOME_WITH_DB_SYSTEM_ID_BASE_SOURCE_DB_BACKUP CreateDbHomeWithDbSystemIdBaseSourceEnum = "DB_BACKUP"
	CREATE_DB_HOME_WITH_DB_SYSTEM_ID_BASE_SOURCE_UNKNOWN   CreateDbHomeWithDbSystemIdBaseSourceEnum = "UNKNOWN"
)

var mapping_createdbhomewithdbsystemidbase_source = map[string]CreateDbHomeWithDbSystemIdBaseSourceEnum{
	"NONE":      CREATE_DB_HOME_WITH_DB_SYSTEM_ID_BASE_SOURCE_NONE,
	"DB_BACKUP": CREATE_DB_HOME_WITH_DB_SYSTEM_ID_BASE_SOURCE_DB_BACKUP,
	"UNKNOWN":   CREATE_DB_HOME_WITH_DB_SYSTEM_ID_BASE_SOURCE_UNKNOWN,
}

func GetCreateDbHomeWithDbSystemIdBaseSourceEnumValues() []CreateDbHomeWithDbSystemIdBaseSourceEnum {
	values := make([]CreateDbHomeWithDbSystemIdBaseSourceEnum, 0)
	for _, v := range mapping_createdbhomewithdbsystemidbase_source {
		if v != CREATE_DB_HOME_WITH_DB_SYSTEM_ID_BASE_SOURCE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
