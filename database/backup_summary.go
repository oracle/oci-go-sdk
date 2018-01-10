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

// BackupSummary A database backup
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type BackupSummary struct {

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
	LifecycleState BackupSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the backup was completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded,omitempty"`

	// The date and time the backup starts.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted,omitempty"`

	// The type of backup.
	Type_ BackupSummaryType_Enum `mandatory:"false" json:"type,omitempty"`
}

func (model BackupSummary) String() string {
	return common.PointerString(model)
}

// BackupSummaryLifecycleStateEnum Enum with underlying type: string
type BackupSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for BackupSummaryLifecycleState
const (
	BackupSummaryLifecycleStateCreating  BackupSummaryLifecycleStateEnum = "CREATING"
	BackupSummaryLifecycleStateActive    BackupSummaryLifecycleStateEnum = "ACTIVE"
	BackupSummaryLifecycleStateDeleting  BackupSummaryLifecycleStateEnum = "DELETING"
	BackupSummaryLifecycleStateDeleted   BackupSummaryLifecycleStateEnum = "DELETED"
	BackupSummaryLifecycleStateFailed    BackupSummaryLifecycleStateEnum = "FAILED"
	BackupSummaryLifecycleStateRestoring BackupSummaryLifecycleStateEnum = "RESTORING"
	BackupSummaryLifecycleStateUnknown   BackupSummaryLifecycleStateEnum = "UNKNOWN"
)

var mappingBackupSummaryLifecycleState = map[string]BackupSummaryLifecycleStateEnum{
	"CREATING":  BackupSummaryLifecycleStateCreating,
	"ACTIVE":    BackupSummaryLifecycleStateActive,
	"DELETING":  BackupSummaryLifecycleStateDeleting,
	"DELETED":   BackupSummaryLifecycleStateDeleted,
	"FAILED":    BackupSummaryLifecycleStateFailed,
	"RESTORING": BackupSummaryLifecycleStateRestoring,
	"UNKNOWN":   BackupSummaryLifecycleStateUnknown,
}

// GetBackupSummaryLifecycleStateEnumValues Enumerates the set of values for BackupSummaryLifecycleState
func GetBackupSummaryLifecycleStateEnumValues() []BackupSummaryLifecycleStateEnum {
	values := make([]BackupSummaryLifecycleStateEnum, 0)
	for _, v := range mappingBackupSummaryLifecycleState {
		if v != BackupSummaryLifecycleStateUnknown {
			values = append(values, v)
		}
	}
	return values
}

// BackupSummaryType_Enum Enum with underlying type: string
type BackupSummaryType_Enum string

// Set of constants representing the allowable values for BackupSummaryType_
const (
	BackupSummaryType_Incremental BackupSummaryType_Enum = "INCREMENTAL"
	BackupSummaryType_Full        BackupSummaryType_Enum = "FULL"
	BackupSummaryType_Unknown     BackupSummaryType_Enum = "UNKNOWN"
)

var mappingBackupSummaryType_ = map[string]BackupSummaryType_Enum{
	"INCREMENTAL": BackupSummaryType_Incremental,
	"FULL":        BackupSummaryType_Full,
	"UNKNOWN":     BackupSummaryType_Unknown,
}

// GetBackupSummaryType_EnumValues Enumerates the set of values for BackupSummaryType_
func GetBackupSummaryType_EnumValues() []BackupSummaryType_Enum {
	values := make([]BackupSummaryType_Enum, 0)
	for _, v := range mappingBackupSummaryType_ {
		if v != BackupSummaryType_Unknown {
			values = append(values, v)
		}
	}
	return values
}
