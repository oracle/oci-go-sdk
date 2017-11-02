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

type UpdatePrivateIpDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// The hostname for the private IP. Used for DNS. The value
	// is the hostname portion of the private IP's fully qualified domain name (FQDN)
	// (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`).
	// Must be unique across all VNICs in the subnet and comply with
	// [RFC 952](https://tools.ietf.org/html/rfc952) and
	// [RFC 1123](https://tools.ietf.org/html/rfc1123).
	// For more information, see
	// [DNS in Your Virtual Cloud Network]({{DOC_SERVER_URL}}/Content/Network/Concepts/dns.htm).
	// Example: `bminstance-1`
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel,omitempty"`

	// The OCID of the VNIC to reassign the private IP to. The VNIC must
	// be in the same subnet as the current VNIC.
	VnicID *string `mandatory:"false" json:"vnicId,omitempty"`
}

func (model UpdatePrivateIpDetails) String() string {
	return common.PointerString(model)
}
