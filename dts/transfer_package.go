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
	"github.com/oracle/oci-go-sdk/v64/common"
	"strings"
)

// TransferPackage The representation of TransferPackage
type TransferPackage struct {
	Label *string `mandatory:"true" json:"label"`

	LifecycleState TransferPackageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	TransferJobId *string `mandatory:"false" json:"transferJobId"`

	CreationTime *common.SDKTime `mandatory:"false" json:"creationTime"`

	OriginalPackageDeliveryTrackingNumber *string `mandatory:"false" json:"originalPackageDeliveryTrackingNumber"`

	ReturnPackageDeliveryTrackingNumber *string `mandatory:"false" json:"returnPackageDeliveryTrackingNumber"`

	PackageDeliveryVendor *string `mandatory:"false" json:"packageDeliveryVendor"`

	TransferSiteShippingAddress *string `mandatory:"false" json:"transferSiteShippingAddress"`

	// Transfer Devices attached to this Transfer Package
	AttachedTransferDeviceLabels []string `mandatory:"false" json:"attachedTransferDeviceLabels"`
}

func (m TransferPackage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TransferPackage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTransferPackageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTransferPackageLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TransferPackageLifecycleStateEnum Enum with underlying type: string
type TransferPackageLifecycleStateEnum string

// Set of constants representing the allowable values for TransferPackageLifecycleStateEnum
const (
	TransferPackageLifecycleStatePreparing         TransferPackageLifecycleStateEnum = "PREPARING"
	TransferPackageLifecycleStateShipping          TransferPackageLifecycleStateEnum = "SHIPPING"
	TransferPackageLifecycleStateReceived          TransferPackageLifecycleStateEnum = "RECEIVED"
	TransferPackageLifecycleStateProcessing        TransferPackageLifecycleStateEnum = "PROCESSING"
	TransferPackageLifecycleStateProcessed         TransferPackageLifecycleStateEnum = "PROCESSED"
	TransferPackageLifecycleStateReturned          TransferPackageLifecycleStateEnum = "RETURNED"
	TransferPackageLifecycleStateDeleted           TransferPackageLifecycleStateEnum = "DELETED"
	TransferPackageLifecycleStateCancelled         TransferPackageLifecycleStateEnum = "CANCELLED"
	TransferPackageLifecycleStateCancelledReturned TransferPackageLifecycleStateEnum = "CANCELLED_RETURNED"
)

var mappingTransferPackageLifecycleStateEnum = map[string]TransferPackageLifecycleStateEnum{
	"PREPARING":          TransferPackageLifecycleStatePreparing,
	"SHIPPING":           TransferPackageLifecycleStateShipping,
	"RECEIVED":           TransferPackageLifecycleStateReceived,
	"PROCESSING":         TransferPackageLifecycleStateProcessing,
	"PROCESSED":          TransferPackageLifecycleStateProcessed,
	"RETURNED":           TransferPackageLifecycleStateReturned,
	"DELETED":            TransferPackageLifecycleStateDeleted,
	"CANCELLED":          TransferPackageLifecycleStateCancelled,
	"CANCELLED_RETURNED": TransferPackageLifecycleStateCancelledReturned,
}

var mappingTransferPackageLifecycleStateEnumLowerCase = map[string]TransferPackageLifecycleStateEnum{
	"preparing":          TransferPackageLifecycleStatePreparing,
	"shipping":           TransferPackageLifecycleStateShipping,
	"received":           TransferPackageLifecycleStateReceived,
	"processing":         TransferPackageLifecycleStateProcessing,
	"processed":          TransferPackageLifecycleStateProcessed,
	"returned":           TransferPackageLifecycleStateReturned,
	"deleted":            TransferPackageLifecycleStateDeleted,
	"cancelled":          TransferPackageLifecycleStateCancelled,
	"cancelled_returned": TransferPackageLifecycleStateCancelledReturned,
}

// GetTransferPackageLifecycleStateEnumValues Enumerates the set of values for TransferPackageLifecycleStateEnum
func GetTransferPackageLifecycleStateEnumValues() []TransferPackageLifecycleStateEnum {
	values := make([]TransferPackageLifecycleStateEnum, 0)
	for _, v := range mappingTransferPackageLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTransferPackageLifecycleStateEnumStringValues Enumerates the set of values in String for TransferPackageLifecycleStateEnum
func GetTransferPackageLifecycleStateEnumStringValues() []string {
	return []string{
		"PREPARING",
		"SHIPPING",
		"RECEIVED",
		"PROCESSING",
		"PROCESSED",
		"RETURNED",
		"DELETED",
		"CANCELLED",
		"CANCELLED_RETURNED",
	}
}

// GetMappingTransferPackageLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTransferPackageLifecycleStateEnum(val string) (TransferPackageLifecycleStateEnum, bool) {
	enum, ok := mappingTransferPackageLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
