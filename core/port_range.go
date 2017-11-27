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

type PortRange struct {

	// The maximum port number. Must not be lower than the minimum port number. To specify
	// a single port number, set both the min and max to the same value.
	Max *int `mandatory:"true" json:"max,omitempty"`

	// The minimum port number. Must not be greater than the maximum port number.
	Min *int `mandatory:"true" json:"min,omitempty"`
}

func (model PortRange) String() string {
	return common.PointerString(model)
}
