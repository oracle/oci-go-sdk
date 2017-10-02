// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

type ModelError struct {

	// A short error code that defines the error, meant for programmatic parsing. See
	// [API Errors](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/API/References/apierrors.htm).
	Code string `mandatory:"true" json:"code,omitempty"`

	// A human-readable error string.
	Message string `mandatory:"true" json:"message,omitempty"`
}
