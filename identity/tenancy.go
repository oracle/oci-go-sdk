// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

// Tenancy. The root compartment that contains all of your organization's compartments and other
// Oracle Bare Metal Cloud Services cloud resources. When you sign up for Oracle Bare Metal Cloud Services,
// Oracle creates a tenancy for your company, which is a secure and isolated partition
// where you can create, organize, and administer your cloud resources.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type Tenancy struct {

	// The OCID of the tenancy.
	ID string `mandatory:"false" json:"id,omitempty"`

	// The name of the tenancy.
	Name string `mandatory:"false" json:"name,omitempty"`

	// The description of the tenancy.
	Description string `mandatory:"false" json:"description,omitempty"`

	// The region key for the tenancy's home region.
	// Allowed values are:
	// - `IAD`
	// - `PHX`
	HomeRegionKey string `mandatory:"false" json:"homeRegionKey,omitempty"`
}
