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

// TransferDevice The representation of TransferDevice
type TransferDevice struct {
	Label *string `mandatory:"true" json:"label"`

	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	IscsiIQN *string `mandatory:"false" json:"iscsiIQN"`

	LifecycleState TransferDeviceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	TransferJobId *string `mandatory:"false" json:"transferJobId"`

	AttachedTransferPackageLabel *string `mandatory:"false" json:"attachedTransferPackageLabel"`

	CreationTime *common.SDKTime `mandatory:"false" json:"creationTime"`

	UploadStatusLogUri *string `mandatory:"false" json:"uploadStatusLogUri"`
}

func (m TransferDevice) String() string {
	return common.PointerString(m)
}

// TransferDeviceLifecycleStateEnum Enum with underlying type: string
type TransferDeviceLifecycleStateEnum string

// Set of constants representing the allowable values for TransferDeviceLifecycleStateEnum
const (
	TransferDeviceLifecycleStatePreparing  TransferDeviceLifecycleStateEnum = "PREPARING"
	TransferDeviceLifecycleStateReady      TransferDeviceLifecycleStateEnum = "READY"
	TransferDeviceLifecycleStatePackaged   TransferDeviceLifecycleStateEnum = "PACKAGED"
	TransferDeviceLifecycleStateActive     TransferDeviceLifecycleStateEnum = "ACTIVE"
	TransferDeviceLifecycleStateProcessing TransferDeviceLifecycleStateEnum = "PROCESSING"
	TransferDeviceLifecycleStateComplete   TransferDeviceLifecycleStateEnum = "COMPLETE"
	TransferDeviceLifecycleStateMissing    TransferDeviceLifecycleStateEnum = "MISSING"
	TransferDeviceLifecycleStateError      TransferDeviceLifecycleStateEnum = "ERROR"
	TransferDeviceLifecycleStateDeleted    TransferDeviceLifecycleStateEnum = "DELETED"
	TransferDeviceLifecycleStateCancelled  TransferDeviceLifecycleStateEnum = "CANCELLED"
)

var mappingTransferDeviceLifecycleState = map[string]TransferDeviceLifecycleStateEnum{
	"PREPARING":  TransferDeviceLifecycleStatePreparing,
	"READY":      TransferDeviceLifecycleStateReady,
	"PACKAGED":   TransferDeviceLifecycleStatePackaged,
	"ACTIVE":     TransferDeviceLifecycleStateActive,
	"PROCESSING": TransferDeviceLifecycleStateProcessing,
	"COMPLETE":   TransferDeviceLifecycleStateComplete,
	"MISSING":    TransferDeviceLifecycleStateMissing,
	"ERROR":      TransferDeviceLifecycleStateError,
	"DELETED":    TransferDeviceLifecycleStateDeleted,
	"CANCELLED":  TransferDeviceLifecycleStateCancelled,
}

// GetTransferDeviceLifecycleStateEnumValues Enumerates the set of values for TransferDeviceLifecycleStateEnum
func GetTransferDeviceLifecycleStateEnumValues() []TransferDeviceLifecycleStateEnum {
	values := make([]TransferDeviceLifecycleStateEnum, 0)
	for _, v := range mappingTransferDeviceLifecycleState {
		values = append(values, v)
	}
	return values
}
