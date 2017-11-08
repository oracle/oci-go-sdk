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

// BackupSummary. A database backup
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

type BackupSummaryLifecycleStateEnum string

const (
	BACKUP_SUMMARY_LIFECYCLE_STATE_CREATING  BackupSummaryLifecycleStateEnum = "CREATING"
	BACKUP_SUMMARY_LIFECYCLE_STATE_ACTIVE    BackupSummaryLifecycleStateEnum = "ACTIVE"
	BACKUP_SUMMARY_LIFECYCLE_STATE_DELETING  BackupSummaryLifecycleStateEnum = "DELETING"
	BACKUP_SUMMARY_LIFECYCLE_STATE_DELETED   BackupSummaryLifecycleStateEnum = "DELETED"
	BACKUP_SUMMARY_LIFECYCLE_STATE_FAILED    BackupSummaryLifecycleStateEnum = "FAILED"
	BACKUP_SUMMARY_LIFECYCLE_STATE_RESTORING BackupSummaryLifecycleStateEnum = "RESTORING"
	BACKUP_SUMMARY_LIFECYCLE_STATE_UNKNOWN   BackupSummaryLifecycleStateEnum = "UNKNOWN"
)

var mapping_backupsummary_lifecycleState = map[string]BackupSummaryLifecycleStateEnum{
	"CREATING":  BACKUP_SUMMARY_LIFECYCLE_STATE_CREATING,
	"ACTIVE":    BACKUP_SUMMARY_LIFECYCLE_STATE_ACTIVE,
	"DELETING":  BACKUP_SUMMARY_LIFECYCLE_STATE_DELETING,
	"DELETED":   BACKUP_SUMMARY_LIFECYCLE_STATE_DELETED,
	"FAILED":    BACKUP_SUMMARY_LIFECYCLE_STATE_FAILED,
	"RESTORING": BACKUP_SUMMARY_LIFECYCLE_STATE_RESTORING,
	"UNKNOWN":   BACKUP_SUMMARY_LIFECYCLE_STATE_UNKNOWN,
}

func GetBackupSummaryLifecycleStateEnumValues() []BackupSummaryLifecycleStateEnum {
	values := make([]BackupSummaryLifecycleStateEnum, 0)
	for _, v := range mapping_backupsummary_lifecycleState {
		if v != BACKUP_SUMMARY_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type BackupSummaryType_Enum string

const (
	BACKUP_SUMMARY_TYPE__INCREMENTAL BackupSummaryType_Enum = "INCREMENTAL"
	BACKUP_SUMMARY_TYPE__FULL        BackupSummaryType_Enum = "FULL"
	BACKUP_SUMMARY_TYPE__UNKNOWN     BackupSummaryType_Enum = "UNKNOWN"
)

var mapping_backupsummary_type = map[string]BackupSummaryType_Enum{
	"INCREMENTAL": BACKUP_SUMMARY_TYPE__INCREMENTAL,
	"FULL":        BACKUP_SUMMARY_TYPE__FULL,
	"UNKNOWN":     BACKUP_SUMMARY_TYPE__UNKNOWN,
}

func GetBackupSummaryType_EnumValues() []BackupSummaryType_Enum {
	values := make([]BackupSummaryType_Enum, 0)
	for _, v := range mapping_backupsummary_type {
		if v != BACKUP_SUMMARY_TYPE__UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
