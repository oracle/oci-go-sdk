// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// GetGroupRequest wrapper for the GetGroup operation
type GetGroupRequest struct {

	// The OCID of the group.
	GroupID string `mandatory:"true" contributesTo:"path" name:"groupId"`
}

// GetGroupResponse wrapper for the GetGroup operation
type GetGroupResponse struct {

	// The Group instance
	Group `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string `presentIn:"header" name:"opcrequestid"`

	// For optimistic concurrency control. See `if-match`.
	Etag string `presentIn:"header" name:"etag"`
}
