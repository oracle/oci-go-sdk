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

// VolumeSourceFromVolumeBackupDetails. Specifies the volume backup.
type VolumeSourceFromVolumeBackupDetails struct {

	// The OCID of the volume backup.
	ID *string `mandatory:"true" json:"id,omitempty"`
}

func (m VolumeSourceFromVolumeBackupDetails) String() string {
	return common.PointerString(m)
}

func (m VolumeSourceFromVolumeBackupDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVolumeSourceFromVolumeBackupDetails VolumeSourceFromVolumeBackupDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeVolumeSourceFromVolumeBackupDetails
	}{
		"volumeBackup",
		(MarshalTypeVolumeSourceFromVolumeBackupDetails)(m),
	}

	return json.Marshal(&s)
}
