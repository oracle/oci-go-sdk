// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Container Engine API

package example

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/core"
	"strings"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/containerengine"
	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/oracle/oci-go-sdk/identity"
)

// Example for how to do CRUD on cluster, how to get kubernets config and
// how to work with WorkRequest
func ExampleClusterCRUD() {
	ctx := context.Background()
	c, clerr := containerengine.NewContainerEngineClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)

	// create network resources for cluster.
	// this sample is to demonstrate how to use cluster APIs
	// for more configuration setup, please refer to the link here:
	// https://docs.cloud.oracle.com/Content/ContEng/Concepts/contengnetworkconfig.htm
	vcnID, subnet1ID, _ := createVCNWithSubnets(ctx)

	defaulKubetVersion := getDefaultKubernetesVersion(c)
	createClusterResp := createCluster(ctx, c, vcnID, defaulKubetVersion, subnet1ID)

	// wait until work request complete
	workReqResp := waitUntilWorkRequestComplete(c, createClusterResp.OpcWorkRequestId)
	fmt.Println("cluster created")

	// update cluster with a new name and upgrade the kubernets version
	updateReq := containerengine.UpdateClusterRequest{}

	// please see the document here for actionType values:
	// https://docs.cloud.oracle.com/api/#/en/containerengine/20180222/datatypes/WorkRequestResource
	clusterID := getResourceID(workReqResp.Resources, containerengine.WorkRequestResourceActionTypeCreated, "CLUSTER")
	updateReq.ClusterId = clusterID
	defer deleteCluster(ctx, c, clusterID)
	updateReq.Name = common.String("GOSDK_Sample_New_CE")

	getReq := containerengine.GetClusterRequest{
		ClusterId: updateReq.ClusterId,
	}

	getResp, err := c.GetCluster(ctx, getReq)
	// check for upgrade versions
	if len(getResp.Cluster.AvailableKubernetesUpgrades) > 0 {
		// if newer version available, set it for upgrade
		updateReq.KubernetesVersion = common.String(getResp.Cluster.AvailableKubernetesUpgrades[0])
	}

	updateResp, err := c.UpdateCluster(ctx, updateReq)
	helpers.FatalIfError(err)
	fmt.Println("updating cluster")

	// wait until update complete
	workReqResp = waitUntilWorkRequestComplete(c, updateResp.OpcWorkRequestId)
	fmt.Println("cluster updated")

	// get cluster
	getResp, err = c.GetCluster(ctx, getReq)
	helpers.FatalIfError(err)

	fmt.Printf("cluster name updated to %s\n", *getResp.Name)

	// Output:
	// create VCN complete
	// create subnet1 complete
	// create subnet2 complete
	// creating cluster
	// cluster created
	// updating cluster
	// cluster updated
	// cluster name updated to GOSDK_Sample_New_CE
	// deleting cluster
}

