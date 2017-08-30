// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// Request wrapper for the GetTenancy operation
type GetTenancyRequest struct {

	// The OCID of the tenancy.
	TenancyID string
}

// Response wrapper for the GetTenancy operation
type GetTenancyResponse struct {

	// The Tenancy instance
	Tenancy

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string
}
