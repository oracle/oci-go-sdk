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

// DetachDevicesDetails The representation of DetachDevicesDetails
type DetachDevicesDetails struct {

	// List of TransferDeviceLabel's
	DeviceLabels []string `mandatory:"false" json:"deviceLabels"`
}

func (m DetachDevicesDetails) String() string {
	return common.PointerString(m)
}
