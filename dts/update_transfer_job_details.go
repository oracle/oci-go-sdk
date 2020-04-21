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

// UpdateTransferJobDetails The representation of UpdateTransferJobDetails
type UpdateTransferJobDetails struct {
	LifecycleState UpdateTransferJobDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	DeviceType UpdateTransferJobDetailsDeviceTypeEnum `mandatory:"false" json:"deviceType,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
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