// Example for NodePool
func ExampleNodePoolCRUD() {
	ctx := context.Background()
	c, clerr := containerengine.NewContainerEngineClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)

	compute, err := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	identityClient, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)
	req := identity.ListAvailabilityDomainsRequest{}
	req.CompartmentId = helpers.CompartmentID()
	ads, err := identityClient.ListAvailabilityDomains(ctx, req)
	helpers.FatalIfError(err)

	// create network resources for cluster
	vcnID, subnet1ID, subnet2ID := createVCNWithSubnets(ctx)

	// create cluster
	kubeVersion := getDefaultKubernetesVersion(c)
	createClusterResp := createCluster(ctx, c, vcnID, kubeVersion, subnet1ID)

	// wait until work request complete
	workReqResp := waitUntilWorkRequestComplete(c, createClusterResp.OpcWorkRequestId)
	fmt.Println("cluster created")
	clusterID := getResourceID(workReqResp.Resources, containerengine.WorkRequestResourceActionTypeCreated, "CLUSTER")

	// get Image Id
	image := getImageID(ctx, compute)

	// create NodePool
	createNodePoolReq := containerengine.CreateNodePoolRequest{}
	createNodePoolReq.CompartmentId = helpers.CompartmentID()
	createNodePoolReq.Name = common.String("GOSDK_SAMPLE_NP")
	createNodePoolReq.ClusterId = clusterID
	createNodePoolReq.KubernetesVersion = common.String(kubeVersion)
	createNodePoolReq.NodeSourceDetails = containerengine.NodeSourceViaImageDetails{ImageId: image.Id}
	createNodePoolReq.NodeShape = common.String("VM.Standard1.1")
	createNodePoolReq.InitialNodeLabels = []containerengine.KeyValue{{Key: common.String("foo"), Value: common.String("bar")}}
	createNodePoolReq.NodeConfigDetails = &containerengine.CreateNodePoolNodeConfigDetails{
		PlacementConfigs: make([]containerengine.NodePoolPlacementConfigDetails, 0, len(ads.Items)),
		Size:             common.Int(len(ads.Items)),
	}

	for i := 0; i < len(ads.Items); i++ {
		createNodePoolReq.NodeConfigDetails.PlacementConfigs = append(createNodePoolReq.NodeConfigDetails.PlacementConfigs,
			containerengine.NodePoolPlacementConfigDetails{
				AvailabilityDomain: ads.Items[i].Name,
				SubnetId:           &subnet2ID,
			})
	}

	createNodePoolResp, err := c.CreateNodePool(ctx, createNodePoolReq)
	helpers.FatalIfError(err)
	fmt.Println("creating nodepool")

	workReqResp = waitUntilWorkRequestComplete(c, createNodePoolResp.OpcWorkRequestId)
	fmt.Println("nodepool created")

	nodePoolID := getResourceID(workReqResp.Resources, containerengine.WorkRequestResourceActionTypeCreated, "NODEPOOL")

	defer func() {
		deleteNodePool(ctx, c, nodePoolID)
		deleteCluster(ctx, c, clusterID)
	}()

	// update NodePool
	updateNodePoolReq := containerengine.UpdateNodePoolRequest{
		NodePoolId: nodePoolID,
	}

	updateNodePoolReq.Name = common.String("GOSDK_SAMPLE_NP_NEW")
	updateNodePoolResp, err := c.UpdateNodePool(ctx, updateNodePoolReq)
	helpers.FatalIfError(err)
	fmt.Println("updating nodepool")

	workReqResp = waitUntilWorkRequestComplete(c, updateNodePoolResp.OpcWorkRequestId)
	fmt.Println("nodepool updated")

	// Output:
	// create VCN complete
	// create subnet1 complete
	// create subnet2 complete
	// creating cluster
	// cluster created
	// creating nodepool
	// nodepool created
	// updating nodepool
	// nodepool updated
	// deleting nodepool
	// deleting cluster
}

func ExampleKubeConfig() {
	ctx := context.Background()
	c, clerr := containerengine.NewContainerEngineClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)

	clusterID := common.String("[YOUR CLUSTER ID]")
	req := containerengine.CreateKubeconfigRequest{
		ClusterId: clusterID,
	}

	_, err := c.CreateKubeconfig(ctx, req)
	helpers.FatalIfError(err)
	fmt.Println("create kubeconfig")

	// Output:
	// create kubeconfig
}

// Example for work request query
func ExampleWorkRequestQuery() {
	ctx := context.Background()
	c, clerr := containerengine.NewContainerEngineClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)

	workRequestID := common.String("[YOUR WORK REQUEST ID]")
	listErrorReq := containerengine.ListWorkRequestErrorsRequest{
		CompartmentId: helpers.CompartmentID(),
		WorkRequestId: workRequestID,
	}

	_, err := c.ListWorkRequestErrors(ctx, listErrorReq)
	helpers.FatalIfError(err)
	fmt.Println("list work request errors")

	listLogReq := containerengine.ListWorkRequestLogsRequest{
		CompartmentId: helpers.CompartmentID(),
		WorkRequestId: workRequestID,
	}

	_, err = c.ListWorkRequestLogs(ctx, listLogReq)
	helpers.FatalIfError(err)
	fmt.Println("list work request logs")

	// Output:
	// list work request errors
	// list work request logs
}

