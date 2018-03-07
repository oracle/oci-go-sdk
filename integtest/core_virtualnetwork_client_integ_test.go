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
	"testing"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/stretchr/testify/assert"
)

var (
	testRegionForVirtualNetwork = common.RegionPHX
)

func TestVirtualNetworkClient_CreateCpe(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreateCpeRequest{}
	r, err := c.CreateCpe(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateCrossConnect(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreateCrossConnectRequest{}
	r, err := c.CreateCrossConnect(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateCrossConnectGroup(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreateCrossConnectGroupRequest{}
	r, err := c.CreateCrossConnectGroup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateDhcpOptions(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreateDhcpOptionsRequest{}
	r, err := c.CreateDhcpOptions(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateDrg(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreateDrgRequest{}
	r, err := c.CreateDrg(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateDrgAttachment(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreateDrgAttachmentRequest{}
	r, err := c.CreateDrgAttachment(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateIPSecConnection(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreateIPSecConnectionRequest{}
	r, err := c.CreateIPSecConnection(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateInternetGateway(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreateInternetGatewayRequest{}
	r, err := c.CreateInternetGateway(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreatePrivateIp(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreatePrivateIpRequest{}
	r, err := c.CreatePrivateIp(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateRouteTable(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreateRouteTableRequest{}
	r, err := c.CreateRouteTable(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateSecurityList(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreateSecurityListRequest{}
	r, err := c.CreateSecurityList(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateSubnet(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreateSubnetRequest{}
	r, err := c.CreateSubnet(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateVcn(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreateVcnRequest{}
	r, err := c.CreateVcn(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_CreateVirtualCircuit(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.CreateVirtualCircuitRequest{}
	r, err := c.CreateVirtualCircuit(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteCpe(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeleteCpeRequest{}
	_, err := c.DeleteCpe(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteCrossConnect(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeleteCrossConnectRequest{}
	_, err := c.DeleteCrossConnect(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteCrossConnectGroup(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeleteCrossConnectGroupRequest{}
	_, err := c.DeleteCrossConnectGroup(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteDhcpOptions(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeleteDhcpOptionsRequest{}
	_, err := c.DeleteDhcpOptions(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteDrg(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeleteDrgRequest{}
	_, err := c.DeleteDrg(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteDrgAttachment(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeleteDrgAttachmentRequest{}
	_, err := c.DeleteDrgAttachment(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteIPSecConnection(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeleteIPSecConnectionRequest{}
	_, err := c.DeleteIPSecConnection(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteInternetGateway(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeleteInternetGatewayRequest{}
	_, err := c.DeleteInternetGateway(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeletePrivateIp(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeletePrivateIpRequest{}
	_, err := c.DeletePrivateIp(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteRouteTable(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeleteRouteTableRequest{}
	_, err := c.DeleteRouteTable(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteSecurityList(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeleteSecurityListRequest{}
	_, err := c.DeleteSecurityList(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteSubnet(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeleteSubnetRequest{}
	_, err := c.DeleteSubnet(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteVcn(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeleteVcnRequest{}
	_, err := c.DeleteVcn(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_DeleteVirtualCircuit(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.DeleteVirtualCircuitRequest{}
	_, err := c.DeleteVirtualCircuit(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetCpe(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetCpeRequest{}
	r, err := c.GetCpe(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetCrossConnect(t *testing.T) {
	crossConnect := createOrGetCrossConnect(t)
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetCrossConnectRequest{}
	request.CrossConnectId = crossConnect.Id
	r, err := c.GetCrossConnect(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetCrossConnectGroup(t *testing.T) {
	connectGroup := createOrGetCrossConnectGroup(t)
	assert.NotEmpty(t, connectGroup)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetCrossConnectGroupRequest{
		CrossConnectGroupId: connectGroup.Id,
	}

	r, err := c.GetCrossConnectGroup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetCrossConnectLetterOfAuthority(t *testing.T) {
	r := getCrossConnectLetterOfAuthority(t)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	return
}

func TestVirtualNetworkClient_GetCrossConnectStatus(t *testing.T) {
	crossConnectStatus := getCrossConnectStatus(t)
	assert.NotEmpty(t, crossConnectStatus, fmt.Sprint(crossConnectStatus))
	assert.NotEmpty(t, crossConnectStatus.CrossConnectId, fmt.Sprint(crossConnectStatus.CrossConnectId))
	return
}

func TestVirtualNetworkClient_GetDhcpOptions(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetDhcpOptionsRequest{}
	r, err := c.GetDhcpOptions(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetDrg(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetDrgRequest{}
	r, err := c.GetDrg(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetDrgAttachment(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetDrgAttachmentRequest{}
	r, err := c.GetDrgAttachment(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetFastConnectProviderService(t *testing.T) {
	r := getFastConnectProviderServices(t)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	return
}

func TestVirtualNetworkClient_GetIPSecConnection(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetIPSecConnectionRequest{}
	r, err := c.GetIPSecConnection(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetIPSecConnectionDeviceConfig(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetIPSecConnectionDeviceConfigRequest{}
	r, err := c.GetIPSecConnectionDeviceConfig(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetIPSecConnectionDeviceStatus(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetIPSecConnectionDeviceStatusRequest{}
	r, err := c.GetIPSecConnectionDeviceStatus(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetInternetGateway(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetInternetGatewayRequest{}
	r, err := c.GetInternetGateway(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetPrivateIp(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetPrivateIpRequest{}
	r, err := c.GetPrivateIp(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetRouteTable(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetRouteTableRequest{}
	r, err := c.GetRouteTable(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetSecurityList(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetSecurityListRequest{}
	r, err := c.GetSecurityList(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetSubnet(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetSubnetRequest{}
	r, err := c.GetSubnet(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetVcn(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetVcnRequest{}
	r, err := c.GetVcn(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetVirtualCircuit(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetVirtualCircuitRequest{}
	r, err := c.GetVirtualCircuit(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetVnic(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetVnicRequest{}
	r, err := c.GetVnic(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListCpes(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListCpesRequest{}
	r, err := c.ListCpes(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListCrossConnectGroups(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListCrossConnectGroupsRequest{}
	r, err := c.ListCrossConnectGroups(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListCrossConnectLocations(t *testing.T) {
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListCrossConnectLocationsRequest{}
	request.CompartmentId = common.String(getCompartmentID())
	r, err := c.ListCrossConnectLocations(context.Background(), request)
	failIfError(t, err)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	return
}

func TestVirtualNetworkClient_ListCrossConnects(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListCrossConnectsRequest{}
	r, err := c.ListCrossConnects(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListCrossconnectPortSpeedShapes(t *testing.T) {
	shapes := listCrossConnectPortSpeedShapes(t)
	assert.NotEmpty(t, shapes, fmt.Sprint(shapes))
	return
}

func TestVirtualNetworkClient_ListDhcpOptions(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListDhcpOptionsRequest{}
	r, err := c.ListDhcpOptions(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListDrgAttachments(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListDrgAttachmentsRequest{}
	r, err := c.ListDrgAttachments(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListDrgs(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListDrgsRequest{}
	r, err := c.ListDrgs(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListFastConnectProviderServices(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListFastConnectProviderServicesRequest{}
	r, err := c.ListFastConnectProviderServices(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListFastConnectProviderVirtualCircuitBandwidthShapes(t *testing.T) {
	r := listFastConnectProviderVirtualCircuitBandwidthShapes(t)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	return
}

func TestVirtualNetworkClient_ListIPSecConnections(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListIPSecConnectionsRequest{}
	r, err := c.ListIPSecConnections(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListInternetGateways(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListInternetGatewaysRequest{}
	r, err := c.ListInternetGateways(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListPrivateIps(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListPrivateIpsRequest{}
	r, err := c.ListPrivateIps(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListRouteTables(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListRouteTablesRequest{}
	r, err := c.ListRouteTables(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListSecurityLists(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListSecurityListsRequest{}
	r, err := c.ListSecurityLists(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListSubnets(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListSubnetsRequest{}
	r, err := c.ListSubnets(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListVcns(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.ListVcnsRequest{}
	r, err := c.ListVcns(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_ListVirtualCircuitBandwidthShapes(t *testing.T) {
	r := listVirtualCircuitBandwidthShapes(t)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	return
}

func TestVirtualNetworkClient_ListVirtualCircuits(t *testing.T) {
	t.Skip("Service return 500, service api bug")
	// make sure the list operation at least return one item
	createOrGetVirtualCircuit(t)
	r := listVirtualCircuits(t)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NotEqual(t, len(r), 0)
	return
}

func TestVirtualNetworkClient_UpdateCpe(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateCpeRequest{}
	r, err := c.UpdateCpe(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateCrossConnect(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateCrossConnectRequest{}
	r, err := c.UpdateCrossConnect(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateCrossConnectGroup(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateCrossConnectGroupRequest{}
	r, err := c.UpdateCrossConnectGroup(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateDhcpOptions(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateDhcpOptionsRequest{}
	r, err := c.UpdateDhcpOptions(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateDrg(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateDrgRequest{}
	r, err := c.UpdateDrg(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateDrgAttachment(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateDrgAttachmentRequest{}
	r, err := c.UpdateDrgAttachment(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateIPSecConnection(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateIPSecConnectionRequest{}
	r, err := c.UpdateIPSecConnection(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateInternetGateway(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateInternetGatewayRequest{}
	r, err := c.UpdateInternetGateway(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdatePrivateIp(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdatePrivateIpRequest{}
	r, err := c.UpdatePrivateIp(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateRouteTable(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateRouteTableRequest{}
	r, err := c.UpdateRouteTable(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateSecurityList(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateSecurityListRequest{}
	r, err := c.UpdateSecurityList(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateSubnet(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateSubnetRequest{}
	r, err := c.UpdateSubnet(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateVcn(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateVcnRequest{}
	r, err := c.UpdateVcn(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateVirtualCircuit(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateVirtualCircuitRequest{}
	r, err := c.UpdateVirtualCircuit(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_UpdateVnic(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.UpdateVnicRequest{}
	r, err := c.UpdateVnic(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestVirtualNetworkClient_GetLocalPeeringGateway(t *testing.T) {
	t.Skip("Not in current version of service spec, update later")
	/*createOrGetSubnet(t)
	gateway := createOrGetLocalPeeringGateway(t)
	assert.NotEmpty(t, gateway)

	c, clerr := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := core.GetLocalPeeringGatewayRequest{
		LocalPeeringGatewayId: gateway.Id,
	}

	r, err := c.GetLocalPeeringGateway(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)*/
	return
}
