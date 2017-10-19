// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import "net/http"

// UpdatePolicyRequest wrapper for the UpdatePolicy operation
type UpdatePolicyRequest struct {

	// The OCID of the policy.
	PolicyID string `mandatory:"true" contributesTo:"path" name:"policyId"`

	// Request object for updating a policy.
	UpdatePolicyDetails UpdatePolicyDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

// UpdatePolicyResponse wrapper for the UpdatePolicy operation
type UpdatePolicyResponse struct {

	// The underlying http response
	RawResponse http.Response

	// The Policy instance
	Policy `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag string `presentIn:"header" name:"etag"`
}
