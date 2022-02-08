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
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// TransferAppliance The representation of TransferAppliance
type TransferAppliance struct {

	// Unique alpha-numeric identifier for a transfer appliance auto generated during create.
	Label *string `mandatory:"true" json:"label"`

	LifecycleState TransferApplianceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	TransferJobId *string `mandatory:"false" json:"transferJobId"`

	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	CreationTime *common.SDKTime `mandatory:"false" json:"creationTime"`

	CustomerReceivedTime *common.SDKTime `mandatory:"false" json:"customerReceivedTime"`

	CustomerReturnedTime *common.SDKTime `mandatory:"false" json:"customerReturnedTime"`

	NextBillingTime *common.SDKTime `mandatory:"false" json:"nextBillingTime"`

	DeliverySecurityTieId *string `mandatory:"false" json:"deliverySecurityTieId"`

	ReturnSecurityTieId *string `mandatory:"false" json:"returnSecurityTieId"`

	ApplianceDeliveryTrackingNumber *string `mandatory:"false" json:"applianceDeliveryTrackingNumber"`

	ApplianceReturnDeliveryTrackingNumber *string `mandatory:"false" json:"applianceReturnDeliveryTrackingNumber"`

	ApplianceDeliveryVendor *string `mandatory:"false" json:"applianceDeliveryVendor"`

	CustomerShippingAddress *ShippingAddress `mandatory:"false" json:"customerShippingAddress"`

	UploadStatusLogUri *string `mandatory:"false" json:"uploadStatusLogUri"`
}

func (m TransferAppliance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TransferAppliance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := mappingTransferApplianceLifecycleStateEnum[string(m.LifecycleState)]; !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTransferApplianceLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TransferApplianceLifecycleStateEnum Enum with underlying type: string
type TransferApplianceLifecycleStateEnum string

// Set of constants representing the allowable values for TransferApplianceLifecycleStateEnum
const (
	TransferApplianceLifecycleStateRequested               TransferApplianceLifecycleStateEnum = "REQUESTED"
	TransferApplianceLifecycleStateOraclePreparing         TransferApplianceLifecycleStateEnum = "ORACLE_PREPARING"
	TransferApplianceLifecycleStateShipping                TransferApplianceLifecycleStateEnum = "SHIPPING"
	TransferApplianceLifecycleStateDelivered               TransferApplianceLifecycleStateEnum = "DELIVERED"
	TransferApplianceLifecycleStatePreparing               TransferApplianceLifecycleStateEnum = "PREPARING"
	TransferApplianceLifecycleStateFinalized               TransferApplianceLifecycleStateEnum = "FINALIZED"
	TransferApplianceLifecycleStateReturnDelayed           TransferApplianceLifecycleStateEnum = "RETURN_DELAYED"
	TransferApplianceLifecycleStateReturnShipped           TransferApplianceLifecycleStateEnum = "RETURN_SHIPPED"
	TransferApplianceLifecycleStateReturnShippedCancelled  TransferApplianceLifecycleStateEnum = "RETURN_SHIPPED_CANCELLED"
	TransferApplianceLifecycleStateOracleReceived          TransferApplianceLifecycleStateEnum = "ORACLE_RECEIVED"
	TransferApplianceLifecycleStateOracleReceivedCancelled TransferApplianceLifecycleStateEnum = "ORACLE_RECEIVED_CANCELLED"
	TransferApplianceLifecycleStateProcessing              TransferApplianceLifecycleStateEnum = "PROCESSING"
	TransferApplianceLifecycleStateComplete                TransferApplianceLifecycleStateEnum = "COMPLETE"
	TransferApplianceLifecycleStateCustomerNeverReceived   TransferApplianceLifecycleStateEnum = "CUSTOMER_NEVER_RECEIVED"
	TransferApplianceLifecycleStateOracleNeverReceived     TransferApplianceLifecycleStateEnum = "ORACLE_NEVER_RECEIVED"
	TransferApplianceLifecycleStateCustomerLost            TransferApplianceLifecycleStateEnum = "CUSTOMER_LOST"
	TransferApplianceLifecycleStateCancelled               TransferApplianceLifecycleStateEnum = "CANCELLED"
	TransferApplianceLifecycleStateDeleted                 TransferApplianceLifecycleStateEnum = "DELETED"
	TransferApplianceLifecycleStateRejected                TransferApplianceLifecycleStateEnum = "REJECTED"
	TransferApplianceLifecycleStateError                   TransferApplianceLifecycleStateEnum = "ERROR"
)

var mappingTransferApplianceLifecycleStateEnum = map[string]TransferApplianceLifecycleStateEnum{
	"REQUESTED":                 TransferApplianceLifecycleStateRequested,
	"ORACLE_PREPARING":          TransferApplianceLifecycleStateOraclePreparing,
	"SHIPPING":                  TransferApplianceLifecycleStateShipping,
	"DELIVERED":                 TransferApplianceLifecycleStateDelivered,
	"PREPARING":                 TransferApplianceLifecycleStatePreparing,
	"FINALIZED":                 TransferApplianceLifecycleStateFinalized,
	"RETURN_DELAYED":            TransferApplianceLifecycleStateReturnDelayed,
	"RETURN_SHIPPED":            TransferApplianceLifecycleStateReturnShipped,
	"RETURN_SHIPPED_CANCELLED":  TransferApplianceLifecycleStateReturnShippedCancelled,
	"ORACLE_RECEIVED":           TransferApplianceLifecycleStateOracleReceived,
	"ORACLE_RECEIVED_CANCELLED": TransferApplianceLifecycleStateOracleReceivedCancelled,
	"PROCESSING":                TransferApplianceLifecycleStateProcessing,
	"COMPLETE":                  TransferApplianceLifecycleStateComplete,
	"CUSTOMER_NEVER_RECEIVED":   TransferApplianceLifecycleStateCustomerNeverReceived,
	"ORACLE_NEVER_RECEIVED":     TransferApplianceLifecycleStateOracleNeverReceived,
	"CUSTOMER_LOST":             TransferApplianceLifecycleStateCustomerLost,
	"CANCELLED":                 TransferApplianceLifecycleStateCancelled,
	"DELETED":                   TransferApplianceLifecycleStateDeleted,
	"REJECTED":                  TransferApplianceLifecycleStateRejected,
	"ERROR":                     TransferApplianceLifecycleStateError,
}

// GetTransferApplianceLifecycleStateEnumValues Enumerates the set of values for TransferApplianceLifecycleStateEnum
func GetTransferApplianceLifecycleStateEnumValues() []TransferApplianceLifecycleStateEnum {
	values := make([]TransferApplianceLifecycleStateEnum, 0)
	for _, v := range mappingTransferApplianceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTransferApplianceLifecycleStateEnumStringValues Enumerates the set of values in String for TransferApplianceLifecycleStateEnum
func GetTransferApplianceLifecycleStateEnumStringValues() []string {
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
