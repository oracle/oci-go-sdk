// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// Request wrapper for the GetCompartment operation
type GetCompartmentRequest struct {

	// The OCID of the compartment.
	CompartmentID string
}

// Response wrapper for the GetCompartment operation
type GetCompartmentResponse struct {

	// The Compartment instance
	Compartment

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string

	// For optimistic concurrency control. See `if-match`.
	Etag string
}
