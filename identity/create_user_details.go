// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity


type CreateUserDetails struct {

    // The OCID of the tenancy containing the user.
    CompartmentId string `json:"compartmentId,omitempty"`

    // The name you assign to the user during creation. This is the user's login for the Console.\nThe name must be unique across all users in the tenancy and cannot be changed.\n
    Name string `json:"name,omitempty"`

    // The description you assign to the user during creation. Does not have to be unique, and it's changeable.
    Description string `json:"description,omitempty"`
}
