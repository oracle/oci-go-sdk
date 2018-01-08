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

// CreateInstanceConsoleConnectionDetails. The details for creating a instance console connection.
// The instance console connection is created in the same compartment as the instance.
type CreateInstanceConsoleConnectionDetails struct {

	// The OCID of the instance to create the console connection to.
	InstanceID *string `mandatory:"true" json:"instanceId,omitempty"`

	// The SSH public key used to authenticate the console connection.
	PublicKey *string `mandatory:"true" json:"publicKey,omitempty"`
}

func (model CreateInstanceConsoleConnectionDetails) String() string {
	return common.PointerString(model)
}
