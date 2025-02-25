// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m TransferJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TransferJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTransferJobDeviceTypeEnum(string(m.DeviceType)); !ok && m.DeviceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeviceType: %s. Supported values are: %s.", m.DeviceType, strings.Join(GetTransferJobDeviceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTransferJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTransferJobLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TransferJobDeviceTypeEnum Enum with underlying type: string
type TransferJobDeviceTypeEnum string

// Set of constants representing the allowable values for TransferJobDeviceTypeEnum
const (
	TransferJobDeviceTypeDisk      TransferJobDeviceTypeEnum = "DISK"
	TransferJobDeviceTypeAppliance TransferJobDeviceTypeEnum = "APPLIANCE"
)

var mappingTransferJobDeviceTypeEnum = map[string]TransferJobDeviceTypeEnum{
	"DISK":      TransferJobDeviceTypeDisk,
	"APPLIANCE": TransferJobDeviceTypeAppliance,
}

var mappingTransferJobDeviceTypeEnumLowerCase = map[string]TransferJobDeviceTypeEnum{
	"disk":      TransferJobDeviceTypeDisk,
	"appliance": TransferJobDeviceTypeAppliance,
}

// GetTransferJobDeviceTypeEnumValues Enumerates the set of values for TransferJobDeviceTypeEnum
func GetTransferJobDeviceTypeEnumValues() []TransferJobDeviceTypeEnum {
	values := make([]TransferJobDeviceTypeEnum, 0)
	for _, v := range mappingTransferJobDeviceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTransferJobDeviceTypeEnumStringValues Enumerates the set of values in String for TransferJobDeviceTypeEnum
func GetTransferJobDeviceTypeEnumStringValues() []string {
	return []string{
		"DISK",
		"APPLIANCE",
	}
}

// GetMappingTransferJobDeviceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTransferJobDeviceTypeEnum(val string) (TransferJobDeviceTypeEnum, bool) {
	enum, ok := mappingTransferJobDeviceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingTransferJobLifecycleStateEnum = map[string]TransferJobLifecycleStateEnum{
	"INITIATED": TransferJobLifecycleStateInitiated,
	"PREPARING": TransferJobLifecycleStatePreparing,
	"ACTIVE":    TransferJobLifecycleStateActive,
	"DELETED":   TransferJobLifecycleStateDeleted,
	"CLOSED":    TransferJobLifecycleStateClosed,
}

var mappingTransferJobLifecycleStateEnumLowerCase = map[string]TransferJobLifecycleStateEnum{
	"initiated": TransferJobLifecycleStateInitiated,
	"preparing": TransferJobLifecycleStatePreparing,
	"active":    TransferJobLifecycleStateActive,
	"deleted":   TransferJobLifecycleStateDeleted,
	"closed":    TransferJobLifecycleStateClosed,
}

// GetTransferJobLifecycleStateEnumValues Enumerates the set of values for TransferJobLifecycleStateEnum
func GetTransferJobLifecycleStateEnumValues() []TransferJobLifecycleStateEnum {
	values := make([]TransferJobLifecycleStateEnum, 0)
	for _, v := range mappingTransferJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTransferJobLifecycleStateEnumStringValues Enumerates the set of values in String for TransferJobLifecycleStateEnum
func GetTransferJobLifecycleStateEnumStringValues() []string {
	return []string{
		"INITIATED",
		"PREPARING",
		"ACTIVE",
		"DELETED",
		"CLOSED",
	}
}

// GetMappingTransferJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTransferJobLifecycleStateEnum(val string) (TransferJobLifecycleStateEnum, bool) {
	enum, ok := mappingTransferJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
