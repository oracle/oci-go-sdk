// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package coreSamples

import (
	"context"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/example"
)

// CreateVcn creates a new Virtual Cloud Network (VCN)
func CreateVcn(compartmentID *string) core.Vcn {
	request := core.CreateVcnRequest{}
	request.CidrBlock = common.String("10.0.0.0/16") // The CIDR IP address block of the VCN.
	request.CompartmentId = compartmentID
	request.DisplayName = common.String("OCI GOSDK VCN DisplayName")

	c, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	example.LogIfError(err)

	r, err := c.CreateVcn(context.Background(), request)
	example.LogIfError(err)
	return r.Vcn
}

// CreateSubnet creates a new subnet in the specified VCN
func CreateSubnet(vcn core.Vcn, availabilityDomain *string, compartmentID *string) core.Subnet {
	request := core.CreateSubnetRequest{}
	request.CompartmentId = compartmentID
	request.CidrBlock = common.String("10.0.0.0/16") // The CIDR IP address range of the subnet.
	request.VcnId = vcn.Id
	request.DisplayName = common.String("OCI GOSDK Subnet DisplayName")
	request.AvailabilityDomain = availabilityDomain

	c, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	example.LogIfError(err)

	r, err := c.CreateSubnet(context.Background(), request)
	example.LogIfError(err)
	return r.Subnet
}
