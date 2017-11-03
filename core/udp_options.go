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

// UdpOptions. Optional object to specify ports for a UDP rule. If you specify UDP as the
// protocol but omit this object, then all ports are allowed.
type UdpOptions struct {

	// An inclusive range of allowed destination ports. Use the same number for the min and max
	// to indicate a single port. Defaults to all ports if not specified.
	DestinationPortRange *PortRange `mandatory:"false" json:"destinationPortRange,omitempty"`

	// An inclusive range of allowed source ports. Use the same number for the min and max to
	// indicate a single port. Defaults to all ports if not specified.
	SourcePortRange *PortRange `mandatory:"false" json:"sourcePortRange,omitempty"`
}

func (model UdpOptions) String() string {
	return common.PointerString(model)
}
