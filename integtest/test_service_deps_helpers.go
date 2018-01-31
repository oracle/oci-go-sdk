// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package integtest

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/database"
	"github.com/oracle/oci-go-sdk/loadbalancer"
)

const (
	vcnDisplayName               = "GOSDK2_Test_Deps_VCN"
	subnetDisplayName            = "GOSDK2_Test_Deps_Subnet"
	instanceDisplayName          = "GOSDK2_Test_Deps_Instance"
	consoleHistoryDisplayName    = "GOSDK2_Test_Deps_ConsoleHistory"
	crossConnectDisplayName      = "GOSDK2_Test_Deps_CrossConnect"
	crossConnectGroupDisplayName = "GOSDK2_Test_Deps_CrossConnectGroup"
	peeringGatewaysDisplayName   = "GOSDK2_Test_Deps_LocalPeeringGateway"
	virtualCircuitDisplayName    = "GOSDK2_Test_Deps_VirtualCircuits"
	dbSystemDisplayName          = "GOSDK2_Test_Deps_DatabaseSystem"
	dbHomeDisplayName            = "GOSDK2_Test_Deps_DatabaseHome"
	loadbalancerDisplayName      = "GOSDK2_Test_Deps_Loadbalancer"
)

/*
// CreateOrGetDatabase creates a new database or return the existing one
// this can be used for tests that depends on a database which can save
// time to run integration test.
// Notes: the database created within this method will not be deleted
// so next time, same tests can run much faster. CreateDatabase API will
// be covered in different test cases which will be deleted
func createOrGetDatabase(t *testing.T) database.DbSystemSummary {
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := database.ListDbSystemsRequest{CompartmentId: common.String(getTenancyID())}

	r, err := c.ListDbSystems(context.Background(), request)
	failIfError(t, err)

	return r.Items[0]
}

func listDatabases(t *testing.T) []database.DatabaseSummary {
	dbHome := createOrGetDBHome(t)

	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := database.ListDatabasesRequest{
		CompartmentId: common.String(getTenancyID()),
		DbHomeId:      dbHome.Id,
	}

	r, err := c.ListDatabases(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, r.Items)

	return r.Items
}

func createOrGetDBHome(t *testing.T) database.CreateDbHomeDetails {}

func listDBHome(t *testing.T) []database.DbHome {
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := database.ListDbHomesRequest{
		CompartmentId: common.String(getTenancyID()),
		DbHomeId:      dbHome.Id,
	}

	r, err := c.ListDatabases(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, r.Items)

	return r.Items
}*/

func createOrGetDBSystem(t *testing.T) database.DbSystem {
	/*dbSystems := listDBSystems(t)

	for _, element := range dbSystems {
		assert.NotEmpty(t, element)
		if *element.DisplayName == dbSystemDisplayName {
			// got the result, return it
			return
		}
	}*/

	// create a new database system
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := database.LaunchDbSystemRequest{}
	request.AvailabilityDomain = common.String(validAD())
	request.CompartmentId = common.String(getCompartmentID())
	request.CpuCoreCount = common.Int(4)
	request.DatabaseEdition = "STANDARD_EDITION"
	request.Hostname = common.String("GOSDK2_Test_Deps_HostName")
	request.DisplayName = common.String(dbSystemDisplayName)
	request.Shape = common.String("VM.Standard1.2") // TODO get it via ListDbSystemShapes API

	pwd, err := os.Getwd()
	failIfError(t, err)
	buffer, err := ioutil.ReadFile(pwd + "/../resources/test_rsa.pub")
	failIfError(t, err)
	request.SshPublicKeys = []string{string(buffer)}

	subnet := createOrGetSubnet(t)
	request.SubnetId = subnet.Id

	request.DbHome.DisplayName = common.String(dbHomeDisplayName)
	request.DbHome.DbVersion = common.String("11.2.0.4") // TODO get it via ListDbVersions

	request.DbHome.Database.AdminPassword = common.String("OraclE12--")
	request.DbHome.Database.DbName = common.String("GOSDK2_Test_Deps_Database_Name")

	r, err := c.LaunchDbSystem(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)

	return r.DbSystem
}

func listDBSystems(t *testing.T) []database.DbSystemSummary {
	c, clerr := database.NewDatabaseClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := database.ListDbSystemsRequest{
		CompartmentId: common.String(getTenancyID()),
	}

	r, err := c.ListDbSystems(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, r.Items)

	return r.Items
}

