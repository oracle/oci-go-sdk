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

type AttachVnicDetails struct {

	// Details for creating a new VNIC.
	CreateVnicDetails *CreateVnicDetails `mandatory:"true" json:"createVnicDetails,omitempty"`

	// The OCID of the instance.
	InstanceID *string `mandatory:"true" json:"instanceId,omitempty"`

	// A user-friendly name for the attachment. Does not have to be unique, and it cannot be changed.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

func (model AttachVnicDetails) String() string {
	return common.PointerString(model)
}
