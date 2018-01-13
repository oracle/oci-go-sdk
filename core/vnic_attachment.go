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

// VnicAttachment Represents an attachment between a VNIC and an instance. For more information, see
// [Virtual Network Interface Cards (VNICs)]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingVNICs.htm).
type VnicAttachment struct {

	// The Availability Domain of the instance.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain,omitempty"`

	// The OCID of the compartment the VNIC attachment is in, which is the same
	// compartment the instance is in.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The OCID of the VNIC attachment.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The OCID of the instance.
	InstanceID *string `mandatory:"true" json:"instanceId,omitempty"`

	// The current state of the VNIC attachment.
	LifecycleState VnicAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The OCID of the VNIC's subnet.
	SubnetID *string `mandatory:"true" json:"subnetId,omitempty"`

	// The date and time the VNIC attachment was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// A user-friendly name. Does not have to be unique.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// Which physical network interface card (NIC) the VNIC uses.
	// Certain bare metal instance shapes have two active physical NICs (0 and 1). If
	// you add a secondary VNIC to one of these instances, you can specify which NIC
	// the VNIC will use. For more information, see
	// [Virtual Network Interface Cards (VNICs)]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingVNICs.htm).
	NicIndex *int `mandatory:"false" json:"nicIndex,omitempty"`

	// The Oracle-assigned VLAN tag of the attached VNIC. Available after the
	// attachment process is complete.
	// Example: `0`
	VlanTag *int `mandatory:"false" json:"vlanTag,omitempty"`

	// The OCID of the VNIC. Available after the attachment process is complete.
	VnicID *string `mandatory:"false" json:"vnicId,omitempty"`
}

func (m VnicAttachment) String() string {
	return common.PointerString(m)
}

// VnicAttachmentLifecycleStateEnum Enum with underlying type: string
type VnicAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for VnicAttachmentLifecycleState
const (
	VnicAttachmentLifecycleStateAttaching VnicAttachmentLifecycleStateEnum = "ATTACHING"
	VnicAttachmentLifecycleStateAttached  VnicAttachmentLifecycleStateEnum = "ATTACHED"
	VnicAttachmentLifecycleStateDetaching VnicAttachmentLifecycleStateEnum = "DETACHING"
	VnicAttachmentLifecycleStateDetached  VnicAttachmentLifecycleStateEnum = "DETACHED"
	VnicAttachmentLifecycleStateUnknown   VnicAttachmentLifecycleStateEnum = "UNKNOWN"
)

var mappingVnicAttachmentLifecycleState = map[string]VnicAttachmentLifecycleStateEnum{
	"ATTACHING": VnicAttachmentLifecycleStateAttaching,
	"ATTACHED":  VnicAttachmentLifecycleStateAttached,
	"DETACHING": VnicAttachmentLifecycleStateDetaching,
	"DETACHED":  VnicAttachmentLifecycleStateDetached,
	"UNKNOWN":   VnicAttachmentLifecycleStateUnknown,
}

// GetVnicAttachmentLifecycleStateEnumValues Enumerates the set of values for VnicAttachmentLifecycleState
func GetVnicAttachmentLifecycleStateEnumValues() []VnicAttachmentLifecycleStateEnum {
	values := make([]VnicAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingVnicAttachmentLifecycleState {
		if v != VnicAttachmentLifecycleStateUnknown {
			values = append(values, v)
		}
	}
	return values
}
