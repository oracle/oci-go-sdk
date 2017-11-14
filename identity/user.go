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

// User. An individual employee or system that needs to manage or use your company's Oracle Cloud Infrastructure
// resources. Users might need to launch instances, manage remote disks, work with your cloud network, etc. Users
// have one or more IAM Service credentials (ApiKey,
// UIPassword, and SwiftPassword).
// For more information, see [User Credentials]({{DOC_SERVER_URL}}/Content/API/Concepts/usercredentials.htm)). End users of your
// application are not typically IAM Service users. For conceptual information about users and other IAM Service
// components, see [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm).
// These users are created directly within the Oracle Cloud Infrastructure system, via the IAM service.
// They are different from *federated users*, who authenticate themselves to the Oracle Cloud Infrastructure
// Console via an identity provider. For more information, see
// [Identity Providers and Federation]({{DOC_SERVER_URL}}/Content/Identity/Concepts/federation.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type User struct {

	// The OCID of the user.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The OCID of the tenancy containing the user.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The name you assign to the user during creation. This is the user's login for the Console.
	// The name must be unique across all users in the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name,omitempty"`

	// The description you assign to the user. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description,omitempty"`

	// Date and time the user was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The user's current state. After creating a user, make sure its `lifecycleState` changes from CREATING to
	// ACTIVE before using it.
	LifecycleState UserLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// Returned only if the user's `lifecycleState` is INACTIVE. A 16-bit value showing the reason why the user
	// is inactive:
	// - bit 0: SUSPENDED (reserved for future use)
	// - bit 1: DISABLED (reserved for future use)
	// - bit 2: BLOCKED (the user has exceeded the maximum number of failed login attempts for the Console)
	InactiveStatus *int `mandatory:"false" json:"inactiveStatus,omitempty"`
}

func (model User) String() string {
	return common.PointerString(model)
}

type UserLifecycleStateEnum string

const (
	USER_LIFECYCLE_STATE_CREATING UserLifecycleStateEnum = "CREATING"
	USER_LIFECYCLE_STATE_ACTIVE   UserLifecycleStateEnum = "ACTIVE"
	USER_LIFECYCLE_STATE_INACTIVE UserLifecycleStateEnum = "INACTIVE"
	USER_LIFECYCLE_STATE_DELETING UserLifecycleStateEnum = "DELETING"
	USER_LIFECYCLE_STATE_DELETED  UserLifecycleStateEnum = "DELETED"
	USER_LIFECYCLE_STATE_UNKNOWN  UserLifecycleStateEnum = "UNKNOWN"
)

var mapping_user_lifecycleState = map[string]UserLifecycleStateEnum{
	"CREATING": USER_LIFECYCLE_STATE_CREATING,
	"ACTIVE":   USER_LIFECYCLE_STATE_ACTIVE,
	"INACTIVE": USER_LIFECYCLE_STATE_INACTIVE,
	"DELETING": USER_LIFECYCLE_STATE_DELETING,
	"DELETED":  USER_LIFECYCLE_STATE_DELETED,
	"UNKNOWN":  USER_LIFECYCLE_STATE_UNKNOWN,
}

func GetUserLifecycleStateEnumValues() []UserLifecycleStateEnum {
	values := make([]UserLifecycleStateEnum, 0)
	for _, v := range mapping_user_lifecycleState {
		if v != USER_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
