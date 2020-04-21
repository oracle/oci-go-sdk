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

// NewTransferDevice The representation of NewTransferDevice
type NewTransferDevice struct {
	Label *string `mandatory:"true" json:"label"`

	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	IscsiIQN *string `mandatory:"false" json:"iscsiIQN"`

	LifecycleState NewTransferDeviceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	EncryptionPassphrase *string `mandatory:"false" json:"encryptionPassphrase"`

	TransferJobId *string `mandatory:"false" json:"transferJobId"`

	CreationTime *common.SDKTime `mandatory:"false" json:"creationTime"`
}

func (m NewTransferDevice) String() string {
	return common.PointerString(m)
}

// NewTransferDeviceLifecycleStateEnum Enum with underlying type: string
type NewTransferDeviceLifecycleStateEnum string

// Set of constants representing the allowable values for NewTransferDeviceLifecycleStateEnum
const (
	NewTransferDeviceLifecycleStatePreparing NewTransferDeviceLifecycleStateEnum = "PREPARING"
)

var mappingNewTransferDeviceLifecycleState = map[string]NewTransferDeviceLifecycleStateEnum{
	"PREPARING": NewTransferDeviceLifecycleStatePreparing,
}

// GetNewTransferDeviceLifecycleStateEnumValues Enumerates the set of values for NewTransferDeviceLifecycleStateEnum
func GetNewTransferDeviceLifecycleStateEnumValues() []NewTransferDeviceLifecycleStateEnum {
	values := make([]NewTransferDeviceLifecycleStateEnum, 0)
	for _, v := range mappingNewTransferDeviceLifecycleState {
		values = append(values, v)
	}
	return values
}
