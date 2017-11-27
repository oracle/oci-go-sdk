// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oci-go-sdk/common"
)

// UiPassword. A text password that enables a user to sign in to the Console, the user interface for interacting with Oracle
// Cloud Infrastructure.
// For more information about user credentials, see [User Credentials]({{DOC_SERVER_URL}}/Content/Identity/Concepts/usercredentials.htm).
type UiPassword struct {

	// The user's password for the Console.
	Password *string `mandatory:"false" json:"password,omitempty"`

	// The OCID of the user.
	UserID *string `mandatory:"false" json:"userId,omitempty"`

	// Date and time the password was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`

	// The password's current state. After creating a password, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState UiPasswordLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int `mandatory:"false" json:"inactiveStatus,omitempty"`
}

func (model UiPassword) String() string {
	return common.PointerString(model)
}

type UiPasswordLifecycleStateEnum string

const (
	UI_PASSWORD_LIFECYCLE_STATE_CREATING UiPasswordLifecycleStateEnum = "CREATING"
	UI_PASSWORD_LIFECYCLE_STATE_ACTIVE   UiPasswordLifecycleStateEnum = "ACTIVE"
	UI_PASSWORD_LIFECYCLE_STATE_INACTIVE UiPasswordLifecycleStateEnum = "INACTIVE"
	UI_PASSWORD_LIFECYCLE_STATE_DELETING UiPasswordLifecycleStateEnum = "DELETING"
	UI_PASSWORD_LIFECYCLE_STATE_DELETED  UiPasswordLifecycleStateEnum = "DELETED"
	UI_PASSWORD_LIFECYCLE_STATE_UNKNOWN  UiPasswordLifecycleStateEnum = "UNKNOWN"
)

var mapping_uipassword_lifecycleState = map[string]UiPasswordLifecycleStateEnum{
	"CREATING": UI_PASSWORD_LIFECYCLE_STATE_CREATING,
	"ACTIVE":   UI_PASSWORD_LIFECYCLE_STATE_ACTIVE,
	"INACTIVE": UI_PASSWORD_LIFECYCLE_STATE_INACTIVE,
	"DELETING": UI_PASSWORD_LIFECYCLE_STATE_DELETING,
	"DELETED":  UI_PASSWORD_LIFECYCLE_STATE_DELETED,
	"UNKNOWN":  UI_PASSWORD_LIFECYCLE_STATE_UNKNOWN,
}

func GetUiPasswordLifecycleStateEnumValues() []UiPasswordLifecycleStateEnum {
	values := make([]UiPasswordLifecycleStateEnum, 0)
	for _, v := range mapping_uipassword_lifecycleState {
		if v != UI_PASSWORD_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
