// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oci-go-sdk/common"
)

// ConsoleHistory. An instance's serial console data. It includes configuration messages that occur when the
// instance boots, such as kernel and BIOS messages, and is useful for checking the status of
// the instance or diagnosing problems. The console data is minimally formatted ASCII text.
type ConsoleHistory struct {

	// The Availability Domain of an instance.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain,omitempty"`

	// The OCID of the compartment.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The OCID of the console history metadata object.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The OCID of the instance this console history was fetched from.
	InstanceID *string `mandatory:"true" json:"instanceId,omitempty"`

	// The current state of the console history.
	LifecycleState ConsoleHistoryLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The date and time the history was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My console history metadata`
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

func (model ConsoleHistory) String() string {
	return common.PointerString(model)
}

type ConsoleHistoryLifecycleStateEnum string

const (
	CONSOLE_HISTORY_LIFECYCLE_STATE_REQUESTED       ConsoleHistoryLifecycleStateEnum = "REQUESTED"
	CONSOLE_HISTORY_LIFECYCLE_STATE_GETTING_HISTORY ConsoleHistoryLifecycleStateEnum = "GETTING-HISTORY"
	CONSOLE_HISTORY_LIFECYCLE_STATE_SUCCEEDED       ConsoleHistoryLifecycleStateEnum = "SUCCEEDED"
	CONSOLE_HISTORY_LIFECYCLE_STATE_FAILED          ConsoleHistoryLifecycleStateEnum = "FAILED"
	CONSOLE_HISTORY_LIFECYCLE_STATE_UNKNOWN         ConsoleHistoryLifecycleStateEnum = "UNKNOWN"
)

var mapping_consolehistory_lifecycleState = map[string]ConsoleHistoryLifecycleStateEnum{
	"REQUESTED":       CONSOLE_HISTORY_LIFECYCLE_STATE_REQUESTED,
	"GETTING-HISTORY": CONSOLE_HISTORY_LIFECYCLE_STATE_GETTING_HISTORY,
	"SUCCEEDED":       CONSOLE_HISTORY_LIFECYCLE_STATE_SUCCEEDED,
	"FAILED":          CONSOLE_HISTORY_LIFECYCLE_STATE_FAILED,
	"UNKNOWN":         CONSOLE_HISTORY_LIFECYCLE_STATE_UNKNOWN,
}

func GetConsoleHistoryLifecycleStateEnumValues() []ConsoleHistoryLifecycleStateEnum {
	values := make([]ConsoleHistoryLifecycleStateEnum, 0)
	for _, v := range mapping_consolehistory_lifecycleState {
		if v != CONSOLE_HISTORY_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
