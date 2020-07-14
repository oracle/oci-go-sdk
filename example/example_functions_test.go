// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for functions API

package example

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/oracle/oci-go-sdk/functions"
)

var (
	fnImage       string
	gwDisplayName = "OCI-GOSDK-Sample-Gateway"
	rtDisplyName  = "Default Route Table for OCI-GOSDK-Sample-VCN"
)

/*
	SETUP:

	This test requires that you have a [DEFAULT] OCI user profile setup e.g. in ~/.oci/config
	the DEFAULT user will be used in these tests, so any variables supplied must be compatible with that user

	This test requires 4 environment variables to be set:
	for these environment variables, see example/example_test.go {
		OCI_COMPARTMENT_ID
		OCI_AVAILABILITY_DOMAIN
		OCI_ROOT_COMPARTMENT_ID
    }
	OCI_FN_IMAGE : The URI of a publicly available image in the Oracle Cloud Infrastructure Registry (OCIR) e.g. phx.ocir.io/<tenancy-name>/<directory>/<image-name>:<image-tag>

	RUN:
	To run this test/example run:
	go test github.com/oracle/oci-go-sdk/example -run ExampleFunctionInvoke

*/
func ExampleFunctionInvoke() {
	managementClient, err := functions.NewFunctionsManagementClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)
	fnImage = os.Getenv("OCI_FN_IMAGE")

	ctx := context.Background()
	subnetID := createOrGetNetworkInfrastructure()
	// A subnet is required to expose and be able invoke functions.
	// In multiple AD regions, subnets can be created in multiple ADs to provide redundancy.
	fmt.Println("Network Layer Created")
	// An application's name must be unique per-compartment
	appName := "Example-Go-SDK-App"
	// A function's name must be unique per-application
	fnName := "Example-Go-SDK-Fn"
	// We must specify which compartment we want to create our Application in
	compartmentID := helpers.CompartmentID()

	createdApp := createApplication(ctx, managementClient, appName, compartmentID, []string{*subnetID})
	fmt.Println("Application Created:", *createdApp.DisplayName)

	gotApp := getReadyApplication(ctx, managementClient, createdApp.Id)
	fmt.Println("Application Got:", *gotApp.DisplayName)

	listedApps := listApplications(ctx, managementClient, compartmentID)
	fmt.Println("Applications Listed:", *listedApps[0].DisplayName)

	createdFn := createFunction(ctx, managementClient, fnName, createdApp.Id)
	fmt.Println("Function Created:", *createdFn.DisplayName)

	gotFn := getReadyFunction(ctx, managementClient, createdFn.Id)
	fmt.Println("Function Got:", *gotFn.DisplayName)

	listedFns := listFunctions(ctx, managementClient, createdApp.Id)
	fmt.Println("Functions Listed:", *listedFns[0].DisplayName)

	invokeClient, err := functions.NewFunctionsInvokeClientWithConfigurationProvider(common.DefaultConfigProvider(), *createdFn.InvokeEndpoint)
	helpers.FatalIfError(err)

	invokeFunction(ctx, invokeClient, createdFn.Id)

	fmt.Println("Function invoked")

	deleteFunction(ctx, managementClient, createdFn.Id)
	fmt.Println("Function Deleted:", *createdFn.DisplayName)

	deleteApplication(ctx, managementClient, createdApp.Id)
	fmt.Println("Application Deleted:", *createdApp.DisplayName)

	// Output:
	// Network Layer Created
	// Application Created: Example-Go-SDK-App
	// Application Got: Example-Go-SDK-App
	// Applications Listed: Example-Go-SDK-App
	// Function Created: Example-Go-SDK-Fn
	// Function Got: Example-Go-SDK-Fn
	// Functions Listed: Example-Go-SDK-Fn
	// Function invoked
	// Function Deleted: Example-Go-SDK-Fn
	// Application Deleted: Example-Go-SDK-App
}

func createApplication(ctx context.Context, client functions.FunctionsManagementClient, appName string, compartmentID *string, subnetIDs []string) functions.Application {
	details := functions.CreateApplicationDetails{
		CompartmentId: compartmentID,
		DisplayName:   &appName,
		SubnetIds:     subnetIDs,
	}

	request := functions.CreateApplicationRequest{CreateApplicationDetails: details}

	response, err := client.CreateApplication(ctx, request)
	helpers.FatalIfError(err)

	return response.Application
}

