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

// IpSecConnectionDeviceConfig Information about the IPSecConnection device configuration.
type IpSecConnectionDeviceConfig struct {

	// The OCID of the compartment containing the IPSec connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The IPSec connection's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id,omitempty"`

	// The date and time the IPSec connection was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`

	// Two TunnelConfig objects.
	Tunnels []TunnelConfig `mandatory:"false" json:"tunnels,omitempty"`
}

func (m IpSecConnectionDeviceConfig) String() string {
	return common.PointerString(m)
}
