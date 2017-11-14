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

// Volume. A detachable block volume device that allows you to dynamically expand
// the storage capacity of an instance. For more information, see
// [Overview of Cloud Volume Storage]({{DOC_SERVER_URL}}/Content/Block/Concepts/overview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type Volume struct {

	// The Availability Domain of the volume.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain,omitempty"`

	// The OCID of the compartment that contains the volume.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName,omitempty"`

	// The volume's Oracle ID (OCID).
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The current state of a volume.
	LifecycleState VolumeLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The size of the volume in MBs. This field is deprecated. Please use sizeInGBs.
	SizeInMBs *int `mandatory:"true" json:"sizeInMBs,omitempty"`

	// The date and time the volume was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The size of the volume in GBs.
	SizeInGBs *int `mandatory:"false" json:"sizeInGBs,omitempty"`
}

func (model Volume) String() string {
	return common.PointerString(model)
}

type VolumeLifecycleStateEnum string

const (
	VOLUME_LIFECYCLE_STATE_PROVISIONING VolumeLifecycleStateEnum = "PROVISIONING"
	VOLUME_LIFECYCLE_STATE_RESTORING    VolumeLifecycleStateEnum = "RESTORING"
	VOLUME_LIFECYCLE_STATE_AVAILABLE    VolumeLifecycleStateEnum = "AVAILABLE"
	VOLUME_LIFECYCLE_STATE_TERMINATING  VolumeLifecycleStateEnum = "TERMINATING"
	VOLUME_LIFECYCLE_STATE_TERMINATED   VolumeLifecycleStateEnum = "TERMINATED"
	VOLUME_LIFECYCLE_STATE_FAULTY       VolumeLifecycleStateEnum = "FAULTY"
	VOLUME_LIFECYCLE_STATE_UNKNOWN      VolumeLifecycleStateEnum = "UNKNOWN"
)

var mapping_volume_lifecycleState = map[string]VolumeLifecycleStateEnum{
	"PROVISIONING": VOLUME_LIFECYCLE_STATE_PROVISIONING,
	"RESTORING":    VOLUME_LIFECYCLE_STATE_RESTORING,
	"AVAILABLE":    VOLUME_LIFECYCLE_STATE_AVAILABLE,
	"TERMINATING":  VOLUME_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   VOLUME_LIFECYCLE_STATE_TERMINATED,
	"FAULTY":       VOLUME_LIFECYCLE_STATE_FAULTY,
	"UNKNOWN":      VOLUME_LIFECYCLE_STATE_UNKNOWN,
}

func GetVolumeLifecycleStateEnumValues() []VolumeLifecycleStateEnum {
	values := make([]VolumeLifecycleStateEnum, 0)
	for _, v := range mapping_volume_lifecycleState {
		if v != VOLUME_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
