// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// NewTransferDevice The representation of NewTransferDevice
type NewTransferDevice struct {
	Label *string `mandatory:"true" json:"label"`

	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	IscsiIQN *string `mandatory:"false" json:"iscsiIQN"`

	LifecycleState NewTransferDeviceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	EncryptionPassphrase *string `mandatory:"false" json:"encryptionPassphrase"`

	TransferJobId *string `mandatory:"false" json:"transferJobId"`

	CreationTime *common.SDKTime `mandatory:"false" json:"creationTime"`
}

func (m NewTransferDevice) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NewTransferDevice) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNewTransferDeviceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNewTransferDeviceLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NewTransferDeviceLifecycleStateEnum Enum with underlying type: string
type NewTransferDeviceLifecycleStateEnum string

// Set of constants representing the allowable values for NewTransferDeviceLifecycleStateEnum
const (
	NewTransferDeviceLifecycleStatePreparing NewTransferDeviceLifecycleStateEnum = "PREPARING"
)

var mappingNewTransferDeviceLifecycleStateEnum = map[string]NewTransferDeviceLifecycleStateEnum{
	"PREPARING": NewTransferDeviceLifecycleStatePreparing,
}

var mappingNewTransferDeviceLifecycleStateEnumLowerCase = map[string]NewTransferDeviceLifecycleStateEnum{
	"preparing": NewTransferDeviceLifecycleStatePreparing,
}

// GetNewTransferDeviceLifecycleStateEnumValues Enumerates the set of values for NewTransferDeviceLifecycleStateEnum
func GetNewTransferDeviceLifecycleStateEnumValues() []NewTransferDeviceLifecycleStateEnum {
	values := make([]NewTransferDeviceLifecycleStateEnum, 0)
	for _, v := range mappingNewTransferDeviceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNewTransferDeviceLifecycleStateEnumStringValues Enumerates the set of values in String for NewTransferDeviceLifecycleStateEnum
func GetNewTransferDeviceLifecycleStateEnumStringValues() []string {
	return []string{
		"PREPARING",
	}
}

// GetMappingNewTransferDeviceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNewTransferDeviceLifecycleStateEnum(val string) (NewTransferDeviceLifecycleStateEnum, bool) {
	enum, ok := mappingNewTransferDeviceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
