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

// UpdateTransferJobDetails The representation of UpdateTransferJobDetails
type UpdateTransferJobDetails struct {
	LifecycleState UpdateTransferJobDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	DeviceType UpdateTransferJobDetailsDeviceTypeEnum `mandatory:"false" json:"deviceType,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateTransferJobDetails) String() string {
	return common.PointerString(m)
}

// UpdateTransferJobDetailsLifecycleStateEnum Enum with underlying type: string
type UpdateTransferJobDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateTransferJobDetailsLifecycleStateEnum
const (
	UpdateTransferJobDetailsLifecycleStateClosed UpdateTransferJobDetailsLifecycleStateEnum = "CLOSED"
)

var mappingUpdateTransferJobDetailsLifecycleState = map[string]UpdateTransferJobDetailsLifecycleStateEnum{
	"CLOSED": UpdateTransferJobDetailsLifecycleStateClosed,
}

// GetUpdateTransferJobDetailsLifecycleStateEnumValues Enumerates the set of values for UpdateTransferJobDetailsLifecycleStateEnum
func GetUpdateTransferJobDetailsLifecycleStateEnumValues() []UpdateTransferJobDetailsLifecycleStateEnum {
	values := make([]UpdateTransferJobDetailsLifecycleStateEnum, 0)
	for _, v := range mappingUpdateTransferJobDetailsLifecycleState {
		values = append(values, v)
	}
	return values
}

// UpdateTransferJobDetailsDeviceTypeEnum Enum with underlying type: string
type UpdateTransferJobDetailsDeviceTypeEnum string

// Set of constants representing the allowable values for UpdateTransferJobDetailsDeviceTypeEnum
const (
	UpdateTransferJobDetailsDeviceTypeDisk      UpdateTransferJobDetailsDeviceTypeEnum = "DISK"
	UpdateTransferJobDetailsDeviceTypeAppliance UpdateTransferJobDetailsDeviceTypeEnum = "APPLIANCE"
)

var mappingUpdateTransferJobDetailsDeviceType = map[string]UpdateTransferJobDetailsDeviceTypeEnum{
	"DISK":      UpdateTransferJobDetailsDeviceTypeDisk,
	"APPLIANCE": UpdateTransferJobDetailsDeviceTypeAppliance,
}

// GetUpdateTransferJobDetailsDeviceTypeEnumValues Enumerates the set of values for UpdateTransferJobDetailsDeviceTypeEnum
func GetUpdateTransferJobDetailsDeviceTypeEnumValues() []UpdateTransferJobDetailsDeviceTypeEnum {
	values := make([]UpdateTransferJobDetailsDeviceTypeEnum, 0)
	for _, v := range mappingUpdateTransferJobDetailsDeviceType {
		values = append(values, v)
	}
	return values
}
