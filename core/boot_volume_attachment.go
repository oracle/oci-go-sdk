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

// BootVolumeAttachment. Represents an attachment between a boot volume and an instance.
type BootVolumeAttachment struct {

	// The Availability Domain of an instance.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain,omitempty"`

	// The OCID of the boot volume.
	BootVolumeID *string `mandatory:"true" json:"bootVolumeId,omitempty"`

	// The OCID of the compartment.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The OCID of the boot volume attachment.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The OCID of the instance the boot volume is attached to.
	InstanceID *string `mandatory:"true" json:"instanceId,omitempty"`

	// The current state of the boot volume attachment.
	LifecycleState BootVolumeAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The date and time the boot volume was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// A user-friendly name. Does not have to be unique, and it cannot be changed.
	// Avoid entering confidential information.
	// Example: `My boot volume`
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

func (model BootVolumeAttachment) String() string {
	return common.PointerString(model)
}

type BootVolumeAttachmentLifecycleStateEnum string

const (
	BOOT_VOLUME_ATTACHMENT_LIFECYCLE_STATE_ATTACHING BootVolumeAttachmentLifecycleStateEnum = "ATTACHING"
	BOOT_VOLUME_ATTACHMENT_LIFECYCLE_STATE_ATTACHED  BootVolumeAttachmentLifecycleStateEnum = "ATTACHED"
	BOOT_VOLUME_ATTACHMENT_LIFECYCLE_STATE_DETACHING BootVolumeAttachmentLifecycleStateEnum = "DETACHING"
	BOOT_VOLUME_ATTACHMENT_LIFECYCLE_STATE_DETACHED  BootVolumeAttachmentLifecycleStateEnum = "DETACHED"
	BOOT_VOLUME_ATTACHMENT_LIFECYCLE_STATE_UNKNOWN   BootVolumeAttachmentLifecycleStateEnum = "UNKNOWN"
)

var mapping_bootvolumeattachment_lifecycleState = map[string]BootVolumeAttachmentLifecycleStateEnum{
	"ATTACHING": BOOT_VOLUME_ATTACHMENT_LIFECYCLE_STATE_ATTACHING,
	"ATTACHED":  BOOT_VOLUME_ATTACHMENT_LIFECYCLE_STATE_ATTACHED,
	"DETACHING": BOOT_VOLUME_ATTACHMENT_LIFECYCLE_STATE_DETACHING,
	"DETACHED":  BOOT_VOLUME_ATTACHMENT_LIFECYCLE_STATE_DETACHED,
	"UNKNOWN":   BOOT_VOLUME_ATTACHMENT_LIFECYCLE_STATE_UNKNOWN,
}

func GetBootVolumeAttachmentLifecycleStateEnumValues() []BootVolumeAttachmentLifecycleStateEnum {
	values := make([]BootVolumeAttachmentLifecycleStateEnum, 0)
	for _, v := range mapping_bootvolumeattachment_lifecycleState {
		if v != BOOT_VOLUME_ATTACHMENT_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
