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

// Instance. A compute host. The image used to launch the instance determines its operating system and other
// software. The shape specified during the launch process determines the number of CPUs and memory
// allocated to the instance. For more information, see
// [Overview of the Compute Service]({{DOC_SERVER_URL}}/Content/Compute/Concepts/computeoverview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type Instance struct {

	// The Availability Domain the instance is running in.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain,omitempty"`

	// The OCID of the compartment that contains the instance.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The OCID of the instance.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The current state of the instance.
	LifecycleState InstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The region that contains the Availability Domain the instance is running in.
	// Example: `phx`
	Region *string `mandatory:"true" json:"region,omitempty"`

	// The shape of the instance. The shape determines the number of CPUs and the amount of memory
	// allocated to the instance. You can enumerate all available shapes by calling
	// ListShapes.
	Shape *string `mandatory:"true" json:"shape,omitempty"`

	// The date and time the instance was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My bare metal instance`
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// Additional metadata key/value pairs that you provide.  They serve a similar purpose and functionality from fields in the 'metadata' object.
	// They are distinguished from 'metadata' fields in that these can be nested JSON objects (whereas 'metadata' fields are string/string maps only).
	// If you don't need nested metadata values, it is strongly advised to avoid using this object and use the Metadata object instead.
	ExtendedMetadata *map[string]interface{} `mandatory:"false" json:"extendedMetadata,omitempty"`

	// The image used to boot the instance. You can enumerate all available images by calling
	// ListImages.
	ImageID *string `mandatory:"false" json:"imageId,omitempty"`

	// When a bare metal or virtual machine
	// instance boots, the iPXE firmware that runs on the instance is
	// configured to run an iPXE script to continue the boot process.
	// If you want more control over the boot process, you can provide
	// your own custom iPXE script that will run when the instance boots;
	// however, you should be aware that the same iPXE script will run
	// every time an instance boots; not only after the initial
	// LaunchInstance call.
	// The default iPXE script connects to the instance's local boot
	// volume over iSCSI and performs a network boot. If you use a custom iPXE
	// script and want to network-boot from the instance's local boot volume
	// over iSCSI the same way as the default iPXE script, you should use the
	// following iSCSI IP address: 169.254.0.2, and boot volume IQN:
	// iqn.2015-02.oracle.boot.
	// For more information about the Bring Your Own Image feature of
	// Oracle Cloud Infrastructure, see
	// [Bring Your Own Image]({{DOC_SERVER_URL}}/Content/Compute/References/bringyourownimage.htm).
	// For more information about iPXE, see http://ipxe.org.
	IpxeScript *string `mandatory:"false" json:"ipxeScript,omitempty"`

	// Custom metadata that you provide.
	Metadata *map[string]string `mandatory:"false" json:"metadata,omitempty"`
}

func (model Instance) String() string {
	return common.PointerString(model)
}

type InstanceLifecycleStateEnum string

const (
	INSTANCE_LIFECYCLE_STATE_PROVISIONING   InstanceLifecycleStateEnum = "PROVISIONING"
	INSTANCE_LIFECYCLE_STATE_RUNNING        InstanceLifecycleStateEnum = "RUNNING"
	INSTANCE_LIFECYCLE_STATE_STARTING       InstanceLifecycleStateEnum = "STARTING"
	INSTANCE_LIFECYCLE_STATE_STOPPING       InstanceLifecycleStateEnum = "STOPPING"
	INSTANCE_LIFECYCLE_STATE_STOPPED        InstanceLifecycleStateEnum = "STOPPED"
	INSTANCE_LIFECYCLE_STATE_CREATING_IMAGE InstanceLifecycleStateEnum = "CREATING_IMAGE"
	INSTANCE_LIFECYCLE_STATE_TERMINATING    InstanceLifecycleStateEnum = "TERMINATING"
	INSTANCE_LIFECYCLE_STATE_TERMINATED     InstanceLifecycleStateEnum = "TERMINATED"
	INSTANCE_LIFECYCLE_STATE_UNKNOWN        InstanceLifecycleStateEnum = "UNKNOWN"
)

var mapping_instance_lifecycleState = map[string]InstanceLifecycleStateEnum{
	"PROVISIONING":   INSTANCE_LIFECYCLE_STATE_PROVISIONING,
	"RUNNING":        INSTANCE_LIFECYCLE_STATE_RUNNING,
	"STARTING":       INSTANCE_LIFECYCLE_STATE_STARTING,
	"STOPPING":       INSTANCE_LIFECYCLE_STATE_STOPPING,
	"STOPPED":        INSTANCE_LIFECYCLE_STATE_STOPPED,
	"CREATING_IMAGE": INSTANCE_LIFECYCLE_STATE_CREATING_IMAGE,
	"TERMINATING":    INSTANCE_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":     INSTANCE_LIFECYCLE_STATE_TERMINATED,
	"UNKNOWN":        INSTANCE_LIFECYCLE_STATE_UNKNOWN,
}

func GetInstanceLifecycleStateEnumValues() []InstanceLifecycleStateEnum {
	values := make([]InstanceLifecycleStateEnum, 0)
	for _, v := range mapping_instance_lifecycleState {
		if v != INSTANCE_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
