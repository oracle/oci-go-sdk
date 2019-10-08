// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TransferAppliancePublicKey The representation of TransferAppliancePublicKey
type TransferAppliancePublicKey struct {
	PublicKey *string `mandatory:"false" json:"publicKey"`
}

func (m TransferAppliancePublicKey) String() string {
	return common.PointerString(m)
}
