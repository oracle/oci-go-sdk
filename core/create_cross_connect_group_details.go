// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

type CreateCrossConnectGroupDetails struct {

	// The OCID of the compartment to contain the cross-connect group.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

func (model CreateCrossConnectGroupDetails) String() string {
	return common.PointerString(model)
}
