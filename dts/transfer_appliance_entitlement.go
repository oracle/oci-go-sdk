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

// TransferApplianceEntitlement The representation of TransferApplianceEntitlement
type TransferApplianceEntitlement struct {
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	LifecycleState TransferApplianceEntitlementLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	Id *string `mandatory:"false" json:"id"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	RequestorName *string `mandatory:"false" json:"requestorName"`

	RequestorEmail *string `mandatory:"false" json:"requestorEmail"`

	// A property that can contain details on the lifecycle.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	CreationTime *common.SDKTime `mandatory:"false" json:"creationTime"`

	UpdateTime *common.SDKTime `mandatory:"false" json:"updateTime"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m TransferApplianceEntitlement) String() string {
	return common.PointerString(m)
}

// TransferApplianceEntitlementLifecycleStateEnum Enum with underlying type: string
type TransferApplianceEntitlementLifecycleStateEnum string

// Set of constants representing the allowable values for TransferApplianceEntitlementLifecycleStateEnum
const (
	TransferApplianceEntitlementLifecycleStateCreating TransferApplianceEntitlementLifecycleStateEnum = "CREATING"
	TransferApplianceEntitlementLifecycleStateActive   TransferApplianceEntitlementLifecycleStateEnum = "ACTIVE"
	TransferApplianceEntitlementLifecycleStateInactive TransferApplianceEntitlementLifecycleStateEnum = "INACTIVE"
	TransferApplianceEntitlementLifecycleStateDeleted  TransferApplianceEntitlementLifecycleStateEnum = "DELETED"
)

var mappingTransferApplianceEntitlementLifecycleState = map[string]TransferApplianceEntitlementLifecycleStateEnum{
	"CREATING": TransferApplianceEntitlementLifecycleStateCreating,
	"ACTIVE":   TransferApplianceEntitlementLifecycleStateActive,
	"INACTIVE": TransferApplianceEntitlementLifecycleStateInactive,
	"DELETED":  TransferApplianceEntitlementLifecycleStateDeleted,
}

// GetTransferApplianceEntitlementLifecycleStateEnumValues Enumerates the set of values for TransferApplianceEntitlementLifecycleStateEnum
func GetTransferApplianceEntitlementLifecycleStateEnumValues() []TransferApplianceEntitlementLifecycleStateEnum {
	values := make([]TransferApplianceEntitlementLifecycleStateEnum, 0)
	for _, v := range mappingTransferApplianceEntitlementLifecycleState {
		values = append(values, v)
	}
	return values
}
