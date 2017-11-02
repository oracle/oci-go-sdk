// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// DhcpOption. A single DHCP option according to [RFC 1533](https://tools.ietf.org/html/rfc1533).
// The two options available to use are DhcpDnsOption
// and DhcpSearchDomainOption. For more
// information, see [DNS in Your Virtual Cloud Network]({{DOC_SERVER_URL}}/Content/Network/Concepts/dns.htm)
// and [DHCP Options]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingDHCP.htm).
type DhcpOption struct {

	// The specific DHCP option. Either `DomainNameServer`
	// (for DhcpDnsOption) or
	// `SearchDomain` (for DhcpSearchDomainOption).
	Type_ *string `mandatory:"true" json:"type,omitempty"`
}

func (model DhcpOption) String() string {
	return common.PointerString(model)
}
