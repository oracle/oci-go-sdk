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

// PathRoute The representation of PathRoute
type PathRoute struct {

	// The name of the associated backend set.
	BackendSetName *string `mandatory:"true" json:"backendSetName,omitempty"`

	// The URI path on which to route requests.
	// Example: `/foo/bar/123`
	Path *string `mandatory:"true" json:"path,omitempty"`

	PathMatchType *PathMatchType `mandatory:"true" json:"pathMatchType,omitempty"`
}

func (m PathRoute) String() string {
	return common.PointerString(m)
}
