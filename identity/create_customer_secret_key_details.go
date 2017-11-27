// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oci-go-sdk/common"
)

type CreateCustomerSecretKeyDetails struct {

	// The name you assign to the secret key during creation. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName,omitempty"`
}

func (model CreateCustomerSecretKeyDetails) String() string {
	return common.PointerString(model)
}
