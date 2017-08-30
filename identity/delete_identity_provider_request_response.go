// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// DeleteIdentityProviderRequest wrapper for the DeleteIdentityProvider operation
type DeleteIdentityProviderRequest struct {

	// The OCID of the identity provider.
	IdentityProviderID string

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch string
}

// DeleteIdentityProviderResponse wrapper for the DeleteIdentityProvider operation
type DeleteIdentityProviderResponse struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string
}
