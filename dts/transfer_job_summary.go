// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"strings"
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

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m TransferJobSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TransferJobSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTransferJobSummaryDeviceTypeEnum(string(m.DeviceType)); !ok && m.DeviceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeviceType: %s. Supported values are: %s.", m.DeviceType, strings.Join(GetTransferJobSummaryDeviceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTransferJobSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTransferJobSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TransferJobSummaryDeviceTypeEnum Enum with underlying type: string
type TransferJobSummaryDeviceTypeEnum string

// Set of constants representing the allowable values for TransferJobSummaryDeviceTypeEnum
const (
	TransferJobSummaryDeviceTypeDisk      TransferJobSummaryDeviceTypeEnum = "DISK"
	TransferJobSummaryDeviceTypeAppliance TransferJobSummaryDeviceTypeEnum = "APPLIANCE"
)

var mappingTransferJobSummaryDeviceTypeEnum = map[string]TransferJobSummaryDeviceTypeEnum{
	"DISK":      TransferJobSummaryDeviceTypeDisk,
	"APPLIANCE": TransferJobSummaryDeviceTypeAppliance,
}

var mappingTransferJobSummaryDeviceTypeEnumLowerCase = map[string]TransferJobSummaryDeviceTypeEnum{
	"disk":      TransferJobSummaryDeviceTypeDisk,
	"appliance": TransferJobSummaryDeviceTypeAppliance,
}

// GetTransferJobSummaryDeviceTypeEnumValues Enumerates the set of values for TransferJobSummaryDeviceTypeEnum
func GetTransferJobSummaryDeviceTypeEnumValues() []TransferJobSummaryDeviceTypeEnum {
	values := make([]TransferJobSummaryDeviceTypeEnum, 0)
	for _, v := range mappingTransferJobSummaryDeviceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTransferJobSummaryDeviceTypeEnumStringValues Enumerates the set of values in String for TransferJobSummaryDeviceTypeEnum
func GetTransferJobSummaryDeviceTypeEnumStringValues() []string {
	return []string{
		"DISK",
		"APPLIANCE",
	}
}

// GetMappingTransferJobSummaryDeviceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTransferJobSummaryDeviceTypeEnum(val string) (TransferJobSummaryDeviceTypeEnum, bool) {
	enum, ok := mappingTransferJobSummaryDeviceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingTransferJobSummaryLifecycleStateEnum = map[string]TransferJobSummaryLifecycleStateEnum{
	"INITIATED": TransferJobSummaryLifecycleStateInitiated,
	"PREPARING": TransferJobSummaryLifecycleStatePreparing,
	"ACTIVE":    TransferJobSummaryLifecycleStateActive,
	"DELETED":   TransferJobSummaryLifecycleStateDeleted,
	"CLOSED":    TransferJobSummaryLifecycleStateClosed,
}

var mappingTransferJobSummaryLifecycleStateEnumLowerCase = map[string]TransferJobSummaryLifecycleStateEnum{
	"initiated": TransferJobSummaryLifecycleStateInitiated,
	"preparing": TransferJobSummaryLifecycleStatePreparing,
	"active":    TransferJobSummaryLifecycleStateActive,
	"deleted":   TransferJobSummaryLifecycleStateDeleted,
	"closed":    TransferJobSummaryLifecycleStateClosed,
}

// GetTransferJobSummaryLifecycleStateEnumValues Enumerates the set of values for TransferJobSummaryLifecycleStateEnum
func GetTransferJobSummaryLifecycleStateEnumValues() []TransferJobSummaryLifecycleStateEnum {
	values := make([]TransferJobSummaryLifecycleStateEnum, 0)
	for _, v := range mappingTransferJobSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTransferJobSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for TransferJobSummaryLifecycleStateEnum
func GetTransferJobSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"INITIATED",
		"PREPARING",
		"ACTIVE",
		"DELETED",
		"CLOSED",
	}
}

// GetMappingTransferJobSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTransferJobSummaryLifecycleStateEnum(val string) (TransferJobSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingTransferJobSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
