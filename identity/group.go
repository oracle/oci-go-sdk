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

//Group A collection of users who all need the same type of access to a particular set of resources or compartment.\nFor conceptual information about groups and other IAM Service components, see\n[Overview of the IAM Service](/Content/Identity/Concepts/overview.htm).\n\nIf you're federating with an identity provider (IdP), you need to create mappings between the groups\ndefined in the IdP and groups you define in the IAM service. For more information, see\n[Identity Providers and Federation](/Content/Identity/Concepts/federation.htm). Also see\n[IdentityProvider](#/en/identity/20160918/IdentityProvider/) and\n[IdpGroupMapping](#/en/identity/20160918/IdpGroupMapping/).\n\nTo use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,\ntalk to an administrator. If you're an administrator who needs to write policies to give users access,\nsee [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).\n

type Group struct {

	// The OCID of the group.
	Id string `json:"id,omitempty"`

	// The OCID of the tenancy containing the group.
	CompartmentId string `json:"compartmentId,omitempty"`

	// The name you assign to the group during creation. The name must be unique across all groups in\nthe tenancy and cannot be changed.\n
	Name string `json:"name,omitempty"`

	// The description you assign to the group. Does not have to be unique, and it's changeable.
	Description string `json:"description,omitempty"`

	// Date and time the group was created, in the format defined by RFC3339.\n\nExample: `2016-08-25T21:10:29.600Z`\n
	TimeCreated time.Time `json:"timeCreated,omitempty"`

	// The group's current state. After creating a group, make sure its `lifecycleState` changes from CREATING to\nACTIVE before using it.\n
	LifecycleState string `json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus int64 `json:"inactiveStatus,omitempty"`
}
