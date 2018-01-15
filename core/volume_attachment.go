// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// VolumeAttachment A base object for all types of attachments between a storage volume and an instance.
// For specific details about iSCSI attachments, see
// IScsiVolumeAttachment.
// For general information about volume attachments, see
// [Overview of Block Volume Storage]({{DOC_SERVER_URL}}/Content/Block/Concepts/overview.htm).
// VolumeAttachment is an interface representing the polymorphic json shape of this model
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
	JsonData           []byte
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

// UnmarshalJSON unmarshals json
func (m *volumeattachment) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalervolumeattachment volumeattachment
	s := struct {
		Model Unmarshalervolumeattachment
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AvailabilityDomain = s.Model.AvailabilityDomain
	m.CompartmentID = s.Model.CompartmentID
	m.ID = s.Model.ID
	m.InstanceID = s.Model.InstanceID
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.VolumeID = s.Model.VolumeID
	m.DisplayName = s.Model.DisplayName
	m.AttachmentType = s.Model.AttachmentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *volumeattachment) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
	switch m.AttachmentType {
	case "iscsi":
		mm := IScsiVolumeAttachment{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

//GetAvailabilityDomain returns AvailabilityDomain
func (m volumeattachment) GetAvailabilityDomain() *string {
	return m.AvailabilityDomain
}

//GetCompartmentID returns CompartmentID
func (m volumeattachment) GetCompartmentID() *string {
	return m.CompartmentID
}

//GetID returns ID
func (m volumeattachment) GetID() *string {
	return m.ID
}

//GetInstanceID returns InstanceID
func (m volumeattachment) GetInstanceID() *string {
	return m.InstanceID
}

//GetLifecycleState returns LifecycleState
func (m volumeattachment) GetLifecycleState() VolumeAttachmentLifecycleStateEnum {
	return m.LifecycleState
}

//GetTimeCreated returns TimeCreated
func (m volumeattachment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetVolumeID returns VolumeID
func (m volumeattachment) GetVolumeID() *string {
	return m.VolumeID
}

//GetDisplayName returns DisplayName
func (m volumeattachment) GetDisplayName() *string {
	return m.DisplayName
}

func (m volumeattachment) String() string {
	return common.PointerString(m)
}

// VolumeAttachmentLifecycleStateEnum Enum with underlying type: string
type VolumeAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for VolumeAttachmentLifecycleState
const (
	VolumeAttachmentLifecycleStateAttaching VolumeAttachmentLifecycleStateEnum = "ATTACHING"
	VolumeAttachmentLifecycleStateAttached  VolumeAttachmentLifecycleStateEnum = "ATTACHED"
	VolumeAttachmentLifecycleStateDetaching VolumeAttachmentLifecycleStateEnum = "DETACHING"
	VolumeAttachmentLifecycleStateDetached  VolumeAttachmentLifecycleStateEnum = "DETACHED"
	VolumeAttachmentLifecycleStateUnknown   VolumeAttachmentLifecycleStateEnum = "UNKNOWN"
)

var mappingVolumeAttachmentLifecycleState = map[string]VolumeAttachmentLifecycleStateEnum{
	"ATTACHING": VolumeAttachmentLifecycleStateAttaching,
	"ATTACHED":  VolumeAttachmentLifecycleStateAttached,
	"DETACHING": VolumeAttachmentLifecycleStateDetaching,
	"DETACHED":  VolumeAttachmentLifecycleStateDetached,
	"UNKNOWN":   VolumeAttachmentLifecycleStateUnknown,
}

// GetVolumeAttachmentLifecycleStateEnumValues Enumerates the set of values for VolumeAttachmentLifecycleState
func GetVolumeAttachmentLifecycleStateEnumValues() []VolumeAttachmentLifecycleStateEnum {
	values := make([]VolumeAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingVolumeAttachmentLifecycleState {
		if v != VolumeAttachmentLifecycleStateUnknown {
			values = append(values, v)
		}
	}
	return values
}
