// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

type CreateApiKeyDetails struct {

	// The public key.  Must be an RSA key in PEM format.
	Key *string `mandatory:"true" json:"key,omitempty"`
}

func (model CreateApiKeyDetails) String() string {
	return common.PointerString(model)
}
