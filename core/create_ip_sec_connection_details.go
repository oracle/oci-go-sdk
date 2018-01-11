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

type CreateIpSecConnectionDetails struct {

	// The OCID of the compartment to contain the IPSec connection.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The OCID of the CPE.
	CpeID *string `mandatory:"true" json:"cpeId,omitempty"`

	// The OCID of the DRG.
	DrgID *string `mandatory:"true" json:"drgId,omitempty"`

	// Static routes to the CPE. At least one route must be included. The CIDR must not be a
	// multicast address or class E address.
	// Example: `10.0.1.0/24`
	StaticRoutes []string `mandatory:"true" json:"staticRoutes,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

func (m CreateIpSecConnectionDetails) String() string {
	return common.PointerString(m)
}
