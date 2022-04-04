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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateTransferPackageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateTransferPackageDetailsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUpdateTransferPackageDetailsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateTransferPackageDetailsLifecycleStateEnum Enum with underlying type: string
type UpdateTransferPackageDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateTransferPackageDetailsLifecycleStateEnum
const (
	UpdateTransferPackageDetailsLifecycleStateShipping  UpdateTransferPackageDetailsLifecycleStateEnum = "SHIPPING"
	UpdateTransferPackageDetailsLifecycleStateCancelled UpdateTransferPackageDetailsLifecycleStateEnum = "CANCELLED"
)

var mappingUpdateTransferPackageDetailsLifecycleStateEnum = map[string]UpdateTransferPackageDetailsLifecycleStateEnum{
	"SHIPPING":  UpdateTransferPackageDetailsLifecycleStateShipping,
	"CANCELLED": UpdateTransferPackageDetailsLifecycleStateCancelled,
}

var mappingUpdateTransferPackageDetailsLifecycleStateEnumLowerCase = map[string]UpdateTransferPackageDetailsLifecycleStateEnum{
	"shipping":  UpdateTransferPackageDetailsLifecycleStateShipping,
	"cancelled": UpdateTransferPackageDetailsLifecycleStateCancelled,
}

// GetUpdateTransferPackageDetailsLifecycleStateEnumValues Enumerates the set of values for UpdateTransferPackageDetailsLifecycleStateEnum
func GetUpdateTransferPackageDetailsLifecycleStateEnumValues() []UpdateTransferPackageDetailsLifecycleStateEnum {
	values := make([]UpdateTransferPackageDetailsLifecycleStateEnum, 0)
	for _, v := range mappingUpdateTransferPackageDetailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTransferPackageDetailsLifecycleStateEnumStringValues Enumerates the set of values in String for UpdateTransferPackageDetailsLifecycleStateEnum
func GetUpdateTransferPackageDetailsLifecycleStateEnumStringValues() []string {
	return []string{
		"SHIPPING",
		"CANCELLED",
	}
}

// GetMappingUpdateTransferPackageDetailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTransferPackageDetailsLifecycleStateEnum(val string) (UpdateTransferPackageDetailsLifecycleStateEnum, bool) {
	enum, ok := mappingUpdateTransferPackageDetailsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