// wait until work request finish
func waitUntilWorkRequestComplete(client containerengine.ContainerEngineClient, workReuqestID *string) containerengine.GetWorkRequestResponse {
	// retry GetWorkRequest call until TimeFinished is set
	shouldRetryFunc := func(r common.OCIOperationResponse) bool {
		return r.Response.(containerengine.GetWorkRequestResponse).TimeFinished == nil
	}

	getWorkReq := containerengine.GetWorkRequestRequest{
		WorkRequestId:   workReuqestID,
		RequestMetadata: helpers.GetRequestMetadataWithCustomizedRetryPolicy(shouldRetryFunc),
	}

	getResp, err := client.GetWorkRequest(context.Background(), getWorkReq)
	helpers.FatalIfError(err)
	return getResp
}

// create a cluster
func createCluster(
	ctx context.Context,
	client containerengine.ContainerEngineClient,
	vcnID, kubernetesVersion, subnet1ID string) containerengine.CreateClusterResponse {
	req := containerengine.CreateClusterRequest{}
	req.Name = common.String("GOSDK_Sample_CE")
	req.CompartmentId = helpers.CompartmentID()
	req.VcnId = common.String(vcnID)
	req.KubernetesVersion = common.String(kubernetesVersion)
	req.Options = &containerengine.ClusterCreateOptions{
		ServiceLbSubnetIds: []string{subnet1ID},
	}

	fmt.Println("creating cluster")
	resp, err := client.CreateCluster(ctx, req)
	helpers.FatalIfError(err)

	return resp
}

// delete a cluster
func deleteCluster(ctx context.Context, client containerengine.ContainerEngineClient, clusterID *string) {
	deleteReq := containerengine.DeleteClusterRequest{
		ClusterId: clusterID,
	}

	client.DeleteCluster(ctx, deleteReq)

	fmt.Println("deleting cluster")
}

// delete a node pool
func deleteNodePool(ctx context.Context, client containerengine.ContainerEngineClient, nodePoolID *string) {
	deleteReq := containerengine.DeleteNodePoolRequest{
		NodePoolId: nodePoolID,
	}

	client.DeleteNodePool(ctx, deleteReq)

	fmt.Println("deleting nodepool")
}

// create VCN and subnets, return id of these three resources
func createVCNWithSubnets(ctx context.Context) (vcnID, subnet1ID, subnet2ID string) {
	// create a new VCN
	vcn := CreateOrGetVcn()
	fmt.Println("create VCN complete")

	subnet1 := CreateOrGetSubnetWithDetails(common.String(subnetDisplayName1), common.String("10.0.1.0/24"), common.String("subnetdns1"), nil)
	fmt.Println("create subnet1 complete")

	subnet2 := CreateOrGetSubnetWithDetails(common.String(subnetDisplayName2), common.String("10.0.2.0/24"), common.String("subnetdns2"), nil)
	fmt.Println("create subnet2 complete")

	return *vcn.Id, *subnet1.Id, *subnet2.Id
}

func getDefaultKubernetesVersion(client containerengine.ContainerEngineClient) string {
	getClusterOptionsReq := containerengine.GetClusterOptionsRequest{
		ClusterOptionId: common.String("all"),
	}
	getClusterOptionsResp, err := client.GetClusterOptions(context.Background(), getClusterOptionsReq)
	helpers.FatalIfError(err)
	kubernetesVersion := getClusterOptionsResp.KubernetesVersions

	if len(kubernetesVersion) < 1 {
		fmt.Println("Kubernetes version not available")
	}

	return kubernetesVersion[0]
}

// getResourceID return a resource ID based on the filter of resource actionType and entityType
func getResourceID(resources []containerengine.WorkRequestResource, actionType containerengine.WorkRequestResourceActionTypeEnum, entityType string) *string {
	for _, resource := range resources {
		if resource.ActionType == actionType && strings.ToUpper(*resource.EntityType) == entityType {
			return resource.Identifier
		}
	}

	fmt.Println("cannot find matched resources")
	return nil
}

func getImageID(ctx context.Context, c core.ComputeClient) core.Image {
	request := core.ListImagesRequest{
		CompartmentId:   helpers.CompartmentID(),
		OperatingSystem: common.String("Oracle Linux"),
		Shape:           common.String("VM.Standard1.1"),
	}

	r, err := c.ListImages(ctx, request)
	helpers.FatalIfError(err)

	return r.Items[0]
}
