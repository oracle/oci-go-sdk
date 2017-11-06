// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// DrgAttachment. A link between a DRG and VCN. For more information, see
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

type DrgAttachmentLifecycleStateEnum string

const (
	DRG_ATTACHMENT_LIFECYCLE_STATE_ATTACHING DrgAttachmentLifecycleStateEnum = "ATTACHING"
	DRG_ATTACHMENT_LIFECYCLE_STATE_ATTACHED  DrgAttachmentLifecycleStateEnum = "ATTACHED"
	DRG_ATTACHMENT_LIFECYCLE_STATE_DETACHING DrgAttachmentLifecycleStateEnum = "DETACHING"
	DRG_ATTACHMENT_LIFECYCLE_STATE_DETACHED  DrgAttachmentLifecycleStateEnum = "DETACHED"
	DRG_ATTACHMENT_LIFECYCLE_STATE_UNKNOWN   DrgAttachmentLifecycleStateEnum = "UNKNOWN"
)

var mapping_drgattachment_lifecycleState = map[string]DrgAttachmentLifecycleStateEnum{
	"ATTACHING": DRG_ATTACHMENT_LIFECYCLE_STATE_ATTACHING,
	"ATTACHED":  DRG_ATTACHMENT_LIFECYCLE_STATE_ATTACHED,
	"DETACHING": DRG_ATTACHMENT_LIFECYCLE_STATE_DETACHING,
	"DETACHED":  DRG_ATTACHMENT_LIFECYCLE_STATE_DETACHED,
	"UNKNOWN":   DRG_ATTACHMENT_LIFECYCLE_STATE_UNKNOWN,
}

func GetDrgAttachmentLifecycleStateEnumValues() []DrgAttachmentLifecycleStateEnum {
	values := make([]DrgAttachmentLifecycleStateEnum, 0)
	for _, v := range mapping_drgattachment_lifecycleState {
		if v != DRG_ATTACHMENT_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
