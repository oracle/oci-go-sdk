// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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
