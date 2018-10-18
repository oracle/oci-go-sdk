// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
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

// InstanceConfigurationBlockVolumeDetails Create new block volumes or attach to an existing volume. Specify either createDetails or volumeId.
type InstanceConfigurationBlockVolumeDetails struct {
	AttachDetails InstanceConfigurationAttachVolumeDetails `mandatory:"false" json:"attachDetails"`

	CreateDetails *InstanceConfigurationCreateVolumeDetails `mandatory:"false" json:"createDetails"`

	// The OCID of the volume.
	VolumeId *string `mandatory:"false" json:"volumeId"`
}

func (m InstanceConfigurationBlockVolumeDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *InstanceConfigurationBlockVolumeDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		AttachDetails instanceconfigurationattachvolumedetails  `json:"attachDetails"`
		CreateDetails *InstanceConfigurationCreateVolumeDetails `json:"createDetails"`
		VolumeId      *string                                   `json:"volumeId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	nn, e := model.AttachDetails.UnmarshalPolymorphicJSON(model.AttachDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AttachDetails = nn.(InstanceConfigurationAttachVolumeDetails)
	} else {
		m.AttachDetails = nil
	}
	m.CreateDetails = model.CreateDetails
	m.VolumeId = model.VolumeId
	return
}
