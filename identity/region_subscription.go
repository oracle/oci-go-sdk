// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

//RegionSubscription An object that represents your tenancy's access to a particular region (i.e., a subscription), the status of that\naccess, and whether that region is the home region. For more information, see [Managing Regions](/Content/Identity/Tasks/managingregions.htm).\n\nTo use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,\ntalk to an administrator. If you're an administrator who needs to write policies to give users access,\nsee [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).\n

type RegionSubscription struct {

	// The region's key.\n\nAllowed values are:\n- `PHX`\n- `IAD`\n
	RegionKey string `json:"regionKey,omitempty"`

	// The region's name.\n\nAllowed values are:\n- `us-phoenix-1`\n- `us-ashburn-1`\n
	RegionName string `json:"regionName,omitempty"`

	// The region subscription status.
	Status string `json:"status,omitempty"`

	// Indicates if the region is the home region or not.
	IsHomeRegion bool `json:"isHomeRegion,omitempty"`
}
