// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

type CreateIdpGroupMappingDetails struct {

	// The name of the IdP group you want to map.
	IdpGroupName string `json:"idpGroupName,omitempty"`

	// The OCID of the IAM Service [group](#/en/identity/20160918/Group/)
	// you want to map to the IdP group.
	GroupID string `json:"groupId,omitempty"`
}
