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

// Backup A database backup
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
	Type BackupType_Enum `mandatory:"false" json:"type,omitempty"`
}

func (model Backup) String() string {
	return common.PointerString(model)
}

// BackupLifecycleStateEnum Enum with underlying type: string
type BackupLifecycleStateEnum string

// Set of constants representing the allowable values for BackupLifecycleState
const (
	BackupLifecycleStateCreating  BackupLifecycleStateEnum = "CREATING"
	BackupLifecycleStateActive    BackupLifecycleStateEnum = "ACTIVE"
	BackupLifecycleStateDeleting  BackupLifecycleStateEnum = "DELETING"
	BackupLifecycleStateDeleted   BackupLifecycleStateEnum = "DELETED"
	BackupLifecycleStateFailed    BackupLifecycleStateEnum = "FAILED"
	BackupLifecycleStateRestoring BackupLifecycleStateEnum = "RESTORING"
	BackupLifecycleStateUnknown   BackupLifecycleStateEnum = "UNKNOWN"
)

var mappingBackupLifecycleState = map[string]BackupLifecycleStateEnum{
	"CREATING":  BackupLifecycleStateCreating,
	"ACTIVE":    BackupLifecycleStateActive,
	"DELETING":  BackupLifecycleStateDeleting,
	"DELETED":   BackupLifecycleStateDeleted,
	"FAILED":    BackupLifecycleStateFailed,
	"RESTORING": BackupLifecycleStateRestoring,
	"UNKNOWN":   BackupLifecycleStateUnknown,
}

// GetBackupLifecycleStateEnumValues Enumerates the set of values for BackupLifecycleState
func GetBackupLifecycleStateEnumValues() []BackupLifecycleStateEnum {
	values := make([]BackupLifecycleStateEnum, 0)
	for _, v := range mappingBackupLifecycleState {
		if v != BackupLifecycleStateUnknown {
			values = append(values, v)
		}
	}
	return values
}

// BackupType_Enum Enum with underlying type: string
type BackupType_Enum string

// Set of constants representing the allowable values for BackupType_
const (
	BackupType_Incremental BackupType_Enum = "INCREMENTAL"
	BackupType_Full        BackupType_Enum = "FULL"
	BackupType_Unknown     BackupType_Enum = "UNKNOWN"
)

var mappingBackupType_ = map[string]BackupType_Enum{
	"INCREMENTAL": BackupType_Incremental,
	"FULL":        BackupType_Full,
	"UNKNOWN":     BackupType_Unknown,
}

// GetBackupType_EnumValues Enumerates the set of values for BackupType_
func GetBackupType_EnumValues() []BackupType_Enum {
	values := make([]BackupType_Enum, 0)
	for _, v := range mappingBackupType_ {
		if v != BackupType_Unknown {
			values = append(values, v)
		}
	}
	return values
}
