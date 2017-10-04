// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// UpdateCompartmentRequest wrapper for the UpdateCompartment operation
type UpdateCompartmentRequest struct {

	// The OCID of the compartment.
	CompartmentID string `mandatory:"true" contributesTo:"path"`

	// Request object for updating a compartment.
	UpdateCompartmentDetails UpdateCompartmentDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

// UpdateCompartmentResponse wrapper for the UpdateCompartment operation
type UpdateCompartmentResponse struct {

	// The Compartment instance
	Compartment

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string

	// For optimistic concurrency control. See `if-match`.
	Etag string
}
