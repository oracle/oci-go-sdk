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

// IpSecConnectionDeviceStatus Status of the IPSec connection.
type IpSecConnectionDeviceStatus struct {

	// The OCID of the compartment containing the IPSec connection.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The IPSec connection's Oracle ID (OCID).
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The date and time the IPSec connection was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`

	// Two TunnelStatus objects.
	Tunnels []TunnelStatus `mandatory:"false" json:"tunnels,omitempty"`
}

func (m IpSecConnectionDeviceStatus) String() string {
	return common.PointerString(m)
}
