// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// GetUserRequest wrapper for the GetUser operation
type GetUserRequest struct {

	// The OCID of the user.
	UserID string `mandatory:"true" contributesTo:"path" name:"userId"`
}

// GetUserResponse wrapper for the GetUser operation
type GetUserResponse struct {

	// The User instance
	User `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string `presentIn:"header" name:"opcrequestid"`

	// For optimistic concurrency control. See `if-match`.
	Etag string `presentIn:"header" name:"etag"`
}
