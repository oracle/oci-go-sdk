// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Group A collection of users who all need the same type of access to a particular set of resources or compartment.
// For conceptual information about groups and other IAM Service components, see
// https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm.
// If you're federating with an identity provider (IdP), you need to create mappings between the groups
// defined in the IdP and groups you define in the IAM service. For more information, see
// https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/federation.htm. Also see
// IdentityProvider and
// IdpGroupMapping.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm.
type Group struct {

	// The OCID of the group.
	Id *string `mandatory:"true" json:"id,omitempty"`

	// The OCID of the tenancy containing the group.
	CompartmentId *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The name you assign to the group during creation. The name must be unique across all groups in
	// the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name,omitempty"`

	// The description you assign to the group. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description,omitempty"`

	// Date and time the group was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The group's current state. After creating a group, make sure its `lifecycleState` changes from CREATING to
	// ACTIVE before using it.
	LifecycleState GroupLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int `mandatory:"false" json:"inactiveStatus,omitempty"`
}

func (m Group) String() string {
	return common.PointerString(m)
}

// GroupLifecycleStateEnum Enum with underlying type: string
type GroupLifecycleStateEnum string

// Set of constants representing the allowable values for GroupLifecycleState
const (
	GroupLifecycleStateCreating GroupLifecycleStateEnum = "CREATING"
	GroupLifecycleStateActive   GroupLifecycleStateEnum = "ACTIVE"
	GroupLifecycleStateInactive GroupLifecycleStateEnum = "INACTIVE"
	GroupLifecycleStateDeleting GroupLifecycleStateEnum = "DELETING"
	GroupLifecycleStateDeleted  GroupLifecycleStateEnum = "DELETED"
	GroupLifecycleStateUnknown  GroupLifecycleStateEnum = "UNKNOWN"
)

var mappingGroupLifecycleState = map[string]GroupLifecycleStateEnum{
	"CREATING": GroupLifecycleStateCreating,
	"ACTIVE":   GroupLifecycleStateActive,
	"INACTIVE": GroupLifecycleStateInactive,
	"DELETING": GroupLifecycleStateDeleting,
	"DELETED":  GroupLifecycleStateDeleted,
	"UNKNOWN":  GroupLifecycleStateUnknown,
}

// GetGroupLifecycleStateEnumValues Enumerates the set of values for GroupLifecycleState
func GetGroupLifecycleStateEnumValues() []GroupLifecycleStateEnum {
	values := make([]GroupLifecycleStateEnum, 0)
	for _, v := range mappingGroupLifecycleState {
		if v != GroupLifecycleStateUnknown {
			values = append(values, v)
		}
	}
	return values
}
