// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

// RegionSubscription. An object that represents your tenancy's access to a particular region (i.e., a subscription), the status of that
// access, and whether that region is the home region. For more information, see [Managing Regions](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/Identity/Tasks/managingregions.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see [Getting Started with Policies](http://lgl-bybliothece-01.virt.lgl.grungy.us/Content/Identity/Concepts/policygetstarted.htm).
type RegionSubscription struct {

	// The region's key.
	// Allowed values are:
	// - `PHX`
	// - `IAD`
	RegionKey string `mandatory:"true" json:"regionKey,omitempty"`

	// The region's name.
	// Allowed values are:
	// - `us-phoenix-1`
	// - `us-ashburn-1`
	RegionName string `mandatory:"true" json:"regionName,omitempty"`

	// The region subscription status.
	Status string `mandatory:"true" json:"status,omitempty"`

	// Indicates if the region is the home region or not.
	IsHomeRegion bool `mandatory:"true" json:"isHomeRegion,omitempty"`
}
