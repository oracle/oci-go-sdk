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

func TestMain(m *testing.M) {
	// parse the arguments defined in .env.sample file
	helpers.ParseAgrs()
	os.Exit(m.Run())
}
