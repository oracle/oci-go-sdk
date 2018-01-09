// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// BootVolume. A detachable boot volume device that contains the image used to boot an Compute instance. For more information, see
// [Overview of Boot Volumes]({{DOC_SERVER_URL}}/Content/Block/Concepts/bootvolumes.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type BootVolume struct {

	// The Availability Domain of the boot volume.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain,omitempty"`

	// The OCID of the compartment that contains the boot volume.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The boot volume's Oracle ID (OCID).
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The current state of a boot volume.
	LifecycleState BootVolumeLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The size of the volume in MBs. The value must be a multiple of 1024.
	// This field is deprecated. Please use sizeInGBs.
	SizeInMBs *int `mandatory:"true" json:"sizeInMBs,omitempty"`

	// The date and time the boot volume was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// The image OCID used to create the boot volume.
	ImageID *string `mandatory:"false" json:"imageId,omitempty"`

	// The size of the boot volume in GBs.
	SizeInGBs *int `mandatory:"false" json:"sizeInGBs,omitempty"`
}

func (model BootVolume) String() string {
	return common.PointerString(model)
}

type BootVolumeLifecycleStateEnum string

const (
	BOOT_VOLUME_LIFECYCLE_STATE_PROVISIONING BootVolumeLifecycleStateEnum = "PROVISIONING"
	BOOT_VOLUME_LIFECYCLE_STATE_RESTORING    BootVolumeLifecycleStateEnum = "RESTORING"
	BOOT_VOLUME_LIFECYCLE_STATE_AVAILABLE    BootVolumeLifecycleStateEnum = "AVAILABLE"
	BOOT_VOLUME_LIFECYCLE_STATE_TERMINATING  BootVolumeLifecycleStateEnum = "TERMINATING"
	BOOT_VOLUME_LIFECYCLE_STATE_TERMINATED   BootVolumeLifecycleStateEnum = "TERMINATED"
	BOOT_VOLUME_LIFECYCLE_STATE_FAULTY       BootVolumeLifecycleStateEnum = "FAULTY"
	BOOT_VOLUME_LIFECYCLE_STATE_UNKNOWN      BootVolumeLifecycleStateEnum = "UNKNOWN"
)

var mapping_bootvolume_lifecycleState = map[string]BootVolumeLifecycleStateEnum{
	"PROVISIONING": BOOT_VOLUME_LIFECYCLE_STATE_PROVISIONING,
	"RESTORING":    BOOT_VOLUME_LIFECYCLE_STATE_RESTORING,
	"AVAILABLE":    BOOT_VOLUME_LIFECYCLE_STATE_AVAILABLE,
	"TERMINATING":  BOOT_VOLUME_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   BOOT_VOLUME_LIFECYCLE_STATE_TERMINATED,
	"FAULTY":       BOOT_VOLUME_LIFECYCLE_STATE_FAULTY,
	"UNKNOWN":      BOOT_VOLUME_LIFECYCLE_STATE_UNKNOWN,
}

func GetBootVolumeLifecycleStateEnumValues() []BootVolumeLifecycleStateEnum {
	values := make([]BootVolumeLifecycleStateEnum, 0)
	for _, v := range mapping_bootvolume_lifecycleState {
		if v != BOOT_VOLUME_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
