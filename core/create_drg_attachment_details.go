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

type CreateDrgAttachmentDetails struct {

	// The OCID of the DRG.
	DrgID *string `mandatory:"true" json:"drgId,omitempty"`

	// The OCID of the VCN.
	VcnID *string `mandatory:"true" json:"vcnId,omitempty"`

	// A user-friendly name. Does not have to be unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

func (model CreateDrgAttachmentDetails) String() string {
	return common.PointerString(model)
}
