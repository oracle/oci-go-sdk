// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// VolumeBackup. A point-in-time copy of a volume that can then be used to create a new block volume
// or recover a block volume. For more information, see
// [Overview of Cloud Volume Storage]({{DOC_SERVER_URL}}/Content/Block/Concepts/overview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type VolumeBackup struct {

	// The OCID of the compartment that contains the volume backup.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// A user-friendly name for the volume backup. Does not have to be unique and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName,omitempty"`

	// The OCID of the volume backup.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The current state of a volume backup.
	LifecycleState VolumeBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The date and time the volume backup was created. This is the time the actual point-in-time image
	// of the volume data was taken. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The size of the volume, in GBs.
	SizeInGBs *int64 `mandatory:"false" json:"sizeInGBs,omitempty"`

	// The size of the volume in MBs. The value must be a multiple of 1024.
	// This field is deprecated. Please use sizeInGBs.
	SizeInMBs *int64 `mandatory:"false" json:"sizeInMBs,omitempty"`

	// The date and time the request to create the volume backup was received. Format defined by RFC3339.
	TimeRequestReceived *common.SDKTime `mandatory:"false" json:"timeRequestReceived,omitempty"`

	// The size used by the backup, in GBs. It is typically smaller than sizeInGBs, depending on the space
	// consumed on the volume and whether the backup is full or incremental.
	UniqueSizeInGBs *int64 `mandatory:"false" json:"uniqueSizeInGBs,omitempty"`

	// The size used by the backup, in MBs. It is typically smaller than sizeInMBs, depending on the space
	// consumed on the volume and whether the backup is full or incremental.
	// This field is deprecated. Please use uniqueSizeInGBs.
	UniqueSizeInMbs *int64 `mandatory:"false" json:"uniqueSizeInMbs,omitempty"`

	// The OCID of the volume.
	VolumeID *string `mandatory:"false" json:"volumeId,omitempty"`
}

func (model VolumeBackup) String() string {
	return common.PointerString(model)
}

type VolumeBackupLifecycleStateEnum string
type VolumeBackupLifecycleState struct{}

const (
	VOLUME_BACKUP_LIFECYCLE_STATE_CREATING         VolumeBackupLifecycleStateEnum = "CREATING"
	VOLUME_BACKUP_LIFECYCLE_STATE_AVAILABLE        VolumeBackupLifecycleStateEnum = "AVAILABLE"
	VOLUME_BACKUP_LIFECYCLE_STATE_TERMINATING      VolumeBackupLifecycleStateEnum = "TERMINATING"
	VOLUME_BACKUP_LIFECYCLE_STATE_TERMINATED       VolumeBackupLifecycleStateEnum = "TERMINATED"
	VOLUME_BACKUP_LIFECYCLE_STATE_FAULTY           VolumeBackupLifecycleStateEnum = "FAULTY"
	VOLUME_BACKUP_LIFECYCLE_STATE_REQUEST_RECEIVED VolumeBackupLifecycleStateEnum = "REQUEST_RECEIVED"
	VOLUME_BACKUP_LIFECYCLE_STATE_UNKNOWN          VolumeBackupLifecycleStateEnum = "UNKNOWN"
)

var mapping_volumebackup_lifecycleState = map[string]VolumeBackupLifecycleStateEnum{
	"CREATING":         VOLUME_BACKUP_LIFECYCLE_STATE_CREATING,
	"AVAILABLE":        VOLUME_BACKUP_LIFECYCLE_STATE_AVAILABLE,
	"TERMINATING":      VOLUME_BACKUP_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":       VOLUME_BACKUP_LIFECYCLE_STATE_TERMINATED,
	"FAULTY":           VOLUME_BACKUP_LIFECYCLE_STATE_FAULTY,
	"REQUEST_RECEIVED": VOLUME_BACKUP_LIFECYCLE_STATE_REQUEST_RECEIVED,
	"UNKNOWN":          VOLUME_BACKUP_LIFECYCLE_STATE_UNKNOWN,
}

func (receiver VolumeBackupLifecycleState) Values() []VolumeBackupLifecycleStateEnum {
	values := make([]VolumeBackupLifecycleStateEnum, 0)
	for _, v := range mapping_volumebackup_lifecycleState {
		if v != VOLUME_BACKUP_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver VolumeBackupLifecycleState) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if VolumeBackupLifecycleStateEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver VolumeBackupLifecycleState) From(toBeConverted string) VolumeBackupLifecycleStateEnum {
	if val, ok := mapping_volumebackup_lifecycleState[toBeConverted]; ok {
		return val
	}
	return VOLUME_BACKUP_LIFECYCLE_STATE_UNKNOWN
}
