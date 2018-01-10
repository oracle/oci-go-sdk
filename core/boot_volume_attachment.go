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

// BootVolumeAttachment Represents an attachment between a boot volume and an instance.
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

// BootVolumeAttachmentLifecycleStateEnum Enum with underlying type: string
type BootVolumeAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for BootVolumeAttachmentLifecycleState
const (
	BootVolumeAttachmentLifecycleStateAttaching BootVolumeAttachmentLifecycleStateEnum = "ATTACHING"
	BootVolumeAttachmentLifecycleStateAttached  BootVolumeAttachmentLifecycleStateEnum = "ATTACHED"
	BootVolumeAttachmentLifecycleStateDetaching BootVolumeAttachmentLifecycleStateEnum = "DETACHING"
	BootVolumeAttachmentLifecycleStateDetached  BootVolumeAttachmentLifecycleStateEnum = "DETACHED"
	BootVolumeAttachmentLifecycleStateUnknown   BootVolumeAttachmentLifecycleStateEnum = "UNKNOWN"
)

var mappingBootVolumeAttachmentLifecycleState = map[string]BootVolumeAttachmentLifecycleStateEnum{
	"ATTACHING": BootVolumeAttachmentLifecycleStateAttaching,
	"ATTACHED":  BootVolumeAttachmentLifecycleStateAttached,
	"DETACHING": BootVolumeAttachmentLifecycleStateDetaching,
	"DETACHED":  BootVolumeAttachmentLifecycleStateDetached,
	"UNKNOWN":   BootVolumeAttachmentLifecycleStateUnknown,
}

// GetBootVolumeAttachmentLifecycleStateEnumValues Enumerates the set of values for BootVolumeAttachmentLifecycleState
func GetBootVolumeAttachmentLifecycleStateEnumValues() []BootVolumeAttachmentLifecycleStateEnum {
	values := make([]BootVolumeAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingBootVolumeAttachmentLifecycleState {
		if v != BootVolumeAttachmentLifecycleStateUnknown {
			values = append(values, v)
		}
	}
	return values
}
