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

// CreateTransferJobDetails The representation of CreateTransferJobDetails
type CreateTransferJobDetails struct {
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	UploadBucketName *string `mandatory:"false" json:"uploadBucketName"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	DeviceType CreateTransferJobDetailsDeviceTypeEnum `mandatory:"false" json:"deviceType,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateTransferJobDetails) String() string {
	return common.PointerString(m)
}

// CreateTransferJobDetailsDeviceTypeEnum Enum with underlying type: string
type CreateTransferJobDetailsDeviceTypeEnum string

// Set of constants representing the allowable values for CreateTransferJobDetailsDeviceTypeEnum
const (
	CreateTransferJobDetailsDeviceTypeDisk      CreateTransferJobDetailsDeviceTypeEnum = "DISK"
	CreateTransferJobDetailsDeviceTypeAppliance CreateTransferJobDetailsDeviceTypeEnum = "APPLIANCE"
)

var mappingCreateTransferJobDetailsDeviceType = map[string]CreateTransferJobDetailsDeviceTypeEnum{
	"DISK":      CreateTransferJobDetailsDeviceTypeDisk,
	"APPLIANCE": CreateTransferJobDetailsDeviceTypeAppliance,
}

// GetCreateTransferJobDetailsDeviceTypeEnumValues Enumerates the set of values for CreateTransferJobDetailsDeviceTypeEnum
func GetCreateTransferJobDetailsDeviceTypeEnumValues() []CreateTransferJobDetailsDeviceTypeEnum {
	values := make([]CreateTransferJobDetailsDeviceTypeEnum, 0)
	for _, v := range mappingCreateTransferJobDetailsDeviceType {
		values = append(values, v)
	}
	return values
}