func createOrGetVcn(t *testing.T) (vcn core.Vcn) {
	vcnItems := listVcns(t)

	for _, element := range vcnItems {
		if *element.DisplayName == vcnDisplayName {
			vcn = element
			assert.NotEmpty(t, vcn.Id)

			// VCN already created, return it
			return
		}
	}

	// create a new VCN
	// Notes: here will not destroy it. so for test cases depends on VCN can reuse it in next run
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := core.CreateVcnRequest{}
	request.CidrBlock = common.String("10.0.0.0/16")
	request.CompartmentId = common.String(getCompartmentID())
	request.DisplayName = common.String(vcnDisplayName)
	request.DnsLabel = common.String("DNSLabel")

	r, err := c.CreateVcn(context.Background(), request)
	failIfError(t, err)
	vcn = r.Vcn

	assert.NotEmpty(t, vcn.Id)
	return
}

func listVcns(t *testing.T) []core.Vcn {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListVcnsRequest{
		CompartmentId: common.String(getCompartmentID()),
	}

	r, err := c.ListVcns(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	return r.Items
}

func createOrGetSubnet(t *testing.T) (subnet core.Subnet) {
	subnets := listSubnets(t)

	// check if the subnet has already been created
	for _, element := range subnets {
		if *element.DisplayName == subnetDisplayName {
			// find the subnet, return it
			subnet = element
			assert.NotEmpty(t, subnet)
			assert.NotEmpty(t, subnet.Id)
			return
		}
	}

	// create a new subnet
	vcn := createOrGetVcn(t)
	request := core.CreateSubnetRequest{}
	request.CompartmentId = common.String(getCompartmentID())
	request.CidrBlock = common.String("10.0.0.0/24")
	request.VcnId = vcn.Id
	request.DisplayName = common.String(subnetDisplayName)
	request.AvailabilityDomain = common.String(validAD())
	request.DnsLabel = common.String("DNSLabel")

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	r, err := c.CreateSubnet(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, r.Subnet)

	subnet = r.Subnet
	assert.NotEmpty(t, subnet.Id)
	return
}

func listSubnets(t *testing.T) []core.Subnet {
	vcn := createOrGetVcn(t)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListSubnetsRequest{
		CompartmentId: common.String(getCompartmentID()),
		VcnId:         vcn.Id,
	}
	r, err := c.ListSubnets(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	return r.Items
}

// list the available images in the test compartment
func listImages(t *testing.T) []core.Image {
	return listImagesByDisplayName(t, nil)
}

func listImagesByDisplayName(t *testing.T, displayName *string) []core.Image {
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListImagesRequest{
		CompartmentId: common.String(getCompartmentID()),
	}

	if displayName != nil {
		request.DisplayName = displayName
	}

	r, err := c.ListImages(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, r.Items)

	return r.Items
}

// list the boot volumes in the specified compartment and Availability Domain.
func listBootVolumes(t *testing.T) []core.BootVolume {
	// this line make sure boot volumes is created
	createOrGetInstance(t)

	c, clerr := core.NewBlockstorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListBootVolumesRequest{
		AvailabilityDomain: common.String(validAD()),
		CompartmentId:      common.String(getCompartmentID()),
	}

	r, err := c.ListBootVolumes(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, r.Items)
	return r.Items
}

func listBootVolumeAttachments(t *testing.T) []core.BootVolumeAttachment {
	// this line make sure boot volumes is created
	createOrGetInstance(t)

	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListBootVolumeAttachmentsRequest{
		AvailabilityDomain: common.String(validAD()),
		CompartmentId:      common.String(getCompartmentID()),
	}

	r, err := c.ListBootVolumeAttachments(context.Background(), request)
	assert.NoError(t, err)
	assert.NotEmpty(t, r.Items)
	return r.Items
}

func listShapes(t *testing.T) []core.Shape {
	return listShapesForImage(t, nil)
}

func listShapesForImage(t *testing.T, imageID *string) []core.Shape {
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	if imageID == nil {
		images := listImages(t)

		assert.NotEmpty(t, images)
		assert.NotEqual(t, len(images), 0)
		imageID = images[0].Id
	}

	request := core.ListShapesRequest{
		CompartmentId: common.String(getCompartmentID()),
		ImageId:       imageID,
	}

	r, err := c.ListShapes(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, r.Items)
	return r.Items
}

func createOrGetInstance(t *testing.T) core.Instance {
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	listRequest := core.ListInstancesRequest{}
	listRequest.CompartmentId = common.String(getCompartmentID())
	listRequest.AvailabilityDomain = common.String(validAD())

	listResp, err := c.ListInstances(context.Background(), listRequest)
	failIfError(t, err)
	assert.NotEmpty(t, listResp)

	// check if instance exists or not
	for _, element := range listResp.Items {
		if *element.DisplayName == instanceDisplayName {
			return element
		}
	}

	// create a new instance
	createRequest := core.LaunchInstanceRequest{}
	createRequest.CompartmentId = common.String(getCompartmentID())
	createRequest.DisplayName = common.String(instanceDisplayName)
	createRequest.AvailabilityDomain = common.String(validAD())
	createRequest.SubnetId = createOrGetSubnet(t).Id

	// search image by display name to make integration test running more relaible
	// i.e. ServiceLimitExeceed error etc...
	images := listImagesByDisplayName(t, common.String("Oracle-Linux-7.4-2018.01.20-0"))
	assert.NotEmpty(t, images)
	assert.NotEqual(t, len(images), 0)
	createRequest.ImageId = images[0].Id

	shapes := listShapesForImage(t, createRequest.ImageId)
	assert.NotEqual(t, len(shapes), 0)
	createRequest.Shape = shapes[0].Shape

	createResp, err := c.LaunchInstance(context.Background(), createRequest)
	assert.NotEmpty(t, createResp, fmt.Sprint(createResp))
	assert.NoError(t, err)

	// TODO wait for instance states
	return createResp.Instance
}

func createOrGetCrossConnectGroup(t *testing.T) core.CrossConnectGroup {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	listRequest := core.ListCrossConnectGroupsRequest{}
	listRequest.CompartmentId = common.String(getCompartmentID())
	listRequest.DisplayName = common.String(crossConnectGroupDisplayName)

	listResp, err := c.ListCrossConnectGroups(context.Background(), listRequest)
	failIfError(t, err)
	assert.NotEmpty(t, listResp)

	// if connect group exist, return it
	for _, element := range listResp.Items {
		if *element.DisplayName == crossConnectGroupDisplayName {
			return element
		}
	}

	// create a new one
	createRequest := core.CreateCrossConnectGroupRequest{}
	createRequest.CompartmentId = common.String(getCompartmentID())
	createRequest.DisplayName = common.String(crossConnectGroupDisplayName)

	createResp, err := c.CreateCrossConnectGroup(context.Background(), createRequest)
	failIfError(t, err)

	assert.NotEmpty(t, createResp)
	return createResp.CrossConnectGroup
}

func createOrGetCrossConnect(t *testing.T) core.CrossConnect {
	crossConnects := listCrossConnects(t)

	// if connect group exist, return it
	for _, element := range crossConnects {
		if *element.DisplayName == crossConnectDisplayName {
			return element
		}
	}

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	// create a new one
	request := core.CreateCrossConnectRequest{}
	request.CompartmentId = common.String(getCompartmentID())

	crossConnectGroup := createOrGetCrossConnectGroup(t)
	request.CrossConnectGroupId = crossConnectGroup.Id
	request.DisplayName = common.String(crossConnectDisplayName)

	locations := listCrossConnectLocations(t)
	assert.NotEmpty(t, locations)
	assert.NotEqual(t, len(locations), 0)
	request.LocationName = locations[0].Name
	request.PortSpeedShapeName = common.String("10 Gbps")

	resp, err := c.CreateCrossConnect(context.Background(), request)
	failIfError(t, err)

	assert.NotEmpty(t, resp)
	return resp.CrossConnect
}

func listCrossConnects(t *testing.T) []core.CrossConnect {
	crossConnectGroup := createOrGetCrossConnectGroup(t)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := core.ListCrossConnectsRequest{}
	request.CompartmentId = common.String(getCompartmentID())
	request.DisplayName = common.String(crossConnectDisplayName)
	request.CrossConnectGroupId = crossConnectGroup.Id

	resp, err := c.ListCrossConnects(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.Items
}

func listCrossConnectLocations(t *testing.T) []core.CrossConnectLocation {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := core.ListCrossConnectLocationsRequest{}
	request.CompartmentId = common.String(getCompartmentID())

	resp, err := c.ListCrossConnectLocations(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.Items
}

func getCrossConnectStatus(t *testing.T) core.CrossConnectStatus {
	crossConnect := createOrGetCrossConnect(t)
	assert.NotEmpty(t, crossConnect)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := core.GetCrossConnectStatusRequest{}
	request.CrossConnectId = crossConnect.Id

	resp, err := c.GetCrossConnectStatus(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.CrossConnectStatus
}

func listCrossConnectPortSpeedShapes(t *testing.T) []core.CrossConnectPortSpeedShape {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := core.ListCrossconnectPortSpeedShapesRequest{}
	request.CompartmentId = common.String(getCompartmentID())

	resp, err := c.ListCrossconnectPortSpeedShapes(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.Items
}

func getCrossConnectLetterOfAuthority(t *testing.T) core.LetterOfAuthority {
	crossConnect := createOrGetCrossConnect(t)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := core.GetCrossConnectLetterOfAuthorityRequest{}
	request.CrossConnectId = crossConnect.Id

	resp, err := c.GetCrossConnectLetterOfAuthority(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.LetterOfAuthority
}

func getFastConnectProviderServices(t *testing.T) core.FastConnectProviderService {
	prividerServices := listFastConnectProviderServices(t)
	assert.NotEqual(t, len(prividerServices), 0)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := core.GetFastConnectProviderServiceRequest{}
	request.ProviderServiceId = prividerServices[0].Id

	resp, err := c.GetFastConnectProviderService(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.FastConnectProviderService
}

func listFastConnectProviderServices(t *testing.T) []core.FastConnectProviderService {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := core.ListFastConnectProviderServicesRequest{}
	request.CompartmentId = common.String(getCompartmentID())

	resp, err := c.ListFastConnectProviderServices(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.Items)
	return resp.Items
}

func listFastConnectProviderVirtualCircuitBandwidthShapes(t *testing.T) []core.VirtualCircuitBandwidthShape {
	providerService := getFastConnectProviderServices(t)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := core.ListFastConnectProviderVirtualCircuitBandwidthShapesRequest{}
	request.ProviderServiceId = providerService.Id

	resp, err := c.ListFastConnectProviderVirtualCircuitBandwidthShapes(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.Items)
	return resp.Items
}

func listVirtualCircuitBandwidthShapes(t *testing.T) []core.VirtualCircuitBandwidthShape {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := core.ListVirtualCircuitBandwidthShapesRequest{}
	request.CompartmentId = common.String(getCompartmentID())

	resp, err := c.ListVirtualCircuitBandwidthShapes(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.Items)
	return resp.Items
}

func createOrGetInstanceConsoleConnection(t *testing.T) core.InstanceConsoleConnection {
	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	listRequest := core.ListInstanceConsoleConnectionsRequest{}
	listRequest.CompartmentId = common.String(getCompartmentID())

	listResp, err := c.ListInstanceConsoleConnections(context.Background(), listRequest)
	failIfError(t, err)
	assert.NotEmpty(t, listResp)

	if listResp.Items != nil && len(listResp.Items) != 0 {
		return listResp.Items[0]
	}

	// create a new one
	instance := createOrGetInstance(t)
	createRequest := core.CreateInstanceConsoleConnectionRequest{}
	createRequest.InstanceId = instance.Id

	// get the public key
	buffer, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa.pub")
	failIfError(t, err)
	createRequest.PublicKey = common.String(string(buffer))

	createResp, err := c.CreateInstanceConsoleConnection(context.Background(), createRequest)
	failIfError(t, err)

	assert.NotEmpty(t, createResp)
	return createResp.InstanceConsoleConnection
}

func getConsoleHistory(t *testing.T, historyID *string) core.ConsoleHistory {
	assert.NotEmpty(t, historyID)

	// create a new console history
	request := core.GetConsoleHistoryRequest{
		InstanceConsoleHistoryId: historyID,
	}

	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	resp, err := c.GetConsoleHistory(context.Background(), request)
	failIfError(t, err)

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.ConsoleHistory)
	return resp.ConsoleHistory
}

func captureOrGetConsoleHistory(t *testing.T) core.ConsoleHistory {
	consoleHistories := listConsoleHistories(t)

	for _, element := range consoleHistories {
		assert.NotEmpty(t, element)
		if *element.DisplayName == consoleHistoryDisplayName {
			// find it, return
			return element
		}
	}

	// create a new console history
	instance := createOrGetInstance(t)
	request := core.CaptureConsoleHistoryRequest{}
	request.InstanceId = instance.Id
	request.DisplayName = common.String(consoleHistoryDisplayName)

	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	resp, err := c.CaptureConsoleHistory(context.Background(), request)
	failIfError(t, err)

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.ConsoleHistory)
	return resp.ConsoleHistory
}

func listConsoleHistories(t *testing.T) []core.ConsoleHistory {
	instance := createOrGetInstance(t)

	request := core.ListConsoleHistoriesRequest{
		CompartmentId:      common.String(getCompartmentID()),
		AvailabilityDomain: common.String(validAD()),
		InstanceId:         instance.Id,
	}

	c, clerr := core.NewComputeClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	resp, err := c.ListConsoleHistories(context.Background(), request)
	failIfError(t, err)

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.Items)
	return resp.Items
}

func createOrGetLocalPeeringGateway(t *testing.T) core.LocalPeeringGateway {
	gateways := listLocalPeeringGateways(t)

	for _, element := range gateways {
		assert.NotEmpty(t, element)
		if *element.DisplayName == peeringGatewaysDisplayName {
			return element
		}
	}

	// create a new one
	vcn := createOrGetVcn(t)
	assert.NotEmpty(t, vcn)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := core.CreateLocalPeeringGatewayRequest{}
	request.CompartmentId = common.String(getCompartmentID())
	request.DisplayName = common.String(peeringGatewaysDisplayName)
	request.VcnId = vcn.Id

	resp, err := c.CreateLocalPeeringGateway(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.LocalPeeringGateway
}

func listLocalPeeringGateways(t *testing.T) []core.LocalPeeringGateway {
	vnc := createOrGetVcn(t)
	assert.NotEmpty(t, vnc)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := core.ListLocalPeeringGatewaysRequest{
		CompartmentId: common.String(getCompartmentID()),
		VcnId:         vnc.Id,
	}

	resp, err := c.ListLocalPeeringGateways(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.Items
}

func createOrGetVirtualCircuit(t *testing.T) core.VirtualCircuit {
	virtualCircuits := listVirtualCircuits(t)

	for _, element := range virtualCircuits {
		assert.NotEmpty(t, element)
		if *element.DisplayName == virtualCircuitDisplayName {
			return element
		}
	}

	// create a new one
	request := core.CreateVirtualCircuitRequest{}
	request.CompartmentId = common.String(getCompartmentID())
	request.DisplayName = common.String(virtualCircuitDisplayName)
	request.Type = "PRIVATE"

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	resp, err := c.CreateVirtualCircuit(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.VirtualCircuit
}

func listVirtualCircuits(t *testing.T) []core.VirtualCircuit {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := core.ListVirtualCircuitsRequest{
		CompartmentId: common.String(getCompartmentID()),
		DisplayName:   common.String(virtualCircuitDisplayName),
	}

	resp, err := c.ListVirtualCircuits(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
	return resp.Items
}

func createLoadBalancerIfNotExist(t *testing.T) *loadbalancer.LoadBalancer {
	loadbalancers := listLoadBalancers(t, loadbalancer.LoadBalancerLifecycleStateActive)

	for _, element := range loadbalancers {
		assert.NotEmpty(t, element)
		if *element.DisplayName == loadbalancerDisplayName {
			// found it, return
			return &element
		}
	}

	// create new load balancer
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.CreateLoadBalancerRequest{}
	request.CompartmentId = common.String(getCompartmentID())
	request.DisplayName = common.String(loadbalancerDisplayName)

	subnet := createOrGetSubnet(t)
	request.SubnetIds = []string{*subnet.Id}

	shapes := listLoadBalancerShapes(t)
	request.ShapeName = shapes[0].Name

	err := c.CreateLoadBalancer(context.Background(), request)
	assert.NoError(t, err)

	loadbalancers = listLoadBalancers(t, "")
	for _, element := range loadbalancers {
		assert.NotEmpty(t, element)
		if *element.DisplayName == loadbalancerDisplayName {
			// found it, return
			// TODO: wati for state change to Active and return
			return &element
		}
	}

	t.Error("create loadbalancer failed")
	t.Fail()
	return nil
}

func listLoadBalancers(t *testing.T, lifecycleState loadbalancer.LoadBalancerLifecycleStateEnum) []loadbalancer.LoadBalancer {
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	request := loadbalancer.ListLoadBalancersRequest{
		CompartmentId: common.String(getCompartmentID()),
		DisplayName:   common.String(loadbalancerDisplayName),
	}

	if lifecycleState != "" {
		request.LifecycleState = lifecycleState
	}

	r, err := c.ListLoadBalancers(context.Background(), request)
	assert.NoError(t, err)
	assert.NotEmpty(t, r)
	return r.Items
}

func listLoadBalancerShapes(t *testing.T) []loadbalancer.LoadBalancerShape {
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.ListShapesRequest{
		CompartmentId: common.String(getCompartmentID()),
	}

	r, err := c.ListShapes(context.Background(), request)
	assert.NoError(t, err)
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, r.Items)
	assert.NotEqual(t, len(r.Items), 0)
	return r.Items
}
