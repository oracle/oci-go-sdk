// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for using retry for SDK APIs

package example

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/oracle/oci-go-sdk/v63/common"
	"github.com/oracle/oci-go-sdk/v63/example/helpers"
	"github.com/oracle/oci-go-sdk/v63/identity"
)

// ExampleRetry shows how to use default retry for Create and Delete groups, please
// refer to example_core_test.go->ExampleLaunchInstance for more examples
// The Retry behavior Precedence (Highest to lowest) is defined as below:
//   Operation level retry policy - setting request.RequestMetadata
//   Client level retry policy - setting client.SetCustomClientConfiguration
//   Global level retry policy - setting common.GlobalRetry

func ExampleRetry() {
	// create global level user customized retry policy
	// globalLevelPolicy := common.NewRetryPolicyWithOptions(
	//	  common.WithMaximumNumberAttempts(5),
	//    common.WithNextDuration(func(r common.OCIOperationResponse) time.Duration {
	//    	return time.Duration(math.Pow(float64(2), float64(r.AttemptNumber-1))) * time.Second
	//    }),
	// )
	// common.GlobalRetry = &globalLevelPolicy

	// create and delete group with retry
	client, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())

	// create a client level defined retry policy
	// clientLevelPolicy := common.NewRetryPolicyWithOptions(
	//	   common.WithMaximumNumberAttempts(6),
	//	   common.WithNextDuration(func(r common.OCIOperationResponse) time.Duration {
	//		   return time.Duration(math.Pow(float64(2), float64(r.AttemptNumber-1))) * time.Second
	//	   }),
	// )
	// client.SetCustomClientConfiguration(common.CustomClientConfiguration{
	//	  RetryPolicy: &clientLevelPolicy,
	// })

	ctx := context.Background()
	helpers.FatalIfError(clerr)

	request := identity.CreateGroupRequest{}
	request.CompartmentId = helpers.RootCompartmentID()
	request.Name = common.String("GoSDK_Sample_Group")
	request.Description = common.String("GoSDK Sample Group Description")

	// use SDK's default retry policy which deals with eventual consistency
	defaultRetryPolicy := common.DefaultRetryPolicy()

	// create request metadata for retry(request level retry)
	request.RequestMetadata = common.RequestMetadata{
		RetryPolicy: &defaultRetryPolicy,
	}

	resp, err := client.CreateGroup(ctx, request)
	helpers.FatalIfError(err)
	fmt.Println("Creating Group")

	// Get with polling
	shouldRetry := func(r common.OCIOperationResponse) bool {
		if _, isServiceError := common.IsServiceError(r.Error); isServiceError {
			// not service error, could be network error or other errors which prevents
			// request send to server, will do retry here
			return true
		}

		if converted, ok := r.Response.(identity.GetGroupResponse); ok {
			// do the retry until lifecycle state become active
			return converted.LifecycleState != identity.GroupLifecycleStateActive
		}

		return true
	}

	// maximum times of retry
	attempts := uint(10)

	lifecycleStateCheckRetryPolicy := common.NewRetryPolicyWithOptions(
		// since this retries on ANY error response, we don't need special handling for eventual consistency
		common.ReplaceWithValuesFromRetryPolicy(common.DefaultRetryPolicyWithoutEventualConsistency()),
		common.WithMaximumNumberAttempts(attempts),
		common.WithShouldRetryOperation(shouldRetry),
	)

	getRequest := identity.GetGroupRequest{
		GroupId: resp.Id,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: &lifecycleStateCheckRetryPolicy,
		},
	}

	_, errAfterPolling := client.GetGroup(ctx, getRequest)
	helpers.FatalIfError(errAfterPolling)
	fmt.Println("Group Created")

	defer func() {
		// if we've successfully created a group, make sure that we delete it
		rDel := identity.DeleteGroupRequest{
			GroupId: resp.Id,
			RequestMetadata: common.RequestMetadata{
				RetryPolicy: &defaultRetryPolicy,
			},
		}

		_, err = client.DeleteGroup(ctx, rDel)
		helpers.FatalIfError(err)
		fmt.Println("Group Deleted")
	}()

	// Output:
	// Creating Group
	// Group Created
	// Group Deleted
}

