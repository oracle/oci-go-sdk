// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DrgAttachment A link between a DRG and VCN. For more information, see
// [Overview of the Networking Service]({{DOC_SERVER_URL}}/Content/Network/Concepts/overview.htm).
type DrgAttachment struct {

	// The OCID of the compartment containing the DRG attachment.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The OCID of the DRG.
	DrgID *string `mandatory:"true" json:"drgId,omitempty"`

	// The DRG attachment's Oracle ID (OCID).
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The DRG attachment's current state.
	LifecycleState DrgAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The OCID of the VCN.
	VcnID *string `mandatory:"true" json:"vcnId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// The date and time the DRG attachment was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`
}

func (model DrgAttachment) String() string {
	return common.PointerString(model)
}

// DrgAttachmentLifecycleStateEnum Enum with underlying type: string
type DrgAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for DrgAttachmentLifecycleState
const (
	DrgAttachmentLifecycleStateAttaching DrgAttachmentLifecycleStateEnum = "ATTACHING"
	DrgAttachmentLifecycleStateAttached  DrgAttachmentLifecycleStateEnum = "ATTACHED"
	DrgAttachmentLifecycleStateDetaching DrgAttachmentLifecycleStateEnum = "DETACHING"
	DrgAttachmentLifecycleStateDetached  DrgAttachmentLifecycleStateEnum = "DETACHED"
	DrgAttachmentLifecycleStateUnknown   DrgAttachmentLifecycleStateEnum = "UNKNOWN"
)

var mappingDrgAttachmentLifecycleState = map[string]DrgAttachmentLifecycleStateEnum{
	"ATTACHING": DrgAttachmentLifecycleStateAttaching,
	"ATTACHED":  DrgAttachmentLifecycleStateAttached,
	"DETACHING": DrgAttachmentLifecycleStateDetaching,
	"DETACHED":  DrgAttachmentLifecycleStateDetached,
	"UNKNOWN":   DrgAttachmentLifecycleStateUnknown,
}

// GetDrgAttachmentLifecycleStateEnumValues Enumerates the set of values for DrgAttachmentLifecycleState
func GetDrgAttachmentLifecycleStateEnumValues() []DrgAttachmentLifecycleStateEnum {
	values := make([]DrgAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingDrgAttachmentLifecycleState {
		if v != DrgAttachmentLifecycleStateUnknown {
			values = append(values, v)
		}
	}
	return values
}
