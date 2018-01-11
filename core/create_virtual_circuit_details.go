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

type CreateVirtualCircuitDetails struct {

	// The OCID of the compartment to contain the virtual circuit.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The type of IP addresses used in this virtual circuit. PRIVATE
	// means [RFC 1918](https://tools.ietf.org/html/rfc1918) addresses
	// (10.0.0.0/8, 172.16/12, and 192.168/16). Only PRIVATE is supported.
	Type CreateVirtualCircuitDetailsType_Enum `mandatory:"true" json:"type,omitempty"`

	// The provisioned data rate of the connection.  To get a list of the
	// available bandwidth levels (that is, shapes), see
	// ListFastConnectProviderVirtualCircuitBandwidthShapes.
	// Example: `10 Gbps`
	BandwidthShapeName *string `mandatory:"false" json:"bandwidthShapeName,omitempty"`

	// Create a `CrossConnectMapping` for each cross-connect or cross-connect
	// group this virtual circuit will run on.
	CrossConnectMappings []CrossConnectMapping `mandatory:"false" json:"crossConnectMappings,omitempty"`

	// Your BGP ASN (either public or private). Provide this value only if
	// there's a BGP session that goes from your edge router to Oracle.
	// Otherwise, leave this empty or null.
	CustomerBgpAsn *int `mandatory:"false" json:"customerBgpAsn,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// For private virtual circuits only. The OCID of the Drg
	// that this virtual circuit uses.
	GatewayID *string `mandatory:"false" json:"gatewayId,omitempty"`

	// Deprecated. Instead use `providerServiceId`.
	// To get a list of the provider names, see
	// ListFastConnectProviderServices.
	ProviderName *string `mandatory:"false" json:"providerName,omitempty"`

	// The OCID of the service offered by the provider (if you're connecting
	// via a provider). To get a list of the available service offerings, see
	// ListFastConnectProviderServices.
	ProviderServiceID *string `mandatory:"false" json:"providerServiceId,omitempty"`

	// Deprecated. Instead use `providerServiceId`.
	// To get a list of the provider names, see
	// ListFastConnectProviderServices.
	ProviderServiceName *string `mandatory:"false" json:"providerServiceName,omitempty"`

	// For a public virtual circuit. The public IP prefixes (CIDRs) the customer wants to
	// advertise across the connection.
	PublicPrefixes []CreateVirtualCircuitPublicPrefixDetails `mandatory:"false" json:"publicPrefixes,omitempty"`

	// The Oracle Cloud Infrastructure region where this virtual
	// circuit is located.
	// Example: `phx`
	Region *string `mandatory:"false" json:"region,omitempty"`
}

func (model CreateVirtualCircuitDetails) String() string {
	return common.PointerString(model)
}

// CreateVirtualCircuitDetailsType_Enum Enum with underlying type: string
type CreateVirtualCircuitDetailsType_Enum string

// Set of constants representing the allowable values for CreateVirtualCircuitDetailsType_
const (
	CreateVirtualCircuitDetailsType_Public  CreateVirtualCircuitDetailsType_Enum = "PUBLIC"
	CreateVirtualCircuitDetailsType_Private CreateVirtualCircuitDetailsType_Enum = "PRIVATE"
	CreateVirtualCircuitDetailsType_Unknown CreateVirtualCircuitDetailsType_Enum = "UNKNOWN"
)

var mappingCreateVirtualCircuitDetailsType_ = map[string]CreateVirtualCircuitDetailsType_Enum{
	"PUBLIC":  CreateVirtualCircuitDetailsType_Public,
	"PRIVATE": CreateVirtualCircuitDetailsType_Private,
	"UNKNOWN": CreateVirtualCircuitDetailsType_Unknown,
}

// GetCreateVirtualCircuitDetailsType_EnumValues Enumerates the set of values for CreateVirtualCircuitDetailsType_
func GetCreateVirtualCircuitDetailsType_EnumValues() []CreateVirtualCircuitDetailsType_Enum {
	values := make([]CreateVirtualCircuitDetailsType_Enum, 0)
	for _, v := range mappingCreateVirtualCircuitDetailsType_ {
		if v != CreateVirtualCircuitDetailsType_Unknown {
			values = append(values, v)
		}
	}
	return values
}