// ExampleCustomRetry shows how to use retry for Create and Delete groups, please
// refer to example_core_test.go->ExampleLaunchInstance for more examples
func ExampleCustomRetry() {
	// create and delete group with retry
	client, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	ctx := context.Background()
	helpers.FatalIfError(clerr)

	request := identity.CreateGroupRequest{}
	request.CompartmentId = helpers.RootCompartmentID()
	request.Name = common.String("GoSDK_Sample_Group")
	request.Description = common.String("GoSDK Sample Group Description")

	// maximum times of retry
	attempts := uint(10)

	// retry for all non-200 status code
	retryOnAllNon200ResponseCodes := func(r common.OCIOperationResponse) bool {
		return !(r.Error == nil && 199 < r.Response.HTTPResponse().StatusCode && r.Response.HTTPResponse().StatusCode < 300)
	}

	customRetryPolicy := common.NewRetryPolicyWithOptions(
		// since this retries on ANY non-2xx response, we don't need special handling for eventual consistency
		common.ReplaceWithValuesFromRetryPolicy(common.DefaultRetryPolicyWithoutEventualConsistency()),
		common.WithMaximumNumberAttempts(attempts),
		common.WithShouldRetryOperation(retryOnAllNon200ResponseCodes),
	)

	// create request metadata for retry
	request.RequestMetadata = common.RequestMetadata{
		RetryPolicy: &customRetryPolicy,
	}

	resp, err := client.CreateGroup(ctx, request)
	helpers.FatalIfError(err)
	fmt.Println("Creating Group")

	// Get with polling
	shouldRetry := func(r common.OCIOperationResponse) bool {
		if _, isServiceError := common.IsServiceError(r.Error); isServiceError {
			// not service error, could be network error or other errors which prevents
			// request send to server, will do retry here
			return true
		}

		if converted, ok := r.Response.(identity.GetGroupResponse); ok {
			// do the retry until lifecycle state become active
			return converted.LifecycleState != identity.GroupLifecycleStateActive
		}

		return true
	}

	lifecycleStateCheckRetryPolicy := common.NewRetryPolicyWithOptions(
		// since this retries on ANY error response, we don't need special handling for eventual consistency
		common.ReplaceWithValuesFromRetryPolicy(common.DefaultRetryPolicyWithoutEventualConsistency()),
		common.WithMaximumNumberAttempts(attempts),
		common.WithShouldRetryOperation(shouldRetry),
	)

	getRequest := identity.GetGroupRequest{
		GroupId: resp.Id,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: &lifecycleStateCheckRetryPolicy,
		},
	}

	_, errAfterPolling := client.GetGroup(ctx, getRequest)
	helpers.FatalIfError(errAfterPolling)
	fmt.Println("Group Created")

	defer func() {
		// if we've successfully created a group, make sure that we delete it
		rDel := identity.DeleteGroupRequest{
			GroupId: resp.Id,
			RequestMetadata: common.RequestMetadata{
				RetryPolicy: &customRetryPolicy,
			},
		}

		_, err = client.DeleteGroup(ctx, rDel)
		helpers.FatalIfError(err)
		fmt.Println("Group Deleted")
	}()

	// Output:
	// Creating Group
	// Group Created
	// Group Deleted
}

// ExampleUnlimitedAttemptsRetry shows how to use retry with unlimited retries, only limited by time,
// for Create and Delete groups, please refer to example_core_test.go->ExampleLaunchInstance for more examples
func ExampleUnlimitedAttemptsRetry() {
	// create and delete group with retry
	client, clerr := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	ctx := context.Background()
	helpers.FatalIfError(clerr)

	request := identity.CreateGroupRequest{}
	request.CompartmentId = helpers.RootCompartmentID()
	request.Name = common.String("GoSDK_Sample_Group")
	request.Description = common.String("GoSDK Sample Group Description")

	maximumCumulativeBackoff := time.Duration(2) * time.Minute

	// retry unlimited number of times, up to two minutes
	customRetryPolicy := common.NewRetryPolicyWithOptions(
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

	// create request metadata for retry
	request.RequestMetadata = common.RequestMetadata{
		RetryPolicy: &customRetryPolicy,
	}

	resp, err := client.CreateGroup(ctx, request)
	helpers.FatalIfError(err)
	fmt.Println("Creating Group")

	// Get with polling
	shouldRetry := func(r common.OCIOperationResponse) bool {
		if _, isServiceError := common.IsServiceError(r.Error); isServiceError {
			// not service error, could be network error or other errors which prevents
			// request send to server, will do retry here
			return true
		}

		if converted, ok := r.Response.(identity.GetGroupResponse); ok {
			// do the retry until lifecycle state become active
			return converted.LifecycleState != identity.GroupLifecycleStateActive
		}

		return true
	}

	// retry unlimited number of times, up to two minutes, until lifecycle state is active
	lifecycleStateCheckRetryPolicy := common.NewRetryPolicyWithOptions(
		// since this retries on ANY error response, we don't need special handling for eventual consistency
		common.ReplaceWithValuesFromRetryPolicy(common.DefaultRetryPolicyWithoutEventualConsistency()),
		common.WithUnlimitedAttempts(maximumCumulativeBackoff),
		common.WithShouldRetryOperation(func(r common.OCIOperationResponse) bool {
			durationSinceInitialAttempt := time.Now().Sub(r.InitialAttemptTime)
			tooLong := durationSinceInitialAttempt > maximumCumulativeBackoff
			return shouldRetry(r) && !tooLong
		}),
		common.WithNextDuration(func(r common.OCIOperationResponse) time.Duration {
			return time.Duration(math.Pow(float64(2), float64(r.AttemptNumber-1))) * time.Second
		}),
	)

	getRequest := identity.GetGroupRequest{
		GroupId: resp.Id,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: &lifecycleStateCheckRetryPolicy,
		},
	}

	_, errAfterPolling := client.GetGroup(ctx, getRequest)
	helpers.FatalIfError(errAfterPolling)
	fmt.Println("Group Created")

	defer func() {
		// if we've successfully created a group, make sure that we delete it
		rDel := identity.DeleteGroupRequest{
			GroupId: resp.Id,
			RequestMetadata: common.RequestMetadata{
				RetryPolicy: &customRetryPolicy,
			},
		}

		_, err = client.DeleteGroup(ctx, rDel)
		helpers.FatalIfError(err)
		fmt.Println("Group Deleted")
	}()

	// Output:
	// Creating Group
	// Group Created
	// Group Deleted
}
