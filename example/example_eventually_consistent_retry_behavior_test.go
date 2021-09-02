// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for using retry for SDK APIs

package example

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/oracle/oci-go-sdk/v47/common"
	"github.com/oracle/oci-go-sdk/v47/core"
	"github.com/oracle/oci-go-sdk/v47/example/helpers"
	"github.com/oracle/oci-go-sdk/v47/identity"
)

// This is a pretend OCID, an instance with this OCID does not actually exist.
// Therefore, GetInstance requests are guaranteed to fail with 404-NotAuthorizedOrNotFound.
const missingInstanceOcid = "ocid1.instance.oc1.phx.<uniqueId>"

// This example simulates the behaviors of retry strategies with respect to eventual consistency.
// The operation that is called that is eventually consistent is CreateGroup in the Identity service.
// After that, this example is making a number of GetInstance requests in the Compute service, which
// are guaranteed to fail with a 404-NotAuthorizedOrNotFound, because the OCID is not a real OCID
// of an instance.
// But it does simulate the behavior of the retries you would see if there were a replication delay
// due to the eventual consistency of the group.
//
// Note: This is a long running example, it takes over 4 minutes. That's why the "Output:" line has
// been changed to prevent the example from automatically running as a test.
func ExampleEventuallyConsistentRetryBehavior_Default() {
	// setup
	ctx := context.Background()

	coreClient, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)
	compartmentID, _ := common.DefaultConfigProvider().TenancyOCID()

	// this will set the eventually consistent timestamp, because the CreateGroup is eventually consistent and sets the timestamp
	groupId := createGroup(ctx, compartmentID)
	deleteGroup(ctx, groupId)

	// test

	defaultRetryPolicy := common.DefaultRetryPolicy()
	nonEcRetryPolicy := common.DefaultRetryPolicyWithoutEventualConsistency()

	fmt.Printf("EC retry policy:     %v\n", defaultRetryPolicy)
	fmt.Printf("Non-EC retry policy: %v\n", nonEcRetryPolicy)

	// Without retry policy, we do not see retries
	fmt.Printf("\nNo retry policy (expect immediate error):\n")
	var elapsed = getInstance(ctx, coreClient, missingInstanceOcid, nil)
	fmt.Printf("No retry policy (expect immediate error), elapsed less than three seconds? %v\n",
		getComparisonMessage(elapsed.String(), elapsed < time.Duration(3)*time.Second))

	// With the non-EC retry policy, we do not see a retry, because it doesn't consider eventual consistency.
	// Without eventual consistency, 404-NotAuthorizedOrNotFound are not retried.
	fmt.Printf("\nNon-EC retry policy (expect immediate error):\n")
	elapsed = getInstance(ctx, coreClient, missingInstanceOcid, &nonEcRetryPolicy)
	fmt.Printf("Non-EC retry policy (expect immediate error), elapsed less than three seconds? %v\n",
		getComparisonMessage(elapsed.String(), elapsed < time.Duration(3)*time.Second))

	// With the default retry policy, we do see retries, and this part takes a long time (about 4 minutes).
	// These retries on 404-NotAuthorizedOrNotFound only happen because there was an eventually consistent
	// operation in the recent past (CreateGroup).
	fmt.Printf("\nDefault retry policy (expect long wait, then error):\n")
	elapsed = getInstance(ctx, coreClient, missingInstanceOcid, &defaultRetryPolicy)
	fmt.Printf("Default retry policy (expect long wait, then error), elapsed about 4 minutes? %v\n",
		getComparisonMessage(elapsed.String(), (time.Duration(239)*time.Second < elapsed) && (elapsed < time.Duration(250)*time.Second)))

	// We use the the EC retry policy again, but by now we're outside the eventually consistent window, so we don't see retries anymore.
	fmt.Printf("\nDefault retry policy, but no more EC (end of window in the past? %v) (expect immediate error):\n",
		getComparisonMessage(fmt.Sprintf("now=%v, eow=%v", time.Now(), common.EcContext.GetEndOfWindow()),
			time.Now().After(*common.EcContext.GetEndOfWindow())))
	elapsed = getInstance(ctx, coreClient, missingInstanceOcid, &defaultRetryPolicy)
	fmt.Printf("Default retry policy, but no more EC (expect immediate error), elapsed less than three seconds? %v\n",
		getComparisonMessage(elapsed.String(), elapsed < time.Duration(3)*time.Second))

	// Output -- to enable this example as a test, change this line to "// Output:"
	// EC retry policy:     {MaximumNumberAttempts=9, MinSleepBetween=0, MaxSleepBetween=45, ExponentialBackoffBase=3.52, NonEventuallyConsistentPolicy={MaximumNumberAttempts=8, MinSleepBetween=0, MaxSleepBetween=30, ExponentialBackoffBase=2, NonEventuallyConsistentPolicy=<nil>}}
	// Non-EC retry policy: {MaximumNumberAttempts=8, MinSleepBetween=0, MaxSleepBetween=30, ExponentialBackoffBase=2, NonEventuallyConsistentPolicy=<nil>}
	//
	// No retry policy (expect immediate error):
	// Service error: NotAuthorizedOrNotFound. Authorization failed or requested resource not found. http status code: 404.
	// No retry policy (expect immediate error), elapsed less than three seconds? true
	//
	// Non-EC retry policy (expect immediate error):
	// Setting retry policy: {MaximumNumberAttempts=8, MinSleepBetween=0, MaxSleepBetween=30, ExponentialBackoffBase=2, NonEventuallyConsistentPolicy=<nil>}
	// Service error: NotAuthorizedOrNotFound. Authorization failed or requested resource not found. http status code: 404.
	// Non-EC retry policy (expect immediate error), elapsed less than three seconds? true
	//
	// Default retry policy (expect long wait, then error):
	// Setting retry policy: {MaximumNumberAttempts=9, MinSleepBetween=0, MaxSleepBetween=45, ExponentialBackoffBase=3.52, NonEventuallyConsistentPolicy={MaximumNumberAttempts=8, MinSleepBetween=0, MaxSleepBetween=30, ExponentialBackoffBase=2, NonEventuallyConsistentPolicy=<nil>}}
	// Service error: NotAuthorizedOrNotFound. Authorization failed or requested resource not found. http status code: 404.
	// Default retry policy (expect long wait, then error), elapsed about 4 minutes? true
	//
	// Default retry policy, but no more EC (end of window in the past? true) (expect immediate error):
	// Setting retry policy: {MaximumNumberAttempts=9, MinSleepBetween=0, MaxSleepBetween=45, ExponentialBackoffBase=3.52, NonEventuallyConsistentPolicy={MaximumNumberAttempts=8, MinSleepBetween=0, MaxSleepBetween=30, ExponentialBackoffBase=2, NonEventuallyConsistentPolicy=<nil>}}
	// Service error: NotAuthorizedOrNotFound. Authorization failed or requested resource not found. http status code: 404.
	// Default retry policy, but no more EC (expect immediate error), elapsed less than three seconds? true
}

