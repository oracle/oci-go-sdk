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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateTransferJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateTransferJobDetailsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUpdateTransferJobDetailsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateTransferJobDetailsDeviceTypeEnum(string(m.DeviceType)); !ok && m.DeviceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeviceType: %s. Supported values are: %s.", m.DeviceType, strings.Join(GetUpdateTransferJobDetailsDeviceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateTransferJobDetailsLifecycleStateEnum Enum with underlying type: string
type UpdateTransferJobDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateTransferJobDetailsLifecycleStateEnum
const (
	UpdateTransferJobDetailsLifecycleStateClosed UpdateTransferJobDetailsLifecycleStateEnum = "CLOSED"
)

var mappingUpdateTransferJobDetailsLifecycleStateEnum = map[string]UpdateTransferJobDetailsLifecycleStateEnum{
	"CLOSED": UpdateTransferJobDetailsLifecycleStateClosed,
}

// GetUpdateTransferJobDetailsLifecycleStateEnumValues Enumerates the set of values for UpdateTransferJobDetailsLifecycleStateEnum
func GetUpdateTransferJobDetailsLifecycleStateEnumValues() []UpdateTransferJobDetailsLifecycleStateEnum {
	values := make([]UpdateTransferJobDetailsLifecycleStateEnum, 0)
	for _, v := range mappingUpdateTransferJobDetailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTransferJobDetailsLifecycleStateEnumStringValues Enumerates the set of values in String for UpdateTransferJobDetailsLifecycleStateEnum
func GetUpdateTransferJobDetailsLifecycleStateEnumStringValues() []string {
	return []string{
		"CLOSED",
	}
}

// GetMappingUpdateTransferJobDetailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTransferJobDetailsLifecycleStateEnum(val string) (UpdateTransferJobDetailsLifecycleStateEnum, bool) {
	mappingUpdateTransferJobDetailsLifecycleStateEnumIgnoreCase := make(map[string]UpdateTransferJobDetailsLifecycleStateEnum)
	for k, v := range mappingUpdateTransferJobDetailsLifecycleStateEnum {
		mappingUpdateTransferJobDetailsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateTransferJobDetailsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateTransferJobDetailsDeviceTypeEnum Enum with underlying type: string
type UpdateTransferJobDetailsDeviceTypeEnum string

// Set of constants representing the allowable values for UpdateTransferJobDetailsDeviceTypeEnum
const (
	UpdateTransferJobDetailsDeviceTypeDisk      UpdateTransferJobDetailsDeviceTypeEnum = "DISK"
	UpdateTransferJobDetailsDeviceTypeAppliance UpdateTransferJobDetailsDeviceTypeEnum = "APPLIANCE"
)

var mappingUpdateTransferJobDetailsDeviceTypeEnum = map[string]UpdateTransferJobDetailsDeviceTypeEnum{
	"DISK":      UpdateTransferJobDetailsDeviceTypeDisk,
	"APPLIANCE": UpdateTransferJobDetailsDeviceTypeAppliance,
}

// GetUpdateTransferJobDetailsDeviceTypeEnumValues Enumerates the set of values for UpdateTransferJobDetailsDeviceTypeEnum
func GetUpdateTransferJobDetailsDeviceTypeEnumValues() []UpdateTransferJobDetailsDeviceTypeEnum {
	values := make([]UpdateTransferJobDetailsDeviceTypeEnum, 0)
	for _, v := range mappingUpdateTransferJobDetailsDeviceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTransferJobDetailsDeviceTypeEnumStringValues Enumerates the set of values in String for UpdateTransferJobDetailsDeviceTypeEnum
func GetUpdateTransferJobDetailsDeviceTypeEnumStringValues() []string {
	return []string{
		"DISK",
		"APPLIANCE",
	}
}

// GetMappingUpdateTransferJobDetailsDeviceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTransferJobDetailsDeviceTypeEnum(val string) (UpdateTransferJobDetailsDeviceTypeEnum, bool) {
	mappingUpdateTransferJobDetailsDeviceTypeEnumIgnoreCase := make(map[string]UpdateTransferJobDetailsDeviceTypeEnum)
	for k, v := range mappingUpdateTransferJobDetailsDeviceTypeEnum {
		mappingUpdateTransferJobDetailsDeviceTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateTransferJobDetailsDeviceTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
