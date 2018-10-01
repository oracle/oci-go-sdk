// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
//
// Email Delivery Service Tests
//
// API spec for managing OCI Email Delivery services.
//

package integtest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/email"
)

// sanity test for email service
func TestEmailClient_EmailSender(t *testing.T) {
	t.Skip("Skipping test")
	client, err := email.NewEmailClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, err)

	testEmailAddress := "gosdkintegtest@test.com"

	ctx := context.Background()

	createReq := email.CreateSenderRequest{
		CreateSenderDetails: email.CreateSenderDetails{
			CompartmentId: common.String(getCompartmentID()),
			EmailAddress:  common.String(testEmailAddress),
		},
	}

	createResp, err := client.CreateSender(ctx, createReq)
	failIfError(t, err)
	assert.Equal(t, testEmailAddress, *createResp.EmailAddress)

	getReq := email.GetSenderRequest{
		SenderId: createResp.Id,
	}

	getResp, err := client.GetSender(ctx, getReq)
	failIfError(t, err)
	assert.Equal(t, testEmailAddress, *getResp.EmailAddress)

	// you can provide additional filters and sorts, here lists all senders
	// sorted by email address and filter by email address
	listReq := email.ListSendersRequest{
		CompartmentId: common.String(getCompartmentID()),
		SortBy:        email.ListSendersSortByEmailaddress,
		SortOrder:     email.ListSendersSortOrderAsc,
	}

	listResp, err := client.ListSenders(ctx, listReq)
	failIfError(t, err)
	assert.Equal(t, 1, len(listResp.Items))

	defer func() {
		deleteReq := email.DeleteSenderRequest{
			SenderId: getReq.SenderId,
		}

		_, err = client.DeleteSender(ctx, deleteReq)
		failIfError(t, err)
	}()
}
