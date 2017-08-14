// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// Request wrapper for the GetUserGroupMembership operation
type GetUserGroupMembershipRequest struct {
        
 // The OCID of the userGroupMembership. 
        UserGroupMembershipID string
}

// Response wrapper for the GetUserGroupMembership operation
type GetUserGroupMembershipResponse struct {
        
 // Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
 // particular request, please provide the request ID.
        OpcRequestID string
        
 // For optimistic concurrency control. See `if-match`.
        Etag string

        
 // The UserGroupMembership instance
        GetUserGroupMembership UserGroupMembership


}


