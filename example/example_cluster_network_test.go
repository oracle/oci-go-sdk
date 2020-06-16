// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Compute Management Services API
// This class provides an example of how you can create and manage a Cluster Network. It will:
//
//   * Create the InstanceConfiguration with an HPC shape
//   * Create a cluster network of size 1 based off of that configuration
//   * Wait for the cluster network to go to Running state
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

var (
	clusterNetwork string
)

// Example to showcase cluster network creation, and eventual teardown
func ExampleCreateAndWaitForRunningClusterNetwork() {
	ClusterNetworkParseEnvironmentVariables()

	ctx := context.Background()

	computeMgmtClient, err := core.NewComputeManagementClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	createInstanceConfigurationResponse, _ := createInstanceConfigurationWithHpcShape(ctx, computeMgmtClient, imageId, compartmentId)
	fmt.Println("Instance configuration created")

	instanceConfiguration := createInstanceConfigurationResponse.InstanceConfiguration

	clusterNetwork, _ := createClusterNetwork(ctx, computeMgmtClient, *instanceConfiguration.Id, subnetId, ad, compartmentId)
	fmt.Println("Cluster Network created")

	// waiting until the cluster network reaches running state
	pollUntilClusterNetworkInDesiredState(ctx, computeMgmtClient, clusterNetwork, core.ClusterNetworkLifecycleStateRunning)

	// clean up resources
	defer func() {
		terminateClusterNetwork(ctx, computeMgmtClient, *clusterNetwork.Id)
		fmt.Println("Terminated Cluster Network")

		deleteInstanceConfiguration(ctx, computeMgmtClient, *instanceConfiguration.Id)
		fmt.Println("Deleted Instance Configuration")
	}()

	// Output:
	// Instance configuration created
	// Cluster Network created
	// Cluster Network is RUNNING
	// Terminated Cluster Network
	// Deleted Instance Configuration
}

// Usage printing
func ClusterNetworkUsage() {
	log.Printf("Please set the following environment variables to run Cluster Network sample")
	log.Printf(" ")
	log.Printf("   IMAGE_ID       # Required: Image Id to use")
	log.Printf("   COMPARTMENT_ID    # Required: Compartment Id to use")
	log.Printf("   AD          # Required: AD to use")
	log.Printf("   SUBNET_ID   # Required: Subnet to use")
	log.Printf(" ")
	os.Exit(1)
}

// Args parser
func ClusterNetworkParseEnvironmentVariables() {

	imageId = os.Getenv("IMAGE_ID")
	compartmentId = os.Getenv("COMPARTMENT_ID")
	ad = os.Getenv("AD")
	subnetId = os.Getenv("SUBNET_ID")

	if imageId == "" ||
		compartmentId == "" ||
		ad == "" ||
		subnetId == "" {
		ClusterNetworkUsage()
	}

	log.Printf("IMAGE_ID     : %s", imageId)
	log.Printf("COMPARTMENT_ID  : %s", compartmentId)
	log.Printf("AD     : %s", ad)
	log.Printf("SUBNET_ID  : %s", subnetId)
}

// helper method to create a cluster network
func createClusterNetwork(ctx context.Context, client core.ComputeManagementClient, instanceConfigurationId string,
	subnetId string, availabilityDomain string, compartmentId string) (response core.CreateClusterNetworkResponse, err error) {

	displayName := "Cluster Network Example"
	size := 1

	placementConfigurationDetails := core.ClusterNetworkPlacementConfigurationDetails{
		AvailabilityDomain: &availabilityDomain,
		PrimarySubnetId:    &subnetId,
	}

	req := core.CreateClusterNetworkRequest{
		CreateClusterNetworkDetails: core.CreateClusterNetworkDetails{
			CompartmentId:          &compartmentId,
			DisplayName:            &displayName,
			PlacementConfiguration: &placementConfigurationDetails,
			InstancePools: []core.CreateClusterNetworkInstancePoolDetails{
				{
					Size: &size,
					InstanceConfigurationId: &instanceConfigurationId,
				},
			},
		},
	}

	response, err = client.CreateClusterNetwork(ctx, req)
	return
}

// helper method to create an instance configuration
func createInstanceConfigurationWithHpcShape(ctx context.Context, client core.ComputeManagementClient, imageId string, compartmentId string) (response core.CreateInstanceConfigurationResponse, err error) {
	vnicDetails := core.InstanceConfigurationCreateVnicDetails{}

	sourceDetails := core.InstanceConfigurationInstanceSourceViaImageDetails{
		ImageId: &imageId,
	}

	displayName := "Instance Configuration HPC Example"
	shape := "BM.HPC2.36"

	launchDetails := core.InstanceConfigurationLaunchInstanceDetails{
		CompartmentId:     &compartmentId,
		DisplayName:       &displayName,
		CreateVnicDetails: &vnicDetails,
		Shape:             &shape,
		SourceDetails:     &sourceDetails,
	}

	instanceDetails := core.ComputeInstanceDetails{
		LaunchDetails: &launchDetails,
	}

	configurationDetails := core.CreateInstanceConfigurationDetails{
		DisplayName:     &displayName,
		CompartmentId:   &compartmentId,
		InstanceDetails: &instanceDetails,
	}

	req := core.CreateInstanceConfigurationRequest{
		CreateInstanceConfiguration: configurationDetails,
	}

	response, err = client.CreateInstanceConfiguration(ctx, req)
	helpers.FatalIfError(err)

	return
}

// helper method to terminate a cluster network
func terminateClusterNetwork(ctx context.Context, client core.ComputeManagementClient,
	clusterNetworkId string) (response core.TerminateClusterNetworkResponse, err error) {

	req := core.TerminateClusterNetworkRequest{
		ClusterNetworkId: &clusterNetworkId,
	}

	response, err = client.TerminateClusterNetwork(ctx, req)
	helpers.FatalIfError(err)

	return
}

// helper method to pool until an cluster network reaches the specified desired state
func pollUntilClusterNetworkInDesiredState(ctx context.Context, computeMgmtClient core.ComputeManagementClient,
	clusterNetwork core.CreateClusterNetworkResponse, desiredState core.ClusterNetworkLifecycleStateEnum) {
	// should retry condition check which returns a bool value indicating whether to do retry or not
	// it checks the lifecycle status equals to Running or not for this case
	shouldRetryFunc := func(r common.OCIOperationResponse) bool {
		if converted, ok := r.Response.(core.GetClusterNetworkResponse); ok {
			return converted.LifecycleState != desiredState
		}
		return true
	}
	// create get cluster network request with a retry policy which takes a function
	// to determine shouldRetry or not
	pollingGetRequest := core.GetClusterNetworkRequest{
		ClusterNetworkId: clusterNetwork.Id,
		RequestMetadata:  helpers.GetRequestMetadataWithCustomizedRetryPolicy(shouldRetryFunc),
	}
	_, pollError := computeMgmtClient.GetClusterNetwork(ctx, pollingGetRequest)
	helpers.FatalIfError(pollError)
	fmt.Println("Cluster Network is", desiredState)
}
