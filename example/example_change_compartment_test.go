// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for changing compartment.
// This script provides a basic example of how to move an instance from one compartment to another using Go SDK.
// This script will:
//
//    * Read user configuration
//    * Construct ComputeClient using user configuration
//    * Construct ChangeInstanceCompartmentDetails()
//    * Call ChangeInstanceCompartment() in core.ComputeClient()
//    * List Instance and its attached resources before and after move operation
//
//  This script takes the following values from environment variables
//
//    * INSTANCE_ID    - The instance id of an instance
//    * COMPARTMENT_ID - The target compartment id
//    * IF_MATCH (Optional)
//          The Instance will be moved only if the etag you provide matches the resource's current etag value.
//    * OPC_RETRY_TOKEN (Optional)
//          A token that uniquely identifies a request so it can be retried in case of a timeout or server error
//          without risk of executing that same action again. Retry tokens expire after 24 hours, but can be
//          invalidated before then due to conflicting operations
//
//

package example

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/oracle/oci-go-sdk/workrequests"
	"log"
	"math"
	"os"
	"time"
)

var (
	instanceId, targetCompartmentId, ifMatch, opcRetryToken string
	retryPolicy                                             common.RetryPolicy
)

func ExampleChangeCompartment() {

	// Parse environment variables to get instanceId, targetCompartmentId, ifMatch and opcRetryToken
	parseEnvironmentVariables()

	// Create ComputeClient with default configuration
	computeClient, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)
	ctx := context.Background()

	// Get Compartment Id of the Instance
	request := core.GetInstanceRequest{
		InstanceId: common.String(instanceId),
	}
	r, err := computeClient.GetInstance(ctx, request)
	helpers.FatalIfError(err)
	availabilityDomain := *r.AvailabilityDomain
	sourceCompartmentID := *r.Instance.CompartmentId
	// Do not attempt compartment move, if the source and target compartment ids are same
	if sourceCompartmentID == targetCompartmentId {
		log.Printf("Source and target compartment ids are same !")
		os.Exit(1)
	}

	log.Printf(" ")
	log.Printf("Instance info before compartment move : ")
	if r.Etag != nil {
		log.Printf("   ETag : %s", *r.Etag)
	}
	printInstanceInfo(sourceCompartmentID, computeClient, availabilityDomain)

	// Create ChangeInstanceCompartmentDetails
	changeInstanceCompartmentDetails := core.ChangeInstanceCompartmentDetails{
		CompartmentId: common.String(targetCompartmentId),
	}

	// Create ChangeInstanceCompartmentRequest
	changeInstanceCompartmentRequest := core.ChangeInstanceCompartmentRequest{
		InstanceId:                       common.String(instanceId),
		ChangeInstanceCompartmentDetails: changeInstanceCompartmentDetails,
	}

	if len(ifMatch) > 0 {
		changeInstanceCompartmentRequest.IfMatch = common.String(ifMatch)
	}

	if len(opcRetryToken) > 0 {
		changeInstanceCompartmentRequest.OpcRetryToken = common.String(opcRetryToken)
	}

	log.Printf(" ")
	log.Printf("Moving Instance to target compartment ...")
	// Perform compartment move operation
	rs, err := computeClient.ChangeInstanceCompartment(ctx, changeInstanceCompartmentRequest)
	helpers.FatalIfError(err)

	//Wait for compartment move operation
	waitUnitlMoveCompletion(rs.OpcWorkRequestId)

	log.Printf(" ")
	log.Printf("Instance info after compartment move : ")
	printInstanceInfo(targetCompartmentId, computeClient, availabilityDomain)

	log.Printf(" ")
	fmt.Println("Change Compartment Completed")
	// Output:
	// Change Compartment Completed

}

func printInstanceInfo(id string, c core.ComputeClient, availabilityDomain string) {
	log.Printf("   Compartment Id : %s", id)
	printVolumeAttachments(c, id)
	printVnicAttachments(c, id)
	printBootVolumecAttachments(c, availabilityDomain, id)
	printConsoleConnections(c, id)
	printConsoleHistories(c, id)

}

func printVolumeAttachments(c core.ComputeClient, compartmentId string) {

	request := core.ListVolumeAttachmentsRequest{
		CompartmentId: common.String(compartmentId),
		InstanceId:    common.String(instanceId),
	}
	lvar, err := c.ListVolumeAttachments(context.Background(), request)
	helpers.FatalIfError(err)
	var volumeAttachments = lvar.Items
	if len(volumeAttachments) > 0 {
		log.Printf("   Volume Attachments:")
		for _, v := range volumeAttachments {
			log.Printf("     Volume id : %s", *v.GetId())
			log.Printf("     Compartment id : %s", *v.GetCompartmentId())
			log.Printf("     State : %s", v.GetLifecycleState())
			log.Printf(" ")
		}
	}

}

func printVnicAttachments(c core.ComputeClient, compartmentId string) {

	request := core.ListVnicAttachmentsRequest{
		CompartmentId: common.String(compartmentId),
		InstanceId:    common.String(instanceId),
	}
	lvar, err := c.ListVnicAttachments(context.Background(), request)
	helpers.FatalIfError(err)
	var vnicAttachments = lvar.Items
	if len(vnicAttachments) > 0 {
		log.Printf("   Vnic Attachments:")
		for _, v := range vnicAttachments {
			log.Printf("     Vnic id : %s", *v.VnicId)
			log.Printf("     Compartment id : %s", *v.CompartmentId)
			log.Printf("     State : %s", v.LifecycleState)
			log.Printf(" ")
		}
	}

}

