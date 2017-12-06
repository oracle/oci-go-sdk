// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

type CreateSecurityListDetails struct {

	// The OCID of the compartment to contain the security list.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// Rules for allowing egress IP packets.
	EgressSecurityRules []EgressSecurityRule `mandatory:"true" json:"egressSecurityRules,omitempty"`

	// Rules for allowing ingress IP packets.
	IngressSecurityRules []IngressSecurityRule `mandatory:"true" json:"ingressSecurityRules,omitempty"`

	// The OCID of the VCN the security list belongs to.
	VcnID *string `mandatory:"true" json:"vcnId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

func (model CreateSecurityListDetails) String() string {
	return common.PointerString(model)
}
