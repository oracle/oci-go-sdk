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

//User An individual employee or system that needs to manage or use your company's Oracle Bare Metal Cloud Services\nresources. Users might need to launch instances, manage remote disks, work with your cloud network, etc. Users\nhave one or more IAM Service credentials ([ApiKey](#/en/identity/20160918/ApiKey/),\n[UIPassword](#/en/identity/20160918/UIPassword/), and [SwiftPassword](#/en/identity/20160918/SwiftPassword/)).\nFor more information, see [User Credentials](/Content/API/Concepts/usercredentials.htm)). End users of your\napplication are not typically IAM Service users. For conceptual information about users and other IAM Service\ncomponents, see [Overview of the IAM Service](/Content/Identity/Concepts/overview.htm).\n\nThese users are created directly within the Oracle Bare Metal Cloud Services system, via the IAM service.\nThey are different from *federated users*, who authenticate themselves to the Oracle Bare Metal\nCloud Services Console via an identity provider. For more information, see\n[Identity Providers and Federation](/Content/Identity/Concepts/federation.htm).\n\nTo use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,\ntalk to an administrator. If you're an administrator who needs to write policies to give users access,\nsee [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).\n

type User struct {

	// The OCID of the user.
	Id string `json:"id,omitempty"`

	// The OCID of the tenancy containing the user.
	CompartmentId string `json:"compartmentId,omitempty"`

	// The name you assign to the user during creation. This is the user's login for the Console.\nThe name must be unique across all users in the tenancy and cannot be changed.\n
	Name string `json:"name,omitempty"`

	// The description you assign to the user. Does not have to be unique, and it's changeable.
	Description string `json:"description,omitempty"`

	// Date and time the user was created, in the format defined by RFC3339.\n\nExample: `2016-08-25T21:10:29.600Z`\n
	TimeCreated time.Time `json:"timeCreated,omitempty"`

	// The user's current state. After creating a user, make sure its `lifecycleState` changes from CREATING to\nACTIVE before using it.\n
	LifecycleState string `json:"lifecycleState,omitempty"`

	// Returned only if the user's `lifecycleState` is INACTIVE. A 16-bit value showing the reason why the user\nis inactive:\n\n- bit 0: SUSPENDED (reserved for future use)\n- bit 1: DISABLED (reserved for future use)\n- bit 2: BLOCKED (the user has exceeded the maximum number of failed login attempts for the Console)\n
	InactiveStatus int64 `json:"inactiveStatus,omitempty"`
}
