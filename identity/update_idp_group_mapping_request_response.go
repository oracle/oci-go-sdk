// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// Request wrapper for the UpdateIdpGroupMapping operation
type UpdateIdpGroupMappingRequest struct {
        
 // The OCID of the identity provider. 
        IdentityProviderID string
        
 // The OCID of the group mapping. 
        MappingID string
        
 // Request object for updating an identity provider group mapping 
        UpdateIdpGroupMappingDetails UpdateIdpGroupMappingDetails
        
 // For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
 // parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
 // will be updated or deleted only if the etag you provide matches the resource's current etag value. 
        IfMatch string
}

// Response wrapper for the UpdateIdpGroupMapping operation
type UpdateIdpGroupMappingResponse struct {
        
 // Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
 // particular request, please provide the request ID.
        OpcRequestID string
        
 // For optimistic concurrency control. See `if-match`.
        Etag string

        
 // The IdpGroupMapping instance
        UpdateIdpGroupMapping IdpGroupMapping


}


