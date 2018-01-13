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

type AttachIScsiVolumeDetails struct {

	// The OCID of the instance.
	InstanceID *string `mandatory:"true" json:"instanceId,omitempty"`

	// The OCID of the volume.
	VolumeID *string `mandatory:"true" json:"volumeId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// Whether to use CHAP authentication for the volume attachment. Defaults to false.
	UseChap *bool `mandatory:"false" json:"useChap,omitempty"`
}

//GetDisplayName returns DisplayName
func (m AttachIScsiVolumeDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetInstanceID returns InstanceID
func (m AttachIScsiVolumeDetails) GetInstanceID() *string {
	return m.InstanceID
}

//GetVolumeID returns VolumeID
func (m AttachIScsiVolumeDetails) GetVolumeID() *string {
	return m.VolumeID
}

func (m AttachIScsiVolumeDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AttachIScsiVolumeDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAttachIScsiVolumeDetails AttachIScsiVolumeDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAttachIScsiVolumeDetails
	}{
		"iscsi",
		(MarshalTypeAttachIScsiVolumeDetails)(m),
	}

	return json.Marshal(&s)
}
