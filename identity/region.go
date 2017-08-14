// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

//Region A localized geographic area, such as Phoenix, AZ. Oracle Bare Metal Cloud Services is hosted in regions and Availability\nDomains. A region is composed of several Availability Domains. An Availability Domain is one or more data centers\nlocated within a region. For more information, see [Regions and Availability Domains](/Content/General/Concepts/regions.htm).\n\nTo use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,\ntalk to an administrator. If you're an administrator who needs to write policies to give users access,\nsee [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).\n

type Region struct {

	// The key of the region.\n\nAllowed values are:\n- `PHX`\n- `IAD`\n
	Key string `json:"key,omitempty"`

	// The name of the region.\n\nAllowed values are:\n- `us-phoenix-1`\n- `us-ashburn-1`\n
	Name string `json:"name,omitempty"`
}
