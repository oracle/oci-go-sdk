// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// UserGroupMembership. An object that represents the membership of a user in a group. When you add a user to a group, the result is a
// `UserGroupMembership` with its own OCID. To remove a user from a group, you delete the `UserGroupMembership` object.
type UserGroupMembership struct {

	// The OCID of the membership.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The OCID of the tenancy containing the user, group, and membership object.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The OCID of the group.
	GroupID *string `mandatory:"true" json:"groupId,omitempty"`

	// The OCID of the user.
	UserID *string `mandatory:"true" json:"userId,omitempty"`

	// Date and time the membership was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The membership's current state.  After creating a membership object, make sure its `lifecycleState` changes
	// from CREATING to ACTIVE before using it.
	LifecycleState *string `mandatory:"true" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int64 `mandatory:"false" json:"inactiveStatus,omitempty"`
}
