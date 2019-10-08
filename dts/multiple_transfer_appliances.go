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

// MultipleTransferAppliances The representation of MultipleTransferAppliances
type MultipleTransferAppliances struct {

	// List of TransferAppliance summary's
	TransferApplianceObjects []TransferApplianceSummary `mandatory:"false" json:"transferApplianceObjects"`
}

func (m MultipleTransferAppliances) String() string {
	return common.PointerString(m)
}
