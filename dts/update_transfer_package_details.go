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

// UpdateTransferPackageDetails The representation of UpdateTransferPackageDetails
type UpdateTransferPackageDetails struct {
	OriginalPackageDeliveryTrackingNumber *string `mandatory:"false" json:"originalPackageDeliveryTrackingNumber"`

	ReturnPackageDeliveryTrackingNumber *string `mandatory:"false" json:"returnPackageDeliveryTrackingNumber"`

	PackageDeliveryVendor *string `mandatory:"false" json:"packageDeliveryVendor"`

	LifecycleState UpdateTransferPackageDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m UpdateTransferPackageDetails) String() string {
	return common.PointerString(m)
}

// UpdateTransferPackageDetailsLifecycleStateEnum Enum with underlying type: string
type UpdateTransferPackageDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateTransferPackageDetailsLifecycleStateEnum
const (
	UpdateTransferPackageDetailsLifecycleStateShipping  UpdateTransferPackageDetailsLifecycleStateEnum = "SHIPPING"
	UpdateTransferPackageDetailsLifecycleStateCancelled UpdateTransferPackageDetailsLifecycleStateEnum = "CANCELLED"
)

var mappingUpdateTransferPackageDetailsLifecycleState = map[string]UpdateTransferPackageDetailsLifecycleStateEnum{
	"SHIPPING":  UpdateTransferPackageDetailsLifecycleStateShipping,
	"CANCELLED": UpdateTransferPackageDetailsLifecycleStateCancelled,
}

// GetUpdateTransferPackageDetailsLifecycleStateEnumValues Enumerates the set of values for UpdateTransferPackageDetailsLifecycleStateEnum
func GetUpdateTransferPackageDetailsLifecycleStateEnumValues() []UpdateTransferPackageDetailsLifecycleStateEnum {
	values := make([]UpdateTransferPackageDetailsLifecycleStateEnum, 0)
	for _, v := range mappingUpdateTransferPackageDetailsLifecycleState {
		values = append(values, v)
	}
	return values
}
