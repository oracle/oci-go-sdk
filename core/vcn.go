// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oci-go-sdk/common"
)

// Vcn. A Virtual Cloud Network (VCN). For more information, see
// [Overview of the Networking Service]({{DOC_SERVER_URL}}/Content/Network/Concepts/overview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type Vcn struct {

	// The CIDR IP address block of the VCN.
	// Example: `172.16.0.0/16`
	CidrBlock *string `mandatory:"true" json:"cidrBlock,omitempty"`

	// The OCID of the compartment containing the VCN.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The VCN's Oracle ID (OCID).
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The VCN's current state.
	LifecycleState VcnLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The OCID for the VCN's default set of DHCP options.
	DefaultDhcpOptionsID *string `mandatory:"false" json:"defaultDhcpOptionsId,omitempty"`

	// The OCID for the VCN's default route table.
	DefaultRouteTableID *string `mandatory:"false" json:"defaultRouteTableId,omitempty"`

	// The OCID for the VCN's default security list.
	DefaultSecurityListID *string `mandatory:"false" json:"defaultSecurityListId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// A DNS label for the VCN, used in conjunction with the VNIC's hostname and
	// subnet's DNS label to form a fully qualified domain name (FQDN) for each VNIC
	// within this subnet (for example, `bminstance-1.subnet123.vcn1.oraclevcn.com`).
	// Must be an alphanumeric string that begins with a letter.
	// The value cannot be changed.
	// The absence of this parameter means the Internet and VCN Resolver will
	// not work for this VCN.
	// For more information, see
	// [DNS in Your Virtual Cloud Network]({{DOC_SERVER_URL}}/Content/Network/Concepts/dns.htm).
	// Example: `vcn1`
	DnsLabel *string `mandatory:"false" json:"dnsLabel,omitempty"`

	// The date and time the VCN was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`

	// The VCN's domain name, which consists of the VCN's DNS label, and the
	// `oraclevcn.com` domain.
	// For more information, see
	// [DNS in Your Virtual Cloud Network]({{DOC_SERVER_URL}}/Content/Network/Concepts/dns.htm).
	// Example: `vcn1.oraclevcn.com`
	VcnDomainName *string `mandatory:"false" json:"vcnDomainName,omitempty"`
}

func (model Vcn) String() string {
	return common.PointerString(model)
}

type VcnLifecycleStateEnum string

const (
	VCN_LIFECYCLE_STATE_PROVISIONING VcnLifecycleStateEnum = "PROVISIONING"
	VCN_LIFECYCLE_STATE_AVAILABLE    VcnLifecycleStateEnum = "AVAILABLE"
	VCN_LIFECYCLE_STATE_TERMINATING  VcnLifecycleStateEnum = "TERMINATING"
	VCN_LIFECYCLE_STATE_TERMINATED   VcnLifecycleStateEnum = "TERMINATED"
	VCN_LIFECYCLE_STATE_UNKNOWN      VcnLifecycleStateEnum = "UNKNOWN"
)

var mapping_vcn_lifecycleState = map[string]VcnLifecycleStateEnum{
	"PROVISIONING": VCN_LIFECYCLE_STATE_PROVISIONING,
	"AVAILABLE":    VCN_LIFECYCLE_STATE_AVAILABLE,
	"TERMINATING":  VCN_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   VCN_LIFECYCLE_STATE_TERMINATED,
	"UNKNOWN":      VCN_LIFECYCLE_STATE_UNKNOWN,
}

func GetVcnLifecycleStateEnumValues() []VcnLifecycleStateEnum {
	values := make([]VcnLifecycleStateEnum, 0)
	for _, v := range mapping_vcn_lifecycleState {
		if v != VCN_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
