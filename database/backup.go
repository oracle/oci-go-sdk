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

// Backup. A database backup
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type Backup struct {

	// The name of the Availability Domain that the backup is located in.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain,omitempty"`

	// The OCID of the compartment.
	CompartmentID *string `mandatory:"false" json:"compartmentId,omitempty"`

	// The OCID of the database.
	DatabaseID *string `mandatory:"false" json:"databaseId,omitempty"`

	// The user-friendly name for the backup. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// The OCID of the backup.
	ID *string `mandatory:"false" json:"id,omitempty"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// The current state of the backup.
	LifecycleState BackupLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the backup was completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded,omitempty"`

	// The date and time the backup starts.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted,omitempty"`

	// The type of backup.
	Type_ BackupType_Enum `mandatory:"false" json:"type,omitempty"`
}

func (model Backup) String() string {
	return common.PointerString(model)
}

type BackupLifecycleStateEnum string

const (
	BACKUP_LIFECYCLE_STATE_CREATING  BackupLifecycleStateEnum = "CREATING"
	BACKUP_LIFECYCLE_STATE_ACTIVE    BackupLifecycleStateEnum = "ACTIVE"
	BACKUP_LIFECYCLE_STATE_DELETING  BackupLifecycleStateEnum = "DELETING"
	BACKUP_LIFECYCLE_STATE_DELETED   BackupLifecycleStateEnum = "DELETED"
	BACKUP_LIFECYCLE_STATE_FAILED    BackupLifecycleStateEnum = "FAILED"
	BACKUP_LIFECYCLE_STATE_RESTORING BackupLifecycleStateEnum = "RESTORING"
	BACKUP_LIFECYCLE_STATE_UNKNOWN   BackupLifecycleStateEnum = "UNKNOWN"
)

var mapping_backup_lifecycleState = map[string]BackupLifecycleStateEnum{
	"CREATING":  BACKUP_LIFECYCLE_STATE_CREATING,
	"ACTIVE":    BACKUP_LIFECYCLE_STATE_ACTIVE,
	"DELETING":  BACKUP_LIFECYCLE_STATE_DELETING,
	"DELETED":   BACKUP_LIFECYCLE_STATE_DELETED,
	"FAILED":    BACKUP_LIFECYCLE_STATE_FAILED,
	"RESTORING": BACKUP_LIFECYCLE_STATE_RESTORING,
	"UNKNOWN":   BACKUP_LIFECYCLE_STATE_UNKNOWN,
}

func GetBackupLifecycleStateEnumValues() []BackupLifecycleStateEnum {
	values := make([]BackupLifecycleStateEnum, 0)
	for _, v := range mapping_backup_lifecycleState {
		if v != BACKUP_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type BackupType_Enum string

const (
	BACKUP_TYPE__INCREMENTAL BackupType_Enum = "INCREMENTAL"
	BACKUP_TYPE__FULL        BackupType_Enum = "FULL"
	BACKUP_TYPE__UNKNOWN     BackupType_Enum = "UNKNOWN"
)

var mapping_backup_type = map[string]BackupType_Enum{
	"INCREMENTAL": BACKUP_TYPE__INCREMENTAL,
	"FULL":        BACKUP_TYPE__FULL,
	"UNKNOWN":     BACKUP_TYPE__UNKNOWN,
}

func GetBackupType_EnumValues() []BackupType_Enum {
	values := make([]BackupType_Enum, 0)
	for _, v := range mapping_backup_type {
		if v != BACKUP_TYPE__UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
