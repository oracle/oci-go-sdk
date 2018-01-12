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

type AttachBootVolumeDetails struct {

	// The OCID of the  boot volume.
	BootVolumeID *string `mandatory:"true" json:"bootVolumeId,omitempty"`

	// The OCID of the instance.
	InstanceID *string `mandatory:"true" json:"instanceId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

func (m AttachBootVolumeDetails) String() string {
	return common.PointerString(m)
}
