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

// VolumeAttachment. A base object for all types of attachments between a storage volume and an instance.
// For specific details about iSCSI attachments, see
// IScsiVolumeAttachment.
// For general information about volume attachments, see
// [Overview of Block Volume Storage]({{DOC_SERVER_URL}}/Content/Block/Concepts/overview.htm).
type VolumeAttachment struct {

	// The type of volume attachment.
	AttachmentType *string `mandatory:"true" json:"attachmentType,omitempty"`

	// The Availability Domain of an instance.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain,omitempty"`

	// The OCID of the compartment.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The OCID of the volume attachment.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The OCID of the instance the volume is attached to.
	InstanceID *string `mandatory:"true" json:"instanceId,omitempty"`

	// The current state of the volume attachment.
	LifecycleState VolumeAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The date and time the volume was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The OCID of the volume.
	VolumeID *string `mandatory:"true" json:"volumeId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it cannot be changed.
	// Avoid entering confidential information.
	// Example: `My volume attachment`
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

func (model VolumeAttachment) String() string {
	return common.PointerString(model)
}

type VolumeAttachmentLifecycleStateEnum string
type VolumeAttachmentLifecycleState struct{}

const (
	VOLUME_ATTACHMENT_LIFECYCLE_STATE_ATTACHING VolumeAttachmentLifecycleStateEnum = "ATTACHING"
	VOLUME_ATTACHMENT_LIFECYCLE_STATE_ATTACHED  VolumeAttachmentLifecycleStateEnum = "ATTACHED"
	VOLUME_ATTACHMENT_LIFECYCLE_STATE_DETACHING VolumeAttachmentLifecycleStateEnum = "DETACHING"
	VOLUME_ATTACHMENT_LIFECYCLE_STATE_DETACHED  VolumeAttachmentLifecycleStateEnum = "DETACHED"
	VOLUME_ATTACHMENT_LIFECYCLE_STATE_UNKNOWN   VolumeAttachmentLifecycleStateEnum = "UNKNOWN"
)

var mapping_volumeattachment_lifecycleState = map[string]VolumeAttachmentLifecycleStateEnum{
	"ATTACHING": VOLUME_ATTACHMENT_LIFECYCLE_STATE_ATTACHING,
	"ATTACHED":  VOLUME_ATTACHMENT_LIFECYCLE_STATE_ATTACHED,
	"DETACHING": VOLUME_ATTACHMENT_LIFECYCLE_STATE_DETACHING,
	"DETACHED":  VOLUME_ATTACHMENT_LIFECYCLE_STATE_DETACHED,
	"UNKNOWN":   VOLUME_ATTACHMENT_LIFECYCLE_STATE_UNKNOWN,
}

func (receiver VolumeAttachmentLifecycleState) Values() []VolumeAttachmentLifecycleStateEnum {
	values := make([]VolumeAttachmentLifecycleStateEnum, 0)
	for _, v := range mapping_volumeattachment_lifecycleState {
		if v != VOLUME_ATTACHMENT_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver VolumeAttachmentLifecycleState) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if VolumeAttachmentLifecycleStateEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver VolumeAttachmentLifecycleState) From(toBeConverted string) VolumeAttachmentLifecycleStateEnum {
	if val, ok := mapping_volumeattachment_lifecycleState[toBeConverted]; ok {
		return val
	}
	return VOLUME_ATTACHMENT_LIFECYCLE_STATE_UNKNOWN
}
