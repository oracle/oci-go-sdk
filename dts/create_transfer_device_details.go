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

// CreateTransferDeviceDetails The representation of CreateTransferDeviceDetails
type CreateTransferDeviceDetails struct {
	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	IscsiIQN *string `mandatory:"false" json:"iscsiIQN"`
}

func (m CreateTransferDeviceDetails) String() string {
	return common.PointerString(m)
}
