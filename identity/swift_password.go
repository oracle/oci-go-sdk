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

// SwiftPassword. Swift is the OpenStack object storage service. A `SwiftPassword` is an Oracle-provided password for using a
// Swift client with the Oracle Cloud Infrastructure Object Storage Service. This password is associated with
// the user's Console login. Swift passwords never expire. A user can have up to two Swift passwords at a time.
// **Note:** The password is always an Oracle-generated string; you can't change it to a string of your choice.
// For more information, see [Managing User Credentials]({{DOC_SERVER_URL}}/Content/Identity/Tasks/managingcredentials.htm).
type SwiftPassword struct {

	// The Swift password. The value is available only in the response for `CreateSwiftPassword`, and not
	// for `ListSwiftPasswords` or `UpdateSwiftPassword`.
	Password *string `mandatory:"false" json:"password,omitempty"`

	// The OCID of the Swift password.
	ID *string `mandatory:"false" json:"id,omitempty"`

	// The OCID of the user the password belongs to.
	UserID *string `mandatory:"false" json:"userId,omitempty"`

	// The description you assign to the Swift password. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"false" json:"description,omitempty"`

	// Date and time the `SwiftPassword` object was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`

	// Date and time when this password will expire, in the format defined by RFC3339.
	// Null if it never expires.
	// Example: `2016-08-25T21:10:29.600Z`
	ExpiresOn *common.SDKTime `mandatory:"false" json:"expiresOn,omitempty"`

	// The password's current state. After creating a password, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState SwiftPasswordLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int `mandatory:"false" json:"inactiveStatus,omitempty"`
}

func (model SwiftPassword) String() string {
	return common.PointerString(model)
}

type SwiftPasswordLifecycleStateEnum string

const (
	SWIFT_PASSWORD_LIFECYCLE_STATE_CREATING SwiftPasswordLifecycleStateEnum = "CREATING"
	SWIFT_PASSWORD_LIFECYCLE_STATE_ACTIVE   SwiftPasswordLifecycleStateEnum = "ACTIVE"
	SWIFT_PASSWORD_LIFECYCLE_STATE_INACTIVE SwiftPasswordLifecycleStateEnum = "INACTIVE"
	SWIFT_PASSWORD_LIFECYCLE_STATE_DELETING SwiftPasswordLifecycleStateEnum = "DELETING"
	SWIFT_PASSWORD_LIFECYCLE_STATE_DELETED  SwiftPasswordLifecycleStateEnum = "DELETED"
	SWIFT_PASSWORD_LIFECYCLE_STATE_UNKNOWN  SwiftPasswordLifecycleStateEnum = "UNKNOWN"
)

var mapping_swiftpassword_lifecycleState = map[string]SwiftPasswordLifecycleStateEnum{
	"CREATING": SWIFT_PASSWORD_LIFECYCLE_STATE_CREATING,
	"ACTIVE":   SWIFT_PASSWORD_LIFECYCLE_STATE_ACTIVE,
	"INACTIVE": SWIFT_PASSWORD_LIFECYCLE_STATE_INACTIVE,
	"DELETING": SWIFT_PASSWORD_LIFECYCLE_STATE_DELETING,
	"DELETED":  SWIFT_PASSWORD_LIFECYCLE_STATE_DELETED,
	"UNKNOWN":  SWIFT_PASSWORD_LIFECYCLE_STATE_UNKNOWN,
}

func GetSwiftPasswordLifecycleStateEnumValues() []SwiftPasswordLifecycleStateEnum {
	values := make([]SwiftPasswordLifecycleStateEnum, 0)
	for _, v := range mapping_swiftpassword_lifecycleState {
		if v != SWIFT_PASSWORD_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
