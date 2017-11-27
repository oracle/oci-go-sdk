// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oci-go-sdk/common"
)

// CreateInstanceConsoleConnectionDetails. The details for creating a serial console connection.
// The serial console connection is created in the same compartment as the instance.
type CreateInstanceConsoleConnectionDetails struct {

	// The OCID of the instance to create the serial console connection to.
	InstanceID *string `mandatory:"true" json:"instanceId,omitempty"`

	// The SSH public key used to authenticate the serial console connection.
	PublicKey *string `mandatory:"true" json:"publicKey,omitempty"`
}

func (model CreateInstanceConsoleConnectionDetails) String() string {
	return common.PointerString(model)
}
