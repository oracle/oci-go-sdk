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

// TransferDeviceSummary The representation of TransferDeviceSummary
type TransferDeviceSummary struct {
	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	IscsiIQN *string `mandatory:"false" json:"iscsiIQN"`

	Label *string `mandatory:"false" json:"label"`

	LifecycleState TransferDeviceSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	AttachedTransferPackageLabel *string `mandatory:"false" json:"attachedTransferPackageLabel"`

	CreationTime *common.SDKTime `mandatory:"false" json:"creationTime"`

	UploadStatusLogUri *string `mandatory:"false" json:"uploadStatusLogUri"`
}

func (m TransferDeviceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TransferDeviceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := mappingTransferDeviceSummaryLifecycleStateEnum[string(m.LifecycleState)]; !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTransferDeviceSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TransferDeviceSummaryLifecycleStateEnum Enum with underlying type: string
type TransferDeviceSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for TransferDeviceSummaryLifecycleStateEnum
const (
	TransferDeviceSummaryLifecycleStatePreparing  TransferDeviceSummaryLifecycleStateEnum = "PREPARING"
	TransferDeviceSummaryLifecycleStateReady      TransferDeviceSummaryLifecycleStateEnum = "READY"
	TransferDeviceSummaryLifecycleStatePackaged   TransferDeviceSummaryLifecycleStateEnum = "PACKAGED"
	TransferDeviceSummaryLifecycleStateActive     TransferDeviceSummaryLifecycleStateEnum = "ACTIVE"
	TransferDeviceSummaryLifecycleStateProcessing TransferDeviceSummaryLifecycleStateEnum = "PROCESSING"
	TransferDeviceSummaryLifecycleStateComplete   TransferDeviceSummaryLifecycleStateEnum = "COMPLETE"
	TransferDeviceSummaryLifecycleStateMissing    TransferDeviceSummaryLifecycleStateEnum = "MISSING"
	TransferDeviceSummaryLifecycleStateError      TransferDeviceSummaryLifecycleStateEnum = "ERROR"
	TransferDeviceSummaryLifecycleStateDeleted    TransferDeviceSummaryLifecycleStateEnum = "DELETED"
	TransferDeviceSummaryLifecycleStateCancelled  TransferDeviceSummaryLifecycleStateEnum = "CANCELLED"
)

var mappingTransferDeviceSummaryLifecycleStateEnum = map[string]TransferDeviceSummaryLifecycleStateEnum{
	"PREPARING":  TransferDeviceSummaryLifecycleStatePreparing,
	"READY":      TransferDeviceSummaryLifecycleStateReady,
	"PACKAGED":   TransferDeviceSummaryLifecycleStatePackaged,
	"ACTIVE":     TransferDeviceSummaryLifecycleStateActive,
	"PROCESSING": TransferDeviceSummaryLifecycleStateProcessing,
	"COMPLETE":   TransferDeviceSummaryLifecycleStateComplete,
	"MISSING":    TransferDeviceSummaryLifecycleStateMissing,
	"ERROR":      TransferDeviceSummaryLifecycleStateError,
	"DELETED":    TransferDeviceSummaryLifecycleStateDeleted,
	"CANCELLED":  TransferDeviceSummaryLifecycleStateCancelled,
}

// GetTransferDeviceSummaryLifecycleStateEnumValues Enumerates the set of values for TransferDeviceSummaryLifecycleStateEnum
func GetTransferDeviceSummaryLifecycleStateEnumValues() []TransferDeviceSummaryLifecycleStateEnum {
	values := make([]TransferDeviceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingTransferDeviceSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTransferDeviceSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for TransferDeviceSummaryLifecycleStateEnum
func GetTransferDeviceSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"PREPARING",
		"READY",
		"PACKAGED",
		"ACTIVE",
		"PROCESSING",
		"COMPLETE",
		"MISSING",
		"ERROR",
		"DELETED",
		"CANCELLED",
	}
}
