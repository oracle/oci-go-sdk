// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

type AddUserToGroupDetails struct {

	// The OCID of the user.
	UserID string `mandatory:"true" json:"userId,omitempty"`

	// The OCID of the group.
	GroupID string `mandatory:"true" json:"groupId,omitempty"`
}
