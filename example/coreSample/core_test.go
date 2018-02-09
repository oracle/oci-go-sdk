// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//
package coreSample

import (
	"testing"

	"github.com/oracle/oci-go-sdk/common"
)

// replace following variables with your instance info
const (
	availabilityDomain  = "[The Availability Domain of the instance. Example: Uocm:PHX-AD-1]"
	compartmentID       = "[The OCID of the compartment.]"
	instanceDisplayName = "[A user-friendly name. Does not have to be unique, and it's changeable.]"
)

func TestCoreSamples_LaunchInstance(t *testing.T) {
	// NOTE: please make sure you will delete the instance after execute this sample code
	// it will create a new instance and VNC. You will be charged if you keep them running
	LaunchInstance(common.String(availabilityDomain), common.String(compartmentID), common.String(instanceDisplayName))
}

func TestCoreSamples_ListShapesWithPagination(t *testing.T) {
	ListShapesWithPagination(common.String(compartmentID))
}
