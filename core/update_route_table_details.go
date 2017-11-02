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

type UpdateRouteTableDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// The collection of rules used for routing destination IPs to network devices.
	RouteRules *[]RouteRule `mandatory:"false" json:"routeRules,omitempty"`
}

func (model UpdateRouteTableDetails) String() string {
	return common.PointerString(model)
}
