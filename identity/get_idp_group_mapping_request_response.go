// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// GetIdpGroupMappingRequest wrapper for the GetIdpGroupMapping operation
type GetIdpGroupMappingRequest struct {

	// The OCID of the identity provider.
	IdentityProviderID string `mandatory:"true" contributesTo:"path"`

	// The OCID of the group mapping.
	MappingID string `mandatory:"true" contributesTo:"path"`
}

// GetIdpGroupMappingResponse wrapper for the GetIdpGroupMapping operation
type GetIdpGroupMappingResponse struct {

	// The IdpGroupMapping instance
	IdpGroupMapping

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string

	// For optimistic concurrency control. See `if-match`.
	Etag string
}
