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

type BulkDeleteVirtualCircuitPublicPrefixesDetails struct {

	// The public IP prefixes (CIDRs) to remove from the public virtual circuit.
	PublicPrefixes []DeleteVirtualCircuitPublicPrefixDetails `mandatory:"true" json:"publicPrefixes,omitempty"`
}

func (m BulkDeleteVirtualCircuitPublicPrefixesDetails) String() string {
	return common.PointerString(m)
}
