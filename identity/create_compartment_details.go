// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

type CreateCompartmentDetails struct {

	// The OCID of the tenancy containing the compartment.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The name you assign to the compartment during creation. The name must be unique across all compartments
	// in the tenancy.
	Name *string `mandatory:"true" json:"name,omitempty"`

	// The description you assign to the compartment during creation. Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description,omitempty"`
}

func (m CreateCompartmentDetails) String() string {
	return common.PointerString(m)
}