// This example simulates the behaviors of retry strategies with respect to eventual consistency.
// The operation that is called that is eventually consistent is CreateGroup in the Identity service.
// After that, this example is making a number of GetInstance requests in the Compute service, which
// are guaranteed to fail with a 404-NotAuthorizedOrNotFound, because the OCID is not a real OCID
// of an instance.
// But it does simulate the behavior of the retries you would see if there were a replication delay
// due to the eventual consistency of the group.
// Instead of using the default retry strategy, which uses exponential backoff and a limited number of
// attempts, the retry strategy here uses unlimited attempts, but a limited amount of time.
//
// Note: This is a long running example, it takes over 4 minutes. That's why the "Output:" line has
// been changed to prevent the example from automatically running as a test.
func ExampleEventuallyConsistentRetryBehavior_UnlimitedAttempts() {
	// setup
	ctx := context.Background()

	coreClient, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)
	compartmentID, _ := common.DefaultConfigProvider().TenancyOCID()

	// this will set the eventually consistent timestamp, because the CreateGroup is eventually consistent and sets the timestamp
	groupId := createGroup(ctx, compartmentID)
	deleteGroup(ctx, groupId)

	// test

	maximumCumulativeBackoff := time.Duration(2) * time.Minute

	// retry unlimited number of times, up to two minutes
	nonEcRetryPolicy := common.NewRetryPolicyWithOptions(
		common.ReplaceWithValuesFromRetryPolicy(common.DefaultRetryPolicyWithoutEventualConsistency()),
		common.WithUnlimitedAttempts(maximumCumulativeBackoff),
		common.WithShouldRetryOperation(func(r common.OCIOperationResponse) bool {
			durationSinceInitialAttempt := time.Now().Sub(r.InitialAttemptTime)
			tooLong := durationSinceInitialAttempt > maximumCumulativeBackoff
			return common.DefaultShouldRetryOperation(r) && !tooLong
		}),
		common.WithNextDuration(func(r common.OCIOperationResponse) time.Duration {
			return time.Duration(math.Pow(float64(2), float64(r.AttemptNumber-1))) * time.Second
		}),
	)

	ecRetryPolicy := common.EventuallyConsistentRetryPolicy(nonEcRetryPolicy)

	fmt.Printf("EC retry policy    : %v\n", ecRetryPolicy)
	fmt.Printf("Non-EC retry policy: %v\n", nonEcRetryPolicy)

	// Without retry policy, we do not see retries
	fmt.Printf("\nNo retry policy (expect immediate error):\n")
	var elapsed = getInstance(ctx, coreClient, missingInstanceOcid, nil)
	fmt.Printf("No retry policy (expect immediate error), elapsed less than three seconds? %v\n",
		getComparisonMessage(elapsed.String(), elapsed < time.Duration(3)*time.Second))

	// With the non-EC retry policy, we do not see a retry, because it doesn't consider eventual consistency.
	// Without eventual consistency, 404-NotAuthorizedOrNotFound are not retried.
	fmt.Printf("\nNon-EC retry policy (expect immediate error):\n")
	elapsed = getInstance(ctx, coreClient, missingInstanceOcid, &nonEcRetryPolicy)
	fmt.Printf("Non-EC retry policy (expect immediate error), elapsed less than three seconds? %v\n",
		getComparisonMessage(elapsed.String(), elapsed < time.Duration(3)*time.Second))

	// With the EC retry policy, we do see retries, and this part takes a long time (about 4 minutes).
	// These retries on 404-NotAuthorizedOrNotFound only happen because there was an eventually consistent
	// operation in the recent past (CreateGroup).
	fmt.Printf("\nEC retry policy (expect long wait, then error):\n")
	elapsed = getInstance(ctx, coreClient, missingInstanceOcid, &ecRetryPolicy)
	fmt.Printf("EC retry policy (expect long wait, then error), elapsed about 4 minutes? %v\n",
		getComparisonMessage(elapsed.String(), (time.Duration(239)*time.Second < elapsed) && (elapsed < time.Duration(250)*time.Second)))

	// We use the the EC retry policy again, but by now we're outside the eventually consistent window, so we don't see retries anymore.
	fmt.Printf("\nEC retry policy, but no more EC (end of window in the past? %v) (expect immediate error):\n",
		getComparisonMessage(fmt.Sprintf("now=%v, eow=%v", time.Now(), common.EcContext.GetEndOfWindow()),
			time.Now().After(*common.EcContext.GetEndOfWindow())))
	elapsed = getInstance(ctx, coreClient, missingInstanceOcid, &ecRetryPolicy)
	fmt.Printf("EC retry policy, but no more EC (expect immediate error), elapsed less than three seconds? %v\n",
		getComparisonMessage(elapsed.String(), elapsed < time.Duration(3)*time.Second))

	// Output -- to enable this example as a test, change this line to "// Output:"
	// EC retry policy    : {MaximumNumberAttempts=9, MinSleepBetween=0, MaxSleepBetween=45, ExponentialBackoffBase=3.52, NonEventuallyConsistentPolicy={MaximumNumberAttempts=0, MinSleepBetween=0, MaxSleepBetween=30, ExponentialBackoffBase=2, NonEventuallyConsistentPolicy=<nil>}}
	// Non-EC retry policy: {MaximumNumberAttempts=0, MinSleepBetween=0, MaxSleepBetween=30, ExponentialBackoffBase=2, NonEventuallyConsistentPolicy=<nil>}
	//
	// No retry policy (expect immediate error):
	// Service error: NotAuthorizedOrNotFound. Authorization failed or requested resource not found. http status code: 404.
	// No retry policy (expect immediate error), elapsed less than three seconds? true
	//
	// Non-EC retry policy (expect immediate error):
	// Setting retry policy: {MaximumNumberAttempts=0, MinSleepBetween=0, MaxSleepBetween=30, ExponentialBackoffBase=2, NonEventuallyConsistentPolicy=<nil>}
	// Service error: NotAuthorizedOrNotFound. Authorization failed or requested resource not found. http status code: 404.
	// Non-EC retry policy (expect immediate error), elapsed less than three seconds? true
	//
	// EC retry policy (expect long wait, then error):
	// Setting retry policy: {MaximumNumberAttempts=9, MinSleepBetween=0, MaxSleepBetween=45, ExponentialBackoffBase=3.52, NonEventuallyConsistentPolicy={MaximumNumberAttempts=0, MinSleepBetween=0, MaxSleepBetween=30, ExponentialBackoffBase=2, NonEventuallyConsistentPolicy=<nil>}}
	// Service error: NotAuthorizedOrNotFound. Authorization failed or requested resource not found. http status code: 404.
	// EC retry policy (expect long wait, then error), elapsed about 4 minutes? true
	//
	// EC retry policy, but no more EC (end of window in the past? true) (expect immediate error):
	// Setting retry policy: {MaximumNumberAttempts=9, MinSleepBetween=0, MaxSleepBetween=45, ExponentialBackoffBase=3.52, NonEventuallyConsistentPolicy={MaximumNumberAttempts=0, MinSleepBetween=0, MaxSleepBetween=30, ExponentialBackoffBase=2, NonEventuallyConsistentPolicy=<nil>}}
	// Service error: NotAuthorizedOrNotFound. Authorization failed or requested resource not found. http status code: 404.
	// EC retry policy, but no more EC (expect immediate error), elapsed less than three seconds? true
}

