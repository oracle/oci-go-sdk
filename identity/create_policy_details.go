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

type CreatePolicyDetails struct {

	// The OCID of the compartment containing the policy (either the tenancy or another compartment).
	CompartmentID string `json:"compartmentId,omitempty"`

	// The name you assign to the policy during creation. The name must be unique across all policies
	// in the tenancy and cannot be changed.
	Name string `json:"name,omitempty"`

	// An array of policy statements written in the policy language. See
	// [How Policies Work](/Content/Identity/Concepts/policies.htm) and
	// [Common Policies](/Content/Identity/Concepts/commonpolicies.htm).
	Statements []string `json:"statements,omitempty"`

	// The description you assign to the policy during creation. Does not have to be unique, and it's changeable.
	Description string `json:"description,omitempty"`

	// The version of the policy. If null or set to an empty string, when a request comes in for authorization, the
	// policy will be evaluated according to the current behavior of the services at that moment. If set to a particular
	// date (YYYY-MM-DD), the policy will be evaluated according to the behavior of the services on that date.
	VersionDate time.Time `json:"versionDate,omitempty"`
}
