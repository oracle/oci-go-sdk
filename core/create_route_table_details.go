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

// CreateRouteTableDetails The representation of CreateRouteTableDetails
type CreateRouteTableDetails struct {

	// The OCID of the compartment to contain the route table.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The collection of rules used for routing destination IPs to network devices.
	RouteRules []RouteRule `mandatory:"true" json:"routeRules,omitempty"`

	// The OCID of the VCN the route table belongs to.
	VcnID *string `mandatory:"true" json:"vcnId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

func (m CreateRouteTableDetails) String() string {
	return common.PointerString(m)
}
