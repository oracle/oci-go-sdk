// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"time"
)

// UiPassword. A text password that enables a user to sign in to the Console, the user interface for interacting with Oracle Bare
// Metal Cloud Services.
// For more information about user credentials, see [User Credentials](/Content/Identity/Concepts/usercredentials.htm).
type UiPassword struct {

	// The user's password for the Console.
	Password string `mandatory:"false" json:"password,omitempty"`

	// The OCID of the user.
	UserID string `mandatory:"false" json:"userId,omitempty"`

	// Date and time the password was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated time.Time `mandatory:"false" json:"timeCreated,omitempty"`

	// The password's current state. After creating a password, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState string `mandatory:"false" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus int64 `mandatory:"false" json:"inactiveStatus,omitempty"`
}
