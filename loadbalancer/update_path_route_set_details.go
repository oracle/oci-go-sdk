// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

type UpdatePathRouteSetDetails struct {
	PathRoutes []PathRoute `mandatory:"true" json:"pathRoutes,omitempty"`
}

func (m UpdatePathRouteSetDetails) String() string {
	return common.PointerString(m)
}
