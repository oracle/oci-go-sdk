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

// TransferJob The representation of TransferJob
type TransferJob struct {
	Id *string `mandatory:"true" json:"id"`

	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	UploadBucketName *string `mandatory:"false" json:"uploadBucketName"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	Label *string `mandatory:"false" json:"label"`

	CreationTime *common.SDKTime `mandatory:"false" json:"creationTime"`

	DeviceType TransferJobDeviceTypeEnum `mandatory:"false" json:"deviceType,omitempty"`

	LifecycleState TransferJobLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Transfer Appliance labels associated with this transfer Job
	AttachedTransferApplianceLabels []string `mandatory:"false" json:"attachedTransferApplianceLabels"`

	// Transfer Package labels associated with this transfer Job
	AttachedTransferPackageLabels []string `mandatory:"false" json:"attachedTransferPackageLabels"`

	// Transfer Device labels associated with this transfer Job
	AttachedTransferDeviceLabels []string `mandatory:"false" json:"attachedTransferDeviceLabels"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m TransferJob) String() string {
	return common.PointerString(m)
}

// TransferJobDeviceTypeEnum Enum with underlying type: string
type TransferJobDeviceTypeEnum string

// Set of constants representing the allowable values for TransferJobDeviceTypeEnum
const (
	TransferJobDeviceTypeDisk      TransferJobDeviceTypeEnum = "DISK"
	TransferJobDeviceTypeAppliance TransferJobDeviceTypeEnum = "APPLIANCE"
)

var mappingTransferJobDeviceType = map[string]TransferJobDeviceTypeEnum{
	"DISK":      TransferJobDeviceTypeDisk,
	"APPLIANCE": TransferJobDeviceTypeAppliance,
}

// GetTransferJobDeviceTypeEnumValues Enumerates the set of values for TransferJobDeviceTypeEnum
func GetTransferJobDeviceTypeEnumValues() []TransferJobDeviceTypeEnum {
	values := make([]TransferJobDeviceTypeEnum, 0)
	for _, v := range mappingTransferJobDeviceType {
		values = append(values, v)
	}
	return values
}

// TransferJobLifecycleStateEnum Enum with underlying type: string
type TransferJobLifecycleStateEnum string

// Set of constants representing the allowable values for TransferJobLifecycleStateEnum
const (
	TransferJobLifecycleStateInitiated TransferJobLifecycleStateEnum = "INITIATED"
	TransferJobLifecycleStatePreparing TransferJobLifecycleStateEnum = "PREPARING"
	TransferJobLifecycleStateActive    TransferJobLifecycleStateEnum = "ACTIVE"
	TransferJobLifecycleStateDeleted   TransferJobLifecycleStateEnum = "DELETED"
	TransferJobLifecycleStateClosed    TransferJobLifecycleStateEnum = "CLOSED"
)

var mappingTransferJobLifecycleState = map[string]TransferJobLifecycleStateEnum{
	"INITIATED": TransferJobLifecycleStateInitiated,
	"PREPARING": TransferJobLifecycleStatePreparing,
	"ACTIVE":    TransferJobLifecycleStateActive,
	"DELETED":   TransferJobLifecycleStateDeleted,
	"CLOSED":    TransferJobLifecycleStateClosed,
}

// GetTransferJobLifecycleStateEnumValues Enumerates the set of values for TransferJobLifecycleStateEnum
func GetTransferJobLifecycleStateEnumValues() []TransferJobLifecycleStateEnum {
	values := make([]TransferJobLifecycleStateEnum, 0)
	for _, v := range mappingTransferJobLifecycleState {
		values = append(values, v)
	}
	return values
}
