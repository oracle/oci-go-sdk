// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Compute Management Services API
// This class provides an example of how you can create and use an Instance Configuration. It will:
//
//   * Create the InstanceConfiguration from input details
//   * Launching an instance with instance configuration
//   * Creating an instance configuration from a running instance
//   * Clean everything up

package example

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/example/helpers"
)

// Example to showcase instance configuration create and operations, and eventual teardown
func ExampleCreateAndUseInstanceConfiguration() {
	InstanceConfigsParseEnvironmentVariables()

	ctx := context.Background()

	computeMgmtClient, err := core.NewComputeManagementClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	computeClient, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create instance configuration
	createInstanceConfigurationResponse, _ := createInstanceConfiguration(ctx, computeMgmtClient, imageId, compartmentId)
	fmt.Println("Instance configuration created")

	instanceConfiguration := createInstanceConfigurationResponse.InstanceConfiguration

	launchInstanceConfigResponse, _ := launchInstanceConfiguration(ctx, computeMgmtClient, *instanceConfiguration.Id, subnetId, ad)

	instance := launchInstanceConfigResponse.Instance
	pollUntilDesiredInstanceState(ctx, computeClient, *instance.Id)
	fmt.Println("Instance launched")

	createInstanceConfigFromInstanceResponse, _ := createInstanceConfigFromInstance(ctx, computeMgmtClient, *instance.Id, compartmentId)
	fmt.Println("Instance configuration created from instance")

	instanceConfigFromInstance := createInstanceConfigFromInstanceResponse.InstanceConfiguration

	// clean up resources
	defer func() {
		_, _ = deleteInstanceConfiguration(ctx, computeMgmtClient, *instanceConfiguration.Id)
		fmt.Println("Deleted instance configuration")

		_, _ = deleteInstanceConfiguration(ctx, computeMgmtClient, *instanceConfigFromInstance.Id)
		fmt.Println("Deleted instance configuration created from instance")

		terminateInstance(ctx, computeClient, instance.Id)
	}()

	// Output:
	// Instance configuration created
	// Instance launched
	// Instance configuration created from instance
	// Deleted instance configuration
	// Deleted instance configuration created from instance
	// terminating instance
	// instance terminated
}

func createInstanceConfigFromInstance(ctx context.Context, client core.ComputeManagementClient, instanceId string,
	compartmentId string) (response core.CreateInstanceConfigurationResponse, err error) {

	displayName := "Instance Config From Instance Example"
	configurationDetails := core.CreateInstanceConfigurationFromInstanceDetails{
		DisplayName:   &displayName,
		CompartmentId: &compartmentId,
		InstanceId:    &instanceId,
	}

	req := core.CreateInstanceConfigurationRequest{
		CreateInstanceConfiguration: configurationDetails,
	}

	response, err = client.CreateInstanceConfiguration(ctx, req)
	helpers.FatalIfError(err)

	return
}

func launchInstanceConfiguration(ctx context.Context, client core.ComputeManagementClient, instanceConfigurationId string,
	subnetId string, availabilityDomain string) (response core.LaunchInstanceConfigurationResponse, err error) {

	req := core.LaunchInstanceConfigurationRequest{
		InstanceConfigurationId: &instanceConfigurationId,
		InstanceConfiguration: core.ComputeInstanceDetails{
			LaunchDetails: &core.InstanceConfigurationLaunchInstanceDetails{
				AvailabilityDomain: &availabilityDomain,
				CreateVnicDetails: &core.InstanceConfigurationCreateVnicDetails{
					SubnetId: &subnetId,
				},
			},
		},
	}

	response, err = client.LaunchInstanceConfiguration(ctx, req)
	helpers.FatalIfError(err)

	return
}

func pollUntilDesiredInstanceState(ctx context.Context, client core.ComputeClient, instanceId string) {
	// should retry condition check which returns a bool value indicating whether to do retry or not
	// it checks the lifecycle status equals to Running or not for this case
	shouldRetryFunc := func(r common.OCIOperationResponse) bool {
		if converted, ok := r.Response.(core.GetInstanceResponse); ok {
			return converted.LifecycleState != core.InstanceLifecycleStateRunning
		}
		return true
	}

	// create get instance request with a retry policy which takes a function
	// to determine shouldRetry or not
	pollingGetRequest := core.GetInstanceRequest{
		InstanceId:      &instanceId,
		RequestMetadata: helpers.GetRequestMetadataWithCustomizedRetryPolicy(shouldRetryFunc),
	}

	_, pollError := client.GetInstance(ctx, pollingGetRequest)
	helpers.FatalIfError(pollError)
}

// Usage printing
func InstanceConfigsUsage() {
	log.Printf("Please set the following environment variables to run Instance Configuration sample")
	log.Printf(" ")
	log.Printf("   IMAGE_ID       # Required: Image Id to use")
	log.Printf("   COMPARTMENT_ID    # Required: Compartment Id to use")
	log.Printf("   AD          # Required: AD to use")
	log.Printf("   SUBNET_ID   # Required: Subnet to use")
	log.Printf(" ")
	os.Exit(1)
}

// Args parser
func InstanceConfigsParseEnvironmentVariables() {

	imageId = os.Getenv("IMAGE_ID")
	compartmentId = os.Getenv("COMPARTMENT_ID")
	ad = os.Getenv("AD")
	subnetId = os.Getenv("SUBNET_ID")

	if imageId == "" ||
		compartmentId == "" ||
		ad == "" ||
		subnetId == "" {
		InstanceConfigsUsage()
	}

	log.Printf("IMAGE_ID     : %s", imageId)
	log.Printf("COMPARTMENT_ID  : %s", compartmentId)
	log.Printf("AD     : %s", ad)
	log.Printf("SUBNET_ID  : %s", subnetId)
}
