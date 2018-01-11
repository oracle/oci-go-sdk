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

type CreateVolumeDetails struct {

	// The Availability Domain of the volume.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain,omitempty"`

	// The OCID of the compartment that contains the volume.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// The size of the volume in GBs.
	SizeInGBs *int `mandatory:"false" json:"sizeInGBs,omitempty"`

	// The size of the volume in MBs. The value must be a multiple of 1024.
	// This field is deprecated. Use sizeInGBs instead.
	SizeInMBs *int `mandatory:"false" json:"sizeInMBs,omitempty"`

	// Specifies the volume source details for a new Block volume. The volume source is either another Block volume in the same Availability Domain or a Block volume backup.
	// This is an optional field. If not specified or set to null, the new Block volume will be empty.
	// When specified, the new Block volume will contain data from the source volume or backup.
	SourceDetails VolumeSourceDetails `mandatory:"false" json:"sourceDetails,omitempty"`

	// The OCID of the volume backup from which the data should be restored on the newly created volume.
	// This field is deprecated. Use the sourceDetails field instead to specify the
	// backup for the volume.
	VolumeBackupID *string `mandatory:"false" json:"volumeBackupId,omitempty"`
}

func (m CreateVolumeDetails) String() string {
	return common.PointerString(m)
}

func (model *CreateVolumeDetails) UnmarshalJSON(data []byte) (e error) {
	m := struct {
		DisplayName        *string             `mandatory:"true" json:"displayName,omitempty"`
		SizeInGBs          *int                `mandatory:"true" json:"sizeInGBs,omitempty"`
		SizeInMBs          *int                `mandatory:"true" json:"sizeInMBs,omitempty"`
		SourceDetails      volumesourcedetails `mandatory:"true" json:"sourceDetails,omitempty"`
		VolumeBackupID     *string             `mandatory:"true" json:"volumeBackupId,omitempty"`
		AvailabilityDomain *string             `mandatory:"true" json:"availabilityDomain,omitempty"`
		CompartmentID      *string             `mandatory:"true" json:"compartmentId,omitempty"`
	}{}

	e = json.Unmarshal(data, &m)
	if e != nil {
		return
	}
	model.DisplayName = m.DisplayName
	model.SizeInGBs = m.SizeInGBs
	model.SizeInMBs = m.SizeInMBs
	nn, e := m.SourceDetails.UnmarshalPolymorphicJSON(m.SourceDetails.JsonData)
	if e != nil {
		return
	}
	model.SourceDetails = nn
	model.VolumeBackupID = m.VolumeBackupID
	model.AvailabilityDomain = m.AvailabilityDomain
	model.CompartmentID = m.CompartmentID
	return
}
