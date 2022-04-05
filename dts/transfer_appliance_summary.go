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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TransferApplianceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTransferApplianceSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTransferApplianceSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingTransferApplianceSummaryLifecycleStateEnum = map[string]TransferApplianceSummaryLifecycleStateEnum{
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

var mappingTransferApplianceSummaryLifecycleStateEnumLowerCase = map[string]TransferApplianceSummaryLifecycleStateEnum{
	"requested":                 TransferApplianceSummaryLifecycleStateRequested,
	"oracle_preparing":          TransferApplianceSummaryLifecycleStateOraclePreparing,
	"shipping":                  TransferApplianceSummaryLifecycleStateShipping,
	"delivered":                 TransferApplianceSummaryLifecycleStateDelivered,
	"preparing":                 TransferApplianceSummaryLifecycleStatePreparing,
	"finalized":                 TransferApplianceSummaryLifecycleStateFinalized,
	"return_delayed":            TransferApplianceSummaryLifecycleStateReturnDelayed,
	"return_shipped":            TransferApplianceSummaryLifecycleStateReturnShipped,
	"return_shipped_cancelled":  TransferApplianceSummaryLifecycleStateReturnShippedCancelled,
	"oracle_received":           TransferApplianceSummaryLifecycleStateOracleReceived,
	"oracle_received_cancelled": TransferApplianceSummaryLifecycleStateOracleReceivedCancelled,
	"processing":                TransferApplianceSummaryLifecycleStateProcessing,
	"complete":                  TransferApplianceSummaryLifecycleStateComplete,
	"customer_never_received":   TransferApplianceSummaryLifecycleStateCustomerNeverReceived,
	"oracle_never_received":     TransferApplianceSummaryLifecycleStateOracleNeverReceived,
	"customer_lost":             TransferApplianceSummaryLifecycleStateCustomerLost,
	"cancelled":                 TransferApplianceSummaryLifecycleStateCancelled,
	"deleted":                   TransferApplianceSummaryLifecycleStateDeleted,
	"rejected":                  TransferApplianceSummaryLifecycleStateRejected,
	"error":                     TransferApplianceSummaryLifecycleStateError,
}

// GetTransferApplianceSummaryLifecycleStateEnumValues Enumerates the set of values for TransferApplianceSummaryLifecycleStateEnum
func GetTransferApplianceSummaryLifecycleStateEnumValues() []TransferApplianceSummaryLifecycleStateEnum {
	values := make([]TransferApplianceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingTransferApplianceSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTransferApplianceSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for TransferApplianceSummaryLifecycleStateEnum
func GetTransferApplianceSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"REQUESTED",
		"ORACLE_PREPARING",
		"SHIPPING",
		"DELIVERED",
		"PREPARING",
		"FINALIZED",
		"RETURN_DELAYED",
		"RETURN_SHIPPED",
		"RETURN_SHIPPED_CANCELLED",
		"ORACLE_RECEIVED",
		"ORACLE_RECEIVED_CANCELLED",
		"PROCESSING",
		"COMPLETE",
		"CUSTOMER_NEVER_RECEIVED",
		"ORACLE_NEVER_RECEIVED",
		"CUSTOMER_LOST",
		"CANCELLED",
		"DELETED",
		"REJECTED",
		"ERROR",
	}
}

// GetMappingTransferApplianceSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTransferApplianceSummaryLifecycleStateEnum(val string) (TransferApplianceSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingTransferApplianceSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
