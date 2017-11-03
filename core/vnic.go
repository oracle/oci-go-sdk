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

// Vnic. A virtual network interface card. Each VNIC resides in a subnet in a VCN.
// An instance attaches to a VNIC to obtain a network connection into the VCN
// through that subnet. Each instance has a *primary VNIC* that is automatically
// created and attached during launch. You can add *secondary VNICs* to an
// instance after it's launched. For more information, see
// [Virtual Network Interface Cards (VNICs)]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingVNICs.htm).
// Each VNIC has a *primary private IP* that is automatically assigned during launch.
// You can add *secondary private IPs* to a VNIC after it's created. For more
// information, see CreatePrivateIp and
// [IP Addresses]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingIPaddresses.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type Vnic struct {

	// The VNIC's Availability Domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain,omitempty"`

	// The OCID of the compartment containing the VNIC.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The OCID of the VNIC.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The current state of the VNIC.
	LifecycleState VnicLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The private IP address of the primary `privateIp` object on the VNIC.
	// The address is within the CIDR of the VNIC's subnet.
	// Example: `10.0.3.3`
	PrivateIp *string `mandatory:"true" json:"privateIp,omitempty"`

	// The OCID of the subnet the VNIC is in.
	SubnetID *string `mandatory:"true" json:"subnetId,omitempty"`

	// The date and time the VNIC was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// A user-friendly name. Does not have to be unique.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// The hostname for the VNIC's primary private IP. Used for DNS. The value is the hostname
	// portion of the primary private IP's fully qualified domain name (FQDN)
	// (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`).
	// Must be unique across all VNICs in the subnet and comply with
	// [RFC 952](https://tools.ietf.org/html/rfc952) and
	// [RFC 1123](https://tools.ietf.org/html/rfc1123).
	// For more information, see
	// [DNS in Your Virtual Cloud Network]({{DOC_SERVER_URL}}/Content/Network/Concepts/dns.htm).
	// Example: `bminstance-1`
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel,omitempty"`

	// Whether the VNIC is the primary VNIC (the VNIC that is automatically created
	// and attached during instance launch).
	IsPrimary *bool `mandatory:"false" json:"isPrimary,omitempty"`

	// The MAC address of the VNIC.
	// Example: `00:00:17:B6:4D:DD`
	MacAddress *string `mandatory:"false" json:"macAddress,omitempty"`

	// The public IP address of the VNIC, if one is assigned.
	PublicIp *string `mandatory:"false" json:"publicIp,omitempty"`

	// Whether the source/destination check is disabled on the VNIC.
	// Defaults to `false`, which means the check is performed. For information
	// about why you would skip the source/destination check, see
	// [Using a Private IP as a Route Target]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingroutetables.htm#privateip).
	// Example: `true`
	SkipSourceDestCheck *bool `mandatory:"false" json:"skipSourceDestCheck,omitempty"`
}

func (model Vnic) String() string {
	return common.PointerString(model)
}

type VnicLifecycleStateEnum string
type VnicLifecycleState struct{}

const (
	VNIC_LIFECYCLE_STATE_PROVISIONING VnicLifecycleStateEnum = "PROVISIONING"
	VNIC_LIFECYCLE_STATE_AVAILABLE    VnicLifecycleStateEnum = "AVAILABLE"
	VNIC_LIFECYCLE_STATE_TERMINATING  VnicLifecycleStateEnum = "TERMINATING"
	VNIC_LIFECYCLE_STATE_TERMINATED   VnicLifecycleStateEnum = "TERMINATED"
	VNIC_LIFECYCLE_STATE_UNKNOWN      VnicLifecycleStateEnum = "UNKNOWN"
)

var mapping_vnic_lifecycleState = map[string]VnicLifecycleStateEnum{
	"PROVISIONING": VNIC_LIFECYCLE_STATE_PROVISIONING,
	"AVAILABLE":    VNIC_LIFECYCLE_STATE_AVAILABLE,
	"TERMINATING":  VNIC_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   VNIC_LIFECYCLE_STATE_TERMINATED,
	"UNKNOWN":      VNIC_LIFECYCLE_STATE_UNKNOWN,
}

func (receiver VnicLifecycleState) Values() []VnicLifecycleStateEnum {
	values := make([]VnicLifecycleStateEnum, 0)
	for _, v := range mapping_vnic_lifecycleState {
		if v != VNIC_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver VnicLifecycleState) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if VnicLifecycleStateEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver VnicLifecycleState) From(toBeConverted string) VnicLifecycleStateEnum {
	if val, ok := mapping_vnic_lifecycleState[toBeConverted]; ok {
		return val
	}
	return VNIC_LIFECYCLE_STATE_UNKNOWN
}
