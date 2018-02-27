// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
//
// Example code for OCI GOSDK
//
package example

import (
	"log"
	"os"
	"testing"

	"github.com/oracle/oci-go-sdk/example/helpers"
)

func TestMain(m *testing.M) {
	// parse the arguments defined in .env.sample file
	helpers.ParseAgrs()

	// run all tests
	retCode := m.Run()

	// each tests should already get the resources cleaned up at this point
	if res := helpers.GetResources(); len(res) != 0 {
		log.Fatalf("Resources need to be cleaned up: %v", res)
	}
	os.Exit(retCode)

}
