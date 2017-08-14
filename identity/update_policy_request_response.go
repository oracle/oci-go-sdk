// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// Request wrapper for the UpdatePolicy operation
type UpdatePolicyRequest struct {

	// The OCID of the policy.
	PolicyID string

	// Request object for updating a policy.
	UpdatePolicyDetails UpdatePolicyDetails

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch string
}

// Response wrapper for the UpdatePolicy operation
type UpdatePolicyResponse struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string

	// For optimistic concurrency control. See `if-match`.
	Etag string

	// The Policy instance
	UpdatePolicy Policy
}
