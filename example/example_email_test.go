// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Email Delivery Service API

package example

import (
	"context"
	"fmt"
	"log"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/email"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
)

const (
	// The address of the email sender
	senderEmailAddress = "sample@sample.com"
)

func ExampleEmailSender() {
	client, err := email.NewEmailClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	ctx := context.Background()

	createReq := email.CreateSenderRequest{
		CreateSenderDetails: email.CreateSenderDetails{
			CompartmentId: helpers.CompartmentID(),
			EmailAddress:  common.String(senderEmailAddress),
		},
	}

	createResp, err := client.CreateSender(ctx, createReq)
	helpers.FatalIfError(err)
	fmt.Println("email sender created")

	getReq := email.GetSenderRequest{
		SenderId: createResp.Id,
	}

	getResp, err := client.GetSender(ctx, getReq)
	helpers.FatalIfError(err)
	fmt.Println("get email sender")
	log.Printf("get email sender with email address %s\n", *getResp.EmailAddress)

	// you can provide additional filters and sorts, here lists all senders
	// sorted by email address and filter by email address
	listReq := email.ListSendersRequest{
		CompartmentId: helpers.CompartmentID(),
		SortBy:        email.ListSendersSortByEmailaddress,
		SortOrder:     email.ListSendersSortOrderAsc,
	}

	listResp, err := client.ListSenders(ctx, listReq)
	helpers.FatalIfError(err)
	log.Printf("list email senders return %v results\n", len(listResp.Items))
	fmt.Println("list email senders")

	defer func() {
		deleteReq := email.DeleteSenderRequest{
			SenderId: getReq.SenderId,
		}

		_, err = client.DeleteSender(ctx, deleteReq)
		helpers.FatalIfError(err)
		fmt.Println("email sender deleted")
	}()

	// Output:
	// email sender created
	// get email sender
	// list email senders
	// email sender deleted
}
