// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// UpdateSwiftPasswordRequest wrapper for the UpdateSwiftPassword operation
type UpdateSwiftPasswordRequest struct {

	// The OCID of the user.
	UserID string

	// The OCID of the Swift password.
	SwiftPasswordID string

	// Request object for updating a Swift password.
	UpdateSwiftPasswordDetails UpdateSwiftPasswordDetails

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch string
}

// UpdateSwiftPasswordResponse wrapper for the UpdateSwiftPassword operation
type UpdateSwiftPasswordResponse struct {

	// The SwiftPassword instance
	SwiftPassword

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string

	// For optimistic concurrency control. See `if-match`.
	Etag string
}
