// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Oracle Cloud Infrastructure Go SDK

package example

import (
	"os"
	"testing"

	"github.com/oracle/oci-go-sdk/v65/example/helpers"
)

// Before run the samples, update the environment variables with your tenancy info.
//
// To run individual sample:
//
//	go test github.com/oracle/oci-go-sdk/example -run ^ExampleLaunchInstance$
//
// To run all samples:
//
//	go test github.com/oracle/oci-go-sdk/example
func TestMain(m *testing.M) {
	// ParseEnvironmentVariables assumes that you have configured your environment variables with following configs
	// Required configs are:
	//  OCI_AVAILABILITY_DOMAIN -- The Availability Domain of the instance. Example: Uocm:PHX-AD-1
	//  OCI_COMPARTMENT_ID      -- The OCID of the compartment
	//  OCI_ROOT_COMPARTMENT_ID -- The OCID of the root compartment
	helpers.ParseEnvironmentVariables()
	os.Exit(m.Run())
}
