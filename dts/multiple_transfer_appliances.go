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

// MultipleTransferAppliances The representation of MultipleTransferAppliances
type MultipleTransferAppliances struct {

	// List of TransferAppliance summary's
	TransferApplianceObjects []TransferApplianceSummary `mandatory:"false" json:"transferApplianceObjects"`
}

func (m MultipleTransferAppliances) String() string {
	return common.PointerString(m)
}
