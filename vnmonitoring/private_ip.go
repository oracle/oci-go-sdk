// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivateIp A *private IP* is a conceptual term that refers to an IPv4 private IP address and related properties.
// The `privateIp` object is the API representation of a private IP.
// **Note:** For information about IPv6 addresses, see Ipv6.
// Each instance has a *primary private IP* that is automatically created and
// assigned to the primary VNIC during instance launch. If you add a secondary
// VNIC to the instance, it also automatically gets a primary private IP. You
// can't remove a primary private IP from its VNIC. The primary private IP is
// automatically deleted when the VNIC is terminated.
// You can add *secondary private IPs* to a VNIC after it's created. For more
// information, see the `privateIp` operations and also
// IP Addresses (https://docs.oracle.com/iaas/Content/Network/Tasks/managingIPaddresses.htm).
// **Note:** Only
// ListPrivateIps and
// GetPrivateIp work with
// *primary* private IPs. To create and update primary private IPs, you instead
// work with instance and VNIC operations. For example, a primary private IP's
// properties come from the values you specify in
// CreateVnicDetails when calling either
// LaunchInstance or
// AttachVnic. To update the hostname
// for a primary private IP, you use UpdateVnic.
// `PrivateIp` objects that are created for use with the Oracle Cloud VMware Solution are
// assigned to a VLAN and not a VNIC in a subnet. See the
// descriptions of the relevant attributes in the `PrivateIp` object. Also see
// Vlan.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
type PrivateIp struct {

	// The private IP's availability domain. This attribute will be null if this is a *secondary*
	// private IP assigned to a VNIC that is in a *regional* subnet.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the private IP.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The hostname for the private IP. Used for DNS. The value is the hostname
	// portion of the private IP's fully qualified domain name (FQDN)
	// (for example, `bminstance1` in FQDN `bminstance1.subnet123.vcn1.oraclevcn.com`).
	// Must be unique across all VNICs in the subnet and comply with
	// RFC 952 (https://tools.ietf.org/html/rfc952) and
	// RFC 1123 (https://tools.ietf.org/html/rfc1123).
	// For more information, see
	// DNS in Your Virtual Cloud Network (https://docs.oracle.com/iaas/Content/Network/Concepts/dns.htm).
	// Example: `bminstance1`
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// The private IP's Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
	Id *string `mandatory:"false" json:"id"`

	// The private IP address of the `privateIp` object. The address is within the CIDR
	// of the VNIC's subnet.
	// However, if the `PrivateIp` object is being used with a VLAN as part of
	// the Oracle Cloud VMware Solution, the address is from the range specified by the
	// `cidrBlock` attribute for the VLAN. See Vlan.
	// Example: `10.0.3.3`
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// Whether this private IP is the primary one on the VNIC. Primary private IPs
	// are unassigned and deleted automatically when the VNIC is terminated.
	// Example: `true`
	IsPrimary *bool `mandatory:"false" json:"isPrimary"`

	// Applicable only if the `PrivateIp` object is being used with a VLAN as part of
	// the Oracle Cloud VMware Solution. The `vlanId` is the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN. See
	// Vlan.
	VlanId *string `mandatory:"false" json:"vlanId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the VNIC is in.
	// However, if the `PrivateIp` object is being used with a VLAN as part of
	// the Oracle Cloud VMware Solution, the `subnetId` is null.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The date and time the private IP was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC the private IP is assigned to. The VNIC and private IP
	// must be in the same subnet.
	// However, if the `PrivateIp` object is being used with a VLAN as part of
	// the Oracle Cloud VMware Solution, the `vnicId` is null.
	VnicId *string `mandatory:"false" json:"vnicId"`
}

func (m PrivateIp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateIp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
