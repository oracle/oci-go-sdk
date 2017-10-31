// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import "net/http"

// GetPolicyRequest wrapper for the GetPolicy operation
type GetPolicyRequest struct {

	// The OCID of the policy.
	PolicyID *string `mandatory:"true" contributesTo:"path" name:"policyId"`
}

// GetPolicyResponse wrapper for the GetPolicy operation
type GetPolicyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Policy instance
	Policy `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}
