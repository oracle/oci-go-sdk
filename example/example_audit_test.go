// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Audit API

package example

import (
	"context"
	"fmt"
	"time"

	"github.com/oracle/oci-go-sdk/v65/audit"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
)

func ExampleListEvents() {
	c, clerr := audit.NewAuditClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)

	// list events for last 5 hour
	req := audit.ListEventsRequest{
		CompartmentId: helpers.CompartmentID(),
		StartTime:     &common.SDKTime{time.Now().Add(time.Hour * -5)},
		EndTime:       &common.SDKTime{time.Now()},
	}

	_, err := c.ListEvents(context.Background(), req)
	helpers.FatalIfError(err)

	//log.Printf("events returned back: %v", resp.Items)
	fmt.Println("list events completed")

	// Output:
	// list events completed
}
