// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
//
// Example code for OCI GOSDK
//
package example

import (
	"os"
	"testing"

	"github.com/oracle/oci-go-sdk/example/helpers"
)

// Before run the samples, update the .env.sample file with your tenancy info.
// To run individual sample:
//   go test github.com/oracle/oci-go-sdk/example -run ^ExampleLaunchInstance$
// To run all samples:
//   go test github.com/oracle/oci-go-sdk/example
func TestMain(m *testing.M) {
	// parse the arguments defined in .env.sample file
	helpers.ParseAgrs()
	os.Exit(m.Run())
}
