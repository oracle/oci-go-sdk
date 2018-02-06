// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateSubnetDetails The representation of CreateSubnetDetails
type CreateSubnetDetails struct {

	// The Availability Domain to contain the subnet.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The CIDR IP address range of the subnet.
	// Example: `172.16.1.0/24`
	CidrBlock *string `mandatory:"true" json:"cidrBlock"`

	// The OCID of the compartment to contain the subnet.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the VCN to contain the subnet.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID of the set of DHCP options the subnet will use. If you don't
	// provide a value, the subnet will use the VCN's default set of DHCP options.
	DhcpOptionsId *string `mandatory:"false" json:"dhcpOptionsId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A DNS label for the subnet, used in conjunction with the VNIC's hostname and
	// VCN's DNS label to form a fully qualified domain name (FQDN) for each VNIC
	// within this subnet (for example, `bminstance-1.subnet123.vcn1.oraclevcn.com`).
	// Must be an alphanumeric string that begins with a letter and is unique within the VCN.
	// The value cannot be changed.
	// This value must be set if you want to use the Internet and VCN Resolver to resolve the
	// hostnames of instances in the subnet. It can only be set if the VCN itself
	// was created with a DNS label.
	// For more information, see
	// https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm.
	// Example: `subnet123`
	DnsLabel *string `mandatory:"false" json:"dnsLabel"`

	// Whether VNICs within this subnet can have public IP addresses.
	// Defaults to false, which means VNICs created in this subnet will
	// automatically be assigned public IP addresses unless specified
	// otherwise during instance launch or VNIC creation (with the
	// `assignPublicIp` flag in CreateVnicDetails).
	// If `prohibitPublicIpOnVnic` is set to true, VNICs created in this
	// subnet cannot have public IP addresses (that is, it's a private
	// subnet).
	// Example: `true`
	ProhibitPublicIpOnVnic *bool `mandatory:"false" json:"prohibitPublicIpOnVnic"`

	// The OCID of the route table the subnet will use. If you don't provide a value,
	// the subnet will use the VCN's default route table.
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// OCIDs for the security lists to associate with the subnet. If you don't
	// provide a value, the VCN's default security list will be associated with
	// the subnet. Remember that security lists are associated at the subnet
	// level, but the rules are applied to the individual VNICs in the subnet.
	SecurityListIds []string `mandatory:"false" json:"securityListIds"`
}

func (m CreateSubnetDetails) String() string {
	return common.PointerString(m)
}
