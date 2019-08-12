// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DTS API
//
// A description of the DTS API
//

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
)

// MultipleTransferPackages The representation of MultipleTransferPackages
type MultipleTransferPackages struct {

	// List of TransferPackage summary's
	TransferPackageObjects []TransferPackageSummary `mandatory:"false" json:"transferPackageObjects"`
}

func (m MultipleTransferPackages) String() string {
	return common.PointerString(m)
}
