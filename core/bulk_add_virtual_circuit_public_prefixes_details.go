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

type BulkAddVirtualCircuitPublicPrefixesDetails struct {

	// The public IP prefixes (CIDRs) to add to the public virtual circuit.
	PublicPrefixes []CreateVirtualCircuitPublicPrefixDetails `mandatory:"true" json:"publicPrefixes,omitempty"`
}

func (model BulkAddVirtualCircuitPublicPrefixesDetails) String() string {
	return common.PointerString(model)
}
