// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for workrequests API
// This script provides a basic example of how to use work requests using Go SDK.
// This script will:
//
//   * Read in user OCI config
//   * Retrieve a list of all Work Requests for the compartment
//   * Get Work Request details
//   * List errors related to a Work Request
//   * List logs related to a Work Request
//
// This script takes no arguments
//
// Usage:
// 		go test -v example/example_work_request_test.go
//

package example

import (
	"context"
	"fmt"
	"log"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/oracle/oci-go-sdk/workrequests"
)

func ExampleWorkRequests() {
	client, err := workrequests.NewWorkRequestClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	compartmentID, err := common.DefaultConfigProvider().TenancyOCID()
	if err != nil {
		log.Println("Coulnd't read Tenancy from OCI config", err)
	}

	log.Println("Compartment ID: ", compartmentID)

	ctx := context.Background()

	workRequests := listWorkRequests(ctx, client, compartmentID)

	log.Println(len(workRequests), " Work Requests found.")

	for _, workRequest := range workRequests {
		getPrintSummary(ctx, client, workRequest.Id)

		getPrintErrors(ctx, client, workRequest.Id)

		getPrintLogs(ctx, client, workRequest.Id)
	}

	fmt.Println("Work Request example Completed")
	// Output: Work Request example Completed
}

func listWorkRequests(ctx context.Context, client workrequests.WorkRequestClient, compartmentID string) []workrequests.WorkRequestSummary {
	request := workrequests.ListWorkRequestsRequest{
		CompartmentId: &compartmentID,
		Limit:         common.Int(5),
	}

	resp, err := client.ListWorkRequests(ctx, request)
	helpers.FatalIfError(err)

	return resp.Items
}

func getPrintSummary(ctx context.Context, client workrequests.WorkRequestClient, workRequestId *string) {
	request := workrequests.GetWorkRequestRequest{
		WorkRequestId: workRequestId,
	}

	resp, err := client.GetWorkRequest(ctx, request)
	helpers.FatalIfError(err)

	printSummary(resp.WorkRequest)
}

func printSummary(w workrequests.WorkRequest) {
	log.Println("")
	log.Println("")
	log.Println("==========================================================")
	log.Printf("Work Request Details: %s\n", *w.Id)
	log.Println("==========================================================")
	log.Println("OperationType: ", *w.OperationType)
	log.Println("Status: ", w.Status)
	log.Println("ID: ", *w.Id)
	log.Println("CompartmentId: ", *w.CompartmentId)
	log.Println("PercentComplete: ", *w.PercentComplete)
	log.Println("TimeAccepted: ", *w.TimeAccepted)
	log.Println("TimeStarted: ", *w.TimeStarted)
	log.Println("TimeFinished: ", *w.TimeFinished)
	log.Println("")
}

func getPrintErrors(ctx context.Context, client workrequests.WorkRequestClient, workRequestId *string) {
	request := workrequests.ListWorkRequestErrorsRequest{
		WorkRequestId: workRequestId,
	}

	resp, err := client.ListWorkRequestErrors(ctx, request)
	helpers.FatalIfError(err)

	log.Println("==========================================================")
	log.Println("Work Request Errors")
	log.Println("==========================================================")

	for _, wrErr := range resp.Items {
		printErrors(wrErr)
	}
	log.Println("")
}

func printErrors(wrErr workrequests.WorkRequestError) {
	log.Println("{")
	log.Println(" Code: ", *wrErr.Code)
	log.Println(" Message: ", *wrErr.Message)
	log.Println(" Timestamp: ", *wrErr.Timestamp)
	log.Println("}")
}

func getPrintLogs(ctx context.Context, client workrequests.WorkRequestClient, workRequestId *string) {
	request := workrequests.ListWorkRequestLogsRequest{
		WorkRequestId: workRequestId,
		Limit:         common.Int(10),
	}

	// example showing how to use the pagination feature.
	// Other work request calls can also be paginated but aren't for simplicity.
	listLogsFunc := func(request workrequests.ListWorkRequestLogsRequest) (workrequests.ListWorkRequestLogsResponse, error) {
		return client.ListWorkRequestLogs(ctx, request)
	}

	log.Println("==========================================================")
	log.Println("Work Request Logs")
	log.Println("==========================================================")

	for resp, err := listLogsFunc(request); ; resp, err = listLogsFunc(request) {
		helpers.FatalIfError(err)

		for _, wrLog := range resp.Items {
			printLogs(wrLog)
		}

		if resp.OpcNextPage != nil {
			// if there are more items in next page, fetch items from next page
			request.Page = resp.OpcNextPage
		} else {
			// no more result, break the loop
			break
		}
	}

	log.Println("")
}

func printLogs(wrLog workrequests.WorkRequestLogEntry) {
	log.Println("{")
	log.Println(" Message: ", *wrLog.Message)
	log.Println(" Timestamp: ", *wrLog.Timestamp)
	log.Println("}")
}
