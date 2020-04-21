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

// UpdateTransferApplianceDetails The representation of UpdateTransferApplianceDetails
type UpdateTransferApplianceDetails struct {
	LifecycleState UpdateTransferApplianceDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	CustomerShippingAddress *ShippingAddress `mandatory:"false" json:"customerShippingAddress"`
}

func (m UpdateTransferApplianceDetails) String() string {
	return common.PointerString(m)
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

var mappingUpdateTransferApplianceDetailsLifecycleState = map[string]UpdateTransferApplianceDetailsLifecycleStateEnum{
	"PREPARING":               UpdateTransferApplianceDetailsLifecycleStatePreparing,
	"FINALIZED":               UpdateTransferApplianceDetailsLifecycleStateFinalized,
	"DELETED":                 UpdateTransferApplianceDetailsLifecycleStateDeleted,
	"CUSTOMER_NEVER_RECEIVED": UpdateTransferApplianceDetailsLifecycleStateCustomerNeverReceived,
	"CANCELLED":               UpdateTransferApplianceDetailsLifecycleStateCancelled,
}

// GetUpdateTransferApplianceDetailsLifecycleStateEnumValues Enumerates the set of values for UpdateTransferApplianceDetailsLifecycleStateEnum
func GetUpdateTransferApplianceDetailsLifecycleStateEnumValues() []UpdateTransferApplianceDetailsLifecycleStateEnum {
	values := make([]UpdateTransferApplianceDetailsLifecycleStateEnum, 0)
	for _, v := range mappingUpdateTransferApplianceDetailsLifecycleState {
		values = append(values, v)
	}
	return values
}
