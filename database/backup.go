// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Backup A database backup.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type Backup struct {

	// The name of the availability domain where the database backup is stored.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The Oracle Database edition of the DB system from which the database backup was taken.
	DatabaseEdition BackupDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// The size of the database in gigabytes at the time the backup was taken.
	DatabaseSizeInGBs *float64 `mandatory:"false" json:"databaseSizeInGBs"`

	// The user-friendly name for the backup. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the backup.
	Id *string `mandatory:"false" json:"id"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the backup.
	LifecycleState BackupLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the backup was completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// The date and time the backup started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The type of backup.
	Type BackupTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m Backup) String() string {
	return common.PointerString(m)
}

// BackupDatabaseEditionEnum Enum with underlying type: string
type BackupDatabaseEditionEnum string

// Set of constants representing the allowable values for BackupDatabaseEditionEnum
const (
	BackupDatabaseEditionStandardEdition                     BackupDatabaseEditionEnum = "STANDARD_EDITION"
	BackupDatabaseEditionEnterpriseEdition                   BackupDatabaseEditionEnum = "ENTERPRISE_EDITION"
	BackupDatabaseEditionEnterpriseEditionHighPerformance    BackupDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	BackupDatabaseEditionEnterpriseEditionExtremePerformance BackupDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingBackupDatabaseEdition = map[string]BackupDatabaseEditionEnum{
	"STANDARD_EDITION":                       BackupDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     BackupDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    BackupDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": BackupDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetBackupDatabaseEditionEnumValues Enumerates the set of values for BackupDatabaseEditionEnum
func GetBackupDatabaseEditionEnumValues() []BackupDatabaseEditionEnum {
	values := make([]BackupDatabaseEditionEnum, 0)
	for _, v := range mappingBackupDatabaseEdition {
		values = append(values, v)
	}
	return values
}

// BackupLifecycleStateEnum Enum with underlying type: string
type BackupLifecycleStateEnum string

// Set of constants representing the allowable values for BackupLifecycleStateEnum
const (
	BackupLifecycleStateCreating  BackupLifecycleStateEnum = "CREATING"
	BackupLifecycleStateActive    BackupLifecycleStateEnum = "ACTIVE"
	BackupLifecycleStateDeleting  BackupLifecycleStateEnum = "DELETING"
	BackupLifecycleStateDeleted   BackupLifecycleStateEnum = "DELETED"
	BackupLifecycleStateFailed    BackupLifecycleStateEnum = "FAILED"
	BackupLifecycleStateRestoring BackupLifecycleStateEnum = "RESTORING"
)

var mappingBackupLifecycleState = map[string]BackupLifecycleStateEnum{
	"CREATING":  BackupLifecycleStateCreating,
	"ACTIVE":    BackupLifecycleStateActive,
	"DELETING":  BackupLifecycleStateDeleting,
	"DELETED":   BackupLifecycleStateDeleted,
	"FAILED":    BackupLifecycleStateFailed,
	"RESTORING": BackupLifecycleStateRestoring,
}

// GetBackupLifecycleStateEnumValues Enumerates the set of values for BackupLifecycleStateEnum
func GetBackupLifecycleStateEnumValues() []BackupLifecycleStateEnum {
	values := make([]BackupLifecycleStateEnum, 0)
	for _, v := range mappingBackupLifecycleState {
		values = append(values, v)
	}
	return values
}

// BackupTypeEnum Enum with underlying type: string
type BackupTypeEnum string

// Set of constants representing the allowable values for BackupTypeEnum
const (
	BackupTypeIncremental BackupTypeEnum = "INCREMENTAL"
	BackupTypeFull        BackupTypeEnum = "FULL"
)

var mappingBackupType = map[string]BackupTypeEnum{
	"INCREMENTAL": BackupTypeIncremental,
	"FULL":        BackupTypeFull,
}

// GetBackupTypeEnumValues Enumerates the set of values for BackupTypeEnum
func GetBackupTypeEnumValues() []BackupTypeEnum {
	values := make([]BackupTypeEnum, 0)
	for _, v := range mappingBackupType {
		values = append(values, v)
	}
	return values
}