func printBootVolumecAttachments(c core.ComputeClient, availabilityDomain string, compartmentId string) {

	request := core.ListBootVolumeAttachmentsRequest{
		CompartmentId:      common.String(compartmentId),
		InstanceId:         common.String(instanceId),
		AvailabilityDomain: common.String(availabilityDomain),
	}
	lvar, err := c.ListBootVolumeAttachments(context.Background(), request)
	helpers.FatalIfError(err)
	var bootVolumeAttachments = lvar.Items
	if len(bootVolumeAttachments) > 0 {
		log.Printf("   Boot Volume Attachments:")
		for _, v := range bootVolumeAttachments {
			log.Printf("     Volume id : %s", *v.BootVolumeId)
			log.Printf("     Compartment id : %s", *v.CompartmentId)
			log.Printf("     State : %s", v.LifecycleState)
			log.Printf(" ")
		}
	}

}

func printConsoleConnections(c core.ComputeClient, compartmentId string) {

	request := core.ListInstanceConsoleConnectionsRequest{
		CompartmentId: common.String(compartmentId),
		InstanceId:    common.String(instanceId),
	}
	lvar, err := c.ListInstanceConsoleConnections(context.Background(), request)
	helpers.FatalIfError(err)
	var consoleConnections = lvar.Items
	if len(consoleConnections) > 0 {
		log.Printf("   Console Connections:")
		for _, v := range consoleConnections {
			log.Printf("     Console Connection Id : %s", *v.Id)
			log.Printf("     Compartment Id : %s", *v.CompartmentId)
			log.Printf("     State : %s", v.LifecycleState)
			log.Printf(" ")
		}
	}

}

func printConsoleHistories(c core.ComputeClient, compartmentId string) {

	request := core.ListConsoleHistoriesRequest{
		CompartmentId: common.String(compartmentId),
		InstanceId:    common.String(instanceId),
	}
	lvar, err := c.ListConsoleHistories(context.Background(), request)
	helpers.FatalIfError(err)
	var consoleHistories = lvar.Items
	if len(consoleHistories) > 0 {
		log.Printf("   Console Histories:")
		for _, v := range consoleHistories {
			log.Printf("     Console Connection Id : %s", *v.Id)
			log.Printf("     Compartment Id : %s", *v.CompartmentId)
			log.Printf("     State : %s", v.LifecycleState)
			log.Printf(" ")
		}
	}

}

func waitUnitlMoveCompletion(opcWorkRequestID *string) {
	if opcWorkRequestID != nil {
		log.Printf("   opc-work-request-id : %s", *opcWorkRequestID)
		log.Printf("   Querying the status of move operation using opc-work-request-id ")
		wc, err := workrequests.NewWorkRequestClientWithConfigurationProvider(common.DefaultConfigProvider())

		retryPolicy = getRetryPolicy()
		// Apply wait until work complete retryPolicy
		workRequest := workrequests.GetWorkRequestRequest{
			WorkRequestId: opcWorkRequestID,
			RequestMetadata: common.RequestMetadata{
				RetryPolicy: &retryPolicy,
			},
		}

		// GetWorkRequest get retried until the work request is in Succeeded status
		wr, err := wc.GetWorkRequest(context.Background(), workRequest)
		helpers.FatalIfError(err)

		if wr.Status != "" {
			log.Printf("   Final Work Status : %s, move operation complete", wr.Status)
		}
	}

}

func getRetryPolicy() common.RetryPolicy {

	// maximum times of retry
	attempts := uint(10)

	nextDuration := func(r common.OCIOperationResponse) time.Duration {
		// you might want wait longer for next retry when your previous one failed
		// this function will return the duration as:
		// 1s, 2s, 4s, 8s, 16s, 32s, 64s etc...
		return time.Duration(math.Pow(float64(2), float64(r.AttemptNumber-1))) * time.Second
	}

	var expectedWorkStatus = workrequests.WorkRequestStatusSucceeded

	// Get shouldRetry function based on GetWorkRequestResponse Status
	shouldRetry := func(r common.OCIOperationResponse) bool {
		if _, isServiceError := common.IsServiceError(r.Error); isServiceError {
			// not service error, could be network error or other errors which prevents
			// request send to server, will do retry here
			return true
		}

		if converted, ok := r.Response.(workrequests.GetWorkRequestResponse); ok {
			log.Printf("     WorkRequest Status : %s", converted.Status)
			// do the retry until WorkReqeut Status is Succeeded  - ignore case (BMI-2652)
			return converted.Status != expectedWorkStatus
		}

		return true
	}

	return common.NewRetryPolicy(attempts, shouldRetry, nextDuration)

}

func usage() {
	log.Printf("Please set the following environment variables to use ChangeInstanceCompartment()")
	log.Printf(" ")
	log.Printf("   INSTANCE_ID       # Required: Instance Id")
	log.Printf("   COMPARTMENT_ID    # Required: Target Compartment Id")
	log.Printf("   IF_MATCH          # Optional: ETag")
	log.Printf("   OPC_RETRY_TOKEN   # Optional: OPC retry token string")
	log.Printf(" ")
	os.Exit(1)
}

func parseEnvironmentVariables() {

	instanceId = os.Getenv("INSTANCE_ID")
	targetCompartmentId = os.Getenv("COMPARTMENT_ID")
	ifMatch = os.Getenv("IF_MATCH")
	opcRetryToken = os.Getenv("OPC_RETRY_TOKEN")

	if instanceId == "" || targetCompartmentId == "" {
		usage()
	}

	log.Printf("INSTANCE_ID     : %s", instanceId)
	log.Printf("COMPARTMENT_ID  : %s", targetCompartmentId)
	if ifMatch != "" {
		log.Printf("IF_MATCH        : %s", ifMatch)
	}
	if opcRetryToken != "" {
		log.Printf("OPC_RETRY_TOKEN : %s", opcRetryToken)
	}
}
