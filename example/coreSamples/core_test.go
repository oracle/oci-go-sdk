// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//
package coreSamples

import (
	"testing"

	"github.com/oracle/oci-go-sdk/common"
)

// replace following variables with your instance info
const (
	availabilityDomain        = "[The Availability Domain of the instance. Example: Uocm:PHX-AD-1]"
	compartmentID             = "[The OCID of the compartment.]"
	instanceDisplayName       = "[A user-friendly name. Does not have to be unique, and it's changeable.]"
	objectStorageURIWtihImage = "[The Object Storage URL for the image which will be used to create an image.]"
)

func TestCoreSamples_LaunchInstance(t *testing.T) {
	// NOTE: launch instance will create a new instance and VNC. please make sure delete the instance
	// after execute this sample code, otherwise, you will be charged for the running instance
	LaunchInstance(common.String(availabilityDomain), common.String(compartmentID), common.String(instanceDisplayName))
}

func TestCoreSamples_ListShapesWithPagination(t *testing.T) {
	// sample shows how pagination works in list API
	ListShapesWithPagination(common.String(compartmentID))
}

func TestCoreSamples_CreateImageDetailsPolymorphic(t *testing.T) {
	// sample about polymorphic types
	CreateImageDetailsPolymorphic(common.String(compartmentID), common.String(objectStorageURIWtihImage))
}
