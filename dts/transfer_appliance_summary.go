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

// TransferApplianceSummary The representation of TransferApplianceSummary
type TransferApplianceSummary struct {
	Label *string `mandatory:"false" json:"label"`

	LifecycleState TransferApplianceSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	CreationTime *common.SDKTime `mandatory:"false" json:"creationTime"`
}

func (m TransferApplianceSummary) String() string {
	return common.PointerString(m)
}

// TransferApplianceSummaryLifecycleStateEnum Enum with underlying type: string
type TransferApplianceSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for TransferApplianceSummaryLifecycleStateEnum
const (
	TransferApplianceSummaryLifecycleStateRequested               TransferApplianceSummaryLifecycleStateEnum = "REQUESTED"
	TransferApplianceSummaryLifecycleStateOraclePreparing         TransferApplianceSummaryLifecycleStateEnum = "ORACLE_PREPARING"
	TransferApplianceSummaryLifecycleStateShipping                TransferApplianceSummaryLifecycleStateEnum = "SHIPPING"
	TransferApplianceSummaryLifecycleStateDelivered               TransferApplianceSummaryLifecycleStateEnum = "DELIVERED"
	TransferApplianceSummaryLifecycleStatePreparing               TransferApplianceSummaryLifecycleStateEnum = "PREPARING"
	TransferApplianceSummaryLifecycleStateFinalized               TransferApplianceSummaryLifecycleStateEnum = "FINALIZED"
	TransferApplianceSummaryLifecycleStateReturnDelayed           TransferApplianceSummaryLifecycleStateEnum = "RETURN_DELAYED"
	TransferApplianceSummaryLifecycleStateReturnShipped           TransferApplianceSummaryLifecycleStateEnum = "RETURN_SHIPPED"
	TransferApplianceSummaryLifecycleStateReturnShippedCancelled  TransferApplianceSummaryLifecycleStateEnum = "RETURN_SHIPPED_CANCELLED"
	TransferApplianceSummaryLifecycleStateOracleReceived          TransferApplianceSummaryLifecycleStateEnum = "ORACLE_RECEIVED"
	TransferApplianceSummaryLifecycleStateOracleReceivedCancelled TransferApplianceSummaryLifecycleStateEnum = "ORACLE_RECEIVED_CANCELLED"
	TransferApplianceSummaryLifecycleStateProcessing              TransferApplianceSummaryLifecycleStateEnum = "PROCESSING"
	TransferApplianceSummaryLifecycleStateComplete                TransferApplianceSummaryLifecycleStateEnum = "COMPLETE"
	TransferApplianceSummaryLifecycleStateCustomerNeverReceived   TransferApplianceSummaryLifecycleStateEnum = "CUSTOMER_NEVER_RECEIVED"
	TransferApplianceSummaryLifecycleStateOracleNeverReceived     TransferApplianceSummaryLifecycleStateEnum = "ORACLE_NEVER_RECEIVED"
	TransferApplianceSummaryLifecycleStateCustomerLost            TransferApplianceSummaryLifecycleStateEnum = "CUSTOMER_LOST"
	TransferApplianceSummaryLifecycleStateCancelled               TransferApplianceSummaryLifecycleStateEnum = "CANCELLED"
	TransferApplianceSummaryLifecycleStateDeleted                 TransferApplianceSummaryLifecycleStateEnum = "DELETED"
	TransferApplianceSummaryLifecycleStateRejected                TransferApplianceSummaryLifecycleStateEnum = "REJECTED"
	TransferApplianceSummaryLifecycleStateError                   TransferApplianceSummaryLifecycleStateEnum = "ERROR"
)

var mappingTransferApplianceSummaryLifecycleState = map[string]TransferApplianceSummaryLifecycleStateEnum{
	"REQUESTED":                 TransferApplianceSummaryLifecycleStateRequested,
	"ORACLE_PREPARING":          TransferApplianceSummaryLifecycleStateOraclePreparing,
	"SHIPPING":                  TransferApplianceSummaryLifecycleStateShipping,
	"DELIVERED":                 TransferApplianceSummaryLifecycleStateDelivered,
	"PREPARING":                 TransferApplianceSummaryLifecycleStatePreparing,
	"FINALIZED":                 TransferApplianceSummaryLifecycleStateFinalized,
	"RETURN_DELAYED":            TransferApplianceSummaryLifecycleStateReturnDelayed,
	"RETURN_SHIPPED":            TransferApplianceSummaryLifecycleStateReturnShipped,
	"RETURN_SHIPPED_CANCELLED":  TransferApplianceSummaryLifecycleStateReturnShippedCancelled,
	"ORACLE_RECEIVED":           TransferApplianceSummaryLifecycleStateOracleReceived,
	"ORACLE_RECEIVED_CANCELLED": TransferApplianceSummaryLifecycleStateOracleReceivedCancelled,
	"PROCESSING":                TransferApplianceSummaryLifecycleStateProcessing,
	"COMPLETE":                  TransferApplianceSummaryLifecycleStateComplete,
	"CUSTOMER_NEVER_RECEIVED":   TransferApplianceSummaryLifecycleStateCustomerNeverReceived,
	"ORACLE_NEVER_RECEIVED":     TransferApplianceSummaryLifecycleStateOracleNeverReceived,
	"CUSTOMER_LOST":             TransferApplianceSummaryLifecycleStateCustomerLost,
	"CANCELLED":                 TransferApplianceSummaryLifecycleStateCancelled,
	"DELETED":                   TransferApplianceSummaryLifecycleStateDeleted,
	"REJECTED":                  TransferApplianceSummaryLifecycleStateRejected,
	"ERROR":                     TransferApplianceSummaryLifecycleStateError,
}

// GetTransferApplianceSummaryLifecycleStateEnumValues Enumerates the set of values for TransferApplianceSummaryLifecycleStateEnum
func GetTransferApplianceSummaryLifecycleStateEnumValues() []TransferApplianceSummaryLifecycleStateEnum {
	values := make([]TransferApplianceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingTransferApplianceSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
