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

type CaptureConsoleHistoryDetails struct {

	// The OCID of the instance to get the console history from.
	InstanceID *string `mandatory:"true" json:"instanceId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

func (m CaptureConsoleHistoryDetails) String() string {
	return common.PointerString(m)
}
