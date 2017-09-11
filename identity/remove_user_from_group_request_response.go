// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// RemoveUserFromGroupRequest wrapper for the RemoveUserFromGroup operation
type RemoveUserFromGroupRequest struct {

	// The OCID of the userGroupMembership.
	UserGroupMembershipID string `mandatory:"true" contributesTo:"path"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

// RemoveUserFromGroupResponse wrapper for the RemoveUserFromGroup operation
type RemoveUserFromGroupResponse struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string
}
