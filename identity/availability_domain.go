// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

// AvailabilityDomain. One or more isolated, fault-tolerant Oracle data centers that host cloud resources such as instances, volumes,
// and subnets. A region contains several Availability Domains. For more information, see
// [Regions and Availability Domains](/Content/General/Concepts/regions.htm).

type AvailabilityDomain struct {

	// The name of the Availability Domain.
	Name string `json:"name,omitempty"`

	// The OCID of the tenancy.
	CompartmentId string `json:"compartmentId,omitempty"`
}
