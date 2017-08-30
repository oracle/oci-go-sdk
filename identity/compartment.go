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

// Compartment. A collection of related resources. Compartments are a fundamental component of Oracle Bare Metal Cloud Services
// for organizing and isolating your cloud resources. You use them to clearly separate resources for the purposes
// of measuring usage and billing, access (through the use of IAM Service policies), and isolation (separating the
// resources for one project or business unit from another). A common approach is to create a compartment for each
// major part of your organization. For more information, see
// [Overview of the IAM Service](/Content/Identity/Concepts/overview.htm) and also
// [Setting Up Your Tenancy](/Content/GSG/Concepts/settinguptenancy.htm).
// To place a resource in a compartment, simply specify the compartment ID in the "Create" request object when
// initially creating the resource. For example, to launch an instance into a particular compartment, specify
// that compartment's OCID in the `LaunchInstance` request. You can't move an existing resource from one
// compartment to another.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).
type Compartment struct {

	// The OCID of the compartment.
	ID string `json:"id,omitempty"`

	// The OCID of the tenancy containing the compartment.
	CompartmentID string `json:"compartmentId,omitempty"`

	// The name you assign to the compartment during creation. The name must be unique across all
	// compartments in the tenancy and cannot be changed.
	Name string `json:"name,omitempty"`

	// The description you assign to the compartment. Does not have to be unique, and it's changeable.
	Description string `json:"description,omitempty"`

	// Date and time the compartment was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated time.Time `json:"timeCreated,omitempty"`

	// The compartment's current state. After creating a compartment, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState string `json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus int64 `json:"inactiveStatus,omitempty"`
}
