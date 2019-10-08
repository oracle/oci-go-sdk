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

// TransferJobSummary The representation of TransferJobSummary
type TransferJobSummary struct {
	Id *string `mandatory:"false" json:"id"`

	UploadBucketName *string `mandatory:"false" json:"uploadBucketName"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	Label *string `mandatory:"false" json:"label"`

	DeviceType TransferJobSummaryDeviceTypeEnum `mandatory:"false" json:"deviceType,omitempty"`

	CreationTime *common.SDKTime `mandatory:"false" json:"creationTime"`

	LifecycleState TransferJobSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m TransferJobSummary) String() string {
	return common.PointerString(m)
}

// TransferJobSummaryDeviceTypeEnum Enum with underlying type: string
type TransferJobSummaryDeviceTypeEnum string

// Set of constants representing the allowable values for TransferJobSummaryDeviceTypeEnum
const (
	TransferJobSummaryDeviceTypeDisk      TransferJobSummaryDeviceTypeEnum = "DISK"
	TransferJobSummaryDeviceTypeAppliance TransferJobSummaryDeviceTypeEnum = "APPLIANCE"
)

var mappingTransferJobSummaryDeviceType = map[string]TransferJobSummaryDeviceTypeEnum{
	"DISK":      TransferJobSummaryDeviceTypeDisk,
	"APPLIANCE": TransferJobSummaryDeviceTypeAppliance,
}

// GetTransferJobSummaryDeviceTypeEnumValues Enumerates the set of values for TransferJobSummaryDeviceTypeEnum
func GetTransferJobSummaryDeviceTypeEnumValues() []TransferJobSummaryDeviceTypeEnum {
	values := make([]TransferJobSummaryDeviceTypeEnum, 0)
	for _, v := range mappingTransferJobSummaryDeviceType {
		values = append(values, v)
	}
	return values
}

// TransferJobSummaryLifecycleStateEnum Enum with underlying type: string
type TransferJobSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for TransferJobSummaryLifecycleStateEnum
const (
	TransferJobSummaryLifecycleStateInitiated TransferJobSummaryLifecycleStateEnum = "INITIATED"
	TransferJobSummaryLifecycleStatePreparing TransferJobSummaryLifecycleStateEnum = "PREPARING"
	TransferJobSummaryLifecycleStateActive    TransferJobSummaryLifecycleStateEnum = "ACTIVE"
	TransferJobSummaryLifecycleStateDeleted   TransferJobSummaryLifecycleStateEnum = "DELETED"
	TransferJobSummaryLifecycleStateClosed    TransferJobSummaryLifecycleStateEnum = "CLOSED"
)

var mappingTransferJobSummaryLifecycleState = map[string]TransferJobSummaryLifecycleStateEnum{
	"INITIATED": TransferJobSummaryLifecycleStateInitiated,
	"PREPARING": TransferJobSummaryLifecycleStatePreparing,
	"ACTIVE":    TransferJobSummaryLifecycleStateActive,
	"DELETED":   TransferJobSummaryLifecycleStateDeleted,
	"CLOSED":    TransferJobSummaryLifecycleStateClosed,
}

// GetTransferJobSummaryLifecycleStateEnumValues Enumerates the set of values for TransferJobSummaryLifecycleStateEnum
func GetTransferJobSummaryLifecycleStateEnumValues() []TransferJobSummaryLifecycleStateEnum {
	values := make([]TransferJobSummaryLifecycleStateEnum, 0)
	for _, v := range mappingTransferJobSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
