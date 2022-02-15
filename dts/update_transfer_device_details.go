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

// UpdateTransferDeviceDetails The representation of UpdateTransferDeviceDetails
type UpdateTransferDeviceDetails struct {
	LifecycleState UpdateTransferDeviceDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m UpdateTransferDeviceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateTransferDeviceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateTransferDeviceDetailsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUpdateTransferDeviceDetailsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateTransferDeviceDetailsLifecycleStateEnum Enum with underlying type: string
type UpdateTransferDeviceDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateTransferDeviceDetailsLifecycleStateEnum
const (
	UpdateTransferDeviceDetailsLifecycleStatePreparing UpdateTransferDeviceDetailsLifecycleStateEnum = "PREPARING"
	UpdateTransferDeviceDetailsLifecycleStateReady     UpdateTransferDeviceDetailsLifecycleStateEnum = "READY"
	UpdateTransferDeviceDetailsLifecycleStateCancelled UpdateTransferDeviceDetailsLifecycleStateEnum = "CANCELLED"
)

var mappingUpdateTransferDeviceDetailsLifecycleStateEnum = map[string]UpdateTransferDeviceDetailsLifecycleStateEnum{
	"PREPARING": UpdateTransferDeviceDetailsLifecycleStatePreparing,
	"READY":     UpdateTransferDeviceDetailsLifecycleStateReady,
	"CANCELLED": UpdateTransferDeviceDetailsLifecycleStateCancelled,
}

// GetUpdateTransferDeviceDetailsLifecycleStateEnumValues Enumerates the set of values for UpdateTransferDeviceDetailsLifecycleStateEnum
func GetUpdateTransferDeviceDetailsLifecycleStateEnumValues() []UpdateTransferDeviceDetailsLifecycleStateEnum {
	values := make([]UpdateTransferDeviceDetailsLifecycleStateEnum, 0)
	for _, v := range mappingUpdateTransferDeviceDetailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTransferDeviceDetailsLifecycleStateEnumStringValues Enumerates the set of values in String for UpdateTransferDeviceDetailsLifecycleStateEnum
func GetUpdateTransferDeviceDetailsLifecycleStateEnumStringValues() []string {
	return []string{
		"PREPARING",
		"READY",
		"CANCELLED",
	}
}

// GetMappingUpdateTransferDeviceDetailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTransferDeviceDetailsLifecycleStateEnum(val string) (UpdateTransferDeviceDetailsLifecycleStateEnum, bool) {
	mappingUpdateTransferDeviceDetailsLifecycleStateEnumIgnoreCase := make(map[string]UpdateTransferDeviceDetailsLifecycleStateEnum)
	for k, v := range mappingUpdateTransferDeviceDetailsLifecycleStateEnum {
		mappingUpdateTransferDeviceDetailsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateTransferDeviceDetailsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
