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


//Compartment A collection of related resources. Compartments are a fundamental component of Oracle Bare Metal Cloud Services\nfor organizing and isolating your cloud resources. You use them to clearly separate resources for the purposes\nof measuring usage and billing, access (through the use of IAM Service policies), and isolation (separating the\nresources for one project or business unit from another). A common approach is to create a compartment for each\nmajor part of your organization. For more information, see\n[Overview of the IAM Service](/Content/Identity/Concepts/overview.htm) and also\n[Setting Up Your Tenancy](/Content/GSG/Concepts/settinguptenancy.htm).\n\nTo place a resource in a compartment, simply specify the compartment ID in the \"Create\" request object when\ninitially creating the resource. For example, to launch an instance into a particular compartment, specify\nthat compartment's OCID in the `LaunchInstance` request. You can't move an existing resource from one\ncompartment to another.\n\nTo use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,\ntalk to an administrator. If you're an administrator who needs to write policies to give users access,\nsee [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).\n

type Compartment struct {

    // The OCID of the compartment.
    Id string `json:"id,omitempty"`

    // The OCID of the tenancy containing the compartment.
    CompartmentId string `json:"compartmentId,omitempty"`

    // The name you assign to the compartment during creation. The name must be unique across all\ncompartments in the tenancy and cannot be changed.\n
    Name string `json:"name,omitempty"`

    // The description you assign to the compartment. Does not have to be unique, and it's changeable.
    Description string `json:"description,omitempty"`

    // Date and time the compartment was created, in the format defined by RFC3339.\n\nExample: `2016-08-25T21:10:29.600Z`\n
    TimeCreated time.Time `json:"timeCreated,omitempty"`

    // The compartment's current state. After creating a compartment, make sure its `lifecycleState` changes from\nCREATING to ACTIVE before using it.\n
    LifecycleState string `json:"lifecycleState,omitempty"`

    // The detailed status of INACTIVE lifecycleState.
    InactiveStatus int64 `json:"inactiveStatus,omitempty"`
}