//Gets an application, if that application is not ready polls until the application is ready
func getReadyApplication(ctx context.Context, client functions.FunctionsManagementClient, appID *string) (app functions.Application) {
	metaWithRetry := helpers.GetRequestMetadataWithCustomizedRetryPolicy(shouldRetryGetApplication)

	request := functions.GetApplicationRequest{
		ApplicationId:   appID,
		RequestMetadata: metaWithRetry,
	}

	response, err := client.GetApplication(ctx, request)
	helpers.FatalIfError(err)

	return response.Application
}

func listApplications(ctx context.Context, client functions.FunctionsManagementClient, compartmentID *string) []functions.ApplicationSummary {
	request := functions.ListApplicationsRequest{CompartmentId: compartmentID}
	response, err := client.ListApplications(ctx, request)
	helpers.FatalIfError(err)
	return response.Items
}

func deleteApplication(ctx context.Context, client functions.FunctionsManagementClient, appID *string) {
	request := functions.DeleteApplicationRequest{ApplicationId: appID}

	_, err := client.DeleteApplication(ctx, request)
	helpers.FatalIfError(err)
	return
}

func createFunction(ctx context.Context, client functions.FunctionsManagementClient, fnName string, appID *string) functions.Function {
	memory := int64(128)
	details := functions.CreateFunctionDetails{
		DisplayName:   &fnName,
		ApplicationId: appID,
		Image:         &fnImage,
		MemoryInMBs:   &memory,
	}

	request := functions.CreateFunctionRequest{CreateFunctionDetails: details}

	response, err := client.CreateFunction(ctx, request)
	helpers.FatalIfError(err)

	return response.Function
}

func getReadyFunction(ctx context.Context, client functions.FunctionsManagementClient, fnID *string) functions.Function {
	metaWithRetry := helpers.GetRequestMetadataWithCustomizedRetryPolicy(shouldRetryGetFunction)

	request := functions.GetFunctionRequest{
		FunctionId:      fnID,
		RequestMetadata: metaWithRetry,
	}

	response, err := client.GetFunction(ctx, request)
	helpers.FatalIfError(err)

	return response.Function
}

func listFunctions(ctx context.Context, client functions.FunctionsManagementClient, appID *string) []functions.FunctionSummary {
	request := functions.ListFunctionsRequest{ApplicationId: appID}

	response, err := client.ListFunctions(ctx, request)
	helpers.FatalIfError(err)

	return response.Items
}

func invokeFunction(ctx context.Context, client functions.FunctionsInvokeClient, fnID *string) *string {
	// Retry function invocation with a standard back-off if we get a 404 in response.
	// This is in case the function creation has not yet completed by the time invocation is attempted
	metaWithRetry := helpers.GetRequestMetadataWithCustomizedRetryPolicy(shouldRetryInvokeFunction)
	requestBody := ioutil.NopCloser(bytes.NewReader([]byte("")))
	request := functions.InvokeFunctionRequest{
		FunctionId:         fnID,
		InvokeFunctionBody: requestBody,
		RequestMetadata:    metaWithRetry,
	}

	response, err := client.InvokeFunction(ctx, request)
	if err != nil {
		fmt.Println("Invoke Error:", err)
		return nil
	}
	resp := response.RawResponse
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Invoke Failed:", resp.StatusCode)
		return nil
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read invoke body:", err)
	}
	responseBody := string(bodyBytes)
	return &responseBody
}

func deleteFunction(ctx context.Context, client functions.FunctionsManagementClient, fnID *string) {
	request := functions.DeleteFunctionRequest{FunctionId: fnID}

	_, err := client.DeleteFunction(ctx, request)
	helpers.FatalIfError(err)
	return
}

func createOrGetNetworkInfrastructure() *string {
	c, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		fmt.Println("Network client request error:", err)
	}
	sn := CreateOrGetSubnet()
	gw := createOrGetInternetGateway(c, sn.VcnId)
	createOrGetRouteTable(c, gw.Id, sn.VcnId)

	return sn.Id
}