func createGroup(ctx context.Context, compartmentID string) string {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := identity.CreateGroupRequest{CreateGroupDetails: identity.CreateGroupDetails{Name: common.String("go-sdk-demo-group"),
		CompartmentId: &compartmentID,
		Description:   common.String("Test for eventually consistent Identity resources"),
	}}

	// Send the request using the service client
	resp, err := client.CreateGroup(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	// fmt.Println(resp)

	return *resp.Group.Id
}

func deleteGroup(ctx context.Context, groupId string) {
	// Create a default authentication provider that uses the DEFAULT
	// profile in the configuration file.
	// Refer to <see href="https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File>the public documentation</see> on how to prepare a configuration file.
	client, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := identity.DeleteGroupRequest{GroupId: &groupId}

	// Send the request using the service client
	_, err = client.DeleteGroup(context.Background(), req)
	helpers.FatalIfError(err)

	// Retrieve value from the response.
	// fmt.Println(resp)
}

func getInstance(ctx context.Context, coreClient core.ComputeClient, instanceId string, retryPolicy *common.RetryPolicy) (elapsed time.Duration) {
	req := core.GetInstanceRequest{InstanceId: &instanceId}
	if retryPolicy != nil {
		fmt.Printf("Setting retry policy: %v\n", retryPolicy)
		req.RequestMetadata = common.RequestMetadata{
			RetryPolicy: retryPolicy,
		}
	}

	begin := time.Now()

	// Send the request using the service client
	_, err := coreClient.GetInstance(context.Background(), req)

	end := time.Now()

	// Retrieve value from the response.
	if failure, isServiceError := common.IsServiceError(err); isServiceError {
		fmt.Printf("Service error: %s. %s http status code: %d.\n", failure.GetCode(), failure.GetMessage(), failure.GetHTTPStatusCode())
	}

	elapsed = end.Sub(begin)
	return
}

func getComparisonMessage(value string, comparison bool) string {
	if comparison {
		return "true"
	}
	return fmt.Sprintf("false (%v)", value)
}
