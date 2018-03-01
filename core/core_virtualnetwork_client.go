// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//VirtualNetworkClient a client for VirtualNetwork
type VirtualNetworkClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewVirtualNetworkClientWithConfigurationProvider Creates a new default VirtualNetwork client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewVirtualNetworkClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client VirtualNetworkClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = VirtualNetworkClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *VirtualNetworkClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "iaas", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *VirtualNetworkClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.config = &configProvider
	client.SetRegion(region)
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *VirtualNetworkClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// BulkAddVirtualCircuitPublicPrefixes Adds one or more customer public IP prefixes to the specified public virtual circuit.
// Use this operation (and not UpdateVirtualCircuit)
// to add prefixes to the virtual circuit. Oracle must verify the customer's ownership
// of each prefix before traffic for that prefix will flow across the virtual circuit.
func (client VirtualNetworkClient) BulkAddVirtualCircuitPublicPrefixes(ctx context.Context, request BulkAddVirtualCircuitPublicPrefixesRequest) (response BulkAddVirtualCircuitPublicPrefixesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.bulkAddVirtualCircuitPublicPrefixes, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(BulkAddVirtualCircuitPublicPrefixesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) bulkAddVirtualCircuitPublicPrefixes(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/virtualCircuits/{virtualCircuitId}/actions/bulkAddPublicPrefixes")
	if err != nil {
		return nil, err
	}

	var response BulkAddVirtualCircuitPublicPrefixesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// BulkDeleteVirtualCircuitPublicPrefixes Removes one or more customer public IP prefixes from the specified public virtual circuit.
// Use this operation (and not UpdateVirtualCircuit)
// to remove prefixes from the virtual circuit. When the virtual circuit's state switches
// back to PROVISIONED, Oracle stops advertising the specified prefixes across the connection.
func (client VirtualNetworkClient) BulkDeleteVirtualCircuitPublicPrefixes(ctx context.Context, request BulkDeleteVirtualCircuitPublicPrefixesRequest) (response BulkDeleteVirtualCircuitPublicPrefixesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.bulkDeleteVirtualCircuitPublicPrefixes, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(BulkDeleteVirtualCircuitPublicPrefixesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) bulkDeleteVirtualCircuitPublicPrefixes(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/virtualCircuits/{virtualCircuitId}/actions/bulkDeletePublicPrefixes")
	if err != nil {
		return nil, err
	}

	var response BulkDeleteVirtualCircuitPublicPrefixesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ConnectLocalPeeringGateways Connects this local peering gateway (LPG) to another one in the same region.
// This operation must be called by the VCN administrator who is designated as
// the *requestor* in the peering relationship. The *acceptor* must implement
// an Identity and Access Management (IAM) policy that gives the requestor permission
// to connect to LPGs in the acceptor's compartment. Without that permission, this
// operation will fail. For more information, see
// [VCN Peering]({{DOC_SERVER_URL}}/Content/Network/Tasks/VCNpeering.htm).
func (client VirtualNetworkClient) ConnectLocalPeeringGateways(ctx context.Context, request ConnectLocalPeeringGatewaysRequest) (response ConnectLocalPeeringGatewaysResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.connectLocalPeeringGateways, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ConnectLocalPeeringGatewaysResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) connectLocalPeeringGateways(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/localPeeringGateways/{localPeeringGatewayId}/actions/connect")
	if err != nil {
		return nil, err
	}

	var response ConnectLocalPeeringGatewaysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateCpe Creates a new virtual Customer-Premises Equipment (CPE) object in the specified compartment. For
// more information, see [IPSec VPNs]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingIPsec.htm).
// For the purposes of access control, you must provide the OCID of the compartment where you want
// the CPE to reside. Notice that the CPE doesn't have to be in the same compartment as the IPSec
// connection or other Networking Service components. If you're not sure which compartment to
// use, put the CPE in the same compartment as the DRG. For more information about
// compartments and access control, see [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm).
// For information about OCIDs, see [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You must provide the public IP address of your on-premises router. See
// [Configuring Your On-Premises Router for an IPSec VPN]({{DOC_SERVER_URL}}/Content/Network/Tasks/configuringCPE.htm).
// You may optionally specify a *display name* for the CPE, otherwise a default is provided. It does not have to
// be unique, and you can change it. Avoid entering confidential information.
func (client VirtualNetworkClient) CreateCpe(ctx context.Context, request CreateCpeRequest) (response CreateCpeResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createCpe, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateCpeResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) createCpe(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/cpes")
	if err != nil {
		return nil, err
	}

	var response CreateCpeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateCrossConnect Creates a new cross-connect. Oracle recommends you create each cross-connect in a
// CrossConnectGroup so you can use link aggregation
// with the connection.
// After creating the `CrossConnect` object, you need to go the FastConnect location
// and request to have the physical cable installed. For more information, see
// [FastConnect Overview]({{DOC_SERVER_URL}}/Content/Network/Concepts/fastconnect.htm).
// For the purposes of access control, you must provide the OCID of the
// compartment where you want the cross-connect to reside. If you're
// not sure which compartment to use, put the cross-connect in the
// same compartment with your VCN. For more information about
// compartments and access control, see
// [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm).
// For information about OCIDs, see
// [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the cross-connect.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
func (client VirtualNetworkClient) CreateCrossConnect(ctx context.Context, request CreateCrossConnectRequest) (response CreateCrossConnectResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createCrossConnect, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateCrossConnectResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) createCrossConnect(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/crossConnects")
	if err != nil {
		return nil, err
	}

	var response CreateCrossConnectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateCrossConnectGroup Creates a new cross-connect group to use with Oracle Cloud Infrastructure
// FastConnect. For more information, see
// [FastConnect Overview]({{DOC_SERVER_URL}}/Content/Network/Concepts/fastconnect.htm).
// For the purposes of access control, you must provide the OCID of the
// compartment where you want the cross-connect group to reside. If you're
// not sure which compartment to use, put the cross-connect group in the
// same compartment with your VCN. For more information about
// compartments and access control, see
// [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm).
// For information about OCIDs, see
// [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the cross-connect group.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
func (client VirtualNetworkClient) CreateCrossConnectGroup(ctx context.Context, request CreateCrossConnectGroupRequest) (response CreateCrossConnectGroupResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createCrossConnectGroup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateCrossConnectGroupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) createCrossConnectGroup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/crossConnectGroups")
	if err != nil {
		return nil, err
	}

	var response CreateCrossConnectGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateDhcpOptions Creates a new set of DHCP options for the specified VCN. For more information, see
// DhcpOptions.
// For the purposes of access control, you must provide the OCID of the compartment where you want the set of
// DHCP options to reside. Notice that the set of options doesn't have to be in the same compartment as the VCN,
// subnets, or other Networking Service components. If you're not sure which compartment to use, put the set
// of DHCP options in the same compartment as the VCN. For more information about compartments and access control, see
// [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
// [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the set of DHCP options, otherwise a default is provided.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
func (client VirtualNetworkClient) CreateDhcpOptions(ctx context.Context, request CreateDhcpOptionsRequest) (response CreateDhcpOptionsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createDhcpOptions, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateDhcpOptionsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) createDhcpOptions(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/dhcps")
	if err != nil {
		return nil, err
	}

	var response CreateDhcpOptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateDrg Creates a new Dynamic Routing Gateway (DRG) in the specified compartment. For more information,
// see [Dynamic Routing Gateways (DRGs)]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingDRGs.htm).
// For the purposes of access control, you must provide the OCID of the compartment where you want
// the DRG to reside. Notice that the DRG doesn't have to be in the same compartment as the VCN,
// the DRG attachment, or other Networking Service components. If you're not sure which compartment
// to use, put the DRG in the same compartment as the VCN. For more information about compartments
// and access control, see [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm).
// For information about OCIDs, see [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the DRG, otherwise a default is provided.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
func (client VirtualNetworkClient) CreateDrg(ctx context.Context, request CreateDrgRequest) (response CreateDrgResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createDrg, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateDrgResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) createDrg(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/drgs")
	if err != nil {
		return nil, err
	}

	var response CreateDrgResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateDrgAttachment Attaches the specified DRG to the specified VCN. A VCN can be attached to only one DRG at a time,
// and vice versa. The response includes a `DrgAttachment` object with its own OCID. For more
// information about DRGs, see
// [Dynamic Routing Gateways (DRGs)]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingDRGs.htm).
// You may optionally specify a *display name* for the attachment, otherwise a default is provided.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
// For the purposes of access control, the DRG attachment is automatically placed into the same compartment
// as the VCN. For more information about compartments and access control, see
// [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm).
func (client VirtualNetworkClient) CreateDrgAttachment(ctx context.Context, request CreateDrgAttachmentRequest) (response CreateDrgAttachmentResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createDrgAttachment, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateDrgAttachmentResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) createDrgAttachment(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/drgAttachments")
	if err != nil {
		return nil, err
	}

	var response CreateDrgAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateIPSecConnection Creates a new IPSec connection between the specified DRG and CPE. For more information, see
// [IPSec VPNs]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingIPsec.htm).
// In the request, you must include at least one static route to the CPE object (you're allowed a maximum
// of 10). For example: 10.0.8.0/16.
// For the purposes of access control, you must provide the OCID of the compartment where you want the
// IPSec connection to reside. Notice that the IPSec connection doesn't have to be in the same compartment
// as the DRG, CPE, or other Networking Service components. If you're not sure which compartment to
// use, put the IPSec connection in the same compartment as the DRG. For more information about
// compartments and access control, see
// [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm).
// For information about OCIDs, see [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the IPSec connection, otherwise a default is provided.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
// After creating the IPSec connection, you need to configure your on-premises router
// with tunnel-specific information returned by
// GetIPSecConnectionDeviceConfig.
// For each tunnel, that operation gives you the IP address of Oracle's VPN headend and the shared secret
// (that is, the pre-shared key). For more information, see
// [Configuring Your On-Premises Router for an IPSec VPN]({{DOC_SERVER_URL}}/Content/Network/Tasks/configuringCPE.htm).
// To get the status of the tunnels (whether they're up or down), use
// GetIPSecConnectionDeviceStatus.
func (client VirtualNetworkClient) CreateIPSecConnection(ctx context.Context, request CreateIPSecConnectionRequest) (response CreateIPSecConnectionResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createIPSecConnection, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateIPSecConnectionResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) createIPSecConnection(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/ipsecConnections")
	if err != nil {
		return nil, err
	}

	var response CreateIPSecConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateInternetGateway Creates a new Internet Gateway for the specified VCN. For more information, see
// [Connectivity to the Internet]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingIGs.htm).
// For the purposes of access control, you must provide the OCID of the compartment where you want the Internet
// Gateway to reside. Notice that the Internet Gateway doesn't have to be in the same compartment as the VCN or
// other Networking Service components. If you're not sure which compartment to use, put the Internet
// Gateway in the same compartment with the VCN. For more information about compartments and access control, see
// [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
// [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the Internet Gateway, otherwise a default is provided. It
// does not have to be unique, and you can change it. Avoid entering confidential information.
// For traffic to flow between a subnet and an Internet Gateway, you must create a route rule accordingly in
// the subnet's route table (for example, 0.0.0.0/0 > Internet Gateway). See
// UpdateRouteTable.
// You must specify whether the Internet Gateway is enabled when you create it. If it's disabled, that means no
// traffic will flow to/from the internet even if there's a route rule that enables that traffic. You can later
// use UpdateInternetGateway to easily disable/enable
// the gateway without changing the route rule.
func (client VirtualNetworkClient) CreateInternetGateway(ctx context.Context, request CreateInternetGatewayRequest) (response CreateInternetGatewayResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createInternetGateway, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateInternetGatewayResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) createInternetGateway(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/internetGateways")
	if err != nil {
		return nil, err
	}

	var response CreateInternetGatewayResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateLocalPeeringGateway Creates a new local peering gateway (LPG) for the specified VCN.
func (client VirtualNetworkClient) CreateLocalPeeringGateway(ctx context.Context, request CreateLocalPeeringGatewayRequest) (response CreateLocalPeeringGatewayResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createLocalPeeringGateway, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateLocalPeeringGatewayResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) createLocalPeeringGateway(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/localPeeringGateways")
	if err != nil {
		return nil, err
	}

	var response CreateLocalPeeringGatewayResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreatePrivateIp Creates a secondary private IP for the specified VNIC.
// For more information about secondary private IPs, see
// [IP Addresses]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingIPaddresses.htm).
func (client VirtualNetworkClient) CreatePrivateIp(ctx context.Context, request CreatePrivateIpRequest) (response CreatePrivateIpResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createPrivateIp, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreatePrivateIpResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) createPrivateIp(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/privateIps")
	if err != nil {
		return nil, err
	}

	var response CreatePrivateIpResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateRouteTable Creates a new route table for the specified VCN. In the request you must also include at least one route
// rule for the new route table. For information on the number of rules you can have in a route table, see
// [Service Limits]({{DOC_SERVER_URL}}/Content/General/Concepts/servicelimits.htm). For general information about route
// tables in your VCN and the types of targets you can use in route rules,
// see [Route Tables]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingroutetables.htm).
// For the purposes of access control, you must provide the OCID of the compartment where you want the route
// table to reside. Notice that the route table doesn't have to be in the same compartment as the VCN, subnets,
// or other Networking Service components. If you're not sure which compartment to use, put the route
// table in the same compartment as the VCN. For more information about compartments and access control, see
// [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
// [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the route table, otherwise a default is provided.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
func (client VirtualNetworkClient) CreateRouteTable(ctx context.Context, request CreateRouteTableRequest) (response CreateRouteTableResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createRouteTable, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateRouteTableResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) createRouteTable(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/routeTables")
	if err != nil {
		return nil, err
	}

	var response CreateRouteTableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateSecurityList Creates a new security list for the specified VCN. For more information
// about security lists, see [Security Lists]({{DOC_SERVER_URL}}/Content/Network/Concepts/securitylists.htm).
// For information on the number of rules you can have in a security list, see
// [Service Limits]({{DOC_SERVER_URL}}/Content/General/Concepts/servicelimits.htm).
// For the purposes of access control, you must provide the OCID of the compartment where you want the security
// list to reside. Notice that the security list doesn't have to be in the same compartment as the VCN, subnets,
// or other Networking Service components. If you're not sure which compartment to use, put the security
// list in the same compartment as the VCN. For more information about compartments and access control, see
// [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
// [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the security list, otherwise a default is provided.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
func (client VirtualNetworkClient) CreateSecurityList(ctx context.Context, request CreateSecurityListRequest) (response CreateSecurityListResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createSecurityList, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateSecurityListResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) createSecurityList(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/securityLists")
	if err != nil {
		return nil, err
	}

	var response CreateSecurityListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateSubnet Creates a new subnet in the specified VCN. You can't change the size of the subnet after creation,
// so it's important to think about the size of subnets you need before creating them.
// For more information, see [VCNs and Subnets]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingVCNs.htm).
// For information on the number of subnets you can have in a VCN, see
// [Service Limits]({{DOC_SERVER_URL}}/Content/General/Concepts/servicelimits.htm).
// For the purposes of access control, you must provide the OCID of the compartment where you want the subnet
// to reside. Notice that the subnet doesn't have to be in the same compartment as the VCN, route tables, or
// other Networking Service components. If you're not sure which compartment to use, put the subnet in
// the same compartment as the VCN. For more information about compartments and access control, see
// [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm). For information about OCIDs,
// see [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You may optionally associate a route table with the subnet. If you don't, the subnet will use the
// VCN's default route table. For more information about route tables, see
// [Route Tables]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingroutetables.htm).
// You may optionally associate a security list with the subnet. If you don't, the subnet will use the
// VCN's default security list. For more information about security lists, see
// [Security Lists]({{DOC_SERVER_URL}}/Content/Network/Concepts/securitylists.htm).
// You may optionally associate a set of DHCP options with the subnet. If you don't, the subnet will use the
// VCN's default set. For more information about DHCP options, see
// [DHCP Options]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingDHCP.htm).
// You may optionally specify a *display name* for the subnet, otherwise a default is provided.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
// You can also add a DNS label for the subnet, which is required if you want the Internet and
// VCN Resolver to resolve hostnames for instances in the subnet. For more information, see
// [DNS in Your Virtual Cloud Network]({{DOC_SERVER_URL}}/Content/Network/Concepts/dns.htm).
func (client VirtualNetworkClient) CreateSubnet(ctx context.Context, request CreateSubnetRequest) (response CreateSubnetResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createSubnet, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateSubnetResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) createSubnet(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/subnets")
	if err != nil {
		return nil, err
	}

	var response CreateSubnetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateVcn Creates a new Virtual Cloud Network (VCN). For more information, see
// [VCNs and Subnets]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingVCNs.htm).
// For the VCN you must specify a single, contiguous IPv4 CIDR block. Oracle recommends using one of the
// private IP address ranges specified in [RFC 1918](https://tools.ietf.org/html/rfc1918) (10.0.0.0/8,
// 172.16/12, and 192.168/16). Example: 172.16.0.0/16. The CIDR block can range from /16 to /30, and it
// must not overlap with your on-premises network. You can't change the size of the VCN after creation.
// For the purposes of access control, you must provide the OCID of the compartment where you want the VCN to
// reside. Consult an Oracle Cloud Infrastructure administrator in your organization if you're not sure which
// compartment to use. Notice that the VCN doesn't have to be in the same compartment as the subnets or other
// Networking Service components. For more information about compartments and access control, see
// [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
// [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the VCN, otherwise a default is provided. It does not have to
// be unique, and you can change it. Avoid entering confidential information.
// You can also add a DNS label for the VCN, which is required if you want the instances to use the
// Interent and VCN Resolver option for DNS in the VCN. For more information, see
// [DNS in Your Virtual Cloud Network]({{DOC_SERVER_URL}}/Content/Network/Concepts/dns.htm).
// The VCN automatically comes with a default route table, default security list, and default set of DHCP options.
// The OCID for each is returned in the response. You can't delete these default objects, but you can change their
// contents (that is, change the route rules, security list rules, and so on).
// The VCN and subnets you create are not accessible until you attach an Internet Gateway or set up an IPSec VPN
// or FastConnect. For more information, see
// [Overview of the Networking Service]({{DOC_SERVER_URL}}/Content/Network/Concepts/overview.htm).
func (client VirtualNetworkClient) CreateVcn(ctx context.Context, request CreateVcnRequest) (response CreateVcnResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createVcn, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateVcnResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) createVcn(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/vcns")
	if err != nil {
		return nil, err
	}

	var response CreateVcnResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateVirtualCircuit Creates a new virtual circuit to use with Oracle Cloud
// Infrastructure FastConnect. For more information, see
// [FastConnect Overview]({{DOC_SERVER_URL}}/Content/Network/Concepts/fastconnect.htm).
// For the purposes of access control, you must provide the OCID of the
// compartment where you want the virtual circuit to reside. If you're
// not sure which compartment to use, put the virtual circuit in the
// same compartment with the DRG it's using. For more information about
// compartments and access control, see
// [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm).
// For information about OCIDs, see
// [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You may optionally specify a *display name* for the virtual circuit.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
// **Important:** When creating a virtual circuit, you specify a DRG for
// the traffic to flow through. Make sure you attach the DRG to your
// VCN and confirm the VCN's routing sends traffic to the DRG. Otherwise
// traffic will not flow. For more information, see
// [Route Tables]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingroutetables.htm).
func (client VirtualNetworkClient) CreateVirtualCircuit(ctx context.Context, request CreateVirtualCircuitRequest) (response CreateVirtualCircuitResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createVirtualCircuit, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateVirtualCircuitResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) createVirtualCircuit(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/virtualCircuits")
	if err != nil {
		return nil, err
	}

	var response CreateVirtualCircuitResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteCpe Deletes the specified CPE object. The CPE must not be connected to a DRG. This is an asynchronous
// operation. The CPE's `lifecycleState` will change to TERMINATING temporarily until the CPE is completely
// removed.
func (client VirtualNetworkClient) DeleteCpe(ctx context.Context, request DeleteCpeRequest) (response DeleteCpeResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteCpe, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteCpeResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) deleteCpe(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/cpes/{cpeId}")
	if err != nil {
		return nil, err
	}

	var response DeleteCpeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteCrossConnect Deletes the specified cross-connect. It must not be mapped to a
// VirtualCircuit.
func (client VirtualNetworkClient) DeleteCrossConnect(ctx context.Context, request DeleteCrossConnectRequest) (response DeleteCrossConnectResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteCrossConnect, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteCrossConnectResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) deleteCrossConnect(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/crossConnects/{crossConnectId}")
	if err != nil {
		return nil, err
	}

	var response DeleteCrossConnectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteCrossConnectGroup Deletes the specified cross-connect group. It must not contain any
// cross-connects, and it cannot be mapped to a
// VirtualCircuit.
func (client VirtualNetworkClient) DeleteCrossConnectGroup(ctx context.Context, request DeleteCrossConnectGroupRequest) (response DeleteCrossConnectGroupResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteCrossConnectGroup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteCrossConnectGroupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) deleteCrossConnectGroup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/crossConnectGroups/{crossConnectGroupId}")
	if err != nil {
		return nil, err
	}

	var response DeleteCrossConnectGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteDhcpOptions Deletes the specified set of DHCP options, but only if it's not associated with a subnet. You can't delete a
// VCN's default set of DHCP options.
// This is an asynchronous operation. The state of the set of options will switch to TERMINATING temporarily
// until the set is completely removed.
func (client VirtualNetworkClient) DeleteDhcpOptions(ctx context.Context, request DeleteDhcpOptionsRequest) (response DeleteDhcpOptionsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteDhcpOptions, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteDhcpOptionsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) deleteDhcpOptions(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/dhcps/{dhcpId}")
	if err != nil {
		return nil, err
	}

	var response DeleteDhcpOptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteDrg Deletes the specified DRG. The DRG must not be attached to a VCN or be connected to your on-premise
// network. Also, there must not be a route table that lists the DRG as a target. This is an asynchronous
// operation. The DRG's `lifecycleState` will change to TERMINATING temporarily until the DRG is completely
// removed.
func (client VirtualNetworkClient) DeleteDrg(ctx context.Context, request DeleteDrgRequest) (response DeleteDrgResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteDrg, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteDrgResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) deleteDrg(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/drgs/{drgId}")
	if err != nil {
		return nil, err
	}

	var response DeleteDrgResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteDrgAttachment Detaches a DRG from a VCN by deleting the corresponding `DrgAttachment`. This is an asynchronous
// operation. The attachment's `lifecycleState` will change to DETACHING temporarily until the attachment
// is completely removed.
func (client VirtualNetworkClient) DeleteDrgAttachment(ctx context.Context, request DeleteDrgAttachmentRequest) (response DeleteDrgAttachmentResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteDrgAttachment, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteDrgAttachmentResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) deleteDrgAttachment(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/drgAttachments/{drgAttachmentId}")
	if err != nil {
		return nil, err
	}

	var response DeleteDrgAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteIPSecConnection Deletes the specified IPSec connection. If your goal is to disable the IPSec VPN between your VCN and
// on-premises network, it's easiest to simply detach the DRG but keep all the IPSec VPN components intact.
// If you were to delete all the components and then later need to create an IPSec VPN again, you would
// need to configure your on-premises router again with the new information returned from
// CreateIPSecConnection.
// This is an asynchronous operation. The connection's `lifecycleState` will change to TERMINATING temporarily
// until the connection is completely removed.
func (client VirtualNetworkClient) DeleteIPSecConnection(ctx context.Context, request DeleteIPSecConnectionRequest) (response DeleteIPSecConnectionResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteIPSecConnection, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteIPSecConnectionResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) deleteIPSecConnection(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/ipsecConnections/{ipscId}")
	if err != nil {
		return nil, err
	}

	var response DeleteIPSecConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteInternetGateway Deletes the specified Internet Gateway. The Internet Gateway does not have to be disabled, but
// there must not be a route table that lists it as a target.
// This is an asynchronous operation. The gateway's `lifecycleState` will change to TERMINATING temporarily
// until the gateway is completely removed.
func (client VirtualNetworkClient) DeleteInternetGateway(ctx context.Context, request DeleteInternetGatewayRequest) (response DeleteInternetGatewayResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteInternetGateway, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteInternetGatewayResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) deleteInternetGateway(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/internetGateways/{igId}")
	if err != nil {
		return nil, err
	}

	var response DeleteInternetGatewayResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteLocalPeeringGateway Deletes the specified local peering gateway (LPG).
// This is an asynchronous operation; the local peering gateway's `lifecycleState` changes to TERMINATING temporarily
// until the local peering gateway is completely removed.
func (client VirtualNetworkClient) DeleteLocalPeeringGateway(ctx context.Context, request DeleteLocalPeeringGatewayRequest) (response DeleteLocalPeeringGatewayResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteLocalPeeringGateway, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteLocalPeeringGatewayResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) deleteLocalPeeringGateway(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/localPeeringGateways/{localPeeringGatewayId}")
	if err != nil {
		return nil, err
	}

	var response DeleteLocalPeeringGatewayResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeletePrivateIp Unassigns and deletes the specified private IP. You must
// specify the object's OCID. The private IP address is returned to
// the subnet's pool of available addresses.
// This operation cannot be used with primary private IPs, which are
// automatically unassigned and deleted when the VNIC is terminated.
// **Important:** If a secondary private IP is the
// [target of a route rule]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingroutetables.htm#privateip),
// unassigning it from the VNIC causes that route rule to blackhole and the traffic
// will be dropped.
func (client VirtualNetworkClient) DeletePrivateIp(ctx context.Context, request DeletePrivateIpRequest) (response DeletePrivateIpResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deletePrivateIp, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeletePrivateIpResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) deletePrivateIp(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/privateIps/{privateIpId}")
	if err != nil {
		return nil, err
	}

	var response DeletePrivateIpResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteRouteTable Deletes the specified route table, but only if it's not associated with a subnet. You can't delete a
// VCN's default route table.
// This is an asynchronous operation. The route table's `lifecycleState` will change to TERMINATING temporarily
// until the route table is completely removed.
func (client VirtualNetworkClient) DeleteRouteTable(ctx context.Context, request DeleteRouteTableRequest) (response DeleteRouteTableResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteRouteTable, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteRouteTableResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) deleteRouteTable(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/routeTables/{rtId}")
	if err != nil {
		return nil, err
	}

	var response DeleteRouteTableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteSecurityList Deletes the specified security list, but only if it's not associated with a subnet. You can't delete
// a VCN's default security list.
// This is an asynchronous operation. The security list's `lifecycleState` will change to TERMINATING temporarily
// until the security list is completely removed.
func (client VirtualNetworkClient) DeleteSecurityList(ctx context.Context, request DeleteSecurityListRequest) (response DeleteSecurityListResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteSecurityList, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteSecurityListResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) deleteSecurityList(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/securityLists/{securityListId}")
	if err != nil {
		return nil, err
	}

	var response DeleteSecurityListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteSubnet Deletes the specified subnet, but only if there are no instances in the subnet. This is an asynchronous
// operation. The subnet's `lifecycleState` will change to TERMINATING temporarily. If there are any
// instances in the subnet, the state will instead change back to AVAILABLE.
func (client VirtualNetworkClient) DeleteSubnet(ctx context.Context, request DeleteSubnetRequest) (response DeleteSubnetResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteSubnet, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteSubnetResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) deleteSubnet(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/subnets/{subnetId}")
	if err != nil {
		return nil, err
	}

	var response DeleteSubnetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteVcn Deletes the specified VCN. The VCN must be empty and have no attached gateways. This is an asynchronous
// operation. The VCN's `lifecycleState` will change to TERMINATING temporarily until the VCN is completely
// removed.
func (client VirtualNetworkClient) DeleteVcn(ctx context.Context, request DeleteVcnRequest) (response DeleteVcnResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteVcn, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteVcnResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) deleteVcn(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/vcns/{vcnId}")
	if err != nil {
		return nil, err
	}

	var response DeleteVcnResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteVirtualCircuit Deletes the specified virtual circuit.
// **Important:** If you're using FastConnect via a provider,
// make sure to also terminate the connection with
// the provider, or else the provider may continue to bill you.
func (client VirtualNetworkClient) DeleteVirtualCircuit(ctx context.Context, request DeleteVirtualCircuitRequest) (response DeleteVirtualCircuitResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteVirtualCircuit, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteVirtualCircuitResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) deleteVirtualCircuit(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/virtualCircuits/{virtualCircuitId}")
	if err != nil {
		return nil, err
	}

	var response DeleteVirtualCircuitResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetCpe Gets the specified CPE's information.
func (client VirtualNetworkClient) GetCpe(ctx context.Context, request GetCpeRequest) (response GetCpeResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getCpe, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetCpeResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getCpe(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/cpes/{cpeId}")
	if err != nil {
		return nil, err
	}

	var response GetCpeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetCrossConnect Gets the specified cross-connect's information.
func (client VirtualNetworkClient) GetCrossConnect(ctx context.Context, request GetCrossConnectRequest) (response GetCrossConnectResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getCrossConnect, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetCrossConnectResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getCrossConnect(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/crossConnects/{crossConnectId}")
	if err != nil {
		return nil, err
	}

	var response GetCrossConnectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetCrossConnectGroup Gets the specified cross-connect group's information.
func (client VirtualNetworkClient) GetCrossConnectGroup(ctx context.Context, request GetCrossConnectGroupRequest) (response GetCrossConnectGroupResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getCrossConnectGroup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetCrossConnectGroupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getCrossConnectGroup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/crossConnectGroups/{crossConnectGroupId}")
	if err != nil {
		return nil, err
	}

	var response GetCrossConnectGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetCrossConnectLetterOfAuthority Gets the Letter of Authority for the specified cross-connect.
func (client VirtualNetworkClient) GetCrossConnectLetterOfAuthority(ctx context.Context, request GetCrossConnectLetterOfAuthorityRequest) (response GetCrossConnectLetterOfAuthorityResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getCrossConnectLetterOfAuthority, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetCrossConnectLetterOfAuthorityResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getCrossConnectLetterOfAuthority(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/crossConnects/{crossConnectId}/letterOfAuthority")
	if err != nil {
		return nil, err
	}

	var response GetCrossConnectLetterOfAuthorityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetCrossConnectStatus Gets the status of the specified cross-connect.
func (client VirtualNetworkClient) GetCrossConnectStatus(ctx context.Context, request GetCrossConnectStatusRequest) (response GetCrossConnectStatusResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getCrossConnectStatus, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetCrossConnectStatusResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getCrossConnectStatus(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/crossConnects/{crossConnectId}/status")
	if err != nil {
		return nil, err
	}

	var response GetCrossConnectStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetDhcpOptions Gets the specified set of DHCP options.
func (client VirtualNetworkClient) GetDhcpOptions(ctx context.Context, request GetDhcpOptionsRequest) (response GetDhcpOptionsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getDhcpOptions, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetDhcpOptionsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getDhcpOptions(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dhcps/{dhcpId}")
	if err != nil {
		return nil, err
	}

	var response GetDhcpOptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetDrg Gets the specified DRG's information.
func (client VirtualNetworkClient) GetDrg(ctx context.Context, request GetDrgRequest) (response GetDrgResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getDrg, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetDrgResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getDrg(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/drgs/{drgId}")
	if err != nil {
		return nil, err
	}

	var response GetDrgResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetDrgAttachment Gets the information for the specified `DrgAttachment`.
func (client VirtualNetworkClient) GetDrgAttachment(ctx context.Context, request GetDrgAttachmentRequest) (response GetDrgAttachmentResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getDrgAttachment, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetDrgAttachmentResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getDrgAttachment(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/drgAttachments/{drgAttachmentId}")
	if err != nil {
		return nil, err
	}

	var response GetDrgAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetFastConnectProviderService Gets the specified provider service.
// For more information, see [FastConnect Overview]({{DOC_SERVER_URL}}/Content/Network/Concepts/fastconnect.htm).
func (client VirtualNetworkClient) GetFastConnectProviderService(ctx context.Context, request GetFastConnectProviderServiceRequest) (response GetFastConnectProviderServiceResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getFastConnectProviderService, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetFastConnectProviderServiceResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getFastConnectProviderService(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/fastConnectProviderServices/{providerServiceId}")
	if err != nil {
		return nil, err
	}

	var response GetFastConnectProviderServiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetIPSecConnection Gets the specified IPSec connection's basic information, including the static routes for the
// on-premises router. If you want the status of the connection (whether it's up or down), use
// GetIPSecConnectionDeviceStatus.
func (client VirtualNetworkClient) GetIPSecConnection(ctx context.Context, request GetIPSecConnectionRequest) (response GetIPSecConnectionResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getIPSecConnection, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetIPSecConnectionResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getIPSecConnection(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/ipsecConnections/{ipscId}")
	if err != nil {
		return nil, err
	}

	var response GetIPSecConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetIPSecConnectionDeviceConfig Gets the configuration information for the specified IPSec connection. For each tunnel, the
// response includes the IP address of Oracle's VPN headend and the shared secret.
func (client VirtualNetworkClient) GetIPSecConnectionDeviceConfig(ctx context.Context, request GetIPSecConnectionDeviceConfigRequest) (response GetIPSecConnectionDeviceConfigResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getIPSecConnectionDeviceConfig, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetIPSecConnectionDeviceConfigResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getIPSecConnectionDeviceConfig(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/ipsecConnections/{ipscId}/deviceConfig")
	if err != nil {
		return nil, err
	}

	var response GetIPSecConnectionDeviceConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetIPSecConnectionDeviceStatus Gets the status of the specified IPSec connection (whether it's up or down).
func (client VirtualNetworkClient) GetIPSecConnectionDeviceStatus(ctx context.Context, request GetIPSecConnectionDeviceStatusRequest) (response GetIPSecConnectionDeviceStatusResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getIPSecConnectionDeviceStatus, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetIPSecConnectionDeviceStatusResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getIPSecConnectionDeviceStatus(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/ipsecConnections/{ipscId}/deviceStatus")
	if err != nil {
		return nil, err
	}

	var response GetIPSecConnectionDeviceStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetInternetGateway Gets the specified Internet Gateway's information.
func (client VirtualNetworkClient) GetInternetGateway(ctx context.Context, request GetInternetGatewayRequest) (response GetInternetGatewayResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getInternetGateway, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetInternetGatewayResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getInternetGateway(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/internetGateways/{igId}")
	if err != nil {
		return nil, err
	}

	var response GetInternetGatewayResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetLocalPeeringGateway Gets the specified local peering gateway's information.
func (client VirtualNetworkClient) GetLocalPeeringGateway(ctx context.Context, request GetLocalPeeringGatewayRequest) (response GetLocalPeeringGatewayResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getLocalPeeringGateway, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetLocalPeeringGatewayResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getLocalPeeringGateway(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/localPeeringGateways/{localPeeringGatewayId}")
	if err != nil {
		return nil, err
	}

	var response GetLocalPeeringGatewayResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetPrivateIp Gets the specified private IP. You must specify the object's OCID.
// Alternatively, you can get the object by using
// ListPrivateIps
// with the private IP address (for example, 10.0.3.3) and subnet OCID.
func (client VirtualNetworkClient) GetPrivateIp(ctx context.Context, request GetPrivateIpRequest) (response GetPrivateIpResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getPrivateIp, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetPrivateIpResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getPrivateIp(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/privateIps/{privateIpId}")
	if err != nil {
		return nil, err
	}

	var response GetPrivateIpResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetRouteTable Gets the specified route table's information.
func (client VirtualNetworkClient) GetRouteTable(ctx context.Context, request GetRouteTableRequest) (response GetRouteTableResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getRouteTable, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetRouteTableResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getRouteTable(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/routeTables/{rtId}")
	if err != nil {
		return nil, err
	}

	var response GetRouteTableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetSecurityList Gets the specified security list's information.
func (client VirtualNetworkClient) GetSecurityList(ctx context.Context, request GetSecurityListRequest) (response GetSecurityListResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getSecurityList, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetSecurityListResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getSecurityList(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/securityLists/{securityListId}")
	if err != nil {
		return nil, err
	}

	var response GetSecurityListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetSubnet Gets the specified subnet's information.
func (client VirtualNetworkClient) GetSubnet(ctx context.Context, request GetSubnetRequest) (response GetSubnetResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getSubnet, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetSubnetResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getSubnet(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/subnets/{subnetId}")
	if err != nil {
		return nil, err
	}

	var response GetSubnetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetVcn Gets the specified VCN's information.
func (client VirtualNetworkClient) GetVcn(ctx context.Context, request GetVcnRequest) (response GetVcnResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getVcn, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetVcnResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getVcn(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/vcns/{vcnId}")
	if err != nil {
		return nil, err
	}

	var response GetVcnResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetVirtualCircuit Gets the specified virtual circuit's information.
func (client VirtualNetworkClient) GetVirtualCircuit(ctx context.Context, request GetVirtualCircuitRequest) (response GetVirtualCircuitResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getVirtualCircuit, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetVirtualCircuitResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getVirtualCircuit(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/virtualCircuits/{virtualCircuitId}")
	if err != nil {
		return nil, err
	}

	var response GetVirtualCircuitResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetVnic Gets the information for the specified virtual network interface card (VNIC).
// You can get the VNIC OCID from the
// ListVnicAttachments
// operation.
func (client VirtualNetworkClient) GetVnic(ctx context.Context, request GetVnicRequest) (response GetVnicResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getVnic, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetVnicResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) getVnic(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/vnics/{vnicId}")
	if err != nil {
		return nil, err
	}

	var response GetVnicResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListCpes Lists the Customer-Premises Equipment objects (CPEs) in the specified compartment.
func (client VirtualNetworkClient) ListCpes(ctx context.Context, request ListCpesRequest) (response ListCpesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listCpes, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListCpesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listCpes(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/cpes")
	if err != nil {
		return nil, err
	}

	var response ListCpesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListCrossConnectGroups Lists the cross-connect groups in the specified compartment.
func (client VirtualNetworkClient) ListCrossConnectGroups(ctx context.Context, request ListCrossConnectGroupsRequest) (response ListCrossConnectGroupsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listCrossConnectGroups, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListCrossConnectGroupsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listCrossConnectGroups(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/crossConnectGroups")
	if err != nil {
		return nil, err
	}

	var response ListCrossConnectGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListCrossConnectLocations Lists the available FastConnect locations for cross-connect installation. You need
// this information so you can specify your desired location when you create a cross-connect.
func (client VirtualNetworkClient) ListCrossConnectLocations(ctx context.Context, request ListCrossConnectLocationsRequest) (response ListCrossConnectLocationsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listCrossConnectLocations, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListCrossConnectLocationsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listCrossConnectLocations(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/crossConnectLocations")
	if err != nil {
		return nil, err
	}

	var response ListCrossConnectLocationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListCrossConnects Lists the cross-connects in the specified compartment. You can filter the list
// by specifying the OCID of a cross-connect group.
func (client VirtualNetworkClient) ListCrossConnects(ctx context.Context, request ListCrossConnectsRequest) (response ListCrossConnectsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listCrossConnects, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListCrossConnectsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listCrossConnects(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/crossConnects")
	if err != nil {
		return nil, err
	}

	var response ListCrossConnectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListCrossconnectPortSpeedShapes Lists the available port speeds for cross-connects. You need this information
// so you can specify your desired port speed (that is, shape) when you create a
// cross-connect.
func (client VirtualNetworkClient) ListCrossconnectPortSpeedShapes(ctx context.Context, request ListCrossconnectPortSpeedShapesRequest) (response ListCrossconnectPortSpeedShapesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listCrossconnectPortSpeedShapes, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListCrossconnectPortSpeedShapesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listCrossconnectPortSpeedShapes(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/crossConnectPortSpeedShapes")
	if err != nil {
		return nil, err
	}

	var response ListCrossconnectPortSpeedShapesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListDhcpOptions Lists the sets of DHCP options in the specified VCN and specified compartment.
// The response includes the default set of options that automatically comes with each VCN,
// plus any other sets you've created.
func (client VirtualNetworkClient) ListDhcpOptions(ctx context.Context, request ListDhcpOptionsRequest) (response ListDhcpOptionsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listDhcpOptions, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListDhcpOptionsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listDhcpOptions(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/dhcps")
	if err != nil {
		return nil, err
	}

	var response ListDhcpOptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListDrgAttachments Lists the `DrgAttachment` objects for the specified compartment. You can filter the
// results by VCN or DRG.
func (client VirtualNetworkClient) ListDrgAttachments(ctx context.Context, request ListDrgAttachmentsRequest) (response ListDrgAttachmentsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listDrgAttachments, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListDrgAttachmentsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listDrgAttachments(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/drgAttachments")
	if err != nil {
		return nil, err
	}

	var response ListDrgAttachmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListDrgs Lists the DRGs in the specified compartment.
func (client VirtualNetworkClient) ListDrgs(ctx context.Context, request ListDrgsRequest) (response ListDrgsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listDrgs, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListDrgsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listDrgs(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/drgs")
	if err != nil {
		return nil, err
	}

	var response ListDrgsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListFastConnectProviderServices Lists the service offerings from supported providers. You need this
// information so you can specify your desired provider and service
// offering when you create a virtual circuit.
// For the compartment ID, provide the OCID of your tenancy (the root compartment).
// For more information, see [FastConnect Overview]({{DOC_SERVER_URL}}/Content/Network/Concepts/fastconnect.htm).
func (client VirtualNetworkClient) ListFastConnectProviderServices(ctx context.Context, request ListFastConnectProviderServicesRequest) (response ListFastConnectProviderServicesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listFastConnectProviderServices, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListFastConnectProviderServicesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listFastConnectProviderServices(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/fastConnectProviderServices")
	if err != nil {
		return nil, err
	}

	var response ListFastConnectProviderServicesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListFastConnectProviderVirtualCircuitBandwidthShapes Gets the list of available virtual circuit bandwidth levels for a provider.
// You need this information so you can specify your desired bandwidth level (shape) when you create a virtual circuit.
// For more information about virtual circuits, see [FastConnect Overview]({{DOC_SERVER_URL}}/Content/Network/Concepts/fastconnect.htm).
func (client VirtualNetworkClient) ListFastConnectProviderVirtualCircuitBandwidthShapes(ctx context.Context, request ListFastConnectProviderVirtualCircuitBandwidthShapesRequest) (response ListFastConnectProviderVirtualCircuitBandwidthShapesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listFastConnectProviderVirtualCircuitBandwidthShapes, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListFastConnectProviderVirtualCircuitBandwidthShapesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listFastConnectProviderVirtualCircuitBandwidthShapes(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/fastConnectProviderServices/{providerServiceId}/virtualCircuitBandwidthShapes")
	if err != nil {
		return nil, err
	}

	var response ListFastConnectProviderVirtualCircuitBandwidthShapesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListIPSecConnections Lists the IPSec connections for the specified compartment. You can filter the
// results by DRG or CPE.
func (client VirtualNetworkClient) ListIPSecConnections(ctx context.Context, request ListIPSecConnectionsRequest) (response ListIPSecConnectionsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listIPSecConnections, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListIPSecConnectionsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listIPSecConnections(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/ipsecConnections")
	if err != nil {
		return nil, err
	}

	var response ListIPSecConnectionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListInternetGateways Lists the Internet Gateways in the specified VCN and the specified compartment.
func (client VirtualNetworkClient) ListInternetGateways(ctx context.Context, request ListInternetGatewaysRequest) (response ListInternetGatewaysResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listInternetGateways, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListInternetGatewaysResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listInternetGateways(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/internetGateways")
	if err != nil {
		return nil, err
	}

	var response ListInternetGatewaysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListLocalPeeringGateways Lists the local peering gateways (LPGs) for the specified VCN and compartment
// (the LPG's compartment).
func (client VirtualNetworkClient) ListLocalPeeringGateways(ctx context.Context, request ListLocalPeeringGatewaysRequest) (response ListLocalPeeringGatewaysResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listLocalPeeringGateways, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListLocalPeeringGatewaysResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listLocalPeeringGateways(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/localPeeringGateways")
	if err != nil {
		return nil, err
	}

	var response ListLocalPeeringGatewaysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListPrivateIps Lists the PrivateIp objects based
// on one of these filters:
//   - Subnet OCID.
//   - VNIC OCID.
//   - Both private IP address and subnet OCID: This lets
//   you get a `privateIP` object based on its private IP
//   address (for example, 10.0.3.3) and not its OCID. For comparison,
//   GetPrivateIp
//   requires the OCID.
// If you're listing all the private IPs associated with a given subnet
// or VNIC, the response includes both primary and secondary private IPs.
func (client VirtualNetworkClient) ListPrivateIps(ctx context.Context, request ListPrivateIpsRequest) (response ListPrivateIpsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listPrivateIps, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListPrivateIpsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listPrivateIps(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/privateIps")
	if err != nil {
		return nil, err
	}

	var response ListPrivateIpsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListRouteTables Lists the route tables in the specified VCN and specified compartment. The response
// includes the default route table that automatically comes with each VCN, plus any route tables
// you've created.
func (client VirtualNetworkClient) ListRouteTables(ctx context.Context, request ListRouteTablesRequest) (response ListRouteTablesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listRouteTables, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListRouteTablesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listRouteTables(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/routeTables")
	if err != nil {
		return nil, err
	}

	var response ListRouteTablesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListSecurityLists Lists the security lists in the specified VCN and compartment.
func (client VirtualNetworkClient) ListSecurityLists(ctx context.Context, request ListSecurityListsRequest) (response ListSecurityListsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listSecurityLists, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListSecurityListsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listSecurityLists(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/securityLists")
	if err != nil {
		return nil, err
	}

	var response ListSecurityListsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListSubnets Lists the subnets in the specified VCN and the specified compartment.
func (client VirtualNetworkClient) ListSubnets(ctx context.Context, request ListSubnetsRequest) (response ListSubnetsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listSubnets, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListSubnetsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listSubnets(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/subnets")
	if err != nil {
		return nil, err
	}

	var response ListSubnetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListVcns Lists the Virtual Cloud Networks (VCNs) in the specified compartment.
func (client VirtualNetworkClient) ListVcns(ctx context.Context, request ListVcnsRequest) (response ListVcnsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listVcns, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListVcnsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listVcns(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/vcns")
	if err != nil {
		return nil, err
	}

	var response ListVcnsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListVirtualCircuitBandwidthShapes The deprecated operation lists available bandwidth levels for virtual circuits. For the compartment ID, provide the OCID of your tenancy (the root compartment).
func (client VirtualNetworkClient) ListVirtualCircuitBandwidthShapes(ctx context.Context, request ListVirtualCircuitBandwidthShapesRequest) (response ListVirtualCircuitBandwidthShapesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listVirtualCircuitBandwidthShapes, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListVirtualCircuitBandwidthShapesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listVirtualCircuitBandwidthShapes(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/virtualCircuitBandwidthShapes")
	if err != nil {
		return nil, err
	}

	var response ListVirtualCircuitBandwidthShapesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListVirtualCircuitPublicPrefixes Lists the public IP prefixes and their details for the specified
// public virtual circuit.
func (client VirtualNetworkClient) ListVirtualCircuitPublicPrefixes(ctx context.Context, request ListVirtualCircuitPublicPrefixesRequest) (response ListVirtualCircuitPublicPrefixesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listVirtualCircuitPublicPrefixes, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListVirtualCircuitPublicPrefixesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listVirtualCircuitPublicPrefixes(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/virtualCircuits/{virtualCircuitId}/publicPrefixes")
	if err != nil {
		return nil, err
	}

	var response ListVirtualCircuitPublicPrefixesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListVirtualCircuits Lists the virtual circuits in the specified compartment.
func (client VirtualNetworkClient) ListVirtualCircuits(ctx context.Context, request ListVirtualCircuitsRequest) (response ListVirtualCircuitsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listVirtualCircuits, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListVirtualCircuitsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) listVirtualCircuits(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/virtualCircuits")
	if err != nil {
		return nil, err
	}

	var response ListVirtualCircuitsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateCpe Updates the specified CPE's display name.
// Avoid entering confidential information.
func (client VirtualNetworkClient) UpdateCpe(ctx context.Context, request UpdateCpeRequest) (response UpdateCpeResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateCpe, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateCpeResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) updateCpe(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/cpes/{cpeId}")
	if err != nil {
		return nil, err
	}

	var response UpdateCpeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateCrossConnect Updates the specified cross-connect.
func (client VirtualNetworkClient) UpdateCrossConnect(ctx context.Context, request UpdateCrossConnectRequest) (response UpdateCrossConnectResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateCrossConnect, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateCrossConnectResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) updateCrossConnect(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/crossConnects/{crossConnectId}")
	if err != nil {
		return nil, err
	}

	var response UpdateCrossConnectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateCrossConnectGroup Updates the specified cross-connect group's display name.
// Avoid entering confidential information.
func (client VirtualNetworkClient) UpdateCrossConnectGroup(ctx context.Context, request UpdateCrossConnectGroupRequest) (response UpdateCrossConnectGroupResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateCrossConnectGroup, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateCrossConnectGroupResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) updateCrossConnectGroup(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/crossConnectGroups/{crossConnectGroupId}")
	if err != nil {
		return nil, err
	}

	var response UpdateCrossConnectGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateDhcpOptions Updates the specified set of DHCP options. You can update the display name or the options
// themselves. Avoid entering confidential information.
// Note that the `options` object you provide replaces the entire existing set of options.
func (client VirtualNetworkClient) UpdateDhcpOptions(ctx context.Context, request UpdateDhcpOptionsRequest) (response UpdateDhcpOptionsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateDhcpOptions, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateDhcpOptionsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) updateDhcpOptions(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/dhcps/{dhcpId}")
	if err != nil {
		return nil, err
	}

	var response UpdateDhcpOptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateDrg Updates the specified DRG's display name. Avoid entering confidential information.
func (client VirtualNetworkClient) UpdateDrg(ctx context.Context, request UpdateDrgRequest) (response UpdateDrgResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateDrg, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateDrgResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) updateDrg(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/drgs/{drgId}")
	if err != nil {
		return nil, err
	}

	var response UpdateDrgResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateDrgAttachment Updates the display name for the specified `DrgAttachment`.
// Avoid entering confidential information.
func (client VirtualNetworkClient) UpdateDrgAttachment(ctx context.Context, request UpdateDrgAttachmentRequest) (response UpdateDrgAttachmentResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateDrgAttachment, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateDrgAttachmentResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) updateDrgAttachment(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/drgAttachments/{drgAttachmentId}")
	if err != nil {
		return nil, err
	}

	var response UpdateDrgAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateIPSecConnection Updates the display name for the specified IPSec connection.
// Avoid entering confidential information.
func (client VirtualNetworkClient) UpdateIPSecConnection(ctx context.Context, request UpdateIPSecConnectionRequest) (response UpdateIPSecConnectionResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateIPSecConnection, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateIPSecConnectionResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) updateIPSecConnection(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/ipsecConnections/{ipscId}")
	if err != nil {
		return nil, err
	}

	var response UpdateIPSecConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateInternetGateway Updates the specified Internet Gateway. You can disable/enable it, or change its display name.
// Avoid entering confidential information.
// If the gateway is disabled, that means no traffic will flow to/from the internet even if there's
// a route rule that enables that traffic.
func (client VirtualNetworkClient) UpdateInternetGateway(ctx context.Context, request UpdateInternetGatewayRequest) (response UpdateInternetGatewayResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateInternetGateway, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateInternetGatewayResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) updateInternetGateway(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/internetGateways/{igId}")
	if err != nil {
		return nil, err
	}

	var response UpdateInternetGatewayResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateLocalPeeringGateway Updates the specified local peering gateway (LPG).
func (client VirtualNetworkClient) UpdateLocalPeeringGateway(ctx context.Context, request UpdateLocalPeeringGatewayRequest) (response UpdateLocalPeeringGatewayResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateLocalPeeringGateway, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateLocalPeeringGatewayResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) updateLocalPeeringGateway(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/localPeeringGateways/{localPeeringGatewayId}")
	if err != nil {
		return nil, err
	}

	var response UpdateLocalPeeringGatewayResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdatePrivateIp Updates the specified private IP. You must specify the object's OCID.
// Use this operation if you want to:
//   - Move a secondary private IP to a different VNIC in the same subnet.
//   - Change the display name for a secondary private IP.
//   - Change the hostname for a secondary private IP.
// This operation cannot be used with primary private IPs.
// To update the hostname for the primary IP on a VNIC, use
// UpdateVnic.
func (client VirtualNetworkClient) UpdatePrivateIp(ctx context.Context, request UpdatePrivateIpRequest) (response UpdatePrivateIpResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updatePrivateIp, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdatePrivateIpResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) updatePrivateIp(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/privateIps/{privateIpId}")
	if err != nil {
		return nil, err
	}

	var response UpdatePrivateIpResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateRouteTable Updates the specified route table's display name or route rules.
// Avoid entering confidential information.
// Note that the `routeRules` object you provide replaces the entire existing set of rules.
func (client VirtualNetworkClient) UpdateRouteTable(ctx context.Context, request UpdateRouteTableRequest) (response UpdateRouteTableResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateRouteTable, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateRouteTableResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) updateRouteTable(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/routeTables/{rtId}")
	if err != nil {
		return nil, err
	}

	var response UpdateRouteTableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateSecurityList Updates the specified security list's display name or rules.
// Avoid entering confidential information.
// Note that the `egressSecurityRules` or `ingressSecurityRules` objects you provide replace the entire
// existing objects.
func (client VirtualNetworkClient) UpdateSecurityList(ctx context.Context, request UpdateSecurityListRequest) (response UpdateSecurityListResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateSecurityList, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateSecurityListResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) updateSecurityList(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/securityLists/{securityListId}")
	if err != nil {
		return nil, err
	}

	var response UpdateSecurityListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateSubnet Updates the specified subnet's display name. Avoid entering confidential information.
func (client VirtualNetworkClient) UpdateSubnet(ctx context.Context, request UpdateSubnetRequest) (response UpdateSubnetResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateSubnet, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateSubnetResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) updateSubnet(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/subnets/{subnetId}")
	if err != nil {
		return nil, err
	}

	var response UpdateSubnetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateVcn Updates the specified VCN's display name.
// Avoid entering confidential information.
func (client VirtualNetworkClient) UpdateVcn(ctx context.Context, request UpdateVcnRequest) (response UpdateVcnResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateVcn, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateVcnResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) updateVcn(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/vcns/{vcnId}")
	if err != nil {
		return nil, err
	}

	var response UpdateVcnResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateVirtualCircuit Updates the specified virtual circuit. This can be called by
// either the customer who owns the virtual circuit, or the
// provider (when provisioning or de-provisioning the virtual
// circuit from their end). The documentation for
// UpdateVirtualCircuitDetails
// indicates who can update each property of the virtual circuit.
// **Important:** If the virtual circuit is working and in the
// PROVISIONED state, updating any of the network-related properties
// (such as the DRG being used, the BGP ASN, and so on) will cause the virtual
// circuit's state to switch to PROVISIONING and the related BGP
// session to go down. After Oracle re-provisions the virtual circuit,
// its state will return to PROVISIONED. Make sure you confirm that
// the associated BGP session is back up. For more information
// about the various states and how to test connectivity, see
// [FastConnect Overview]({{DOC_SERVER_URL}}/Content/Network/Concepts/fastconnect.htm).
// To change the list of public IP prefixes for a public virtual circuit,
// use BulkAddVirtualCircuitPublicPrefixes
// and
// BulkDeleteVirtualCircuitPublicPrefixes.
// Updating the list of prefixes does NOT cause the BGP session to go down. However,
// Oracle must verify the customer's ownership of each added prefix before
// traffic for that prefix will flow across the virtual circuit.
func (client VirtualNetworkClient) UpdateVirtualCircuit(ctx context.Context, request UpdateVirtualCircuitRequest) (response UpdateVirtualCircuitResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateVirtualCircuit, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateVirtualCircuitResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) updateVirtualCircuit(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/virtualCircuits/{virtualCircuitId}")
	if err != nil {
		return nil, err
	}

	var response UpdateVirtualCircuitResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateVnic Updates the specified VNIC.
func (client VirtualNetworkClient) UpdateVnic(ctx context.Context, request UpdateVnicRequest) (response UpdateVnicResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateVnic, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateVnicResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client VirtualNetworkClient) updateVnic(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/vnics/{vnicId}")
	if err != nil {
		return nil, err
	}

	var response UpdateVnicResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}
