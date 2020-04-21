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

// TransferPackageSummary The representation of TransferPackageSummary
type TransferPackageSummary struct {
	Label *string `mandatory:"false" json:"label"`

	LifecycleState TransferPackageSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	CreationTime *common.SDKTime `mandatory:"false" json:"creationTime"`
}

func (m TransferPackageSummary) String() string {
	return common.PointerString(m)
}

// TransferPackageSummaryLifecycleStateEnum Enum with underlying type: string
type TransferPackageSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for TransferPackageSummaryLifecycleStateEnum
const (
	TransferPackageSummaryLifecycleStatePreparing         TransferPackageSummaryLifecycleStateEnum = "PREPARING"
	TransferPackageSummaryLifecycleStateShipping          TransferPackageSummaryLifecycleStateEnum = "SHIPPING"
	TransferPackageSummaryLifecycleStateReceived          TransferPackageSummaryLifecycleStateEnum = "RECEIVED"
	TransferPackageSummaryLifecycleStateProcessing        TransferPackageSummaryLifecycleStateEnum = "PROCESSING"
	TransferPackageSummaryLifecycleStateProcessed         TransferPackageSummaryLifecycleStateEnum = "PROCESSED"
	TransferPackageSummaryLifecycleStateReturned          TransferPackageSummaryLifecycleStateEnum = "RETURNED"
	TransferPackageSummaryLifecycleStateDeleted           TransferPackageSummaryLifecycleStateEnum = "DELETED"
	TransferPackageSummaryLifecycleStateCancelled         TransferPackageSummaryLifecycleStateEnum = "CANCELLED"
	TransferPackageSummaryLifecycleStateCancelledReturned TransferPackageSummaryLifecycleStateEnum = "CANCELLED_RETURNED"
)

var mappingTransferPackageSummaryLifecycleState = map[string]TransferPackageSummaryLifecycleStateEnum{
	"PREPARING":          TransferPackageSummaryLifecycleStatePreparing,
	"SHIPPING":           TransferPackageSummaryLifecycleStateShipping,
	"RECEIVED":           TransferPackageSummaryLifecycleStateReceived,
	"PROCESSING":         TransferPackageSummaryLifecycleStateProcessing,
	"PROCESSED":          TransferPackageSummaryLifecycleStateProcessed,
	"RETURNED":           TransferPackageSummaryLifecycleStateReturned,
	"DELETED":            TransferPackageSummaryLifecycleStateDeleted,
	"CANCELLED":          TransferPackageSummaryLifecycleStateCancelled,
	"CANCELLED_RETURNED": TransferPackageSummaryLifecycleStateCancelledReturned,
}

// GetTransferPackageSummaryLifecycleStateEnumValues Enumerates the set of values for TransferPackageSummaryLifecycleStateEnum
func GetTransferPackageSummaryLifecycleStateEnumValues() []TransferPackageSummaryLifecycleStateEnum {
	values := make([]TransferPackageSummaryLifecycleStateEnum, 0)
	for _, v := range mappingTransferPackageSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
