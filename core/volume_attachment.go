// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"encoding/json"
)

// VolumeAttachment. A base object for all types of attachments between a storage volume and an instance.
// For specific details about iSCSI attachments, see
// IScsiVolumeAttachment.
// For general information about volume attachments, see
// [Overview of Block Volume Storage]({{DOC_SERVER_URL}}/Content/Block/Concepts/overview.htm).
type VolumeAttachment interface {

	// The Availability Domain of an instance.
	// Example: `Uocm:PHX-AD-1`
	GetAvailabilityDomain() *string

	// The OCID of the compartment.
	GetCompartmentID() *string

	// The OCID of the volume attachment.
	GetID() *string

	// The OCID of the instance the volume is attached to.
	GetInstanceID() *string

	// The current state of the volume attachment.
	GetLifecycleState() VolumeAttachmentLifecycleStateEnum

	// The date and time the volume was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	// The OCID of the volume.
	GetVolumeID() *string

	// A user-friendly name. Does not have to be unique, and it cannot be changed.
	// Avoid entering confidential information.
	// Example: `My volume attachment`
	GetDisplayName() *string
}

type volumeattachment struct {
	AvailabilityDomain *string                            `mandatory:"true" json:"availabilityDomain,omitempty"`
	CompartmentID      *string                            `mandatory:"true" json:"compartmentId,omitempty"`
	ID                 *string                            `mandatory:"true" json:"id,omitempty"`
	InstanceID         *string                            `mandatory:"true" json:"instanceId,omitempty"`
	LifecycleState     VolumeAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`
	TimeCreated        *common.SDKTime                    `mandatory:"true" json:"timeCreated,omitempty"`
	VolumeID           *string                            `mandatory:"true" json:"volumeId,omitempty"`
	DisplayName        *string                            `mandatory:"false" json:"displayName,omitempty"`
	AttachmentType     string                             `json:"attachmentType"`
}

func (m *volumeattachment) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	err := json.Unmarshal(data, m)
	if err != nil {
		return nil, err
	}

	switch m.AttachmentType {
	case "iscsi":
		mm := IScsiVolumeAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

func (m volumeattachment) GetAvailabilityDomain() *string {
	return m.AvailabilityDomain
}
func (m volumeattachment) GetCompartmentID() *string {
	return m.CompartmentID
}
func (m volumeattachment) GetID() *string {
	return m.ID
}
func (m volumeattachment) GetInstanceID() *string {
	return m.InstanceID
}
func (m volumeattachment) GetLifecycleState() VolumeAttachmentLifecycleStateEnum {
	return m.LifecycleState
}
func (m volumeattachment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}
func (m volumeattachment) GetVolumeID() *string {
	return m.VolumeID
}
func (m volumeattachment) GetDisplayName() *string {
	return m.DisplayName
}

func (model volumeattachment) String() string {
	return common.PointerString(model)
}

type VolumeAttachmentLifecycleStateEnum string

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

func GetVolumeAttachmentLifecycleStateEnumValues() []VolumeAttachmentLifecycleStateEnum {
	values := make([]VolumeAttachmentLifecycleStateEnum, 0)
	for _, v := range mapping_volumeattachment_lifecycleState {
		if v != VOLUME_ATTACHMENT_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
