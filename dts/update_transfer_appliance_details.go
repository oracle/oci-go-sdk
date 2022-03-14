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
	"github.com/oracle/oci-go-sdk/v62/common"
	"strings"
)

// UpdateTransferApplianceDetails The representation of UpdateTransferApplianceDetails
type UpdateTransferApplianceDetails struct {
	LifecycleState UpdateTransferApplianceDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	CustomerShippingAddress *ShippingAddress `mandatory:"false" json:"customerShippingAddress"`
}

func (m UpdateTransferApplianceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateTransferApplianceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateTransferApplianceDetailsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUpdateTransferApplianceDetailsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateTransferApplianceDetailsLifecycleStateEnum Enum with underlying type: string
type UpdateTransferApplianceDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateTransferApplianceDetailsLifecycleStateEnum
const (
	UpdateTransferApplianceDetailsLifecycleStatePreparing             UpdateTransferApplianceDetailsLifecycleStateEnum = "PREPARING"
	UpdateTransferApplianceDetailsLifecycleStateFinalized             UpdateTransferApplianceDetailsLifecycleStateEnum = "FINALIZED"
	UpdateTransferApplianceDetailsLifecycleStateDeleted               UpdateTransferApplianceDetailsLifecycleStateEnum = "DELETED"
	UpdateTransferApplianceDetailsLifecycleStateCustomerNeverReceived UpdateTransferApplianceDetailsLifecycleStateEnum = "CUSTOMER_NEVER_RECEIVED"
	UpdateTransferApplianceDetailsLifecycleStateCancelled             UpdateTransferApplianceDetailsLifecycleStateEnum = "CANCELLED"
)

var mappingUpdateTransferApplianceDetailsLifecycleStateEnum = map[string]UpdateTransferApplianceDetailsLifecycleStateEnum{
	"PREPARING":               UpdateTransferApplianceDetailsLifecycleStatePreparing,
	"FINALIZED":               UpdateTransferApplianceDetailsLifecycleStateFinalized,
	"DELETED":                 UpdateTransferApplianceDetailsLifecycleStateDeleted,
	"CUSTOMER_NEVER_RECEIVED": UpdateTransferApplianceDetailsLifecycleStateCustomerNeverReceived,
	"CANCELLED":               UpdateTransferApplianceDetailsLifecycleStateCancelled,
}

var mappingUpdateTransferApplianceDetailsLifecycleStateEnumLowerCase = map[string]UpdateTransferApplianceDetailsLifecycleStateEnum{
	"preparing":               UpdateTransferApplianceDetailsLifecycleStatePreparing,
	"finalized":               UpdateTransferApplianceDetailsLifecycleStateFinalized,
	"deleted":                 UpdateTransferApplianceDetailsLifecycleStateDeleted,
	"customer_never_received": UpdateTransferApplianceDetailsLifecycleStateCustomerNeverReceived,
	"cancelled":               UpdateTransferApplianceDetailsLifecycleStateCancelled,
}

// GetUpdateTransferApplianceDetailsLifecycleStateEnumValues Enumerates the set of values for UpdateTransferApplianceDetailsLifecycleStateEnum
func GetUpdateTransferApplianceDetailsLifecycleStateEnumValues() []UpdateTransferApplianceDetailsLifecycleStateEnum {
	values := make([]UpdateTransferApplianceDetailsLifecycleStateEnum, 0)
	for _, v := range mappingUpdateTransferApplianceDetailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTransferApplianceDetailsLifecycleStateEnumStringValues Enumerates the set of values in String for UpdateTransferApplianceDetailsLifecycleStateEnum
func GetUpdateTransferApplianceDetailsLifecycleStateEnumStringValues() []string {
	return []string{
		"PREPARING",
		"FINALIZED",
		"DELETED",
		"CUSTOMER_NEVER_RECEIVED",
		"CANCELLED",
	}
}

// GetMappingUpdateTransferApplianceDetailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTransferApplianceDetailsLifecycleStateEnum(val string) (UpdateTransferApplianceDetailsLifecycleStateEnum, bool) {
	enum, ok := mappingUpdateTransferApplianceDetailsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
