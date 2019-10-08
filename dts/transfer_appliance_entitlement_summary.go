// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TransferApplianceEntitlementSummary The representation of TransferApplianceEntitlementSummary
type TransferApplianceEntitlementSummary struct {
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	LifecycleState TransferApplianceEntitlementLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	Id *string `mandatory:"false" json:"id"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	RequestorName *string `mandatory:"false" json:"requestorName"`

	RequestorEmail *string `mandatory:"false" json:"requestorEmail"`

	// A property that can contain details on the lifecycle.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m TransferApplianceEntitlementSummary) String() string {
	return common.PointerString(m)
}