func createOrGetInternetGateway(c core.VirtualNetworkClient, vcnID *string) (gateway core.InternetGateway) {
	ctx := context.Background()

	//List Gateways
	listGWRequest := core.ListInternetGatewaysRequest{
		CompartmentId: helpers.CompartmentID(),
		VcnId:         vcnID,
		DisplayName:   &gwDisplayName,
	}

	listGWRespone, err := c.ListInternetGateways(ctx, listGWRequest)
	if err != nil {
		fmt.Println("Internet gateway list error:", err)
	}

	if len(listGWRespone.Items) >= 1 {
		//Gateway with name already exists
		gateway = listGWRespone.Items[0]
	} else {
		//Create new Gateway
		enabled := true
		createGWDetails := core.CreateInternetGatewayDetails{
			CompartmentId: helpers.CompartmentID(),
			IsEnabled:     &enabled,
			VcnId:         vcnID,
			DisplayName:   &gwDisplayName,
		}

		createGWRequest := core.CreateInternetGatewayRequest{CreateInternetGatewayDetails: createGWDetails}

		createGWResponse, err := c.CreateInternetGateway(ctx, createGWRequest)
		if err != nil {
			fmt.Println("Internet gateway create error:", err)
		}
		gateway = createGWResponse.InternetGateway
	}
	return
}

func createOrGetRouteTable(c core.VirtualNetworkClient, gatewayID, VcnID *string) (routeTable core.RouteTable) {
	ctx := context.Background()

	//List Route Table
	listRTRequest := core.ListRouteTablesRequest{
		CompartmentId: helpers.CompartmentID(),
		VcnId:         VcnID,
		DisplayName:   &rtDisplyName,
	}

	listRTResponse, err := c.ListRouteTables(ctx, listRTRequest)
	if err != nil {
		fmt.Println("Route table list error", err)
	}

	cidrRange := "0.0.0.0/0"
	rr := core.RouteRule{
		NetworkEntityId: gatewayID,
		Destination:     &cidrRange,
		DestinationType: core.RouteRuleDestinationTypeCidrBlock,
	}

	if len(listRTResponse.Items) >= 1 {

		//Default Route Table found and has at least 1 route rule
		if len(listRTResponse.Items[0].RouteRules) >= 1 {
			routeTable = listRTResponse.Items[0]
			//Default Route table needs route rule adding
		} else {

			updateRTDetails := core.UpdateRouteTableDetails{
				RouteRules: []core.RouteRule{rr},
			}

			updateRTRequest := core.UpdateRouteTableRequest{
				RtId: listRTResponse.Items[0].Id,
				UpdateRouteTableDetails: updateRTDetails,
			}

			updateRTResponse, err := c.UpdateRouteTable(ctx, updateRTRequest)
			if err != nil {
				fmt.Println("Error updating route table:", err)
			}
			routeTable = updateRTResponse.RouteTable
		}

	} else {
		//No default route table found
		fmt.Println("Error could not find VCN default route table, VCN OCID:", *VcnID, "Could not find route table:", rtDisplyName)
	}
	return
}

func shouldRetryGetApplication(response common.OCIOperationResponse) bool {
	createResponse, correctType := response.Response.(functions.GetApplicationResponse)
	if !correctType {
		fmt.Println("Retry attempt used incompatible response type, expected GetApplicationResponse, found:", reflect.TypeOf(response.Response))
	}
	if createResponse.LifecycleState != functions.ApplicationLifecycleStateActive {
		return true
	}
	return false
}

func shouldRetryGetFunction(response common.OCIOperationResponse) bool {
	createResponse, correctType := response.Response.(functions.GetFunctionResponse)
	if !correctType {
		fmt.Println("Retry attempt used incompatible response type, expected GetFunctionResponse, found:", reflect.TypeOf(response.Response))
	}
	if createResponse.LifecycleState != functions.FunctionLifecycleStateActive {
		return true
	}
	return false
}

func shouldRetryInvokeFunction(response common.OCIOperationResponse) bool {
	invokeResponse, correctType := response.Response.(functions.InvokeFunctionResponse)
	if !correctType {
		fmt.Println("Retry attempt used incompatible response type, expected InvokeFunctionResponse, found:", reflect.TypeOf(response.Response))
	}
	if invokeResponse.RawResponse.StatusCode == 404 {
		return true
	}
	return false
}
