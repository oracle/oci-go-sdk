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

// AttachDevicesDetails The representation of AttachDevicesDetails
type AttachDevicesDetails struct {

	// List of TransferDeviceLabel's
	DeviceLabels []string `mandatory:"false" json:"deviceLabels"`
}

func (m AttachDevicesDetails) String() string {
	return common.PointerString(m)
}
