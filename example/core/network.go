package core

import (
	"context"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/example/helper"
	"github.com/oracle/oci-go-sdk/example/identity"
)

// ListVcns list all VCNs under compartment
func ListVcns() []core.Vcn {
	// create a compartment, otherwise return the id of existing compartment
	compartmentID, err := identity.CreateCompartment()
	helper.LogIfError(err)

	request := core.ListVcnsRequest{
		CompartmentId: compartmentID,
	}

	c, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helper.LogIfError(err)

	r, err := c.ListVcns(context.Background(), request)
	helper.LogIfError(err)

	return r.Items
}

// CreateVcn create virtual cloud network
func CreateVcn() core.Vcn {
	var vcn core.Vcn
	vcnItems := ListVcns()
	vcnDisplayName := "GOSDK_Sample_CreateVcn"

	// check if the VCN has already been created
	for _, element := range vcnItems {
		if *element.DisplayName == vcnDisplayName {
			vcn = element
			break
		}
	}

	// just return the VCN
	if vcn.Id != nil {
		return vcn
	}

	// create a compartment, otherwise return the id of existing compartment
	compartmentID, err := identity.CreateCompartment()
	helper.LogIfError(err)

	request := core.CreateVcnRequest{}
	request.CidrBlock = common.String("10.0.0.0/16")
	request.CompartmentId = compartmentID
	request.DisplayName = common.String(vcnDisplayName)

	c, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helper.LogIfError(err)

	r, err := c.CreateVcn(context.Background(), request)

	if r.RawResponse.StatusCode == 409 {
		err = nil // ignore the error if VNC already exist
	}

	helper.LogIfError(err)
	return r.Vcn
}

// ListSubnets lists subnets under a compartment
func ListSubnets() []core.Subnet {
	vcn := CreateVcn()
	// create a compartment, otherwise return the id of existing compartment
	compartmentID, err := identity.CreateCompartment()
	helper.LogIfError(err)

	request := core.ListSubnetsRequest{
		CompartmentId: compartmentID,
		VcnId:         vcn.Id,
	}

	c, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helper.LogIfError(err)

	r, err := c.ListSubnets(context.Background(), request)
	helper.LogIfError(err)

	return r.Items
}

// CreateSubnet creates a subnet
func CreateSubnet() core.Subnet {
	var subnet core.Subnet
	subnets := ListSubnets()
	subnetDisplayName := "OCI_GOSDK_Sample_Subnet"

	// check if the subnet has already been created
	for _, element := range subnets {
		if *element.DisplayName == subnetDisplayName {
			subnet = element
			break
		}
	}

	// just return the VCN
	if subnet.Id != nil {
		return subnet
	}

	// create a compartment, otherwise return the id of existing compartment

	vcn := CreateVcn()
	availabilityDomains := identity.ListAvailabilityDomains()

	request := core.CreateSubnetRequest{}
	request.CompartmentId = vcn.CompartmentId
	request.CidrBlock = vcn.CidrBlock
	request.VcnId = vcn.Id
	request.DisplayName = common.String(subnetDisplayName)
	request.AvailabilityDomain = availabilityDomains[0].Name

	c, err := core.NewVirtualNetworkClientWithConfigurationProvider(common.DefaultConfigProvider())
	helper.LogIfError(err)

	r, err := c.CreateSubnet(context.Background(), request)
	helper.LogIfError(err)
	return r.Subnet
}
