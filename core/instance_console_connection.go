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

// InstanceConsoleConnection. The `InstanceConsoleConnection` API provides you with console access to virtual machine (VM) instances,
// enabling you to troubleshoot malfunctioning instances remotely.
// For more information about console access, see
// [Accessing the Console]({{DOC_SERVER_URL}}/Content/Compute/References/serialconsole.htm).
type InstanceConsoleConnection struct {

	// The OCID of the compartment to contain the console connection.
	CompartmentID *string `mandatory:"false" json:"compartmentId,omitempty"`

	// The SSH connection string for the console connection.
	ConnectionString *string `mandatory:"false" json:"connectionString,omitempty"`

	// The SSH public key fingerprint for the console connection.
	Fingerprint *string `mandatory:"false" json:"fingerprint,omitempty"`

	// The OCID of the console connection.
	ID *string `mandatory:"false" json:"id,omitempty"`

	// The OCID of the instance the console connection connects to.
	InstanceID *string `mandatory:"false" json:"instanceId,omitempty"`

	// The current state of the console connection.
	LifecycleState InstanceConsoleConnectionLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The SSH connection string for the SSH tunnel used to
	// connect to the console connection over VNC.
	VncConnectionString *string `mandatory:"false" json:"vncConnectionString,omitempty"`
}

func (model InstanceConsoleConnection) String() string {
	return common.PointerString(model)
}

type InstanceConsoleConnectionLifecycleStateEnum string

const (
	INSTANCE_CONSOLE_CONNECTION_LIFECYCLE_STATE_ACTIVE   InstanceConsoleConnectionLifecycleStateEnum = "ACTIVE"
	INSTANCE_CONSOLE_CONNECTION_LIFECYCLE_STATE_CREATING InstanceConsoleConnectionLifecycleStateEnum = "CREATING"
	INSTANCE_CONSOLE_CONNECTION_LIFECYCLE_STATE_DELETED  InstanceConsoleConnectionLifecycleStateEnum = "DELETED"
	INSTANCE_CONSOLE_CONNECTION_LIFECYCLE_STATE_DELETING InstanceConsoleConnectionLifecycleStateEnum = "DELETING"
	INSTANCE_CONSOLE_CONNECTION_LIFECYCLE_STATE_FAILED   InstanceConsoleConnectionLifecycleStateEnum = "FAILED"
	INSTANCE_CONSOLE_CONNECTION_LIFECYCLE_STATE_UNKNOWN  InstanceConsoleConnectionLifecycleStateEnum = "UNKNOWN"
)

var mapping_instanceconsoleconnection_lifecycleState = map[string]InstanceConsoleConnectionLifecycleStateEnum{
	"ACTIVE":   INSTANCE_CONSOLE_CONNECTION_LIFECYCLE_STATE_ACTIVE,
	"CREATING": INSTANCE_CONSOLE_CONNECTION_LIFECYCLE_STATE_CREATING,
	"DELETED":  INSTANCE_CONSOLE_CONNECTION_LIFECYCLE_STATE_DELETED,
	"DELETING": INSTANCE_CONSOLE_CONNECTION_LIFECYCLE_STATE_DELETING,
	"FAILED":   INSTANCE_CONSOLE_CONNECTION_LIFECYCLE_STATE_FAILED,
	"UNKNOWN":  INSTANCE_CONSOLE_CONNECTION_LIFECYCLE_STATE_UNKNOWN,
}

func GetInstanceConsoleConnectionLifecycleStateEnumValues() []InstanceConsoleConnectionLifecycleStateEnum {
	values := make([]InstanceConsoleConnectionLifecycleStateEnum, 0)
	for _, v := range mapping_instanceconsoleconnection_lifecycleState {
		if v != INSTANCE_CONSOLE_CONNECTION_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
