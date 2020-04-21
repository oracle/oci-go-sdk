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

// CreateTransferJobDetails The representation of CreateTransferJobDetails
type CreateTransferJobDetails struct {
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	UploadBucketName *string `mandatory:"false" json:"uploadBucketName"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	DeviceType CreateTransferJobDetailsDeviceTypeEnum `mandatory:"false" json:"deviceType,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
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
