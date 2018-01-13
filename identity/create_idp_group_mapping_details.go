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

// CreateIdpGroupMappingDetails The representation of CreateIdpGroupMappingDetails
type CreateIdpGroupMappingDetails struct {

	// The name of the IdP group you want to map.
	IdpGroupName *string `mandatory:"true" json:"idpGroupName,omitempty"`

	// The OCID of the IAM Service Group
	// you want to map to the IdP group.
	GroupID *string `mandatory:"true" json:"groupId,omitempty"`
}

func (m CreateIdpGroupMappingDetails) String() string {
	return common.PointerString(m)
}
