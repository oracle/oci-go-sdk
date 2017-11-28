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

type UpdateIdpGroupMappingDetails struct {

	// The idp group name.
	IdpGroupName *string `mandatory:"false" json:"idpGroupName,omitempty"`

	// The OCID of the group.
	GroupID *string `mandatory:"false" json:"groupId,omitempty"`
}

func (model UpdateIdpGroupMappingDetails) String() string {
	return common.PointerString(model)
}
