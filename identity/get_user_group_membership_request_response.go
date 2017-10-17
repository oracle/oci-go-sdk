// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// GetUserGroupMembershipRequest wrapper for the GetUserGroupMembership operation
type GetUserGroupMembershipRequest struct {

	// The OCID of the userGroupMembership.
	UserGroupMembershipID string `mandatory:"true" contributesTo:"path" name:"userGroupMembershipId"`
}

// GetUserGroupMembershipResponse wrapper for the GetUserGroupMembership operation
type GetUserGroupMembershipResponse struct {

	// The UserGroupMembership instance
	UserGroupMembership `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string `presentIn:"header" name:"opcrequestid"`

	// For optimistic concurrency control. See `if-match`.
	Etag string `presentIn:"header" name:"etag"`
}
