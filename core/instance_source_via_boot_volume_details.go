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

// InstanceSourceViaBootVolumeDetails
type InstanceSourceViaBootVolumeDetails struct {

	// The OCID of the boot volume used to boot the instance.
	BootVolumeID *string `mandatory:"true" json:"bootVolumeId,omitempty"`
}

func (m InstanceSourceViaBootVolumeDetails) String() string {
	return common.PointerString(m)
}

func (m InstanceSourceViaBootVolumeDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceSourceViaBootVolumeDetails InstanceSourceViaBootVolumeDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeInstanceSourceViaBootVolumeDetails
	}{
		"bootVolume",
		(MarshalTypeInstanceSourceViaBootVolumeDetails)(m),
	}

	return json.Marshal(&s)
}
