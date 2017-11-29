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

type AttachVolumeDetails interface {

	// The OCID of the instance.
	GetInstanceID() *string

	// The OCID of the volume.
	GetVolumeID() *string

	// A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information.
	GetDisplayName() *string
}

type attachvolumedetails struct {
	JsonData    []byte
	InstanceID  *string `mandatory:"true" json:"instanceId,omitempty"`
	VolumeID    *string `mandatory:"true" json:"volumeId,omitempty"`
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
	Type        string  `json:"type"`
}

func (m *attachvolumedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerattachvolumedetails attachvolumedetails
	s := struct {
		Model Unmarshalerattachvolumedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.InstanceID = s.Model.InstanceID
	m.VolumeID = s.Model.VolumeID
	m.DisplayName = s.Model.DisplayName
	m.Type = s.Model.Type

	return err
}

func (m *attachvolumedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
	switch m.Type {
	case "iscsi":
		mm := AttachIScsiVolumeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

func (m attachvolumedetails) GetInstanceID() *string {
	return m.InstanceID
}
func (m attachvolumedetails) GetVolumeID() *string {
	return m.VolumeID
}
func (m attachvolumedetails) GetDisplayName() *string {
	return m.DisplayName
}

func (model attachvolumedetails) String() string {
	return common.PointerString(model)
}
