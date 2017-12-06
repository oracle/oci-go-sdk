// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// DhcpDnsOption. DHCP option for specifying how DNS (hostname resolution) is handled in the subnets in the VCN.
// For more information, see
// [DNS in Your Virtual Cloud Network]({{DOC_SERVER_URL}}/Content/Network/Concepts/dns.htm).
type DhcpDnsOption struct {

	// If you set `serverType` to `CustomDnsServer`, specify the IP address
	// of at least one DNS server of your choice (three maximum).
	CustomDnsServers []string `mandatory:"false" json:"customDnsServers,omitempty"`

	// - **VcnLocal:** Reserved for future use.
	// - **VcnLocalPlusInternet:** Also referred to as "Internet and VCN Resolver".
	// Instances can resolve internet hostnames (no Internet Gateway is required),
	// and can resolve hostnames of instances in the VCN. This is the default
	// value in the default set of DHCP options in the VCN. For the Internet and
	// VCN Resolver to work across the VCN, there must also be a DNS label set for
	// the VCN, a DNS label set for each subnet, and a hostname for each instance.
	// The Internet and VCN Resolver also enables reverse DNS lookup, which lets
	// you determine the hostname corresponding to the private IP address. For more
	// information, see
	// [DNS in Your Virtual Cloud Network]({{DOC_SERVER_URL}}/Content/Network/Concepts/dns.htm).
	// - **CustomDnsServer:** Instances use a DNS server of your choice (three maximum).
	ServerType DhcpDnsOptionServerTypeEnum `mandatory:"true" json:"serverType,omitempty"`
}

func (m DhcpDnsOption) String() string {
	return common.PointerString(m)
}

func (m DhcpDnsOption) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDhcpDnsOption DhcpDnsOption
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDhcpDnsOption
	}{
		"DomainNameServer",
		(MarshalTypeDhcpDnsOption)(m),
	}

	return json.Marshal(&s)
}

type DhcpDnsOptionServerTypeEnum string

const (
	DHCP_DNS_OPTION_SERVER_TYPE_VCNLOCAL             DhcpDnsOptionServerTypeEnum = "VcnLocal"
	DHCP_DNS_OPTION_SERVER_TYPE_VCNLOCALPLUSINTERNET DhcpDnsOptionServerTypeEnum = "VcnLocalPlusInternet"
	DHCP_DNS_OPTION_SERVER_TYPE_CUSTOMDNSSERVER      DhcpDnsOptionServerTypeEnum = "CustomDnsServer"
	DHCP_DNS_OPTION_SERVER_TYPE_UNKNOWN              DhcpDnsOptionServerTypeEnum = "UNKNOWN"
)

var mapping_dhcpdnsoption_serverType = map[string]DhcpDnsOptionServerTypeEnum{
	"VcnLocal":             DHCP_DNS_OPTION_SERVER_TYPE_VCNLOCAL,
	"VcnLocalPlusInternet": DHCP_DNS_OPTION_SERVER_TYPE_VCNLOCALPLUSINTERNET,
	"CustomDnsServer":      DHCP_DNS_OPTION_SERVER_TYPE_CUSTOMDNSSERVER,
	"UNKNOWN":              DHCP_DNS_OPTION_SERVER_TYPE_UNKNOWN,
}

func GetDhcpDnsOptionServerTypeEnumValues() []DhcpDnsOptionServerTypeEnum {
	values := make([]DhcpDnsOptionServerTypeEnum, 0)
	for _, v := range mapping_dhcpdnsoption_serverType {
		if v != DHCP_DNS_OPTION_SERVER_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
