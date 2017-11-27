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

// Compartment. A collection of related resources. Compartments are a fundamental component of Oracle Cloud Infrastructure
// for organizing and isolating your cloud resources. You use them to clearly separate resources for the purposes
// of measuring usage and billing, access (through the use of IAM Service policies), and isolation (separating the
// resources for one project or business unit from another). A common approach is to create a compartment for each
// major part of your organization. For more information, see
// [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm) and also
// [Setting Up Your Tenancy]({{DOC_SERVER_URL}}/Content/GSG/Concepts/settinguptenancy.htm).
//
// To place a resource in a compartment, simply specify the compartment ID in the "Create" request object when
// initially creating the resource. For example, to launch an instance into a particular compartment, specify
// that compartment's OCID in the `LaunchInstance` request. You can't move an existing resource from one
// compartment to another.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type Compartment struct {

	// The OCID of the compartment.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The OCID of the tenancy containing the compartment.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The name you assign to the compartment during creation. The name must be unique across all
	// compartments in the tenancy.
	Name *string `mandatory:"true" json:"name,omitempty"`

	// The description you assign to the compartment. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description,omitempty"`

	// Date and time the compartment was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The compartment's current state. After creating a compartment, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	LifecycleState CompartmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus *int `mandatory:"false" json:"inactiveStatus,omitempty"`
}

func (model Compartment) String() string {
	return common.PointerString(model)
}

type CompartmentLifecycleStateEnum string

const (
	COMPARTMENT_LIFECYCLE_STATE_CREATING CompartmentLifecycleStateEnum = "CREATING"
	COMPARTMENT_LIFECYCLE_STATE_ACTIVE   CompartmentLifecycleStateEnum = "ACTIVE"
	COMPARTMENT_LIFECYCLE_STATE_INACTIVE CompartmentLifecycleStateEnum = "INACTIVE"
	COMPARTMENT_LIFECYCLE_STATE_DELETING CompartmentLifecycleStateEnum = "DELETING"
	COMPARTMENT_LIFECYCLE_STATE_DELETED  CompartmentLifecycleStateEnum = "DELETED"
	COMPARTMENT_LIFECYCLE_STATE_UNKNOWN  CompartmentLifecycleStateEnum = "UNKNOWN"
)

var mapping_compartment_lifecycleState = map[string]CompartmentLifecycleStateEnum{
	"CREATING": COMPARTMENT_LIFECYCLE_STATE_CREATING,
	"ACTIVE":   COMPARTMENT_LIFECYCLE_STATE_ACTIVE,
	"INACTIVE": COMPARTMENT_LIFECYCLE_STATE_INACTIVE,
	"DELETING": COMPARTMENT_LIFECYCLE_STATE_DELETING,
	"DELETED":  COMPARTMENT_LIFECYCLE_STATE_DELETED,
	"UNKNOWN":  COMPARTMENT_LIFECYCLE_STATE_UNKNOWN,
}

func GetCompartmentLifecycleStateEnumValues() []CompartmentLifecycleStateEnum {
	values := make([]CompartmentLifecycleStateEnum, 0)
	for _, v := range mapping_compartment_lifecycleState {
		if v != COMPARTMENT_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
